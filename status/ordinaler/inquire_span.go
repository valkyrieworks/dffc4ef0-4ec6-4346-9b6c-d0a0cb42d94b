package ordinaler

import (
	"math/big"
	"time"

	"github.com/valkyrieworks/utils/broadcast/inquire/grammar"
	"github.com/valkyrieworks/kinds"
)

//
//
//
type InquireSpans map[string]InquireSpan

//
type InquireSpan struct {
	LesserLimited        any //
	UpperLimited        any //
	Key               string
	EncompassLesserLimited bool
	EncompassUpperLimited bool
}

//
func (qr InquireSpan) AnyLimited() any {
	if qr.LesserLimited != nil {
		return qr.LesserLimited
	}

	return qr.UpperLimited
}

//
//
func (qr InquireSpan) LesserLimitedItem() any {
	if qr.LesserLimited == nil {
		return nil
	}

	if qr.EncompassLesserLimited {
		return qr.LesserLimited
	}

	switch t := qr.LesserLimited.(type) {
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
func (qr InquireSpan) UpperLimitedItem() any {
	if qr.UpperLimited == nil {
		return nil
	}

	if qr.EncompassUpperLimited {
		return qr.UpperLimited
	}

	switch t := qr.UpperLimited.(type) {
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
func ScanForSpansWithLevel(states []grammar.Status) (inquireSpan InquireSpans, listings []int, levelSpan InquireSpan) {
	inquireSpan = make(InquireSpans)
	for i, c := range states {
		if IsSpanProcess(c.Op) {
			levelKey := c.Tag == kinds.LedgerLevelKey || c.Tag == kinds.TransferLevelKey
			r, ok := inquireSpan[c.Tag]
			if !ok {
				r = InquireSpan{Key: c.Tag}
				if c.Tag == kinds.LedgerLevelKey || c.Tag == kinds.TransferLevelKey {
					levelSpan = InquireSpan{Key: c.Tag}
				}
			}

			switch c.Op {
			case grammar.TGt:
				if levelKey {
					levelSpan.LesserLimited = statusArgument(c)
				}
				r.LesserLimited = statusArgument(c)

			case grammar.TGeq:
				r.EncompassLesserLimited = true
				r.LesserLimited = statusArgument(c)
				if levelKey {
					levelSpan.EncompassLesserLimited = true
					levelSpan.LesserLimited = statusArgument(c)
				}

			case grammar.TLt:
				r.UpperLimited = statusArgument(c)
				if levelKey {
					levelSpan.UpperLimited = statusArgument(c)
				}

			case grammar.TLeq:
				r.EncompassUpperLimited = true
				r.UpperLimited = statusArgument(c)
				if levelKey {
					levelSpan.EncompassUpperLimited = true
					levelSpan.UpperLimited = statusArgument(c)
				}
			}

			inquireSpan[c.Tag] = r
			listings = append(listings, i)
		}
	}

	return inquireSpan, listings, levelSpan
}

//
func ScanForSpans(states []grammar.Status) (spans InquireSpans, listings []int) {
	spans = make(InquireSpans)
	for i, c := range states {
		if IsSpanProcess(c.Op) {
			r, ok := spans[c.Tag]
			if !ok {
				r = InquireSpan{Key: c.Tag}
			}

			switch c.Op {
			case grammar.TGt:
				r.LesserLimited = statusArgument(c)

			case grammar.TGeq:
				r.EncompassLesserLimited = true
				r.LesserLimited = statusArgument(c)

			case grammar.TLt:
				r.UpperLimited = statusArgument(c)

			case grammar.TLeq:
				r.EncompassUpperLimited = true
				r.UpperLimited = statusArgument(c)
			}

			spans[c.Tag] = r
			listings = append(listings, i)
		}
	}

	return spans, listings
}

//
//
func IsSpanProcess(op grammar.Symbol) bool {
	switch op {
	case grammar.TGt, grammar.TGeq, grammar.TLt, grammar.TLeq:
		return true

	default:
		return false
	}
}

func statusArgument(c grammar.Status) any {
	if c.Arg == nil {
		return nil
	}
	switch c.Arg.Kind {
	case grammar.TAmount:
		return c.Arg.Amount()
	case grammar.TTime, grammar.TDate:
		return c.Arg.Time()
	default:
		return c.Arg.Item() //
	}
}
