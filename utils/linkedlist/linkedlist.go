package linkedlist

/**

.
.
.
.
y
.

*/

import (
	"fmt"
	"sync"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
//
//
const MaximumMagnitude = int(^uint(0) >> 1)

/**
t
.

n
n
e
o
a
r
.

n
e
,
s
.
*/
type CNComponent struct {
	mtx        commitchronize.ReadwriteExclusion
	previous       *CNComponent
	previousGroup     *sync.WaitGroup
	previousPauseChnl chan struct{}
	following       *CNComponent
	followingGroup     *sync.WaitGroup
	followingPauseChnl chan struct{}
	discarded    bool

	Datum any //
}

//
//
func (e *CNComponent) FollowingPause() *CNComponent {
	for {
		e.mtx.RLock()
		following := e.following
		followingGroup := e.followingGroup
		discarded := e.discarded
		e.mtx.RUnlock()

		if following != nil || discarded {
			return following
		}

		followingGroup.Wait()
		//
		//
	}
}

//
//
func (e *CNComponent) PreviousPause() *CNComponent {
	for {
		e.mtx.RLock()
		previous := e.previous
		previousGroup := e.previousGroup
		discarded := e.discarded
		e.mtx.RUnlock()

		if previous != nil || discarded {
			return previous
		}

		previousGroup.Wait()
	}
}

//
//
func (e *CNComponent) PreviousPauseChn() <-chan struct{} {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	return e.previousPauseChnl
}

//
//
func (e *CNComponent) FollowingPauseChnl() <-chan struct{} {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	return e.followingPauseChnl
}

//
func (e *CNComponent) Following() *CNComponent {
	e.mtx.RLock()
	val := e.following
	e.mtx.RUnlock()
	return val
}

//
func (e *CNComponent) Previous() *CNComponent {
	e.mtx.RLock()
	previous := e.previous
	e.mtx.RUnlock()
	return previous
}

func (e *CNComponent) Discarded() bool {
	e.mtx.RLock()
	equalsDiscarded := e.discarded
	e.mtx.RUnlock()
	return equalsDiscarded
}

func (e *CNComponent) UncoupleFollowing() {
	e.mtx.Lock()
	if !e.discarded {
		e.mtx.Unlock()
		panic("REDACTED")
	}
	e.following = nil
	e.mtx.Unlock()
}

func (e *CNComponent) UncouplePrevious() {
	e.mtx.Lock()
	if !e.discarded {
		e.mtx.Unlock()
		panic("REDACTED")
	}
	e.previous = nil
	e.mtx.Unlock()
}

//
//
func (e *CNComponent) AssignFollowing(freshFollowing *CNComponent) {
	e.mtx.Lock()

	agedFollowing := e.following
	e.following = freshFollowing
	if agedFollowing != nil && freshFollowing == nil {
		//
		//
		//
		//
		//
		e.followingGroup = pauseCluster1() //
		e.followingPauseChnl = make(chan struct{})
	}
	if agedFollowing == nil && freshFollowing != nil {
		e.followingGroup.Done()
		close(e.followingPauseChnl)
	}
	e.mtx.Unlock()
}

//
//
func (e *CNComponent) AssignPrevious(freshPrevious *CNComponent) {
	e.mtx.Lock()

	agedPrevious := e.previous
	e.previous = freshPrevious
	if agedPrevious != nil && freshPrevious == nil {
		e.previousGroup = pauseCluster1() //
		e.previousPauseChnl = make(chan struct{})
	}
	if agedPrevious == nil && freshPrevious != nil {
		e.previousGroup.Done()
		close(e.previousPauseChnl)
	}
	e.mtx.Unlock()
}

func (e *CNComponent) AssignDiscarded() {
	e.mtx.Lock()

	e.discarded = true

	//
	if e.previous == nil {
		e.previousGroup.Done()
		close(e.previousPauseChnl)
	}
	if e.following == nil {
		e.followingGroup.Done()
		close(e.followingPauseChnl)
	}
	e.mtx.Unlock()
}

//

//
//
//
//
type CNCatalog struct {
	mtx    commitchronize.ReadwriteExclusion
	wg     *sync.WaitGroup
	pauseChnl chan struct{}
	header   *CNComponent //
	end   *CNComponent //
	currentLength int       //
	maximumLength int       //
}

func (l *CNCatalog) Initialize() *CNCatalog {
	l.mtx.Lock()

	l.wg = pauseCluster1()
	l.pauseChnl = make(chan struct{})
	l.header = nil
	l.end = nil
	l.currentLength = 0
	l.mtx.Unlock()
	return l
}

//
func New() *CNCatalog { return freshUsingMaximum(MaximumMagnitude) }

//
//
func freshUsingMaximum(maximumMagnitude int) *CNCatalog {
	l := new(CNCatalog)
	l.maximumLength = maximumMagnitude
	return l.Initialize()
}

func (l *CNCatalog) Len() int {
	l.mtx.RLock()
	currentLength := l.currentLength
	l.mtx.RUnlock()
	return currentLength
}

func (l *CNCatalog) Leading() *CNComponent {
	l.mtx.RLock()
	header := l.header
	l.mtx.RUnlock()
	return header
}

func (l *CNCatalog) LeadingPause() *CNComponent {
	//
	for {
		l.mtx.RLock()
		header := l.header
		wg := l.wg
		l.mtx.RUnlock()

		if header != nil {
			return header
		}
		wg.Wait()
		//
	}
}

func (l *CNCatalog) Rear() *CNComponent {
	l.mtx.RLock()
	rear := l.end
	l.mtx.RUnlock()
	return rear
}

func (l *CNCatalog) RearPause() *CNComponent {
	for {
		l.mtx.RLock()
		end := l.end
		wg := l.wg
		l.mtx.RUnlock()

		if end != nil {
			return end
		}
		wg.Wait()
		//
		//
	}
}

//
//
func (l *CNCatalog) PauseChnl() <-chan struct{} {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	return l.pauseChnl
}

//
func (l *CNCatalog) PropelRear(v any) *CNComponent {
	l.mtx.Lock()

	//
	e := &CNComponent{
		previous:       nil,
		previousGroup:     pauseCluster1(),
		previousPauseChnl: make(chan struct{}),
		following:       nil,
		followingGroup:     pauseCluster1(),
		followingPauseChnl: make(chan struct{}),
		discarded:    false,
		Datum:      v,
	}

	//
	if l.currentLength == 0 {
		l.wg.Done()
		close(l.pauseChnl)
	}
	if l.currentLength >= l.maximumLength {
		panic(fmt.Sprintf("REDACTED", l.maximumLength))
	}
	l.currentLength++

	//
	if l.end == nil {
		l.header = e
		l.end = e
	} else {
		e.AssignPrevious(l.end) //
		l.end.AssignFollowing(e) //
		l.end = e        //
	}
	l.mtx.Unlock()
	return e
}

//
//
func (l *CNCatalog) Discard(e *CNComponent) any {
	l.mtx.Lock()

	previous := e.Previous()
	following := e.Following()

	if l.header == nil || l.end == nil {
		l.mtx.Unlock()
		panic("REDACTED")
	}
	if previous == nil && l.header != e {
		l.mtx.Unlock()
		panic("REDACTED")
	}
	if following == nil && l.end != e {
		l.mtx.Unlock()
		panic("REDACTED")
	}

	//
	if l.currentLength == 1 {
		l.wg = pauseCluster1() //
		l.pauseChnl = make(chan struct{})
	}

	//
	l.currentLength--

	//
	if previous == nil {
		l.header = following
	} else {
		previous.AssignFollowing(following)
	}
	if following == nil {
		l.end = previous
	} else {
		following.AssignPrevious(previous)
	}

	//
	e.AssignDiscarded()

	l.mtx.Unlock()
	return e.Datum
}

func pauseCluster1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}
