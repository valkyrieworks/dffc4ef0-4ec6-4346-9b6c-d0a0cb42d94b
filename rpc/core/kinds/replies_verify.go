package basetypes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/p2p"
)

func VerifyStateOrdinaler(t *testing.T) {
	var state *OutcomeState
	assert.False(t, state.TransferOrdinalActivated())

	state = &OutcomeState{}
	assert.False(t, state.TransferOrdinalActivated())

	state.MemberDetails = p2p.StandardMemberDetails{}
	assert.False(t, state.TransferOrdinalActivated())

	scenarios := []struct {
		anticipated bool
		another    p2p.StandardMemberDetailsAnother
	}{
		{false, p2p.StandardMemberDetailsAnother{}},
		{false, p2p.StandardMemberDetailsAnother{TransOrdinal: "REDACTED"}},
		{false, p2p.StandardMemberDetailsAnother{TransOrdinal: "REDACTED"}},
		{true, p2p.StandardMemberDetailsAnother{TransOrdinal: "REDACTED"}},
	}

	for _, tc := range scenarios {
		state.MemberDetails.Another = tc.another
		assert.Equal(t, tc.anticipated, state.TransferOrdinalActivated())
	}
}
