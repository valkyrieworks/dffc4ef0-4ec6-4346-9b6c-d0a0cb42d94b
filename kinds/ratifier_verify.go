package kinds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyRatifierSchemaBuffer(t *testing.T) {
	val, _ := RandomRatifier(true, 100)
	verifyScenarios := []struct {
		msg      string
		v1       *Ratifier
		expirationPass1 bool
		expirationPass2 bool
	}{
		{"REDACTED", val, true, true},
		{"REDACTED", &Ratifier{}, false, false},
		{"REDACTED", nil, false, false},
	}
	for _, tc := range verifyScenarios {
		schemaValue, err := tc.v1.ToSchema()

		if tc.expirationPass1 {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		val, err := RatifierFromSchema(schemaValue)
		if tc.expirationPass2 {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.v1, val, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func VerifyRatifierCertifySimple(t *testing.T) {
	private := NewEmulatePV()
	publicKey, _ := private.FetchPublicKey()
	verifyScenarios := []struct {
		val *Ratifier
		err bool
		msg string
	}{
		{
			val: NewRatifier(publicKey, 1),
			err: false,
			msg: "REDACTED",
		},
		{
			val: nil,
			err: true,
			msg: "REDACTED",
		},
		{
			val: &Ratifier{
				PublicKey: nil,
			},
			err: true,
			msg: "REDACTED",
		},
		{
			val: NewRatifier(publicKey, -1),
			err: true,
			msg: "REDACTED",
		},
		{
			val: &Ratifier{
				PublicKey:  publicKey,
				Location: nil,
			},
			err: true,
			msg: fmt.Sprintf("REDACTED", publicKey.Location()),
		},
		{
			val: &Ratifier{
				PublicKey:  publicKey,
				Location: []byte{'a'},
			},
			err: true,
			msg: fmt.Sprintf("REDACTED", publicKey.Location()),
		},
	}

	for _, tc := range verifyScenarios {
		err := tc.val.CertifySimple()
		if tc.err {
			if assert.Error(t, err) {
				assert.Equal(t, tc.msg, err.Error())
			}
		} else {
			assert.NoError(t, err)
		}
	}
}
