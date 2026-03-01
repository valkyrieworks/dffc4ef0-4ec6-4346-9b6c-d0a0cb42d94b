package netp2p

import (
	"errors"
	"fmt"
)

//
type FailureFleeting struct {
	Err error
}

func (e *FailureFleeting) Failure() string {
	return fmt.Sprintf("REDACTED", e.Err.Error())
}

func (e *FailureFleeting) Disclose() error {
	return e.Err
}

func FleetingFailureOriginatingSome(v any) (*FailureFleeting, bool) {
	err, ok := v.(error)
	if !ok {
		return nil, false
	}

	var te *FailureFleeting
	if !errors.As(err, &te) {
		return nil, false
	}

	return te, true
}
