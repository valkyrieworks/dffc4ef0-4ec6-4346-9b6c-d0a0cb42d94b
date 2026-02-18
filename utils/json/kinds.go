package json

import (
	"errors"
	"fmt"
	"reflect"

	engineconnect "github.com/valkyrieworks/utils/align"
)

//
var kindRepository = newKinds()

//
//
//
//
//
//
//
func EnrollKind(_kind any, label string) {
	if _kind == nil {
		panic("REDACTED")
	}
	err := kindRepository.enroll(label, reflect.ValueOf(_kind).Type())
	if err != nil {
		panic(err)
	}
}

//
type kindDetails struct {
	label      string
	rt        reflect.Type
	yieldPointer bool
}

//
type kinds struct {
	engineconnect.ReadwriteLock
	byKind map[reflect.Type]*kindDetails
	byLabel map[string]*kindDetails
}

//
func newKinds() kinds {
	return kinds{
		byKind: map[reflect.Type]*kindDetails{},
		byLabel: map[string]*kindDetails{},
	}
}

//
func (t *kinds) enroll(label string, rt reflect.Type) error {
	if label == "REDACTED" {
		return errors.New("REDACTED")
	}
	//
	//
	yieldPointer := false
	for rt.Kind() == reflect.Ptr {
		yieldPointer = true
		rt = rt.Elem()
	}
	tDetails := &kindDetails{
		label:      label,
		rt:        rt,
		yieldPointer: yieldPointer,
	}

	t.Lock()
	defer t.Unlock()
	if _, ok := t.byLabel[tDetails.label]; ok {
		return fmt.Errorf("REDACTED", label)
	}
	if _, ok := t.byKind[tDetails.rt]; ok {
		return fmt.Errorf("REDACTED", rt)
	}
	t.byLabel[label] = tDetails
	t.byKind[rt] = tDetails
	return nil
}

//
func (t *kinds) search(label string) (reflect.Type, bool) {
	t.RLock()
	defer t.RUnlock()
	tDetails := t.byLabel[label]
	if tDetails == nil {
		return nil, false
	}
	return tDetails.rt, tDetails.yieldPointer
}

//
func (t *kinds) label(rt reflect.Type) string {
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	t.RLock()
	defer t.RUnlock()
	tDetails := t.byKind[rt]
	if tDetails == nil {
		return "REDACTED"
	}
	return tDetails.label
}
