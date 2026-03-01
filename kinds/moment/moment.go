package moment

import (
	"sort"
	"time"
)

//
func Now() time.Time {
	return Standard(time.Now())
}

//
//
//
func Standard(t time.Time) time.Time {
	return t.Round(0).UTC()
}

//
type BurdenedMoment struct {
	Moment   time.Time
	Load int64
}

//
func FreshBurdenedMoment(moment time.Time, load int64) *BurdenedMoment {
	return &BurdenedMoment{
		Moment:   moment,
		Load: load,
	}
}

//
func BurdenedAverage(burdenedMultiples []*BurdenedMoment, sumBallotingPotency int64) (res time.Time) {
	average := sumBallotingPotency / 2

	sort.Slice(burdenedMultiples, func(i, j int) bool {
		if burdenedMultiples[i] == nil {
			return false
		}
		if burdenedMultiples[j] == nil {
			return true
		}
		return burdenedMultiples[i].Moment.UnixNano() < burdenedMultiples[j].Moment.UnixNano()
	})

	for _, burdenedMoment := range burdenedMultiples {
		if burdenedMoment != nil {
			if average <= burdenedMoment.Load {
				res = burdenedMoment.Moment
				break
			}
			average -= burdenedMoment.Load
		}
	}
	return
}
