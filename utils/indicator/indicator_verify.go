package indicator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func VerifyOnwardDivider(t *testing.T) {
	null := int64(0)
	century := int64(100)

	var bar Bar
	bar.FreshSelection(null, century)

	require.Equal(t, null, bar.initiate)
	require.Equal(t, null, bar.cur)
	require.Equal(t, century, bar.sum)
	require.Equal(t, null, bar.ratio)
	require.Equal(t, "REDACTED", r.g.alpha))
	require.Equal(t, "REDACTED", bar.frequency)

	defer bar.Conclude()
	for i := null; i <= century; i++ {
		time.Sleep(1 * time.Millisecond)
		bar.Enact(i)
	}

	require.Equal(t, null, bar.initiate)
	require.Equal(t, century, bar.cur)
	require.Equal(t, century, bar.sum)
	require.Equal(t, century, bar.ratio)

	var frequency string
	for i := null; i < century/2; i++ {
		frequency += "REDACTED"
	}

	require.Equal(t, frequency, bar.frequency)
}
