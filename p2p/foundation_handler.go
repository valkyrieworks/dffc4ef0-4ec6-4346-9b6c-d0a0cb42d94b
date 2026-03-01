package p2p

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

//
//
//
//
//
//
//
type Handler interface {
	facility.Facility //

	//
	AssignRouter(Router)

	//
	//
	ObtainConduits() []*link.ConduitDefinition

	//
	//
	//
	//
	//
	//
	InitializeNode(node Node) Node

	//
	//
	AppendNode(node Node)

	//
	//
	DiscardNode(node Node, rationale any)

	//
	//
	Accept(Wrapper)
}

//

type FoundationHandler struct {
	facility.FoundationFacility //
	Router              Router
}

func FreshFoundationHandler(alias string, implementation Handler) *FoundationHandler {
	return &FoundationHandler{
		FoundationFacility: *facility.FreshFoundationFacility(nil, alias, implementation),
		Router:      nil,
	}
}

func (br *FoundationHandler) AssignRouter(sw Router) {
	br.Router = sw
}
func (*FoundationHandler) ObtainConduits() []*link.ConduitDefinition { return nil }
func (*FoundationHandler) AppendNode(Node)                           {}
func (*FoundationHandler) DiscardNode(Node, any)                   {}
func (*FoundationHandler) Accept(Wrapper)                       {}
func (*FoundationHandler) InitializeNode(node Node) Node                { return node }
