package p2p

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

//
//
//
//
type Router interface {
	facility.Facility

	HandlerAdministrator
	NodeAdministrator
	Publisher

	PeerDetails() PeerDetails
	Log() log.Tracer
}

type HandlerAdministrator interface {
	Handler(alias string) (handler Handler, present bool)
	AppendHandler(alias string, handler Handler) Handler
	DiscardHandler(alias string, handler Handler)
}

type NodeAdministrator interface {
	Nodes() IDXNodeAssign
	CountNodes() (outgoing, incoming, calling int)
	MaximumCountOutgoingNodes() int

	AppendEnduringNodes(locations []string) error
	AppendSecludedNodeIDXDstore(ids []string) error
	AppendAbsoluteNodeIDXDstore(ids []string) error

	CallNodeUsingLocator(location *NetworkLocator) error
	CallNodesAsyncronous(nodes []string) error

	HaltNodeForeachFailure(node Node, rationale any)
	HaltNodeSmoothly(node Node)

	EqualsCallingEitherCurrentLocator(location *NetworkLocator) bool
	EqualsNodeEnduring(location *NetworkLocator) bool
	EqualsNodeAbsolute(id ID) bool

	LabelNodeLikeValid(node Node)
}

type Publisher interface {
	Multicast(e Wrapper) (triumphChn chan bool)
	MulticastAsyncronous(e Wrapper)
	AttemptMulticast(e Wrapper)
}
