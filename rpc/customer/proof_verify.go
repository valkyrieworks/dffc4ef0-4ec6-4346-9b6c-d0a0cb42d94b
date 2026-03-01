package cust_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	rpcoverify "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
var fallbackVerifyMoment = time.Date(2018, 10, 10, 8, 20, 13, 695936996, time.UTC)

func freshProof(t *testing.T, val *privatevalue.RecordPRV,
	ballot *kinds.Ballot, ballot2 *kinds.Ballot,
	successionUUID string,
) *kinds.ReplicatedBallotProof {
	var err error

	v := ballot.TowardSchema()
	v2 := ballot2.TowardSchema()

	ballot.Notation, err = val.Key.PrivateToken.Attest(kinds.BallotAttestOctets(successionUUID, v))
	require.NoError(t, err)

	ballot2.Notation, err = val.Key.PrivateToken.Attest(kinds.BallotAttestOctets(successionUUID, v2))
	require.NoError(t, err)

	assessor := kinds.FreshAssessor(val.Key.PublicToken, 10)
	itemAssign := kinds.FreshAssessorAssign([]*kinds.Assessor{assessor})

	ev, err := kinds.FreshReplicatedBallotProof(ballot, ballot2, fallbackVerifyMoment, itemAssign)
	require.NoError(t, err)
	return ev
}

func createProofs(
	t *testing.T,
	val *privatevalue.RecordPRV,
	successionUUID string,
) (precise *kinds.ReplicatedBallotProof, simulations []*kinds.ReplicatedBallotProof) {
	ballot := kinds.Ballot{
		AssessorLocation: val.Key.Location,
		AssessorOrdinal:   0,
		Altitude:           1,
		Iteration:            0,
		Kind:             commitchema.PreballotKind,
		Timestamp:        fallbackVerifyMoment,
		LedgerUUID: kinds.LedgerUUID{
			Digest: tenderminthash.Sum(commitrand.Octets(tenderminthash.Extent)),
			FragmentAssignHeading: kinds.FragmentAssignHeading{
				Sum: 1000,
				Digest:  tenderminthash.Sum([]byte("REDACTED")),
			},
		},
	}

	ballot2 := ballot
	ballot2.LedgerUUID.Digest = tenderminthash.Sum([]byte("REDACTED"))
	precise = freshProof(t, val, &ballot, &ballot2, successionUUID)

	simulations = make([]*kinds.ReplicatedBallotProof, 0)

	//
	{
		v := ballot2
		v.AssessorLocation = []byte("REDACTED")
		simulations = append(simulations, freshProof(t, val, &ballot, &v, successionUUID))
	}

	//
	{
		v := ballot2
		v.Altitude = ballot.Altitude + 1
		simulations = append(simulations, freshProof(t, val, &ballot, &v, successionUUID))
	}

	//
	{
		v := ballot2
		v.Iteration = ballot.Iteration + 1
		simulations = append(simulations, freshProof(t, val, &ballot, &v, successionUUID))
	}

	//
	{
		v := ballot2
		v.Kind = commitchema.PreendorseKind
		simulations = append(simulations, freshProof(t, val, &ballot, &v, successionUUID))
	}

	//
	{
		v := ballot
		simulations = append(simulations, freshProof(t, val, &ballot, &v, successionUUID))
	}

	return precise, simulations
}

func Verifyballotevidence_Replicatedballotevidence(t *testing.T) {
	var (
		settings  = rpcoverify.FetchSettings()
		successionUUID = verify.FallbackVerifySuccessionUUID
		pv      = privatevalue.FetchEitherProduceRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord())
	)

	for i, c := range FetchCustomers() {
		precise, simulations := createProofs(t, pv, successionUUID)
		t.Logf("REDACTED", i)

		outcome, err := c.MulticastProof(context.Background(), precise)
		require.NoError(t, err, "REDACTED", precise)
		assert.Equal(t, precise.Digest(), outcome.Digest, "REDACTED")

		condition, err := c.Condition(context.Background())
		require.NoError(t, err)
		err = customer.PauseForeachAltitude(c, condition.ChronizeDetails.NewestLedgerAltitude+2, nil)
		require.NoError(t, err)

		curve25519key := pv.Key.PublicToken.(edwards25519.PublicToken)
		plainkey := curve25519key.Octets()
		outcome2, err := c.IfaceInquire(context.Background(), "REDACTED", plainkey)
		require.NoError(t, err)
		queryresp := outcome2.Reply
		require.True(t, queryresp.EqualsOKAY())

		var v iface.AssessorRevise
		err = iface.FetchArtifact(bytes.NewReader(queryresp.Datum), &v)
		require.NoError(t, err, "REDACTED", queryresp.Datum)

		pk, err := cryptocode.PublicTokenOriginatingSchema(v.PublicToken)
		require.NoError(t, err)

		require.EqualValues(t, plainkey, pk, "REDACTED", string(queryresp.Datum))
		require.Equal(t, int64(9), v.Potency, "REDACTED", string(queryresp.Datum))

		for _, mocked := range simulations {
			_, err := c.MulticastProof(context.Background(), mocked)
			require.Error(t, err, "REDACTED", mocked)
		}
	}
}

func VerifyMulticastBlankProof(t *testing.T) {
	for _, c := range FetchCustomers() {
		_, err := c.MulticastProof(context.Background(), nil)
		assert.Error(t, err)
	}
}
