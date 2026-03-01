package status

import (
	"errors"
	"fmt"

	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//
//
//
func Revert(bs LedgerDepot, ss Depot, discardLedger bool) (int64, []byte, error) {
	unfitStatus, err := ss.Fetch()
	if err != nil {
		return -1, nil, err
	}
	if unfitStatus.EqualsBlank() {
		return -1, nil, errors.New("REDACTED")
	}

	altitude := bs.Altitude()

	//
	//
	//
	if altitude == unfitStatus.FinalLedgerAltitude+1 {
		if discardLedger {
			if err := bs.EraseNewestLedger(); err != nil {
				return -1, nil, fmt.Errorf("REDACTED", err)
			}
		}
		return unfitStatus.FinalLedgerAltitude, unfitStatus.PlatformDigest, nil
	}

	//
	//
	if altitude != unfitStatus.FinalLedgerAltitude {
		return -1, nil, fmt.Errorf("REDACTED",
			unfitStatus.FinalLedgerAltitude, altitude)
	}

	//
	revertAltitude := unfitStatus.FinalLedgerAltitude - 1
	revertLedger := bs.FetchLedgerSummary(revertAltitude)
	if revertLedger == nil {
		return -1, nil, fmt.Errorf("REDACTED", revertAltitude)
	}
	//
	//
	newestLedger := bs.FetchLedgerSummary(unfitStatus.FinalLedgerAltitude)
	if newestLedger == nil {
		return -1, nil, fmt.Errorf("REDACTED", unfitStatus.FinalLedgerAltitude)
	}

	precedingFinalAssessorAssign, err := ss.FetchAssessors(revertAltitude)
	if err != nil {
		return -1, nil, err
	}

	precedingParameters, err := ss.FetchAgreementParameters(revertAltitude + 1)
	if err != nil {
		return -1, nil, err
	}

	followingAltitude := revertAltitude + 1
	itemAlterationAltitude := unfitStatus.FinalAltitudeAssessorsAltered
	//
	if itemAlterationAltitude > followingAltitude+1 {
		itemAlterationAltitude = followingAltitude + 1
	}

	parametersAlterationAltitude := unfitStatus.FinalAltitudeAgreementParametersAltered
	//
	if parametersAlterationAltitude > revertAltitude {
		parametersAlterationAltitude = revertAltitude + 1
	}

	//
	morphedRearStatus := Status{
		Edition: strongstatus.Edition{
			Agreement: strongmindedition.Agreement{
				Ledger: edition.LedgerScheme,
				App:   precedingParameters.Edition.App,
			},
			Package: edition.TEMPBaseSemaphoreEdtn,
		},
		//
		SuccessionUUID:       unfitStatus.SuccessionUUID,
		PrimaryAltitude: unfitStatus.PrimaryAltitude,

		FinalLedgerAltitude: revertLedger.Heading.Altitude,
		FinalLedgerUUID:     revertLedger.LedgerUUID,
		FinalLedgerMoment:   revertLedger.Heading.Moment,

		FollowingAssessors:              unfitStatus.Assessors,
		Assessors:                  unfitStatus.FinalAssessors,
		FinalAssessors:              precedingFinalAssessorAssign,
		FinalAltitudeAssessorsAltered: itemAlterationAltitude,

		AgreementSettings:                  precedingParameters,
		FinalAltitudeAgreementParametersAltered: parametersAlterationAltitude,

		FinalOutcomesDigest: newestLedger.Heading.FinalOutcomesDigest,
		PlatformDigest:         newestLedger.Heading.PlatformDigest,
	}

	//
	//
	//
	if err := ss.Persist(morphedRearStatus); err != nil {
		return -1, nil, fmt.Errorf("REDACTED", err)
	}

	//
	//
	if discardLedger {
		if err := bs.EraseNewestLedger(); err != nil {
			return -1, nil, fmt.Errorf("REDACTED", err)
		}
	}

	return morphedRearStatus.FinalLedgerAltitude, morphedRearStatus.PlatformDigest, nil
}
