package json__test

import (
	"time"

	"github.com/valkyrieworks/utils/json"
)

//
func init() {
	json.EnrollKind(&Car{}, "REDACTED")
	json.EnrollKind(Vessel{}, "REDACTED")
	json.EnrollKind(PublicKey{}, "REDACTED")
	json.EnrollKind(InternalKey{}, "REDACTED")
}

type Automobile interface {
	Propel() error
}

//
type Car struct {
	Revolutions int32
}

func (c *Car) Propel() error { return nil }

//
type Vessel struct {
	Navigate bool
}

func (b Vessel) Propel() error { return nil }

//
type (
	PublicKey  [8]byte
	InternalKey [8]byte
)

//
type BespokePointer struct {
	Item string
}

func (c *BespokePointer) SerializeJSON() ([]byte, error) {
	return []byte("REDACTED"), nil
}

func (c *BespokePointer) UnserializeJSON(_ []byte) error {
	c.Item = "REDACTED"
	return nil
}

//
//
type BespokeItem struct {
	Item string
}

func (c BespokeItem) SerializeJSON() ([]byte, error) {
	return []byte("REDACTED"), nil
}

func (c BespokeItem) UnserializeJSON(_ []byte) error {
	return nil
}

//
type Markers struct {
	JSONLabel  string `json:"label"`
	IgnoreEmpty string `json:",omitempty"`
	Concealed    string `json:"-"`
	Markers      *Markers  `json:"labels,omitempty"`
}

//
type Struct struct {
	Bool         bool
	Float64      float64
	Int32        int32
	Int64        int64
	Int64pointer     *int64
	String       string
	StringPointerPointer **string
	Octets        []byte
	Time         time.Time
	Car          *Car
	Vessel         Vessel
	Automobiles     []Automobile
	Offspring        *Struct
	internal      string
}
