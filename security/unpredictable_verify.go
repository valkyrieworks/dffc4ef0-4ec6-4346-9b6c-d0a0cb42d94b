package crypto_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
)

//
//
func VerifyUnpredictableCoherence(t *testing.T) {
	x1 := security.CHARArbitraryOctets(256)
	x2 := security.CHARArbitraryOctets(256)
	x3 := security.CHARArbitraryOctets(256)
	x4 := security.CHARArbitraryOctets(256)
	x5 := security.CHARArbitraryOctets(256)
	require.NotEqual(t, x1, x2)
	require.NotEqual(t, x3, x4)
	require.NotEqual(t, x4, x5)
	require.NotEqual(t, x1, x5)
}
