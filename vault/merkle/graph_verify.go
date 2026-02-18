package merkle

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/utils/verify"

	"github.com/valkyrieworks/vault/comethash"
)

type verifyItem []byte

func (tI verifyItem) Digest() []byte {
	return []byte(tI)
}

func VerifyDigestFromOctetSegments(t *testing.T) {
	verifyscenarios := map[string]struct {
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
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			digest := DigestFromOctetSegments(tc.segments)
			assert.Equal(t, tc.anticipateDigest, hex.EncodeToString(digest))
		})
	}
}

func VerifyEvidence(t *testing.T) {
	//
	originDigest, evidences := EvidencesFromOctetSegments([][]byte{})
	require.Equal(t, "REDACTED", hex.EncodeToString(originDigest))
	require.Empty(t, evidences)

	sum := 100

	items := make([][]byte, sum)
	for i := 0; i < sum; i++ {
		items[i] = verifyItem(engineseed.Octets(comethash.Volume))
	}

	originDigest = DigestFromOctetSegments(items)

	originDigest2, evidences := EvidencesFromOctetSegments(items)

	require.Equal(t, originDigest, originDigest2, "REDACTED", originDigest, originDigest2)

	//
	for i, item := range items {
		evidence := evidences[i]

		//
		require.EqualValues(t, evidence.Ordinal, i, "REDACTED", evidence.Ordinal, i)

		require.EqualValues(t, evidence.Sum, sum, "REDACTED", evidence.Sum, sum)

		//
		err := evidence.Validate(originDigest, item)
		require.NoError(t, err, "REDACTED", err)

		//
		origKin := evidence.Kin
		evidence.Kin = append(evidence.Kin, engineseed.Octets(32))
		err = evidence.Validate(originDigest, item)
		require.Error(t, err, "REDACTED")

		evidence.Kin = origKin

		//
		evidence.Kin = evidence.Kin[0 : len(evidence.Kin)-1]
		err = evidence.Validate(originDigest, item)
		require.Error(t, err, "REDACTED")

		evidence.Kin = origKin

		//
		err = evidence.Validate(originDigest, verify.TransformOctetSegment(item))
		require.Error(t, err, "REDACTED")

		//
		err = evidence.Validate(verify.TransformOctetSegment(originDigest), item)
		require.Error(t, err, "REDACTED")
	}
}

func VerifyDigestOptions(t *testing.T) {
	sum := 100

	items := make([][]byte, sum)
	for i := 0; i < sum; i++ {
		items[i] = verifyItem(engineseed.Octets(comethash.Volume))
	}

	originDigest1 := DigestFromOctetSegmentsRecursive(items)
	originDigest2 := DigestFromOctetSegments(items)
	require.Equal(t, originDigest1, originDigest2, "REDACTED", originDigest1, originDigest2)
}

func CriterionDigestOptions(b *testing.B) {
	sum := 100

	items := make([][]byte, sum)
	for i := 0; i < sum; i++ {
		items[i] = verifyItem(engineseed.Octets(comethash.Volume))
	}

	b.ResetTimer()
	b.Run("REDACTED", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DigestFromOctetSegments(items)
		}
	})

	b.Run("REDACTED", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = DigestFromOctetSegmentsRecursive(items)
		}
	})
}

func Verify_fetchdivisionpoint(t *testing.T) {
	verifies := []struct {
		extent int64
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
		got := fetchDivideSpot(tt.extent)
		require.EqualValues(t, tt.desire, got, "REDACTED", tt.extent, got, tt.desire)
	}
}
