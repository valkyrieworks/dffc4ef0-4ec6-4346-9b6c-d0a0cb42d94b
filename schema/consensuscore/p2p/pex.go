package p2p

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"
)

func (m *PexLocations) Enclose() proto.Message {
	pm := &Signal{}
	pm.Sum = &Signal_Pexlocations{PexLocations: m}
	return pm
}

func (m *PexQuery) Enclose() proto.Message {
	pm := &Signal{}
	pm.Sum = &Signal_Pexquery{PexQuery: m}
	return pm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Signal_Pexquery:
		return msg.PexQuery, nil
	case *Signal_Pexlocations:
		return msg.PexLocations, nil
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
