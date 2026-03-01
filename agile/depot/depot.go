package depot

import "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"

//
type Depot interface {
	//
	//
	//
	//
	PersistAgileLedger(lb *kinds.AgileLedger) error

	//
	//
	//
	//
	EraseAgileLedger(altitude int64) error

	//
	//
	//
	//
	//
	//
	AgileLedger(altitude int64) (*kinds.AgileLedger, error)

	//
	//
	//
	FinalAgileLedgerAltitude() (int64, error)

	//
	//
	//
	InitialAgileLedgerAltitude() (int64, error)

	//
	//
	//
	AgileLedgerPrior(altitude int64) (*kinds.AgileLedger, error)

	//
	//
	Trim(extent uint16) error

	//
	Extent() uint16
}
