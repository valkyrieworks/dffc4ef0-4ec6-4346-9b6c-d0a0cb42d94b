package privatevalue

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	privatevalueschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
)

//
type EndorserObserverGatewaySelection func(*EndorserObserverGateway)

//
//
//
//
func EndorserObserverGatewayDeadlineRetrievePersist(deadline time.Duration) EndorserObserverGatewaySelection {
	return func(sl *EndorserObserverGateway) { sl.deadlineRetrievePersist = deadline }
}

//
//
//
//
//
type EndorserObserverGateway struct {
	endorserGateway

	observer              net.Listener
	relateSolicitChnl      chan struct{}
	linkageAccessibleChnl chan net.Conn

	deadlineEmbrace   time.Duration
	embraceMishapTally atomic.Uint32
	pingClock       *time.Ticker
	pingDuration    time.Duration

	replicaMutex commitchronize.Exclusion //
}

//
func FreshEndorserObserverGateway(
	tracer log.Tracer,
	observer net.Listener,
	choices ...EndorserObserverGatewaySelection,
) *EndorserObserverGateway {
	sl := &EndorserObserverGateway{
		observer:      observer,
		deadlineEmbrace: fallbackDeadlineEmbraceMoments * time.Second,
	}

	sl.FoundationFacility = *facility.FreshFoundationFacility(tracer, "REDACTED", sl)
	sl.deadlineRetrievePersist = fallbackDeadlineRetrievePersistMoments * time.Second

	for _, selectionMethod := range choices {
		selectionMethod(sl)
	}

	return sl
}

//
func (sl *EndorserObserverGateway) UponInitiate() error {
	sl.relateSolicitChnl = make(chan struct{}, 1) //
	sl.linkageAccessibleChnl = make(chan net.Conn)

	//
	sl.pingDuration = time.Duration(sl.deadlineRetrievePersist.Milliseconds()*2/3) * time.Millisecond
	sl.pingClock = time.NewTicker(sl.pingDuration)

	go sl.facilityCycle()
	go sl.pingCycle()

	sl.relateSolicitChnl <- struct{}{}

	return nil
}

//
func (sl *EndorserObserverGateway) UponHalt() {
	sl.replicaMutex.Lock()
	defer sl.replicaMutex.Unlock()
	_ = sl.Shutdown()

	//
	if sl.observer != nil {
		if err := sl.observer.Close(); err != nil {
			sl.Tracer.Failure("REDACTED", "REDACTED", err)
			sl.observer = nil
		}
	}

	sl.pingClock.Stop()
}

//
func (sl *EndorserObserverGateway) PauseForeachLinkage(maximumPause time.Duration) error {
	sl.replicaMutex.Lock()
	defer sl.replicaMutex.Unlock()
	return sl.assureLinkage(maximumPause)
}

//
func (sl *EndorserObserverGateway) TransmitSolicit(solicit privatevalueschema.Signal) (*privatevalueschema.Signal, error) {
	sl.replicaMutex.Lock()
	defer sl.replicaMutex.Unlock()

	err := sl.assureLinkage(sl.deadlineEmbrace)
	if err != nil {
		return nil, err
	}

	err = sl.PersistArtifact(solicit)
	if err != nil {
		return nil, err
	}

	res, err := sl.FetchArtifact()
	if err != nil {
		return nil, err
	}

	//
	sl.pingClock.Reset(sl.pingDuration)

	return &res, nil
}

func (sl *EndorserObserverGateway) assureLinkage(maximumPause time.Duration) error {
	if sl.EqualsAssociated() {
		return nil
	}

	//
	if sl.ObtainAccessibleLinkage(sl.linkageAccessibleChnl) {
		return nil
	}

	//
	sl.Tracer.Details("REDACTED")
	sl.activateRelate()
	err := sl.PauseLinkage(sl.linkageAccessibleChnl, maximumPause)
	if err != nil {
		return err
	}

	return nil
}

func (sl *EndorserObserverGateway) embraceFreshLinkage() (net.Conn, error) {
	if !sl.EqualsActive() || sl.observer == nil {
		return nil, fmt.Errorf("REDACTED")
	}

	//
	sl.Tracer.Details("REDACTED")
	link, err := sl.observer.Accept()
	if err != nil {
		sl.embraceMishapTally.Add(1)
		return nil, err
	}

	sl.embraceMishapTally.Store(0)
	return link, nil
}

func (sl *EndorserObserverGateway) activateRelate() {
	select {
	case sl.relateSolicitChnl <- struct{}{}:
	default:
	}
}

func (sl *EndorserObserverGateway) activateReestablish() {
	sl.DiscardLinkage()
	sl.activateRelate()
}

func (sl *EndorserObserverGateway) facilityCycle() {
	for {
		select {
		case <-sl.relateSolicitChnl:
			//
			//
			if sl.EqualsAssociated() {
				sl.Tracer.Diagnose("REDACTED")
				continue
			}

			//
			link, err := sl.embraceFreshLinkage()
			if err != nil {
				sl.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", sl.embraceMishapTally.Load())
				sl.activateRelate()
				continue
			}

			//
			sl.Tracer.Details("REDACTED")
			select {
			case sl.linkageAccessibleChnl <- link:
			case <-sl.Exit():
				return
			}
		case <-sl.Exit():
			return
		}
	}
}

func (sl *EndorserObserverGateway) pingCycle() {
	for {
		select {
		case <-sl.pingClock.C:
			{
				_, err := sl.TransmitSolicit(shouldEncloseSignal(&privatevalueschema.PingSolicit{}))
				if err != nil {
					sl.Tracer.Failure("REDACTED")
					sl.activateReestablish()
				}
			}
		case <-sl.Exit():
			return
		}
	}
}
