package ed255_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
)

func VerifyAttestAndCertifyEd25519(t *testing.T) {
	privateKey := ed25519.GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := vault.CRandomOctets(128)
	sig, err := privateKey.Attest(msg)
	require.Nil(t, err)

	//
	assert.True(t, publicKey.ValidateAutograph(msg, sig))

	//
	//
	sig[7] ^= byte(0x01)

	assert.False(t, publicKey.ValidateAutograph(msg, sig))
}

func VerifyGroupSecure(t *testing.T) {
	v := ed25519.NewGroupValidator()

	for i := 0; i <= 38; i++ {
		private := ed25519.GeneratePrivateKey()
		pub := private.PublicKey()

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
