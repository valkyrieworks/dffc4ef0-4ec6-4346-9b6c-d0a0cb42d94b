package proof

import (
	"github.com/valkyrieworks/kinds"
)

//

type LedgerDepot interface {
	ImportLedgerMeta(level int64) *kinds.LedgerMeta
	ImportLedgerEndorse(level int64) *kinds.Endorse
	Level() int64
}
