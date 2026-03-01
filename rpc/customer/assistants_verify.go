package cust_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/simulate"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
)

func VerifyPauseForeachAltitude(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	m := &simulate.ConditionSimulate{
		Invocation: simulate.Invocation{
			Failure: errors.New("REDACTED"),
		},
	}
	r := simulate.FreshConditionTracer(m)

	//
	err := customer.PauseForeachAltitude(r, 8, nil)
	require.NotNil(err)
	require.Equal("REDACTED", err.Error())
	//
	require.Equal(1, len(r.Invocations))

	//
	m.Invocation = simulate.Invocation{
		Reply: &ktypes.OutcomeCondition{ChronizeDetails: ktypes.ChronizeDetails{NewestLedgerAltitude: 10}},
	}

	//
	err = customer.PauseForeachAltitude(r, 40, nil)
	require.NotNil(err)
	require.True(strings.Contains(err.Error(), "REDACTED"))
	//
	require.Equal(2, len(r.Invocations))

	//
	err = customer.PauseForeachAltitude(r, 5, nil)
	require.Nil(err)
	//
	require.Equal(3, len(r.Invocations))

	//
	//
	selfPauser := func(variation int64) error {
		//
		m.Reply = &ktypes.OutcomeCondition{ChronizeDetails: ktypes.ChronizeDetails{NewestLedgerAltitude: 15}}
		return customer.FallbackPauseTactic(variation)
	}

	//
	err = customer.PauseForeachAltitude(r, 12, selfPauser)
	require.Nil(err)
	//
	require.Equal(5, len(r.Invocations))

	pre := r.Invocations[3]
	require.Nil(pre.Failure)
	precheckr, ok := pre.Reply.(*ktypes.OutcomeCondition)
	require.True(ok)
	assert.Equal(int64(10), precheckr.ChronizeDetails.NewestLedgerAltitude)

	submit := r.Invocations[4]
	require.Nil(submit.Failure)
	submitr, ok := submit.Reply.(*ktypes.OutcomeCondition)
	require.True(ok)
	assert.Equal(int64(15), submitr.ChronizeDetails.NewestLedgerAltitude)
}
