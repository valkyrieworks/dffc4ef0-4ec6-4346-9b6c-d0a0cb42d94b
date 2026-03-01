package faults

import "fmt"

type (
	//
	FaultAdverseAttribute struct {
		Attribute string
	}

	//
	FaultMandatoryAttribute struct {
		Attribute string
	}

	//
	FaultUnfitAttribute struct {
		Attribute  string
		Rationale string
	}

	//
	FaultIncorrectAttribute struct {
		Attribute string
		Err   error
	}
)

func (e FaultAdverseAttribute) Failure() string {
	return fmt.Sprintf("REDACTED", e.Attribute)
}

func (e FaultMandatoryAttribute) Failure() string {
	return fmt.Sprintf("REDACTED", e.Attribute)
}

func (e FaultUnfitAttribute) Failure() string {
	return fmt.Sprintf("REDACTED", e.Attribute, e.Rationale)
}

func (e FaultIncorrectAttribute) Failure() string {
	return fmt.Sprintf("REDACTED", e.Attribute, e.Err)
}
