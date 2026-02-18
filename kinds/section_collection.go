package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/utils/bits"
	cometbytes "github.com/valkyrieworks/utils/octets"
	cometjson "github.com/valkyrieworks/utils/json"
	cometmath "github.com/valkyrieworks/utils/math"
	engineconnect "github.com/valkyrieworks/utils/align"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

var (
	ErrSegmentCollectionUnforeseenOrdinal = errors.New("REDACTED")
	ErrSegmentCollectionCorruptAttestation    = errors.New("REDACTED")
	ErrSectionTooLarge             = errors.New("REDACTED")
	ErrSectionCorruptVolume        = errors.New("REDACTED")
)

//
type ErrCorruptSection struct {
	Cause error
}

func (e ErrCorruptSection) Fault() string {
	return fmt.Sprintf("REDACTED", e.Cause)
}

func (e ErrCorruptSection) Disclose() error {
	return e.Cause
}

type Segment struct {
	Ordinal uint32            `json:"ordinal"`
	Octets cometbytes.HexOctets `json:"octets"`
	Attestation merkle.Attestation      `json:"evidence"`
}

//
func (segment *Segment) CertifySimple() error {
	if len(segment.Octets) > int(LedgerSegmentVolumeOctets) {
		return ErrSectionTooLarge
	}
	//
	if int64(segment.Ordinal) < segment.Attestation.Sum-1 && len(segment.Octets) != int(LedgerSegmentVolumeOctets) {
		return ErrSectionCorruptVolume
	}
	if int64(segment.Ordinal) != segment.Attestation.Ordinal {
		return ErrCorruptSection{Cause: fmt.Errorf("REDACTED", segment.Ordinal, segment.Attestation.Ordinal)}
	}
	if err := segment.Attestation.CertifySimple(); err != nil {
		return ErrCorruptSection{Cause: fmt.Errorf("REDACTED", err)}
	}
	return nil
}

//
//
//
func (segment *Segment) String() string {
	return segment.StringIndented("REDACTED")
}

//
//
//
func (segment *Segment) StringIndented(indent string) string {
	return fmt.Sprintf(`REDACTEDv
REDACTED.
REDACTEDv
REDACTED`,
		segment.Ordinal,
		indent, cometbytes.Footprint(segment.Octets),
		indent, segment.Attestation.StringIndented(indent+"REDACTED"),
		indent)
}

func (segment *Segment) ToSchema() (*engineproto.Segment, error) {
	if segment == nil {
		return nil, errors.New("REDACTED")
	}
	pb := new(engineproto.Segment)
	evidence := segment.Attestation.ToSchema()

	pb.Ordinal = segment.Ordinal
	pb.Octets = segment.Octets
	pb.Attestation = *evidence

	return pb, nil
}

func SegmentFromSchema(pb *engineproto.Segment) (*Segment, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	segment := new(Segment)
	evidence, err := merkle.EvidenceFromSchema(&pb.Attestation)
	if err != nil {
		return nil, err
	}
	segment.Ordinal = pb.Ordinal
	segment.Octets = pb.Octets
	segment.Attestation = *evidence

	return segment, segment.CertifySimple()
}

//

type SegmentAssignHeading struct {
	Sum uint32            `json:"sum"`
	Digest  cometbytes.HexOctets `json:"digest"`
}

//
//
//
//
func (psh SegmentAssignHeading) String() string {
	return fmt.Sprintf("REDACTED", psh.Sum, cometbytes.Footprint(psh.Digest))
}

func (psh SegmentAssignHeading) IsNil() bool {
	return psh.Sum == 0 && len(psh.Digest) == 0
}

func (psh SegmentAssignHeading) Matches(another SegmentAssignHeading) bool {
	return psh.Sum == another.Sum && bytes.Equal(psh.Digest, another.Digest)
}

//
func (psh SegmentAssignHeading) CertifySimple() error {
	//
	if err := CertifyDigest(psh.Digest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (psh *SegmentAssignHeading) ToSchema() engineproto.SegmentAssignHeading {
	if psh == nil {
		return engineproto.SegmentAssignHeading{}
	}

	return engineproto.SegmentAssignHeading{
		Sum: psh.Sum,
		Digest:  psh.Digest,
	}
}

//
func SegmentAssignHeadingFromSchema(rush *engineproto.SegmentAssignHeading) (*SegmentAssignHeading, error) {
	if rush == nil {
		return nil, errors.New("REDACTED")
	}
	psh := new(SegmentAssignHeading)
	psh.Sum = rush.Sum
	psh.Digest = rush.Digest

	return psh, psh.CertifySimple()
}

//
//
func SchemaSectionCollectionHeadingIsNil(rush *engineproto.SegmentAssignHeading) bool {
	return rush.Sum == 0 && len(rush.Digest) == 0
}

//

type SegmentCollection struct {
	sum uint32
	digest  []byte

	mtx           engineconnect.Lock
	segments         []*Segment
	sectionsBitList *bits.BitList
	tally         uint32
	//
	//
	octetVolume int64
}

//
//
//
func NewSegmentCollectionFromData(data []byte, segmentVolume uint32) *SegmentCollection {
	//
	sum := (uint32(len(data)) + segmentVolume - 1) / segmentVolume
	segments := make([]*Segment, sum)
	sectionsOctets := make([][]byte, sum)
	for i := uint32(0); i < sum; i++ {
		segment := &Segment{
			Ordinal: i,
			Octets: data[i*segmentVolume : cometmath.MinimumInteger(len(data), int((i+1)*segmentVolume))],
		}
		segments[i] = segment
		sectionsOctets[i] = segment.Octets
	}
	//
	origin, evidences := merkle.EvidencesFromOctetSegments(sectionsOctets)
	for i := uint32(0); i < sum; i++ {
		segments[i].Attestation = *evidences[i]
	}
	sectionsBitList := bits.NewBitListFromFn(int(sum), func(int) bool { return true })
	return &SegmentCollection{
		sum:         sum,
		digest:          origin,
		segments:         segments,
		sectionsBitList: sectionsBitList,
		tally:         sum,
		octetVolume:      int64(len(data)),
	}
}

//
func NewSegmentCollectionFromHeading(heading SegmentAssignHeading) *SegmentCollection {
	return &SegmentCollection{
		sum:         heading.Sum,
		digest:          heading.Digest,
		segments:         make([]*Segment, heading.Sum),
		sectionsBitList: bits.NewBitList(int(heading.Sum)),
		tally:         0,
		octetVolume:      0,
	}
}

func (ps *SegmentCollection) Heading() SegmentAssignHeading {
	if ps == nil {
		return SegmentAssignHeading{}
	}
	return SegmentAssignHeading{
		Sum: ps.sum,
		Digest:  ps.digest,
	}
}

func (ps *SegmentCollection) HasHeading(heading SegmentAssignHeading) bool {
	if ps == nil {
		return false
	}
	return ps.Heading().Matches(heading)
}

func (ps *SegmentCollection) BitList() *bits.BitList {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.sectionsBitList.Clone()
}

func (ps *SegmentCollection) Digest() []byte {
	if ps == nil {
		return merkle.DigestFromOctetSegments(nil)
	}
	return ps.digest
}

func (ps *SegmentCollection) DigestsTo(digest []byte) bool {
	if ps == nil {
		return false
	}
	return bytes.Equal(ps.digest, digest)
}

func (ps *SegmentCollection) Number() uint32 {
	if ps == nil {
		return 0
	}
	return ps.tally
}

func (ps *SegmentCollection) OctetVolume() int64 {
	if ps == nil {
		return 0
	}
	return ps.octetVolume
}

func (ps *SegmentCollection) Sum() uint32 {
	if ps == nil {
		return 0
	}
	return ps.sum
}

//
func (ps *SegmentCollection) AppendSegment(segment *Segment) (bool, error) {
	//
	//
	if ps == nil {
		return false, nil
	}

	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	//
	if segment.Ordinal >= ps.sum {
		return false, ErrSegmentCollectionUnforeseenOrdinal
	}

	//
	if ps.segments[segment.Ordinal] != nil {
		return false, nil
	}

	//
	if segment.Attestation.Sum != int64(ps.sum) {
		return false, ErrSegmentCollectionCorruptAttestation
	}

	//
	if segment.Attestation.Validate(ps.Digest(), segment.Octets) != nil {
		return false, ErrSegmentCollectionCorruptAttestation
	}

	//
	ps.segments[segment.Ordinal] = segment
	ps.sectionsBitList.AssignOrdinal(int(segment.Ordinal), true)
	ps.tally++
	ps.octetVolume += int64(len(segment.Octets))
	return true, nil
}

func (ps *SegmentCollection) FetchSegment(ordinal int) *Segment {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.segments[ordinal]
}

func (ps *SegmentCollection) IsFinished() bool {
	return ps.tally == ps.sum
}

func (ps *SegmentCollection) FetchScanner() io.Reader {
	if !ps.IsFinished() {
		panic("REDACTED")
	}
	return NewSectionCollectionScanner(ps.segments)
}

type SectionCollectionScanner struct {
	i      int
	segments  []*Segment
	scanner *bytes.Reader
}

func NewSectionCollectionScanner(segments []*Segment) *SectionCollectionScanner {
	return &SectionCollectionScanner{
		i:      0,
		segments:  segments,
		scanner: bytes.NewReader(segments[0].Octets),
	}
}

func (psr *SectionCollectionScanner) Scan(p []byte) (n int, err error) {
	scannerSize := psr.scanner.Len()
	if scannerSize >= len(p) {
		return psr.scanner.Read(p)
	} else if scannerSize > 0 {
		n1, err := psr.Scan(p[:scannerSize])
		if err != nil {
			return n1, err
		}
		n2, err := psr.Scan(p[scannerSize:])
		return n1 + n2, err
	}

	psr.i++
	if psr.i >= len(psr.segments) {
		return 0, io.EOF
	}
	psr.scanner = bytes.NewReader(psr.segments[psr.i].Octets)
	return psr.Scan(p)
}

//
//
//
func (ps *SegmentCollection) StringBrief() string {
	if ps == nil {
		return "REDACTED"
	}
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return fmt.Sprintf("REDACTED", ps.Number(), ps.Sum())
}

func (ps *SegmentCollection) SerializeJSON() ([]byte, error) {
	if ps == nil {
		return []byte("REDACTED"), nil
	}

	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	return cometjson.Serialize(struct {
		NumberSum    string         `json:"number/sum"`
		SectionsBitList *bits.BitList `json:"sections_bit_list"`
	}{
		fmt.Sprintf("REDACTED", ps.Number(), ps.Sum()),
		ps.sectionsBitList,
	})
}
