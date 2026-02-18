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

	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire/grammar"
)

//
var All *Inquire

//
type Inquire struct {
	ast   grammar.Inquire
	states []state
}

//
func New(inquire string) (*Inquire, error) {
	ast, err := grammar.Analyze(inquire)
	if err != nil {
		return nil, err
	}
	return Build(ast)
}

//
//
//
//
//
func ShouldBuild(inquire string) *Inquire {
	q, err := New(inquire)
	if err != nil {
		panic(err)
	}
	return q
}

//
func Build(ast grammar.Inquire) (*Inquire, error) {
	states := make([]state, len(ast))
	for i, q := range ast {
		stat, err := buildState(q)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", q, err)
		}
		states[i] = stat
	}
	return &Inquire{ast: ast, states: states}, nil
}

func ExtendEvents(unfoldedEvents map[string][]string) []kinds.Event {
	events := make([]kinds.Event, 0)

	for compound, items := range unfoldedEvents {
		symbols := strings.Split(compound, "REDACTED")

		props := make([]kinds.EventProperty, len(items))
		for i, v := range items {
			props[i] = kinds.EventProperty{
				Key:   symbols[len(symbols)-1],
				Item: v,
			}
		}

		events = append(events, kinds.Event{
			Kind:       strings.Join(symbols[:len(symbols)-1], "REDACTED"),
			Properties: props,
		})
	}

	return events
}

//
//
func (q *Inquire) Aligns(events map[string][]string) (bool, error) {
	if q == nil {
		return true, nil
	}
	return q.alignsEvents(ExtendEvents(events)), nil
}

//
func (q *Inquire) String() string {
	if q == nil {
		return "REDACTED"
	}
	return q.ast.String()
}

//
func (q *Inquire) Grammar() grammar.Inquire {
	if q == nil {
		return nil
	}
	return q.ast
}

//
func (q *Inquire) alignsEvents(events []kinds.Event) bool {
	for _, stat := range q.states {
		if !stat.alignsAny(events) {
			return false
		}
	}
	return len(events) != 0
}

//
//
//
type state struct {
	tag   string //
	align func(s string) bool
}

//
//
//
func (c state) locateProp(event kinds.Event) ([]string, bool) {
	if !strings.HasPrefix(c.tag, event.Kind) {
		return nil, false //
	} else if len(c.tag) == len(event.Kind) {
		return nil, true //
	}
	var values []string
	for _, property := range event.Properties {
		completeLabel := event.Kind + "REDACTED" + property.Key
		if completeLabel == c.tag {
			values = append(values, property.Item)
		}
	}
	return values, false
}

//
func (c state) alignsAny(events []kinds.Event) bool {
	for _, event := range events {
		if c.alignsEvent(event) {
			return true
		}
	}
	return false
}

//
func (c state) alignsEvent(event kinds.Event) bool {
	vs, markerMatchesKind := c.locateProp(event)
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

func buildState(stat grammar.State) (state, error) {
	out := state{tag: stat.Tag}

	//
	//
	if stat.Op == grammar.TPresent {
		out.align = func(string) bool { return true }
		return out, nil
	}

	//
	if stat.Arg == nil {
		return state{}, fmt.Errorf("REDACTED", stat.Op)
	}

	//
	argumentKind := stat.Arg.Kind
	var argumentItem any

	switch argumentKind {
	case grammar.TString:
		argumentItem = stat.Arg.Item()
	case grammar.TAmount:
		argumentItem = stat.Arg.Amount()
	case grammar.TTime, grammar.TDate:
		argumentItem = stat.Arg.Time()
	default:
		return state{}, fmt.Errorf("REDACTED", argumentKind)
	}

	mconfig := actKindIndex[stat.Op][argumentKind]
	if mconfig == nil {
		return state{}, fmt.Errorf("REDACTED", stat.Op, argumentKind)
	}
	out.align = mconfig(argumentItem)
	return out, nil
}

//
//
//
var retrieveCount = regexp.MustCompile("REDACTED")

func analyzeAmount(s string) (*big.Float, error) {
	integerValue := new(big.Int)
	if _, ok := integerValue.SetString(s, 10); !ok {
		f, _, err := big.ParseFloat(retrieveCount.FindString(s), 10, 125, big.ToNearestEven)
		if err != nil {
			return nil, err
		}
		return f, err
	}
	f, _, err := big.ParseFloat(retrieveCount.FindString(s), 10, uint(integerValue.BitLen()), big.ToNearestEven)
	return f, err
}

//
//
//
//
//
//
var actKindIndex = map[grammar.Symbol]map[grammar.Symbol]func(any) func(string) bool{
	grammar.TIncludes: {
		grammar.TString: func(v any) func(string) bool {
			return func(s string) bool {
				return strings.Contains(s, v.(string))
			}
		},
	},
	grammar.TEq: {
		grammar.TString: func(v any) func(string) bool {
			return func(s string) bool { return s == v.(string) }
		},
		grammar.TAmount: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeAmount(s)
				return err == nil && w.Cmp(v.(*big.Float)) == 0
			}
		},
		grammar.TDate: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeDate(s)
				return err == nil && ts.Equal(v.(time.Time))
			}
		},
		grammar.TTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && ts.Equal(v.(time.Time))
			}
		},
	},
	grammar.TLt: {
		grammar.TAmount: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeAmount(s)
				return err == nil && w.Cmp(v.(*big.Float)) < 0
			}
		},
		grammar.TDate: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeDate(s)
				return err == nil && ts.Before(v.(time.Time))
			}
		},
		grammar.TTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && ts.Before(v.(time.Time))
			}
		},
	},
	grammar.TLeq: {
		grammar.TAmount: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeAmount(s)
				return err == nil && w.Cmp(v.(*big.Float)) <= 0
			}
		},
		grammar.TDate: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeDate(s)
				return err == nil && !ts.After(v.(time.Time))
			}
		},
		grammar.TTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && !ts.After(v.(time.Time))
			}
		},
	},
	grammar.TGt: {
		grammar.TAmount: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeAmount(s)
				return err == nil && w.Cmp(v.(*big.Float)) > 0
			}
		},
		grammar.TDate: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeDate(s)
				return err == nil && ts.After(v.(time.Time))
			}
		},
		grammar.TTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && ts.After(v.(time.Time))
			}
		},
	},
	grammar.TGeq: {
		grammar.TAmount: func(v any) func(string) bool {
			return func(s string) bool {
				w, err := analyzeAmount(s)
				return err == nil && w.Cmp(v.(*big.Float)) >= 0
			}
		},
		grammar.TDate: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeDate(s)
				return err == nil && !ts.Before(v.(time.Time))
			}
		},
		grammar.TTime: func(v any) func(string) bool {
			return func(s string) bool {
				ts, err := grammar.AnalyzeTime(s)
				return err == nil && !ts.Before(v.(time.Time))
			}
		},
	},
}
