package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
)

func VerifyPeerDetailsCertify(t *testing.T) {
	//
	ni := FallbackPeerDetails{}
	assert.Error(t, ni.Certify())

	conduits := make([]byte, maximumCountConduits)
	for i := 0; i < maximumCountConduits; i++ {
		conduits[i] = byte(i)
	}
	cloneConduits := make([]byte, 5)
	copy(cloneConduits, conduits[:5])
	cloneConduits = append(cloneConduits, verifyChnl)

	unCODE := "REDACTED"
	blankTable := "REDACTED"
	blankCapacity := "REDACTED"

	verifyScenarios := []struct {
		verifyAlias         string
		distortPeerDetails func(*FallbackPeerDetails)
		anticipateFault        bool
	}{
		{
			"REDACTED",
			func(ni *FallbackPeerDetails) { ni.Conduits = append(conduits, byte(maximumCountConduits)) }, //
			true,
		},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Conduits = cloneConduits }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Conduits = ni.Conduits[:5] }, false},

		{"REDACTED", func(ni *FallbackPeerDetails) { ni.OverhearLocation = "REDACTED" }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.OverhearLocation = "REDACTED" }, false},

		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Edition = unCODE }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Edition = blankTable }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Edition = blankCapacity }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Edition = "REDACTED" }, false},

		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Pseudonym = unCODE }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Pseudonym = blankTable }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Pseudonym = blankCapacity }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Pseudonym = "REDACTED" }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Pseudonym = "REDACTED" }, false},

		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.TransferOrdinal = unCODE }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.TransferOrdinal = blankTable }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.TransferOrdinal = blankCapacity }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.TransferOrdinal = "REDACTED" }, false},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.TransferOrdinal = "REDACTED" }, false},

		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.RemoteLocator = unCODE }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.RemoteLocator = blankTable }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.RemoteLocator = blankCapacity }, true},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.RemoteLocator = "REDACTED" }, false},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Another.RemoteLocator = "REDACTED" }, false},
	}

	peerToken := PeerToken{PrivateToken: edwards25519.ProducePrivateToken()}
	alias := "REDACTED"

	//
	ni = verifyPeerDetails(peerToken.ID(), alias).(FallbackPeerDetails)
	ni.Conduits = conduits
	assert.NoError(t, ni.Certify())

	for _, tc := range verifyScenarios {
		ni := verifyPeerDetails(peerToken.ID(), alias).(FallbackPeerDetails)
		ni.Conduits = conduits
		tc.distortPeerDetails(&ni)
		err := ni.Certify()
		if tc.anticipateFault {
			assert.Error(t, err, tc.verifyAlias)
		} else {
			assert.NoError(t, err, tc.verifyAlias)
		}
	}
}

func VerifyPeerDetailsMatched(t *testing.T) {
	peerToken1 := PeerToken{PrivateToken: edwards25519.ProducePrivateToken()}
	peerToken2 := PeerToken{PrivateToken: edwards25519.ProducePrivateToken()}
	alias := "REDACTED"

	var freshVerifyConduit byte = 0x2

	//
	ni1 := verifyPeerDetails(peerToken1.ID(), alias).(FallbackPeerDetails)
	ni2 := verifyPeerDetails(peerToken2.ID(), alias).(FallbackPeerDetails)
	assert.NoError(t, ni1.MatchedUsing(ni2))

	//
	ni2.Conduits = append(ni2.Conduits, freshVerifyConduit)
	assert.True(t, ni2.OwnsConduit(freshVerifyConduit))
	assert.NoError(t, ni1.MatchedUsing(ni2))

	//
	_, networkLocation := GenerateDirectableLocation()
	ni3 := simulatePeerDetails{networkLocation}
	assert.Error(t, ni1.MatchedUsing(ni3))

	verifyScenarios := []struct {
		verifyAlias         string
		distortPeerDetails func(*FallbackPeerDetails)
	}{
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.SchemeEdition.Ledger++ }},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Fabric += "REDACTED" }},
		{"REDACTED", func(ni *FallbackPeerDetails) { ni.Conduits = []byte{freshVerifyConduit} }},
	}

	for _, tc := range verifyScenarios {
		ni := verifyPeerDetails(peerToken2.ID(), alias).(FallbackPeerDetails)
		tc.distortPeerDetails(&ni)
		assert.Error(t, ni1.MatchedUsing(ni))
	}
}
