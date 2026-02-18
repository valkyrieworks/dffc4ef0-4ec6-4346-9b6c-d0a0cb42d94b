package statusconnect

import (
	"encoding/hex"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/p2p"
	statusproto "github.com/valkyrieworks/schema/consensuscore/statusconnect"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

func VerifyCertifyMessage(t *testing.T) {
	verifyscenarios := map[string]struct {
		msg   proto.Message
		sound bool
	}{
		"REDACTED":       {nil, false},
		"REDACTED": {&engineproto.Ledger{}, false},

		"REDACTED":    {&statusproto.SegmentQuery{Level: 1, Layout: 1, Ordinal: 1}, true},
		"REDACTED": {&statusproto.SegmentQuery{Level: 0, Layout: 1, Ordinal: 1}, false},
		"REDACTED": {&statusproto.SegmentQuery{Level: 1, Layout: 0, Ordinal: 1}, true},
		"REDACTED":  {&statusproto.SegmentQuery{Level: 1, Layout: 1, Ordinal: 0}, true},

		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Segment: []byte{1}},
			true,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 0, Layout: 1, Ordinal: 1, Segment: []byte{1}},
			false,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 0, Ordinal: 1, Segment: []byte{1}},
			true,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 0, Segment: []byte{1}},
			true,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Segment: []byte{}},
			true,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Segment: nil},
			false,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Absent: true},
			true,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Absent: true, Segment: []byte{}},
			true,
		},
		"REDACTED": {
			&statusproto.SegmentReply{Level: 1, Layout: 1, Ordinal: 1, Absent: true, Segment: []byte{1}},
			false,
		},

		"REDACTED": {&statusproto.MirrorsQuery{}, true},

		"REDACTED": {
			&statusproto.MirrorsReply{Level: 1, Layout: 1, Segments: 2, Digest: []byte{1}},
			true,
		},
		"REDACTED": {
			&statusproto.MirrorsReply{Level: 0, Layout: 1, Segments: 2, Digest: []byte{1}},
			false,
		},
		"REDACTED": {
			&statusproto.MirrorsReply{Level: 1, Layout: 0, Segments: 2, Digest: []byte{1}},
			true,
		},
		"REDACTED": {
			&statusproto.MirrorsReply{Level: 1, Layout: 1, Digest: []byte{1}},
			false,
		},
		"REDACTED": {
			&statusproto.MirrorsReply{Level: 1, Layout: 1, Segments: 2, Digest: []byte{}},
			false,
		},
		"REDACTED": {
			&statusproto.MirrorsReply{Level: 1, Layout: 1, Segments: 100001, Digest: []byte{1}},
			false,
		},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			err := certifyMessage(tc.msg, 100000)
			if tc.sound {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

//
func VerifyStatusAlignArrays(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &statusproto.MirrorsQuery{}, "REDACTED"},
		{"REDACTED", &statusproto.MirrorsReply{Level: 1, Layout: 2, Segments: 3, Digest: []byte("REDACTED"), Metainfo: []byte("REDACTED")}, "REDACTED"},
		{"REDACTED", &statusproto.SegmentQuery{Level: 1, Layout: 2, Ordinal: 3}, "REDACTED"},
		{"REDACTED", &statusproto.SegmentReply{Level: 1, Layout: 2, Ordinal: 3, Segment: []byte("REDACTED")}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		w := tc.msg.(p2p.Adapter).Enclose()
		bz, err := proto.Marshal(w)
		require.NoError(t, err)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyLabel)
	}
}
