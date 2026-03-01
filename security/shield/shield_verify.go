package shield

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyShield(t *testing.T) {
	ledgerKind := "REDACTED"
	data := []byte("REDACTED")
	shieldTxt := SerializeShield(ledgerKind, nil, data)

	//
	ledgerKind2, _, datum2, err := DeserializeShield(shieldTxt)
	require.Nil(t, err, "REDACTED", err)
	assert.Equal(t, ledgerKind, ledgerKind2)
	assert.Equal(t, data, datum2)
}
