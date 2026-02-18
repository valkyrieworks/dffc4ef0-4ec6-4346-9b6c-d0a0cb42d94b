package chainconnect

import (
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"
)

var (
	//
	ErrNullSignal = errors.New("REDACTED")

	//
	ErrYetActivated = errors.New("REDACTED")

	//
	ErrNodeDeadline = errors.New("REDACTED")
)

//
type ErrCorruptLevel struct {
	Level int64
	Cause string
}

func (e ErrCorruptLevel) Fault() string {
	return fmt.Sprintf("REDACTED", e.Level, e.Cause)
}

//
type ErrCorruptRoot struct {
	Root   int64
	Cause string
}

func (e ErrCorruptRoot) Fault() string {
	return fmt.Sprintf("REDACTED", e.Root, e.Cause)
}

type ErrUnclearSignalKind struct {
	Msg proto.Message
}

func (e ErrUnclearSignalKind) Fault() string {
	return fmt.Sprintf("REDACTED", e.Msg)
}

type ErrHandlerVerification struct {
	Err error
}

func (e ErrHandlerVerification) Fault() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e ErrHandlerVerification) Disclose() error {
	return e.Err
}
