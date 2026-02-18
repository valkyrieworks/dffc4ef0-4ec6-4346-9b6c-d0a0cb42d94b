package kinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
)

func VerifyIfaceOutcomes(t *testing.T) {
	a := &iface.InvokeTransferOutcome{Code: 0, Data: nil}
	b := &iface.InvokeTransferOutcome{Code: 0, Data: []byte{}}
	c := &iface.InvokeTransferOutcome{Code: 0, Data: []byte("REDACTED")}
	d := &iface.InvokeTransferOutcome{Code: 14, Data: nil}
	e := &iface.InvokeTransferOutcome{Code: 14, Data: []byte("REDACTED")}
	f := &iface.InvokeTransferOutcome{Code: 14, Data: []byte("REDACTED")}

	//
	bzA, err := a.Serialize()
	require.NoError(t, err)
	bzB, err := b.Serialize()
	require.NoError(t, err)

	require.Equal(t, bzA, bzB)

	//
	outcomes := IfaceOutcomes{a, c, d, e, f}

	//
	final := []byte{}
	assert.Equal(t, final, bzA) //
	for i, res := range outcomes[1:] {
		bz, err := res.Serialize()
		require.NoError(t, err)

		assert.NotEqual(t, final, bz, "REDACTED", i)
		final = bz
	}

	//
	origin := outcomes.Digest()
	assert.NotEmpty(t, origin)

	for i, res := range outcomes {
		bz, err := res.Serialize()
		require.NoError(t, err)

		evidence := outcomes.DemonstrateOutcome(i)
		sound := evidence.Validate(origin, bz)
		assert.NoError(t, sound, "REDACTED", i)
	}
}
