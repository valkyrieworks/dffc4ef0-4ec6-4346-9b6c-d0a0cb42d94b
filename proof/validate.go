package proof

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/rapid"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
//
//
func (eventpool *Depository) validate(proof kinds.Proof) error {
	var (
		status          = eventpool.Status()
		level         = status.FinalLedgerLevel
		proofOptions = status.AgreementOptions.Proof
	)

	//
	ledgerMeta := eventpool.ledgerDepot.ImportLedgerMeta(proof.Level())
	if ledgerMeta == nil {
		return fmt.Errorf("REDACTED", proof.Level())
	}
	evtTime := ledgerMeta.Heading.Time
	if proof.Time() != evtTime {
		return fmt.Errorf("REDACTED",
			proof.Time(), evtTime)
	}

	//
	if IsProofLapsed(level, status.FinalLedgerTime, proof.Level(), evtTime, proofOptions) {
		return fmt.Errorf(
			"REDACTED",
			proof.Level(),
			evtTime,
			level-proofOptions.MaximumDurationCountLedgers,
			status.FinalLedgerTime.Add(proofOptions.MaximumDurationPeriod),
		)
	}

	//
	switch ev := proof.(type) {
	case *kinds.ReplicatedBallotProof:
		valueCollection, err := eventpool.statusStore.ImportRatifiers(proof.Level())
		if err != nil {
			return err
		}
		return ValidateReplicatedBallot(ev, status.LedgerUID, valueCollection)

	case *kinds.RapidCustomerAssaultProof:
		sharedHeading, err := fetchAttestedHeading(eventpool.ledgerDepot, proof.Level())
		if err != nil {
			return err
		}
		sharedValues, err := eventpool.statusStore.ImportRatifiers(proof.Level())
		if err != nil {
			return err
		}
		validatedHeading := sharedHeading
		//
		if proof.Level() != ev.ClashingLedger.Level {
			validatedHeading, err = fetchAttestedHeading(eventpool.ledgerDepot, ev.ClashingLedger.Level)
			if err != nil {
				//
				//

				//
				//
				newestLevel := eventpool.ledgerDepot.Level()
				validatedHeading, err = fetchAttestedHeading(eventpool.ledgerDepot, newestLevel)
				if err != nil {
					return err
				}
				if validatedHeading.Time.Before(ev.ClashingLedger.Time) {
					return fmt.Errorf("REDACTED",
						validatedHeading.Time, ev.ClashingLedger.Time,
					)
				}
			}
		}

		err = ValidateRapidCustomerAssault(ev, sharedHeading, validatedHeading, sharedValues, status.FinalLedgerTime,
			status.AgreementOptions.Proof.MaximumDurationPeriod)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("REDACTED", proof)
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
//
//
func ValidateRapidCustomerAssault(
	e *kinds.RapidCustomerAssaultProof,
	sharedHeading, validatedHeading *kinds.AttestedHeading,
	sharedValues *kinds.RatifierAssign,
	now time.Time, //
	relianceDuration time.Duration, //
) error {
	//
	//

	//
	//
	if sharedHeading.Level != e.ClashingLedger.Level {
		err := sharedValues.ValidateEndorseRapidValidatingAllEndorsements(validatedHeading.LedgerUID, e.ClashingLedger.Endorse, rapid.StandardRelianceLayer)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		//
	} else if e.ClashingHeadingIsCorrupt(validatedHeading.Heading) {
		return errors.New("REDACTED" +
			"REDACTED")
	}

	//
	if err := e.ClashingLedger.RatifierAssign.ValidateEndorseRapidAllEndorsements(validatedHeading.LedgerUID, e.ClashingLedger.Endorse.LedgerUID,
		e.ClashingLedger.Level, e.ClashingLedger.Endorse); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if evtSum, valuesSum := e.SumPollingEnergy, sharedValues.SumPollingEnergy(); evtSum != valuesSum {
		return fmt.Errorf("REDACTED",
			evtSum, valuesSum)
	}

	//
	if e.ClashingLedger.Level > validatedHeading.Level && e.ClashingLedger.Time.After(validatedHeading.Time) {
		return fmt.Errorf("REDACTED",
			e.ClashingLedger.Time, validatedHeading.Time,
		)

		//
	} else if bytes.Equal(validatedHeading.Digest(), e.ClashingLedger.Digest()) {
		return fmt.Errorf("REDACTED",
			validatedHeading.Digest())
	}

	return certifyIfaceProof(e, sharedValues, validatedHeading)
}

//
//
//
//
//
//
func ValidateReplicatedBallot(e *kinds.ReplicatedBallotProof, ledgerUID string, valueCollection *kinds.RatifierAssign) error {
	_, val := valueCollection.FetchByLocation(e.BallotA.RatifierLocation)
	if val == nil {
		return fmt.Errorf("REDACTED", e.BallotA.RatifierLocation, e.Level())
	}
	publicKey := val.PublicKey

	//
	if e.BallotA.Level != e.BallotBYTE.Level ||
		e.BallotA.Cycle != e.BallotBYTE.Cycle ||
		e.BallotA.Kind != e.BallotBYTE.Kind {
		return fmt.Errorf("REDACTED",
			e.BallotA.Level, e.BallotA.Cycle, e.BallotA.Kind,
			e.BallotBYTE.Level, e.BallotBYTE.Cycle, e.BallotBYTE.Kind)
	}

	//
	if !bytes.Equal(e.BallotA.RatifierLocation, e.BallotBYTE.RatifierLocation) {
		return fmt.Errorf("REDACTED",
			e.BallotA.RatifierLocation,
			e.BallotBYTE.RatifierLocation,
		)
	}

	//
	if e.BallotA.LedgerUID.Matches(e.BallotBYTE.LedgerUID) {
		return fmt.Errorf(
			"REDACTED",
			e.BallotA.LedgerUID,
		)
	}

	//
	address := e.BallotA.RatifierLocation
	if !bytes.Equal(publicKey.Location(), address) {
		return fmt.Errorf("REDACTED",
			address, publicKey, publicKey.Location())
	}

	//
	if val.PollingEnergy != e.RatifierEnergy {
		return fmt.Errorf("REDACTED",
			e.RatifierEnergy, val.PollingEnergy)
	}
	if valueCollection.SumPollingEnergy() != e.SumPollingEnergy {
		return fmt.Errorf("REDACTED",
			e.SumPollingEnergy, valueCollection.SumPollingEnergy())
	}

	va := e.BallotA.ToSchema()
	vb := e.BallotBYTE.ToSchema()
	//
	if !publicKey.ValidateAutograph(kinds.BallotAttestOctets(ledgerUID, va), e.BallotA.Autograph) {
		return fmt.Errorf("REDACTED", kinds.ErrBallotCorruptAutograph)
	}
	if !publicKey.ValidateAutograph(kinds.BallotAttestOctets(ledgerUID, vb), e.BallotBYTE.Autograph) {
		return fmt.Errorf("REDACTED", kinds.ErrBallotCorruptAutograph)
	}

	return nil
}

//
//
func certifyIfaceProof(
	ev *kinds.RapidCustomerAssaultProof,
	sharedValues *kinds.RatifierAssign,
	validatedHeading *kinds.AttestedHeading,
) error {
	if evtSum, valuesSum := ev.SumPollingEnergy, sharedValues.SumPollingEnergy(); evtSum != valuesSum {
		return fmt.Errorf("REDACTED",
			evtSum, valuesSum)
	}

	//
	//
	//
	ratifiers := ev.FetchFaultyRatifiers(sharedValues, validatedHeading)

	//
	//
	if ratifiers == nil && ev.FaultyRatifiers != nil {
		return fmt.Errorf(
			"REDACTED",
			len(ev.FaultyRatifiers),
		)
	}

	if exp, got := len(ratifiers), len(ev.FaultyRatifiers); exp != got {
		return fmt.Errorf("REDACTED", exp, got)
	}

	for idx, val := range ratifiers {
		if !bytes.Equal(ev.FaultyRatifiers[idx].Location, val.Location) {
			return fmt.Errorf(
				"REDACTED",
				val.Location, ev.FaultyRatifiers[idx].Location,
			)
		}

		if ev.FaultyRatifiers[idx].PollingEnergy != val.PollingEnergy {
			return fmt.Errorf(
				"REDACTED",
				val.PollingEnergy, ev.FaultyRatifiers[idx].PollingEnergy,
			)
		}
	}

	return nil
}

func fetchAttestedHeading(ledgerDepot LedgerDepot, level int64) (*kinds.AttestedHeading, error) {
	ledgerMeta := ledgerDepot.ImportLedgerMeta(level)
	if ledgerMeta == nil {
		return nil, fmt.Errorf("REDACTED", level)
	}
	endorse := ledgerDepot.ImportLedgerEndorse(level)
	if endorse == nil {
		return nil, fmt.Errorf("REDACTED", level)
	}
	return &kinds.AttestedHeading{
		Heading: &ledgerMeta.Heading,
		Endorse: endorse,
	}, nil
}

//
func IsProofLapsed(levelPresent int64, timePresent time.Time, levelEvt int64, timeEvt time.Time, proofOptions kinds.ProofOptions) bool {
	eraPeriod := timePresent.Sub(timeEvt)
	eraCountLedgers := levelPresent - levelEvt

	if eraPeriod > proofOptions.MaximumDurationPeriod && eraCountLedgers > proofOptions.MaximumDurationCountLedgers {
		return true
	}
	return false
}
