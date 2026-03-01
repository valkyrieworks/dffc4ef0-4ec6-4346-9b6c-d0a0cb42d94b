package simulate__test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/simulate"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
)

func VerifyCondition(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	m := &simulate.ConditionSimulate{
		Invocation: simulate.Invocation{
			Reply: &ktypes.OutcomeCondition{
				ChronizeDetails: ktypes.ChronizeDetails{
					NewestLedgerDigest:   octets.HexadecimalOctets("REDACTED"),
					NewestApplicationDigest:     octets.HexadecimalOctets("REDACTED"),
					NewestLedgerAltitude: 10,
				},
			},
		},
	}

	r := simulate.FreshConditionScribe(m)
	require.Equal(0, len(r.Invocations))

	//
	condition, err := r.Condition(context.Background())
	require.Nil(err, "REDACTED", err)
	assert.EqualValues("REDACTED", condition.ChronizeDetails.NewestLedgerDigest)
	assert.EqualValues(10, condition.ChronizeDetails.NewestLedgerAltitude)

	//
	require.Equal(1, len(r.Invocations))
	rs := r.Invocations[0]
	assert.Equal("REDACTED", rs.Alias)
	assert.Nil(rs.Arguments)
	assert.Nil(rs.Failure)
	require.NotNil(rs.Reply)
	st, ok := rs.Reply.(*ktypes.OutcomeCondition)
	require.True(ok)
	assert.EqualValues("REDACTED", st.ChronizeDetails.NewestLedgerDigest)
	assert.EqualValues(10, st.ChronizeDetails.NewestLedgerAltitude)
}
