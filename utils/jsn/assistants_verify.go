package jsn__test

import (
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

//
func initialize() {
	jsn.EnrollKind(&Car{}, "REDACTED")
	jsn.EnrollKind(Vessel{}, "REDACTED")
	jsn.EnrollKind(CommonToken{}, "REDACTED")
	jsn.EnrollKind(SecludedToken{}, "REDACTED")
}

type Carriage interface {
	Operate() error
}

//
type Car struct {
	Rotations int32
}

func (c *Car) Operate() error { return nil }

//
type Vessel struct {
	Navigate bool
}

func (b Vessel) Operate() error { return nil }

//
type (
	CommonToken  [8]byte
	SecludedToken [8]byte
)

//
type BespokeReference struct {
	Datum string
}

func (c *BespokeReference) SerializeJSN() ([]byte, error) {
	return []byte("REDACTED"), nil
}

func (c *BespokeReference) DecodeJSN(_ []byte) error {
	c.Datum = "REDACTED"
	return nil
}

//
//
type BespokeDatum struct {
	Datum string
}

func (c BespokeDatum) SerializeJSN() ([]byte, error) {
	return []byte("REDACTED"), nil
}

func (c BespokeDatum) DecodeJSN(_ []byte) error {
	return nil
}

//
type Labels struct {
	JSNAlias  string `json:"alias"`
	ExcludeBlank string `json:",omitempty"`
	Concealed    string `json:"-"`
	Labels      *Labels  `json:"labels,omitempty"`
}

//
type Schema struct {
	Flag         bool
	Float64      float64
	Integer32        int32
	Int64n        int64
	Int64reference     *int64
	Text       string
	TextReferenceReference **string
	Octets        []byte
	Moment         time.Time
	Car          *Car
	Vessel         Vessel
	Carriages     []Carriage
	Offspring        *Schema
	secluded      string
}
