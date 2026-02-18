package status_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/simulations"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

func VerifyRevert(t *testing.T) {
	var (
		level     int64 = 100
		followingLevel int64 = 101
	)
	ledgerDepot := &simulations.LedgerDepot{}
	statusDepot := configureStatusDepot(t, level)
	primaryStatus, err := statusDepot.Import()
	require.NoError(t, err)

	//
	newOptions := kinds.StandardAgreementOptions()
	newOptions.Release.App = 11
	newOptions.Ledger.MaximumOctets = 1000
	followingStatus := primaryStatus.Clone()
	followingStatus.FinalLedgerLevel = followingLevel
	followingStatus.Release.Agreement.App = 11
	followingStatus.FinalLedgerUID = createLedgerUIDArbitrary()
	followingStatus.ApplicationDigest = comethash.Sum([]byte("REDACTED"))
	followingStatus.FinalRatifiers = primaryStatus.Ratifiers
	followingStatus.Ratifiers = primaryStatus.FollowingRatifiers
	followingStatus.FollowingRatifiers = primaryStatus.FollowingRatifiers.CloneAugmentRecommenderUrgency(1)
	followingStatus.AgreementOptions = *newOptions
	followingStatus.FinalLevelAgreementOptionsModified = followingLevel + 1
	followingStatus.FinalLevelRatifiersModified = followingLevel + 1

	//
	require.NoError(t, statusDepot.Persist(followingStatus))

	ledger := &kinds.LedgerMeta{
		LedgerUID: primaryStatus.FinalLedgerUID,
		Heading: kinds.Heading{
			Level:          primaryStatus.FinalLedgerLevel,
			Time:            primaryStatus.FinalLedgerTime,
			ApplicationDigest:         vault.CRandomOctets(comethash.Volume),
			FinalLedgerUID:     createLedgerUIDArbitrary(),
			FinalOutcomesDigest: primaryStatus.FinalOutcomesDigest,
		},
	}
	followingLedger := &kinds.LedgerMeta{
		LedgerUID: primaryStatus.FinalLedgerUID,
		Heading: kinds.Heading{
			Level:          followingStatus.FinalLedgerLevel,
			ApplicationDigest:         primaryStatus.ApplicationDigest,
			FinalLedgerUID:     ledger.LedgerUID,
			Time:            followingStatus.FinalLedgerTime,
			FinalOutcomesDigest: followingStatus.FinalOutcomesDigest,
		},
	}
	ledgerDepot.On("REDACTED", level).Return(ledger)
	ledgerDepot.On("REDACTED", followingLevel).Return(followingLedger)
	ledgerDepot.On("REDACTED").Return(followingLevel)

	//
	revertLevel, revertDigest, err := status.Revert(ledgerDepot, statusDepot, false)
	require.NoError(t, err)
	require.EqualValues(t, level, revertLevel)
	require.EqualValues(t, primaryStatus.ApplicationDigest, revertDigest)
	ledgerDepot.AssertExpectations(t)

	//
	retrievedStatus, err := statusDepot.Import()
	require.NoError(t, err)
	require.EqualValues(t, primaryStatus, retrievedStatus)
}

func VerifyRevertRigid(t *testing.T) {
	const level int64 = 100
	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	statusDepot := status.NewDepot(dbm.NewMemoryStore(), status.DepotSettings{DropIfaceReplies: false})

	valueCollection, _ := kinds.RandomRatifierCollection(5, 10)

	options := kinds.StandardAgreementOptions()
	options.Release.App = 10
	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	ledger := &kinds.Ledger{
		Heading: kinds.Heading{
			Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1},
			LedgerUID:            "REDACTED",
			Time:               now,
			Level:             level,
			ApplicationDigest:            vault.CRandomOctets(comethash.Volume),
			FinalLedgerUID:        createLedgerUIDArbitrary(),
			FinalEndorseDigest:     vault.CRandomOctets(comethash.Volume),
			DataDigest:           vault.CRandomOctets(comethash.Volume),
			RatifiersDigest:     valueCollection.Digest(),
			FollowingRatifiersDigest: valueCollection.CloneAugmentRecommenderUrgency(1).Digest(),
			AgreementDigest:      options.Digest(),
			FinalOutcomesDigest:    vault.CRandomOctets(comethash.Volume),
			ProofDigest:       vault.CRandomOctets(comethash.Volume),
			RecommenderLocation:    vault.CRandomOctets(vault.LocationVolume),
		},
		FinalEndorse: &kinds.Endorse{Level: level - 1},
	}

	sectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	ledgerDepot.PersistLedger(ledger, sectionCollection, &kinds.Endorse{Level: ledger.Level})

	currentStatus := status.Status{
		Release: cometstatus.Release{
			Agreement: ledger.Release,
			Software:  release.TMCoreSemaphoreRev,
		},
		FinalLedgerLevel:                  ledger.Level,
		FinalLedgerTime:                    ledger.Time,
		ApplicationDigest:                          vault.CRandomOctets(comethash.Volume),
		FinalRatifiers:                   valueCollection,
		Ratifiers:                       valueCollection.CloneAugmentRecommenderUrgency(1),
		FollowingRatifiers:                   valueCollection.CloneAugmentRecommenderUrgency(2),
		AgreementOptions:                  *options,
		FinalLevelAgreementOptionsModified: level + 1,
		FinalLevelRatifiersModified:      level + 1,
		FinalOutcomesDigest:                  vault.CRandomOctets(comethash.Volume),
	}
	require.NoError(t, statusDepot.Onboard(currentStatus))

	followingLedger := &kinds.Ledger{
		Heading: kinds.Heading{
			Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1},
			LedgerUID:            ledger.LedgerUID,
			Time:               ledger.Time,
			Level:             currentStatus.FinalLedgerLevel + 1,
			ApplicationDigest:            currentStatus.ApplicationDigest,
			FinalLedgerUID:        kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: sectionCollection.Heading()},
			FinalEndorseDigest:     vault.CRandomOctets(comethash.Volume),
			DataDigest:           vault.CRandomOctets(comethash.Volume),
			RatifiersDigest:     valueCollection.CloneAugmentRecommenderUrgency(1).Digest(),
			FollowingRatifiersDigest: valueCollection.CloneAugmentRecommenderUrgency(2).Digest(),
			AgreementDigest:      options.Digest(),
			FinalOutcomesDigest:    currentStatus.FinalOutcomesDigest,
			ProofDigest:       vault.CRandomOctets(comethash.Volume),
			RecommenderLocation:    vault.CRandomOctets(vault.LocationVolume),
		},
		FinalEndorse: &kinds.Endorse{Level: currentStatus.FinalLedgerLevel},
	}

	followingSectionCollection, err := followingLedger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	ledgerDepot.PersistLedger(followingLedger, followingSectionCollection, &kinds.Endorse{Level: followingLedger.Level})

	revertLevel, revertDigest, err := status.Revert(ledgerDepot, statusDepot, true)
	require.NoError(t, err)
	require.Equal(t, revertLevel, currentStatus.FinalLedgerLevel)
	require.Equal(t, revertDigest, currentStatus.ApplicationDigest)

	//
	retrievedStatus, err := statusDepot.Import()
	require.NoError(t, err)
	require.Equal(t, currentStatus, retrievedStatus)

	//
	ledgerDepot.PersistLedger(followingLedger, followingSectionCollection, &kinds.Endorse{Level: followingLedger.Level})

	options.Release.App = 11

	followingStatus := status.Status{
		Release: cometstatus.Release{
			Agreement: ledger.Release,
			Software:  release.TMCoreSemaphoreRev,
		},
		FinalLedgerLevel:                  followingLedger.Level,
		FinalLedgerTime:                    followingLedger.Time,
		ApplicationDigest:                          vault.CRandomOctets(comethash.Volume),
		FinalRatifiers:                   valueCollection.CloneAugmentRecommenderUrgency(1),
		Ratifiers:                       valueCollection.CloneAugmentRecommenderUrgency(2),
		FollowingRatifiers:                   valueCollection.CloneAugmentRecommenderUrgency(3),
		AgreementOptions:                  *options,
		FinalLevelAgreementOptionsModified: followingLedger.Level + 1,
		FinalLevelRatifiersModified:      followingLedger.Level + 1,
		FinalOutcomesDigest:                  vault.CRandomOctets(comethash.Volume),
	}
	require.NoError(t, statusDepot.Persist(followingStatus))

	revertLevel, revertDigest, err = status.Revert(ledgerDepot, statusDepot, true)
	require.NoError(t, err)
	require.Equal(t, revertLevel, currentStatus.FinalLedgerLevel)
	require.Equal(t, revertDigest, currentStatus.ApplicationDigest)
}

func VerifyRevertNoStatus(t *testing.T) {
	statusDepot := status.NewDepot(dbm.NewMemoryStore(),
		status.DepotSettings{
			DropIfaceReplies: false,
		})
	ledgerDepot := &simulations.LedgerDepot{}

	_, _, err := status.Revert(ledgerDepot, statusDepot, false)
	require.Error(t, err)
	require.Contains(t, err.Error(), "REDACTED")
}

func VerifyRevertNoLedgers(t *testing.T) {
	const level = int64(100)
	statusDepot := configureStatusDepot(t, level)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED").Return(level)
	ledgerDepot.On("REDACTED", level).Return(nil)
	ledgerDepot.On("REDACTED", level-1).Return(nil)

	_, _, err := status.Revert(ledgerDepot, statusDepot, false)
	require.Error(t, err)
	require.Contains(t, err.Error(), "REDACTED")
}

func VerifyRevertDistinctStatusLevel(t *testing.T) {
	const level = int64(100)
	statusDepot := configureStatusDepot(t, level)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED").Return(level + 2)

	_, _, err := status.Revert(ledgerDepot, statusDepot, false)
	require.Error(t, err)
	require.Equal(t, err.Error(), "REDACTED")
}

func configureStatusDepot(t *testing.T, level int64) status.Depot {
	statusDepot := status.NewDepot(dbm.NewMemoryStore(), status.DepotSettings{DropIfaceReplies: false})
	valueCollection, _ := kinds.RandomRatifierCollection(5, 10)

	options := kinds.StandardAgreementOptions()
	options.Release.App = 10

	primaryStatus := status.Status{
		Release: cometstatus.Release{
			Agreement: cometrelease.Agreement{
				Ledger: release.LedgerProtocol,
				App:   10,
			},
			Software: release.TMCoreSemaphoreRev,
		},
		LedgerUID:                          "REDACTED",
		PrimaryLevel:                    10,
		FinalLedgerUID:                      createLedgerUIDArbitrary(),
		ApplicationDigest:                          comethash.Sum([]byte("REDACTED")),
		FinalOutcomesDigest:                  comethash.Sum([]byte("REDACTED")),
		FinalLedgerLevel:                  level,
		FinalLedgerTime:                    time.Now(),
		FinalRatifiers:                   valueCollection,
		Ratifiers:                       valueCollection.CloneAugmentRecommenderUrgency(1),
		FollowingRatifiers:                   valueCollection.CloneAugmentRecommenderUrgency(2),
		FinalLevelRatifiersModified:      level + 1 + 1,
		AgreementOptions:                  *options,
		FinalLevelAgreementOptionsModified: level + 1,
	}
	require.NoError(t, statusDepot.Onboard(primaryStatus))
	return statusDepot
}

func createLedgerUIDArbitrary() kinds.LedgerUID {
	var (
		ledgerDigest   = make([]byte, comethash.Volume)
		sectionCollectionDigest = make([]byte, comethash.Volume)
	)
	rand.Read(ledgerDigest)   //
	rand.Read(sectionCollectionDigest) //
	return kinds.LedgerUID{
		Digest: ledgerDigest,
		SegmentAssignHeading: kinds.SegmentAssignHeading{
			Sum: 123,
			Digest:  sectionCollectionDigest,
		},
	}
}
