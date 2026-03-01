package statuschronize

import (
	"encoding/hex"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	sschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/statuschronize"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

func VerifyCertifySignal(t *testing.T) {
	verifycases := map[string]struct {
		msg   proto.Message
		sound bool
	}{
		"REDACTED":       {nil, false},
		"REDACTED": {&commitchema.Ledger{}, false},

		"REDACTED":    {&sschema.SegmentSolicit{Altitude: 1, Layout: 1, Ordinal: 1}, true},
		"REDACTED": {&sschema.SegmentSolicit{Altitude: 0, Layout: 1, Ordinal: 1}, false},
		"REDACTED": {&sschema.SegmentSolicit{Altitude: 1, Layout: 0, Ordinal: 1}, true},
		"REDACTED":  {&sschema.SegmentSolicit{Altitude: 1, Layout: 1, Ordinal: 0}, true},

		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Segment: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 0, Layout: 1, Ordinal: 1, Segment: []byte{1}},
			false,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 0, Ordinal: 1, Segment: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 0, Segment: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Segment: []byte{}},
			true,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Segment: nil},
			false,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Absent: true},
			true,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Absent: true, Segment: []byte{}},
			true,
		},
		"REDACTED": {
			&sschema.SegmentReply{Altitude: 1, Layout: 1, Ordinal: 1, Absent: true, Segment: []byte{1}},
			false,
		},

		"REDACTED": {&sschema.ImagesSolicit{}, true},

		"REDACTED": {
			&sschema.ImagesReply{Altitude: 1, Layout: 1, Segments: 2, Digest: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.ImagesReply{Altitude: 0, Layout: 1, Segments: 2, Digest: []byte{1}},
			false,
		},
		"REDACTED": {
			&sschema.ImagesReply{Altitude: 1, Layout: 0, Segments: 2, Digest: []byte{1}},
			true,
		},
		"REDACTED": {
			&sschema.ImagesReply{Altitude: 1, Layout: 1, Digest: []byte{1}},
			false,
		},
		"REDACTED": {
			&sschema.ImagesReply{Altitude: 1, Layout: 1, Segments: 2, Digest: []byte{}},
			false,
		},
		"REDACTED": {
			&sschema.ImagesReply{Altitude: 1, Layout: 1, Segments: 100001, Digest: []byte{1}},
			false,
		},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			err := certifySignal(tc.msg, 100000)
			if tc.sound {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

//
func VerifyStatusChronizeArrays(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &sschema.ImagesSolicit{}, "REDACTED"},
		{"REDACTED", &sschema.ImagesReply{Altitude: 1, Layout: 2, Segments: 3, Digest: []byte("REDACTED"), Attributes: []byte("REDACTED")}, "REDACTED"},
		{"REDACTED", &sschema.SegmentSolicit{Altitude: 1, Layout: 2, Ordinal: 3}, "REDACTED"},
		{"REDACTED", &sschema.SegmentReply{Altitude: 1, Layout: 2, Ordinal: 3, Segment: []byte("REDACTED")}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		w := tc.msg.(p2p.Encapsulator).Enclose()
		bz, err := proto.Marshal(w)
		require.NoError(t, err)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyAlias)
	}
}
