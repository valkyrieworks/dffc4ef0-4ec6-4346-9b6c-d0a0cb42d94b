package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/vault/ed25519"
)

func VerifyMemberDetailsCertify(t *testing.T) {
	//
	ni := StandardMemberDetails{}
	assert.Error(t, ni.Certify())

	streams := make([]byte, maximumCountStreams)
	for i := 0; i < maximumCountStreams; i++ {
		streams[i] = byte(i)
	}
	cloneStreams := make([]byte, 5)
	copy(cloneStreams, streams[:5])
	cloneStreams = append(cloneStreams, verifyChan)

	notAscii := "REDACTED"
	emptyTab := "REDACTED"
	emptyArea := "REDACTED"

	verifyScenarios := []struct {
		verifyLabel         string
		distortMemberDetails func(*StandardMemberDetails)
		anticipateErr        bool
	}{
		{
			"REDACTED",
			func(ni *StandardMemberDetails) { ni.Streams = append(streams, byte(maximumCountStreams)) }, //
			true,
		},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Streams = cloneStreams }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Streams = ni.Streams[:5] }, false},

		{"REDACTED", func(ni *StandardMemberDetails) { ni.ObserveAddress = "REDACTED" }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.ObserveAddress = "REDACTED" }, false},

		{"REDACTED", func(ni *StandardMemberDetails) { ni.Release = notAscii }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Release = emptyTab }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Release = emptyArea }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Release = "REDACTED" }, false},

		{"REDACTED", func(ni *StandardMemberDetails) { ni.Moniker = notAscii }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Moniker = emptyTab }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Moniker = emptyArea }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Moniker = "REDACTED" }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Moniker = "REDACTED" }, false},

		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.TransOrdinal = notAscii }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.TransOrdinal = emptyTab }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.TransOrdinal = emptyArea }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.TransOrdinal = "REDACTED" }, false},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.TransOrdinal = "REDACTED" }, false},

		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.RPCLocation = notAscii }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.RPCLocation = emptyTab }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.RPCLocation = emptyArea }, true},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.RPCLocation = "REDACTED" }, false},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Another.RPCLocation = "REDACTED" }, false},
	}

	memberKey := MemberKey{PrivateKey: ed25519.GeneratePrivateKey()}
	label := "REDACTED"

	//
	ni = verifyMemberDetails(memberKey.ID(), label).(StandardMemberDetails)
	ni.Streams = streams
	assert.NoError(t, ni.Certify())

	for _, tc := range verifyScenarios {
		ni := verifyMemberDetails(memberKey.ID(), label).(StandardMemberDetails)
		ni.Streams = streams
		tc.distortMemberDetails(&ni)
		err := ni.Certify()
		if tc.anticipateErr {
			assert.Error(t, err, tc.verifyLabel)
		} else {
			assert.NoError(t, err, tc.verifyLabel)
		}
	}
}

func VerifyMemberDetailsHarmonious(t *testing.T) {
	memberKey1 := MemberKey{PrivateKey: ed25519.GeneratePrivateKey()}
	memberKey2 := MemberKey{PrivateKey: ed25519.GeneratePrivateKey()}
	label := "REDACTED"

	var newVerifyConduit byte = 0x2

	//
	ni1 := verifyMemberDetails(memberKey1.ID(), label).(StandardMemberDetails)
	ni2 := verifyMemberDetails(memberKey2.ID(), label).(StandardMemberDetails)
	assert.NoError(t, ni1.HarmoniousWith(ni2))

	//
	ni2.Streams = append(ni2.Streams, newVerifyConduit)
	assert.True(t, ni2.HasConduit(newVerifyConduit))
	assert.NoError(t, ni1.HarmoniousWith(ni2))

	//
	_, netAddress := InstantiateForwardableAddress()
	ni3 := emulateMemberDetails{netAddress}
	assert.Error(t, ni1.HarmoniousWith(ni3))

	verifyScenarios := []struct {
		verifyLabel         string
		distortMemberDetails func(*StandardMemberDetails)
	}{
		{"REDACTED", func(ni *StandardMemberDetails) { ni.ProtocolRelease.Ledger++ }},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Fabric += "REDACTED" }},
		{"REDACTED", func(ni *StandardMemberDetails) { ni.Streams = []byte{newVerifyConduit} }},
	}

	for _, tc := range verifyScenarios {
		ni := verifyMemberDetails(memberKey2.ID(), label).(StandardMemberDetails)
		tc.distortMemberDetails(&ni)
		assert.Error(t, ni1.HarmoniousWith(ni))
	}
}
