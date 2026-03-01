package kinds

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

func Assessmentvoteset_Appendvote_Valid(t *testing.T) {
	altitude, iteration := int64(1), int32(0)
	ballotAssign, _, privateAssessors := arbitraryBallotAssign(altitude, iteration, commitchema.PreballotKind, 10, 1, false)
	item0 := privateAssessors[0]

	item0p, err := item0.ObtainPublicToken()
	require.NoError(t, err)
	item0address := item0p.Location()

	assert.Nil(t, ballotAssign.ObtainViaLocation(item0address))
	assert.False(t, ballotAssign.DigitSeries().ObtainOrdinal(0))
	ledgerUUID, ok := ballotAssign.CoupleTrinityPreponderance()
	assert.False(t, ok || !ledgerUUID.EqualsNull(), "REDACTED")

	ballot := &Ballot{
		AssessorLocation: item0address,
		AssessorOrdinal:   0, //
		Altitude:           altitude,
		Iteration:            iteration,
		Kind:             commitchema.PreballotKind,
		Timestamp:        committime.Now(),
		LedgerUUID:          LedgerUUID{nil, FragmentAssignHeading{}},
	}
	_, err = attestAppendBallot(item0, ballot, ballotAssign)
	require.NoError(t, err)

	assert.NotNil(t, ballotAssign.ObtainViaLocation(item0address))
	assert.True(t, ballotAssign.DigitSeries().ObtainOrdinal(0))
	ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
	assert.False(t, ok || !ledgerUUID.EqualsNull(), "REDACTED")
}

func Assessmentvoteset_Appendvote_Flawed(t *testing.T) {
	altitude, iteration := int64(1), int32(0)
	ballotAssign, _, privateAssessors := arbitraryBallotAssign(altitude, iteration, commitchema.PreballotKind, 10, 1, false)

	ballotSchema := &Ballot{
		AssessorLocation: nil,
		AssessorOrdinal:   -1,
		Altitude:           altitude,
		Iteration:            iteration,
		Timestamp:        committime.Now(),
		Kind:             commitchema.PreballotKind,
		LedgerUUID:          LedgerUUID{nil, FragmentAssignHeading{}},
	}

	//
	{
		publicToken, err := privateAssessors[0].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 0)
		appended, err := attestAppendBallot(privateAssessors[0], ballot, ballotAssign)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicToken, err := privateAssessors[0].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 0)
		appended, err := attestAppendBallot(privateAssessors[0], usingLedgerDigest(ballot, commitrand.Octets(32)), ballotAssign)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicToken, err := privateAssessors[1].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 1)
		appended, err := attestAppendBallot(privateAssessors[1], usingAltitude(ballot, altitude+1), ballotAssign)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicToken, err := privateAssessors[2].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 2)
		appended, err := attestAppendBallot(privateAssessors[2], usingIteration(ballot, iteration+1), ballotAssign)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicToken, err := privateAssessors[3].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 3)
		appended, err := attestAppendBallot(privateAssessors[3], usingKind(ballot, byte(commitchema.PreendorseKind)), ballotAssign)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}
}

func Assessment_two_three_Major(b *testing.B) {
	altitude, iteration := int64(1), int32(0)

	ballotSchema := &Ballot{
		AssessorLocation: nil, //
		AssessorOrdinal:   -1,  //
		Altitude:           altitude,
		Iteration:            iteration,
		Kind:             commitchema.PreballotKind,
		Timestamp:        committime.Now(),
		LedgerUUID:          LedgerUUID{nil, FragmentAssignHeading{}},
	}
	ledgerFragmentsSum := uint32(123)
	ledgerFragmentAssignHeading := FragmentAssignHeading{ledgerFragmentsSum, security.CHARArbitraryOctets(32)}
	for b.Loop() {
		ballotAssign, _, privateAssessors := arbitraryBallotAssign(altitude, iteration, commitchema.PreballotKind, 100, 1, false)
		for i := int32(0); i < int32(100); i += 4 {
			publicToken, _ := privateAssessors[i].ObtainPublicToken()
			address := publicToken.Location()
			ballot := usingAssessor(ballotSchema, address, i)
			_, err := attestAppendBallot(privateAssessors[i], usingLedgerDigest(ballot, nil), ballotAssign)
			require.NoError(b, err)
			_, _ = ballotAssign.CoupleTrinityPreponderance()

			publicToken, _ = privateAssessors[i+1].ObtainPublicToken()
			address = publicToken.Location()
			ballot = usingAssessor(ballotSchema, address, i+1)
			_, err = attestAppendBallot(privateAssessors[i+1], ballot, ballotAssign)
			require.NoError(b, err)
			_, _ = ballotAssign.CoupleTrinityPreponderance()

			publicToken, _ = privateAssessors[i+2].ObtainPublicToken()
			address = publicToken.Location()
			ballot = usingAssessor(ballotSchema, address, i+2)
			ledgerFragmentsHeading := FragmentAssignHeading{ledgerFragmentsSum, security.CHARArbitraryOctets(32)}
			_, err = attestAppendBallot(privateAssessors[i+2], usingLedgerFragmentAssignHeading(ballot, ledgerFragmentsHeading), ballotAssign)
			require.NoError(b, err)
			_, _ = ballotAssign.CoupleTrinityPreponderance()

			publicToken, _ = privateAssessors[i+3].ObtainPublicToken()
			address = publicToken.Location()
			ballot = usingAssessor(ballotSchema, address, i+3)
			ledgerFragmentsHeading = FragmentAssignHeading{ledgerFragmentsSum + 1, ledgerFragmentAssignHeading.Digest}
			_, err = attestAppendBallot(privateAssessors[i+3], usingLedgerFragmentAssignHeading(ballot, ledgerFragmentsHeading), ballotAssign)
			require.NoError(b, err)
			_, _ = ballotAssign.CoupleTrinityPreponderance()
		}
	}
}

func Assessmentvoteset_two_threemajority(t *testing.T) {
	altitude, iteration := int64(1), int32(0)
	ballotAssign, _, privateAssessors := arbitraryBallotAssign(altitude, iteration, commitchema.PreballotKind, 10, 1, false)

	ballotSchema := &Ballot{
		AssessorLocation: nil, //
		AssessorOrdinal:   -1,  //
		Altitude:           altitude,
		Iteration:            iteration,
		Kind:             commitchema.PreballotKind,
		Timestamp:        committime.Now(),
		LedgerUUID:          LedgerUUID{nil, FragmentAssignHeading{}},
	}
	//
	for i := int32(0); i < 6; i++ {
		publicToken, err := privateAssessors[i].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, i)
		_, err = attestAppendBallot(privateAssessors[i], ballot, ballotAssign)
		require.NoError(t, err)
	}
	ledgerUUID, ok := ballotAssign.CoupleTrinityPreponderance()
	assert.False(t, ok || !ledgerUUID.EqualsNull(), "REDACTED")

	//
	{
		publicToken, err := privateAssessors[6].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 6)
		_, err = attestAppendBallot(privateAssessors[6], usingLedgerDigest(ballot, commitrand.Octets(32)), ballotAssign)
		require.NoError(t, err)
		ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
		assert.False(t, ok || !ledgerUUID.EqualsNull(), "REDACTED")
	}

	//
	{
		publicToken, err := privateAssessors[7].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 7)
		_, err = attestAppendBallot(privateAssessors[7], ballot, ballotAssign)
		require.NoError(t, err)
		ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
		assert.True(t, ok || ledgerUUID.EqualsNull(), "REDACTED")
	}
}

func Assessmentvoteset_two_threemajorityreduced(t *testing.T) {
	altitude, iteration := int64(1), int32(0)
	ballotAssign, _, privateAssessors := arbitraryBallotAssign(altitude, iteration, commitchema.PreballotKind, 100, 1, false)

	ledgerDigest := security.CHARArbitraryOctets(32)
	ledgerFragmentsSum := uint32(123)
	ledgerFragmentAssignHeading := FragmentAssignHeading{ledgerFragmentsSum, security.CHARArbitraryOctets(32)}

	ballotSchema := &Ballot{
		AssessorLocation: nil, //
		AssessorOrdinal:   -1,  //
		Altitude:           altitude,
		Iteration:            iteration,
		Timestamp:        committime.Now(),
		Kind:             commitchema.PreballotKind,
		LedgerUUID:          LedgerUUID{ledgerDigest, ledgerFragmentAssignHeading},
	}

	//
	for i := int32(0); i < 66; i++ {
		publicToken, err := privateAssessors[i].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, i)
		_, err = attestAppendBallot(privateAssessors[i], ballot, ballotAssign)
		require.NoError(t, err)
	}
	ledgerUUID, ok := ballotAssign.CoupleTrinityPreponderance()
	assert.False(t, ok || !ledgerUUID.EqualsNull(),
		"REDACTED")

	//
	{
		publicToken, err := privateAssessors[66].ObtainPublicToken()
		require.NoError(t, err)
		address := publicToken.Location()
		ballot := usingAssessor(ballotSchema, address, 66)
		_, err = attestAppendBallot(privateAssessors[66], usingLedgerDigest(ballot, nil), ballotAssign)
		require.NoError(t, err)
		ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
		assert.False(t, ok || !ledgerUUID.EqualsNull(),
			"REDACTED")
	}

	//
	{
		publicToken, err := privateAssessors[67].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 67)
		ledgerFragmentsHeading := FragmentAssignHeading{ledgerFragmentsSum, security.CHARArbitraryOctets(32)}
		_, err = attestAppendBallot(privateAssessors[67], usingLedgerFragmentAssignHeading(ballot, ledgerFragmentsHeading), ballotAssign)
		require.NoError(t, err)
		ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
		assert.False(t, ok || !ledgerUUID.EqualsNull(),
			"REDACTED")
	}

	//
	{
		publicToken, err := privateAssessors[68].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 68)
		ledgerFragmentsHeading := FragmentAssignHeading{ledgerFragmentsSum + 1, ledgerFragmentAssignHeading.Digest}
		_, err = attestAppendBallot(privateAssessors[68], usingLedgerFragmentAssignHeading(ballot, ledgerFragmentsHeading), ballotAssign)
		require.NoError(t, err)
		ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
		assert.False(t, ok || !ledgerUUID.EqualsNull(),
			"REDACTED")
	}

	//
	{
		publicToken, err := privateAssessors[69].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 69)
		_, err = attestAppendBallot(privateAssessors[69], usingLedgerDigest(ballot, commitrand.Octets(32)), ballotAssign)
		require.NoError(t, err)
		ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
		assert.False(t, ok || !ledgerUUID.EqualsNull(),
			"REDACTED")
	}

	//
	{
		publicToken, err := privateAssessors[70].ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		ballot := usingAssessor(ballotSchema, location, 70)
		_, err = attestAppendBallot(privateAssessors[70], ballot, ballotAssign)
		require.NoError(t, err)
		ledgerUUID, ok = ballotAssign.CoupleTrinityPreponderance()
		assert.True(t, ok && ledgerUUID.Matches(LedgerUUID{ledgerDigest, ledgerFragmentAssignHeading}),
			"REDACTED")
	}
}

func Assessmentvoteset_Disagreements(t *testing.T) {
	altitude, iteration := int64(1), int32(0)
	ballotAssign, _, privateAssessors := arbitraryBallotAssign(altitude, iteration, commitchema.PreballotKind, 4, 1, false)
	ledgerDigest1 := commitrand.Octets(32)
	ledgerDigest2 := commitrand.Octets(32)

	ballotSchema := &Ballot{
		AssessorLocation: nil,
		AssessorOrdinal:   -1,
		Altitude:           altitude,
		Iteration:            iteration,
		Timestamp:        committime.Now(),
		Kind:             commitchema.PreballotKind,
		LedgerUUID:          LedgerUUID{nil, FragmentAssignHeading{}},
	}

	item0, err := privateAssessors[0].ObtainPublicToken()
	require.NoError(t, err)
	item0address := item0.Location()

	//
	{
		ballot := usingAssessor(ballotSchema, item0address, 0)
		appended, err := attestAppendBallot(privateAssessors[0], ballot, ballotAssign)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		ballot := usingAssessor(ballotSchema, item0address, 0)
		appended, err := attestAppendBallot(privateAssessors[0], usingLedgerDigest(ballot, ledgerDigest1), ballotAssign)
		assert.False(t, appended, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}

	//
	err = ballotAssign.AssignNodeMajor23("REDACTED", LedgerUUID{ledgerDigest1, FragmentAssignHeading{}})
	require.NoError(t, err)

	//
	{
		ballot := usingAssessor(ballotSchema, item0address, 0)
		appended, err := attestAppendBallot(privateAssessors[0], usingLedgerDigest(ballot, ledgerDigest1), ballotAssign)
		assert.True(t, appended, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}

	//
	err = ballotAssign.AssignNodeMajor23("REDACTED", LedgerUUID{ledgerDigest2, FragmentAssignHeading{}})
	require.Error(t, err)

	//
	{
		ballot := usingAssessor(ballotSchema, item0address, 0)
		appended, err := attestAppendBallot(privateAssessors[0], usingLedgerDigest(ballot, ledgerDigest2), ballotAssign)
		assert.False(t, appended, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}

	//
	{
		pv, err := privateAssessors[1].ObtainPublicToken()
		assert.NoError(t, err)
		location := pv.Location()
		ballot := usingAssessor(ballotSchema, location, 1)
		appended, err := attestAppendBallot(privateAssessors[1], usingLedgerDigest(ballot, ledgerDigest1), ballotAssign)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	if ballotAssign.OwnsCoupleTrinityPreponderance() {
		t.Errorf("REDACTED")
	}
	if ballotAssign.OwnsCoupleTrinitySome() {
		t.Errorf("REDACTED")
	}

	//
	{
		pv, err := privateAssessors[2].ObtainPublicToken()
		assert.NoError(t, err)
		location := pv.Location()
		ballot := usingAssessor(ballotSchema, location, 2)
		appended, err := attestAppendBallot(privateAssessors[2], usingLedgerDigest(ballot, ledgerDigest2), ballotAssign)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	if ballotAssign.OwnsCoupleTrinityPreponderance() {
		t.Errorf("REDACTED")
	}
	if !ballotAssign.OwnsCoupleTrinitySome() {
		t.Errorf("REDACTED")
	}

	//
	err = ballotAssign.AssignNodeMajor23("REDACTED", LedgerUUID{ledgerDigest1, FragmentAssignHeading{}})
	require.NoError(t, err)

	//
	{
		pv, err := privateAssessors[2].ObtainPublicToken()
		assert.NoError(t, err)
		location := pv.Location()
		ballot := usingAssessor(ballotSchema, location, 2)
		appended, err := attestAppendBallot(privateAssessors[2], usingLedgerDigest(ballot, ledgerDigest1), ballotAssign)
		assert.True(t, appended)
		assert.Error(t, err, "REDACTED")
	}

	//
	if !ballotAssign.OwnsCoupleTrinityPreponderance() {
		t.Errorf("REDACTED")
	}
	ledgerUUIDMajor23, _ := ballotAssign.CoupleTrinityPreponderance()
	if !bytes.Equal(ledgerUUIDMajor23.Digest, ledgerDigest1) {
		t.Errorf("REDACTED")
	}
	if !ballotAssign.OwnsCoupleTrinitySome() {
		t.Errorf("REDACTED")
	}
}

func Assessmentvoteset_Formcommit(t *testing.T) {
	altitude, iteration := int64(1), int32(0)
	ballotAssign, _, privateAssessors := arbitraryBallotAssign(altitude, iteration, commitchema.PreendorseKind, 10, 1, true)
	ledgerDigest, ledgerFragmentAssignHeading := security.CHARArbitraryOctets(32), FragmentAssignHeading{123, security.CHARArbitraryOctets(32)}

	ballotSchema := &Ballot{
		AssessorLocation: nil,
		AssessorOrdinal:   -1,
		Altitude:           altitude,
		Iteration:            iteration,
		Timestamp:        committime.Now(),
		Kind:             commitchema.PreendorseKind,
		LedgerUUID:          LedgerUUID{ledgerDigest, ledgerFragmentAssignHeading},
	}

	//
	for i := int32(0); i < 6; i++ {
		pv, err := privateAssessors[i].ObtainPublicToken()
		assert.NoError(t, err)
		location := pv.Location()
		ballot := usingAssessor(ballotSchema, location, i)
		_, err = attestAppendBallot(privateAssessors[i], ballot, ballotAssign)
		if err != nil {
			t.Error(err)
		}
	}

	//
	verAltitudeArgument := IfaceParameters{BallotAdditionsActivateAltitude: altitude}
	assert.Panics(t, func() { ballotAssign.CreateExpandedEndorse(verAltitudeArgument) }, "REDACTED")

	//
	{
		pv, err := privateAssessors[6].ObtainPublicToken()
		assert.NoError(t, err)
		location := pv.Location()
		ballot := usingAssessor(ballotSchema, location, 6)
		ballot = usingLedgerDigest(ballot, commitrand.Octets(32))
		ballot = usingLedgerFragmentAssignHeading(ballot, FragmentAssignHeading{123, commitrand.Octets(32)})

		_, err = attestAppendBallot(privateAssessors[6], ballot, ballotAssign)
		require.NoError(t, err)
	}

	//
	{
		pv, err := privateAssessors[7].ObtainPublicToken()
		assert.NoError(t, err)
		location := pv.Location()
		ballot := usingAssessor(ballotSchema, location, 7)
		_, err = attestAppendBallot(privateAssessors[7], ballot, ballotAssign)
		require.NoError(t, err)
	}

	//
	{
		pv, err := privateAssessors[8].ObtainPublicToken()
		assert.NoError(t, err)
		location := pv.Location()
		ballot := usingAssessor(ballotSchema, location, 8)
		ballot.LedgerUUID = LedgerUUID{}

		_, err = attestAppendBallot(privateAssessors[8], ballot, ballotAssign)
		require.NoError(t, err)
	}

	addnEndorse := ballotAssign.CreateExpandedEndorse(verAltitudeArgument)

	//
	assert.Equal(t, 10, len(addnEndorse.ExpandedNotations))

	//
	if err := addnEndorse.CertifyFundamental(); err != nil {
		t.Errorf("REDACTED", err)
	}
}

//
//
func Assessmentvoteset_Voteadditionsenabled(t *testing.T) {
	for _, tc := range []struct {
		alias              string
		demandAdditions bool
		appendAddition      bool
		anticipateFailure      bool
	}{
		{
			alias:              "REDACTED",
			demandAdditions: true,
			appendAddition:      false,
			anticipateFailure:      true,
		},
		{
			alias:              "REDACTED",
			demandAdditions: true,
			appendAddition:      false,
			anticipateFailure:      true,
		},
		{
			alias:              "REDACTED",
			demandAdditions: false,
			appendAddition:      false,
			anticipateFailure:      false,
		},
		{
			alias:              "REDACTED",
			demandAdditions: true,
			appendAddition:      true,
			anticipateFailure:      false,
		},
	} {
		t.Run(tc.alias, func(t *testing.T) {
			altitude, iteration := int64(1), int32(0)
			itemAssign, privateAssessors := ArbitraryAssessorAssign(5, 10)
			var ballotAssign *BallotAssign
			if tc.demandAdditions {
				ballotAssign = FreshExpandedBallotAssign("REDACTED", altitude, iteration, commitchema.PreendorseKind, itemAssign)
			} else {
				ballotAssign = FreshBallotAssign("REDACTED", altitude, iteration, commitchema.PreendorseKind, itemAssign)
			}

			item0 := privateAssessors[0]

			item0p, err := item0.ObtainPublicToken()
			require.NoError(t, err)
			item0address := item0p.Location()
			ledgerDigest := security.CHARArbitraryOctets(32)
			ledgerFragmentsSum := uint32(123)
			ledgerFragmentAssignHeading := FragmentAssignHeading{ledgerFragmentsSum, security.CHARArbitraryOctets(32)}

			ballot := &Ballot{
				AssessorLocation: item0address,
				AssessorOrdinal:   0,
				Altitude:           altitude,
				Iteration:            iteration,
				Kind:             commitchema.PreendorseKind,
				Timestamp:        committime.Now(),
				LedgerUUID:          LedgerUUID{ledgerDigest, ledgerFragmentAssignHeading},
			}
			v := ballot.TowardSchema()
			err = item0.AttestBallot(ballotAssign.SuccessionUUID(), v)
			require.NoError(t, err)
			ballot.Notation = v.Notation

			if tc.appendAddition {
				ballot.AdditionNotation = v.AdditionNotation
			}

			appended, err := ballotAssign.AppendBallot(ballot)
			if tc.anticipateFailure {
				require.Error(t, err)
				require.False(t, appended)
			} else {
				require.NoError(t, err)
				require.True(t, appended)
			}
		})
	}
}

//
func arbitraryBallotAssign(
	altitude int64,
	iteration int32,
	notatedSignalKind commitchema.AttestedSignalKind,
	countAssessors int,
	ballotingPotency int64,
	addnActivated bool,
) (*BallotAssign, *AssessorAssign, []PrivateAssessor) {
	itemAssign, privateAssessors := ArbitraryAssessorAssign(countAssessors, ballotingPotency)
	if addnActivated {
		if notatedSignalKind != commitchema.PreendorseKind {
			return nil, nil, nil
		}
		return FreshExpandedBallotAssign("REDACTED", altitude, iteration, notatedSignalKind, itemAssign), itemAssign, privateAssessors
	}
	return FreshBallotAssign("REDACTED", altitude, iteration, notatedSignalKind, itemAssign), itemAssign, privateAssessors
}

//
func usingAssessor(ballot *Ballot, location []byte, idx int32) *Ballot {
	ballot = ballot.Duplicate()
	ballot.AssessorLocation = location
	ballot.AssessorOrdinal = idx
	return ballot
}

//
func usingAltitude(ballot *Ballot, altitude int64) *Ballot {
	ballot = ballot.Duplicate()
	ballot.Altitude = altitude
	return ballot
}

//
func usingIteration(ballot *Ballot, iteration int32) *Ballot {
	ballot = ballot.Duplicate()
	ballot.Iteration = iteration
	return ballot
}

//
func usingKind(ballot *Ballot, notatedSignalKind byte) *Ballot {
	ballot = ballot.Duplicate()
	ballot.Kind = commitchema.AttestedSignalKind(notatedSignalKind)
	return ballot
}

//
func usingLedgerDigest(ballot *Ballot, ledgerDigest []byte) *Ballot {
	ballot = ballot.Duplicate()
	ballot.LedgerUUID.Digest = ledgerDigest
	return ballot
}

//
func usingLedgerFragmentAssignHeading(ballot *Ballot, ledgerFragmentsHeading FragmentAssignHeading) *Ballot {
	ballot = ballot.Duplicate()
	ballot.LedgerUUID.FragmentAssignHeading = ledgerFragmentsHeading
	return ballot
}
