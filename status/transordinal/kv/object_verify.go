package kv

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	db "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/kinds"
)

func VerifyTransferOrdinal(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	tx := kinds.Tx("REDACTED")
	transOutcome := &iface.TransOutcome{
		Level: 1,
		Ordinal:  0,
		Tx:     tx,
		Outcome: iface.InvokeTransferOutcome{
			Data: []byte{0},
			Code: iface.CodeKindSuccess, Log: "REDACTED", Events: nil,
		},
	}
	digest := tx.Digest()

	group := transordinal.NewGroup(1)
	if err := group.Add(transOutcome); err != nil {
		t.Error(err)
	}
	err := ordinaler.AppendGroup(group)
	require.NoError(t, err)

	retrievedTransferOutcome, err := ordinaler.Get(digest)
	require.NoError(t, err)
	assert.True(t, proto.Equal(transOutcome, retrievedTransferOutcome))

	tx2 := kinds.Tx("REDACTED")
	transferOutcome2 := &iface.TransOutcome{
		Level: 1,
		Ordinal:  0,
		Tx:     tx2,
		Outcome: iface.InvokeTransferOutcome{
			Data: []byte{0},
			Code: iface.CodeKindSuccess, Log: "REDACTED", Events: nil,
		},
	}
	digest2 := tx2.Digest()

	err = ordinaler.Ordinal(transferOutcome2)
	require.NoError(t, err)

	retrievedTransferOutcome2, err := ordinaler.Get(digest2)
	require.NoError(t, err)
	assert.True(t, proto.Equal(transferOutcome2, retrievedTransferOutcome2))
}

func VerifyTransferScan(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	transOutcome := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	digest := kinds.Tx(transOutcome.Tx).Digest()

	err := ordinaler.Ordinal(transOutcome)
	require.NoError(t, err)

	verifyScenarios := []struct {
		q             string
		outcomesExtent int
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
			outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild(tc.q))
			assert.NoError(t, err)

			assert.Len(t, outcomes, tc.outcomesExtent)
			if tc.outcomesExtent > 0 {
				for _, txr := range outcomes {
					assert.True(t, proto.Equal(transOutcome, txr))
				}
			}
		})
	}
}

func VerifyTransferScanEventAlign(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	transOutcome := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}, {Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}, {Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: false}, {Key: "REDACTED", Item: "REDACTED", Ordinal: false}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})

	err := ordinaler.Ordinal(transOutcome)
	require.NoError(t, err)

	verifyScenarios := map[string]struct {
		q             string
		outcomesExtent int
	}{
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild(tc.q))
			assert.NoError(t, err)

			assert.Len(t, outcomes, tc.outcomesExtent)
			if tc.outcomesExtent > 0 {
				for _, txr := range outcomes {
					assert.True(t, proto.Equal(transOutcome, txr))
				}
			}
		})
	}
}

func VerifyTransferScanEventAlignByLevel(t *testing.T) {

	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	transOutcome := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}, {Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})

	err := ordinaler.Ordinal(transOutcome)
	require.NoError(t, err)

	transferOutcome10 := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}, {Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	transferOutcome10.Tx = kinds.Tx("REDACTED")
	transferOutcome10.Level = 10

	err = ordinaler.Ordinal(transferOutcome10)
	require.NoError(t, err)

	verifyScenarios := map[string]struct {
		q             string
		outcomesExtent int
	}{
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 0,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 2,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 1,
		},
		"REDACTED": {
			q:             "REDACTED",
			outcomesExtent: 2,
		},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild(tc.q))
			assert.NoError(t, err)

			assert.Len(t, outcomes, tc.outcomesExtent)
			if tc.outcomesExtent > 0 {
				for _, txr := range outcomes {
					switch txr.Level {
					case 1:
						assert.True(t, proto.Equal(transOutcome, txr))
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

func VerifyTransferScanWithAbortion(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	transOutcome := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	err := ordinaler.Ordinal(transOutcome)
	require.NoError(t, err)

	ctx, revoke := context.WithCancel(context.Background())
	revoke()
	outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild("REDACTED"))
	assert.NoError(t, err)
	assert.Empty(t, outcomes)
}

func VerifyTransferScanObsoleteCataloging(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	//
	transferOutcome1 := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	digest1 := kinds.Tx(transferOutcome1.Tx).Digest()

	err := ordinaler.Ordinal(transferOutcome1)
	require.NoError(t, err)

	//
	transferOutcome2 := transferOutcomeWithEvents(nil)
	transferOutcome2.Tx = kinds.Tx("REDACTED")

	digest2 := kinds.Tx(transferOutcome2.Tx).Digest()
	b := ordinaler.depot.NewGroup()

	crudeOctets, err := proto.Marshal(transferOutcome2)
	require.NoError(t, err)

	dependKey := []byte(fmt.Sprintf("REDACTED",
		"REDACTED",
		"REDACTED",
		transferOutcome2.Level,
		transferOutcome2.Ordinal,
	))

	err = b.Set(dependKey, digest2)
	require.NoError(t, err)
	err = b.Set(keyForLevel(transferOutcome2), digest2)
	require.NoError(t, err)
	err = b.Set(digest2, crudeOctets)
	require.NoError(t, err)
	err = b.Record()
	require.NoError(t, err)

	verifyScenarios := []struct {
		q       string
		outcomes []*iface.TransOutcome
	}{
		//
		{fmt.Sprintf("REDACTED", digest1), []*iface.TransOutcome{transferOutcome1}},
		//
		{fmt.Sprintf("REDACTED", digest2), []*iface.TransOutcome{transferOutcome2}},
		//
		{"REDACTED", []*iface.TransOutcome{transferOutcome1}},
		{"REDACTED", []*iface.TransOutcome{transferOutcome1}},
		//
		{"REDACTED", []*iface.TransOutcome{transferOutcome1}},
		//
		{"REDACTED", []*iface.TransOutcome{transferOutcome1}},
		//
		{"REDACTED", []*iface.TransOutcome{}},
		//
		{"REDACTED", []*iface.TransOutcome{}},
		//
		{"REDACTED", []*iface.TransOutcome{}},
		//
		{"REDACTED", []*iface.TransOutcome{transferOutcome2}},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild(tc.q))
			require.NoError(t, err)
			for _, txr := range outcomes {
				for _, tr := range tc.outcomes {
					assert.True(t, proto.Equal(tr, txr))
				}
			}
		})
	}
}

func VerifyTransferScanOneTransferWithVariedIdenticalMarkersButDistinctItems(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	transOutcome := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: false}}},
	})

	err := ordinaler.Ordinal(transOutcome)
	require.NoError(t, err)

	verifyScenarios := []struct {
		label  string
		q     string
		located bool
	}{
		{
			q:     "REDACTED",
			located: true,
		},
		{
			q:     "REDACTED",
			located: false,
		},
		{
			q:     "REDACTED",
			located: true,
		},
		{
			q:     "REDACTED",
			located: true,
		},

		{
			q:     "REDACTED",
			located: true,
		},

		{
			q:     "REDACTED",
			located: false,
		},
		{
			q:     "REDACTED",
			located: false,
		},
		{
			q:     "REDACTED",
			located: true,
		},
		{
			q:     "REDACTED",
			located: true,
		},
		{
			q:     "REDACTED",
			located: true,
		},
		{
			q:     "REDACTED",
			located: false,
		},
		{
			q:     "REDACTED",
			located: false,
		},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {
		outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild(tc.q))
		assert.NoError(t, err)
		n := 0
		if tc.located {
			n = 1
		}
		assert.Len(t, outcomes, n)
		assert.True(t, !tc.located || proto.Equal(transOutcome, outcomes[0]))

	}
}

func VerifyTransferOrdinalReplicatedEarlierTriumphant(t *testing.T) {
	emulateTransfer := kinds.Tx("REDACTED")

	verifyScenarios := []struct {
		label         string
		tx1          *iface.TransOutcome
		tx2          *iface.TransOutcome
		expirationOverride bool //
	}{
		{
			"REDACTED",
			&iface.TransOutcome{
				Level: 1,
				Ordinal:  0,
				Tx:     emulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Code: iface.CodeKindSuccess,
				},
			},
			&iface.TransOutcome{
				Level: 2,
				Ordinal:  0,
				Tx:     emulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Code: iface.CodeKindSuccess + 1,
				},
			},
			false,
		},
		{
			"REDACTED",
			&iface.TransOutcome{
				Level: 1,
				Ordinal:  0,
				Tx:     emulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Code: iface.CodeKindSuccess + 1,
				},
			},
			&iface.TransOutcome{
				Level: 2,
				Ordinal:  0,
				Tx:     emulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Code: iface.CodeKindSuccess + 1,
				},
			},
			true,
		},
		{
			"REDACTED",
			&iface.TransOutcome{
				Level: 1,
				Ordinal:  0,
				Tx:     emulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Code: iface.CodeKindSuccess,
				},
			},
			&iface.TransOutcome{
				Level: 2,
				Ordinal:  0,
				Tx:     emulateTransfer,
				Outcome: iface.InvokeTransferOutcome{
					Code: iface.CodeKindSuccess,
				},
			},
			true,
		},
	}

	digest := emulateTransfer.Digest()

	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			ordinaler := NewTransOrdinal(db.NewMemoryStore())

			//
			err := ordinaler.Ordinal(tc.tx1)
			require.NoError(t, err)

			//
			err = ordinaler.Ordinal(tc.tx2)
			require.NoError(t, err)

			res, err := ordinaler.Get(digest)
			require.NoError(t, err)

			if tc.expirationOverride {
				require.Equal(t, tc.tx2, res)
			} else {
				require.Equal(t, tc.tx1, res)
			}
		})
	}
}

func VerifyTransferScanVariedTrans(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	//
	transOutcome := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})

	transOutcome.Tx = kinds.Tx("REDACTED")
	transOutcome.Level = 2
	transOutcome.Ordinal = 1
	err := ordinaler.Ordinal(transOutcome)
	require.NoError(t, err)

	//
	transferOutcome2 := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	transferOutcome2.Tx = kinds.Tx("REDACTED")
	transferOutcome2.Level = 1
	transferOutcome2.Ordinal = 2

	err = ordinaler.Ordinal(transferOutcome2)
	require.NoError(t, err)

	//
	transferOutcome3 := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	transferOutcome3.Tx = kinds.Tx("REDACTED")
	transferOutcome3.Level = 1
	transferOutcome3.Ordinal = 1
	err = ordinaler.Ordinal(transferOutcome3)
	require.NoError(t, err)

	//
	//
	transferOutcome4 := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	transferOutcome4.Tx = kinds.Tx("REDACTED")
	transferOutcome4.Level = 2
	transferOutcome4.Ordinal = 2
	err = ordinaler.Ordinal(transferOutcome4)
	require.NoError(t, err)

	ctx := context.Background()

	outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild("REDACTED"))
	assert.NoError(t, err)

	require.Len(t, outcomes, 3)
}

func transferOutcomeWithEvents(events []iface.Event) *iface.TransOutcome {
	tx := kinds.Tx("REDACTED")
	return &iface.TransOutcome{
		Level: 1,
		Ordinal:  0,
		Tx:     tx,
		Outcome: iface.InvokeTransferOutcome{
			Data:   []byte{0},
			Code:   iface.CodeKindSuccess,
			Log:    "REDACTED",
			Events: events,
		},
	}
}

func criterionTransferOrdinal(transNumber int64, b *testing.B) {
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(b, err)
	defer os.RemoveAll(dir)

	depot, err := db.NewStore("REDACTED", "REDACTED", dir)
	require.NoError(b, err)
	ordinaler := NewTransOrdinal(depot)

	group := transordinal.NewGroup(transNumber)
	transOrdinal := uint32(0)
	for i := int64(0); i < transNumber; i++ {
		tx := engineseed.Octets(250)
		transOutcome := &iface.TransOutcome{
			Level: 1,
			Ordinal:  transOrdinal,
			Tx:     tx,
			Outcome: iface.InvokeTransferOutcome{
				Data:   []byte{0},
				Code:   iface.CodeKindSuccess,
				Log:    "REDACTED",
				Events: []iface.Event{},
			},
		}
		if err := group.Add(transOutcome); err != nil {
			b.Fatal(err)
		}
		transOrdinal++
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		err = ordinaler.AppendGroup(group)
	}
	if err != nil {
		b.Fatal(err)
	}
}

func VerifyLargeInteger(t *testing.T) {
	ordinaler := NewTransOrdinal(db.NewMemoryStore())

	largeInteger := "REDACTED"
	largeIntegerAdd1 := "REDACTED"
	largeFloat := largeInteger + "REDACTED"
	largeFloatLesser := largeInteger + "REDACTED"
	largeFloatLess := "REDACTED" + "REDACTED"
	largeIntegerLess := "REDACTED"

	transOutcome := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeInteger, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeFloatLess, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeIntegerAdd1, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeFloatLesser, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
	})
	digest := kinds.Tx(transOutcome.Tx).Digest()

	err := ordinaler.Ordinal(transOutcome)

	require.NoError(t, err)

	transferOutcome2 := transferOutcomeWithEvents([]iface.Event{
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeFloat, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeFloat, Ordinal: true}, {Key: "REDACTED", Item: "REDACTED", Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeIntegerLess, Ordinal: true}}},
		{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: largeInteger, Ordinal: true}, {Key: "REDACTED", Item: "REDACTED", Ordinal: true}}}})

	transferOutcome2.Tx = kinds.Tx("REDACTED")
	transferOutcome2.Level = 2
	transferOutcome2.Ordinal = 2

	digest2 := kinds.Tx(transferOutcome2.Tx).Digest()

	err = ordinaler.Ordinal(transferOutcome2)
	require.NoError(t, err)
	verifyScenarios := []struct {
		q             string
		transferOutput         *iface.TransOutcome
		outcomesExtent int
	}{
		//
		{fmt.Sprintf("REDACTED", digest), transOutcome, 1},
		//
		{fmt.Sprintf("REDACTED", digest), transOutcome, 1},
		{fmt.Sprintf("REDACTED", digest2), transferOutcome2, 1},
		//
		{"REDACTED" + largeInteger, nil, 2},
		//
		{"REDACTED" + largeInteger + "REDACTED", nil, 2},
		{"REDACTED" + largeInteger + "REDACTED", nil, 0},
		//
		{"REDACTED" + largeInteger + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + largeInteger + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + largeInteger + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + largeFloatLess + "REDACTED", transferOutcome2, 1},
		{"REDACTED" + largeInteger + "REDACTED", nil, 2},
		{"REDACTED" + largeInteger + "REDACTED", nil, 1},
		{"REDACTED" + largeInteger + "REDACTED", nil, 1},
	}

	ctx := context.Background()

	for _, tc := range verifyScenarios {

		t.Run(tc.q, func(t *testing.T) {
			outcomes, err := ordinaler.Scan(ctx, inquire.ShouldBuild(tc.q))
			assert.NoError(t, err)
			assert.Len(t, outcomes, tc.outcomesExtent)
			if tc.outcomesExtent > 0 && tc.transferOutput != nil {
				assert.True(t, proto.Equal(outcomes[0], tc.transferOutput))
			}
		})
	}
}

func CriterionTransferOrdinal1(b *testing.B)     { criterionTransferOrdinal(1, b) }
func CriterionTransferOrdinal500(b *testing.B)   { criterionTransferOrdinal(500, b) }
func CriterionTransferOrdinal1000(b *testing.B)  { criterionTransferOrdinal(1000, b) }
func CriterionTransferOrdinal2000(b *testing.B)  { criterionTransferOrdinal(2000, b) }
func CriterionTransferOrdinal10000(b *testing.B) { criterionTransferOrdinal(10000, b) }
