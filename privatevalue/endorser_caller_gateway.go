package privatevalue

import (
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

const (
	fallbackMaximumCallAttempts        = 10
	fallbackReissuePauseMillis = 100
)

//
type EndorserFacilityGatewaySelection func(*EndorserCallerGateway)

//
//
func EndorserCallerGatewayDeadlineRetrievePersist(deadline time.Duration) EndorserFacilityGatewaySelection {
	return func(ss *EndorserCallerGateway) { ss.deadlineRetrievePersist = deadline }
}

//
//
func EndorserCallerGatewayLinkAttempts(attempts int) EndorserFacilityGatewaySelection {
	return func(ss *EndorserCallerGateway) { ss.maximumLinkAttempts = attempts }
}

//
//
func EndorserCallerGatewayReissuePauseDuration(duration time.Duration) EndorserFacilityGatewaySelection {
	return func(ss *EndorserCallerGateway) { ss.reissuePause = duration }
}

//
//
type EndorserCallerGateway struct {
	endorserGateway

	caller PortCaller

	reissuePause      time.Duration
	maximumLinkAttempts int
}

//
//
//
func FreshEndorserCallerGateway(
	tracer log.Tracer,
	caller PortCaller,
	choices ...EndorserFacilityGatewaySelection,
) *EndorserCallerGateway {
	sd := &EndorserCallerGateway{
		caller:         caller,
		reissuePause:      fallbackReissuePauseMillis * time.Millisecond,
		maximumLinkAttempts: fallbackMaximumCallAttempts,
	}

	sd.FoundationFacility = *facility.FreshFoundationFacility(tracer, "REDACTED", sd)
	sd.deadlineRetrievePersist = fallbackDeadlineRetrievePersistMoments * time.Second

	for _, selectionMethod := range choices {
		selectionMethod(sd)
	}

	return sd
}

func (sd *EndorserCallerGateway) assureLinkage() error {
	if sd.EqualsAssociated() {
		return nil
	}

	attempts := 0
	for attempts < sd.maximumLinkAttempts {
		link, err := sd.caller()

		if err != nil {
			attempts++
			sd.Tracer.Diagnose("REDACTED", "REDACTED", attempts, "REDACTED", sd.maximumLinkAttempts, "REDACTED", err)
			//
			time.Sleep(sd.reissuePause)
		} else {
			sd.AssignLinkage(link)
			sd.Tracer.Diagnose("REDACTED")
			return nil
		}
	}

	sd.Tracer.Diagnose("REDACTED", "REDACTED", attempts, "REDACTED", sd.maximumLinkAttempts)

	return FaultNegativeLinkage
}
