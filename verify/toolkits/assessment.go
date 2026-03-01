package toolkits

import (
	"math"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const EnvironmentPeer2peerBenchmark = "REDACTED"

func ShieldPeer2peerBenchmarkVerify(t *testing.T) {
	if os.Getenv(EnvironmentPeer2peerBenchmark) == "REDACTED" {
		t.Skip(EnvironmentPeer2peerBenchmark + "REDACTED")
	}
}

func RecordIntervalMetrics(t *testing.T, heading string, intervals []time.Duration) {
	require.NotEmpty(t, intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i] < intervals[j]
	})

	t.Log(heading)
	t.Logf(
		"REDACTED",
		intervals[0].String(),
		quantile(intervals, 50).String(),
		quantile(intervals, 90).String(),
		quantile(intervals, 95).String(),
		quantile(intervals, 99).String(),
		intervals[len(intervals)-1].String(),
	)
}

func quantile(intervals []time.Duration, p float64) time.Duration {
	switch {
	case len(intervals) == 0:
		return 0
	case p <= 0:
		return intervals[0]
	case p >= 100:
		return intervals[len(intervals)-1]
	}

	position := (p / 100) * float64(len(intervals)-1)
	low := int(math.Floor(position))
	tall := int(math.Ceil(position))

	if low == tall {
		return intervals[low]
	}

	//
	load := position - float64(low)
	deltaShort := float64(intervals[low])
	deltaTall := float64(intervals[tall])

	return time.Duration(deltaShort + (deltaTall-deltaShort)*load)
}
