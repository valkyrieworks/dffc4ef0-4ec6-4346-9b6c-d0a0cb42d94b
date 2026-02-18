package privatekey

import (
	"time"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
)

const (
	standardMaximumCallAttempts        = 10
	standardReprocessWaitMillis = 100
)

//
type NotaryDaemonTerminusSetting func(*NotaryCallerTerminus)

//
//
func NotaryCallerTerminusDeadlineScanRecord(deadline time.Duration) NotaryDaemonTerminusSetting {
	return func(ss *NotaryCallerTerminus) { ss.deadlineScanRecord = deadline }
}

//
//
func NotaryCallerTerminusLinkAttempts(attempts int) NotaryDaemonTerminusSetting {
	return func(ss *NotaryCallerTerminus) { ss.maximumLinkAttempts = attempts }
}

//
//
func NotaryCallerTerminusReprocessWaitCadence(cadence time.Duration) NotaryDaemonTerminusSetting {
	return func(ss *NotaryCallerTerminus) { ss.reprocessWait = cadence }
}

//
//
type NotaryCallerTerminus struct {
	notaryTerminus

	caller SocketCaller

	reprocessWait      time.Duration
	maximumLinkAttempts int
}

//
//
//
func NewNotaryCallerGateway(
	tracer log.Tracer,
	caller SocketCaller,
	options ...NotaryDaemonTerminusSetting,
) *NotaryCallerTerminus {
	sd := &NotaryCallerTerminus{
		caller:         caller,
		reprocessWait:      standardReprocessWaitMillis * time.Millisecond,
		maximumLinkAttempts: standardMaximumCallAttempts,
	}

	sd.RootDaemon = *daemon.NewRootDaemon(tracer, "REDACTED", sd)
	sd.deadlineScanRecord = standardDeadlineScanRecordMoments * time.Second

	for _, settingFunction := range options {
		settingFunction(sd)
	}

	return sd
}

func (sd *NotaryCallerTerminus) assureLinkage() error {
	if sd.IsLinked() {
		return nil
	}

	attempts := 0
	for attempts < sd.maximumLinkAttempts {
		link, err := sd.caller()

		if err != nil {
			attempts++
			sd.Tracer.Diagnose("REDACTED", "REDACTED", attempts, "REDACTED", sd.maximumLinkAttempts, "REDACTED", err)
			//
			time.Sleep(sd.reprocessWait)
		} else {
			sd.CollectionLinkage(link)
			sd.Tracer.Diagnose("REDACTED")
			return nil
		}
	}

	sd.Tracer.Diagnose("REDACTED", "REDACTED", attempts, "REDACTED", sd.maximumLinkAttempts)

	return ErrNoLinkage
}
