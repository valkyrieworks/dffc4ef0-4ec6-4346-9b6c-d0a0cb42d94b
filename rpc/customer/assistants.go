package customer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/kinds"
)

//
type Observer func(variance int64) (cancel error)

//
//
func StandardWaitTactic(variance int64) (cancel error) {
	if variance > 10 {
		return fmt.Errorf("REDACTED", variance)
	} else if variance > 0 {
		//
		//
		//
		deferral := time.Duration(variance-1)*time.Second + 500*time.Millisecond
		time.Sleep(deferral)
	}
	return nil
}

//
//
//
//
//
func WaitForLevel(c StateCustomer, h int64, observer Observer) error {
	if observer == nil {
		observer = StandardWaitTactic
	}
	variance := int64(1)
	for variance > 0 {
		s, err := c.Status(context.Background())
		if err != nil {
			return err
		}
		variance = h - s.AlignDetails.NewestLedgerLevel
		//
		if err := observer(variance); err != nil {
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
func WaitForOneEvent(c EventsCustomer, occurrenceType string, deadline time.Duration) (kinds.TMEventData, error) {
	const enrollee = "REDACTED"
	ctx, revoke := context.WithTimeout(context.Background(), deadline)
	defer revoke()

	//
	eventChan, err := c.Enrol(ctx, enrollee, kinds.InquireForEvent(occurrenceType).String())
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	//
	defer func() {
		if delayErr := c.DeenrollAll(ctx, enrollee); delayErr != nil {
			panic(delayErr)
		}
	}()

	select {
	case event := <-eventChan:
		return event.Data, nil
	case <-ctx.Done():
		return nil, errors.New("REDACTED")
	}
}
