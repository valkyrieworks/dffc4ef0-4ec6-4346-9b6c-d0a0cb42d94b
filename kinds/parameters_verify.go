package kinds

import (
	"bytes"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

var (
	itemEdwards25519   = []string{IfacePublicTokenKindEdwards25519}
	itemEllipticp256 = []string{IfacePublicTokenKindEllipticp256}
)

func VerifyAgreementParametersCertification(t *testing.T) {
	verifyScenarios := []struct {
		parameters AgreementSettings
		sound  bool
	}{
		//
		0: {createParameters(1, 0, 2, 0, itemEdwards25519, 0), true},
		1: {createParameters(0, 0, 2, 0, itemEdwards25519, 0), false},
		2: {createParameters(47*1024*1024, 0, 2, 0, itemEdwards25519, 0), true},
		3: {createParameters(10, 0, 2, 0, itemEdwards25519, 0), true},
		4: {createParameters(100*1024*1024, 0, 2, 0, itemEdwards25519, 0), true},
		5: {createParameters(101*1024*1024, 0, 2, 0, itemEdwards25519, 0), false},
		6: {createParameters(1024*1024*1024, 0, 2, 0, itemEdwards25519, 0), false},
		//
		7:  {createParameters(1, 0, 0, 0, itemEdwards25519, 0), false},
		8:  {createParameters(1, 0, 2, 2, itemEdwards25519, 0), false},
		9:  {createParameters(1000, 0, 2, 1, itemEdwards25519, 0), true},
		10: {createParameters(1, 0, -1, 0, itemEdwards25519, 0), false},
		//
		11: {createParameters(1, 0, 2, 0, []string{}, 0), false},
		//
		12: {createParameters(1, 0, 2, 0, []string{"REDACTED"}, 0), false},
		13: {createParameters(-1, 0, 2, 0, itemEdwards25519, 0), true},
		14: {createParameters(-2, 0, 2, 0, itemEdwards25519, 0), false},
	}
	for i, tc := range verifyScenarios {
		if tc.sound {
			assert.NoErrorf(t, tc.parameters.CertifyFundamental(), "REDACTED", i)
		} else {
			assert.Errorf(t, tc.parameters.CertifyFundamental(), "REDACTED", i)
		}
	}
}

func createParameters(
	ledgerOctets, ledgerFuel int64,
	proofLifespan int64,
	maximumProofOctets int64,
	publickeyKinds []string,
	ifaceAdditionAltitude int64,
) AgreementSettings {
	return AgreementSettings{
		Ledger: LedgerParameters{
			MaximumOctets: ledgerOctets,
			MaximumFuel:   ledgerFuel,
		},
		Proof: ProofParameters{
			MaximumLifespanCountLedgers: proofLifespan,
			MaximumLifespanInterval:  time.Duration(proofLifespan),
			MaximumOctets:        maximumProofOctets,
		},
		Assessor: AssessorParameters{
			PublicTokenKinds: publickeyKinds,
		},
		Iface: IfaceParameters{
			BallotAdditionsActivateAltitude: ifaceAdditionAltitude,
		},
	}
}

func VerifyAgreementParametersDigest(t *testing.T) {
	parameters := []AgreementSettings{
		createParameters(4, 2, 3, 1, itemEdwards25519, 0),
		createParameters(1, 4, 3, 1, itemEdwards25519, 0),
		createParameters(1, 2, 4, 1, itemEdwards25519, 0),
		createParameters(2, 5, 7, 1, itemEdwards25519, 0),
		createParameters(1, 7, 6, 1, itemEdwards25519, 0),
		createParameters(9, 5, 4, 1, itemEdwards25519, 0),
		createParameters(7, 8, 9, 1, itemEdwards25519, 0),
		createParameters(4, 6, 5, 1, itemEdwards25519, 0),
	}

	digests := make([][]byte, len(parameters))
	for i := range parameters {
		digests[i] = parameters[i].Digest()
	}

	//
	//
	sort.Slice(digests, func(i, j int) bool {
		return bytes.Compare(digests[i], digests[j]) < 0
	})
	for i := 0; i < len(digests)-1; i++ {
		assert.NotEqual(t, digests[i], digests[i+1])
	}
}

func VerifyAgreementParametersRevise(t *testing.T) {
	verifyScenarios := []struct {
		parameters        AgreementSettings
		revisions       *commitchema.AgreementSettings
		modifiedParameters AgreementSettings
	}{
		//
		{
			createParameters(1, 2, 3, 0, itemEdwards25519, 0),
			&commitchema.AgreementSettings{},
			createParameters(1, 2, 3, 0, itemEdwards25519, 0),
		},
		//
		{
			createParameters(1, 2, 3, 0, itemEdwards25519, 0),
			&commitchema.AgreementSettings{
				Ledger: &commitchema.LedgerParameters{
					MaximumOctets: 100,
					MaximumFuel:   200,
				},
				Proof: &commitchema.ProofParameters{
					MaximumLifespanCountLedgers: 300,
					MaximumLifespanInterval:  time.Duration(300),
					MaximumOctets:        50,
				},
				Assessor: &commitchema.AssessorParameters{
					PublicTokenKinds: itemEllipticp256,
				},
			},
			createParameters(100, 200, 300, 50, itemEllipticp256, 0),
		},
	}

	for _, tc := range verifyScenarios {
		assert.Equal(t, tc.modifiedParameters, tc.parameters.Revise(tc.revisions))
	}
}

func Testagreementparamsrevision_Applicationversion(t *testing.T) {
	parameters := createParameters(1, 2, 3, 0, itemEdwards25519, 0)

	assert.EqualValues(t, 0, parameters.Edition.App)

	modified := parameters.Revise(
		&commitchema.AgreementSettings{Edition: &commitchema.EditionParameters{App: 1}})

	assert.EqualValues(t, 1, modified.Edition.App)
}

func Testagreementparamsrevision_Ballotadditionsenableheight(t *testing.T) {
	const voidVerify = -10000000
	verifyScenarios := []struct {
		alias        string
		prevailing     int64
		originating        int64
		to          int64
		anticipatedFault bool
	}{
		//
		{"REDACTED", 3, 0, 0, false},
		{"REDACTED", 3, 100, 100, false},
		{"REDACTED", 100, 100, 100, false},
		{"REDACTED", 300, 100, 100, false},
		//
		{"REDACTED", 3, 0, 5, false},
		{"REDACTED", 4, 0, 5, false},
		{"REDACTED", 5, 0, 5, true},
		{"REDACTED", 6, 0, 5, true},
		{"REDACTED", 50, 0, 5, true},
		//
		{"REDACTED", 4, 5, 0, false},
		{"REDACTED", 5, 5, 0, true},
		{"REDACTED", 6, 5, 0, true},
		{"REDACTED", 10, 5, 0, true},
		//
		{"REDACTED", 1, 10, 5, false},
		{"REDACTED", 4, 10, 5, false},
		{"REDACTED", 5, 10, 5, true},
		{"REDACTED", 6, 10, 5, true},
		{"REDACTED", 9, 10, 5, true},
		{"REDACTED", 10, 10, 5, true},
		{"REDACTED", 11, 10, 5, true},
		{"REDACTED", 100, 10, 5, true},
		//
		{"REDACTED", 3, 10, 15, false},
		{"REDACTED", 9, 10, 15, false},
		{"REDACTED", 10, 10, 15, true},
		{"REDACTED", 11, 10, 15, true},
		{"REDACTED", 14, 10, 15, true},
		{"REDACTED", 15, 10, 15, true},
		{"REDACTED", 16, 10, 15, true},
		{"REDACTED", 100, 10, 15, true},
		//
		{"REDACTED", 3, 0, -5, true},
		{"REDACTED", 3, -5, 100, false},
		{"REDACTED", 3, -10, 3, true},
		{"REDACTED", 3, -3, -3, true},
		{"REDACTED", 100, -8, -9, true},
		{"REDACTED", 300, -10, -8, true},
		//
		{"REDACTED", 300, 400, voidVerify, false},
		{"REDACTED", 300, 200, voidVerify, false},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(*testing.T) {
			primaryParameters := createParameters(1, 0, 2, 0, itemEdwards25519, tc.originating)
			revise := &commitchema.AgreementSettings{}
			if tc.to == voidVerify {
				revise.Iface = nil
			} else {
				revise.Iface = &commitchema.IfaceParameters{
					BallotAdditionsActivateAltitude: tc.to,
				}
			}
			if tc.anticipatedFault {
				require.Error(t, primaryParameters.CertifyRevise(revise, tc.prevailing))
			} else {
				require.NoError(t, primaryParameters.CertifyRevise(revise, tc.prevailing))
			}
		})
	}
}

func VerifySchema(t *testing.T) {
	parameters := []AgreementSettings{
		createParameters(4, 2, 3, 1, itemEdwards25519, 1),
		createParameters(1, 4, 3, 1, itemEdwards25519, 1),
		createParameters(1, 2, 4, 1, itemEdwards25519, 1),
		createParameters(2, 5, 7, 1, itemEdwards25519, 1),
		createParameters(1, 7, 6, 1, itemEdwards25519, 1),
		createParameters(9, 5, 4, 1, itemEdwards25519, 1),
		createParameters(7, 8, 9, 1, itemEdwards25519, 1),
		createParameters(4, 6, 5, 1, itemEdwards25519, 1),
	}

	for i := range parameters {
		bufferParameters := parameters[i].TowardSchema()

		sourceParameters := AgreementParametersOriginatingSchema(bufferParameters)

		assert.Equal(t, parameters[i], sourceParameters)

	}
}
