package basetypes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

func VerifyConditionOrdinalizer(t *testing.T) {
	var condition *OutcomeCondition
	assert.False(t, condition.TransferPositionActivated())

	condition = &OutcomeCondition{}
	assert.False(t, condition.TransferPositionActivated())

	condition.PeerDetails = p2p.FallbackPeerDetails{}
	assert.False(t, condition.TransferPositionActivated())

	scenarios := []struct {
		anticipated bool
		another    p2p.FallbackPeerDetailsAnother
	}{
		{false, p2p.FallbackPeerDetailsAnother{}},
		{false, p2p.FallbackPeerDetailsAnother{TransferOrdinal: "REDACTED"}},
		{false, p2p.FallbackPeerDetailsAnother{TransferOrdinal: "REDACTED"}},
		{true, p2p.FallbackPeerDetailsAnother{TransferOrdinal: "REDACTED"}},
	}

	for _, tc := range scenarios {
		condition.PeerDetails.Another = tc.another
		assert.Equal(t, tc.anticipated, condition.TransferPositionActivated())
	}
}
