//
//
//
//
//
//
//
//
package inquire

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
)

//
var All *Inquire

//
type Inquire struct {
	ast   grammar.Inquire
	stipulations []stipulation
}

//
func New(inquire string) (*Inquire, error) {
	ast, err := grammar.Analyze(inquire)
	if err != nil {
		return nil, err
	}
	return Assemble(ast)
}

//
//
//
//
//
func ShouldAssemble(inquire string) *Inquire {
	q, err := New(inquire)
	if err != nil {
		panic(err)
	}
	return q
}

//
func Assemble(ast grammar.Inquire) (*Inquire, error) {
	stipulations := make([]stipulation, len(ast))
	for i, q := range ast {
		stipulation, err := assembleStipulation(q)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", q, err)
		}
		stipulations[i] = stipulation
	}
	return &Inquire{ast: ast, stipulations: stipulations}, nil
}

func AugmentIncidents(collapsedIncidents map[string][]string) []kinds.Incident {
	incidents := make([]kinds.Incident, 0)

	for complex, items := range collapsedIncidents {
		symbols := strings.Split(complex, "REDACTED")

		traits := make([]kinds.IncidentProperty, len(items))
		for i, v := range items {
			traits[i] = kinds.IncidentProperty{
				Key:   symbols[len(symbols)-1],
				Datum: v,
			}
		}

		incidents = append(incidents, kinds.Incident{
			Kind:       strings.Join(symbols[:len(symbols)-1], "REDACTED"),
			Properties: traits,
		})
	}

	return incidents
}

//
//
func (q *Inquire) Aligns(incidents map[string][]string) (bool, error) {
	if q == nil {
		return true, nil
	}
	return q.alignsIncidents(AugmentIncidents(incidents)), nil
}

//
func (q *Inquire) Text() string {
	if q == nil {
		return "REDACTED"
	}
	return q.ast.Text()
}

//
func (q *Inquire) Grammar() grammar.Inquire {
	if q == nil {
		return nil
	}
	return q.ast
}

//
func (q *Inquire) alignsIncidents(incidents []kinds.Incident) bool {
	for _, stipulation := range q.stipulations {
		if !stipulation.alignsSome(incidents) {
			return false
		}
	}
	return len(incidents) != 0
}

//
//
//
type stipulation struct {
	tag   string //
	align func(s string) bool
}

//
//
//
func (c stipulation) locateTrait(incident kinds.Incident) ([]string, bool) {
	if !strings.HasPrefix(c.tag, incident.Kind) {
		return nil, false //
	} else if len(c.tag) == len(incident.Kind) {
		return nil, true //
	}
	var values []string
	for _, property := range incident.Properties {
		completeAlias := incident.Kind + "REDACTED" + property.Key
		if completeAlias == c.tag {
			values = append(values, property.Datum)
		}
	}
	return values, false
}

//
func (c stipulation) alignsSome(incidents []kinds.Incident) bool {
	for _, incident := range incidents {
		if c.alignsIncident(incident) {
			return true
		}
	}
	return false
}

//
func (c stipulation) alignsIncident(incident kinds.Incident) bool {
	vs, markerMatchesKind := c.locateTrait(incident)
	if len(vs) == 0 {
		//
		//
		//
		if markerMatchesKind {
			return c.align("REDACTED")
		}
		return false
	}

	//
	for _, v := range vs {
		if c.align(v) {
			return true
		}
	}
	return false
}

func assembleStipulation(stipulation grammar.Stipulation) (stipulation, error) {
	out := stipulation{tag: stipulation.Tag}

	//
	//
	if stipulation.Op == grammar.TYPPresent {
		out.align = func(string) bool { return true }
		return out, nil
	}

	//
	if stipulation.Arg == nil {
		return stipulation{}, fmt.Errorf("REDACTED", stipulation.Op)
	}

	//
	argumentKind := stipulation.Arg.Kind
	var argumentDatum any

	switch argumentKind {
	case grammar.TYPText:
		argumentDatum = stipulation.Arg.Datum()
	case grammar.TYPNumeral:
		argumentDatum = stipulation.Arg.Numeral()
	case grammar.TYPMoment, grammar.TYPTime:
		argumentDatum = stipulation.Arg.Moment()
	default:
		return stipulation{}, fmt.Errorf("REDACTED", argumentKind)
	}

	mledger := actionKindIndex[stipulation.Op][argumentKind]
	if mledger == nil {
		return stipulation{}, fmt.Errorf("REDACTED", stipulation.Op, argumentKind)
	}
	out.align = mledger(argumentDatum)
	return out, nil
}

//
//
//
var deriveCount = regexp.MustCompile("REDACTED")

func analyzeNumeral(s string) (*big.Float, error) {
	integerItem := new(big.Int)
	if _, ok := integerItem.SetString(s, 10); !ok {
		f, _, err := big.ParseFloat(deriveCount.FindString(s), 10, 125, big.ToNearestEven)
		if err != nil {
			return nil, err
		}
		return f, err
	}
	f, _, err := big.ParseFloat(deriveCount.FindString(s), 10, uint(integerItem.BitLen()), big.ToNearestEven)
	return f, err
}

//
//
//
//
//
//
var actionKindIndex = map[grammar.Symbol]map[grammar.Symbol]func(any) func(string) bool{
	grammar.TYPIncludes: {
		grammar.TYPText: func(v any) func(string) bool {
			return func(s string) bool {
				return strings.Contains(s, v.(string))
			}
		},
	},
	grammar.TEq: {
		grammar.TYPText: func(v any) func(string) bool {
			return func(s string) bool { return s == v.(string) }
		},
		grammar.TYPNumeral: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeNumeral(s)
				return err == nil && w.Cmp(v.(*big.Float)) == 0
			}
		},
		grammar.TYPTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && ts.Equal(v.(time.Time))
			}
		},
		grammar.TYPMoment: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeMoment(s)
				return err == nil && ts.Equal(v.(time.Time))
			}
		},
	},
	grammar.TLt: {
		grammar.TYPNumeral: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeNumeral(s)
				return err == nil && w.Cmp(v.(*big.Float)) < 0
			}
		},
		grammar.TYPTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && ts.Before(v.(time.Time))
			}
		},
		grammar.TYPMoment: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeMoment(s)
				return err == nil && ts.Before(v.(time.Time))
			}
		},
	},
	grammar.TYPLesseq: {
		grammar.TYPNumeral: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeNumeral(s)
				return err == nil && w.Cmp(v.(*big.Float)) <= 0
			}
		},
		grammar.TYPTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && !ts.After(v.(time.Time))
			}
		},
		grammar.TYPMoment: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeMoment(s)
				return err == nil && !ts.After(v.(time.Time))
			}
		},
	},
	grammar.TGt: {
		grammar.TYPNumeral: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeNumeral(s)
				return err == nil && w.Cmp(v.(*big.Float)) > 0
			}
		},
		grammar.TYPTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && ts.After(v.(time.Time))
			}
		},
		grammar.TYPMoment: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeMoment(s)
				return err == nil && ts.After(v.(time.Time))
			}
		},
	},
	grammar.TYPGreatereq: {
		grammar.TYPNumeral: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeNumeral(s)
				return err == nil && w.Cmp(v.(*big.Float)) >= 0
			}
		},
		grammar.TYPTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && !ts.Before(v.(time.Time))
			}
		},
		grammar.TYPMoment: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeMoment(s)
				return err == nil && !ts.Before(v.(time.Time))
			}
		},
	},
}
