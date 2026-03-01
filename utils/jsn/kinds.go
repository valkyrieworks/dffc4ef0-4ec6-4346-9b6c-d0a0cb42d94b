package jsn

import (
	"errors"
	"fmt"
	"reflect"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
var kindDirectory = freshKinds()

//
//
//
//
//
//
//
func EnrollKind(_kind any, alias string) {
	if _kind == nil {
		panic("REDACTED")
	}
	err := kindDirectory.enroll(alias, reflect.ValueOf(_kind).Type())
	if err != nil {
		panic(err)
	}
}

//
type kindDetails struct {
	alias      string
	rt        reflect.Type
	yieldReference bool
}

//
type kinds struct {
	commitchronize.ReadwriteExclusion
	viaKind map[reflect.Type]*kindDetails
	viaAlias map[string]*kindDetails
}

//
func freshKinds() kinds {
	return kinds{
		viaKind: map[reflect.Type]*kindDetails{},
		viaAlias: map[string]*kindDetails{},
	}
}

//
func (t *kinds) enroll(alias string, rt reflect.Type) error {
	if alias == "REDACTED" {
		return errors.New("REDACTED")
	}
	//
	//
	yieldReference := false
	for rt.Kind() == reflect.Ptr {
		yieldReference = true
		rt = rt.Elem()
	}
	typDetails := &kindDetails{
		alias:      alias,
		rt:        rt,
		yieldReference: yieldReference,
	}

	t.Lock()
	defer t.Unlock()
	if _, ok := t.viaAlias[typDetails.alias]; ok {
		return fmt.Errorf("REDACTED", alias)
	}
	if _, ok := t.viaKind[typDetails.rt]; ok {
		return fmt.Errorf("REDACTED", rt)
	}
	t.viaAlias[alias] = typDetails
	t.viaKind[rt] = typDetails
	return nil
}

//
func (t *kinds) search(alias string) (reflect.Type, bool) {
	t.RLock()
	defer t.RUnlock()
	typDetails := t.viaAlias[alias]
	if typDetails == nil {
		return nil, false
	}
	return typDetails.rt, typDetails.yieldReference
}

//
func (t *kinds) alias(rt reflect.Type) string {
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	t.RLock()
	defer t.RUnlock()
	typDetails := t.viaKind[rt]
	if typDetails == nil {
		return "REDACTED"
	}
	return typDetails.alias
}
