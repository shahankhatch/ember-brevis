package main

import (
	"fmt"
	"os"

	"github.com/brevis-network/brevis-sdk/ember"
	"github.com/brevis-network/brevis-sdk/sdk/prover"
)

func main() {
	proverService, err := prover.NewService(
		&ember.AppCircuit{},
		prover.ServiceConfig{
			SetupDir: "$HOME/circuitOut",
			SrsDir:   "$HOME/kzgsrs",
			RpcURL:   "http://localhost:8547",
			ChainId:  412346,
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	proverService.Serve("localhost", 33247)
}
