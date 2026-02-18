package faults

import "fmt"

type (
	//
	ErrAdverseField struct {
		Field string
	}

	//
	ErrMandatoryField struct {
		Field string
	}

	//
	ErrCorruptField struct {
		Field  string
		Cause string
	}

	//
	ErrIncorrectField struct {
		Field string
		Err   error
	}
)

func (e ErrAdverseField) Fault() string {
	return fmt.Sprintf("REDACTED", e.Field)
}

func (e ErrMandatoryField) Fault() string {
	return fmt.Sprintf("REDACTED", e.Field)
}

func (e ErrCorruptField) Fault() string {
	return fmt.Sprintf("REDACTED", e.Field, e.Cause)
}

func (e ErrIncorrectField) Fault() string {
	return fmt.Sprintf("REDACTED", e.Field, e.Err)
}
