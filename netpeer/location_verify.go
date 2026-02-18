package netpeer

import (
	"net"
	"testing"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func VerifyLocationToMultipleAddress(t *testing.T) {
	for _, tt := range []struct {
		label        string
		address        string
		carrier   string
		desire        string
		errIncludes string
	}{
		{
			label:      "REDACTED",
			address:      "REDACTED",
			carrier: CarrierQUIC,
			desire:      "REDACTED",
		},
		{
			label:      "REDACTED",
			address:      "REDACTED",
			carrier: CarrierQUIC,
			desire:      "REDACTED",
		},
		{
			label:        "REDACTED",
			address:        "REDACTED",
			carrier:   CarrierQUIC,
			errIncludes: "REDACTED",
		},
		{
			label:      "REDACTED",
			address:      "REDACTED",
			carrier: CarrierQUIC,
			desire:      "REDACTED",
		},
		{
			label:      "REDACTED",
			address:      "REDACTED",
			carrier: CarrierQUIC,
			desire:      "REDACTED",
		},
		{
			label:      "REDACTED",
			address:      "REDACTED",
			carrier: CarrierQUIC,
			desire:      "REDACTED",
		},
	} {
		t.Run(tt.label, func(t *testing.T) {
			got, err := LocationToMultipleAddress(tt.address, tt.carrier)
			if tt.errIncludes != "REDACTED" {
				require.ErrorContains(t, err, tt.errIncludes)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.desire, got.String())
		})
	}
}

func VerifyAddressDetailsFromMachineAndUID(t *testing.T) {
	//
	generateNodeUID := func(t *testing.T) string {
		t.Helper()
		pk := ed25519.GeneratePrivateKey()
		id, err := UIDFromPrivateKey(pk)
		require.NoError(t, err)
		return id.String()
	}

	fixedString := func(s string) func(*testing.T) string {
		return func(*testing.T) string { return s }
	}

	for _, tt := range []struct {
		label        string
		machine        string
		id          func(*testing.T) string
		errIncludes string
		affirm      func(t *testing.T, addressDetails peer.AddrInfo)
	}{
		{
			label: "REDACTED",
			machine: "REDACTED",
			id:   generateNodeUID,
			affirm: func(t *testing.T, addressDetails peer.AddrInfo) {
				require.NotEmpty(t, addressDetails.ID)
				require.Len(t, addressDetails.Addrs, 1)
				require.Equal(t, "REDACTED", addressDetails.Addrs[0].String())
			},
		},
		{
			label: "REDACTED",
			machine: "REDACTED",
			id:   generateNodeUID,
			affirm: func(t *testing.T, addressDetails peer.AddrInfo) {
				require.NotEmpty(t, addressDetails.ID)
				require.Len(t, addressDetails.Addrs, 1)
				require.Equal(t, "REDACTED", addressDetails.Addrs[0].String())
			},
		},
		{
			label:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedString("REDACTED"),
			errIncludes: "REDACTED",
		},
		{
			label:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedString("REDACTED"),
			errIncludes: "REDACTED",
		},
		{
			label:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedString("REDACTED"),
			errIncludes: "REDACTED",
		},
		{
			label:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedString("REDACTED"),
			errIncludes: "REDACTED",
		},
		{
			label:        "REDACTED",
			machine:        "REDACTED",
			id:          fixedString("REDACTED"),
			errIncludes: "REDACTED",
		},
	} {
		t.Run(tt.label, func(t *testing.T) {
			//
			nodeUID := tt.id(t)

			//
			addressDetails, err := AddressDetailsFromMachineAndUID(tt.machine, nodeUID)

			//
			if tt.errIncludes != "REDACTED" {
				require.ErrorContains(t, err, tt.errIncludes)
				require.Empty(t, addressDetails.ID)
				require.Empty(t, addressDetails.Addrs)
				return
			}

			require.NoError(t, err)
			if tt.affirm != nil {
				tt.affirm(t, addressDetails)
			}
		})
	}
}

func VerifyNetLocationFromNode(t *testing.T) {
	nodeUID, err := UIDFromPrivateKey(ed25519.GeneratePrivateKey())
	require.NoError(t, err)

	t.Run("REDACTED", func(t *testing.T) {
		//
		address, err := LocationToMultipleAddress("REDACTED", CarrierQUIC)
		require.NoError(t, err)
		require.Equal(t, "REDACTED", address.String())

		addressDetails := peer.AddrInfo{ID: nodeUID, Addrs: []ma.Multiaddr{address}}

		netAddress, err := netLocationFromNode(addressDetails)
		require.NoError(t, err)
		require.NotNil(t, netAddress.IP)
		require.Equal(t, uint16(5678), netAddress.Port)
	})

	t.Run("REDACTED", func(t *testing.T) {
		address, err := LocationToMultipleAddress("REDACTED", CarrierQUIC)
		require.NoError(t, err)

		addressDetails := peer.AddrInfo{ID: nodeUID, Addrs: []ma.Multiaddr{address}}

		_, err = netLocationFromNode(addressDetails)
		require.Error(t, err)
		require.ErrorContains(t, err, "REDACTED")
	})
}

func VerifyFavorIDXPv4(t *testing.T) {
	v4 := net.ParseIP("REDACTED")
	v6 := net.ParseIP("REDACTED")

	require.Equal(t, v4, favorIDXPv4([]net.IP{v4}), "REDACTED")
	require.Equal(t, v6, favorIDXPv4([]net.IP{v6}), "REDACTED")
	require.Equal(t, v4, favorIDXPv4([]net.IP{v6, v4}), "REDACTED")
	require.Equal(t, v4, favorIDXPv4([]net.IP{v4, v6}), "REDACTED")
}
