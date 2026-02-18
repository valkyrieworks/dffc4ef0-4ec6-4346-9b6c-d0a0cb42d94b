package txpool

import (
	"errors"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
type NoopTxpool struct{}

//
var errNegatePermitted = errors.New("REDACTED")

var _ Txpool = &NoopTxpool{}

//
func (*NoopTxpool) InspectTransfer(kinds.Tx, func(*iface.ReplyInspectTransfer), TransferDetails) error {
	return errNegatePermitted
}

//
func (*NoopTxpool) DeleteTransferByKey(kinds.TransferKey) error { return errNegatePermitted }

//
func (*NoopTxpool) HarvestMaximumOctetsMaximumFuel(int64, int64) kinds.Txs { return nil }

//
func (*NoopTxpool) HarvestMaximumTrans(int) kinds.Txs { return nil }

//
func (*NoopTxpool) Secure() {}

//
func (*NoopTxpool) Release() {}

//
func (*NoopTxpool) Modify(
	int64,
	kinds.Txs,
	[]*iface.InvokeTransferOutcome,
	PreInspectFunction,
	SubmitInspectFunction,
) error {
	return nil
}

//
func (*NoopTxpool) PurgeApplicationLink() error { return nil }

//
func (*NoopTxpool) Purge() {}

//
func (*NoopTxpool) TransAccessible() <-chan struct{} {
	return nil
}

//
func (*NoopTxpool) ActivateTransAccessible() {}

//
func (*NoopTxpool) CollectionTransferDeletedResponse(func(transferKey kinds.TransferKey)) {}

//
func (*NoopTxpool) Volume() int { return 0 }

//
func (*NoopTxpool) VolumeOctets() int64 { return 0 }

//
type NoopTxpoolHandler struct {
	daemon.RootDaemon
}

//
//
//
func NewNoopTxpoolHandler() *NoopTxpoolHandler {
	return &NoopTxpoolHandler{*daemon.NewRootDaemon(nil, "REDACTED", nil)}
}

var _ p2p.Handler = &NoopTxpoolHandler{}

//
func (*NoopTxpoolHandler) WaitAlign() bool { return false }

//
func (*NoopTxpoolHandler) FetchStreams() []*p2p.StreamDefinition { return nil }

//
func (*NoopTxpoolHandler) AppendNode(p2p.Node) {}

//
func (*NoopTxpoolHandler) InitNode(p2p.Node) p2p.Node { return nil }

//
func (*NoopTxpoolHandler) DeleteNode(p2p.Node, any) {}

//
func (*NoopTxpoolHandler) Accept(p2p.Packet) {}

//
func (*NoopTxpoolHandler) CollectionRouter(p2p.Toggeler) {}
