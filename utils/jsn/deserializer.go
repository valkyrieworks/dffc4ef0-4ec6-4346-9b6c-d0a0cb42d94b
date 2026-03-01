package jsn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

//
//
func Decode(bz []byte, v any) error {
	return deserialize(bz, v)
}

func deserialize(bz []byte, v any) error {
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
	if kindDirectory.alias(rv.Type()) != "REDACTED" {
		return deserializeMirrorContract(bz, rv)
	}

	return deserializeMirror(bz, rv)
}

func deserializeMirror(bz []byte, rv reflect.Value) error {
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
	if rv.Type() == momentKind {
		switch {
		case len(bz) < 2 || bz[0] != '"' || bz[len(bz)-1] != '"':
			return fmt.Errorf("REDACTED", bz)
		case bz[len(bz)-2] != 'Z':
			return fmt.Errorf("REDACTED", bz)
		}
	}

	//
	if rv.Addr().Type().Implements(jsnDeserializerKind) {
		return rv.Addr().Interface().(json.Unmarshaler).UnmarshalJSON(bz)
	}

	switch rv.Type().Kind() {
	//
	case reflect.Slice, reflect.Array:
		return deserializeMirrorCatalog(bz, rv)

	case reflect.Map:
		return deserializeMirrorIndex(bz, rv)

	case reflect.Struct:
		return deserializeMirrorSchema(bz, rv)

	case reflect.Interface:
		return deserializeMirrorContract(bz, rv)

	//
	case reflect.Int64, reflect.Int, reflect.Uint64, reflect.Uint:
		if bz[0] != '"' || bz[len(bz)-1] != '"' {
			return fmt.Errorf("REDACTED", string(bz))
		}
		bz = bz[1 : len(bz)-1]
		fallthrough

	//
	default:
		return deserializeStandardlib(bz, rv)
	}
}

func deserializeMirrorCatalog(bz []byte, rv reflect.Value) error {
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

		} else if err := deserializeStandardlib(bz, rv); err != nil {
			return err
		}

	//
	default:
		var crudeSection []json.RawMessage
		if err := json.Unmarshal(bz, &crudeSection); err != nil {
			return err
		}
		if rv.Type().Kind() == reflect.Slice {
			rv.Set(reflect.MakeSlice(reflect.SliceOf(rv.Type().Elem()), len(crudeSection), len(crudeSection)))
		}
		if rv.Len() != len(crudeSection) { //
			return fmt.Errorf("REDACTED", len(crudeSection), rv.Len())
		}
		for i, bz := range crudeSection {
			if err := deserializeMirror(bz, rv.Index(i)); err != nil {
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

func deserializeMirrorIndex(bz []byte, rv reflect.Value) error {
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
		datum := reflect.New(rv.Type().Elem()).Elem()
		if err := deserializeMirror(bz, datum); err != nil {
			return err
		}
		rv.SetMapIndex(reflect.ValueOf(key), datum)
	}
	return nil
}

func deserializeMirrorSchema(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() {
		return errors.New("REDACTED")
	}
	strDetails := createSchemaDetails(rv.Type())

	//
	crudeIndex := make(map[string]json.RawMessage)
	if err := json.Unmarshal(bz, &crudeIndex); err != nil {
		return err
	}
	for i, funcDetails := range strDetails.areas {
		if !funcDetails.concealed {
			frv := rv.Field(i)
			bz := crudeIndex[funcDetails.jsnAlias]
			if len(bz) > 0 {
				if err := deserializeMirror(bz, frv); err != nil {
					return err
				}
			} else if !funcDetails.excludeBlank {
				frv.Set(reflect.Zero(frv.Type()))
			}
		}
	}

	return nil
}

func deserializeMirrorContract(bz []byte, rv reflect.Value) error {
	if !rv.CanAddr() {
		return errors.New("REDACTED")
	}

	//
	encapsulator := contractEncapsulator{}
	if err := json.Unmarshal(bz, &encapsulator); err != nil {
		return err
	}
	if encapsulator.Kind == "REDACTED" {
		return errors.New("REDACTED")
	}
	if len(encapsulator.Datum) == 0 {
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
	rt, yieldReference := kindDirectory.search(encapsulator.Kind)
	if rt == nil {
		return fmt.Errorf("REDACTED", encapsulator.Kind)
	}

	creference := reflect.New(rt)
	crv := creference.Elem()
	if err := deserializeMirror(encapsulator.Datum, crv); err != nil {
		return err
	}

	//
	//
	//
	if rv.Type().Kind() == reflect.Interface && yieldReference {
		if !creference.Type().AssignableTo(rv.Type()) {
			return fmt.Errorf("REDACTED", encapsulator.Kind)
		}
		rv.Set(creference)
	} else {
		if !crv.Type().AssignableTo(rv.Type()) {
			return fmt.Errorf("REDACTED", encapsulator.Kind)
		}
		rv.Set(crv)
	}
	return nil
}

func deserializeStandardlib(bz []byte, rv reflect.Value) error {
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

type contractEncapsulator struct {
	Kind  string          `json:"kind"`
	Datum json.RawMessage `json:"datum"`
}
