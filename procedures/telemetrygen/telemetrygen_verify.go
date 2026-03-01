package primary__test

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	telemetrygen "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/procedures/telemetrygen"
)

const verifyDataPath = "REDACTED"

func VerifyPlainBlueprint(t *testing.T) {
	m := telemetrygen.ProcessedMeasurementAttribute{
		KindAlias:    "REDACTED",
		AttributeAlias:   "REDACTED",
		MeasurementAlias:  "REDACTED",
		Characterization: "REDACTED",
		Tags:      "REDACTED",
	}
	td := telemetrygen.BlueprintData{
		Bundle:       "REDACTED",
		ProcessedTelemetry: []telemetrygen.ProcessedMeasurementAttribute{m},
	}
	b := bytes.NewBuffer([]byte{})
	err := telemetrygen.ComposeTelemetryRecord(b, td)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
}

func VerifyOriginatingData(t *testing.T) {
	insights, err := os.ReadDir(verifyDataPath)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	for _, dir := range insights {
		t.Run(dir.Name(), func(t *testing.T) {
			if !dir.IsDir() {
				t.Fatalf("REDACTED", dir.Name())
			}
			pathAlias := path.Join(verifyDataPath, dir.Name())
			pt, err := telemetrygen.AnalyzeTelemetryPath(pathAlias, "REDACTED")
			if err != nil {
				t.Fatalf("REDACTED", dir, err)
			}
			outputRecord := path.Join(pathAlias, "REDACTED")
			if err != nil {
				t.Fatalf("REDACTED", outputRecord, err)
			}
			of, err := os.Create(outputRecord)
			if err != nil {
				t.Fatalf("REDACTED", outputRecord, err)
			}
			defer os.Remove(outputRecord)
			if err := telemetrygen.ComposeTelemetryRecord(of, pt); err != nil {
				t.Fatalf("REDACTED", outputRecord, err)
			}
			if _, err := parser.ParseFile(token.NewFileSet(), outputRecord, nil, parser.AllErrors); err != nil {
				t.Fatalf("REDACTED", outputRecord, err)
			}
			byteFresh, err := os.ReadFile(outputRecord)
			if err != nil {
				t.Fatalf("REDACTED", outputRecord, err)
			}
			primeRecord := path.Join(pathAlias, "REDACTED")
			byteAged, err := os.ReadFile(primeRecord)
			if err != nil {
				t.Fatalf("REDACTED", primeRecord, err)
			}
			if !bytes.Equal(byteFresh, byteAged) {
				t.Fatalf("REDACTED"+
					"REDACTED"+
					"REDACTED", outputRecord, primeRecord)
			}
		})
	}
}

func VerifyAnalyzeTelemetryRecord(t *testing.T) {
	const libraryAlias = "REDACTED"
	telemetryVerifies := []struct {
		alias          string
		mustFailure   bool
		telemetryRecord string
		anticipated      telemetrygen.BlueprintData
	}{
		{
			alias: "REDACTED",
			telemetryRecord: `REDACTED{
REDACTEDe
REDACTED`,
			anticipated: telemetrygen.BlueprintData{
				Bundle: libraryAlias,
				ProcessedTelemetry: []telemetrygen.ProcessedMeasurementAttribute{
					{
						KindAlias:   "REDACTED",
						AttributeAlias:  "REDACTED",
						MeasurementAlias: "REDACTED",
					},
				},
			},
		},
		{
			alias: "REDACTED",
			telemetryRecord: "REDACTED" +
				"REDACTED" +
				"REDACTED",
			anticipated: telemetrygen.BlueprintData{
				Bundle: libraryAlias,
				ProcessedTelemetry: []telemetrygen.ProcessedMeasurementAttribute{
					{
						KindAlias:   "REDACTED",
						AttributeAlias:  "REDACTED",
						MeasurementAlias: "REDACTED",

						HistogramChoices: telemetrygen.HistogramOptions{
							SegmentKind:  "REDACTED",
							SegmentExtents: "REDACTED",
						},
					},
				},
			},
		},
		{
			alias: "REDACTED",
			telemetryRecord: "REDACTED" +
				"REDACTED" +
				"REDACTED",
			anticipated: telemetrygen.BlueprintData{
				Bundle: libraryAlias,
				ProcessedTelemetry: []telemetrygen.ProcessedMeasurementAttribute{
					{
						KindAlias:   "REDACTED",
						AttributeAlias:  "REDACTED",
						MeasurementAlias: "REDACTED",
					},
				},
			},
		},
		{
			alias: "REDACTED",
			telemetryRecord: "REDACTED" +
				"REDACTED" +
				"REDACTED",
			anticipated: telemetrygen.BlueprintData{
				Bundle: libraryAlias,
				ProcessedTelemetry: []telemetrygen.ProcessedMeasurementAttribute{
					{
						KindAlias:   "REDACTED",
						AttributeAlias:  "REDACTED",
						MeasurementAlias: "REDACTED",
						Tags:     "REDACTED",
					},
				},
			},
		},
		{
			alias: "REDACTED",
			telemetryRecord: `REDACTED{
REDACTEDr
REDACTEDg
REDACTED`,
			anticipated: telemetrygen.BlueprintData{
				Bundle: libraryAlias,
				ProcessedTelemetry: []telemetrygen.ProcessedMeasurementAttribute{
					{
						KindAlias:   "REDACTED",
						AttributeAlias:  "REDACTED",
						MeasurementAlias: "REDACTED",
					},
				},
			},
		},
	}
	for _, verifyInstance := range telemetryVerifies {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			dir, err := os.MkdirTemp(os.TempDir(), "REDACTED")
			if err != nil {
				t.Fatalf("REDACTED", err)
			}
			defer os.Remove(dir)
			f, err := os.Create(filepath.Join(dir, "REDACTED"))
			if err != nil {
				t.Fatalf("REDACTED", err)
			}
			libraryRow := fmt.Sprintf("REDACTED", libraryAlias)
			ingestCondition := `
REDACTED(
REDACTED"
REDACTED)
REDACTED`

			_, err = io.WriteString(f, libraryRow)
			require.NoError(t, err)
			_, err = io.WriteString(f, ingestCondition)
			require.NoError(t, err)
			_, err = io.WriteString(f, verifyInstance.telemetryRecord)
			require.NoError(t, err)

			td, err := telemetrygen.AnalyzeTelemetryPath(dir, "REDACTED")
			if verifyInstance.mustFailure {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, verifyInstance.anticipated, td)
			}
		})
	}
}

func VerifyAnalyzeRenamedMeasurement(t *testing.T) {
	renamedData := `
REDACTEDg

REDACTED(
REDACTED"
REDACTED)
REDACTED{
REDACTEDe
REDACTED}
REDACTED`
	dir, err := os.MkdirTemp(os.TempDir(), "REDACTED")
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	defer os.Remove(dir)
	f, err := os.Create(filepath.Join(dir, "REDACTED"))
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	_, err = io.WriteString(f, renamedData)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	td, err := telemetrygen.AnalyzeTelemetryPath(dir, "REDACTED")
	require.NoError(t, err)

	anticipated := telemetrygen.BlueprintData{
		Bundle: "REDACTED",
		ProcessedTelemetry: []telemetrygen.ProcessedMeasurementAttribute{
			{
				KindAlias:   "REDACTED",
				AttributeAlias:  "REDACTED",
				MeasurementAlias: "REDACTED",
			},
		},
	}
	require.Equal(t, anticipated, td)
}
