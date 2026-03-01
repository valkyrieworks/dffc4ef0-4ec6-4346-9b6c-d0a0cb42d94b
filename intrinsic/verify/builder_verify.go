package verify

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyCreateHeadline(t *testing.T) {
	heading := CreateHeadline(t, &kinds.Heading{})
	require.NotNil(t, heading)

	require.NoError(t, heading.CertifyFundamental())
}
