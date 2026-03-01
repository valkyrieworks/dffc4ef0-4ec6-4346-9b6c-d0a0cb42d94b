package netp2p

import (
	"context"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/toolkits"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/stretchr/testify/require"
)

func VerifySchemeUUID(t *testing.T) {
	for _, tt := range []struct {
		conduit  byte
		anticipated string
	}{
		{conduit: 0x00, anticipated: "REDACTED"},
		{conduit: 0x01, anticipated: "REDACTED"},
		{conduit: 0x10, anticipated: "REDACTED"},
		{conduit: 0xff, anticipated: "REDACTED"},
	} {
		require.Equal(t, protocol.ID(tt.anticipated), SchemeUUID(tt.conduit))
	}
}

func VerifyInfluxFetch(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		var (
			ctx     = context.Background()
			schemaUUID = SchemeUUID(0xAA)
			channels   = toolkits.ObtainReleaseChannels(t, 2)
			machine1   = createVerifyMachine(t, channels[0], usingJournaling())
			machine2   = createVerifyMachine(t, channels[1], usingJournaling())
		)

		t.Cleanup(func() {
			machine2.Close()
			machine1.Close()
		})

		//
		require.NoError(t, machine2.Connect(ctx, machine1.LocationDetails()))

		fetchFault := make(chan error, 1)
		machine1.SetStreamHandler(schemaUUID, func(influx network.Stream) {
			defer influx.Close()

			_, err := InfluxFetch(influx)
			fetchFault <- err
		})

		//
		influx, err := machine2.NewStream(ctx, machine1.ID(), schemaUUID)
		require.NoError(t, err)
		t.Cleanup(func() { _ = influx.Close() })

		excessivelyBulkyHeadline := uinttobyteVaruint(MaximumInfluxExtent + 1)

		//
		_, err = influx.Write(excessivelyBulkyHeadline)
		require.NoError(t, err)
		require.NoError(t, influx.Close())

		//
		err = <-fetchFault

		require.Error(t, err)
		require.ErrorContains(t, err, "REDACTED")
	})
}
