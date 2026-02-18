package statusconnect

import (
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/p2p"
	netpeersims "github.com/valkyrieworks/p2p/simulations"
	statusproto "github.com/valkyrieworks/schema/consensuscore/statusconnect"
	gatewaysims "github.com/valkyrieworks/gateway/simulations"
)

func Verifyhandler_Accept_Segmentrequest(t *testing.T) {
	verifyscenarios := map[string]struct {
		query        *statusproto.SegmentQuery
		segment          []byte
		anticipateReply *statusproto.SegmentReply
	}{
		"REDACTED": {
			&statusproto.SegmentQuery{Level: 1, Layout: 1, Ordinal: 1},
			[]byte{1, 2, 3},
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Segment: []byte{1, 2, 3}},
		},
		"REDACTED": {
			&statusproto.SegmentQuery{Level: 1, Layout: 1, Ordinal: 1},
			[]byte{},
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Segment: nil},
		},
		"REDACTED": {
			&statusproto.SegmentQuery{Level: 1, Layout: 1, Ordinal: 1},
			nil,
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Absent: true},
		},
	}

	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			//
			link := &gatewaysims.ApplicationLinkMirror{}
			link.On("REDACTED", mock.Anything, &iface.QueryImportMirrorSegment{
				Level: tc.query.Level,
				Layout: tc.query.Layout,
				Segment:  tc.query.Ordinal,
			}).Return(&iface.ReplyImportMirrorSegment{Segment: tc.segment}, nil)

			//
			node := &netpeersims.Node{}
			node.On("REDACTED").Return(p2p.ID("REDACTED"))
			var reply *statusproto.SegmentReply
			if tc.anticipateReply != nil {
				node.On("REDACTED", mock.MatchedBy(func(i any) bool {
					e, ok := i.(p2p.Packet)
					return ok && e.StreamUID == SegmentStream
				})).Run(func(args mock.Arguments) {
					e := args[0].(p2p.Packet)

					//
					bz, err := proto.Marshal(e.Signal)
					require.NoError(t, err)
					err = proto.Unmarshal(bz, e.Signal)
					require.NoError(t, err)
					reply = e.Signal.(*statusproto.SegmentReply)
				}).Return(true)
			}

			//
			cfg := settings.StandardStatusAlignSettings()
			r := NewHandler(*cfg, link, nil, NoopStats())
			err := r.Begin()
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := r.Halt(); err != nil {
					t.Error(err)
				}
			})

			r.Accept(p2p.Packet{
				StreamUID: SegmentStream,
				Src:       node,
				Signal:   tc.query,
			})
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, tc.anticipateReply, reply)

			link.AssertExpectations(t)
			node.AssertExpectations(t)
		})
	}
}

func Verifyhandler_Accept_Mirrorsrequest(t *testing.T) {
	verifyscenarios := map[string]struct {
		mirrors       []*iface.Mirror
		anticipateReplies []*statusproto.MirrorsReply
	}{
		"REDACTED": {nil, []*statusproto.MirrorsReply{}},
		"REDACTED": {
			[]*iface.Mirror{
				{Level: 1, Layout: 2, Segments: 7, Digest: []byte{1, 2}, Metainfo: []byte{1}},
				{Level: 2, Layout: 2, Segments: 7, Digest: []byte{2, 2}, Metainfo: []byte{2}},
				{Level: 3, Layout: 2, Segments: 7, Digest: []byte{3, 2}, Metainfo: []byte{3}},
				{Level: 1, Layout: 1, Segments: 7, Digest: []byte{1, 1}, Metainfo: []byte{4}},
				{Level: 2, Layout: 1, Segments: 7, Digest: []byte{2, 1}, Metainfo: []byte{5}},
				{Level: 3, Layout: 1, Segments: 7, Digest: []byte{3, 1}, Metainfo: []byte{6}},
				{Level: 1, Layout: 4, Segments: 7, Digest: []byte{1, 4}, Metainfo: []byte{7}},
				{Level: 2, Layout: 4, Segments: 7, Digest: []byte{2, 4}, Metainfo: []byte{8}},
				{Level: 3, Layout: 4, Segments: 7, Digest: []byte{3, 4}, Metainfo: []byte{9}},
				{Level: 1, Layout: 3, Segments: 7, Digest: []byte{1, 3}, Metainfo: []byte{10}},
				{Level: 2, Layout: 3, Segments: 7, Digest: []byte{2, 3}, Metainfo: []byte{11}},
				{Level: 3, Layout: 3, Segments: 7, Digest: []byte{3, 3}, Metainfo: []byte{12}},
			},
			[]*statusproto.MirrorsReply{
				{Level: 3, Layout: 4, Segments: 7, Digest: []byte{3, 4}, Metainfo: []byte{9}},
				{Level: 3, Layout: 3, Segments: 7, Digest: []byte{3, 3}, Metainfo: []byte{12}},
				{Level: 3, Layout: 2, Segments: 7, Digest: []byte{3, 2}, Metainfo: []byte{3}},
				{Level: 3, Layout: 1, Segments: 7, Digest: []byte{3, 1}, Metainfo: []byte{6}},
				{Level: 2, Layout: 4, Segments: 7, Digest: []byte{2, 4}, Metainfo: []byte{8}},
				{Level: 2, Layout: 3, Segments: 7, Digest: []byte{2, 3}, Metainfo: []byte{11}},
				{Level: 2, Layout: 2, Segments: 7, Digest: []byte{2, 2}, Metainfo: []byte{2}},
				{Level: 2, Layout: 1, Segments: 7, Digest: []byte{2, 1}, Metainfo: []byte{5}},
				{Level: 1, Layout: 4, Segments: 7, Digest: []byte{1, 4}, Metainfo: []byte{7}},
				{Level: 1, Layout: 3, Segments: 7, Digest: []byte{1, 3}, Metainfo: []byte{10}},
			},
		},
	}

	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			//
			link := &gatewaysims.ApplicationLinkMirror{}
			link.On("REDACTED", mock.Anything, &iface.QueryCatalogMirrors{}).Return(&iface.ReplyCatalogMirrors{
				Mirrors: tc.mirrors,
			}, nil)

			//
			replies := []*statusproto.MirrorsReply{}
			node := &netpeersims.Node{}
			if len(tc.anticipateReplies) > 0 {
				node.On("REDACTED").Return(p2p.ID("REDACTED"))
				node.On("REDACTED", mock.MatchedBy(func(i any) bool {
					e, ok := i.(p2p.Packet)
					return ok && e.StreamUID == MirrorStream
				})).Run(func(args mock.Arguments) {
					e := args[0].(p2p.Packet)

					//
					bz, err := proto.Marshal(e.Signal)
					require.NoError(t, err)
					err = proto.Unmarshal(bz, e.Signal)
					require.NoError(t, err)
					replies = append(replies, e.Signal.(*statusproto.MirrorsReply))
				}).Return(true)
			}

			//
			cfg := settings.StandardStatusAlignSettings()
			r := NewHandler(*cfg, link, nil, NoopStats())
			err := r.Begin()
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := r.Halt(); err != nil {
					t.Error(err)
				}
			})

			r.Accept(p2p.Packet{
				StreamUID: MirrorStream,
				Src:       node,
				Signal:   &statusproto.MirrorsQuery{},
			})
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, tc.anticipateReplies, replies)

			link.AssertExpectations(t)
			node.AssertExpectations(t)
		})
	}
}
