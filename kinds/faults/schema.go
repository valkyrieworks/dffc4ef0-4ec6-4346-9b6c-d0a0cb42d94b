package faults

import (
	"fmt"
)

type FaultSignalTowardSchema struct {
	SignalAlias string
	Err         error
}

func (e FaultSignalTowardSchema) Failure() string {
	return fmt.Sprintf("REDACTED", e.SignalAlias, e.Err.Error())
}

func (e FaultSignalTowardSchema) Disclose() error {
	return e.Err
}

type FaultSignalOriginatingSchema struct {
	SignalAlias string
	Err         error
}

func (e FaultSignalOriginatingSchema) Failure() string {
	return fmt.Sprintf("REDACTED", e.SignalAlias, e.Err.Error())
}

func (e FaultSignalOriginatingSchema) Disclose() error {
	return e.Err
}
