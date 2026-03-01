//
//
package primary

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

func initialize() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `REDACTED>

REDACTEDn
REDACTEDy
REDACTED.

REDACTED:
REDACTED`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

const telemetryBundleAlias = "REDACTED"

const (
	measurementAliasMarker = "REDACTED"
	tagsMarker     = "REDACTED"
	containerKindMarker = "REDACTED"
	containerExtentMarker = "REDACTED"
)

var (
	dir   = flag.String("REDACTED", "REDACTED", "REDACTED")
	layout = flag.String("REDACTED", "REDACTED", "REDACTED")
)

var containerKind = map[string]string{
	"REDACTED": "REDACTED",
	"REDACTED":      "REDACTED",
	"REDACTED":      "REDACTED",
}

var layout = template.Must(template.New("REDACTED").Parse(`REDACTED.

REDACTED}

REDACTED(
REDACTED"
REDACTED"
REDACTED"
)

REDACTED{
REDACTED}
REDACTED{
REDACTED)
REDACTED}
REDACTED{
REDACTED}
REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED}
REDACTED,
REDACTED}
REDACTED,
REDACTED}
REDACTED}
REDACTED,
REDACTED}
REDACTED,
REDACTED}
REDACTED}
REDACTED}
}


REDACTED{
REDACTED{
REDACTED}
REDACTED,
REDACTED}
REDACTED}
}
REDACTED`))

//
type ProcessedMeasurementAttribute struct {
	KindAlias    string
	AttributeAlias   string
	MeasurementAlias  string
	Characterization string
	Tags      string

	HistogramChoices HistogramOptions
}

type HistogramOptions struct {
	ContainerKind  string
	ContainerExtents string
}

//
type BlueprintData struct {
	Bundle       string
	ProcessedTelemetry []ProcessedMeasurementAttribute
}

func primary() {
	flag.Parse()
	if *layout == "REDACTED" {
		log.Fatal("REDACTED")
	}
	td, err := AnalyzeTelemetryPath("REDACTED", *layout)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	out := filepath.Join(*dir, "REDACTED")
	f, err := os.Create(out)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	err = ComposeTelemetryRecord(f, td)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
}

func bypassVerifyRecords(f fs.FileInfo) bool {
	return !strings.Contains(f.Name(), "REDACTED")
}

//
//
//
func AnalyzeTelemetryPath(dir, patternAlias string) (BlueprintData, error) {
	fs := token.NewFileSet()
	d, err := parser.ParseDir(fs, dir, bypassVerifyRecords, parser.ParseComments)
	if err != nil {
		return BlueprintData{}, err
	}
	if len(d) > 1 {
		return BlueprintData{}, fmt.Errorf("REDACTED", dir)
	}
	if len(d) == 0 {
		return BlueprintData{}, fmt.Errorf("REDACTED", dir)
	}

	//
	var (
		modAlias string
		pkg     *ast.Package //
	)
	for modAlias, pkg = range d {
	}
	td := BlueprintData{
		Bundle: modAlias,
	}
	//
	m, moduleModAlias, err := locateTelemetryPattern(pkg.Files, patternAlias)
	if err != nil {
		return BlueprintData{}, err
	}
	for _, f := range m.Fields.List {
		if !equalsMeasurement(f.Type, moduleModAlias) {
			continue
		}
		pmf := analyzeMeasurementAttribute(f)
		td.ProcessedTelemetry = append(td.ProcessedTelemetry, pmf)
	}

	return td, err
}

//
//
func ComposeTelemetryRecord(w io.Writer, td BlueprintData) error {
	b := []byte{}
	buf := bytes.NewBuffer(b)
	err := layout.Execute(buf, td)
	if err != nil {
		return err
	}
	b, err = format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	_, err = io.Copy(w, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	return nil
}

func locateTelemetryPattern(records map[string]*ast.File, patternAlias string) (*ast.StructType, string, error) {
	var st *ast.StructType
	for _, record := range records {
		moduleModAlias, err := deriveTelemetryBundleAlias(record.Imports)
		if err != nil {
			return nil, "REDACTED", fmt.Errorf("REDACTED", err)
		}
		if !ast.FilterFile(record, func(alias string) bool {
			return alias == patternAlias
		}) {
			continue
		}
		ast.Inspect(record, func(n ast.Node) bool {
			switch f := n.(type) {
			case *ast.TypeSpec:
				if f.Name.Name == patternAlias {
					var ok bool
					st, ok = f.Type.(*ast.StructType)
					if !ok {
						err = fmt.Errorf("REDACTED", patternAlias)
					}
				}
				return false
			default:
				return true
			}
		})
		if err != nil {
			return nil, "REDACTED", err
		}
		if st != nil {
			return st, moduleModAlias, nil
		}
	}
	return nil, "REDACTED", fmt.Errorf("REDACTED", patternAlias)
}

func analyzeMeasurementAttribute(f *ast.Field) ProcessedMeasurementAttribute {
	pmf := ProcessedMeasurementAttribute{
		Characterization: deriveGuidanceArtifact(f.Doc),
		MeasurementAlias:  deriveAttributeAlias(f.Names[0].String(), f.Tag),
		AttributeAlias:   f.Names[0].String(),
		KindAlias:    deriveKindAlias(f.Type),
		Tags:      deriveTags(f.Tag),
	}
	if pmf.KindAlias == "REDACTED" {
		pmf.HistogramChoices = deriveHistogramChoices(f.Tag)
	}
	return pmf
}

func deriveKindAlias(e ast.Expr) string {
	return strings.TrimPrefix(path.Ext(types.ExprString(e)), "REDACTED")
}

func deriveGuidanceArtifact(cg *ast.CommentGroup) string {
	if cg == nil {
		return "REDACTED"
	}
	var guidance []string
	for _, c := range cg.List {
		mt := strings.TrimPrefix(c.Text, "REDACTED")
		if mt != c.Text {
			return strings.TrimSpace(mt)
		}
		guidance = append(guidance, strings.TrimSpace(strings.TrimPrefix(c.Text, "REDACTED")))
	}
	return strings.Join(guidance, "REDACTED")
}

func equalsMeasurement(e ast.Expr, moduleModAlias string) bool {
	return strings.Contains(types.ExprString(e), fmt.Sprintf("REDACTED", moduleModAlias))
}

func deriveTags(bl *ast.BasicLit) string {
	if bl != nil {
		t := reflect.StructTag(strings.Trim(bl.Value, "REDACTED"))
		if v := t.Get(tagsMarker); v != "REDACTED" {
			var res []string
			for _, s := range strings.Split(v, "REDACTED") {
				res = append(res, strconv.Quote(strings.TrimSpace(s)))
			}
			return strings.Join(res, "REDACTED")
		}
	}
	return "REDACTED"
}

func deriveAttributeAlias(alias string, tag *ast.BasicLit) string {
	if tag != nil {
		t := reflect.StructTag(strings.Trim(tag.Value, "REDACTED"))
		if v := t.Get(measurementAliasMarker); v != "REDACTED" {
			return v
		}
	}
	return towardSerpentineScenario(alias)
}

func deriveHistogramChoices(tag *ast.BasicLit) HistogramOptions {
	h := HistogramOptions{}
	if tag != nil {
		t := reflect.StructTag(strings.Trim(tag.Value, "REDACTED"))
		if v := t.Get(containerKindMarker); v != "REDACTED" {
			h.ContainerKind = containerKind[v]
		}
		if v := t.Get(containerExtentMarker); v != "REDACTED" {
			h.ContainerExtents = v
		}
	}
	return h
}

func deriveTelemetryBundleAlias(modules []*ast.ImportSpec) (string, error) {
	for _, i := range modules {
		u, err := strconv.Unquote(i.Path.Value)
		if err != nil {
			return "REDACTED", err
		}
		if u == telemetryBundleAlias {
			if i.Name != nil {
				return i.Name.Name, nil
			}
			return path.Base(u), nil
		}
	}
	return "REDACTED", nil
}

var principalAlteration = regexp.MustCompile("REDACTED")

func towardSerpentineScenario(str string) string {
	serpentine := principalAlteration.ReplaceAllString(str, "REDACTED")
	return strings.ToLower(serpentine)
}
