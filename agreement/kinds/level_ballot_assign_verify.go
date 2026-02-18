package kinds

import (
	"os"
	"testing"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/intrinsic/verify"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
	"github.com/stretchr/testify/require"
)

var settings *cfg.Settings //

func VerifyMain(m *testing.M) {
	settings = verify.RestoreVerifyOrigin("REDACTED")
	code := m.Run()
	os.RemoveAll(settings.OriginFolder)
	os.Exit(code)
}

func VerifyNodeOvertakeIterations(t *testing.T) {
	valueCollection, privateValues := kinds.RandomRatifierCollection(10, 1)

	hvs := NewExpandedLevelBallotCollection(verify.StandardVerifyLedgerUID, 1, valueCollection)

	ballot999_0 := createBallotHR(1, 0, 999, privateValues)
	appended, err := hvs.AppendBallot(ballot999_0, "REDACTED", true)
	if !appended || err != nil {
		t.Error("REDACTED", appended, err)
	}

	ballot1000_0 := createBallotHR(1, 0, 1000, privateValues)
	appended, err = hvs.AppendBallot(ballot1000_0, "REDACTED", true)
	if !appended || err != nil {
		t.Error("REDACTED", appended, err)
	}

	ballot1001_0 := createBallotHR(1, 0, 1001, privateValues)
	appended, err = hvs.AppendBallot(ballot1001_0, "REDACTED", true)
	if err != ErrAcquiredBallotFromUndesiredEpoch {
		t.Errorf("REDACTED", err)
	}
	if appended {
		t.Error("REDACTED")
	}

	appended, err = hvs.AppendBallot(ballot1001_0, "REDACTED", true)
	if !appended || err != nil {
		t.Error("REDACTED")
	}
}

func VerifyDiscordantAdditionData(t *testing.T) {
	valueCollection, privateValues := kinds.RandomRatifierCollection(10, 1)

	hvsE := NewExpandedLevelBallotCollection(verify.StandardVerifyLedgerUID, 1, valueCollection)
	ballotNoExtension := createBallotHR(1, 0, 20, privateValues)
	ballotNoExtension.Addition, ballotNoExtension.AdditionAutograph = nil, nil
	require.Panics(t, func() {
		_, _ = hvsE.AppendBallot(ballotNoExtension, "REDACTED", false)
	})

	hvsNoE := NewLevelBallotCollection(verify.StandardVerifyLedgerUID, 1, valueCollection)
	ballotExtension := createBallotHR(1, 0, 20, privateValues)
	require.Panics(t, func() {
		_, _ = hvsNoE.AppendBallot(ballotExtension, "REDACTED", true)
	})
}

func createBallotHR(
	level int64,
	valueOrdinal,
	duration int32,
	privateValues []kinds.PrivateRatifier,
) *kinds.Ballot {
	privateValue := privateValues[valueOrdinal]
	randomOctets := engineseed.Octets(comethash.Volume)

	ballot, err := kinds.CreateBallot(
		privateValue,
		verify.StandardVerifyLedgerUID,
		valueOrdinal,
		level,
		duration,
		engineproto.PreendorseKind,
		kinds.LedgerUID{Digest: randomOctets, SegmentAssignHeading: kinds.SegmentAssignHeading{}},
		engineclock.Now(),
	)
	if err != nil {
		panic(err)
	}

	return ballot
}
