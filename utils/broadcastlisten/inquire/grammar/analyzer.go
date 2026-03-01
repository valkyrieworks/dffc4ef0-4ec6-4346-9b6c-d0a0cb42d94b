package grammar

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
	"unicode"
)

//
type Symbol byte

const (
	TYPUnfit  = iota //
	TYPMarker             //
	TYPText          //
	TYPNumeral          //
	TYPMoment            //
	TYPTime            //
	TYPAlso             //
	TYPIncludes        //
	TYPPresent          //
	TEq              //
	TLt              //
	TYPLesseq             //
	TGt              //
	TYPGreatereq             //

	//
)

var typText = [...]string{
	TYPUnfit:  "REDACTED",
	TYPMarker:      "REDACTED",
	TYPText:   "REDACTED",
	TYPNumeral:   "REDACTED",
	TYPMoment:     "REDACTED",
	TYPTime:     "REDACTED",
	TYPAlso:      "REDACTED",
	TYPIncludes: "REDACTED",
	TYPPresent:   "REDACTED",
	TEq:       "REDACTED",
	TLt:       "REDACTED",
	TYPLesseq:      "REDACTED",
	TGt:       "REDACTED",
	TYPGreatereq:      "REDACTED",
}

func (t Symbol) Text() string {
	v := int(t)
	if v > len(typText) {
		return "REDACTED"
	}
	return typText[v]
}

const (
	//
	MomentLayout = time.RFC3339

	//
	TimeLayout = "REDACTED"
)

//
//
//
type Analyzer struct {
	r   *bufio.Reader
	buf bytes.Buffer
	tok Symbol
	err error

	pos, final, end int
}

//
func FreshAnalyzer(r io.Reader) *Analyzer { return &Analyzer{r: bufio.NewReader(r)} }

//
//
func (s *Analyzer) Following() error {
	s.buf.Reset()
	s.pos = s.end
	s.tok = TYPUnfit
	s.err = nil

	for {
		ch, err := s.character()
		if err != nil {
			return s.abort(err)
		}
		if unicode.IsSpace(ch) {
			s.pos = s.end
			continue //
		}
		if '0' <= ch && ch <= '9' {
			return s.probeNumeral(ch)
		} else if equalsInitialMarkerCharacter(ch) {
			return s.probeMarkerAlike(ch)
		}
		switch ch {
		case '\'':
			return s.probeText(ch)
		case '<', '>', '=':
			return s.probeContrast(ch)
		default:
			return s.unfit(ch)
		}
	}
}

//
func (s *Analyzer) Symbol() Symbol { return s.tok }

//
func (s *Analyzer) String() string { return s.buf.String() }

//
func (s *Analyzer) Pos() int { return s.pos }

//
func (s *Analyzer) Err() error { return s.err }

//
//
func (s *Analyzer) probeNumeral(initial rune) error {
	s.buf.WriteRune(initial)
	if err := s.probeDuring(equalsNumber); err != nil {
		return err
	}

	ch, err := s.character()
	if err != nil && err != io.EOF {
		return err
	}
	if ch == '.' {
		s.buf.WriteRune(ch)
		if err := s.probeDuring(equalsNumber); err != nil {
			return err
		}
	} else {
		s.uncharacter()
	}
	s.tok = TYPNumeral
	return nil
}

func (s *Analyzer) probeText(initial rune) error {
	//
	for {
		ch, err := s.character()
		if err != nil {
			return s.abort(err)
		} else if ch == initial {
			//
			s.tok = TYPText
			return nil
		}
		s.buf.WriteRune(ch)
	}
}

func (s *Analyzer) probeContrast(initial rune) error {
	s.buf.WriteRune(initial)
	switch initial {
	case '=':
		s.tok = TEq
		return nil
	case '<':
		s.tok = TLt
	case '>':
		s.tok = TGt
	default:
		return s.unfit(initial)
	}

	ch, err := s.character()
	if err == io.EOF {
		return nil //
	} else if err != nil {
		return s.abort(err)
	}
	if ch == '=' {
		s.buf.WriteRune(ch)
		s.tok++ //
		return nil
	}
	s.uncharacter()
	return nil
}

func (s *Analyzer) probeMarkerAlike(initial rune) error {
	s.buf.WriteRune(initial)
	var ownsDomain bool
	for {
		ch, err := s.character()
		if err == io.EOF {
			break
		} else if err != nil {
			return s.abort(err)
		}
		if !equalsMarkerCharacter(ch) {
			ownsDomain = ch == ' ' //
			break
		}
		s.buf.WriteRune(ch)
	}

	string := s.buf.String()
	switch string {
	case "REDACTED":
		if ownsDomain {
			return s.probeTimestamp()
		}
		s.tok = TYPMarker
	case "REDACTED":
		if ownsDomain {
			return s.probeDatamark()
		}
		s.tok = TYPMarker
	case "REDACTED":
		s.tok = TYPAlso
	case "REDACTED":
		s.tok = TYPPresent
	case "REDACTED":
		s.tok = TYPIncludes
	default:
		s.tok = TYPMarker
	}
	s.uncharacter()
	return nil
}

func (s *Analyzer) probeTimestamp() error {
	s.buf.Reset() //
	if err := s.probeDuring(equalsMomentCharacter); err != nil {
		return err
	}
	if ts, err := time.Parse(MomentLayout, s.buf.String()); err != nil {
		return s.abort(fmt.Errorf("REDACTED", err))
	} else if y := ts.Year(); y < 1900 || y > 2999 {
		return s.abort(fmt.Errorf("REDACTED", ts.Year()))
	}
	s.tok = TYPMoment
	return nil
}

func (s *Analyzer) probeDatamark() error {
	s.buf.Reset() //
	if err := s.probeDuring(equalsTimeCharacter); err != nil {
		return err
	}
	if ts, err := time.Parse(TimeLayout, s.buf.String()); err != nil {
		return s.abort(fmt.Errorf("REDACTED", err))
	} else if y := ts.Year(); y < 1900 || y > 2999 {
		return s.abort(fmt.Errorf("REDACTED", ts.Year()))
	}
	s.tok = TYPTime
	return nil
}

func (s *Analyzer) probeDuring(ok func(rune) bool) error {
	for {
		ch, err := s.character()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return s.abort(err)
		case !ok(ch):
			s.uncharacter()
			return nil
		default:
			s.buf.WriteRune(ch)
		}
	}
}

func (s *Analyzer) character() (rune, error) {
	ch, nb, err := s.r.ReadRune()
	s.final = nb
	s.end += nb
	return ch, err
}

func (s *Analyzer) uncharacter() {
	_ = s.r.UnreadRune()
	s.end -= s.final
}

func (s *Analyzer) abort(err error) error {
	s.err = err
	return err
}

func (s *Analyzer) unfit(ch rune) error {
	return s.abort(fmt.Errorf("REDACTED", ch, s.end))
}

func equalsNumber(r rune) bool { return '0' <= r && r <= '9' }

func equalsMarkerCharacter(r rune) bool {
	return r == '.' || r == '_' || r == '-' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func equalsInitialMarkerCharacter(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func equalsMomentCharacter(r rune) bool {
	return strings.ContainsRune("REDACTED", r) || equalsNumber(r)
}

func equalsTimeCharacter(r rune) bool { return equalsNumber(r) || r == '-' }
