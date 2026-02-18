package main__test

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

	statsgen "github.com/valkyrieworks/procedures/statsgen"
)

const verifyDataFolder = "REDACTED"

func VerifyBasicPrototype(t *testing.T) {
	m := statsgen.AnalyzedIndicatorField{
		KindLabel:    "REDACTED",
		FieldLabel:   "REDACTED",
		IndicatorLabel:  "REDACTED",
		Summary: "REDACTED",
		Tags:      "REDACTED",
	}
	td := statsgen.PrototypeData{
		Module:       "REDACTED",
		AnalyzedStats: []statsgen.AnalyzedIndicatorField{m},
	}
	b := bytes.NewBuffer([]byte{})
	err := statsgen.ComposeStatsEntry(b, td)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
}

func VerifyFromData(t *testing.T) {
	details, err := os.ReadDir(verifyDataFolder)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	for _, dir := range details {
		t.Run(dir.Name(), func(t *testing.T) {
			if !dir.IsDir() {
				t.Fatalf("REDACTED", dir.Name())
			}
			folderLabel := path.Join(verifyDataFolder, dir.Name())
			pt, err := statsgen.AnalyzeStatsFolder(folderLabel, "REDACTED")
			if err != nil {
				t.Fatalf("REDACTED", dir, err)
			}
			outEntry := path.Join(folderLabel, "REDACTED")
			if err != nil {
				t.Fatalf("REDACTED", outEntry, err)
			}
			of, err := os.Create(outEntry)
			if err != nil {
				t.Fatalf("REDACTED", outEntry, err)
			}
			defer os.Remove(outEntry)
			if err := statsgen.ComposeStatsEntry(of, pt); err != nil {
				t.Fatalf("REDACTED", outEntry, err)
			}
			if _, err := parser.ParseFile(token.NewFileSet(), outEntry, nil, parser.AllErrors); err != nil {
				t.Fatalf("REDACTED", outEntry, err)
			}
			byteNew, err := os.ReadFile(outEntry)
			if err != nil {
				t.Fatalf("REDACTED", outEntry, err)
			}
			validatedEntry := path.Join(folderLabel, "REDACTED")
			byteAged, err := os.ReadFile(validatedEntry)
			if err != nil {
				t.Fatalf("REDACTED", validatedEntry, err)
			}
			if !bytes.Equal(byteNew, byteAged) {
				t.Fatalf("REDACTED"+
					"REDACTED"+
					"REDACTED", outEntry, validatedEntry)
			}
		})
	}
}

func VerifyAnalyzeStatsStruct(t *testing.T) {
	const pkgLabel = "REDACTED"
	statsVerifies := []struct {
		label          string
		mustFault   bool
		statsStruct string
		anticipated      statsgen.PrototypeData
	}{
		{
			label: "REDACTED",
			statsStruct: `REDACTED{
REDACTEDe
REDACTED`,
			anticipated: statsgen.PrototypeData{
				Module: pkgLabel,
				AnalyzedStats: []statsgen.AnalyzedIndicatorField{
					{
						KindLabel:   "REDACTED",
						FieldLabel:  "REDACTED",
						IndicatorLabel: "REDACTED",
					},
				},
			},
		},
		{
			label: "REDACTED",
			statsStruct: "REDACTED" +
				"REDACTED" +
				"REDACTED",
			anticipated: statsgen.PrototypeData{
				Module: pkgLabel,
				AnalyzedStats: []statsgen.AnalyzedIndicatorField{
					{
						KindLabel:   "REDACTED",
						FieldLabel:  "REDACTED",
						IndicatorLabel: "REDACTED",

						HistogramSettings: statsgen.HistogramOpts{
							SegmentKind:  "REDACTED",
							ContainerExtents: "REDACTED",
						},
					},
				},
			},
		},
		{
			label: "REDACTED",
			statsStruct: "REDACTED" +
				"REDACTED" +
				"REDACTED",
			anticipated: statsgen.PrototypeData{
				Module: pkgLabel,
				AnalyzedStats: []statsgen.AnalyzedIndicatorField{
					{
						KindLabel:   "REDACTED",
						FieldLabel:  "REDACTED",
						IndicatorLabel: "REDACTED",
					},
				},
			},
		},
		{
			label: "REDACTED",
			statsStruct: "REDACTED" +
				"REDACTED" +
				"REDACTED",
			anticipated: statsgen.PrototypeData{
				Module: pkgLabel,
				AnalyzedStats: []statsgen.AnalyzedIndicatorField{
					{
						KindLabel:   "REDACTED",
						FieldLabel:  "REDACTED",
						IndicatorLabel: "REDACTED",
						Tags:     "REDACTED",
					},
				},
			},
		},
		{
			label: "REDACTED",
			statsStruct: `REDACTED{
REDACTEDr
REDACTEDg
REDACTED`,
			anticipated: statsgen.PrototypeData{
				Module: pkgLabel,
				AnalyzedStats: []statsgen.AnalyzedIndicatorField{
					{
						KindLabel:   "REDACTED",
						FieldLabel:  "REDACTED",
						IndicatorLabel: "REDACTED",
					},
				},
			},
		},
	}
	for _, verifyInstance := range statsVerifies {
		t.Run(verifyInstance.label, func(t *testing.T) {
			dir, err := os.MkdirTemp(os.TempDir(), "REDACTED")
			if err != nil {
				t.Fatalf("REDACTED", err)
			}
			defer os.Remove(dir)
			f, err := os.Create(filepath.Join(dir, "REDACTED"))
			if err != nil {
				t.Fatalf("REDACTED", err)
			}
			pkgRow := fmt.Sprintf("REDACTED", pkgLabel)
			includeStipulation := `
REDACTED(
REDACTED"
REDACTED)
REDACTED`

			_, err = io.WriteString(f, pkgRow)
			require.NoError(t, err)
			_, err = io.WriteString(f, includeStipulation)
			require.NoError(t, err)
			_, err = io.WriteString(f, verifyInstance.statsStruct)
			require.NoError(t, err)

			td, err := statsgen.AnalyzeStatsFolder(dir, "REDACTED")
			if verifyInstance.mustFault {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, verifyInstance.anticipated, td)
			}
		})
	}
}

func VerifyAnalyzeRenamedIndicator(t *testing.T) {
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
	td, err := statsgen.AnalyzeStatsFolder(dir, "REDACTED")
	require.NoError(t, err)

	anticipated := statsgen.PrototypeData{
		Module: "REDACTED",
		AnalyzedStats: []statsgen.AnalyzedIndicatorField{
			{
				KindLabel:   "REDACTED",
				FieldLabel:  "REDACTED",
				IndicatorLabel: "REDACTED",
			},
		},
	}
	require.Equal(t, anticipated, td)
}
