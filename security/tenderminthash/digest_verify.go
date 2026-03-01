package tendermintdigest_test

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
)

func VerifyDigest(t *testing.T) {
	verifyArray := []byte("REDACTED")
	digester := tenderminthash.New()
	_, err := digester.Write(verifyArray)
	require.NoError(t, err)
	bz := digester.Sum(nil)

	bz2 := tenderminthash.Sum(verifyArray)

	digester = sha256.New()
	_, err = digester.Write(verifyArray)
	require.NoError(t, err)
	bz3 := digester.Sum(nil)

	assert.Equal(t, bz, bz2)
	assert.Equal(t, bz, bz3)
}

func VerifyDigestAbridged(t *testing.T) {
	verifyArray := []byte("REDACTED")
	digester := tenderminthash.FreshAbridged()
	_, err := digester.Write(verifyArray)
	require.NoError(t, err)
	bz := digester.Sum(nil)

	bz2 := tenderminthash.TotalAbridged(verifyArray)

	digester = sha256.New()
	_, err = digester.Write(verifyArray)
	require.NoError(t, err)
	bz3 := digester.Sum(nil)
	bz3 = bz3[:tenderminthash.AbridgedExtent]

	assert.Equal(t, bz, bz2)
	assert.Equal(t, bz, bz3)
}
