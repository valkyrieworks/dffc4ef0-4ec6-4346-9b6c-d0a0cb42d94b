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
type ScaledTime struct {
	Time   time.Time
	Magnitude int64
}

//
func NewScaledTime(moment time.Time, magnitude int64) *ScaledTime {
	return &ScaledTime{
		Time:   moment,
		Magnitude: magnitude,
	}
}

//
func ScaledAverage(scaledInstances []*ScaledTime, sumPollingEnergy int64) (res time.Time) {
	average := sumPollingEnergy / 2

	sort.Slice(scaledInstances, func(i, j int) bool {
		if scaledInstances[i] == nil {
			return false
		}
		if scaledInstances[j] == nil {
			return true
		}
		return scaledInstances[i].Time.UnixNano() < scaledInstances[j].Time.UnixNano()
	})

	for _, scaledTime := range scaledInstances {
		if scaledTime != nil {
			if average <= scaledTime.Magnitude {
				res = scaledTime.Time
				break
			}
			average -= scaledTime.Magnitude
		}
	}
	return
}
