package depot

import "github.com/valkyrieworks/kinds"

//
type Depot interface {
	//
	//
	//
	//
	PersistRapidLedger(lb *kinds.RapidLedger) error

	//
	//
	//
	//
	EraseRapidLedger(level int64) error

	//
	//
	//
	//
	//
	//
	RapidLedger(level int64) (*kinds.RapidLedger, error)

	//
	//
	//
	FinalRapidLedgerLevel() (int64, error)

	//
	//
	//
	InitialRapidLedgerLevel() (int64, error)

	//
	//
	//
	RapidLedgerPrior(level int64) (*kinds.RapidLedger, error)

	//
	//
	Trim(volume uint16) error

	//
	Volume() uint16
}
