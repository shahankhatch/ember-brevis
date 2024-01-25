package sdk

import (
	"context"
	"fmt"
	"github.com/celer-network/brevis-sdk/sdk/proto"
	"google.golang.org/grpc"
)

type GatewayClient struct {
	c proto.WebClient
}

func NewGatewayClient(url ...string) (*GatewayClient, error) {
	if len(url) > 1 {
		panic("must supply at most one url")
	}
	gatewayUrl := "appsdk.brevis.network:11082"
	if len(url) > 0 {
		gatewayUrl = url[0]
	}
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(gatewayUrl, opts...)
	if err != nil {
		return nil, err
	}
	return &GatewayClient{
		c: proto.NewWebClient(conn),
	}, nil
}

func (c *GatewayClient) PrepareQuery(req *proto.PrepareQueryRequest) (resp *proto.PrepareQueryResponse, err error) {
	resp, err = c.c.PrepareQuery(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("invalid resp, err: %v", resp.Err)
	}
	return
}

func (c *GatewayClient) GetQueryStatus(req *proto.GetQueryStatusRequest) (resp *proto.GetQueryStatusResponse, err error) {
	resp, err = c.c.GetQueryStatus(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("invalid resp, err: %v", resp.Err)
	}
	return
}

func (c *GatewayClient) SubmitProof(req *proto.SubmitAppCircuitProofRequest) (resp *proto.SubmitAppCircuitProofResponse, err error) {
	resp, err = c.c.SubmitAppCircuitProof(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("invalid resp, err: %v", resp.Err)
	}
	return
}
