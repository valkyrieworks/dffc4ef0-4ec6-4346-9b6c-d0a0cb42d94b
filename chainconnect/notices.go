package chainconnect

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	chainproto "github.com/valkyrieworks/schema/consensuscore/chainconnect"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	LedgerReplySignalHeadingVolume   = 4
	LedgerReplySignalFieldKeyVolume = 1
	MaximumMessageVolume                       = kinds.MaximumLedgerVolumeOctets +
		LedgerReplySignalHeadingVolume +
		LedgerReplySignalFieldKeyVolume
)

//
func CertifyMessage(pb proto.Message) error {
	if pb == nil {
		return ErrNullSignal
	}

	switch msg := pb.(type) {
	case *chainproto.LedgerQuery:
		if msg.Level < 0 {
			return ErrCorruptLevel{Level: msg.Level, Cause: "REDACTED"}
		}
	case *chainproto.LedgerReply:
		//
		//
		//
		return nil
	case *chainproto.NoLedgerReply:
		if msg.Level < 0 {
			return ErrCorruptLevel{Level: msg.Level, Cause: "REDACTED"}
		}
	case *chainproto.StatusQuery:
		return nil
	case *chainproto.StatusReply:
		if msg.Root < 0 {
			return ErrCorruptRoot{Root: msg.Root, Cause: "REDACTED"}
		}
		if msg.Level < 0 {
			return ErrCorruptLevel{Level: msg.Level, Cause: "REDACTED"}
		}
		if msg.Root > msg.Level {
			return ErrCorruptLevel{Level: msg.Level, Cause: fmt.Sprintf("REDACTED", msg.Root)}
		}
	default:
		return ErrUnclearSignalKind{Msg: msg}
	}
	return nil
}
