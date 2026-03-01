package primary__test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	telemetryvariance "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/procedures/telemetrygen/telemetryvariance"
)

func VerifyVariance(t *testing.T) {
	for _, tc := range []struct {
		alias      string
		anMaterial string
		byteMaterial string

		desire string
	}{
		{
			alias: "REDACTED",
			anMaterial: `
REDACTED0
REDACTED`,
			byteMaterial: `
REDACTED0
REDACTED`,
			desire: `REDACTED:
REDACTEDe
REDACTEDe
REDACTEDr
REDACTEDe
REDACTEDo
REDACTED`,
		},
		{
			alias: "REDACTED",
			anMaterial: `
REDACTED0
REDACTED`,
			byteMaterial: `
REDACTED0
REDACTED`,
			desire: `REDACTED:
REDACTEDo
REDACTEDe
REDACTED`,
		},
	} {
		t.Run(tc.alias, func(t *testing.T) {
			areaAN := bytes.NewBuffer([]byte{})
			areaBYTE := bytes.NewBuffer([]byte{})
			_, err := io.WriteString(areaAN, tc.anMaterial)
			require.NoError(t, err)
			_, err = io.WriteString(areaBYTE, tc.byteMaterial)
			require.NoError(t, err)
			md, err := telemetryvariance.VarianceOriginatingFetchers(areaAN, areaBYTE)
			require.NoError(t, err)
			require.Equal(t, tc.desire, md.Text())
		})
	}
}
