package kinds

import (
	"errors"
	"fmt"
	"time"

	cometbytes "github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/protoio"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

var (
	ErrCorruptLedgerSectionAutograph = errors.New("REDACTED")
	ErrCorruptLedgerSectionDigest      = errors.New("REDACTED")
)

//
//
//
//
//
//
type Nomination struct {
	Kind      engineproto.AttestedMessageKind
	Level    int64     `json:"level"`
	Cycle     int32     `json:"epoch"`     //
	POLDuration  int32     `json:"pol_epoch"` //
	LedgerUID   LedgerUID   `json:"ledger_uid"`
	Timestamp time.Time `json:"timestamp"`
	Autograph []byte    `json:"autograph"`
}

//
//
func NewNomination(level int64, epoch int32, polEpoch int32, ledgerUID LedgerUID) *Nomination {
	return &Nomination{
		Kind:      engineproto.NominationKind,
		Level:    level,
		Cycle:     epoch,
		LedgerUID:   ledgerUID,
		POLDuration:  polEpoch,
		Timestamp: engineclock.Now(),
	}
}

//
func (p *Nomination) CertifySimple() error {
	if p.Kind != engineproto.NominationKind {
		return errors.New("REDACTED")
	}
	if p.Level < 0 {
		return errors.New("REDACTED")
	}
	if p.Cycle < 0 {
		return errors.New("REDACTED")
	}
	if p.POLDuration < -1 {
		return errors.New("REDACTED")
	}
	if err := p.LedgerUID.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	//
	if !p.LedgerUID.IsFinished() {
		return fmt.Errorf("REDACTED", p.LedgerUID)
	}

	//

	if len(p.Autograph) == 0 {
		return errors.New("REDACTED")
	}

	if len(p.Autograph) > MaximumAutographVolume {
		return fmt.Errorf("REDACTED", MaximumAutographVolume)
	}
	return nil
}

//
//
//
//
func (p *Nomination) CertifyLedgerVolume(maximumLedgerVolumeOctets int64) error {
	if maximumLedgerVolumeOctets == -1 {
		maximumLedgerVolumeOctets = int64(MaximumLedgerVolumeOctets)
	}
	sumSections := int64(p.LedgerUID.SegmentAssignHeading.Sum)
	maximumSections := (maximumLedgerVolumeOctets-1)/int64(LedgerSegmentVolumeOctets) + 1
	if sumSections > maximumSections {
		return fmt.Errorf("REDACTED", sumSections, maximumSections)
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
//
//
//
func (p *Nomination) String() string {
	return fmt.Sprintf("REDACTED",
		p.Level,
		p.Cycle,
		p.LedgerUID,
		p.POLDuration,
		cometbytes.Footprint(p.Autograph),
		StandardTime(p.Timestamp))
}

//
//
//
//
//
//
//
//
func NominationAttestOctets(ledgerUID string, p *engineproto.Nomination) []byte {
	pb := StandardizeNomination(ledgerUID, p)
	bz, err := protoio.SerializeSeparated(&pb)
	if err != nil {
		panic(err)
	}

	return bz
}

//
func (p *Nomination) ToSchema() *engineproto.Nomination {
	if p == nil {
		return &engineproto.Nomination{}
	}
	pb := new(engineproto.Nomination)

	pb.LedgerUID = p.LedgerUID.ToSchema()
	pb.Kind = p.Kind
	pb.Level = p.Level
	pb.Cycle = p.Cycle
	pb.PolEpoch = p.POLDuration
	pb.Timestamp = p.Timestamp
	pb.Autograph = p.Autograph

	return pb
}

//
//
func NominationFromSchema(pp *engineproto.Nomination) (*Nomination, error) {
	if pp == nil {
		return nil, errors.New("REDACTED")
	}

	p := new(Nomination)

	ledgerUID, err := LedgerUIDFromSchema(&pp.LedgerUID)
	if err != nil {
		return nil, err
	}

	p.LedgerUID = *ledgerUID
	p.Kind = pp.Kind
	p.Level = pp.Level
	p.Cycle = pp.Cycle
	p.POLDuration = pp.PolEpoch
	p.Timestamp = pp.Timestamp
	p.Autograph = pp.Autograph

	return p, p.CertifySimple()
}
