package status

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/gogoproto/proto"

	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
	"github.com/valkyrieworks/release"
)

//
var (
	statusKey = []byte("REDACTED")
)

//

//
//
//
//
var InitStatusRelease = cometstatus.Release{
	Agreement: cometrelease.Agreement{
		Ledger: release.LedgerProtocol,
		App:   0,
	},
	Software: release.TMCoreSemaphoreRev,
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
	Release cometstatus.Release

	//
	LedgerUID       string
	PrimaryLevel int64 //

	//
	FinalLedgerLevel int64
	FinalLedgerUID     kinds.LedgerUID
	FinalLedgerTime   time.Time

	//
	//
	//
	//
	//
	//
	FollowingRatifiers              *kinds.RatifierAssign
	Ratifiers                  *kinds.RatifierAssign
	FinalRatifiers              *kinds.RatifierAssign
	FinalLevelRatifiersModified int64

	//
	//
	AgreementOptions                  kinds.AgreementOptions
	FinalLevelAgreementOptionsModified int64

	//
	FinalOutcomesDigest []byte

	//
	ApplicationDigest []byte
}

//
func (status Status) Clone() Status {
	return Status{
		Release:       status.Release,
		LedgerUID:       status.LedgerUID,
		PrimaryLevel: status.PrimaryLevel,

		FinalLedgerLevel: status.FinalLedgerLevel,
		FinalLedgerUID:     status.FinalLedgerUID,
		FinalLedgerTime:   status.FinalLedgerTime,

		FollowingRatifiers:              status.FollowingRatifiers.Clone(),
		Ratifiers:                  status.Ratifiers.Clone(),
		FinalRatifiers:              status.FinalRatifiers.Clone(),
		FinalLevelRatifiersModified: status.FinalLevelRatifiersModified,

		AgreementOptions:                  status.AgreementOptions,
		FinalLevelAgreementOptionsModified: status.FinalLevelAgreementOptionsModified,

		ApplicationDigest: status.ApplicationDigest,

		FinalOutcomesDigest: status.FinalOutcomesDigest,
	}
}

//
func (status Status) Matches(status2 Status) bool {
	sbz, s2bz := status.Octets(), status2.Octets()
	return bytes.Equal(sbz, s2bz)
}

//
//
func (status Status) Octets() []byte {
	sm, err := status.ToSchema()
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
func (status Status) IsEmpty() bool {
	return status.Ratifiers == nil //
}

//
func (status *Status) ToSchema() (*cometstatus.Status, error) {
	if status == nil {
		return nil, errors.New("REDACTED")
	}

	sm := new(cometstatus.Status)

	sm.Release = status.Release
	sm.LedgerUID = status.LedgerUID
	sm.PrimaryLevel = status.PrimaryLevel
	sm.FinalLedgerLevel = status.FinalLedgerLevel

	sm.FinalLedgerUID = status.FinalLedgerUID.ToSchema()
	sm.FinalLedgerTime = status.FinalLedgerTime
	values, err := status.Ratifiers.ToSchema()
	if err != nil {
		return nil, err
	}
	sm.Ratifiers = values

	nValues, err := status.FollowingRatifiers.ToSchema()
	if err != nil {
		return nil, err
	}
	sm.FollowingRatifiers = nValues

	if status.FinalLedgerLevel >= 1 { //
		lValues, err := status.FinalRatifiers.ToSchema()
		if err != nil {
			return nil, err
		}
		sm.FinalRatifiers = lValues
	}

	sm.FinalLevelRatifiersModified = status.FinalLevelRatifiersModified
	sm.AgreementOptions = status.AgreementOptions.ToSchema()
	sm.FinalLevelAgreementOptionsModified = status.FinalLevelAgreementOptionsModified
	sm.FinalOutcomesDigest = status.FinalOutcomesDigest
	sm.ApplicationDigest = status.ApplicationDigest

	return sm, nil
}

//
func FromSchema(pb *cometstatus.Status) (*Status, error) { //
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	status := new(Status)

	status.Release = pb.Release
	status.LedgerUID = pb.LedgerUID
	status.PrimaryLevel = pb.PrimaryLevel

	bi, err := kinds.LedgerUIDFromSchema(&pb.FinalLedgerUID)
	if err != nil {
		return nil, err
	}
	status.FinalLedgerUID = *bi
	status.FinalLedgerLevel = pb.FinalLedgerLevel
	status.FinalLedgerTime = pb.FinalLedgerTime

	values, err := kinds.RatifierCollectionFromSchema(pb.Ratifiers)
	if err != nil {
		return nil, err
	}
	status.Ratifiers = values

	nValues, err := kinds.RatifierCollectionFromSchema(pb.FollowingRatifiers)
	if err != nil {
		return nil, err
	}
	status.FollowingRatifiers = nValues

	if status.FinalLedgerLevel >= 1 { //
		lValues, err := kinds.RatifierCollectionFromSchema(pb.FinalRatifiers)
		if err != nil {
			return nil, err
		}
		status.FinalRatifiers = lValues
	} else {
		status.FinalRatifiers = kinds.NewRatifierCollection(nil)
	}

	status.FinalLevelRatifiersModified = pb.FinalLevelRatifiersModified
	status.AgreementOptions = kinds.AgreementOptionsFromSchema(pb.AgreementOptions)
	status.FinalLevelAgreementOptionsModified = pb.FinalLevelAgreementOptionsModified
	status.FinalOutcomesDigest = pb.FinalOutcomesDigest
	status.ApplicationDigest = pb.ApplicationDigest

	return status, nil
}

//
//

//
//
//
func (status Status) CreateLedger(
	level int64,
	txs []kinds.Tx,
	finalEndorse *kinds.Endorse,
	proof []kinds.Proof,
	recommenderLocation []byte,
) (*kinds.Ledger, error) {

	//
	ledger := kinds.CreateLedger(level, txs, finalEndorse, proof)

	//
	var timestamp time.Time
	if level == status.PrimaryLevel {
		timestamp = status.FinalLedgerTime //
	} else {
		ts, err := MidpointTime(finalEndorse, status.FinalRatifiers)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		timestamp = ts
	}

	//
	ledger.Fill(
		status.Release.Agreement, status.LedgerUID,
		timestamp, status.FinalLedgerUID,
		status.Ratifiers.Digest(), status.FollowingRatifiers.Digest(),
		status.AgreementOptions.Digest(), status.ApplicationDigest, status.FinalOutcomesDigest,
		recommenderLocation,
	)

	return ledger, nil
}

//
//
//
//
func MidpointTime(endorse *kinds.Endorse, ratifiers *kinds.RatifierAssign) (time.Time, error) {
	scaledInstances := make([]*engineclock.ScaledTime, len(endorse.Endorsements))
	sumPollingEnergy := int64(0)

	for i, endorseSignature := range endorse.Endorsements {
		if endorseSignature.LedgerUIDMark == kinds.LedgerUIDMarkMissing {
			continue
		}
		_, ratifier := ratifiers.FetchByLocation(endorseSignature.RatifierLocation)
		//
		if ratifier == nil {
			return time.Time{}, fmt.Errorf("REDACTED",
				endorseSignature.RatifierLocation)
		}
		sumPollingEnergy += ratifier.PollingEnergy
		scaledInstances[i] = engineclock.NewScaledTime(endorseSignature.Timestamp, ratifier.PollingEnergy)
	}

	return engineclock.ScaledAverage(scaledInstances, sumPollingEnergy), nil
}

//
//

//
//
//
//
func CreateOriginStatusFromEntry(generatePaperEntry string) (Status, error) {
	generatePaper, err := CreateOriginPaperFromEntry(generatePaperEntry)
	if err != nil {
		return Status{}, err
	}
	return CreateOriginStatus(generatePaper)
}

//
func CreateOriginPaperFromEntry(generatePaperEntry string) (*kinds.OriginPaper, error) {
	generatePaperJSON, err := os.ReadFile(generatePaperEntry)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	generatePaper, err := kinds.OriginPaperFromJSON(generatePaperJSON)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return generatePaper, nil
}

//
func CreateOriginStatus(generatePaper *kinds.OriginPaper) (Status, error) {
	err := generatePaper.CertifyAndFinished()
	if err != nil {
		return Status{}, fmt.Errorf("REDACTED", err)
	}

	var ratifierCollection, followingRatifierCollection *kinds.RatifierAssign
	if generatePaper.Ratifiers == nil {
		ratifierCollection = kinds.NewRatifierCollection(nil)
		followingRatifierCollection = kinds.NewRatifierCollection(nil)
	} else {
		ratifiers := make([]*kinds.Ratifier, len(generatePaper.Ratifiers))
		for i, val := range generatePaper.Ratifiers {
			ratifiers[i] = kinds.NewRatifier(val.PublicKey, val.Energy)
		}
		ratifierCollection = kinds.NewRatifierCollection(ratifiers)
		followingRatifierCollection = kinds.NewRatifierCollection(ratifiers).CloneAugmentRecommenderUrgency(1)
	}

	return Status{
		Release:       InitStatusRelease,
		LedgerUID:       generatePaper.LedgerUID,
		PrimaryLevel: generatePaper.PrimaryLevel,

		FinalLedgerLevel: 0,
		FinalLedgerUID:     kinds.LedgerUID{},
		FinalLedgerTime:   generatePaper.OriginMoment,

		FollowingRatifiers:              followingRatifierCollection,
		Ratifiers:                  ratifierCollection,
		FinalRatifiers:              kinds.NewRatifierCollection(nil),
		FinalLevelRatifiersModified: generatePaper.PrimaryLevel,

		AgreementOptions:                  *generatePaper.AgreementOptions,
		FinalLevelAgreementOptionsModified: generatePaper.PrimaryLevel,

		ApplicationDigest: generatePaper.ApplicationDigest,
	}, nil
}
