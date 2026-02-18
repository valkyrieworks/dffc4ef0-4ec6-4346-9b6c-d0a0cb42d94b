package txpool

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/p2p"
)

var (
	_ p2p.Adapter   = &Txs{}
	_ p2p.Extractor = &Signal{}
)

//
func (m *Txs) Enclose() proto.Message {
	mm := &Signal{}
	mm.Sum = &Signal_Trans{Txs: m}
	return mm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Signal_Trans:
		return m.FetchTrans(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
