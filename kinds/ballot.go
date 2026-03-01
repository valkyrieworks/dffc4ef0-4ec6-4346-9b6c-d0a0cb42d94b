package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

const (
	voidBallotTxt string = "REDACTED"

	//
	MaximumBallotAdditionExtent int = 1024 * 1024
)

var (
	FaultBallotUnforeseenPhase            = errors.New("REDACTED")
	FaultBallotUnfitAssessorPosition     = errors.New("REDACTED")
	FaultBallotUnfitAssessorLocator   = errors.New("REDACTED")
	FaultBallotUnfitSigning          = errors.New("REDACTED")
	FaultBallotUnfitLedgerDigest          = errors.New("REDACTED")
	FaultBallotUnCertainNotation = errors.New("REDACTED")
	FaultBallotVoid                       = errors.New("REDACTED")
	FaultBallotAdditionMissing           = errors.New("REDACTED")
	FaultUnfitBallotAddition          = errors.New("REDACTED")
)

type FaultBallotDiscordantBallots struct {
	BallotAN *Ballot
	BallotBYTE *Ballot
}

func (err *FaultBallotDiscordantBallots) Failure() string {
	return fmt.Sprintf("REDACTED", err.BallotAN.AssessorLocation)
}

func FreshDiscordantBallotFailure(ballot1, ballot2 *Ballot) *FaultBallotDiscordantBallots {
	return &FaultBallotDiscordantBallots{
		BallotAN: ballot1,
		BallotBYTE: ballot2,
	}
}

//
type FaultBallotAdditionUnfit struct {
	AddnSigning []byte
}

func (err *FaultBallotAdditionUnfit) Failure() string {
	return fmt.Sprintf("REDACTED", err.AddnSigning)
}

//
type Location = security.Location

//
//
type Ballot struct {
	Kind               commitchema.AttestedSignalKind `json:"kind"`
	Altitude             int64                  `json:"altitude"`
	Iteration              int32                  `json:"iteration"`    //
	LedgerUUID            LedgerUUID                `json:"ledger_uuid"` //
	Timestamp          time.Time              `json:"timestamp"`
	AssessorLocation   Location                `json:"assessor_location"`
	AssessorOrdinal     int32                  `json:"assessor_position"`
	Notation          []byte                 `json:"signing"`
	Addition          []byte                 `json:"addition"`
	AdditionNotation []byte                 `json:"addition_signing"`
}

//
//
//
//
func BallotOriginatingSchema(pv *commitchema.Ballot) (*Ballot, error) {
	ledgerUUID, err := LedgerUUIDOriginatingSchema(&pv.LedgerUUID)
	if err != nil {
		return nil, err
	}

	return &Ballot{
		Kind:               pv.Kind,
		Altitude:             pv.Altitude,
		Iteration:              pv.Iteration,
		LedgerUUID:            *ledgerUUID,
		Timestamp:          pv.Timestamp,
		AssessorLocation:   pv.AssessorLocation,
		AssessorOrdinal:     pv.AssessorOrdinal,
		Notation:          pv.Notation,
		Addition:          pv.Addition,
		AdditionNotation: pv.AdditionNotation,
	}, nil
}

//
func (ballot *Ballot) EndorseSignature() EndorseSignature {
	if ballot == nil {
		return FreshEndorseSignatureMissing()
	}

	var ledgerUUIDMarker LedgerUUIDMarker
	switch {
	case ballot.LedgerUUID.EqualsFinish():
		ledgerUUIDMarker = LedgerUUIDMarkerEndorse
	case ballot.LedgerUUID.EqualsNull():
		ledgerUUIDMarker = LedgerUUIDMarkerVoid
	default:
		panic(fmt.Sprintf("REDACTED", ballot))
	}

	return EndorseSignature{
		LedgerUUIDMarker:      ledgerUUIDMarker,
		AssessorLocation: ballot.AssessorLocation,
		Timestamp:        ballot.Timestamp,
		Notation:        ballot.Notation,
	}
}

//
//
//
func (ballot *Ballot) ExpandedEndorseSignature() ExpandedEndorseSignature {
	if ballot == nil {
		return FreshExpandedEndorseSignatureMissing()
	}

	return ExpandedEndorseSignature{
		EndorseSignature:          ballot.EndorseSignature(),
		Addition:          ballot.Addition,
		AdditionNotation: ballot.AdditionNotation,
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
func BallotAttestOctets(successionUUID string, ballot *commitchema.Ballot) []byte {
	pb := NormalizeBallot(successionUUID, ballot)
	bz, err := protocolio.SerializeSeparated(&pb)
	if err != nil {
		panic(err)
	}

	return bz
}

//
//
//
//
//
func BallotAdditionAttestOctets(successionUUID string, ballot *commitchema.Ballot) []byte {
	pb := NormalizeBallotAddition(successionUUID, ballot)
	bz, err := protocolio.SerializeSeparated(&pb)
	if err != nil {
		panic(err)
	}

	return bz
}

func (ballot *Ballot) Duplicate() *Ballot {
	ballotDuplicate := *ballot
	return &ballotDuplicate
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
func (ballot *Ballot) Text() string {
	if ballot == nil {
		return voidBallotTxt
	}

	var kindText string
	switch ballot.Kind {
	case commitchema.PreballotKind:
		kindText = "REDACTED"
	case commitchema.PreendorseKind:
		kindText = "REDACTED"
	default:
		panic("REDACTED")
	}

	return fmt.Sprintf("REDACTED",
		ballot.AssessorOrdinal,
		tendermintoctets.Identifier(ballot.AssessorLocation),
		ballot.Altitude,
		ballot.Iteration,
		ballot.Kind,
		kindText,
		tendermintoctets.Identifier(ballot.LedgerUUID.Digest),
		tendermintoctets.Identifier(ballot.Notation),
		tendermintoctets.Identifier(ballot.Addition),
		StandardMoment(ballot.Timestamp),
	)
}

func (ballot *Ballot) validateAlsoYieldSchema(successionUUID string, publicToken security.PublicToken) (*commitchema.Ballot, error) {
	if !bytes.Equal(publicToken.Location(), ballot.AssessorLocation) {
		return nil, FaultBallotUnfitAssessorLocator
	}
	v := ballot.TowardSchema()
	if !publicToken.ValidateNotation(BallotAttestOctets(successionUUID, v), ballot.Notation) {
		return nil, FaultBallotUnfitSigning
	}
	return v, nil
}

//
//
//
func (ballot *Ballot) Validate(successionUUID string, publicToken security.PublicToken) error {
	_, err := ballot.validateAlsoYieldSchema(successionUUID, publicToken)
	return err
}

//
//
//
//
func (ballot *Ballot) ValidateBallotAlsoAddition(successionUUID string, publicToken security.PublicToken) error {
	v, err := ballot.validateAlsoYieldSchema(successionUUID, publicToken)
	if err != nil {
		return err
	}
	//
	if ballot.Kind == commitchema.PreendorseKind && !SchemaLedgerUUIDEqualsVoid(&v.LedgerUUID) {
		if len(ballot.AdditionNotation) == 0 {
			return errors.New("REDACTED")
		}

		addnAttestOctets := BallotAdditionAttestOctets(successionUUID, v)
		if !publicToken.ValidateNotation(addnAttestOctets, ballot.AdditionNotation) {
			return FaultBallotUnfitSigning
		}
	}
	return nil
}

//
//
func (ballot *Ballot) ValidateAddition(successionUUID string, publicToken security.PublicToken) error {
	if ballot.Kind != commitchema.PreendorseKind || ballot.LedgerUUID.EqualsNull() {
		return nil
	}
	v := ballot.TowardSchema()
	addnAttestOctets := BallotAdditionAttestOctets(successionUUID, v)
	if !publicToken.ValidateNotation(addnAttestOctets, ballot.AdditionNotation) {
		return FaultBallotUnfitSigning
	}
	return nil
}

//
//
//
func (ballot *Ballot) CertifyFundamental() error {
	if !EqualsBallotKindSound(ballot.Kind) {
		return errors.New("REDACTED")
	}

	if ballot.Altitude <= 0 {
		return errors.New("REDACTED")
	}

	if ballot.Iteration < 0 {
		return errors.New("REDACTED")
	}

	//

	if err := ballot.LedgerUUID.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	//
	if !ballot.LedgerUUID.EqualsNull() && !ballot.LedgerUUID.EqualsFinish() {
		return fmt.Errorf("REDACTED", ballot.LedgerUUID)
	}

	if len(ballot.AssessorLocation) != security.LocatorExtent {
		return fmt.Errorf("REDACTED",
			security.LocatorExtent,
			len(ballot.AssessorLocation),
		)
	}
	if ballot.AssessorOrdinal < 0 {
		return errors.New("REDACTED")
	}
	if len(ballot.Notation) == 0 {
		return errors.New("REDACTED")
	}

	if len(ballot.Notation) > MaximumSigningExtent {
		return fmt.Errorf("REDACTED", MaximumSigningExtent)
	}

	//
	//
	//
	if ballot.Kind != commitchema.PreendorseKind || ballot.LedgerUUID.EqualsNull() {
		if len(ballot.Addition) > 0 {
			return fmt.Errorf(
				"REDACTED",
				ballot.Kind, ballot.LedgerUUID.EqualsNull(),
			)
		}
		if len(ballot.AdditionNotation) > 0 {
			return errors.New("REDACTED")
		}
	}

	if ballot.Kind == commitchema.PreendorseKind && !ballot.LedgerUUID.EqualsNull() {
		//
		//
		//
		if len(ballot.AdditionNotation) > MaximumSigningExtent {
			return fmt.Errorf("REDACTED", MaximumSigningExtent)
		}

		//
		//
		//
		//
		if len(ballot.AdditionNotation) == 0 && len(ballot.Addition) != 0 {
			return fmt.Errorf("REDACTED")
		}
	}

	return nil
}

//
//
func (ballot *Ballot) AssureAddition() error {
	//
	if ballot.Kind != commitchema.PreendorseKind {
		return nil
	}
	if ballot.LedgerUUID.EqualsNull() {
		return nil
	}
	if len(ballot.AdditionNotation) > 0 {
		return nil
	}
	return FaultBallotAdditionMissing
}

//
//
func (ballot *Ballot) TowardSchema() *commitchema.Ballot {
	if ballot == nil {
		return nil
	}

	return &commitchema.Ballot{
		Kind:               ballot.Kind,
		Altitude:             ballot.Altitude,
		Iteration:              ballot.Iteration,
		LedgerUUID:            ballot.LedgerUUID.TowardSchema(),
		Timestamp:          ballot.Timestamp,
		AssessorLocation:   ballot.AssessorLocation,
		AssessorOrdinal:     ballot.AssessorOrdinal,
		Notation:          ballot.Notation,
		Addition:          ballot.Addition,
		AdditionNotation: ballot.AdditionNotation,
	}
}

func BallotsTowardSchema(ballots []*Ballot) []*commitchema.Ballot {
	if ballots == nil {
		return nil
	}

	res := make([]*commitchema.Ballot, 0, len(ballots))
	for _, ballot := range ballots {
		v := ballot.TowardSchema()
		//
		if v != nil {
			res = append(res, v)
		}
	}
	return res
}

//
//
//
func AttestAlsoInspectBallot(
	ballot *Ballot,
	privateItem PrivateAssessor,
	successionUUID string,
	additionsActivated bool,
) (bool, error) {
	v := ballot.TowardSchema()
	if err := privateItem.AttestBallot(successionUUID, v); err != nil {
		//
		//
		return true, err
	}
	ballot.Notation = v.Notation

	equalsPreendorse := ballot.Kind == commitchema.PreendorseKind
	if !equalsPreendorse && additionsActivated {
		//
		return false, &FaultBallotAdditionUnfit{AddnSigning: v.AdditionNotation}
	}

	equalsVoid := ballot.LedgerUUID.EqualsNull()
	addnSigning := (len(v.AdditionNotation) > 0)

	//
	if addnSigning && (!equalsPreendorse || equalsVoid) {
		//
		return false, &FaultBallotAdditionUnfit{AddnSigning: v.AdditionNotation}
	}

	ballot.AdditionNotation = nil
	if additionsActivated {
		//
		if !addnSigning && equalsPreendorse && !equalsVoid {
			//
			return false, &FaultBallotAdditionUnfit{AddnSigning: v.AdditionNotation}
		}

		ballot.AdditionNotation = v.AdditionNotation
	}

	ballot.Timestamp = v.Timestamp

	return true, nil
}
