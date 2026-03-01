package kinds

import (
	"errors"
	"fmt"
	"time"

	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

var (
	FaultUnfitLedgerFragmentSigning = errors.New("REDACTED")
	FaultUnfitLedgerFragmentDigest      = errors.New("REDACTED")
)

//
//
//
//
//
//
type Nomination struct {
	Kind      commitchema.AttestedSignalKind
	Altitude    int64     `json:"altitude"`
	Iteration     int32     `json:"iteration"`     //
	PolicyIteration  int32     `json:"policy_iteration"` //
	LedgerUUID   LedgerUUID   `json:"ledger_uuid"`
	Timestamp time.Time `json:"timestamp"`
	Notation []byte    `json:"signing"`
}

//
//
func FreshNomination(altitude int64, iteration int32, policyIteration int32, ledgerUUID LedgerUUID) *Nomination {
	return &Nomination{
		Kind:      commitchema.NominationKind,
		Altitude:    altitude,
		Iteration:     iteration,
		LedgerUUID:   ledgerUUID,
		PolicyIteration:  policyIteration,
		Timestamp: committime.Now(),
	}
}

//
func (p *Nomination) CertifyFundamental() error {
	if p.Kind != commitchema.NominationKind {
		return errors.New("REDACTED")
	}
	if p.Altitude < 0 {
		return errors.New("REDACTED")
	}
	if p.Iteration < 0 {
		return errors.New("REDACTED")
	}
	if p.PolicyIteration < -1 {
		return errors.New("REDACTED")
	}
	if err := p.LedgerUUID.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	//
	if !p.LedgerUUID.EqualsFinish() {
		return fmt.Errorf("REDACTED", p.LedgerUUID)
	}

	//

	if len(p.Notation) == 0 {
		return errors.New("REDACTED")
	}

	if len(p.Notation) > MaximumSigningExtent {
		return fmt.Errorf("REDACTED", MaximumSigningExtent)
	}
	return nil
}

//
//
//
//
func (p *Nomination) CertifyLedgerExtent(maximumLedgerExtentOctets int64) error {
	if maximumLedgerExtentOctets == -1 {
		maximumLedgerExtentOctets = int64(MaximumLedgerExtentOctets)
	}
	sumFragments := int64(p.LedgerUUID.FragmentAssignHeading.Sum)
	maximumFragments := (maximumLedgerExtentOctets-1)/int64(LedgerFragmentExtentOctets) + 1
	if sumFragments > maximumFragments {
		return fmt.Errorf("REDACTED", sumFragments, maximumFragments)
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
func (p *Nomination) Text() string {
	return fmt.Sprintf("REDACTED",
		p.Altitude,
		p.Iteration,
		p.LedgerUUID,
		p.PolicyIteration,
		tendermintoctets.Identifier(p.Notation),
		StandardMoment(p.Timestamp))
}

//
//
//
//
//
//
//
//
func NominationAttestOctets(successionUUID string, p *commitchema.Nomination) []byte {
	pb := NormalizeNomination(successionUUID, p)
	bz, err := protocolio.SerializeSeparated(&pb)
	if err != nil {
		panic(err)
	}

	return bz
}

//
func (p *Nomination) TowardSchema() *commitchema.Nomination {
	if p == nil {
		return &commitchema.Nomination{}
	}
	pb := new(commitchema.Nomination)

	pb.LedgerUUID = p.LedgerUUID.TowardSchema()
	pb.Kind = p.Kind
	pb.Altitude = p.Altitude
	pb.Iteration = p.Iteration
	pb.PolicyIteration = p.PolicyIteration
	pb.Timestamp = p.Timestamp
	pb.Notation = p.Notation

	return pb
}

//
//
func NominationOriginatingSchema(pp *commitchema.Nomination) (*Nomination, error) {
	if pp == nil {
		return nil, errors.New("REDACTED")
	}

	p := new(Nomination)

	ledgerUUID, err := LedgerUUIDOriginatingSchema(&pp.LedgerUUID)
	if err != nil {
		return nil, err
	}

	p.LedgerUUID = *ledgerUUID
	p.Kind = pp.Kind
	p.Altitude = pp.Altitude
	p.Iteration = pp.Iteration
	p.PolicyIteration = pp.PolicyIteration
	p.Timestamp = pp.Timestamp
	p.Notation = pp.Notation

	return p, p.CertifyFundamental()
}
