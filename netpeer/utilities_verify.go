package netpeer

import (
	"testing"

	"github.com/valkyrieworks/iface/kinds"
	"github.com/stretchr/testify/require"
)

func VerifyNewPreSerializedSignal(t *testing.T) {
	//
	//
	reverberate := &kinds.QueryReverberate{Signal: "REDACTED"}
	msg := &kinds.Query{
		Item: &kinds.Query_Reverberate{Replicate: reverberate},
	}

	//
	stored := newPreSerializedSignal(msg)

	//
	bzSource, err := serializeSchema(msg)
	require.NoError(t, err)

	bzStored, err := serializeSchema(stored)
	require.NoError(t, err)

	//
	//
	require.Equal(t, bzSource, bzStored)

	//
	//
	reverberate.Signal = "REDACTED"

	bzSource, err = serializeSchema(msg)
	require.NoError(t, err)

	//
	msg.Item = nil
	require.Nil(t, stored.Signal.(*kinds.Query).Item)

	bzStored, err = serializeSchema(stored)
	require.NoError(t, err)

	//
	//
	require.NotEqual(t, bzSource, bzStored)
}
