package simulate

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

type Handler struct {
	p2p.FoundationHandler

	Conduits []*link.ConduitDefinition
}

func FreshHandler() *Handler {
	r := &Handler{}
	r.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", r)
	r.AssignTracer(log.VerifyingTracer())
	return r
}

func (r *Handler) ObtainConduits() []*link.ConduitDefinition { return r.Conduits }
func (r *Handler) AppendNode(_ p2p.Node)                     {}
func (r *Handler) DiscardNode(_ p2p.Node, _ any)           {}
func (r *Handler) Accept(_ p2p.Wrapper)                 {}
