package kinds

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/group"
	"github.com/valkyrieworks/vault/comethash"
	cometmath "github.com/valkyrieworks/utils/math"
	cometfaults "github.com/valkyrieworks/kinds/faults"
)

const groupValidateLimit = 2

func mustGroupValidate(values *RatifierAssign, endorse *Endorse) bool {
	return len(endorse.Endorsements) >= groupValidateLimit &&
		group.SustainsGroupValidator(values.FetchRecommender().PublicKey) &&
		values.AllKeysPossessIdenticalKind()
}

//
//
//
//
//
//
//
func ValidateEndorse(ledgerUID string, values *RatifierAssign, ledgerUID LedgerUID,
	level int64, endorse *Endorse,
) error {
	//
	if err := validateSimpleValuesAndEndorse(values, endorse, level, ledgerUID); err != nil {
		return err
	}

	//
	//
	pollingEnergyRequired := values.SumPollingEnergy() * 2 / 3

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUIDMark == LedgerUIDMarkMissing }

	//
	tally := func(c EndorseSignature) bool { return c.LedgerUIDMark == LedgerUIDMarkEndorse }

	//
	if mustGroupValidate(values, endorse) {
		return validateEndorseGroup(ledgerUID, values, endorse,
			pollingEnergyRequired, bypass, tally, true, true, nil, nil)
	}

	//
	return validateEndorseUnique(ledgerUID, values, endorse, pollingEnergyRequired,
		bypass, tally, true, true, nil)
}

//

//
//
//
//
func ValidateEndorseRapid(
	ledgerUID string,
	values *RatifierAssign,
	ledgerUID LedgerUID,
	level int64,
	endorse *Endorse,
) error {
	return validateEndorseRapidIntrinsic(ledgerUID, values, ledgerUID, level, endorse, false, nil)
}

//
//
//
//
//
//
//
func ValidateEndorseRapidWithRepository(
	ledgerUID string,
	values *RatifierAssign,
	ledgerUID LedgerUID,
	level int64,
	endorse *Endorse,
	certifiedAutographRepository *AutographRepository,
) error {
	return validateEndorseRapidIntrinsic(ledgerUID, values, ledgerUID, level, endorse, false, certifiedAutographRepository)
}

//
//
//
func ValidateEndorseRapidAllEndorsements(
	ledgerUID string,
	values *RatifierAssign,
	ledgerUID LedgerUID,
	level int64,
	endorse *Endorse,
) error {
	return validateEndorseRapidIntrinsic(ledgerUID, values, ledgerUID, level, endorse, true, nil)
}

func validateEndorseRapidIntrinsic(
	ledgerUID string,
	values *RatifierAssign,
	ledgerUID LedgerUID,
	level int64,
	endorse *Endorse,
	numberAllEndorsements bool,
	certifiedAutographRepository *AutographRepository,
) error {
	//
	if err := validateSimpleValuesAndEndorse(values, endorse, level, ledgerUID); err != nil {
		return err
	}

	//
	pollingEnergyRequired := values.SumPollingEnergy() * 2 / 3

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUIDMark != LedgerUIDMarkEndorse }

	//
	tally := func(c EndorseSignature) bool { return true }

	//
	if mustGroupValidate(values, endorse) {
		return validateEndorseGroup(ledgerUID, values, endorse,
			pollingEnergyRequired, bypass, tally, numberAllEndorsements, true, nil, certifiedAutographRepository)
	}

	//
	return validateEndorseUnique(ledgerUID, values, endorse, pollingEnergyRequired,
		bypass, tally, numberAllEndorsements, true, certifiedAutographRepository)
}

//
//
//
//
//
//
//
//
func ValidateEndorseRapidValidating(
	ledgerUID string,
	values *RatifierAssign,
	endorse *Endorse,
	validateLayer cometmath.Portion,
) error {
	return validateEndorseRapidValidatingIntrinsic(ledgerUID, values, endorse, validateLayer, false, nil)
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
func ValidateEndorseRapidValidatingWithRepository(
	ledgerUID string,
	values *RatifierAssign,
	endorse *Endorse,
	validateLayer cometmath.Portion,
	certifiedAutographRepository *AutographRepository,
) error {
	return validateEndorseRapidValidatingIntrinsic(ledgerUID, values, endorse, validateLayer, false, certifiedAutographRepository)
}

//
//
//
//
//
//
//
func ValidateEndorseRapidValidatingAllEndorsements(
	ledgerUID string,
	values *RatifierAssign,
	endorse *Endorse,
	validateLayer cometmath.Portion,
) error {
	return validateEndorseRapidValidatingIntrinsic(ledgerUID, values, endorse, validateLayer, true, nil)
}

func validateEndorseRapidValidatingIntrinsic(
	ledgerUID string,
	values *RatifierAssign,
	endorse *Endorse,
	validateLayer cometmath.Portion,
	numberAllEndorsements bool,
	certifiedAutographRepository *AutographRepository,
) error {
	//
	if values == nil {
		return errors.New("REDACTED")
	}
	if validateLayer.Divisor == 0 {
		return errors.New("REDACTED")
	}
	if endorse == nil {
		return errors.New("REDACTED")
	}

	//
	sumPollingEnergyMultiplyByDividend, overload := secureMultiply(values.SumPollingEnergy(), int64(validateLayer.Dividend))
	if overload {
		return errors.New("REDACTED")
	}
	pollingEnergyRequired := sumPollingEnergyMultiplyByDividend / int64(validateLayer.Divisor)

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUIDMark != LedgerUIDMarkEndorse }

	//
	tally := func(c EndorseSignature) bool { return true }

	//
	//
	//
	if mustGroupValidate(values, endorse) {
		return validateEndorseGroup(ledgerUID, values, endorse,
			pollingEnergyRequired, bypass, tally, numberAllEndorsements, false, nil, certifiedAutographRepository)
	}

	//
	return validateEndorseUnique(ledgerUID, values, endorse, pollingEnergyRequired,
		bypass, tally, numberAllEndorsements, false, certifiedAutographRepository)
}

//
//
func CertifyDigest(h []byte) error {
	if len(h) > 0 && len(h) != comethash.Volume {
		return fmt.Errorf("REDACTED",
			comethash.Volume,
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
func validateEndorseGroup(
	ledgerUID string,
	values *RatifierAssign,
	endorse *Endorse,
	pollingEnergyRequired int64,
	bypassSignature func(EndorseSignature) bool,
	numberSignature func(EndorseSignature) bool,
	numberAllEndorsements bool,
	scanUpByOrdinal bool,
	groupValidator vault.GroupValidator,
	certifiedAutographRepository *AutographRepository,
) error {
	var (
		val                *Ratifier
		valueIdx             int32
		viewedValues           = make(map[int32]int, len(endorse.Endorsements))
		groupSignatureIndices       = make([]int, 0, len(endorse.Endorsements))
		calculatedPollingEnergy int64
	)
	//
	bv, ok := groupValidator, true
	if groupValidator == nil {
		bv, ok = group.InstantiateGroupValidator(values.FetchRecommender().PublicKey)
	}
	//
	if !ok || len(endorse.Endorsements) < groupValidateLimit {
		//
		return fmt.Errorf("REDACTED")
	}

	for idx, endorseSignature := range endorse.Endorsements {
		//
		if bypassSignature(endorseSignature) {
			continue
		}

		//
		//
		if scanUpByOrdinal {
			val = values.Ratifiers[idx]
			if !bytes.Equal(val.Location, endorseSignature.RatifierLocation) {
				return fmt.Errorf("REDACTED",
					idx, val.Location, endorseSignature.RatifierLocation)
			}
		} else {
			valueIdx, val = values.FetchByLocationMut(endorseSignature.RatifierLocation)

			//
			//
			if val == nil {
				continue
			}

			//
			//
			if initialOrdinal, ok := viewedValues[valueIdx]; ok {
				momentOrdinal := idx
				return fmt.Errorf("REDACTED", val, initialOrdinal, momentOrdinal)
			}
			viewedValues[valueIdx] = idx
		}

		//
		ballotAttestOctets := endorse.BallotAttestOctets(ledgerUID, int32(idx))

		repositoryStrike := false
		if certifiedAutographRepository != nil {
			repositoryValue, signatureIsInRepository := certifiedAutographRepository.Get(string(endorseSignature.Autograph))
			repositoryStrike = signatureIsInRepository && bytes.Equal(repositoryValue.RatifierLocation, val.PublicKey.Location()) && bytes.Equal(repositoryValue.BallotAttestOctets, ballotAttestOctets)
		}

		if !repositoryStrike {
			//
			if err := bv.Add(val.PublicKey, ballotAttestOctets, endorseSignature.Autograph); err != nil {
				return err
			}
			groupSignatureIndices = append(groupSignatureIndices, idx)
		}

		//
		//
		if numberSignature(endorseSignature) {
			calculatedPollingEnergy += val.PollingEnergy
		}

		//
		//
		if !numberAllEndorsements && calculatedPollingEnergy > pollingEnergyRequired {
			break
		}
	}

	//
	//
	if got, required := calculatedPollingEnergy, pollingEnergyRequired; got <= required {
		return ErrNoSufficientPollingEnergyAttested{Got: got, Required: required}
	}

	//
	if len(groupSignatureIndices) == 0 {
		return nil
	}

	//
	ok, soundAutographs := bv.Validate()
	if ok {
		//
		if certifiedAutographRepository != nil {
			for i := range soundAutographs {
				idx := groupSignatureIndices[i]
				sig := endorse.Endorsements[idx]
				certifiedAutographRepository.Add(string(sig.Autograph), AutographRepositoryItem{
					RatifierLocation: sig.RatifierLocation,
					BallotAttestOctets:    endorse.BallotAttestOctets(ledgerUID, int32(idx)),
				})
			}
		}

		return nil
	}

	//
	//
	for i, ok := range soundAutographs {
		//
		idx := groupSignatureIndices[i]
		sig := endorse.Endorsements[idx]
		if !ok {
			return fmt.Errorf("REDACTED", idx, sig)
		}
		if certifiedAutographRepository != nil {
			certifiedAutographRepository.Add(string(sig.Autograph), AutographRepositoryItem{
				RatifierLocation: sig.RatifierLocation,
				BallotAttestOctets:    endorse.BallotAttestOctets(ledgerUID, int32(idx)),
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
	ledgerUID string,
	values *RatifierAssign,
	endorse *Endorse,
	pollingEnergyRequired int64,
	bypassSignature func(EndorseSignature) bool,
	numberSignature func(EndorseSignature) bool,
	numberAllEndorsements bool,
	scanUpByOrdinal bool,
	certifiedAutographRepository *AutographRepository,
) error {
	var (
		val                *Ratifier
		valueIdx             int32
		viewedValues           = make(map[int32]int, len(endorse.Endorsements))
		calculatedPollingEnergy int64
		ballotAttestOctets      []byte
	)
	for idx, endorseSignature := range endorse.Endorsements {
		if bypassSignature(endorseSignature) {
			continue
		}

		if endorseSignature.CertifySimple() != nil {
			return fmt.Errorf("REDACTED", val, idx)
		}

		//
		//
		if scanUpByOrdinal {
			val = values.Ratifiers[idx]
			if !bytes.Equal(val.Location, endorseSignature.RatifierLocation) {
				return fmt.Errorf("REDACTED",
					idx, val.Location, endorseSignature.RatifierLocation)
			}
		} else {
			valueIdx, val = values.FetchByLocation(endorseSignature.RatifierLocation)

			//
			//
			if val == nil {
				continue
			}

			//
			//
			if initialOrdinal, ok := viewedValues[valueIdx]; ok {
				momentOrdinal := idx
				return fmt.Errorf("REDACTED", val, initialOrdinal, momentOrdinal)
			}
			viewedValues[valueIdx] = idx
		}

		if val.PublicKey == nil {
			return fmt.Errorf("REDACTED", val, idx)
		}

		ballotAttestOctets = endorse.BallotAttestOctets(ledgerUID, int32(idx))

		repositoryKey, repositoryStrike := "REDACTED", false
		if certifiedAutographRepository != nil {
			repositoryKey = string(endorseSignature.Autograph)
			repositoryValue, signatureIsInRepository := certifiedAutographRepository.Get(repositoryKey)
			repositoryStrike = signatureIsInRepository && bytes.Equal(repositoryValue.RatifierLocation, val.PublicKey.Location()) && bytes.Equal(repositoryValue.BallotAttestOctets, ballotAttestOctets)
		}

		if !repositoryStrike {
			if !val.PublicKey.ValidateAutograph(ballotAttestOctets, endorseSignature.Autograph) {
				return fmt.Errorf("REDACTED", idx, endorseSignature.Autograph)
			}
			if certifiedAutographRepository != nil {
				certifiedAutographRepository.Add(repositoryKey, AutographRepositoryItem{
					RatifierLocation: val.PublicKey.Location(),
					BallotAttestOctets:    ballotAttestOctets,
				})
			}
		}

		//
		//
		if numberSignature(endorseSignature) {
			calculatedPollingEnergy += val.PollingEnergy
		}

		//
		if !numberAllEndorsements && calculatedPollingEnergy > pollingEnergyRequired {
			return nil
		}
	}

	if got, required := calculatedPollingEnergy, pollingEnergyRequired; got <= required {
		return ErrNoSufficientPollingEnergyAttested{Got: got, Required: required}
	}

	return nil
}

func validateSimpleValuesAndEndorse(values *RatifierAssign, endorse *Endorse, level int64, ledgerUID LedgerUID) error {
	if values == nil {
		return errors.New("REDACTED")
	}

	if endorse == nil {
		return errors.New("REDACTED")
	}

	if values.Volume() != len(endorse.Endorsements) {
		return cometfaults.NewErrCorruptEndorseEndorsements(values.Volume(), len(endorse.Endorsements))
	}

	//
	if level != endorse.Level {
		return cometfaults.NewErrCorruptEndorseLevel(level, endorse.Level)
	}
	if !ledgerUID.Matches(endorse.LedgerUID) {
		return fmt.Errorf("REDACTED",
			ledgerUID, endorse.LedgerUID)
	}

	return nil
}
