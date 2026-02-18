package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/utils/bits"
	cometbytes "github.com/valkyrieworks/utils/octets"
	cometmath "github.com/valkyrieworks/utils/math"
	engineconnect "github.com/valkyrieworks/utils/align"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/release"
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
	MaximumBurdenForLedger int64 = 11
)

//
type Ledger struct {
	mtx engineconnect.Lock

	certifiedDigest cometbytes.HexOctets //
	Heading       `json:"heading"`
	Data         `json:"data"`
	Proof     ProofData `json:"proof"`
	FinalEndorse   *Endorse      `json:"final_endorse"`
}

//
//
//
func (b *Ledger) CertifySimple() error {
	if b == nil {
		return errors.New("REDACTED")
	}

	b.mtx.Lock()
	defer b.mtx.Unlock()

	if err := b.Heading.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if b.FinalEndorse == nil {
		return errors.New("REDACTED")
	}
	if err := b.FinalEndorse.CertifySimple(); err != nil {
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
		if err := ev.CertifySimple(); err != nil {
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
func (b *Ledger) Digest() cometbytes.HexOctets {
	if b == nil {
		return nil
	}
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if b.FinalEndorse == nil {
		return nil
	}
	if b.certifiedDigest != nil {
		return b.certifiedDigest
	}
	b.populateHeading()
	digest := b.Heading.Digest()
	b.certifiedDigest = digest
	return digest
}

//
//
//
func (b *Ledger) CreateSegmentAssign(segmentVolume uint32) (*SegmentCollection, error) {
	if b == nil {
		return nil, errors.New("REDACTED")
	}
	b.mtx.Lock()
	defer b.mtx.Unlock()

	pbb, err := b.ToSchema()
	if err != nil {
		return nil, err
	}
	bz, err := proto.Marshal(pbb)
	if err != nil {
		return nil, err
	}
	return NewSegmentCollectionFromData(bz, segmentVolume), nil
}

//
//
func (b *Ledger) DigestsTo(digest []byte) bool {
	if len(digest) == 0 {
		return false
	}
	if b == nil {
		return false
	}
	return bytes.Equal(b.Digest(), digest)
}

//
func (b *Ledger) Volume() int {
	pbb, err := b.ToSchema()
	if err != nil {
		return 0
	}

	return pbb.Volume()
}

//
//
//
func (b *Ledger) String() string {
	return b.StringIndented("REDACTED")
}

//
//
//
//
//
//
//
func (b *Ledger) StringIndented(indent string) string {
	if b == nil {
		return "REDACTED"
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		indent, b.Heading.StringIndented(indent+"REDACTED"),
		indent, b.Data.StringIndented(indent+"REDACTED"),
		indent, b.Proof.StringIndented(indent+"REDACTED"),
		indent, b.FinalEndorse.StringIndented(indent+"REDACTED"),
		indent, b.Digest())
}

//
func (b *Ledger) StringBrief() string {
	if b == nil {
		return "REDACTED"
	}
	return fmt.Sprintf("REDACTED", b.Digest())
}

//
func (b *Ledger) ToSchema() (*engineproto.Ledger, error) {
	if b == nil {
		return nil, errors.New("REDACTED")
	}

	pb := new(engineproto.Ledger)

	pb.Heading = *b.Heading.ToSchema()
	pb.FinalEndorse = b.FinalEndorse.ToSchema()
	pb.Data = b.Data.ToSchema()

	schemaProof, err := b.Proof.ToSchema()
	if err != nil {
		return nil, err
	}
	pb.Proof = *schemaProof

	return pb, nil
}

//
//
func LedgerFromSchema(bp *engineproto.Ledger) (*Ledger, error) {
	if bp == nil {
		return nil, errors.New("REDACTED")
	}

	b := new(Ledger)
	h, err := HeadingFromSchema(&bp.Heading)
	if err != nil {
		return nil, err
	}
	b.Heading = h
	data, err := DataFromSchema(&bp.Data)
	if err != nil {
		return nil, err
	}
	b.Data = data
	if err := b.Proof.FromSchema(&bp.Proof); err != nil {
		return nil, err
	}

	if bp.FinalEndorse != nil {
		lc, err := EndorseFromSchema(bp.FinalEndorse)
		if err != nil {
			return nil, err
		}
		b.FinalEndorse = lc
	}

	return b, b.CertifySimple()
}

//

//
//
//
func MaximumDataOctets(maximumOctets, proofOctets int64, valuesTally int) int64 {
	maximumDataOctets := maximumOctets -
		MaximumBurdenForLedger -
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
func MaximumDataOctetsNoProof(maximumOctets int64, valuesTally int) int64 {
	maximumDataOctets := maximumOctets -
		MaximumBurdenForLedger -
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
	Release cometrelease.Agreement `json:"release"`
	LedgerUID string               `json:"series_uid"`
	Level  int64                `json:"level"`
	Time    time.Time            `json:"moment"`

	//
	FinalLedgerUID LedgerUID `json:"final_ledger_uid"`

	//
	FinalEndorseDigest cometbytes.HexOctets `json:"final_endorse_digest"` //
	DataDigest       cometbytes.HexOctets `json:"data_digest"`        //

	//
	RatifiersDigest     cometbytes.HexOctets `json:"ratifiers_digest"`      //
	FollowingRatifiersDigest cometbytes.HexOctets `json:"following_ratifiers_digest"` //
	AgreementDigest      cometbytes.HexOctets `json:"agreement_digest"`       //
	ApplicationDigest            cometbytes.HexOctets `json:"application_digest"`             //
	//
	//
	FinalOutcomesDigest cometbytes.HexOctets `json:"final_outcomes_digest"`

	//
	ProofDigest    cometbytes.HexOctets `json:"proof_digest"`    //
	RecommenderLocation Location           `json:"recommender_location"` //
}

//
//
func (h *Heading) Fill(
	release cometrelease.Agreement, ledgerUID string,
	timestamp time.Time, finalLedgerUID LedgerUID,
	valueDigest, followingValueDigest []byte,
	agreementDigest, applicationDigest, finalOutcomesDigest []byte,
	recommenderLocation Location,
) {
	h.Release = release
	h.LedgerUID = ledgerUID
	h.Time = timestamp
	h.FinalLedgerUID = finalLedgerUID
	h.RatifiersDigest = valueDigest
	h.FollowingRatifiersDigest = followingValueDigest
	h.AgreementDigest = agreementDigest
	h.ApplicationDigest = applicationDigest
	h.FinalOutcomesDigest = finalOutcomesDigest
	h.RecommenderLocation = recommenderLocation
}

//
//
//
//
func (h Heading) CertifySimple() error {
	if h.Release.Ledger != release.LedgerProtocol {
		return fmt.Errorf("REDACTED", h.Release.Ledger, release.LedgerProtocol)
	}
	if len(h.LedgerUID) > MaximumSeriesUIDSize {
		return fmt.Errorf("REDACTED", len(h.LedgerUID), MaximumSeriesUIDSize)
	}

	if h.Level < 0 {
		return errors.New("REDACTED")
	} else if h.Level == 0 {
		return errors.New("REDACTED")
	}

	if err := h.FinalLedgerUID.CertifySimple(); err != nil {
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

	if len(h.RecommenderLocation) != vault.LocationVolume {
		return fmt.Errorf(
			"REDACTED",
			len(h.RecommenderLocation), vault.LocationVolume,
		)
	}

	//
	//
	if err := CertifyDigest(h.RatifiersDigest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := CertifyDigest(h.FollowingRatifiersDigest); err != nil {
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
func (h *Heading) Digest() cometbytes.HexOctets {
	if h == nil || len(h.RatifiersDigest) == 0 {
		return nil
	}
	hbz, err := h.Release.Serialize()
	if err != nil {
		return nil
	}

	pbt, err := gogotypes.StdTimeMarshal(h.Time)
	if err != nil {
		return nil
	}

	pbbi := h.FinalLedgerUID.ToSchema()
	bzbi, err := pbbi.Serialize()
	if err != nil {
		return nil
	}
	return merkle.DigestFromOctetSegments([][]byte{
		hbz,
		cdcEncode(h.LedgerUID),
		cdcEncode(h.Level),
		pbt,
		bzbi,
		cdcEncode(h.FinalEndorseDigest),
		cdcEncode(h.DataDigest),
		cdcEncode(h.RatifiersDigest),
		cdcEncode(h.FollowingRatifiersDigest),
		cdcEncode(h.AgreementDigest),
		cdcEncode(h.ApplicationDigest),
		cdcEncode(h.FinalOutcomesDigest),
		cdcEncode(h.ProofDigest),
		cdcEncode(h.RecommenderLocation),
	})
}

//
func (h *Heading) StringIndented(indent string) string {
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
		indent, h.Release,
		indent, h.LedgerUID,
		indent, h.Level,
		indent, h.Time,
		indent, h.FinalLedgerUID,
		indent, h.FinalEndorseDigest,
		indent, h.DataDigest,
		indent, h.RatifiersDigest,
		indent, h.FollowingRatifiersDigest,
		indent, h.ApplicationDigest,
		indent, h.AgreementDigest,
		indent, h.FinalOutcomesDigest,
		indent, h.ProofDigest,
		indent, h.RecommenderLocation,
		indent, h.Digest(),
	)
}

//
func (h *Heading) ToSchema() *engineproto.Heading {
	if h == nil {
		return nil
	}

	return &engineproto.Heading{
		Release:            h.Release,
		LedgerUID:            h.LedgerUID,
		Level:             h.Level,
		Time:               h.Time,
		FinalLedgerUid:        h.FinalLedgerUID.ToSchema(),
		RatifiersDigest:     h.RatifiersDigest,
		FollowingRatifiersDigest: h.FollowingRatifiersDigest,
		AgreementDigest:      h.AgreementDigest,
		ApplicationDigest:            h.ApplicationDigest,
		DataDigest:           h.DataDigest,
		ProofDigest:       h.ProofDigest,
		FinalOutcomesDigest:    h.FinalOutcomesDigest,
		FinalEndorseDigest:     h.FinalEndorseDigest,
		RecommenderLocation:    h.RecommenderLocation,
	}
}

//
//
func HeadingFromSchema(ph *engineproto.Heading) (Heading, error) {
	if ph == nil {
		return Heading{}, errors.New("REDACTED")
	}

	h := new(Heading)

	bi, err := LedgerUIDFromSchema(&ph.FinalLedgerUid)
	if err != nil {
		return Heading{}, err
	}

	h.Release = ph.Release
	h.LedgerUID = ph.LedgerUID
	h.Level = ph.Level
	h.Time = ph.Time
	h.FinalLedgerUID = *bi
	h.RatifiersDigest = ph.RatifiersDigest
	h.FollowingRatifiersDigest = ph.FollowingRatifiersDigest
	h.AgreementDigest = ph.AgreementDigest
	h.ApplicationDigest = ph.ApplicationDigest
	h.DataDigest = ph.DataDigest
	h.ProofDigest = ph.ProofDigest
	h.FinalOutcomesDigest = ph.FinalOutcomesDigest
	h.FinalEndorseDigest = ph.FinalEndorseDigest
	h.RecommenderLocation = ph.RecommenderLocation

	return *h, h.CertifySimple()
}

//

//
type LedgerUIDMark byte

const (
	//
	LedgerUIDMarkMissing LedgerUIDMark = iota + 1
	//
	LedgerUIDMarkEndorse
	//
	LedgerUIDMarkNull
)

const (
	//
	MaximumEndorseBurdenOctets int64 = 94

	//
	//
	maximumEndorseSignatureSchemaEncodeBurden = 4 + 1 + 1 + 1 + 3 //
	//
	//
	//
	MaximumEndorseSignatureOctets = 131 + maximumEndorseSignatureSchemaEncodeBurden
)

//
type EndorseSignature struct {
	LedgerUIDMark      LedgerUIDMark `json:"ledger_uid_mark"`
	RatifierLocation Location     `json:"ratifier_location"`
	Timestamp        time.Time   `json:"timestamp"`
	Autograph        []byte      `json:"autograph"`
}

func MaximumEndorseOctets(valueNumber int) int64 {
	//
	const schemaIteratedFieldSizeBurden int64 = 3
	//
	return MaximumEndorseBurdenOctets + ((MaximumEndorseSignatureOctets + schemaIteratedFieldSizeBurden) * int64(valueNumber))
}

//
//
func NewEndorseSignatureMissing() EndorseSignature {
	return EndorseSignature{
		LedgerUIDMark: LedgerUIDMarkMissing,
	}
}

//
//
//
//
//
//
func (cs EndorseSignature) String() string {
	return fmt.Sprintf("REDACTED",
		cometbytes.Footprint(cs.Autograph),
		cometbytes.Footprint(cs.RatifierLocation),
		cs.LedgerUIDMark,
		StandardTime(cs.Timestamp))
}

//
//
func (cs EndorseSignature) LedgerUID(endorseLedgerUID LedgerUID) LedgerUID {
	var ledgerUID LedgerUID
	switch cs.LedgerUIDMark {
	case LedgerUIDMarkMissing:
		ledgerUID = LedgerUID{}
	case LedgerUIDMarkEndorse:
		ledgerUID = endorseLedgerUID
	case LedgerUIDMarkNull:
		ledgerUID = LedgerUID{}
	default:
		panic(fmt.Sprintf("REDACTED", cs.LedgerUIDMark))
	}
	return ledgerUID
}

//
func (cs EndorseSignature) CertifySimple() error {
	switch cs.LedgerUIDMark {
	case LedgerUIDMarkMissing:
	case LedgerUIDMarkEndorse:
	case LedgerUIDMarkNull:
	default:
		return fmt.Errorf("REDACTED", cs.LedgerUIDMark)
	}

	switch cs.LedgerUIDMark {
	case LedgerUIDMarkMissing:
		if len(cs.RatifierLocation) != 0 {
			return errors.New("REDACTED")
		}
		if !cs.Timestamp.IsZero() {
			return errors.New("REDACTED")
		}
		if len(cs.Autograph) != 0 {
			return errors.New("REDACTED")
		}
	default:
		if len(cs.RatifierLocation) != vault.LocationVolume {
			return fmt.Errorf("REDACTED",
				vault.LocationVolume,
				len(cs.RatifierLocation),
			)
		}
		//
		if len(cs.Autograph) == 0 {
			return errors.New("REDACTED")
		}
		if len(cs.Autograph) > MaximumAutographVolume {
			return fmt.Errorf("REDACTED", MaximumAutographVolume)
		}
	}

	return nil
}

//
func (cs *EndorseSignature) ToSchema() *engineproto.EndorseSignature {
	if cs == nil {
		return nil
	}

	return &engineproto.EndorseSignature{
		LedgerUidMark:      engineproto.LedgerUIDMark(cs.LedgerUIDMark),
		RatifierLocation: cs.RatifierLocation,
		Timestamp:        cs.Timestamp,
		Autograph:        cs.Autograph,
	}
}

//
//
func (cs *EndorseSignature) FromSchema(csp engineproto.EndorseSignature) error {
	cs.LedgerUIDMark = LedgerUIDMark(csp.LedgerUidMark)
	cs.RatifierLocation = csp.RatifierLocation
	cs.Timestamp = csp.Timestamp
	cs.Autograph = csp.Autograph

	return cs.CertifySimple()
}

//

//
//
type ExpandedEndorseSignature struct {
	EndorseSignature                 //
	Addition          []byte //
	AdditionAutograph []byte //
}

//
//
func NewExpandedEndorseSignatureMissing() ExpandedEndorseSignature {
	return ExpandedEndorseSignature{EndorseSignature: NewEndorseSignatureMissing()}
}

//
//
//
//
//
func (ecs ExpandedEndorseSignature) String() string {
	return fmt.Sprintf("REDACTED",
		ecs.EndorseSignature,
		cometbytes.Footprint(ecs.Addition),
		cometbytes.Footprint(ecs.AdditionAutograph),
	)
}

//
func (ecs ExpandedEndorseSignature) CertifySimple() error {
	if err := ecs.EndorseSignature.CertifySimple(); err != nil {
		return err
	}

	if ecs.LedgerUIDMark == LedgerUIDMarkEndorse {
		if len(ecs.Addition) > MaximumBallotAdditionVolume {
			return fmt.Errorf("REDACTED", MaximumBallotAdditionVolume)
		}
		if len(ecs.AdditionAutograph) > MaximumAutographVolume {
			return fmt.Errorf("REDACTED", MaximumAutographVolume)
		}
		return nil
	}

	if len(ecs.AdditionAutograph) == 0 && len(ecs.Addition) != 0 {
		return errors.New("REDACTED")
	}
	return nil
}

//
//
func (ecs ExpandedEndorseSignature) AssureAddition(extensionActivated bool) error {
	if extensionActivated {
		if ecs.LedgerUIDMark == LedgerUIDMarkEndorse && len(ecs.AdditionAutograph) == 0 {
			return fmt.Errorf("REDACTED",
				ecs.RatifierLocation.String(),
				ecs.Timestamp,
			)
		}
		if ecs.LedgerUIDMark != LedgerUIDMarkEndorse && len(ecs.Addition) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.RatifierLocation.String(),
				ecs.Timestamp,
			)
		}
		if ecs.LedgerUIDMark != LedgerUIDMarkEndorse && len(ecs.AdditionAutograph) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.RatifierLocation.String(),
				ecs.Timestamp,
			)
		}
	} else {
		if len(ecs.Addition) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.RatifierLocation.String(),
				ecs.Timestamp,
			)
		}
		if len(ecs.AdditionAutograph) != 0 {
			return fmt.Errorf("REDACTED",
				ecs.RatifierLocation.String(),
				ecs.Timestamp,
			)
		}
	}
	return nil
}

//
func (ecs *ExpandedEndorseSignature) ToSchema() *engineproto.ExpandedEndorseSignature {
	if ecs == nil {
		return nil
	}

	return &engineproto.ExpandedEndorseSignature{
		LedgerUidMark:        engineproto.LedgerUIDMark(ecs.LedgerUIDMark),
		RatifierLocation:   ecs.RatifierLocation,
		Timestamp:          ecs.Timestamp,
		Autograph:          ecs.Autograph,
		Addition:          ecs.Addition,
		AdditionAutograph: ecs.AdditionAutograph,
	}
}

//
//
//
func (ecs *ExpandedEndorseSignature) FromSchema(ecsp engineproto.ExpandedEndorseSignature) error {
	ecs.LedgerUIDMark = LedgerUIDMark(ecsp.LedgerUidMark)
	ecs.RatifierLocation = ecsp.RatifierLocation
	ecs.Timestamp = ecsp.Timestamp
	ecs.Autograph = ecsp.Autograph
	ecs.Addition = ecsp.Addition
	ecs.AdditionAutograph = ecsp.AdditionAutograph

	return ecs.CertifySimple()
}

//

//
//
type Endorse struct {
	//
	//
	//
	//
	Level     int64       `json:"level"`
	Cycle      int32       `json:"epoch"`
	LedgerUID    LedgerUID     `json:"ledger_uid"`
	Endorsements []EndorseSignature `json:"endorsements"`

	//
	//
	//
	digest cometbytes.HexOctets
}

//
func (endorse *Endorse) Replicate() *Endorse {
	autographs := make([]EndorseSignature, len(endorse.Endorsements))
	copy(autographs, endorse.Endorsements)
	xferClone := *endorse
	xferClone.Endorsements = autographs
	return &xferClone
}

//
//
//
//
//
func (endorse *Endorse) FetchBallot(valueIdx int32) *Ballot {
	endorseSignature := endorse.Endorsements[valueIdx]
	return &Ballot{
		Kind:             engineproto.PreendorseKind,
		Level:           endorse.Level,
		Cycle:            endorse.Cycle,
		LedgerUID:          endorseSignature.LedgerUID(endorse.LedgerUID),
		Timestamp:        endorseSignature.Timestamp,
		RatifierLocation: endorseSignature.RatifierLocation,
		RatifierOrdinal:   valueIdx,
		Autograph:        endorseSignature.Autograph,
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
func (endorse *Endorse) BallotAttestOctets(ledgerUID string, valueIdx int32) []byte {
	v := endorse.FetchBallot(valueIdx).ToSchema()
	return BallotAttestOctets(ledgerUID, v)
}

//
func (endorse *Endorse) Volume() int {
	if endorse == nil {
		return 0
	}
	return len(endorse.Endorsements)
}

//
//
func (endorse *Endorse) CertifySimple() error {
	if endorse.Level < 0 {
		return errors.New("REDACTED")
	}
	if endorse.Cycle < 0 {
		return errors.New("REDACTED")
	}

	if endorse.Level >= 1 {
		if endorse.LedgerUID.IsNil() {
			return errors.New("REDACTED")
		}

		if len(endorse.Endorsements) == 0 {
			return errors.New("REDACTED")
		}
		for i, endorseSignature := range endorse.Endorsements {
			if err := endorseSignature.CertifySimple(); err != nil {
				return fmt.Errorf("REDACTED", i, err)
			}
		}
	}
	return nil
}

//
func (endorse *Endorse) Digest() cometbytes.HexOctets {
	if endorse == nil {
		return nil
	}
	if endorse.digest == nil {
		bs := make([][]byte, len(endorse.Endorsements))
		for i, endorseSignature := range endorse.Endorsements {
			pbft := endorseSignature.ToSchema()
			bz, err := pbft.Serialize()
			if err != nil {
				panic(err)
			}

			bs[i] = bz
		}
		endorse.digest = merkle.DigestFromOctetSegments(bs)
	}
	return endorse.digest
}

//
//
//
//
//
func (endorse *Endorse) EncapsulatedExpandedEndorse() *ExpandedEndorse {
	cs := make([]ExpandedEndorseSignature, len(endorse.Endorsements))
	for idx, s := range endorse.Endorsements {
		cs[idx] = ExpandedEndorseSignature{
			EndorseSignature: s,
		}
	}
	return &ExpandedEndorse{
		Level:             endorse.Level,
		Cycle:              endorse.Cycle,
		LedgerUID:            endorse.LedgerUID,
		ExpandedEndorsements: cs,
	}
}

//
func (endorse *Endorse) StringIndented(indent string) string {
	if endorse == nil {
		return "REDACTED"
	}
	endorseSignatureStrings := make([]string, len(endorse.Endorsements))
	for i, endorseSignature := range endorse.Endorsements {
		endorseSignatureStrings[i] = endorseSignature.String()
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDd
REDACTEDd
REDACTEDv
REDACTED:
REDACTEDv
REDACTED`,
		indent, endorse.Level,
		indent, endorse.Cycle,
		indent, endorse.LedgerUID,
		indent,
		indent, strings.Join(endorseSignatureStrings, "REDACTED"+indent+"REDACTED"),
		indent, endorse.digest)
}

//
func (endorse *Endorse) ToSchema() *engineproto.Endorse {
	if endorse == nil {
		return nil
	}

	c := new(engineproto.Endorse)
	autographs := make([]engineproto.EndorseSignature, len(endorse.Endorsements))
	for i := range endorse.Endorsements {
		autographs[i] = *endorse.Endorsements[i].ToSchema()
	}
	c.Endorsements = autographs

	c.Level = endorse.Level
	c.Cycle = endorse.Cycle
	c.LedgerUID = endorse.LedgerUID.ToSchema()

	return c
}

//
//
func EndorseFromSchema(cp *engineproto.Endorse) (*Endorse, error) {
	if cp == nil {
		return nil, errors.New("REDACTED")
	}

	endorse := new(Endorse)

	bi, err := LedgerUIDFromSchema(&cp.LedgerUID)
	if err != nil {
		return nil, err
	}

	autographs := make([]EndorseSignature, len(cp.Endorsements))
	for i := range cp.Endorsements {
		if err := autographs[i].FromSchema(cp.Endorsements[i]); err != nil {
			return nil, err
		}
	}
	endorse.Endorsements = autographs

	endorse.Level = cp.Level
	endorse.Cycle = cp.Cycle
	endorse.LedgerUID = *bi

	return endorse, endorse.CertifySimple()
}

//

//
//
type ExpandedEndorse struct {
	Level             int64
	Cycle              int32
	LedgerUID            LedgerUID
	ExpandedEndorsements []ExpandedEndorseSignature

	bitList *bits.BitList
}

//
func (ec *ExpandedEndorse) Replicate() *ExpandedEndorse {
	autographs := make([]ExpandedEndorseSignature, len(ec.ExpandedEndorsements))
	copy(autographs, ec.ExpandedEndorsements)
	ecc := *ec
	ecc.ExpandedEndorsements = autographs
	return &ecc
}

//
//
//
//
func (ec *ExpandedEndorse) ToExpandedBallotCollection(ledgerUID string, values *RatifierAssign) *BallotCollection {
	ballotCollection := NewExpandedBallotCollection(ledgerUID, ec.Level, ec.Cycle, engineproto.PreendorseKind, values)
	ec.appendAutographsToBallotCollection(ballotCollection)
	return ballotCollection
}

//
func (ec *ExpandedEndorse) appendAutographsToBallotCollection(ballotCollection *BallotCollection) {
	for idx, ecs := range ec.ExpandedEndorsements {
		if ecs.LedgerUIDMark == LedgerUIDMarkMissing {
			continue //
		}
		ballot := ec.FetchExpandedBallot(int32(idx))
		if err := ballot.CertifySimple(); err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
		appended, err := ballotCollection.AppendBallot(ballot)
		if !appended || err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
	}
}

//
//
//
func (endorse *Endorse) ToBallotCollection(ledgerUID string, values *RatifierAssign) *BallotCollection {
	ballotCollection := NewBallotCollection(ledgerUID, endorse.Level, endorse.Cycle, engineproto.PreendorseKind, values)
	for idx, cs := range endorse.Endorsements {
		if cs.LedgerUIDMark == LedgerUIDMarkMissing {
			continue //
		}
		ballot := endorse.FetchBallot(int32(idx))
		if err := ballot.CertifySimple(); err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
		appended, err := ballotCollection.AppendBallot(ballot)
		if !appended || err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}
	}
	return ballotCollection
}

//
//
func (ec *ExpandedEndorse) AssurePlugins(extensionActivated bool) error {
	for _, ecs := range ec.ExpandedEndorsements {
		if err := ecs.AssureAddition(extensionActivated); err != nil {
			return err
		}
	}
	return nil
}

//
//
func (ec *ExpandedEndorse) ToEndorse() *Endorse {
	cs := make([]EndorseSignature, len(ec.ExpandedEndorsements))
	for idx, ecs := range ec.ExpandedEndorsements {
		cs[idx] = ecs.EndorseSignature
	}
	return &Endorse{
		Level:     ec.Level,
		Cycle:      ec.Cycle,
		LedgerUID:    ec.LedgerUID,
		Endorsements: cs,
	}
}

//
//
//
func (ec *ExpandedEndorse) FetchExpandedBallot(valueOrdinal int32) *Ballot {
	ecs := ec.ExpandedEndorsements[valueOrdinal]
	return &Ballot{
		Kind:               engineproto.PreendorseKind,
		Level:             ec.Level,
		Cycle:              ec.Cycle,
		LedgerUID:            ecs.LedgerUID(ec.LedgerUID),
		Timestamp:          ecs.Timestamp,
		RatifierLocation:   ecs.RatifierLocation,
		RatifierOrdinal:     valueOrdinal,
		Autograph:          ecs.Autograph,
		Addition:          ecs.Addition,
		AdditionAutograph: ecs.AdditionAutograph,
	}
}

//
//
//
func (ec *ExpandedEndorse) Kind() byte { return byte(engineproto.PreendorseKind) }

//
//
func (ec *ExpandedEndorse) FetchLevel() int64 { return ec.Level }

//
//
func (ec *ExpandedEndorse) FetchDuration() int32 { return ec.Cycle }

//
//
func (ec *ExpandedEndorse) Volume() int {
	if ec == nil {
		return 0
	}
	return len(ec.ExpandedEndorsements)
}

//
//
//
func (ec *ExpandedEndorse) BitList() *bits.BitList {
	if ec.bitList == nil {
		primaryBitFn := func(i int) bool {
			//
			//
			return ec.ExpandedEndorsements[i].LedgerUIDMark != LedgerUIDMarkMissing
		}
		ec.bitList = bits.NewBitListFromFn(len(ec.ExpandedEndorsements), primaryBitFn)
	}
	return ec.bitList
}

//
//
//
func (ec *ExpandedEndorse) FetchByOrdinal(valueIdx int32) *Ballot {
	return ec.FetchExpandedBallot(valueIdx)
}

//
//
func (ec *ExpandedEndorse) IsEndorse() bool {
	return len(ec.ExpandedEndorsements) != 0
}

//
//
func (ec *ExpandedEndorse) CertifySimple() error {
	if ec.Level < 0 {
		return errors.New("REDACTED")
	}
	if ec.Cycle < 0 {
		return errors.New("REDACTED")
	}

	if ec.Level >= 1 {
		if ec.LedgerUID.IsNil() {
			return errors.New("REDACTED")
		}

		if len(ec.ExpandedEndorsements) == 0 {
			return errors.New("REDACTED")
		}
		for i, extensionEndorseSignature := range ec.ExpandedEndorsements {
			if err := extensionEndorseSignature.CertifySimple(); err != nil {
				return fmt.Errorf("REDACTED", i, err)
			}
		}
	}
	return nil
}

//
func (ec *ExpandedEndorse) ToSchema() *engineproto.ExpandedEndorse {
	if ec == nil {
		return nil
	}

	c := new(engineproto.ExpandedEndorse)
	autographs := make([]engineproto.ExpandedEndorseSignature, len(ec.ExpandedEndorsements))
	for i := range ec.ExpandedEndorsements {
		autographs[i] = *ec.ExpandedEndorsements[i].ToSchema()
	}
	c.ExpandedEndorsements = autographs

	c.Level = ec.Level
	c.Cycle = ec.Cycle
	c.LedgerUID = ec.LedgerUID.ToSchema()

	return c
}

//
//
func ExpandedEndorseFromSchema(ecp *engineproto.ExpandedEndorse) (*ExpandedEndorse, error) {
	if ecp == nil {
		return nil, errors.New("REDACTED")
	}

	extensionEndorse := new(ExpandedEndorse)

	bi, err := LedgerUIDFromSchema(&ecp.LedgerUID)
	if err != nil {
		return nil, err
	}

	autographs := make([]ExpandedEndorseSignature, len(ecp.ExpandedEndorsements))
	for i := range ecp.ExpandedEndorsements {
		if err := autographs[i].FromSchema(ecp.ExpandedEndorsements[i]); err != nil {
			return nil, err
		}
	}
	extensionEndorse.ExpandedEndorsements = autographs
	extensionEndorse.Level = ecp.Level
	extensionEndorse.Cycle = ecp.Cycle
	extensionEndorse.LedgerUID = *bi

	return extensionEndorse, extensionEndorse.CertifySimple()
}

//

//
type Data struct {
	//
	//
	//
	Txs Txs `json:"txs"`

	//
	digest cometbytes.HexOctets
}

//
func (data *Data) Digest() cometbytes.HexOctets {
	if data == nil {
		return (Txs{}).Digest()
	}
	if data.digest == nil {
		data.digest = data.Txs.Digest() //
	}
	return data.digest
}

//
func (data *Data) StringIndented(indent string) string {
	if data == nil {
		return "REDACTED"
	}
	transferStrings := make([]string, cometmath.MinimumInteger(len(data.Txs), 21))
	for i, tx := range data.Txs {
		if i == 20 {
			transferStrings[i] = fmt.Sprintf("REDACTED", len(data.Txs))
			break
		}
		transferStrings[i] = fmt.Sprintf("REDACTED", tx.Digest(), len(tx))
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTED`,
		indent, strings.Join(transferStrings, "REDACTED"+indent+"REDACTED"),
		indent, data.digest)
}

//
func (data *Data) ToSchema() engineproto.Data {
	tp := new(engineproto.Data)

	if len(data.Txs) > 0 {
		transferBzs := make([][]byte, len(data.Txs))
		for i := range data.Txs {
			transferBzs[i] = data.Txs[i]
		}
		tp.Txs = transferBzs
	}

	return *tp
}

//
//
func DataFromSchema(dp *engineproto.Data) (Data, error) {
	if dp == nil {
		return Data{}, errors.New("REDACTED")
	}
	data := new(Data)

	if len(dp.Txs) > 0 {
		transferBzs := make(Txs, len(dp.Txs))
		for i := range dp.Txs {
			transferBzs[i] = Tx(dp.Txs[i])
		}
		data.Txs = transferBzs
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
	digest     cometbytes.HexOctets
	octetVolume int64
}

//
func (data *ProofData) Digest() cometbytes.HexOctets {
	if data.digest == nil {
		data.digest = data.Proof.Digest()
	}
	return data.digest
}

//
func (data *ProofData) OctetVolume() int64 {
	if data.octetVolume == 0 && len(data.Proof) != 0 {
		pb, err := data.ToSchema()
		if err != nil {
			panic(err)
		}
		data.octetVolume = int64(pb.Volume())
	}
	return data.octetVolume
}

//
func (data *ProofData) StringIndented(indent string) string {
	if data == nil {
		return "REDACTED"
	}
	evtStrings := make([]string, cometmath.MinimumInteger(len(data.Proof), 21))
	for i, ev := range data.Proof {
		if i == 20 {
			evtStrings[i] = fmt.Sprintf("REDACTED", len(data.Proof))
			break
		}
		evtStrings[i] = "REDACTED" + ev.String()
	}
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTED`,
		indent, strings.Join(evtStrings, "REDACTED"+indent+"REDACTED"),
		indent, data.digest)
}

//
func (data *ProofData) ToSchema() (*engineproto.ProofCatalog, error) {
	if data == nil {
		return nil, errors.New("REDACTED")
	}

	evi := new(engineproto.ProofCatalog)
	eviBzs := make([]engineproto.Proof, len(data.Proof))
	for i := range data.Proof {
		schemaEvi, err := ProofToSchema(data.Proof[i])
		if err != nil {
			return nil, err
		}
		eviBzs[i] = *schemaEvi
	}
	evi.Proof = eviBzs

	return evi, nil
}

//
func (data *ProofData) FromSchema(eviData *engineproto.ProofCatalog) error {
	if eviData == nil {
		return errors.New("REDACTED")
	}

	eviBzs := make(ProofCatalog, len(eviData.Proof))
	for i := range eviData.Proof {
		evi, err := ProofFromSchema(&eviData.Proof[i])
		if err != nil {
			return err
		}
		eviBzs[i] = evi
	}
	data.Proof = eviBzs
	data.octetVolume = int64(eviData.Volume())

	return nil
}

//

//
type LedgerUID struct {
	Digest          cometbytes.HexOctets `json:"digest"`
	SegmentAssignHeading SegmentAssignHeading     `json:"segments"`
}

//
func (ledgerUID LedgerUID) Matches(another LedgerUID) bool {
	return bytes.Equal(ledgerUID.Digest, another.Digest) &&
		ledgerUID.SegmentAssignHeading.Matches(another.SegmentAssignHeading)
}

//
func (ledgerUID LedgerUID) Key() string {
	pbph := ledgerUID.SegmentAssignHeading.ToSchema()
	bz, err := pbph.Serialize()
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(string(ledgerUID.Digest), string(bz))
}

//
func (ledgerUID LedgerUID) CertifySimple() error {
	//
	if err := CertifyDigest(ledgerUID.Digest); err != nil {
		return fmt.Errorf("REDACTED")
	}
	if err := ledgerUID.SegmentAssignHeading.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (ledgerUID LedgerUID) IsNil() bool {
	return len(ledgerUID.Digest) == 0 &&
		ledgerUID.SegmentAssignHeading.IsNil()
}

//
func (ledgerUID LedgerUID) IsFinished() bool {
	return len(ledgerUID.Digest) == comethash.Volume &&
		ledgerUID.SegmentAssignHeading.Sum > 0 &&
		len(ledgerUID.SegmentAssignHeading.Digest) == comethash.Volume
}

//
//
//
//
//
//
func (ledgerUID LedgerUID) String() string {
	return fmt.Sprintf("REDACTED", ledgerUID.Digest, ledgerUID.SegmentAssignHeading)
}

//
func (ledgerUID *LedgerUID) ToSchema() engineproto.LedgerUID {
	if ledgerUID == nil {
		return engineproto.LedgerUID{}
	}

	return engineproto.LedgerUID{
		Digest:          ledgerUID.Digest,
		SegmentAssignHeading: ledgerUID.SegmentAssignHeading.ToSchema(),
	}
}

//
//
func LedgerUIDFromSchema(bID *engineproto.LedgerUID) (*LedgerUID, error) {
	if bID == nil {
		return nil, errors.New("REDACTED")
	}

	ledgerUID := new(LedgerUID)
	ph, err := SegmentAssignHeadingFromSchema(&bID.SegmentAssignHeading)
	if err != nil {
		return nil, err
	}

	ledgerUID.SegmentAssignHeading = *ph
	ledgerUID.Digest = bID.Digest

	return ledgerUID, ledgerUID.CertifySimple()
}

//
//
func SchemaLedgerUIDIsNull(bID *engineproto.LedgerUID) bool {
	return len(bID.Digest) == 0 && SchemaSectionCollectionHeadingIsNil(&bID.SegmentAssignHeading)
}
