package proof

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//

type LedgerDepot interface {
	FetchLedgerSummary(altitude int64) *kinds.LedgerSummary
	FetchLedgerEndorse(altitude int64) *kinds.Endorse
	Altitude() int64
}
