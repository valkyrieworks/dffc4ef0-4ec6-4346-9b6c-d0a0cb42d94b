package status

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/gogoproto/proto"

	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//
var (
	statusToken = []byte("REDACTED")
)

//

//
//
//
//
var InitializeStatusEdition = strongstatus.Edition{
	Agreement: strongmindedition.Agreement{
		Ledger: edition.LedgerScheme,
		App:   0,
	},
	Package: edition.TEMPBaseSemaphoreEdtn,
}

//

//
//
//
//
//
//
//
type Status struct {
	Edition strongstatus.Edition

	//
	SuccessionUUID       string
	PrimaryAltitude int64 //

	//
	FinalLedgerAltitude int64
	FinalLedgerUUID     kinds.LedgerUUID
	FinalLedgerMoment   time.Time

	//
	//
	//
	//
	//
	//
	FollowingAssessors              *kinds.AssessorAssign
	Assessors                  *kinds.AssessorAssign
	FinalAssessors              *kinds.AssessorAssign
	FinalAltitudeAssessorsAltered int64

	//
	//
	AgreementSettings                  kinds.AgreementSettings
	FinalAltitudeAgreementParametersAltered int64

	//
	FinalOutcomesDigest []byte

	//
	PlatformDigest []byte
}

//
func (status Status) Duplicate() Status {
	return Status{
		Edition:       status.Edition,
		SuccessionUUID:       status.SuccessionUUID,
		PrimaryAltitude: status.PrimaryAltitude,

		FinalLedgerAltitude: status.FinalLedgerAltitude,
		FinalLedgerUUID:     status.FinalLedgerUUID,
		FinalLedgerMoment:   status.FinalLedgerMoment,

		FollowingAssessors:              status.FollowingAssessors.Duplicate(),
		Assessors:                  status.Assessors.Duplicate(),
		FinalAssessors:              status.FinalAssessors.Duplicate(),
		FinalAltitudeAssessorsAltered: status.FinalAltitudeAssessorsAltered,

		AgreementSettings:                  status.AgreementSettings,
		FinalAltitudeAgreementParametersAltered: status.FinalAltitudeAgreementParametersAltered,

		PlatformDigest: status.PlatformDigest,

		FinalOutcomesDigest: status.FinalOutcomesDigest,
	}
}

//
func (status Status) Matches(status2 Status) bool {
	sbz, s2by := status.Octets(), status2.Octets()
	return bytes.Equal(sbz, s2by)
}

//
//
func (status Status) Octets() []byte {
	sm, err := status.TowardSchema()
	if err != nil {
		panic(err)
	}
	bz, err := proto.Marshal(sm)
	if err != nil {
		panic(err)
	}
	return bz
}

//
func (status Status) EqualsBlank() bool {
	return status.Assessors == nil //
}

//
func (status *Status) TowardSchema() (*strongstatus.Status, error) {
	if status == nil {
		return nil, errors.New("REDACTED")
	}

	sm := new(strongstatus.Status)

	sm.Edition = status.Edition
	sm.SuccessionUUID = status.SuccessionUUID
	sm.PrimaryAltitude = status.PrimaryAltitude
	sm.FinalLedgerAltitude = status.FinalLedgerAltitude

	sm.FinalLedgerUUID = status.FinalLedgerUUID.TowardSchema()
	sm.FinalLedgerMoment = status.FinalLedgerMoment
	values, err := status.Assessors.TowardSchema()
	if err != nil {
		return nil, err
	}
	sm.Assessors = values

	nthValues, err := status.FollowingAssessors.TowardSchema()
	if err != nil {
		return nil, err
	}
	sm.FollowingAssessors = nthValues

	if status.FinalLedgerAltitude >= 1 { //
		lengthValues, err := status.FinalAssessors.TowardSchema()
		if err != nil {
			return nil, err
		}
		sm.FinalAssessors = lengthValues
	}

	sm.FinalAltitudeAssessorsAltered = status.FinalAltitudeAssessorsAltered
	sm.AgreementSettings = status.AgreementSettings.TowardSchema()
	sm.FinalAltitudeAgreementParametersAltered = status.FinalAltitudeAgreementParametersAltered
	sm.FinalOutcomesDigest = status.FinalOutcomesDigest
	sm.PlatformDigest = status.PlatformDigest

	return sm, nil
}

//
func OriginatingSchema(pb *strongstatus.Status) (*Status, error) { //
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	status := new(Status)

	status.Edition = pb.Edition
	status.SuccessionUUID = pb.SuccessionUUID
	status.PrimaryAltitude = pb.PrimaryAltitude

	bi, err := kinds.LedgerUUIDOriginatingSchema(&pb.FinalLedgerUUID)
	if err != nil {
		return nil, err
	}
	status.FinalLedgerUUID = *bi
	status.FinalLedgerAltitude = pb.FinalLedgerAltitude
	status.FinalLedgerMoment = pb.FinalLedgerMoment

	values, err := kinds.AssessorAssignOriginatingSchema(pb.Assessors)
	if err != nil {
		return nil, err
	}
	status.Assessors = values

	nthValues, err := kinds.AssessorAssignOriginatingSchema(pb.FollowingAssessors)
	if err != nil {
		return nil, err
	}
	status.FollowingAssessors = nthValues

	if status.FinalLedgerAltitude >= 1 { //
		lengthValues, err := kinds.AssessorAssignOriginatingSchema(pb.FinalAssessors)
		if err != nil {
			return nil, err
		}
		status.FinalAssessors = lengthValues
	} else {
		status.FinalAssessors = kinds.FreshAssessorAssign(nil)
	}

	status.FinalAltitudeAssessorsAltered = pb.FinalAltitudeAssessorsAltered
	status.AgreementSettings = kinds.AgreementParametersOriginatingSchema(pb.AgreementSettings)
	status.FinalAltitudeAgreementParametersAltered = pb.FinalAltitudeAgreementParametersAltered
	status.FinalOutcomesDigest = pb.FinalOutcomesDigest
	status.PlatformDigest = pb.PlatformDigest

	return status, nil
}

//
//

//
//
//
func (status Status) CreateLedger(
	altitude int64,
	txs []kinds.Tx,
	finalEndorse *kinds.Endorse,
	proof []kinds.Proof,
	nominatorLocator []byte,
) (*kinds.Ledger, error) {

	//
	ledger := kinds.CreateLedger(altitude, txs, finalEndorse, proof)

	//
	var timestamp time.Time
	if altitude == status.PrimaryAltitude {
		timestamp = status.FinalLedgerMoment //
	} else {
		ts, err := AverageMoment(finalEndorse, status.FinalAssessors)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		timestamp = ts
	}

	//
	ledger.Inhabit(
		status.Edition.Agreement, status.SuccessionUUID,
		timestamp, status.FinalLedgerUUID,
		status.Assessors.Digest(), status.FollowingAssessors.Digest(),
		status.AgreementSettings.Digest(), status.PlatformDigest, status.FinalOutcomesDigest,
		nominatorLocator,
	)

	return ledger, nil
}

//
func (status Status) CertifyLedger(ledger *kinds.Ledger) error {
	return certifyLedger(status, ledger)
}

//
//
//
//
func AverageMoment(endorse *kinds.Endorse, assessors *kinds.AssessorAssign) (time.Time, error) {
	burdenedMultiples := make([]*committime.BurdenedMoment, len(endorse.Notations))
	sumBallotingPotency := int64(0)

	for i, endorseSignature := range endorse.Notations {
		if endorseSignature.LedgerUUIDMarker == kinds.LedgerUUIDMarkerMissing {
			continue
		}
		_, assessor := assessors.ObtainViaLocation(endorseSignature.AssessorLocation)
		//
		if assessor == nil {
			return time.Time{}, fmt.Errorf("REDACTED",
				endorseSignature.AssessorLocation)
		}
		sumBallotingPotency += assessor.BallotingPotency
		burdenedMultiples[i] = committime.FreshBurdenedMoment(endorseSignature.Timestamp, assessor.BallotingPotency)
	}

	return committime.BurdenedAverage(burdenedMultiples, sumBallotingPotency), nil
}

//
//

//
//
//
//
func CreateInaugurationStatusOriginatingRecord(producePaperRecord string) (Status, error) {
	producePaper, err := CreateInaugurationPaperOriginatingRecord(producePaperRecord)
	if err != nil {
		return Status{}, err
	}
	return CreateInaugurationStatus(producePaper)
}

//
func CreateInaugurationPaperOriginatingRecord(producePaperRecord string) (*kinds.OriginPaper, error) {
	producePaperJSN, err := os.ReadFile(producePaperRecord)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	producePaper, err := kinds.InaugurationPaperOriginatingJSN(producePaperJSN)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return producePaper, nil
}

//
func CreateInaugurationStatus(producePaper *kinds.OriginPaper) (Status, error) {
	err := producePaper.CertifyAlsoFinish()
	if err != nil {
		return Status{}, fmt.Errorf("REDACTED", err)
	}

	var assessorAssign, followingAssessorCollection *kinds.AssessorAssign
	if producePaper.Assessors == nil {
		assessorAssign = kinds.FreshAssessorAssign(nil)
		followingAssessorCollection = kinds.FreshAssessorAssign(nil)
	} else {
		assessors := make([]*kinds.Assessor, len(producePaper.Assessors))
		for i, val := range producePaper.Assessors {
			assessors[i] = kinds.FreshAssessor(val.PublicToken, val.Potency)
		}
		assessorAssign = kinds.FreshAssessorAssign(assessors)
		followingAssessorCollection = kinds.FreshAssessorAssign(assessors).DuplicateAdvanceNominatorUrgency(1)
	}

	return Status{
		Edition:       InitializeStatusEdition,
		SuccessionUUID:       producePaper.SuccessionUUID,
		PrimaryAltitude: producePaper.PrimaryAltitude,

		FinalLedgerAltitude: 0,
		FinalLedgerUUID:     kinds.LedgerUUID{},
		FinalLedgerMoment:   producePaper.OriginMoment,

		FollowingAssessors:              followingAssessorCollection,
		Assessors:                  assessorAssign,
		FinalAssessors:              kinds.FreshAssessorAssign(nil),
		FinalAltitudeAssessorsAltered: producePaper.PrimaryAltitude,

		AgreementSettings:                  *producePaper.AgreementSettings,
		FinalAltitudeAgreementParametersAltered: producePaper.PrimaryAltitude,

		PlatformDigest: producePaper.PlatformDigest,
	}, nil
}
