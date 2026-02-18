package emulate__test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/rpc/customer/emulate"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
)

func VerifyState(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	m := &emulate.StateEmulate{
		Invoke: emulate.Invoke{
			Reply: &ctypes.OutcomeState{
				AlignDetails: ctypes.AlignDetails{
					NewestLedgerDigest:   octets.HexOctets("REDACTED"),
					NewestApplicationDigest:     octets.HexOctets("REDACTED"),
					NewestLedgerLevel: 10,
				},
			},
		},
	}

	r := emulate.NewStateTracer(m)
	require.Equal(0, len(r.Invocations))

	//
	state, err := r.Status(context.Background())
	require.Nil(err, "REDACTED", err)
	assert.EqualValues("REDACTED", state.AlignDetails.NewestLedgerDigest)
	assert.EqualValues(10, state.AlignDetails.NewestLedgerLevel)

	//
	require.Equal(1, len(r.Invocations))
	rs := r.Invocations[0]
	assert.Equal("REDACTED", rs.Label)
	assert.Nil(rs.Args)
	assert.Nil(rs.Fault)
	require.NotNil(rs.Reply)
	st, ok := rs.Reply.(*ctypes.OutcomeState)
	require.True(ok)
	assert.EqualValues("REDACTED", st.AlignDetails.NewestLedgerDigest)
	assert.EqualValues(10, st.AlignDetails.NewestLedgerLevel)
}
