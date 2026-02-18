package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/bls12381"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/secp256k1"
)

func VerifyPublicKeyToFromSchema(t *testing.T) {
	//
	pk := ed25519.GeneratePrivateKey().PublicKey()
	schema, err := PublicKeyToSchema(pk)
	require.NoError(t, err)

	publickey, err := PublicKeyFromSchema(schema)
	require.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")))

	//
	pk = secp256k1.GeneratePrivateKey().PublicKey()
	schema, err = PublicKeyToSchema(pk)
	require.NoError(t, err)

	publickey, err = PublicKeyFromSchema(schema)
	require.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")))

	//
	if bls12381.Activated {
		privateKey, err := bls12381.GeneratePrivateKey()
		require.NoError(t, err)
		defer privateKey.Nullify()
		pk = privateKey.PublicKey()
		schema, err := PublicKeyToSchema(pk)
		require.NoError(t, err)

		publickey, err := PublicKeyFromSchema(schema)
		require.NoError(t, err)
		assert.Equal(t, pk.Kind(), publickey.Kind())
		assert.Equal(t, pk.Octets(), publickey.Octets())
		assert.Equal(t, pk.Location(), publickey.Location())
		assert.Equal(t, pk.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")))
	} else {
		_, err = PublicKeyToSchema(bls12381.PublicKey{})
		assert.Error(t, err)
	}
}

func VerifyPublicKeyFromKindAndOctets(t *testing.T) {
	//
	pk := ed25519.GeneratePrivateKey().PublicKey()
	publickey, err := PublicKeyFromKindAndOctets(pk.Kind(), pk.Octets())
	assert.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")))

	//
	_, err = PublicKeyFromKindAndOctets(pk.Kind(), pk.Octets()[:10])
	assert.Error(t, err)

	//
	pk = secp256k1.GeneratePrivateKey().PublicKey()
	publickey, err = PublicKeyFromKindAndOctets(pk.Kind(), pk.Octets())
	assert.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")))

	//
	_, err = PublicKeyFromKindAndOctets(pk.Kind(), pk.Octets()[:10])
	assert.Error(t, err)

	//
	if bls12381.Activated {
		privateKey, err := bls12381.GeneratePrivateKey()
		require.NoError(t, err)
		pk := privateKey.PublicKey()
		publickey, err = PublicKeyFromKindAndOctets(pk.Kind(), pk.Octets())
		assert.NoError(t, err)
		assert.Equal(t, pk.Kind(), publickey.Kind())
		assert.Equal(t, pk.Octets(), publickey.Octets())
		assert.Equal(t, pk.Location(), publickey.Location())
		assert.Equal(t, pk.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateAutograph([]byte("REDACTED"), []byte("REDACTED")))

		//
		_, err = PublicKeyFromKindAndOctets(pk.Kind(), pk.Octets()[:10])
		assert.Error(t, err)
	} else {
		_, err = PublicKeyFromKindAndOctets(bls12381.KeyKind, []byte{})
		assert.Error(t, err)
	}
}
