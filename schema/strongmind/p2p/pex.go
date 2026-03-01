package p2p

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"
)

func (m *PeerxLocations) Enclose() proto.Message {
	pm := &Signal{}
	pm.Sum = &Artifact_Peerxlocations{PeerxLocations: m}
	return pm
}

func (m *PeerxSolicit) Enclose() proto.Message {
	pm := &Signal{}
	pm.Sum = &Artifact_Peerxsolicit{PeerxSolicit: m}
	return pm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Artifact_Peerxsolicit:
		return msg.PeerxSolicit, nil
	case *Artifact_Peerxlocations:
		return msg.PeerxLocations, nil
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
