package db

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

func Verifylast_Initialledgerpeak(t *testing.T) {
	datastoreDepot := New(dbm.FreshMemoryDatastore(), "REDACTED")

	//
	altitude, err := datastoreDepot.FinalAgileLedgerAltitude()
	require.NoError(t, err)
	assert.EqualValues(t, -1, altitude)

	altitude, err = datastoreDepot.InitialAgileLedgerAltitude()
	require.NoError(t, err)
	assert.EqualValues(t, -1, altitude)

	//
	err = datastoreDepot.PersistAgileLedger(arbitraryAgileLedger(int64(1)))
	require.NoError(t, err)

	altitude, err = datastoreDepot.FinalAgileLedgerAltitude()
	require.NoError(t, err)
	assert.EqualValues(t, 1, altitude)

	altitude, err = datastoreDepot.InitialAgileLedgerAltitude()
	require.NoError(t, err)
	assert.EqualValues(t, 1, altitude)
}

func Verify_Persistledger(t *testing.T) {
	datastoreDepot := New(dbm.FreshMemoryDatastore(), "REDACTED")

	//
	h, err := datastoreDepot.AgileLedger(1)
	require.Error(t, err)
	assert.Nil(t, h)

	//
	err = datastoreDepot.PersistAgileLedger(arbitraryAgileLedger(1))
	require.NoError(t, err)

	extent := datastoreDepot.Extent()
	assert.Equal(t, uint16(1), extent)
	t.Log(extent)

	h, err = datastoreDepot.AgileLedger(1)
	require.NoError(t, err)
	assert.NotNil(t, h)

	//
	err = datastoreDepot.EraseAgileLedger(1)
	require.NoError(t, err)

	h, err = datastoreDepot.AgileLedger(1)
	require.Error(t, err)
	assert.Nil(t, h)
}

func Verify_Ledgerbefore(t *testing.T) {
	datastoreDepot := New(dbm.FreshMemoryDatastore(), "REDACTED")

	assert.Panics(t, func() {
		_, _ = datastoreDepot.AgileLedgerPrior(0)
		_, _ = datastoreDepot.AgileLedgerPrior(100)
	})

	err := datastoreDepot.PersistAgileLedger(arbitraryAgileLedger(int64(2)))
	require.NoError(t, err)

	h, err := datastoreDepot.AgileLedgerPrior(3)
	require.NoError(t, err)
	if assert.NotNil(t, h) {
		assert.EqualValues(t, 2, h.Altitude)
	}
}

func Verify_Trim(t *testing.T) {
	datastoreDepot := New(dbm.FreshMemoryDatastore(), "REDACTED")

	//
	assert.EqualValues(t, 0, datastoreDepot.Extent())
	err := datastoreDepot.Trim(0)
	require.NoError(t, err)

	//
	err = datastoreDepot.PersistAgileLedger(arbitraryAgileLedger(2))
	require.NoError(t, err)

	assert.EqualValues(t, 1, datastoreDepot.Extent())

	err = datastoreDepot.Trim(1)
	require.NoError(t, err)
	assert.EqualValues(t, 1, datastoreDepot.Extent())

	err = datastoreDepot.Trim(0)
	require.NoError(t, err)
	assert.EqualValues(t, 0, datastoreDepot.Extent())

	//
	for i := 1; i <= 10; i++ {
		err = datastoreDepot.PersistAgileLedger(arbitraryAgileLedger(int64(i)))
		require.NoError(t, err)
	}

	err = datastoreDepot.Trim(11)
	require.NoError(t, err)
	assert.EqualValues(t, 10, datastoreDepot.Extent())

	err = datastoreDepot.Trim(7)
	require.NoError(t, err)
	assert.EqualValues(t, 7, datastoreDepot.Extent())
}

func Verify_Parallelism(t *testing.T) {
	datastoreDepot := New(dbm.FreshMemoryDatastore(), "REDACTED")

	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()

			err := datastoreDepot.PersistAgileLedger(arbitraryAgileLedger(i))
			require.NoError(t, err)

			_, err = datastoreDepot.AgileLedger(i)
			if err != nil {
				t.Log(err)
			}

			_, err = datastoreDepot.FinalAgileLedgerAltitude()
			if err != nil {
				t.Log(err)
			}
			_, err = datastoreDepot.InitialAgileLedgerAltitude()
			if err != nil {
				t.Log(err)
			}

			err = datastoreDepot.Trim(2)
			if err != nil {
				t.Log(err)
			}
			_ = datastoreDepot.Extent()

			err = datastoreDepot.EraseAgileLedger(1)
			if err != nil {
				t.Log(err)
			}
		}(int64(i))
	}

	wg.Wait()
}

func arbitraryAgileLedger(altitude int64) *kinds.AgileLedger {
	values, _ := kinds.ArbitraryAssessorAssign(2, 1)
	return &kinds.AgileLedger{
		NotatedHeading: &kinds.NotatedHeading{
			Heading: &kinds.Heading{
				Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 0},
				SuccessionUUID:            commitrand.Str(12),
				Altitude:             altitude,
				Moment:               time.Now(),
				FinalLedgerUUID:        kinds.LedgerUUID{},
				FinalEndorseDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
				DataDigest:           security.CHARArbitraryOctets(tenderminthash.Extent),
				AssessorsDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
				FollowingAssessorsDigest: security.CHARArbitraryOctets(tenderminthash.Extent),
				AgreementDigest:      security.CHARArbitraryOctets(tenderminthash.Extent),
				PlatformDigest:            security.CHARArbitraryOctets(tenderminthash.Extent),
				FinalOutcomesDigest:    security.CHARArbitraryOctets(tenderminthash.Extent),
				ProofDigest:       security.CHARArbitraryOctets(tenderminthash.Extent),
				NominatorLocation:    security.CHARArbitraryOctets(security.LocatorExtent),
			},
			Endorse: &kinds.Endorse{},
		},
		AssessorAssign: values,
	}
}
