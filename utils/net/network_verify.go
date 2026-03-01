package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifySchemeAlsoLocator(t *testing.T) {
	scenarios := []struct {
		completeLocation string
		schema    string
		location     string
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
		schema, location := SchemeAlsoLocation(c.completeLocation)
		assert.Equal(t, schema, c.schema)
		assert.Equal(t, location, c.location)
	}
}
