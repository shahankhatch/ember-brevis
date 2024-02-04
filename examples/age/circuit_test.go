package age

import (
	"context"
	"fmt"
	"github.com/celer-network/brevis-sdk/sdk"
	"github.com/celer-network/brevis-sdk/test"
	"github.com/ethereum/go-ethereum/common"
	"path/filepath"
	"testing"
)

func TestCircuit(t *testing.T) {
	app, err := sdk.NewBrevisApp("https://eth-mainnet.nodereal.io/v1/0af795b55d124a61b86836461ece1dee") // TODO use your eth rpc
	check(err)

	//txHash := common.HexToHash(
	//	"954b01a12e0846eca75751b248796597b6b1715f5a23ada7f2009a8930ce10ad")
	txHash := common.HexToHash(
		"0x6dc75e61220cc775aafa17796c20e49ac08030020fce710e3e546aa4e003454c")

	app.AddTransaction(sdk.TransactionQuery{TxHash: txHash})

	appCircuit := &AppCircuit{}
	appCircuitAssignment := &AppCircuit{}

	circuitInput, err := app.BuildCircuitInput(context.Background(), appCircuit)
	check(err)

	///////////////////////////////////////////////////////////////////////////////
	// Testing
	///////////////////////////////////////////////////////////////////////////////

	test.IsSolved(t, appCircuit, appCircuitAssignment, circuitInput)

	///////////////////////////////////////////////////////////////////////////////
	// Compiling and Setup
	///////////////////////////////////////////////////////////////////////////////

	outDir := "$HOME/circuitOut/age"
	//srsDir := "$HOME/kzgsrs"

	// The compilation output is the description of the circuit's constraint system.
	// You should use sdk.WriteTo to serialize and save your circuit so that it can
	// be used in the proving step later.
	//ccs, err := sdk.Compile(appCircuit, circuitInput)
	//check(err)
	//err = sdk.WriteTo(ccs, filepath.Join(outDir, "ccs"))
	//check(err)

	// Setup is a one-time effort per circuit. A cache dir can be provided to output
	// external dependencies. Once you have the verifying key you should also save
	// its hash in your contract so that when a proof via Brevis is submitted
	// on-chain you can verify that Brevis indeed used your verifying key to verify
	// your circuit computations
	//pk, vk, err := sdk.Setup(ccs, srsDir)
	//check(err)
	//err = sdk.WriteTo(pk, filepath.Join(outDir, "pk"))
	//check(err)
	//err = sdk.WriteTo(vk, filepath.Join(outDir, "vk"))
	//check(err)

	fmt.Println("compilation/setup complete")

	///////////////////////////////////////////////////////////////////////////////
	// Proving
	///////////////////////////////////////////////////////////////////////////////

	ccs, err := sdk.ReadCircuitFrom(filepath.Join(outDir, "ccs"))
	check(err)
	pk, err := sdk.ReadPkFrom(filepath.Join(outDir, "pk"))
	check(err)
	vk, err := sdk.ReadVkFrom(filepath.Join(outDir, "vk"))
	check(err)

	witness, _, err := sdk.NewFullWitness(appCircuitAssignment, circuitInput)
	check(err)
	proof, err := sdk.Prove(ccs, pk, witness)
	check(err)

	///////////////////////////////////////////////////////////////////////////////
	// Initiating Brevis Request
	///////////////////////////////////////////////////////////////////////////////

	fmt.Println(">> Initiate Brevis Request")
	appContract := common.HexToAddress("0x73090023b8D731c4e87B3Ce9Ac4A9F4837b4C1bd")
	refundee := common.HexToAddress("0x164Ef8f77e1C88Fb2C724D3755488bE4a3ba4342")

	calldata, _, feeValue, err := app.PrepareRequest(vk, 1, 11155111, refundee, appContract)
	check(err)
	fmt.Printf("calldata 0x%x\nfeeValue %d Wei\n", calldata, feeValue)

	///////////////////////////////////////////////////////////////////////////////
	// Submit Proof to Brevis
	///////////////////////////////////////////////////////////////////////////////

	fmt.Println(">> Submit Proof to Brevis")
	err = app.SubmitProof(proof)
	check(err)

	// [Call BrevisProof.sendRequest() with the above calldata]

	// Poll Brevis gateway for query status till the final proof is submitted
	// on-chain by Brevis and your contract is called
	tx, err := app.WaitFinalProofSubmitted(context.Background())
	check(err)
	fmt.Printf("tx hash %s\n", tx)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
