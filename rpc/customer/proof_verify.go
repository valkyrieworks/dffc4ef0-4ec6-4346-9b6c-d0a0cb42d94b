package agent_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/intrinsic/verify"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/rpc/customer"
	rpctest "github.com/valkyrieworks/rpc/verify"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
var standardVerifyTime = time.Date(2018, 10, 10, 8, 20, 13, 695936996, time.UTC)

func newProof(t *testing.T, val *privatekey.EntryPV,
	ballot *kinds.Ballot, ballot2 *kinds.Ballot,
	ledgerUID string,
) *kinds.ReplicatedBallotProof {
	var err error

	v := ballot.ToSchema()
	v2 := ballot2.ToSchema()

	ballot.Autograph, err = val.Key.PrivateKey.Attest(kinds.BallotAttestOctets(ledgerUID, v))
	require.NoError(t, err)

	ballot2.Autograph, err = val.Key.PrivateKey.Attest(kinds.BallotAttestOctets(ledgerUID, v2))
	require.NoError(t, err)

	ratifier := kinds.NewRatifier(val.Key.PublicKey, 10)
	valueCollection := kinds.NewRatifierCollection([]*kinds.Ratifier{ratifier})

	ev, err := kinds.NewReplicatedBallotProof(ballot, ballot2, standardVerifyTime, valueCollection)
	require.NoError(t, err)
	return ev
}

func createProofs(
	t *testing.T,
	val *privatekey.EntryPV,
	ledgerUID string,
) (accurate *kinds.ReplicatedBallotProof, mocks []*kinds.ReplicatedBallotProof) {
	ballot := kinds.Ballot{
		RatifierLocation: val.Key.Location,
		RatifierOrdinal:   0,
		Level:           1,
		Cycle:            0,
		Kind:             engineproto.PreballotKind,
		Timestamp:        standardVerifyTime,
		LedgerUID: kinds.LedgerUID{
			Digest: comethash.Sum(engineseed.Octets(comethash.Volume)),
			SegmentAssignHeading: kinds.SegmentAssignHeading{
				Sum: 1000,
				Digest:  comethash.Sum([]byte("REDACTED")),
			},
		},
	}

	ballot2 := ballot
	ballot2.LedgerUID.Digest = comethash.Sum([]byte("REDACTED"))
	accurate = newProof(t, val, &ballot, &ballot2, ledgerUID)

	mocks = make([]*kinds.ReplicatedBallotProof, 0)

	//
	{
		v := ballot2
		v.RatifierLocation = []byte("REDACTED")
		mocks = append(mocks, newProof(t, val, &ballot, &v, ledgerUID))
	}

	//
	{
		v := ballot2
		v.Level = ballot.Level + 1
		mocks = append(mocks, newProof(t, val, &ballot, &v, ledgerUID))
	}

	//
	{
		v := ballot2
		v.Cycle = ballot.Cycle + 1
		mocks = append(mocks, newProof(t, val, &ballot, &v, ledgerUID))
	}

	//
	{
		v := ballot2
		v.Kind = engineproto.PreendorseKind
		mocks = append(mocks, newProof(t, val, &ballot, &v, ledgerUID))
	}

	//
	{
		v := ballot
		mocks = append(mocks, newProof(t, val, &ballot, &v, ledgerUID))
	}

	return accurate, mocks
}

func Verifybroadcastevidence_Duplicateballotevidence(t *testing.T) {
	var (
		settings  = rpctest.FetchSettings()
		ledgerUID = verify.StandardVerifyLedgerUID
		pv      = privatekey.ImportOrGenerateEntryPV(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry())
	)

	for i, c := range FetchAgents() {
		accurate, mocks := createProofs(t, pv, ledgerUID)
		t.Logf("REDACTED", i)

		outcome, err := c.MulticastProof(context.Background(), accurate)
		require.NoError(t, err, "REDACTED", accurate)
		assert.Equal(t, accurate.Digest(), outcome.Digest, "REDACTED")

		state, err := c.Status(context.Background())
		require.NoError(t, err)
		err = customer.WaitForLevel(c, state.AlignDetails.NewestLedgerLevel+2, nil)
		require.NoError(t, err)

		ed25519key := pv.Key.PublicKey.(ed25519.PublicKey)
		rawkey := ed25519key.Octets()
		outcome2, err := c.IfaceInquire(context.Background(), "REDACTED", rawkey)
		require.NoError(t, err)
		inquiryout := outcome2.Reply
		require.True(t, inquiryout.IsOK())

		var v iface.RatifierModify
		err = iface.ScanSignal(bytes.NewReader(inquiryout.Item), &v)
		require.NoError(t, err, "REDACTED", inquiryout.Item)

		pk, err := cryptocode.PublicKeyFromSchema(v.PublicKey)
		require.NoError(t, err)

		require.EqualValues(t, rawkey, pk, "REDACTED", string(inquiryout.Item))
		require.Equal(t, int64(9), v.Energy, "REDACTED", string(inquiryout.Item))

		for _, mock := range mocks {
			_, err := c.MulticastProof(context.Background(), mock)
			require.Error(t, err, "REDACTED", mock)
		}
	}
}

func VerifyMulticastEmptyProof(t *testing.T) {
	for _, c := range FetchAgents() {
		_, err := c.MulticastProof(context.Background(), nil)
		assert.Error(t, err)
	}
}
