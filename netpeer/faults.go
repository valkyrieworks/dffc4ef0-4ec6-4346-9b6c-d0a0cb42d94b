package netpeer

import (
	"errors"
	"fmt"
)

//
type FaultTemporary struct {
	Err error
}

func (e *FaultTemporary) Fault() string {
	return fmt.Sprintf("REDACTED", e.Err.Error())
}

func (e *FaultTemporary) Disclose() error {
	return e.Err
}

func TemporaryFaultFromAny(v any) (*FaultTemporary, bool) {
	err, ok := v.(error)
	if !ok {
		return nil, false
	}

	var te *FaultTemporary
	if !errors.As(err, &te) {
		return nil, false
	}

	return te, true
}
