package crypt_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
)

//
//
func VerifyArbitraryCoherence(t *testing.T) {
	x1 := vault.CRandomOctets(256)
	x2 := vault.CRandomOctets(256)
	x3 := vault.CRandomOctets(256)
	x4 := vault.CRandomOctets(256)
	x5 := vault.CRandomOctets(256)
	require.NotEqual(t, x1, x2)
	require.NotEqual(t, x3, x4)
	require.NotEqual(t, x4, x5)
	require.NotEqual(t, x1, x5)
}
