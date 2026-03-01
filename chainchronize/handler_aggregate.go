package chainchronize

import (
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type LedgerAbsorber interface {
	AbsorbAttestedLedger(ledgerNominee agreement.AbsorbNominee) error
}

func (r *Handler) obtainLedgerAbsorber() (LedgerAbsorber, error) {
	cr, ok := r.Router.Handler("REDACTED")
	if !ok {
		return nil, fmt.Errorf("REDACTED")
	}

	bi, ok := cr.(LedgerAbsorber)
	if !ok {
		return nil, fmt.Errorf("REDACTED", cr)
	}

	return bi, nil
}

//
//
//
//
func (r *Handler) ledgerAbsorberProcedure(ledgerAbsorber LedgerAbsorber) {
	r.Tracer.Details("REDACTED")

	attemptChronizeMetronome := time.NewTicker(durationAttemptChronize)
	defer attemptChronizeMetronome.Stop()

	chronizeRepetitionStream := make(chan struct{}, 1)
	defer close(chronizeRepetitionStream)

	for {
		select {
		case <-r.Exit():
			return
		case <-r.hub.Exit():
			return
		case <-attemptChronizeMetronome.C:
			select {
			case chronizeRepetitionStream <- struct{}{}:
			default:
				//
			}
		case <-chronizeRepetitionStream:
			//
			//
			ledger, followingLedger, addnEndorse := r.hub.GlanceCoupleLedgers()
			if ledger == nil || followingLedger == nil {
				continue
			}

			//
			if ledger.Altitude+1 != followingLedger.Altitude {
				panic(fmt.Errorf(
					"REDACTED",
					ledger.Altitude+1,
					followingLedger.Altitude,
				))
			}

			//
			//
			status, err := r.ledgerExecute.Depot().Fetch()
			if err != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", err)
				return
			}

			newestAltitude := status.FinalLedgerAltitude

			//
			//
			if ledger.Altitude <= newestAltitude {
				r.hub.ExtractSolicit()
				r.telemetry.EarlierComprisedLedgers.Add(1)

				r.Tracer.Diagnose(
					"REDACTED",
					"REDACTED", ledger.Altitude,
					"REDACTED", newestAltitude,
				)

				continue
			}

			//
			//
			if ledger.Altitude != newestAltitude+1 {
				panic(fmt.Errorf(
					"REDACTED",
					ledger.Altitude,
					newestAltitude+1,
				))
			}

			if !r.EqualsActive() || !r.hub.EqualsActive() {
				return
			}

			//
			chronizeRepetitionStream <- struct{}{}

			ledgerFragments, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
			if err != nil {
				//
				r.Tracer.Failure("REDACTED", "REDACTED", ledger.Altitude, "REDACTED", err)
				return
			}

			//
			ic, err := agreement.FreshAbsorbNominee(ledger, ledgerFragments, followingLedger.FinalEndorse, addnEndorse)
			if err != nil {
				r.processCertificationBreakdown(ledger, followingLedger, fmt.Errorf("REDACTED", err))
				continue
			}

			//
			if err := ic.Validate(status); err != nil {
				r.processCertificationBreakdown(ledger, followingLedger, fmt.Errorf("REDACTED", err))
				continue
			}

			//
			r.hub.ExtractSolicit()

			//
			//
			initiate := time.Now()
			err = ledgerAbsorber.AbsorbAttestedLedger(ic)
			passed := time.Since(initiate)

			switch {
			case errors.Is(err, agreement.FaultEarlierComprised):
				r.Tracer.Details("REDACTED", "REDACTED", ledger.Altitude)
				r.telemetry.EarlierComprisedLedgers.Add(1)
			case err != nil:
				//
				//
				//
				r.Tracer.Failure("REDACTED", "REDACTED", ledger.Altitude, "REDACTED", err)
				return
			default:
				r.telemetry.logLedgerTelemetry(ledger)
				r.telemetry.AbsorbedLedgers.Add(1)
				r.telemetry.AbsorbedLedgerInterval.Observe(passed.Seconds())
			}
		}
	}
}
