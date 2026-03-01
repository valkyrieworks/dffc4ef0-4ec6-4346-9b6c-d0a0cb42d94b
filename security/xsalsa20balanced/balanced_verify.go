package xsalsa20balanced

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang.org/x/crypto/bcrypt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
)

func VerifyPlain(t *testing.T) {
	cleartext := []byte("REDACTED")
	credential := []byte("REDACTED")
	sealedtext := SealBalanced(cleartext, credential)
	cleartext2, err := UnsealBalanced(sealedtext, credential)

	require.Nil(t, err, "REDACTED", err)
	assert.Equal(t, cleartext, cleartext2)
}

func VerifyPlainUsingDerivation(t *testing.T) {
	cleartext := []byte("REDACTED")
	credentialPhrase := []byte("REDACTED")
	credential, err := bcrypt.GenerateFromPassword(credentialPhrase, 12)
	if err != nil {
		t.Error(err)
	}
	credential = security.Hash256(credential)

	sealedtext := SealBalanced(cleartext, credential)
	cleartext2, err := UnsealBalanced(sealedtext, credential)

	require.Nil(t, err, "REDACTED", err)
	assert.Equal(t, cleartext, cleartext2)
}
