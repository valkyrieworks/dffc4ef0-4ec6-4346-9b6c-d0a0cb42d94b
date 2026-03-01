package kinds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyAssessorSchemaArea(t *testing.T) {
	val, _ := ArbitraryAssessor(true, 100)
	verifyScenarios := []struct {
		msg      string
		v1       *Assessor
		expirationPhase1 bool
		expirationPhase2 bool
	}{
		{"REDACTED", val, true, true},
		{"REDACTED", &Assessor{}, false, false},
		{"REDACTED", nil, false, false},
	}
	for _, tc := range verifyScenarios {
		schemaItem, err := tc.v1.TowardSchema()

		if tc.expirationPhase1 {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		val, err := AssessorOriginatingSchema(schemaItem)
		if tc.expirationPhase2 {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.v1, val, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func VerifyAssessorCertifyFundamental(t *testing.T) {
	private := FreshSimulatePRV()
	publicToken, _ := private.ObtainPublicToken()
	verifyScenarios := []struct {
		val *Assessor
		err bool
		msg string
	}{
		{
			val: FreshAssessor(publicToken, 1),
			err: false,
			msg: "REDACTED",
		},
		{
			val: nil,
			err: true,
			msg: "REDACTED",
		},
		{
			val: &Assessor{
				PublicToken: nil,
			},
			err: true,
			msg: "REDACTED",
		},
		{
			val: FreshAssessor(publicToken, -1),
			err: true,
			msg: "REDACTED",
		},
		{
			val: &Assessor{
				PublicToken:  publicToken,
				Location: nil,
			},
			err: true,
			msg: fmt.Sprintf("REDACTED", publicToken.Location()),
		},
		{
			val: &Assessor{
				PublicToken:  publicToken,
				Location: []byte{'a'},
			},
			err: true,
			msg: fmt.Sprintf("REDACTED", publicToken.Location()),
		},
	}

	for _, tc := range verifyScenarios {
		err := tc.val.CertifyFundamental()
		if tc.err {
			if assert.Error(t, err) {
				assert.Equal(t, tc.msg, err.Error())
			}
		} else {
			assert.NoError(t, err)
		}
	}
}
