package kinds

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
)

//
//
type Proof interface {
	Iface() []iface.Malpractice //
	Octets() []byte            //
	Digest() []byte             //
	Altitude() int64            //
	Text() string           //
	Moment() time.Time          //
	CertifyFundamental() error     //
}

//

//
type ReplicatedBallotProof struct {
	BallotAN *Ballot `json:"ballot_an"`
	BallotBYTE *Ballot `json:"ballot_byte"`

	//
	SumBallotingPotency int64
	AssessorPotency   int64
	Timestamp        time.Time
}

var _ Proof = &ReplicatedBallotProof{}

//
//
//
func FreshReplicatedBallotProof(ballot1, ballot2 *Ballot, ledgerMoment time.Time, itemAssign *AssessorAssign,
) (*ReplicatedBallotProof, error) {
	var ballotAN, ballotBYTE *Ballot
	if ballot1 == nil || ballot2 == nil {
		return nil, errors.New("REDACTED")
	}
	if itemAssign == nil {
		return nil, errors.New("REDACTED")
	}
	idx, val := itemAssign.ObtainViaLocation(ballot1.AssessorLocation)
	if idx == -1 {
		return nil, fmt.Errorf("REDACTED", ballot1.AssessorLocation.Text())
	}

	if strings.Compare(ballot1.LedgerUUID.Key(), ballot2.LedgerUUID.Key()) == -1 {
		ballotAN = ballot1
		ballotBYTE = ballot2
	} else {
		ballotAN = ballot2
		ballotBYTE = ballot1
	}
	return &ReplicatedBallotProof{
		BallotAN:            ballotAN,
		BallotBYTE:            ballotBYTE,
		SumBallotingPotency: itemAssign.SumBallotingPotency(),
		AssessorPotency:   val.BallotingPotency,
		Timestamp:        ledgerMoment,
	}, nil
}

//
func (dve *ReplicatedBallotProof) Iface() []iface.Malpractice {
	return []iface.Malpractice{{
		Kind: iface.Malfunctionkind_REPLICATED_BALLOT,
		Assessor: iface.Assessor{
			Location: dve.BallotAN.AssessorLocation,
			Potency:   dve.AssessorPotency,
		},
		Altitude:           dve.BallotAN.Altitude,
		Moment:             dve.Timestamp,
		SumBallotingPotency: dve.SumBallotingPotency,
	}}
}

//
func (dve *ReplicatedBallotProof) Octets() []byte {
	pbe := dve.TowardSchema()
	bz, err := pbe.Serialize()
	if err != nil {
		panic(err)
	}

	return bz
}

//
func (dve *ReplicatedBallotProof) Digest() []byte {
	return tenderminthash.Sum(dve.Octets())
}

//
func (dve *ReplicatedBallotProof) Altitude() int64 {
	return dve.BallotAN.Altitude
}

//
func (dve *ReplicatedBallotProof) Text() string {
	return fmt.Sprintf("REDACTED", dve.BallotAN, dve.BallotBYTE)
}

//
func (dve *ReplicatedBallotProof) Moment() time.Time {
	return dve.Timestamp
}

//
func (dve *ReplicatedBallotProof) CertifyFundamental() error {
	if dve == nil {
		return strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}

	if dve.BallotAN == nil || dve.BallotBYTE == nil {
		return fmt.Errorf("REDACTED", dve.BallotAN, dve.BallotBYTE)
	}
	if err := dve.BallotAN.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := dve.BallotBYTE.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	//
	if strings.Compare(dve.BallotAN.LedgerUUID.Key(), dve.BallotBYTE.LedgerUUID.Key()) >= 0 {
		return errors.New("REDACTED")
	}
	return nil
}

//
func (dve *ReplicatedBallotProof) TowardSchema() *commitchema.ReplicatedBallotProof {
	ballotBYTE := dve.BallotBYTE.TowardSchema()
	ballotAN := dve.BallotAN.TowardSchema()
	tp := commitchema.ReplicatedBallotProof{
		BallotAN:            ballotAN,
		BallotBYTE:            ballotBYTE,
		SumBallotingPotency: dve.SumBallotingPotency,
		AssessorPotency:   dve.AssessorPotency,
		Timestamp:        dve.Timestamp,
	}
	return &tp
}

//
func ReplicatedBallotProofOriginatingSchema(pb *commitchema.ReplicatedBallotProof) (*ReplicatedBallotProof, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	var vA *Ballot
	if pb.BallotAN != nil {
		var err error
		vA, err = BallotOriginatingSchema(pb.BallotAN)
		if err != nil {
			return nil, err
		}
		if err = vA.CertifyFundamental(); err != nil {
			return nil, err
		}
	}

	var vB *Ballot
	if pb.BallotBYTE != nil {
		var err error
		vB, err = BallotOriginatingSchema(pb.BallotBYTE)
		if err != nil {
			return nil, err
		}
		if err = vB.CertifyFundamental(); err != nil {
			return nil, err
		}
	}

	dve := &ReplicatedBallotProof{
		BallotAN:            vA,
		BallotBYTE:            vB,
		SumBallotingPotency: pb.SumBallotingPotency,
		AssessorPotency:   pb.AssessorPotency,
		Timestamp:        pb.Timestamp,
	}

	return dve, dve.CertifyFundamental()
}

//

//
//
//
//
//
type AgileCustomerOnslaughtProof struct {
	DiscordantLedger *AgileLedger
	SharedAltitude     int64

	//
	TreacherousAssessors []*Assessor //
	SumBallotingPotency    int64        //
	Timestamp           time.Time    //
}

var _ Proof = &AgileCustomerOnslaughtProof{}

//
func (l *AgileCustomerOnslaughtProof) Iface() []iface.Malpractice {
	ifaceOccurence := make([]iface.Malpractice, len(l.TreacherousAssessors))
	for idx, val := range l.TreacherousAssessors {
		ifaceOccurence[idx] = iface.Malpractice{
			Kind:             iface.Malfunctionkind_AGILE_CUSTOMER_ONSLAUGHT,
			Assessor:        Temp2buffer.Assessor(val),
			Altitude:           l.Altitude(),
			Moment:             l.Timestamp,
			SumBallotingPotency: l.SumBallotingPotency,
		}
	}
	return ifaceOccurence
}

//
func (l *AgileCustomerOnslaughtProof) Octets() []byte {
	pbe, err := l.TowardSchema()
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
func (l *AgileCustomerOnslaughtProof) ObtainTreacherousAssessors(sharedValues *AssessorAssign,
	reliable *NotatedHeading,
) []*Assessor {
	var assessors []*Assessor
	//
	//
	if l.DiscordantHeadingEqualsUnfit(reliable.Heading) {
		for _, endorseSignature := range l.DiscordantLedger.Endorse.Notations {
			if endorseSignature.LedgerUUIDMarker != LedgerUUIDMarkerEndorse {
				continue
			}

			_, val := sharedValues.ObtainViaLocation(endorseSignature.AssessorLocation)
			if val == nil {
				//
				continue
			}
			assessors = append(assessors, val)
		}
		sort.Sort(AssessorsViaBallotingPotency(assessors))
		return assessors
	} else if reliable.Endorse.Iteration == l.DiscordantLedger.Endorse.Iteration {
		//
		//
		//
		//
		for i := 0; i < len(l.DiscordantLedger.Endorse.Notations); i++ {
			signatureAN := l.DiscordantLedger.Endorse.Notations[i]
			if signatureAN.LedgerUUIDMarker != LedgerUUIDMarkerEndorse {
				continue
			}

			signatureBYTE := reliable.Endorse.Notations[i]
			if signatureBYTE.LedgerUUIDMarker != LedgerUUIDMarkerEndorse {
				continue
			}

			_, val := l.DiscordantLedger.AssessorAssign.ObtainViaLocation(signatureAN.AssessorLocation)
			assessors = append(assessors, val)
		}
		sort.Sort(AssessorsViaBallotingPotency(assessors))
		return assessors
	}
	//
	//
	//
	return assessors
}

//
//
//
//
func (l *AgileCustomerOnslaughtProof) DiscordantHeadingEqualsUnfit(reliableHeading *Heading) bool {
	return !bytes.Equal(reliableHeading.AssessorsDigest, l.DiscordantLedger.AssessorsDigest) ||
		!bytes.Equal(reliableHeading.FollowingAssessorsDigest, l.DiscordantLedger.FollowingAssessorsDigest) ||
		!bytes.Equal(reliableHeading.AgreementDigest, l.DiscordantLedger.AgreementDigest) ||
		!bytes.Equal(reliableHeading.PlatformDigest, l.DiscordantLedger.PlatformDigest) ||
		!bytes.Equal(reliableHeading.FinalOutcomesDigest, l.DiscordantLedger.FinalOutcomesDigest)
}

//
//
//
//
//
//
//
//
func (l *AgileCustomerOnslaughtProof) Digest() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, l.SharedAltitude)
	bz := make([]byte, tenderminthash.Extent+n)
	copy(bz[:tenderminthash.Extent-1], l.DiscordantLedger.Digest().Octets())
	copy(bz[tenderminthash.Extent:], buf)
	return tenderminthash.Sum(bz)
}

//
//
//
func (l *AgileCustomerOnslaughtProof) Altitude() int64 {
	return l.SharedAltitude
}

//
func (l *AgileCustomerOnslaughtProof) Text() string {
	return fmt.Sprintf(`REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED`,
		l.DiscordantLedger.Text(), l.SharedAltitude, l.TreacherousAssessors,
		l.SumBallotingPotency, l.Timestamp, l.Digest())
}

//
func (l *AgileCustomerOnslaughtProof) Moment() time.Time {
	return l.Timestamp
}

//
func (l *AgileCustomerOnslaughtProof) CertifyFundamental() error {
	if l.DiscordantLedger == nil {
		return errors.New("REDACTED")
	}

	//
	if l.DiscordantLedger.Heading == nil {
		return errors.New("REDACTED")
	}

	if l.SumBallotingPotency <= 0 {
		return errors.New("REDACTED")
	}

	if l.SharedAltitude <= 0 {
		return errors.New("REDACTED")
	}

	//
	//
	//
	if l.SharedAltitude > l.DiscordantLedger.Altitude {
		return fmt.Errorf("REDACTED",
			l.SharedAltitude, l.DiscordantLedger.Altitude)
	}

	if err := l.DiscordantLedger.CertifyFundamental(l.DiscordantLedger.SuccessionUUID); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return nil
}

//
func (l *AgileCustomerOnslaughtProof) TowardSchema() (*commitchema.AgileCustomerOnslaughtProof, error) {
	discordantLedger, err := l.DiscordantLedger.TowardSchema()
	if err != nil {
		return nil, err
	}

	byzantineValues := make([]*commitchema.Assessor, len(l.TreacherousAssessors))
	for idx, val := range l.TreacherousAssessors {
		valuepb, err := val.TowardSchema()
		if err != nil {
			return nil, err
		}
		byzantineValues[idx] = valuepb
	}

	return &commitchema.AgileCustomerOnslaughtProof{
		DiscordantLedger:    discordantLedger,
		SharedAltitude:        l.SharedAltitude,
		TreacherousAssessors: byzantineValues,
		SumBallotingPotency:    l.SumBallotingPotency,
		Timestamp:           l.Timestamp,
	}, nil
}

//
func AgileCustomerOnslaughtProofOriginatingSchema(lpb *commitchema.AgileCustomerOnslaughtProof) (*AgileCustomerOnslaughtProof, error) {
	if lpb == nil {
		return nil, strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}

	discordantLedger, err := AgileLedgerOriginatingSchema(lpb.DiscordantLedger)
	if err != nil {
		return nil, err
	}

	byzantineValues := make([]*Assessor, len(lpb.TreacherousAssessors))
	for idx, valuepb := range lpb.TreacherousAssessors {
		val, err := AssessorOriginatingSchema(valuepb)
		if err != nil {
			return nil, err
		}
		byzantineValues[idx] = val
	}

	l := &AgileCustomerOnslaughtProof{
		DiscordantLedger:    discordantLedger,
		SharedAltitude:        lpb.SharedAltitude,
		TreacherousAssessors: byzantineValues,
		SumBallotingPotency:    lpb.SumBallotingPotency,
		Timestamp:           lpb.Timestamp,
	}

	return l, l.CertifyFundamental()
}

//

//
type ProofCatalog []Proof

//
func (evl ProofCatalog) Digest() []byte {
	//
	//
	//
	proofByteslices := make([][]byte, len(evl))
	for i := 0; i < len(evl); i++ {
		//
		//
		proofByteslices[i] = evl[i].Octets()
	}
	return hashmap.DigestOriginatingOctetSegments(proofByteslices)
}

func (evl ProofCatalog) Text() string {
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
func (evl ProofCatalog) TowardIface() []iface.Malpractice {
	var el []iface.Malpractice
	for _, e := range evl {
		el = append(el, e.Iface()...)
	}
	return el
}

//

//
//
func ProofTowardSchema(proof Proof) (*commitchema.Proof, error) {
	if proof == nil {
		return nil, errors.New("REDACTED")
	}

	switch evi := proof.(type) {
	case *ReplicatedBallotProof:
		schemab := evi.TowardSchema()
		return &commitchema.Proof{
			Sum: &commitchema.Proof_Replicatedballotevidence{
				ReplicatedBallotProof: schemab,
			},
		}, nil

	case *AgileCustomerOnslaughtProof:
		schemab, err := evi.TowardSchema()
		if err != nil {
			return nil, err
		}
		return &commitchema.Proof{
			Sum: &commitchema.Proof_Agilecustomerattackproof{
				AgileCustomerOnslaughtProof: schemab,
			},
		}, nil

	default:
		return nil, fmt.Errorf("REDACTED", evi)
	}
}

//
//
func ProofOriginatingSchema(proof *commitchema.Proof) (Proof, error) {
	if proof == nil {
		return nil, errors.New("REDACTED")
	}

	switch evi := proof.Sum.(type) {
	case *commitchema.Proof_Replicatedballotevidence:
		return ReplicatedBallotProofOriginatingSchema(evi.ReplicatedBallotProof)
	case *commitchema.Proof_Agilecustomerattackproof:
		return AgileCustomerOnslaughtProofOriginatingSchema(evi.AgileCustomerOnslaughtProof)
	default:
		return nil, errors.New("REDACTED")
	}
}

func initialize() {
	strongmindjson.EnrollKind(&ReplicatedBallotProof{}, "REDACTED")
	strongmindjson.EnrollKind(&AgileCustomerOnslaughtProof{}, "REDACTED")
}

//

//
type FaultUnfitProof struct {
	Proof Proof
	Rationale   error
}

//
func FreshFaultUnfitProof(ev Proof, err error) *FaultUnfitProof {
	return &FaultUnfitProof{ev, err}
}

//
func (err *FaultUnfitProof) Failure() string {
	return fmt.Sprintf("REDACTED", err.Rationale, err.Proof)
}

//
type FaultProofOverrun struct {
	Max int64
	Got int64
}

//
func FreshFaultProofOverrun(max, got int64) *FaultProofOverrun {
	return &FaultProofOverrun{max, got}
}

//
func (err *FaultProofOverrun) Failure() string {
	return fmt.Sprintf("REDACTED", err.Max, err.Got)
}

//

//

//
//
func FreshSimulateReplicatedBallotProof(altitude int64, moment time.Time, successionUUID string) (*ReplicatedBallotProof, error) {
	val := FreshSimulatePRV()
	return FreshSimulateReplicatedBallotProofUsingAssessor(altitude, moment, val, successionUUID)
}

//
//
func FreshSimulateReplicatedBallotProofUsingAssessor(altitude int64, moment time.Time,
	pv PrivateAssessor, successionUUID string,
) (*ReplicatedBallotProof, error) {
	publicToken, err := pv.ObtainPublicToken()
	if err != nil {
		return nil, err
	}
	val := FreshAssessor(publicToken, 10)
	ballotAN := createSimulateBallot(altitude, 0, 0, publicToken.Location(), arbitraryLedgerUUID(), moment)
	vA := ballotAN.TowardSchema()
	err = pv.AttestBallot(successionUUID, vA)
	if err != nil {
		return nil, err
	}
	ballotAN.Notation = vA.Notation
	ballotBYTE := createSimulateBallot(altitude, 0, 0, publicToken.Location(), arbitraryLedgerUUID(), moment)
	vB := ballotBYTE.TowardSchema()
	err = pv.AttestBallot(successionUUID, vB)
	if err != nil {
		return nil, err
	}
	ballotBYTE.Notation = vB.Notation
	return FreshReplicatedBallotProof(ballotAN, ballotBYTE, moment, FreshAssessorAssign([]*Assessor{val}))
}

func createSimulateBallot(altitude int64, iteration, ordinal int32, location Location,
	ledgerUUID LedgerUUID, moment time.Time,
) *Ballot {
	return &Ballot{
		Kind:             commitchema.AttestedSignalKind(2),
		Altitude:           altitude,
		Iteration:            iteration,
		LedgerUUID:          ledgerUUID,
		Timestamp:        moment,
		AssessorLocation: location,
		AssessorOrdinal:   ordinal,
	}
}

func arbitraryLedgerUUID() LedgerUUID {
	return LedgerUUID{
		Digest: commitrand.Octets(tenderminthash.Extent),
		FragmentAssignHeading: FragmentAssignHeading{
			Sum: 1,
			Digest:  commitrand.Octets(tenderminthash.Extent),
		},
	}
}
