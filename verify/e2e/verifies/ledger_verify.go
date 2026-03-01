package end2end_typ_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
)

//
func Testledger_Headline(t *testing.T) {
	ledgers := acquireLedgerSuccession(t)
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		if peer.Style == e2e.StyleGerm {
			return
		}

		customer, err := peer.Customer()
		require.NoError(t, err)
		condition, err := customer.Condition(ctx)
		require.NoError(t, err)

		initial := condition.ChronizeDetails.InitialLedgerAltitude
		final := condition.ChronizeDetails.NewestLedgerAltitude
		if peer.PreserveLedgers > 0 {
			initial++ //
		}

		for _, ledger := range ledgers {
			if ledger.Altitude < initial {
				continue
			}
			if ledger.Altitude > final {
				break
			}
			reply, err := customer.Ledger(ctx, &ledger.Altitude)
			require.NoError(t, err)

			require.Equal(t, ledger, reply.Ledger,
				"REDACTED", ledger.Altitude)

			require.NoError(t, reply.Ledger.CertifyFundamental(),
				"REDACTED", ledger.Altitude)
		}
	})
}

//
func Testledger_Scope(t *testing.T) {
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		if peer.Style == e2e.StyleGerm {
			return
		}

		customer, err := peer.Customer()
		require.NoError(t, err)
		condition, err := customer.Condition(ctx)
		require.NoError(t, err)

		initial := condition.ChronizeDetails.InitialLedgerAltitude
		final := condition.ChronizeDetails.NewestLedgerAltitude

		switch {
		case peer.StatusChronize:
			assert.Greater(t, initial, peer.Simnet.PrimaryAltitude,
				"REDACTED")

		case peer.PreserveLedgers > 0 && int64(peer.PreserveLedgers) < (final-peer.Simnet.PrimaryAltitude+1):
			//
			assert.InDelta(t, peer.PreserveLedgers, final-initial+1, 1,
				"REDACTED")

		default:
			assert.Equal(t, peer.Simnet.PrimaryAltitude, initial,
				"REDACTED")
		}

		for h := initial; h <= final; h++ {
			reply, err := customer.Ledger(ctx, &(h))
			if err != nil && peer.PreserveLedgers > 0 && h == initial {
				//
				continue
			}
			require.NoError(t, err)
			assert.Equal(t, h, reply.Ledger.Altitude)
		}

		for h := peer.Simnet.PrimaryAltitude; h < initial; h++ {
			_, err := customer.Ledger(ctx, &(h))
			require.Error(t, err)
		}
	})
}
