package host

import (
	"fmt"

	"github.com/valkyrieworks/iface/kinds"
)

//
type ErrUnclearHostKind struct {
	HostKind string
}

func (e ErrUnclearHostKind) Fault() string {
	return fmt.Sprintf("REDACTED", e.HostKind)
}

//
type ErrLinkageDoesNotOccur struct {
	LinkUID int
}

func (e ErrLinkageDoesNotOccur) Fault() string {
	return fmt.Sprintf("REDACTED", e.LinkUID)
}

type ErrUnclearQuery struct {
	Query kinds.Query
}

func (e ErrUnclearQuery) Fault() string {
	return fmt.Sprintf("REDACTED", e.Query)
}
