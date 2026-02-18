package transordinal

import (
	"context"

	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/kinds"
)

//

const (
	enrollee = "REDACTED"
)

//
//
type OrdinalerDaemon struct {
	daemon.RootDaemon

	transferOrdinalr           TransOrdinaler
	ledgerOrdinalr        ordinaler.LedgerOrdinaler
	eventBus         *kinds.EventBus
	concludeOnFault bool
}

//
func NewOrdinalerDaemon(
	transferOrdinalr TransOrdinaler,
	ledgerOrdinalr ordinaler.LedgerOrdinaler,
	eventBus *kinds.EventBus,
	concludeOnFault bool,
) *OrdinalerDaemon {

	is := &OrdinalerDaemon{transferOrdinalr: transferOrdinalr, ledgerOrdinalr: ledgerOrdinalr, eventBus: eventBus, concludeOnFault: concludeOnFault}
	is.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", is)
	return is
}

//
//
func (is *OrdinalerDaemon) OnBegin() error {
	//
	//
	//
	ledgerSubtract, err := is.eventBus.EnrolUnbuffered(
		context.Background(),
		enrollee,
		kinds.EventInquireNewLedgerEvents)
	if err != nil {
		return err
	}

	transSubtract, err := is.eventBus.EnrolUnbuffered(context.Background(), enrollee, kinds.EventInquireTransfer)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ledgerSubtract.Revoked():
				return
			case msg := <-ledgerSubtract.Out():
				eventNewLedgerEvents := msg.Data().(kinds.EventDataNewLedgerEvents)
				level := eventNewLedgerEvents.Level
				countTrans := eventNewLedgerEvents.CountTrans

				group := NewGroup(countTrans)

				for i := int64(0); i < countTrans; i++ {
					message2 := <-transSubtract.Out()
					transOutcome := message2.Data().(kinds.EventDataTransfer).TransOutcome

					if err = group.Add(&transOutcome); err != nil {
						is.Tracer.Fault(
							"REDACTED",
							"REDACTED", level,
							"REDACTED", transOutcome.Ordinal,
							"REDACTED", err,
						)

						if is.concludeOnFault {
							if err := is.Halt(); err != nil {
								is.Tracer.Fault("REDACTED", "REDACTED", err)
							}
							return
						}
					}
				}

				if err := is.ledgerOrdinalr.Ordinal(eventNewLedgerEvents); err != nil {
					is.Tracer.Fault("REDACTED", "REDACTED", level, "REDACTED", err)
					if is.concludeOnFault {
						if err := is.Halt(); err != nil {
							is.Tracer.Fault("REDACTED", "REDACTED", err)
						}
						return
					}
				} else {
					is.Tracer.Details("REDACTED", "REDACTED", level)
				}

				if err = is.transferOrdinalr.AppendGroup(group); err != nil {
					is.Tracer.Fault("REDACTED", "REDACTED", level, "REDACTED", err)
					if is.concludeOnFault {
						if err := is.Halt(); err != nil {
							is.Tracer.Fault("REDACTED", "REDACTED", err)
						}
						return
					}
				} else {
					is.Tracer.Diagnose("REDACTED", "REDACTED", level, "REDACTED", countTrans)
				}
			}
		}
	}()
	return nil
}

//
func (is *OrdinalerDaemon) OnHalt() {
	if is.eventBus.IsActive() {
		_ = is.eventBus.DeenrollAll(context.Background(), enrollee)
	}
}
