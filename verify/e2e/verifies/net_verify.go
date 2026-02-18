package integration_t_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
)

//
func Verifychain_Nodes(t *testing.T) {
	//
	t.SkipNow()

	verifyMember(t, func(t *testing.T, member e2e.Member) {
		//
		if member.Style == e2e.StyleOrigin {
			return
		}

		customer, err := member.Customer()
		require.NoError(t, err)
		netDetails, err := customer.NetDetails(ctx)
		require.NoError(t, err)

		require.Equal(t, len(member.Verifychain.Instances)-1, netDetails.NNodes,
			"REDACTED")

		viewed := map[string]bool{}
		for _, n := range member.Verifychain.Instances {
			viewed[n.Label] = (n.Label == member.Label) //
		}
		for _, nodeDetails := range netDetails.Nodes {
			node := member.Verifychain.SearchMember(nodeDetails.MemberDetails.Moniker)
			require.NotNil(t, node, "REDACTED", nodeDetails.MemberDetails.Moniker)
			require.Equal(t, node.IntrinsicIP.String(), nodeDetails.DistantIP,
				"REDACTED", node.Label)
			viewed[nodeDetails.MemberDetails.Moniker] = true
		}

		for label := range viewed {
			require.True(t, viewed[label], "REDACTED", member.Label, label)
		}
	})
}
