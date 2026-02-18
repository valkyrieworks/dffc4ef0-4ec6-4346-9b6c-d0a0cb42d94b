package statusconnect

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/rapid"
	"github.com/valkyrieworks/p2p"
	statusproto "github.com/valkyrieworks/schema/consensuscore/statusconnect"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	segmentDeadline = 2 * time.Minute

	//
	//
	smallestDetectionMoment = 5 * time.Second
)

var (
	//
	errCancel = errors.New("REDACTED")
	//
	errReprocessMirror = errors.New("REDACTED")
	//
	errDeclineMirror = errors.New("REDACTED")
	//
	errDeclineLayout = errors.New("REDACTED")
	//
	errDeclineEmitter = errors.New("REDACTED")
	//
	errValidateErrored = errors.New("REDACTED")
	//
	errDeadline = errors.New("REDACTED")
	//
	errNoMirrors = errors.New("REDACTED")
)

//
//
//
type aligner struct {
	tracer        log.Tracer
	statusSource StatusSource
	link          gateway.ApplicationLinkMirror
	linkInquire     gateway.ApplicationLinkInquire
	mirrors     *mirrorDepository
	temporaryFolder       string
	segmentAcquirers int32
	reprocessDeadline  time.Duration

	mtx    engineconnect.ReadwriteLock
	segments *segmentBuffer
}

//
func newAligner(
	cfg settings.StatusAlignSettings,
	tracer log.Tracer,
	link gateway.ApplicationLinkMirror,
	linkInquire gateway.ApplicationLinkInquire,
	statusSource StatusSource,
	temporaryFolder string,
) *aligner {
	return &aligner{
		tracer:        tracer,
		statusSource: statusSource,
		link:          link,
		linkInquire:     linkInquire,
		mirrors:     newMirrorDepository(),
		temporaryFolder:       temporaryFolder,
		segmentAcquirers: cfg.SegmentAcquirers,
		reprocessDeadline:  cfg.SegmentQueryDeadline,
	}
}

//
//
func (s *aligner) AppendSegment(segment *segment) (bool, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s.segments == nil {
		return false, errors.New("REDACTED")
	}
	appended, err := s.segments.Add(segment)
	if err != nil {
		return false, err
	}
	if appended {
		s.tracer.Diagnose("REDACTED", "REDACTED", segment.Level, "REDACTED", segment.Layout,
			"REDACTED", segment.Ordinal)
	} else {
		s.tracer.Diagnose("REDACTED", "REDACTED", segment.Level, "REDACTED", segment.Layout,
			"REDACTED", segment.Ordinal)
	}
	return appended, nil
}

//
//
func (s *aligner) AppendMirror(node p2p.Node, mirror *mirror) (bool, error) {
	appended, err := s.mirrors.Add(node, mirror)
	if err != nil {
		return false, err
	}
	if appended {
		s.tracer.Details("REDACTED", "REDACTED", mirror.Level, "REDACTED", mirror.Layout,
			"REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))
	}
	return appended, nil
}

//
//
func (s *aligner) AppendNode(node p2p.Node) {
	s.tracer.Diagnose("REDACTED", "REDACTED", node.ID())
	e := p2p.Packet{
		StreamUID: MirrorStream,
		Signal:   &statusproto.MirrorsQuery{},
	}
	node.Transmit(e)
}

//
func (s *aligner) DeleteNode(node p2p.Node) {
	s.tracer.Diagnose("REDACTED", "REDACTED", node.ID())
	s.mirrors.DeleteNode(node.ID())
}

//
func (s *aligner) DeclineNode(node p2p.Node) {
	s.tracer.Diagnose("REDACTED", "REDACTED", node.ID())
	s.mirrors.DeclineNode(node.ID())
}

//
//
//
func (s *aligner) AlignAny(detectionMoment time.Duration, reprocessCallback func()) (sm.Status, *kinds.Endorse, error) {
	if detectionMoment != 0 && detectionMoment < smallestDetectionMoment {
		detectionMoment = 5 * smallestDetectionMoment
	}

	if detectionMoment > 0 {
		s.tracer.Details("REDACTED", "REDACTED", detectionMoment)
		time.Sleep(detectionMoment)
	}

	//
	//
	var (
		mirror *mirror
		segments   *segmentBuffer
		err      error
	)
	for {
		//
		if mirror == nil {
			mirror = s.mirrors.Optimal()
			segments = nil
		}
		if mirror == nil {
			if detectionMoment == 0 {
				return sm.Status{}, nil, errNoMirrors
			}
			reprocessCallback()
			s.tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", detectionMoment))
			time.Sleep(detectionMoment)
			continue
		}
		if segments == nil {
			segments, err = newSegmentBuffer(mirror, s.temporaryFolder)
			if err != nil {
				return sm.Status{}, nil, fmt.Errorf("REDACTED", err)
			}
			defer segments.End() //
		}

		newStatus, endorse, err := s.Align(mirror, segments)
		switch {
		case err == nil:
			return newStatus, endorse, nil

		case errors.Is(err, errCancel):
			return sm.Status{}, nil, err

		case errors.Is(err, errReprocessMirror):
			segments.ReprocessAll()
			s.tracer.Details("REDACTED", "REDACTED", mirror.Level, "REDACTED", mirror.Layout,
				"REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))
			continue

		case errors.Is(err, errDeadline):
			s.mirrors.Decline(mirror)
			s.tracer.Fault("REDACTED",
				"REDACTED", mirror.Level, "REDACTED", mirror.Layout, "REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))

		case errors.Is(err, errDeclineMirror):
			s.mirrors.Decline(mirror)
			s.tracer.Details("REDACTED", "REDACTED", mirror.Level, "REDACTED", mirror.Layout,
				"REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))

		case errors.Is(err, errDeclineLayout):
			s.mirrors.DeclineLayout(mirror.Layout)
			s.tracer.Details("REDACTED", "REDACTED", mirror.Layout)

		case errors.Is(err, errDeclineEmitter):
			s.tracer.Details("REDACTED", "REDACTED", mirror.Level, "REDACTED", mirror.Layout,
				"REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))
			for _, node := range s.mirrors.FetchNodes(mirror) {
				s.mirrors.DeclineNode(node.ID())
				s.tracer.Details("REDACTED", "REDACTED", node.ID())
			}

		case errors.Is(err, context.DeadlineExceeded):
			s.tracer.Details("REDACTED", "REDACTED", mirror.Level, "REDACTED", err)
			s.mirrors.Decline(mirror)

		default:
			return sm.Status{}, nil, fmt.Errorf("REDACTED", err)
		}

		//
		err = segments.End()
		if err != nil {
			s.tracer.Fault("REDACTED", "REDACTED", err)
		}
		mirror = nil
		segments = nil
	}
}

//
//
func (s *aligner) Align(mirror *mirror, segments *segmentBuffer) (sm.Status, *kinds.Endorse, error) {
	s.mtx.Lock()
	if s.segments != nil {
		s.mtx.Unlock()
		return sm.Status{}, nil, errors.New("REDACTED")
	}
	s.segments = segments
	s.mtx.Unlock()
	defer func() {
		s.mtx.Lock()
		s.segments = nil
		s.mtx.Unlock()
	}()

	hcontext, revoke := context.WithTimeout(context.TODO(), 30*time.Second)
	defer revoke()

	applicationDigest, err := s.statusSource.ApplicationDigest(hcontext, mirror.Level)
	if err != nil {
		s.tracer.Details("REDACTED", "REDACTED", err)
		if errors.Is(err, rapid.ErrNoAttestors) {
			return sm.Status{}, nil, err
		}
		return sm.Status{}, nil, errDeclineMirror
	}
	mirror.validatedApplicationDigest = applicationDigest

	//
	err = s.proposalMirror(mirror)
	if err != nil {
		return sm.Status{}, nil, err
	}

	//
	acquireCtx, revoke := context.WithCancel(context.TODO())
	defer revoke()
	for i := int32(0); i < s.segmentAcquirers; i++ {
		go s.acquireSegments(acquireCtx, mirror, segments)
	}

	pcontext, pabort := context.WithTimeout(context.TODO(), 30*time.Second)
	defer pabort()

	//
	status, err := s.statusSource.Status(pcontext, mirror.Level)
	if err != nil {
		s.tracer.Details("REDACTED", "REDACTED", err)
		if errors.Is(err, rapid.ErrNoAttestors) {
			return sm.Status{}, nil, err
		}
		return sm.Status{}, nil, errDeclineMirror
	}
	endorse, err := s.statusSource.Endorse(pcontext, mirror.Level)
	if err != nil {
		s.tracer.Details("REDACTED", "REDACTED", err)
		if errors.Is(err, rapid.ErrNoAttestors) {
			return sm.Status{}, nil, err
		}
		return sm.Status{}, nil, errDeclineMirror
	}

	//
	err = s.executeSegments(segments)
	if err != nil {
		return sm.Status{}, nil, err
	}

	//
	if err := s.validateApplication(mirror, status.Release.Agreement.App); err != nil {
		return sm.Status{}, nil, err
	}

	//
	s.tracer.Details("REDACTED", "REDACTED", mirror.Level, "REDACTED", mirror.Layout,
		"REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))

	return status, endorse, nil
}

//
//
func (s *aligner) proposalMirror(mirror *mirror) error {
	s.tracer.Details("REDACTED", "REDACTED", mirror.Level,
		"REDACTED", mirror.Layout, "REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))
	reply, err := s.link.ProposalMirror(context.TODO(), &iface.QueryProposalMirror{
		Mirror: &iface.Mirror{
			Level:   mirror.Level,
			Layout:   mirror.Layout,
			Segments:   mirror.Segments,
			Digest:     mirror.Digest,
			Metainfo: mirror.Metainfo,
		},
		ApplicationDigest: mirror.validatedApplicationDigest,
	})
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	switch reply.Outcome {
	case iface.Replymirrorsnapshot_ALLOW:
		s.tracer.Details("REDACTED", "REDACTED", mirror.Level,
			"REDACTED", mirror.Layout, "REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))
		return nil
	case iface.Replymirrorsnapshot_CANCEL:
		return errCancel
	case iface.Replymirrorsnapshot_DECLINE:
		return errDeclineMirror
	case iface.Replymirrorsnapshot_DECLINE_LAYOUT:
		return errDeclineLayout
	case iface.Replymirrorsnapshot_DECLINE_EMITTER:
		return errDeclineEmitter
	default:
		return fmt.Errorf("REDACTED", reply.Outcome)
	}
}

//
//
func (s *aligner) executeSegments(segments *segmentBuffer) error {
	for {
		segment, err := segments.Following()
		if errors.Is(err, errDone) {
			return nil
		} else if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		reply, err := s.link.ExecuteMirrorSegment(context.TODO(), &iface.QueryExecuteMirrorSegment{
			Ordinal:  segment.Ordinal,
			Segment:  segment.Segment,
			Emitter: string(segment.Emitter),
		})
		if err != nil {
			return fmt.Errorf("REDACTED", segment.Ordinal, err)
		}
		s.tracer.Details("REDACTED", "REDACTED", segment.Level,
			"REDACTED", segment.Layout, "REDACTED", segment.Ordinal, "REDACTED", segments.Volume())

		//
		for _, ordinal := range reply.ReacquireSegments {
			err := segments.Drop(ordinal)
			if err != nil {
				return fmt.Errorf("REDACTED", ordinal, err)
			}
		}

		//
		for _, emitter := range reply.DeclineEmitters {
			if emitter != "REDACTED" {
				s.mirrors.DeclineNode(p2p.ID(emitter))
				err := segments.DropEmitter(p2p.ID(emitter))
				if err != nil {
					return fmt.Errorf("REDACTED", err)
				}
			}
		}

		switch reply.Outcome {
		case iface.Replyexecutemirrorsegment_ALLOW:
		case iface.Replyexecutemirrorsegment_CANCEL:
			return errCancel
		case iface.Replyexecutemirrorsegment_REPROCESS:
			segments.Reprocess(segment.Ordinal)
		case iface.Replyexecutemirrorsegment_REPROCESS_MIRROR:
			return errReprocessMirror
		case iface.Replyexecutemirrorsegment_DECLINE_MIRROR:
			return errDeclineMirror
		default:
			return fmt.Errorf("REDACTED", reply.Outcome)
		}
	}
}

//
//
func (s *aligner) acquireSegments(ctx context.Context, mirror *mirror, segments *segmentBuffer) {
	var (
		following  = true
		ordinal uint32
		err   error
	)

	for {
		if following {
			ordinal, err = segments.Assign()
			if errors.Is(err, errDone) {
				//
				//
				select {
				case <-ctx.Done():
					return
				default:
				}
				time.Sleep(2 * time.Second)
				continue
			}
			if err != nil {
				s.tracer.Fault("REDACTED", "REDACTED", err)
				return
			}
		}
		s.tracer.Details("REDACTED", "REDACTED", mirror.Level,
			"REDACTED", mirror.Layout, "REDACTED", ordinal, "REDACTED", segments.Volume())

		timer := time.NewTicker(s.reprocessDeadline)
		defer timer.Stop()

		s.querySegment(mirror, ordinal)

		select {
		case <-segments.WaitFor(ordinal):
			following = true

		case <-timer.C:
			following = false

		case <-ctx.Done():
			return
		}

		timer.Stop()
	}
}

//
func (s *aligner) querySegment(mirror *mirror, segment uint32) {
	node := s.mirrors.FetchNode(mirror)
	if node == nil {
		s.tracer.Fault("REDACTED", "REDACTED", mirror.Level,
			"REDACTED", mirror.Layout, "REDACTED", log.NewIdleFormat("REDACTED", mirror.Digest))
		return
	}
	s.tracer.Diagnose("REDACTED", "REDACTED", mirror.Level,
		"REDACTED", mirror.Layout, "REDACTED", segment, "REDACTED", node.ID())
	node.Transmit(p2p.Packet{
		StreamUID: SegmentStream,
		Signal: &statusproto.SegmentQuery{
			Level: mirror.Level,
			Layout: mirror.Layout,
			Ordinal:  segment,
		},
	})
}

//
func (s *aligner) validateApplication(mirror *mirror, applicationRelease uint64) error {
	reply, err := s.linkInquire.Details(context.TODO(), gateway.QueryDetails)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	//
	if reply.ApplicationRelease != applicationRelease {
		//
		//
		return fmt.Errorf("REDACTED",
			applicationRelease, reply.ApplicationRelease)
	}
	if !bytes.Equal(mirror.validatedApplicationDigest, reply.FinalLedgerApplicationDigest) {
		s.tracer.Fault("REDACTED",
			"REDACTED", fmt.Sprintf("REDACTED", mirror.validatedApplicationDigest),
			"REDACTED", fmt.Sprintf("REDACTED", reply.FinalLedgerApplicationDigest))
		return errValidateErrored
	}
	if uint64(reply.FinalLedgerLevel) != mirror.Level {
		s.tracer.Fault(
			"REDACTED",
			"REDACTED", mirror.Level,
			"REDACTED", reply.FinalLedgerLevel,
		)
		return errValidateErrored
	}

	s.tracer.Details("REDACTED", "REDACTED", mirror.Level, "REDACTED", log.NewIdleFormat("REDACTED", mirror.validatedApplicationDigest))
	return nil
}
