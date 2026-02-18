package blockreplication

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	bcschema "github.com/valkyrieworks/schema/consensuscore/blockreplication"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	BlockResponseMessagePrefixSize   = 4
	BlockResponseMessageFieldKeySize = 1
	MaxMsgSize                       = kinds.MaxBlockSizeBytes +
		BlockResponseMessagePrefixSize +
		BlockResponseMessageFieldKeySize
)

//
func ValidateMsg(pb proto.Message) error {
	if pb == nil {
		return ErrNilMessage
	}

	switch msg := pb.(type) {
	case *bcschema.BlockRequest:
		if msg.Height < 0 {
			return ErrInvalidHeight{Height: msg.Height, Reason: "REDACTED"}
		}
	case *bcschema.BlockResponse:
		//
		//
		//
		return nil
	case *bcschema.NoBlockResponse:
		if msg.Height < 0 {
			return ErrInvalidHeight{Height: msg.Height, Reason: "REDACTED"}
		}
	case *bcschema.StatusRequest:
		return nil
	case *bcschema.StatusResponse:
		if msg.Base < 0 {
			return ErrInvalidBase{Base: msg.Base, Reason: "REDACTED"}
		}
		if msg.Height < 0 {
			return ErrInvalidHeight{Height: msg.Height, Reason: "REDACTED"}
		}
		if msg.Base > msg.Height {
			return ErrInvalidHeight{Height: msg.Height, Reason: fmt.Sprintf("REDACTED", msg.Base)}
		}
	default:
		return ErrUnknownMessageType{Msg: msg}
	}
	return nil
}
