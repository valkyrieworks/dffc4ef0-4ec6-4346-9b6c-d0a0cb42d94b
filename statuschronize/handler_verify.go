package statuschronize

import (
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	netmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulations"
	sschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/statuschronize"
	delegatesimulate "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate/simulations"
)

func Reactortest_Accept_Fragmentsolicit(t *testing.T) {
	verifycases := map[string]struct {
		solicit        *sschema.SegmentSolicit
		segment          []byte
		anticipateReply *sschema.SegmentReply
	}{
		"REDACTED": {
			&sschema.SegmentSolicit{Altitude: 1, Layout: 1, Ordinal: 1},
			[]byte{1, 2, 3},
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Segment: []byte{1, 2, 3}},
		},
		"REDACTED": {
			&sschema.SegmentSolicit{Altitude: 1, Layout: 1, Ordinal: 1},
			[]byte{},
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Segment: nil},
		},
		"REDACTED": {
			&sschema.SegmentSolicit{Altitude: 1, Layout: 1, Ordinal: 1},
			nil,
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Absent: true},
		},
	}

	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			//
			link := &delegatesimulate.PlatformLinkImage{}
			link.On("REDACTED", mock.Anything, &iface.SolicitFetchImageSegment{
				Altitude: tc.solicit.Altitude,
				Layout: tc.solicit.Layout,
				Segment:  tc.solicit.Ordinal,
			}).Return(&iface.ReplyFetchImageSegment{Segment: tc.segment}, nil)

			//
			node := &netmocks.Node{}
			node.On("REDACTED").Return(p2p.ID("REDACTED"))
			var reply *sschema.SegmentReply
			if tc.anticipateReply != nil {
				node.On("REDACTED", mock.MatchedBy(func(i any) bool {
					e, ok := i.(p2p.Wrapper)
					return ok && e.ConduitUUID == SegmentConduit
				})).Run(func(arguments mock.Arguments) {
					e := arguments[0].(p2p.Wrapper)

					//
					bz, err := proto.Marshal(e.Signal)
					require.NoError(t, err)
					err = proto.Unmarshal(bz, e.Signal)
					require.NoError(t, err)
					reply = e.Signal.(*sschema.SegmentReply)
				}).Return(true)
			}

			//
			cfg := settings.FallbackStatusChronizeSettings()
			r := FreshHandler(*cfg, link, nil, NooperationTelemetry())
			err := r.Initiate()
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := r.Halt(); err != nil {
					t.Error(err)
				}
			})

			r.Accept(p2p.Wrapper{
				ConduitUUID: SegmentConduit,
				Src:       node,
				Signal:   tc.solicit,
			})
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, tc.anticipateReply, reply)

			link.AssertExpectations(t)
			node.AssertExpectations(t)
		})
	}
}

func Reactortest_Accept_Imagessolicit(t *testing.T) {
	verifycases := map[string]struct {
		images       []*iface.Image
		anticipateReplies []*sschema.ImagesReply
	}{
		"REDACTED": {nil, []*sschema.ImagesReply{}},
		"REDACTED": {
			[]*iface.Image{
				{Altitude: 1, Layout: 2, Segments: 7, Digest: []byte{1, 2}, Attributes: []byte{1}},
				{Altitude: 2, Layout: 2, Segments: 7, Digest: []byte{2, 2}, Attributes: []byte{2}},
				{Altitude: 3, Layout: 2, Segments: 7, Digest: []byte{3, 2}, Attributes: []byte{3}},
				{Altitude: 1, Layout: 1, Segments: 7, Digest: []byte{1, 1}, Attributes: []byte{4}},
				{Altitude: 2, Layout: 1, Segments: 7, Digest: []byte{2, 1}, Attributes: []byte{5}},
				{Altitude: 3, Layout: 1, Segments: 7, Digest: []byte{3, 1}, Attributes: []byte{6}},
				{Altitude: 1, Layout: 4, Segments: 7, Digest: []byte{1, 4}, Attributes: []byte{7}},
				{Altitude: 2, Layout: 4, Segments: 7, Digest: []byte{2, 4}, Attributes: []byte{8}},
				{Altitude: 3, Layout: 4, Segments: 7, Digest: []byte{3, 4}, Attributes: []byte{9}},
				{Altitude: 1, Layout: 3, Segments: 7, Digest: []byte{1, 3}, Attributes: []byte{10}},
				{Altitude: 2, Layout: 3, Segments: 7, Digest: []byte{2, 3}, Attributes: []byte{11}},
				{Altitude: 3, Layout: 3, Segments: 7, Digest: []byte{3, 3}, Attributes: []byte{12}},
			},
			[]*sschema.ImagesReply{
				{Altitude: 3, Layout: 4, Segments: 7, Digest: []byte{3, 4}, Attributes: []byte{9}},
				{Altitude: 3, Layout: 3, Segments: 7, Digest: []byte{3, 3}, Attributes: []byte{12}},
				{Altitude: 3, Layout: 2, Segments: 7, Digest: []byte{3, 2}, Attributes: []byte{3}},
				{Altitude: 3, Layout: 1, Segments: 7, Digest: []byte{3, 1}, Attributes: []byte{6}},
				{Altitude: 2, Layout: 4, Segments: 7, Digest: []byte{2, 4}, Attributes: []byte{8}},
				{Altitude: 2, Layout: 3, Segments: 7, Digest: []byte{2, 3}, Attributes: []byte{11}},
				{Altitude: 2, Layout: 2, Segments: 7, Digest: []byte{2, 2}, Attributes: []byte{2}},
				{Altitude: 2, Layout: 1, Segments: 7, Digest: []byte{2, 1}, Attributes: []byte{5}},
				{Altitude: 1, Layout: 4, Segments: 7, Digest: []byte{1, 4}, Attributes: []byte{7}},
				{Altitude: 1, Layout: 3, Segments: 7, Digest: []byte{1, 3}, Attributes: []byte{10}},
			},
		},
	}

	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			//
			link := &delegatesimulate.PlatformLinkImage{}
			link.On("REDACTED", mock.Anything, &iface.SolicitCollectionImages{}).Return(&iface.ReplyCatalogImages{
				Images: tc.images,
			}, nil)

			//
			replies := []*sschema.ImagesReply{}
			node := &netmocks.Node{}
			if len(tc.anticipateReplies) > 0 {
				node.On("REDACTED").Return(p2p.ID("REDACTED"))
				node.On("REDACTED", mock.MatchedBy(func(i any) bool {
					e, ok := i.(p2p.Wrapper)
					return ok && e.ConduitUUID == ImageConduit
				})).Run(func(arguments mock.Arguments) {
					e := arguments[0].(p2p.Wrapper)

					//
					bz, err := proto.Marshal(e.Signal)
					require.NoError(t, err)
					err = proto.Unmarshal(bz, e.Signal)
					require.NoError(t, err)
					replies = append(replies, e.Signal.(*sschema.ImagesReply))
				}).Return(true)
			}

			//
			cfg := settings.FallbackStatusChronizeSettings()
			r := FreshHandler(*cfg, link, nil, NooperationTelemetry())
			err := r.Initiate()
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := r.Halt(); err != nil {
					t.Error(err)
				}
			})

			r.Accept(p2p.Wrapper{
				ConduitUUID: ImageConduit,
				Src:       node,
				Signal:   &sschema.ImagesSolicit{},
			})
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, tc.anticipateReplies, replies)

			link.AssertExpectations(t)
			node.AssertExpectations(t)
		})
	}
}
