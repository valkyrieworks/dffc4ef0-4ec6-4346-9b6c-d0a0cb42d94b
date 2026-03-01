package ledger_test

import (
	"encoding/hex"
	"math"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/chainchronize"
	chainchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/chainchronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyBchainLedgerSolicitSignalCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias      string
		solicitAltitude int64
		anticipateFault     bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyAlias, func(t *testing.T) {
			solicit := chainchema.LedgerSolicit{Altitude: tc.solicitAltitude}
			assert.Equal(t, tc.anticipateFault, chainchronize.CertifySignal(&solicit) != nil, "REDACTED")
		})
	}
}

func VerifyBchainNegativeLedgerReplySignalCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias          string
		unReplyAltitude int64
		anticipateFault         bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyAlias, func(t *testing.T) {
			unReply := chainchema.NegativeLedgerReply{Altitude: tc.unReplyAltitude}
			assert.Equal(t, tc.anticipateFault, chainchronize.CertifySignal(&unReply) != nil, "REDACTED")
		})
	}
}

func VerifyBchainConditionSolicitSignalCertifyFundamental(t *testing.T) {
	solicit := chainchema.ConditionSolicit{}
	assert.NoError(t, chainchronize.CertifySignal(&solicit))
}

func VerifyBchainConditionReplySignalCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias       string
		replyAltitude int64
		anticipateFault      bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyAlias, func(t *testing.T) {
			reply := chainchema.ConditionReply{Altitude: tc.replyAltitude}
			assert.Equal(t, tc.anticipateFault, chainchronize.CertifySignal(&reply) != nil, "REDACTED")
		})
	}
}

//
func VerifyChainchronizeSignalArrays(t *testing.T) {
	ledger := kinds.CreateLedger(int64(3), []kinds.Tx{kinds.Tx("REDACTED")}, nil, nil)
	ledger.Edition.Ledger = 11 //

	bpb, err := ledger.TowardSchema()
	require.NoError(t, err)

	verifyScenarios := []struct {
		verifyAlias string
		bsignal     proto.Message
		expirationOctets string
	}{
		{"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Ledgerrequest{
			LedgerSolicit: &chainchema.LedgerSolicit{Altitude: 1},
		}}, "REDACTED"},
		{
			"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Ledgerrequest{
				LedgerSolicit: &chainchema.LedgerSolicit{Altitude: math.MaxInt64},
			}},
			"REDACTED",
		},
		{"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Ledgerreply{
			LedgerReply: &chainchema.LedgerReply{Ledger: bpb},
		}}, "REDACTED"},
		{"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Noledgerreply{
			NegativeLedgerReply: &chainchema.NegativeLedgerReply{Altitude: 1},
		}}, "REDACTED"},
		{
			"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Noledgerreply{
				NegativeLedgerReply: &chainchema.NegativeLedgerReply{Altitude: math.MaxInt64},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Conditionrequest{
				ConditionSolicit: &chainchema.ConditionSolicit{},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Conditionreply{
				ConditionReply: &chainchema.ConditionReply{Altitude: 1, Foundation: 2},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &chainchema.Signal{Sum: &chainchema.Signal_Conditionreply{
				ConditionReply: &chainchema.ConditionReply{Altitude: math.MaxInt64, Foundation: math.MaxInt64},
			}},
			"REDACTED",
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyAlias, func(t *testing.T) {
			bz, _ := proto.Marshal(tc.bsignal)

			require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz))
		})
	}
}
