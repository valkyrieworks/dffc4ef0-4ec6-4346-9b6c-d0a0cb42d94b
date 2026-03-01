package agreement

import (
	"context"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/linkedlist"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//

type blankTxpool struct{}

var _ txpooll.Txpool = blankTxpool{}

func (blankTxpool) Secure()            {}
func (blankTxpool) Release()          {}
func (blankTxpool) Extent() int        { return 0 }
func (blankTxpool) ExtentOctets() int64 { return 0 }
func (blankTxpool) InspectTransfer(kinds.Tx, func(*iface.ReplyInspectTransfer), txpooll.TransferDetails) error {
	return nil
}

func (txpoolmp blankTxpool) DiscardTransferViaToken(kinds.TransferToken) error {
	return nil
}

func (blankTxpool) HarvestMaximumOctetsMaximumFuel(int64, int64) kinds.Txs { return kinds.Txs{} }
func (blankTxpool) HarvestMaximumTrans(int) kinds.Txs                  { return kinds.Txs{} }
func (blankTxpool) Revise(
	int64,
	kinds.Txs,
	[]*iface.InvokeTransferOutcome,
	txpooll.PriorInspectMethod,
	txpooll.RelayInspectMethod,
) error {
	return nil
}
func (blankTxpool) Purge()                        {}
func (blankTxpool) PurgeApplicationLink() error           { return nil }
func (blankTxpool) TransAccessible() <-chan struct{} { return make(chan struct{}) }
func (blankTxpool) ActivateTransAccessible()           {}
func (blankTxpool) TransOctets() int64               { return 0 }

func (blankTxpool) TransLeading() *linkedlist.CNComponent    { return nil }
func (blankTxpool) TransPauseChannel() <-chan struct{} { return nil }

func (blankTxpool) InitializeJournal() error { return nil }
func (blankTxpool) ShutdownJournal()      {}

//
//
//
//
//

func freshSimulateDelegateApplication(culminateLedgerReply *iface.ReplyCulminateLedger) delegate.ApplicationLinkAgreement {
	customerOriginator := delegate.FreshRegionalCustomerOriginator(&simulateDelegateApplication{
		culminateLedgerReply: culminateLedgerReply,
	})
	cli, _ := customerOriginator.FreshIfaceCustomer()
	err := cli.Initiate()
	if err != nil {
		panic(err)
	}
	return delegate.FreshApplicationLinkAgreement(cli, delegate.NooperationTelemetry())
}

type simulateDelegateApplication struct {
	iface.FoundationPlatform
	culminateLedgerReply *iface.ReplyCulminateLedger
}

func (simulate *simulateDelegateApplication) CulminateLedger(context.Context, *iface.SolicitCulminateLedger) (*iface.ReplyCulminateLedger, error) {
	return simulate.culminateLedgerReply, nil
}
