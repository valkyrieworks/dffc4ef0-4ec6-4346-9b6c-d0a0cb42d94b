package statusconnect

import (
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	statusproto "github.com/valkyrieworks/schema/consensuscore/statusconnect"
)

const (
	//
	mirrorMessageVolume = int(4e6)
	//
	segmentMessageVolume = int(16e6)
)

var (
	ErrSurpassesMaximumMirrorSegments = errors.New("REDACTED")
)

//
func certifyMessage(pb proto.Message, maximumMirrorSegments uint32) error {
	if pb == nil {
		return errors.New("REDACTED")
	}
	switch msg := pb.(type) {
	case *statusproto.SegmentQuery:
		if msg.Level == 0 {
			return errors.New("REDACTED")
		}
	case *statusproto.SegmentReply:
		if msg.Level == 0 {
			return errors.New("REDACTED")
		}
		if msg.Absent && len(msg.Segment) > 0 {
			return errors.New("REDACTED")
		}
		if !msg.Absent && msg.Segment == nil {
			return errors.New("REDACTED")
		}
	case *statusproto.MirrorsQuery:
	case *statusproto.MirrorsReply:
		if msg.Level == 0 {
			return errors.New("REDACTED")
		}
		if len(msg.Digest) == 0 {
			return errors.New("REDACTED")
		}
		if msg.Segments == 0 {
			return errors.New("REDACTED")
		}
		if msg.Segments > maximumMirrorSegments {
			return fmt.Errorf("REDACTED", ErrSurpassesMaximumMirrorSegments, msg.Segments, maximumMirrorSegments)
		}
	default:
		return fmt.Errorf("REDACTED", msg)
	}
	return nil
}
