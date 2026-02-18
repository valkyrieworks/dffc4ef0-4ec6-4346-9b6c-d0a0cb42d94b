package statereplication

import (
	"errors"
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	sschema "github.com/valkyrieworks/schema/consensuscore/statereplication"
)

const (
	//
	snapshotMsgSize = int(4e6)
	//
	chunkMsgSize = int(16e6)
)

var (
	ErrExceedsMaxSnapshotChunks = errors.New("REDACTED")
)

//
func validateMsg(pb proto.Message, maxSnapshotChunks uint32) error {
	if pb == nil {
		return errors.New("REDACTED")
	}
	switch msg := pb.(type) {
	case *sschema.ChunkRequest:
		if msg.Height == 0 {
			return errors.New("REDACTED")
		}
	case *sschema.ChunkResponse:
		if msg.Height == 0 {
			return errors.New("REDACTED")
		}
		if msg.Missing && len(msg.Chunk) > 0 {
			return errors.New("REDACTED")
		}
		if !msg.Missing && msg.Chunk == nil {
			return errors.New("REDACTED")
		}
	case *sschema.SnapshotsRequest:
	case *sschema.SnapshotsResponse:
		if msg.Height == 0 {
			return errors.New("REDACTED")
		}
		if len(msg.Hash) == 0 {
			return errors.New("REDACTED")
		}
		if msg.Chunks == 0 {
			return errors.New("REDACTED")
		}
		if msg.Chunks > maxSnapshotChunks {
			return fmt.Errorf("REDACTED", ErrExceedsMaxSnapshotChunks, msg.Chunks, maxSnapshotChunks)
		}
	default:
		return fmt.Errorf("REDACTED", msg)
	}
	return nil
}
