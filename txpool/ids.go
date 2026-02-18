package txpool

import (
	"fmt"

	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
)

type txpoolIDXDatastore struct {
	mtx       engineconnect.ReadwriteLock
	nodeIndex   map[p2p.ID]uint16
	followingUID    uint16              //
	enabledIDXDatastore map[uint16]struct{} //
}

//
//
func (ids *txpoolIDXDatastore) AllocateForNode(node p2p.Node) {
	ids.mtx.Lock()
	defer ids.mtx.Unlock()

	currentUID := ids.followingNodeUID()
	ids.nodeIndex[node.ID()] = currentUID
	ids.enabledIDXDatastore[currentUID] = struct{}{}
}

//
//
func (ids *txpoolIDXDatastore) followingNodeUID() uint16 {
	if len(ids.enabledIDXDatastore) == MaximumEnabledIDXDatastore {
		panic(fmt.Sprintf("REDACTED", MaximumEnabledIDXDatastore))
	}

	_, uidPresent := ids.enabledIDXDatastore[ids.followingUID]
	for uidPresent {
		ids.followingUID++
		_, uidPresent = ids.enabledIDXDatastore[ids.followingUID]
	}
	currentUID := ids.followingUID
	ids.followingUID++
	return currentUID
}

//
func (ids *txpoolIDXDatastore) Recover(node p2p.Node) {
	ids.mtx.Lock()
	defer ids.mtx.Unlock()

	deletedUID, ok := ids.nodeIndex[node.ID()]
	if ok {
		delete(ids.enabledIDXDatastore, deletedUID)
		delete(ids.nodeIndex, node.ID())
	}
}

//
func (ids *txpoolIDXDatastore) FetchForNode(node p2p.Node) uint16 {
	ids.mtx.RLock()
	defer ids.mtx.RUnlock()

	return ids.nodeIndex[node.ID()]
}

func newTxpoolIDXDatastore() *txpoolIDXDatastore {
	return &txpoolIDXDatastore{
		nodeIndex:   make(map[p2p.ID]uint16),
		enabledIDXDatastore: map[uint16]struct{}{0: {}},
		followingUID:    1, //
	}
}
