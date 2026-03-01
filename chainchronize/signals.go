package chainchronize

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	chainchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/chainchronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	//
	LedgerReplySignalHeadingExtent   = 4
	LedgerReplySignalAttributeTokenExtent = 1
	MaximumSignalExtent                       = kinds.MaximumLedgerExtentOctets +
		LedgerReplySignalHeadingExtent +
		LedgerReplySignalAttributeTokenExtent
)

//
func CertifySignal(pb proto.Message) error {
	if pb == nil {
		return FaultVoidSignal
	}

	switch msg := pb.(type) {
	case *chainchema.LedgerSolicit:
		if msg.Altitude < 0 {
			return FaultUnfitAltitude{Altitude: msg.Altitude, Rationale: "REDACTED"}
		}
	case *chainchema.LedgerReply:
		//
		//
		//
		return nil
	case *chainchema.NegativeLedgerReply:
		if msg.Altitude < 0 {
			return FaultUnfitAltitude{Altitude: msg.Altitude, Rationale: "REDACTED"}
		}
	case *chainchema.ConditionSolicit:
		return nil
	case *chainchema.ConditionReply:
		if msg.Foundation < 0 {
			return FaultUnfitFoundation{Foundation: msg.Foundation, Rationale: "REDACTED"}
		}
		if msg.Altitude < 0 {
			return FaultUnfitAltitude{Altitude: msg.Altitude, Rationale: "REDACTED"}
		}
		if msg.Foundation > msg.Altitude {
			return FaultUnfitAltitude{Altitude: msg.Altitude, Rationale: fmt.Sprintf("REDACTED", msg.Foundation)}
		}
	default:
		return FaultUnfamiliarSignalKind{Msg: msg}
	}
	return nil
}
