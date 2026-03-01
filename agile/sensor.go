package agile

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
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
func (c *Customer) senseDeviation(ctx context.Context, leadingLogging []*kinds.AgileLedger, now time.Time) error {
	if len(leadingLogging) < 2 {
		return errors.New("REDACTED")
	}
	var (
		headlineAligned      bool
		finalAttestedLedger  = leadingLogging[len(leadingLogging)-1]
		finalAttestedHeadline = finalAttestedLedger.NotatedHeading
		attestorsTowardDiscard  = make([]int, 0)
	)
	c.tracer.Diagnose("REDACTED", "REDACTED", finalAttestedHeadline.Altitude,
		"REDACTED", finalAttestedHeadline.Digest, "REDACTED", len(leadingLogging))

	c.supplierExclusion.Lock()
	defer c.supplierExclusion.Unlock()

	if len(c.attestors) == 0 {
		return FaultNegativeAttestors
	}

	//
	//
	faultchnl := make(chan error, len(c.attestors))
	for i, attestor := range c.attestors {
		go c.contrastFreshAgileLedgerUsingAttestor(ctx, faultchnl, finalAttestedLedger, attestor, i)
	}

	//
	for i := 0; i < cap(faultchnl); i++ {
		err := <-faultchnl

		switch e := err.(type) {
		case nil: //
			headlineAligned = true
		case FaultDiscordantHeadings:
			//
			//
			//
			//
			//
			//
			err := c.processDiscordantHeadings(ctx, leadingLogging, e.Ledger, e.AttestorPosition, now)
			if err != nil {
				//
				return err
			}
			//
			attestorsTowardDiscard = append(attestorsTowardDiscard, e.AttestorPosition)

		case faultFlawedAttestor:
			//
			//
			c.tracer.Details("REDACTED",
				"REDACTED", c.attestors[e.AttestorPosition], "REDACTED", err)
			attestorsTowardDiscard = append(attestorsTowardDiscard, e.AttestorPosition)
		case FaultNominatorUrgenciesDeviate:
			c.tracer.Details("REDACTED",
				"REDACTED", c.attestors[e.AttestorPosition], "REDACTED", err)
			return e
		default:
			//
			//
			if errors.Is(e, context.Canceled) || errors.Is(e, context.DeadlineExceeded) {
				return e
			}
			c.tracer.Details("REDACTED", "REDACTED", err)
		}
	}

	//
	if err := c.discardAttestors(attestorsTowardDiscard); err != nil {
		return err
	}

	//
	//
	if headlineAligned {
		return nil
	}

	//
	return FaultUnsuccessfulHeadlineIntersectAlluding
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
func (c *Customer) contrastFreshAgileLedgerUsingAttestor(ctx context.Context, faultchnl chan error, l *kinds.AgileLedger,
	attestor supplier.Supplier, attestorPosition int,
) {
	h := l.NotatedHeading

	agileLedger, err := attestor.AgileLedger(ctx, h.Altitude)
	switch err {
	//
	case nil:
		break

	//
	//
	case supplier.FaultNegativeReply, supplier.FaultAgileLedgerNegationDetected, context.DeadlineExceeded, context.Canceled:
		faultchnl <- err
		return

	//
	//
	//
	//
	case supplier.FaultAltitudeExcessivelyTall:
		//
		var equalsObjectiveAltitude bool
		equalsObjectiveAltitude, agileLedger, err = c.fetchObjectiveLedgerEitherNewest(ctx, h.Altitude, attestor)
		if err != nil {
			faultchnl <- err
			return
		}

		//
		//
		if equalsObjectiveAltitude {
			break
		}

		//
		//
		if !agileLedger.Moment.Before(h.Moment) {
			faultchnl <- FaultDiscordantHeadings{Ledger: agileLedger, AttestorPosition: attestorPosition}
			return
		}

		//
		//
		//
		//
		time.Sleep(2*c.maximumTimerDeviation + c.maximumLedgerDelay)
		equalsObjectiveAltitude, agileLedger, err = c.fetchObjectiveLedgerEitherNewest(ctx, h.Altitude, attestor)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				faultchnl <- err
			} else {
				faultchnl <- faultFlawedAttestor{Rationale: err, AttestorPosition: attestorPosition}
			}
			return
		}
		if equalsObjectiveAltitude {
			break
		}

		//
		//
		if !agileLedger.Moment.Before(h.Moment) {
			faultchnl <- FaultDiscordantHeadings{Ledger: agileLedger, AttestorPosition: attestorPosition}
			return
		}

		//
		//
		//
		//
		//
		//
		//
		faultchnl <- supplier.FaultNegativeReply
		return

	default:
		//
		//
		faultchnl <- faultFlawedAttestor{Rationale: err, AttestorPosition: attestorPosition}
		return
	}

	if !bytes.Equal(h.Digest(), agileLedger.Digest()) {
		faultchnl <- FaultDiscordantHeadings{Ledger: agileLedger, AttestorPosition: attestorPosition}
	}

	//
	desired, got := l.AssessorAssign.NominatorUrgencyDigest(), agileLedger.AssessorAssign.NominatorUrgencyDigest()
	if !bytes.Equal(desired, got) {
		faultchnl <- FaultNominatorUrgenciesDeviate{AttestorDigest: got, AttestorPosition: attestorPosition, LeadingDigest: desired}
	}

	c.tracer.Diagnose("REDACTED", "REDACTED", h.Altitude, "REDACTED", attestorPosition)
	faultchnl <- nil
}

//
func (c *Customer) transmitProof(ctx context.Context, ev *kinds.AgileCustomerOnslaughtProof, acceptor supplier.Supplier) {
	err := acceptor.NotifyProof(ctx, ev)
	if err != nil {
		c.tracer.Failure("REDACTED", "REDACTED", ev, "REDACTED", acceptor)
	}
}

//
//
func (c *Customer) processDiscordantHeadings(
	ctx context.Context,
	leadingLogging []*kinds.AgileLedger,
	disputingLedger *kinds.AgileLedger,
	attestorPosition int,
	now time.Time,
) error {
	endorsingAttestor := c.attestors[attestorPosition]
	attestorLogging, leadingLedger, err := c.scrutinizeDiscordantHeadlineVersusLogging(
		ctx,
		leadingLogging,
		disputingLedger,
		endorsingAttestor,
		now,
	)
	if err != nil {
		c.tracer.Details("REDACTED", "REDACTED", endorsingAttestor, "REDACTED", err)
		return nil
	}

	//
	//
	sharedLedger, reliableLedger := attestorLogging[0], attestorLogging[len(attestorLogging)-1]
	proofVersusLeading := freshAgileCustomerOnslaughtProof(leadingLedger, reliableLedger, sharedLedger)
	c.tracer.Failure("REDACTED", "REDACTED", proofVersusLeading,
		"REDACTED", c.leading, "REDACTED", endorsingAttestor)
	c.transmitProof(ctx, proofVersusLeading, endorsingAttestor)

	if leadingLedger.Endorse.Iteration != attestorLogging[len(attestorLogging)-1].Endorse.Iteration {
		c.tracer.Details("REDACTED" +
			"REDACTED" +
			"REDACTED")
	}

	//
	//
	//
	leadingLogging, attestorLedger, err := c.scrutinizeDiscordantHeadlineVersusLogging(
		ctx,
		attestorLogging,
		leadingLedger,
		c.leading,
		now,
	)
	if err != nil {
		c.tracer.Details("REDACTED", "REDACTED", c.leading, "REDACTED", err)
		return FaultAgileCustomerOnslaught
	}

	//
	sharedLedger, reliableLedger = leadingLogging[0], leadingLogging[len(leadingLogging)-1]
	proofVersusAttestor := freshAgileCustomerOnslaughtProof(attestorLedger, reliableLedger, sharedLedger)
	c.tracer.Failure("REDACTED", "REDACTED", proofVersusAttestor,
		"REDACTED", c.leading, "REDACTED", endorsingAttestor)
	c.transmitProof(ctx, proofVersusAttestor, c.leading)
	//
	return FaultAgileCustomerOnslaught
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
//
//
//
//
func (c *Customer) scrutinizeDiscordantHeadlineVersusLogging(
	ctx context.Context,
	logging []*kinds.AgileLedger,
	objectiveLedger *kinds.AgileLedger,
	origin supplier.Supplier, now time.Time,
) ([]*kinds.AgileLedger, *kinds.AgileLedger, error) {
	var (
		formerlyAttestedLedger, originLedger *kinds.AgileLedger
		originLogging                          []*kinds.AgileLedger
		err                                  error
	)

	if objectiveLedger.Altitude < logging[0].Altitude {
		return nil, nil, fmt.Errorf("REDACTED",
			objectiveLedger.Altitude, logging[0].Altitude)
	}

	for idx, loggingLedger := range logging {
		//
		//
		if loggingLedger.Altitude > objectiveLedger.Altitude {
			//
			//
			//
			//
			if loggingLedger.Moment.After(objectiveLedger.Moment) {
				return nil, nil,
					errors.New("REDACTED")
			}

			//
			//
			if formerlyAttestedLedger.Altitude != objectiveLedger.Altitude {
				originLogging, err = c.validateOmitting(ctx, origin, formerlyAttestedLedger, objectiveLedger, now)
				if err != nil {
					return nil, nil, fmt.Errorf("REDACTED", err)
				}
			}
			return originLogging, loggingLedger, nil
		}

		//
		if loggingLedger.Altitude == objectiveLedger.Altitude {
			originLedger = objectiveLedger
		} else {
			originLedger, err = origin.AgileLedger(ctx, loggingLedger.Altitude)
			if err != nil {
				return nil, nil, fmt.Errorf("REDACTED", err)
			}
		}

		//
		//
		if idx == 0 {
			if ydigest, xdigest := originLedger.Digest(), loggingLedger.Digest(); !bytes.Equal(ydigest, xdigest) {
				return nil, nil, fmt.Errorf("REDACTED",
					xdigest, ydigest)
			}
			formerlyAttestedLedger = originLedger
			continue
		}

		//
		//
		originLogging, err = c.validateOmitting(ctx, origin, formerlyAttestedLedger, originLedger, now)
		if err != nil {
			return nil, nil, fmt.Errorf("REDACTED", err)
		}
		//
		if ydigest, xdigest := originLedger.Digest(), loggingLedger.Digest(); !bytes.Equal(ydigest, xdigest) {
			//
			return originLogging, loggingLedger, nil
		}

		//
		formerlyAttestedLedger = originLedger
	}

	//
	//
	//
	return nil, nil, faultNegativeDeviation
}

//
//
//
func (c *Customer) fetchObjectiveLedgerEitherNewest(
	ctx context.Context,
	altitude int64,
	attestor supplier.Supplier,
) (bool, *kinds.AgileLedger, error) {
	agileLedger, err := attestor.AgileLedger(ctx, 0)
	if err != nil {
		return false, nil, err
	}

	if agileLedger.Altitude == altitude {
		//
		//
		return true, agileLedger, nil
	}

	if agileLedger.Altitude > altitude {
		//
		//
		//
		agileLedger, err := attestor.AgileLedger(ctx, altitude)
		return true, agileLedger, err
	}

	return false, agileLedger, nil
}

//
//
func freshAgileCustomerOnslaughtProof(contested, reliable, shared *kinds.AgileLedger) *kinds.AgileCustomerOnslaughtProof {
	ev := &kinds.AgileCustomerOnslaughtProof{DiscordantLedger: contested}
	//
	//
	//
	if ev.DiscordantHeadingEqualsUnfit(reliable.Heading) {
		ev.SharedAltitude = shared.Altitude
		ev.Timestamp = shared.Moment
		ev.SumBallotingPotency = shared.AssessorAssign.SumBallotingPotency()
	} else {
		ev.SharedAltitude = reliable.Altitude
		ev.Timestamp = reliable.Moment
		ev.SumBallotingPotency = reliable.AssessorAssign.SumBallotingPotency()
	}
	ev.TreacherousAssessors = ev.ObtainTreacherousAssessors(shared.AssessorAssign, reliable.NotatedHeading)
	return ev
}
