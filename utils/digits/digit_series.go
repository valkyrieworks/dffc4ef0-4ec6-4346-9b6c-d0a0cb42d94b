package digits

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"regexp"
	"strings"
	"sync"

	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	schemaoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/utils/digits"
)

//
type DigitSeries struct {
	mtx   sync.Mutex
	Digits  int      `json:"digits"`  //
	Components []uint64 `json:"components"` //
}

//
//
func FreshDigitCollection(digits int) *DigitSeries {
	if digits <= 0 {
		return nil
	}
	return &DigitSeries{
		Digits:  digits,
		Components: make([]uint64, countConstituents(digits)),
	}
}

//
//
//
func FreshDigitSeriesOriginatingProc(digits int, fn func(int) bool) *DigitSeries {
	if digits <= 0 {
		return nil
	}
	bA := &DigitSeries{
		Digits:  digits,
		Components: make([]uint64, countConstituents(digits)),
	}
	for i := 0; i < digits; i++ {
		v := fn(i)
		if v {
			bA.Components[i/64] |= (uint64(1) << uint(i%64))
		}
	}
	return bA
}

//
func (bA *DigitSeries) Extent() int {
	if bA == nil {
		return 0
	}
	return bA.Digits
}

//
//
func (bA *DigitSeries) ObtainOrdinal(i int) bool {
	if bA == nil {
		return false
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.fetchPosition(i)
}

func (bA *DigitSeries) fetchPosition(i int) bool {
	if i >= bA.Digits {
		return false
	}
	return bA.Components[i/64]&(uint64(1)<<uint(i%64)) > 0
}

//
//
func (bA *DigitSeries) AssignOrdinal(i int, v bool) bool {
	if bA == nil {
		return false
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.assignPosition(i, v)
}

func (bA *DigitSeries) assignPosition(i int, v bool) bool {
	if i >= bA.Digits || i/64 >= len(bA.Components) {
		return false
	}
	if v {
		bA.Components[i/64] |= (uint64(1) << uint(i%64))
	} else {
		bA.Components[i/64] &= ^(uint64(1) << uint(i%64))
	}
	return true
}

//
func (bA *DigitSeries) Duplicate() *DigitSeries {
	if bA == nil {
		return nil
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.duplicate()
}

func (bA *DigitSeries) duplicate() *DigitSeries {
	c := make([]uint64, len(bA.Components))
	copy(c, bA.Components)
	return &DigitSeries{
		Digits:  bA.Digits,
		Components: c,
	}
}

func (bA *DigitSeries) duplicateDigits(digits int) *DigitSeries {
	c := make([]uint64, countConstituents(digits))
	copy(c, bA.Components)
	return &DigitSeries{
		Digits:  digits,
		Components: c,
	}
}

//
//
//
func (bA *DigitSeries) Or(o *DigitSeries) *DigitSeries {
	if bA == nil && o == nil {
		return nil
	}
	if bA == nil && o != nil {
		return o.Duplicate()
	}
	if o == nil {
		return bA.Duplicate()
	}
	bA.mtx.Lock()
	o.mtx.Lock()
	c := bA.duplicateDigits(strongarithmetic.MaximumInteger(bA.Digits, o.Digits))
	tinier := strongarithmetic.MinimumInteger(len(bA.Components), len(o.Components))
	for i := 0; i < tinier; i++ {
		c.Components[i] |= o.Components[i]
	}
	bA.mtx.Unlock()
	o.mtx.Unlock()
	return c
}

//
//
//
func (bA *DigitSeries) And(o *DigitSeries) *DigitSeries {
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

func (bA *DigitSeries) and(o *DigitSeries) *DigitSeries {
	c := bA.duplicateDigits(strongarithmetic.MinimumInteger(bA.Digits, o.Digits))
	for i := 0; i < len(c.Components); i++ {
		c.Components[i] &= o.Components[i]
	}
	return c
}

//
func (bA *DigitSeries) Not() *DigitSeries {
	if bA == nil {
		return nil //
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.not()
}

func (bA *DigitSeries) not() *DigitSeries {
	c := bA.duplicate()
	for i := 0; i < len(c.Components); i++ {
		c.Components[i] = ^c.Components[i]
	}
	return c
}

//
//
//
//
func (bA *DigitSeries) Sub(o *DigitSeries) *DigitSeries {
	if bA == nil || o == nil {
		//
		return nil
	}
	bA.mtx.Lock()
	o.mtx.Lock()
	//
	c := bA.duplicateDigits(bA.Digits)
	//
	//
	//
	//
	tinier := strongarithmetic.MinimumInteger(len(bA.Components), len(o.Components))
	for i := 0; i < tinier; i++ {
		//
		c.Components[i] &^= o.Components[i]
	}
	bA.mtx.Unlock()
	o.mtx.Unlock()
	return c
}

//
func (bA *DigitSeries) EqualsBlank() bool {
	if bA == nil {
		return true //
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	for _, e := range bA.Components {
		if e > 0 {
			return false
		}
	}
	return true
}

//
func (bA *DigitSeries) EqualsComplete() bool {
	if bA == nil {
		return true
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()

	//
	for _, member := range bA.Components[:len(bA.Components)-1] {
		if (^member) != 0 {
			return false
		}
	}

	//
	finalMemberDigits := (bA.Digits+63)%64 + 1
	finalMember := bA.Components[len(bA.Components)-1]
	return (finalMember+1)&((uint64(1)<<uint(finalMemberDigits))-1) == 0
}

//
//
//
func (bA *DigitSeries) SelectArbitrary() (int, bool) {
	if bA == nil {
		return 0, false
	}

	bA.mtx.Lock()
	countSuccessPositions := bA.fetchCountSuccessPositions()
	if countSuccessPositions == 0 { //
		bA.mtx.Unlock()
		return 0, false
	}
	ordinal := bA.fetchOrdinalSuccessPosition(commitrand.Integern(countSuccessPositions))
	bA.mtx.Unlock()
	if ordinal == -1 {
		return 0, false
	}
	return ordinal, true
}

func (bA *DigitSeries) fetchCountSuccessPositions() int {
	if bA.Extent() == 0 || len(bA.Components) == 0 || len(bA.Components) != countConstituents(bA.Extent()) {
		//
		return 0
	}

	tally := 0
	countComponents := len(bA.Components)
	//
	for i := 0; i < countComponents-1; i++ {
		tally += bits.OnesCount64(bA.Components[i])
	}
	//
	countUltimateDigits := bA.Digits - (countComponents-1)*64
	for i := 0; i < countUltimateDigits; i++ {
		if (bA.Components[countComponents-1] & (uint64(1) << uint64(i))) > 0 {
			tally++
		}
	}
	return tally
}

//
//
//
func (bA *DigitSeries) fetchOrdinalSuccessPosition(n int) int {
	countComponents := len(bA.Components)
	tally := 0

	//
	for i := 0; i < countComponents; i++ {
		//
		assignDigits := bits.OnesCount64(bA.Components[i])

		//
		//
		if tally+assignDigits >= n {
			//
			for j := 0; j < 64; j++ {
				if bA.Components[i]&(1<<uint(j)) != 0 {
					if tally == n {
						//
						return i*64 + j
					}
					tally++
				}
			}
		} else {
			//
			tally += assignDigits
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
func (bA *DigitSeries) Text() string {
	return bA.TextFormatted("REDACTED")
}

//
//
func (bA *DigitSeries) TextFormatted(format string) string {
	if bA == nil {
		return "REDACTED"
	}
	bA.mtx.Lock()
	defer bA.mtx.Unlock()
	return bA.textFormatted(format)
}

func (bA *DigitSeries) textFormatted(format string) string {
	traces := []string{}
	digits := "REDACTED"
	for i := 0; i < bA.Digits; i++ {
		if bA.fetchPosition(i) {
			digits += "REDACTED"
		} else {
			digits += "REDACTED"
		}
		if i%100 == 99 {
			traces = append(traces, digits)
			digits = "REDACTED"
		}
		if i%10 == 9 {
			digits += format
		}
		if i%50 == 49 {
			digits += format
		}
	}
	if len(digits) > 0 {
		traces = append(traces, digits)
	}
	return fmt.Sprintf("REDACTED", bA.Digits, strings.Join(traces, format))
}

//
func (bA *DigitSeries) Octets() []byte {
	bA.mtx.Lock()
	defer bA.mtx.Unlock()

	countOctets := (bA.Digits + 7) / 8
	octets := make([]byte, countOctets)
	for i := 0; i < len(bA.Components); i++ {
		memberOctets := [8]byte{}
		binary.LittleEndian.PutUint64(memberOctets[:], bA.Components[i])
		copy(octets[i*8:], memberOctets[:])
	}
	return octets
}

//
//
func (bA *DigitSeries) Revise(o *DigitSeries) {
	if bA == nil || o == nil {
		return
	}

	bA.mtx.Lock()
	o.mtx.Lock()
	copy(bA.Components, o.Components)
	o.mtx.Unlock()
	bA.mtx.Unlock()
}

//
//
func (bA *DigitSeries) SerializeJSN() ([]byte, error) {
	if bA == nil {
		return []byte("REDACTED"), nil
	}

	bA.mtx.Lock()
	defer bA.mtx.Unlock()

	digits := "REDACTED"
	for i := 0; i < bA.Digits; i++ {
		if bA.fetchPosition(i) {
			digits += "REDACTED"
		} else {
			digits += "REDACTED"
		}
	}
	digits += "REDACTED"
	return []byte(digits), nil
}

var digitSeriesJSNPattern = regexp.MustCompile("REDACTED")

//
//
func (bA *DigitSeries) DecodeJSN(bz []byte) error {
	b := string(bz)
	if b == "REDACTED" {
		//
		//
		bA.Digits = 0
		bA.Components = nil
		return nil
	}

	//
	align := digitSeriesJSNPattern.FindStringSubmatch(b)
	if align == nil {
		return fmt.Errorf("REDACTED", digitSeriesJSNPattern.String(), b)
	}
	digits := align[1]

	//
	countDigits := len(digits)
	bA2 := FreshDigitCollection(countDigits)
	if bA2 == nil {
		//
		bA.Digits = 0
		bA.Components = nil
		return nil
	}

	for i := 0; i < countDigits; i++ {
		if digits[i] == 'x' {
			bA2.AssignOrdinal(i, true)
		}
	}
	*bA = *bA2 //
	return nil
}

//
func (bA *DigitSeries) TowardSchema() *schemaoctets.DigitSeries {
	if bA == nil || len(bA.Components) == 0 {
		return nil
	}

	return &schemaoctets.DigitSeries{
		Digits:  int64(bA.Digits),
		Components: bA.Components,
	}
}

//
func (bA *DigitSeries) OriginatingSchema(schemaDigitSeries *schemaoctets.DigitSeries) {
	if schemaDigitSeries == nil {
		bA = nil
		return
	}

	bA.Digits = int(schemaDigitSeries.Digits)
	if len(schemaDigitSeries.Components) > 0 {
		bA.Components = schemaDigitSeries.Components
	}
}

//
//
//
func (bA *DigitSeries) CertifyFundamental() error {
	if bA == nil {
		return nil
	}

	anticipatedComponents := countConstituents(bA.Extent())
	if anticipatedComponents != len(bA.Components) {
		return fmt.Errorf("REDACTED", bA.Extent(), len(bA.Components), anticipatedComponents)
	}
	return nil
}

func countConstituents(digits int) int {
	return (digits + 63) / 64
}
