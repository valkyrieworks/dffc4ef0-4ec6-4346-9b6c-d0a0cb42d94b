package kinds

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

const (
	verifyFragmentExtent = 65536 //
)

func VerifyFundamentalFragmentAssign(t *testing.T) {
	//
	nthFragments := 100
	data := commitrand.Octets(verifyFragmentExtent * nthFragments)
	fragmentAssign := FreshFragmentAssignOriginatingData(data, verifyFragmentExtent)

	assert.NotEmpty(t, fragmentAssign.Digest())
	assert.EqualValues(t, nthFragments, fragmentAssign.Sum())
	assert.Equal(t, nthFragments, fragmentAssign.DigitSeries().Extent())
	assert.True(t, fragmentAssign.DigestsToward(fragmentAssign.Digest()))
	assert.True(t, fragmentAssign.EqualsFinish())
	assert.EqualValues(t, nthFragments, fragmentAssign.Tally())
	assert.EqualValues(t, verifyFragmentExtent*nthFragments, fragmentAssign.OctetExtent())

	//
	fragmentGroup2 := FreshFragmentAssignOriginatingHeading(fragmentAssign.Heading())

	assert.True(t, fragmentGroup2.OwnsHeading(fragmentAssign.Heading()))
	for i := 0; i < int(fragmentAssign.Sum()); i++ {
		fragment := fragmentAssign.ObtainFragment(i)
		//
		appended, err := fragmentGroup2.AppendFragment(fragment)
		if !appended || err != nil {
			t.Errorf("REDACTED", i, err)
		}
	}
	//
	appended, err := fragmentGroup2.AppendFragment(&Fragment{Ordinal: 10000})
	assert.False(t, appended)
	assert.Error(t, err)
	//
	appended, err = fragmentGroup2.AppendFragment(fragmentGroup2.ObtainFragment(0))
	assert.False(t, appended)
	assert.Nil(t, err)

	assert.Equal(t, fragmentAssign.Digest(), fragmentGroup2.Digest())
	assert.EqualValues(t, nthFragments, fragmentGroup2.Sum())
	assert.EqualValues(t, nthFragments*verifyFragmentExtent, fragmentAssign.OctetExtent())
	assert.True(t, fragmentGroup2.EqualsFinish())

	//
	data2fetcher := fragmentGroup2.ObtainFetcher()
	datum2, err := io.ReadAll(data2fetcher)
	require.NoError(t, err)

	assert.Equal(t, data, datum2)
}

func VerifyIncorrectAttestation(t *testing.T) {
	//
	data := commitrand.Octets(verifyFragmentExtent * 100)
	fragmentAssign := FreshFragmentAssignOriginatingData(data, verifyFragmentExtent)

	//
	fragmentGroup2 := FreshFragmentAssignOriginatingHeading(fragmentAssign.Heading())

	//
	fragment := fragmentAssign.ObtainFragment(0)
	fragment.Attestation.Kin[0][0] += byte(0x01)
	appended, err := fragmentGroup2.AppendFragment(fragment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}

	//
	fragment = fragmentAssign.ObtainFragment(1)
	fragment.Octets[0] += byte(0x01)
	appended, err = fragmentGroup2.AppendFragment(fragment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}

	//
	fragment = fragmentAssign.ObtainFragment(2)
	fragment.Attestation.Ordinal = 1
	appended, err = fragmentGroup2.AppendFragment(fragment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}

	//
	fragment = fragmentAssign.ObtainFragment(3)
	fragment.Attestation.Sum = int64(fragmentAssign.Sum() - 1)
	appended, err = fragmentGroup2.AppendFragment(fragment)
	if appended || err == nil {
		t.Errorf("REDACTED")
	}
}

func VerifyFragmentAssignHeadingCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias              string
		distortFragmentAssignHeading func(*FragmentAssignHeading)
		anticipateFault             bool
	}{
		{"REDACTED", func(processesHeading *FragmentAssignHeading) {}, false},
		{"REDACTED", func(processesHeading *FragmentAssignHeading) { processesHeading.Digest = make([]byte, 1) }, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			data := commitrand.Octets(verifyFragmentExtent * 100)
			ps := FreshFragmentAssignOriginatingData(data, verifyFragmentExtent)
			processesHeading := ps.Heading()
			tc.distortFragmentAssignHeading(&processesHeading)
			assert.Equal(t, tc.anticipateFault, processesHeading.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func Testsection_Certifyfundamental(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias     string
		distortFragment func(*Fragment)
		anticipateFault    bool
	}{
		{"REDACTED", func(pt *Fragment) {}, false},
		{"REDACTED", func(pt *Fragment) { pt.Octets = make([]byte, LedgerFragmentExtentOctets+1) }, true},
		{"REDACTED", func(pt *Fragment) {
			pt.Ordinal = 1
			pt.Octets = make([]byte, LedgerFragmentExtentOctets-1)
			pt.Attestation.Sum = 2
			pt.Attestation.Ordinal = 1
		}, false},
		{"REDACTED", func(pt *Fragment) {
			pt.Ordinal = 0
			pt.Octets = make([]byte, LedgerFragmentExtentOctets-1)
			pt.Attestation.Sum = 2
		}, true},
		{"REDACTED", func(pt *Fragment) {
			pt.Attestation = hashmap.Attestation{
				Sum:    2,
				Ordinal:    1,
				NodeDigest: make([]byte, 1024*1024),
			}
			pt.Ordinal = 1
		}, true},
		{"REDACTED", func(pt *Fragment) {
			pt.Ordinal = 1
			pt.Attestation.Ordinal = 0
		}, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			data := commitrand.Octets(verifyFragmentExtent * 100)
			ps := FreshFragmentAssignOriginatingData(data, verifyFragmentExtent)
			fragment := ps.ObtainFragment(0)
			tc.distortFragment(fragment)
			assert.Equal(t, tc.anticipateFault, fragment.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func VerifyParallelAssignHeadingSchemaArea(t *testing.T) {
	verifyScenarios := []struct {
		msg     string
		ps1     *FragmentAssignHeading
		expirationPhrase bool
	}{
		{"REDACTED", &FragmentAssignHeading{}, true},
		{
			"REDACTED",
			&FragmentAssignHeading{Sum: 1, Digest: []byte("REDACTED")}, true,
		},
	}

	for _, tc := range verifyScenarios {
		schemaLedgerUUID := tc.ps1.TowardSchema()

		psh, err := FragmentAssignHeadingOriginatingSchema(&schemaLedgerUUID)
		if tc.expirationPhrase {
			require.Equal(t, tc.ps1, psh, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func VerifyFragmentSchemaArea(t *testing.T) {
	attestation := hashmap.Attestation{
		Sum:    1,
		Ordinal:    1,
		NodeDigest: commitrand.Octets(32),
	}
	verifyScenarios := []struct {
		msg     string
		ps1     *Fragment
		expirationPhrase bool
	}{
		{"REDACTED", &Fragment{}, false},
		{"REDACTED", nil, false},
		{
			"REDACTED",
			&Fragment{Ordinal: 1, Octets: commitrand.Octets(32), Attestation: attestation}, true,
		},
	}

	for _, tc := range verifyScenarios {
		schema, err := tc.ps1.TowardSchema()
		if tc.expirationPhrase {
			require.NoError(t, err, tc.msg)
		}

		p, err := FragmentOriginatingSchema(schema)
		if tc.expirationPhrase {
			require.NoError(t, err)
			require.Equal(t, tc.ps1, p, tc.msg)
		}
	}
}

func AssessmentCreateFragmentAssign(b *testing.B) {
	for nthFragments := 1; nthFragments <= 5; nthFragments++ {
		b.Run(fmt.Sprintf("REDACTED", nthFragments), func(b *testing.B) {
			data := commitrand.Octets(verifyFragmentExtent * nthFragments)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				FreshFragmentAssignOriginatingData(data, verifyFragmentExtent)
			}
		})
	}
}
