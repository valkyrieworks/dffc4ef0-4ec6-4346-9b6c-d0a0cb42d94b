package xsalsa20balanced

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang.org/x/crypto/bcrypt"

	"github.com/valkyrieworks/vault"
)

func VerifyBasic(t *testing.T) {
	cleartext := []byte("REDACTED")
	key := []byte("REDACTED")
	cyphertext := EncodeBalanced(cleartext, key)
	cleartext2, err := DecodeBalanced(cyphertext, key)

	require.Nil(t, err, "REDACTED", err)
	assert.Equal(t, cleartext, cleartext2)
}

func VerifyBasicWithKDF(t *testing.T) {
	cleartext := []byte("REDACTED")
	privatePass := []byte("REDACTED")
	key, err := bcrypt.GenerateFromPassword(privatePass, 12)
	if err != nil {
		t.Error(err)
	}
	key = vault.Sha256(key)

	cyphertext := EncodeBalanced(cleartext, key)
	cleartext2, err := DecodeBalanced(cyphertext, key)

	require.Nil(t, err, "REDACTED", err)
	assert.Equal(t, cleartext, cleartext2)
}
