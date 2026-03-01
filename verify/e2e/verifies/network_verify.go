package end2end_typ_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
)

//
func Simnet_Nodes(t *testing.T) {
	//
	t.SkipNow()

	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		//
		if peer.Style == e2e.StyleGerm {
			return
		}

		customer, err := peer.Customer()
		require.NoError(t, err)
		networkDetails, err := customer.NetworkDetails(ctx)
		require.NoError(t, err)

		require.Equal(t, len(peer.Simnet.Peers)-1, networkDetails.NTHNodes,
			"REDACTED")

		observed := map[string]bool{}
		for _, n := range peer.Simnet.Peers {
			observed[n.Alias] = (n.Alias == peer.Alias) //
		}
		for _, nodeDetails := range networkDetails.Nodes {
			node := peer.Simnet.SearchPeer(nodeDetails.PeerDetails.Pseudonym)
			require.NotNil(t, node, "REDACTED", nodeDetails.PeerDetails.Pseudonym)
			require.Equal(t, node.IntrinsicINET.String(), nodeDetails.DistantINET,
				"REDACTED", node.Alias)
			observed[nodeDetails.PeerDetails.Pseudonym] = true
		}

		for alias := range observed {
			require.True(t, observed[alias], "REDACTED", peer.Alias, alias)
		}
	})
}
