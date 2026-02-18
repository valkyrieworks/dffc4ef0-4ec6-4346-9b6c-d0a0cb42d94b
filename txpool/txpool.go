package txpool

import (
	"crypto/sha256"
	"fmt"
	"math"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/kinds"
)

const (
	TxpoolConduit = byte(0x30)

	//
	NodeOvertakePauseCadenceMillis = 100

	//
	//
	UnclearNodeUID uint16 = 0

	MaximumEnabledIDXDatastore = math.MaxUint16
)

//

//
//
//
//
type Txpool interface {
	//
	//
	InspectTransfer(tx kinds.Tx, callback func(*iface.ReplyInspectTransfer), transferDetails TransferDetails) error

	//
	//
	DeleteTransferByKey(transferKey kinds.TransferKey) error

	//
	//
	//
	//
	//
	//
	HarvestMaximumOctetsMaximumFuel(maximumOctets, maximumFuel int64) kinds.Txs

	//
	//
	//
	HarvestMaximumTrans(max int) kinds.Txs

	//
	//
	//
	//
	Secure()

	//
	Release()

	//
	//
	//
	//
	//
	//
	Modify(
		ledgerLevel int64,
		ledgerTrans kinds.Txs,
		dispatchTransferReplies []*iface.InvokeTransferOutcome,
		newPreFn PreInspectFunction,
		newSubmitFn SubmitInspectFunction,
	) error

	//
	//
	//
	//
	//
	PurgeApplicationLink() error

	//
	Purge()

	//
	//
	//
	//
	//
	TransAccessible() <-chan struct{}

	//
	//
	ActivateTransAccessible()

	//
	Volume() int

	//
	VolumeOctets() int64
}

//
//
//
type PreInspectFunction func(kinds.Tx) error

//
//
//
type SubmitInspectFunction func(kinds.Tx, *iface.ReplyInspectTransfer) error

//
//
func PreInspectMaximumOctets(maximumOctets int64) PreInspectFunction {
	return func(tx kinds.Tx) error {
		transferVolume := kinds.CalculateSchemaVolumeForTrans([]kinds.Tx{tx})

		if transferVolume > maximumOctets {
			return fmt.Errorf("REDACTED", transferVolume, maximumOctets)
		}

		return nil
	}
}

//
//
func SubmitInspectMaximumFuel(maximumFuel int64) SubmitInspectFunction {
	return func(tx kinds.Tx, res *iface.ReplyInspectTransfer) error {
		if maximumFuel == -1 {
			return nil
		}
		if res.FuelDesired < 0 {
			return fmt.Errorf("REDACTED",
				res.FuelDesired)
		}
		if res.FuelDesired > maximumFuel {
			return fmt.Errorf("REDACTED",
				res.FuelDesired, maximumFuel)
		}

		return nil
	}
}

//
type TransferKey [sha256.Size]byte
