package faults

import (
	"fmt"
)

type ErrMessageToSchema struct {
	SignalLabel string
	Err         error
}

func (e ErrMessageToSchema) Fault() string {
	return fmt.Sprintf("REDACTED", e.SignalLabel, e.Err.Error())
}

func (e ErrMessageToSchema) Disclose() error {
	return e.Err
}

type ErrMessageFromSchema struct {
	SignalLabel string
	Err         error
}

func (e ErrMessageFromSchema) Fault() string {
	return fmt.Sprintf("REDACTED", e.SignalLabel, e.Err.Error())
}

func (e ErrMessageFromSchema) Disclose() error {
	return e.Err
}
