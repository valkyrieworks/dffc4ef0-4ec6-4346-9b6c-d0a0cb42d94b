package abciend

import (
	"fmt"

	"github.com/valkyrieworks/iface/kinds"
)

//
type ErrUnclearIfaceCarrier struct {
	Carrier string
}

func (e ErrUnclearIfaceCarrier) Fault() string {
	return fmt.Sprintf("REDACTED", e.Carrier)
}

type ErrUnforeseenReply struct {
	Reply kinds.Reply
	Cause   string
}

func (e ErrUnforeseenReply) Fault() string {
	return fmt.Sprintf("REDACTED", e.Reply.Item, e.Cause)
}
