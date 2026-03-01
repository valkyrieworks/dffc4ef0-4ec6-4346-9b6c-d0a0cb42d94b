package txpool

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

var (
	_ p2p.Encapsulator   = &Txs{}
	_ p2p.Unwrapper = &Signal{}
)

//
func (m *Txs) Enclose() proto.Message {
	mm := &Signal{}
	mm.Sum = &Artifact_Trans{Txs: m}
	return mm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Artifact_Trans:
		return m.ObtainTrans(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
