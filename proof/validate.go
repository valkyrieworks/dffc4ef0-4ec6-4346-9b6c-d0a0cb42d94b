package proof

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
//
//
func (incidentpool *Hub) validate(proof kinds.Proof) error {
	var (
		status          = incidentpool.Status()
		altitude         = status.FinalLedgerAltitude
		proofParameters = status.AgreementSettings.Proof
	)

	//
	ledgerSummary := incidentpool.ledgerDepot.FetchLedgerSummary(proof.Altitude())
	if ledgerSummary == nil {
		return fmt.Errorf("REDACTED", proof.Altitude())
	}
	occurenceMoment := ledgerSummary.Heading.Moment
	if proof.Moment() != occurenceMoment {
		return fmt.Errorf("REDACTED",
			proof.Moment(), occurenceMoment)
	}

	//
	if EqualsProofLapsed(altitude, status.FinalLedgerMoment, proof.Altitude(), occurenceMoment, proofParameters) {
		return fmt.Errorf(
			"REDACTED",
			proof.Altitude(),
			occurenceMoment,
			altitude-proofParameters.MaximumLifespanCountLedgers,
			status.FinalLedgerMoment.Add(proofParameters.MaximumLifespanInterval),
		)
	}

	//
	switch ev := proof.(type) {
	case *kinds.ReplicatedBallotProof:
		itemAssign, err := incidentpool.statusDatastore.FetchAssessors(proof.Altitude())
		if err != nil {
			return err
		}
		return ValidateReplicatedBallot(ev, status.SuccessionUUID, itemAssign)

	case *kinds.AgileCustomerOnslaughtProof:
		sharedHeadline, err := obtainNotatedHeadline(incidentpool.ledgerDepot, proof.Altitude())
		if err != nil {
			return err
		}
		sharedValues, err := incidentpool.statusDatastore.FetchAssessors(proof.Altitude())
		if err != nil {
			return err
		}
		reliableHeading := sharedHeadline
		//
		if proof.Altitude() != ev.DiscordantLedger.Altitude {
			reliableHeading, err = obtainNotatedHeadline(incidentpool.ledgerDepot, ev.DiscordantLedger.Altitude)
			if err != nil {
				//
				//

				//
				//
				newestAltitude := incidentpool.ledgerDepot.Altitude()
				reliableHeading, err = obtainNotatedHeadline(incidentpool.ledgerDepot, newestAltitude)
				if err != nil {
					return err
				}
				if reliableHeading.Moment.Before(ev.DiscordantLedger.Moment) {
					return fmt.Errorf("REDACTED",
						reliableHeading.Moment, ev.DiscordantLedger.Moment,
					)
				}
			}
		}

		err = ValidateAgileCustomerOnslaught(ev, sharedHeadline, reliableHeading, sharedValues, status.FinalLedgerMoment,
			status.AgreementSettings.Proof.MaximumLifespanInterval)
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
func ValidateAgileCustomerOnslaught(
	e *kinds.AgileCustomerOnslaughtProof,
	sharedHeadline, reliableHeading *kinds.NotatedHeading,
	sharedValues *kinds.AssessorAssign,
	now time.Time, //
	relianceCycle time.Duration, //
) error {
	//
	//

	//
	//
	if sharedHeadline.Altitude != e.DiscordantLedger.Altitude {
		err := sharedValues.ValidateEndorseAgileRelyingEveryNotations(reliableHeading.SuccessionUUID, e.DiscordantLedger.Endorse, agile.FallbackRelianceStratum)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		//
	} else if e.DiscordantHeadingEqualsUnfit(reliableHeading.Heading) {
		return errors.New("REDACTED" +
			"REDACTED")
	}

	//
	if err := e.DiscordantLedger.AssessorAssign.ValidateEndorseAgileEveryNotations(reliableHeading.SuccessionUUID, e.DiscordantLedger.Endorse.LedgerUUID,
		e.DiscordantLedger.Altitude, e.DiscordantLedger.Endorse); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if occurenceSum, valuesSum := e.SumBallotingPotency, sharedValues.SumBallotingPotency(); occurenceSum != valuesSum {
		return fmt.Errorf("REDACTED",
			occurenceSum, valuesSum)
	}

	//
	if e.DiscordantLedger.Altitude > reliableHeading.Altitude && e.DiscordantLedger.Moment.After(reliableHeading.Moment) {
		return fmt.Errorf("REDACTED",
			e.DiscordantLedger.Moment, reliableHeading.Moment,
		)

		//
	} else if bytes.Equal(reliableHeading.Digest(), e.DiscordantLedger.Digest()) {
		return fmt.Errorf("REDACTED",
			reliableHeading.Digest())
	}

	return certifyIfaceProof(e, sharedValues, reliableHeading)
}

//
//
//
//
//
//
func ValidateReplicatedBallot(e *kinds.ReplicatedBallotProof, successionUUID string, itemAssign *kinds.AssessorAssign) error {
	_, val := itemAssign.ObtainViaLocation(e.BallotAN.AssessorLocation)
	if val == nil {
		return fmt.Errorf("REDACTED", e.BallotAN.AssessorLocation, e.Altitude())
	}
	publicToken := val.PublicToken

	//
	if e.BallotAN.Altitude != e.BallotBYTE.Altitude ||
		e.BallotAN.Iteration != e.BallotBYTE.Iteration ||
		e.BallotAN.Kind != e.BallotBYTE.Kind {
		return fmt.Errorf("REDACTED",
			e.BallotAN.Altitude, e.BallotAN.Iteration, e.BallotAN.Kind,
			e.BallotBYTE.Altitude, e.BallotBYTE.Iteration, e.BallotBYTE.Kind)
	}

	//
	if !bytes.Equal(e.BallotAN.AssessorLocation, e.BallotBYTE.AssessorLocation) {
		return fmt.Errorf("REDACTED",
			e.BallotAN.AssessorLocation,
			e.BallotBYTE.AssessorLocation,
		)
	}

	//
	if e.BallotAN.LedgerUUID.Matches(e.BallotBYTE.LedgerUUID) {
		return fmt.Errorf(
			"REDACTED",
			e.BallotAN.LedgerUUID,
		)
	}

	//
	location := e.BallotAN.AssessorLocation
	if !bytes.Equal(publicToken.Location(), location) {
		return fmt.Errorf("REDACTED",
			location, publicToken, publicToken.Location())
	}

	//
	if val.BallotingPotency != e.AssessorPotency {
		return fmt.Errorf("REDACTED",
			e.AssessorPotency, val.BallotingPotency)
	}
	if itemAssign.SumBallotingPotency() != e.SumBallotingPotency {
		return fmt.Errorf("REDACTED",
			e.SumBallotingPotency, itemAssign.SumBallotingPotency())
	}

	va := e.BallotAN.TowardSchema()
	vb := e.BallotBYTE.TowardSchema()
	//
	if !publicToken.ValidateNotation(kinds.BallotAttestOctets(successionUUID, va), e.BallotAN.Notation) {
		return fmt.Errorf("REDACTED", kinds.FaultBallotUnfitSigning)
	}
	if !publicToken.ValidateNotation(kinds.BallotAttestOctets(successionUUID, vb), e.BallotBYTE.Notation) {
		return fmt.Errorf("REDACTED", kinds.FaultBallotUnfitSigning)
	}

	return nil
}

//
//
func certifyIfaceProof(
	ev *kinds.AgileCustomerOnslaughtProof,
	sharedValues *kinds.AssessorAssign,
	reliableHeading *kinds.NotatedHeading,
) error {
	if occurenceSum, valuesSum := ev.SumBallotingPotency, sharedValues.SumBallotingPotency(); occurenceSum != valuesSum {
		return fmt.Errorf("REDACTED",
			occurenceSum, valuesSum)
	}

	//
	//
	//
	assessors := ev.ObtainTreacherousAssessors(sharedValues, reliableHeading)

	//
	//
	if assessors == nil && ev.TreacherousAssessors != nil {
		return fmt.Errorf(
			"REDACTED",
			len(ev.TreacherousAssessors),
		)
	}

	if exp, got := len(assessors), len(ev.TreacherousAssessors); exp != got {
		return fmt.Errorf("REDACTED", exp, got)
	}

	for idx, val := range assessors {
		occurenceByzantine := ev.TreacherousAssessors[idx]
		if !bytes.Equal(occurenceByzantine.Location, val.Location) {
			return fmt.Errorf(
				"REDACTED",
				val.Location, occurenceByzantine.Location,
			)
		}

		if occurenceByzantine.BallotingPotency != val.BallotingPotency {
			return fmt.Errorf(
				"REDACTED",
				val.BallotingPotency, occurenceByzantine.BallotingPotency,
			)
		}

		//
		//
		//
		if occurenceByzantine.PublicToken == nil {
			return fmt.Errorf("REDACTED", idx)
		}
		if !bytes.Equal(occurenceByzantine.Location, occurenceByzantine.PublicToken.Location()) {
			return fmt.Errorf(
				"REDACTED",
				idx, occurenceByzantine.Location, occurenceByzantine.PublicToken.Location(),
			)
		}
	}

	return nil
}

func obtainNotatedHeadline(ledgerDepot LedgerDepot, altitude int64) (*kinds.NotatedHeading, error) {
	ledgerSummary := ledgerDepot.FetchLedgerSummary(altitude)
	if ledgerSummary == nil {
		return nil, fmt.Errorf("REDACTED", altitude)
	}
	endorse := ledgerDepot.FetchLedgerEndorse(altitude)
	if endorse == nil {
		return nil, fmt.Errorf("REDACTED", altitude)
	}
	return &kinds.NotatedHeading{
		Heading: &ledgerSummary.Heading,
		Endorse: endorse,
	}, nil
}

//
func EqualsProofLapsed(altitudePresent int64, momentPresent time.Time, altitudeOccurence int64, momentOccurence time.Time, proofParameters kinds.ProofParameters) bool {
	lifespanInterval := momentPresent.Sub(momentOccurence)
	lifespanCountLedgers := altitudePresent - altitudeOccurence

	if lifespanInterval > proofParameters.MaximumLifespanInterval && lifespanCountLedgers > proofParameters.MaximumLifespanCountLedgers {
		return true
	}
	return false
}
