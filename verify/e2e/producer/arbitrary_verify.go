package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyGroupings(t *testing.T) {
	influx := map[string][]any{
		"REDACTED":   {false, true},
		"REDACTED":    {1, 2, 3},
		"REDACTED": {"REDACTED", "REDACTED"},
	}

	c := groupings(influx)
	assert.Equal(t, []map[string]any{
		{"REDACTED": false, "REDACTED": 1, "REDACTED": "REDACTED"},
		{"REDACTED": false, "REDACTED": 1, "REDACTED": "REDACTED"},
		{"REDACTED": false, "REDACTED": 2, "REDACTED": "REDACTED"},
		{"REDACTED": false, "REDACTED": 2, "REDACTED": "REDACTED"},
		{"REDACTED": false, "REDACTED": 3, "REDACTED": "REDACTED"},
		{"REDACTED": false, "REDACTED": 3, "REDACTED": "REDACTED"},
		{"REDACTED": true, "REDACTED": 1, "REDACTED": "REDACTED"},
		{"REDACTED": true, "REDACTED": 1, "REDACTED": "REDACTED"},
		{"REDACTED": true, "REDACTED": 2, "REDACTED": "REDACTED"},
		{"REDACTED": true, "REDACTED": 2, "REDACTED": "REDACTED"},
		{"REDACTED": true, "REDACTED": 3, "REDACTED": "REDACTED"},
		{"REDACTED": true, "REDACTED": 3, "REDACTED": "REDACTED"},
	}, c)
}
