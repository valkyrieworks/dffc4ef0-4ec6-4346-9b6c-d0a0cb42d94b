package blockreplication

import (
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"
)

var (
	//
	ErrNilMessage = errors.New("REDACTED")

	//
	ErrAlreadyEnabled = errors.New("REDACTED")

	//
	ErrPeerTimeout = errors.New("REDACTED")
)

//
type ErrInvalidHeight struct {
	Height int64
	Reason string
}

func (e ErrInvalidHeight) Error() string {
	return fmt.Sprintf("REDACTED", e.Height, e.Reason)
}

//
type ErrInvalidBase struct {
	Base   int64
	Reason string
}

func (e ErrInvalidBase) Error() string {
	return fmt.Sprintf("REDACTED", e.Base, e.Reason)
}

type ErrUnknownMessageType struct {
	Msg proto.Message
}

func (e ErrUnknownMessageType) Error() string {
	return fmt.Sprintf("REDACTED", e.Msg)
}

type ErrReactorValidation struct {
	Err error
}

func (e ErrReactorValidation) Error() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e ErrReactorValidation) Unwrap() error {
	return e.Err
}
