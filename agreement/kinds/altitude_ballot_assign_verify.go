package kinds

import (
	"os"
	"testing"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/stretchr/testify/require"
)

var settings *cfg.Settings //

func VerifyPrimary(m *testing.M) {
	settings = verify.RestoreVerifyOrigin("REDACTED")
	cipher := m.Run()
	os.RemoveAll(settings.OriginPath)
	os.Exit(cipher)
}

func VerifyNodeOvertakeCycles(t *testing.T) {
	itemAssign, privateItems := kinds.ArbitraryAssessorAssign(10, 1)

	hvs := FreshExpandedAltitudeBallotAssign(verify.FallbackVerifySuccessionUUID, 1, itemAssign)

	ballot999_zero := createBallotResource(1, 0, 999, privateItems)
	appended, err := hvs.AppendBallot(ballot999_zero, "REDACTED", true)
	if !appended || err != nil {
		t.Error("REDACTED", appended, err)
	}

	ballot1000_zero := createBallotResource(1, 0, 1000, privateItems)
	appended, err = hvs.AppendBallot(ballot1000_zero, "REDACTED", true)
	if !appended || err != nil {
		t.Error("REDACTED", appended, err)
	}

	ballot1001_zero := createBallotResource(1, 0, 1001, privateItems)
	appended, err = hvs.AppendBallot(ballot1001_zero, "REDACTED", true)
	if err != FaultAttainedBallotOriginatingUndesiredIteration {
		t.Errorf("REDACTED", err)
	}
	if appended {
		t.Error("REDACTED")
	}

	appended, err = hvs.AppendBallot(ballot1001_zero, "REDACTED", true)
	if !appended || err != nil {
		t.Error("REDACTED")
	}
}

func VerifyUnstableAdditionData(t *testing.T) {
	itemAssign, privateItems := kinds.ArbitraryAssessorAssign(10, 1)

	hvschemeEX := FreshExpandedAltitudeBallotAssign(verify.FallbackVerifySuccessionUUID, 1, itemAssign)
	ballotNegativeAddn := createBallotResource(1, 0, 20, privateItems)
	ballotNegativeAddn.Addition, ballotNegativeAddn.AdditionNotation = nil, nil
	require.Panics(t, func() {
		_, _ = hvschemeEX.AppendBallot(ballotNegativeAddn, "REDACTED", false)
	})

	hvschemeNegativeEX := FreshAltitudeBallotAssign(verify.FallbackVerifySuccessionUUID, 1, itemAssign)
	ballotAddn := createBallotResource(1, 0, 20, privateItems)
	require.Panics(t, func() {
		_, _ = hvschemeNegativeEX.AppendBallot(ballotAddn, "REDACTED", true)
	})
}

func createBallotResource(
	altitude int64,
	itemOrdinal,
	iteration int32,
	privateItems []kinds.PrivateAssessor,
) *kinds.Ballot {
	privateItem := privateItems[itemOrdinal]
	arbitraryOctets := commitrand.Octets(tenderminthash.Extent)

	ballot, err := kinds.CreateBallot(
		privateItem,
		verify.FallbackVerifySuccessionUUID,
		itemOrdinal,
		altitude,
		iteration,
		commitchema.PreendorseKind,
		kinds.LedgerUUID{Digest: arbitraryOctets, FragmentAssignHeading: kinds.FragmentAssignHeading{}},
		committime.Now(),
	)
	if err != nil {
		panic(err)
	}

	return ballot
}
