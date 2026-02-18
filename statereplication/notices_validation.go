package statereplication

import (
	"encoding/hex"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/p2p"
	sschema "github.com/valkyrieworks/schema/consensuscore/statereplication"
	ctschema "github.com/valkyrieworks/schema/consensuscore/kinds"
)

func TestValidateMsg(t *testing.T) {
	testcases := map[string]struct {
		msg   proto.Message
		valid bool
	}{
		"REDACTED":       {nil, false},
		"REDACTED": {&ctschema.Block{}, false},

		"REDACTED":    {&sschema.ChunkRequest{Height: 1, Format: 1, Index: 1}, true},
		"REDACTED": {&sschema.ChunkRequest{Height: 0, Format: 1, Index: 1}, false},
		"REDACTED": {&sschema.ChunkRequest{Height: 1, Format: 0, Index: 1}, true},
		"REDACTED":  {&sschema.ChunkRequest{Height: 1, Format: 1, Index: 0}, true},

		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Chunk: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 0, Format: 1, Index: 1, Chunk: []byte{1}},
			false,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 0, Index: 1, Chunk: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 0, Chunk: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Chunk: []byte{}},
			true,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Chunk: nil},
			false,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Missing: true},
			true,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Missing: true, Chunk: []byte{}},
			true,
		},
		"REDACTED": {
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Missing: true, Chunk: []byte{1}},
			false,
		},

		"REDACTED": {&sschema.SnapshotsRequest{}, true},

		"REDACTED": {
			&sschema.SnapshotsResponse{Height: 1, Format: 1, Chunks: 2, Hash: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.SnapshotsResponse{Height: 0, Format: 1, Chunks: 2, Hash: []byte{1}},
			false,
		},
		"REDACTED": {
			&sschema.SnapshotsResponse{Height: 1, Format: 0, Chunks: 2, Hash: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.SnapshotsResponse{Height: 1, Format: 1, Hash: []byte{1}},
			false,
		},
		"REDACTED": {
			&sschema.SnapshotsResponse{Height: 1, Format: 1, Chunks: 2, Hash: []byte{}},
			false,
		},
		"REDACTED": {
			&sschema.SnapshotsResponse{Height: 1, Format: 1, Chunks: 100001, Hash: []byte{1}},
			false,
		},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			err := validateMsg(tc.msg, 100000)
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

//
func TestStateSyncVectors(t *testing.T) {
	testCases := []struct {
		testName string
		msg      proto.Message
		expBytes string
	}{
		{"REDACTED", &sschema.SnapshotsRequest{}, "REDACTED"},
		{"REDACTED", &sschema.SnapshotsResponse{Height: 1, Format: 2, Chunks: 3, Hash: []byte("REDACTED"), Metadata: []byte("REDACTED")}, "REDACTED"},
		{"REDACTED", &sschema.ChunkRequest{Height: 1, Format: 2, Index: 3}, "REDACTED"},
		{"REDACTED", &sschema.ChunkResponse{Height: 1, Format: 2, Index: 3, Chunk: []byte("REDACTED")}, "REDACTED"},
	}

	for _, tc := range testCases {

		w := tc.msg.(p2p.Wrapper).Wrap()
		bz, err := proto.Marshal(w)
		require.NoError(t, err)

		require.Equal(t, tc.expBytes, hex.EncodeToString(bz), tc.testName)
	}
}
