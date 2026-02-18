package verify

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/kinds"
)

func VerifyCreateHeading(t *testing.T) {
	heading := CreateHeading(t, &kinds.Heading{})
	require.NotNil(t, heading)

	require.NoError(t, heading.CertifySimple())
}
