package context

import (
	"fabric-asset-transfer-basic-server/pkg/service"
	"log"
)

/*
	Context Type
*/
type GlobalContext struct {
	MyFabric *service.FabricContext
}

/*
New Context
*/
func New(
	userName string,
	walletName string,
	channelName string,
	contractName string,
	mspName string,
	orgName string,
) *GlobalContext {
	log.Println("========[ new context ]==========")
	log.Println("userName: " + userName)
	log.Println("walletName: " + walletName)
	log.Println("channelName: " + channelName)
	log.Println("contractName: " + contractName)
	log.Println("orgName: " + mspName)

	ctx := &service.FabricContext{}
	ctx.Build(userName, walletName, channelName, contractName, mspName)
	ctx.BuildDetail(channelName, userName, orgName)
	return &GlobalContext{
		MyFabric: ctx,
	}
}

var SharedDataContext *GlobalContext
