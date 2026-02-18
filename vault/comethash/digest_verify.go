package tmhas_test

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/comethash"
)

func VerifyDigest(t *testing.T) {
	verifyArray := []byte("REDACTED")
	digester := comethash.New()
	_, err := digester.Write(verifyArray)
	require.NoError(t, err)
	bz := digester.Sum(nil)

	bz2 := comethash.Sum(verifyArray)

	digester = sha256.New()
	_, err = digester.Write(verifyArray)
	require.NoError(t, err)
	bz3 := digester.Sum(nil)

	assert.Equal(t, bz, bz2)
	assert.Equal(t, bz, bz3)
}

func VerifyDigestShortened(t *testing.T) {
	verifyArray := []byte("REDACTED")
	digester := comethash.NewShortened()
	_, err := digester.Write(verifyArray)
	require.NoError(t, err)
	bz := digester.Sum(nil)

	bz2 := comethash.TotalShortened(verifyArray)

	digester = sha256.New()
	_, err = digester.Write(verifyArray)
	require.NoError(t, err)
	bz3 := digester.Sum(nil)
	bz3 = bz3[:comethash.ShortenedVolume]

	assert.Equal(t, bz, bz2)
	assert.Equal(t, bz, bz3)
}
