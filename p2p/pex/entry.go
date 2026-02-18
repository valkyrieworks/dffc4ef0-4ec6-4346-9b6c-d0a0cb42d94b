package pex

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/valkyrieworks/utils/tempentry"
)

//

type addressRegistryJSON struct {
	Key   string          `json:"key"`
	Locations []*recognizedLocation `json:"locations"`
}

func (a *addressLedger) persistToEntry(entryRoute string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.Tracer.Details("REDACTED", "REDACTED", a.volume())

	locations := make([]*recognizedLocation, 0, len(a.addressSearch))
	for _, ka := range a.addressSearch {
		locations = append(locations, ka)
	}
	aJSON := &addressRegistryJSON{
		Key:   a.key,
		Locations: locations,
	}

	jsonOctets, err := json.MarshalIndent(aJSON, "REDACTED", "REDACTED")
	if err != nil {
		a.Tracer.Fault("REDACTED", "REDACTED", err)
		return
	}
	err = tempentry.RecordEntryAtomic(entryRoute, jsonOctets, 0o644)
	if err != nil {
		a.Tracer.Fault("REDACTED", "REDACTED", entryRoute, "REDACTED", err)
	}
}

//
//
func (a *addressLedger) importFromEntry(entryRoute string) bool {
	//
	_, err := os.Stat(entryRoute)
	if os.IsNotExist(err) {
		return false
	}

	//
	r, err := os.Open(entryRoute)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", entryRoute, err))
	}
	defer r.Close()
	aJSON := &addressRegistryJSON{}
	dec := json.NewDecoder(r)
	err = dec.Decode(aJSON)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", entryRoute, err))
	}

	//
	//
	a.key = aJSON.Key
	//
	for _, ka := range aJSON.Locations {
		for _, containerOrdinal := range ka.Containers {
			container := a.fetchContainer(ka.ContainerKind, containerOrdinal)
			container[ka.Address.String()] = ka
		}
		a.addressSearch[ka.ID()] = ka
		if ka.ContainerKind == containerKindNew {
			a.nNew++
		} else {
			a.nAged++
		}
	}
	return true
}
