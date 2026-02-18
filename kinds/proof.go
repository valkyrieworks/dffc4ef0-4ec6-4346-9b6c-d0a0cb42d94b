package kinds

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/vault/comethash"
	cometjson "github.com/valkyrieworks/utils/json"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometfaults "github.com/valkyrieworks/kinds/faults"
)

//
//
type Proof interface {
	Iface() []iface.Malpractice //
	Octets() []byte            //
	Digest() []byte             //
	Level() int64            //
	String() string           //
	Time() time.Time          //
	CertifySimple() error     //
}

//

//
type ReplicatedBallotProof struct {
	BallotA *Ballot `json:"ballot_a"`
	BallotBYTE *Ballot `json:"ballot_byte"`

	//
	SumPollingEnergy int64
	RatifierEnergy   int64
	Timestamp        time.Time
}

var _ Proof = &ReplicatedBallotProof{}

//
//
//
func NewReplicatedBallotProof(vote1, ballot2 *Ballot, ledgerTime time.Time, valueCollection *RatifierAssign,
) (*ReplicatedBallotProof, error) {
	var ballotA, ballotBYTE *Ballot
	if vote1 == nil || ballot2 == nil {
		return nil, errors.New("REDACTED")
	}
	if valueCollection == nil {
		return nil, errors.New("REDACTED")
	}
	idx, val := valueCollection.FetchByLocation(vote1.RatifierLocation)
	if idx == -1 {
		return nil, fmt.Errorf("REDACTED", vote1.RatifierLocation.String())
	}

	if strings.Compare(vote1.LedgerUID.Key(), ballot2.LedgerUID.Key()) == -1 {
		ballotA = vote1
		ballotBYTE = ballot2
	} else {
		ballotA = ballot2
		ballotBYTE = vote1
	}
	return &ReplicatedBallotProof{
		BallotA:            ballotA,
		BallotBYTE:            ballotBYTE,
		SumPollingEnergy: valueCollection.SumPollingEnergy(),
		RatifierEnergy:   val.PollingEnergy,
		Timestamp:        ledgerTime,
	}, nil
}

//
func (dve *ReplicatedBallotProof) Iface() []iface.Malpractice {
	return []iface.Malpractice{{
		Kind: iface.Misconductkind_REPLICATED_BALLOT,
		Ratifier: iface.Ratifier{
			Location: dve.BallotA.RatifierLocation,
			Energy:   dve.RatifierEnergy,
		},
		Level:           dve.BallotA.Level,
		Time:             dve.Timestamp,
		SumPollingEnergy: dve.SumPollingEnergy,
	}}
}

//
func (dve *ReplicatedBallotProof) Octets() []byte {
	pbe := dve.ToSchema()
	bz, err := pbe.Serialize()
	if err != nil {
		panic(err)
	}

	return bz
}

//
func (dve *ReplicatedBallotProof) Digest() []byte {
	return comethash.Sum(dve.Octets())
}

//
func (dve *ReplicatedBallotProof) Level() int64 {
	return dve.BallotA.Level
}

//
func (dve *ReplicatedBallotProof) String() string {
	return fmt.Sprintf("REDACTED", dve.BallotA, dve.BallotBYTE)
}

//
func (dve *ReplicatedBallotProof) Time() time.Time {
	return dve.Timestamp
}

//
func (dve *ReplicatedBallotProof) CertifySimple() error {
	if dve == nil {
		return cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}

	if dve.BallotA == nil || dve.BallotBYTE == nil {
		return fmt.Errorf("REDACTED", dve.BallotA, dve.BallotBYTE)
	}
	if err := dve.BallotA.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := dve.BallotBYTE.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	//
	if strings.Compare(dve.BallotA.LedgerUID.Key(), dve.BallotBYTE.LedgerUID.Key()) >= 0 {
		return errors.New("REDACTED")
	}
	return nil
}

//
func (dve *ReplicatedBallotProof) ToSchema() *engineproto.ReplicatedBallotProof {
	ballotBYTE := dve.BallotBYTE.ToSchema()
	ballotA := dve.BallotA.ToSchema()
	tp := engineproto.ReplicatedBallotProof{
		BallotA:            ballotA,
		BallotBYTE:            ballotBYTE,
		SumPollingEnergy: dve.SumPollingEnergy,
		RatifierEnergy:   dve.RatifierEnergy,
		Timestamp:        dve.Timestamp,
	}
	return &tp
}

//
func ReplicatedBallotProofFromSchema(pb *engineproto.ReplicatedBallotProof) (*ReplicatedBallotProof, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	var vA *Ballot
	if pb.BallotA != nil {
		var err error
		vA, err = BallotFromSchema(pb.BallotA)
		if err != nil {
			return nil, err
		}
		if err = vA.CertifySimple(); err != nil {
			return nil, err
		}
	}

	var vB *Ballot
	if pb.BallotBYTE != nil {
		var err error
		vB, err = BallotFromSchema(pb.BallotBYTE)
		if err != nil {
			return nil, err
		}
		if err = vB.CertifySimple(); err != nil {
			return nil, err
		}
	}

	dve := &ReplicatedBallotProof{
		BallotA:            vA,
		BallotBYTE:            vB,
		SumPollingEnergy: pb.SumPollingEnergy,
		RatifierEnergy:   pb.RatifierEnergy,
		Timestamp:        pb.Timestamp,
	}

	return dve, dve.CertifySimple()
}

//

//
//
//
//
//
type RapidCustomerAssaultProof struct {
	ClashingLedger *RapidLedger
	SharedLevel     int64

	//
	FaultyRatifiers []*Ratifier //
	SumPollingEnergy    int64        //
	Timestamp           time.Time    //
}

var _ Proof = &RapidCustomerAssaultProof{}

//
func (l *RapidCustomerAssaultProof) Iface() []iface.Malpractice {
	ifaceEvt := make([]iface.Malpractice, len(l.FaultyRatifiers))
	for idx, val := range l.FaultyRatifiers {
		ifaceEvt[idx] = iface.Malpractice{
			Kind:             iface.Misconductkind_RAPID_CUSTOMER_ASSAULT,
			Ratifier:        Tm2schema.Ratifier(val),
			Level:           l.Level(),
			Time:             l.Timestamp,
			SumPollingEnergy: l.SumPollingEnergy,
		}
	}
	return ifaceEvt
}

//
func (l *RapidCustomerAssaultProof) Octets() []byte {
	pbe, err := l.ToSchema()
	if err != nil {
		panic(err)
	}
	bz, err := pbe.Serialize()
	if err != nil {
		panic(err)
	}
	return bz
}

//
//
//
func (l *RapidCustomerAssaultProof) FetchFaultyRatifiers(sharedValues *RatifierAssign,
	validated *AttestedHeading,
) []*Ratifier {
	var ratifiers []*Ratifier
	//
	//
	if l.ClashingHeadingIsCorrupt(validated.Heading) {
		for _, endorseSignature := range l.ClashingLedger.Endorse.Endorsements {
			if endorseSignature.LedgerUIDMark != LedgerUIDMarkEndorse {
				continue
			}

			_, val := sharedValues.FetchByLocation(endorseSignature.RatifierLocation)
			if val == nil {
				//
				continue
			}
			ratifiers = append(ratifiers, val)
		}
		sort.Sort(RatifiersByPollingEnergy(ratifiers))
		return ratifiers
	} else if validated.Endorse.Cycle == l.ClashingLedger.Endorse.Cycle {
		//
		//
		//
		//
		for i := 0; i < len(l.ClashingLedger.Endorse.Endorsements); i++ {
			signatureA := l.ClashingLedger.Endorse.Endorsements[i]
			if signatureA.LedgerUIDMark != LedgerUIDMarkEndorse {
				continue
			}

			signatureBYTE := validated.Endorse.Endorsements[i]
			if signatureBYTE.LedgerUIDMark != LedgerUIDMarkEndorse {
				continue
			}

			_, val := l.ClashingLedger.RatifierAssign.FetchByLocation(signatureA.RatifierLocation)
			ratifiers = append(ratifiers, val)
		}
		sort.Sort(RatifiersByPollingEnergy(ratifiers))
		return ratifiers
	}
	//
	//
	//
	return ratifiers
}

//
//
//
//
func (l *RapidCustomerAssaultProof) ClashingHeadingIsCorrupt(validatedHeading *Heading) bool {
	return !bytes.Equal(validatedHeading.RatifiersDigest, l.ClashingLedger.RatifiersDigest) ||
		!bytes.Equal(validatedHeading.FollowingRatifiersDigest, l.ClashingLedger.FollowingRatifiersDigest) ||
		!bytes.Equal(validatedHeading.AgreementDigest, l.ClashingLedger.AgreementDigest) ||
		!bytes.Equal(validatedHeading.ApplicationDigest, l.ClashingLedger.ApplicationDigest) ||
		!bytes.Equal(validatedHeading.FinalOutcomesDigest, l.ClashingLedger.FinalOutcomesDigest)
}

//
//
//
//
//
//
//
//
func (l *RapidCustomerAssaultProof) Digest() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, l.SharedLevel)
	bz := make([]byte, comethash.Volume+n)
	copy(bz[:comethash.Volume-1], l.ClashingLedger.Digest().Octets())
	copy(bz[comethash.Volume:], buf)
	return comethash.Sum(bz)
}

//
//
//
func (l *RapidCustomerAssaultProof) Level() int64 {
	return l.SharedLevel
}

//
func (l *RapidCustomerAssaultProof) String() string {
	return fmt.Sprintf(`REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED`,
		l.ClashingLedger.String(), l.SharedLevel, l.FaultyRatifiers,
		l.SumPollingEnergy, l.Timestamp, l.Digest())
}

//
func (l *RapidCustomerAssaultProof) Time() time.Time {
	return l.Timestamp
}

//
func (l *RapidCustomerAssaultProof) CertifySimple() error {
	if l.ClashingLedger == nil {
		return errors.New("REDACTED")
	}

	//
	if l.ClashingLedger.Heading == nil {
		return errors.New("REDACTED")
	}

	if l.SumPollingEnergy <= 0 {
		return errors.New("REDACTED")
	}

	if l.SharedLevel <= 0 {
		return errors.New("REDACTED")
	}

	//
	//
	//
	if l.SharedLevel > l.ClashingLedger.Level {
		return fmt.Errorf("REDACTED",
			l.SharedLevel, l.ClashingLedger.Level)
	}

	if err := l.ClashingLedger.CertifySimple(l.ClashingLedger.LedgerUID); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return nil
}

//
func (l *RapidCustomerAssaultProof) ToSchema() (*engineproto.RapidCustomerAssaultProof, error) {
	clashingLedger, err := l.ClashingLedger.ToSchema()
	if err != nil {
		return nil, err
	}

	byzValues := make([]*engineproto.Ratifier, len(l.FaultyRatifiers))
	for idx, val := range l.FaultyRatifiers {
		valueschema, err := val.ToSchema()
		if err != nil {
			return nil, err
		}
		byzValues[idx] = valueschema
	}

	return &engineproto.RapidCustomerAssaultProof{
		ClashingLedger:    clashingLedger,
		SharedLevel:        l.SharedLevel,
		FaultyRatifiers: byzValues,
		SumPollingEnergy:    l.SumPollingEnergy,
		Timestamp:           l.Timestamp,
	}, nil
}

//
func RapidCustomerAssaultProofFromSchema(lpb *engineproto.RapidCustomerAssaultProof) (*RapidCustomerAssaultProof, error) {
	if lpb == nil {
		return nil, cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}

	clashingLedger, err := RapidLedgerFromSchema(lpb.ClashingLedger)
	if err != nil {
		return nil, err
	}

	byzValues := make([]*Ratifier, len(lpb.FaultyRatifiers))
	for idx, valueschema := range lpb.FaultyRatifiers {
		val, err := RatifierFromSchema(valueschema)
		if err != nil {
			return nil, err
		}
		byzValues[idx] = val
	}

	l := &RapidCustomerAssaultProof{
		ClashingLedger:    clashingLedger,
		SharedLevel:        lpb.SharedLevel,
		FaultyRatifiers: byzValues,
		SumPollingEnergy:    lpb.SumPollingEnergy,
		Timestamp:           lpb.Timestamp,
	}

	return l, l.CertifySimple()
}

//

//
type ProofCatalog []Proof

//
func (evl ProofCatalog) Digest() []byte {
	//
	//
	//
	proofBzs := make([][]byte, len(evl))
	for i := 0; i < len(evl); i++ {
		//
		//
		proofBzs[i] = evl[i].Octets()
	}
	return merkle.DigestFromOctetSegments(proofBzs)
}

func (evl ProofCatalog) String() string {
	s := "REDACTED"
	for _, e := range evl {
		s += fmt.Sprintf("REDACTED", e)
	}
	return s
}

//
func (evl ProofCatalog) Has(proof Proof) bool {
	for _, ev := range evl {
		if bytes.Equal(proof.Digest(), ev.Digest()) {
			return true
		}
	}
	return false
}

//
//
func (evl ProofCatalog) ToIface() []iface.Malpractice {
	var el []iface.Malpractice
	for _, e := range evl {
		el = append(el, e.Iface()...)
	}
	return el
}

//

//
//
func ProofToSchema(proof Proof) (*engineproto.Proof, error) {
	if proof == nil {
		return nil, errors.New("REDACTED")
	}

	switch evi := proof.(type) {
	case *ReplicatedBallotProof:
		schemav := evi.ToSchema()
		return &engineproto.Proof{
			Sum: &engineproto.Proof_Duplicateballotevidence{
				ReplicatedBallotProof: schemav,
			},
		}, nil

	case *RapidCustomerAssaultProof:
		schemav, err := evi.ToSchema()
		if err != nil {
			return nil, err
		}
		return &engineproto.Proof{
			Sum: &engineproto.Proof_Rapidcustomerevidence{
				RapidCustomerAssaultProof: schemav,
			},
		}, nil

	default:
		return nil, fmt.Errorf("REDACTED", evi)
	}
}

//
//
func ProofFromSchema(proof *engineproto.Proof) (Proof, error) {
	if proof == nil {
		return nil, errors.New("REDACTED")
	}

	switch evi := proof.Sum.(type) {
	case *engineproto.Proof_Duplicateballotevidence:
		return ReplicatedBallotProofFromSchema(evi.ReplicatedBallotProof)
	case *engineproto.Proof_Rapidcustomerevidence:
		return RapidCustomerAssaultProofFromSchema(evi.RapidCustomerAssaultProof)
	default:
		return nil, errors.New("REDACTED")
	}
}

func init() {
	cometjson.EnrollKind(&ReplicatedBallotProof{}, "REDACTED")
	cometjson.EnrollKind(&RapidCustomerAssaultProof{}, "REDACTED")
}

//

//
type ErrCorruptProof struct {
	Proof Proof
	Cause   error
}

//
func NewErrCorruptProof(ev Proof, err error) *ErrCorruptProof {
	return &ErrCorruptProof{ev, err}
}

//
func (err *ErrCorruptProof) Fault() string {
	return fmt.Sprintf("REDACTED", err.Cause, err.Proof)
}

//
type ErrProofOverload struct {
	Max int64
	Got int64
}

//
func NewErrProofOverload(max, got int64) *ErrProofOverload {
	return &ErrProofOverload{max, got}
}

//
func (err *ErrProofOverload) Fault() string {
	return fmt.Sprintf("REDACTED", err.Max, err.Got)
}

//

//

//
//
func NewEmulateReplicatedBallotProof(level int64, moment time.Time, ledgerUID string) (*ReplicatedBallotProof, error) {
	val := NewEmulatePV()
	return NewEmulateReplicatedBallotProofWithRatifier(level, moment, val, ledgerUID)
}

//
//
func NewEmulateReplicatedBallotProofWithRatifier(level int64, moment time.Time,
	pv PrivateRatifier, ledgerUID string,
) (*ReplicatedBallotProof, error) {
	publicKey, err := pv.FetchPublicKey()
	if err != nil {
		return nil, err
	}
	val := NewRatifier(publicKey, 10)
	ballotA := createEmulateBallot(level, 0, 0, publicKey.Location(), randomLedgerUID(), moment)
	vA := ballotA.ToSchema()
	err = pv.AttestBallot(ledgerUID, vA)
	if err != nil {
		return nil, err
	}
	ballotA.Autograph = vA.Autograph
	ballotBYTE := createEmulateBallot(level, 0, 0, publicKey.Location(), randomLedgerUID(), moment)
	vB := ballotBYTE.ToSchema()
	err = pv.AttestBallot(ledgerUID, vB)
	if err != nil {
		return nil, err
	}
	ballotBYTE.Autograph = vB.Autograph
	return NewReplicatedBallotProof(ballotA, ballotBYTE, moment, NewRatifierCollection([]*Ratifier{val}))
}

func createEmulateBallot(level int64, epoch, ordinal int32, address Location,
	ledgerUID LedgerUID, moment time.Time,
) *Ballot {
	return &Ballot{
		Kind:             engineproto.AttestedMessageKind(2),
		Level:           level,
		Cycle:            epoch,
		LedgerUID:          ledgerUID,
		Timestamp:        moment,
		RatifierLocation: address,
		RatifierOrdinal:   ordinal,
	}
}

func randomLedgerUID() LedgerUID {
	return LedgerUID{
		Digest: engineseed.Octets(comethash.Volume),
		SegmentAssignHeading: SegmentAssignHeading{
			Sum: 1,
			Digest:  engineseed.Octets(comethash.Volume),
		},
	}
}
