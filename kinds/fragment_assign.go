package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

var (
	FaultFragmentAssignUnforeseenOrdinal = errors.New("REDACTED")
	FaultFragmentAssignUnfitAttestation    = errors.New("REDACTED")
	FaultFragmentExcessivelyAmple             = errors.New("REDACTED")
	FaultFragmentUnfitExtent        = errors.New("REDACTED")
)

//
type FaultUnfitFragment struct {
	Rationale error
}

func (e FaultUnfitFragment) Failure() string {
	return fmt.Sprintf("REDACTED", e.Rationale)
}

func (e FaultUnfitFragment) Disclose() error {
	return e.Rationale
}

type Fragment struct {
	Ordinal uint32            `json:"ordinal"`
	Octets tendermintoctets.HexadecimalOctets `json:"octets"`
	Attestation hashmap.Attestation      `json:"attestation"`
}

//
func (fragment *Fragment) CertifyFundamental() error {
	if len(fragment.Octets) > int(LedgerFragmentExtentOctets) {
		return FaultFragmentExcessivelyAmple
	}
	//
	if int64(fragment.Ordinal) < fragment.Attestation.Sum-1 && len(fragment.Octets) != int(LedgerFragmentExtentOctets) {
		return FaultFragmentUnfitExtent
	}
	if int64(fragment.Ordinal) != fragment.Attestation.Ordinal {
		return FaultUnfitFragment{Rationale: fmt.Errorf("REDACTED", fragment.Ordinal, fragment.Attestation.Ordinal)}
	}
	if err := fragment.Attestation.CertifyFundamental(); err != nil {
		return FaultUnfitFragment{Rationale: fmt.Errorf("REDACTED", err)}
	}
	return nil
}

//
//
//
func (fragment *Fragment) Text() string {
	return fragment.TextFormatted("REDACTED")
}

//
//
//
func (fragment *Fragment) TextFormatted(format string) string {
	return fmt.Sprintf(`REDACTEDv
REDACTED.
REDACTEDv
REDACTED`,
		fragment.Ordinal,
		format, tendermintoctets.Identifier(fragment.Octets),
		format, fragment.Attestation.TextFormatted(format+"REDACTED"),
		format)
}

func (fragment *Fragment) TowardSchema() (*commitchema.Fragment, error) {
	if fragment == nil {
		return nil, errors.New("REDACTED")
	}
	pb := new(commitchema.Fragment)
	attestation := fragment.Attestation.TowardSchema()

	pb.Ordinal = fragment.Ordinal
	pb.Octets = fragment.Octets
	pb.Attestation = *attestation

	return pb, nil
}

func FragmentOriginatingSchema(pb *commitchema.Fragment) (*Fragment, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	fragment := new(Fragment)
	attestation, err := hashmap.AttestationOriginatingSchema(&pb.Attestation)
	if err != nil {
		return nil, err
	}
	fragment.Ordinal = pb.Ordinal
	fragment.Octets = pb.Octets
	fragment.Attestation = *attestation

	return fragment, fragment.CertifyFundamental()
}

//

type FragmentAssignHeading struct {
	Sum uint32            `json:"sum"`
	Digest  tendermintoctets.HexadecimalOctets `json:"digest"`
}

//
//
//
//
func (psh FragmentAssignHeading) Text() string {
	return fmt.Sprintf("REDACTED", psh.Sum, tendermintoctets.Identifier(psh.Digest))
}

func (psh FragmentAssignHeading) EqualsNull() bool {
	return psh.Sum == 0 && len(psh.Digest) == 0
}

func (psh FragmentAssignHeading) Matches(another FragmentAssignHeading) bool {
	return psh.Sum == another.Sum && bytes.Equal(psh.Digest, another.Digest)
}

//
func (psh FragmentAssignHeading) CertifyFundamental() error {
	//
	if err := CertifyDigest(psh.Digest); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
func (psh *FragmentAssignHeading) TowardSchema() commitchema.FragmentAssignHeading {
	if psh == nil {
		return commitchema.FragmentAssignHeading{}
	}

	return commitchema.FragmentAssignHeading{
		Sum: psh.Sum,
		Digest:  psh.Digest,
	}
}

//
func FragmentAssignHeadingOriginatingSchema(props *commitchema.FragmentAssignHeading) (*FragmentAssignHeading, error) {
	if props == nil {
		return nil, errors.New("REDACTED")
	}
	psh := new(FragmentAssignHeading)
	psh.Sum = props.Sum
	psh.Digest = props.Digest

	return psh, psh.CertifyFundamental()
}

//
//
func SchemaFragmentAssignHeadingEqualsNull(props *commitchema.FragmentAssignHeading) bool {
	return props.Sum == 0 && len(props.Digest) == 0
}

//

type FragmentAssign struct {
	sum uint32
	digest  []byte

	mtx           commitchronize.Exclusion
	fragments         []*Fragment
	fragmentsDigitSeries *digits.DigitSeries
	tally         uint32
	//
	//
	octetExtent int64
}

//
//
//
func FreshFragmentAssignOriginatingData(data []byte, fragmentExtent uint32) *FragmentAssign {
	//
	sum := (uint32(len(data)) + fragmentExtent - 1) / fragmentExtent
	fragments := make([]*Fragment, sum)
	fragmentsOctets := make([][]byte, sum)
	for i := uint32(0); i < sum; i++ {
		fragment := &Fragment{
			Ordinal: i,
			Octets: data[i*fragmentExtent : strongarithmetic.MinimumInteger(len(data), int((i+1)*fragmentExtent))],
		}
		fragments[i] = fragment
		fragmentsOctets[i] = fragment.Octets
	}
	//
	origin, attestations := hashmap.AttestationsOriginatingOctetSegments(fragmentsOctets)
	for i := uint32(0); i < sum; i++ {
		fragments[i].Attestation = *attestations[i]
	}
	fragmentsDigitSeries := digits.FreshDigitSeriesOriginatingProc(int(sum), func(int) bool { return true })
	return &FragmentAssign{
		sum:         sum,
		digest:          origin,
		fragments:         fragments,
		fragmentsDigitSeries: fragmentsDigitSeries,
		tally:         sum,
		octetExtent:      int64(len(data)),
	}
}

//
func FreshFragmentAssignOriginatingHeading(heading FragmentAssignHeading) *FragmentAssign {
	return &FragmentAssign{
		sum:         heading.Sum,
		digest:          heading.Digest,
		fragments:         make([]*Fragment, heading.Sum),
		fragmentsDigitSeries: digits.FreshDigitCollection(int(heading.Sum)),
		tally:         0,
		octetExtent:      0,
	}
}

func (ps *FragmentAssign) Heading() FragmentAssignHeading {
	if ps == nil {
		return FragmentAssignHeading{}
	}
	return FragmentAssignHeading{
		Sum: ps.sum,
		Digest:  ps.digest,
	}
}

func (ps *FragmentAssign) OwnsHeading(heading FragmentAssignHeading) bool {
	if ps == nil {
		return false
	}
	return ps.Heading().Matches(heading)
}

func (ps *FragmentAssign) DigitSeries() *digits.DigitSeries {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.fragmentsDigitSeries.Duplicate()
}

func (ps *FragmentAssign) Digest() []byte {
	if ps == nil {
		return hashmap.DigestOriginatingOctetSegments(nil)
	}
	return ps.digest
}

func (ps *FragmentAssign) DigestsToward(digest []byte) bool {
	if ps == nil {
		return false
	}
	return bytes.Equal(ps.digest, digest)
}

func (ps *FragmentAssign) Tally() uint32 {
	if ps == nil {
		return 0
	}
	return ps.tally
}

func (ps *FragmentAssign) OctetExtent() int64 {
	if ps == nil {
		return 0
	}
	return ps.octetExtent
}

func (ps *FragmentAssign) Sum() uint32 {
	if ps == nil {
		return 0
	}
	return ps.sum
}

//
func (ps *FragmentAssign) AppendFragment(fragment *Fragment) (bool, error) {
	//
	//
	if ps == nil {
		return false, nil
	}

	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	//
	if fragment.Ordinal >= ps.sum {
		return false, FaultFragmentAssignUnforeseenOrdinal
	}

	//
	if ps.fragments[fragment.Ordinal] != nil {
		return false, nil
	}

	//
	if fragment.Attestation.Sum != int64(ps.sum) {
		return false, FaultFragmentAssignUnfitAttestation
	}

	//
	if fragment.Attestation.Validate(ps.Digest(), fragment.Octets) != nil {
		return false, FaultFragmentAssignUnfitAttestation
	}

	//
	ps.fragments[fragment.Ordinal] = fragment
	ps.fragmentsDigitSeries.AssignOrdinal(int(fragment.Ordinal), true)
	ps.tally++
	ps.octetExtent += int64(len(fragment.Octets))
	return true, nil
}

func (ps *FragmentAssign) ObtainFragment(ordinal int) *Fragment {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.fragments[ordinal]
}

func (ps *FragmentAssign) EqualsFinish() bool {
	return ps.tally == ps.sum
}

func (ps *FragmentAssign) ObtainFetcher() io.Reader {
	if !ps.EqualsFinish() {
		panic("REDACTED")
	}
	return FreshFragmentAssignFetcher(ps.fragments)
}

type FragmentAssignFetcher struct {
	i      int
	fragments  []*Fragment
	fetcher *bytes.Reader
}

func FreshFragmentAssignFetcher(fragments []*Fragment) *FragmentAssignFetcher {
	return &FragmentAssignFetcher{
		i:      0,
		fragments:  fragments,
		fetcher: bytes.NewReader(fragments[0].Octets),
	}
}

func (psr *FragmentAssignFetcher) Obtain(p []byte) (n int, err error) {
	fetcherLength := psr.fetcher.Len()
	if fetcherLength >= len(p) {
		return psr.fetcher.Read(p)
	} else if fetcherLength > 0 {
		n1, err := psr.Obtain(p[:fetcherLength])
		if err != nil {
			return n1, err
		}
		n2, err := psr.Obtain(p[fetcherLength:])
		return n1 + n2, err
	}

	psr.i++
	if psr.i >= len(psr.fragments) {
		return 0, io.EOF
	}
	psr.fetcher = bytes.NewReader(psr.fragments[psr.i].Octets)
	return psr.Obtain(p)
}

//
//
//
func (ps *FragmentAssign) TextBrief() string {
	if ps == nil {
		return "REDACTED"
	}
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return fmt.Sprintf("REDACTED", ps.Tally(), ps.Sum())
}

func (ps *FragmentAssign) SerializeJSN() ([]byte, error) {
	if ps == nil {
		return []byte("REDACTED"), nil
	}

	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	return strongmindjson.Serialize(struct {
		TallySum    string         `json:"tally/sum"`
		FragmentsDigitSeries *digits.DigitSeries `json:"fragments_digit_series"`
	}{
		fmt.Sprintf("REDACTED", ps.Tally(), ps.Sum()),
		ps.fragmentsDigitSeries,
	})
}
