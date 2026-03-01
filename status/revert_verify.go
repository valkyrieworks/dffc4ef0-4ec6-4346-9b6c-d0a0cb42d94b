package status_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

func VerifyRevert(t *testing.T) {
	var (
		altitude     int64 = 100
		followingAltitude int64 = 101
	)
	ledgerDepot := &simulations.LedgerDepot{}
	statusDepot := configureStatusDepot(t, altitude)
	primaryStatus, err := statusDepot.Fetch()
	require.NoError(t, err)

	//
	freshParameters := kinds.FallbackAgreementSettings()
	freshParameters.Edition.App = 11
	freshParameters.Ledger.MaximumOctets = 1000
	followingStatus := primaryStatus.Duplicate()
	followingStatus.FinalLedgerAltitude = followingAltitude
	followingStatus.Edition.Agreement.App = 11
	followingStatus.FinalLedgerUUID = createLedgerUUIDUnpredictable()
	followingStatus.PlatformDigest = tenderminthash.Sum([]byte("REDACTED"))
	followingStatus.FinalAssessors = primaryStatus.Assessors
	followingStatus.Assessors = primaryStatus.FollowingAssessors
	followingStatus.FollowingAssessors = primaryStatus.FollowingAssessors.DuplicateAdvanceNominatorUrgency(1)
	followingStatus.AgreementSettings = *freshParameters
	followingStatus.FinalAltitudeAgreementParametersAltered = followingAltitude + 1
	followingStatus.FinalAltitudeAssessorsAltered = followingAltitude + 1

	//
	require.NoError(t, statusDepot.Persist(followingStatus))

	ledger := &kinds.LedgerSummary{
		LedgerUUID: primaryStatus.FinalLedgerUUID,
		Heading: kinds.Heading{
			Altitude:          primaryStatus.FinalLedgerAltitude,
			Moment:            primaryStatus.FinalLedgerMoment,
			PlatformDigest:         security.CHARArbitraryOctets(tenderminthash.Extent),
			FinalLedgerUUID:     createLedgerUUIDUnpredictable(),
			FinalOutcomesDigest: primaryStatus.FinalOutcomesDigest,
		},
	}
	followingLedger := &kinds.LedgerSummary{
		LedgerUUID: primaryStatus.FinalLedgerUUID,
		Heading: kinds.Heading{
			Altitude:          followingStatus.FinalLedgerAltitude,
			PlatformDigest:         primaryStatus.PlatformDigest,
			FinalLedgerUUID:     ledger.LedgerUUID,
			Moment:            followingStatus.FinalLedgerMoment,
			FinalOutcomesDigest: followingStatus.FinalOutcomesDigest,
		},
	}
	ledgerDepot.On("REDACTED", altitude).Return(ledger)
	ledgerDepot.On("REDACTED", followingAltitude).Return(followingLedger)
	ledgerDepot.On("REDACTED").Return(followingAltitude)

	//
	revertAltitude, revertDigest, err := status.Revert(ledgerDepot, statusDepot, false)
	require.NoError(t, err)
	require.EqualValues(t, altitude, revertAltitude)
	require.EqualValues(t, primaryStatus.PlatformDigest, revertDigest)
	ledgerDepot.AssertExpectations(t)

	//
	retrievedStatus, err := statusDepot.Fetch()
	require.NoError(t, err)
	require.EqualValues(t, primaryStatus, retrievedStatus)
}

func VerifyRevertRigid(t *testing.T) {
	const altitude int64 = 100
	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	statusDepot := status.FreshDepot(dbm.FreshMemoryDatastore(), status.DepotChoices{EjectIfaceReplies: false})

	itemAssign, _ := kinds.ArbitraryAssessorAssign(5, 10)

	parameters := kinds.FallbackAgreementSettings()
	parameters.Edition.App = 10
	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	ledger := &kinds.Ledger{
		Heading: kinds.Heading{
			Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1},
			SuccessionUUID:            "REDACTED",
			Moment:               now,
			Altitude:             altitude,
			PlatformDigest:            security.CHARArbitraryOctets(tenderminthash.Extent),
			FinalLedgerUUID:        createLedgerUUIDUnpredictable(),
			FinalEndorseDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
			DataDigest:           security.CHARArbitraryOctets(tenderminthash.Extent),
			AssessorsDigest:     itemAssign.Digest(),
			FollowingAssessorsDigest: itemAssign.DuplicateAdvanceNominatorUrgency(1).Digest(),
			AgreementDigest:      parameters.Digest(),
			FinalOutcomesDigest:    security.CHARArbitraryOctets(tenderminthash.Extent),
			ProofDigest:       security.CHARArbitraryOctets(tenderminthash.Extent),
			NominatorLocation:    security.CHARArbitraryOctets(security.LocatorExtent),
		},
		FinalEndorse: &kinds.Endorse{Altitude: altitude - 1},
	}

	fragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	ledgerDepot.PersistLedger(ledger, fragmentAssign, &kinds.Endorse{Altitude: ledger.Altitude})

	presentStatus := status.Status{
		Edition: strongstatus.Edition{
			Agreement: ledger.Edition,
			Package:  edition.TEMPBaseSemaphoreEdtn,
		},
		FinalLedgerAltitude:                  ledger.Altitude,
		FinalLedgerMoment:                    ledger.Moment,
		PlatformDigest:                          security.CHARArbitraryOctets(tenderminthash.Extent),
		FinalAssessors:                   itemAssign,
		Assessors:                       itemAssign.DuplicateAdvanceNominatorUrgency(1),
		FollowingAssessors:                   itemAssign.DuplicateAdvanceNominatorUrgency(2),
		AgreementSettings:                  *parameters,
		FinalAltitudeAgreementParametersAltered: altitude + 1,
		FinalAltitudeAssessorsAltered:      altitude + 1,
		FinalOutcomesDigest:                  security.CHARArbitraryOctets(tenderminthash.Extent),
	}
	require.NoError(t, statusDepot.Onboard(presentStatus))

	followingLedger := &kinds.Ledger{
		Heading: kinds.Heading{
			Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1},
			SuccessionUUID:            ledger.SuccessionUUID,
			Moment:               ledger.Moment,
			Altitude:             presentStatus.FinalLedgerAltitude + 1,
			PlatformDigest:            presentStatus.PlatformDigest,
			FinalLedgerUUID:        kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: fragmentAssign.Heading()},
			FinalEndorseDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
			DataDigest:           security.CHARArbitraryOctets(tenderminthash.Extent),
			AssessorsDigest:     itemAssign.DuplicateAdvanceNominatorUrgency(1).Digest(),
			FollowingAssessorsDigest: itemAssign.DuplicateAdvanceNominatorUrgency(2).Digest(),
			AgreementDigest:      parameters.Digest(),
			FinalOutcomesDigest:    presentStatus.FinalOutcomesDigest,
			ProofDigest:       security.CHARArbitraryOctets(tenderminthash.Extent),
			NominatorLocation:    security.CHARArbitraryOctets(security.LocatorExtent),
		},
		FinalEndorse: &kinds.Endorse{Altitude: presentStatus.FinalLedgerAltitude},
	}

	followingFragmentAssign, err := followingLedger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	ledgerDepot.PersistLedger(followingLedger, followingFragmentAssign, &kinds.Endorse{Altitude: followingLedger.Altitude})

	revertAltitude, revertDigest, err := status.Revert(ledgerDepot, statusDepot, true)
	require.NoError(t, err)
	require.Equal(t, revertAltitude, presentStatus.FinalLedgerAltitude)
	require.Equal(t, revertDigest, presentStatus.PlatformDigest)

	//
	retrievedStatus, err := statusDepot.Fetch()
	require.NoError(t, err)
	require.Equal(t, presentStatus, retrievedStatus)

	//
	ledgerDepot.PersistLedger(followingLedger, followingFragmentAssign, &kinds.Endorse{Altitude: followingLedger.Altitude})

	parameters.Edition.App = 11

	followingStatus := status.Status{
		Edition: strongstatus.Edition{
			Agreement: ledger.Edition,
			Package:  edition.TEMPBaseSemaphoreEdtn,
		},
		FinalLedgerAltitude:                  followingLedger.Altitude,
		FinalLedgerMoment:                    followingLedger.Moment,
		PlatformDigest:                          security.CHARArbitraryOctets(tenderminthash.Extent),
		FinalAssessors:                   itemAssign.DuplicateAdvanceNominatorUrgency(1),
		Assessors:                       itemAssign.DuplicateAdvanceNominatorUrgency(2),
		FollowingAssessors:                   itemAssign.DuplicateAdvanceNominatorUrgency(3),
		AgreementSettings:                  *parameters,
		FinalAltitudeAgreementParametersAltered: followingLedger.Altitude + 1,
		FinalAltitudeAssessorsAltered:      followingLedger.Altitude + 1,
		FinalOutcomesDigest:                  security.CHARArbitraryOctets(tenderminthash.Extent),
	}
	require.NoError(t, statusDepot.Persist(followingStatus))

	revertAltitude, revertDigest, err = status.Revert(ledgerDepot, statusDepot, true)
	require.NoError(t, err)
	require.Equal(t, revertAltitude, presentStatus.FinalLedgerAltitude)
	require.Equal(t, revertDigest, presentStatus.PlatformDigest)
}

func VerifyRevertNegativeStatus(t *testing.T) {
	statusDepot := status.FreshDepot(dbm.FreshMemoryDatastore(),
		status.DepotChoices{
			EjectIfaceReplies: false,
		})
	ledgerDepot := &simulations.LedgerDepot{}

	_, _, err := status.Revert(ledgerDepot, statusDepot, false)
	require.Error(t, err)
	require.Contains(t, err.Error(), "REDACTED")
}

func VerifyRevertNegativeLedgers(t *testing.T) {
	const altitude = int64(100)
	statusDepot := configureStatusDepot(t, altitude)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED").Return(altitude)
	ledgerDepot.On("REDACTED", altitude).Return(nil)
	ledgerDepot.On("REDACTED", altitude-1).Return(nil)

	_, _, err := status.Revert(ledgerDepot, statusDepot, false)
	require.Error(t, err)
	require.Contains(t, err.Error(), "REDACTED")
}

func VerifyRevertDistinctStatusAltitude(t *testing.T) {
	const altitude = int64(100)
	statusDepot := configureStatusDepot(t, altitude)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED").Return(altitude + 2)

	_, _, err := status.Revert(ledgerDepot, statusDepot, false)
	require.Error(t, err)
	require.Equal(t, err.Error(), "REDACTED")
}

func configureStatusDepot(t *testing.T, altitude int64) status.Depot {
	statusDepot := status.FreshDepot(dbm.FreshMemoryDatastore(), status.DepotChoices{EjectIfaceReplies: false})
	itemAssign, _ := kinds.ArbitraryAssessorAssign(5, 10)

	parameters := kinds.FallbackAgreementSettings()
	parameters.Edition.App = 10

	primaryStatus := status.Status{
		Edition: strongstatus.Edition{
			Agreement: strongmindedition.Agreement{
				Ledger: edition.LedgerScheme,
				App:   10,
			},
			Package: edition.TEMPBaseSemaphoreEdtn,
		},
		SuccessionUUID:                          "REDACTED",
		PrimaryAltitude:                    10,
		FinalLedgerUUID:                      createLedgerUUIDUnpredictable(),
		PlatformDigest:                          tenderminthash.Sum([]byte("REDACTED")),
		FinalOutcomesDigest:                  tenderminthash.Sum([]byte("REDACTED")),
		FinalLedgerAltitude:                  altitude,
		FinalLedgerMoment:                    time.Now(),
		FinalAssessors:                   itemAssign,
		Assessors:                       itemAssign.DuplicateAdvanceNominatorUrgency(1),
		FollowingAssessors:                   itemAssign.DuplicateAdvanceNominatorUrgency(2),
		FinalAltitudeAssessorsAltered:      altitude + 1 + 1,
		AgreementSettings:                  *parameters,
		FinalAltitudeAgreementParametersAltered: altitude + 1,
	}
	require.NoError(t, statusDepot.Onboard(primaryStatus))
	return statusDepot
}

func createLedgerUUIDUnpredictable() kinds.LedgerUUID {
	var (
		ledgerDigest   = make([]byte, tenderminthash.Extent)
		fragmentAssignDigest = make([]byte, tenderminthash.Extent)
	)
	rand.Read(ledgerDigest)   //
	rand.Read(fragmentAssignDigest) //
	return kinds.LedgerUUID{
		Digest: ledgerDigest,
		FragmentAssignHeading: kinds.FragmentAssignHeading{
			Sum: 123,
			Digest:  fragmentAssignDigest,
		},
	}
}
