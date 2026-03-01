package edwards255_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
)

func VerifyAttestAlsoCertifyEdwards25519(t *testing.T) {
	privateToken := edwards25519.ProducePrivateToken()
	publicToken := privateToken.PublicToken()

	msg := security.CHARArbitraryOctets(128)
	sig, err := privateToken.Attest(msg)
	require.Nil(t, err)

	//
	assert.True(t, publicToken.ValidateNotation(msg, sig))

	//
	//
	sig[7] ^= byte(0x01)

	assert.False(t, publicToken.ValidateNotation(msg, sig))
}

func VerifyClusterSecure(t *testing.T) {
	v := edwards25519.FreshClusterValidator()

	for i := 0; i <= 38; i++ {
		private := edwards25519.ProducePrivateToken()
		pub := private.PublicToken()

		var msg []byte
		if i%2 == 0 {
			msg = []byte("REDACTED")
		} else {
			msg = []byte("REDACTED")
		}

		sig, err := private.Attest(msg)
		require.NoError(t, err)

		err = v.Add(pub, msg, sig)
		require.NoError(t, err)
	}

	ok, _ := v.Validate()
	require.True(t, ok)
}
