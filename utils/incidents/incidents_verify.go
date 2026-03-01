package incidents

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

//
//
func VerifyAppendObserverForeachIncidentTriggerOnetime(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	signals := make(chan IncidentData)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			//
			incidentctl.DiscardObserver("REDACTED")
			signals <- data
		})
	require.NoError(t, err)
	go incidentctl.TriggerIncident("REDACTED", "REDACTED")
	accepted := <-signals
	if accepted != "REDACTED" {
		t.Errorf("REDACTED", accepted)
	}
}

//
//
func VerifyAppendObserverForeachIncidentTriggerMultiple(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	completeTotal := make(chan uint64)
	completeRelaying := make(chan uint64)
	numerals := make(chan uint64, 4)
	//
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedNumerals(numerals, completeTotal)
	//
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying, uint64(1))
	inspectTotal := <-completeRelaying
	close(numerals)
	incidentTotal := <-completeTotal
	if inspectTotal != incidentTotal {
		t.Errorf("REDACTED")
	}
}

//
//
//
func VerifyAppendObserverForeachDistinctIncidents(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	completeTotal := make(chan uint64)
	completeRelaying1 := make(chan uint64)
	completeRelaying2 := make(chan uint64)
	completeRelaying3 := make(chan uint64)
	numerals := make(chan uint64, 4)
	//
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedNumerals(numerals, completeTotal)
	//
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying1, uint64(1))
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying2, uint64(1))
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying3, uint64(1))
	var inspectTotal uint64
	inspectTotal += <-completeRelaying1
	inspectTotal += <-completeRelaying2
	inspectTotal += <-completeRelaying3
	close(numerals)
	incidentTotal := <-completeTotal
	if inspectTotal != incidentTotal {
		t.Errorf("REDACTED")
	}
}

//
//
//
//
func VerifyAppendDistinctObserverForeachDistinctIncidents(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	completeTotal1 := make(chan uint64)
	completeTotal2 := make(chan uint64)
	completeRelaying1 := make(chan uint64)
	completeRelaying2 := make(chan uint64)
	completeRelaying3 := make(chan uint64)
	numerals1 := make(chan uint64, 4)
	numerals2 := make(chan uint64, 4)
	//
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals2 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals2 <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedNumerals(numerals1, completeTotal1)
	//
	go totalAcceptedNumerals(numerals2, completeTotal2)
	//
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying1, uint64(1))
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying2, uint64(1001))
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying3, uint64(2001))
	inspectTotalIncident1 := <-completeRelaying1
	inspectTotalIncident2 := <-completeRelaying2
	inspectTotalIncident3 := <-completeRelaying3
	inspectTotal1 := inspectTotalIncident1 + inspectTotalIncident2 + inspectTotalIncident3
	inspectTotal2 := inspectTotalIncident2 + inspectTotalIncident3
	close(numerals1)
	close(numerals2)
	incidentTotal1 := <-completeTotal1
	incidentTotal2 := <-completeTotal2
	if inspectTotal1 != incidentTotal1 ||
		inspectTotal2 != incidentTotal2 {
		t.Errorf("REDACTED")
	}
}

func VerifyAppendAlsoDiscardObserverParallelism(t *testing.T) {
	var (
		haltInfluxIncident = false
		iterationTally     = 2000
	)

	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	complete1 := make(chan struct{})
	complete2 := make(chan struct{})

	//
	//
	go func() {
		defer close(complete1)
		for i := 0; i < iterationTally; i++ {
			incidentctl.DiscardObserver("REDACTED")
		}
	}()

	//
	go func() {
		defer close(complete2)
		for i := 0; i < iterationTally; i++ {
			ordinal := i
			//
			//
			_ = incidentctl.AppendObserverForeachIncident("REDACTED", fmt.Sprintf("REDACTED", ordinal),
				func(data IncidentData) {
					t.Errorf("REDACTED", ordinal)
					haltInfluxIncident = true
				})
		}
	}()

	<-complete1
	<-complete2

	incidentctl.DiscardObserver("REDACTED") //

	for i := 0; i < iterationTally && !haltInfluxIncident; i++ {
		incidentctl.TriggerIncident(fmt.Sprintf("REDACTED", i), uint64(1001))
	}
}

//
//
//
func VerifyAppendAlsoDiscardObserver(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	completeTotal1 := make(chan uint64)
	completeTotal2 := make(chan uint64)
	completeRelaying1 := make(chan uint64)
	completeRelaying2 := make(chan uint64)
	numerals1 := make(chan uint64, 4)
	numerals2 := make(chan uint64, 4)
	//
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals2 <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedNumerals(numerals1, completeTotal1)
	//
	go totalAcceptedNumerals(numerals2, completeTotal2)
	//
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying1, uint64(1))
	inspectTotalIncident1 := <-completeRelaying1
	//
	incidentctl.DiscardObserver("REDACTED")
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying2, uint64(1001))
	inspectTotalIncident2 := <-completeRelaying2
	close(numerals1)
	close(numerals2)
	incidentTotal1 := <-completeTotal1
	incidentTotal2 := <-completeTotal2
	if inspectTotalIncident1 != incidentTotal1 ||
		//
		inspectTotalIncident2 == uint64(0) ||
		incidentTotal2 != uint64(0) {
		t.Errorf("REDACTED")
	}
}

//
func VerifyDiscardObserver(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	tally := 10
	total1, total2 := 0, 0
	//
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			total1++
		})
	require.NoError(t, err)

	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			total2++
		})
	require.NoError(t, err)

	for i := 0; i < tally; i++ {
		incidentctl.TriggerIncident("REDACTED", true)
		incidentctl.TriggerIncident("REDACTED", true)
	}
	assert.Equal(t, tally, total1)
	assert.Equal(t, tally, total2)

	//
	incidentctl.DiscardObserverForeachIncident("REDACTED", "REDACTED")
	for i := 0; i < tally; i++ {
		incidentctl.TriggerIncident("REDACTED", true)
		incidentctl.TriggerIncident("REDACTED", true)
	}
	assert.Equal(t, tally*2, total1)
	assert.Equal(t, tally, total2)

	//
	incidentctl.DiscardObserver("REDACTED")
	for i := 0; i < tally; i++ {
		incidentctl.TriggerIncident("REDACTED", true)
		incidentctl.TriggerIncident("REDACTED", true)
	}
	assert.Equal(t, tally*2, total1)
	assert.Equal(t, tally, total2)
}

//
//
//
//
//
//
//
//
//
//
func VerifyDiscardObserversAsyncronous(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentctl.Halt(); err != nil {
			t.Error(err)
		}
	})

	completeTotal1 := make(chan uint64)
	completeTotal2 := make(chan uint64)
	completeRelaying1 := make(chan uint64)
	completeRelaying2 := make(chan uint64)
	completeRelaying3 := make(chan uint64)
	numerals1 := make(chan uint64, 4)
	numerals2 := make(chan uint64, 4)
	//
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals2 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals2 <- data.(uint64)
		})
	require.NoError(t, err)
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED",
		func(data IncidentData) {
			numerals2 <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedNumerals(numerals1, completeTotal1)
	//
	go totalAcceptedNumerals(numerals2, completeTotal2)
	appendObserversPressure := func() {
		r1 := arbitrary.FreshArbitrary()
		r1.Germ(time.Now().UnixNano())
		for k := uint16(0); k < 400; k++ {
			observerNumeral := r1.Integern(100) + 3
			incidentNumeral := r1.Integern(3) + 1
			go incidentctl.AppendObserverForeachIncident(fmt.Sprintf("REDACTED", observerNumeral), //
				fmt.Sprintf("REDACTED", incidentNumeral),
				func(_ IncidentData) {})
		}
	}
	discardObserversPressure := func() {
		r2 := arbitrary.FreshArbitrary()
		r2.Germ(time.Now().UnixNano())
		for k := uint16(0); k < 80; k++ {
			observerNumeral := r2.Integern(100) + 3
			go incidentctl.DiscardObserver(fmt.Sprintf("REDACTED", observerNumeral))
		}
	}
	appendObserversPressure()
	//
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying1, uint64(1))
	discardObserversPressure()
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying2, uint64(1001))
	go triggerIncidents(incidentctl, "REDACTED", completeRelaying3, uint64(2001))
	inspectTotalIncident1 := <-completeRelaying1
	inspectTotalIncident2 := <-completeRelaying2
	inspectTotalIncident3 := <-completeRelaying3
	inspectTotal := inspectTotalIncident1 + inspectTotalIncident2 + inspectTotalIncident3
	close(numerals1)
	close(numerals2)
	incidentTotal1 := <-completeTotal1
	incidentTotal2 := <-completeTotal2
	if inspectTotal != incidentTotal1 ||
		inspectTotal != incidentTotal2 {
		t.Errorf("REDACTED")
	}
}

//
//

//
//
//
func totalAcceptedNumerals(numerals, completeTotal chan uint64) {
	var sum uint64
	for {
		j, extra := <-numerals
		sum += j
		if !extra {
			completeTotal <- sum
			close(completeTotal)
			return
		}
	}
}

//
//
//
//
//
func triggerIncidents(incidentctl Triggerable, incident string, completeChn chan uint64,
	displacement uint64,
) {
	var relayedTotal uint64
	for i := displacement; i <= displacement+uint64(999); i++ {
		relayedTotal += i
		incidentctl.TriggerIncident(incident, i)
	}
	completeChn <- relayedTotal
	close(completeChn)
}
