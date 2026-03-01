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
	return FreshTokenizer(strings.NewReader(s)).Analyze()
}

//
//
type Inquire []Stipulation

func (q Inquire) Text() string {
	ss := make([]string, len(q))
	for i, stipulation := range q {
		ss[i] = stipulation.Text()
	}
	return strings.Join(ss, "REDACTED")
}

//
//
//
type Stipulation struct {
	Tag string
	Op  Symbol
	Arg *Arg

	actionString string
}

func (c Stipulation) Text() string {
	s := c.Tag + "REDACTED" + c.actionString
	if c.Arg != nil {
		return s + "REDACTED" + c.Arg.Text()
	}
	return s
}

//
type Arg struct {
	Kind Symbol
	string string
}

func (a *Arg) Text() string {
	if a == nil {
		return "REDACTED"
	}
	switch a.Kind {
	case TYPText:
		return "REDACTED" + a.string + "REDACTED"
	case TYPMoment:
		return "REDACTED" + a.string
	case TYPTime:
		return "REDACTED" + a.string
	default:
		return a.string
	}
}

//
//
func (a *Arg) Numeral() *big.Float {
	if a == nil {
		return nil
	}
	integerItem := new(big.Int)
	if _, ok := integerItem.SetString(a.string, 10); !ok {
		f, _, err := big.ParseFloat(a.string, 10, 125, big.ToNearestEven)
		if err != nil {
			return nil
		}
		return f
	}
	//
	//
	digitLength := uint(integerItem.BitLen())
	var f *big.Float
	var err error
	if digitLength <= 64 {
		f, _, err = big.ParseFloat(a.string, 10, 0, big.ToNearestEven)
	} else {
		f, _, err = big.ParseFloat(a.string, 10, digitLength, big.ToNearestEven)
	}
	if err != nil {
		return nil
	}
	return f
}

//
//
func (a *Arg) Moment() time.Time {
	var ts time.Time
	if a == nil {
		return ts
	}
	var err error
	switch a.Kind {
	case TYPTime:
		ts, err = AnalyzeTime(a.string)
	case TYPMoment:
		ts, err = AnalyzeMoment(a.string)
	}
	if err == nil {
		return ts
	}
	return time.Time{}
}

//
func (a *Arg) Datum() string {
	if a == nil {
		return "REDACTED"
	}
	return a.string
}

//
//
type Tokenizer struct {
	analyzer *Analyzer
}

//
func FreshTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{analyzer: FreshAnalyzer(r)}
}

//
func (p *Tokenizer) Analyze() (Inquire, error) {
	stipulation, err := p.analyzeStipulation()
	if err != nil {
		return nil, err
	}
	stipulations := []Stipulation{stipulation}
	for p.analyzer.Following() != io.EOF {
		if tok := p.analyzer.Symbol(); tok != TYPAlso {
			return nil, fmt.Errorf("REDACTED", p.analyzer.Pos(), tok, TYPAlso)
		}
		stipulation, err := p.analyzeStipulation()
		if err != nil {
			return nil, err
		}
		stipulations = append(stipulations, stipulation)
	}
	return stipulations, nil
}

//
func (p *Tokenizer) analyzeStipulation() (Stipulation, error) {
	var stipulation Stipulation
	if err := p.demand(TYPMarker); err != nil {
		return stipulation, err
	}
	stipulation.Tag = p.analyzer.String()
	if err := p.demand(TYPLesseq, TYPGreatereq, TLt, TGt, TEq, TYPIncludes, TYPPresent); err != nil {
		return stipulation, err
	}
	stipulation.Op = p.analyzer.Symbol()
	stipulation.actionString = p.analyzer.String()

	var err error
	switch stipulation.Op {
	case TYPLesseq, TYPGreatereq, TLt, TGt:
		err = p.demand(TYPNumeral, TYPMoment, TYPTime)
	case TEq:
		err = p.demand(TYPNumeral, TYPMoment, TYPTime, TYPText)
	case TYPIncludes:
		err = p.demand(TYPText)
	case TYPPresent:
		//
		return stipulation, nil
	default:
		return stipulation, fmt.Errorf("REDACTED", p.analyzer.Pos(), stipulation.Op)
	}
	if err != nil {
		return stipulation, err
	}
	stipulation.Arg = &Arg{Kind: p.analyzer.Symbol(), string: p.analyzer.String()}
	return stipulation, nil
}

//
//
func (p *Tokenizer) demand(symbols ...Symbol) error {
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
		return symbols[0].Text()
	}
	final := len(symbols) - 1
	ss := make([]string, len(symbols)-1)
	for i, tok := range symbols[:final] {
		ss[i] = tok.Text()
	}
	return strings.Join(ss, "REDACTED") + "REDACTED" + symbols[final].Text()
}

//
func AnalyzeTime(s string) (time.Time, error) {
	return time.Parse("REDACTED", s)
}

//
func AnalyzeMoment(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}
