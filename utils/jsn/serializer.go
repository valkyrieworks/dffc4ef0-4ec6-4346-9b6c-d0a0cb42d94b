package jsn

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
	momentKind            = reflect.TypeOf(time.Time{})
	jsnSerializerKind   = reflect.TypeOf(new(json.Marshaler)).Elem()
	jsnDeserializerKind = reflect.TypeOf(new(json.Unmarshaler)).Elem()
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
func SerializeRecess(v any, heading, format string) ([]byte, error) {
	bz, err := Serialize(v)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bz, heading, format)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(w *bytes.Buffer, v any) error {
	//
	if v == nil {
		return persistTxt(w, "REDACTED")
	}
	rv := reflect.ValueOf(v)

	//
	//
	//
	//
	if kindDirectory.alias(rv.Type()) != "REDACTED" {
		return encodeMirrorContract(w, rv)
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
			return persistTxt(w, "REDACTED")
		}
		rv = rv.Elem()
	}

	//
	if rv.Type() == momentKind {
		rv = reflect.ValueOf(rv.Interface().(time.Time).Round(0).UTC())
	}

	//
	//
	//
	if rv.Type().Implements(jsnSerializerKind) {
		return encodeStandardlib(w, rv.Interface())
	} else if rv.CanAddr() && rv.Addr().Type().Implements(jsnSerializerKind) {
		return encodeStandardlib(w, rv.Addr().Interface())
	}

	switch rv.Type().Kind() {
	//
	case reflect.Interface:
		return encodeMirrorContract(w, rv)

	case reflect.Array, reflect.Slice:
		return encodeMirrorCatalog(w, rv)

	case reflect.Map:
		return encodeMirrorIndex(w, rv)

	case reflect.Struct:
		return encodeMirrorSchema(w, rv)

	//
	//
	case reflect.Int64, reflect.Int:
		return persistTxt(w, "REDACTED"+strconv.FormatInt(rv.Int(), 10)+"REDACTED")

	case reflect.Uint64, reflect.Uint:
		return persistTxt(w, "REDACTED"+strconv.FormatUint(rv.Uint(), 10)+"REDACTED")

	//
	default:
		return encodeStandardlib(w, rv.Interface())
	}
}

func encodeMirrorCatalog(w *bytes.Buffer, rv reflect.Value) error {
	//
	if rv.Kind() == reflect.Slice && rv.IsNil() {
		return persistTxt(w, "REDACTED")
	}

	//
	if rv.Type().Elem().Kind() == reflect.Uint8 {
		//
		if rv.Type().Kind() == reflect.Array {
			section := reflect.MakeSlice(reflect.SliceOf(rv.Type().Elem()), rv.Len(), rv.Len())
			reflect.Copy(section, rv)
			rv = section
		}
		return encodeStandardlib(w, rv.Interface())
	}

	//
	magnitude := rv.Len()
	if err := persistTxt(w, "REDACTED"); err != nil {
		return err
	}
	for i := 0; i < magnitude; i++ {
		if err := encodeMirror(w, rv.Index(i)); err != nil {
			return err
		}
		if i < magnitude-1 {
			if err := persistTxt(w, "REDACTED"); err != nil {
				return err
			}
		}
	}
	return persistTxt(w, "REDACTED")
}

func encodeMirrorIndex(w *bytes.Buffer, rv reflect.Value) error {
	if rv.Type().Key().Kind() != reflect.String {
		return errors.New("REDACTED")
	}

	//

	if err := persistTxt(w, "REDACTED"); err != nil {
		return err
	}
	persistSeparator := false
	for _, tokrv := range rv.MapKeys() {
		if persistSeparator {
			if err := persistTxt(w, "REDACTED"); err != nil {
				return err
			}
		}
		if err := encodeStandardlib(w, tokrv.Interface()); err != nil {
			return err
		}
		if err := persistTxt(w, "REDACTED"); err != nil {
			return err
		}
		if err := encodeMirror(w, rv.MapIndex(tokrv)); err != nil {
			return err
		}
		persistSeparator = true
	}
	return persistTxt(w, "REDACTED")
}

func encodeMirrorSchema(w *bytes.Buffer, rv reflect.Value) error {
	strDetails := createSchemaDetails(rv.Type())
	if err := persistTxt(w, "REDACTED"); err != nil {
		return err
	}
	persistSeparator := false
	for i, funcDetails := range strDetails.areas {
		frv := rv.Field(i)
		if funcDetails.concealed || (funcDetails.excludeBlank && frv.IsZero()) {
			continue
		}

		if persistSeparator {
			if err := persistTxt(w, "REDACTED"); err != nil {
				return err
			}
		}
		if err := encodeStandardlib(w, funcDetails.jsnAlias); err != nil {
			return err
		}
		if err := persistTxt(w, "REDACTED"); err != nil {
			return err
		}
		if err := encodeMirror(w, frv); err != nil {
			return err
		}
		persistSeparator = true
	}
	return persistTxt(w, "REDACTED")
}

func encodeMirrorContract(w *bytes.Buffer, rv reflect.Value) error {
	//
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return persistTxt(w, "REDACTED")
		}
		rv = rv.Elem()
	}

	//
	alias := kindDirectory.alias(rv.Type())
	if alias == "REDACTED" {
		return fmt.Errorf("REDACTED", rv.Type())
	}

	//
	if err := persistTxt(w, fmt.Sprintf("REDACTED", alias)); err != nil {
		return err
	}
	if err := encodeMirror(w, rv); err != nil {
		return err
	}
	return persistTxt(w, "REDACTED")
}

func encodeStandardlib(w *bytes.Buffer, v any) error {
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

func persistTxt(w io.Writer, s string) error {
	_, err := w.Write([]byte(s))
	return err
}
