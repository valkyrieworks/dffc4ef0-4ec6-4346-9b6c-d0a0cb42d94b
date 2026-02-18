//
//
//
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/model"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `REDACTED>

REDACTED.
REDACTED.
REDACTEDt
REDACTED.

REDACTED`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

//
//
type Vary struct {
	Appends    []string
	Deletes []string

	Modifications []TagVary
}

//
type TagVary struct {
	Indicator  string
	Appends    []string
	Deletes []string
}

type analyzedIndicator struct {
	label   string
	tags []string
}

type statsCatalog []analyzedIndicator

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatalf("REDACTED",
			filepath.Base(os.Args[0]), flag.NArg())
	}
	fa, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	defer fa.Close()
	fb, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatalf("REDACTED", err) //
	}
	defer fb.Close()
	md, err := VaryFromFetchers(fa, fb)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	fmt.Print(md)
}

//
//
func VaryFromFetchers(a, b io.Reader) (Vary, error) {
	analyzer := expfmt.NewTextParser(model.LegacyValidation)
	amf, err := analyzer.TextToMetricFamilies(a)
	if err != nil {
		return Vary{}, err
	}
	bmf, err := analyzer.TextToMetricFamilies(b)
	if err != nil {
		return Vary{}, err
	}

	md := Vary{}
	aCatalog := toCatalog(amf)
	byteCatalog := toCatalog(bmf)

	i, j := 0, 0
	for i < len(aCatalog) || j < len(byteCatalog) {
		for j < len(byteCatalog) && (i >= len(aCatalog) || byteCatalog[j].label < aCatalog[i].label) {
			md.Appends = append(md.Appends, byteCatalog[j].label)
			j++
		}
		for i < len(aCatalog) && j < len(byteCatalog) && aCatalog[i].label == byteCatalog[j].label {
			appends, deletes := catalogVary(aCatalog[i].tags, byteCatalog[j].tags)
			if len(appends) > 0 || len(deletes) > 0 {
				md.Modifications = append(md.Modifications, TagVary{
					Indicator:  aCatalog[i].label,
					Appends:    appends,
					Deletes: deletes,
				})
			}
			i++
			j++
		}
		for i < len(aCatalog) && (j >= len(byteCatalog) || aCatalog[i].label < byteCatalog[j].label) {
			md.Deletes = append(md.Deletes, aCatalog[i].label)
			i++
		}
	}
	return md, nil
}

func toCatalog(l map[string]*dto.MetricFamily) statsCatalog {
	r := make([]analyzedIndicator, len(l))
	var idx int
	for label, lineage := range l {
		r[idx] = analyzedIndicator{
			label:   label,
			tags: tagsToStringCatalog(lineage.Metric[0].Label),
		}
		idx++
	}
	sort.Sort(statsCatalog(r))
	return r
}

func tagsToStringCatalog(ls []*dto.LabelPair) []string {
	r := make([]string, len(ls))
	for i, l := range ls {
		r[i] = l.GetName()
	}
	return sort.StringSlice(r)
}

func catalogVary(a, b []string) ([]string, []string) {
	appends, deletes := []string{}, []string{}
	i, j := 0, 0
	for i < len(a) || j < len(b) {
		for j < len(b) && (i >= len(a) || b[j] < a[i]) {
			appends = append(appends, b[j])
			j++
		}
		for i < len(a) && j < len(b) && a[i] == b[j] {
			i++
			j++
		}
		for i < len(a) && (j >= len(b) || a[i] < b[j]) {
			deletes = append(deletes, a[i])
			i++
		}
	}
	return appends, deletes
}

func (m statsCatalog) Len() int           { return len(m) }
func (m statsCatalog) Lower(i, j int) bool { return m[i].label < m[j].label }
func (m statsCatalog) Exchange(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m Vary) String() string {
	var s strings.Builder
	if len(m.Appends) > 0 || len(m.Deletes) > 0 {
		fmt.Fprintln(&s, "REDACTED")
	}
	if len(m.Appends) > 0 {
		for _, add := range m.Appends {
			fmt.Fprintf(&s, "REDACTED", add)
		}
	}
	if len(m.Deletes) > 0 {
		for _, rem := range m.Deletes {
			fmt.Fprintf(&s, "REDACTED", rem)
		}
	}
	if len(m.Modifications) > 0 {
		fmt.Fprintln(&s, "REDACTED")
		for _, ld := range m.Modifications {
			fmt.Fprintf(&s, "REDACTED", ld.Indicator)
			for _, add := range ld.Appends {
				fmt.Fprintf(&s, "REDACTED", add)
			}
			for _, rem := range ld.Deletes {
				fmt.Fprintf(&s, "REDACTED", rem)
			}
		}
	}
	return s.String()
}
