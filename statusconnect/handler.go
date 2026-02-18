package statusconnect

import (
	"context"
	"errors"
	"sort"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
	statusproto "github.com/valkyrieworks/schema/consensuscore/statusconnect"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	MirrorStream = byte(0x60)
	//
	SegmentStream = byte(0x61)
	//
	currentMirrors = 10
)

//
//
type Handler struct {
	p2p.RootHandler

	cfg       settings.StatusAlignSettings
	link      gateway.ApplicationLinkMirror
	linkInquire gateway.ApplicationLinkInquire
	temporaryFolder   string
	stats   *Stats

	//
	//
	mtx    engineconnect.ReadwriteLock
	aligner *aligner
}

//
func NewHandler(
	cfg settings.StatusAlignSettings,
	link gateway.ApplicationLinkMirror,
	linkInquire gateway.ApplicationLinkInquire,
	stats *Stats,
) *Handler {
	r := &Handler{
		cfg:       cfg,
		link:      link,
		linkInquire: linkInquire,
		stats:   stats,
	}
	r.RootHandler = *p2p.NewRootHandler("REDACTED", r)

	return r
}

//
func (r *Handler) FetchStreams() []*p2p.StreamDefinition {
	return []*p2p.StreamDefinition{
		{
			ID:                  MirrorStream,
			Urgency:            5,
			TransmitBufferVolume:   10,
			AcceptSignalVolume: mirrorMessageVolume,
			SignalKind:         &statusproto.Signal{},
		},
		{
			ID:                  SegmentStream,
			Urgency:            3,
			TransmitBufferVolume:   10,
			AcceptSignalVolume: segmentMessageVolume,
			SignalKind:         &statusproto.Signal{},
		},
	}
}

//
func (r *Handler) OnBegin() error {
	return nil
}

//
func (r *Handler) AppendNode(node p2p.Node) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.aligner != nil {
		r.aligner.AppendNode(node)
	}
}

//
func (r *Handler) DeleteNode(node p2p.Node, _ any) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.aligner != nil {
		r.aligner.DeleteNode(node)
	}
}

//
func (r *Handler) Accept(e p2p.Packet) {
	if !r.IsActive() {
		return
	}

	err := certifyMessage(e.Signal, r.cfg.MaximumMirrorSegments)
	if err != nil {
		if errors.Is(err, ErrSurpassesMaximumMirrorSegments) {
			r.aligner.DeclineNode(e.Src)
		}
		r.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		r.Router.HaltNodeForFault(e.Src, err)
		return
	}

	switch e.StreamUID {
	case MirrorStream:
		switch msg := e.Signal.(type) {
		case *statusproto.MirrorsQuery:
			mirrors, err := r.currentMirrors(currentMirrors)
			if err != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", err)
				return
			}
			for _, mirror := range mirrors {
				r.Tracer.Diagnose("REDACTED", "REDACTED", mirror.Level,
					"REDACTED", mirror.Layout, "REDACTED", e.Src.ID())
				e.Src.Transmit(p2p.Packet{
					StreamUID: e.StreamUID,
					Signal: &statusproto.MirrorsReply{
						Level:   mirror.Level,
						Layout:   mirror.Layout,
						Segments:   mirror.Segments,
						Digest:     mirror.Digest,
						Metainfo: mirror.Metainfo,
					},
				})
			}

		case *statusproto.MirrorsReply:
			r.mtx.RLock()
			defer r.mtx.RUnlock()
			if r.aligner == nil {
				r.Tracer.Diagnose("REDACTED")
				return
			}
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Layout, "REDACTED", e.Src.ID())
			_, err := r.aligner.AppendMirror(e.Src, &mirror{
				Level:   msg.Level,
				Layout:   msg.Layout,
				Segments:   msg.Segments,
				Digest:     msg.Digest,
				Metainfo: msg.Metainfo,
			})
			//
			if err != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Layout,
					"REDACTED", e.Src.ID(), "REDACTED", err)
				return
			}

		default:
			r.Tracer.Fault("REDACTED", msg)
		}

	case SegmentStream:
		switch msg := e.Signal.(type) {
		case *statusproto.SegmentQuery:
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Layout,
				"REDACTED", msg.Ordinal, "REDACTED", e.Src.ID())
			reply, err := r.link.ImportMirrorSegment(context.TODO(), &iface.QueryImportMirrorSegment{
				Level: msg.Level,
				Layout: msg.Layout,
				Segment:  msg.Ordinal,
			})
			if err != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Layout,
					"REDACTED", msg.Ordinal, "REDACTED", err)
				return
			}
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Layout,
				"REDACTED", msg.Ordinal, "REDACTED", e.Src.ID())
			e.Src.Transmit(p2p.Packet{
				StreamUID: SegmentStream,
				Signal: &statusproto.SegmentReply{
					Level:  msg.Level,
					Layout:  msg.Layout,
					Ordinal:   msg.Ordinal,
					Segment:   reply.Segment,
					Absent: reply.Segment == nil,
				},
			})

		case *statusproto.SegmentReply:
			r.mtx.RLock()
			defer r.mtx.RUnlock()
			if r.aligner == nil {
				r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src.ID())
				return
			}
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Layout,
				"REDACTED", msg.Ordinal, "REDACTED", e.Src.ID())
			_, err := r.aligner.AppendSegment(&segment{
				Level: msg.Level,
				Layout: msg.Layout,
				Ordinal:  msg.Ordinal,
				Segment:  msg.Segment,
				Emitter: e.Src.ID(),
			})
			if err != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Layout,
					"REDACTED", msg.Ordinal, "REDACTED", err)
				return
			}

		default:
			r.Tracer.Fault("REDACTED", msg)
		}

	default:
		r.Tracer.Fault("REDACTED", e.StreamUID)
	}
}

//
func (r *Handler) currentMirrors(n uint32) ([]*mirror, error) {
	reply, err := r.link.CatalogMirrors(context.TODO(), &iface.QueryCatalogMirrors{})
	if err != nil {
		return nil, err
	}
	sort.Slice(reply.Mirrors, func(i, j int) bool {
		a := reply.Mirrors[i]
		b := reply.Mirrors[j]
		switch {
		case a.Level > b.Level:
			return true
		case a.Level == b.Level && a.Layout > b.Layout:
			return true
		default:
			return false
		}
	})
	mirrors := make([]*mirror, 0, n)
	for i, s := range reply.Mirrors {
		if i >= currentMirrors {
			break
		}
		mirrors = append(mirrors, &mirror{
			Level:   s.Level,
			Layout:   s.Layout,
			Segments:   s.Segments,
			Digest:     s.Digest,
			Metainfo: s.Metainfo,
		})
	}
	return mirrors, nil
}

//
//
func (r *Handler) Align(statusSource StatusSource, detectionMoment time.Duration) (sm.Status, *kinds.Endorse, error) {
	r.mtx.Lock()
	if r.aligner != nil {
		r.mtx.Unlock()
		return sm.Status{}, nil, errors.New("REDACTED")
	}
	r.stats.Aligning.Set(1)
	r.aligner = newAligner(r.cfg, r.Tracer, r.link, r.linkInquire, statusSource, r.temporaryFolder)
	r.mtx.Unlock()

	callback := func() {
		r.Tracer.Diagnose("REDACTED")
		//

		r.Router.MulticastAsync(p2p.Packet{
			StreamUID: MirrorStream,
			Signal:   &statusproto.MirrorsQuery{},
		})
	}

	callback()

	status, endorse, err := r.aligner.AlignAny(detectionMoment, callback)

	r.mtx.Lock()
	r.aligner = nil
	r.stats.Aligning.Set(0)
	r.mtx.Unlock()
	return status, endorse, err
}
