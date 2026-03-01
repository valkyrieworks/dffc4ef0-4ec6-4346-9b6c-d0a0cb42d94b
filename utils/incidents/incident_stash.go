package incidents

//
//
type IncidentStash struct {
	incidentctl   Triggerable
	incidents []incidentDetails
}

//
func FreshIncidentStash(incidentctl Triggerable) *IncidentStash {
	return &IncidentStash{
		incidentctl: incidentctl,
	}
}

//
type incidentDetails struct {
	incident string
	data  IncidentData
}

//
func (evc *IncidentStash) TriggerIncident(incident string, data IncidentData) {
	//
	evc.incidents = append(evc.incidents, incidentDetails{incident, data})
}

//
//
func (evc *IncidentStash) Purge() {
	for _, ei := range evc.incidents {
		evc.incidentctl.TriggerIncident(ei.incident, ei.data)
	}
	//
	evc.incidents = nil
}
