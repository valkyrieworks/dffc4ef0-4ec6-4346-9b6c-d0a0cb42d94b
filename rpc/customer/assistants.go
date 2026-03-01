package customer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type Pauser func(variation int64) (cancel error)

//
//
func FallbackPauseTactic(variation int64) (cancel error) {
	if variation > 10 {
		return fmt.Errorf("REDACTED", variation)
	} else if variation > 0 {
		//
		//
		//
		deferral := time.Duration(variation-1)*time.Second + 500*time.Millisecond
		time.Sleep(deferral)
	}
	return nil
}

//
//
//
//
//
func PauseForeachAltitude(c ConditionCustomer, h int64, observer Pauser) error {
	if observer == nil {
		observer = FallbackPauseTactic
	}
	variation := int64(1)
	for variation > 0 {
		s, err := c.Condition(context.Background())
		if err != nil {
			return err
		}
		variation = h - s.ChronizeDetails.NewestLedgerAltitude
		//
		if err := observer(variation); err != nil {
			return err
		}
	}

	return nil
}

//
//
//
//
//
func PauseForeachSingleIncident(c IncidentsCustomer, signalKind string, deadline time.Duration) (kinds.TEMPIncidentData, error) {
	const listener = "REDACTED"
	ctx, abort := context.WithTimeout(context.Background(), deadline)
	defer abort()

	//
	incidentChnl, err := c.Listen(ctx, listener, kinds.InquireForeachIncident(signalKind).Text())
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	//
	defer func() {
		if delayFault := c.UnlistenEvery(ctx, listener); delayFault != nil {
			panic(delayFault)
		}
	}()

	select {
	case incident := <-incidentChnl:
		return incident.Data, nil
	case <-ctx.Done():
		return nil, errors.New("REDACTED")
	}
}
