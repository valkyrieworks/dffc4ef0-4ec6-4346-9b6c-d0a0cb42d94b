package kinds

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	engineseed "github.com/valkyrieworks/utils/random"
	vtest "github.com/valkyrieworks/utils/verify"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

func createTrans(cnt, volume int) Txs {
	txs := make(Txs, cnt)
	for i := 0; i < cnt; i++ {
		txs[i] = engineseed.Octets(volume)
	}
	return txs
}

func VerifyTransferOrdinal(t *testing.T) {
	for i := 0; i < 20; i++ {
		txs := createTrans(15, 60)
		for j := 0; j < len(txs); j++ {
			tx := txs[j]
			idx := txs.Ordinal(tx)
			assert.Equal(t, j, idx)
		}
		assert.Equal(t, -1, txs.Ordinal(nil))
		assert.Equal(t, -1, txs.Ordinal(Tx("REDACTED")))
	}
}

func VerifyTransferOrdinalByDigest(t *testing.T) {
	for i := 0; i < 20; i++ {
		txs := createTrans(15, 60)
		for j := 0; j < len(txs); j++ {
			tx := txs[j]
			idx := txs.OrdinalByDigest(tx.Digest())
			assert.Equal(t, j, idx)
		}
		assert.Equal(t, -1, txs.OrdinalByDigest(nil))
		assert.Equal(t, -1, txs.OrdinalByDigest(Tx("REDACTED").Digest()))
	}
}

func VerifySoundTransferEvidence(t *testing.T) {
	scenarios := []struct {
		txs Txs
	}{
		{Txs{{1, 4, 34, 87, 163, 1}}},
		{Txs{{5, 56, 165, 2}, {4, 77}}},
		{Txs{Tx("REDACTED"), Tx("REDACTED"), Tx("REDACTED")}},
		{createTrans(20, 5)},
		{createTrans(7, 81)},
		{createTrans(61, 15)},
	}

	for h, tc := range scenarios {
		txs := tc.txs
		origin := txs.Digest()
		//
		for i := range txs {
			tx := []byte(txs[i])
			evidence := txs.Attestation(i)
			assert.EqualValues(t, i, evidence.Attestation.Ordinal, "REDACTED", h, i)
			assert.EqualValues(t, len(txs), evidence.Attestation.Sum, "REDACTED", h, i)
			assert.EqualValues(t, origin, evidence.OriginDigest, "REDACTED", h, i)
			assert.EqualValues(t, tx, evidence.Data, "REDACTED", h, i)
			assert.EqualValues(t, txs[i].Digest(), evidence.Element(), "REDACTED", h, i)
			assert.Nil(t, evidence.Certify(origin), "REDACTED", h, i)
			assert.NotNil(t, evidence.Certify([]byte("REDACTED")), "REDACTED", h, i)

			//
			var (
				p2  TransferEvidence
				pb2 engineproto.TransferEvidence
			)
			pbEvidence := evidence.ToSchema()
			bin, err := pbEvidence.Serialize()
			require.NoError(t, err)

			err = pb2.Unserialize(bin)
			require.NoError(t, err)

			p2, err = TransferEvidenceFromSchema(pb2)
			if assert.Nil(t, err, "REDACTED", h, i, err) {
				assert.Nil(t, p2.Certify(origin), "REDACTED", h, i)
			}
		}
	}
}

func VerifyTransferEvidenceImmutable(t *testing.T) {
	//
	for i := 0; i < 40; i++ {
		verifyTransferEvidenceImmutable(t)
	}
}

func verifyTransferEvidenceImmutable(t *testing.T) {
	//
	txs := createTrans(randomInteger(2, 100), randomInteger(16, 128))
	origin := txs.Digest()
	i := randomInteger(0, len(txs)-1)
	evidence := txs.Attestation(i)

	//
	assert.Nil(t, evidence.Certify(origin))
	pbEvidence := evidence.ToSchema()
	bin, err := pbEvidence.Serialize()
	require.NoError(t, err)

	//
	for j := 0; j < 500; j++ {
		bad := vtest.TransformOctetSegment(bin)
		if !bytes.Equal(bad, bin) {
			affirmFlawedEvidence(t, origin, bad, evidence)
		}
	}
}

//
func affirmFlawedEvidence(t *testing.T, origin []byte, bad []byte, sound TransferEvidence) {
	var (
		evidence   TransferEvidence
		pbEvidence engineproto.TransferEvidence
	)
	err := pbEvidence.Unserialize(bad)
	if err == nil {
		evidence, err = TransferEvidenceFromSchema(pbEvidence)
		if err == nil {
			err = evidence.Certify(origin)
			if err == nil {
				//
				//
				//
				//
				assert.NotEqual(t, evidence.Attestation.Sum, sound.Attestation.Sum, "REDACTED", evidence, sound)
			}
		}
	}
}

func randomInteger(low, elevated int) int {
	return rand.Intn(elevated-low) + low
}
