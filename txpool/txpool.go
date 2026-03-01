package txpool

import (
	"crypto/sha256"
	"fmt"
	"math"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	TxpoolConduit = byte(0x30)

	//
	NodeOvertakeSnoozeDurationMSEC = 100

	//
	//
	UnfamiliarNodeUUID uint16 = 0

	MaximumDynamicIDXDstore = math.MaxUint16
)

//

//
//
//
//
type Txpool interface {
	//
	//
	InspectTransfer(tx kinds.Tx, clbk func(*iface.ReplyInspectTransfer), transferDetails TransferDetails) error

	//
	//
	DiscardTransferViaToken(transferToken kinds.TransferToken) error

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
	Revise(
		ledgerAltitude int64,
		ledgerTrans kinds.Txs,
		dispatchTransferReplies []*iface.InvokeTransferOutcome,
		freshAnteProc PriorInspectMethod,
		freshSubmitProc RelayInspectMethod,
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
	Extent() int

	//
	ExtentOctets() int64
}

//
//
//
type PriorInspectMethod func(kinds.Tx) error

//
//
//
type RelayInspectMethod func(kinds.Tx, *iface.ReplyInspectTransfer) error

//
//
func PriorInspectMaximumOctets(maximumOctets int64) PriorInspectMethod {
	return func(tx kinds.Tx) error {
		transferExtent := kinds.CalculateSchemaExtentForeachTrans([]kinds.Tx{tx})

		if transferExtent > maximumOctets {
			return fmt.Errorf("REDACTED", transferExtent, maximumOctets)
		}

		return nil
	}
}

//
//
func SubmitInspectMaximumFuel(maximumFuel int64) RelayInspectMethod {
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
type TransferToken [sha256.Size]byte
