package integration_t_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
)

//
func Verifyblock_Heading(t *testing.T) {
	ledgers := acquireLedgerSeries(t)
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		if member.Style == e2e.StyleOrigin {
			return
		}

		customer, err := member.Customer()
		require.NoError(t, err)
		state, err := customer.Status(ctx)
		require.NoError(t, err)

		initial := state.AlignDetails.OldestLedgerLevel
		final := state.AlignDetails.NewestLedgerLevel
		if member.PreserveLedgers > 0 {
			initial++ //
		}

		for _, ledger := range ledgers {
			if ledger.Level < initial {
				continue
			}
			if ledger.Level > final {
				break
			}
			reply, err := customer.Ledger(ctx, &ledger.Level)
			require.NoError(t, err)

			require.Equal(t, ledger, reply.Ledger,
				"REDACTED", ledger.Level)

			require.NoError(t, reply.Ledger.CertifySimple(),
				"REDACTED", ledger.Level)
		}
	})
}

//
func Verifyblock_Scope(t *testing.T) {
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		if member.Style == e2e.StyleOrigin {
			return
		}

		customer, err := member.Customer()
		require.NoError(t, err)
		state, err := customer.Status(ctx)
		require.NoError(t, err)

		initial := state.AlignDetails.OldestLedgerLevel
		final := state.AlignDetails.NewestLedgerLevel

		switch {
		case member.StatusAlign:
			assert.Greater(t, initial, member.Verifychain.PrimaryLevel,
				"REDACTED")

		case member.PreserveLedgers > 0 && int64(member.PreserveLedgers) < (final-member.Verifychain.PrimaryLevel+1):
			//
			assert.InDelta(t, member.PreserveLedgers, final-initial+1, 1,
				"REDACTED")

		default:
			assert.Equal(t, member.Verifychain.PrimaryLevel, initial,
				"REDACTED")
		}

		for h := initial; h <= final; h++ {
			reply, err := customer.Ledger(ctx, &(h))
			if err != nil && member.PreserveLedgers > 0 && h == initial {
				//
				continue
			}
			require.NoError(t, err)
			assert.Equal(t, h, reply.Ledger.Level)
		}

		for h := member.Verifychain.PrimaryLevel; h < initial; h++ {
			_, err := customer.Ledger(ctx, &(h))
			require.Error(t, err)
		}
	})
}
