package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
)

func VerifyLedgerchainDetails(t *testing.T) {
	scenarios := []struct {
		min, max     int64
		foundation, altitude int64
		threshold        int64
		outcomeMagnitude int64
		desireFault      bool
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
		scenarioText := fmt.Sprintf("REDACTED", i)
		min, max, err := refineMinimumMaximum(c.foundation, c.altitude, c.min, c.max, c.threshold)
		if c.desireFault {
			require.Error(t, err, scenarioText)
		} else {
			require.NoError(t, err, scenarioText)
			require.Equal(t, 1+max-min, c.outcomeMagnitude, scenarioText)
		}
	}
}

func VerifyLedgerOutcomes(t *testing.T) {
	outcomes := &iface.ReplyCulminateLedger{
		TransferOutcomes: []*iface.InvokeTransferOutcome{
			{Cipher: 0, Data: []byte{0x01}, Log: "REDACTED"},
			{Cipher: 0, Data: []byte{0x02}, Log: "REDACTED"},
			{Cipher: 1, Log: "REDACTED"},
		},
		PlatformDigest: make([]byte, 1),
	}

	env := &Context{}
	env.StatusDepot = sm.FreshDepot(dbm.FreshMemoryDatastore(), sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	err := env.StatusDepot.PersistCulminateLedgerReply(100, outcomes)
	require.NoError(t, err)
	teststore := &simulations.LedgerDepot{}
	teststore.On("REDACTED").Return(int64(100))
	teststore.On("REDACTED").Return(int64(1))
	env.LedgerDepot = teststore

	verifyScenarios := []struct {
		altitude  int64
		desireFault bool
		desireResult *ktypes.OutcomeLedgerOutcomes
	}{
		{-1, true, nil},
		{0, true, nil},
		{101, true, nil},
		{100, false, &ktypes.OutcomeLedgerOutcomes{
			Altitude:                100,
			TransOutcomes:            outcomes.TransferOutcomes,
			CulminateLedgerIncidents:   outcomes.Incidents,
			AssessorRevisions:      outcomes.AssessorRevisions,
			AgreementArgumentRevisions: outcomes.AgreementArgumentRevisions,
			PlatformDigest:               make([]byte, 1),
		}},
	}

	for _, tc := range verifyScenarios {
		res, err := env.LedgerOutcomes(&remoteifacetypes.Env{}, &tc.altitude)
		if tc.desireFault {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.desireResult, res)
		}
	}
}
