package kv

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	db "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyTransferPosition(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	tx := kinds.Tx("REDACTED")
	transferOutcome := &iface.TransferOutcome{
		Altitude: 1,
		Ordinal:  0,
		Tx:     tx,
		Outcome: iface.InvokeTransferOutcome{
			Data: []byte{0},
			Cipher: iface.CipherKindOKAY, Log: "REDACTED", Incidents: nil,
		},
	}
	digest := tx.Digest()

	cluster := transferordinal.FreshCluster(1)
	if err := cluster.Add(transferOutcome); err != nil {
		t.Error(err)
	}
	err := ordinalizer.AppendCluster(cluster)
	require.NoError(t, err)

	retrievedTransferOutcome, err := ordinalizer.Get(digest)
	require.NoError(t, err)
	assert.True(t, proto.Equal(transferOutcome, retrievedTransferOutcome))

	tx2 := kinds.Tx("REDACTED")
	transferOutcome2 := &iface.TransferOutcome{
		Altitude: 1,
		Ordinal:  0,
		Tx:     tx2,
		Outcome: iface.InvokeTransferOutcome{
			Data: []byte{0},
			Cipher: iface.CipherKindOKAY, Log: "REDACTED", Incidents: nil,
		},
	}
	digest2 := tx2.Digest()

	err = ordinalizer.Ordinal(transferOutcome2)
	require.NoError(t, err)

	retrievedTransferOutcome2, err := ordinalizer.Get(digest2)
	require.NoError(t, err)
	assert.True(t, proto.Equal(transferOutcome2, retrievedTransferOutcome2))
}

func VerifyTransferLookup(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	digest := kinds.Tx(transferOutcome.Tx).Digest()

	err := ordinalizer.Ordinal(transferOutcome)
	require.NoError(t, err)

	verifyScenarios := []struct {
		q             string
		outcomesMagnitude int
	}{
		//
		{fmt.Sprintf("REDACTED", digest), 1},
		//
		{fmt.Sprintf("REDACTED", digest), 1},
		//
		{"REDACTED", 1},
		//
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 1},
		//
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		//
		{"REDACTED", 0},
		//
		{"REDACTED", 1},
		//
		{"REDACTED", 0},
		//
		{"REDACTED", 1},
		//
		{"REDACTED", 1},
		{"REDACTED", 1},
		//
		{"REDACTED", 0},
		{"REDACTED", 0},
		//
		{"REDACTED", 0},
		//
		{"REDACTED", 0},
		//
		{"REDACTED", 1},
		//
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		{"REDACTED", 0},
		//
		{"REDACTED", 0},
		//
		{"REDACTED", 1},
		//
		{"REDACTED", 0},
		{"REDACTED", 0},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble(tc.q))
			assert.NoError(t, err)

			assert.Len(t, outcomes, tc.outcomesMagnitude)
			if tc.outcomesMagnitude > 0 {
				for _, txr := range outcomes {
					assert.True(t, proto.Equal(transferOutcome, txr))
				}
			}
		})
	}
}

func VerifyTransferLookupIncidentAlign(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}, {Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}, {Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: false}, {Key: "REDACTED", Datum: "REDACTED", Ordinal: false}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})

	err := ordinalizer.Ordinal(transferOutcome)
	require.NoError(t, err)

	verifyScenarios := map[string]struct {
		q             string
		outcomesMagnitude int
	}{
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble(tc.q))
			assert.NoError(t, err)

			assert.Len(t, outcomes, tc.outcomesMagnitude)
			if tc.outcomesMagnitude > 0 {
				for _, txr := range outcomes {
					assert.True(t, proto.Equal(transferOutcome, txr))
				}
			}
		})
	}
}

func VerifyTransferLookupIncidentAlignViaAltitude(t *testing.T) {

	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}, {Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})

	err := ordinalizer.Ordinal(transferOutcome)
	require.NoError(t, err)

	transferOutcome10 := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}, {Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	transferOutcome10.Tx = kinds.Tx("REDACTED")
	transferOutcome10.Altitude = 10

	err = ordinalizer.Ordinal(transferOutcome10)
	require.NoError(t, err)

	verifyScenarios := map[string]struct {
		q             string
		outcomesMagnitude int
	}{
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 2,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesMagnitude: 2,
		},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble(tc.q))
			assert.NoError(t, err)

			assert.Len(t, outcomes, tc.outcomesMagnitude)
			if tc.outcomesMagnitude > 0 {
				for _, txr := range outcomes {
					switch txr.Altitude {
					case 1:
						assert.True(t, proto.Equal(transferOutcome, txr))
					case 10:
						assert.True(t, proto.Equal(transferOutcome10, txr))
					default:
						assert.True(t, false)
					}
				}
			}
		})
	}
}

func VerifyTransferLookupUsingRevocation(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	err := ordinalizer.Ordinal(transferOutcome)
	require.NoError(t, err)

	ctx, abort := context.WithCancel(context.Background())
	abort()
	outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble("REDACTED"))
	assert.NoError(t, err)
	assert.Empty(t, outcomes)
}

func VerifyTransferLookupObsoleteCataloging(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	//
	transferOutcome1 := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	digest1 := kinds.Tx(transferOutcome1.Tx).Digest()

	err := ordinalizer.Ordinal(transferOutcome1)
	require.NoError(t, err)

	//
	transferOutcome2 := transferOutcomeUsingIncidents(nil)
	transferOutcome2.Tx = kinds.Tx("REDACTED")

	digest2 := kinds.Tx(transferOutcome2.Tx).Digest()
	b := ordinalizer.depot.FreshCluster()

	crudeOctets, err := proto.Marshal(transferOutcome2)
	require.NoError(t, err)

	dependToken := []byte(fmt.Sprintf("REDACTED",
		"REDACTED",
		"REDACTED",
		transferOutcome2.Altitude,
		transferOutcome2.Ordinal,
	))

	err = b.Set(dependToken, digest2)
	require.NoError(t, err)
	err = b.Set(tokenForeachAltitude(transferOutcome2), digest2)
	require.NoError(t, err)
	err = b.Set(digest2, crudeOctets)
	require.NoError(t, err)
	err = b.Record()
	require.NoError(t, err)

	verifyScenarios := []struct {
		q       string
		outcomes []*iface.TransferOutcome
	}{
		//
		{fmt.Sprintf("REDACTED", digest1), []*iface.TransferOutcome{transferOutcome1}},
		//
		{fmt.Sprintf("REDACTED", digest2), []*iface.TransferOutcome{transferOutcome2}},
		//
		{"REDACTED", []*iface.TransferOutcome{transferOutcome1}},
		{"REDACTED", []*iface.TransferOutcome{transferOutcome1}},
		//
		{"REDACTED", []*iface.TransferOutcome{transferOutcome1}},
		//
		{"REDACTED", []*iface.TransferOutcome{transferOutcome1}},
		//
		{"REDACTED", []*iface.TransferOutcome{}},
		//
		{"REDACTED", []*iface.TransferOutcome{}},
		//
		{"REDACTED", []*iface.TransferOutcome{}},
		//
		{"REDACTED", []*iface.TransferOutcome{transferOutcome2}},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble(tc.q))
			require.NoError(t, err)
			for _, txr := range outcomes {
				for _, tr := range tc.outcomes {
					assert.True(t, proto.Equal(tr, txr))
				}
			}
		})
	}
}

func VerifyTransferLookupSingleTransferUsingVariousIdenticalLabelsHoweverDistinctItems(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: false}}},
	})

	err := ordinalizer.Ordinal(transferOutcome)
	require.NoError(t, err)

	verifyScenarios := []struct {
		alias  string
		q     string
		detected bool
	}{
		{
			q:     "REDACTED",
			detected: true,
		},
		{
			q:     "REDACTED",
			detected: false,
		},
		{
			q:     "REDACTED",
			detected: true,
		},
		{
			q:     "REDACTED",
			detected: true,
		},

		{
			q:     "REDACTED",
			detected: true,
		},

		{
			q:     "REDACTED",
			detected: false,
		},
		{
			q:     "REDACTED",
			detected: false,
		},
		{
			q:     "REDACTED",
			detected: true,
		},
		{
			q:     "REDACTED",
			detected: true,
		},
		{
			q:     "REDACTED",
			detected: true,
		},
		{
			q:     "REDACTED",
			detected: false,
		},
		{
			q:     "REDACTED",
			detected: false,
		},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {
		outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble(tc.q))
		assert.NoError(t, err)
		n := 0
		if tc.detected {
			n = 1
		}
		assert.Len(t, outcomes, n)
		assert.True(t, !tc.detected || proto.Equal(transferOutcome, outcomes[0]))

	}
}

func VerifyTransferPositionReplicatedFormerlyTriumphant(t *testing.T) {
	simulateTransfer := kinds.Tx("REDACTED")

	verifyScenarios := []struct {
		alias         string
		tx1          *iface.TransferOutcome
		tx2          *iface.TransferOutcome
		expirationSupersede bool //
	}{
		{
			"REDACTED",
			&iface.TransferOutcome{
				Altitude: 1,
				Ordinal:  0,
				Tx:     simulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Cipher: iface.CipherKindOKAY,
				},
			},
			&iface.TransferOutcome{
				Altitude: 2,
				Ordinal:  0,
				Tx:     simulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Cipher: iface.CipherKindOKAY + 1,
				},
			},
			false,
		},
		{
			"REDACTED",
			&iface.TransferOutcome{
				Altitude: 1,
				Ordinal:  0,
				Tx:     simulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Cipher: iface.CipherKindOKAY + 1,
				},
			},
			&iface.TransferOutcome{
				Altitude: 2,
				Ordinal:  0,
				Tx:     simulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Cipher: iface.CipherKindOKAY + 1,
				},
			},
			true,
		},
		{
			"REDACTED",
			&iface.TransferOutcome{
				Altitude: 1,
				Ordinal:  0,
				Tx:     simulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Cipher: iface.CipherKindOKAY,
				},
			},
			&iface.TransferOutcome{
				Altitude: 2,
				Ordinal:  0,
				Tx:     simulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Cipher: iface.CipherKindOKAY,
				},
			},
			true,
		},
	}

	digest := simulateTransfer.Digest()

	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(t *testing.T) {
			ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

			//
			err := ordinalizer.Ordinal(tc.tx1)
			require.NoError(t, err)

			//
			err = ordinalizer.Ordinal(tc.tx2)
			require.NoError(t, err)

			res, err := ordinalizer.Get(digest)
			require.NoError(t, err)

			if tc.expirationSupersede {
				require.Equal(t, tc.tx2, res)
			} else {
				require.Equal(t, tc.tx1, res)
			}
		})
	}
}

func VerifyTransferLookupVariousTrans(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	//
	transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})

	transferOutcome.Tx = kinds.Tx("REDACTED")
	transferOutcome.Altitude = 2
	transferOutcome.Ordinal = 1
	err := ordinalizer.Ordinal(transferOutcome)
	require.NoError(t, err)

	//
	transferOutcome2 := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	transferOutcome2.Tx = kinds.Tx("REDACTED")
	transferOutcome2.Altitude = 1
	transferOutcome2.Ordinal = 2

	err = ordinalizer.Ordinal(transferOutcome2)
	require.NoError(t, err)

	//
	transferOutcome3 := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	transferOutcome3.Tx = kinds.Tx("REDACTED")
	transferOutcome3.Altitude = 1
	transferOutcome3.Ordinal = 1
	err = ordinalizer.Ordinal(transferOutcome3)
	require.NoError(t, err)

	//
	//
	transferOutcome4 := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	transferOutcome4.Tx = kinds.Tx("REDACTED")
	transferOutcome4.Altitude = 2
	transferOutcome4.Ordinal = 2
	err = ordinalizer.Ordinal(transferOutcome4)
	require.NoError(t, err)

	ctx := context.Background()

	outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble("REDACTED"))
	assert.NoError(t, err)

	require.Len(t, outcomes, 3)
}

func transferOutcomeUsingIncidents(incidents []iface.Incident) *iface.TransferOutcome {
	tx := kinds.Tx("REDACTED")
	return &iface.TransferOutcome{
		Altitude: 1,
		Ordinal:  0,
		Tx:     tx,
		Outcome: iface.InvokeTransferOutcome{
			Data:   []byte{0},
			Cipher:   iface.CipherKindOKAY,
			Log:    "REDACTED",
			Incidents: incidents,
		},
	}
}

func assessmentTransferPosition(transTally int64, b *testing.B) {
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(b, err)
	defer os.RemoveAll(dir)

	depot, err := db.FreshDatastore("REDACTED", "REDACTED", dir)
	require.NoError(b, err)
	ordinalizer := FreshTransferOrdinal(depot)

	cluster := transferordinal.FreshCluster(transTally)
	transferOrdinal := uint32(0)
	for i := int64(0); i < transTally; i++ {
		tx := commitrand.Octets(250)
		transferOutcome := &iface.TransferOutcome{
			Altitude: 1,
			Ordinal:  transferOrdinal,
			Tx:     tx,
			Outcome: iface.InvokeTransferOutcome{
				Data:   []byte{0},
				Cipher:   iface.CipherKindOKAY,
				Log:    "REDACTED",
				Incidents: []iface.Incident{},
			},
		}
		if err := cluster.Add(transferOutcome); err != nil {
			b.Fatal(err)
		}
		transferOrdinal++
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		err = ordinalizer.AppendCluster(cluster)
	}
	if err != nil {
		b.Fatal(err)
	}
}

func VerifyAmpleInteger(t *testing.T) {
	ordinalizer := FreshTransferOrdinal(db.FreshMemoryDatastore())

	ampleInteger := "REDACTED"
	ampleIntegerAdd1 := "REDACTED"
	ampleDecimal := ampleInteger + "REDACTED"
	ampleDecimalLesser := ampleInteger + "REDACTED"
	ampleDecimalTinier := "REDACTED" + "REDACTED"
	ampleIntegerTinier := "REDACTED"

	transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleInteger, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleDecimalTinier, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleIntegerAdd1, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleDecimalLesser, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
	})
	digest := kinds.Tx(transferOutcome.Tx).Digest()

	err := ordinalizer.Ordinal(transferOutcome)

	require.NoError(t, err)

	transferOutcome2 := transferOutcomeUsingIncidents([]iface.Incident{
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleDecimal, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleDecimal, Ordinal: true}, {Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleIntegerTinier, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: ampleInteger, Ordinal: true}, {Key: "REDACTED", Datum: "REDACTED", Ordinal: true}}}})

	transferOutcome2.Tx = kinds.Tx("REDACTED")
	transferOutcome2.Altitude = 2
	transferOutcome2.Ordinal = 2

	digest2 := kinds.Tx(transferOutcome2.Tx).Digest()

	err = ordinalizer.Ordinal(transferOutcome2)
	require.NoError(t, err)
	verifyScenarios := []struct {
		q             string
		transferResult         *iface.TransferOutcome
		outcomesMagnitude int
	}{
		//
		{fmt.Sprintf("REDACTED", digest), transferOutcome, 1},
		//
		{fmt.Sprintf("REDACTED", digest), transferOutcome, 1},
		{fmt.Sprintf("REDACTED", digest2), transferOutcome2, 1},
		//
		{"REDACTED" + ampleInteger, nil, 2},
		//
		{"REDACTED" + ampleInteger + "REDACTED", nil, 2},
		{"REDACTED" + ampleInteger + "REDACTED", nil, 0},
		//
		{"REDACTED" + ampleInteger + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + ampleInteger + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + ampleInteger + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + ampleDecimalTinier + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + ampleInteger + "REDACTED", nil, 2},
		{"REDACTED" + ampleInteger + "REDACTED", nil, 1},
		{"REDACTED" + ampleInteger + "REDACTED", nil, 1},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(ctx, inquire.ShouldAssemble(tc.q))
			assert.NoError(t, err)
			assert.Len(t, outcomes, tc.outcomesMagnitude)
			if tc.outcomesMagnitude > 0 && tc.transferResult != nil {
				assert.True(t, proto.Equal(outcomes[0], tc.transferResult))
			}
		})
	}
}

func AssessmentTransferCatalog1(b *testing.B)     { assessmentTransferPosition(1, b) }
func AssessmentTransferCatalog500(b *testing.B)   { assessmentTransferPosition(500, b) }
func AssessmentTransferCatalog1000(b *testing.B)  { assessmentTransferPosition(1000, b) }
func AssessmentTransferCatalog2000(b *testing.B)  { assessmentTransferPosition(2000, b) }
func AssessmentTransferCatalog10000(b *testing.B) { assessmentTransferPosition(10000, b) }
