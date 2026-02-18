package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyProtocolAndLocation(t *testing.T) {
	scenarios := []struct {
		completeAddress string
		schema    string
		address     string
	}{
		{
			"REDACTED",
			"REDACTED",
			"REDACTED",
		},
		{
			"REDACTED",
			"REDACTED",
			"REDACTED",
		},
		{
			"REDACTED",
			"REDACTED",
			"REDACTED",
		},
	}

	for _, c := range scenarios {
		schema, address := ProtocolAndLocation(c.completeAddress)
		assert.Equal(t, schema, c.schema)
		assert.Equal(t, address, c.address)
	}
}
