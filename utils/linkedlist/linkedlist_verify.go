package linkedlist

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

func VerifyAlarmUponMaximumMagnitude(t *testing.T) {
	maximumMagnitude := 1000

	l := freshUsingMaximum(maximumMagnitude)
	for i := 0; i < maximumMagnitude; i++ {
		l.PropelRear(1)
	}
	assert.Panics(t, func() {
		l.PropelRear(1)
	})
}

func VerifyMinor(t *testing.T) {
	l := New()
	el1 := l.PropelRear(1)
	el2 := l.PropelRear(2)
	el3 := l.PropelRear(3)
	if l.Len() != 3 {
		t.Error("REDACTED", l.Len())
	}

	//
	//
	//

	r1 := l.Discard(el1)

	//
	//
	//

	r2 := l.Discard(el2)

	//
	//
	//

	r3 := l.Discard(el3)

	if r1 != 1 {
		t.Error("REDACTED", r1)
	}
	if r2 != 2 {
		t.Error("REDACTED", r2)
	}
	if r3 != 3 {
		t.Error("REDACTED", r3)
	}
	if l.Len() != 0 {
		t.Error("REDACTED", l.Len())
	}
}

//
//
//
//
func _Verifygcfifo(t *testing.T) {
	if runtime.GOARCH != "REDACTED" {
		t.Skipf("REDACTED")
	}

	const countConstituents = 1000000
	l := New()
	gcollectTally := new(uint64)

	//
	//
	//
	type datum struct {
		Int int
	}
	complete := make(chan struct{})

	for i := 0; i < countConstituents; i++ {
		v := new(datum)
		v.Int = i
		l.PropelRear(v)
		runtime.SetFinalizer(v, func(v *datum) {
			atomic.AddUint64(gcollectTally, 1)
		})
	}

	for el := l.Leading(); el != nil; {
		l.Discard(el)
		//
		el = el.Following()
		//
		//
	}

	runtime.GC()
	time.Sleep(time.Second * 3)
	runtime.GC()
	time.Sleep(time.Second * 3)
	_ = complete

	if *gcollectTally != countConstituents {
		t.Errorf("REDACTED", countConstituents,
			*gcollectTally)
	}
}

//
//
//
//
func _Verifygcarbitrary(t *testing.T) {
	if runtime.GOARCH != "REDACTED" {
		t.Skipf("REDACTED")
	}

	const countConstituents = 1000000
	l := New()
	gcollectTally := 0

	//
	//
	//
	type datum struct {
		Int int
	}

	for i := 0; i < countConstituents; i++ {
		v := new(datum)
		v.Int = i
		l.PropelRear(v)
		runtime.SetFinalizer(v, func(v *datum) {
			gcollectTally++
		})
	}

	els := make([]*CNComponent, 0, countConstituents)
	for el := l.Leading(); el != nil; el = el.Following() {
		els = append(els, el)
	}

	for _, i := range commitrand.Mode(countConstituents) {
		el := els[i]
		l.Discard(el)
		_ = el.Following()
	}

	runtime.GC()
	time.Sleep(time.Second * 3)

	if gcollectTally != countConstituents {
		t.Errorf("REDACTED", countConstituents,
			gcollectTally)
	}
}

func VerifyProbeTrailingEraseUnpredictable(t *testing.T) {
	const countConstituents = 1000
	const countMultiples = 100
	const countAnalyzers = 10

	l := New()
	halt := make(chan struct{})

	els := make([]*CNComponent, countConstituents)
	for i := 0; i < countConstituents; i++ {
		el := l.PropelRear(i)
		els[i] = el
	}

	//
	for i := 0; i < countAnalyzers; i++ {
		go func(analyzerUUID int) {
			var el *CNComponent
			rebootTally := 0
			tally := 0
		FOREACH_CYCLE:
			for {
				select {
				case <-halt:
					fmt.Println("REDACTED")
					break FOREACH_CYCLE
				default:
				}
				if el == nil {
					el = l.LeadingPause()
					rebootTally++
				}
				el = el.Following()
				tally++
			}
			fmt.Printf("REDACTED", analyzerUUID, rebootTally, tally)
		}(i)
	}

	//
	for i := 0; i < countMultiples; i++ {
		//
		delElementOffset := commitrand.Integern(len(els))
		delElement := els[delElementOffset]

		//
		l.Discard(delElement)
		//

		//
		freshElement := l.PropelRear(-1*i - 1)
		els[delElementOffset] = freshElement

		if i%100000 == 0 {
			fmt.Printf("REDACTED", i/1000)
		}

	}

	//
	close(halt)
	//

	//
	for el := l.Leading(); el != nil; el = el.Following() {
		l.Discard(el)
	}
	if l.Len() != 0 {
		t.Fatal("REDACTED")
	}
}

func VerifyPauseChn(t *testing.T) {
	l := New()
	ch := l.PauseChnl()

	//
	go l.PropelRear(1)
	<-ch

	//
	el := l.Leading()
	v := l.Discard(el)
	if v != 1 {
		t.Fatal("REDACTED")
	}

	//
	el = l.PropelRear(0)

	complete := make(chan struct{})
	propelled := 0
	go func() {
		for i := 1; i < 100; i++ {
			l.PropelRear(i)
			propelled++
			time.Sleep(time.Duration(commitrand.Integern(25)) * time.Millisecond)
		}
		//
		time.Sleep(25 * time.Millisecond)
		close(complete)
	}()

	following := el
	observed := 0
FOREACH_CYCLE:
	for {
		select {
		case <-following.FollowingPauseChnl():
			following = following.Following()
			observed++
			if following == nil {
				t.Fatal("REDACTED")
			}
		case <-complete:
			break FOREACH_CYCLE
		case <-time.After(10 * time.Second):
			t.Fatal("REDACTED")
		}
	}

	if propelled != observed {
		t.Fatalf("REDACTED", propelled, observed)
	}

	//
	previous := following
	observed = 0
FOREACH_CYCLE2:
	for {
		select {
		case <-previous.PreviousPauseChn():
			previous = previous.Previous()
			observed++
			if previous == nil {
				t.Fatal("REDACTED")
			}
		case <-time.After(3 * time.Second):
			break FOREACH_CYCLE2
		}
	}

	if propelled != observed {
		t.Fatalf("REDACTED", propelled, observed)
	}
}
