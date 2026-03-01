//
//
//

package throughput

import (
	"math"
	"strconv"
	"time"
)

//
const timerFrequency = 20 * time.Millisecond

//
//
var cnull = time.Now().Round(timerFrequency)

//
func timer() time.Duration {
	return time.Now().Round(timerFrequency).Sub(cnull)
}

//
func timerTowardMoment(c time.Duration) time.Time {
	return cnull.Add(c)
}

//
func timerIteration(d time.Duration) time.Duration {
	return (d + timerFrequency>>1) / timerFrequency * timerFrequency
}

//
func iteration(x float64) int64 {
	if _, division := math.Modf(x); division >= 0.5 {
		return int64(math.Ceil(x))
	}
	return int64(math.Floor(x))
}

//
type Ratio uint32

//
func ratioBelonging(x, sum float64) Ratio {
	if x < 0 || sum <= 0 {
		return 0
	} else if p := iteration(x / sum * 1e5); p <= math.MaxUint32 {
		return Ratio(p)
	}
	return Ratio(math.MaxUint32)
}

func (p Ratio) Decimal() float64 {
	return float64(p) * 1e-3
}

func (p Ratio) Text() string {
	var buf [12]byte
	b := strconv.AppendUint(buf[:0], uint64(p)/1000, 10)
	n := len(b)
	b = strconv.AppendUint(b, 1000+uint64(p)%1000, 10)
	b[n] = '.'
	return string(append(b, '%'))
}
