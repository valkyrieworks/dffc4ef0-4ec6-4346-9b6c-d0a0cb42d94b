package chainconnect

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/netpeer"
	"github.com/valkyrieworks/p2p"
	chainproto "github.com/valkyrieworks/schema/consensuscore/chainconnect"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

//
const ChainconnectStream = byte(0x40)

const (
	standardCadenceStatusModify  = 10 * time.Second
	replicaCadenceStatusModify = 1 * time.Second

	standardCadenceRouterToAgreement = 1 * time.Second

	//
	cadenceAttemptAlign = 10 * time.Millisecond
)

type agreementHandler interface {
	//
	//
	RouterToAgreement(status sm.Status, omitJournal bool)
}

type txpoolHandler interface {
	//
	ActivateInOutTrans()
}

type nodeFault struct {
	err    error
	nodeUID p2p.ID
}

func (e nodeFault) Fault() string {
	return fmt.Sprintf("REDACTED", e.nodeUID, e.err.Error())
}

//
type Handler struct {
	p2p.RootHandler

	//
	primaryStatus sm.Status

	//
	//
	activated *atomic.Bool

	//
	replicaStyle bool

	ledgerExecute     *sm.LedgerRunner
	depot         sm.LedgerDepot
	depository          *LedgerDepository
	nativeAddress     vault.Location
	depositoryProcessGroup sync.WaitGroup

	queriesChan <-chan LedgerQuery
	faultsChan   <-chan nodeFault

	//
	cadenceRouterToAgreement time.Duration

	//
	cadenceStatusModify time.Duration

	stats *Stats
}

//
func NewHandler(
	activated bool,
	replicaStyle bool,
	status sm.Status,
	ledgerExecute *sm.LedgerRunner,
	depot *depot.LedgerDepot,
	nativeAddress vault.Location,
	inactiveStatusAlignLevel int64,
	stats *Stats,
) *Handler {
	depotLevel := depot.Level()
	if depotLevel == 0 {
		//
		//
		//
		//
		//
		//
		depotLevel = inactiveStatusAlignLevel
	}

	if status.FinalLedgerLevel != depotLevel {
		panic(fmt.Sprintf(
			"REDACTED",
			status.FinalLedgerLevel,
			depotLevel,
		))
	}

	//
	//
	queriesChan := make(chan LedgerQuery)

	const volume = 1000                      //
	faultsChan := make(chan nodeFault, volume) //

	beginLevel := depotLevel + 1
	if beginLevel == 1 {
		beginLevel = status.PrimaryLevel
	}
	depository := NewLedgerDepository(beginLevel, queriesChan, faultsChan)

	activatedMark := &atomic.Bool{}
	activatedMark.Store(activated)

	cadenceStatusModify := standardCadenceStatusModify
	if replicaStyle {
		cadenceStatusModify = replicaCadenceStatusModify
	}

	r := &Handler{
		primaryStatus:              status,
		ledgerExecute:                 ledgerExecute,
		depot:                     depot,
		depository:                      depository,
		activated:                   activatedMark,
		replicaStyle:              replicaStyle,
		nativeAddress:                 nativeAddress,
		queriesChan:                queriesChan,
		faultsChan:                  faultsChan,
		stats:                   stats,
		cadenceRouterToAgreement: standardCadenceRouterToAgreement,
		cadenceStatusModify:      cadenceStatusModify,
	}

	r.RootHandler = *p2p.NewRootHandler("REDACTED", r)

	return r
}

//
func (r *Handler) AssignTracer(l log.Tracer) {
	r.Tracer = l
	r.depository.Tracer = l
}

//
func (r *Handler) OnBegin() error {
	//
	if !r.activated.Load() {
		return nil
	}

	return r.executeDepository(false)
}

func (r *Handler) executeDepository(statusAligned bool) error {
	if err := r.depository.Begin(); err != nil {
		return err
	}

	r.depositoryProcessGroup.Add(1)
	go func() {
		defer r.depositoryProcessGroup.Done()
		r.depositoryProcess(statusAligned)
	}()

	return nil
}

//
func (r *Handler) Activate(status sm.Status) error {
	if !r.activated.CompareAndSwap(false, true) {
		return ErrYetActivated
	}

	r.primaryStatus = status
	r.depository.level = status.FinalLedgerLevel + 1

	return r.executeDepository(true)
}

//
func (r *Handler) OnHalt() {
	if !r.activated.Load() {
		return
	}

	if err := r.depository.Halt(); err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	r.depositoryProcessGroup.Wait()
}

//
func (r *Handler) FetchStreams() []*p2p.StreamDefinition {
	return []*p2p.StreamDefinition{
		{
			ID:                  ChainconnectStream,
			Urgency:            5,
			TransmitBufferVolume:   1000,
			AcceptBufferVolume:  50 * 4096,
			AcceptSignalVolume: MaximumMessageVolume,
			SignalKind:         &chainproto.Signal{},
		},
	}
}

//
func (r *Handler) AppendNode(node p2p.Node) {
	node.Transmit(p2p.Packet{
		StreamUID: ChainconnectStream,
		Signal: &chainproto.StatusReply{
			Root:   r.depot.Root(),
			Level: r.depot.Level(),
		},
	})
	//

	//
	//
}

//
func (r *Handler) DeleteNode(node p2p.Node, _ any) {
	r.depository.DeleteNode(node.ID())
}

//
//
func (r *Handler) answerToNode(msg *chainproto.LedgerQuery, src p2p.Node) {
	ledger := r.depot.ImportLedger(msg.Level)
	if ledger == nil {
		r.Tracer.Details("REDACTED", "REDACTED", src, "REDACTED", msg.Level)
		src.AttemptTransmit(p2p.Packet{
			StreamUID: ChainconnectStream,
			Signal:   &chainproto.NoLedgerReply{Level: msg.Level},
		})

		return
	}

	status, err := r.ledgerExecute.Depot().Import()
	if err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", err)
		return
	}

	var extensionEndorse *kinds.ExpandedEndorse
	if status.AgreementOptions.Iface.BallotPluginsActivated(msg.Level) {
		extensionEndorse = r.depot.ImportLedgerExpandedEndorse(msg.Level)
		if extensionEndorse == nil {
			r.Tracer.Fault("REDACTED", "REDACTED", ledger)
			return
		}
	}

	bl, err := ledger.ToSchema()
	if err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", err)
		return
	}

	src.AttemptTransmit(p2p.Packet{
		StreamUID: ChainconnectStream,
		Signal: &chainproto.LedgerReply{
			Ledger:     bl,
			ExtensionEndorse: extensionEndorse.ToSchema(),
		},
	})
}

func (r *Handler) processNodeReply(msg *chainproto.LedgerReply, src p2p.Node) {
	bi, err := kinds.LedgerFromSchema(msg.Ledger)
	if err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", src, "REDACTED", msg, "REDACTED", err)
		r.haltNodeForFault(src, err)
		return
	}

	var extensionEndorse *kinds.ExpandedEndorse
	if msg.ExtensionEndorse != nil {
		extensionEndorse, err = kinds.ExpandedEndorseFromSchema(msg.ExtensionEndorse)
		if err != nil {
			r.Tracer.Fault("REDACTED", "REDACTED", src, "REDACTED", err)
			r.haltNodeForFault(src, err)
			return
		}
	}

	if err := r.depository.AppendLedger(src.ID(), bi, extensionEndorse, msg.Ledger.Volume()); err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", src, "REDACTED", err)
	}
}

//
func (r *Handler) Accept(e p2p.Packet) {
	if err := CertifyMessage(e.Signal); err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		r.haltNodeForFault(e.Src, err)
		return
	}

	r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", e.Signal)

	switch msg := e.Signal.(type) {
	case *chainproto.LedgerQuery:
		//
		r.answerToNode(msg, e.Src)
	case *chainproto.LedgerReply:
		//
		go r.processNodeReply(msg, e.Src)
	case *chainproto.StatusQuery:
		//
		e.Src.AttemptTransmit(p2p.Packet{
			StreamUID: ChainconnectStream,
			Signal: &chainproto.StatusReply{
				Level: r.depot.Level(),
				Root:   r.depot.Root(),
			},
		})
	case *chainproto.StatusReply:
		//
		r.depository.AssignNodeScope(e.Src.ID(), msg.Root, msg.Level)
	case *chainproto.NoLedgerReply:
		r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Level)
		r.depository.ReworkQueryFrom(msg.Level, e.Src.ID())
	default:
		r.Tracer.Fault(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
	}
}

func (r *Handler) nativeMemberLedgersTheLedger(status sm.Status) bool {
	_, val := status.Ratifiers.FetchByLocation(r.nativeAddress)
	if val == nil {
		return false
	}
	sum := status.Ratifiers.SumPollingEnergy()
	return val.PollingEnergy >= sum/3
}

//
//
func (r *Handler) depositoryProcess(statusAligned bool) {
	r.stats.Aligning.Set(1)
	defer r.stats.Aligning.Set(0)

	attemptAlignTimer := time.NewTicker(cadenceAttemptAlign)
	defer attemptAlignTimer.Stop()

	statusModifyTimer := time.NewTicker(r.cadenceStatusModify)
	defer statusModifyTimer.Stop()

	routerToAgreementTimer := time.NewTicker(r.cadenceRouterToAgreement)
	defer routerToAgreementTimer.Stop()

	go r.depositorySignalsProcess(statusModifyTimer)

	var (
		ledgerUID                    = r.primaryStatus.LedgerUID
		status                      = r.primaryStatus
		primaryEndorseHasPlugins = status.FinalLedgerLevel > 0 &&
			r.depot.ImportLedgerExpandedEndorse(status.FinalLedgerLevel) != nil

		didHandleChan = make(chan struct{}, 1)

		//
		ledgersAligned = 0
		finalCentury  = time.Now()
		finalRatio     = 0.0
	)

FOR_CYCLE:
	for {
		select {
		case <-r.Exit():
			break FOR_CYCLE
		case <-r.depository.Exit():
			break FOR_CYCLE
		case <-routerToAgreementTimer.C:
			level, countAwaiting, sizeInquirers := r.depository.FetchStatus()
			outgoing, incoming, _ := r.Router.CountNodes()

			r.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", countAwaiting,
				"REDACTED", sizeInquirers,
				"REDACTED", outgoing,
				"REDACTED", incoming,
				"REDACTED", status.FinalLedgerLevel,
			)

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
			//
			//
			//
			//
			//
			//
			//
			absentAddition := true
			if status.FinalLedgerLevel == 0 ||
				!status.AgreementOptions.Iface.BallotPluginsActivated(status.FinalLedgerLevel) ||
				ledgersAligned > 0 ||
				primaryEndorseHasPlugins {
				absentAddition = false
			}

			//
			if absentAddition {
				r.Tracer.Details(
					"REDACTED",
					"REDACTED", level,
					"REDACTED", status.FinalLedgerLevel,
					"REDACTED", status.PrimaryLevel,
					"REDACTED", r.depository.MaximumNodeLevel(),
				)
				continue FOR_CYCLE
			}

			//
			if !r.depository.IsSeizedUp() && !r.nativeMemberLedgersTheLedger(status) {
				continue FOR_CYCLE
			}

			if r.replicaStyle {
				r.Tracer.Diagnose("REDACTED", "REDACTED", status.FinalLedgerLevel)
				continue FOR_CYCLE
			}

			r.Tracer.Details("REDACTED", "REDACTED", level)
			if err := r.depository.Halt(); err != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", err)
			}

			memoryReader, present := r.Router.Handler("REDACTED")
			if present {
				if memoryReader, ok := memoryReader.(txpoolHandler); ok {
					memoryReader.ActivateInOutTrans()
				}
			}

			connectReader, present := r.Router.Handler("REDACTED")
			if present {
				if connectReader, ok := connectReader.(agreementHandler); ok {
					connectReader.RouterToAgreement(status, ledgersAligned > 0 || statusAligned)
				}
			}

			break FOR_CYCLE
		case <-attemptAlignTimer.C:
			select {
			case didHandleChan <- struct{}{}:
			default:
			}
		case <-didHandleChan:
			//
			//
			//
			//
			//
			//
			//

			//
			initial, moment, extensionEndorse := r.depository.GlanceDualLedgers()
			if initial == nil || moment == nil {
				//
				//
				continue FOR_CYCLE
			}
			//
			if status.FinalLedgerLevel > 0 && status.FinalLedgerLevel+1 != initial.Level {
				//
				panic(fmt.Errorf("REDACTED", status.FinalLedgerLevel+1, initial.Level))
			}
			if initial.Level+1 != moment.Level {
				//
				panic(fmt.Errorf("REDACTED", status.FinalLedgerLevel, initial.Level))
			}

			//
			//
			//
			//
			if !r.IsActive() || !r.depository.IsActive() {
				break FOR_CYCLE
			}
			//
			didHandleChan <- struct{}{}

			initialSegments, err := initial.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
			if err != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", initial.Level, "REDACTED", err.Error())
				break FOR_CYCLE
			}

			initialSegmentAssignHeading := initialSegments.Heading()
			initialUID := kinds.LedgerUID{Digest: initial.Digest(), SegmentAssignHeading: initialSegmentAssignHeading}

			//
			//
			//
			//
			//
			err = status.Ratifiers.ValidateEndorseRapid(ledgerUID, initialUID, initial.Level, moment.FinalEndorse)

			if err == nil {
				//
				err = r.ledgerExecute.CertifyLedger(status, initial)
			}
			existingExtensionEndorse := extensionEndorse != nil
			pluginsActivated := status.AgreementOptions.Iface.BallotPluginsActivated(initial.Level)
			if existingExtensionEndorse != pluginsActivated {
				err = fmt.Errorf("REDACTED"+
					"REDACTED",
					initial.Level, existingExtensionEndorse, pluginsActivated,
				)
			}
			if err == nil && pluginsActivated {
				//
				err = extensionEndorse.AssurePlugins(true)
			}
			if err != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", err)
				nodeUID := r.depository.DeleteNodeAndReworkAllNodeQueries(initial.Level)
				node := r.Router.Nodes().Get(nodeUID)
				if node != nil {
					//
					//
					r.haltNodeForFault(node, ErrHandlerVerification{Err: err})
				}
				nodeUidtwo := r.depository.DeleteNodeAndReworkAllNodeQueries(moment.Level)
				nodetwo := r.Router.Nodes().Get(nodeUidtwo)
				if nodetwo != nil && nodetwo != node {
					//
					//
					r.haltNodeForFault(nodetwo, ErrHandlerVerification{Err: err})
				}
				continue FOR_CYCLE
			}

			r.depository.EjectQuery()

			//
			if pluginsActivated {
				r.depot.PersistLedgerWithExpandedEndorse(initial, initialSegments, extensionEndorse)
			} else {
				//
				//
				//
				//
				r.depot.PersistLedger(initial, initialSegments, moment.FinalEndorse)
			}

			//
			//
			status, err = r.ledgerExecute.ExecuteValidatedLedger(status, initialUID, initial)
			if err != nil {
				//
				panic(fmt.Sprintf("REDACTED", initial.Level, initial.Digest(), err))
			}

			r.stats.logLedgerStats(initial)
			ledgersAligned++

			if ledgersAligned%100 == 0 {
				finalRatio = 0.9*finalRatio + 0.1*(100/time.Since(finalCentury).Seconds())
				finalCentury = time.Now()
				r.Tracer.Details(
					"REDACTED",
					"REDACTED", r.depository.level,
					"REDACTED", r.depository.MaximumNodeLevel(),
					"REDACTED", finalRatio,
				)
			}

			continue FOR_CYCLE
		}
	}
}

func (r *Handler) depositorySignalsProcess(statusModifyTimer *time.Ticker) {
	for {
		select {
		case <-r.Exit():
			return
		case <-r.depository.Exit():
			return
		case query := <-r.queriesChan:
			//
			node := r.Router.Nodes().Get(query.NodeUID)
			if node == nil {
				continue
			}

			buffered := node.AttemptTransmit(p2p.Packet{
				StreamUID: ChainconnectStream,
				Signal:   &chainproto.LedgerQuery{Level: query.Level},
			})

			if !buffered {
				r.Tracer.Diagnose("REDACTED", "REDACTED", node.ID(), "REDACTED", query.Level)
			}
		case err := <-r.faultsChan:
			//
			if node := r.Router.Nodes().Get(err.nodeUID); node != nil {
				r.haltNodeForFault(node, err.err)
			}
		case <-statusModifyTimer.C:
			//
			r.Router.MulticastAsync(p2p.Packet{
				StreamUID: ChainconnectStream,
				Signal:   &chainproto.StatusQuery{},
			})
		}
	}
}

func (r *Handler) haltNodeForFault(node p2p.Node, err error) {
	if r.replicaStyle && mustBeRelined(err) {
		err = &netpeer.FaultTemporary{Err: err}
	}

	r.Router.HaltNodeForFault(node, err)
}

//
//
//
//
func mustBeRelined(err error) bool {
	//
	return errors.Is(err, ErrNodeDeadline)
}
