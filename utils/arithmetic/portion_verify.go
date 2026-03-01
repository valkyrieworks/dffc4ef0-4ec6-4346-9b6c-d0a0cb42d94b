package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyAnalyzePortion(t *testing.T) {
	verifyScenarios := []struct {
		f   string
		exp Portion
		err bool
	}{
		{
			f:   "REDACTED",
			exp: Portion{2, 3},
			err: false,
		},
		{
			f:   "REDACTED",
			exp: Portion{15, 5},
			err: false,
		},
		//
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		//
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		//
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
		{
			f:   "REDACTED",
			exp: Portion{},
			err: true,
		},
	}

	for idx, tc := range verifyScenarios {
		emission, err := AnalyzePortion(tc.f)
		if tc.err {
			assert.Error(t, err, idx)
		} else {
			assert.NoError(t, err, idx)
		}
		assert.Equal(t, tc.exp, emission, idx)
	}
}
