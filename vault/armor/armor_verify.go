package armor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyArmor(t *testing.T) {
	ledgerKind := "REDACTED"
	data := []byte("REDACTED")
	armorStr := MarshalArmor(ledgerKind, nil, data)

	//
	ledgerKind2, _, data2, err := ParseArmor(armorStr)
	require.Nil(t, err, "REDACTED", err)
	assert.Equal(t, ledgerKind, ledgerKind2)
	assert.Equal(t, data, data2)
}
