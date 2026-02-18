package events

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/random"
)

//
//
func VerifyAppendObserverForEventTriggerOnce(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	signals := make(chan EventData)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			//
			evsw.DeleteObserver("REDACTED")
			signals <- data
		})
	require.NoError(t, err)
	go evsw.TriggerEvent("REDACTED", "REDACTED")
	accepted := <-signals
	if accepted != "REDACTED" {
		t.Errorf("REDACTED", accepted)
	}
}

//
//
func VerifyAppendObserverForEventTriggerNumerous(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	doneTotal := make(chan uint64)
	doneDispatching := make(chan uint64)
	figures := make(chan uint64, 4)
	//
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			figures <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedFigures(figures, doneTotal)
	//
	go triggerEvents(evsw, "REDACTED", doneDispatching, uint64(1))
	inspectTotal := <-doneDispatching
	close(figures)
	eventTotal := <-doneTotal
	if inspectTotal != eventTotal {
		t.Errorf("REDACTED")
	}
}

//
//
//
func VerifyAppendObserverForDistinctEvents(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	doneTotal := make(chan uint64)
	doneDispatching1 := make(chan uint64)
	doneDispatching2 := make(chan uint64)
	doneDispatching3 := make(chan uint64)
	figures := make(chan uint64, 4)
	//
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			figures <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			figures <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			figures <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedFigures(figures, doneTotal)
	//
	go triggerEvents(evsw, "REDACTED", doneDispatching1, uint64(1))
	go triggerEvents(evsw, "REDACTED", doneDispatching2, uint64(1))
	go triggerEvents(evsw, "REDACTED", doneDispatching3, uint64(1))
	var inspectTotal uint64
	inspectTotal += <-doneDispatching1
	inspectTotal += <-doneDispatching2
	inspectTotal += <-doneDispatching3
	close(figures)
	eventTotal := <-doneTotal
	if inspectTotal != eventTotal {
		t.Errorf("REDACTED")
	}
}

//
//
//
//
func VerifyAppendDistinctObserverForDistinctEvents(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	doneTotal1 := make(chan uint64)
	doneTotal2 := make(chan uint64)
	doneDispatching1 := make(chan uint64)
	doneDispatching2 := make(chan uint64)
	doneDispatching3 := make(chan uint64)
	numbers1 := make(chan uint64, 4)
	numbers2 := make(chan uint64, 4)
	//
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers2 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers2 <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedFigures(numbers1, doneTotal1)
	//
	go totalAcceptedFigures(numbers2, doneTotal2)
	//
	go triggerEvents(evsw, "REDACTED", doneDispatching1, uint64(1))
	go triggerEvents(evsw, "REDACTED", doneDispatching2, uint64(1001))
	go triggerEvents(evsw, "REDACTED", doneDispatching3, uint64(2001))
	inspectTotalEvent1 := <-doneDispatching1
	inspectTotalEvent2 := <-doneDispatching2
	inspectTotalEvent3 := <-doneDispatching3
	inspectTotal1 := inspectTotalEvent1 + inspectTotalEvent2 + inspectTotalEvent3
	inspectTotal2 := inspectTotalEvent2 + inspectTotalEvent3
	close(numbers1)
	close(numbers2)
	eventTotal1 := <-doneTotal1
	eventTotal2 := <-doneTotal2
	if inspectTotal1 != eventTotal1 ||
		inspectTotal2 != eventTotal2 {
		t.Errorf("REDACTED")
	}
}

func VerifyAppendAndDeleteObserverParallelism(t *testing.T) {
	var (
		haltInfluxEvent = false
		epochNumber     = 2000
	)

	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	done1 := make(chan struct{})
	done2 := make(chan struct{})

	//
	//
	go func() {
		defer close(done1)
		for i := 0; i < epochNumber; i++ {
			evsw.DeleteObserver("REDACTED")
		}
	}()

	//
	go func() {
		defer close(done2)
		for i := 0; i < epochNumber; i++ {
			ordinal := i
			//
			//
			_ = evsw.AppendObserverForEvent("REDACTED", fmt.Sprintf("REDACTED", ordinal),
				func(data EventData) {
					t.Errorf("REDACTED", ordinal)
					haltInfluxEvent = true
				})
		}
	}()

	<-done1
	<-done2

	evsw.DeleteObserver("REDACTED") //

	for i := 0; i < epochNumber && !haltInfluxEvent; i++ {
		evsw.TriggerEvent(fmt.Sprintf("REDACTED", i), uint64(1001))
	}
}

//
//
//
func VerifyAppendAndDeleteObserver(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	doneTotal1 := make(chan uint64)
	doneTotal2 := make(chan uint64)
	doneDispatching1 := make(chan uint64)
	doneDispatching2 := make(chan uint64)
	numbers1 := make(chan uint64, 4)
	numbers2 := make(chan uint64, 4)
	//
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers2 <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedFigures(numbers1, doneTotal1)
	//
	go totalAcceptedFigures(numbers2, doneTotal2)
	//
	go triggerEvents(evsw, "REDACTED", doneDispatching1, uint64(1))
	inspectTotalEvent1 := <-doneDispatching1
	//
	evsw.DeleteObserver("REDACTED")
	go triggerEvents(evsw, "REDACTED", doneDispatching2, uint64(1001))
	inspectTotalEvent2 := <-doneDispatching2
	close(numbers1)
	close(numbers2)
	eventTotal1 := <-doneTotal1
	eventTotal2 := <-doneTotal2
	if inspectTotalEvent1 != eventTotal1 ||
		//
		inspectTotalEvent2 == uint64(0) ||
		eventTotal2 != uint64(0) {
		t.Errorf("REDACTED")
	}
}

//
func VerifyDeleteObserver(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	tally := 10
	total1, total2 := 0, 0
	//
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			total1++
		})
	require.NoError(t, err)

	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			total2++
		})
	require.NoError(t, err)

	for i := 0; i < tally; i++ {
		evsw.TriggerEvent("REDACTED", true)
		evsw.TriggerEvent("REDACTED", true)
	}
	assert.Equal(t, tally, total1)
	assert.Equal(t, tally, total2)

	//
	evsw.DeleteObserverForEvent("REDACTED", "REDACTED")
	for i := 0; i < tally; i++ {
		evsw.TriggerEvent("REDACTED", true)
		evsw.TriggerEvent("REDACTED", true)
	}
	assert.Equal(t, tally*2, total1)
	assert.Equal(t, tally, total2)

	//
	evsw.DeleteObserver("REDACTED")
	for i := 0; i < tally; i++ {
		evsw.TriggerEvent("REDACTED", true)
		evsw.TriggerEvent("REDACTED", true)
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
func VerifyDeleteObserversAsync(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := evsw.Halt(); err != nil {
			t.Error(err)
		}
	})

	doneTotal1 := make(chan uint64)
	doneTotal2 := make(chan uint64)
	doneDispatching1 := make(chan uint64)
	doneDispatching2 := make(chan uint64)
	doneDispatching3 := make(chan uint64)
	numbers1 := make(chan uint64, 4)
	numbers2 := make(chan uint64, 4)
	//
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers1 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers2 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers2 <- data.(uint64)
		})
	require.NoError(t, err)
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED",
		func(data EventData) {
			numbers2 <- data.(uint64)
		})
	require.NoError(t, err)
	//
	go totalAcceptedFigures(numbers1, doneTotal1)
	//
	go totalAcceptedFigures(numbers2, doneTotal2)
	appendObserversLoad := func() {
		r1 := random.NewRandom()
		r1.Source(time.Now().UnixNano())
		for k := uint16(0); k < 400; k++ {
			observerAmount := r1.Intn(100) + 3
			eventAmount := r1.Intn(3) + 1
			go evsw.AppendObserverForEvent(fmt.Sprintf("REDACTED", observerAmount), //
				fmt.Sprintf("REDACTED", eventAmount),
				func(_ EventData) {})
		}
	}
	deleteObserversLoad := func() {
		r2 := random.NewRandom()
		r2.Source(time.Now().UnixNano())
		for k := uint16(0); k < 80; k++ {
			observerAmount := r2.Intn(100) + 3
			go evsw.DeleteObserver(fmt.Sprintf("REDACTED", observerAmount))
		}
	}
	appendObserversLoad()
	//
	go triggerEvents(evsw, "REDACTED", doneDispatching1, uint64(1))
	deleteObserversLoad()
	go triggerEvents(evsw, "REDACTED", doneDispatching2, uint64(1001))
	go triggerEvents(evsw, "REDACTED", doneDispatching3, uint64(2001))
	inspectTotalEvent1 := <-doneDispatching1
	inspectTotalEvent2 := <-doneDispatching2
	inspectTotalEvent3 := <-doneDispatching3
	inspectTotal := inspectTotalEvent1 + inspectTotalEvent2 + inspectTotalEvent3
	close(numbers1)
	close(numbers2)
	eventTotal1 := <-doneTotal1
	eventTotal2 := <-doneTotal2
	if inspectTotal != eventTotal1 ||
		inspectTotal != eventTotal2 {
		t.Errorf("REDACTED")
	}
}

//
//

//
//
//
func totalAcceptedFigures(figures, doneTotal chan uint64) {
	var sum uint64
	for {
		j, additional := <-figures
		sum += j
		if !additional {
			doneTotal <- sum
			close(doneTotal)
			return
		}
	}
}

//
//
//
//
//
func triggerEvents(evsw Triggerable, event string, doneChannel chan uint64,
	displacement uint64,
) {
	var relayedTotal uint64
	for i := displacement; i <= displacement+uint64(999); i++ {
		relayedTotal += i
		evsw.TriggerEvent(event, i)
	}
	doneChannel <- relayedTotal
	close(doneChannel)
}
