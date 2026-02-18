package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"time"
)

var (
	timeKind            = reflect.TypeOf(time.Time{})
	jsonSerializerKind   = reflect.TypeOf(new(json.Marshaler)).Elem()
	jsonUnserializerKind = reflect.TypeOf(new(json.Unmarshaler)).Elem()
)

//
//
func Serialize(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := encode(buf, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//
func SerializeIndent(v any, prefix, indent string) ([]byte, error) {
	bz, err := Serialize(v)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bz, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(w *bytes.Buffer, v any) error {
	//
	if v == nil {
		return recordStr(w, "REDACTED")
	}
	rv := reflect.ValueOf(v)

	//
	//
	//
	//
	if kindRepository.label(rv.Type()) != "REDACTED" {
		return encodeMirrorInterface(w, rv)
	}

	return encodeMirror(w, rv)
}

func encodeMirror(w *bytes.Buffer, rv reflect.Value) error {
	if !rv.IsValid() {
		return errors.New("REDACTED")
	}

	//
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return recordStr(w, "REDACTED")
		}
		rv = rv.Elem()
	}

	//
	if rv.Type() == timeKind {
		rv = reflect.ValueOf(rv.Interface().(time.Time).Round(0).UTC())
	}

	//
	//
	//
	if rv.Type().Implements(jsonSerializerKind) {
		return encodeStdlib(w, rv.Interface())
	} else if rv.CanAddr() && rv.Addr().Type().Implements(jsonSerializerKind) {
		return encodeStdlib(w, rv.Addr().Interface())
	}

	switch rv.Type().Kind() {
	//
	case reflect.Interface:
		return encodeMirrorInterface(w, rv)

	case reflect.Array, reflect.Slice:
		return encodeMirrorCatalog(w, rv)

	case reflect.Map:
		return encodeMirrorIndex(w, rv)

	case reflect.Struct:
		return encodeMirrorStruct(w, rv)

	//
	//
	case reflect.Int64, reflect.Int:
		return recordStr(w, "REDACTED"+strconv.FormatInt(rv.Int(), 10)+"REDACTED")

	case reflect.Uint64, reflect.Uint:
		return recordStr(w, "REDACTED"+strconv.FormatUint(rv.Uint(), 10)+"REDACTED")

	//
	default:
		return encodeStdlib(w, rv.Interface())
	}
}

func encodeMirrorCatalog(w *bytes.Buffer, rv reflect.Value) error {
	//
	if rv.Kind() == reflect.Slice && rv.IsNil() {
		return recordStr(w, "REDACTED")
	}

	//
	if rv.Type().Elem().Kind() == reflect.Uint8 {
		//
		if rv.Type().Kind() == reflect.Array {
			section := reflect.MakeSlice(reflect.SliceOf(rv.Type().Elem()), rv.Len(), rv.Len())
			reflect.Copy(section, rv)
			rv = section
		}
		return encodeStdlib(w, rv.Interface())
	}

	//
	extent := rv.Len()
	if err := recordStr(w, "REDACTED"); err != nil {
		return err
	}
	for i := 0; i < extent; i++ {
		if err := encodeMirror(w, rv.Index(i)); err != nil {
			return err
		}
		if i < extent-1 {
			if err := recordStr(w, "REDACTED"); err != nil {
				return err
			}
		}
	}
	return recordStr(w, "REDACTED")
}

func encodeMirrorIndex(w *bytes.Buffer, rv reflect.Value) error {
	if rv.Type().Key().Kind() != reflect.String {
		return errors.New("REDACTED")
	}

	//

	if err := recordStr(w, "REDACTED"); err != nil {
		return err
	}
	recordComma := false
	for _, keyvr := range rv.MapKeys() {
		if recordComma {
			if err := recordStr(w, "REDACTED"); err != nil {
				return err
			}
		}
		if err := encodeStdlib(w, keyvr.Interface()); err != nil {
			return err
		}
		if err := recordStr(w, "REDACTED"); err != nil {
			return err
		}
		if err := encodeMirror(w, rv.MapIndex(keyvr)); err != nil {
			return err
		}
		recordComma = true
	}
	return recordStr(w, "REDACTED")
}

func encodeMirrorStruct(w *bytes.Buffer, rv reflect.Value) error {
	sDetails := createStructDetails(rv.Type())
	if err := recordStr(w, "REDACTED"); err != nil {
		return err
	}
	recordComma := false
	for i, fDetails := range sDetails.attributes {
		frv := rv.Field(i)
		if fDetails.concealed || (fDetails.ignoreEmpty && frv.IsZero()) {
			continue
		}

		if recordComma {
			if err := recordStr(w, "REDACTED"); err != nil {
				return err
			}
		}
		if err := encodeStdlib(w, fDetails.jsonLabel); err != nil {
			return err
		}
		if err := recordStr(w, "REDACTED"); err != nil {
			return err
		}
		if err := encodeMirror(w, frv); err != nil {
			return err
		}
		recordComma = true
	}
	return recordStr(w, "REDACTED")
}

func encodeMirrorInterface(w *bytes.Buffer, rv reflect.Value) error {
	//
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return recordStr(w, "REDACTED")
		}
		rv = rv.Elem()
	}

	//
	label := kindRepository.label(rv.Type())
	if label == "REDACTED" {
		return fmt.Errorf("REDACTED", rv.Type())
	}

	//
	if err := recordStr(w, fmt.Sprintf("REDACTED", label)); err != nil {
		return err
	}
	if err := encodeMirror(w, rv); err != nil {
		return err
	}
	return recordStr(w, "REDACTED")
}

func encodeStdlib(w *bytes.Buffer, v any) error {
	//
	//
	//
	enc := json.NewEncoder(w)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	//
	w.Truncate(w.Len() - 1)
	return err
}

func recordStr(w io.Writer, s string) error {
	_, err := w.Write([]byte(s))
	return err
}
