package kinds

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/merkle"
	engineseed "github.com/valkyrieworks/utils/random"
)

const (
	verifySegmentVolume = 65536 //
)

func VerifySimpleSectionCollection(t *testing.T) {
	//
	nSections := 100
	data := engineseed.Octets(verifySegmentVolume * nSections)
	sectionCollection := NewSegmentCollectionFromData(data, verifySegmentVolume)

	assert.NotEmpty(t, sectionCollection.Digest())
	assert.EqualValues(t, nSections, sectionCollection.Sum())
	assert.Equal(t, nSections, sectionCollection.BitList().Volume())
	assert.True(t, sectionCollection.DigestsTo(sectionCollection.Digest()))
	assert.True(t, sectionCollection.IsFinished())
	assert.EqualValues(t, nSections, sectionCollection.Number())
	assert.EqualValues(t, verifySegmentVolume*nSections, sectionCollection.OctetVolume())

	//
	sectionSet2 := NewSegmentCollectionFromHeading(sectionCollection.Heading())

	assert.True(t, sectionSet2.HasHeading(sectionCollection.Heading()))
	for i := 0; i < int(sectionCollection.Sum()); i++ {
		segment := sectionCollection.FetchSegment(i)
		//
		appended, err := sectionSet2.AppendSegment(segment)
		if !appended || err != nil {
			t.Errorf("REDACTED", i, err)
		}
	}
	//
	appended, err := sectionSet2.AppendSegment(&Segment{Ordinal: 10000})
	assert.False(t, appended)
	assert.Error(t, err)
	//
	appended, err = sectionSet2.AppendSegment(sectionSet2.FetchSegment(0))
	assert.False(t, appended)
	assert.Nil(t, err)

	assert.Equal(t, sectionCollection.Digest(), sectionSet2.Digest())
	assert.EqualValues(t, nSections, sectionSet2.Sum())
	assert.EqualValues(t, nSections*verifySegmentVolume, sectionCollection.OctetVolume())
	assert.True(t, sectionSet2.IsFinished())

	//
	data2scanner := sectionSet2.FetchScanner()
	data2, err := io.ReadAll(data2scanner)
	require.NoError(t, err)

	assert.Equal(t, data, data2)
}

func VerifyIncorrectEvidence(t *testing.T) {
	//
	data := engineseed.Octets(verifySegmentVolume * 100)
	sectionCollection := NewSegmentCollectionFromData(data, verifySegmentVolume)

	//
	sectionSet2 := NewSegmentCollectionFromHeading(sectionCollection.Heading())

	//
	segment := sectionCollection.FetchSegment(0)
	segment.Attestation.Kin[0][0] += byte(0x01)
	appended, err := sectionSet2.AppendSegment(segment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}

	//
	segment = sectionCollection.FetchSegment(1)
	segment.Octets[0] += byte(0x01)
	appended, err = sectionSet2.AppendSegment(segment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}

	//
	segment = sectionCollection.FetchSegment(2)
	segment.Attestation.Ordinal = 1
	appended, err = sectionSet2.AppendSegment(segment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}

	//
	segment = sectionCollection.FetchSegment(3)
	segment.Attestation.Sum = int64(sectionCollection.Sum() - 1)
	appended, err = sectionSet2.AppendSegment(segment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}
}

func VerifySectionCollectionHeadingCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel              string
		distortSectionCollectionHeading func(*SegmentAssignHeading)
		anticipateErr             bool
	}{
		{"REDACTED", func(psHeading *SegmentAssignHeading) {}, false},
		{"REDACTED", func(psHeading *SegmentAssignHeading) { psHeading.Digest = make([]byte, 1) }, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			data := engineseed.Octets(verifySegmentVolume * 100)
			ps := NewSegmentCollectionFromData(data, verifySegmentVolume)
			psHeading := ps.Heading()
			tc.distortSectionCollectionHeading(&psHeading)
			assert.Equal(t, tc.anticipateErr, psHeading.CertifySimple() != nil, "REDACTED")
		})
	}
}

func Verifypart_Verifybasic(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel     string
		distortSection func(*Segment)
		anticipateErr    bool
	}{
		{"REDACTED", func(pt *Segment) {}, false},
		{"REDACTED", func(pt *Segment) { pt.Octets = make([]byte, LedgerSegmentVolumeOctets+1) }, true},
		{"REDACTED", func(pt *Segment) {
			pt.Ordinal = 1
			pt.Octets = make([]byte, LedgerSegmentVolumeOctets-1)
			pt.Attestation.Sum = 2
			pt.Attestation.Ordinal = 1
		}, false},
		{"REDACTED", func(pt *Segment) {
			pt.Ordinal = 0
			pt.Octets = make([]byte, LedgerSegmentVolumeOctets-1)
			pt.Attestation.Sum = 2
		}, true},
		{"REDACTED", func(pt *Segment) {
			pt.Attestation = merkle.Attestation{
				Sum:    2,
				Ordinal:    1,
				NodeDigest: make([]byte, 1024*1024),
			}
			pt.Ordinal = 1
		}, true},
		{"REDACTED", func(pt *Segment) {
			pt.Ordinal = 1
			pt.Attestation.Ordinal = 0
		}, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			data := engineseed.Octets(verifySegmentVolume * 100)
			ps := NewSegmentCollectionFromData(data, verifySegmentVolume)
			segment := ps.FetchSegment(0)
			tc.distortSection(segment)
			assert.Equal(t, tc.anticipateErr, segment.CertifySimple() != nil, "REDACTED")
		})
	}
}

func VerifyParallelCollectionHeadingSchemaBuffer(t *testing.T) {
	verifyScenarios := []struct {
		msg     string
		ps1     *SegmentAssignHeading
		expirationPass bool
	}{
		{"REDACTED", &SegmentAssignHeading{}, true},
		{
			"REDACTED",
			&SegmentAssignHeading{Sum: 1, Digest: []byte("REDACTED")}, true,
		},
	}

	for _, tc := range verifyScenarios {
		schemaLedgerUID := tc.ps1.ToSchema()

		psh, err := SegmentAssignHeadingFromSchema(&schemaLedgerUID)
		if tc.expirationPass {
			require.Equal(t, tc.ps1, psh, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func VerifySectionSchemaBuffer(t *testing.T) {
	evidence := merkle.Attestation{
		Sum:    1,
		Ordinal:    1,
		NodeDigest: engineseed.Octets(32),
	}
	verifyScenarios := []struct {
		msg     string
		ps1     *Segment
		expirationPass bool
	}{
		{"REDACTED", &Segment{}, false},
		{"REDACTED", nil, false},
		{
			"REDACTED",
			&Segment{Ordinal: 1, Octets: engineseed.Octets(32), Attestation: evidence}, true,
		},
	}

	for _, tc := range verifyScenarios {
		schema, err := tc.ps1.ToSchema()
		if tc.expirationPass {
			require.NoError(t, err, tc.msg)
		}

		p, err := SegmentFromSchema(schema)
		if tc.expirationPass {
			require.NoError(t, err)
			require.Equal(t, tc.ps1, p, tc.msg)
		}
	}
}

func CriterionCreateSectionCollection(b *testing.B) {
	for nSections := 1; nSections <= 5; nSections++ {
		b.Run(fmt.Sprintf("REDACTED", nSections), func(b *testing.B) {
			data := engineseed.Octets(verifySegmentVolume * nSections)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				NewSegmentCollectionFromData(data, verifySegmentVolume)
			}
		})
	}
}
