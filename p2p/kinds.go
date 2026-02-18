package p2p

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/p2p/link"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

type (
	StreamDefinition = link.StreamDefinition
	LinkageState  = link.LinkageState
)

//
type Packet struct {
	Src       Node          //
	Signal   proto.Message //
	StreamUID byte
}

//
//
//
type Extractor interface {
	proto.Message

	//
	Disclose() (proto.Message, error)
}

//
type Adapter interface {
	proto.Message

	//
	Enclose() proto.Message
}

var (
	_ Adapter = &tmp2p.PexQuery{}
	_ Adapter = &tmp2p.PexLocations{}
)
