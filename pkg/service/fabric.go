package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

/*
	FabricContext
*/
type FabricContext struct {
	Wallet       *gateway.Wallet
	Gateway      *gateway.Gateway
	Network      *gateway.Network
	Contract     *gateway.Contract
	ContractName string
	FabricDetail
}

type FabricDetail struct {
	Sdk     *fabsdk.FabricSDK
	Channel *channel.Client
}

func (f *FabricContext) BuildDetail(channelID string, userName string, orgName string) {
	log.Println("============ BuildDetail starts ============")
	sdk, err := fabsdk.New(config.FromFile(getCcpPath()))
	if err != nil {
		//log.Println(err.Error())
		log.Printf("Failed to create new channel client: %s \n ", err)
	}
	log.Printf("sdk : %+v \n ", sdk)
	f.Sdk = sdk
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	ctx, err := channel.New(clientChannelContext)
	if err != nil {
		log.Printf("Failed to create new channel client: %s \n", err)
	}
	log.Printf("ctx : %+v \n ", ctx)
	f.Channel = ctx
}

/*
Build FabricContext
*/
func (f *FabricContext) Build(
	userName string,
	walletName string,
	channelName string,
	contractName string,
	mspName string,
) {
	f.ContractName = contractName
	log.Println("============ application-golang starts ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet(walletName)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists(userName) {
		err = populateWallet(wallet, userName, mspName)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := getCcpPath()

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, userName),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channelName)
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract(contractName)
	f.Wallet = wallet
	f.Gateway = gw
	f.Network = network
	f.Contract = contract

}

func getCcpPath() string {
	return filepath.Join(
		//"..",
		//"..",
		// os.Getenv("Network"),           //"test-network",
		// os.Getenv("Organizations"),     //"organizations",
		// os.Getenv("PeerOrganizations"), //"peerOrganizations",
		// os.Getenv("ORG"),               //"org1.example.com",
		// os.Getenv("CONN-YAML"),         //"connection-org1.yaml",
		os.Getenv("CCP_PATH"),
		"ccp.yaml",
	)
}

func getCredPath() (string, string) {

	credPath := filepath.Join(
		os.Getenv("CRED_PATH"),
		//"..",
		// "..",
		// os.Getenv("Network"),           //"test-network"
		// os.Getenv("Organizations"),     //"organizations"
		// os.Getenv("PeerOrganizations"), //"peerOrganizations",
		// os.Getenv("ORG"),               //"org1.example.com",
		// os.Getenv("Users"),             //"users",
		// os.Getenv("FULL_ACCOUNT"),       //"User1@org1.example.com",
		// os.Getenv("MSP"),               //"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", os.Getenv("FULL_ACCOUNT")+"-cert.pem")
	return credPath, certPath
}

func populateWallet(wallet *gateway.Wallet, userName string, orgName string) error {
	log.Println("============ Populating wallet ============")

	credPath, certPath := getCredPath()
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(orgName, string(cert), string(key))

	return wallet.Put(userName, identity)
}
