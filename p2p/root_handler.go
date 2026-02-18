package p2p

import (
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/p2p/link"
)

//
//
//
//
//
//
//
type Handler interface {
	daemon.Daemon //

	//
	CollectionRouter(Toggeler)

	//
	//
	FetchStreams() []*link.StreamDefinition

	//
	//
	//
	//
	//
	//
	InitNode(node Node) Node

	//
	//
	AppendNode(node Node)

	//
	//
	DeleteNode(node Node, cause any)

	//
	//
	Accept(Packet)
}

//

type RootHandler struct {
	daemon.RootDaemon //
	Router              Toggeler
}

func NewRootHandler(label string, impl Handler) *RootHandler {
	return &RootHandler{
		RootDaemon: *daemon.NewRootDaemon(nil, label, impl),
		Router:      nil,
	}
}

func (br *RootHandler) CollectionRouter(sw Toggeler) {
	br.Router = sw
}
func (*RootHandler) FetchStreams() []*link.StreamDefinition { return nil }
func (*RootHandler) AppendNode(Node)                           {}
func (*RootHandler) DeleteNode(Node, any)                   {}
func (*RootHandler) Accept(Packet)                       {}
func (*RootHandler) InitNode(node Node) Node                { return node }
