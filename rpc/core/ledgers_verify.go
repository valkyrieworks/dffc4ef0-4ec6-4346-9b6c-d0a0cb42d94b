package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/simulations"
)

func VerifyLedgerchainDetails(t *testing.T) {
	scenarios := []struct {
		min, max     int64
		root, level int64
		ceiling        int64
		outcomeExtent int64
		desireErr      bool
	}{
		//
		{0, 0, 0, 0, 10, 0, true},  //
		{0, 1, 0, 0, 10, 0, true},  //
		{0, 0, 0, 1, 10, 1, false}, //
		{2, 0, 0, 1, 10, 0, true},  //
		{2, 1, 0, 5, 10, 0, true},

		//
		{1, 10, 0, 14, 10, 10, false}, //
		{-1, 10, 0, 14, 10, 0, true},
		{1, -10, 0, 14, 10, 0, true},
		{-9223372036854775808, -9223372036854775788, 0, 100, 20, 0, true},

		//
		{1, 1, 1, 1, 1, 1, false},
		{2, 5, 3, 5, 5, 3, false},

		//
		{1, 1, 0, 1, 10, 1, false},
		{1, 1, 0, 5, 10, 1, false},
		{2, 2, 0, 5, 10, 1, false},
		{1, 2, 0, 5, 10, 2, false},
		{1, 5, 0, 1, 10, 1, false},
		{1, 5, 0, 10, 10, 5, false},
		{1, 15, 0, 10, 10, 10, false},
		{1, 15, 0, 15, 10, 10, false},
		{1, 15, 0, 15, 20, 15, false},
		{1, 20, 0, 15, 20, 15, false},
		{1, 20, 0, 20, 20, 20, false},
	}

	for i, c := range scenarios {
		scenarioString := fmt.Sprintf("REDACTED", i)
		min, max, err := refineMinimumMaximum(c.root, c.level, c.min, c.max, c.ceiling)
		if c.desireErr {
			require.Error(t, err, scenarioString)
		} else {
			require.NoError(t, err, scenarioString)
			require.Equal(t, 1+max-min, c.outcomeExtent, scenarioString)
		}
	}
}

func VerifyLedgerOutcomes(t *testing.T) {
	outcomes := &iface.ReplyCompleteLedger{
		TransOutcomes: []*iface.InvokeTransferOutcome{
			{Code: 0, Data: []byte{0x01}, Log: "REDACTED"},
			{Code: 0, Data: []byte{0x02}, Log: "REDACTED"},
			{Code: 1, Log: "REDACTED"},
		},
		ApplicationDigest: make([]byte, 1),
	}

	env := &Context{}
	env.StatusDepot = sm.NewDepot(dbm.NewMemoryStore(), sm.DepotSettings{
		DropIfaceReplies: false,
	})
	err := env.StatusDepot.PersistCompleteLedgerReply(100, outcomes)
	require.NoError(t, err)
	stubstore := &simulations.LedgerDepot{}
	stubstore.On("REDACTED").Return(int64(100))
	stubstore.On("REDACTED").Return(int64(1))
	env.LedgerDepot = stubstore

	verifyScenarios := []struct {
		level  int64
		desireErr bool
		desireOutput *ctypes.OutcomeLedgerOutcomes
	}{
		{-1, true, nil},
		{0, true, nil},
		{101, true, nil},
		{100, false, &ctypes.OutcomeLedgerOutcomes{
			Level:                100,
			TransOutcomes:            outcomes.TransOutcomes,
			CompleteLedgerEvents:   outcomes.Events,
			RatifierRefreshes:      outcomes.RatifierRefreshes,
			AgreementArgumentRefreshes: outcomes.AgreementArgumentRefreshes,
			ApplicationDigest:               make([]byte, 1),
		}},
	}

	for _, tc := range verifyScenarios {
		res, err := env.LedgerOutcomes(&rpctypes.Context{}, &tc.level)
		if tc.desireErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.desireOutput, res)
		}
	}
}
