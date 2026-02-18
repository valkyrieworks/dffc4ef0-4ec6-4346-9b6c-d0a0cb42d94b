package object_te_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	db "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	ledgerordinalkv "github.com/valkyrieworks/status/ordinaler/ledger/kv"
	"github.com/valkyrieworks/kinds"
)

func VerifyLedgerOrdinaler(t *testing.T) {
	depot := db.NewHeadingStore(db.NewMemoryStore(), []byte("REDACTED"))
	ordinaler := ledgerordinalkv.New(depot)

	require.NoError(t, ordinaler.Ordinal(kinds.EventDataNewLedgerEvents{
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
	}))

	for i := 2; i < 12; i++ {
		var ordinal bool
		if i%2 == 0 {
			ordinal = true
		}

		require.NoError(t, ordinaler.Ordinal(kinds.EventDataNewLedgerEvents{
			Level: int64(i),
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
				{
					Kind: "REDACTED",
					Properties: []iface.EventProperty{
						{
							Key:   "REDACTED",
							Item: fmt.Sprintf("REDACTED", i),
							Ordinal: ordinal,
						},
					},
				},
			},
		}))
	}

	verifyScenarios := map[string]struct {
		q       *inquire.Inquire
		outcomes []int64
	}{
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{5},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{2, 4},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{4, 6, 8},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{2, 4, 6},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1, 10},
		},
	}

	for label, tc := range verifyScenarios {

		t.Run(label, func(t *testing.T) {
			outcomes, err := ordinaler.Scan(context.Background(), tc.q)
			require.NoError(t, err)
			require.Equal(t, tc.outcomes, outcomes)
		})
	}
}

func VerifyLedgerOrdinalerMultiple(t *testing.T) {
	depot := db.NewHeadingStore(db.NewMemoryStore(), []byte("REDACTED"))
	ordinaler := ledgerordinalkv.New(depot)

	require.NoError(t, ordinaler.Ordinal(kinds.EventDataNewLedgerEvents{
		Level: 1,
		Events: []iface.Event{
			{},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
		},
	}))

	require.NoError(t, ordinaler.Ordinal(kinds.EventDataNewLedgerEvents{
		Level: 2,
		Events: []iface.Event{
			{},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
		},
	}))

	verifyScenarios := map[string]struct {
		q       *inquire.Inquire
		outcomes []int64
	}{
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1, 2},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{2},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{2},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1, 2},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{},
		},
	}

	for label, tc := range verifyScenarios {

		t.Run(label, func(t *testing.T) {
			outcomes, err := ordinaler.Scan(context.Background(), tc.q)
			require.NoError(t, err)
			require.Equal(t, tc.outcomes, outcomes)
		})
	}
}

func VerifyLargeInteger(t *testing.T) {
	largeInteger := "REDACTED"
	largeFloat := largeInteger + "REDACTED"
	largeFloatLesser := largeInteger + "REDACTED"
	largeIntegerLess := "REDACTED"
	depot := db.NewHeadingStore(db.NewMemoryStore(), []byte("REDACTED"))
	ordinaler := ledgerordinalkv.New(depot)

	require.NoError(t, ordinaler.Ordinal(kinds.EventDataNewLedgerEvents{
		Level: 1,
		Events: []iface.Event{
			{},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: largeFloat,
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: largeFloatLesser,
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: largeInteger,
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
		},
	},
	))

	verifyScenarios := map[string]struct {
		q       *inquire.Inquire
		outcomes []int64
	}{
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeFloat),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldBuild("REDACTED" + largeIntegerLess),
			outcomes: []int64{1},
		},
	}
	for label, tc := range verifyScenarios {

		t.Run(label, func(t *testing.T) {
			outcomes, err := ordinaler.Scan(context.Background(), tc.q)
			require.NoError(t, err)
			require.Equal(t, tc.outcomes, outcomes)
		})
	}
}
