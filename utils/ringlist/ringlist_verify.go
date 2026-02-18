package ringlist

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	engineseed "github.com/valkyrieworks/utils/random"
)

func VerifyAlarmOnMaximumExtent(t *testing.T) {
	maximumExtent := 1000

	l := newWithMaximum(maximumExtent)
	for i := 0; i < maximumExtent; i++ {
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

	r1 := l.Delete(el1)

	//
	//
	//

	r2 := l.Delete(el2)

	//
	//
	//

	r3 := l.Delete(el3)

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
func _Verifycollectfifo(t *testing.T) {
	if runtime.GOARCH != "REDACTED" {
		t.Skipf("REDACTED")
	}

	const countMembers = 1000000
	l := New()
	collectNumber := new(uint64)

	//
	//
	//
	type item struct {
		Int int
	}
	done := make(chan struct{})

	for i := 0; i < countMembers; i++ {
		v := new(item)
		v.Int = i
		l.PropelRear(v)
		runtime.SetFinalizer(v, func(v *item) {
			atomic.AddUint64(collectNumber, 1)
		})
	}

	for el := l.Head(); el != nil; {
		l.Delete(el)
		//
		el = el.Following()
		//
		//
	}

	runtime.GC()
	time.Sleep(time.Second * 3)
	runtime.GC()
	time.Sleep(time.Second * 3)
	_ = done

	if *collectNumber != countMembers {
		t.Errorf("REDACTED", countMembers,
			*collectNumber)
	}
}

//
//
//
//
func _Verifycollectarbitrary(t *testing.T) {
	if runtime.GOARCH != "REDACTED" {
		t.Skipf("REDACTED")
	}

	const countMembers = 1000000
	l := New()
	collectNumber := 0

	//
	//
	//
	type item struct {
		Int int
	}

	for i := 0; i < countMembers; i++ {
		v := new(item)
		v.Int = i
		l.PropelRear(v)
		runtime.SetFinalizer(v, func(v *item) {
			collectNumber++
		})
	}

	els := make([]*CComponent, 0, countMembers)
	for el := l.Head(); el != nil; el = el.Following() {
		els = append(els, el)
	}

	for _, i := range engineseed.Mode(countMembers) {
		el := els[i]
		l.Delete(el)
		_ = el.Following()
	}

	runtime.GC()
	time.Sleep(time.Second * 3)

	if collectNumber != countMembers {
		t.Errorf("REDACTED", countMembers,
			collectNumber)
	}
}

func VerifyProbeCorrectEraseArbitrary(t *testing.T) {
	const countMembers = 1000
	const countInstances = 100
	const countAnalyzers = 10

	l := New()
	halt := make(chan struct{})

	els := make([]*CComponent, countMembers)
	for i := 0; i < countMembers; i++ {
		el := l.PropelRear(i)
		els[i] = el
	}

	//
	for i := 0; i < countAnalyzers; i++ {
		go func(analyzerUID int) {
			var el *CComponent
			rebootTally := 0
			tally := 0
		FOR_CYCLE:
			for {
				select {
				case <-halt:
					fmt.Println("REDACTED")
					break FOR_CYCLE
				default:
				}
				if el == nil {
					el = l.HeadWait()
					rebootTally++
				}
				el = el.Following()
				tally++
			}
			fmt.Printf("REDACTED", analyzerUID, rebootTally, tally)
		}(i)
	}

	//
	for i := 0; i < countInstances; i++ {
		//
		removeItemIndex := engineseed.Intn(len(els))
		removeItem := els[removeItemIndex]

		//
		l.Delete(removeItem)
		//

		//
		newItem := l.PropelRear(-1*i - 1)
		els[removeItemIndex] = newItem

		if i%100000 == 0 {
			fmt.Printf("REDACTED", i/1000)
		}

	}

	//
	close(halt)
	//

	//
	for el := l.Head(); el != nil; el = el.Following() {
		l.Delete(el)
	}
	if l.Len() != 0 {
		t.Fatal("REDACTED")
	}
}

func VerifyWaitChannel(t *testing.T) {
	l := New()
	ch := l.WaitChan()

	//
	go l.PropelRear(1)
	<-ch

	//
	el := l.Head()
	v := l.Delete(el)
	if v != 1 {
		t.Fatal("REDACTED")
	}

	//
	el = l.PropelRear(0)

	done := make(chan struct{})
	impelled := 0
	go func() {
		for i := 1; i < 100; i++ {
			l.PropelRear(i)
			impelled++
			time.Sleep(time.Duration(engineseed.Intn(25)) * time.Millisecond)
		}
		//
		time.Sleep(25 * time.Millisecond)
		close(done)
	}()

	following := el
	viewed := 0
FOR_CYCLE:
	for {
		select {
		case <-following.FollowingWaitChan():
			following = following.Following()
			viewed++
			if following == nil {
				t.Fatal("REDACTED")
			}
		case <-done:
			break FOR_CYCLE
		case <-time.After(10 * time.Second):
			t.Fatal("REDACTED")
		}
	}

	if impelled != viewed {
		t.Fatalf("REDACTED", impelled, viewed)
	}

	//
	previous := following
	viewed = 0
FOR_CYCLE2:
	for {
		select {
		case <-previous.PreviousWaitChannel():
			previous = previous.Previous()
			viewed++
			if previous == nil {
				t.Fatal("REDACTED")
			}
		case <-time.After(3 * time.Second):
			break FOR_CYCLE2
		}
	}

	if impelled != viewed {
		t.Fatalf("REDACTED", impelled, viewed)
	}
}
