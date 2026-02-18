package bits

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"regexp"
	"strings"
	"sync"

	cometmath "github.com/valkyrieworks/utils/math"
	engineseed "github.com/valkyrieworks/utils/random"
	cmtprotobits "github.com/valkyrieworks/schema/consensuscore/utils/bits"
)

//
type BitList struct {
	mtx   sync.Mutex
	Bits  int      `json:"bits"`  //
	Elements []uint64 `json:"elements"` //
}

//
//
func NewBitList(bits int) *BitList {
	if bits <= 0 {
		return nil
	}
	return &BitList{
		Bits:  bits,
		Elements: make([]uint64, countMembers(bits)),
	}
}

//
//
//
func NewBitListFromFn(bits int, fn func(int) bool) *BitList {
	if bits <= 0 {
		return nil
	}
	bA := &BitList{
		Bits:  bits,
		Elements: make([]uint64, countMembers(bits)),
	}
	for i := 0; i < bits; i++ {
		v := fn(i)
		if v {
			bA.Elements[i/64] |= (uint64(1) << uint(i%64))
		}
	}
	return bA
}

//
func (bA *BitList) Volume() int {
	if bA == nil {
		return 0
	}
	return bA.Bits
}

//
//
func (bA *BitList) FetchOrdinal(i int) bool {
	if bA == nil {
		return false
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.fetchOrdinal(i)
}

func (bA *BitList) fetchOrdinal(i int) bool {
	if i >= bA.Bits {
		return false
	}
	return bA.Elements[i/64]&(uint64(1)<<uint(i%64)) > 0
}

//
//
func (bA *BitList) AssignOrdinal(i int, v bool) bool {
	if bA == nil {
		return false
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.collectionOrdinal(i, v)
}

func (bA *BitList) collectionOrdinal(i int, v bool) bool {
	if i >= bA.Bits || i/64 >= len(bA.Elements) {
		return false
	}
	if v {
		bA.Elements[i/64] |= (uint64(1) << uint(i%64))
	} else {
		bA.Elements[i/64] &= ^(uint64(1) << uint(i%64))
	}
	return true
}

//
func (bA *BitList) Clone() *BitList {
	if bA == nil {
		return nil
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.clone()
}

func (bA *BitList) clone() *BitList {
	c := make([]uint64, len(bA.Elements))
	copy(c, bA.Elements)
	return &BitList{
		Bits:  bA.Bits,
		Elements: c,
	}
}

func (bA *BitList) cloneBits(bits int) *BitList {
	c := make([]uint64, countMembers(bits))
	copy(c, bA.Elements)
	return &BitList{
		Bits:  bits,
		Elements: c,
	}
}

//
//
//
func (bA *BitList) Or(o *BitList) *BitList {
	if bA == nil && o == nil {
		return nil
	}
	if bA == nil && o != nil {
		return o.Clone()
	}
	if o == nil {
		return bA.Clone()
	}
	bA.mtx.Lock()
	o.mtx.Lock()
	c := bA.cloneBits(cometmath.MaximumInteger(bA.Bits, o.Bits))
	less := cometmath.MinimumInteger(len(bA.Elements), len(o.Elements))
	for i := 0; i < less; i++ {
		c.Elements[i] |= o.Elements[i]
	}
	bA.mtx.Unlock()
	o.mtx.Unlock()
	return c
}

//
//
//
func (bA *BitList) And(o *BitList) *BitList {
	if bA == nil || o == nil {
		return nil
	}
	bA.mtx.Lock()
	o.mtx.Lock()
	defer func() {
		bA.mtx.Unlock()
		o.mtx.Unlock()
	}()
	return bA.and(o)
}

func (bA *BitList) and(o *BitList) *BitList {
	c := bA.cloneBits(cometmath.MinimumInteger(bA.Bits, o.Bits))
	for i := 0; i < len(c.Elements); i++ {
		c.Elements[i] &= o.Elements[i]
	}
	return c
}

//
func (bA *BitList) Not() *BitList {
	if bA == nil {
		return nil //
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.not()
}

func (bA *BitList) not() *BitList {
	c := bA.clone()
	for i := 0; i < len(c.Elements); i++ {
		c.Elements[i] = ^c.Elements[i]
	}
	return c
}

//
//
//
//
func (bA *BitList) Sub(o *BitList) *BitList {
	if bA == nil || o == nil {
		//
		return nil
	}
	bA.mtx.Lock()
	o.mtx.Lock()
	//
	c := bA.cloneBits(bA.Bits)
	//
	//
	//
	//
	less := cometmath.MinimumInteger(len(bA.Elements), len(o.Elements))
	for i := 0; i < less; i++ {
		//
		c.Elements[i] &^= o.Elements[i]
	}
	bA.mtx.Unlock()
	o.mtx.Unlock()
	return c
}

//
func (bA *BitList) IsEmpty() bool {
	if bA == nil {
		return true //
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	for _, e := range bA.Elements {
		if e > 0 {
			return false
		}
	}
	return true
}

//
func (bA *BitList) IsComplete() bool {
	if bA == nil {
		return true
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()

	//
	for _, element := range bA.Elements[:len(bA.Elements)-1] {
		if (^element) != 0 {
			return false
		}
	}

	//
	finalElementBits := (bA.Bits+63)%64 + 1
	finalElement := bA.Elements[len(bA.Elements)-1]
	return (finalElement+1)&((uint64(1)<<uint(finalElementBits))-1) == 0
}

//
//
//
func (bA *BitList) SelectArbitrary() (int, bool) {
	if bA == nil {
		return 0, false
	}

	bA.mtx.Lock()
	countTrueOrdinals := bA.fetchCountTrueOrdinals()
	if countTrueOrdinals == 0 { //
		bA.mtx.Unlock()
		return 0, false
	}
	ordinal := bA.fetchNthTrueOrdinal(engineseed.Intn(countTrueOrdinals))
	bA.mtx.Unlock()
	if ordinal == -1 {
		return 0, false
	}
	return ordinal, true
}

func (bA *BitList) fetchCountTrueOrdinals() int {
	if bA.Volume() == 0 || len(bA.Elements) == 0 || len(bA.Elements) != countMembers(bA.Volume()) {
		//
		return 0
	}

	tally := 0
	countElements := len(bA.Elements)
	//
	for i := 0; i < countElements-1; i++ {
		tally += bits.OnesCount64(bA.Elements[i])
	}
	//
	countUltimateBits := bA.Bits - (countElements-1)*64
	for i := 0; i < countUltimateBits; i++ {
		if (bA.Elements[countElements-1] & (uint64(1) << uint64(i))) > 0 {
			tally++
		}
	}
	return tally
}

//
//
//
func (bA *BitList) fetchNthTrueOrdinal(n int) int {
	countElements := len(bA.Elements)
	tally := 0

	//
	for i := 0; i < countElements; i++ {
		//
		collectionBits := bits.OnesCount64(bA.Elements[i])

		//
		//
		if tally+collectionBits >= n {
			//
			for j := 0; j < 64; j++ {
				if bA.Elements[i]&(1<<uint(j)) != 0 {
					if tally == n {
						//
						return i*64 + j
					}
					tally++
				}
			}
		} else {
			//
			tally += collectionBits
		}
	}

	//
	return -1
}

//
//
//
//
//
//
func (bA *BitList) String() string {
	return bA.StringIndented("REDACTED")
}

//
//
func (bA *BitList) StringIndented(indent string) string {
	if bA == nil {
		return "REDACTED"
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.stringIndented(indent)
}

func (bA *BitList) stringIndented(indent string) string {
	rows := []string{}
	bits := "REDACTED"
	for i := 0; i < bA.Bits; i++ {
		if bA.fetchOrdinal(i) {
			bits += "REDACTED"
		} else {
			bits += "REDACTED"
		}
		if i%100 == 99 {
			rows = append(rows, bits)
			bits = "REDACTED"
		}
		if i%10 == 9 {
			bits += indent
		}
		if i%50 == 49 {
			bits += indent
		}
	}
	if len(bits) > 0 {
		rows = append(rows, bits)
	}
	return fmt.Sprintf("REDACTED", bA.Bits, strings.Join(rows, indent))
}

//
func (bA *BitList) Octets() []byte {
	bA.mtx.Lock()
	defer bA.mtx.Unlock()

	countOctets := (bA.Bits + 7) / 8
	octets := make([]byte, countOctets)
	for i := 0; i < len(bA.Elements); i++ {
		elementOctets := [8]byte{}
		binary.LittleEndian.PutUint64(elementOctets[:], bA.Elements[i])
		copy(octets[i*8:], elementOctets[:])
	}
	return octets
}

//
//
func (bA *BitList) Modify(o *BitList) {
	if bA == nil || o == nil {
		return
	}

	bA.mtx.Lock()
	o.mtx.Lock()
	copy(bA.Elements, o.Elements)
	o.mtx.Unlock()
	bA.mtx.Unlock()
}

//
//
func (bA *BitList) SerializeJSON() ([]byte, error) {
	if bA == nil {
		return []byte("REDACTED"), nil
	}

	bA.mtx.Lock()
	defer bA.mtx.Unlock()

	bits := "REDACTED"
	for i := 0; i < bA.Bits; i++ {
		if bA.fetchOrdinal(i) {
			bits += "REDACTED"
		} else {
			bits += "REDACTED"
		}
	}
	bits += "REDACTED"
	return []byte(bits), nil
}

var bitListJSONPattern = regexp.MustCompile("REDACTED")

//
//
func (bA *BitList) UnserializeJSON(bz []byte) error {
	b := string(bz)
	if b == "REDACTED" {
		//
		//
		bA.Bits = 0
		bA.Elements = nil
		return nil
	}

	//
	align := bitListJSONPattern.FindStringSubmatch(b)
	if align == nil {
		return fmt.Errorf("REDACTED", bitListJSONPattern.String(), b)
	}
	bits := align[1]

	//
	countBits := len(bits)
	bA2 := NewBitList(countBits)
	if bA2 == nil {
		//
		bA.Bits = 0
		bA.Elements = nil
		return nil
	}

	for i := 0; i < countBits; i++ {
		if bits[i] == 'x' {
			bA2.AssignOrdinal(i, true)
		}
	}
	*bA = *bA2 //
	return nil
}

//
func (bA *BitList) ToSchema() *cmtprotobits.BitList {
	if bA == nil || len(bA.Elements) == 0 {
		return nil
	}

	return &cmtprotobits.BitList{
		Bits:  int64(bA.Bits),
		Elements: bA.Elements,
	}
}

//
func (bA *BitList) FromSchema(schemaBitList *cmtprotobits.BitList) {
	if schemaBitList == nil {
		bA = nil
		return
	}

	bA.Bits = int(schemaBitList.Bits)
	if len(schemaBitList.Elements) > 0 {
		bA.Elements = schemaBitList.Elements
	}
}

//
//
//
func (bA *BitList) CertifySimple() error {
	if bA == nil {
		return nil
	}

	anticipatedElements := countMembers(bA.Volume())
	if anticipatedElements != len(bA.Elements) {
		return fmt.Errorf("REDACTED", bA.Volume(), len(bA.Elements), anticipatedElements)
	}
	return nil
}

func countMembers(bits int) int {
	return (bits + 63) / 64
}
