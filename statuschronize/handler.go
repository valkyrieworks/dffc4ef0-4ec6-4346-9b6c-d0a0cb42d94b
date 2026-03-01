package statuschronize

import (
	"context"
	"errors"
	"sort"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	sschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/statuschronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	//
	ImageConduit = byte(0x60)
	//
	SegmentConduit = byte(0x61)
	//
	currentImages = 10
)

//
//
type Handler struct {
	p2p.FoundationHandler

	cfg       settings.StatusChronizeSettings
	link      delegate.PlatformLinkImage
	linkInquire delegate.PlatformLinkInquire
	transientPath   string
	telemetry   *Telemetry

	//
	//
	mtx    commitchronize.ReadwriteExclusion
	chronizer *chronizer
}

//
func FreshHandler(
	cfg settings.StatusChronizeSettings,
	link delegate.PlatformLinkImage,
	linkInquire delegate.PlatformLinkInquire,
	telemetry *Telemetry,
) *Handler {
	r := &Handler{
		cfg:       cfg,
		link:      link,
		linkInquire: linkInquire,
		telemetry:   telemetry,
	}
	r.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", r)

	return r
}

//
func (r *Handler) ObtainConduits() []*p2p.ConduitDefinition {
	return []*p2p.ConduitDefinition{
		{
			ID:                  ImageConduit,
			Urgency:            5,
			TransmitStagingVolume:   10,
			ObtainSignalVolume: imageSignalExtent,
			SignalKind:         &sschema.Signal{},
		},
		{
			ID:                  SegmentConduit,
			Urgency:            3,
			TransmitStagingVolume:   10,
			ObtainSignalVolume: segmentSignalExtent,
			SignalKind:         &sschema.Signal{},
		},
	}
}

//
func (r *Handler) UponInitiate() error {
	return nil
}

//
func (r *Handler) AppendNode(node p2p.Node) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.chronizer != nil {
		r.chronizer.AppendNode(node)
	}
}

//
func (r *Handler) DiscardNode(node p2p.Node, _ any) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.chronizer != nil {
		r.chronizer.DiscardNode(node)
	}
}

//
func (r *Handler) Accept(e p2p.Wrapper) {
	if !r.EqualsActive() {
		return
	}

	err := certifySignal(e.Signal, r.cfg.MaximumImageSegments)
	if err != nil {
		if errors.Is(err, FaultSurpassesMaximumImageSegments) {
			r.chronizer.DeclineNode(e.Src)
		}
		r.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		r.Router.HaltNodeForeachFailure(e.Src, err)
		return
	}

	switch e.ConduitUUID {
	case ImageConduit:
		switch msg := e.Signal.(type) {
		case *sschema.ImagesSolicit:
			images, err := r.currentImages(currentImages)
			if err != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", err)
				return
			}
			for _, image := range images {
				r.Tracer.Diagnose("REDACTED", "REDACTED", image.Altitude,
					"REDACTED", image.Layout, "REDACTED", e.Src.ID())
				e.Src.Transmit(p2p.Wrapper{
					ConduitUUID: e.ConduitUUID,
					Signal: &sschema.ImagesReply{
						Altitude:   image.Altitude,
						Layout:   image.Layout,
						Segments:   image.Segments,
						Digest:     image.Digest,
						Attributes: image.Attributes,
					},
				})
			}

		case *sschema.ImagesReply:
			r.mtx.RLock()
			defer r.mtx.RUnlock()
			if r.chronizer == nil {
				r.Tracer.Diagnose("REDACTED")
				return
			}
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Layout, "REDACTED", e.Src.ID())
			_, err := r.chronizer.AppendImage(e.Src, &image{
				Altitude:   msg.Altitude,
				Layout:   msg.Layout,
				Segments:   msg.Segments,
				Digest:     msg.Digest,
				Attributes: msg.Attributes,
			})
			//
			if err != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Layout,
					"REDACTED", e.Src.ID(), "REDACTED", err)
				return
			}

		default:
			r.Tracer.Failure("REDACTED", msg)
		}

	case SegmentConduit:
		switch msg := e.Signal.(type) {
		case *sschema.SegmentSolicit:
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Layout,
				"REDACTED", msg.Ordinal, "REDACTED", e.Src.ID())
			reply, err := r.link.FetchImageSegment(context.TODO(), &iface.SolicitFetchImageSegment{
				Altitude: msg.Altitude,
				Layout: msg.Layout,
				Segment:  msg.Ordinal,
			})
			if err != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Layout,
					"REDACTED", msg.Ordinal, "REDACTED", err)
				return
			}
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Layout,
				"REDACTED", msg.Ordinal, "REDACTED", e.Src.ID())
			e.Src.Transmit(p2p.Wrapper{
				ConduitUUID: SegmentConduit,
				Signal: &sschema.SegmentReply{
					Altitude:  msg.Altitude,
					Layout:  msg.Layout,
					Ordinal:   msg.Ordinal,
					Segment:   reply.Segment,
					Absent: reply.Segment == nil,
				},
			})

		case *sschema.SegmentReply:
			r.mtx.RLock()
			defer r.mtx.RUnlock()
			if r.chronizer == nil {
				r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src.ID())
				return
			}
			r.Tracer.Diagnose("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Layout,
				"REDACTED", msg.Ordinal, "REDACTED", e.Src.ID())
			_, err := r.chronizer.AppendSegment(&segment{
				Altitude: msg.Altitude,
				Layout: msg.Layout,
				Ordinal:  msg.Ordinal,
				Segment:  msg.Segment,
				Originator: e.Src.ID(),
			})
			if err != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Layout,
					"REDACTED", msg.Ordinal, "REDACTED", err)
				return
			}

		default:
			r.Tracer.Failure("REDACTED", msg)
		}

	default:
		r.Tracer.Failure("REDACTED", e.ConduitUUID)
	}
}

//
func (r *Handler) currentImages(n uint32) ([]*image, error) {
	reply, err := r.link.CollectionImages(context.TODO(), &iface.SolicitCollectionImages{})
	if err != nil {
		return nil, err
	}
	sort.Slice(reply.Images, func(i, j int) bool {
		a := reply.Images[i]
		b := reply.Images[j]
		switch {
		case a.Altitude > b.Altitude:
			return true
		case a.Altitude == b.Altitude && a.Layout > b.Layout:
			return true
		default:
			return false
		}
	})
	images := make([]*image, 0, n)
	for i, s := range reply.Images {
		if i >= currentImages {
			break
		}
		images = append(images, &image{
			Altitude:   s.Altitude,
			Layout:   s.Layout,
			Segments:   s.Segments,
			Digest:     s.Digest,
			Attributes: s.Attributes,
		})
	}
	return images, nil
}

//
//
func (r *Handler) Chronize(statusSupplier StatusSupplier, explorationMoment time.Duration) (sm.Status, *kinds.Endorse, error) {
	r.mtx.Lock()
	if r.chronizer != nil {
		r.mtx.Unlock()
		return sm.Status{}, nil, errors.New("REDACTED")
	}
	r.telemetry.Chronizing.Set(1)
	r.chronizer = freshChronizer(r.cfg, r.Tracer, r.link, r.linkInquire, statusSupplier, r.transientPath)
	r.mtx.Unlock()

	callback := func() {
		r.Tracer.Diagnose("REDACTED")
		//

		r.Router.MulticastAsyncronous(p2p.Wrapper{
			ConduitUUID: ImageConduit,
			Signal:   &sschema.ImagesSolicit{},
		})
	}

	callback()

	status, endorse, err := r.chronizer.ChronizeSome(explorationMoment, callback)

	r.mtx.Lock()
	r.chronizer = nil
	r.telemetry.Chronizing.Set(0)
	r.mtx.Unlock()
	return status, endorse, err
}
