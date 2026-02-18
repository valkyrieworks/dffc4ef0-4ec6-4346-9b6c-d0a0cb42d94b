package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyReleaseLocator(t *testing.T) {
	verifyScenarios := []struct {
		rootRev        string
		labels           []string
		anticipatedNewest string
	}{
		{
			rootRev:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			rootRev:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			rootRev:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			rootRev:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			rootRev:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
	}
	for _, tc := range verifyScenarios {
		factualNewest, err := locateNewestDeliveryLabel(tc.rootRev, tc.labels)
		require.NoError(t, err)
		assert.Equal(t, tc.anticipatedNewest, factualNewest)
	}
}
