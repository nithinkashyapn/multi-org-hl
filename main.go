package main

import (
	"fmt"
	"os"

	"github.com/cubereum/mew/blockchain"
	"github.com/cubereum/mew/web"
	"github.com/cubereum/mew/web/controllers"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.example.com",

		// Channel parameters
		ChannelID:     "mychannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/cubereum/mew/network/artifacts/channel.tx",

		// Chaincode parameters
		ChainCodeID:     "service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/cubereum/mew/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}
