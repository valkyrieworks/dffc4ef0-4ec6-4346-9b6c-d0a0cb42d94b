package statereplication

import (
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	atci "github.com/valkyrieworks/atci/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/p2p"
	netmocks "github.com/valkyrieworks/p2p/simulations"
	sschema "github.com/valkyrieworks/schema/consensuscore/statereplication"
	gatwaymocks "github.com/valkyrieworks/gateway/simulations"
)

func TestReactor_Receive_ChunkRequest(t *testing.T) {
	testcases := map[string]struct {
		request        *sschema.ChunkRequest
		chunk          []byte
		expectResponse *sschema.ChunkResponse
	}{
		"REDACTED": {
			&sschema.ChunkRequest{Height: 1, Format: 1, Index: 1},
			[]byte{1, 2, 3},
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Chunk: []byte{1, 2, 3}},
		},
		"REDACTED": {
			&sschema.ChunkRequest{Height: 1, Format: 1, Index: 1},
			[]byte{},
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Chunk: nil},
		},
		"REDACTED": {
			&sschema.ChunkRequest{Height: 1, Format: 1, Index: 1},
			nil,
			&sschema.ChunkResponse{Height: 1, Format: 1, Index: 1, Missing: true},
		},
	}

	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			//
			conn := &gatwaymocks.AppConnSnapshot{}
			conn.On("REDACTED", mock.Anything, &atci.RequestLoadSnapshotChunk{
				Height: tc.request.Height,
				Format: tc.request.Format,
				Chunk:  tc.request.Index,
			}).Return(&atci.ResponseLoadSnapshotChunk{Chunk: tc.chunk}, nil)

			//
			peer := &netmocks.Peer{}
			peer.On("REDACTED").Return(p2p.ID("REDACTED"))
			var response *sschema.ChunkResponse
			if tc.expectResponse != nil {
				peer.On("REDACTED", mock.MatchedBy(func(i any) bool {
					e, ok := i.(p2p.Envelope)
					return ok && e.ChannelID == ChunkChannel
				})).Run(func(args mock.Arguments) {
					e := args[0].(p2p.Envelope)

					//
					bz, err := proto.Marshal(e.Message)
					require.NoError(t, err)
					err = proto.Unmarshal(bz, e.Message)
					require.NoError(t, err)
					response = e.Message.(*sschema.ChunkResponse)
				}).Return(true)
			}

			//
			cfg := settings.DefaultStateSyncConfig()
			r := NewReactor(*cfg, conn, nil, NopMetrics())
			err := r.Start()
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := r.Stop(); err != nil {
					t.Error(err)
				}
			})

			r.Receive(p2p.Envelope{
				ChannelID: ChunkChannel,
				Src:       peer,
				Message:   tc.request,
			})
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, tc.expectResponse, response)

			conn.AssertExpectations(t)
			peer.AssertExpectations(t)
		})
	}
}

func TestReactor_Receive_SnapshotsRequest(t *testing.T) {
	testcases := map[string]struct {
		snapshots       []*atci.Snapshot
		expectResponses []*sschema.SnapshotsResponse
	}{
		"REDACTED": {nil, []*sschema.SnapshotsResponse{}},
		"REDACTED": {
			[]*atci.Snapshot{
				{Height: 1, Format: 2, Chunks: 7, Hash: []byte{1, 2}, Metadata: []byte{1}},
				{Height: 2, Format: 2, Chunks: 7, Hash: []byte{2, 2}, Metadata: []byte{2}},
				{Height: 3, Format: 2, Chunks: 7, Hash: []byte{3, 2}, Metadata: []byte{3}},
				{Height: 1, Format: 1, Chunks: 7, Hash: []byte{1, 1}, Metadata: []byte{4}},
				{Height: 2, Format: 1, Chunks: 7, Hash: []byte{2, 1}, Metadata: []byte{5}},
				{Height: 3, Format: 1, Chunks: 7, Hash: []byte{3, 1}, Metadata: []byte{6}},
				{Height: 1, Format: 4, Chunks: 7, Hash: []byte{1, 4}, Metadata: []byte{7}},
				{Height: 2, Format: 4, Chunks: 7, Hash: []byte{2, 4}, Metadata: []byte{8}},
				{Height: 3, Format: 4, Chunks: 7, Hash: []byte{3, 4}, Metadata: []byte{9}},
				{Height: 1, Format: 3, Chunks: 7, Hash: []byte{1, 3}, Metadata: []byte{10}},
				{Height: 2, Format: 3, Chunks: 7, Hash: []byte{2, 3}, Metadata: []byte{11}},
				{Height: 3, Format: 3, Chunks: 7, Hash: []byte{3, 3}, Metadata: []byte{12}},
			},
			[]*sschema.SnapshotsResponse{
				{Height: 3, Format: 4, Chunks: 7, Hash: []byte{3, 4}, Metadata: []byte{9}},
				{Height: 3, Format: 3, Chunks: 7, Hash: []byte{3, 3}, Metadata: []byte{12}},
				{Height: 3, Format: 2, Chunks: 7, Hash: []byte{3, 2}, Metadata: []byte{3}},
				{Height: 3, Format: 1, Chunks: 7, Hash: []byte{3, 1}, Metadata: []byte{6}},
				{Height: 2, Format: 4, Chunks: 7, Hash: []byte{2, 4}, Metadata: []byte{8}},
				{Height: 2, Format: 3, Chunks: 7, Hash: []byte{2, 3}, Metadata: []byte{11}},
				{Height: 2, Format: 2, Chunks: 7, Hash: []byte{2, 2}, Metadata: []byte{2}},
				{Height: 2, Format: 1, Chunks: 7, Hash: []byte{2, 1}, Metadata: []byte{5}},
				{Height: 1, Format: 4, Chunks: 7, Hash: []byte{1, 4}, Metadata: []byte{7}},
				{Height: 1, Format: 3, Chunks: 7, Hash: []byte{1, 3}, Metadata: []byte{10}},
			},
		},
	}

	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			//
			conn := &gatwaymocks.AppConnSnapshot{}
			conn.On("REDACTED", mock.Anything, &atci.RequestListSnapshots{}).Return(&atci.ResponseListSnapshots{
				Snapshots: tc.snapshots,
			}, nil)

			//
			responses := []*sschema.SnapshotsResponse{}
			peer := &netmocks.Peer{}
			if len(tc.expectResponses) > 0 {
				peer.On("REDACTED").Return(p2p.ID("REDACTED"))
				peer.On("REDACTED", mock.MatchedBy(func(i any) bool {
					e, ok := i.(p2p.Envelope)
					return ok && e.ChannelID == SnapshotChannel
				})).Run(func(args mock.Arguments) {
					e := args[0].(p2p.Envelope)

					//
					bz, err := proto.Marshal(e.Message)
					require.NoError(t, err)
					err = proto.Unmarshal(bz, e.Message)
					require.NoError(t, err)
					responses = append(responses, e.Message.(*sschema.SnapshotsResponse))
				}).Return(true)
			}

			//
			cfg := settings.DefaultStateSyncConfig()
			r := NewReactor(*cfg, conn, nil, NopMetrics())
			err := r.Start()
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := r.Stop(); err != nil {
					t.Error(err)
				}
			})

			r.Receive(p2p.Envelope{
				ChannelID: SnapshotChannel,
				Src:       peer,
				Message:   &sschema.SnapshotsRequest{},
			})
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, tc.expectResponses, responses)

			conn.AssertExpectations(t)
			peer.AssertExpectations(t)
		})
	}
}
