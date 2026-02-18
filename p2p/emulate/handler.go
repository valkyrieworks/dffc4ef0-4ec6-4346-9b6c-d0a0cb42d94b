package emulate

import (
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/link"
)

type Handler struct {
	p2p.RootHandler

	Streams []*link.StreamDefinition
}

func NewHandler() *Handler {
	r := &Handler{}
	r.RootHandler = *p2p.NewRootHandler("REDACTED", r)
	r.AssignTracer(log.VerifyingTracer())
	return r
}

func (r *Handler) FetchStreams() []*link.StreamDefinition { return r.Streams }
func (r *Handler) AppendNode(_ p2p.Node)                     {}
func (r *Handler) DeleteNode(_ p2p.Node, _ any)           {}
func (r *Handler) Accept(_ p2p.Packet)                 {}
