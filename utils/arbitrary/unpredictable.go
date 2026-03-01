package arbitrary

import (
	crand "crypto/rand"
	mrand "math/rand"
	"time"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

const (
	txtSymbols = "REDACTED" //
)

//
//
//
//
//
//
//
type Arbitrary struct {
	commitchronize.Exclusion
	arbitrary *mrand.Rand
}

var majestic *Arbitrary

func initialize() {
	majestic = FreshArbitrary()
	majestic.initialize()
}

func FreshArbitrary() *Arbitrary {
	arbitrary := &Arbitrary{}
	arbitrary.initialize()
	return arbitrary
}

func (r *Arbitrary) initialize() {
	bz := cnArbitraryOctets(8)
	var germ uint64
	for i := 0; i < 8; i++ {
		germ |= uint64(bz[i])
		germ <<= 8
	}
	r.restore(int64(germ))
}

func (r *Arbitrary) restore(germ int64) {
	//
	//
	r.arbitrary = mrand.New(mrand.NewSource(germ))
}

//
//

func Germ(germ int64) {
	majestic.Germ(germ)
}

func Str(magnitude int) string {
	return majestic.Str(magnitude)
}

func Uint16() uint16 {
	return majestic.Uint16()
}

func Uint32n() uint32 {
	return majestic.Uint32n()
}

func Uint64n() uint64 {
	return majestic.Uint64n()
}

func Uintn() uint {
	return majestic.Uintn()
}

func Integer16() int16 {
	return majestic.Integer16()
}

func Integer32() int32 {
	return majestic.Integer32()
}

func Int64n() int64 {
	return majestic.Int64n()
}

func Int() int {
	return majestic.Int()
}

func Int31n() int32 {
	return majestic.Int31n()
}

func Integer31n(n int32) int32 {
	return majestic.Integer31n(n)
}

func Int63n() int64 {
	return majestic.Int63n()
}

func Int63num(n int64) int64 {
	return majestic.Int63num(n)
}

func Flag() bool {
	return majestic.Flag()
}

func Float32() float32 {
	return majestic.Float32()
}

func Float64() float64 {
	return majestic.Float64()
}

func Moment() time.Time {
	return majestic.Moment()
}

func Octets(n int) []byte {
	return majestic.Octets(n)
}

func Integern(n int) int {
	return majestic.Integern(n)
}

func Mode(n int) []int {
	return majestic.Mode(n)
}

//
//

func (r *Arbitrary) Germ(germ int64) {
	r.Lock()
	r.restore(germ)
	r.Unlock()
}

//
func (r *Arbitrary) Str(magnitude int) string {
	if magnitude <= 0 {
		return "REDACTED"
	}

	symbols := []byte{}
PRIMARY_CYCLE:
	for {
		val := r.Int63n()
		for i := 0; i < 10; i++ {
			v := int(val & 0x3f) //
			if v >= 62 {         //
				val >>= 6
				continue
			}
			symbols = append(symbols, txtSymbols[v])
			if len(symbols) == magnitude {
				break PRIMARY_CYCLE
			}
			val >>= 6
		}
	}

	return string(symbols)
}

func (r *Arbitrary) Uint16() uint16 {
	return uint16(r.Uint32n() & (1<<16 - 1))
}

func (r *Arbitrary) Uint32n() uint32 {
	r.Lock()
	u32 := r.arbitrary.Uint32()
	r.Unlock()
	return u32
}

func (r *Arbitrary) Uint64n() uint64 {
	return uint64(r.Uint32n())<<32 + uint64(r.Uint32n())
}

func (r *Arbitrary) Uintn() uint {
	r.Lock()
	i := r.arbitrary.Int()
	r.Unlock()
	return uint(i)
}

func (r *Arbitrary) Integer16() int16 {
	return int16(r.Uint32n() & (1<<16 - 1))
}

func (r *Arbitrary) Integer32() int32 {
	return int32(r.Uint32n())
}

func (r *Arbitrary) Int64n() int64 {
	return int64(r.Uint64n())
}

func (r *Arbitrary) Int() int {
	r.Lock()
	i := r.arbitrary.Int()
	r.Unlock()
	return i
}

func (r *Arbitrary) Int31n() int32 {
	r.Lock()
	i31 := r.arbitrary.Int31()
	r.Unlock()
	return i31
}

func (r *Arbitrary) Integer31n(n int32) int32 {
	r.Lock()
	integer31n := r.arbitrary.Int31n(n)
	r.Unlock()
	return integer31n
}

func (r *Arbitrary) Int63n() int64 {
	r.Lock()
	i63 := r.arbitrary.Int63()
	r.Unlock()
	return i63
}

func (r *Arbitrary) Int63num(n int64) int64 {
	r.Lock()
	integer63n := r.arbitrary.Int63n(n)
	r.Unlock()
	return integer63n
}

func (r *Arbitrary) Float32() float32 {
	r.Lock()
	f32 := r.arbitrary.Float32()
	r.Unlock()
	return f32
}

func (r *Arbitrary) Float64() float64 {
	r.Lock()
	f64 := r.arbitrary.Float64()
	r.Unlock()
	return f64
}

func (r *Arbitrary) Moment() time.Time {
	return time.Unix(int64(r.Uint64n()), 0)
}

//
//
func (r *Arbitrary) Octets(n int) []byte {
	//
	//
	bs := make([]byte, n)
	for i := 0; i < len(bs); i++ {
		bs[i] = byte(r.Int() & 0xFF)
	}
	return bs
}

//
//
func (r *Arbitrary) Integern(n int) int {
	r.Lock()
	i := r.arbitrary.Intn(n)
	r.Unlock()
	return i
}

//
func (r *Arbitrary) Flag() bool {
	//
	//
	return r.Int63n()%2 == 0
}

//
func (r *Arbitrary) Mode(n int) []int {
	r.Lock()
	mode := r.arbitrary.Perm(n)
	r.Unlock()
	return mode
}

//
//
//
func cnArbitraryOctets(countOctets int) []byte {
	b := make([]byte, countOctets)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
