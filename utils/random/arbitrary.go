package random

import (
	crand "crypto/rand"
	mrand "math/rand"
	"time"

	engineconnect "github.com/valkyrieworks/utils/align"
)

const (
	strRunes = "REDACTED" //
)

//
//
//
//
//
//
//
type Random struct {
	engineconnect.Lock
	random *mrand.Rand
}

var major *Random

func init() {
	major = NewRandom()
	major.init()
}

func NewRandom() *Random {
	random := &Random{}
	random.init()
	return random
}

func (r *Random) init() {
	bz := cRandomOctets(8)
	var origin uint64
	for i := 0; i < 8; i++ {
		origin |= uint64(bz[i])
		origin <<= 8
	}
	r.restore(int64(origin))
}

func (r *Random) restore(origin int64) {
	//
	//
	r.random = mrand.New(mrand.NewSource(origin))
}

//
//

func Source(origin int64) {
	major.Source(origin)
}

func Str(extent int) string {
	return major.Str(extent)
}

func Uint16() uint16 {
	return major.Uint16()
}

func Uint32() uint32 {
	return major.Uint32()
}

func Uint64() uint64 {
	return major.Uint64()
}

func Uint() uint {
	return major.Uint()
}

func Int16() int16 {
	return major.Int16()
}

func Int32() int32 {
	return major.Int32()
}

func Int64() int64 {
	return major.Int64()
}

func Int() int {
	return major.Int()
}

func Int31() int32 {
	return major.Int31()
}

func Int31n(n int32) int32 {
	return major.Int31n(n)
}

func Int63() int64 {
	return major.Int63()
}

func Int64count(n int64) int64 {
	return major.Int64count(n)
}

func Bool() bool {
	return major.Bool()
}

func Float32() float32 {
	return major.Float32()
}

func Float64() float64 {
	return major.Float64()
}

func Time() time.Time {
	return major.Time()
}

func Octets(n int) []byte {
	return major.Octets(n)
}

func Intn(n int) int {
	return major.Intn(n)
}

func Mode(n int) []int {
	return major.Mode(n)
}

//
//

func (r *Random) Source(origin int64) {
	r.Lock()
	r.restore(origin)
	r.Unlock()
}

//
func (r *Random) Str(extent int) string {
	if extent <= 0 {
		return "REDACTED"
	}

	runes := []byte{}
MAIN_CYCLE:
	for {
		val := r.Int63()
		for i := 0; i < 10; i++ {
			v := int(val & 0x3f) //
			if v >= 62 {         //
				val >>= 6
				continue
			}
			runes = append(runes, strRunes[v])
			if len(runes) == extent {
				break MAIN_CYCLE
			}
			val >>= 6
		}
	}

	return string(runes)
}

func (r *Random) Uint16() uint16 {
	return uint16(r.Uint32() & (1<<16 - 1))
}

func (r *Random) Uint32() uint32 {
	r.Lock()
	u32 := r.random.Uint32()
	r.Unlock()
	return u32
}

func (r *Random) Uint64() uint64 {
	return uint64(r.Uint32())<<32 + uint64(r.Uint32())
}

func (r *Random) Uint() uint {
	r.Lock()
	i := r.random.Int()
	r.Unlock()
	return uint(i)
}

func (r *Random) Int16() int16 {
	return int16(r.Uint32() & (1<<16 - 1))
}

func (r *Random) Int32() int32 {
	return int32(r.Uint32())
}

func (r *Random) Int64() int64 {
	return int64(r.Uint64())
}

func (r *Random) Int() int {
	r.Lock()
	i := r.random.Int()
	r.Unlock()
	return i
}

func (r *Random) Int31() int32 {
	r.Lock()
	i31 := r.random.Int31()
	r.Unlock()
	return i31
}

func (r *Random) Int31n(n int32) int32 {
	r.Lock()
	i31n := r.random.Int31n(n)
	r.Unlock()
	return i31n
}

func (r *Random) Int63() int64 {
	r.Lock()
	i63 := r.random.Int63()
	r.Unlock()
	return i63
}

func (r *Random) Int64count(n int64) int64 {
	r.Lock()
	i63n := r.random.Int63n(n)
	r.Unlock()
	return i63n
}

func (r *Random) Float32() float32 {
	r.Lock()
	f32 := r.random.Float32()
	r.Unlock()
	return f32
}

func (r *Random) Float64() float64 {
	r.Lock()
	f64 := r.random.Float64()
	r.Unlock()
	return f64
}

func (r *Random) Time() time.Time {
	return time.Unix(int64(r.Uint64()), 0)
}

//
//
func (r *Random) Octets(n int) []byte {
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
func (r *Random) Intn(n int) int {
	r.Lock()
	i := r.random.Intn(n)
	r.Unlock()
	return i
}

//
func (r *Random) Bool() bool {
	//
	//
	return r.Int63()%2 == 0
}

//
func (r *Random) Mode(n int) []int {
	r.Lock()
	mode := r.random.Perm(n)
	r.Unlock()
	return mode
}

//
//
//
func cRandomOctets(countOctets int) []byte {
	b := make([]byte, countOctets)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
