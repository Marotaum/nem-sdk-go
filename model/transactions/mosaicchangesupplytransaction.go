package transactions

import (
	"github.com/Marotaum/nem-sdk-go/base"
	"github.com/Marotaum/nem-sdk-go/model"
	"github.com/Marotaum/nem-sdk-go/utils"
	"github.com/pkg/errors"
)

type MosaicSupplyChange struct {
	NamespaceID     string
	MosaicName      string
	SupplyType      int
	Delta           float64
	IsMultisig      bool
	MultisigAccount string
	Due             int64
	Network         int
	SenderPublicKey string
}

func (r MosaicSupplyChange) Prepare(common Common, network int) *base.MosaicSupplyChangeTransaction {
	if !utils.IsPrivateKeyValid(common.PrivateKey) {
		panic(nil)
	}
	kp := model.KeyPairCreate(common.PrivateKey)
	if r.IsMultisig {
		if r.MultisigAccount != "" {
			if !utils.IsPublicKeyValid(r.MultisigAccount) {
				panic(nil)
			}
			r.SenderPublicKey = r.MultisigAccount
		} else {
			err := errors.New("must place a publickey of the multifirm account")
			panic(err)
		}
	} else {
		r.SenderPublicKey = kp.PublicString()
	}

	if network == model.Data.Testnet.ID {
		r.Due = 60
	} else {
		r.Due = 24 * 60
	}
	r.Network = network
	return constructMSC(r)
}

// Change a mosaic supply transaction struct
// param msc - A MosaicSupplyChange struct
// return - A [MosaicSupplyChangeTransaction] struct
// link https://nemproject.github.io/#mosaicSupplyChangeTransaction
func constructMSC(msc MosaicSupplyChange) *base.MosaicSupplyChangeTransaction {
	timeStamp := utils.CreateNEMTimeStamp()
	version := model.GetVersion(1, msc.Network)
	data := CommonPart(model.MosaicSupply, version, timeStamp, msc.Due, msc.SenderPublicKey)
	fee := model.NamespaceAndMosaicCommon

	custom := base.MosaicSupplyChangeTransaction{
		TimeStamp:       data.TimeStamp,
		Type:            data.Type,
		Deadline:        data.Deadline,
		Version:         data.Version,
		Signer:          data.Signer,
		Fee:             fee,

		SupplyType:      msc.SupplyType,
		Delta:           msc.Delta,
		MosaicID:        base.MosaicID {
			NamespaceID: msc.NamespaceID,
			Name:        msc.MosaicName,
		},
	}
	return &custom
}
