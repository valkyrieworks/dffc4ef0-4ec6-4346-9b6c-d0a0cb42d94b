package hashmap

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/verify"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
)

type verifyRecord []byte

func (tI verifyRecord) Digest() []byte {
	return []byte(tI)
}

func VerifyDigestOriginatingOctetSegments(t *testing.T) {
	verifycases := map[string]struct {
		segments     [][]byte
		anticipateDigest string //
	}{
		"REDACTED":          {nil, "REDACTED"},
		"REDACTED":        {[][]byte{}, "REDACTED"},
		"REDACTED":       {[][]byte{{1, 2, 3}}, "REDACTED"},
		"REDACTED": {[][]byte{{}}, "REDACTED"},
		"REDACTED":          {[][]byte{{1, 2, 3}, {4, 5, 6}}, "REDACTED"},
		"REDACTED": {
			[][]byte{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			"REDACTED",
		},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			digest := DigestOriginatingOctetSegments(tc.segments)
			assert.Equal(t, tc.anticipateDigest, hex.EncodeToString(digest))
		})
	}
}

func VerifyAttestation(t *testing.T) {
	//
	originDigest, attestations := AttestationsOriginatingOctetSegments([][]byte{})
	require.Equal(t, "REDACTED", hex.EncodeToString(originDigest))
	require.Empty(t, attestations)

	sum := 100

	elements := make([][]byte, sum)
	for i := 0; i < sum; i++ {
		elements[i] = verifyRecord(commitrand.Octets(tenderminthash.Extent))
	}

	originDigest = DigestOriginatingOctetSegments(elements)

	originDigest2, attestations := AttestationsOriginatingOctetSegments(elements)

	require.Equal(t, originDigest, originDigest2, "REDACTED", originDigest, originDigest2)

	//
	for i, record := range elements {
		attestation := attestations[i]

		//
		require.EqualValues(t, attestation.Ordinal, i, "REDACTED", attestation.Ordinal, i)

		require.EqualValues(t, attestation.Sum, sum, "REDACTED", attestation.Sum, sum)

		//
		err := attestation.Validate(originDigest, record)
		require.NoError(t, err, "REDACTED", err)

		//
		initialKin := attestation.Kin
		attestation.Kin = append(attestation.Kin, commitrand.Octets(32))
		err = attestation.Validate(originDigest, record)
		require.Error(t, err, "REDACTED")

		attestation.Kin = initialKin

		//
		attestation.Kin = attestation.Kin[0 : len(attestation.Kin)-1]
		err = attestation.Validate(originDigest, record)
		require.Error(t, err, "REDACTED")

		attestation.Kin = initialKin

		//
		err = attestation.Validate(originDigest, verify.TransformOctetSegment(record))
		require.Error(t, err, "REDACTED")

		//
		err = attestation.Validate(verify.TransformOctetSegment(originDigest), record)
		require.Error(t, err, "REDACTED")
	}
}

func VerifyDigestOptions(t *testing.T) {
	sum := 100

	elements := make([][]byte, sum)
	for i := 0; i < sum; i++ {
		elements[i] = verifyRecord(commitrand.Octets(tenderminthash.Extent))
	}

	originDigest1 := DigestOriginatingOctetSegmentsRecursive(elements)
	originDigest2 := DigestOriginatingOctetSegments(elements)
	require.Equal(t, originDigest1, originDigest2, "REDACTED", originDigest1, originDigest2)
}

func AssessmentDigestOptions(b *testing.B) {
	sum := 100

	elements := make([][]byte, sum)
	for i := 0; i < sum; i++ {
		elements[i] = verifyRecord(commitrand.Octets(tenderminthash.Extent))
	}

	b.ResetTimer()
	b.Run("REDACTED", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DigestOriginatingOctetSegments(elements)
		}
	})

	b.Run("REDACTED", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DigestOriginatingOctetSegmentsRecursive(elements)
		}
	})
}

func Verify_obtaindivisionpoint(t *testing.T) {
	verifies := []struct {
		magnitude int64
		desire   int64
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 4},
		{10, 8},
		{20, 16},
		{100, 64},
		{255, 128},
		{256, 128},
		{257, 256},
	}
	for _, tt := range verifies {
		got := obtainPartitionNode(tt.magnitude)
		require.EqualValues(t, tt.desire, got, "REDACTED", tt.magnitude, got, tt.desire)
	}
}
