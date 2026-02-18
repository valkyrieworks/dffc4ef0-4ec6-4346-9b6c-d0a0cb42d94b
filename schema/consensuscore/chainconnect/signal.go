package chainconnect

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/p2p"
)

var (
	_ p2p.Adapter = &StatusQuery{}
	_ p2p.Adapter = &StatusReply{}
	_ p2p.Adapter = &NoLedgerReply{}
	_ p2p.Adapter = &LedgerReply{}
	_ p2p.Adapter = &LedgerQuery{}
)

const (
	LedgerReplySignalHeadingVolume   = 4
	LedgerReplySignalFieldKeyVolume = 1
)

func (m *LedgerQuery) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Chainrequest{LedgerQuery: m}
	return bm
}

func (m *LedgerReply) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Chainreply{LedgerReply: m}
	return bm
}

func (m *NoLedgerReply) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Nonledgerreply{NoLedgerReply: m}
	return bm
}

func (m *StatusQuery) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Staterequest{StatusQuery: m}
	return bm
}

func (m *StatusReply) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Statereply{StatusReply: m}
	return bm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Signal_Chainrequest:
		return m.FetchLedgerQuery(), nil

	case *Signal_Chainreply:
		return m.FetchLedgerReply(), nil

	case *Signal_Nonledgerreply:
		return m.FetchNoLedgerReply(), nil

	case *Signal_Staterequest:
		return m.FetchStateQuery(), nil

	case *Signal_Statereply:
		return m.FetchStateReply(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
