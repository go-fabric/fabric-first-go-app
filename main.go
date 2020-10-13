package main

import (
	"fmt"
	"github.com/shuizhongmose/go-fabric/fabric-first-go-app/sdkenv"
	"os"
	"github.com/shuizhongmose/go-fabric/fabric-first-go-app/service"
	"github.com/shuizhongmose/go-fabric/fabric-first-go-app/web/controllers"
	"github.com/shuizhongmose/go-fabric/fabric-first-go-app/web"
)

const (
	cc_name = "simplecc"
	cc_version = "1.0.0"
)
func main() {
	// init orgs information
	orgs := []*sdkenv.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-samples/test-network/channel-artifacts/Org1MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org2",
			OrgMspId:      "Org2MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-samples/test-network/channel-artifacts/Org2MSPanchors.tx",
		},
	}

	// init sdk env info
	info := sdkenv.SdkEnvInfo{
		ChannelID:        "testchannel",
		ChannelConfig:    os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-samples/test-network/channel-artifacts/testchannel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    os.Getenv("GOPATH")+"/src/github.com/shuizhongmose/go-fabric/fabric-first-go-app/chaincode/",
		ChaincodeVersion: cc_version,
	}

	// sdk setup
	sdk, err := sdkenv.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and jion
	if err := sdkenv.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkenv.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")
	serviceHandler, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk)
	if err!=nil{
		fmt.Println()
		os.Exit(-1)
	}

	msg, err := serviceHandler.SetInfo("name", "verayy")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	fmt.Println(">> 设置链码状态完成")

	// start web service
	fmt.Println(">> 启动web服务......")
	app := controllers.Application{
		Fabric: serviceHandler,
	}
	web.WebStart(&app)
	fmt.Println(">> 启动web服务......")
}
