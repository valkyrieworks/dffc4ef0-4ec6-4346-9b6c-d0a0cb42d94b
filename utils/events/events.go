//
package events

import (
	"fmt"

	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
)

//
type ErrObserverWasDeleted struct {
	observerUID string
}

//
func (e ErrObserverWasDeleted) Fault() string {
	return fmt.Sprintf("REDACTED", e.observerUID)
}

//
//
type EventData any

//
//
type Observable interface {
	CollectionEventRouter(evsw EventRouter)
}

//
//
//
type Triggerable interface {
	TriggerEvent(event string, data EventData)
}

//
//
//
//
//
//
//
type EventRouter interface {
	daemon.Daemon
	Triggerable

	AppendObserverForEvent(observerUID, event string, cb EventResponse) error
	DeleteObserverForEvent(event string, observerUID string)
	DeleteObserver(observerUID string)
}

type eventRouter struct {
	daemon.RootDaemon

	mtx        engineconnect.ReadwriteLock
	eventSegments map[string]*eventSegment
	observers  map[string]*eventObserver
}

func NewEventRouter() EventRouter {
	evsw := &eventRouter{
		eventSegments: make(map[string]*eventSegment),
		observers:  make(map[string]*eventObserver),
	}
	evsw.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", evsw)
	return evsw
}

func (evsw *eventRouter) OnBegin() error {
	return nil
}

func (evsw *eventRouter) OnHalt() {}

func (evsw *eventRouter) AppendObserverForEvent(observerUID, event string, cb EventResponse) error {
	//
	evsw.mtx.Lock()
	eventSegment := evsw.eventSegments[event]
	if eventSegment == nil {
		eventSegment = newEventSegment()
		evsw.eventSegments[event] = eventSegment
	}
	observer := evsw.observers[observerUID]
	if observer == nil {
		observer = newEventObserver(observerUID)
		evsw.observers[observerUID] = observer
	}
	evsw.mtx.Unlock()

	//
	if err := observer.AppendEvent(event); err != nil {
		return err
	}
	eventSegment.AppendObserver(observerUID, cb)

	return nil
}

func (evsw *eventRouter) DeleteObserver(observerUID string) {
	//
	evsw.mtx.RLock()
	observer := evsw.observers[observerUID]
	evsw.mtx.RUnlock()
	if observer == nil {
		return
	}

	evsw.mtx.Lock()
	delete(evsw.observers, observerUID)
	evsw.mtx.Unlock()

	//
	observer.CollectionDeleted()
	for _, event := range observer.FetchEvents() {
		evsw.DeleteObserverForEvent(event, observerUID)
	}
}

func (evsw *eventRouter) DeleteObserverForEvent(event string, observerUID string) {
	//
	evsw.mtx.Lock()
	eventSegment := evsw.eventSegments[event]
	evsw.mtx.Unlock()

	if eventSegment == nil {
		return
	}

	//
	countObservers := eventSegment.DeleteObserver(observerUID)

	//
	if countObservers == 0 {
		//
		evsw.mtx.Lock()      //
		eventSegment.mtx.Lock() //
		if len(eventSegment.observers) == 0 {
			delete(evsw.eventSegments, event)
		}
		eventSegment.mtx.Unlock() //
		evsw.mtx.Unlock()      //
	}
}

func (evsw *eventRouter) TriggerEvent(event string, data EventData) {
	//
	evsw.mtx.RLock()
	eventSegment := evsw.eventSegments[event]
	evsw.mtx.RUnlock()

	if eventSegment == nil {
		return
	}

	//
	eventSegment.TriggerEvent(data)
}

//

//
type eventSegment struct {
	mtx       engineconnect.ReadwriteLock
	observers map[string]EventResponse
}

func newEventSegment() *eventSegment {
	return &eventSegment{
		observers: make(map[string]EventResponse),
	}
}

func (segment *eventSegment) AppendObserver(observerUID string, cb EventResponse) {
	segment.mtx.Lock()
	segment.observers[observerUID] = cb
	segment.mtx.Unlock()
}

func (segment *eventSegment) DeleteObserver(observerUID string) int {
	segment.mtx.Lock()
	delete(segment.observers, observerUID)
	countObservers := len(segment.observers)
	segment.mtx.Unlock()
	return countObservers
}

func (segment *eventSegment) TriggerEvent(data EventData) {
	segment.mtx.RLock()
	eventResponses := make([]EventResponse, 0, len(segment.observers))
	for _, cb := range segment.observers {
		eventResponses = append(eventResponses, cb)
	}
	segment.mtx.RUnlock()

	for _, cb := range eventResponses {
		cb(data)
	}
}

//

type EventResponse func(data EventData)

type eventObserver struct {
	id string

	mtx     engineconnect.ReadwriteLock
	deleted bool
	events  []string
}

func newEventObserver(id string) *eventObserver {
	return &eventObserver{
		id:      id,
		deleted: false,
		events:  nil,
	}
}

func (evl *eventObserver) AppendEvent(event string) error {
	evl.mtx.Lock()

	if evl.deleted {
		evl.mtx.Unlock()
		return ErrObserverWasDeleted{observerUID: evl.id}
	}

	evl.events = append(evl.events, event)
	evl.mtx.Unlock()
	return nil
}

func (evl *eventObserver) FetchEvents() []string {
	evl.mtx.RLock()
	events := make([]string, len(evl.events))
	copy(events, evl.events)
	evl.mtx.RUnlock()
	return events
}

func (evl *eventObserver) CollectionDeleted() {
	evl.mtx.Lock()
	evl.deleted = true
	evl.mtx.Unlock()
}
