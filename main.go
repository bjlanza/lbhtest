package main

import (
	"fmt"
	"os"

	"github.com/bjlanza/lbhtest/blockchain"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.example.com",

		// Channel parameters
		ChannelID:     "filandon",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/bjlanza/lbhtest/network/channel-artifacts/channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:  "github.com/bjlanza/lbhtest/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1"
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()
}
