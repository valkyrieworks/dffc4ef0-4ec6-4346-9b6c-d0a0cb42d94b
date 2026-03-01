package log

import "fmt"

type stratum byte

const (
	stratumDiagnose stratum = 1 << iota
	stratumDetails
	stratumFailure
)

type refine struct {
	following             Tracer
	permitted          stratum            //
	primarilyPermitted stratum            //
	permittedTokvals   map[tokval]stratum //
}

type tokval struct {
	key   any
	datum any
}

//
//
//
//
func FreshRefine(following Tracer, choices ...Selection) Tracer {
	l := &refine{
		following:           following,
		permittedTokvals: make(map[tokval]stratum),
	}
	for _, selection := range choices {
		selection(l)
	}
	l.primarilyPermitted = l.permitted
	return l
}

func (l *refine) Details(msg string, tokvals ...any) {
	stratumPermitted := l.permitted&stratumDetails != 0
	if !stratumPermitted {
		return
	}
	l.following.Details(msg, tokvals...)
}

func (l *refine) Diagnose(msg string, tokvals ...any) {
	if ReportDiagnose {
		stratumPermitted := l.permitted&stratumDiagnose != 0
		if !stratumPermitted {
			return
		}
		l.following.Diagnose(msg, tokvals...)
	}
}

func (l *refine) Failure(msg string, tokvals ...any) {
	stratumPermitted := l.permitted&stratumFailure != 0
	if !stratumPermitted {
		return
	}
	l.following.Failure(msg, tokvals...)
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
func (l *refine) Using(tokvals ...any) Tracer {
	tokenInsidePermittedTokvals := false

	for i := len(tokvals) - 2; i >= 0; i -= 2 {
		for kv, permitted := range l.permittedTokvals {
			if tokvals[i] == kv.key {
				tokenInsidePermittedTokvals = true
				//
				//
				//
				if tokvals[i+1] == kv.datum {
					return &refine{
						following:             l.following.Using(tokvals...),
						permitted:          permitted, //
						permittedTokvals:   l.permittedTokvals,
						primarilyPermitted: l.primarilyPermitted,
					}
				}
			}
		}
	}

	//
	//
	//
	if tokenInsidePermittedTokvals {
		return &refine{
			following:             l.following.Using(tokvals...),
			permitted:          l.primarilyPermitted, //
			permittedTokvals:   l.permittedTokvals,
			primarilyPermitted: l.primarilyPermitted,
		}
	}

	return &refine{
		following:             l.following.Using(tokvals...),
		permitted:          l.permitted, //
		permittedTokvals:   l.permittedTokvals,
		primarilyPermitted: l.primarilyPermitted,
	}
}

//

//
type Selection func(*refine)

//
//
func PermitStratum(lvl string) (Selection, error) {
	switch lvl {
	case "REDACTED":
		return PermitDiagnose(), nil
	case "REDACTED":
		return PermitDetails(), nil
	case "REDACTED":
		return PermitFailure(), nil
	case "REDACTED":
		return PermitNil(), nil
	default:
		return nil, fmt.Errorf("REDACTED", lvl)
	}
}

//
func PermitEvery() Selection {
	return PermitDiagnose()
}

//
func PermitDiagnose() Selection {
	return permitted(stratumFailure | stratumDetails | stratumDiagnose)
}

//
func PermitDetails() Selection {
	return permitted(stratumFailure | stratumDetails)
}

//
func PermitFailure() Selection {
	return permitted(stratumFailure)
}

//
func PermitNil() Selection {
	return permitted(0)
}

func permitted(permitted stratum) Selection {
	return func(l *refine) { l.permitted = permitted }
}

//
func PermitDiagnoseUsing(key any, datum any) Selection {
	return func(l *refine) { l.permittedTokvals[tokval{key, datum}] = stratumFailure | stratumDetails | stratumDiagnose }
}

//
func PermitDetailsUsing(key any, datum any) Selection {
	return func(l *refine) { l.permittedTokvals[tokval{key, datum}] = stratumFailure | stratumDetails }
}

//
func PermitFailureUsing(key any, datum any) Selection {
	return func(l *refine) { l.permittedTokvals[tokval{key, datum}] = stratumFailure }
}

//
func PermitNilUsing(key any, datum any) Selection {
	return func(l *refine) { l.permittedTokvals[tokval{key, datum}] = 0 }
}
