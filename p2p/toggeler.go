package p2p

import (
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
)

//
//
//
//
type Toggeler interface {
	daemon.Daemon

	HandlerAdministrator
	NodeAdministrator
	Dispatcher

	MemberDetails() MemberDetails
	Log() log.Tracer
}

type HandlerAdministrator interface {
	Handler(label string) (handler Handler, present bool)
	AppendHandler(label string, handler Handler) Handler
	DeleteHandler(label string, handler Handler)
}

type NodeAdministrator interface {
	Nodes() IDXNodeCollection
	CountNodes() (outgoing, incoming, calling int)
	MaximumCountOutgoingNodes() int

	AppendDurableNodes(locations []string) error
	AppendInternalNodeIDXDatastore(ids []string) error
	AppendAbsoluteNodeIDXDatastore(ids []string) error

	CallNodeWithLocation(address *NetLocation) error
	CallNodesAsync(nodes []string) error

	HaltNodeForFault(node Node, cause any)
	HaltNodeSmoothly(node Node)

	IsCallingOrCurrentLocation(address *NetLocation) bool
	IsNodeDurable(address *NetLocation) bool
	IsNodeAbsolute(id ID) bool

	StampNodeAsSound(node Node)
}

type Dispatcher interface {
	Multicast(e Packet) (successChannel chan bool)
	MulticastAsync(e Packet)
	AttemptMulticast(e Packet)
}
