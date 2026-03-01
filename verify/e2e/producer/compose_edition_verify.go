package primary

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyEditionLocator(t *testing.T) {
	verifyScenarios := []struct {
		foundationEdtn        string
		labels           []string
		anticipatedNewest string
	}{
		{
			foundationEdtn:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			foundationEdtn:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			foundationEdtn:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			foundationEdtn:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
		{
			foundationEdtn:        "REDACTED",
			labels:           []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
			anticipatedNewest: "REDACTED",
		},
	}
	for _, tc := range verifyScenarios {
		veritableNewest, err := locateNewestDeliveryLabel(tc.foundationEdtn, tc.labels)
		require.NoError(t, err)
		assert.Equal(t, tc.anticipatedNewest, veritableNewest)
	}
}
