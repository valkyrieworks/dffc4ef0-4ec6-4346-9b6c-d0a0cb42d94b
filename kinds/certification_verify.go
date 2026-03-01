package kinds

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cryptosimulate "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/simulations"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

//
//
func Testvalidset_Verifyendorsement_Every(t *testing.T) {
	var (
		iteration  = int32(0)
		altitude = int64(100)

		ledgerUUID    = createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
		successionUUID    = "REDACTED"
		relianceStratum = strongarithmetic.Portion{Dividend: 2, Divisor: 3}
	)

	verifyScenarios := []struct {
		definition, definition2 string //
		//
		successionUUID string
		//
		ledgerUUID LedgerUUID
		itemExtent int

		//
		altitude int64

		//
		ledgerBallots  int
		voidBallots    int
		missingBallots int

		expirationFault bool
	}{
		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 3, altitude, 3, 0, 0, false},
		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 1, altitude, 1, 0, 0, false},

		{"REDACTED", "REDACTED", "REDACTED", ledgerUUID, 2, altitude, 2, 0, 0, true},
		{"REDACTED", "REDACTED", successionUUID, createLedgerUUIDUnpredictable(), 2, altitude, 2, 0, 0, true},
		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 1, altitude - 1, 1, 0, 0, true},

		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 4, altitude, 3, 0, 0, true},
		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 1, altitude, 2, 0, 0, true},

		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 10, altitude, 3, 2, 5, true},
		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 1, altitude, 0, 0, 1, true}, //
		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 1, altitude, 0, 1, 0, true}, //
		{"REDACTED", "REDACTED", successionUUID, ledgerUUID, 9, altitude, 6, 3, 0, true},
	}

	for _, tc := range verifyScenarios {
		tallyEveryNotations := false
		f := func(t *testing.T) {
			_, itemAssign, values := arbitraryBallotAssign(tc.altitude, iteration, commitchema.PreendorseKind, tc.itemExtent, 10, false)
			sumBallots := tc.ledgerBallots + tc.missingBallots + tc.voidBallots
			signatures := make([]EndorseSignature, sumBallots)
			vi := 0
			//
			for i := 0; i < tc.missingBallots; i++ {
				signatures[vi] = FreshEndorseSignatureMissing()
				vi++
			}
			for i := 0; i < tc.ledgerBallots+tc.voidBallots; i++ {

				publicToken, err := values[vi%len(values)].ObtainPublicToken()
				require.NoError(t, err)
				ballot := &Ballot{
					AssessorLocation: publicToken.Location(),
					AssessorOrdinal:   int32(vi),
					Altitude:           tc.altitude,
					Iteration:            iteration,
					Kind:             commitchema.PreendorseKind,
					LedgerUUID:          tc.ledgerUUID,
					Timestamp:        time.Now(),
				}
				if i >= tc.ledgerBallots {
					ballot.LedgerUUID = LedgerUUID{}
				}

				v := ballot.TowardSchema()

				require.NoError(t, values[vi%len(values)].AttestBallot(tc.successionUUID, v))
				ballot.Notation = v.Notation

				signatures[vi] = ballot.EndorseSignature()

				vi++
			}
			endorse := &Endorse{
				Altitude:     tc.altitude,
				Iteration:      iteration,
				LedgerUUID:    tc.ledgerUUID,
				Notations: signatures,
			}

			err := itemAssign.ValidateEndorse(successionUUID, ledgerUUID, altitude, endorse)
			if tc.expirationFault {
				if assert.Error(t, err, "REDACTED") {
					assert.Contains(t, err.Error(), tc.definition, "REDACTED")
				}
			} else {
				assert.NoError(t, err, "REDACTED")
			}

			if tallyEveryNotations {
				err = itemAssign.ValidateEndorseAgileEveryNotations(successionUUID, ledgerUUID, altitude, endorse)
			} else {
				err = itemAssign.ValidateEndorseAgile(successionUUID, ledgerUUID, altitude, endorse)
			}
			if tc.expirationFault {
				if assert.Error(t, err, "REDACTED") {
					assert.Contains(t, err.Error(), tc.definition, "REDACTED")
				}
			} else {
				assert.NoError(t, err, "REDACTED")
			}

			//
			expirationFault := tc.expirationFault
			if (!tallyEveryNotations && sumBallots != tc.itemExtent) || sumBallots < tc.itemExtent || !tc.ledgerUUID.Matches(ledgerUUID) || tc.altitude != altitude {
				expirationFault = false
			}
			if tallyEveryNotations {
				err = itemAssign.ValidateEndorseAgileRelyingEveryNotations(successionUUID, endorse, relianceStratum)
			} else {
				err = itemAssign.ValidateEndorseAgileRelying(successionUUID, endorse, relianceStratum)
			}
			if expirationFault {
				if assert.Error(t, err, "REDACTED") {
					faultTxt := tc.definition2
					if len(faultTxt) == 0 {
						faultTxt = tc.definition
					}
					assert.Contains(t, err.Error(), faultTxt, "REDACTED")
				}
			} else {
				assert.NoError(t, err, "REDACTED")
			}
		}
		t.Run(tc.definition+"REDACTED"+strconv.FormatBool(tallyEveryNotations), f)
		tallyEveryNotations = true
		t.Run(tc.definition+"REDACTED"+strconv.FormatBool(tallyEveryNotations), f)
	}
}

func Testvalidset_Verifyendorsement_Verifyallequivalents(t *testing.T) {
	var (
		successionUUID = "REDACTED"
		h       = int64(3)
		ledgerUUID = createLedgerUUIDUnpredictable()
	)

	ballotAssign, itemAssign, values := arbitraryBallotAssign(h, 0, commitchema.PreendorseKind, 4, 10, false)
	addnEndorse, err := CreateAddnEndorse(ledgerUUID, h, 0, ballotAssign, values, time.Now(), false)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()
	require.NoError(t, itemAssign.ValidateEndorse(successionUUID, ledgerUUID, h, endorse))

	//
	ballot := ballotAssign.ObtainViaOrdinal(3)
	v := ballot.TowardSchema()
	err = values[3].AttestBallot("REDACTED", v)
	require.NoError(t, err)
	ballot.Notation = v.Notation
	ballot.AdditionNotation = v.AdditionNotation
	endorse.Notations[3] = ballot.EndorseSignature()

	err = itemAssign.ValidateEndorse(successionUUID, ledgerUUID, h, endorse)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "REDACTED")
	}
}

func Testvalidset_Verifyendorsementlite_Returnsimmediatelyifmajorityballotingpowernotallequivalents(t *testing.T) {
	var (
		successionUUID = "REDACTED"
		h       = int64(3)
		ledgerUUID = createLedgerUUIDUnpredictable()
	)

	ballotAssign, itemAssign, values := arbitraryBallotAssign(h, 0, commitchema.PreendorseKind, 4, 10, false)
	addnEndorse, err := CreateAddnEndorse(ledgerUUID, h, 0, ballotAssign, values, time.Now(), false)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()
	require.NoError(t, itemAssign.ValidateEndorse(successionUUID, ledgerUUID, h, endorse))

	err = itemAssign.ValidateEndorseAgileEveryNotations(successionUUID, ledgerUUID, h, endorse)
	assert.NoError(t, err)

	//
	ballot := ballotAssign.ObtainViaOrdinal(3)
	v := ballot.TowardSchema()
	err = values[3].AttestBallot("REDACTED", v)
	require.NoError(t, err)
	ballot.Notation = v.Notation
	ballot.AdditionNotation = v.AdditionNotation
	endorse.Notations[3] = ballot.EndorseSignature()

	err = itemAssign.ValidateEndorseAgile(successionUUID, ledgerUUID, h, endorse)
	assert.NoError(t, err)
	err = itemAssign.ValidateEndorseAgileEveryNotations(successionUUID, ledgerUUID, h, endorse)
	assert.Error(t, err) //
}

func Testvalidset_Verifyendorsementlitetrusting_Returnsimmediatelyiftrustlevelnotallequivalents(t *testing.T) {
	var (
		successionUUID = "REDACTED"
		h       = int64(3)
		ledgerUUID = createLedgerUUIDUnpredictable()
	)

	ballotAssign, itemAssign, values := arbitraryBallotAssign(h, 0, commitchema.PreendorseKind, 4, 10, false)
	addnEndorse, err := CreateAddnEndorse(ledgerUUID, h, 0, ballotAssign, values, time.Now(), false)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()
	require.NoError(t, itemAssign.ValidateEndorse(successionUUID, ledgerUUID, h, endorse))

	err = itemAssign.ValidateEndorseAgileRelyingEveryNotations(
		successionUUID,
		endorse,
		strongarithmetic.Portion{Dividend: 1, Divisor: 3},
	)
	assert.NoError(t, err)

	//
	ballot := ballotAssign.ObtainViaOrdinal(2)
	v := ballot.TowardSchema()
	err = values[2].AttestBallot("REDACTED", v)
	require.NoError(t, err)
	ballot.Notation = v.Notation
	ballot.AdditionNotation = v.AdditionNotation
	endorse.Notations[2] = ballot.EndorseSignature()

	err = itemAssign.ValidateEndorseAgileRelying(successionUUID, endorse, strongarithmetic.Portion{Dividend: 1, Divisor: 3})
	assert.NoError(t, err)
	err = itemAssign.ValidateEndorseAgileRelyingEveryNotations(
		successionUUID,
		endorse,
		strongarithmetic.Portion{Dividend: 1, Divisor: 3},
	)
	assert.Error(t, err) //
}

func Testvalidset_Verifyendorsementlitetrusting(t *testing.T) {
	var (
		ledgerUUID                       = createLedgerUUIDUnpredictable()
		ballotAssign, initialValidset, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 6, 1, false)
		addnEndorse, err                = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, time.Now(), false)
		freshItemAssign, _                  = ArbitraryAssessorAssign(2, 1)
	)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	verifyScenarios := []struct {
		itemAssign *AssessorAssign
		err    bool
	}{
		//
		0: {
			itemAssign: initialValidset,
			err:    false,
		},
		//
		1: {
			itemAssign: freshItemAssign,
			err:    true,
		},
		//
		2: {
			itemAssign: FreshAssessorAssign(append(freshItemAssign.Assessors, initialValidset.Assessors...)),
			err:    false,
		},
	}

	for _, tc := range verifyScenarios {
		err = tc.itemAssign.ValidateEndorseAgileRelying("REDACTED", endorse,
			strongarithmetic.Portion{Dividend: 1, Divisor: 3})
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func Testvalidset_Verifyendorsementlitetrustingwithdepot_Refreshescache(t *testing.T) {
	var (
		ledgerUUID                       = createLedgerUUIDUnpredictable()
		ballotAssign, initialValidset, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 6, 1, false)
		addnEndorse, err                = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, committime.Now(), false)
		freshItemAssign, _                  = ArbitraryAssessorAssign(2, 1)
	)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	itemAssign := FreshAssessorAssign(append(initialValidset.Assessors, freshItemAssign.Assessors...))
	stash := FreshSigningStash()
	err = itemAssign.ValidateEndorseAgileRelyingUsingStash("REDACTED", endorse, strongarithmetic.Portion{Dividend: 1, Divisor: 3}, stash)
	require.NoError(t, err)
	require.Equal(t, 3, stash.Len()) //

	stashItem, ok := stash.Get(string(endorse.Notations[0].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[0].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 0), stashItem.BallotAttestOctets)

	stashItem, ok = stash.Get(string(endorse.Notations[1].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[1].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 1), stashItem.BallotAttestOctets)

	stashItem, ok = stash.Get(string(endorse.Notations[2].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[2].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 2), stashItem.BallotAttestOctets)
}

func Testvalidset_Verifyendorsementlitetrustingwithdepot_Utilizesdepot(t *testing.T) {
	var (
		ledgerUUID                       = createLedgerUUIDUnpredictable()
		ballotAssign, initialValidset, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 6, 1, false)
		addnEndorse, err                = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, committime.Now(), false)
		freshItemAssign, _                  = ArbitraryAssessorAssign(2, 1)
	)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	itemAssign := FreshAssessorAssign(append(freshItemAssign.Assessors, initialValidset.Assessors...))

	stash := FreshSigningStash()
	stash.Add(string(endorse.Notations[0].Notation), NotationStashDatum{
		AssessorLocation: itemAssign.Assessors[0].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	stash.Add(string(endorse.Notations[1].Notation), NotationStashDatum{
		AssessorLocation: itemAssign.Assessors[1].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	stash.Add(string(endorse.Notations[2].Notation), NotationStashDatum{
		AssessorLocation: itemAssign.Assessors[2].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})

	err = itemAssign.ValidateEndorseAgileRelyingUsingStash("REDACTED", endorse, strongarithmetic.Portion{Dividend: 1, Divisor: 3}, stash)
	require.NoError(t, err)
	require.Equal(t, 3, stash.Len()) //
}

func Testvalidset_Verifyendorsementlitewithdepot_Refreshescache(t *testing.T) {
	var (
		ledgerUUID                       = createLedgerUUIDUnpredictable()
		ballotAssign, initialValidset, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 6, 1, false)
		addnEndorse, err                = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, committime.Now(), false)
	)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	stash := FreshSigningStash()
	err = initialValidset.ValidateEndorseAgileUsingStash("REDACTED", ledgerUUID, 1, endorse, stash)
	require.NoError(t, err)

	require.Equal(t, 5, stash.Len()) //

	stashItem, ok := stash.Get(string(endorse.Notations[0].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[0].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 0), stashItem.BallotAttestOctets)

	stashItem, ok = stash.Get(string(endorse.Notations[1].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[1].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 1), stashItem.BallotAttestOctets)

	stashItem, ok = stash.Get(string(endorse.Notations[2].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[2].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 2), stashItem.BallotAttestOctets)

	stashItem, ok = stash.Get(string(endorse.Notations[3].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[3].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 3), stashItem.BallotAttestOctets)

	stashItem, ok = stash.Get(string(endorse.Notations[4].Notation))
	require.True(t, ok)
	require.Equal(t, initialValidset.Assessors[4].PublicToken.Location().Octets(), stashItem.AssessorLocation)
	require.Equal(t, endorse.BallotAttestOctets("REDACTED", 4), stashItem.BallotAttestOctets)
}

func Testvalidset_Verifyendorsementlitewithdepot_Utilizesdepot(t *testing.T) {
	var (
		ledgerUUID                       = createLedgerUUIDUnpredictable()
		ballotAssign, initialValidset, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 6, 1, false)
		addnEndorse, err                = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, committime.Now(), false)
	)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	stash := FreshSigningStash()
	stash.Add(string(endorse.Notations[0].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[0].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	stash.Add(string(endorse.Notations[1].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[1].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	stash.Add(string(endorse.Notations[2].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[2].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})
	stash.Add(string(endorse.Notations[3].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[3].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 3),
	})
	stash.Add(string(endorse.Notations[4].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[4].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 4),
	})

	err = initialValidset.ValidateEndorseAgileUsingStash("REDACTED", ledgerUUID, 1, endorse, stash)
	require.NoError(t, err)
	require.Equal(t, 5, stash.Len()) //
}

func Testvalidset_Verifyendorsementlitetrustingfaultsonexcess(t *testing.T) {
	var (
		ledgerUUID               = createLedgerUUIDUnpredictable()
		ballotAssign, itemAssign, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 1, MaximumSumBallotingPotency, false)
		addnEndorse, err        = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, time.Now(), false)
	)
	require.NoError(t, err)

	err = itemAssign.ValidateEndorseAgileRelying("REDACTED", addnEndorse.TowardEndorse(),
		strongarithmetic.Portion{Dividend: 25, Divisor: 55})
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "REDACTED")
	}
}

func Verifyvalidity_verifyendorsementgroup_Utilizesdepot(t *testing.T) {
	var (
		ledgerUUID                       = createLedgerUUIDUnpredictable()
		ballotAssign, initialValidset, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 6, 1, false)
		addnEndorse, err                = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, committime.Now(), false)
	)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	stash := FreshSigningStash()
	stash.Add(string(endorse.Notations[0].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[0].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	stash.Add(string(endorse.Notations[1].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[1].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	stash.Add(string(endorse.Notations[2].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[2].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})
	stash.Add(string(endorse.Notations[3].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[3].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 3),
	})
	stash.Add(string(endorse.Notations[4].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[4].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 4),
	})

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUUIDMarker != LedgerUUIDMarkerEndorse }

	//
	tally := func(_ EndorseSignature) bool { return true }

	bv := cryptosimulate.FreshClusterValidator(t)

	err = validateEndorseCluster("REDACTED", initialValidset, endorse, 4, bypass, tally, false, true, bv, stash)
	require.NoError(t, err)
	bv.AssertNotCalled(t, "REDACTED")
	bv.AssertNotCalled(t, "REDACTED")
}

func Verifyvalidity_verifyendorsementsingle_Utilizesdepot(t *testing.T) {
	var (
		ledgerUUID                       = createLedgerUUIDUnpredictable()
		ballotAssign, initialValidset, values = arbitraryBallotAssign(1, 1, commitchema.PreendorseKind, 6, 1, false)
		addnEndorse, err                = CreateAddnEndorse(ledgerUUID, 1, 1, ballotAssign, values, committime.Now(), false)
	)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	stash := FreshSigningStash()
	stash.Add(string(endorse.Notations[0].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[0].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 0),
	})
	stash.Add(string(endorse.Notations[1].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[1].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 1),
	})
	stash.Add(string(endorse.Notations[2].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[2].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 2),
	})
	stash.Add(string(endorse.Notations[3].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[3].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 3),
	})
	stash.Add(string(endorse.Notations[4].Notation), NotationStashDatum{
		AssessorLocation: initialValidset.Assessors[4].PublicToken.Location(),
		BallotAttestOctets:    endorse.BallotAttestOctets("REDACTED", 4),
	})

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUUIDMarker != LedgerUUIDMarkerEndorse }

	//
	tally := func(_ EndorseSignature) bool { return true }

	simulateItemPublictokens := []*cryptosimulate.PublicToken{
		cryptosimulate.FreshPublicToken(t),
		cryptosimulate.FreshPublicToken(t),
		cryptosimulate.FreshPublicToken(t),
		cryptosimulate.FreshPublicToken(t),
		cryptosimulate.FreshPublicToken(t),
	}

	simulateItemPublictokens[0].On("REDACTED").Return(initialValidset.Assessors[0].PublicToken.Location())
	simulateItemPublictokens[1].On("REDACTED").Return(initialValidset.Assessors[1].PublicToken.Location())
	simulateItemPublictokens[2].On("REDACTED").Return(initialValidset.Assessors[2].PublicToken.Location())
	simulateItemPublictokens[3].On("REDACTED").Return(initialValidset.Assessors[3].PublicToken.Location())
	simulateItemPublictokens[4].On("REDACTED").Return(initialValidset.Assessors[4].PublicToken.Location())

	initialValidset.Assessors[0].PublicToken = simulateItemPublictokens[0]
	initialValidset.Assessors[1].PublicToken = simulateItemPublictokens[1]
	initialValidset.Assessors[2].PublicToken = simulateItemPublictokens[2]
	initialValidset.Assessors[3].PublicToken = simulateItemPublictokens[3]
	initialValidset.Assessors[4].PublicToken = simulateItemPublictokens[4]

	err = validateEndorseUnique("REDACTED", initialValidset, endorse, 4, bypass, tally, false, true, stash)
	require.NoError(t, err)

	simulateItemPublictokens[0].AssertCalled(t, "REDACTED")
	simulateItemPublictokens[1].AssertCalled(t, "REDACTED")
	simulateItemPublictokens[2].AssertCalled(t, "REDACTED")
	simulateItemPublictokens[3].AssertCalled(t, "REDACTED")
	simulateItemPublictokens[4].AssertCalled(t, "REDACTED")

	simulateItemPublictokens[0].AssertNotCalled(t, "REDACTED")
	simulateItemPublictokens[1].AssertNotCalled(t, "REDACTED")
	simulateItemPublictokens[2].AssertNotCalled(t, "REDACTED")
	simulateItemPublictokens[3].AssertNotCalled(t, "REDACTED")
	simulateItemPublictokens[4].AssertNotCalled(t, "REDACTED")
}
