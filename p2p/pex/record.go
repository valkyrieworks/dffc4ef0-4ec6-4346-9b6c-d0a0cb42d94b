package pex

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/scratchfile"
)

//

type locationRegisterJSN struct {
	Key   string          `json:"key"`
	Locations []*recognizedLocator `json:"locations"`
}

func (a *locationRegister) persistTowardRecord(recordRoute string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.Tracer.Details("REDACTED", "REDACTED", a.extent())

	locations := make([]*recognizedLocator, 0, len(a.locationSearch))
	for _, ka := range a.locationSearch {
		locations = append(locations, ka)
	}
	anJSN := &locationRegisterJSN{
		Key:   a.key,
		Locations: locations,
	}

	jsnOctets, err := json.MarshalIndent(anJSN, "REDACTED", "REDACTED")
	if err != nil {
		a.Tracer.Failure("REDACTED", "REDACTED", err)
		return
	}
	err = scratchfile.PersistRecordIndivisible(recordRoute, jsnOctets, 0o644)
	if err != nil {
		a.Tracer.Failure("REDACTED", "REDACTED", recordRoute, "REDACTED", err)
	}
}

//
//
func (a *locationRegister) fetchOriginatingRecord(recordRoute string) bool {
	//
	_, err := os.Stat(recordRoute)
	if os.IsNotExist(err) {
		return false
	}

	//
	r, err := os.Open(recordRoute)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", recordRoute, err))
	}
	defer r.Close()
	anJSN := &locationRegisterJSN{}
	dec := json.NewDecoder(r)
	err = dec.Decode(anJSN)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", recordRoute, err))
	}

	//
	//
	a.key = anJSN.Key
	//
	for _, ka := range anJSN.Locations {
		for _, segmentPosition := range ka.Segments {
			segment := a.fetchSegment(ka.SegmentKind, segmentPosition)
			segment[ka.Location.Text()] = ka
		}
		a.locationSearch[ka.ID()] = ka
		if ka.SegmentKind == segmentKindFresh {
			a.nthFresh++
		} else {
			a.nthAged++
		}
	}
	return true
}
