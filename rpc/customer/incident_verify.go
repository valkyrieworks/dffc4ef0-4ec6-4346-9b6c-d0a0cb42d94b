package cust_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var pauseForeachIncidentDeadline = 8 * time.Second

//
func CreateTransferTokval() ([]byte, []byte, []byte) {
	k := []byte(commitrand.Str(8))
	v := []byte(commitrand.Str(8))
	return k, v, append(k, append([]byte("REDACTED"), v...)...)
}

func VerifyHeadlineIncidents(t *testing.T) {
	for i, c := range FetchCustomers() {
		t.Run(reflect.TypeOf(c).String(), func(t *testing.T) {
			//
			if !c.EqualsActive() {
				//
				err := c.Initiate()
				require.Nil(t, err, "REDACTED", i, err)
				t.Cleanup(func() {
					if err := c.Halt(); err != nil {
						t.Error(err)
					}
				})
			}

			signalKind := kinds.IncidentFreshLedgerHeading
			evt, err := customer.PauseForeachSingleIncident(c, signalKind, pauseForeachIncidentDeadline)
			require.Nil(t, err, "REDACTED", i, err)
			_, ok := evt.(kinds.IncidentDataFreshLedgerHeading)
			require.True(t, ok, "REDACTED", i, evt)
			//
		})
	}
}

//
func VerifyLedgerIncidents(t *testing.T) {
	for _, c := range FetchCustomers() {

		t.Run(reflect.TypeOf(c).String(), func(t *testing.T) {
			//
			if !c.EqualsActive() {
				//
				err := c.Initiate()
				require.Nil(t, err)
				t.Cleanup(func() {
					if err := c.Halt(); err != nil {
						t.Error(err)
					}
				})
			}

			const listener = "REDACTED"

			incidentChnl, err := c.Listen(context.Background(), listener, kinds.InquireForeachIncident(kinds.IncidentFreshLedger).Text())
			require.NoError(t, err)
			t.Cleanup(func() {
				if err := c.UnlistenEvery(context.Background(), listener); err != nil {
					t.Error(err)
				}
			})

			var initialLedgerAltitude int64
			for i := int64(0); i < 3; i++ {
				incident := <-incidentChnl
				ledgerIncident, ok := incident.Data.(kinds.IncidentDataFreshLedger)
				require.True(t, ok)

				ledger := ledgerIncident.Ledger

				if initialLedgerAltitude == 0 {
					initialLedgerAltitude = ledger.Altitude
				}

				require.Equal(t, initialLedgerAltitude+i, ledger.Altitude)
			}
		})
	}
}

func VerifyTransferIncidentsRelayedUsingMulticastTransferAsyncronous(t *testing.T) { verifyTransferIncidentsRelayed(t, "REDACTED") }
func VerifyTransferIncidentsRelayedUsingMulticastTransferChronize(t *testing.T)  { verifyTransferIncidentsRelayed(t, "REDACTED") }

func verifyTransferIncidentsRelayed(t *testing.T, multicastApproach string) {
	for _, c := range FetchCustomers() {

		t.Run(reflect.TypeOf(c).String(), func(t *testing.T) {
			//
			if !c.EqualsActive() {
				//
				err := c.Initiate()
				require.Nil(t, err)
				t.Cleanup(func() {
					if err := c.Halt(); err != nil {
						t.Error(err)
					}
				})
			}

			//
			_, _, tx := CreateTransferTokval()

			//
			go func() {
				var (
					transferresp *ktypes.OutcomeMulticastTransfer
					err   error
					ctx   = context.Background()
				)
				switch multicastApproach {
				case "REDACTED":
					transferresp, err = c.MulticastTransferAsyncronous(ctx, tx)
				case "REDACTED":
					transferresp, err = c.MulticastTransferChronize(ctx, tx)
				default:
					panic(fmt.Sprintf("REDACTED", multicastApproach))
				}
				if assert.NoError(t, err) {
					assert.Equal(t, transferresp.Cipher, iface.CipherKindOKAY)
				}
			}()

			//
			evt, err := customer.PauseForeachSingleIncident(c, kinds.IncidentTransfer, pauseForeachIncidentDeadline)
			require.Nil(t, err)

			//
			txe, ok := evt.(kinds.IncidentDataTransfer)
			require.True(t, ok)

			//
			require.EqualValues(t, tx, txe.Tx)
			require.True(t, txe.Outcome.EqualsOKAY())
		})
	}
}

func VerifyHttpsvcYieldsFailureConditionalCustomerEqualsNegationActive(t *testing.T) {
	c := fetchHttpsvcCustomer()

	//
	_, err := c.Listen(context.Background(), "REDACTED",
		kinds.InquireForeachIncident(kinds.IncidentFreshLedgerHeading).Text())
	assert.Error(t, err)

	//
	err = c.Unlisten(context.Background(), "REDACTED",
		kinds.InquireForeachIncident(kinds.IncidentFreshLedgerHeading).Text())
	assert.Error(t, err)

	//
	err = c.UnlistenEvery(context.Background(), "REDACTED")
	assert.Error(t, err)
}
