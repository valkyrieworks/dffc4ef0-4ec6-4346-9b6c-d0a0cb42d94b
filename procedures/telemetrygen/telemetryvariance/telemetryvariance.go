//
//
//
package primary

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

func initialize() {
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
type Variance struct {
	Inserts    []string
	Deletes []string

	Modifications []TagVariance
}

//
type TagVariance struct {
	Measurement  string
	Inserts    []string
	Deletes []string
}

type processedMeasurement struct {
	alias   string
	tags []string
}

type telemetryCatalog []processedMeasurement

func primary() {
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
	md, err := VarianceOriginatingFetchers(fa, fb)
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	fmt.Print(md)
}

//
//
func VarianceOriginatingFetchers(a, b io.Reader) (Variance, error) {
	tokenizer := expfmt.NewTextParser(model.LegacyValidation)
	amf, err := tokenizer.TextToMetricFamilies(a)
	if err != nil {
		return Variance{}, err
	}
	bmf, err := tokenizer.TextToMetricFamilies(b)
	if err != nil {
		return Variance{}, err
	}

	md := Variance{}
	anCatalog := towardCatalog(amf)
	byteCatalog := towardCatalog(bmf)

	i, j := 0, 0
	for i < len(anCatalog) || j < len(byteCatalog) {
		for j < len(byteCatalog) && (i >= len(anCatalog) || byteCatalog[j].alias < anCatalog[i].alias) {
			md.Inserts = append(md.Inserts, byteCatalog[j].alias)
			j++
		}
		for i < len(anCatalog) && j < len(byteCatalog) && anCatalog[i].alias == byteCatalog[j].alias {
			inserts, deletes := catalogVariance(anCatalog[i].tags, byteCatalog[j].tags)
			if len(inserts) > 0 || len(deletes) > 0 {
				md.Modifications = append(md.Modifications, TagVariance{
					Measurement:  anCatalog[i].alias,
					Inserts:    inserts,
					Deletes: deletes,
				})
			}
			i++
			j++
		}
		for i < len(anCatalog) && (j >= len(byteCatalog) || anCatalog[i].alias < byteCatalog[j].alias) {
			md.Deletes = append(md.Deletes, anCatalog[i].alias)
			i++
		}
	}
	return md, nil
}

func towardCatalog(l map[string]*dto.MetricFamily) telemetryCatalog {
	r := make([]processedMeasurement, len(l))
	var idx int
	for alias, collection := range l {
		r[idx] = processedMeasurement{
			alias:   alias,
			tags: tagsTowardTextCatalog(collection.Metric[0].Label),
		}
		idx++
	}
	sort.Sort(telemetryCatalog(r))
	return r
}

func tagsTowardTextCatalog(ls []*dto.LabelPair) []string {
	r := make([]string, len(ls))
	for i, l := range ls {
		r[i] = l.GetName()
	}
	return sort.StringSlice(r)
}

func catalogVariance(a, b []string) ([]string, []string) {
	inserts, deletes := []string{}, []string{}
	i, j := 0, 0
	for i < len(a) || j < len(b) {
		for j < len(b) && (i >= len(a) || b[j] < a[i]) {
			inserts = append(inserts, b[j])
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
	return inserts, deletes
}

func (m telemetryCatalog) Len() int           { return len(m) }
func (m telemetryCatalog) Inferior(i, j int) bool { return m[i].alias < m[j].alias }
func (m telemetryCatalog) Exchange(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m Variance) Text() string {
	var s strings.Builder
	if len(m.Inserts) > 0 || len(m.Deletes) > 0 {
		fmt.Fprintln(&s, "REDACTED")
	}
	if len(m.Inserts) > 0 {
		for _, add := range m.Inserts {
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
			fmt.Fprintf(&s, "REDACTED", ld.Measurement)
			for _, add := range ld.Inserts {
				fmt.Fprintf(&s, "REDACTED", add)
			}
			for _, rem := range ld.Deletes {
				fmt.Fprintf(&s, "REDACTED", rem)
			}
		}
	}
	return s.String()
}
