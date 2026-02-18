package transferordinal_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	db "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/log"
	ledgerordinalkv "github.com/valkyrieworks/status/ordinaler/ledger/kv"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/status/transordinal/kv"
	"github.com/valkyrieworks/kinds"
)

func VerifyOrdinalerDaemonListingsLedgers(t *testing.T) {
	//
	eventBus := kinds.NewEventBus()
	eventBus.AssignTracer(log.VerifyingTracer())
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	depot := db.NewMemoryStore()
	transOrdinaler := kv.NewTransOrdinal(depot)
	ledgerOrdinaler := ledgerordinalkv.New(db.NewHeadingStore(depot, []byte("REDACTED")))

	daemon := transordinal.NewOrdinalerDaemon(transOrdinaler, ledgerOrdinaler, eventBus, false)
	daemon.AssignTracer(log.VerifyingTracer())
	err = daemon.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := daemon.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	err = eventBus.BroadcastEventNewLedgerEvents(kinds.EventDataNewLedgerEvents{
		Level: 1,
		Events: []iface.Event{
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
		},
		CountTrans: int64(2),
	})
	require.NoError(t, err)
	transferOutcome1 := &iface.TransOutcome{
		Level: 1,
		Ordinal:  uint32(0),
		Tx:     kinds.Tx("REDACTED"),
		Outcome: iface.InvokeTransferOutcome{Code: 0},
	}
	err = eventBus.BroadcastEventTransfer(kinds.EventDataTransfer{TransOutcome: *transferOutcome1})
	require.NoError(t, err)
	transferOutcome2 := &iface.TransOutcome{
		Level: 1,
		Ordinal:  uint32(1),
		Tx:     kinds.Tx("REDACTED"),
		Outcome: iface.InvokeTransferOutcome{Code: 0},
	}
	err = eventBus.BroadcastEventTransfer(kinds.EventDataTransfer{TransOutcome: *transferOutcome2})
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	res, err := transOrdinaler.Get(kinds.Tx("REDACTED").Digest())
	require.NoError(t, err)
	require.Equal(t, transferOutcome1, res)

	ok, err := ledgerOrdinaler.Has(1)
	require.NoError(t, err)
	require.True(t, ok)

	res, err = transOrdinaler.Get(kinds.Tx("REDACTED").Digest())
	require.NoError(t, err)
	require.Equal(t, transferOutcome2, res)
}
