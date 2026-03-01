package chainchronize

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

var (
	_ p2p.Encapsulator = &ConditionSolicit{}
	_ p2p.Encapsulator = &ConditionReply{}
	_ p2p.Encapsulator = &NegativeLedgerReply{}
	_ p2p.Encapsulator = &LedgerReply{}
	_ p2p.Encapsulator = &LedgerSolicit{}
)

const (
	LedgerReplySignalHeadingExtent   = 4
	LedgerReplySignalAttributeTokenExtent = 1
)

func (m *LedgerSolicit) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Ledgerrequest{LedgerSolicit: m}
	return bm
}

func (m *LedgerReply) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Ledgerreply{LedgerReply: m}
	return bm
}

func (m *NegativeLedgerReply) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Noledgerreply{NegativeLedgerReply: m}
	return bm
}

func (m *ConditionSolicit) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Conditionrequest{ConditionSolicit: m}
	return bm
}

func (m *ConditionReply) Enclose() proto.Message {
	bm := &Signal{}
	bm.Sum = &Signal_Conditionreply{ConditionReply: m}
	return bm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Signal_Ledgerrequest:
		return m.ObtainLedgerSolicit(), nil

	case *Signal_Ledgerreply:
		return m.ObtainLedgerReply(), nil

	case *Signal_Noledgerreply:
		return m.ObtainNegativeLedgerReply(), nil

	case *Signal_Conditionrequest:
		return m.ObtainConditionSolicit(), nil

	case *Signal_Conditionreply:
		return m.ObtainConditionReply(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
