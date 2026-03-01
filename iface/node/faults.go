package node

import (
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

//
type FaultUnfamiliarDaemonKind struct {
	DaemonKind string
}

func (e FaultUnfamiliarDaemonKind) Failure() string {
	return fmt.Sprintf("REDACTED", e.DaemonKind)
}

//
type FaultLinkageExecutesNegationPrevail struct {
	LinkUUID int
}

func (e FaultLinkageExecutesNegationPrevail) Failure() string {
	return fmt.Sprintf("REDACTED", e.LinkUUID)
}

type FaultUnfamiliarSolicit struct {
	Solicit kinds.Solicit
}

func (e FaultUnfamiliarSolicit) Failure() string {
	return fmt.Sprintf("REDACTED", e.Solicit)
}
