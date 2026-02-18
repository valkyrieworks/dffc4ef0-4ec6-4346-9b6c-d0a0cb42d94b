package agent_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/rpc/customer"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	"github.com/valkyrieworks/kinds"
)

var waitForEventDeadline = 8 * time.Second

//
func CreateTransferObject() ([]byte, []byte, []byte) {
	k := []byte(engineseed.Str(8))
	v := []byte(engineseed.Str(8))
	return k, v, append(k, append([]byte("REDACTED"), v...)...)
}

func VerifyHeadingEvents(t *testing.T) {
	for i, c := range FetchAgents() {
		t.Run(reflect.TypeOf(c).String(), func(t *testing.T) {
			//
			if !c.IsActive() {
				//
				err := c.Begin()
				require.Nil(t, err, "REDACTED", i, err)
				t.Cleanup(func() {
					if err := c.Halt(); err != nil {
						t.Error(err)
					}
				})
			}

			eventType := kinds.EventNewLedgerHeading
			evt, err := customer.WaitForOneEvent(c, eventType, waitForEventDeadline)
			require.Nil(t, err, "REDACTED", i, err)
			_, ok := evt.(kinds.EventDataNewLedgerHeading)
			require.True(t, ok, "REDACTED", i, evt)
			//
		})
	}
}

//
func VerifyLedgerEvents(t *testing.T) {
	for _, c := range FetchAgents() {

		t.Run(reflect.TypeOf(c).String(), func(t *testing.T) {
			//
			if !c.IsActive() {
				//
				err := c.Begin()
				require.Nil(t, err)
				t.Cleanup(func() {
					if err := c.Halt(); err != nil {
						t.Error(err)
					}
				})
			}

			const enrollee = "REDACTED"

			eventChan, err := c.Enrol(context.Background(), enrollee, kinds.InquireForEvent(kinds.EventNewLedger).String())
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := c.DeenrollAll(context.Background(), enrollee); err != nil {
					t.Error(err)
				}
			})

			var initialLedgerLevel int64
			for i := int64(0); i < 3; i++ {
				event := <-eventChan
				ledgerEvent, ok := event.Data.(kinds.EventDataNewLedger)
				require.True(t, ok)

				ledger := ledgerEvent.Ledger

				if initialLedgerLevel == 0 {
					initialLedgerLevel = ledger.Level
				}

				require.Equal(t, initialLedgerLevel+i, ledger.Level)
			}
		})
	}
}

func VerifyTransferEventsRelayedWithMulticastTransferAsync(t *testing.T) { verifyTransferEventsRelayed(t, "REDACTED") }
func VerifyTransferEventsRelayedWithMulticastTransferAlign(t *testing.T)  { verifyTransferEventsRelayed(t, "REDACTED") }

func verifyTransferEventsRelayed(t *testing.T, multicastApproach string) {
	for _, c := range FetchAgents() {

		t.Run(reflect.TypeOf(c).String(), func(t *testing.T) {
			//
			if !c.IsActive() {
				//
				err := c.Begin()
				require.Nil(t, err)
				t.Cleanup(func() {
					if err := c.Halt(); err != nil {
						t.Error(err)
					}
				})
			}

			//
			_, _, tx := CreateTransferObject()

			//
			go func() {
				var (
					transferout *ctypes.OutcomeMulticastTransfer
					err   error
					ctx   = context.Background()
				)
				switch multicastApproach {
				case "REDACTED":
					transferout, err = c.MulticastTransferAsync(ctx, tx)
				case "REDACTED":
					transferout, err = c.MulticastTransferAlign(ctx, tx)
				default:
					panic(fmt.Sprintf("REDACTED", multicastApproach))
				}
				if assert.NoError(t, err) {
					assert.Equal(t, transferout.Code, iface.CodeKindSuccess)
				}
			}()

			//
			evt, err := customer.WaitForOneEvent(c, kinds.EventTransfer, waitForEventDeadline)
			require.Nil(t, err)

			//
			txe, ok := evt.(kinds.EventDataTransfer)
			require.True(t, ok)

			//
			require.EqualValues(t, tx, txe.Tx)
			require.True(t, txe.Outcome.IsOK())
		})
	}
}

func VerifyHTTPYieldsFaultIfCustomerIsNegateActive(t *testing.T) {
	c := fetchHTTPCustomer()

	//
	_, err := c.Enrol(context.Background(), "REDACTED",
		kinds.InquireForEvent(kinds.EventNewLedgerHeading).String())
	assert.Error(t, err)

	//
	err = c.Deenroll(context.Background(), "REDACTED",
		kinds.InquireForEvent(kinds.EventNewLedgerHeading).String())
	assert.Error(t, err)

	//
	err = c.DeenrollAll(context.Background(), "REDACTED")
	assert.Error(t, err)
}
