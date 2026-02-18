package statusconnect

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/p2p"
)

var (
	_ p2p.Adapter = &SegmentQuery{}
	_ p2p.Adapter = &SegmentReply{}
	_ p2p.Adapter = &MirrorsQuery{}
	_ p2p.Adapter = &MirrorsReply{}
)

func (m *MirrorsReply) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Signal_Mirrorsreply{MirrorsReply: m}
	return sm
}

func (m *MirrorsQuery) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Signal_Mirrorsrequest{MirrorsQuery: m}
	return sm
}

func (m *SegmentReply) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Signal_Segmentreply{SegmentReply: m}
	return sm
}

func (m *SegmentQuery) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Signal_Segmentrequest{SegmentQuery: m}
	return sm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Signal_Segmentrequest:
		return m.FetchSegmentQuery(), nil

	case *Signal_Segmentreply:
		return m.FetchSegmentReply(), nil

	case *Signal_Mirrorsrequest:
		return m.FetchMirrorsQuery(), nil

	case *Signal_Mirrorsreply:
		return m.FetchMirrorsReply(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
