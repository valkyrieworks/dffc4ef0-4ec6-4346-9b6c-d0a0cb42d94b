package agreement

import (
	"context"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/ringlist"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/gateway"
	"github.com/valkyrieworks/kinds"
)

//

type emptyTxpool struct{}

var _ txpool.Txpool = emptyTxpool{}

func (emptyTxpool) Secure()            {}
func (emptyTxpool) Release()          {}
func (emptyTxpool) Volume() int        { return 0 }
func (emptyTxpool) VolumeOctets() int64 { return 0 }
func (emptyTxpool) InspectTransfer(kinds.Tx, func(*iface.ReplyInspectTransfer), txpool.TransferDetails) error {
	return nil
}

func (transmp emptyTxpool) DeleteTransferByKey(kinds.TransferKey) error {
	return nil
}

func (emptyTxpool) HarvestMaximumOctetsMaximumFuel(int64, int64) kinds.Txs { return kinds.Txs{} }
func (emptyTxpool) HarvestMaximumTrans(int) kinds.Txs                  { return kinds.Txs{} }
func (emptyTxpool) Modify(
	int64,
	kinds.Txs,
	[]*iface.InvokeTransferOutcome,
	txpool.PreInspectFunction,
	txpool.SubmitInspectFunction,
) error {
	return nil
}
func (emptyTxpool) Purge()                        {}
func (emptyTxpool) PurgeApplicationLink() error           { return nil }
func (emptyTxpool) TransAccessible() <-chan struct{} { return make(chan struct{}) }
func (emptyTxpool) ActivateTransAccessible()           {}
func (emptyTxpool) TransOctets() int64               { return 0 }

func (emptyTxpool) TransHead() *ringlist.CComponent    { return nil }
func (emptyTxpool) TransWaitChan() <-chan struct{} { return nil }

func (emptyTxpool) InitJournal() error { return nil }
func (emptyTxpool) EndJournal()      {}

//
//
//
//
//

func newEmulateGatewayApplication(completeLedgerReply *iface.ReplyCompleteLedger) gateway.ApplicationLinkAgreement {
	customerOriginator := gateway.NewNativeCustomerOriginator(&emulateGatewayApplication{
		completeLedgerReply: completeLedgerReply,
	})
	cli, _ := customerOriginator.NewIfaceCustomer()
	err := cli.Begin()
	if err != nil {
		panic(err)
	}
	return gateway.NewApplicationLinkAgreement(cli, gateway.NoopStats())
}

type emulateGatewayApplication struct {
	iface.RootSoftware
	completeLedgerReply *iface.ReplyCompleteLedger
}

func (emulate *emulateGatewayApplication) CompleteLedger(context.Context, *iface.QueryCompleteLedger) (*iface.ReplyCompleteLedger, error) {
	return emulate.completeLedgerReply, nil
}
