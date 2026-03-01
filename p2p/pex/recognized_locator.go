package pex

import (
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

//
//
type recognizedLocator struct {
	Location        *p2p.NetworkLocator `json:"location"`
	Src         *p2p.NetworkLocator `json:"src"`
	Segments     []int           `json:"segments"`
	Endeavors    int32           `json:"endeavors"`
	SegmentKind  byte            `json:"segment_kind"`
	FinalEffort time.Time       `json:"final_effort"`
	FinalTriumph time.Time       `json:"final_triumph"`
	FinalProhibitMoment time.Time       `json:"final_prohibit_moment"`
}

func freshRecognizedLocator(location *p2p.NetworkLocator, src *p2p.NetworkLocator) *recognizedLocator {
	return &recognizedLocator{
		Location:        location,
		Src:         src,
		Endeavors:    0,
		FinalEffort: time.Now(),
		SegmentKind:  segmentKindFresh,
		Segments:     nil,
	}
}

func (ka *recognizedLocator) ID() p2p.ID {
	return ka.Location.ID
}

func (ka *recognizedLocator) equalsAged() bool {
	return ka.SegmentKind == segmentKindAged
}

func (ka *recognizedLocator) equalsFresh() bool {
	return ka.SegmentKind == segmentKindFresh
}

func (ka *recognizedLocator) labelEffort() {
	now := time.Now()
	ka.FinalEffort = now
	ka.Endeavors++
}

func (ka *recognizedLocator) labelValid() {
	now := time.Now()
	ka.FinalEffort = now
	ka.Endeavors = 0
	ka.FinalTriumph = now
}

func (ka *recognizedLocator) ban(prohibitMoment time.Duration) {
	if ka.FinalProhibitMoment.Before(time.Now().Add(prohibitMoment)) {
		ka.FinalProhibitMoment = time.Now().Add(prohibitMoment)
	}
}

func (ka *recognizedLocator) equalsProhibited() bool {
	return ka.FinalProhibitMoment.After(time.Now())
}

func (ka *recognizedLocator) appendSegmentPointer(segmentOffset int) int {
	for _, segment := range ka.Segments {
		if segment == segmentOffset {
			//
			//
			return -1
		}
	}
	ka.Segments = append(ka.Segments, segmentOffset)
	return len(ka.Segments)
}

func (ka *recognizedLocator) discardSegmentPointer(segmentOffset int) int {
	segments := []int{}
	for _, segment := range ka.Segments {
		if segment != segmentOffset {
			segments = append(segments, segment)
		}
	}
	if len(segments) != len(ka.Segments)-1 {
		//
		//
		return -1
	}
	ka.Segments = segments
	return len(ka.Segments)
}

/**
t
:

e
k
d
k

t
.
*/
func (ka *recognizedLocator) equalsFlawed() bool {
	//
	if ka.SegmentKind == segmentKindAged {
		return false
	}

	//
	if ka.FinalEffort.After(time.Now().Add(-1 * time.Minute)) {
		return false
	}

	//

	//
	//
	if ka.FinalEffort.Before(time.Now().Add(-1 * countAbsentEpochs * time.Hour * 24)) {
		return true
	}

	//
	if ka.FinalTriumph.IsZero() && ka.Endeavors >= countAttempts {
		return true
	}

	//
	if ka.FinalTriumph.Before(time.Now().Add(-1*minimumFlawedEpochs*time.Hour*24)) &&
		ka.Endeavors >= maximumMishaps {
		return true
	}

	return false
}
