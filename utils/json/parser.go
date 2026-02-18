package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

//
//
func Unserialize(bz []byte, v any) error {
	return parse(bz, v)
}

func parse(bz []byte, v any) error {
	if len(bz) == 0 {
		return errors.New("REDACTED")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return errors.New("REDACTED")
	}
	rv = rv.Elem()

	//
	//
	//
	//
	if kindRepository.label(rv.Type()) != "REDACTED" {
		return parseMirrorInterface(bz, rv)
	}

	return parseMirror(bz, rv)
}

func parseMirror(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() {
		return errors.New("REDACTED")
	}

	//
	if bytes.Equal(bz, []byte("REDACTED")) {
		rv.Set(reflect.Zero(rv.Type()))
		return nil
	}

	//
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}

	//
	if rv.Type() == timeKind {
		switch {
		case len(bz) < 2 || bz[0] != '"' || bz[len(bz)-1] != '"':
			return fmt.Errorf("REDACTED", bz)
		case bz[len(bz)-2] != 'Z':
			return fmt.Errorf("REDACTED", bz)
		}
	}

	//
	if rv.Addr().Type().Implements(jsonUnserializerKind) {
		return rv.Addr().Interface().(json.Unmarshaler).UnmarshalJSON(bz)
	}

	switch rv.Type().Kind() {
	//
	case reflect.Slice, reflect.Array:
		return parseMirrorCatalog(bz, rv)

	case reflect.Map:
		return parseMirrorIndex(bz, rv)

	case reflect.Struct:
		return parseMirrorStruct(bz, rv)

	case reflect.Interface:
		return parseMirrorInterface(bz, rv)

	//
	case reflect.Int64, reflect.Int, reflect.Uint64, reflect.Uint:
		if bz[0] != '"' || bz[len(bz)-1] != '"' {
			return fmt.Errorf("REDACTED", string(bz))
		}
		bz = bz[1 : len(bz)-1]
		fallthrough

	//
	default:
		return parseStdlib(bz, rv)
	}
}

func parseMirrorCatalog(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() {
		return errors.New("REDACTED")
	}

	switch rv.Type().Elem().Kind() {
	//
	case reflect.Uint8:
		if rv.Type().Kind() == reflect.Array {
			var buf []byte
			if err := json.Unmarshal(bz, &buf); err != nil {
				return err
			}
			if len(buf) != rv.Len() {
				return fmt.Errorf("REDACTED", len(buf), rv.Len())
			}
			reflect.Copy(rv, reflect.ValueOf(buf))

		} else if err := parseStdlib(bz, rv); err != nil {
			return err
		}

	//
	default:
		var crudeSegment []json.RawMessage
		if err := json.Unmarshal(bz, &crudeSegment); err != nil {
			return err
		}
		if rv.Type().Kind() == reflect.Slice {
			rv.Set(reflect.MakeSlice(reflect.SliceOf(rv.Type().Elem()), len(crudeSegment), len(crudeSegment)))
		}
		if rv.Len() != len(crudeSegment) { //
			return fmt.Errorf("REDACTED", len(crudeSegment), rv.Len())
		}
		for i, bz := range crudeSegment {
			if err := parseMirror(bz, rv.Index(i)); err != nil {
				return err
			}
		}
	}

	//
	if rv.Type().Kind() == reflect.Slice && rv.Len() == 0 {
		rv.Set(reflect.Zero(rv.Type()))
	}

	return nil
}

func parseMirrorIndex(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() {
		return errors.New("REDACTED")
	}

	//
	crudeIndex := make(map[string]json.RawMessage)
	if err := json.Unmarshal(bz, &crudeIndex); err != nil {
		return err
	}
	if rv.Type().Key().Kind() != reflect.String {
		return fmt.Errorf("REDACTED", rv.Type().Key().String())
	}

	//
	rv.Set(reflect.MakeMapWithSize(rv.Type(), len(crudeIndex)))
	for key, bz := range crudeIndex {
		item := reflect.New(rv.Type().Elem()).Elem()
		if err := parseMirror(bz, item); err != nil {
			return err
		}
		rv.SetMapIndex(reflect.ValueOf(key), item)
	}
	return nil
}

func parseMirrorStruct(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() {
		return errors.New("REDACTED")
	}
	sDetails := createStructDetails(rv.Type())

	//
	crudeIndex := make(map[string]json.RawMessage)
	if err := json.Unmarshal(bz, &crudeIndex); err != nil {
		return err
	}
	for i, fDetails := range sDetails.attributes {
		if !fDetails.concealed {
			frv := rv.Field(i)
			bz := crudeIndex[fDetails.jsonLabel]
			if len(bz) > 0 {
				if err := parseMirror(bz, frv); err != nil {
					return err
				}
			} else if !fDetails.ignoreEmpty {
				frv.Set(reflect.Zero(frv.Type()))
			}
		}
	}

	return nil
}

func parseMirrorInterface(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() {
		return errors.New("REDACTED")
	}

	//
	adapter := interfaceAdapter{}
	if err := json.Unmarshal(bz, &adapter); err != nil {
		return err
	}
	if adapter.Kind == "REDACTED" {
		return errors.New("REDACTED")
	}
	if len(adapter.Item) == 0 {
		return errors.New("REDACTED")
	}

	//
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}

	//
	rt, yieldPointer := kindRepository.search(adapter.Kind)
	if rt == nil {
		return fmt.Errorf("REDACTED", adapter.Kind)
	}

	cpointer := reflect.New(rt)
	crv := cpointer.Elem()
	if err := parseMirror(adapter.Item, crv); err != nil {
		return err
	}

	//
	//
	//
	if rv.Type().Kind() == reflect.Interface && yieldPointer {
		if !cpointer.Type().AssignableTo(rv.Type()) {
			return fmt.Errorf("REDACTED", adapter.Kind)
		}
		rv.Set(cpointer)
	} else {
		if !crv.Type().AssignableTo(rv.Type()) {
			return fmt.Errorf("REDACTED", adapter.Kind)
		}
		rv.Set(crv)
	}
	return nil
}

func parseStdlib(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() && rv.Kind() != reflect.Ptr {
		return errors.New("REDACTED")
	}

	//
	objective := rv
	if rv.Kind() != reflect.Ptr {
		objective = reflect.New(rv.Type())
	}
	if err := json.Unmarshal(bz, objective.Interface()); err != nil {
		return err
	}
	rv.Set(objective.Elem())
	return nil
}

type interfaceAdapter struct {
	Kind  string          `json:"kind"`
	Item json.RawMessage `json:"item"`
}
