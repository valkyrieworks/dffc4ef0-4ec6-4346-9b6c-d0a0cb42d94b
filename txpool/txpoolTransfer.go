package txpool

import (
	"sync"
	"sync/atomic"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type txpoolTransfer struct {
	altitude    int64    //
	fuelDesired int64    //
	tx        kinds.Tx //

	//
	//
	originators sync.Map
}

//
func (memoryTransfer *txpoolTransfer) Altitude() int64 {
	return atomic.LoadInt64(&memoryTransfer.altitude)
}

func (memoryTransfer *txpoolTransfer) equalsOriginator(nodeUUID uint16) bool {
	_, ok := memoryTransfer.originators.Load(nodeUUID)
	return ok
}

func (memoryTransfer *txpoolTransfer) appendOriginator(originatorUUID uint16) bool {
	_, appended := memoryTransfer.originators.LoadOrStore(originatorUUID, true)
	return appended
}
