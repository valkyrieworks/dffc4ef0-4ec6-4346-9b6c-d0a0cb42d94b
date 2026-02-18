package events

//
//
type EventRepository struct {
	evsw   Triggerable
	events []eventDetails
}

//
func NewEventRepository(evsw Triggerable) *EventRepository {
	return &EventRepository{
		evsw: evsw,
	}
}

//
type eventDetails struct {
	event string
	data  EventData
}

//
func (evc *EventRepository) TriggerEvent(event string, data EventData) {
	//
	evc.events = append(evc.events, eventDetails{event, data})
}

//
//
func (evc *EventRepository) Purge() {
	for _, ei := range evc.events {
		evc.evsw.TriggerEvent(ei.event, ei.data)
	}
	//
	evc.events = nil
}
