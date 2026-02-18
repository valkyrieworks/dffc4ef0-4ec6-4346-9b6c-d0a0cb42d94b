package db

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	engineseed "github.com/valkyrieworks/utils/random"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

func Verifylast_Firstlightledgerheight(t *testing.T) {
	storeDepot := New(dbm.NewMemoryStore(), "REDACTED")

	//
	level, err := storeDepot.FinalRapidLedgerLevel()
	require.NoError(t, err)
	assert.EqualValues(t, -1, level)

	level, err = storeDepot.InitialRapidLedgerLevel()
	require.NoError(t, err)
	assert.EqualValues(t, -1, level)

	//
	err = storeDepot.PersistRapidLedger(randomRapidLedger(int64(1)))
	require.NoError(t, err)

	level, err = storeDepot.FinalRapidLedgerLevel()
	require.NoError(t, err)
	assert.EqualValues(t, 1, level)

	level, err = storeDepot.InitialRapidLedgerLevel()
	require.NoError(t, err)
	assert.EqualValues(t, 1, level)
}

func Verify_Savelightledger(t *testing.T) {
	storeDepot := New(dbm.NewMemoryStore(), "REDACTED")

	//
	h, err := storeDepot.RapidLedger(1)
	require.Error(t, err)
	assert.Nil(t, h)

	//
	err = storeDepot.PersistRapidLedger(randomRapidLedger(1))
	require.NoError(t, err)

	volume := storeDepot.Volume()
	assert.Equal(t, uint16(1), volume)
	t.Log(volume)

	h, err = storeDepot.RapidLedger(1)
	require.NoError(t, err)
	assert.NotNil(t, h)

	//
	err = storeDepot.EraseRapidLedger(1)
	require.NoError(t, err)

	h, err = storeDepot.RapidLedger(1)
	require.Error(t, err)
	assert.Nil(t, h)
}

func Verify_Lightledgerbefore(t *testing.T) {
	storeDepot := New(dbm.NewMemoryStore(), "REDACTED")

	assert.Panics(t, func() {
		_, _ = storeDepot.RapidLedgerPrior(0)
		_, _ = storeDepot.RapidLedgerPrior(100)
	})

	err := storeDepot.PersistRapidLedger(randomRapidLedger(int64(2)))
	require.NoError(t, err)

	h, err := storeDepot.RapidLedgerPrior(3)
	require.NoError(t, err)
	if assert.NotNil(t, h) {
		assert.EqualValues(t, 2, h.Level)
	}
}

func Verify_Trim(t *testing.T) {
	storeDepot := New(dbm.NewMemoryStore(), "REDACTED")

	//
	assert.EqualValues(t, 0, storeDepot.Volume())
	err := storeDepot.Trim(0)
	require.NoError(t, err)

	//
	err = storeDepot.PersistRapidLedger(randomRapidLedger(2))
	require.NoError(t, err)

	assert.EqualValues(t, 1, storeDepot.Volume())

	err = storeDepot.Trim(1)
	require.NoError(t, err)
	assert.EqualValues(t, 1, storeDepot.Volume())

	err = storeDepot.Trim(0)
	require.NoError(t, err)
	assert.EqualValues(t, 0, storeDepot.Volume())

	//
	for i := 1; i <= 10; i++ {
		err = storeDepot.PersistRapidLedger(randomRapidLedger(int64(i)))
		require.NoError(t, err)
	}

	err = storeDepot.Trim(11)
	require.NoError(t, err)
	assert.EqualValues(t, 10, storeDepot.Volume())

	err = storeDepot.Trim(7)
	require.NoError(t, err)
	assert.EqualValues(t, 7, storeDepot.Volume())
}

func Verify_Parallelism(t *testing.T) {
	storeDepot := New(dbm.NewMemoryStore(), "REDACTED")

	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()

			err := storeDepot.PersistRapidLedger(randomRapidLedger(i))
			require.NoError(t, err)

			_, err = storeDepot.RapidLedger(i)
			if err != nil {
				t.Log(err)
			}

			_, err = storeDepot.FinalRapidLedgerLevel()
			if err != nil {
				t.Log(err)
			}
			_, err = storeDepot.InitialRapidLedgerLevel()
			if err != nil {
				t.Log(err)
			}

			err = storeDepot.Trim(2)
			if err != nil {
				t.Log(err)
			}
			_ = storeDepot.Volume()

			err = storeDepot.EraseRapidLedger(1)
			if err != nil {
				t.Log(err)
			}
		}(int64(i))
	}

	wg.Wait()
}

func randomRapidLedger(level int64) *kinds.RapidLedger {
	values, _ := kinds.RandomRatifierCollection(2, 1)
	return &kinds.RapidLedger{
		AttestedHeading: &kinds.AttestedHeading{
			Heading: &kinds.Heading{
				Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 0},
				LedgerUID:            engineseed.Str(12),
				Level:             level,
				Time:               time.Now(),
				FinalLedgerUID:        kinds.LedgerUID{},
				FinalEndorseDigest:     vault.CRandomOctets(comethash.Volume),
				DataDigest:           vault.CRandomOctets(comethash.Volume),
				RatifiersDigest:     vault.CRandomOctets(comethash.Volume),
				FollowingRatifiersDigest: vault.CRandomOctets(comethash.Volume),
				AgreementDigest:      vault.CRandomOctets(comethash.Volume),
				ApplicationDigest:            vault.CRandomOctets(comethash.Volume),
				FinalOutcomesDigest:    vault.CRandomOctets(comethash.Volume),
				ProofDigest:       vault.CRandomOctets(comethash.Volume),
				RecommenderLocation:    vault.CRandomOctets(vault.LocationVolume),
			},
			Endorse: &kinds.Endorse{},
		},
		RatifierAssign: values,
	}
}
