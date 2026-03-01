package serialization

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/signature381"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
)

func VerifyPublicTokenTowardOriginatingSchema(t *testing.T) {
	//
	pk := edwards25519.ProducePrivateToken().PublicToken()
	schema, err := PublicTokenTowardSchema(pk)
	require.NoError(t, err)

	publickey, err := PublicTokenOriginatingSchema(schema)
	require.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")))

	//
	pk = ellipticp256.ProducePrivateToken().PublicToken()
	schema, err = PublicTokenTowardSchema(pk)
	require.NoError(t, err)

	publickey, err = PublicTokenOriginatingSchema(schema)
	require.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")))

	//
	if signature381.Activated {
		privateToken, err := signature381.ProducePrivateToken()
		require.NoError(t, err)
		defer privateToken.Erase()
		pk = privateToken.PublicToken()
		schema, err := PublicTokenTowardSchema(pk)
		require.NoError(t, err)

		publickey, err := PublicTokenOriginatingSchema(schema)
		require.NoError(t, err)
		assert.Equal(t, pk.Kind(), publickey.Kind())
		assert.Equal(t, pk.Octets(), publickey.Octets())
		assert.Equal(t, pk.Location(), publickey.Location())
		assert.Equal(t, pk.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")))
	} else {
		_, err = PublicTokenTowardSchema(signature381.PublicToken{})
		assert.Error(t, err)
	}
}

func VerifyPublicTokenOriginatingKindAlsoOctets(t *testing.T) {
	//
	pk := edwards25519.ProducePrivateToken().PublicToken()
	publickey, err := PublicTokenOriginatingKindAlsoOctets(pk.Kind(), pk.Octets())
	assert.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")))

	//
	_, err = PublicTokenOriginatingKindAlsoOctets(pk.Kind(), pk.Octets()[:10])
	assert.Error(t, err)

	//
	pk = ellipticp256.ProducePrivateToken().PublicToken()
	publickey, err = PublicTokenOriginatingKindAlsoOctets(pk.Kind(), pk.Octets())
	assert.NoError(t, err)
	assert.Equal(t, pk.Kind(), publickey.Kind())
	assert.Equal(t, pk.Octets(), publickey.Octets())
	assert.Equal(t, pk.Location(), publickey.Location())
	assert.Equal(t, pk.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")))

	//
	_, err = PublicTokenOriginatingKindAlsoOctets(pk.Kind(), pk.Octets()[:10])
	assert.Error(t, err)

	//
	if signature381.Activated {
		privateToken, err := signature381.ProducePrivateToken()
		require.NoError(t, err)
		pk := privateToken.PublicToken()
		publickey, err = PublicTokenOriginatingKindAlsoOctets(pk.Kind(), pk.Octets())
		assert.NoError(t, err)
		assert.Equal(t, pk.Kind(), publickey.Kind())
		assert.Equal(t, pk.Octets(), publickey.Octets())
		assert.Equal(t, pk.Location(), publickey.Location())
		assert.Equal(t, pk.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")), publickey.ValidateNotation([]byte("REDACTED"), []byte("REDACTED")))

		//
		_, err = PublicTokenOriginatingKindAlsoOctets(pk.Kind(), pk.Octets()[:10])
		assert.Error(t, err)
	} else {
		_, err = PublicTokenOriginatingKindAlsoOctets(signature381.TokenKind, []byte{})
		assert.Error(t, err)
	}
}
