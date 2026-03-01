package txpool

import (
	"fmt"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

type txpoolIDXDstore struct {
	mtx       commitchronize.ReadwriteExclusion
	nodeIndex   map[p2p.ID]uint16
	followingUUID    uint16              //
	dynamicIDXDstore map[uint16]struct{} //
}

//
//
func (ids *txpoolIDXDstore) AllocateForeachNode(node p2p.Node) {
	ids.mtx.Lock()
	defer ids.mtx.Unlock()

	currentUUID := ids.followingNodeUUID()
	ids.nodeIndex[node.ID()] = currentUUID
	ids.dynamicIDXDstore[currentUUID] = struct{}{}
}

//
//
func (ids *txpoolIDXDstore) followingNodeUUID() uint16 {
	if len(ids.dynamicIDXDstore) == MaximumDynamicIDXDstore {
		panic(fmt.Sprintf("REDACTED", MaximumDynamicIDXDstore))
	}

	_, uuidPresent := ids.dynamicIDXDstore[ids.followingUUID]
	for uuidPresent {
		ids.followingUUID++
		_, uuidPresent = ids.dynamicIDXDstore[ids.followingUUID]
	}
	currentUUID := ids.followingUUID
	ids.followingUUID++
	return currentUUID
}

//
func (ids *txpoolIDXDstore) Recover(node p2p.Node) {
	ids.mtx.Lock()
	defer ids.mtx.Unlock()

	discardedUUID, ok := ids.nodeIndex[node.ID()]
	if ok {
		delete(ids.dynamicIDXDstore, discardedUUID)
		delete(ids.nodeIndex, node.ID())
	}
}

//
func (ids *txpoolIDXDstore) ObtainForeachNode(node p2p.Node) uint16 {
	ids.mtx.RLock()
	defer ids.mtx.RUnlock()

	return ids.nodeIndex[node.ID()]
}

func freshTxpoolIDXDstore() *txpoolIDXDstore {
	return &txpoolIDXDstore{
		nodeIndex:   make(map[p2p.ID]uint16),
		dynamicIDXDstore: map[uint16]struct{}{0: {}},
		followingUUID:    1, //
	}
}
