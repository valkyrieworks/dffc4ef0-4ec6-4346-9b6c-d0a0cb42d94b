package kinds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/ed25519"
	cometjson "github.com/valkyrieworks/utils/json"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

func VerifyOriginFlawed(t *testing.T) {
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
		_, err := OriginPaperFromJSON(verifyInstance)
		assert.Error(t, err, "REDACTED")
	}
}

func VerifyOriginSound(t *testing.T) {
	//
	generatePaperOctets := []byte(
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
	_, err := OriginPaperFromJSON(generatePaperOctets)
	assert.NoError(t, err, "REDACTED")

	publickey := ed25519.GeneratePrivateKey().PublicKey()
	//
	rootGeneratePaper := &OriginPaper{
		LedgerUID:    "REDACTED",
		Ratifiers: []OriginRatifier{{publickey.Location(), publickey, 10, "REDACTED"}},
	}
	generatePaperOctets, err = cometjson.Serialize(rootGeneratePaper)
	assert.NoError(t, err, "REDACTED")

	//
	generatePaper, err := OriginPaperFromJSON(generatePaperOctets)
	assert.NoError(t, err, "REDACTED")
	assert.NotNil(t, generatePaper.AgreementOptions, "REDACTED")

	//
	assert.NotNil(t, generatePaper.Ratifiers[0].Location, "REDACTED")

	//
	generatePaperOctets, err = cometjson.Serialize(generatePaper)
	assert.NoError(t, err, "REDACTED")
	generatePaper, err = OriginPaperFromJSON(generatePaperOctets)
	assert.NoError(t, err, "REDACTED")

	//
	generatePaper.AgreementOptions.Ledger.MaximumOctets = 0
	generatePaperOctets, err = cometjson.Serialize(generatePaper)
	assert.NoError(t, err, "REDACTED")
	_, err = OriginPaperFromJSON(generatePaperOctets)
	assert.Error(t, err, "REDACTED")

	//
	absentRatifiersVerifyScenarios := [][]byte{
		[]byte("REDACTED"),                   //
		[]byte("REDACTED"),   //
		[]byte("REDACTED"), //
		[]byte("REDACTED"),                   //
	}

	for _, tc := range absentRatifiersVerifyScenarios {
		_, err := OriginPaperFromJSON(tc)
		assert.NoError(t, err)
	}
}

func VerifyOriginPersistAs(t *testing.T) {
	tempfile, err := os.CreateTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.Remove(tempfile.Name())

	generatePaper := arbitraryOriginPaper()

	//
	err = generatePaper.PersistAs(tempfile.Name())
	require.NoError(t, err)
	status, err := tempfile.Stat()
	require.NoError(t, err)
	if err != nil && status.Size() <= 0 {
		t.Fatalf("REDACTED", tempfile.Name())
	}

	err = tempfile.Close()
	require.NoError(t, err)

	//
	generateDoc2, err := OriginPaperFromEntry(tempfile.Name())
	require.NoError(t, err)
	assert.EqualValues(t, generateDoc2, generatePaper)
	assert.Equal(t, generateDoc2.Ratifiers, generatePaper.Ratifiers)
}

func VerifyOriginRatifierDigest(t *testing.T) {
	generatePaper := arbitraryOriginPaper()
	assert.NotEmpty(t, generatePaper.RatifierDigest())
}

func arbitraryOriginPaper() *OriginPaper {
	publickey := ed25519.GeneratePrivateKey().PublicKey()
	return &OriginPaper{
		OriginMoment:     engineclock.Now(),
		LedgerUID:         "REDACTED",
		PrimaryLevel:   1000,
		Ratifiers:      []OriginRatifier{{publickey.Location(), publickey, 10, "REDACTED"}},
		AgreementOptions: StandardAgreementOptions(),
		ApplicationDigest:         []byte{1, 2, 3},
	}
}
