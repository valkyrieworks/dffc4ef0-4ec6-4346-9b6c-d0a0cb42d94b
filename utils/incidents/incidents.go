//
package incidents

import (
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
type FaultObserverExistedDiscarded struct {
	observerUUID string
}

//
func (e FaultObserverExistedDiscarded) Failure() string {
	return fmt.Sprintf("REDACTED", e.observerUUID)
}

//
//
type IncidentData any

//
//
type Incidental interface {
	GroupIncidentRouter(incidentctl IncidentRouter)
}

//
//
//
type Triggerable interface {
	TriggerIncident(incident string, data IncidentData)
}

//
//
//
//
//
//
//
type IncidentRouter interface {
	facility.Facility
	Triggerable

	AppendObserverForeachIncident(observerUUID, incident string, cb IncidentReact) error
	DiscardObserverForeachIncident(incident string, observerUUID string)
	DiscardObserver(observerUUID string)
}

type incidentRouter struct {
	facility.FoundationFacility

	mtx        commitchronize.ReadwriteExclusion
	incidentNodes map[string]*incidentNode
	observers  map[string]*incidentObserver
}

func FreshIncidentRouter() IncidentRouter {
	incidentctl := &incidentRouter{
		incidentNodes: make(map[string]*incidentNode),
		observers:  make(map[string]*incidentObserver),
	}
	incidentctl.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", incidentctl)
	return incidentctl
}

func (incidentctl *incidentRouter) UponInitiate() error {
	return nil
}

func (incidentctl *incidentRouter) UponHalt() {}

func (incidentctl *incidentRouter) AppendObserverForeachIncident(observerUUID, incident string, cb IncidentReact) error {
	//
	incidentctl.mtx.Lock()
	incidentNode := incidentctl.incidentNodes[incident]
	if incidentNode == nil {
		incidentNode = freshIncidentNode()
		incidentctl.incidentNodes[incident] = incidentNode
	}
	observer := incidentctl.observers[observerUUID]
	if observer == nil {
		observer = freshIncidentObserver(observerUUID)
		incidentctl.observers[observerUUID] = observer
	}
	incidentctl.mtx.Unlock()

	//
	if err := observer.AppendIncident(incident); err != nil {
		return err
	}
	incidentNode.AppendObserver(observerUUID, cb)

	return nil
}

func (incidentctl *incidentRouter) DiscardObserver(observerUUID string) {
	//
	incidentctl.mtx.RLock()
	observer := incidentctl.observers[observerUUID]
	incidentctl.mtx.RUnlock()
	if observer == nil {
		return
	}

	incidentctl.mtx.Lock()
	delete(incidentctl.observers, observerUUID)
	incidentctl.mtx.Unlock()

	//
	observer.AssignDiscarded()
	for _, incident := range observer.ObtainIncidents() {
		incidentctl.DiscardObserverForeachIncident(incident, observerUUID)
	}
}

func (incidentctl *incidentRouter) DiscardObserverForeachIncident(incident string, observerUUID string) {
	//
	incidentctl.mtx.Lock()
	incidentNode := incidentctl.incidentNodes[incident]
	incidentctl.mtx.Unlock()

	if incidentNode == nil {
		return
	}

	//
	countObservers := incidentNode.DiscardObserver(observerUUID)

	//
	if countObservers == 0 {
		//
		incidentctl.mtx.Lock()      //
		incidentNode.mtx.Lock() //
		if len(incidentNode.observers) == 0 {
			delete(incidentctl.incidentNodes, incident)
		}
		incidentNode.mtx.Unlock() //
		incidentctl.mtx.Unlock()      //
	}
}

func (incidentctl *incidentRouter) TriggerIncident(incident string, data IncidentData) {
	//
	incidentctl.mtx.RLock()
	incidentNode := incidentctl.incidentNodes[incident]
	incidentctl.mtx.RUnlock()

	if incidentNode == nil {
		return
	}

	//
	incidentNode.TriggerIncident(data)
}

//

//
type incidentNode struct {
	mtx       commitchronize.ReadwriteExclusion
	observers map[string]IncidentReact
}

func freshIncidentNode() *incidentNode {
	return &incidentNode{
		observers: make(map[string]IncidentReact),
	}
}

func (node *incidentNode) AppendObserver(observerUUID string, cb IncidentReact) {
	node.mtx.Lock()
	node.observers[observerUUID] = cb
	node.mtx.Unlock()
}

func (node *incidentNode) DiscardObserver(observerUUID string) int {
	node.mtx.Lock()
	delete(node.observers, observerUUID)
	countObservers := len(node.observers)
	node.mtx.Unlock()
	return countObservers
}

func (node *incidentNode) TriggerIncident(data IncidentData) {
	node.mtx.RLock()
	incidentReacts := make([]IncidentReact, 0, len(node.observers))
	for _, cb := range node.observers {
		incidentReacts = append(incidentReacts, cb)
	}
	node.mtx.RUnlock()

	for _, cb := range incidentReacts {
		cb(data)
	}
}

//

type IncidentReact func(data IncidentData)

type incidentObserver struct {
	id string

	mtx     commitchronize.ReadwriteExclusion
	discarded bool
	incidents  []string
}

func freshIncidentObserver(id string) *incidentObserver {
	return &incidentObserver{
		id:      id,
		discarded: false,
		incidents:  nil,
	}
}

func (evl *incidentObserver) AppendIncident(incident string) error {
	evl.mtx.Lock()

	if evl.discarded {
		evl.mtx.Unlock()
		return FaultObserverExistedDiscarded{observerUUID: evl.id}
	}

	evl.incidents = append(evl.incidents, incident)
	evl.mtx.Unlock()
	return nil
}

func (evl *incidentObserver) ObtainIncidents() []string {
	evl.mtx.RLock()
	incidents := make([]string, len(evl.incidents))
	copy(incidents, evl.incidents)
	evl.mtx.RUnlock()
	return incidents
}

func (evl *incidentObserver) AssignDiscarded() {
	evl.mtx.Lock()
	evl.discarded = true
	evl.mtx.Unlock()
}
