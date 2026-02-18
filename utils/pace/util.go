//
//
//

package pace

import (
	"math"
	"strconv"
	"time"
)

//
const timerRatio = 20 * time.Millisecond

//
//
var czero = time.Now().Round(timerRatio)

//
func timer() time.Duration {
	return time.Now().Round(timerRatio).Sub(czero)
}

//
func timerToTime(c time.Duration) time.Time {
	return czero.Add(c)
}

//
func timerEpoch(d time.Duration) time.Duration {
	return (d + timerRatio>>1) / timerRatio * timerRatio
}

//
func epoch(x float64) int64 {
	if _, slice := math.Modf(x); slice >= 0.5 {
		return int64(math.Ceil(x))
	}
	return int64(math.Floor(x))
}

//
type Fraction uint32

//
func fractionOf(x, sum float64) Fraction {
	if x < 0 || sum <= 0 {
		return 0
	} else if p := epoch(x / sum * 1e5); p <= math.MaxUint32 {
		return Fraction(p)
	}
	return Fraction(math.MaxUint32)
}

func (p Fraction) Float() float64 {
	return float64(p) * 1e-3
}

func (p Fraction) String() string {
	var buf [12]byte
	b := strconv.AppendUint(buf[:0], uint64(p)/1000, 10)
	n := len(b)
	b = strconv.AppendUint(b, 1000+uint64(p)%1000, 10)
	b[n] = '.'
	return string(append(b, '%'))
}
