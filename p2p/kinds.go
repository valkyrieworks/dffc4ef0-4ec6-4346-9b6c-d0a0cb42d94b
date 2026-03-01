package p2p

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

type (
	ConduitDefinition = link.ConduitDefinition
	LinkageCondition  = link.LinkageCondition
)

//
type Wrapper struct {
	Src       Node          //
	Signal   proto.Message //
	ConduitUUID byte
}

//
//
//
type Unwrapper interface {
	proto.Message

	//
	Disclose() (proto.Message, error)
}

//
type Encapsulator interface {
	proto.Message

	//
	Enclose() proto.Message
}

var (
	_ Encapsulator = &tmpfabric.PeerxSolicit{}
	_ Encapsulator = &tmpfabric.PeerxLocations{}
)
