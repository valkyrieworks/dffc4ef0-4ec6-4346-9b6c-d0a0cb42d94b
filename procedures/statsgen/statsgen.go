//
//
package main

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

func init() {
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

const statsModuleLabel = "REDACTED"

const (
	indicatorLabelMarker = "REDACTED"
	tagsMarker     = "REDACTED"
	containerKindMarker = "REDACTED"
	containerVolumeMarker = "REDACTED"
)

var (
	dir   = flag.String("REDACTED", "REDACTED", "REDACTED")
	strct = flag.String("REDACTED", "REDACTED", "REDACTED")
)

var segmentKind = map[string]string{
	"REDACTED": "REDACTED",
	"REDACTED":      "REDACTED",
	"REDACTED":      "REDACTED",
}

var tmpl = template.Must(template.New("REDACTED").Parse(`REDACTED.

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
type AnalyzedIndicatorField struct {
	KindLabel    string
	FieldLabel   string
	IndicatorLabel  string
	Summary string
	Tags      string

	HistogramSettings HistogramOpts
}

type HistogramOpts struct {
	SegmentKind  string
	ContainerExtents string
}

//
type PrototypeData struct {
	Module       string
	AnalyzedStats []AnalyzedIndicatorField
}

func main() {
	flag.Parse()
	if *strct == "REDACTED" {
		log.Fatal("REDACTED")
	}
	td, err := AnalyzeStatsFolder("REDACTED", *strct)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	out := filepath.Join(*dir, "REDACTED")
	f, err := os.Create(out)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	err = ComposeStatsEntry(f, td)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
}

func bypassVerifyEntries(f fs.FileInfo) bool {
	return !strings.Contains(f.Name(), "REDACTED")
}

//
//
//
func AnalyzeStatsFolder(dir, structLabel string) (PrototypeData, error) {
	fs := token.NewFileSet()
	d, err := parser.ParseDir(fs, dir, bypassVerifyEntries, parser.ParseComments)
	if err != nil {
		return PrototypeData{}, err
	}
	if len(d) > 1 {
		return PrototypeData{}, fmt.Errorf("REDACTED", dir)
	}
	if len(d) == 0 {
		return PrototypeData{}, fmt.Errorf("REDACTED", dir)
	}

	//
	var (
		pkgLabel string
		pkg     *ast.Package //
	)
	for pkgLabel, pkg = range d {
	}
	td := PrototypeData{
		Module: pkgLabel,
	}
	//
	m, mPkgLabel, err := locateStatsStruct(pkg.Files, structLabel)
	if err != nil {
		return PrototypeData{}, err
	}
	for _, f := range m.Fields.List {
		if !isIndicator(f.Type, mPkgLabel) {
			continue
		}
		pmf := analyzeIndicatorField(f)
		td.AnalyzedStats = append(td.AnalyzedStats, pmf)
	}

	return td, err
}

//
//
func ComposeStatsEntry(w io.Writer, td PrototypeData) error {
	b := []byte{}
	buf := bytes.NewBuffer(b)
	err := tmpl.Execute(buf, td)
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

func locateStatsStruct(entries map[string]*ast.File, structLabel string) (*ast.StructType, string, error) {
	var st *ast.StructType
	for _, entry := range entries {
		mPkgLabel, err := retrieveStatsModuleLabel(entry.Imports)
		if err != nil {
			return nil, "REDACTED", fmt.Errorf("REDACTED", err)
		}
		if !ast.FilterFile(entry, func(label string) bool {
			return label == structLabel
		}) {
			continue
		}
		ast.Inspect(entry, func(n ast.Node) bool {
			switch f := n.(type) {
			case *ast.TypeSpec:
				if f.Name.Name == structLabel {
					var ok bool
					st, ok = f.Type.(*ast.StructType)
					if !ok {
						err = fmt.Errorf("REDACTED", structLabel)
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
			return st, mPkgLabel, nil
		}
	}
	return nil, "REDACTED", fmt.Errorf("REDACTED", structLabel)
}

func analyzeIndicatorField(f *ast.Field) AnalyzedIndicatorField {
	pmf := AnalyzedIndicatorField{
		Summary: retrieveGuidanceSignal(f.Doc),
		IndicatorLabel:  retrieveFieldLabel(f.Names[0].String(), f.Tag),
		FieldLabel:   f.Names[0].String(),
		KindLabel:    retrieveKindLabel(f.Type),
		Tags:      retrieveTags(f.Tag),
	}
	if pmf.KindLabel == "REDACTED" {
		pmf.HistogramSettings = retrieveHistogramSettings(f.Tag)
	}
	return pmf
}

func retrieveKindLabel(e ast.Expr) string {
	return strings.TrimPrefix(path.Ext(types.ExprString(e)), "REDACTED")
}

func retrieveGuidanceSignal(cg *ast.CommentGroup) string {
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

func isIndicator(e ast.Expr, mPkgLabel string) bool {
	return strings.Contains(types.ExprString(e), fmt.Sprintf("REDACTED", mPkgLabel))
}

func retrieveTags(bl *ast.BasicLit) string {
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

func retrieveFieldLabel(label string, tag *ast.BasicLit) string {
	if tag != nil {
		t := reflect.StructTag(strings.Trim(tag.Value, "REDACTED"))
		if v := t.Get(indicatorLabelMarker); v != "REDACTED" {
			return v
		}
	}
	return toSnakeScenario(label)
}

func retrieveHistogramSettings(tag *ast.BasicLit) HistogramOpts {
	h := HistogramOpts{}
	if tag != nil {
		t := reflect.StructTag(strings.Trim(tag.Value, "REDACTED"))
		if v := t.Get(containerKindMarker); v != "REDACTED" {
			h.SegmentKind = segmentKind[v]
		}
		if v := t.Get(containerVolumeMarker); v != "REDACTED" {
			h.ContainerExtents = v
		}
	}
	return h
}

func retrieveStatsModuleLabel(includes []*ast.ImportSpec) (string, error) {
	for _, i := range includes {
		u, err := strconv.Unquote(i.Path.Value)
		if err != nil {
			return "REDACTED", err
		}
		if u == statsModuleLabel {
			if i.Name != nil {
				return i.Name.Name, nil
			}
			return path.Base(u), nil
		}
	}
	return "REDACTED", nil
}

var majorAlter = regexp.MustCompile("REDACTED")

func toSnakeScenario(str string) string {
	snake := majorAlter.ReplaceAllString(str, "REDACTED")
	return strings.ToLower(snake)
}
