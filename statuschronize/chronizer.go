package statuschronize

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	sschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/statuschronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	//
	segmentDeadline = 2 * time.Minute

	//
	//
	minimalExplorationMoment = 5 * time.Second
)

var (
	//
	faultCancel = errors.New("REDACTED")
	//
	faultReissueImage = errors.New("REDACTED")
	//
	faultDeclineImage = errors.New("REDACTED")
	//
	faultDeclineLayout = errors.New("REDACTED")
	//
	faultDeclineOriginator = errors.New("REDACTED")
	//
	faultValidateUnsuccessful = errors.New("REDACTED")
	//
	faultDeadline = errors.New("REDACTED")
	//
	faultNegativeImages = errors.New("REDACTED")
)

//
//
//
type chronizer struct {
	tracer        log.Tracer
	statusSupplier StatusSupplier
	link          delegate.PlatformLinkImage
	linkInquire     delegate.PlatformLinkInquire
	images     *imageHub
	transientPath       string
	segmentRetrievers int32
	reissueDeadline  time.Duration

	mtx    commitchronize.ReadwriteExclusion
	segments *segmentStaging
}

//
func freshChronizer(
	cfg settings.StatusChronizeSettings,
	tracer log.Tracer,
	link delegate.PlatformLinkImage,
	linkInquire delegate.PlatformLinkInquire,
	statusSupplier StatusSupplier,
	transientPath string,
) *chronizer {
	return &chronizer{
		tracer:        tracer,
		statusSupplier: statusSupplier,
		link:          link,
		linkInquire:     linkInquire,
		images:     freshImageHub(),
		transientPath:       transientPath,
		segmentRetrievers: cfg.SegmentRetrievers,
		reissueDeadline:  cfg.SegmentSolicitDeadline,
	}
}

//
//
func (s *chronizer) AppendSegment(segment *segment) (bool, error) {
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
		s.tracer.Diagnose("REDACTED", "REDACTED", segment.Altitude, "REDACTED", segment.Layout,
			"REDACTED", segment.Ordinal)
	} else {
		s.tracer.Diagnose("REDACTED", "REDACTED", segment.Altitude, "REDACTED", segment.Layout,
			"REDACTED", segment.Ordinal)
	}
	return appended, nil
}

//
//
func (s *chronizer) AppendImage(node p2p.Node, image *image) (bool, error) {
	appended, err := s.images.Add(node, image)
	if err != nil {
		return false, err
	}
	if appended {
		s.tracer.Details("REDACTED", "REDACTED", image.Altitude, "REDACTED", image.Layout,
			"REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))
	}
	return appended, nil
}

//
//
func (s *chronizer) AppendNode(node p2p.Node) {
	s.tracer.Diagnose("REDACTED", "REDACTED", node.ID())
	e := p2p.Wrapper{
		ConduitUUID: ImageConduit,
		Signal:   &sschema.ImagesSolicit{},
	}
	node.Transmit(e)
}

//
func (s *chronizer) DiscardNode(node p2p.Node) {
	s.tracer.Diagnose("REDACTED", "REDACTED", node.ID())
	s.images.DiscardNode(node.ID())
}

//
func (s *chronizer) DeclineNode(node p2p.Node) {
	s.tracer.Diagnose("REDACTED", "REDACTED", node.ID())
	s.images.DeclineNode(node.ID())
}

//
//
//
func (s *chronizer) ChronizeSome(explorationMoment time.Duration, reissueCallback func()) (sm.Status, *kinds.Endorse, error) {
	if explorationMoment != 0 && explorationMoment < minimalExplorationMoment {
		explorationMoment = 5 * minimalExplorationMoment
	}

	if explorationMoment > 0 {
		s.tracer.Details("REDACTED", "REDACTED", explorationMoment)
		time.Sleep(explorationMoment)
	}

	//
	//
	var (
		image *image
		segments   *segmentStaging
		err      error
	)
	for {
		//
		if image == nil {
			image = s.images.Optimal()
			segments = nil
		}
		if image == nil {
			if explorationMoment == 0 {
				return sm.Status{}, nil, faultNegativeImages
			}
			reissueCallback()
			s.tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", explorationMoment))
			time.Sleep(explorationMoment)
			continue
		}
		if segments == nil {
			segments, err = freshSegmentStaging(image, s.transientPath)
			if err != nil {
				return sm.Status{}, nil, fmt.Errorf("REDACTED", err)
			}
			defer segments.Shutdown() //
		}

		freshStatus, endorse, err := s.Chronize(image, segments)
		switch {
		case err == nil:
			return freshStatus, endorse, nil

		case errors.Is(err, faultCancel):
			return sm.Status{}, nil, err

		case errors.Is(err, faultReissueImage):
			segments.ReissueEvery()
			s.tracer.Details("REDACTED", "REDACTED", image.Altitude, "REDACTED", image.Layout,
				"REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))
			continue

		case errors.Is(err, faultDeadline):
			s.images.Decline(image)
			s.tracer.Failure("REDACTED",
				"REDACTED", image.Altitude, "REDACTED", image.Layout, "REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))

		case errors.Is(err, faultDeclineImage):
			s.images.Decline(image)
			s.tracer.Details("REDACTED", "REDACTED", image.Altitude, "REDACTED", image.Layout,
				"REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))

		case errors.Is(err, faultDeclineLayout):
			s.images.DeclineLayout(image.Layout)
			s.tracer.Details("REDACTED", "REDACTED", image.Layout)

		case errors.Is(err, faultDeclineOriginator):
			s.tracer.Details("REDACTED", "REDACTED", image.Altitude, "REDACTED", image.Layout,
				"REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))
			for _, node := range s.images.ObtainNodes(image) {
				s.images.DeclineNode(node.ID())
				s.tracer.Details("REDACTED", "REDACTED", node.ID())
			}

		case errors.Is(err, context.DeadlineExceeded):
			s.tracer.Details("REDACTED", "REDACTED", image.Altitude, "REDACTED", err)
			s.images.Decline(image)

		default:
			return sm.Status{}, nil, fmt.Errorf("REDACTED", err)
		}

		//
		err = segments.Shutdown()
		if err != nil {
			s.tracer.Failure("REDACTED", "REDACTED", err)
		}
		image = nil
		segments = nil
	}
}

//
//
func (s *chronizer) Chronize(image *image, segments *segmentStaging) (sm.Status, *kinds.Endorse, error) {
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

	hcontext, abort := context.WithTimeout(context.TODO(), 30*time.Second)
	defer abort()

	platformDigest, err := s.statusSupplier.PlatformDigest(hcontext, image.Altitude)
	if err != nil {
		s.tracer.Details("REDACTED", "REDACTED", err)
		if errors.Is(err, agile.FaultNegativeAttestors) {
			return sm.Status{}, nil, err
		}
		return sm.Status{}, nil, faultDeclineImage
	}
	image.reliablePlatformDigest = platformDigest

	//
	err = s.extendImage(image)
	if err != nil {
		return sm.Status{}, nil, err
	}

	//
	retrieveContext, abort := context.WithCancel(context.TODO())
	defer abort()
	for i := int32(0); i < s.segmentRetrievers; i++ {
		go s.retrieveSegments(retrieveContext, image, segments)
	}

	pcontext, preemptcancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer preemptcancel()

	//
	status, err := s.statusSupplier.Status(pcontext, image.Altitude)
	if err != nil {
		s.tracer.Details("REDACTED", "REDACTED", err)
		if errors.Is(err, agile.FaultNegativeAttestors) {
			return sm.Status{}, nil, err
		}
		return sm.Status{}, nil, faultDeclineImage
	}
	endorse, err := s.statusSupplier.Endorse(pcontext, image.Altitude)
	if err != nil {
		s.tracer.Details("REDACTED", "REDACTED", err)
		if errors.Is(err, agile.FaultNegativeAttestors) {
			return sm.Status{}, nil, err
		}
		return sm.Status{}, nil, faultDeclineImage
	}

	//
	err = s.executeSegments(segments)
	if err != nil {
		return sm.Status{}, nil, err
	}

	//
	if err := s.validatePlatform(image, status.Edition.Agreement.App); err != nil {
		return sm.Status{}, nil, err
	}

	//
	s.tracer.Details("REDACTED", "REDACTED", image.Altitude, "REDACTED", image.Layout,
		"REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))

	return status, endorse, nil
}

//
//
func (s *chronizer) extendImage(image *image) error {
	s.tracer.Details("REDACTED", "REDACTED", image.Altitude,
		"REDACTED", image.Layout, "REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))
	reply, err := s.link.ExtendImage(context.TODO(), &iface.SolicitExtendImage{
		Image: &iface.Image{
			Altitude:   image.Altitude,
			Layout:   image.Layout,
			Segments:   image.Segments,
			Digest:     image.Digest,
			Attributes: image.Attributes,
		},
		PlatformDigest: image.reliablePlatformDigest,
	})
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	switch reply.Outcome {
	case iface.Replyextendimage_EMBRACE:
		s.tracer.Details("REDACTED", "REDACTED", image.Altitude,
			"REDACTED", image.Layout, "REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))
		return nil
	case iface.Replyextendimage_CANCEL:
		return faultCancel
	case iface.Replyextendimage_DECLINE:
		return faultDeclineImage
	case iface.Replyextendimage_DECLINE_LAYOUT:
		return faultDeclineLayout
	case iface.Replyextendimage_DECLINE_ORIGINATOR:
		return faultDeclineOriginator
	default:
		return fmt.Errorf("REDACTED", reply.Outcome)
	}
}

//
//
func (s *chronizer) executeSegments(segments *segmentStaging) error {
	for {
		segment, err := segments.Following()
		if errors.Is(err, faultComplete) {
			return nil
		} else if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		reply, err := s.link.ExecuteImageSegment(context.TODO(), &iface.SolicitExecuteImageSegment{
			Ordinal:  segment.Ordinal,
			Segment:  segment.Segment,
			Originator: string(segment.Originator),
		})
		if err != nil {
			return fmt.Errorf("REDACTED", segment.Ordinal, err)
		}
		s.tracer.Details("REDACTED", "REDACTED", segment.Altitude,
			"REDACTED", segment.Layout, "REDACTED", segment.Ordinal, "REDACTED", segments.Extent())

		//
		for _, ordinal := range reply.RetrieveSegments {
			err := segments.Eject(ordinal)
			if err != nil {
				return fmt.Errorf("REDACTED", ordinal, err)
			}
		}

		//
		for _, originator := range reply.DeclineOriginators {
			if originator != "REDACTED" {
				s.images.DeclineNode(p2p.ID(originator))
				err := segments.EjectOriginator(p2p.ID(originator))
				if err != nil {
					return fmt.Errorf("REDACTED", err)
				}
			}
		}

		switch reply.Outcome {
		case iface.Replyapplyimagefragment_EMBRACE:
		case iface.Replyapplyimagefragment_CANCEL:
			return faultCancel
		case iface.Replyapplyimagefragment_REISSUE:
			segments.Reissue(segment.Ordinal)
		case iface.Replyapplyimagefragment_REISSUE_IMAGE:
			return faultReissueImage
		case iface.Replyapplyimagefragment_DECLINE_IMAGE:
			return faultDeclineImage
		default:
			return fmt.Errorf("REDACTED", reply.Outcome)
		}
	}
}

//
//
func (s *chronizer) retrieveSegments(ctx context.Context, image *image, segments *segmentStaging) {
	var (
		following  = true
		ordinal uint32
		err   error
	)

	for {
		if following {
			ordinal, err = segments.Assign()
			if errors.Is(err, faultComplete) {
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
				s.tracer.Failure("REDACTED", "REDACTED", err)
				return
			}
		}
		s.tracer.Details("REDACTED", "REDACTED", image.Altitude,
			"REDACTED", image.Layout, "REDACTED", ordinal, "REDACTED", segments.Extent())

		metronome := time.NewTicker(s.reissueDeadline)
		defer metronome.Stop()

		s.solicitSegment(image, ordinal)

		select {
		case <-segments.AwaitForeach(ordinal):
			following = true

		case <-metronome.C:
			following = false

		case <-ctx.Done():
			return
		}

		metronome.Stop()
	}
}

//
func (s *chronizer) solicitSegment(image *image, segment uint32) {
	node := s.images.ObtainNode(image)
	if node == nil {
		s.tracer.Failure("REDACTED", "REDACTED", image.Altitude,
			"REDACTED", image.Layout, "REDACTED", log.FreshIdleFormat("REDACTED", image.Digest))
		return
	}
	s.tracer.Diagnose("REDACTED", "REDACTED", image.Altitude,
		"REDACTED", image.Layout, "REDACTED", segment, "REDACTED", node.ID())
	node.Transmit(p2p.Wrapper{
		ConduitUUID: SegmentConduit,
		Signal: &sschema.SegmentSolicit{
			Altitude: image.Altitude,
			Layout: image.Layout,
			Ordinal:  segment,
		},
	})
}

//
func (s *chronizer) validatePlatform(image *image, platformEdition uint64) error {
	reply, err := s.linkInquire.Details(context.TODO(), delegate.SolicitDetails)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	//
	if reply.PlatformEdition != platformEdition {
		//
		//
		return fmt.Errorf("REDACTED",
			platformEdition, reply.PlatformEdition)
	}
	if !bytes.Equal(image.reliablePlatformDigest, reply.FinalLedgerPlatformDigest) {
		s.tracer.Failure("REDACTED",
			"REDACTED", fmt.Sprintf("REDACTED", image.reliablePlatformDigest),
			"REDACTED", fmt.Sprintf("REDACTED", reply.FinalLedgerPlatformDigest))
		return faultValidateUnsuccessful
	}
	if uint64(reply.FinalLedgerAltitude) != image.Altitude {
		s.tracer.Failure(
			"REDACTED",
			"REDACTED", image.Altitude,
			"REDACTED", reply.FinalLedgerAltitude,
		)
		return faultValidateUnsuccessful
	}

	s.tracer.Details("REDACTED", "REDACTED", image.Altitude, "REDACTED", log.FreshIdleFormat("REDACTED", image.reliablePlatformDigest))
	return nil
}
