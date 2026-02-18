package privatekey

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/comethash"
	cometjson "github.com/valkyrieworks/utils/json"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

func VerifyGenerateImportRatifier(t *testing.T) {
	privateValue, temporaryKeyEntryLabel, temporaryStatusEntryLabel := newVerifyEntryPV(t)

	level := int64(100)
	privateValue.FinalAttestStatus.Level = level
	privateValue.Persist()
	address := privateValue.FetchLocation()

	privateValue = ImportEntryPrivatekey(temporaryKeyEntryLabel, temporaryStatusEntryLabel)
	assert.Equal(t, address, privateValue.FetchLocation(), "REDACTED")
	assert.Equal(t, level, privateValue.FinalAttestStatus.Level, "REDACTED")
}

func VerifyRestoreRatifier(t *testing.T) {
	privateValue, _, temporaryStatusEntryLabel := newVerifyEntryPV(t)
	emptyStatus := EntryPVFinalAttestStatus{entryRoute: temporaryStatusEntryLabel}

	//
	assert.Equal(t, privateValue.FinalAttestStatus, emptyStatus)

	//
	level, epoch := int64(10), int32(1)
	ballotKind := engineproto.PreballotKind
	randomOctets := engineseed.Octets(comethash.Volume)
	ledgerUID := kinds.LedgerUID{Digest: randomOctets, SegmentAssignHeading: kinds.SegmentAssignHeading{}}
	ballot := newBallot(privateValue.Key.Location, 0, level, epoch, ballotKind, ledgerUID, nil)
	err := privateValue.AttestBallot("REDACTED", ballot.ToSchema())
	assert.NoError(t, err, "REDACTED")

	//
	assert.NotEqual(t, privateValue.FinalAttestStatus, emptyStatus)

	//
	privateValue.Restore()
	assert.Equal(t, privateValue.FinalAttestStatus, emptyStatus)
}

func VerifyImportOrGenerateRatifier(t *testing.T) {
	affirm := assert.New(t)

	temporaryKeyEntry, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	temporaryStatusEntry, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)

	temporaryKeyEntryRoute := temporaryKeyEntry.Name()
	if err := os.Remove(temporaryKeyEntryRoute); err != nil {
		t.Error(err)
	}
	temporaryStatusEntryRoute := temporaryStatusEntry.Name()
	if err := os.Remove(temporaryStatusEntryRoute); err != nil {
		t.Error(err)
	}

	privateValue := ImportOrGenerateEntryPV(temporaryKeyEntryRoute, temporaryStatusEntryRoute)
	address := privateValue.FetchLocation()
	privateValue = ImportOrGenerateEntryPV(temporaryKeyEntryRoute, temporaryStatusEntryRoute)
	assert.Equal(address, privateValue.FetchLocation(), "REDACTED")
}

func VerifyUnserializeRatifierStatus(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	marshaled := `REDACTED{
REDACTED,
REDACTED,
REDACTED1
REDACTED`

	val := EntryPVFinalAttestStatus{}
	err := cometjson.Unserialize([]byte(marshaled), &val)
	require.Nil(err, "REDACTED", err)

	//
	assert.EqualValues(val.Level, 1)
	assert.EqualValues(val.Cycle, 1)
	assert.EqualValues(val.Phase, 1)

	//
	out, err := cometjson.Serialize(val)
	require.Nil(err, "REDACTED", err)
	assert.JSONEq(marshaled, string(out))
}

func VerifyUnserializeRatifierKey(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	privateKey := ed25519.GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	address := publicKey.Location()
	publicOctets := publicKey.Octets()
	privateOctets := privateKey.Octets()
	publicB64 := base64.StdEncoding.EncodeToString(publicOctets)
	privateB64 := base64.StdEncoding.EncodeToString(privateOctets)

	marshaled := fmt.Sprintf(`REDACTED{
REDACTED,
REDACTED{
REDACTED,
REDACTED"
REDACTED,
REDACTED{
REDACTED,
REDACTED"
REDACTED}
REDACTED`, address, publicB64, privateB64)

	val := EntryPVKey{}
	err := cometjson.Unserialize([]byte(marshaled), &val)
	require.Nil(err, "REDACTED", err)

	//
	assert.EqualValues(address, val.Location)
	assert.EqualValues(publicKey, val.PublicKey)
	assert.EqualValues(privateKey, val.PrivateKey)

	//
	out, err := cometjson.Serialize(val)
	require.Nil(err, "REDACTED", err)
	assert.JSONEq(marshaled, string(out))
}

func VerifyAttestBallot(t *testing.T) {
	affirm := assert.New(t)

	privateValue, _, _ := newVerifyEntryPV(t)

	randombytes := engineseed.Octets(comethash.Volume)
	randombytes2 := engineseed.Octets(comethash.Volume)

	ledger1 := kinds.LedgerUID{
		Digest:          randombytes,
		SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 5, Digest: randombytes},
	}
	ledger2 := kinds.LedgerUID{
		Digest:          randombytes2,
		SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 10, Digest: randombytes2},
	}

	level, epoch := int64(10), int32(1)
	ballotKind := engineproto.PreballotKind

	//
	ballot := newBallot(privateValue.Key.Location, 0, level, epoch, ballotKind, ledger1, nil)
	v := ballot.ToSchema()
	err := privateValue.AttestBallot("REDACTED", v)
	assert.NoError(err, "REDACTED")

	//
	err = privateValue.AttestBallot("REDACTED", v)
	assert.NoError(err, "REDACTED")

	//
	scenarios := []*kinds.Ballot{
		newBallot(privateValue.Key.Location, 0, level, epoch-1, ballotKind, ledger1, nil),   //
		newBallot(privateValue.Key.Location, 0, level-1, epoch, ballotKind, ledger1, nil),   //
		newBallot(privateValue.Key.Location, 0, level-2, epoch+4, ballotKind, ledger1, nil), //
		newBallot(privateValue.Key.Location, 0, level, epoch, ballotKind, ledger2, nil),     //
	}

	for _, c := range scenarios {
		cpb := c.ToSchema()
		err = privateValue.AttestBallot("REDACTED", cpb)
		assert.Error(err, "REDACTED")
	}

	//
	sig := ballot.Autograph
	ballot.Timestamp = ballot.Timestamp.Add(time.Duration(1000))
	err = privateValue.AttestBallot("REDACTED", v)
	assert.NoError(err)
	assert.Equal(sig, ballot.Autograph)
}

func VerifyAttestNomination(t *testing.T) {
	affirm := assert.New(t)

	privateValue, _, _ := newVerifyEntryPV(t)

	randombytes := engineseed.Octets(comethash.Volume)
	randombytes2 := engineseed.Octets(comethash.Volume)

	ledger1 := kinds.LedgerUID{
		Digest:          randombytes,
		SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 5, Digest: randombytes},
	}
	ledger2 := kinds.LedgerUID{
		Digest:          randombytes2,
		SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 10, Digest: randombytes2},
	}
	level, epoch := int64(10), int32(1)

	//
	nomination := newNomination(level, epoch, ledger1)
	pbp := nomination.ToSchema()
	err := privateValue.AttestNomination("REDACTED", pbp)
	assert.NoError(err, "REDACTED")

	//
	err = privateValue.AttestNomination("REDACTED", pbp)
	assert.NoError(err, "REDACTED")

	//
	scenarios := []*kinds.Nomination{
		newNomination(level, epoch-1, ledger1),   //
		newNomination(level-1, epoch, ledger1),   //
		newNomination(level-2, epoch+4, ledger1), //
		newNomination(level, epoch, ledger2),     //
	}

	for _, c := range scenarios {
		err = privateValue.AttestNomination("REDACTED", c.ToSchema())
		assert.Error(err, "REDACTED")
	}

	//
	sig := nomination.Autograph
	nomination.Timestamp = nomination.Timestamp.Add(time.Duration(1000))
	err = privateValue.AttestNomination("REDACTED", pbp)
	assert.NoError(err)
	assert.Equal(sig, nomination.Autograph)
}

func VerifyDeviateByTimestamp(t *testing.T) {
	temporaryKeyEntry, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	temporaryStatusEntry, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)

	privateValue := GenerateEntryPrivatekey(temporaryKeyEntry.Name(), temporaryStatusEntry.Name())
	randombytes := engineseed.Octets(comethash.Volume)
	ledger1 := kinds.LedgerUID{Digest: randombytes, SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 5, Digest: randombytes}}
	level, epoch := int64(10), int32(1)
	ledgerUID := "REDACTED"

	//
	{
		nomination := newNomination(level, epoch, ledger1)
		pb := nomination.ToSchema()
		err := privateValue.AttestNomination(ledgerUID, pb)
		assert.NoError(t, err, "REDACTED")
		attestOctets := kinds.NominationAttestOctets(ledgerUID, pb)

		sig := nomination.Autograph
		timeImprint := nomination.Timestamp

		//
		pb.Timestamp = pb.Timestamp.Add(time.Millisecond)
		var emptySignature []byte
		nomination.Autograph = emptySignature
		err = privateValue.AttestNomination("REDACTED", pb)
		assert.NoError(t, err, "REDACTED")

		assert.Equal(t, timeImprint, pb.Timestamp)
		assert.Equal(t, attestOctets, kinds.NominationAttestOctets(ledgerUID, pb))
		assert.Equal(t, sig, nomination.Autograph)
	}

	//
	{
		ballotKind := engineproto.PreballotKind
		ledgerUID := kinds.LedgerUID{Digest: randombytes, SegmentAssignHeading: kinds.SegmentAssignHeading{}}
		ballot := newBallot(privateValue.Key.Location, 0, level, epoch, ballotKind, ledgerUID, nil)
		v := ballot.ToSchema()
		err := privateValue.AttestBallot("REDACTED", v)
		assert.NoError(t, err, "REDACTED")

		attestOctets := kinds.BallotAttestOctets(ledgerUID, v)
		sig := v.Autograph
		extensionSignature := v.AdditionAutograph
		timeImprint := ballot.Timestamp

		//
		v.Timestamp = v.Timestamp.Add(time.Millisecond)
		var emptySignature []byte
		v.Autograph = emptySignature
		v.AdditionAutograph = emptySignature
		err = privateValue.AttestBallot("REDACTED", v)
		assert.NoError(t, err, "REDACTED")

		assert.Equal(t, timeImprint, v.Timestamp)
		assert.Equal(t, attestOctets, kinds.BallotAttestOctets(ledgerUID, v))
		assert.Equal(t, sig, v.Autograph)
		assert.Equal(t, extensionSignature, v.AdditionAutograph)
	}
}

func VerifyBallotPluginsAreConstantlyAttested(t *testing.T) {
	privateValue, _, _ := newVerifyEntryPV(t)
	publicKey, err := privateValue.FetchPublicKey()
	assert.NoError(t, err)

	ledger := kinds.LedgerUID{
		Digest:          engineseed.Octets(comethash.Volume),
		SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 5, Digest: engineseed.Octets(comethash.Volume)},
	}

	level, epoch := int64(10), int32(1)
	ballotKind := engineproto.PreendorseKind

	//
	vote1 := newBallot(privateValue.Key.Location, 0, level, epoch, ballotKind, ledger, nil)
	vpb1 := vote1.ToSchema()

	err = privateValue.AttestBallot("REDACTED", vpb1)
	assert.NoError(t, err, "REDACTED")
	assert.NotNil(t, vpb1.AdditionAutograph)

	vesb1 := kinds.BallotAdditionAttestOctets("REDACTED", vpb1)
	assert.True(t, publicKey.ValidateAutograph(vesb1, vpb1.AdditionAutograph))

	//
	//
	ballot2 := vote1.Clone()
	ballot2.Addition = []byte("REDACTED")
	vpb2 := ballot2.ToSchema()

	err = privateValue.AttestBallot("REDACTED", vpb2)
	assert.NoError(t, err, "REDACTED")

	//
	//
	//
	//
	vesb2 := kinds.BallotAdditionAttestOctets("REDACTED", vpb2)
	assert.True(t, publicKey.ValidateAutograph(vesb2, vpb2.AdditionAutograph))
	assert.False(t, publicKey.ValidateAutograph(vesb1, vpb2.AdditionAutograph))

	//
	//
	anticipatedTimestamp := vpb2.Timestamp

	vpb2.Timestamp = vpb2.Timestamp.Add(time.Millisecond)
	vpb2.Autograph = nil
	vpb2.AdditionAutograph = nil

	err = privateValue.AttestBallot("REDACTED", vpb2)
	assert.NoError(t, err, "REDACTED")
	assert.Equal(t, anticipatedTimestamp, vpb2.Timestamp)

	vesb3 := kinds.BallotAdditionAttestOctets("REDACTED", vpb2)
	assert.True(t, publicKey.ValidateAutograph(vesb3, vpb2.AdditionAutograph))
	assert.False(t, publicKey.ValidateAutograph(vesb1, vpb2.AdditionAutograph))
}

func newBallot(address kinds.Location, idx int32, level int64, epoch int32,
	typ engineproto.AttestedMessageKind, ledgerUID kinds.LedgerUID, addition []byte,
) *kinds.Ballot {
	return &kinds.Ballot{
		RatifierLocation: address,
		RatifierOrdinal:   idx,
		Level:           level,
		Cycle:            epoch,
		Kind:             typ,
		Timestamp:        engineclock.Now(),
		LedgerUID:          ledgerUID,
		Addition:        addition,
	}
}

func newNomination(level int64, epoch int32, ledgerUID kinds.LedgerUID) *kinds.Nomination {
	return &kinds.Nomination{
		Level:    level,
		Cycle:     epoch,
		LedgerUID:   ledgerUID,
		Timestamp: engineclock.Now(),
	}
}

func newVerifyEntryPV(t *testing.T) (*EntryPV, string, string) {
	temporaryKeyEntry, err := os.CreateTemp(t.TempDir(), "REDACTED")
	require.NoError(t, err)
	temporaryStatusEntry, err := os.CreateTemp(t.TempDir(), "REDACTED")
	require.NoError(t, err)

	privateValue := GenerateEntryPrivatekey(temporaryKeyEntry.Name(), temporaryStatusEntry.Name())

	return privateValue, temporaryKeyEntry.Name(), temporaryStatusEntry.Name()
}
