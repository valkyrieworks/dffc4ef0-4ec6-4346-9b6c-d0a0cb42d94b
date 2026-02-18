package rapid

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/rapid/source"
	"github.com/valkyrieworks/kinds"
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
func (c *Customer) perceiveDeviation(ctx context.Context, leadingTrack []*kinds.RapidLedger, now time.Time) error {
	if len(leadingTrack) < 2 {
		return errors.New("REDACTED")
	}
	var (
		headingAligned      bool
		finalCertifiedLedger  = leadingTrack[len(leadingTrack)-1]
		finalCertifiedHeading = finalCertifiedLedger.AttestedHeading
		attestorsToDelete  = make([]int, 0)
	)
	c.tracer.Diagnose("REDACTED", "REDACTED", finalCertifiedHeading.Level,
		"REDACTED", finalCertifiedHeading.Digest, "REDACTED", len(leadingTrack))

	c.sourceLock.Lock()
	defer c.sourceLock.Unlock()

	if len(c.attestors) == 0 {
		return ErrNoAttestors
	}

	//
	//
	faultc := make(chan error, len(c.attestors))
	for i, attestor := range c.attestors {
		go c.contrastNewRapidLedgerWithAttestor(ctx, faultc, finalCertifiedLedger, attestor, i)
	}

	//
	for i := 0; i < cap(faultc); i++ {
		err := <-faultc

		switch e := err.(type) {
		case nil: //
			headingAligned = true
		case ErrClashingHeadings:
			//
			//
			//
			//
			//
			//
			err := c.processClashingHeadings(ctx, leadingTrack, e.Ledger, e.AttestorOrdinal, now)
			if err != nil {
				//
				return err
			}
			//
			attestorsToDelete = append(attestorsToDelete, e.AttestorOrdinal)

		case errFlawedAttestor:
			//
			//
			c.tracer.Details("REDACTED",
				"REDACTED", c.attestors[e.AttestorOrdinal], "REDACTED", err)
			attestorsToDelete = append(attestorsToDelete, e.AttestorOrdinal)
		case ErrRecommenderUrgenciesDeviate:
			c.tracer.Details("REDACTED",
				"REDACTED", c.attestors[e.AttestorOrdinal], "REDACTED", err)
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
	if err := c.deleteAttestors(attestorsToDelete); err != nil {
		return err
	}

	//
	//
	if headingAligned {
		return nil
	}

	//
	return ErrErroredHeadingIntersectPointing
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
func (c *Customer) contrastNewRapidLedgerWithAttestor(ctx context.Context, faultc chan error, l *kinds.RapidLedger,
	attestor source.Source, attestorOrdinal int,
) {
	h := l.AttestedHeading

	rapidLedger, err := attestor.RapidLedger(ctx, h.Level)
	switch err {
	//
	case nil:
		break

	//
	//
	case source.ErrNoReply, source.ErrRapidLedgerNegateLocated, context.DeadlineExceeded, context.Canceled:
		faultc <- err
		return

	//
	//
	//
	//
	case source.ErrLevelTooElevated:
		//
		var isObjectiveLevel bool
		isObjectiveLevel, rapidLedger, err = c.fetchObjectiveLedgerOrNewest(ctx, h.Level, attestor)
		if err != nil {
			faultc <- err
			return
		}

		//
		//
		if isObjectiveLevel {
			break
		}

		//
		//
		if !rapidLedger.Time.Before(h.Time) {
			faultc <- ErrClashingHeadings{Ledger: rapidLedger, AttestorOrdinal: attestorOrdinal}
			return
		}

		//
		//
		//
		//
		time.Sleep(2*c.maximumTimerDeviation + c.maximumLedgerDelay)
		isObjectiveLevel, rapidLedger, err = c.fetchObjectiveLedgerOrNewest(ctx, h.Level, attestor)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				faultc <- err
			} else {
				faultc <- errFlawedAttestor{Cause: err, AttestorOrdinal: attestorOrdinal}
			}
			return
		}
		if isObjectiveLevel {
			break
		}

		//
		//
		if !rapidLedger.Time.Before(h.Time) {
			faultc <- ErrClashingHeadings{Ledger: rapidLedger, AttestorOrdinal: attestorOrdinal}
			return
		}

		//
		//
		//
		//
		//
		//
		//
		faultc <- source.ErrNoReply
		return

	default:
		//
		//
		faultc <- errFlawedAttestor{Cause: err, AttestorOrdinal: attestorOrdinal}
		return
	}

	if !bytes.Equal(h.Digest(), rapidLedger.Digest()) {
		faultc <- ErrClashingHeadings{Ledger: rapidLedger, AttestorOrdinal: attestorOrdinal}
	}

	//
	desired, got := l.RatifierAssign.RecommenderUrgencyDigest(), rapidLedger.RatifierAssign.RecommenderUrgencyDigest()
	if !bytes.Equal(desired, got) {
		faultc <- ErrRecommenderUrgenciesDeviate{AttestorDigest: got, AttestorOrdinal: attestorOrdinal, LeadingDigest: desired}
	}

	c.tracer.Diagnose("REDACTED", "REDACTED", h.Level, "REDACTED", attestorOrdinal)
	faultc <- nil
}

//
func (c *Customer) transmitProof(ctx context.Context, ev *kinds.RapidCustomerAssaultProof, subscriber source.Source) {
	err := subscriber.NotifyProof(ctx, ev)
	if err != nil {
		c.tracer.Fault("REDACTED", "REDACTED", ev, "REDACTED", subscriber)
	}
}

//
//
func (c *Customer) processClashingHeadings(
	ctx context.Context,
	leadingTrack []*kinds.RapidLedger,
	demandingLedger *kinds.RapidLedger,
	attestorOrdinal int,
	now time.Time,
) error {
	aidingAttestor := c.attestors[attestorOrdinal]
	attestorTrack, leadingLedger, err := c.scrutinizeClashingHeadingVersusTrack(
		ctx,
		leadingTrack,
		demandingLedger,
		aidingAttestor,
		now,
	)
	if err != nil {
		c.tracer.Details("REDACTED", "REDACTED", aidingAttestor, "REDACTED", err)
		return nil
	}

	//
	//
	sharedLedger, validatedLedger := attestorTrack[0], attestorTrack[len(attestorTrack)-1]
	proofVersusLeading := newRapidCustomerAssaultProof(leadingLedger, validatedLedger, sharedLedger)
	c.tracer.Fault("REDACTED", "REDACTED", proofVersusLeading,
		"REDACTED", c.leading, "REDACTED", aidingAttestor)
	c.transmitProof(ctx, proofVersusLeading, aidingAttestor)

	if leadingLedger.Endorse.Cycle != attestorTrack[len(attestorTrack)-1].Endorse.Cycle {
		c.tracer.Details("REDACTED" +
			"REDACTED" +
			"REDACTED")
	}

	//
	//
	//
	leadingTrack, attestorLedger, err := c.scrutinizeClashingHeadingVersusTrack(
		ctx,
		attestorTrack,
		leadingLedger,
		c.leading,
		now,
	)
	if err != nil {
		c.tracer.Details("REDACTED", "REDACTED", c.leading, "REDACTED", err)
		return ErrRapidCustomerAssault
	}

	//
	sharedLedger, validatedLedger = leadingTrack[0], leadingTrack[len(leadingTrack)-1]
	proofVersusAttestor := newRapidCustomerAssaultProof(attestorLedger, validatedLedger, sharedLedger)
	c.tracer.Fault("REDACTED", "REDACTED", proofVersusAttestor,
		"REDACTED", c.leading, "REDACTED", aidingAttestor)
	c.transmitProof(ctx, proofVersusAttestor, c.leading)
	//
	return ErrRapidCustomerAssault
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
func (c *Customer) scrutinizeClashingHeadingVersusTrack(
	ctx context.Context,
	track []*kinds.RapidLedger,
	objectiveLedger *kinds.RapidLedger,
	origin source.Source, now time.Time,
) ([]*kinds.RapidLedger, *kinds.RapidLedger, error) {
	var (
		earlierCertifiedLedger, originLedger *kinds.RapidLedger
		originTrack                          []*kinds.RapidLedger
		err                                  error
	)

	if objectiveLedger.Level < track[0].Level {
		return nil, nil, fmt.Errorf("REDACTED",
			objectiveLedger.Level, track[0].Level)
	}

	for idx, trackLedger := range track {
		//
		//
		if trackLedger.Level > objectiveLedger.Level {
			//
			//
			//
			//
			if trackLedger.Time.After(objectiveLedger.Time) {
				return nil, nil,
					errors.New("REDACTED")
			}

			//
			//
			if earlierCertifiedLedger.Level != objectiveLedger.Level {
				originTrack, err = c.validateOmitting(ctx, origin, earlierCertifiedLedger, objectiveLedger, now)
				if err != nil {
					return nil, nil, fmt.Errorf("REDACTED", err)
				}
			}
			return originTrack, trackLedger, nil
		}

		//
		if trackLedger.Level == objectiveLedger.Level {
			originLedger = objectiveLedger
		} else {
			originLedger, err = origin.RapidLedger(ctx, trackLedger.Level)
			if err != nil {
				return nil, nil, fmt.Errorf("REDACTED", err)
			}
		}

		//
		//
		if idx == 0 {
			if statehash, transhash := originLedger.Digest(), trackLedger.Digest(); !bytes.Equal(statehash, transhash) {
				return nil, nil, fmt.Errorf("REDACTED",
					transhash, statehash)
			}
			earlierCertifiedLedger = originLedger
			continue
		}

		//
		//
		originTrack, err = c.validateOmitting(ctx, origin, earlierCertifiedLedger, originLedger, now)
		if err != nil {
			return nil, nil, fmt.Errorf("REDACTED", err)
		}
		//
		if statehash, transhash := originLedger.Digest(), trackLedger.Digest(); !bytes.Equal(statehash, transhash) {
			//
			return originTrack, trackLedger, nil
		}

		//
		earlierCertifiedLedger = originLedger
	}

	//
	//
	//
	return nil, nil, errNoDeviation
}

//
//
//
func (c *Customer) fetchObjectiveLedgerOrNewest(
	ctx context.Context,
	level int64,
	attestor source.Source,
) (bool, *kinds.RapidLedger, error) {
	rapidLedger, err := attestor.RapidLedger(ctx, 0)
	if err != nil {
		return false, nil, err
	}

	if rapidLedger.Level == level {
		//
		//
		return true, rapidLedger, nil
	}

	if rapidLedger.Level > level {
		//
		//
		//
		rapidLedger, err := attestor.RapidLedger(ctx, level)
		return true, rapidLedger, err
	}

	return false, rapidLedger, nil
}

//
//
func newRapidCustomerAssaultProof(disputed, validated, shared *kinds.RapidLedger) *kinds.RapidCustomerAssaultProof {
	ev := &kinds.RapidCustomerAssaultProof{ClashingLedger: disputed}
	//
	//
	//
	if ev.ClashingHeadingIsCorrupt(validated.Heading) {
		ev.SharedLevel = shared.Level
		ev.Timestamp = shared.Time
		ev.SumPollingEnergy = shared.RatifierAssign.SumPollingEnergy()
	} else {
		ev.SharedLevel = validated.Level
		ev.Timestamp = validated.Time
		ev.SumPollingEnergy = validated.RatifierAssign.SumPollingEnergy()
	}
	ev.FaultyRatifiers = ev.FetchFaultyRatifiers(shared.RatifierAssign, validated.AttestedHeading)
	return ev
}
