package abcicustomer

import (
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

//
type FaultUnfamiliarIfaceCarrier struct {
	Carrier string
}

func (e FaultUnfamiliarIfaceCarrier) Failure() string {
	return fmt.Sprintf("REDACTED", e.Carrier)
}

type FaultUnforeseenReply struct {
	Reply kinds.Reply
	Rationale   string
}

func (e FaultUnforeseenReply) Failure() string {
	return fmt.Sprintf("REDACTED", e.Reply.Datum, e.Rationale)
}
