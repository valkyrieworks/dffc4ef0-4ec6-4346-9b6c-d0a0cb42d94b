package netp2p

import (
	"net"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func VerifyLocatorTowardVariedLocation(t *testing.T) {
	for _, tt := range []struct {
		alias        string
		location        string
		carrier   string
		desire        string
		faultIncludes string
	}{
		{
			alias:      "REDACTED",
			location:      "REDACTED",
			carrier: CarrierQuicprotocol,
			desire:      "REDACTED",
		},
		{
			alias:      "REDACTED",
			location:      "REDACTED",
			carrier: CarrierQuicprotocol,
			desire:      "REDACTED",
		},
		{
			alias:        "REDACTED",
			location:        "REDACTED",
			carrier:   CarrierQuicprotocol,
			faultIncludes: "REDACTED",
		},
		{
			alias:      "REDACTED",
			location:      "REDACTED",
			carrier: CarrierQuicprotocol,
			desire:      "REDACTED",
		},
		{
			alias:      "REDACTED",
			location:      "REDACTED",
			carrier: CarrierQuicprotocol,
			desire:      "REDACTED",
		},
		{
			alias:      "REDACTED",
			location:      "REDACTED",
			carrier: CarrierQuicprotocol,
			desire:      "REDACTED",
		},
	} {
		t.Run(tt.alias, func(t *testing.T) {
			got, err := LocatorTowardVariedLocation(tt.location, tt.carrier)
			if tt.faultIncludes != "REDACTED" {
				require.ErrorContains(t, err, tt.faultIncludes)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.desire, got.String())
		})
	}
}

func VerifyLocationDetailsOriginatingMachineAlsoUUID(t *testing.T) {
	//
	produceNodeUUID := func(t *testing.T) string {
		t.Helper()
		pk := edwards25519.ProducePrivateToken()
		id, err := UUIDOriginatingSecludedToken(pk)
		require.NoError(t, err)
		return id.String()
	}

	fixedText := func(s string) func(*testing.T) string {
		return func(*testing.T) string { return s }
	}

	for _, tt := range []struct {
		alias        string
		machine        string
		id          func(*testing.T) string
		faultIncludes string
		affirm      func(t *testing.T, locationDetails peer.AddrInfo)
	}{
		{
			alias: "REDACTED",
			machine: "REDACTED",
			id:   produceNodeUUID,
			affirm: func(t *testing.T, locationDetails peer.AddrInfo) {
				require.NotEmpty(t, locationDetails.ID)
				require.Len(t, locationDetails.Addrs, 1)
				require.Equal(t, "REDACTED", locationDetails.Addrs[0].String())
			},
		},
		{
			alias: "REDACTED",
			machine: "REDACTED",
			id:   produceNodeUUID,
			affirm: func(t *testing.T, locationDetails peer.AddrInfo) {
				require.NotEmpty(t, locationDetails.ID)
				require.Len(t, locationDetails.Addrs, 1)
				require.Equal(t, "REDACTED", locationDetails.Addrs[0].String())
			},
		},
		{
			alias:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedText("REDACTED"),
			faultIncludes: "REDACTED",
		},
		{
			alias:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedText("REDACTED"),
			faultIncludes: "REDACTED",
		},
		{
			alias:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedText("REDACTED"),
			faultIncludes: "REDACTED",
		},
		{
			alias:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedText("REDACTED"),
			faultIncludes: "REDACTED",
		},
		{
			alias:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedText("REDACTED"),
			faultIncludes: "REDACTED",
		},
	} {
		t.Run(tt.alias, func(t *testing.T) {
			//
			nodeUUID := tt.id(t)

			//
			locationDetails, err := LocationDetailsOriginatingMachineAlsoUUID(tt.machine, nodeUUID)

			//
			if tt.faultIncludes != "REDACTED" {
				require.ErrorContains(t, err, tt.faultIncludes)
				require.Empty(t, locationDetails.ID)
				require.Empty(t, locationDetails.Addrs)
				return
			}

			require.NoError(t, err)
			if tt.affirm != nil {
				tt.affirm(t, locationDetails)
			}
		})
	}
}

func VerifyNetworkLocatorOriginatingNode(t *testing.T) {
	nodeUUID, err := UUIDOriginatingSecludedToken(edwards25519.ProducePrivateToken())
	require.NoError(t, err)

	t.Run("REDACTED", func(t *testing.T) {
		//
		location, err := LocatorTowardVariedLocation("REDACTED", CarrierQuicprotocol)
		require.NoError(t, err)
		require.Equal(t, "REDACTED", location.String())

		locationDetails := peer.AddrInfo{ID: nodeUUID, Addrs: []ma.Multiaddr{location}}

		networkLocation, err := networkLocatorOriginatingNode(locationDetails)
		require.NoError(t, err)
		require.NotNil(t, networkLocation.IP)
		require.Equal(t, uint16(5678), networkLocation.Channel)
	})

	t.Run("REDACTED", func(t *testing.T) {
		location, err := LocatorTowardVariedLocation("REDACTED", CarrierQuicprotocol)
		require.NoError(t, err)

		locationDetails := peer.AddrInfo{ID: nodeUUID, Addrs: []ma.Multiaddr{location}}

		_, err = networkLocatorOriginatingNode(locationDetails)
		require.Error(t, err)
		require.ErrorContains(t, err, "REDACTED")
	})
}

func VerifyFavorIDXPrivatevalue4(t *testing.T) {
	v4 := net.ParseIP("REDACTED")
	v6 := net.ParseIP("REDACTED")

	require.Equal(t, v4, favorIDXPrivatevalue4([]net.IP{v4}), "REDACTED")
	require.Equal(t, v6, favorIDXPrivatevalue4([]net.IP{v6}), "REDACTED")
	require.Equal(t, v4, favorIDXPrivatevalue4([]net.IP{v6, v4}), "REDACTED")
	require.Equal(t, v4, favorIDXPrivatevalue4([]net.IP{v4, v6}), "REDACTED")
}
