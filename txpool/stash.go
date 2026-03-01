package txpool

import (
	"container/list"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
//
type TransferStash interface {
	//
	Restore()

	//
	//
	Propel(tx kinds.Tx) bool

	//
	Discard(tx kinds.Tx)

	//
	//
	Has(tx kinds.Tx) bool
}

var _ TransferStash = (*LeastusedTransferStash)(nil)

//
//
type LeastusedTransferStash struct {
	mtx      commitchronize.Exclusion
	extent     int
	stashIndex map[kinds.TransferToken]*list.Element
	catalog     *list.List
}

func FreshLeastusedTransferStash(stashExtent int) *LeastusedTransferStash {
	return &LeastusedTransferStash{
		extent:     stashExtent,
		stashIndex: make(map[kinds.TransferToken]*list.Element, stashExtent),
		catalog:     list.New(),
	}
}

//
//
func (c *LeastusedTransferStash) FetchCatalog() *list.List {
	return c.catalog
}

//
func (c *LeastusedTransferStash) Restore() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	clear(c.stashIndex)
	c.catalog.Init()
}

func (c *LeastusedTransferStash) Propel(tx kinds.Tx) bool {
	key := tx.Key()
	c.mtx.Lock()
	defer c.mtx.Unlock()

	shifted, ok := c.stashIndex[key]
	if ok {
		c.catalog.MoveToBack(shifted)
		return false
	}

	if c.catalog.Len() >= c.extent {
		leading := c.catalog.Front()
		if leading != nil {
			leadingToken := leading.Value.(kinds.TransferToken)
			delete(c.stashIndex, leadingToken)
			c.catalog.Remove(leading)
		}
	}

	e := c.catalog.PushBack(key)
	c.stashIndex[key] = e

	return true
}

func (c *LeastusedTransferStash) Discard(tx kinds.Tx) {
	key := tx.Key()
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if e, ok := c.stashIndex[key]; ok {
		delete(c.stashIndex, key)
		c.catalog.Remove(e)
	}
}

func (c *LeastusedTransferStash) Has(tx kinds.Tx) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, ok := c.stashIndex[tx.Key()]
	return ok
}

//
type NooperationTransferStash struct{}

var _ TransferStash = (*NooperationTransferStash)(nil)

func (NooperationTransferStash) Restore()             {}
func (NooperationTransferStash) Propel(kinds.Tx) bool { return true }
func (NooperationTransferStash) Discard(kinds.Tx)    {}
func (NooperationTransferStash) Has(kinds.Tx) bool  { return false }
