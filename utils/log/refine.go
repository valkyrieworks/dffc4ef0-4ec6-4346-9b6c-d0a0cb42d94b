package log

import "fmt"

type layer byte

const (
	layerDiagnose layer = 1 << iota
	layerDetails
	layerFault
)

type refine struct {
	following             Tracer
	permitted          layer            //
	primarilyPermitted layer            //
	permittedKeyvalues   map[property]layer //
}

type property struct {
	key   any
	item any
}

//
//
//
//
func NewRefine(following Tracer, options ...Setting) Tracer {
	l := &refine{
		following:           following,
		permittedKeyvalues: make(map[property]layer),
	}
	for _, setting := range options {
		setting(l)
	}
	l.primarilyPermitted = l.permitted
	return l
}

func (l *refine) Details(msg string, keyvalues ...any) {
	layerPermitted := l.permitted&layerDetails != 0
	if !layerPermitted {
		return
	}
	l.following.Details(msg, keyvalues...)
}

func (l *refine) Diagnose(msg string, keyvalues ...any) {
	if TraceDiagnose {
		layerPermitted := l.permitted&layerDiagnose != 0
		if !layerPermitted {
			return
		}
		l.following.Diagnose(msg, keyvalues...)
	}
}

func (l *refine) Fault(msg string, keyvalues ...any) {
	layerPermitted := l.permitted&layerFault != 0
	if !layerPermitted {
		return
	}
	l.following.Fault(msg, keyvalues...)
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
func (l *refine) With(keyvalues ...any) Tracer {
	keyInPermittedKeyvalues := false

	for i := len(keyvalues) - 2; i >= 0; i -= 2 {
		for kv, permitted := range l.permittedKeyvalues {
			if keyvalues[i] == kv.key {
				keyInPermittedKeyvalues = true
				//
				//
				//
				if keyvalues[i+1] == kv.item {
					return &refine{
						following:             l.following.With(keyvalues...),
						permitted:          permitted, //
						permittedKeyvalues:   l.permittedKeyvalues,
						primarilyPermitted: l.primarilyPermitted,
					}
				}
			}
		}
	}

	//
	//
	//
	if keyInPermittedKeyvalues {
		return &refine{
			following:             l.following.With(keyvalues...),
			permitted:          l.primarilyPermitted, //
			permittedKeyvalues:   l.permittedKeyvalues,
			primarilyPermitted: l.primarilyPermitted,
		}
	}

	return &refine{
		following:             l.following.With(keyvalues...),
		permitted:          l.permitted, //
		permittedKeyvalues:   l.permittedKeyvalues,
		primarilyPermitted: l.primarilyPermitted,
	}
}

//

//
type Setting func(*refine)

//
//
func PermitLayer(lvl string) (Setting, error) {
	switch lvl {
	case "REDACTED":
		return PermitDiagnose(), nil
	case "REDACTED":
		return PermitDetails(), nil
	case "REDACTED":
		return PermitFault(), nil
	case "REDACTED":
		return PermitNone(), nil
	default:
		return nil, fmt.Errorf("REDACTED", lvl)
	}
}

//
func PermitAll() Setting {
	return PermitDiagnose()
}

//
func PermitDiagnose() Setting {
	return permitted(layerFault | layerDetails | layerDiagnose)
}

//
func PermitDetails() Setting {
	return permitted(layerFault | layerDetails)
}

//
func PermitFault() Setting {
	return permitted(layerFault)
}

//
func PermitNone() Setting {
	return permitted(0)
}

func permitted(permitted layer) Setting {
	return func(l *refine) { l.permitted = permitted }
}

//
func PermitDiagnoseWith(key any, item any) Setting {
	return func(l *refine) { l.permittedKeyvalues[property{key, item}] = layerFault | layerDetails | layerDiagnose }
}

//
func PermitDetailsWith(key any, item any) Setting {
	return func(l *refine) { l.permittedKeyvalues[property{key, item}] = layerFault | layerDetails }
}

//
func PermitFaultWith(key any, item any) Setting {
	return func(l *refine) { l.permittedKeyvalues[property{key, item}] = layerFault }
}

//
func PermitNoneWith(key any, item any) Setting {
	return func(l *refine) { l.permittedKeyvalues[property{key, item}] = 0 }
}
