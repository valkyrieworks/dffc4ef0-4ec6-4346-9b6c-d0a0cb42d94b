package txpool

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/ringlist"
	"github.com/valkyrieworks/utils/log"
	cometmath "github.com/valkyrieworks/utils/math"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/gateway"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
//
type CCatalogTxpool struct {
	level   atomic.Int64 //
	transOctets atomic.Int64 //

	//
	alertedTransAccessible atomic.Bool
	transAccessible         chan struct{} //

	settings *settings.TxpoolSettings

	//
	//
	modifyMutex engineconnect.ReadwriteLock
	preInspect  PreInspectFunction
	submitInspect SubmitInspectFunction

	txs          *ringlist.CCatalog //
	gatewayApplicationLink gateway.ApplicationLinkTxpool

	//
	revalidate *revalidate

	//
	//
	transIndex sync.Map

	//
	//
	repository TransferRepository

	tracer  log.Tracer
	stats *Stats
}

var _ Txpool = &CCatalogTxpool{}

//
type CCatalogTxpoolSetting func(*CCatalogTxpool)

//
//
func NewCCatalogTxpool(
	cfg *settings.TxpoolSettings,
	gatewayApplicationLink gateway.ApplicationLinkTxpool,
	level int64,
	options ...CCatalogTxpoolSetting,
) *CCatalogTxpool {
	mp := &CCatalogTxpool{
		settings:       cfg,
		gatewayApplicationLink: gatewayApplicationLink,
		txs:          ringlist.New(),
		revalidate:      newRevalidate(),
		tracer:       log.NewNoopTracer(),
		stats:      NoopStats(),
	}
	mp.level.Store(level)

	if cfg.RepositoryVolume > 0 {
		mp.repository = NewLRUTransferRepository(cfg.RepositoryVolume)
	} else {
		mp.repository = NoopTransferRepository{}
	}

	gatewayApplicationLink.CollectionReplyCallback(mp.universalCallbackfn)

	for _, setting := range options {
		setting(mp)
	}

	return mp
}

func (mem *CCatalogTxpool) fetchCComponent(transferKey kinds.TransferKey) (*ringlist.CComponent, bool) {
	if e, ok := mem.transIndex.Load(transferKey); ok {
		return e.(*ringlist.CComponent), true
	}
	return nil, false
}

func (mem *CCatalogTxpool) fetchMemoryTransfer(transferKey kinds.TransferKey) *txpoolTransfer {
	if e, ok := mem.fetchCComponent(transferKey); ok {
		return e.Item.(*txpoolTransfer)
	}
	return nil
}

func (mem *CCatalogTxpool) deleteAllTrans() {
	for e := mem.txs.Head(); e != nil; e = e.Following() {
		mem.txs.Delete(e)
		e.UnplugPrevious()
	}

	mem.transIndex.Range(func(key, _ any) bool {
		mem.transIndex.Delete(key)
		return true
	})
}

//
func (mem *CCatalogTxpool) ActivateTransAccessible() {
	mem.transAccessible = make(chan struct{}, 1)
}

//
func (mem *CCatalogTxpool) AssignTracer(l log.Tracer) {
	mem.tracer = l
}

//
//
//
func WithPreInspect(f PreInspectFunction) CCatalogTxpoolSetting {
	return func(mem *CCatalogTxpool) { mem.preInspect = f }
}

//
//
//
func WithSubmitInspect(f SubmitInspectFunction) CCatalogTxpoolSetting {
	return func(mem *CCatalogTxpool) { mem.submitInspect = f }
}

//
func WithStats(stats *Stats) CCatalogTxpoolSetting {
	return func(mem *CCatalogTxpool) { mem.stats = stats }
}

//
func (mem *CCatalogTxpool) Secure() {
	if mem.revalidate.collectionRevalidateComplete() {
		mem.tracer.Diagnose("REDACTED")
	}
	mem.modifyMutex.Lock()
}

//
func (mem *CCatalogTxpool) Release() {
	mem.modifyMutex.Unlock()
}

//
func (mem *CCatalogTxpool) Volume() int {
	return mem.txs.Len()
}

//
func (mem *CCatalogTxpool) VolumeOctets() int64 {
	return mem.transOctets.Load()
}

//
func (mem *CCatalogTxpool) PurgeApplicationLink() error {
	err := mem.gatewayApplicationLink.Purge(context.TODO())
	if err != nil {
		return ErrPurgeApplicationLink{Err: err}
	}

	return nil
}

//
func (mem *CCatalogTxpool) Purge() {
	mem.modifyMutex.Lock()
	defer mem.modifyMutex.Unlock()

	mem.transOctets.Store(0)
	mem.repository.Restore()

	mem.deleteAllTrans()
}

//
//
//
//
//
func (mem *CCatalogTxpool) TransHead() *ringlist.CComponent {
	return mem.txs.Head()
}

//
//
//
//
//
func (mem *CCatalogTxpool) TransWaitChan() <-chan struct{} {
	return mem.txs.WaitChan()
}

//
//
//
//
//
//
//
//
func (mem *CCatalogTxpool) InspectTransfer(
	tx kinds.Tx,
	cb func(*iface.ReplyInspectTransfer),
	transferDetails TransferDetails,
) error {
	mem.modifyMutex.RLock()
	//
	defer mem.modifyMutex.RUnlock()

	transferVolume := len(tx)

	if err := mem.isComplete(transferVolume); err != nil {
		mem.stats.DeclinedTrans.Add(1)
		return err
	}

	if transferVolume > mem.settings.MaximumTransferOctets {
		return ErrTransferTooBulky{
			Max:    mem.settings.MaximumTransferOctets,
			Factual: transferVolume,
		}
	}

	if mem.preInspect != nil {
		if err := mem.preInspect(tx); err != nil {
			return ErrPreInspect{Err: err}
		}
	}

	//
	if err := mem.gatewayApplicationLink.Fault(); err != nil {
		return ErrApplicationLinkTxpool{Err: err}
	}

	if !mem.repository.Propel(tx) { //
		mem.stats.YetAcceptedTrans.Add(1)
		//
		//
		//
		//
		if memoryTransfer := mem.fetchMemoryTransfer(tx.Key()); memoryTransfer != nil {
			memoryTransfer.appendEmitter(transferDetails.EmitterUID)
			//
			//
			//
		}
		return ErrTransferInRepository
	}

	requestOutput, err := mem.gatewayApplicationLink.InspectTransferAsync(context.TODO(), &iface.QueryInspectTransfer{Tx: tx})
	if err != nil {
		panic(fmt.Errorf("REDACTED", log.NewIdleFormat("REDACTED", tx.Digest()), err))
	}
	requestOutput.CollectionCallback(mem.requestOutputCallbackfn(tx, transferDetails, cb))

	return nil
}

//
//
//
//
//
//
//
//
//
func (mem *CCatalogTxpool) universalCallbackfn(req *iface.Query, res *iface.Reply) {
	switch r := req.Item.(type) {
	case *iface.Query_Transfercheck:
		//
		if r.InspectTransfer.Kind != iface.Transfercheckkind_Revalidate {
			return
		}
	default:
		//
		return
	}

	switch r := res.Item.(type) {
	case *iface.Reply_Transfercheck:
		tx := kinds.Tx(req.FetchInspectTransfer().Tx)
		if mem.revalidate.done() {
			mem.tracer.Fault("REDACTED",
				"REDACTED", log.NewIdleFormat("REDACTED", tx.Key()))
			return
		}
		mem.stats.RevalidateInstances.Add(1)
		mem.outputCallbackfnRevalidate(tx, r.InspectTransfer)

		//
		mem.stats.Volume.Set(float64(mem.Volume()))

	default:
		//
	}
}

//
//
//
//
//
//
//
//
//
func (mem *CCatalogTxpool) requestOutputCallbackfn(
	tx []byte,
	transferDetails TransferDetails,
	outsideCallbackfn func(*iface.ReplyInspectTransfer),
) func(res *iface.Reply) {
	return func(res *iface.Reply) {
		if !mem.revalidate.done() {
			panic(log.NewIdleFormat("REDACTED",
				kinds.Tx(tx).Digest()))
		}

		mem.outputCallbackfnInitialTime(tx, transferDetails, res)

		//
		mem.stats.Volume.Set(float64(mem.Volume()))
		mem.stats.VolumeOctets.Set(float64(mem.VolumeOctets()))

		//
		if outsideCallbackfn != nil {
			outsideCallbackfn(res.FetchInspectTransfer())
		}
	}
}

//
//
func (mem *CCatalogTxpool) appendTransfer(memoryTransfer *txpoolTransfer) {
	e := mem.txs.PropelRear(memoryTransfer)
	mem.transIndex.Store(memoryTransfer.tx.Key(), e)
	mem.transOctets.Add(int64(len(memoryTransfer.tx)))
	mem.stats.TransferVolumeOctets.Observe(float64(len(memoryTransfer.tx)))
}

//
//
//
//
func (mem *CCatalogTxpool) DeleteTransferByKey(transferKey kinds.TransferKey) error {
	if element, ok := mem.fetchCComponent(transferKey); ok {
		mem.txs.Delete(element)
		element.UnplugPrevious()
		mem.transIndex.Delete(transferKey)
		tx := element.Item.(*txpoolTransfer).tx
		mem.transOctets.Add(int64(-len(tx)))
		return nil
	}
	return ErrTransferNegateLocated
}

func (mem *CCatalogTxpool) isComplete(transferVolume int) error {
	memoryVolume := mem.Volume()
	transOctets := mem.VolumeOctets()
	if memoryVolume >= mem.settings.Volume || int64(transferVolume)+transOctets > mem.settings.MaximumTransOctets {
		return ErrTxpoolIsComplete{
			CountTrans:      memoryVolume,
			MaximumTrans:      mem.settings.Volume,
			TransOctets:    transOctets,
			MaximumTransOctets: mem.settings.MaximumTransOctets,
		}
	}

	if mem.revalidate.regardedComplete() {
		return ErrRevalidateComplete
	}

	return nil
}

//
//
//
//
func (mem *CCatalogTxpool) outputCallbackfnInitialTime(
	tx []byte,
	transferDetails TransferDetails,
	res *iface.Reply,
) {
	switch r := res.Item.(type) {
	case *iface.Reply_Transfercheck:
		var submitInspectErr error
		if mem.submitInspect != nil {
			submitInspectErr = mem.submitInspect(tx, r.InspectTransfer)
		}
		if (r.InspectTransfer.Code == iface.CodeKindSuccess) && submitInspectErr == nil {
			//
			//
			if err := mem.isComplete(len(tx)); err != nil {
				//
				mem.repository.Delete(tx)
				//
				mem.tracer.Diagnose(err.Error())
				mem.stats.DeclinedTrans.Add(1)
				return
			}

			//
			if e, ok := mem.transIndex.Load(kinds.Tx(tx).Key()); ok {
				memoryTransfer := e.(*ringlist.CComponent).Item.(*txpoolTransfer)
				memoryTransfer.appendEmitter(transferDetails.EmitterUID)
				mem.tracer.Diagnose(
					"REDACTED",
					"REDACTED", kinds.Tx(tx).Digest(),
					"REDACTED", r,
					"REDACTED", mem.level.Load(),
					"REDACTED", mem.Volume(),
				)
				mem.stats.DeclinedTrans.Add(1)
				return
			}

			memoryTransfer := &txpoolTransfer{
				level:    mem.level.Load(),
				fuelDesired: r.InspectTransfer.FuelDesired,
				tx:        tx,
			}
			memoryTransfer.appendEmitter(transferDetails.EmitterUID)
			mem.appendTransfer(memoryTransfer)
			mem.tracer.Diagnose(
				"REDACTED",
				"REDACTED", kinds.Tx(tx).Digest(),
				"REDACTED", r,
				"REDACTED", mem.level.Load(),
				"REDACTED", mem.Volume(),
			)
			mem.adviseTransAccessible()
		} else {
			//
			mem.tracer.Diagnose(
				"REDACTED",
				"REDACTED", kinds.Tx(tx).Digest(),
				"REDACTED", transferDetails.EmitterP2pid,
				"REDACTED", r,
				"REDACTED", submitInspectErr,
			)
			mem.stats.ErroredTrans.Add(1)

			if !mem.settings.RetainCorruptTransInRepository {
				//
				mem.repository.Delete(tx)
			}
		}

	default:
		//
	}
}

//
//
//
//
func (mem *CCatalogTxpool) outputCallbackfnRevalidate(tx kinds.Tx, res *iface.ReplyInspectTransfer) {
	//
	if !mem.revalidate.locateFollowingElementCoordinating(&tx) {
		//
		return
	}

	var submitInspectErr error
	if mem.submitInspect != nil {
		submitInspectErr = mem.submitInspect(tx, res)
	}

	if (res.Code != iface.CodeKindSuccess) || submitInspectErr != nil {
		//
		mem.tracer.Diagnose("REDACTED", "REDACTED", tx.Digest(), "REDACTED", res, "REDACTED", submitInspectErr)
		if err := mem.DeleteTransferByKey(tx.Key()); err != nil {
			mem.tracer.Diagnose("REDACTED", "REDACTED", err)
		}
		if !mem.settings.RetainCorruptTransInRepository {
			mem.repository.Delete(tx)
			mem.stats.ExpelledTrans.Add(1)
		}
	}
}

//
func (mem *CCatalogTxpool) TransAccessible() <-chan struct{} {
	return mem.transAccessible
}

func (mem *CCatalogTxpool) adviseTransAccessible() {
	if mem.Volume() == 0 {
		panic("REDACTED")
	}
	if mem.transAccessible != nil && mem.alertedTransAccessible.CompareAndSwap(false, true) {
		//
		select {
		case mem.transAccessible <- struct{}{}:
		default:
		}
	}
}

//
func (mem *CCatalogTxpool) HarvestMaximumOctetsMaximumFuel(maximumOctets, maximumFuel int64) kinds.Txs {
	mem.modifyMutex.RLock()
	defer mem.modifyMutex.RUnlock()

	var (
		sumFuel    int64
		activeVolume int64
	)

	//
	//
	//
	txs := make([]kinds.Tx, 0, mem.txs.Len())
	for e := mem.txs.Head(); e != nil; e = e.Following() {
		memoryTransfer := e.Item.(*txpoolTransfer)

		txs = append(txs, memoryTransfer.tx)

		dataVolume := kinds.CalculateSchemaVolumeForTrans([]kinds.Tx{memoryTransfer.tx})

		//
		if maximumOctets > -1 && activeVolume+dataVolume > maximumOctets {
			return txs[:len(txs)-1]
		}

		activeVolume += dataVolume

		//
		//
		//
		//
		newSumFuel := sumFuel + memoryTransfer.fuelDesired
		if maximumFuel > -1 && newSumFuel > maximumFuel {
			return txs[:len(txs)-1]
		}
		sumFuel = newSumFuel
	}
	return txs
}

//
func (mem *CCatalogTxpool) HarvestMaximumTrans(max int) kinds.Txs {
	mem.modifyMutex.RLock()
	defer mem.modifyMutex.RUnlock()

	if max < 0 {
		max = mem.txs.Len()
	}

	txs := make([]kinds.Tx, 0, cometmath.MinimumInteger(mem.txs.Len(), max))
	for e := mem.txs.Head(); e != nil && len(txs) < max; e = e.Following() {
		memoryTransfer := e.Item.(*txpoolTransfer)
		txs = append(txs, memoryTransfer.tx)
	}
	return txs
}

//
func (mem *CCatalogTxpool) Modify(
	level int64,
	txs kinds.Txs,
	transferOutcomes []*iface.InvokeTransferOutcome,
	preInspect PreInspectFunction,
	submitInspect SubmitInspectFunction,
) error {
	mem.tracer.Diagnose("REDACTED", "REDACTED", level, "REDACTED", len(txs))

	//
	mem.level.Store(level)
	mem.alertedTransAccessible.Store(false)

	if preInspect != nil {
		mem.preInspect = preInspect
	}
	if submitInspect != nil {
		mem.submitInspect = submitInspect
	}

	for i, tx := range txs {
		if transferOutcomes[i].Code == iface.CodeKindSuccess {
			//
			_ = mem.repository.Propel(tx)
		} else if !mem.settings.RetainCorruptTransInRepository {
			//
			mem.repository.Delete(tx)
		}

		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		if err := mem.DeleteTransferByKey(tx.Key()); err != nil {
			mem.tracer.Diagnose("REDACTED",
				"REDACTED", tx.Key(),
				"REDACTED", err.Error())
		}
	}

	//
	if mem.settings.Revalidate {
		mem.revalidateTrans()
	}

	//
	if mem.Volume() > 0 {
		mem.adviseTransAccessible()
	}

	//
	mem.stats.Volume.Set(float64(mem.Volume()))
	mem.stats.VolumeOctets.Set(float64(mem.VolumeOctets()))

	return nil
}

//
//
func (mem *CCatalogTxpool) revalidateTrans() {
	mem.tracer.Diagnose("REDACTED", "REDACTED", mem.level.Load(), "REDACTED", mem.Volume())

	if mem.Volume() <= 0 {
		return
	}

	mem.revalidate.init(mem.txs.Head(), mem.txs.Rear())

	//
	//
	for e := mem.txs.Head(); e != nil; e = e.Following() {
		tx := e.Item.(*txpoolTransfer).tx
		mem.revalidate.countAwaitingTrans.Add(1)

		//
		//
		_, err := mem.gatewayApplicationLink.InspectTransferAsync(context.TODO(), &iface.QueryInspectTransfer{
			Tx:   tx,
			Kind: iface.Transfercheckkind_Revalidate,
		})
		if err != nil {
			panic(fmt.Errorf("REDACTED", log.NewIdleFormat("REDACTED", tx.Digest()), err))
		}
	}

	//
	mem.gatewayApplicationLink.Purge(context.TODO())

	//
	//
	select {
	case <-time.After(mem.settings.RevalidateDeadline):
		mem.revalidate.collectionDone()
		mem.tracer.Fault("REDACTED")
	case <-mem.revalidate.doneRevalidating():
	}

	if n := mem.revalidate.countAwaitingTrans.Load(); n > 0 {
		mem.tracer.Fault("REDACTED", "REDACTED", n)
	}
	mem.tracer.Diagnose("REDACTED", "REDACTED", mem.level.Load(), "REDACTED", mem.Volume())
}

//
//
//
//
//
//
type revalidate struct {
	locator        *ringlist.CComponent //
	end           *ringlist.CComponent //
	doneChan        chan struct{}   //
	countAwaitingTrans atomic.Int32    //
	isRevalidating  atomic.Bool     //
	revalidateComplete   atomic.Bool     //
}

func newRevalidate() *revalidate {
	return &revalidate{
		doneChan: make(chan struct{}, 1),
	}
}

func (rc *revalidate) init(initial, final *ringlist.CComponent) {
	if !rc.done() {
		panic("REDACTED")
	}
	rc.locator = initial
	rc.end = final
	rc.countAwaitingTrans.Store(0)
	rc.isRevalidating.Store(true)
}

//
//
func (rc *revalidate) done() bool {
	return !rc.isRevalidating.Load()
}

//
func (rc *revalidate) collectionDone() {
	rc.locator = nil
	rc.revalidateComplete.Store(false)
	rc.isRevalidating.Store(false)
}

//
func (rc *revalidate) collectionFollowingElement() {
	rc.locator = rc.locator.Following()
}

//
//
func (rc *revalidate) attemptConclude() bool {
	if rc.locator == rc.end {
		//
		rc.collectionDone()
	}
	if rc.done() {
		//
		select {
		case rc.doneChan <- struct{}{}:
		default:
		}
		return true
	}
	return false
}

//
//
//
//
//
//
//
func (rc *revalidate) locateFollowingElementCoordinating(tx *kinds.Tx) bool {
	located := false
	for ; !rc.done(); rc.collectionFollowingElement() {
		anticipatedTransfer := rc.locator.Item.(*txpoolTransfer).tx
		if bytes.Equal(*tx, anticipatedTransfer) {
			//
			located = true
			rc.countAwaitingTrans.Add(-1)
			break
		}
	}

	if !rc.attemptConclude() {
		//
		rc.collectionFollowingElement()
	}
	return located
}

//
func (rc *revalidate) doneRevalidating() <-chan struct{} {
	return rc.doneChan
}

//
//
func (rc *revalidate) collectionRevalidateComplete() bool {
	revalidating := !rc.done()
	revalidateComplete := rc.revalidateComplete.Swap(revalidating)
	return revalidating != revalidateComplete
}

//
//
func (rc *revalidate) regardedComplete() bool {
	return rc.revalidateComplete.Load()
}
