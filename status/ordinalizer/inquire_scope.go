package ordinalizer

import (
	"math/big"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
type InquireExtents map[string]InquireScope

//
type InquireScope struct {
	LesserRestricted        any //
	HigherRestricted        any //
	Key               string
	EncompassLesserRestricted bool
	EncompassHigherRestricted bool
}

//
func (qr InquireScope) SomeRestricted() any {
	if qr.LesserRestricted != nil {
		return qr.LesserRestricted
	}

	return qr.HigherRestricted
}

//
//
func (qr InquireScope) LesserRestrictedDatum() any {
	if qr.LesserRestricted == nil {
		return nil
	}

	if qr.EncompassLesserRestricted {
		return qr.LesserRestricted
	}

	switch t := qr.LesserRestricted.(type) {
	case int64:
		return t + 1
	case *big.Int:
		tmp := new(big.Int)
		return tmp.Add(t, big.NewInt(1))

	case *big.Float:
		//
		//
		//
		//
		//
		return t
	case time.Time:
		return t.Unix() + 1

	default:
		panic("REDACTED")
	}
}

//
//
func (qr InquireScope) HigherRestrictedDatum() any {
	if qr.HigherRestricted == nil {
		return nil
	}

	if qr.EncompassHigherRestricted {
		return qr.HigherRestricted
	}

	switch t := qr.HigherRestricted.(type) {
	case int64:
		return t - 1
	case *big.Int:
		tmp := new(big.Int)
		return tmp.Sub(t, big.NewInt(1))
	case *big.Float:
		return t
	case time.Time:
		return t.Unix() - 1

	default:
		panic("REDACTED")
	}
}

//
//
func ScanForeachExtentsUsingAltitude(terms []grammar.Stipulation) (inquireScope InquireExtents, indices []int, altitudeScope InquireScope) {
	inquireScope = make(InquireExtents)
	for i, c := range terms {
		if EqualsScopeAction(c.Op) {
			altitudeToken := c.Tag == kinds.LedgerAltitudeToken || c.Tag == kinds.TransferAltitudeToken
			r, ok := inquireScope[c.Tag]
			if !ok {
				r = InquireScope{Key: c.Tag}
				if c.Tag == kinds.LedgerAltitudeToken || c.Tag == kinds.TransferAltitudeToken {
					altitudeScope = InquireScope{Key: c.Tag}
				}
			}

			switch c.Op {
			case grammar.TGt:
				if altitudeToken {
					altitudeScope.LesserRestricted = stipulationParam(c)
				}
				r.LesserRestricted = stipulationParam(c)

			case grammar.TYPGreatereq:
				r.EncompassLesserRestricted = true
				r.LesserRestricted = stipulationParam(c)
				if altitudeToken {
					altitudeScope.EncompassLesserRestricted = true
					altitudeScope.LesserRestricted = stipulationParam(c)
				}

			case grammar.TLt:
				r.HigherRestricted = stipulationParam(c)
				if altitudeToken {
					altitudeScope.HigherRestricted = stipulationParam(c)
				}

			case grammar.TYPLesseq:
				r.EncompassHigherRestricted = true
				r.HigherRestricted = stipulationParam(c)
				if altitudeToken {
					altitudeScope.EncompassHigherRestricted = true
					altitudeScope.HigherRestricted = stipulationParam(c)
				}
			}

			inquireScope[c.Tag] = r
			indices = append(indices, i)
		}
	}

	return inquireScope, indices, altitudeScope
}

//
func ScanForeachExtents(terms []grammar.Stipulation) (extents InquireExtents, indices []int) {
	extents = make(InquireExtents)
	for i, c := range terms {
		if EqualsScopeAction(c.Op) {
			r, ok := extents[c.Tag]
			if !ok {
				r = InquireScope{Key: c.Tag}
			}

			switch c.Op {
			case grammar.TGt:
				r.LesserRestricted = stipulationParam(c)

			case grammar.TYPGreatereq:
				r.EncompassLesserRestricted = true
				r.LesserRestricted = stipulationParam(c)

			case grammar.TLt:
				r.HigherRestricted = stipulationParam(c)

			case grammar.TYPLesseq:
				r.EncompassHigherRestricted = true
				r.HigherRestricted = stipulationParam(c)
			}

			extents[c.Tag] = r
			indices = append(indices, i)
		}
	}

	return extents, indices
}

//
//
func EqualsScopeAction(op grammar.Symbol) bool {
	switch op {
	case grammar.TGt, grammar.TYPGreatereq, grammar.TLt, grammar.TYPLesseq:
		return true

	default:
		return false
	}
}

func stipulationParam(c grammar.Stipulation) any {
	if c.Arg == nil {
		return nil
	}
	switch c.Arg.Kind {
	case grammar.TYPNumeral:
		return c.Arg.Numeral()
	case grammar.TYPMoment, grammar.TYPTime:
		return c.Arg.Moment()
	default:
		return c.Arg.Datum() //
	}
}
