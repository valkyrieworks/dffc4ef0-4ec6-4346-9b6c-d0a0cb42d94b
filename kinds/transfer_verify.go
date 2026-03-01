package kinds

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	agreementtest "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/verify"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

func createTrans(cnt, extent int) Txs {
	txs := make(Txs, cnt)
	for i := 0; i < cnt; i++ {
		txs[i] = commitrand.Octets(extent)
	}
	return txs
}

func VerifyTransferPosition(t *testing.T) {
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

func VerifyTransferPositionViaDigest(t *testing.T) {
	for i := 0; i < 20; i++ {
		txs := createTrans(15, 60)
		for j := 0; j < len(txs); j++ {
			tx := txs[j]
			idx := txs.PositionViaDigest(tx.Digest())
			assert.Equal(t, j, idx)
		}
		assert.Equal(t, -1, txs.PositionViaDigest(nil))
		assert.Equal(t, -1, txs.PositionViaDigest(Tx("REDACTED").Digest()))
	}
}

func VerifySoundTransferAttestation(t *testing.T) {
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
			attestation := txs.Attestation(i)
			assert.EqualValues(t, i, attestation.Attestation.Ordinal, "REDACTED", h, i)
			assert.EqualValues(t, len(txs), attestation.Attestation.Sum, "REDACTED", h, i)
			assert.EqualValues(t, origin, attestation.OriginDigest, "REDACTED", h, i)
			assert.EqualValues(t, tx, attestation.Data, "REDACTED", h, i)
			assert.EqualValues(t, txs[i].Digest(), attestation.Node(), "REDACTED", h, i)
			assert.Nil(t, attestation.Certify(origin), "REDACTED", h, i)
			assert.NotNil(t, attestation.Certify([]byte("REDACTED")), "REDACTED", h, i)

			//
			var (
				p2  TransferAttestation
				pb2 commitchema.TransferAttestation
			)
			bufferAttestation := attestation.TowardSchema()
			bin, err := bufferAttestation.Serialize()
			require.NoError(t, err)

			err = pb2.Decode(bin)
			require.NoError(t, err)

			p2, err = TransferAttestationOriginatingSchema(pb2)
			if assert.Nil(t, err, "REDACTED", h, i, err) {
				assert.Nil(t, p2.Certify(origin), "REDACTED", h, i)
			}
		}
	}
}

func VerifyTransferAttestationImmutable(t *testing.T) {
	//
	for i := 0; i < 40; i++ {
		verifyTransferAttestationImmutable(t)
	}
}

func verifyTransferAttestationImmutable(t *testing.T) {
	//
	txs := createTrans(arbitraryInteger(2, 100), arbitraryInteger(16, 128))
	origin := txs.Digest()
	i := arbitraryInteger(0, len(txs)-1)
	attestation := txs.Attestation(i)

	//
	assert.Nil(t, attestation.Certify(origin))
	bufferAttestation := attestation.TowardSchema()
	bin, err := bufferAttestation.Serialize()
	require.NoError(t, err)

	//
	for j := 0; j < 500; j++ {
		bad := agreementtest.TransformOctetSegment(bin)
		if !bytes.Equal(bad, bin) {
			attestFlawedAttestation(t, origin, bad, attestation)
		}
	}
}

//
func attestFlawedAttestation(t *testing.T, origin []byte, bad []byte, valid TransferAttestation) {
	var (
		attestation   TransferAttestation
		bufferAttestation commitchema.TransferAttestation
	)
	err := bufferAttestation.Decode(bad)
	if err == nil {
		attestation, err = TransferAttestationOriginatingSchema(bufferAttestation)
		if err == nil {
			err = attestation.Certify(origin)
			if err == nil {
				//
				//
				//
				//
				assert.NotEqual(t, attestation.Attestation.Sum, valid.Attestation.Sum, "REDACTED", attestation, valid)
			}
		}
	}
}

func arbitraryInteger(low, tall int) int {
	return rand.Intn(tall-low) + low
}
