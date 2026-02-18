package pex

import (
	"time"

	"github.com/valkyrieworks/p2p"
)

//
//
type recognizedLocation struct {
	Address        *p2p.NetLocation `json:"address"`
	Src         *p2p.NetLocation `json:"src"`
	Segments     []int           `json:"segments"`
	Tries    int32           `json:"tries"`
	SegmentKind  byte            `json:"container_kind"`
	FinalEndeavor time.Time       `json:"final_endeavor"`
	FinalSuccess time.Time       `json:"final_success"`
	FinalProhibitTime time.Time       `json:"final_prohibit_time"`
}

func newRecognizedLocation(address *p2p.NetLocation, src *p2p.NetLocation) *recognizedLocation {
	return &recognizedLocation{
		Address:        address,
		Src:         src,
		Tries:    0,
		FinalEndeavor: time.Now(),
		SegmentKind:  segmentKindNew,
		Segments:     nil,
	}
}

func (ka *recognizedLocation) ID() p2p.ID {
	return ka.Address.ID
}

func (ka *recognizedLocation) isAged() bool {
	return ka.SegmentKind == segmentKindAged
}

func (ka *recognizedLocation) isNew() bool {
	return ka.SegmentKind == segmentKindNew
}

func (ka *recognizedLocation) stampEndeavor() {
	now := time.Now()
	ka.FinalEndeavor = now
	ka.Tries++
}

func (ka *recognizedLocation) stampValid() {
	now := time.Now()
	ka.FinalEndeavor = now
	ka.Tries = 0
	ka.FinalSuccess = now
}

func (ka *recognizedLocation) ban(prohibitTime time.Duration) {
	if ka.FinalProhibitTime.Before(time.Now().Add(prohibitTime)) {
		ka.FinalProhibitTime = time.Now().Add(prohibitTime)
	}
}

func (ka *recognizedLocation) isProhibited() bool {
	return ka.FinalProhibitTime.After(time.Now())
}

func (ka *recognizedLocation) appendSegmentReference(segmentIndex int) int {
	for _, segment := range ka.Segments {
		if segment == segmentIndex {
			//
			//
			return -1
		}
	}
	ka.Segments = append(ka.Segments, segmentIndex)
	return len(ka.Segments)
}

func (ka *recognizedLocation) deleteSegmentReference(segmentIndex int) int {
	segments := []int{}
	for _, segment := range ka.Segments {
		if segment != segmentIndex {
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
func (ka *recognizedLocation) isFlawed() bool {
	//
	if ka.SegmentKind == segmentKindAged {
		return false
	}

	//
	if ka.FinalEndeavor.After(time.Now().Add(-1 * time.Minute)) {
		return false
	}

	//

	//
	//
	if ka.FinalEndeavor.Before(time.Now().Add(-1 * countAbsentPeriods * time.Hour * 24)) {
		return true
	}

	//
	if ka.FinalSuccess.IsZero() && ka.Tries >= countAttempts {
		return true
	}

	//
	if ka.FinalSuccess.Before(time.Now().Add(-1*minimumFlawedPeriods*time.Hour*24)) &&
		ka.Tries >= maximumBreakdowns {
		return true
	}

	return false
}
