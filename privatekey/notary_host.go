package privatekey

import (
	"io"

	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	privatekeyproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
	"github.com/valkyrieworks/kinds"
)

//
type VerificationQueryManagerFunction func(
	privateValue kinds.PrivateRatifier,
	querySignal privatekeyproto.Signal,
	ledgerUID string) (privatekeyproto.Signal, error)

type NotaryHost struct {
	daemon.RootDaemon

	gateway *NotaryCallerTerminus
	ledgerUID  string
	privateValue  kinds.PrivateRatifier

	managerMutex               engineconnect.Lock
	verificationQueryManager VerificationQueryManagerFunction
}

func NewNotaryHost(gateway *NotaryCallerTerminus, ledgerUID string, privateValue kinds.PrivateRatifier) *NotaryHost {
	ss := &NotaryHost{
		gateway:                 gateway,
		ledgerUID:                  ledgerUID,
		privateValue:                  privateValue,
		verificationQueryManager: StandardVerificationQueryManager,
	}

	ss.RootDaemon = *daemon.NewRootDaemon(gateway.Tracer, "REDACTED", ss)

	return ss
}

//
func (ss *NotaryHost) OnBegin() error {
	go ss.daemonCycle()
	return nil
}

//
func (ss *NotaryHost) OnHalt() {
	ss.gateway.Tracer.Diagnose("REDACTED")
	_ = ss.gateway.End()
}

//
func (ss *NotaryHost) CollectionQueryManager(verificationQueryManager VerificationQueryManagerFunction) {
	ss.managerMutex.Lock()
	defer ss.managerMutex.Unlock()
	ss.verificationQueryManager = verificationQueryManager
}

func (ss *NotaryHost) daemonAwaitingQuery() {
	if !ss.IsActive() {
		return //
	}

	req, err := ss.gateway.ScanSignal()
	if err != nil {
		if err != io.EOF {
			ss.Tracer.Fault("REDACTED", "REDACTED", err)
		}
		return
	}

	var res privatekeyproto.Signal
	{
		//
		ss.managerMutex.Lock()
		defer ss.managerMutex.Unlock()
		res, err = ss.verificationQueryManager(ss.privateValue, req, ss.ledgerUID)
		if err != nil {
			//
			ss.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}

	err = ss.gateway.RecordSignal(res)
	if err != nil {
		ss.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

func (ss *NotaryHost) daemonCycle() {
	for {
		select {
		default:
			err := ss.gateway.assureLinkage()
			if err != nil {
				return
			}
			ss.daemonAwaitingQuery()

		case <-ss.Exit():
			return
		}
	}
}
