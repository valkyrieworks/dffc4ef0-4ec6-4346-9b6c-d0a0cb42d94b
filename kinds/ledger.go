package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

const (
	//
	//
	//
	MaximumHeadingOctets int64 = 626

	//
	//
	//
	//
	//
	//
	//
	//
	MaximumMarginForeachLedger int64 = 11
)

//
type Ledger struct {
	mtx commitchronize.Exclusion

	attestedDigest tendermintoctets.HexadecimalOctets //
	Heading       `json:"heading"`
	Data         `json:"data"`
	Proof     ProofData `json:"proof"`
	FinalEndorse   *Endorse      `json:"final_endorse"`
}

//
//
//
func (b *Ledger) CertifyFundamental() error {
	if b == nil {
		return errors.New("REDACTED")
	}

	b.mtx.Lock()
	defer b.mtx.Unlock()

	if err := b.Heading.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if b.FinalEndorse == nil {
		return errors.New("REDACTED")
	}
	if err := b.FinalEndorse.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if !bytes.Equal(b.FinalEndorseDigest, b.FinalEndorse.Digest()) {
		return fmt.Errorf("REDACTED",
			b.FinalEndorse.Digest(),
			b.FinalEndorseDigest,
		)
	}

	//
	if !bytes.Equal(b.DataDigest, b.Data.Digest()) {
		return fmt.Errorf(
			"REDACTED",
			b.Data.Digest(),
			b.DataDigest,
		)
	}

	//
	for i, ev := range b.Proof.Proof {
		if err := ev.CertifyFundamental(); err != nil {
			return fmt.Errorf("REDACTED", i, err)
		}
	}

	if !bytes.Equal(b.ProofDigest, b.Proof.Digest()) {
		return fmt.Errorf("REDACTED",
			b.ProofDigest,
			b.Proof.Digest(),
		)
	}

	return nil
}

//
func (b *Ledger) populateHeading() {
	if b.FinalEndorseDigest == nil {
		b.FinalEndorseDigest = b.FinalEndorse.Digest()
	}
	if b.DataDigest == nil {
		b.DataDigest = b.Data.Digest()
	}
	if b.ProofDigest == nil {
		b.ProofDigest = b.Proof.Digest()
	}
}

//
//
func (b *Ledger) Digest() tendermintoctets.HexadecimalOctets {
	if b == nil {
		return nil
	}
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if b.FinalEndorse == nil {
		return nil
	}
	if b.attestedDigest != nil {
		return b.attestedDigest
	}
	b.populateHeading()
	digest := b.Heading.Digest()
	b.attestedDigest = digest
	return digest
}

//
//
//
func (b *Ledger) CreateFragmentAssign(fragmentExtent uint32) (*FragmentAssign, error) {
	if b == nil {
		return nil, errors.New("REDACTED")
	}
	b.mtx.Lock()
	defer b.mtx.Unlock()

	pbb, err := b.TowardSchema()
	if err != nil {
		return nil, err
	}
	bz, err := proto.Marshal(pbb)
	if err != nil {
		return nil, err
	}
	return FreshFragmentAssignOriginatingData(bz, fragmentExtent), nil
}

//
//
func (b *Ledger) DigestsToward(digest []byte) bool {
	if len(digest) == 0 {
		return false
	}
	if b == nil {
		return false
	}
	return bytes.Equal(b.Digest(), digest)
}

//
func (b *Ledger) Extent() int {
	pbb, err := b.TowardSchema()
	if err != nil {
		return 0
	}

	return pbb.Extent()
}

//
//
//
func (b *Ledger) Text() string {
	return b.TextFormatted("REDACTED")
}

//
//
//
//
//
//
//
func (b *Ledger) TextFormatted(format string) string {
	if b == nil {
		return "REDACTED"
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		format, b.Heading.TextFormatted(format+"REDACTED"),
		format, b.Data.TextFormatted(format+"REDACTED"),
		format, b.Proof.TextFormatted(format+"REDACTED"),
		format, b.FinalEndorse.TextFormatted(format+"REDACTED"),
		format, b.Digest())
}

//
func (b *Ledger) TextBrief() string {
	if b == nil {
		return "REDACTED"
	}
	return fmt.Sprintf("REDACTED", b.Digest())
}

//
func (b *Ledger) TowardSchema() (*commitchema.Ledger, error) {
	if b == nil {
		return nil, errors.New("REDACTED")
	}

	pb := new(commitchema.Ledger)

	pb.Heading = *b.Heading.TowardSchema()
	pb.FinalEndorse = b.FinalEndorse.TowardSchema()
	pb.Data = b.Data.TowardSchema()

	schemaProof, err := b.Proof.TowardSchema()
	if err != nil {
		return nil, err
	}
	pb.Proof = *schemaProof

	return pb, nil
}

//
//
func LedgerOriginatingSchema(bp *commitchema.Ledger) (*Ledger, error) {
	if bp == nil {
		return nil, errors.New("REDACTED")
	}

	b := new(Ledger)
	h, err := HeadingOriginatingSchema(&bp.Heading)
	if err != nil {
		return nil, err
	}
	b.Heading = h
	data, err := DataOriginatingSchema(&bp.Data)
	if err != nil {
		return nil, err
	}
	b.Data = data
	if err := b.Proof.OriginatingSchema(&bp.Proof); err != nil {
		return nil, err
	}

	if bp.FinalEndorse != nil {
		lc, err := EndorseOriginatingSchema(bp.FinalEndorse)
		if err != nil {
			return nil, err
		}
		b.FinalEndorse = lc
	}

	return b, b.CertifyFundamental()
}

//

//
//
//
func MaximumDataOctets(maximumOctets, proofOctets int64, valuesTally int) int64 {
	maximumDataOctets := maximumOctets -
		MaximumMarginForeachLedger -
		MaximumHeadingOctets -
		MaximumEndorseOctets(valuesTally) -
		proofOctets

	if maximumDataOctets < 0 {
		panic(fmt.Sprintf(
			"REDACTED",
			maximumOctets,
			-(maximumDataOctets - maximumOctets),
		))
	}

	return maximumDataOctets
}

//
//
//
//
func MaximumDataOctetsNegativeProof(maximumOctets int64, valuesTally int) int64 {
	maximumDataOctets := maximumOctets -
		MaximumMarginForeachLedger -
		MaximumHeadingOctets -
		MaximumEndorseOctets(valuesTally)

	if maximumDataOctets < 0 {
		panic(fmt.Sprintf(
			"REDACTED",
			maximumOctets,
			-(maximumDataOctets - maximumOctets),
		))
	}

	return maximumDataOctets
}

//

//
//
//
//
//
type Heading struct {
	//
	Edition strongmindedition.Agreement `json:"edition"`
	SuccessionUUID string               `json:"succession_uuid"`
	Altitude  int64                `json:"altitude"`
	Moment    time.Time            `json:"moment"`

	//
	FinalLedgerUUID LedgerUUID `json:"final_ledger_uuid"`

	//
	FinalEndorseDigest tendermintoctets.HexadecimalOctets `json:"final_endorse_digest"` //
	DataDigest       tendermintoctets.HexadecimalOctets `json:"data_digest"`        //

	//
	AssessorsDigest     tendermintoctets.HexadecimalOctets `json:"assessors_digest"`      //
	FollowingAssessorsDigest tendermintoctets.HexadecimalOctets `json:"following_assessors_digest"` //
	AgreementDigest      tendermintoctets.HexadecimalOctets `json:"agreement_digest"`       //
	PlatformDigest            tendermintoctets.HexadecimalOctets `json:"application_digest"`             //
	//
	//
	FinalOutcomesDigest tendermintoctets.HexadecimalOctets `json:"final_outcomes_digest"`

	//
	ProofDigest    tendermintoctets.HexadecimalOctets `json:"proof_digest"`    //
	NominatorLocation Location           `json:"nominator_location"` //
}

//
//
func (h *Heading) Inhabit(
	edition strongmindedition.Agreement, successionUUID string,
	timestamp time.Time, finalLedgerUUID LedgerUUID,
	itemDigest, followingItemDigest []byte,
	agreementDigest, platformDigest, finalOutcomesDigest []byte,
	nominatorLocator Location,
) {
	h.Edition = edition
	h.SuccessionUUID = successionUUID
	h.Moment = timestamp
	h.FinalLedgerUUID = finalLedgerUUID
	h.AssessorsDigest = itemDigest
	h.FollowingAssessorsDigest = followingItemDigest
	h.AgreementDigest = agreementDigest
	h.PlatformDigest = platformDigest
	h.FinalOutcomesDigest = finalOutcomesDigest
	h.NominatorLocation = nominatorLocator
}

//
//
//
//
func (h Heading) CertifyFundamental() error {
	if h.Edition.Ledger != edition.LedgerScheme {
		return fmt.Errorf("REDACTED", h.Edition.Ledger, edition.LedgerScheme)
	}
	if len(h.SuccessionUUID) > MaximumSuccessionUUIDSize {
		return fmt.Errorf("REDACTED", len(h.SuccessionUUID), MaximumSuccessionUUIDSize)
	}

	if h.Altitude < 0 {
		return errors.New("REDACTED")
	} else if h.Altitude == 0 {
		return errors.New("REDACTED")
	}

	if err := h.FinalLedgerUUID.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if err := CertifyDigest(h.FinalEndorseDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if err := CertifyDigest(h.DataDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if err := CertifyDigest(h.ProofDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if len(h.NominatorLocation) != security.LocatorExtent {
		return fmt.Errorf(
			"REDACTED",
			len(h.NominatorLocation), security.LocatorExtent,
		)
	}

	//
	//
	if err := CertifyDigest(h.AssessorsDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := CertifyDigest(h.FollowingAssessorsDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := CertifyDigest(h.AgreementDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	//
	if err := CertifyDigest(h.FinalOutcomesDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return nil
}

//
//
//
//
//
//
func (h *Heading) Digest() tendermintoctets.HexadecimalOctets {
	if h == nil || len(h.AssessorsDigest) == 0 {
		return nil
	}
	hbz, err := h.Edition.Serialize()
	if err != nil {
		return nil
	}

	pbt, err := gogotypes.StdTimeMarshal(h.Moment)
	if err != nil {
		return nil
	}

	bufferbi := h.FinalLedgerUUID.TowardSchema()
	byzinfra, err := bufferbi.Serialize()
	if err != nil {
		return nil
	}
	return hashmap.DigestOriginatingOctetSegments([][]byte{
		hbz,
		codecSerialize(h.SuccessionUUID),
		codecSerialize(h.Altitude),
		pbt,
		byzinfra,
		codecSerialize(h.FinalEndorseDigest),
		codecSerialize(h.DataDigest),
		codecSerialize(h.AssessorsDigest),
		codecSerialize(h.FollowingAssessorsDigest),
		codecSerialize(h.AgreementDigest),
		codecSerialize(h.PlatformDigest),
		codecSerialize(h.FinalOutcomesDigest),
		codecSerialize(h.ProofDigest),
		codecSerialize(h.NominatorLocation),
	})
}

//
func (h *Heading) TextFormatted(format string) string {
	if h == nil {
		return "REDACTED"
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		format, h.Edition,
		format, h.SuccessionUUID,
		format, h.Altitude,
		format, h.Moment,
		format, h.FinalLedgerUUID,
		format, h.FinalEndorseDigest,
		format, h.DataDigest,
		format, h.AssessorsDigest,
		format, h.FollowingAssessorsDigest,
		format, h.PlatformDigest,
		format, h.AgreementDigest,
		format, h.FinalOutcomesDigest,
		format, h.ProofDigest,
		format, h.NominatorLocation,
		format, h.Digest(),
	)
}

//
func (h *Heading) TowardSchema() *commitchema.Heading {
	if h == nil {
		return nil
	}

	return &commitchema.Heading{
		Edition:            h.Edition,
		SuccessionUUID:            h.SuccessionUUID,
		Altitude:             h.Altitude,
		Moment:               h.Moment,
		FinalLedgerUuid:        h.FinalLedgerUUID.TowardSchema(),
		AssessorsDigest:     h.AssessorsDigest,
		FollowingAssessorsDigest: h.FollowingAssessorsDigest,
		AgreementDigest:      h.AgreementDigest,
		PlatformDigest:            h.PlatformDigest,
		DataDigest:           h.DataDigest,
		ProofDigest:       h.ProofDigest,
		FinalOutcomesDigest:    h.FinalOutcomesDigest,
		FinalEndorseDigest:     h.FinalEndorseDigest,
		NominatorLocation:    h.NominatorLocation,
	}
}

//
//
func HeadingOriginatingSchema(ph *commitchema.Heading) (Heading, error) {
	if ph == nil {
		return Heading{}, errors.New("REDACTED")
	}

	h := new(Heading)

	bi, err := LedgerUUIDOriginatingSchema(&ph.FinalLedgerUuid)
	if err != nil {
		return Heading{}, err
	}

	h.Edition = ph.Edition
	h.SuccessionUUID = ph.SuccessionUUID
	h.Altitude = ph.Altitude
	h.Moment = ph.Moment
	h.FinalLedgerUUID = *bi
	h.AssessorsDigest = ph.AssessorsDigest
	h.FollowingAssessorsDigest = ph.FollowingAssessorsDigest
	h.AgreementDigest = ph.AgreementDigest
	h.PlatformDigest = ph.PlatformDigest
	h.DataDigest = ph.DataDigest
	h.ProofDigest = ph.ProofDigest
	h.FinalOutcomesDigest = ph.FinalOutcomesDigest
	h.FinalEndorseDigest = ph.FinalEndorseDigest
	h.NominatorLocation = ph.NominatorLocation

	return *h, h.CertifyFundamental()
}

//

//
type LedgerUUIDMarker byte

const (
	//
	LedgerUUIDMarkerMissing LedgerUUIDMarker = iota + 1
	//
	LedgerUUIDMarkerEndorse
	//
	LedgerUUIDMarkerVoid
)

const (
	//
	MaximumEndorseMarginOctets int64 = 94

	//
	//
	maximumEndorseSignatureSchemaSerMargin = 4 + 1 + 1 + 1 + 3 //
	//
	//
	//
	MaximumEndorseSignatureOctets = 131 + maximumEndorseSignatureSchemaSerMargin
)

//
type EndorseSignature struct {
	LedgerUUIDMarker      LedgerUUIDMarker `json:"ledger_uuid_marker"`
	AssessorLocation Location     `json:"assessor_location"`
	Timestamp        time.Time   `json:"timestamp"`
	Notation        []byte      `json:"signing"`
}

func MaximumEndorseOctets(itemTally int) int64 {
	//
	const schemaIteratedAttributeLengthMargin int64 = 3
	//
	return MaximumEndorseMarginOctets + ((MaximumEndorseSignatureOctets + schemaIteratedAttributeLengthMargin) * int64(itemTally))
}

//
//
func FreshEndorseSignatureMissing() EndorseSignature {
	return EndorseSignature{
		LedgerUUIDMarker: LedgerUUIDMarkerMissing,
	}
}

//
//
//
//
//
//
func (cs EndorseSignature) Text() string {
	return fmt.Sprintf("REDACTED",
		tendermintoctets.Identifier(cs.Notation),
		tendermintoctets.Identifier(cs.AssessorLocation),
		cs.LedgerUUIDMarker,
		StandardMoment(cs.Timestamp))
}

//
//
func (cs EndorseSignature) LedgerUUID(endorseLedgerUUID LedgerUUID) LedgerUUID {
	var ledgerUUID LedgerUUID
	switch cs.LedgerUUIDMarker {
	case LedgerUUIDMarkerMissing:
		ledgerUUID = LedgerUUID{}
	case LedgerUUIDMarkerEndorse:
		ledgerUUID = endorseLedgerUUID
	case LedgerUUIDMarkerVoid:
		ledgerUUID = LedgerUUID{}
	default:
		panic(fmt.Sprintf("REDACTED", cs.LedgerUUIDMarker))
	}
	return ledgerUUID
}

//
func (cs EndorseSignature) CertifyFundamental() error {
	switch cs.LedgerUUIDMarker {
	case LedgerUUIDMarkerMissing:
	case LedgerUUIDMarkerEndorse:
	case LedgerUUIDMarkerVoid:
	default:
		return fmt.Errorf("REDACTED", cs.LedgerUUIDMarker)
	}

	switch cs.LedgerUUIDMarker {
	case LedgerUUIDMarkerMissing:
		if len(cs.AssessorLocation) != 0 {
			return errors.New("REDACTED")
		}
		if !cs.Timestamp.IsZero() {
			return errors.New("REDACTED")
		}
		if len(cs.Notation) != 0 {
			return errors.New("REDACTED")
		}
	default:
		if len(cs.AssessorLocation) != security.LocatorExtent {
			return fmt.Errorf("REDACTED",
				security.LocatorExtent,
				len(cs.AssessorLocation),
			)
		}
		//
		if len(cs.Notation) == 0 {
			return errors.New("REDACTED")
		}
		if len(cs.Notation) > MaximumNotationExtent {
			return fmt.Errorf("REDACTED", MaximumNotationExtent)
		}
	}

	return nil
}

//
func (cs *EndorseSignature) TowardSchema() *commitchema.EndorseSignature {
	if cs == nil {
		return nil
	}

	return &commitchema.EndorseSignature{
		LedgerUuidMarker:      commitchema.LedgerUUIDMarker(cs.LedgerUUIDMarker),
		AssessorLocation: cs.AssessorLocation,
		Timestamp:        cs.Timestamp,
		Notation:        cs.Notation,
	}
}

//
//
func (cs *EndorseSignature) OriginatingSchema(csp commitchema.EndorseSignature) error {
	cs.LedgerUUIDMarker = LedgerUUIDMarker(csp.LedgerUuidMarker)
	cs.AssessorLocation = csp.AssessorLocation
	cs.Timestamp = csp.Timestamp
	cs.Notation = csp.Notation

	return cs.CertifyFundamental()
}

//

//
//
type ExpandedEndorseSignature struct {
	EndorseSignature                 //
	Addition          []byte //
	AdditionNotation []byte //
}

//
//
func FreshExpandedEndorseSignatureMissing() ExpandedEndorseSignature {
	return ExpandedEndorseSignature{EndorseSignature: FreshEndorseSignatureMissing()}
}

//
//
//
//
//
func (ecs ExpandedEndorseSignature) Text() string {
	return fmt.Sprintf("REDACTED",
		ecs.EndorseSignature,
		tendermintoctets.Identifier(ecs.Addition),
		tendermintoctets.Identifier(ecs.AdditionNotation),
	)
}

//
func (ecs ExpandedEndorseSignature) CertifyFundamental() error {
	if err := ecs.EndorseSignature.CertifyFundamental(); err != nil {
		return err
	}

	if ecs.LedgerUUIDMarker == LedgerUUIDMarkerEndorse {
		if len(ecs.Addition) > MaximumBallotAdditionExtent {
			return fmt.Errorf("REDACTED", MaximumBallotAdditionExtent)
		}
		if len(ecs.AdditionNotation) > MaximumNotationExtent {
			return fmt.Errorf("REDACTED", MaximumNotationExtent)
		}
		return nil
	}

	if len(ecs.AdditionNotation) == 0 && len(ecs.Addition) != 0 {
		return errors.New("REDACTED")
	}
	return nil
}

//
//
func (ecs ExpandedEndorseSignature) AssureAddition(addnActivated bool) error {
	if addnActivated {
		if ecs.LedgerUUIDMarker == LedgerUUIDMarkerEndorse && len(ecs.AdditionNotation) == 0 {
			return fmt.Errorf("REDACTED",
				ecs.AssessorLocation.Text(),
				ecs.Timestamp,
			)
		}
		if ecs.LedgerUUIDMarker != LedgerUUIDMarkerEndorse && len(ecs.Addition) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.AssessorLocation.Text(),
				ecs.Timestamp,
			)
		}
		if ecs.LedgerUUIDMarker != LedgerUUIDMarkerEndorse && len(ecs.AdditionNotation) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.AssessorLocation.Text(),
				ecs.Timestamp,
			)
		}
	} else {
		if len(ecs.Addition) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.AssessorLocation.Text(),
				ecs.Timestamp,
			)
		}
		if len(ecs.AdditionNotation) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.AssessorLocation.Text(),
				ecs.Timestamp,
			)
		}
	}
	return nil
}

//
func (ecs *ExpandedEndorseSignature) TowardSchema() *commitchema.ExpandedEndorseSignature {
	if ecs == nil {
		return nil
	}

	return &commitchema.ExpandedEndorseSignature{
		LedgerUuidMarker:        commitchema.LedgerUUIDMarker(ecs.LedgerUUIDMarker),
		AssessorLocation:   ecs.AssessorLocation,
		Timestamp:          ecs.Timestamp,
		Notation:          ecs.Notation,
		Addition:          ecs.Addition,
		AdditionNotation: ecs.AdditionNotation,
	}
}

//
//
//
func (ecs *ExpandedEndorseSignature) OriginatingSchema(endcontextswitchproc commitchema.ExpandedEndorseSignature) error {
	ecs.LedgerUUIDMarker = LedgerUUIDMarker(endcontextswitchproc.LedgerUuidMarker)
	ecs.AssessorLocation = endcontextswitchproc.AssessorLocation
	ecs.Timestamp = endcontextswitchproc.Timestamp
	ecs.Notation = endcontextswitchproc.Notation
	ecs.Addition = endcontextswitchproc.Addition
	ecs.AdditionNotation = endcontextswitchproc.AdditionNotation

	return ecs.CertifyFundamental()
}

//

//
//
type Endorse struct {
	//
	//
	//
	//
	Altitude     int64       `json:"altitude"`
	Iteration      int32       `json:"iteration"`
	LedgerUUID    LedgerUUID     `json:"ledger_uuid"`
	Notations []EndorseSignature `json:"notations"`

	//
	//
	//
	digest tendermintoctets.HexadecimalOctets
}

//
func (endorse *Endorse) Replicate() *Endorse {
	signatures := make([]EndorseSignature, len(endorse.Notations))
	copy(signatures, endorse.Notations)
	xchangeDuplicate := *endorse
	xchangeDuplicate.Notations = signatures
	return &xchangeDuplicate
}

//
//
//
//
//
func (endorse *Endorse) FetchBallot(itemOffset int32) *Ballot {
	endorseSignature := endorse.Notations[itemOffset]
	return &Ballot{
		Kind:             commitchema.PreendorseKind,
		Altitude:           endorse.Altitude,
		Iteration:            endorse.Iteration,
		LedgerUUID:          endorseSignature.LedgerUUID(endorse.LedgerUUID),
		Timestamp:        endorseSignature.Timestamp,
		AssessorLocation: endorseSignature.AssessorLocation,
		AssessorOrdinal:   itemOffset,
		Notation:        endorseSignature.Notation,
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
//
func (endorse *Endorse) BallotAttestOctets(successionUUID string, itemOffset int32) []byte {
	v := endorse.FetchBallot(itemOffset).TowardSchema()
	return BallotAttestOctets(successionUUID, v)
}

//
func (endorse *Endorse) Extent() int {
	if endorse == nil {
		return 0
	}
	return len(endorse.Notations)
}

//
//
func (endorse *Endorse) CertifyFundamental() error {
	if endorse.Altitude < 0 {
		return errors.New("REDACTED")
	}
	if endorse.Iteration < 0 {
		return errors.New("REDACTED")
	}

	if endorse.Altitude >= 1 {
		if endorse.LedgerUUID.EqualsNull() {
			return errors.New("REDACTED")
		}

		if len(endorse.Notations) == 0 {
			return errors.New("REDACTED")
		}
		for i, endorseSignature := range endorse.Notations {
			if err := endorseSignature.CertifyFundamental(); err != nil {
				return fmt.Errorf("REDACTED", i, err)
			}
		}
	}
	return nil
}

//
func (endorse *Endorse) Digest() tendermintoctets.HexadecimalOctets {
	if endorse == nil {
		return nil
	}
	if endorse.digest == nil {
		bs := make([][]byte, len(endorse.Notations))
		for i, endorseSignature := range endorse.Notations {
			buffercontextswitch := endorseSignature.TowardSchema()
			bz, err := buffercontextswitch.Serialize()
			if err != nil {
				panic(err)
			}

			bs[i] = bz
		}
		endorse.digest = hashmap.DigestOriginatingOctetSegments(bs)
	}
	return endorse.digest
}

//
//
//
//
//
func (endorse *Endorse) EncapsulatedExpandedEndorse() *ExpandedEndorse {
	cs := make([]ExpandedEndorseSignature, len(endorse.Notations))
	for idx, s := range endorse.Notations {
		cs[idx] = ExpandedEndorseSignature{
			EndorseSignature: s,
		}
	}
	return &ExpandedEndorse{
		Altitude:             endorse.Altitude,
		Iteration:              endorse.Iteration,
		LedgerUUID:            endorse.LedgerUUID,
		ExpandedNotations: cs,
	}
}

//
func (endorse *Endorse) TextFormatted(format string) string {
	if endorse == nil {
		return "REDACTED"
	}
	endorseSignatureTexts := make([]string, len(endorse.Notations))
	for i, endorseSignature := range endorse.Notations {
		endorseSignatureTexts[i] = endorseSignature.Text()
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDd
REDACTEDd
REDACTEDv
REDACTED:
REDACTEDv
REDACTED`,
		format, endorse.Altitude,
		format, endorse.Iteration,
		format, endorse.LedgerUUID,
		format,
		format, strings.Join(endorseSignatureTexts, "REDACTED"+format+"REDACTED"),
		format, endorse.digest)
}

//
func (endorse *Endorse) TowardSchema() *commitchema.Endorse {
	if endorse == nil {
		return nil
	}

	c := new(commitchema.Endorse)
	signatures := make([]commitchema.EndorseSignature, len(endorse.Notations))
	for i := range endorse.Notations {
		signatures[i] = *endorse.Notations[i].TowardSchema()
	}
	c.Notations = signatures

	c.Altitude = endorse.Altitude
	c.Iteration = endorse.Iteration
	c.LedgerUUID = endorse.LedgerUUID.TowardSchema()

	return c
}

//
//
func EndorseOriginatingSchema(cp *commitchema.Endorse) (*Endorse, error) {
	if cp == nil {
		return nil, errors.New("REDACTED")
	}

	endorse := new(Endorse)

	bi, err := LedgerUUIDOriginatingSchema(&cp.LedgerUUID)
	if err != nil {
		return nil, err
	}

	signatures := make([]EndorseSignature, len(cp.Notations))
	for i := range cp.Notations {
		if err := signatures[i].OriginatingSchema(cp.Notations[i]); err != nil {
			return nil, err
		}
	}
	endorse.Notations = signatures

	endorse.Altitude = cp.Altitude
	endorse.Iteration = cp.Iteration
	endorse.LedgerUUID = *bi

	return endorse, endorse.CertifyFundamental()
}

//

//
//
type ExpandedEndorse struct {
	Altitude             int64
	Iteration              int32
	LedgerUUID            LedgerUUID
	ExpandedNotations []ExpandedEndorseSignature

	digitSeries *digits.DigitSeries
}

//
func (ec *ExpandedEndorse) Replicate() *ExpandedEndorse {
	signatures := make([]ExpandedEndorseSignature, len(ec.ExpandedNotations))
	copy(signatures, ec.ExpandedNotations)
	ecc := *ec
	ecc.ExpandedNotations = signatures
	return &ecc
}

//
//
//
//
func (ec *ExpandedEndorse) TowardExpandedBallotAssign(successionUUID string, values *AssessorAssign) *BallotAssign {
	ballotAssign := FreshExpandedBallotAssign(successionUUID, ec.Altitude, ec.Iteration, commitchema.PreendorseKind, values)
	ec.appendSignaturesTowardBallotAssign(ballotAssign)
	return ballotAssign
}

//
func (ec *ExpandedEndorse) appendSignaturesTowardBallotAssign(ballotAssign *BallotAssign) {
	for idx, ecs := range ec.ExpandedNotations {
		if ecs.LedgerUUIDMarker == LedgerUUIDMarkerMissing {
			continue //
		}
		ballot := ec.ObtainExpandedBallot(int32(idx))
		if err := ballot.CertifyFundamental(); err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
		appended, err := ballotAssign.AppendBallot(ballot)
		if !appended || err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
	}
}

//
//
//
func (endorse *Endorse) TowardBallotAssign(successionUUID string, values *AssessorAssign) *BallotAssign {
	ballotAssign := FreshBallotAssign(successionUUID, endorse.Altitude, endorse.Iteration, commitchema.PreendorseKind, values)
	for idx, cs := range endorse.Notations {
		if cs.LedgerUUIDMarker == LedgerUUIDMarkerMissing {
			continue //
		}
		ballot := endorse.FetchBallot(int32(idx))
		if err := ballot.CertifyFundamental(); err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
		appended, err := ballotAssign.AppendBallot(ballot)
		if !appended || err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
	}
	return ballotAssign
}

//
//
func (ec *ExpandedEndorse) AssureAdditions(addnActivated bool) error {
	for _, ecs := range ec.ExpandedNotations {
		if err := ecs.AssureAddition(addnActivated); err != nil {
			return err
		}
	}
	return nil
}

//
//
func (ec *ExpandedEndorse) TowardEndorse() *Endorse {
	cs := make([]EndorseSignature, len(ec.ExpandedNotations))
	for idx, ecs := range ec.ExpandedNotations {
		cs[idx] = ecs.EndorseSignature
	}
	return &Endorse{
		Altitude:     ec.Altitude,
		Iteration:      ec.Iteration,
		LedgerUUID:    ec.LedgerUUID,
		Notations: cs,
	}
}

//
//
//
func (ec *ExpandedEndorse) ObtainExpandedBallot(itemOrdinal int32) *Ballot {
	ecs := ec.ExpandedNotations[itemOrdinal]
	return &Ballot{
		Kind:               commitchema.PreendorseKind,
		Altitude:             ec.Altitude,
		Iteration:              ec.Iteration,
		LedgerUUID:            ecs.LedgerUUID(ec.LedgerUUID),
		Timestamp:          ecs.Timestamp,
		AssessorLocation:   ecs.AssessorLocation,
		AssessorOrdinal:     itemOrdinal,
		Notation:          ecs.Notation,
		Addition:          ecs.Addition,
		AdditionNotation: ecs.AdditionNotation,
	}
}

//
//
//
func (ec *ExpandedEndorse) Kind() byte { return byte(commitchema.PreendorseKind) }

//
//
func (ec *ExpandedEndorse) ObtainAltitude() int64 { return ec.Altitude }

//
//
func (ec *ExpandedEndorse) ObtainIteration() int32 { return ec.Iteration }

//
//
func (ec *ExpandedEndorse) Extent() int {
	if ec == nil {
		return 0
	}
	return len(ec.ExpandedNotations)
}

//
//
//
func (ec *ExpandedEndorse) DigitSeries() *digits.DigitSeries {
	if ec.digitSeries == nil {
		primaryDigitProc := func(i int) bool {
			//
			//
			return ec.ExpandedNotations[i].LedgerUUIDMarker != LedgerUUIDMarkerMissing
		}
		ec.digitSeries = digits.FreshDigitSeriesOriginatingProc(len(ec.ExpandedNotations), primaryDigitProc)
	}
	return ec.digitSeries
}

//
//
//
func (ec *ExpandedEndorse) ObtainViaOrdinal(itemOffset int32) *Ballot {
	return ec.ObtainExpandedBallot(itemOffset)
}

//
//
func (ec *ExpandedEndorse) EqualsEndorse() bool {
	return len(ec.ExpandedNotations) != 0
}

//
//
func (ec *ExpandedEndorse) CertifyFundamental() error {
	if ec.Altitude < 0 {
		return errors.New("REDACTED")
	}
	if ec.Iteration < 0 {
		return errors.New("REDACTED")
	}

	if ec.Altitude >= 1 {
		if ec.LedgerUUID.EqualsNull() {
			return errors.New("REDACTED")
		}

		if len(ec.ExpandedNotations) == 0 {
			return errors.New("REDACTED")
		}
		for i, addnEndorseSignature := range ec.ExpandedNotations {
			if err := addnEndorseSignature.CertifyFundamental(); err != nil {
				return fmt.Errorf("REDACTED", i, err)
			}
		}
	}
	return nil
}

//
func (ec *ExpandedEndorse) TowardSchema() *commitchema.ExpandedEndorse {
	if ec == nil {
		return nil
	}

	c := new(commitchema.ExpandedEndorse)
	signatures := make([]commitchema.ExpandedEndorseSignature, len(ec.ExpandedNotations))
	for i := range ec.ExpandedNotations {
		signatures[i] = *ec.ExpandedNotations[i].TowardSchema()
	}
	c.ExpandedNotations = signatures

	c.Altitude = ec.Altitude
	c.Iteration = ec.Iteration
	c.LedgerUUID = ec.LedgerUUID.TowardSchema()

	return c
}

//
//
func ExpandedEndorseOriginatingSchema(ecp *commitchema.ExpandedEndorse) (*ExpandedEndorse, error) {
	if ecp == nil {
		return nil, errors.New("REDACTED")
	}

	addnEndorse := new(ExpandedEndorse)

	bi, err := LedgerUUIDOriginatingSchema(&ecp.LedgerUUID)
	if err != nil {
		return nil, err
	}

	signatures := make([]ExpandedEndorseSignature, len(ecp.ExpandedNotations))
	for i := range ecp.ExpandedNotations {
		if err := signatures[i].OriginatingSchema(ecp.ExpandedNotations[i]); err != nil {
			return nil, err
		}
	}
	addnEndorse.ExpandedNotations = signatures
	addnEndorse.Altitude = ecp.Altitude
	addnEndorse.Iteration = ecp.Iteration
	addnEndorse.LedgerUUID = *bi

	return addnEndorse, addnEndorse.CertifyFundamental()
}

//

//
type Data struct {
	//
	//
	//
	Txs Txs `json:"txs"`

	//
	digest tendermintoctets.HexadecimalOctets
}

//
func (data *Data) Digest() tendermintoctets.HexadecimalOctets {
	if data == nil {
		return (Txs{}).Digest()
	}
	if data.digest == nil {
		data.digest = data.Txs.Digest() //
	}
	return data.digest
}

//
func (data *Data) TextFormatted(format string) string {
	if data == nil {
		return "REDACTED"
	}
	transferTexts := make([]string, strongarithmetic.MinimumInteger(len(data.Txs), 21))
	for i, tx := range data.Txs {
		if i == 20 {
			transferTexts[i] = fmt.Sprintf("REDACTED", len(data.Txs))
			break
		}
		transferTexts[i] = fmt.Sprintf("REDACTED", tx.Digest(), len(tx))
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTED`,
		format, strings.Join(transferTexts, "REDACTED"+format+"REDACTED"),
		format, data.digest)
}

//
func (data *Data) TowardSchema() commitchema.Data {
	tp := new(commitchema.Data)

	if len(data.Txs) > 0 {
		transferByteslices := make([][]byte, len(data.Txs))
		for i := range data.Txs {
			transferByteslices[i] = data.Txs[i]
		}
		tp.Txs = transferByteslices
	}

	return *tp
}

//
//
func DataOriginatingSchema(dp *commitchema.Data) (Data, error) {
	if dp == nil {
		return Data{}, errors.New("REDACTED")
	}
	data := new(Data)

	if len(dp.Txs) > 0 {
		transferByteslices := make(Txs, len(dp.Txs))
		for i := range dp.Txs {
			transferByteslices[i] = Tx(dp.Txs[i])
		}
		data.Txs = transferByteslices
	} else {
		data.Txs = Txs{}
	}

	return *data, nil
}

//

//
type ProofData struct {
	Proof ProofCatalog `json:"proof"`

	//
	digest     tendermintoctets.HexadecimalOctets
	octetExtent int64
}

//
func (data *ProofData) Digest() tendermintoctets.HexadecimalOctets {
	if data.digest == nil {
		data.digest = data.Proof.Digest()
	}
	return data.digest
}

//
func (data *ProofData) OctetExtent() int64 {
	if data.octetExtent == 0 && len(data.Proof) != 0 {
		pb, err := data.TowardSchema()
		if err != nil {
			panic(err)
		}
		data.octetExtent = int64(pb.Extent())
	}
	return data.octetExtent
}

//
func (data *ProofData) TextFormatted(format string) string {
	if data == nil {
		return "REDACTED"
	}
	occurenceTexts := make([]string, strongarithmetic.MinimumInteger(len(data.Proof), 21))
	for i, ev := range data.Proof {
		if i == 20 {
			occurenceTexts[i] = fmt.Sprintf("REDACTED", len(data.Proof))
			break
		}
		occurenceTexts[i] = "REDACTED" + ev.Text()
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTED`,
		format, strings.Join(occurenceTexts, "REDACTED"+format+"REDACTED"),
		format, data.digest)
}

//
func (data *ProofData) TowardSchema() (*commitchema.ProofCatalog, error) {
	if data == nil {
		return nil, errors.New("REDACTED")
	}

	evi := new(commitchema.ProofCatalog)
	evidenceByteslices := make([]commitchema.Proof, len(data.Proof))
	for i := range data.Proof {
		schemaEvidence, err := ProofTowardSchema(data.Proof[i])
		if err != nil {
			return nil, err
		}
		evidenceByteslices[i] = *schemaEvidence
	}
	evi.Proof = evidenceByteslices

	return evi, nil
}

//
func (data *ProofData) OriginatingSchema(evidenceData *commitchema.ProofCatalog) error {
	if evidenceData == nil {
		return errors.New("REDACTED")
	}

	evidenceByteslices := make(ProofCatalog, len(evidenceData.Proof))
	for i := range evidenceData.Proof {
		evi, err := ProofOriginatingSchema(&evidenceData.Proof[i])
		if err != nil {
			return err
		}
		evidenceByteslices[i] = evi
	}
	data.Proof = evidenceByteslices
	data.octetExtent = int64(evidenceData.Extent())

	return nil
}

//

//
type LedgerUUID struct {
	Digest          tendermintoctets.HexadecimalOctets `json:"digest"`
	FragmentAssignHeading FragmentAssignHeading     `json:"fragments"`
}

//
func (ledgerUUID LedgerUUID) Matches(another LedgerUUID) bool {
	return bytes.Equal(ledgerUUID.Digest, another.Digest) &&
		ledgerUUID.FragmentAssignHeading.Matches(another.FragmentAssignHeading)
}

//
func (ledgerUUID LedgerUUID) Key() string {
	bufferprocess := ledgerUUID.FragmentAssignHeading.TowardSchema()
	bz, err := bufferprocess.Serialize()
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(string(ledgerUUID.Digest), string(bz))
}

//
func (ledgerUUID LedgerUUID) CertifyFundamental() error {
	//
	if err := CertifyDigest(ledgerUUID.Digest); err != nil {
		return fmt.Errorf("REDACTED")
	}
	if err := ledgerUUID.FragmentAssignHeading.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (ledgerUUID LedgerUUID) EqualsNull() bool {
	return len(ledgerUUID.Digest) == 0 &&
		ledgerUUID.FragmentAssignHeading.EqualsNull()
}

//
func (ledgerUUID LedgerUUID) EqualsFinish() bool {
	return len(ledgerUUID.Digest) == tenderminthash.Extent &&
		ledgerUUID.FragmentAssignHeading.Sum > 0 &&
		len(ledgerUUID.FragmentAssignHeading.Digest) == tenderminthash.Extent
}

//
//
//
//
//
//
func (ledgerUUID LedgerUUID) Text() string {
	return fmt.Sprintf("REDACTED", ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading)
}

//
func (ledgerUUID *LedgerUUID) TowardSchema() commitchema.LedgerUUID {
	if ledgerUUID == nil {
		return commitchema.LedgerUUID{}
	}

	return commitchema.LedgerUUID{
		Digest:          ledgerUUID.Digest,
		FragmentAssignHeading: ledgerUUID.FragmentAssignHeading.TowardSchema(),
	}
}

//
//
func LedgerUUIDOriginatingSchema(bID *commitchema.LedgerUUID) (*LedgerUUID, error) {
	if bID == nil {
		return nil, errors.New("REDACTED")
	}

	ledgerUUID := new(LedgerUUID)
	ph, err := FragmentAssignHeadingOriginatingSchema(&bID.FragmentAssignHeading)
	if err != nil {
		return nil, err
	}

	ledgerUUID.FragmentAssignHeading = *ph
	ledgerUUID.Digest = bID.Digest

	return ledgerUUID, ledgerUUID.CertifyFundamental()
}

//
//
func SchemaLedgerUUIDEqualsVoid(bID *commitchema.LedgerUUID) bool {
	return len(bID.Digest) == 0 && SchemaFragmentAssignHeadingEqualsNull(&bID.FragmentAssignHeading)
}
