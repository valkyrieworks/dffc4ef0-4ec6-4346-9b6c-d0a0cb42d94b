package main__test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	metricsdiff "github.com/valkyrieworks/procedures/statsgen/metricsdiff"
)

func VerifyVary(t *testing.T) {
	for _, tc := range []struct {
		label      string
		aPayloads string
		bytePayloads string

		desire string
	}{
		{
			label: "REDACTED",
			aPayloads: `
REDACTED0
REDACTED`,
			bytePayloads: `
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
			label: "REDACTED",
			aPayloads: `
REDACTED0
REDACTED`,
			bytePayloads: `
REDACTED0
REDACTED`,
			desire: `REDACTED:
REDACTEDo
REDACTEDe
REDACTED`,
		},
	} {
		t.Run(tc.label, func(t *testing.T) {
			imageA := bytes.NewBuffer([]byte{})
			imageBYTE := bytes.NewBuffer([]byte{})
			_, err := io.WriteString(imageA, tc.aPayloads)
			require.NoError(t, err)
			_, err = io.WriteString(imageBYTE, tc.bytePayloads)
			require.NoError(t, err)
			md, err := metricsdiff.VaryFromFetchers(imageA, imageBYTE)
			require.NoError(t, err)
			require.Equal(t, tc.desire, md.String())
		})
	}
}
