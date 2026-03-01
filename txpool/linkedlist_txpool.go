package txpool

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/linkedlist"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
//
type CNCatalogTxpool struct {
	altitude   atomic.Int64 //
	transOctets atomic.Int64 //

	//
	alertedTransAccessible atomic.Bool
	transAccessible         chan struct{} //

	settings *settings.TxpoolSettings

	//
	//
	reviseMutex commitchronize.ReadwriteExclusion
	priorInspect  PriorInspectMethod
	submitInspect RelayInspectMethod

	txs          *linkedlist.CNCatalog //
	delegateApplicationLink delegate.ApplicationLinkTxpool

	//
	reinspect *reinspect

	//
	//
	transIndex sync.Map

	//
	//
	stash TransferStash

	tracer  log.Tracer
	telemetry *Telemetry
}

var _ Txpool = &CNCatalogTxpool{}

//
type CNCatalogTxpoolSelection func(*CNCatalogTxpool)

//
//
func FreshCNCatalogTxpool(
	cfg *settings.TxpoolSettings,
	delegateApplicationLink delegate.ApplicationLinkTxpool,
	altitude int64,
	choices ...CNCatalogTxpoolSelection,
) *CNCatalogTxpool {
	mp := &CNCatalogTxpool{
		settings:       cfg,
		delegateApplicationLink: delegateApplicationLink,
		txs:          linkedlist.New(),
		reinspect:      freshReinspect(),
		tracer:       log.FreshNooperationTracer(),
		telemetry:      NooperationTelemetry(),
	}
	mp.altitude.Store(altitude)

	if cfg.StashExtent > 0 {
		mp.stash = FreshLeastusedTransferStash(cfg.StashExtent)
	} else {
		mp.stash = NooperationTransferStash{}
	}

	delegateApplicationLink.AssignReplyClbk(mp.universalClbk)

	for _, selection := range choices {
		selection(mp)
	}

	return mp
}

func (mem *CNCatalogTxpool) fetchCNComponent(transferToken kinds.TransferToken) (*linkedlist.CNComponent, bool) {
	if e, ok := mem.transIndex.Load(transferToken); ok {
		return e.(*linkedlist.CNComponent), true
	}
	return nil, false
}

func (mem *CNCatalogTxpool) fetchMemoryTransfer(transferToken kinds.TransferToken) *txpoolTransfer {
	if e, ok := mem.fetchCNComponent(transferToken); ok {
		return e.Datum.(*txpoolTransfer)
	}
	return nil
}

func (mem *CNCatalogTxpool) discardEveryTrans() {
	for e := mem.txs.Leading(); e != nil; e = e.Following() {
		mem.txs.Discard(e)
		e.UncouplePrevious()
	}

	mem.transIndex.Range(func(key, _ any) bool {
		mem.transIndex.Delete(key)
		return true
	})
}

//
func (mem *CNCatalogTxpool) ActivateTransAccessible() {
	mem.transAccessible = make(chan struct{}, 1)
}

//
func (mem *CNCatalogTxpool) AssignTracer(l log.Tracer) {
	mem.tracer = l
}

//
//
//
func UsingPriorInspect(f PriorInspectMethod) CNCatalogTxpoolSelection {
	return func(mem *CNCatalogTxpool) { mem.priorInspect = f }
}

//
//
//
func UsingRelayInspect(f RelayInspectMethod) CNCatalogTxpoolSelection {
	return func(mem *CNCatalogTxpool) { mem.submitInspect = f }
}

//
func UsingTelemetry(telemetry *Telemetry) CNCatalogTxpoolSelection {
	return func(mem *CNCatalogTxpool) { mem.telemetry = telemetry }
}

//
func (mem *CNCatalogTxpool) Secure() {
	if mem.reinspect.assignReinspectComplete() {
		mem.tracer.Diagnose("REDACTED")
	}
	mem.reviseMutex.Lock()
}

//
func (mem *CNCatalogTxpool) Release() {
	mem.reviseMutex.Unlock()
}

//
func (mem *CNCatalogTxpool) Extent() int {
	return mem.txs.Len()
}

//
func (mem *CNCatalogTxpool) ExtentOctets() int64 {
	return mem.transOctets.Load()
}

//
func (mem *CNCatalogTxpool) PurgeApplicationLink() error {
	err := mem.delegateApplicationLink.Purge(context.TODO())
	if err != nil {
		return FaultPurgeApplicationLink{Err: err}
	}

	return nil
}

//
func (mem *CNCatalogTxpool) Purge() {
	mem.reviseMutex.Lock()
	defer mem.reviseMutex.Unlock()

	mem.transOctets.Store(0)
	mem.stash.Restore()

	mem.discardEveryTrans()
}

//
//
//
//
//
func (mem *CNCatalogTxpool) TransLeading() *linkedlist.CNComponent {
	return mem.txs.Leading()
}

//
//
//
//
//
func (mem *CNCatalogTxpool) TransPauseChannel() <-chan struct{} {
	return mem.txs.PauseChnl()
}

//
//
//
//
//
//
//
//
func (mem *CNCatalogTxpool) InspectTransfer(
	tx kinds.Tx,
	cb func(*iface.ReplyInspectTransfer),
	transferDetails TransferDetails,
) error {
	mem.reviseMutex.RLock()
	//
	defer mem.reviseMutex.RUnlock()

	transferExtent := len(tx)

	if err := mem.equalsComplete(transferExtent); err != nil {
		mem.telemetry.DeclinedTrans.Add(1)
		return err
	}

	if transferExtent > mem.settings.MaximumTransferOctets {
		return FaultTransferExcessivelyAmple{
			Max:    mem.settings.MaximumTransferOctets,
			Existing: transferExtent,
		}
	}

	if mem.priorInspect != nil {
		if err := mem.priorInspect(tx); err != nil {
			return FaultPriorInspect{Err: err}
		}
	}

	//
	if err := mem.delegateApplicationLink.Failure(); err != nil {
		return FaultApplicationLinkTxpool{Err: err}
	}

	if !mem.stash.Propel(tx) { //
		mem.telemetry.EarlierAcceptedTrans.Add(1)
		//
		//
		//
		//
		if memoryTransfer := mem.fetchMemoryTransfer(tx.Key()); memoryTransfer != nil {
			memoryTransfer.appendOriginator(transferDetails.OriginatorUUID)
			//
			//
			//
		}
		return FaultTransferInsideStash
	}

	requestResult, err := mem.delegateApplicationLink.InspectTransferAsyncronous(context.TODO(), &iface.SolicitInspectTransfer{Tx: tx})
	if err != nil {
		panic(fmt.Errorf("REDACTED", log.FreshIdleFormat("REDACTED", tx.Digest()), err))
	}
	requestResult.AssignClbk(mem.requestResultClbk(tx, transferDetails, cb))

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
func (mem *CNCatalogTxpool) universalClbk(req *iface.Solicit, res *iface.Reply) {
	switch r := req.Datum.(type) {
	case *iface.Solicit_Inspecttrans:
		//
		if r.InspectTransfer.Kind != iface.Inspecttranskind_Reinspect {
			return
		}
	default:
		//
		return
	}

	switch r := res.Datum.(type) {
	case *iface.Reply_Inspecttrans:
		tx := kinds.Tx(req.ObtainInspectTransfer().Tx)
		if mem.reinspect.complete() {
			mem.tracer.Failure("REDACTED",
				"REDACTED", log.FreshIdleFormat("REDACTED", tx.Key()))
			return
		}
		mem.telemetry.ReinspectMultiples.Add(1)
		mem.resultClbkReinspect(tx, r.InspectTransfer)

		//
		mem.telemetry.Extent.Set(float64(mem.Extent()))

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
func (mem *CNCatalogTxpool) requestResultClbk(
	tx []byte,
	transferDetails TransferDetails,
	outsideClbk func(*iface.ReplyInspectTransfer),
) func(res *iface.Reply) {
	return func(res *iface.Reply) {
		if !mem.reinspect.complete() {
			panic(log.FreshIdleFormat("REDACTED",
				kinds.Tx(tx).Digest()))
		}

		mem.resultClbkInitialMoment(tx, transferDetails, res)

		//
		mem.telemetry.Extent.Set(float64(mem.Extent()))
		mem.telemetry.ExtentOctets.Set(float64(mem.ExtentOctets()))

		//
		if outsideClbk != nil {
			outsideClbk(res.ObtainInspectTransfer())
		}
	}
}

//
//
func (mem *CNCatalogTxpool) appendTransfer(memoryTransfer *txpoolTransfer) {
	e := mem.txs.PropelRear(memoryTransfer)
	mem.transIndex.Store(memoryTransfer.tx.Key(), e)
	mem.transOctets.Add(int64(len(memoryTransfer.tx)))
	mem.telemetry.TransferExtentOctets.Observe(float64(len(memoryTransfer.tx)))
}

//
//
//
//
func (mem *CNCatalogTxpool) DiscardTransferViaToken(transferToken kinds.TransferToken) error {
	if member, ok := mem.fetchCNComponent(transferToken); ok {
		mem.txs.Discard(member)
		member.UncouplePrevious()
		mem.transIndex.Delete(transferToken)
		tx := member.Datum.(*txpoolTransfer).tx
		mem.transOctets.Add(int64(-len(tx)))
		return nil
	}
	return FaultTransferNegationDetected
}

func (mem *CNCatalogTxpool) equalsComplete(transferExtent int) error {
	memoryExtent := mem.Extent()
	transOctets := mem.ExtentOctets()
	if memoryExtent >= mem.settings.Extent || int64(transferExtent)+transOctets > mem.settings.MaximumTransOctets {
		return FaultTxpoolEqualsComplete{
			CountTrans:      memoryExtent,
			MaximumTrans:      mem.settings.Extent,
			TransOctets:    transOctets,
			MaximumTransOctets: mem.settings.MaximumTransOctets,
		}
	}

	if mem.reinspect.deemedComplete() {
		return FaultReinspectComplete
	}

	return nil
}

//
//
//
//
func (mem *CNCatalogTxpool) resultClbkInitialMoment(
	tx []byte,
	transferDetails TransferDetails,
	res *iface.Reply,
) {
	switch r := res.Datum.(type) {
	case *iface.Reply_Inspecttrans:
		var submitInspectFault error
		if mem.submitInspect != nil {
			submitInspectFault = mem.submitInspect(tx, r.InspectTransfer)
		}
		if (r.InspectTransfer.Cipher == iface.CipherKindOKAY) && submitInspectFault == nil {
			//
			//
			if err := mem.equalsComplete(len(tx)); err != nil {
				//
				mem.stash.Discard(tx)
				//
				mem.tracer.Diagnose(err.Error())
				mem.telemetry.DeclinedTrans.Add(1)
				return
			}

			//
			if e, ok := mem.transIndex.Load(kinds.Tx(tx).Key()); ok {
				memoryTransfer := e.(*linkedlist.CNComponent).Datum.(*txpoolTransfer)
				memoryTransfer.appendOriginator(transferDetails.OriginatorUUID)
				mem.tracer.Diagnose(
					"REDACTED",
					"REDACTED", kinds.Tx(tx).Digest(),
					"REDACTED", r,
					"REDACTED", mem.altitude.Load(),
					"REDACTED", mem.Extent(),
				)
				mem.telemetry.DeclinedTrans.Add(1)
				return
			}

			memoryTransfer := &txpoolTransfer{
				altitude:    mem.altitude.Load(),
				fuelDesired: r.InspectTransfer.FuelDesired,
				tx:        tx,
			}
			memoryTransfer.appendOriginator(transferDetails.OriginatorUUID)
			mem.appendTransfer(memoryTransfer)
			mem.tracer.Diagnose(
				"REDACTED",
				"REDACTED", kinds.Tx(tx).Digest(),
				"REDACTED", r,
				"REDACTED", mem.altitude.Load(),
				"REDACTED", mem.Extent(),
			)
			mem.alertTransAccessible()
		} else {
			//
			mem.tracer.Diagnose(
				"REDACTED",
				"REDACTED", kinds.Tx(tx).Digest(),
				"REDACTED", transferDetails.OriginatorNodeid,
				"REDACTED", r,
				"REDACTED", submitInspectFault,
			)
			mem.telemetry.UnsuccessfulTrans.Add(1)

			if !mem.settings.RetainUnfitTransInsideStash {
				//
				mem.stash.Discard(tx)
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
func (mem *CNCatalogTxpool) resultClbkReinspect(tx kinds.Tx, res *iface.ReplyInspectTransfer) {
	//
	if !mem.reinspect.locateFollowingRecordAligning(&tx) {
		//
		return
	}

	var submitInspectFault error
	if mem.submitInspect != nil {
		submitInspectFault = mem.submitInspect(tx, res)
	}

	if (res.Cipher != iface.CipherKindOKAY) || submitInspectFault != nil {
		//
		mem.tracer.Diagnose("REDACTED", "REDACTED", tx.Digest(), "REDACTED", res, "REDACTED", submitInspectFault)
		if err := mem.DiscardTransferViaToken(tx.Key()); err != nil {
			mem.tracer.Diagnose("REDACTED", "REDACTED", err)
		}
		if !mem.settings.RetainUnfitTransInsideStash {
			mem.stash.Discard(tx)
			mem.telemetry.ExpelledTrans.Add(1)
		}
	}
}

//
func (mem *CNCatalogTxpool) TransAccessible() <-chan struct{} {
	return mem.transAccessible
}

func (mem *CNCatalogTxpool) alertTransAccessible() {
	if mem.Extent() == 0 {
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
func (mem *CNCatalogTxpool) HarvestMaximumOctetsMaximumFuel(maximumOctets, maximumFuel int64) kinds.Txs {
	mem.reviseMutex.RLock()
	defer mem.reviseMutex.RUnlock()

	var (
		sumFuel    int64
		activeExtent int64
	)

	//
	//
	//
	txs := make([]kinds.Tx, 0, mem.txs.Len())
	for e := mem.txs.Leading(); e != nil; e = e.Following() {
		memoryTransfer := e.Datum.(*txpoolTransfer)

		txs = append(txs, memoryTransfer.tx)

		dataExtent := kinds.CalculateSchemaExtentForeachTrans([]kinds.Tx{memoryTransfer.tx})

		//
		if maximumOctets > -1 && activeExtent+dataExtent > maximumOctets {
			return txs[:len(txs)-1]
		}

		activeExtent += dataExtent

		//
		//
		//
		//
		freshSumFuel := sumFuel + memoryTransfer.fuelDesired
		if maximumFuel > -1 && freshSumFuel > maximumFuel {
			return txs[:len(txs)-1]
		}
		sumFuel = freshSumFuel
	}
	return txs
}

//
func (mem *CNCatalogTxpool) HarvestMaximumTrans(max int) kinds.Txs {
	mem.reviseMutex.RLock()
	defer mem.reviseMutex.RUnlock()

	if max < 0 {
		max = mem.txs.Len()
	}

	txs := make([]kinds.Tx, 0, strongarithmetic.MinimumInteger(mem.txs.Len(), max))
	for e := mem.txs.Leading(); e != nil && len(txs) < max; e = e.Following() {
		memoryTransfer := e.Datum.(*txpoolTransfer)
		txs = append(txs, memoryTransfer.tx)
	}
	return txs
}

//
func (mem *CNCatalogTxpool) Revise(
	altitude int64,
	txs kinds.Txs,
	transferOutcomes []*iface.InvokeTransferOutcome,
	priorInspect PriorInspectMethod,
	submitInspect RelayInspectMethod,
) error {
	mem.tracer.Diagnose("REDACTED", "REDACTED", altitude, "REDACTED", len(txs))

	//
	mem.altitude.Store(altitude)
	mem.alertedTransAccessible.Store(false)

	if priorInspect != nil {
		mem.priorInspect = priorInspect
	}
	if submitInspect != nil {
		mem.submitInspect = submitInspect
	}

	for i, tx := range txs {
		if transferOutcomes[i].Cipher == iface.CipherKindOKAY {
			//
			_ = mem.stash.Propel(tx)
		} else if !mem.settings.RetainUnfitTransInsideStash {
			//
			mem.stash.Discard(tx)
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
		if err := mem.DiscardTransferViaToken(tx.Key()); err != nil {
			mem.tracer.Diagnose("REDACTED",
				"REDACTED", tx.Key(),
				"REDACTED", err.Error())
		}
	}

	//
	if mem.settings.Reinspect {
		mem.reinspectTrans()
	}

	//
	if mem.Extent() > 0 {
		mem.alertTransAccessible()
	}

	//
	mem.telemetry.Extent.Set(float64(mem.Extent()))
	mem.telemetry.ExtentOctets.Set(float64(mem.ExtentOctets()))

	return nil
}

//
//
func (mem *CNCatalogTxpool) reinspectTrans() {
	mem.tracer.Diagnose("REDACTED", "REDACTED", mem.altitude.Load(), "REDACTED", mem.Extent())

	if mem.Extent() <= 0 {
		return
	}

	mem.reinspect.initialize(mem.txs.Leading(), mem.txs.Rear())

	//
	//
	for e := mem.txs.Leading(); e != nil; e = e.Following() {
		tx := e.Datum.(*txpoolTransfer).tx
		mem.reinspect.countAwaitingTrans.Add(1)

		//
		//
		_, err := mem.delegateApplicationLink.InspectTransferAsyncronous(context.TODO(), &iface.SolicitInspectTransfer{
			Tx:   tx,
			Kind: iface.Inspecttranskind_Reinspect,
		})
		if err != nil {
			panic(fmt.Errorf("REDACTED", log.FreshIdleFormat("REDACTED", tx.Digest()), err))
		}
	}

	//
	mem.delegateApplicationLink.Purge(context.TODO())

	//
	//
	select {
	case <-time.After(mem.settings.ReinspectDeadline):
		mem.reinspect.assignComplete()
		mem.tracer.Failure("REDACTED")
	case <-mem.reinspect.completeReexamining():
	}

	if n := mem.reinspect.countAwaitingTrans.Load(); n > 0 {
		mem.tracer.Failure("REDACTED", "REDACTED", n)
	}
	mem.tracer.Diagnose("REDACTED", "REDACTED", mem.altitude.Load(), "REDACTED", mem.Extent())
}

//
//
//
//
//
//
type reinspect struct {
	locator        *linkedlist.CNComponent //
	end           *linkedlist.CNComponent //
	completeChnl        chan struct{}   //
	countAwaitingTrans atomic.Int32    //
	equalsReexamining  atomic.Bool     //
	reinspectComplete   atomic.Bool     //
}

func freshReinspect() *reinspect {
	return &reinspect{
		completeChnl: make(chan struct{}, 1),
	}
}

func (rc *reinspect) initialize(initial, final *linkedlist.CNComponent) {
	if !rc.complete() {
		panic("REDACTED")
	}
	rc.locator = initial
	rc.end = final
	rc.countAwaitingTrans.Store(0)
	rc.equalsReexamining.Store(true)
}

//
//
func (rc *reinspect) complete() bool {
	return !rc.equalsReexamining.Load()
}

//
func (rc *reinspect) assignComplete() {
	rc.locator = nil
	rc.reinspectComplete.Store(false)
	rc.equalsReexamining.Store(false)
}

//
func (rc *reinspect) assignFollowingRecord() {
	rc.locator = rc.locator.Following()
}

//
//
func (rc *reinspect) attemptConclude() bool {
	if rc.locator == rc.end {
		//
		rc.assignComplete()
	}
	if rc.complete() {
		//
		select {
		case rc.completeChnl <- struct{}{}:
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
func (rc *reinspect) locateFollowingRecordAligning(tx *kinds.Tx) bool {
	detected := false
	for ; !rc.complete(); rc.assignFollowingRecord() {
		anticipatedTransfer := rc.locator.Datum.(*txpoolTransfer).tx
		if bytes.Equal(*tx, anticipatedTransfer) {
			//
			detected = true
			rc.countAwaitingTrans.Add(-1)
			break
		}
	}

	if !rc.attemptConclude() {
		//
		rc.assignFollowingRecord()
	}
	return detected
}

//
func (rc *reinspect) completeReexamining() <-chan struct{} {
	return rc.completeChnl
}

//
//
func (rc *reinspect) assignReinspectComplete() bool {
	reexamining := !rc.complete()
	reinspectComplete := rc.reinspectComplete.Swap(reexamining)
	return reexamining != reinspectComplete
}

//
//
func (rc *reinspect) deemedComplete() bool {
	return rc.reinspectComplete.Load()
}
