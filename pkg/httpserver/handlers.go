package httpserver

import (
	"encoding/json"
	"fabric-asset-transfer-basic-server/pkg/context"
	"fabric-asset-transfer-basic-server/pkg/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
)

// InitLedger godoc
// GET
func InitLedger(w http.ResponseWriter, r *http.Request) {
	contract := context.SharedDataContext.MyFabric.Contract
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Printf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))

}

// UpdateData godoc
// POST
// @Summary https://tender-chain.fishlab.com.tw/UpdateData
// @Description <br> Id: 透過DB查出 <br> TenderID: 標案號 <br> Accountcode: 銀行代碼 <br> Account: 銀行代碼 <br> Name: 戶名 <br> Currency: 幣別 <br> Branch: 分行 <br> Amount: 金額 <br> Status: 憑證狀態
// @Tags DApp
// @Accept  json
// @Produce  json
// @Param dapp body model.Data true "UpdateData"
// @Success 200 "OK"
// @Router /UpdateData [post]
func UpdateData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===[ start Update Data]===")
	cctx := context.SharedDataContext.MyFabric.Channel
	ccID := context.SharedDataContext.MyFabric.ContractName
	var tender model.Data
	err := json.NewDecoder(r.Body).Decode(&tender)
	if err != nil {
		log.Println("Failed to unMarshal")
		return
	}
	var txArgs = [][]byte{
		[]byte(tender.Id),
		[]byte(tender.TenderID),
		[]byte(tender.Accountcode),
		[]byte(tender.Account),
		[]byte(tender.Name),
		[]byte(tender.Currency),
		[]byte(tender.Branch),
		[]byte(tender.Amount),
		[]byte(tender.Status)}

	result, err := cctx.Execute(
		channel.Request{ChaincodeID: ccID, Fcn: "UpdateData", Args: txArgs},
		channel.WithRetry(retry.DefaultChannelOpts))
	log.Printf("result: %+v \n", result)
	if err != nil {
		fmt.Println("Failed to Update hash: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Payload))

}

// CreateData godoc
// POST

// CreateData godoc
// @Summary https://tender-chain.fishlab.com.tw/CreateData
// @Description 傳入以下參數 <br> Id: 透過DB查出 <br> TenderID: 標案號 <br> Accountcode: 銀行代碼 <br> Account: 銀行代碼 <br> Name: 戶名 <br> Currency: 幣別 <br> Branch: 分行 <br> Amount: 金額 <br> Status: 憑證狀態
// @Tags DApp
// @Accept  json
// @Produce  json
// @Param dapp body model.Data true "CreateData"
// @Success 200 "OK"
// @Failure 500 "Data already exists"
// @Router /CreateData [post]
func CreateData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===[ start Create  Data]===")
	cctx := context.SharedDataContext.MyFabric.Channel
	ccID := context.SharedDataContext.MyFabric.ContractName
	var tender model.Data
	err := json.NewDecoder(r.Body).Decode(&tender)
	if err != nil {
		log.Printf("Failed to check TransferData  car: %s", err)
		log.Println("--------")
		s, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "%s", s)
		//log.Printf("Failed to unMarshal when create => r.Body:  %+v \n", r.Body)
		return
	}
	var txArgs = [][]byte{
		[]byte(tender.Id),
		[]byte(tender.TenderID),
		[]byte(tender.Accountcode),
		[]byte(tender.Account),
		[]byte(tender.Name),
		[]byte(tender.Currency),
		[]byte(tender.Branch),
		[]byte(tender.Amount),
		[]byte(tender.Status)}

	result, err := cctx.Execute(
		channel.Request{ChaincodeID: ccID, Fcn: "CreateData", Args: txArgs},
		channel.WithRetry(retry.DefaultChannelOpts))

	log.Printf("result: %+v \n", result)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Printf("Failed to create Data: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result.Payload))
	}

	w.Write([]byte(result.Payload))

}

// ReadData godoc
// GET

// ReadData godoc
// @Summary https://tender-chain.fishlab.com.tw/ReadData
// @Description 傳入以下參數 <br> Id: 透過DB查出
// @Tags DApp
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Data
// @Param Id query string true "Id"
// @Router /ReadData [get]
func ReadData(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := vars["Id"]
	if !ok {
		fmt.Printf("param id does not exist\n")
	}
	log.Println("id[0]=" + id[0])
	fmt.Println("===[ start Read  Data]===")
	cctx := context.SharedDataContext.MyFabric.Channel
	ccID := context.SharedDataContext.MyFabric.ContractName

	var txArgs = [][]byte{
		[]byte(id[0]),
	}

	result, err := cctx.Query(
		channel.Request{ChaincodeID: ccID, Fcn: "ReadData", Args: txArgs},
		channel.WithRetry(retry.DefaultChannelOpts))

	log.Printf("result: %+v \n", result)
	if err != nil {
		log.Printf("Failed to read car: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Payload))

}

// GET

// GetHistory godoc
// @Summary https://tender-chain.fishlab.com.tw/GetHistory
// @Description 傳入以下參數 <br> Id: 透過DB查出
// @Tags DApp
// @Accept  json
// @Produce  json
// @Success 200 "OK"
// @Param Id query string true "Id"
// @Router /GetHistory [get]
func GetHistory(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := vars["Id"]
	if !ok {
		fmt.Printf("param id does not exist\n")
	}
	log.Println("id[0]=" + id[0])
	fmt.Println("===[ start Read id]===")
	cctx := context.SharedDataContext.MyFabric.Channel
	ccID := context.SharedDataContext.MyFabric.ContractName

	var txArgs = [][]byte{
		[]byte(id[0]),
	}

	result, err := cctx.Query(
		channel.Request{ChaincodeID: ccID, Fcn: "GetHistory", Args: txArgs},
		channel.WithRetry(retry.DefaultChannelOpts))

	log.Printf("result: %+v \n", result)
	if err != nil {
		log.Printf("Failed to read data: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Payload))

}

// DeleteData godoc
// GET

func DeleteData(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	hash, ok := vars["Hash"]
	if !ok {
		fmt.Printf("param hash does not exist\n")
	}
	log.Println("hash[0]=" + hash[0])
	fmt.Println("===[ start Delete Data]===")
	cctx := context.SharedDataContext.MyFabric.Channel
	ccID := context.SharedDataContext.MyFabric.ContractName

	var txArgs = [][]byte{
		[]byte(hash[0]),
	}

	result, err := cctx.Execute(
		channel.Request{ChaincodeID: ccID, Fcn: "DeleteData", Args: txArgs},
		channel.WithRetry(retry.DefaultChannelOpts))

	log.Printf("result: %+v \n", result)
	if err != nil {
		log.Printf("Failed to delete data: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Payload))

}
