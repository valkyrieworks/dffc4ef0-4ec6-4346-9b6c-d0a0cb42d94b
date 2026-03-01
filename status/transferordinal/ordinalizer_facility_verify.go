package txmark_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	db "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	ledgeridxkv "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger/kv"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/kv"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyOrdinalizerFacilityCatalogsLedgers(t *testing.T) {
	//
	incidentPipeline := kinds.FreshIncidentPipeline()
	incidentPipeline.AssignTracer(log.VerifyingTracer())
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	depot := db.FreshMemoryDatastore()
	transferOrdinalizer := kv.FreshTransferOrdinal(depot)
	ledgerOrdinalizer := ledgeridxkv.New(db.FreshHeadingDatastore(depot, []byte("REDACTED")))

	facility := transferordinal.FreshOrdinalizerFacility(transferOrdinalizer, ledgerOrdinalizer, incidentPipeline, false)
	facility.AssignTracer(log.VerifyingTracer())
	err = facility.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := facility.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	err = incidentPipeline.BroadcastIncidentFreshLedgerIncidents(kinds.IncidentDataFreshLedgerIncidents{
		Altitude: 1,
		Incidents: []iface.Incident{
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
				},
			},
		},
		CountTrans: int64(2),
	})
	require.NoError(t, err)
	transferOutcome1 := &iface.TransferOutcome{
		Altitude: 1,
		Ordinal:  uint32(0),
		Tx:     kinds.Tx("REDACTED"),
		Outcome: iface.InvokeTransferOutcome{Cipher: 0},
	}
	err = incidentPipeline.BroadcastIncidentTransfer(kinds.IncidentDataTransfer{TransferOutcome: *transferOutcome1})
	require.NoError(t, err)
	transferOutcome2 := &iface.TransferOutcome{
		Altitude: 1,
		Ordinal:  uint32(1),
		Tx:     kinds.Tx("REDACTED"),
		Outcome: iface.InvokeTransferOutcome{Cipher: 0},
	}
	err = incidentPipeline.BroadcastIncidentTransfer(kinds.IncidentDataTransfer{TransferOutcome: *transferOutcome2})
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	res, err := transferOrdinalizer.Get(kinds.Tx("REDACTED").Digest())
	require.NoError(t, err)
	require.Equal(t, transferOutcome1, res)

	ok, err := ledgerOrdinalizer.Has(1)
	require.NoError(t, err)
	require.True(t, ok)

	res, err = transferOrdinalizer.Get(kinds.Tx("REDACTED").Digest())
	require.NoError(t, err)
	require.Equal(t, transferOutcome2, res)
}
