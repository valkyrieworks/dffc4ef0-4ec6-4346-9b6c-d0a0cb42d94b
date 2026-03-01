package faults

import "fmt"

type (
	//
	//
	FaultUnfitEndorseAltitude struct {
		Anticipated int64
		Existing   int64
	}

	//
	//
	FaultUnfitEndorseNotations struct {
		Anticipated int
		Existing   int
	}
)

func FreshFaultUnfitEndorseAltitude(anticipated, existing int64) FaultUnfitEndorseAltitude {
	return FaultUnfitEndorseAltitude{
		Anticipated: anticipated,
		Existing:   existing,
	}
}

func (e FaultUnfitEndorseAltitude) Failure() string {
	return fmt.Sprintf("REDACTED", e.Anticipated, e.Existing)
}

func FreshFaultUnfitEndorseNotations(anticipated, existing int) FaultUnfitEndorseNotations {
	return FaultUnfitEndorseNotations{
		Anticipated: anticipated,
		Existing:   existing,
	}
}

func (e FaultUnfitEndorseNotations) Failure() string {
	return fmt.Sprintf("REDACTED", e.Anticipated, e.Existing)
}
