package ringlist

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

	engineconnect "github.com/valkyrieworks/utils/align"
)

//
//
//
const MaximumExtent = int(^uint(0) >> 1)

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
type CComponent struct {
	mtx        engineconnect.ReadwriteLock
	previous       *CComponent
	previousGroup     *sync.WaitGroup
	previousWaitChan chan struct{}
	following       *CComponent
	followingGroup     *sync.WaitGroup
	followingWaitChan chan struct{}
	deleted    bool

	Item any //
}

//
//
func (e *CComponent) FollowingWait() *CComponent {
	for {
		e.mtx.RLock()
		following := e.following
		followingGroup := e.followingGroup
		deleted := e.deleted
		e.mtx.RUnlock()

		if following != nil || deleted {
			return following
		}

		followingGroup.Wait()
		//
		//
	}
}

//
//
func (e *CComponent) PreviousWait() *CComponent {
	for {
		e.mtx.RLock()
		previous := e.previous
		previousGroup := e.previousGroup
		deleted := e.deleted
		e.mtx.RUnlock()

		if previous != nil || deleted {
			return previous
		}

		previousGroup.Wait()
	}
}

//
//
func (e *CComponent) PreviousWaitChannel() <-chan struct{} {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	return e.previousWaitChan
}

//
//
func (e *CComponent) FollowingWaitChan() <-chan struct{} {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	return e.followingWaitChan
}

//
func (e *CComponent) Following() *CComponent {
	e.mtx.RLock()
	val := e.following
	e.mtx.RUnlock()
	return val
}

//
func (e *CComponent) Previous() *CComponent {
	e.mtx.RLock()
	previous := e.previous
	e.mtx.RUnlock()
	return previous
}

func (e *CComponent) Deleted() bool {
	e.mtx.RLock()
	isDeleted := e.deleted
	e.mtx.RUnlock()
	return isDeleted
}

func (e *CComponent) UnplugFollowing() {
	e.mtx.Lock()
	if !e.deleted {
		e.mtx.Unlock()
		panic("REDACTED")
	}
	e.following = nil
	e.mtx.Unlock()
}

func (e *CComponent) UnplugPrevious() {
	e.mtx.Lock()
	if !e.deleted {
		e.mtx.Unlock()
		panic("REDACTED")
	}
	e.previous = nil
	e.mtx.Unlock()
}

//
//
func (e *CComponent) CollectionFollowing(newFollowing *CComponent) {
	e.mtx.Lock()

	agedFollowing := e.following
	e.following = newFollowing
	if agedFollowing != nil && newFollowing == nil {
		//
		//
		//
		//
		//
		e.followingGroup = waitCluster1() //
		e.followingWaitChan = make(chan struct{})
	}
	if agedFollowing == nil && newFollowing != nil {
		e.followingGroup.Done()
		close(e.followingWaitChan)
	}
	e.mtx.Unlock()
}

//
//
func (e *CComponent) CollectionPrevious(newPrevious *CComponent) {
	e.mtx.Lock()

	agedPrevious := e.previous
	e.previous = newPrevious
	if agedPrevious != nil && newPrevious == nil {
		e.previousGroup = waitCluster1() //
		e.previousWaitChan = make(chan struct{})
	}
	if agedPrevious == nil && newPrevious != nil {
		e.previousGroup.Done()
		close(e.previousWaitChan)
	}
	e.mtx.Unlock()
}

func (e *CComponent) CollectionDeleted() {
	e.mtx.Lock()

	e.deleted = true

	//
	if e.previous == nil {
		e.previousGroup.Done()
		close(e.previousWaitChan)
	}
	if e.following == nil {
		e.followingGroup.Done()
		close(e.followingWaitChan)
	}
	e.mtx.Unlock()
}

//

//
//
//
//
type CCatalog struct {
	mtx    engineconnect.ReadwriteLock
	wg     *sync.WaitGroup
	waitChan chan struct{}
	front   *CComponent //
	end   *CComponent //
	currentSize int       //
	maximumSize int       //
}

func (l *CCatalog) Init() *CCatalog {
	l.mtx.Lock()

	l.wg = waitCluster1()
	l.waitChan = make(chan struct{})
	l.front = nil
	l.end = nil
	l.currentSize = 0
	l.mtx.Unlock()
	return l
}

//
func New() *CCatalog { return newWithMaximum(MaximumExtent) }

//
//
func newWithMaximum(maximumExtent int) *CCatalog {
	l := new(CCatalog)
	l.maximumSize = maximumExtent
	return l.Init()
}

func (l *CCatalog) Len() int {
	l.mtx.RLock()
	currentSize := l.currentSize
	l.mtx.RUnlock()
	return currentSize
}

func (l *CCatalog) Head() *CComponent {
	l.mtx.RLock()
	front := l.front
	l.mtx.RUnlock()
	return front
}

func (l *CCatalog) HeadWait() *CComponent {
	//
	for {
		l.mtx.RLock()
		front := l.front
		wg := l.wg
		l.mtx.RUnlock()

		if front != nil {
			return front
		}
		wg.Wait()
		//
	}
}

func (l *CCatalog) Rear() *CComponent {
	l.mtx.RLock()
	rear := l.end
	l.mtx.RUnlock()
	return rear
}

func (l *CCatalog) RearWait() *CComponent {
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
func (l *CCatalog) WaitChan() <-chan struct{} {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	return l.waitChan
}

//
func (l *CCatalog) PropelRear(v any) *CComponent {
	l.mtx.Lock()

	//
	e := &CComponent{
		previous:       nil,
		previousGroup:     waitCluster1(),
		previousWaitChan: make(chan struct{}),
		following:       nil,
		followingGroup:     waitCluster1(),
		followingWaitChan: make(chan struct{}),
		deleted:    false,
		Item:      v,
	}

	//
	if l.currentSize == 0 {
		l.wg.Done()
		close(l.waitChan)
	}
	if l.currentSize >= l.maximumSize {
		panic(fmt.Sprintf("REDACTED", l.maximumSize))
	}
	l.currentSize++

	//
	if l.end == nil {
		l.front = e
		l.end = e
	} else {
		e.CollectionPrevious(l.end) //
		l.end.CollectionFollowing(e) //
		l.end = e        //
	}
	l.mtx.Unlock()
	return e
}

//
//
func (l *CCatalog) Delete(e *CComponent) any {
	l.mtx.Lock()

	previous := e.Previous()
	following := e.Following()

	if l.front == nil || l.end == nil {
		l.mtx.Unlock()
		panic("REDACTED")
	}
	if previous == nil && l.front != e {
		l.mtx.Unlock()
		panic("REDACTED")
	}
	if following == nil && l.end != e {
		l.mtx.Unlock()
		panic("REDACTED")
	}

	//
	if l.currentSize == 1 {
		l.wg = waitCluster1() //
		l.waitChan = make(chan struct{})
	}

	//
	l.currentSize--

	//
	if previous == nil {
		l.front = following
	} else {
		previous.CollectionFollowing(following)
	}
	if following == nil {
		l.end = previous
	} else {
		following.CollectionPrevious(previous)
	}

	//
	e.CollectionDeleted()

	l.mtx.Unlock()
	return e.Item
}

func waitCluster1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}
