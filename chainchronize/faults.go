package chainchronize

import (
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"
)

var (
	//
	FaultVoidSignal = errors.New("REDACTED")

	//
	FaultEarlierActivated = errors.New("REDACTED")

	//
	FaultNodeDeadline = errors.New("REDACTED")
)

//
type FaultUnfitAltitude struct {
	Altitude int64
	Rationale string
}

func (e FaultUnfitAltitude) Failure() string {
	return fmt.Sprintf("REDACTED", e.Altitude, e.Rationale)
}

//
type FaultUnfitFoundation struct {
	Foundation   int64
	Rationale string
}

func (e FaultUnfitFoundation) Failure() string {
	return fmt.Sprintf("REDACTED", e.Foundation, e.Rationale)
}

type FaultUnfamiliarSignalKind struct {
	Msg proto.Message
}

func (e FaultUnfamiliarSignalKind) Failure() string {
	return fmt.Sprintf("REDACTED", e.Msg)
}

type FaultHandlerCertification struct {
	Err error
}

func (e FaultHandlerCertification) Failure() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e FaultHandlerCertification) Disclose() error {
	return e.Err
}
