package kinds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

func VerifyInaugurationFlawed(t *testing.T) {
	//
	verifyScenarios := [][]byte{
		{},              //
		{1, 1, 1, 1, 1}, //
		[]byte("REDACTED"),    //
		[]byte("REDACTED"),   //
		[]byte("REDACTED"), //
		//
		[]byte(
			"REDACTED",
		),
		//
		[]byte(
			"REDACTED" +
				"REDACTED" +
				"REDACTED" +
				"REDACTED" +
				"REDACTED",
		),
		//
		[]byte(
			"REDACTED" +
				"REDACTED" +
				"REDACTED" +
				"REDACTED" +
				"REDACTED",
		),
		//
		[]byte(
			"REDACTED" +
				"REDACTED" +
				"REDACTED" +
				"REDACTED" +
				"REDACTED",
		),
	}

	for _, verifyInstance := range verifyScenarios {
		_, err := InaugurationPaperOriginatingJSN(verifyInstance)
		assert.Error(t, err, "REDACTED")
	}
}

func VerifyInaugurationValid(t *testing.T) {
	//
	producePaperOctets := []byte(
		`REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED{
REDACTED,
REDACTED,
REDACTED"
REDACTED,
REDACTED,
REDACTED}
REDACTED`,
	)
	_, err := InaugurationPaperOriginatingJSN(producePaperOctets)
	assert.NoError(t, err, "REDACTED")

	publickey := edwards25519.ProducePrivateToken().PublicToken()
	//
	foundationProducePaper := &OriginPaper{
		SuccessionUUID:    "REDACTED",
		Assessors: []OriginAssessor{{publickey.Location(), publickey, 10, "REDACTED"}},
	}
	producePaperOctets, err = strongmindjson.Serialize(foundationProducePaper)
	assert.NoError(t, err, "REDACTED")

	//
	producePaper, err := InaugurationPaperOriginatingJSN(producePaperOctets)
	assert.NoError(t, err, "REDACTED")
	assert.NotNil(t, producePaper.AgreementSettings, "REDACTED")

	//
	assert.NotNil(t, producePaper.Assessors[0].Location, "REDACTED")

	//
	producePaperOctets, err = strongmindjson.Serialize(producePaper)
	assert.NoError(t, err, "REDACTED")
	producePaper, err = InaugurationPaperOriginatingJSN(producePaperOctets)
	assert.NoError(t, err, "REDACTED")

	//
	producePaper.AgreementSettings.Ledger.MaximumOctets = 0
	producePaperOctets, err = strongmindjson.Serialize(producePaper)
	assert.NoError(t, err, "REDACTED")
	_, err = InaugurationPaperOriginatingJSN(producePaperOctets)
	assert.Error(t, err, "REDACTED")

	//
	absentAssessorsVerifyScenarios := [][]byte{
		[]byte("REDACTED"),                   //
		[]byte("REDACTED"),   //
		[]byte("REDACTED"), //
		[]byte("REDACTED"),                   //
	}

	for _, tc := range absentAssessorsVerifyScenarios {
		_, err := InaugurationPaperOriginatingJSN(tc)
		assert.NoError(t, err)
	}
}

func VerifyInaugurationPersistLike(t *testing.T) {
	tempfile, err := os.CreateTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.Remove(tempfile.Name())

	producePaper := unpredictableInaugurationPaper()

	//
	err = producePaper.PersistLike(tempfile.Name())
	require.NoError(t, err)
	summary, err := tempfile.Stat()
	require.NoError(t, err)
	if err != nil && summary.Size() <= 0 {
		t.Fatalf("REDACTED", tempfile.Name())
	}

	err = tempfile.Close()
	require.NoError(t, err)

	//
	producePaper2, err := InaugurationPaperOriginatingRecord(tempfile.Name())
	require.NoError(t, err)
	assert.EqualValues(t, producePaper2, producePaper)
	assert.Equal(t, producePaper2.Assessors, producePaper.Assessors)
}

func VerifyInaugurationAssessorDigest(t *testing.T) {
	producePaper := unpredictableInaugurationPaper()
	assert.NotEmpty(t, producePaper.AssessorDigest())
}

func unpredictableInaugurationPaper() *OriginPaper {
	publickey := edwards25519.ProducePrivateToken().PublicToken()
	return &OriginPaper{
		OriginMoment:     committime.Now(),
		SuccessionUUID:         "REDACTED",
		PrimaryAltitude:   1000,
		Assessors:      []OriginAssessor{{publickey.Location(), publickey, 10, "REDACTED"}},
		AgreementSettings: FallbackAgreementSettings(),
		PlatformDigest:         []byte{1, 2, 3},
	}
}
