package kinds

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cryptomocks "github.com/valkyrieworks/vault/simulations"
	cometmath "github.com/valkyrieworks/utils/math"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

//
//
func Verifyratifierset_Validateendorse_All(t *testing.T) {
	var (
		epoch  = int32(0)
		level = int64(100)

		ledgerUID    = createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
		ledgerUID    = "REDACTED"
		validateLayer = cometmath.Portion{Dividend: 2, Divisor: 3}
	)

	verifyScenarios := []struct {
		summary, summary2 string //
		//
		ledgerUID string
		//
		ledgerUID LedgerUID
		valueVolume int

		//
		level int64

		//
		ledgerBallots  int
		nullBallots    int
		missingBallots int

		expirationErr bool
	}{
		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 3, level, 3, 0, 0, false},
		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 1, level, 1, 0, 0, false},

		{"REDACTED", "REDACTED", "REDACTED", ledgerUID, 2, level, 2, 0, 0, true},
		{"REDACTED", "REDACTED", ledgerUID, createLedgerUIDArbitrary(), 2, level, 2, 0, 0, true},
		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 1, level - 1, 1, 0, 0, true},

		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 4, level, 3, 0, 0, true},
		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 1, level, 2, 0, 0, true},

		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 10, level, 3, 2, 5, true},
		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 1, level, 0, 0, 1, true}, //
		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 1, level, 0, 1, 0, true}, //
		{"REDACTED", "REDACTED", ledgerUID, ledgerUID, 9, level, 6, 3, 0, true},
	}

	for _, tc := range verifyScenarios {
		numberAllEndorsements := false
		f := func(t *testing.T) {
			_, valueCollection, values := randomBallotCollection(tc.level, epoch, engineproto.PreendorseKind, tc.valueVolume, 10, false)
			sumBallots := tc.ledgerBallots + tc.missingBallots + tc.nullBallots
			autographs := make([]EndorseSignature, sumBallots)
			vi := 0
			//
			for i := 0; i < tc.missingBallots; i++ {
				autographs[vi] = NewEndorseSignatureMissing()
				vi++
			}
			for i := 0; i < tc.ledgerBallots+tc.nullBallots; i++ {

				publicKey, err := values[vi%len(values)].FetchPublicKey()
				require.NoError(t, err)
				ballot := &Ballot{
					RatifierLocation: publicKey.Location(),
					RatifierOrdinal:   int32(vi),
					Level:           tc.level,
					Cycle:            epoch,
					Kind:             engineproto.PreendorseKind,
					LedgerUID:          tc.ledgerUID,
					Timestamp:        time.Now(),
				}
				if i >= tc.ledgerBallots {
					ballot.LedgerUID = LedgerUID{}
				}

				v := ballot.ToSchema()

				require.NoError(t, values[vi%len(values)].AttestBallot(tc.ledgerUID, v))
				ballot.Autograph = v.Autograph

				autographs[vi] = ballot.EndorseSignature()

				vi++
			}
			endorse := &Endorse{
				Level:     tc.level,
				Cycle:      epoch,
				LedgerUID:    tc.ledgerUID,
				Endorsements: autographs,
			}

			err := valueCollection.ValidateEndorse(ledgerUID, ledgerUID, level, endorse)
			if tc.expirationErr {
				if assert.Error(t, err, "REDACTED") {
					assert.Contains(t, err.Error(), tc.summary, "REDACTED")
				}
			} else {
				assert.NoError(t, err, "REDACTED")
			}

			if numberAllEndorsements {
				err = valueCollection.ValidateEndorseRapidAllEndorsements(ledgerUID, ledgerUID, level, endorse)
			} else {
				err = valueCollection.ValidateEndorseRapid(ledgerUID, ledgerUID, level, endorse)
			}
			if tc.expirationErr {
				if assert.Error(t, err, "REDACTED") {
					assert.Contains(t, err.Error(), tc.summary, "REDACTED")
				}
			} else {
				assert.NoError(t, err, "REDACTED")
			}

			//
			expirationErr := tc.expirationErr
			if (!numberAllEndorsements && sumBallots != tc.valueVolume) || sumBallots < tc.valueVolume || !tc.ledgerUID.Matches(ledgerUID) || tc.level != level {
				expirationErr = false
			}
			if numberAllEndorsements {
				err = valueCollection.ValidateEndorseRapidValidatingAllEndorsements(ledgerUID, endorse, validateLayer)
			} else {
				err = valueCollection.ValidateEndorseRapidValidating(ledgerUID, endorse, validateLayer)
			}
			if expirationErr {
				if assert.Error(t, err, "REDACTED") {
					errStr := tc.summary2
					if len(errStr) == 0 {
						errStr = tc.summary
					}
					assert.Contains(t, err.Error(), errStr, "REDACTED")
				}
			} else {
				assert.NoError(t, err, "REDACTED")
			}
		}
		t.Run(tc.summary+"REDACTED"+strconv.FormatBool(numberAllEndorsements), f)
		numberAllEndorsements = true
		t.Run(tc.summary+"REDACTED"+strconv.FormatBool(numberAllEndorsements), f)
	}
}

func Verifyratifierset_Validateendorse_Validateallautographs(t *testing.T) {
	var (
		ledgerUID = "REDACTED"
		h       = int64(3)
		ledgerUID = createLedgerUIDArbitrary()
	)

	ballotCollection, valueCollection, values := randomBallotCollection(h, 0, engineproto.PreendorseKind, 4, 10, false)
	extensionEndorse, err := CreateExtensionEndorse(ledgerUID, h, 0, ballotCollection, values, time.Now(), false)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()
	require.NoError(t, valueCollection.ValidateEndorse(ledgerUID, ledgerUID, h, endorse))

	//
	ballot := ballotCollection.FetchByOrdinal(3)
	v := ballot.ToSchema()
	err = values[3].AttestBallot("REDACTED", v)
	require.NoError(t, err)
	ballot.Autograph = v.Autograph
	ballot.AdditionAutograph = v.AdditionAutograph
	endorse.Endorsements[3] = ballot.EndorseSignature()

	err = valueCollection.ValidateEndorse(ledgerUID, ledgerUID, h, endorse)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "REDACTED")
	}
}

func Verifyratifierset_Validateendorsefast_Returnsassoonasmajormagnitudesignediffnotallsigs(t *testing.T) {
	var (
		ledgerUID = "REDACTED"
		h       = int64(3)
		ledgerUID = createLedgerUIDArbitrary()
	)

	ballotCollection, valueCollection, values := randomBallotCollection(h, 0, engineproto.PreendorseKind, 4, 10, false)
	extensionEndorse, err := CreateExtensionEndorse(ledgerUID, h, 0, ballotCollection, values, time.Now(), false)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()
	require.NoError(t, valueCollection.ValidateEndorse(ledgerUID, ledgerUID, h, endorse))

	err = valueCollection.ValidateEndorseRapidAllEndorsements(ledgerUID, ledgerUID, h, endorse)
	assert.NoError(t, err)

	//
	ballot := ballotCollection.FetchByOrdinal(3)
	v := ballot.ToSchema()
	err = values[3].AttestBallot("REDACTED", v)
	require.NoError(t, err)
	ballot.Autograph = v.Autograph
	ballot.AdditionAutograph = v.AdditionAutograph
	endorse.Endorsements[3] = ballot.EndorseSignature()

	err = valueCollection.ValidateEndorseRapid(ledgerUID, ledgerUID, h, endorse)
	assert.NoError(t, err)
	err = valueCollection.ValidateEndorseRapidAllEndorsements(ledgerUID, ledgerUID, h, endorse)
	assert.Error(t, err) //
}

func Verifyratifierset_Validateendorsefasttrusting_Returnsassoonastrustmagnitudesignediffnotallsigs(t *testing.T) {
	var (
		ledgerUID = "REDACTED"
		h       = int64(3)
		ledgerUID = createLedgerUIDArbitrary()
	)

	ballotCollection, valueCollection, values := randomBallotCollection(h, 0, engineproto.PreendorseKind, 4, 10, false)
	extensionEndorse, err := CreateExtensionEndorse(ledgerUID, h, 0, ballotCollection, values, time.Now(), false)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()
	require.NoError(t, valueCollection.ValidateEndorse(ledgerUID, ledgerUID, h, endorse))

	err = valueCollection.ValidateEndorseRapidValidatingAllEndorsements(
		ledgerUID,
		endorse,
		cometmath.Portion{Dividend: 1, Divisor: 3},
	)
	assert.NoError(t, err)

	//
	ballot := ballotCollection.FetchByOrdinal(2)
	v := ballot.ToSchema()
	err = values[2].AttestBallot("REDACTED", v)
	require.NoError(t, err)
	ballot.Autograph = v.Autograph
	ballot.AdditionAutograph = v.AdditionAutograph
	endorse.Endorsements[2] = ballot.EndorseSignature()

	err = valueCollection.ValidateEndorseRapidValidating(ledgerUID, endorse, cometmath.Portion{Dividend: 1, Divisor: 3})
	assert.NoError(t, err)
	err = valueCollection.ValidateEndorseRapidValidatingAllEndorsements(
		ledgerUID,
		endorse,
		cometmath.Portion{Dividend: 1, Divisor: 3},
	)
	assert.Error(t, err) //
}

func Verifyratifierset_Validateendorsefasttrusting(t *testing.T) {
	var (
		ledgerUID                       = createLedgerUIDArbitrary()
		ballotCollection, sourceRatifierset, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 6, 1, false)
		extensionEndorse, err                = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, time.Now(), false)
		newValueCollection, _                  = RandomRatifierCollection(2, 1)
	)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	verifyScenarios := []struct {
		valueCollection *RatifierAssign
		err    bool
	}{
		//
		0: {
			valueCollection: sourceRatifierset,
			err:    false,
		},
		//
		1: {
			valueCollection: newValueCollection,
			err:    true,
		},
		//
		2: {
			valueCollection: NewRatifierCollection(append(newValueCollection.Ratifiers, sourceRatifierset.Ratifiers...)),
			err:    false,
		},
	}

	for _, tc := range verifyScenarios {
		err = tc.valueCollection.ValidateEndorseRapidValidating("REDACTED", endorse,
			cometmath.Portion{Dividend: 1, Divisor: 3})
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func Verifyratifierset_Validateendorsefasttrustingwithrepository_Updatesrepository(t *testing.T) {
	var (
		ledgerUID                       = createLedgerUIDArbitrary()
		ballotCollection, sourceRatifierset, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 6, 1, false)
		extensionEndorse, err                = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, engineclock.Now(), false)
		newValueCollection, _                  = RandomRatifierCollection(2, 1)
	)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	valueCollection := NewRatifierCollection(append(sourceRatifierset.Ratifiers, newValueCollection.Ratifiers...))
	repository := NewAutographRepository()
	err = valueCollection.ValidateEndorseRapidValidatingWithRepository("REDACTED", endorse, cometmath.Portion{Dividend: 1, Divisor: 3}, repository)
	require.NoError(t, err)
	require.Equal(t, 3, repository.Len()) //

	repositoryValue, ok := repository.Get(string(endorse.Endorsements[0].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[0].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 0), repositoryValue.BallotAttestOctets)

	repositoryValue, ok = repository.Get(string(endorse.Endorsements[1].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[1].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 1), repositoryValue.BallotAttestOctets)

	repositoryValue, ok = repository.Get(string(endorse.Endorsements[2].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[2].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 2), repositoryValue.BallotAttestOctets)
}

func Verifyratifierset_Validateendorsefasttrustingwithrepository_Usesrepository(t *testing.T) {
	var (
		ledgerUID                       = createLedgerUIDArbitrary()
		ballotCollection, sourceRatifierset, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 6, 1, false)
		extensionEndorse, err                = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, engineclock.Now(), false)
		newValueCollection, _                  = RandomRatifierCollection(2, 1)
	)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	valueCollection := NewRatifierCollection(append(newValueCollection.Ratifiers, sourceRatifierset.Ratifiers...))

	repository := NewAutographRepository()
	repository.Add(string(endorse.Endorsements[0].Autograph), AutographRepositoryItem{
		RatifierLocation: valueCollection.Ratifiers[0].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	repository.Add(string(endorse.Endorsements[1].Autograph), AutographRepositoryItem{
		RatifierLocation: valueCollection.Ratifiers[1].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	repository.Add(string(endorse.Endorsements[2].Autograph), AutographRepositoryItem{
		RatifierLocation: valueCollection.Ratifiers[2].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})

	err = valueCollection.ValidateEndorseRapidValidatingWithRepository("REDACTED", endorse, cometmath.Portion{Dividend: 1, Divisor: 3}, repository)
	require.NoError(t, err)
	require.Equal(t, 3, repository.Len()) //
}

func Verifyratifierset_Validateendorsefastwithrepository_Updatesrepository(t *testing.T) {
	var (
		ledgerUID                       = createLedgerUIDArbitrary()
		ballotCollection, sourceRatifierset, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 6, 1, false)
		extensionEndorse, err                = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, engineclock.Now(), false)
	)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	repository := NewAutographRepository()
	err = sourceRatifierset.ValidateEndorseRapidWithRepository("REDACTED", ledgerUID, 1, endorse, repository)
	require.NoError(t, err)

	require.Equal(t, 5, repository.Len()) //

	repositoryValue, ok := repository.Get(string(endorse.Endorsements[0].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[0].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 0), repositoryValue.BallotAttestOctets)

	repositoryValue, ok = repository.Get(string(endorse.Endorsements[1].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[1].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 1), repositoryValue.BallotAttestOctets)

	repositoryValue, ok = repository.Get(string(endorse.Endorsements[2].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[2].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 2), repositoryValue.BallotAttestOctets)

	repositoryValue, ok = repository.Get(string(endorse.Endorsements[3].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[3].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 3), repositoryValue.BallotAttestOctets)

	repositoryValue, ok = repository.Get(string(endorse.Endorsements[4].Autograph))
	require.True(t, ok)
	require.Equal(t, sourceRatifierset.Ratifiers[4].PublicKey.Location().Octets(), repositoryValue.RatifierLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 4), repositoryValue.BallotAttestOctets)
}

func Verifyratifierset_Validateendorsefastwithrepository_Usesrepository(t *testing.T) {
	var (
		ledgerUID                       = createLedgerUIDArbitrary()
		ballotCollection, sourceRatifierset, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 6, 1, false)
		extensionEndorse, err                = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, engineclock.Now(), false)
	)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	repository := NewAutographRepository()
	repository.Add(string(endorse.Endorsements[0].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[0].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	repository.Add(string(endorse.Endorsements[1].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[1].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	repository.Add(string(endorse.Endorsements[2].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[2].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})
	repository.Add(string(endorse.Endorsements[3].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[3].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 3),
	})
	repository.Add(string(endorse.Endorsements[4].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[4].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 4),
	})

	err = sourceRatifierset.ValidateEndorseRapidWithRepository("REDACTED", ledgerUID, 1, endorse, repository)
	require.NoError(t, err)
	require.Equal(t, 5, repository.Len()) //
}

func Verifyratifierset_Validateendorsefasttrustingerrorsonoverflow(t *testing.T) {
	var (
		ledgerUID               = createLedgerUIDArbitrary()
		ballotCollection, valueCollection, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 1, MaximumSumPollingEnergy, false)
		extensionEndorse, err        = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, time.Now(), false)
	)
	require.NoError(t, err)

	err = valueCollection.ValidateEndorseRapidValidating("REDACTED", extensionEndorse.ToEndorse(),
		cometmath.Portion{Dividend: 25, Divisor: 55})
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "REDACTED")
	}
}

func Testverification_validateendorsegroup_Usesrepository(t *testing.T) {
	var (
		ledgerUID                       = createLedgerUIDArbitrary()
		ballotCollection, sourceRatifierset, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 6, 1, false)
		extensionEndorse, err                = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, engineclock.Now(), false)
	)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	repository := NewAutographRepository()
	repository.Add(string(endorse.Endorsements[0].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[0].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	repository.Add(string(endorse.Endorsements[1].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[1].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	repository.Add(string(endorse.Endorsements[2].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[2].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})
	repository.Add(string(endorse.Endorsements[3].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[3].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 3),
	})
	repository.Add(string(endorse.Endorsements[4].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[4].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 4),
	})

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUIDMark != LedgerUIDMarkEndorse }

	//
	tally := func(_ EndorseSignature) bool { return true }

	bv := cryptomocks.NewGroupValidator(t)

	err = validateEndorseGroup("REDACTED", sourceRatifierset, endorse, 4, bypass, tally, false, true, bv, repository)
	require.NoError(t, err)
	bv.AssertNotCalled(t, "REDACTED")
	bv.AssertNotCalled(t, "REDACTED")
}

func Testverification_validateendorseone_Usesrepository(t *testing.T) {
	var (
		ledgerUID                       = createLedgerUIDArbitrary()
		ballotCollection, sourceRatifierset, values = randomBallotCollection(1, 1, engineproto.PreendorseKind, 6, 1, false)
		extensionEndorse, err                = CreateExtensionEndorse(ledgerUID, 1, 1, ballotCollection, values, engineclock.Now(), false)
	)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	repository := NewAutographRepository()
	repository.Add(string(endorse.Endorsements[0].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[0].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	repository.Add(string(endorse.Endorsements[1].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[1].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	repository.Add(string(endorse.Endorsements[2].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[2].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})
	repository.Add(string(endorse.Endorsements[3].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[3].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 3),
	})
	repository.Add(string(endorse.Endorsements[4].Autograph), AutographRepositoryItem{
		RatifierLocation: sourceRatifierset.Ratifiers[4].PublicKey.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 4),
	})

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUIDMark != LedgerUIDMarkEndorse }

	//
	tally := func(_ EndorseSignature) bool { return true }

	emulateValueAuthkeys := []*cryptomocks.PublicKey{
		cryptomocks.NewPublicKey(t),
		cryptomocks.NewPublicKey(t),
		cryptomocks.NewPublicKey(t),
		cryptomocks.NewPublicKey(t),
		cryptomocks.NewPublicKey(t),
	}

	emulateValueAuthkeys[0].On("REDACTED").Return(sourceRatifierset.Ratifiers[0].PublicKey.Location())
	emulateValueAuthkeys[1].On("REDACTED").Return(sourceRatifierset.Ratifiers[1].PublicKey.Location())
	emulateValueAuthkeys[2].On("REDACTED").Return(sourceRatifierset.Ratifiers[2].PublicKey.Location())
	emulateValueAuthkeys[3].On("REDACTED").Return(sourceRatifierset.Ratifiers[3].PublicKey.Location())
	emulateValueAuthkeys[4].On("REDACTED").Return(sourceRatifierset.Ratifiers[4].PublicKey.Location())

	sourceRatifierset.Ratifiers[0].PublicKey = emulateValueAuthkeys[0]
	sourceRatifierset.Ratifiers[1].PublicKey = emulateValueAuthkeys[1]
	sourceRatifierset.Ratifiers[2].PublicKey = emulateValueAuthkeys[2]
	sourceRatifierset.Ratifiers[3].PublicKey = emulateValueAuthkeys[3]
	sourceRatifierset.Ratifiers[4].PublicKey = emulateValueAuthkeys[4]

	err = validateEndorseUnique("REDACTED", sourceRatifierset, endorse, 4, bypass, tally, false, true, repository)
	require.NoError(t, err)

	emulateValueAuthkeys[0].AssertCalled(t, "REDACTED")
	emulateValueAuthkeys[1].AssertCalled(t, "REDACTED")
	emulateValueAuthkeys[2].AssertCalled(t, "REDACTED")
	emulateValueAuthkeys[3].AssertCalled(t, "REDACTED")
	emulateValueAuthkeys[4].AssertCalled(t, "REDACTED")

	emulateValueAuthkeys[0].AssertNotCalled(t, "REDACTED")
	emulateValueAuthkeys[1].AssertNotCalled(t, "REDACTED")
	emulateValueAuthkeys[2].AssertNotCalled(t, "REDACTED")
	emulateValueAuthkeys[3].AssertNotCalled(t, "REDACTED")
	emulateValueAuthkeys[4].AssertNotCalled(t, "REDACTED")
}
