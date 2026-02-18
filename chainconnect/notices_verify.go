package ledger_test

import (
	"encoding/hex"
	"math"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/chainconnect"
	chainproto "github.com/valkyrieworks/schema/consensuscore/chainconnect"
	"github.com/valkyrieworks/kinds"
)

func VerifyBcodeLedgerQuerySignalCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel      string
		queryLevel int64
		anticipateErr     bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyLabel, func(t *testing.T) {
			query := chainproto.LedgerQuery{Level: tc.queryLevel}
			assert.Equal(t, tc.anticipateErr, chainconnect.CertifyMessage(&query) != nil, "REDACTED")
		})
	}
}

func VerifyBcodeNoLedgerReplySignalCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel          string
		notReplyLevel int64
		anticipateErr         bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyLabel, func(t *testing.T) {
			notReply := chainproto.NoLedgerReply{Level: tc.notReplyLevel}
			assert.Equal(t, tc.anticipateErr, chainconnect.CertifyMessage(&notReply) != nil, "REDACTED")
		})
	}
}

func VerifyBcodeStatusQuerySignalCertifySimple(t *testing.T) {
	query := chainproto.StatusQuery{}
	assert.NoError(t, chainconnect.CertifyMessage(&query))
}

func VerifyBcodeStatusReplySignalCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel       string
		replyLevel int64
		anticipateErr      bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyLabel, func(t *testing.T) {
			reply := chainproto.StatusReply{Level: tc.replyLevel}
			assert.Equal(t, tc.anticipateErr, chainconnect.CertifyMessage(&reply) != nil, "REDACTED")
		})
	}
}

//
func VerifyChainconnectSignalArrays(t *testing.T) {
	ledger := kinds.CreateLedger(int64(3), []kinds.Tx{kinds.Tx("REDACTED")}, nil, nil)
	ledger.Release.Ledger = 11 //

	bpb, err := ledger.ToSchema()
	require.NoError(t, err)

	verifyScenarios := []struct {
		verifyLabel string
		bmessage     proto.Message
		expirationOctets string
	}{
		{"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Chainrequest{
			LedgerQuery: &chainproto.LedgerQuery{Level: 1},
		}}, "REDACTED"},
		{
			"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Chainrequest{
				LedgerQuery: &chainproto.LedgerQuery{Level: math.MaxInt64},
			}},
			"REDACTED",
		},
		{"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Chainreply{
			LedgerReply: &chainproto.LedgerReply{Ledger: bpb},
		}}, "REDACTED"},
		{"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Nonledgerreply{
			NoLedgerReply: &chainproto.NoLedgerReply{Level: 1},
		}}, "REDACTED"},
		{
			"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Nonledgerreply{
				NoLedgerReply: &chainproto.NoLedgerReply{Level: math.MaxInt64},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Staterequest{
				StatusQuery: &chainproto.StatusQuery{},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Statereply{
				StatusReply: &chainproto.StatusReply{Level: 1, Root: 2},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &chainproto.Signal{Sum: &chainproto.Signal_Statereply{
				StatusReply: &chainproto.StatusReply{Level: math.MaxInt64, Root: math.MaxInt64},
			}},
			"REDACTED",
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.verifyLabel, func(t *testing.T) {
			bz, _ := proto.Marshal(tc.bmessage)

			require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz))
		})
	}
}
