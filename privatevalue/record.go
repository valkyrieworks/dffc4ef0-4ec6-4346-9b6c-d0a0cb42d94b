package privatevalue

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/scratchfile"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

//
const (
	phaseVoid      int8 = 0 //
	phaseNominate   int8 = 1
	phasePreballot   int8 = 2
	phasePreendorse int8 = 3
)

//
func ballotTowardPhase(ballot *commitchema.Ballot) int8 {
	switch ballot.Kind {
	case commitchema.PreballotKind:
		return phasePreballot
	case commitchema.PreendorseKind:
		return phasePreendorse
	default:
		panic(fmt.Sprintf("REDACTED", ballot.Kind))
	}
}

//

//
type RecordPRVToken struct {
	Location kinds.Location  `json:"location"`
	PublicToken  security.PublicToken  `json:"public_token"`
	PrivateToken security.PrivateToken `json:"private_token"`

	recordRoute string
}

//
func (prvToken RecordPRVToken) Persist() {
	outputRecord := prvToken.recordRoute
	if outputRecord == "REDACTED" {
		panic("REDACTED")
	}

	jsnOctets, err := strongmindjson.SerializeRecess(prvToken, "REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}

	if err := scratchfile.PersistRecordIndivisible(outputRecord, jsnOctets, 0o600); err != nil {
		panic(err)
	}
}

//

//
type RecordPRVFinalAttestStatus struct {
	Altitude    int64             `json:"altitude"`
	Iteration     int32             `json:"iteration"`
	Phase      int8              `json:"phase"`
	Notation []byte            `json:"signing,omitempty"`
	AttestOctets tendermintoctets.HexadecimalOctets `json:"attestoctets,omitempty"`

	recordRoute string
}

func (lss *RecordPRVFinalAttestStatus) restore() {
	lss.Altitude = 0
	lss.Iteration = 0
	lss.Phase = 0
	lss.Notation = nil
	lss.AttestOctets = nil
}

//
//
//
//
//
//
//
func (lss *RecordPRVFinalAttestStatus) InspectHours(altitude int64, iteration int32, phase int8) (bool, error) {
	if lss.Altitude > altitude {
		return false, fmt.Errorf("REDACTED", altitude, lss.Altitude)
	}

	if lss.Altitude == altitude {
		if lss.Iteration > iteration {
			return false, fmt.Errorf("REDACTED", altitude, iteration, lss.Iteration)
		}

		if lss.Iteration == iteration {
			if lss.Phase > phase {
				return false, fmt.Errorf(
					"REDACTED",
					altitude,
					iteration,
					phase,
					lss.Phase,
				)
			} else if lss.Phase == phase {
				if lss.AttestOctets != nil {
					if lss.Notation == nil {
						panic("REDACTED")
					}
					return true, nil
				}
				return false, errors.New("REDACTED")
			}
		}
	}
	return false, nil
}

//
func (lss *RecordPRVFinalAttestStatus) Persist() {
	outputRecord := lss.recordRoute
	if outputRecord == "REDACTED" {
		panic("REDACTED")
	}
	jsnOctets, err := strongmindjson.SerializeRecess(lss, "REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	err = scratchfile.PersistRecordIndivisible(outputRecord, jsnOctets, 0o600)
	if err != nil {
		panic(err)
	}
}

//

//
//
//
//
//
type RecordPRV struct {
	Key           RecordPRVToken
	FinalAttestStatus RecordPRVFinalAttestStatus
}

//
func FreshRecordPRV(privateToken security.PrivateToken, tokenRecordRoute, statusRecordRoute string) *RecordPRV {
	return &RecordPRV{
		Key: RecordPRVToken{
			Location:  privateToken.PublicToken().Location(),
			PublicToken:   privateToken.PublicToken(),
			PrivateToken:  privateToken,
			recordRoute: tokenRecordRoute,
		},
		FinalAttestStatus: RecordPRVFinalAttestStatus{
			Phase:     phaseVoid,
			recordRoute: statusRecordRoute,
		},
	}
}

//
//
func ProduceRecordPRV(tokenRecordRoute, statusRecordRoute string) *RecordPRV {
	return FreshRecordPRV(edwards25519.ProducePrivateToken(), tokenRecordRoute, statusRecordRoute)
}

//
//
//
func FetchRecordPRV(tokenRecordRoute, statusRecordRoute string) *RecordPRV {
	return fetchRecordPRV(tokenRecordRoute, statusRecordRoute, true)
}

//
//
func FetchRecordPRVBlankStatus(tokenRecordRoute, statusRecordRoute string) *RecordPRV {
	return fetchRecordPRV(tokenRecordRoute, statusRecordRoute, false)
}

//
func fetchRecordPRV(tokenRecordRoute, statusRecordRoute string, fetchStatus bool) *RecordPRV {
	tokenJSNOctets, err := os.ReadFile(tokenRecordRoute)
	if err != nil {
		strongos.Quit(err.Error())
	}
	prvToken := RecordPRVToken{}
	err = strongmindjson.Decode(tokenJSNOctets, &prvToken)
	if err != nil {
		strongos.Quit(fmt.Sprintf("REDACTED", tokenRecordRoute, err))
	}

	//
	prvToken.PublicToken = prvToken.PrivateToken.PublicToken()
	prvToken.Location = prvToken.PublicToken.Location()
	prvToken.recordRoute = tokenRecordRoute

	prvStatus := RecordPRVFinalAttestStatus{}

	if fetchStatus {
		statusJSNOctets, err := os.ReadFile(statusRecordRoute)
		if err != nil {
			strongos.Quit(err.Error())
		}
		err = strongmindjson.Decode(statusJSNOctets, &prvStatus)
		if err != nil {
			strongos.Quit(fmt.Sprintf("REDACTED", statusRecordRoute, err))
		}
	}

	prvStatus.recordRoute = statusRecordRoute

	return &RecordPRV{
		Key:           prvToken,
		FinalAttestStatus: prvStatus,
	}
}

//
//
func FetchEitherProduceRecordPRV(tokenRecordRoute, statusRecordRoute string) *RecordPRV {
	var pv *RecordPRV
	if strongos.RecordPresent(tokenRecordRoute) {
		pv = FetchRecordPRV(tokenRecordRoute, statusRecordRoute)
	} else {
		pv = ProduceRecordPRV(tokenRecordRoute, statusRecordRoute)
		pv.Persist()
	}
	return pv
}

//
//
func (pv *RecordPRV) ObtainLocator() kinds.Location {
	return pv.Key.Location
}

//
//
func (pv *RecordPRV) ObtainPublicToken() (security.PublicToken, error) {
	return pv.Key.PublicToken, nil
}

//
//
func (pv *RecordPRV) AttestBallot(successionUUID string, ballot *commitchema.Ballot) error {
	if err := pv.attestBallot(successionUUID, ballot); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
//
func (pv *RecordPRV) AttestNomination(successionUUID string, nomination *commitchema.Nomination) error {
	if err := pv.attestNomination(successionUUID, nomination); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (pv *RecordPRV) Persist() {
	pv.Key.Persist()
	pv.FinalAttestStatus.Persist()
}

//
//
func (pv *RecordPRV) Restore() {
	pv.FinalAttestStatus.restore()
	pv.Persist()
}

//
func (pv *RecordPRV) Text() string {
	return fmt.Sprintf(
		"REDACTED",
		pv.ObtainLocator(),
		pv.FinalAttestStatus.Altitude,
		pv.FinalAttestStatus.Iteration,
		pv.FinalAttestStatus.Phase,
	)
}

//

//
//
//
//
func (pv *RecordPRV) attestBallot(successionUUID string, ballot *commitchema.Ballot) error {
	altitude, iteration, phase := ballot.Altitude, ballot.Iteration, ballotTowardPhase(ballot)

	lss := pv.FinalAttestStatus

	identicalHours, err := lss.InspectHours(altitude, iteration, phase)
	if err != nil {
		return err
	}

	attestOctets := kinds.BallotAttestOctets(successionUUID, ballot)

	//
	//
	//
	//
	//
	var addnSignature []byte
	if ballot.Kind == commitchema.PreendorseKind && !kinds.SchemaLedgerUUIDEqualsVoid(&ballot.LedgerUUID) {
		addnAttestOctets := kinds.BallotAdditionAttestOctets(successionUUID, ballot)
		addnSignature, err = pv.Key.PrivateToken.Attest(addnAttestOctets)
		if err != nil {
			return err
		}
	} else if len(ballot.Addition) > 0 {
		return errors.New("REDACTED")
	}

	//
	//
	//
	//
	//
	if identicalHours {
		if bytes.Equal(attestOctets, lss.AttestOctets) {
			ballot.Notation = lss.Notation
		} else if timestamp, ok := inspectBallotsSolelyDeviateViaTimestamp(lss.AttestOctets, attestOctets); ok {
			//
			//
			ballot.Timestamp = timestamp
			ballot.Notation = lss.Notation
		} else {
			err = fmt.Errorf("REDACTED")
		}

		ballot.AdditionNotation = addnSignature

		return err
	}

	//
	sig, err := pv.Key.PrivateToken.Attest(attestOctets)
	if err != nil {
		return err
	}
	pv.persistNotated(altitude, iteration, phase, attestOctets, sig)
	ballot.Notation = sig
	ballot.AdditionNotation = addnSignature

	return nil
}

//
//
//
func (pv *RecordPRV) attestNomination(successionUUID string, nomination *commitchema.Nomination) error {
	altitude, iteration, phase := nomination.Altitude, nomination.Iteration, phaseNominate

	lss := pv.FinalAttestStatus

	identicalHours, err := lss.InspectHours(altitude, iteration, phase)
	if err != nil {
		return err
	}

	attestOctets := kinds.NominationAttestOctets(successionUUID, nomination)

	//
	//
	//
	//
	//
	if identicalHours {
		if bytes.Equal(attestOctets, lss.AttestOctets) {
			nomination.Notation = lss.Notation
		} else if timestamp, ok := inspectNominationsSolelyDeviateViaTimestamp(lss.AttestOctets, attestOctets); ok {
			nomination.Timestamp = timestamp
			nomination.Notation = lss.Notation
		} else {
			err = fmt.Errorf("REDACTED")
		}
		return err
	}

	//
	sig, err := pv.Key.PrivateToken.Attest(attestOctets)
	if err != nil {
		return err
	}
	pv.persistNotated(altitude, iteration, phase, attestOctets, sig)
	nomination.Notation = sig
	return nil
}

//
func (pv *RecordPRV) persistNotated(altitude int64, iteration int32, phase int8,
	attestOctets []byte, sig []byte,
) {
	pv.FinalAttestStatus.Altitude = altitude
	pv.FinalAttestStatus.Iteration = iteration
	pv.FinalAttestStatus.Phase = phase
	pv.FinalAttestStatus.Notation = sig
	pv.FinalAttestStatus.AttestOctets = attestOctets
	pv.FinalAttestStatus.Persist()
}

//

//
//
//
//
func inspectBallotsSolelyDeviateViaTimestamp(finalAttestOctets, freshAttestOctets []byte) (time.Time, bool) {
	var finalBallot, freshBallot commitchema.StandardBallot
	if err := protocolio.DecodeSeparated(finalAttestOctets, &finalBallot); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if err := protocolio.DecodeSeparated(freshAttestOctets, &freshBallot); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	finalMoment := finalBallot.Timestamp
	//
	now := committime.Now()
	finalBallot.Timestamp = now
	freshBallot.Timestamp = now

	return finalMoment, proto.Equal(&freshBallot, &finalBallot)
}

//
//
func inspectNominationsSolelyDeviateViaTimestamp(finalAttestOctets, freshAttestOctets []byte) (time.Time, bool) {
	var finalNomination, freshNomination commitchema.StandardNomination
	if err := protocolio.DecodeSeparated(finalAttestOctets, &finalNomination); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if err := protocolio.DecodeSeparated(freshAttestOctets, &freshNomination); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	finalMoment := finalNomination.Timestamp
	//
	now := committime.Now()
	finalNomination.Timestamp = now
	freshNomination.Timestamp = now

	return finalMoment, proto.Equal(&freshNomination, &finalNomination)
}
