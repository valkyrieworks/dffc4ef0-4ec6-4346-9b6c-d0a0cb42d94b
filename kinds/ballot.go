package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/vault"
	cometbytes "github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/protoio"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

const (
	nullBallotStr string = "REDACTED"

	//
	MaximumBallotAdditionVolume int = 1024 * 1024
)

var (
	ErrBallotUnforeseenPhase            = errors.New("REDACTED")
	ErrBallotCorruptRatifierOrdinal     = errors.New("REDACTED")
	ErrBallotCorruptRatifierLocation   = errors.New("REDACTED")
	ErrBallotCorruptAutograph          = errors.New("REDACTED")
	ErrBallotCorruptLedgerDigest          = errors.New("REDACTED")
	ErrBallotNotCertainAutograph = errors.New("REDACTED")
	ErrBallotNull                       = errors.New("REDACTED")
	ErrBallotAdditionMissing           = errors.New("REDACTED")
	ErrCorruptBallotAddition          = errors.New("REDACTED")
)

type ErrBallotClashingBallots struct {
	BallotA *Ballot
	BallotBYTE *Ballot
}

func (err *ErrBallotClashingBallots) Fault() string {
	return fmt.Sprintf("REDACTED", err.BallotA.RatifierLocation)
}

func NewClashingBallotFault(vote1, ballot2 *Ballot) *ErrBallotClashingBallots {
	return &ErrBallotClashingBallots{
		BallotA: vote1,
		BallotBYTE: ballot2,
	}
}

//
type ErrBallotAdditionCorrupt struct {
	ExtensionAutograph []byte
}

func (err *ErrBallotAdditionCorrupt) Fault() string {
	return fmt.Sprintf("REDACTED", err.ExtensionAutograph)
}

//
type Location = vault.Location

//
//
type Ballot struct {
	Kind               engineproto.AttestedMessageKind `json:"kind"`
	Level             int64                  `json:"level"`
	Cycle              int32                  `json:"epoch"`    //
	LedgerUID            LedgerUID                `json:"ledger_uid"` //
	Timestamp          time.Time              `json:"timestamp"`
	RatifierLocation   Location                `json:"ratifier_location"`
	RatifierOrdinal     int32                  `json:"ratifier_ordinal"`
	Autograph          []byte                 `json:"autograph"`
	Addition          []byte                 `json:"addition"`
	AdditionAutograph []byte                 `json:"addition_autograph"`
}

//
//
//
//
func BallotFromSchema(pv *engineproto.Ballot) (*Ballot, error) {
	ledgerUID, err := LedgerUIDFromSchema(&pv.LedgerUID)
	if err != nil {
		return nil, err
	}

	return &Ballot{
		Kind:               pv.Kind,
		Level:             pv.Level,
		Cycle:              pv.Cycle,
		LedgerUID:            *ledgerUID,
		Timestamp:          pv.Timestamp,
		RatifierLocation:   pv.RatifierLocation,
		RatifierOrdinal:     pv.RatifierOrdinal,
		Autograph:          pv.Autograph,
		Addition:          pv.Addition,
		AdditionAutograph: pv.AdditionAutograph,
	}, nil
}

//
func (ballot *Ballot) EndorseSignature() EndorseSignature {
	if ballot == nil {
		return NewEndorseSignatureMissing()
	}

	var ledgerUIDMark LedgerUIDMark
	switch {
	case ballot.LedgerUID.IsFinished():
		ledgerUIDMark = LedgerUIDMarkEndorse
	case ballot.LedgerUID.IsNil():
		ledgerUIDMark = LedgerUIDMarkNull
	default:
		panic(fmt.Sprintf("REDACTED", ballot))
	}

	return EndorseSignature{
		LedgerUIDMark:      ledgerUIDMark,
		RatifierLocation: ballot.RatifierLocation,
		Timestamp:        ballot.Timestamp,
		Autograph:        ballot.Autograph,
	}
}

//
//
//
func (ballot *Ballot) ExpandedEndorseSignature() ExpandedEndorseSignature {
	if ballot == nil {
		return NewExpandedEndorseSignatureMissing()
	}

	return ExpandedEndorseSignature{
		EndorseSignature:          ballot.EndorseSignature(),
		Addition:          ballot.Addition,
		AdditionAutograph: ballot.AdditionAutograph,
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
func BallotAttestOctets(ledgerUID string, ballot *engineproto.Ballot) []byte {
	pb := StandardizeBallot(ledgerUID, ballot)
	bz, err := protoio.SerializeSeparated(&pb)
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
func BallotAdditionAttestOctets(ledgerUID string, ballot *engineproto.Ballot) []byte {
	pb := StandardizeBallotAddition(ledgerUID, ballot)
	bz, err := protoio.SerializeSeparated(&pb)
	if err != nil {
		panic(err)
	}

	return bz
}

func (ballot *Ballot) Clone() *Ballot {
	ballotClone := *ballot
	return &ballotClone
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
func (ballot *Ballot) String() string {
	if ballot == nil {
		return nullBallotStr
	}

	var kindString string
	switch ballot.Kind {
	case engineproto.PreballotKind:
		kindString = "REDACTED"
	case engineproto.PreendorseKind:
		kindString = "REDACTED"
	default:
		panic("REDACTED")
	}

	return fmt.Sprintf("REDACTED",
		ballot.RatifierOrdinal,
		cometbytes.Footprint(ballot.RatifierLocation),
		ballot.Level,
		ballot.Cycle,
		ballot.Kind,
		kindString,
		cometbytes.Footprint(ballot.LedgerUID.Digest),
		cometbytes.Footprint(ballot.Autograph),
		cometbytes.Footprint(ballot.Addition),
		StandardTime(ballot.Timestamp),
	)
}

func (ballot *Ballot) validateAndYieldSchema(ledgerUID string, publicKey vault.PublicKey) (*engineproto.Ballot, error) {
	if !bytes.Equal(publicKey.Location(), ballot.RatifierLocation) {
		return nil, ErrBallotCorruptRatifierLocation
	}
	v := ballot.ToSchema()
	if !publicKey.ValidateAutograph(BallotAttestOctets(ledgerUID, v), ballot.Autograph) {
		return nil, ErrBallotCorruptAutograph
	}
	return v, nil
}

//
//
//
func (ballot *Ballot) Validate(ledgerUID string, publicKey vault.PublicKey) error {
	_, err := ballot.validateAndYieldSchema(ledgerUID, publicKey)
	return err
}

//
//
//
//
func (ballot *Ballot) ValidateBallotAndAddition(ledgerUID string, publicKey vault.PublicKey) error {
	v, err := ballot.validateAndYieldSchema(ledgerUID, publicKey)
	if err != nil {
		return err
	}
	//
	if ballot.Kind == engineproto.PreendorseKind && !SchemaLedgerUIDIsNull(&v.LedgerUID) {
		if len(ballot.AdditionAutograph) == 0 {
			return errors.New("REDACTED")
		}

		extensionAttestOctets := BallotAdditionAttestOctets(ledgerUID, v)
		if !publicKey.ValidateAutograph(extensionAttestOctets, ballot.AdditionAutograph) {
			return ErrBallotCorruptAutograph
		}
	}
	return nil
}

//
//
func (ballot *Ballot) ValidateAddition(ledgerUID string, publicKey vault.PublicKey) error {
	if ballot.Kind != engineproto.PreendorseKind || ballot.LedgerUID.IsNil() {
		return nil
	}
	v := ballot.ToSchema()
	extensionAttestOctets := BallotAdditionAttestOctets(ledgerUID, v)
	if !publicKey.ValidateAutograph(extensionAttestOctets, ballot.AdditionAutograph) {
		return ErrBallotCorruptAutograph
	}
	return nil
}

//
//
//
func (ballot *Ballot) CertifySimple() error {
	if !IsBallotKindSound(ballot.Kind) {
		return errors.New("REDACTED")
	}

	if ballot.Level <= 0 {
		return errors.New("REDACTED")
	}

	if ballot.Cycle < 0 {
		return errors.New("REDACTED")
	}

	//

	if err := ballot.LedgerUID.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	//
	if !ballot.LedgerUID.IsNil() && !ballot.LedgerUID.IsFinished() {
		return fmt.Errorf("REDACTED", ballot.LedgerUID)
	}

	if len(ballot.RatifierLocation) != vault.LocationVolume {
		return fmt.Errorf("REDACTED",
			vault.LocationVolume,
			len(ballot.RatifierLocation),
		)
	}
	if ballot.RatifierOrdinal < 0 {
		return errors.New("REDACTED")
	}
	if len(ballot.Autograph) == 0 {
		return errors.New("REDACTED")
	}

	if len(ballot.Autograph) > MaximumAutographVolume {
		return fmt.Errorf("REDACTED", MaximumAutographVolume)
	}

	//
	//
	//
	if ballot.Kind != engineproto.PreendorseKind || ballot.LedgerUID.IsNil() {
		if len(ballot.Addition) > 0 {
			return fmt.Errorf(
				"REDACTED",
				ballot.Kind, ballot.LedgerUID.IsNil(),
			)
		}
		if len(ballot.AdditionAutograph) > 0 {
			return errors.New("REDACTED")
		}
	}

	if ballot.Kind == engineproto.PreendorseKind && !ballot.LedgerUID.IsNil() {
		//
		//
		//
		if len(ballot.AdditionAutograph) > MaximumAutographVolume {
			return fmt.Errorf("REDACTED", MaximumAutographVolume)
		}

		//
		//
		//
		//
		if len(ballot.AdditionAutograph) == 0 && len(ballot.Addition) != 0 {
			return fmt.Errorf("REDACTED")
		}
	}

	return nil
}

//
//
func (ballot *Ballot) AssureAddition() error {
	//
	if ballot.Kind != engineproto.PreendorseKind {
		return nil
	}
	if ballot.LedgerUID.IsNil() {
		return nil
	}
	if len(ballot.AdditionAutograph) > 0 {
		return nil
	}
	return ErrBallotAdditionMissing
}

//
//
func (ballot *Ballot) ToSchema() *engineproto.Ballot {
	if ballot == nil {
		return nil
	}

	return &engineproto.Ballot{
		Kind:               ballot.Kind,
		Level:             ballot.Level,
		Cycle:              ballot.Cycle,
		LedgerUID:            ballot.LedgerUID.ToSchema(),
		Timestamp:          ballot.Timestamp,
		RatifierLocation:   ballot.RatifierLocation,
		RatifierOrdinal:     ballot.RatifierOrdinal,
		Autograph:          ballot.Autograph,
		Addition:          ballot.Addition,
		AdditionAutograph: ballot.AdditionAutograph,
	}
}

func BallotsToSchema(ballots []*Ballot) []*engineproto.Ballot {
	if ballots == nil {
		return nil
	}

	res := make([]*engineproto.Ballot, 0, len(ballots))
	for _, ballot := range ballots {
		v := ballot.ToSchema()
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
func AttestAndInspectBallot(
	ballot *Ballot,
	privateValue PrivateRatifier,
	ledgerUID string,
	pluginsActivated bool,
) (bool, error) {
	v := ballot.ToSchema()
	if err := privateValue.AttestBallot(ledgerUID, v); err != nil {
		//
		//
		return true, err
	}
	ballot.Autograph = v.Autograph

	isPreendorse := ballot.Kind == engineproto.PreendorseKind
	if !isPreendorse && pluginsActivated {
		//
		return false, &ErrBallotAdditionCorrupt{ExtensionAutograph: v.AdditionAutograph}
	}

	isNull := ballot.LedgerUID.IsNil()
	extensionAutograph := (len(v.AdditionAutograph) > 0)

	//
	if extensionAutograph && (!isPreendorse || isNull) {
		//
		return false, &ErrBallotAdditionCorrupt{ExtensionAutograph: v.AdditionAutograph}
	}

	ballot.AdditionAutograph = nil
	if pluginsActivated {
		//
		if !extensionAutograph && isPreendorse && !isNull {
			//
			return false, &ErrBallotAdditionCorrupt{ExtensionAutograph: v.AdditionAutograph}
		}

		ballot.AdditionAutograph = v.AdditionAutograph
	}

	ballot.Timestamp = v.Timestamp

	return true, nil
}
