package statuschronize

import (
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	sschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/statuschronize"
)

const (
	//
	imageSignalExtent = int(4e6)
	//
	segmentSignalExtent = int(16e6)
)

var (
	FaultSurpassesMaximumImageSegments = errors.New("REDACTED")
)

//
func certifySignal(pb proto.Message, maximumImageSegments uint32) error {
	if pb == nil {
		return errors.New("REDACTED")
	}
	switch msg := pb.(type) {
	case *sschema.SegmentSolicit:
		if msg.Altitude == 0 {
			return errors.New("REDACTED")
		}
	case *sschema.SegmentReply:
		if msg.Altitude == 0 {
			return errors.New("REDACTED")
		}
		if msg.Absent && len(msg.Segment) > 0 {
			return errors.New("REDACTED")
		}
		if !msg.Absent && msg.Segment == nil {
			return errors.New("REDACTED")
		}
	case *sschema.ImagesSolicit:
	case *sschema.ImagesReply:
		if msg.Altitude == 0 {
			return errors.New("REDACTED")
		}
		if len(msg.Digest) == 0 {
			return errors.New("REDACTED")
		}
		if msg.Segments == 0 {
			return errors.New("REDACTED")
		}
		if msg.Segments > maximumImageSegments {
			return fmt.Errorf("REDACTED", FaultSurpassesMaximumImageSegments, msg.Segments, maximumImageSegments)
		}
	default:
		return fmt.Errorf("REDACTED", msg)
	}
	return nil
}
