package privatekey

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	privatekeyproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
)

//
type NotaryObserverTerminusSetting func(*NotaryObserverTerminus)

//
//
//
//
func NotaryObserverTerminusDeadlineScanRecord(deadline time.Duration) NotaryObserverTerminusSetting {
	return func(sl *NotaryObserverTerminus) { sl.deadlineScanRecord = deadline }
}

//
//
//
//
//
type NotaryObserverTerminus struct {
	notaryTerminus

	observer              net.Listener
	joinQueryChan      chan struct{}
	linkageAccessibleChan chan net.Conn

	deadlineAllow   time.Duration
	allowAbortNumber atomic.Uint32
	pingClock       *time.Ticker
	pingCadence    time.Duration

	occurrenceMutex engineconnect.Lock //
}

//
func NewNotaryObserverTerminus(
	tracer log.Tracer,
	observer net.Listener,
	options ...NotaryObserverTerminusSetting,
) *NotaryObserverTerminus {
	sl := &NotaryObserverTerminus{
		observer:      observer,
		deadlineAllow: standardDeadlineAllowMoments * time.Second,
	}

	sl.RootDaemon = *daemon.NewRootDaemon(tracer, "REDACTED", sl)
	sl.deadlineScanRecord = standardDeadlineScanRecordMoments * time.Second

	for _, settingFunction := range options {
		settingFunction(sl)
	}

	return sl
}

//
func (sl *NotaryObserverTerminus) OnBegin() error {
	sl.joinQueryChan = make(chan struct{}, 1) //
	sl.linkageAccessibleChan = make(chan net.Conn)

	//
	sl.pingCadence = time.Duration(sl.deadlineScanRecord.Milliseconds()*2/3) * time.Millisecond
	sl.pingClock = time.NewTicker(sl.pingCadence)

	go sl.daemonCycle()
	go sl.pingCycle()

	sl.joinQueryChan <- struct{}{}

	return nil
}

//
func (sl *NotaryObserverTerminus) OnHalt() {
	sl.occurrenceMutex.Lock()
	defer sl.occurrenceMutex.Unlock()
	_ = sl.End()

	//
	if sl.observer != nil {
		if err := sl.observer.Close(); err != nil {
			sl.Tracer.Fault("REDACTED", "REDACTED", err)
			sl.observer = nil
		}
	}

	sl.pingClock.Stop()
}

//
func (sl *NotaryObserverTerminus) WaitForLinkage(maximumWait time.Duration) error {
	sl.occurrenceMutex.Lock()
	defer sl.occurrenceMutex.Unlock()
	return sl.assureLinkage(maximumWait)
}

//
func (sl *NotaryObserverTerminus) TransmitQuery(query privatekeyproto.Signal) (*privatekeyproto.Signal, error) {
	sl.occurrenceMutex.Lock()
	defer sl.occurrenceMutex.Unlock()

	err := sl.assureLinkage(sl.deadlineAllow)
	if err != nil {
		return nil, err
	}

	err = sl.RecordSignal(query)
	if err != nil {
		return nil, err
	}

	res, err := sl.ScanSignal()
	if err != nil {
		return nil, err
	}

	//
	sl.pingClock.Reset(sl.pingCadence)

	return &res, nil
}

func (sl *NotaryObserverTerminus) assureLinkage(maximumWait time.Duration) error {
	if sl.IsLinked() {
		return nil
	}

	//
	if sl.FetchAccessibleLinkage(sl.linkageAccessibleChan) {
		return nil
	}

	//
	sl.Tracer.Details("REDACTED")
	sl.activateJoin()
	err := sl.WaitLinkage(sl.linkageAccessibleChan, maximumWait)
	if err != nil {
		return err
	}

	return nil
}

func (sl *NotaryObserverTerminus) allowNewLinkage() (net.Conn, error) {
	if !sl.IsActive() || sl.observer == nil {
		return nil, fmt.Errorf("REDACTED")
	}

	//
	sl.Tracer.Details("REDACTED")
	link, err := sl.observer.Accept()
	if err != nil {
		sl.allowAbortNumber.Add(1)
		return nil, err
	}

	sl.allowAbortNumber.Store(0)
	return link, nil
}

func (sl *NotaryObserverTerminus) activateJoin() {
	select {
	case sl.joinQueryChan <- struct{}{}:
	default:
	}
}

func (sl *NotaryObserverTerminus) activateReestablish() {
	sl.DiscardLinkage()
	sl.activateJoin()
}

func (sl *NotaryObserverTerminus) daemonCycle() {
	for {
		select {
		case <-sl.joinQueryChan:
			//
			//
			if sl.IsLinked() {
				sl.Tracer.Diagnose("REDACTED")
				continue
			}

			//
			link, err := sl.allowNewLinkage()
			if err != nil {
				sl.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", sl.allowAbortNumber.Load())
				sl.activateJoin()
				continue
			}

			//
			sl.Tracer.Details("REDACTED")
			select {
			case sl.linkageAccessibleChan <- link:
			case <-sl.Exit():
				return
			}
		case <-sl.Exit():
			return
		}
	}
}

func (sl *NotaryObserverTerminus) pingCycle() {
	for {
		select {
		case <-sl.pingClock.C:
			{
				_, err := sl.TransmitQuery(shouldEncloseMessage(&privatekeyproto.PingQuery{}))
				if err != nil {
					sl.Tracer.Fault("REDACTED")
					sl.activateReestablish()
				}
			}
		case <-sl.Exit():
			return
		}
	}
}
