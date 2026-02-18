package netpeer

import (
	"testing"

	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/stretchr/testify/require"
)

func VerifyProtocolUID(t *testing.T) {
	for _, tt := range []struct {
		conduit  byte
		anticipated string
	}{
		{conduit: 0x00, anticipated: "REDACTED"},
		{conduit: 0x01, anticipated: "REDACTED"},
		{conduit: 0x10, anticipated: "REDACTED"},
		{conduit: 0xff, anticipated: "REDACTED"},
	} {
		require.Equal(t, protocol.ID(tt.anticipated), ProtocolUID(tt.conduit))
	}
}
