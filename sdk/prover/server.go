package prover

import (
	"bytes"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/sdk/proto/sdkproto"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/constraint"
	"google.golang.org/grpc"
)

type Service struct {
	svr *server
}

// NewService creates a new prover server instance that automatically manages
// compilation & setup, and serves as a GRPC server that interoperates with
// brevis sdk in other languages.
func NewService(app sdk.AppCircuit, config ServiceConfig) (*Service, error) {
	pk, vk, ccs, err := readOrSetup(app, config.SetupDir, config.GetSrsDir())
	if err != nil {
		return nil, err
	}
	return &Service{
		svr: newServer(app, pk, vk, ccs),
	}, nil
}

func (s *Service) Serve(port uint) {
	go s.serveGrpc(port)
	s.serveGrpcGateway(port, port+10)
}

func (s *Service) serveGrpc(port uint) {
	grpcServer := grpc.NewServer()
	sdkproto.RegisterProverServer(grpcServer, s.svr)
	address := fmt.Sprintf("localhost:%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("failed to start prover server:", err)
		os.Exit(1)
	}
	fmt.Println(">> serving prover GRPC at port", port)
	if err = grpcServer.Serve(lis); err != nil {
		fmt.Println("grpc server crashed", err)
		os.Exit(1)
	}
}

func (s *Service) serveGrpcGateway(grpcPort, restPort uint) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := fmt.Sprintf("localhost:%d", grpcPort)

	err := sdkproto.RegisterProverHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		fmt.Println("failed to start prover server:", err)
		os.Exit(1)
	}

	handler := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(mux)

	fmt.Println(">> serving prover REST API at port", restPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", restPort), handler); err != nil {
		fmt.Println("REST server crashed", err)
		os.Exit(1)
	}
}

type server struct {
	sdkproto.UnimplementedProverServer

	app sdk.AppCircuit

	pk  plonk.ProvingKey
	vk  plonk.VerifyingKey
	ccs constraint.ConstraintSystem

	vkBytes string
}

func newServer(
	app sdk.AppCircuit,
	pk plonk.ProvingKey,
	vk plonk.VerifyingKey,
	ccs constraint.ConstraintSystem,
) *server {
	var buf bytes.Buffer
	_, err := vk.WriteRawTo(&buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &server{
		app:     app,
		pk:      pk,
		vk:      vk,
		ccs:     ccs,
		vkBytes: hexutil.Encode(buf.Bytes()),
	}
}

func (s *server) Prove(ctx context.Context, req *sdkproto.ProveRequest) (*sdkproto.ProveResponse, error) {
	fmt.Println(req.String())
	brevisApp, err := sdk.NewBrevisApp()

	if err != nil {
		msg := "failed to new brevis app: " + err.Error()
		fmt.Println(msg)
		return prepareErrorResponse(sdkproto.ErrCode_ERROR_DEFAULT, msg), nil
	}

	for _, receipt := range req.Receipts {
		sdkReceipt, err := convertProtoReceiptToSdkReceipt(receipt.Data)
		if err != nil {
			msg := fmt.Sprintf("invalid sdk receipt: %+v, %s", receipt.Data, err.Error())
			fmt.Println(msg)
			return prepareErrorResponse(sdkproto.ErrCode_ERROR_INVALID_INPUT, msg), nil
		}
		brevisApp.AddReceipt(sdkReceipt, int(receipt.Index))
	}

	for _, storage := range req.Storages {
		sdkStorage, err := convertProtoStorageToSdkStorage(storage.Data)
		if err != nil {
			msg := fmt.Sprintf("invalid sdk storage: %+v, %s", storage.Data, err.Error())
			fmt.Println(msg)
			return prepareErrorResponse(sdkproto.ErrCode_ERROR_INVALID_INPUT, msg), nil
		}

		brevisApp.AddStorage(sdkStorage, int(storage.Index))
	}

	for _, transaction := range req.Transactions {
		sdkTx, err := convertProtoTxToSdkTx(transaction.Data)
		if err != nil {
			msg := fmt.Sprintf("invalid sdk transaction: %+v, %s", transaction.Data, err.Error())
			fmt.Println(msg)
			return prepareErrorResponse(sdkproto.ErrCode_ERROR_INVALID_INPUT, msg), nil
		}

		brevisApp.AddTransaction(sdkTx, int(transaction.Index))
	}

	guest, err := assignCustomInput(s.app, req.CustomInput)
	if err != nil {
		fmt.Printf("invalid custom input %s\n", err.Error())
		return prepareErrorResponse(sdkproto.ErrCode_ERROR_INVALID_CUSTOM_INPUT, err.Error()), nil
	}

	input, err := brevisApp.BuildCircuitInput(guest)
	if err != nil {
		msg := fmt.Sprintf("failed to build circuit input: %+v, %s", req, err.Error())
		fmt.Println(msg)
		return prepareErrorResponse(sdkproto.ErrCode_ERROR_FAILED_TO_PROVE, msg), nil
	}

	witness, publicWitness, err := sdk.NewFullWitness(guest, input)
	if err != nil {
		msg := fmt.Sprintf("failed to get full witness: %+v, %s", req, err.Error())
		fmt.Println(msg)
		return prepareErrorResponse(sdkproto.ErrCode_ERROR_FAILED_TO_PROVE, msg), nil
	}

	proof, err := sdk.Prove(s.ccs, s.pk, witness)
	if err != nil {
		msg := fmt.Sprintf("failed to prove: %+v, %s", req, err.Error())
		fmt.Println(msg)
		return prepareErrorResponse(sdkproto.ErrCode_ERROR_FAILED_TO_PROVE, msg), nil
	}

	err = sdk.Verify(s.vk, publicWitness, proof)
	if err != nil {
		msg := fmt.Sprintf("failed to test verifying after proving: %+v, %s", req, err.Error())
		fmt.Println(msg)
		return prepareErrorResponse(sdkproto.ErrCode_ERROR_FAILED_TO_PROVE, msg), nil
	}

	var buf bytes.Buffer
	_, err = proof.WriteRawTo(&buf)
	if err != nil {
		msg := fmt.Sprintf("failed to write proof bytes: %+v, %s", req, err.Error())
		fmt.Println(msg)
		return prepareErrorResponse(sdkproto.ErrCode_ERROR_DEFAULT, msg), nil
	}

	return &sdkproto.ProveResponse{
		Proof:       hexutil.Encode(buf.Bytes()),
		CircuitInfo: buildAppCircuitInfo(input, s.vkBytes),
	}, nil
}

func prepareErrorResponse(code sdkproto.ErrCode, errorMessage string) *sdkproto.ProveResponse {
	msg := ""
	switch code {
	case sdkproto.ErrCode_ERROR_UNDEFINED:
		msg = "unknown error"
	case sdkproto.ErrCode_ERROR_DEFAULT:
		msg = "internal server error"
	case sdkproto.ErrCode_ERROR_INVALID_INPUT:
		msg = "invalid input"
	case sdkproto.ErrCode_ERROR_INVALID_CUSTOM_INPUT:
		msg = "invalid custom input"
	case sdkproto.ErrCode_ERROR_FAILED_TO_PROVE:
		msg = "failed to prove"
	default:
		fmt.Sprintln("found unknown code usage", code)
		msg = "unknown error"
	}

	return &sdkproto.ProveResponse{
		Err: &sdkproto.Err{
			Code: code,
			Msg:  msg + ": " + errorMessage,
		},
	}
}
