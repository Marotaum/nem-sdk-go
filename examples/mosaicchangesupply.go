package main

import (
	"github.com/Marotaum/nem-sdk-go/com/requests"
	"github.com/Marotaum/nem-sdk-go/model"
	"github.com/Marotaum/nem-sdk-go/model/objects"
	"github.com/Marotaum/nem-sdk-go/utils"

	"fmt"
	"github.com/Marotaum/nem-sdk-go/model/transactions"
)

func main() {
	// Create an NIS endpoint
	endpoint := objects.Endpoint(model.DefaultTestnet, model.DefaultPort)
	client := requests.NewClient(endpoint)

	// Create a common object holding key
	common := objects.GetCommon("", "265087519502bd6f6c93f74b189ecdea18da9f58ba9d83a425821e714ea2aeea", false)

	msc := objects.MosaicSupplyChange()
	msc.Delta = 1
	msc.IsMultisig = false
	msc.NamespaceID = "xem"
	msc.MosaicName = "token"
	msc.MultisigAccount = ""
	msc.SupplyType = 2
	msc.Network = model.Data.Testnet.ID

	// Prepare the change transaction
	transactionEntity := msc.Prepare(common, model.Data.Testnet.ID)

	res, err := transactions.Send(common, transactionEntity, *client)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Transfer\n%s", utils.Struc2Json(res))
}
