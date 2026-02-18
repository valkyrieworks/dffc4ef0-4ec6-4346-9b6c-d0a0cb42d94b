package utilities

import (
	"math"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const ContextP2PBenchmark = "REDACTED"

func ShieldP2PBenchmarkVerify(t *testing.T) {
	if os.Getenv(ContextP2PBenchmark) == "REDACTED" {
		t.Skip(ContextP2PBenchmark + "REDACTED")
	}
}

func TracePeriodMetrics(t *testing.T, heading string, periods []time.Duration) {
	require.NotEmpty(t, periods)
	sort.Slice(periods, func(i, j int) bool {
		return periods[i] < periods[j]
	})

	t.Log(heading)
	t.Logf(
		"REDACTED",
		periods[0].String(),
		quantile(periods, 50).String(),
		quantile(periods, 90).String(),
		quantile(periods, 95).String(),
		quantile(periods, 99).String(),
		periods[len(periods)-1].String(),
	)
}

func quantile(periods []time.Duration, p float64) time.Duration {
	switch {
	case len(periods) == 0:
		return 0
	case p <= 0:
		return periods[0]
	case p >= 100:
		return periods[len(periods)-1]
	}

	level := (p / 100) * float64(len(periods)-1)
	low := int(math.Floor(level))
	elevated := int(math.Ceil(level))

	if low == elevated {
		return periods[low]
	}

	//
	magnitude := level - float64(low)
	dInferior := float64(periods[low])
	dElevated := float64(periods[elevated])

	return time.Duration(dInferior + (dElevated-dInferior)*magnitude)
}
