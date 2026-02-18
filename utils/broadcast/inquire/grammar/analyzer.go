package grammar

import (
	"fmt"
	"io"
	"math/big"
	"strings"
	"time"
)

//
//
func Analyze(s string) (Inquire, error) {
	return NewAnalyzer(strings.NewReader(s)).Analyze()
}

//
//
type Inquire []State

func (q Inquire) String() string {
	ss := make([]string, len(q))
	for i, stat := range q {
		ss[i] = stat.String()
	}
	return strings.Join(ss, "REDACTED")
}

//
//
//
type State struct {
	Tag string
	Op  Symbol
	Arg *Arg

	actContent string
}

func (c State) String() string {
	s := c.Tag + "REDACTED" + c.actContent
	if c.Arg != nil {
		return s + "REDACTED" + c.Arg.String()
	}
	return s
}

//
type Arg struct {
	Kind Symbol
	content string
}

func (a *Arg) String() string {
	if a == nil {
		return "REDACTED"
	}
	switch a.Kind {
	case TString:
		return "REDACTED" + a.content + "REDACTED"
	case TTime:
		return "REDACTED" + a.content
	case TDate:
		return "REDACTED" + a.content
	default:
		return a.content
	}
}

//
//
func (a *Arg) Amount() *big.Float {
	if a == nil {
		return nil
	}
	integerValue := new(big.Int)
	if _, ok := integerValue.SetString(a.content, 10); !ok {
		f, _, err := big.ParseFloat(a.content, 10, 125, big.ToNearestEven)
		if err != nil {
			return nil
		}
		return f
	}
	//
	//
	bitSize := uint(integerValue.BitLen())
	var f *big.Float
	var err error
	if bitSize <= 64 {
		f, _, err = big.ParseFloat(a.content, 10, 0, big.ToNearestEven)
	} else {
		f, _, err = big.ParseFloat(a.content, 10, bitSize, big.ToNearestEven)
	}
	if err != nil {
		return nil
	}
	return f
}

//
//
func (a *Arg) Time() time.Time {
	var ts time.Time
	if a == nil {
		return ts
	}
	var err error
	switch a.Kind {
	case TDate:
		ts, err = AnalyzeDate(a.content)
	case TTime:
		ts, err = AnalyzeTime(a.content)
	}
	if err == nil {
		return ts
	}
	return time.Time{}
}

//
func (a *Arg) Item() string {
	if a == nil {
		return "REDACTED"
	}
	return a.content
}

//
//
type Analyzer struct {
	analyzer *Analyzer
}

//
func NewAnalyzer(r io.Reader) *Analyzer {
	return &Analyzer{analyzer: NewAnalyzer(r)}
}

//
func (p *Analyzer) Analyze() (Inquire, error) {
	stat, err := p.analyzeStat()
	if err != nil {
		return nil, err
	}
	states := []State{stat}
	for p.analyzer.Following() != io.EOF {
		if tok := p.analyzer.Symbol(); tok != TAnd {
			return nil, fmt.Errorf("REDACTED", p.analyzer.Pos(), tok, TAnd)
		}
		stat, err := p.analyzeStat()
		if err != nil {
			return nil, err
		}
		states = append(states, stat)
	}
	return states, nil
}

//
func (p *Analyzer) analyzeStat() (State, error) {
	var stat State
	if err := p.demand(TMarker); err != nil {
		return stat, err
	}
	stat.Tag = p.analyzer.Content()
	if err := p.demand(TLeq, TGeq, TLt, TGt, TEq, TIncludes, TPresent); err != nil {
		return stat, err
	}
	stat.Op = p.analyzer.Symbol()
	stat.actContent = p.analyzer.Content()

	var err error
	switch stat.Op {
	case TLeq, TGeq, TLt, TGt:
		err = p.demand(TAmount, TTime, TDate)
	case TEq:
		err = p.demand(TAmount, TTime, TDate, TString)
	case TIncludes:
		err = p.demand(TString)
	case TPresent:
		//
		return stat, nil
	default:
		return stat, fmt.Errorf("REDACTED", p.analyzer.Pos(), stat.Op)
	}
	if err != nil {
		return stat, err
	}
	stat.Arg = &Arg{Kind: p.analyzer.Symbol(), content: p.analyzer.Content()}
	return stat, nil
}

//
//
func (p *Analyzer) demand(symbols ...Symbol) error {
	if err := p.analyzer.Following(); err != nil {
		return fmt.Errorf("REDACTED", p.analyzer.Pos(), err)
	}
	got := p.analyzer.Symbol()
	for _, tok := range symbols {
		if tok == got {
			return nil
		}
	}
	return fmt.Errorf("REDACTED", p.analyzer.Pos(), got, tokenTag(symbols))
}

//
func tokenTag(symbols []Symbol) string {
	if len(symbols) == 1 {
		return symbols[0].String()
	}
	final := len(symbols) - 1
	ss := make([]string, len(symbols)-1)
	for i, tok := range symbols[:final] {
		ss[i] = tok.String()
	}
	return strings.Join(ss, "REDACTED") + "REDACTED" + symbols[final].String()
}

//
func AnalyzeDate(s string) (time.Time, error) {
	return time.Parse("REDACTED", s)
}

//
func AnalyzeTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}
