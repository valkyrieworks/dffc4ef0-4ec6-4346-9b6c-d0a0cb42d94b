package kinds

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/cluster"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
)

const clusterValidateLimit = 2

func mustClusterValidate(values *AssessorAssign, endorse *Endorse) bool {
	return len(endorse.Notations) >= clusterValidateLimit &&
		cluster.SustainsClusterValidator(values.ObtainNominator().PublicToken) &&
		values.EveryTokensPossessIdenticalKind()
}

//
//
//
//
//
//
//
func ValidateEndorse(successionUUID string, values *AssessorAssign, ledgerUUID LedgerUUID,
	altitude int64, endorse *Endorse,
) error {
	//
	if err := validateFundamentalValuesAlsoEndorse(values, endorse, altitude, ledgerUUID); err != nil {
		return err
	}

	//
	//
	ballotingPotencyRequired := values.SumBallotingPotency() * 2 / 3

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUUIDMarker == LedgerUUIDMarkerMissing }

	//
	tally := func(c EndorseSignature) bool { return c.LedgerUUIDMarker == LedgerUUIDMarkerEndorse }

	//
	if mustClusterValidate(values, endorse) {
		return validateEndorseCluster(successionUUID, values, endorse,
			ballotingPotencyRequired, bypass, tally, true, true, nil, nil)
	}

	//
	return validateEndorseUnique(successionUUID, values, endorse, ballotingPotencyRequired,
		bypass, tally, true, true, nil)
}

//

//
//
//
//
func ValidateEndorseAgile(
	successionUUID string,
	values *AssessorAssign,
	ledgerUUID LedgerUUID,
	altitude int64,
	endorse *Endorse,
) error {
	return validateEndorseAgileIntrinsic(successionUUID, values, ledgerUUID, altitude, endorse, false, nil)
}

//
//
//
//
//
//
//
func ValidateEndorseAgileUsingStash(
	successionUUID string,
	values *AssessorAssign,
	ledgerUUID LedgerUUID,
	altitude int64,
	endorse *Endorse,
	attestedSigningStash *SigningStash,
) error {
	return validateEndorseAgileIntrinsic(successionUUID, values, ledgerUUID, altitude, endorse, false, attestedSigningStash)
}

//
//
//
func ValidateEndorseAgileEveryNotations(
	successionUUID string,
	values *AssessorAssign,
	ledgerUUID LedgerUUID,
	altitude int64,
	endorse *Endorse,
) error {
	return validateEndorseAgileIntrinsic(successionUUID, values, ledgerUUID, altitude, endorse, true, nil)
}

func validateEndorseAgileIntrinsic(
	successionUUID string,
	values *AssessorAssign,
	ledgerUUID LedgerUUID,
	altitude int64,
	endorse *Endorse,
	tallyEveryNotations bool,
	attestedSigningStash *SigningStash,
) error {
	//
	if err := validateFundamentalValuesAlsoEndorse(values, endorse, altitude, ledgerUUID); err != nil {
		return err
	}

	//
	ballotingPotencyRequired := values.SumBallotingPotency() * 2 / 3

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUUIDMarker != LedgerUUIDMarkerEndorse }

	//
	tally := func(c EndorseSignature) bool { return true }

	//
	if mustClusterValidate(values, endorse) {
		return validateEndorseCluster(successionUUID, values, endorse,
			ballotingPotencyRequired, bypass, tally, tallyEveryNotations, true, nil, attestedSigningStash)
	}

	//
	return validateEndorseUnique(successionUUID, values, endorse, ballotingPotencyRequired,
		bypass, tally, tallyEveryNotations, true, attestedSigningStash)
}

//
//
//
//
//
//
//
//
func ValidateEndorseAgileRelying(
	successionUUID string,
	values *AssessorAssign,
	endorse *Endorse,
	relianceStratum strongarithmetic.Portion,
) error {
	return validateEndorseAgileRelyingIntrinsic(successionUUID, values, endorse, relianceStratum, false, nil)
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
func ValidateEndorseAgileRelyingUsingStash(
	successionUUID string,
	values *AssessorAssign,
	endorse *Endorse,
	relianceStratum strongarithmetic.Portion,
	attestedSigningStash *SigningStash,
) error {
	return validateEndorseAgileRelyingIntrinsic(successionUUID, values, endorse, relianceStratum, false, attestedSigningStash)
}

//
//
//
//
//
//
//
func ValidateEndorseAgileRelyingEveryNotations(
	successionUUID string,
	values *AssessorAssign,
	endorse *Endorse,
	relianceStratum strongarithmetic.Portion,
) error {
	return validateEndorseAgileRelyingIntrinsic(successionUUID, values, endorse, relianceStratum, true, nil)
}

func validateEndorseAgileRelyingIntrinsic(
	successionUUID string,
	values *AssessorAssign,
	endorse *Endorse,
	relianceStratum strongarithmetic.Portion,
	tallyEveryNotations bool,
	attestedSigningStash *SigningStash,
) error {
	//
	if values == nil {
		return errors.New("REDACTED")
	}
	if relianceStratum.Divisor == 0 {
		return errors.New("REDACTED")
	}
	if endorse == nil {
		return errors.New("REDACTED")
	}

	//
	sumBallotingPotencyMultiplyViaDividend, overrun := secureMultiply(values.SumBallotingPotency(), int64(relianceStratum.Dividend))
	if overrun {
		return errors.New("REDACTED")
	}
	ballotingPotencyRequired := sumBallotingPotencyMultiplyViaDividend / int64(relianceStratum.Divisor)

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUUIDMarker != LedgerUUIDMarkerEndorse }

	//
	tally := func(c EndorseSignature) bool { return true }

	//
	//
	//
	if mustClusterValidate(values, endorse) {
		return validateEndorseCluster(successionUUID, values, endorse,
			ballotingPotencyRequired, bypass, tally, tallyEveryNotations, false, nil, attestedSigningStash)
	}

	//
	return validateEndorseUnique(successionUUID, values, endorse, ballotingPotencyRequired,
		bypass, tally, tallyEveryNotations, false, attestedSigningStash)
}

//
//
func CertifyDigest(h []byte) error {
	if len(h) > 0 && len(h) != tenderminthash.Extent {
		return fmt.Errorf("REDACTED",
			tenderminthash.Extent,
			len(h),
		)
	}
	return nil
}

//

//
//
//
//
//
//
func validateEndorseCluster(
	successionUUID string,
	values *AssessorAssign,
	endorse *Endorse,
	ballotingPotencyRequired int64,
	bypassSignature func(EndorseSignature) bool,
	tallySignature func(EndorseSignature) bool,
	tallyEveryNotations bool,
	scanAscendViaPosition bool,
	clusterValidator security.ClusterValidator,
	attestedSigningStash *SigningStash,
) error {
	var (
		val                *Assessor
		itemOffset             int32
		observedValues           = make(map[int32]int, len(endorse.Notations))
		clusterSignatureIndices       = make([]int, 0, len(endorse.Notations))
		accountedBallotingPotency int64
	)
	//
	bv, ok := clusterValidator, true
	if clusterValidator == nil {
		bv, ok = cluster.GenerateClusterValidator(values.ObtainNominator().PublicToken)
	}
	//
	if !ok || len(endorse.Notations) < clusterValidateLimit {
		//
		return fmt.Errorf("REDACTED")
	}

	for idx, endorseSignature := range endorse.Notations {
		//
		if bypassSignature(endorseSignature) {
			continue
		}

		//
		//
		if scanAscendViaPosition {
			val = values.Assessors[idx]
			if !bytes.Equal(val.Location, endorseSignature.AssessorLocation) {
				return fmt.Errorf("REDACTED",
					idx, val.Location, endorseSignature.AssessorLocation)
			}
		} else {
			itemOffset, val = values.ObtainViaLocationAlterable(endorseSignature.AssessorLocation)

			//
			//
			if val == nil {
				continue
			}

			//
			//
			if initialPosition, ok := observedValues[itemOffset]; ok {
				ordinalPosition := idx
				return fmt.Errorf("REDACTED", val, initialPosition, ordinalPosition)
			}
			observedValues[itemOffset] = idx
		}

		//
		ballotAttestOctets := endorse.BallotAttestOctets(successionUUID, int32(idx))

		stashStrike := false
		if attestedSigningStash != nil {
			stashItem, signatureEqualsInsideStash := attestedSigningStash.Get(string(endorseSignature.Notation))
			stashStrike = signatureEqualsInsideStash && bytes.Equal(stashItem.AssessorLocation, val.PublicToken.Location()) && bytes.Equal(stashItem.BallotAttestOctets, ballotAttestOctets)
		}

		if !stashStrike {
			//
			if err := bv.Add(val.PublicToken, ballotAttestOctets, endorseSignature.Notation); err != nil {
				return err
			}
			clusterSignatureIndices = append(clusterSignatureIndices, idx)
		}

		//
		//
		if tallySignature(endorseSignature) {
			accountedBallotingPotency += val.BallotingPotency
		}

		//
		//
		if !tallyEveryNotations && accountedBallotingPotency > ballotingPotencyRequired {
			break
		}
	}

	//
	//
	if got, required := accountedBallotingPotency, ballotingPotencyRequired; got <= required {
		return FaultNegationAmpleBallotingPotencyNotated{Got: got, Required: required}
	}

	//
	if len(clusterSignatureIndices) == 0 {
		return nil
	}

	//
	ok, soundSignatures := bv.Validate()
	if ok {
		//
		if attestedSigningStash != nil {
			for i := range soundSignatures {
				idx := clusterSignatureIndices[i]
				sig := endorse.Notations[idx]
				attestedSigningStash.Add(string(sig.Notation), NotationStashDatum{
					AssessorLocation: sig.AssessorLocation,
					BallotAttestOctets:    endorse.BallotAttestOctets(successionUUID, int32(idx)),
				})
			}
		}

		return nil
	}

	//
	//
	for i, ok := range soundSignatures {
		//
		idx := clusterSignatureIndices[i]
		sig := endorse.Notations[idx]
		if !ok {
			return fmt.Errorf("REDACTED", idx, sig)
		}
		if attestedSigningStash != nil {
			attestedSigningStash.Add(string(sig.Notation), NotationStashDatum{
				AssessorLocation: sig.AssessorLocation,
				BallotAttestOctets:    endorse.BallotAttestOctets(successionUUID, int32(idx)),
			})
		}
	}

	//
	//
	//
	//
	return fmt.Errorf("REDACTED")
}

//

//
//
//
//
//
func validateEndorseUnique(
	successionUUID string,
	values *AssessorAssign,
	endorse *Endorse,
	ballotingPotencyRequired int64,
	bypassSignature func(EndorseSignature) bool,
	tallySignature func(EndorseSignature) bool,
	tallyEveryNotations bool,
	scanAscendViaPosition bool,
	attestedSigningStash *SigningStash,
) error {
	var (
		val                *Assessor
		itemOffset             int32
		observedValues           = make(map[int32]int, len(endorse.Notations))
		accountedBallotingPotency int64
		ballotAttestOctets      []byte
	)
	for idx, endorseSignature := range endorse.Notations {
		if bypassSignature(endorseSignature) {
			continue
		}

		if endorseSignature.CertifyFundamental() != nil {
			return fmt.Errorf("REDACTED", val, idx)
		}

		//
		//
		if scanAscendViaPosition {
			val = values.Assessors[idx]
			if !bytes.Equal(val.Location, endorseSignature.AssessorLocation) {
				return fmt.Errorf("REDACTED",
					idx, val.Location, endorseSignature.AssessorLocation)
			}
		} else {
			itemOffset, val = values.ObtainViaLocation(endorseSignature.AssessorLocation)

			//
			//
			if val == nil {
				continue
			}

			//
			//
			if initialPosition, ok := observedValues[itemOffset]; ok {
				ordinalPosition := idx
				return fmt.Errorf("REDACTED", val, initialPosition, ordinalPosition)
			}
			observedValues[itemOffset] = idx
		}

		if val.PublicToken == nil {
			return fmt.Errorf("REDACTED", val, idx)
		}

		ballotAttestOctets = endorse.BallotAttestOctets(successionUUID, int32(idx))

		stashToken, stashStrike := "REDACTED", false
		if attestedSigningStash != nil {
			stashToken = string(endorseSignature.Notation)
			stashItem, signatureEqualsInsideStash := attestedSigningStash.Get(stashToken)
			stashStrike = signatureEqualsInsideStash && bytes.Equal(stashItem.AssessorLocation, val.PublicToken.Location()) && bytes.Equal(stashItem.BallotAttestOctets, ballotAttestOctets)
		}

		if !stashStrike {
			if !val.PublicToken.ValidateNotation(ballotAttestOctets, endorseSignature.Notation) {
				return fmt.Errorf("REDACTED", idx, endorseSignature.Notation)
			}
			if attestedSigningStash != nil {
				attestedSigningStash.Add(stashToken, NotationStashDatum{
					AssessorLocation: val.PublicToken.Location(),
					BallotAttestOctets:    ballotAttestOctets,
				})
			}
		}

		//
		//
		if tallySignature(endorseSignature) {
			accountedBallotingPotency += val.BallotingPotency
		}

		//
		if !tallyEveryNotations && accountedBallotingPotency > ballotingPotencyRequired {
			return nil
		}
	}

	if got, required := accountedBallotingPotency, ballotingPotencyRequired; got <= required {
		return FaultNegationAmpleBallotingPotencyNotated{Got: got, Required: required}
	}

	return nil
}

func validateFundamentalValuesAlsoEndorse(values *AssessorAssign, endorse *Endorse, altitude int64, ledgerUUID LedgerUUID) error {
	if values == nil {
		return errors.New("REDACTED")
	}

	if endorse == nil {
		return errors.New("REDACTED")
	}

	if values.Extent() != len(endorse.Notations) {
		return strongminderrors.FreshFaultUnfitEndorseNotations(values.Extent(), len(endorse.Notations))
	}

	//
	if altitude != endorse.Altitude {
		return strongminderrors.FreshFaultUnfitEndorseAltitude(altitude, endorse.Altitude)
	}
	if !ledgerUUID.Matches(endorse.LedgerUUID) {
		return fmt.Errorf("REDACTED",
			ledgerUUID, endorse.LedgerUUID)
	}

	return nil
}
