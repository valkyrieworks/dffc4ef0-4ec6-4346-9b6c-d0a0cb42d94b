package agent_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/rpc/customer"
	"github.com/valkyrieworks/rpc/customer/emulate"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
)

func VerifyWaitForLevel(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	m := &emulate.StateEmulate{
		Invoke: emulate.Invoke{
			Fault: errors.New("REDACTED"),
		},
	}
	r := emulate.NewStateTracer(m)

	//
	err := customer.WaitForLevel(r, 8, nil)
	require.NotNil(err)
	require.Equal("REDACTED", err.Error())
	//
	require.Equal(1, len(r.Invocations))

	//
	m.Invoke = emulate.Invoke{
		Reply: &ctypes.OutcomeState{AlignDetails: ctypes.AlignDetails{NewestLedgerLevel: 10}},
	}

	//
	err = customer.WaitForLevel(r, 40, nil)
	require.NotNil(err)
	require.True(strings.Contains(err.Error(), "REDACTED"))
	//
	require.Equal(2, len(r.Invocations))

	//
	err = customer.WaitForLevel(r, 5, nil)
	require.Nil(err)
	//
	require.Equal(3, len(r.Invocations))

	//
	//
	mineObserver := func(variance int64) error {
		//
		m.Reply = &ctypes.OutcomeState{AlignDetails: ctypes.AlignDetails{NewestLedgerLevel: 15}}
		return customer.StandardWaitTactic(variance)
	}

	//
	err = customer.WaitForLevel(r, 12, mineObserver)
	require.Nil(err)
	//
	require.Equal(5, len(r.Invocations))

	pre := r.Invocations[3]
	require.Nil(pre.Fault)
	pout, ok := pre.Reply.(*ctypes.OutcomeState)
	require.True(ok)
	assert.Equal(int64(10), pout.AlignDetails.NewestLedgerLevel)

	submit := r.Invocations[4]
	require.Nil(submit.Fault)
	submitter, ok := submit.Reply.(*ctypes.OutcomeState)
	require.True(ok)
	assert.Equal(int64(15), submitter.AlignDetails.NewestLedgerLevel)
}
