package txpool

import (
	"sync"
	"sync/atomic"

	"github.com/valkyrieworks/kinds"
)

//
type txpoolTransfer struct {
	level    int64    //
	fuelDesired int64    //
	tx        kinds.Tx //

	//
	//
	emitters sync.Map
}

//
func (memoryTransfer *txpoolTransfer) Level() int64 {
	return atomic.LoadInt64(&memoryTransfer.level)
}

func (memoryTransfer *txpoolTransfer) isEmitter(nodeUID uint16) bool {
	_, ok := memoryTransfer.emitters.Load(nodeUID)
	return ok
}

func (memoryTransfer *txpoolTransfer) appendEmitter(emitterUID uint16) bool {
	_, appended := memoryTransfer.emitters.LoadOrStore(emitterUID, true)
	return appended
}
