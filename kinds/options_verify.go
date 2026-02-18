package kinds

import (
	"bytes"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

var (
	valueEd25519   = []string{IfacePublicKeyKindEd25519}
	valueSecp256k1 = []string{IfacePublicKeyKindSecp256k1}
)

func VerifyAgreementOptionsVerification(t *testing.T) {
	verifyScenarios := []struct {
		options AgreementOptions
		sound  bool
	}{
		//
		0: {createOptions(1, 0, 2, 0, valueEd25519, 0), true},
		1: {createOptions(0, 0, 2, 0, valueEd25519, 0), false},
		2: {createOptions(47*1024*1024, 0, 2, 0, valueEd25519, 0), true},
		3: {createOptions(10, 0, 2, 0, valueEd25519, 0), true},
		4: {createOptions(100*1024*1024, 0, 2, 0, valueEd25519, 0), true},
		5: {createOptions(101*1024*1024, 0, 2, 0, valueEd25519, 0), false},
		6: {createOptions(1024*1024*1024, 0, 2, 0, valueEd25519, 0), false},
		//
		7:  {createOptions(1, 0, 0, 0, valueEd25519, 0), false},
		8:  {createOptions(1, 0, 2, 2, valueEd25519, 0), false},
		9:  {createOptions(1000, 0, 2, 1, valueEd25519, 0), true},
		10: {createOptions(1, 0, -1, 0, valueEd25519, 0), false},
		//
		11: {createOptions(1, 0, 2, 0, []string{}, 0), false},
		//
		12: {createOptions(1, 0, 2, 0, []string{"REDACTED"}, 0), false},
		13: {createOptions(-1, 0, 2, 0, valueEd25519, 0), true},
		14: {createOptions(-2, 0, 2, 0, valueEd25519, 0), false},
	}
	for i, tc := range verifyScenarios {
		if tc.sound {
			assert.NoErrorf(t, tc.options.CertifySimple(), "REDACTED", i)
		} else {
			assert.Errorf(t, tc.options.CertifySimple(), "REDACTED", i)
		}
	}
}

func createOptions(
	ledgerOctets, ledgerFuel int64,
	proofPeriod int64,
	maximumProofOctets int64,
	publickeyKinds []string,
	ifaceAdditionLevel int64,
) AgreementOptions {
	return AgreementOptions{
		Ledger: LedgerOptions{
			MaximumOctets: ledgerOctets,
			MaximumFuel:   ledgerFuel,
		},
		Proof: ProofOptions{
			MaximumDurationCountLedgers: proofPeriod,
			MaximumDurationPeriod:  time.Duration(proofPeriod),
			MaximumOctets:        maximumProofOctets,
		},
		Ratifier: RatifierOptions{
			PublicKeyKinds: publickeyKinds,
		},
		Iface: IfaceOptions{
			BallotPluginsActivateLevel: ifaceAdditionLevel,
		},
	}
}

func VerifyAgreementOptionsDigest(t *testing.T) {
	options := []AgreementOptions{
		createOptions(4, 2, 3, 1, valueEd25519, 0),
		createOptions(1, 4, 3, 1, valueEd25519, 0),
		createOptions(1, 2, 4, 1, valueEd25519, 0),
		createOptions(2, 5, 7, 1, valueEd25519, 0),
		createOptions(1, 7, 6, 1, valueEd25519, 0),
		createOptions(9, 5, 4, 1, valueEd25519, 0),
		createOptions(7, 8, 9, 1, valueEd25519, 0),
		createOptions(4, 6, 5, 1, valueEd25519, 0),
	}

	digests := make([][]byte, len(options))
	for i := range options {
		digests[i] = options[i].Digest()
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

func VerifyAgreementOptionsModify(t *testing.T) {
	verifyScenarios := []struct {
		options        AgreementOptions
		refreshes       *engineproto.AgreementOptions
		refreshedOptions AgreementOptions
	}{
		//
		{
			createOptions(1, 2, 3, 0, valueEd25519, 0),
			&engineproto.AgreementOptions{},
			createOptions(1, 2, 3, 0, valueEd25519, 0),
		},
		//
		{
			createOptions(1, 2, 3, 0, valueEd25519, 0),
			&engineproto.AgreementOptions{
				Ledger: &engineproto.LedgerOptions{
					MaximumOctets: 100,
					MaximumFuel:   200,
				},
				Proof: &engineproto.ProofOptions{
					MaximumDurationCountLedgers: 300,
					MaximumDurationPeriod:  time.Duration(300),
					MaximumOctets:        50,
				},
				Ratifier: &engineproto.RatifierOptions{
					PublicKeyKinds: valueSecp256k1,
				},
			},
			createOptions(100, 200, 300, 50, valueSecp256k1, 0),
		},
	}

	for _, tc := range verifyScenarios {
		assert.Equal(t, tc.refreshedOptions, tc.options.Modify(tc.refreshes))
	}
}

func Verifyconsensusoptionsupdate_Applicationversion(t *testing.T) {
	options := createOptions(1, 2, 3, 0, valueEd25519, 0)

	assert.EqualValues(t, 0, options.Release.App)

	refreshed := options.Modify(
		&engineproto.AgreementOptions{Release: &engineproto.ReleaseOptions{App: 1}})

	assert.EqualValues(t, 1, refreshed.Release.App)
}

func Verifyconsensusoptionsupdate_Ballotpluginsenablelevel(t *testing.T) {
	const nullVerify = -10000000
	verifyScenarios := []struct {
		label        string
		ongoing     int64
		from        int64
		to          int64
		anticipatedErr bool
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
		{"REDACTED", 300, 400, nullVerify, false},
		{"REDACTED", 300, 200, nullVerify, false},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(*testing.T) {
			primaryOptions := createOptions(1, 0, 2, 0, valueEd25519, tc.from)
			modify := &engineproto.AgreementOptions{}
			if tc.to == nullVerify {
				modify.Iface = nil
			} else {
				modify.Iface = &engineproto.IfaceOptions{
					BallotPluginsActivateLevel: tc.to,
				}
			}
			if tc.anticipatedErr {
				require.Error(t, primaryOptions.CertifyModify(modify, tc.ongoing))
			} else {
				require.NoError(t, primaryOptions.CertifyModify(modify, tc.ongoing))
			}
		})
	}
}

func VerifySchema(t *testing.T) {
	options := []AgreementOptions{
		createOptions(4, 2, 3, 1, valueEd25519, 1),
		createOptions(1, 4, 3, 1, valueEd25519, 1),
		createOptions(1, 2, 4, 1, valueEd25519, 1),
		createOptions(2, 5, 7, 1, valueEd25519, 1),
		createOptions(1, 7, 6, 1, valueEd25519, 1),
		createOptions(9, 5, 4, 1, valueEd25519, 1),
		createOptions(7, 8, 9, 1, valueEd25519, 1),
		createOptions(4, 6, 5, 1, valueEd25519, 1),
	}

	for i := range options {
		pbOptions := options[i].ToSchema()

		originOptions := AgreementOptionsFromSchema(pbOptions)

		assert.Equal(t, options[i], originOptions)

	}
}
