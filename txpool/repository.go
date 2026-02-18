package txpool

import (
	"container/list"

	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
//
type TransferRepository interface {
	//
	Restore()

	//
	//
	Propel(tx kinds.Tx) bool

	//
	Delete(tx kinds.Tx)

	//
	//
	Has(tx kinds.Tx) bool
}

var _ TransferRepository = (*LRUTransferRepository)(nil)

//
//
type LRUTransferRepository struct {
	mtx      engineconnect.Lock
	volume     int
	repositoryIndex map[kinds.TransferKey]*list.Element
	catalog     *list.List
}

func NewLRUTransferRepository(storeVolume int) *LRUTransferRepository {
	return &LRUTransferRepository{
		volume:     storeVolume,
		repositoryIndex: make(map[kinds.TransferKey]*list.Element, storeVolume),
		catalog:     list.New(),
	}
}

//
//
func (c *LRUTransferRepository) FetchCatalog() *list.List {
	return c.catalog
}

//
func (c *LRUTransferRepository) Restore() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	clear(c.repositoryIndex)
	c.catalog.Init()
}

func (c *LRUTransferRepository) Propel(tx kinds.Tx) bool {
	key := tx.Key()
	c.mtx.Lock()
	defer c.mtx.Unlock()

	relocated, ok := c.repositoryIndex[key]
	if ok {
		c.catalog.MoveToBack(relocated)
		return false
	}

	if c.catalog.Len() >= c.volume {
		head := c.catalog.Front()
		if head != nil {
			headKey := head.Value.(kinds.TransferKey)
			delete(c.repositoryIndex, headKey)
			c.catalog.Remove(head)
		}
	}

	e := c.catalog.PushBack(key)
	c.repositoryIndex[key] = e

	return true
}

func (c *LRUTransferRepository) Delete(tx kinds.Tx) {
	key := tx.Key()
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if e, ok := c.repositoryIndex[key]; ok {
		delete(c.repositoryIndex, key)
		c.catalog.Remove(e)
	}
}

func (c *LRUTransferRepository) Has(tx kinds.Tx) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, ok := c.repositoryIndex[tx.Key()]
	return ok
}

//
type NoopTransferRepository struct{}

var _ TransferRepository = (*NoopTransferRepository)(nil)

func (NoopTransferRepository) Restore()             {}
func (NoopTransferRepository) Propel(kinds.Tx) bool { return true }
func (NoopTransferRepository) Delete(kinds.Tx)    {}
func (NoopTransferRepository) Has(kinds.Tx) bool  { return false }
