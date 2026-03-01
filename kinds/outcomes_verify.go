package kinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

func VerifyIfaceOutcomes(t *testing.T) {
	a := &iface.InvokeTransferOutcome{Cipher: 0, Data: nil}
	b := &iface.InvokeTransferOutcome{Cipher: 0, Data: []byte{}}
	c := &iface.InvokeTransferOutcome{Cipher: 0, Data: []byte("REDACTED")}
	d := &iface.InvokeTransferOutcome{Cipher: 14, Data: nil}
	e := &iface.InvokeTransferOutcome{Cipher: 14, Data: []byte("REDACTED")}
	f := &iface.InvokeTransferOutcome{Cipher: 14, Data: []byte("REDACTED")}

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

		attestation := outcomes.AscertainOutcome(i)
		sound := attestation.Validate(origin, bz)
		assert.NoError(t, sound, "REDACTED", i)
	}
}
