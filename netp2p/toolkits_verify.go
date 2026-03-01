package netp2p

import (
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

func VerifyFreshPriorSerializedArtifact(t *testing.T) {
	//
	//
	reverberate := &kinds.SolicitReverberate{Signal: "REDACTED"}
	msg := &kinds.Solicit{
		Datum: &kinds.Solicit_Reverberate{Reverberate: reverberate},
	}

	//
	buffered := freshPriorSerializedArtifact(msg)

	//
	byzInitial, err := serializeSchema(msg)
	require.NoError(t, err)

	byzBuffered, err := serializeSchema(buffered)
	require.NoError(t, err)

	//
	//
	require.Equal(t, byzInitial, byzBuffered)

	//
	//
	reverberate.Signal = "REDACTED"

	byzInitial, err = serializeSchema(msg)
	require.NoError(t, err)

	//
	msg.Datum = nil
	require.Nil(t, buffered.Signal.(*kinds.Solicit).Datum)

	byzBuffered, err = serializeSchema(buffered)
	require.NoError(t, err)

	//
	//
	require.NotEqual(t, byzInitial, byzBuffered)
}

func VerifySchemaKindAlias(t *testing.T) {
	var (
		reverberateRequest       = &kinds.SolicitReverberate{Signal: "REDACTED"}
		reverberateRequestBuffered = freshPriorSerializedArtifact(reverberateRequest)
	)

	//
	for _, tt := range []struct {
		msg  proto.Message
		desire string
	}{
		{
			msg:  reverberateRequest,
			desire: "REDACTED",
		},
		{
			msg:  reverberateRequestBuffered,
			desire: "REDACTED",
		},
	} {
		got := schemaKindAlias(tt.msg)
		require.Equal(t, tt.desire, got)
	}
}
