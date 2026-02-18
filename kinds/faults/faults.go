package faults

import "fmt"

type (
	//
	//
	ErrCorruptEndorseLevel struct {
		Anticipated int64
		Factual   int64
	}

	//
	//
	ErrCorruptEndorseEndorsements struct {
		Anticipated int
		Factual   int
	}
)

func NewErrCorruptEndorseLevel(anticipated, factual int64) ErrCorruptEndorseLevel {
	return ErrCorruptEndorseLevel{
		Anticipated: anticipated,
		Factual:   factual,
	}
}

func (e ErrCorruptEndorseLevel) Fault() string {
	return fmt.Sprintf("REDACTED", e.Anticipated, e.Factual)
}

func NewErrCorruptEndorseEndorsements(anticipated, factual int) ErrCorruptEndorseEndorsements {
	return ErrCorruptEndorseEndorsements{
		Anticipated: anticipated,
		Factual:   factual,
	}
}

func (e ErrCorruptEndorseEndorsements) Fault() string {
	return fmt.Sprintf("REDACTED", e.Anticipated, e.Factual)
}
