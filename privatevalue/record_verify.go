package privatevalue

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

func VerifyProduceFetchAssessor(t *testing.T) {
	privateItem, transientTokenRecordAlias, transientStatusRecordAlias := freshVerifyRecordPRV(t)

	altitude := int64(100)
	privateItem.FinalAttestStatus.Altitude = altitude
	privateItem.Persist()
	location := privateItem.ObtainLocator()

	privateItem = FetchRecordPRV(transientTokenRecordAlias, transientStatusRecordAlias)
	assert.Equal(t, location, privateItem.ObtainLocator(), "REDACTED")
	assert.Equal(t, altitude, privateItem.FinalAttestStatus.Altitude, "REDACTED")
}

func VerifyRestoreAssessor(t *testing.T) {
	privateItem, _, transientStatusRecordAlias := freshVerifyRecordPRV(t)
	blankStatus := RecordPRVFinalAttestStatus{recordRoute: transientStatusRecordAlias}

	//
	assert.Equal(t, privateItem.FinalAttestStatus, blankStatus)

	//
	altitude, iteration := int64(10), int32(1)
	ballotKind := commitchema.PreballotKind
	arbitraryOctets := commitrand.Octets(tenderminthash.Extent)
	ledgerUUID := kinds.LedgerUUID{Digest: arbitraryOctets, FragmentAssignHeading: kinds.FragmentAssignHeading{}}
	ballot := freshBallot(privateItem.Key.Location, 0, altitude, iteration, ballotKind, ledgerUUID, nil)
	err := privateItem.AttestBallot("REDACTED", ballot.TowardSchema())
	assert.NoError(t, err, "REDACTED")

	//
	assert.NotEqual(t, privateItem.FinalAttestStatus, blankStatus)

	//
	privateItem.Restore()
	assert.Equal(t, privateItem.FinalAttestStatus, blankStatus)
}

func VerifyFetchEitherProduceAssessor(t *testing.T) {
	affirm := assert.New(t)

	transientTokenRecord, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	transientStatusRecord, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)

	transientTokenRecordRoute := transientTokenRecord.Name()
	if err := os.Remove(transientTokenRecordRoute); err != nil {
		t.Error(err)
	}
	transientStatusRecordRoute := transientStatusRecord.Name()
	if err := os.Remove(transientStatusRecordRoute); err != nil {
		t.Error(err)
	}

	privateItem := FetchEitherProduceRecordPRV(transientTokenRecordRoute, transientStatusRecordRoute)
	location := privateItem.ObtainLocator()
	privateItem = FetchEitherProduceRecordPRV(transientTokenRecordRoute, transientStatusRecordRoute)
	assert.Equal(location, privateItem.ObtainLocator(), "REDACTED")
}

func VerifyDecodeAssessorStatus(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	marshaled := `REDACTED{
REDACTED,
REDACTED,
REDACTED1
REDACTED`

	val := RecordPRVFinalAttestStatus{}
	err := strongmindjson.Decode([]byte(marshaled), &val)
	require.Nil(err, "REDACTED", err)

	//
	assert.EqualValues(val.Altitude, 1)
	assert.EqualValues(val.Iteration, 1)
	assert.EqualValues(val.Phase, 1)

	//
	out, err := strongmindjson.Serialize(val)
	require.Nil(err, "REDACTED", err)
	assert.JSONEq(marshaled, string(out))
}

func VerifyDecodeAssessorToken(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	privateToken := edwards25519.ProducePrivateToken()
	publicToken := privateToken.PublicToken()
	location := publicToken.Location()
	publicOctets := publicToken.Octets()
	privateOctets := privateToken.Octets()
	publicBase64 := base64.StdEncoding.EncodeToString(publicOctets)
	privateBase64 := base64.StdEncoding.EncodeToString(privateOctets)

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
REDACTED`, location, publicBase64, privateBase64)

	val := RecordPRVToken{}
	err := strongmindjson.Decode([]byte(marshaled), &val)
	require.Nil(err, "REDACTED", err)

	//
	assert.EqualValues(location, val.Location)
	assert.EqualValues(publicToken, val.PublicToken)
	assert.EqualValues(privateToken, val.PrivateToken)

	//
	out, err := strongmindjson.Serialize(val)
	require.Nil(err, "REDACTED", err)
	assert.JSONEq(marshaled, string(out))
}

func VerifyAttestBallot(t *testing.T) {
	affirm := assert.New(t)

	privateItem, _, _ := freshVerifyRecordPRV(t)

	randomoctets := commitrand.Octets(tenderminthash.Extent)
	randomoctets2 := commitrand.Octets(tenderminthash.Extent)

	ledger1 := kinds.LedgerUUID{
		Digest:          randomoctets,
		FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 5, Digest: randomoctets},
	}
	ledger2 := kinds.LedgerUUID{
		Digest:          randomoctets2,
		FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 10, Digest: randomoctets2},
	}

	altitude, iteration := int64(10), int32(1)
	ballotKind := commitchema.PreballotKind

	//
	ballot := freshBallot(privateItem.Key.Location, 0, altitude, iteration, ballotKind, ledger1, nil)
	v := ballot.TowardSchema()
	err := privateItem.AttestBallot("REDACTED", v)
	assert.NoError(err, "REDACTED")

	//
	err = privateItem.AttestBallot("REDACTED", v)
	assert.NoError(err, "REDACTED")

	//
	scenarios := []*kinds.Ballot{
		freshBallot(privateItem.Key.Location, 0, altitude, iteration-1, ballotKind, ledger1, nil),   //
		freshBallot(privateItem.Key.Location, 0, altitude-1, iteration, ballotKind, ledger1, nil),   //
		freshBallot(privateItem.Key.Location, 0, altitude-2, iteration+4, ballotKind, ledger1, nil), //
		freshBallot(privateItem.Key.Location, 0, altitude, iteration, ballotKind, ledger2, nil),     //
	}

	for _, c := range scenarios {
		cpb := c.TowardSchema()
		err = privateItem.AttestBallot("REDACTED", cpb)
		assert.Error(err, "REDACTED")
	}

	//
	sig := ballot.Notation
	ballot.Timestamp = ballot.Timestamp.Add(time.Duration(1000))
	err = privateItem.AttestBallot("REDACTED", v)
	assert.NoError(err)
	assert.Equal(sig, ballot.Notation)
}

func VerifyAttestNomination(t *testing.T) {
	affirm := assert.New(t)

	privateItem, _, _ := freshVerifyRecordPRV(t)

	randomoctets := commitrand.Octets(tenderminthash.Extent)
	randomoctets2 := commitrand.Octets(tenderminthash.Extent)

	ledger1 := kinds.LedgerUUID{
		Digest:          randomoctets,
		FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 5, Digest: randomoctets},
	}
	ledger2 := kinds.LedgerUUID{
		Digest:          randomoctets2,
		FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 10, Digest: randomoctets2},
	}
	altitude, iteration := int64(10), int32(1)

	//
	nomination := freshNomination(altitude, iteration, ledger1)
	pbp := nomination.TowardSchema()
	err := privateItem.AttestNomination("REDACTED", pbp)
	assert.NoError(err, "REDACTED")

	//
	err = privateItem.AttestNomination("REDACTED", pbp)
	assert.NoError(err, "REDACTED")

	//
	scenarios := []*kinds.Nomination{
		freshNomination(altitude, iteration-1, ledger1),   //
		freshNomination(altitude-1, iteration, ledger1),   //
		freshNomination(altitude-2, iteration+4, ledger1), //
		freshNomination(altitude, iteration, ledger2),     //
	}

	for _, c := range scenarios {
		err = privateItem.AttestNomination("REDACTED", c.TowardSchema())
		assert.Error(err, "REDACTED")
	}

	//
	sig := nomination.Notation
	nomination.Timestamp = nomination.Timestamp.Add(time.Duration(1000))
	err = privateItem.AttestNomination("REDACTED", pbp)
	assert.NoError(err)
	assert.Equal(sig, nomination.Notation)
}

func VerifyDeviateViaTimestamp(t *testing.T) {
	transientTokenRecord, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	transientStatusRecord, err := os.CreateTemp("REDACTED", "REDACTED")
	require.Nil(t, err)

	privateItem := ProduceRecordPRV(transientTokenRecord.Name(), transientStatusRecord.Name())
	randomoctets := commitrand.Octets(tenderminthash.Extent)
	ledger1 := kinds.LedgerUUID{Digest: randomoctets, FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 5, Digest: randomoctets}}
	altitude, iteration := int64(10), int32(1)
	successionUUID := "REDACTED"

	//
	{
		nomination := freshNomination(altitude, iteration, ledger1)
		pb := nomination.TowardSchema()
		err := privateItem.AttestNomination(successionUUID, pb)
		assert.NoError(t, err, "REDACTED")
		attestOctets := kinds.NominationAttestOctets(successionUUID, pb)

		sig := nomination.Notation
		momentImprint := nomination.Timestamp

		//
		pb.Timestamp = pb.Timestamp.Add(time.Millisecond)
		var blankSignature []byte
		nomination.Notation = blankSignature
		err = privateItem.AttestNomination("REDACTED", pb)
		assert.NoError(t, err, "REDACTED")

		assert.Equal(t, momentImprint, pb.Timestamp)
		assert.Equal(t, attestOctets, kinds.NominationAttestOctets(successionUUID, pb))
		assert.Equal(t, sig, nomination.Notation)
	}

	//
	{
		ballotKind := commitchema.PreballotKind
		ledgerUUID := kinds.LedgerUUID{Digest: randomoctets, FragmentAssignHeading: kinds.FragmentAssignHeading{}}
		ballot := freshBallot(privateItem.Key.Location, 0, altitude, iteration, ballotKind, ledgerUUID, nil)
		v := ballot.TowardSchema()
		err := privateItem.AttestBallot("REDACTED", v)
		assert.NoError(t, err, "REDACTED")

		attestOctets := kinds.BallotAttestOctets(successionUUID, v)
		sig := v.Notation
		addnSignature := v.AdditionNotation
		momentImprint := ballot.Timestamp

		//
		v.Timestamp = v.Timestamp.Add(time.Millisecond)
		var blankSignature []byte
		v.Notation = blankSignature
		v.AdditionNotation = blankSignature
		err = privateItem.AttestBallot("REDACTED", v)
		assert.NoError(t, err, "REDACTED")

		assert.Equal(t, momentImprint, v.Timestamp)
		assert.Equal(t, attestOctets, kinds.BallotAttestOctets(successionUUID, v))
		assert.Equal(t, sig, v.Notation)
		assert.Equal(t, addnSignature, v.AdditionNotation)
	}
}

func VerifyBallotAdditionsExistInvariablyNotated(t *testing.T) {
	privateItem, _, _ := freshVerifyRecordPRV(t)
	publicToken, err := privateItem.ObtainPublicToken()
	assert.NoError(t, err)

	ledger := kinds.LedgerUUID{
		Digest:          commitrand.Octets(tenderminthash.Extent),
		FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 5, Digest: commitrand.Octets(tenderminthash.Extent)},
	}

	altitude, iteration := int64(10), int32(1)
	ballotKind := commitchema.PreendorseKind

	//
	ballot1 := freshBallot(privateItem.Key.Location, 0, altitude, iteration, ballotKind, ledger, nil)
	verificationschema1 := ballot1.TowardSchema()

	err = privateItem.AttestBallot("REDACTED", verificationschema1)
	assert.NoError(t, err, "REDACTED")
	assert.NotNil(t, verificationschema1.AdditionNotation)

	verificationenvelope1 := kinds.BallotAdditionAttestOctets("REDACTED", verificationschema1)
	assert.True(t, publicToken.ValidateNotation(verificationenvelope1, verificationschema1.AdditionNotation))

	//
	//
	ballot2 := ballot1.Duplicate()
	ballot2.Addition = []byte("REDACTED")
	verificationschema2 := ballot2.TowardSchema()

	err = privateItem.AttestBallot("REDACTED", verificationschema2)
	assert.NoError(t, err, "REDACTED")

	//
	//
	//
	//
	verificationenvelope2 := kinds.BallotAdditionAttestOctets("REDACTED", verificationschema2)
	assert.True(t, publicToken.ValidateNotation(verificationenvelope2, verificationschema2.AdditionNotation))
	assert.False(t, publicToken.ValidateNotation(verificationenvelope1, verificationschema2.AdditionNotation))

	//
	//
	anticipatedTimestamp := verificationschema2.Timestamp

	verificationschema2.Timestamp = verificationschema2.Timestamp.Add(time.Millisecond)
	verificationschema2.Notation = nil
	verificationschema2.AdditionNotation = nil

	err = privateItem.AttestBallot("REDACTED", verificationschema2)
	assert.NoError(t, err, "REDACTED")
	assert.Equal(t, anticipatedTimestamp, verificationschema2.Timestamp)

	verificationenvelope3 := kinds.BallotAdditionAttestOctets("REDACTED", verificationschema2)
	assert.True(t, publicToken.ValidateNotation(verificationenvelope3, verificationschema2.AdditionNotation))
	assert.False(t, publicToken.ValidateNotation(verificationenvelope1, verificationschema2.AdditionNotation))
}

func freshBallot(location kinds.Location, idx int32, altitude int64, iteration int32,
	typ commitchema.AttestedSignalKind, ledgerUUID kinds.LedgerUUID, addition []byte,
) *kinds.Ballot {
	return &kinds.Ballot{
		AssessorLocation: location,
		AssessorOrdinal:   idx,
		Altitude:           altitude,
		Iteration:            iteration,
		Kind:             typ,
		Timestamp:        committime.Now(),
		LedgerUUID:          ledgerUUID,
		Addition:        addition,
	}
}

func freshNomination(altitude int64, iteration int32, ledgerUUID kinds.LedgerUUID) *kinds.Nomination {
	return &kinds.Nomination{
		Altitude:    altitude,
		Iteration:     iteration,
		LedgerUUID:   ledgerUUID,
		Timestamp: committime.Now(),
	}
}

func freshVerifyRecordPRV(t *testing.T) (*RecordPRV, string, string) {
	transientTokenRecord, err := os.CreateTemp(t.TempDir(), "REDACTED")
	require.NoError(t, err)
	transientStatusRecord, err := os.CreateTemp(t.TempDir(), "REDACTED")
	require.NoError(t, err)

	privateItem := ProduceRecordPRV(transientTokenRecord.Name(), transientStatusRecord.Name())

	return privateItem, transientTokenRecord.Name(), transientStatusRecord.Name()
}
