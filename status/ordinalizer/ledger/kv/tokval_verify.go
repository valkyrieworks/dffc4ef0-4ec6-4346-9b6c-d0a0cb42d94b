package tokval_rend_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	db "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	ledgeridxkv "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger/kv"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyLedgerOrdinalizer(t *testing.T) {
	depot := db.FreshHeadingDatastore(db.FreshMemoryDatastore(), []byte("REDACTED"))
	ordinalizer := ledgeridxkv.New(depot)

	require.NoError(t, ordinalizer.Ordinal(kinds.IncidentDataFreshLedgerIncidents{
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
	}))

	for i := 2; i < 12; i++ {
		var ordinal bool
		if i%2 == 0 {
			ordinal = true
		}

		require.NoError(t, ordinalizer.Ordinal(kinds.IncidentDataFreshLedgerIncidents{
			Altitude: int64(i),
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
				{
					Kind: "REDACTED",
					Properties: []iface.IncidentProperty{
						{
							Key:   "REDACTED",
							Datum: fmt.Sprintf("REDACTED", i),
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
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{5},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{2, 4},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{4, 6, 8},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{2, 4, 6},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1, 10},
		},
	}

	for alias, tc := range verifyScenarios {

		t.Run(alias, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(context.Background(), tc.q)
			require.NoError(t, err)
			require.Equal(t, tc.outcomes, outcomes)
		})
	}
}

func VerifyLedgerOrdinalizerVaried(t *testing.T) {
	depot := db.FreshHeadingDatastore(db.FreshMemoryDatastore(), []byte("REDACTED"))
	ordinalizer := ledgeridxkv.New(depot)

	require.NoError(t, ordinalizer.Ordinal(kinds.IncidentDataFreshLedgerIncidents{
		Altitude: 1,
		Incidents: []iface.Incident{
			{},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
				},
			},
		},
	}))

	require.NoError(t, ordinalizer.Ordinal(kinds.IncidentDataFreshLedgerIncidents{
		Altitude: 2,
		Incidents: []iface.Incident{
			{},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
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
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1, 2},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{2},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{2},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1, 2},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{},
		},
	}

	for alias, tc := range verifyScenarios {

		t.Run(alias, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(context.Background(), tc.q)
			require.NoError(t, err)
			require.Equal(t, tc.outcomes, outcomes)
		})
	}
}

func VerifyLargeInteger(t *testing.T) {
	largeInteger := "REDACTED"
	largeDecimal := largeInteger + "REDACTED"
	largeDecimalLesser := largeInteger + "REDACTED"
	largeIntegerTinier := "REDACTED"
	depot := db.FreshHeadingDatastore(db.FreshMemoryDatastore(), []byte("REDACTED"))
	ordinalizer := ledgeridxkv.New(depot)

	require.NoError(t, ordinalizer.Ordinal(kinds.IncidentDataFreshLedgerIncidents{
		Altitude: 1,
		Incidents: []iface.Incident{
			{},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: largeDecimal,
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: largeDecimalLesser,
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: largeInteger,
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
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
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeInteger + "REDACTED"),
			outcomes: []int64{},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED"),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeDecimal),
			outcomes: []int64{1},
		},
		"REDACTED": {
			q:       inquire.ShouldAssemble("REDACTED" + largeIntegerTinier),
			outcomes: []int64{1},
		},
	}
	for alias, tc := range verifyScenarios {

		t.Run(alias, func(t *testing.T) {
			outcomes, err := ordinalizer.Lookup(context.Background(), tc.q)
			require.NoError(t, err)
			require.Equal(t, tc.outcomes, outcomes)
		})
	}
}
