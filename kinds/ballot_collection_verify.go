package kinds

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

func Verifyballotcollection_Appendballot_Sound(t *testing.T) {
	level, epoch := int64(1), int32(0)
	ballotCollection, _, privateRatifiers := randomBallotCollection(level, epoch, engineproto.PreballotKind, 10, 1, false)
	node0 := privateRatifiers[0]

	node0p, err := node0.FetchPublicKey()
	require.NoError(t, err)
	node0address := node0p.Location()

	assert.Nil(t, ballotCollection.FetchByLocation(node0address))
	assert.False(t, ballotCollection.BitList().FetchOrdinal(0))
	ledgerUID, ok := ballotCollection.DualThirdsBulk()
	assert.False(t, ok || !ledgerUID.IsNil(), "REDACTED")

	ballot := &Ballot{
		RatifierLocation: node0address,
		RatifierOrdinal:   0, //
		Level:           level,
		Cycle:            epoch,
		Kind:             engineproto.PreballotKind,
		Timestamp:        engineclock.Now(),
		LedgerUID:          LedgerUID{nil, SegmentAssignHeading{}},
	}
	_, err = attestAppendBallot(node0, ballot, ballotCollection)
	require.NoError(t, err)

	assert.NotNil(t, ballotCollection.FetchByLocation(node0address))
	assert.True(t, ballotCollection.BitList().FetchOrdinal(0))
	ledgerUID, ok = ballotCollection.DualThirdsBulk()
	assert.False(t, ok || !ledgerUID.IsNil(), "REDACTED")
}

func Verifyballotcollection_Appendballot_Flawed(t *testing.T) {
	level, epoch := int64(1), int32(0)
	ballotCollection, _, privateRatifiers := randomBallotCollection(level, epoch, engineproto.PreballotKind, 10, 1, false)

	ballotSchema := &Ballot{
		RatifierLocation: nil,
		RatifierOrdinal:   -1,
		Level:           level,
		Cycle:            epoch,
		Timestamp:        engineclock.Now(),
		Kind:             engineproto.PreballotKind,
		LedgerUID:          LedgerUID{nil, SegmentAssignHeading{}},
	}

	//
	{
		publicKey, err := privateRatifiers[0].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 0)
		appended, err := attestAppendBallot(privateRatifiers[0], ballot, ballotCollection)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicKey, err := privateRatifiers[0].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 0)
		appended, err := attestAppendBallot(privateRatifiers[0], withLedgerDigest(ballot, engineseed.Octets(32)), ballotCollection)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicKey, err := privateRatifiers[1].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 1)
		appended, err := attestAppendBallot(privateRatifiers[1], withLevel(ballot, level+1), ballotCollection)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicKey, err := privateRatifiers[2].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 2)
		appended, err := attestAppendBallot(privateRatifiers[2], withEpoch(ballot, epoch+1), ballotCollection)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		publicKey, err := privateRatifiers[3].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 3)
		appended, err := attestAppendBallot(privateRatifiers[3], withKind(ballot, byte(engineproto.PreendorseKind)), ballotCollection)
		if appended || err == nil {
			t.Errorf("REDACTED")
		}
	}
}

func Criterion_two_three_Major(b *testing.B) {
	level, epoch := int64(1), int32(0)

	ballotSchema := &Ballot{
		RatifierLocation: nil, //
		RatifierOrdinal:   -1,  //
		Level:           level,
		Cycle:            epoch,
		Kind:             engineproto.PreballotKind,
		Timestamp:        engineclock.Now(),
		LedgerUID:          LedgerUID{nil, SegmentAssignHeading{}},
	}
	ledgerSegmentsSum := uint32(123)
	ledgerSectionCollectionHeading := SegmentAssignHeading{ledgerSegmentsSum, vault.CRandomOctets(32)}
	for b.Loop() {
		ballotCollection, _, privateRatifiers := randomBallotCollection(level, epoch, engineproto.PreballotKind, 100, 1, false)
		for i := int32(0); i < int32(100); i += 4 {
			publicKey, _ := privateRatifiers[i].FetchPublicKey()
			address := publicKey.Location()
			ballot := withRatifier(ballotSchema, address, i)
			_, err := attestAppendBallot(privateRatifiers[i], withLedgerDigest(ballot, nil), ballotCollection)
			require.NoError(b, err)
			_, _ = ballotCollection.DualThirdsBulk()

			publicKey, _ = privateRatifiers[i+1].FetchPublicKey()
			address = publicKey.Location()
			ballot = withRatifier(ballotSchema, address, i+1)
			_, err = attestAppendBallot(privateRatifiers[i+1], ballot, ballotCollection)
			require.NoError(b, err)
			_, _ = ballotCollection.DualThirdsBulk()

			publicKey, _ = privateRatifiers[i+2].FetchPublicKey()
			address = publicKey.Location()
			ballot = withRatifier(ballotSchema, address, i+2)
			ledgerSegmentsHeading := SegmentAssignHeading{ledgerSegmentsSum, vault.CRandomOctets(32)}
			_, err = attestAppendBallot(privateRatifiers[i+2], withLedgerSectionCollectionHeading(ballot, ledgerSegmentsHeading), ballotCollection)
			require.NoError(b, err)
			_, _ = ballotCollection.DualThirdsBulk()

			publicKey, _ = privateRatifiers[i+3].FetchPublicKey()
			address = publicKey.Location()
			ballot = withRatifier(ballotSchema, address, i+3)
			ledgerSegmentsHeading = SegmentAssignHeading{ledgerSegmentsSum + 1, ledgerSectionCollectionHeading.Digest}
			_, err = attestAppendBallot(privateRatifiers[i+3], withLedgerSectionCollectionHeading(ballot, ledgerSegmentsHeading), ballotCollection)
			require.NoError(b, err)
			_, _ = ballotCollection.DualThirdsBulk()
		}
	}
}

func Verifyballotcollection_two_threemajority(t *testing.T) {
	level, epoch := int64(1), int32(0)
	ballotCollection, _, privateRatifiers := randomBallotCollection(level, epoch, engineproto.PreballotKind, 10, 1, false)

	ballotSchema := &Ballot{
		RatifierLocation: nil, //
		RatifierOrdinal:   -1,  //
		Level:           level,
		Cycle:            epoch,
		Kind:             engineproto.PreballotKind,
		Timestamp:        engineclock.Now(),
		LedgerUID:          LedgerUID{nil, SegmentAssignHeading{}},
	}
	//
	for i := int32(0); i < 6; i++ {
		publicKey, err := privateRatifiers[i].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, i)
		_, err = attestAppendBallot(privateRatifiers[i], ballot, ballotCollection)
		require.NoError(t, err)
	}
	ledgerUID, ok := ballotCollection.DualThirdsBulk()
	assert.False(t, ok || !ledgerUID.IsNil(), "REDACTED")

	//
	{
		publicKey, err := privateRatifiers[6].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 6)
		_, err = attestAppendBallot(privateRatifiers[6], withLedgerDigest(ballot, engineseed.Octets(32)), ballotCollection)
		require.NoError(t, err)
		ledgerUID, ok = ballotCollection.DualThirdsBulk()
		assert.False(t, ok || !ledgerUID.IsNil(), "REDACTED")
	}

	//
	{
		publicKey, err := privateRatifiers[7].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 7)
		_, err = attestAppendBallot(privateRatifiers[7], ballot, ballotCollection)
		require.NoError(t, err)
		ledgerUID, ok = ballotCollection.DualThirdsBulk()
		assert.True(t, ok || ledgerUID.IsNil(), "REDACTED")
	}
}

func Verifyballotcollection_two_threemajorityreduced(t *testing.T) {
	level, epoch := int64(1), int32(0)
	ballotCollection, _, privateRatifiers := randomBallotCollection(level, epoch, engineproto.PreballotKind, 100, 1, false)

	ledgerDigest := vault.CRandomOctets(32)
	ledgerSegmentsSum := uint32(123)
	ledgerSectionCollectionHeading := SegmentAssignHeading{ledgerSegmentsSum, vault.CRandomOctets(32)}

	ballotSchema := &Ballot{
		RatifierLocation: nil, //
		RatifierOrdinal:   -1,  //
		Level:           level,
		Cycle:            epoch,
		Timestamp:        engineclock.Now(),
		Kind:             engineproto.PreballotKind,
		LedgerUID:          LedgerUID{ledgerDigest, ledgerSectionCollectionHeading},
	}

	//
	for i := int32(0); i < 66; i++ {
		publicKey, err := privateRatifiers[i].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, i)
		_, err = attestAppendBallot(privateRatifiers[i], ballot, ballotCollection)
		require.NoError(t, err)
	}
	ledgerUID, ok := ballotCollection.DualThirdsBulk()
	assert.False(t, ok || !ledgerUID.IsNil(),
		"REDACTED")

	//
	{
		publicKey, err := privateRatifiers[66].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 66)
		_, err = attestAppendBallot(privateRatifiers[66], withLedgerDigest(ballot, nil), ballotCollection)
		require.NoError(t, err)
		ledgerUID, ok = ballotCollection.DualThirdsBulk()
		assert.False(t, ok || !ledgerUID.IsNil(),
			"REDACTED")
	}

	//
	{
		publicKey, err := privateRatifiers[67].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 67)
		ledgerSegmentsHeading := SegmentAssignHeading{ledgerSegmentsSum, vault.CRandomOctets(32)}
		_, err = attestAppendBallot(privateRatifiers[67], withLedgerSectionCollectionHeading(ballot, ledgerSegmentsHeading), ballotCollection)
		require.NoError(t, err)
		ledgerUID, ok = ballotCollection.DualThirdsBulk()
		assert.False(t, ok || !ledgerUID.IsNil(),
			"REDACTED")
	}

	//
	{
		publicKey, err := privateRatifiers[68].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 68)
		ledgerSegmentsHeading := SegmentAssignHeading{ledgerSegmentsSum + 1, ledgerSectionCollectionHeading.Digest}
		_, err = attestAppendBallot(privateRatifiers[68], withLedgerSectionCollectionHeading(ballot, ledgerSegmentsHeading), ballotCollection)
		require.NoError(t, err)
		ledgerUID, ok = ballotCollection.DualThirdsBulk()
		assert.False(t, ok || !ledgerUID.IsNil(),
			"REDACTED")
	}

	//
	{
		publicKey, err := privateRatifiers[69].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 69)
		_, err = attestAppendBallot(privateRatifiers[69], withLedgerDigest(ballot, engineseed.Octets(32)), ballotCollection)
		require.NoError(t, err)
		ledgerUID, ok = ballotCollection.DualThirdsBulk()
		assert.False(t, ok || !ledgerUID.IsNil(),
			"REDACTED")
	}

	//
	{
		publicKey, err := privateRatifiers[70].FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		ballot := withRatifier(ballotSchema, address, 70)
		_, err = attestAppendBallot(privateRatifiers[70], ballot, ballotCollection)
		require.NoError(t, err)
		ledgerUID, ok = ballotCollection.DualThirdsBulk()
		assert.True(t, ok && ledgerUID.Matches(LedgerUID{ledgerDigest, ledgerSectionCollectionHeading}),
			"REDACTED")
	}
}

func Verifyballotcollection_Clashes(t *testing.T) {
	level, epoch := int64(1), int32(0)
	ballotCollection, _, privateRatifiers := randomBallotCollection(level, epoch, engineproto.PreballotKind, 4, 1, false)
	ledgerDigest1 := engineseed.Octets(32)
	ledgerDigest2 := engineseed.Octets(32)

	ballotSchema := &Ballot{
		RatifierLocation: nil,
		RatifierOrdinal:   -1,
		Level:           level,
		Cycle:            epoch,
		Timestamp:        engineclock.Now(),
		Kind:             engineproto.PreballotKind,
		LedgerUID:          LedgerUID{nil, SegmentAssignHeading{}},
	}

	node0, err := privateRatifiers[0].FetchPublicKey()
	require.NoError(t, err)
	node0address := node0.Location()

	//
	{
		ballot := withRatifier(ballotSchema, node0address, 0)
		appended, err := attestAppendBallot(privateRatifiers[0], ballot, ballotCollection)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	{
		ballot := withRatifier(ballotSchema, node0address, 0)
		appended, err := attestAppendBallot(privateRatifiers[0], withLedgerDigest(ballot, ledgerDigest1), ballotCollection)
		assert.False(t, appended, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}

	//
	err = ballotCollection.AssignNodeMaj23("REDACTED", LedgerUID{ledgerDigest1, SegmentAssignHeading{}})
	require.NoError(t, err)

	//
	{
		ballot := withRatifier(ballotSchema, node0address, 0)
		appended, err := attestAppendBallot(privateRatifiers[0], withLedgerDigest(ballot, ledgerDigest1), ballotCollection)
		assert.True(t, appended, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}

	//
	err = ballotCollection.AssignNodeMaj23("REDACTED", LedgerUID{ledgerDigest2, SegmentAssignHeading{}})
	require.Error(t, err)

	//
	{
		ballot := withRatifier(ballotSchema, node0address, 0)
		appended, err := attestAppendBallot(privateRatifiers[0], withLedgerDigest(ballot, ledgerDigest2), ballotCollection)
		assert.False(t, appended, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}

	//
	{
		pv, err := privateRatifiers[1].FetchPublicKey()
		assert.NoError(t, err)
		address := pv.Location()
		ballot := withRatifier(ballotSchema, address, 1)
		appended, err := attestAppendBallot(privateRatifiers[1], withLedgerDigest(ballot, ledgerDigest1), ballotCollection)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	if ballotCollection.HasDualThirdsBulk() {
		t.Errorf("REDACTED")
	}
	if ballotCollection.HasDualThirdsAny() {
		t.Errorf("REDACTED")
	}

	//
	{
		pv, err := privateRatifiers[2].FetchPublicKey()
		assert.NoError(t, err)
		address := pv.Location()
		ballot := withRatifier(ballotSchema, address, 2)
		appended, err := attestAppendBallot(privateRatifiers[2], withLedgerDigest(ballot, ledgerDigest2), ballotCollection)
		if !appended || err != nil {
			t.Errorf("REDACTED")
		}
	}

	//
	if ballotCollection.HasDualThirdsBulk() {
		t.Errorf("REDACTED")
	}
	if !ballotCollection.HasDualThirdsAny() {
		t.Errorf("REDACTED")
	}

	//
	err = ballotCollection.AssignNodeMaj23("REDACTED", LedgerUID{ledgerDigest1, SegmentAssignHeading{}})
	require.NoError(t, err)

	//
	{
		pv, err := privateRatifiers[2].FetchPublicKey()
		assert.NoError(t, err)
		address := pv.Location()
		ballot := withRatifier(ballotSchema, address, 2)
		appended, err := attestAppendBallot(privateRatifiers[2], withLedgerDigest(ballot, ledgerDigest1), ballotCollection)
		assert.True(t, appended)
		assert.Error(t, err, "REDACTED")
	}

	//
	if !ballotCollection.HasDualThirdsBulk() {
		t.Errorf("REDACTED")
	}
	ledgerUIDMaj23, _ := ballotCollection.DualThirdsBulk()
	if !bytes.Equal(ledgerUIDMaj23.Digest, ledgerDigest1) {
		t.Errorf("REDACTED")
	}
	if !ballotCollection.HasDualThirdsAny() {
		t.Errorf("REDACTED")
	}
}

func Verifyballotcollection_Createcommit(t *testing.T) {
	level, epoch := int64(1), int32(0)
	ballotCollection, _, privateRatifiers := randomBallotCollection(level, epoch, engineproto.PreendorseKind, 10, 1, true)
	ledgerDigest, ledgerSectionCollectionHeading := vault.CRandomOctets(32), SegmentAssignHeading{123, vault.CRandomOctets(32)}

	ballotSchema := &Ballot{
		RatifierLocation: nil,
		RatifierOrdinal:   -1,
		Level:           level,
		Cycle:            epoch,
		Timestamp:        engineclock.Now(),
		Kind:             engineproto.PreendorseKind,
		LedgerUID:          LedgerUID{ledgerDigest, ledgerSectionCollectionHeading},
	}

	//
	for i := int32(0); i < 6; i++ {
		pv, err := privateRatifiers[i].FetchPublicKey()
		assert.NoError(t, err)
		address := pv.Location()
		ballot := withRatifier(ballotSchema, address, i)
		_, err = attestAppendBallot(privateRatifiers[i], ballot, ballotCollection)
		if err != nil {
			t.Error(err)
		}
	}

	//
	veLevelArgument := IfaceOptions{BallotPluginsActivateLevel: level}
	assert.Panics(t, func() { ballotCollection.CreateExpandedEndorse(veLevelArgument) }, "REDACTED")

	//
	{
		pv, err := privateRatifiers[6].FetchPublicKey()
		assert.NoError(t, err)
		address := pv.Location()
		ballot := withRatifier(ballotSchema, address, 6)
		ballot = withLedgerDigest(ballot, engineseed.Octets(32))
		ballot = withLedgerSectionCollectionHeading(ballot, SegmentAssignHeading{123, engineseed.Octets(32)})

		_, err = attestAppendBallot(privateRatifiers[6], ballot, ballotCollection)
		require.NoError(t, err)
	}

	//
	{
		pv, err := privateRatifiers[7].FetchPublicKey()
		assert.NoError(t, err)
		address := pv.Location()
		ballot := withRatifier(ballotSchema, address, 7)
		_, err = attestAppendBallot(privateRatifiers[7], ballot, ballotCollection)
		require.NoError(t, err)
	}

	//
	{
		pv, err := privateRatifiers[8].FetchPublicKey()
		assert.NoError(t, err)
		address := pv.Location()
		ballot := withRatifier(ballotSchema, address, 8)
		ballot.LedgerUID = LedgerUID{}

		_, err = attestAppendBallot(privateRatifiers[8], ballot, ballotCollection)
		require.NoError(t, err)
	}

	extensionEndorse := ballotCollection.CreateExpandedEndorse(veLevelArgument)

	//
	assert.Equal(t, 10, len(extensionEndorse.ExpandedEndorsements))

	//
	if err := extensionEndorse.CertifySimple(); err != nil {
		t.Errorf("REDACTED", err)
	}
}

//
//
func Verifyballotcollection_Ballotextensionsenabled(t *testing.T) {
	for _, tc := range []struct {
		label              string
		demandPlugins bool
		appendAddition      bool
		anticipateFault      bool
	}{
		{
			label:              "REDACTED",
			demandPlugins: true,
			appendAddition:      false,
			anticipateFault:      true,
		},
		{
			label:              "REDACTED",
			demandPlugins: true,
			appendAddition:      false,
			anticipateFault:      true,
		},
		{
			label:              "REDACTED",
			demandPlugins: false,
			appendAddition:      false,
			anticipateFault:      false,
		},
		{
			label:              "REDACTED",
			demandPlugins: true,
			appendAddition:      true,
			anticipateFault:      false,
		},
	} {
		t.Run(tc.label, func(t *testing.T) {
			level, epoch := int64(1), int32(0)
			valueCollection, privateRatifiers := RandomRatifierCollection(5, 10)
			var ballotCollection *BallotCollection
			if tc.demandPlugins {
				ballotCollection = NewExpandedBallotCollection("REDACTED", level, epoch, engineproto.PreendorseKind, valueCollection)
			} else {
				ballotCollection = NewBallotCollection("REDACTED", level, epoch, engineproto.PreendorseKind, valueCollection)
			}

			node0 := privateRatifiers[0]

			node0p, err := node0.FetchPublicKey()
			require.NoError(t, err)
			node0address := node0p.Location()
			ledgerDigest := vault.CRandomOctets(32)
			ledgerSegmentsSum := uint32(123)
			ledgerSectionCollectionHeading := SegmentAssignHeading{ledgerSegmentsSum, vault.CRandomOctets(32)}

			ballot := &Ballot{
				RatifierLocation: node0address,
				RatifierOrdinal:   0,
				Level:           level,
				Cycle:            epoch,
				Kind:             engineproto.PreendorseKind,
				Timestamp:        engineclock.Now(),
				LedgerUID:          LedgerUID{ledgerDigest, ledgerSectionCollectionHeading},
			}
			v := ballot.ToSchema()
			err = node0.AttestBallot(ballotCollection.LedgerUID(), v)
			require.NoError(t, err)
			ballot.Autograph = v.Autograph

			if tc.appendAddition {
				ballot.AdditionAutograph = v.AdditionAutograph
			}

			appended, err := ballotCollection.AppendBallot(ballot)
			if tc.anticipateFault {
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
func randomBallotCollection(
	level int64,
	epoch int32,
	attestedMessageKind engineproto.AttestedMessageKind,
	countRatifiers int,
	pollingEnergy int64,
	extensionActivated bool,
) (*BallotCollection, *RatifierAssign, []PrivateRatifier) {
	valueCollection, privateRatifiers := RandomRatifierCollection(countRatifiers, pollingEnergy)
	if extensionActivated {
		if attestedMessageKind != engineproto.PreendorseKind {
			return nil, nil, nil
		}
		return NewExpandedBallotCollection("REDACTED", level, epoch, attestedMessageKind, valueCollection), valueCollection, privateRatifiers
	}
	return NewBallotCollection("REDACTED", level, epoch, attestedMessageKind, valueCollection), valueCollection, privateRatifiers
}

//
func withRatifier(ballot *Ballot, address []byte, idx int32) *Ballot {
	ballot = ballot.Clone()
	ballot.RatifierLocation = address
	ballot.RatifierOrdinal = idx
	return ballot
}

//
func withLevel(ballot *Ballot, level int64) *Ballot {
	ballot = ballot.Clone()
	ballot.Level = level
	return ballot
}

//
func withEpoch(ballot *Ballot, epoch int32) *Ballot {
	ballot = ballot.Clone()
	ballot.Cycle = epoch
	return ballot
}

//
func withKind(ballot *Ballot, attestedMessageKind byte) *Ballot {
	ballot = ballot.Clone()
	ballot.Kind = engineproto.AttestedMessageKind(attestedMessageKind)
	return ballot
}

//
func withLedgerDigest(ballot *Ballot, ledgerDigest []byte) *Ballot {
	ballot = ballot.Clone()
	ballot.LedgerUID.Digest = ledgerDigest
	return ballot
}

//
func withLedgerSectionCollectionHeading(ballot *Ballot, ledgerSegmentsHeading SegmentAssignHeading) *Ballot {
	ballot = ballot.Clone()
	ballot.LedgerUID.SegmentAssignHeading = ledgerSegmentsHeading
	return ballot
}
