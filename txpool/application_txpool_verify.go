package txpool

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	abciemulate "github.com/valkyrieworks/iface/customer/simulations"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/kinds"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func VerifyApplicationTxpool(t *testing.T) {
	tx := func(v string) kinds.Tx { return kinds.Tx(v) }

	t.Run("REDACTED", func(t *testing.T) {
		//
		appended := atomic.Uint64{}

		//
		app := abciemulate.NewCustomer(t)
		app.
			On("REDACTED", mock.Anything, mock.Anything).
			Return(func(_ context.Context, req *iface.QueryEmbedTransfer) (*iface.ReplyEmbedTransfer, error) {
				if string(req.Tx) == "REDACTED" {
					t.Logf("REDACTED")
					return &iface.ReplyEmbedTransfer{Code: iface.CodeKindReprocess}, nil
				}

				appended.Add(1)
				return &iface.ReplyEmbedTransfer{Code: iface.CodeKindSuccess}, nil
			})

		//
		m := NewApplicationTxpool(settings.StandardTxpoolSettings(), app)

		//
		txs := []kinds.Tx{tx("REDACTED"), tx("REDACTED"), tx("REDACTED"), tx("REDACTED")}

		//
		fault1 := m.EmbedTransfer(txs[0])
		err2 := m.EmbedTransfer(txs[1])
		err3 := m.EmbedTransfer(txs[0]) //
		err4 := m.EmbedTransfer(txs[2]) //
		fault5 := m.EmbedTransfer(txs[3]) //

		//
		require.NoError(t, fault1)
		require.NoError(t, err2)

		require.ErrorIs(t, err3, ErrViewedTransfer)
		require.ErrorIs(t, err4, ErrEmptyTransfer)

		require.ErrorContains(t, fault5, "REDACTED")
		require.False(t, m.viewed.Has(txs[3]), "REDACTED")

		require.Equal(t, uint64(2), appended.Load())

		t.Run("REDACTED", func(t *testing.T) {
			for _, tt := range []struct {
				label        string
				tx          kinds.Tx
				errIncludes string
				noResponse  bool
				affirm      func(t *testing.T, res *iface.ReplyInspectTransfer)
			}{
				{
					label:        "REDACTED",
					tx:          tx("REDACTED"),
					errIncludes: "REDACTED",
				},
				{
					label: "REDACTED",
					tx:   tx("REDACTED"),
					affirm: func(t *testing.T, res *iface.ReplyInspectTransfer) {
						require.Equal(t, iface.CodeKindReprocess, res.Code)
					},
				},
				{
					label: "REDACTED",
					tx:   tx("REDACTED"),
					affirm: func(t *testing.T, res *iface.ReplyInspectTransfer) {
						require.Equal(t, iface.CodeKindSuccess, res.Code)
					},
				},
				{
					label: "REDACTED",
					tx:   tx("REDACTED"),
				},
			} {
				t.Run(tt.label, func(t *testing.T) {
					//
					var (
						outcome       = atomic.Pointer[iface.ReplyInspectTransfer]{}
						callback     = func(res *iface.ReplyInspectTransfer) { outcome.Store(res) }
						assureOutcome = func() bool { return outcome.Load() != nil }
					)

					if tt.noResponse {
						callback = nil
					}

					//
					err := m.InspectTransfer(tt.tx, callback, TransferDetails{})

					//
					if tt.errIncludes != "REDACTED" {
						require.ErrorContains(t, err, tt.errIncludes)
						return
					}

					require.NoError(t, err)
					require.Eventually(t, assureOutcome, time.Second, time.Millisecond*50)

					if tt.affirm != nil {
						tt.affirm(t, outcome.Load())
					}
				})
			}
		})
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		const quantity = 100
		const invocationsToRevoke = 4

		//
		ctx, revoke := context.WithCancel(context.Background())
		invocations := atomic.Uint64{}

		//
		allTxpoolTrans := [][]byte{}

		app := abciemulate.NewCustomer(t)
		app.
			On("REDACTED", mock.Anything, mock.Anything).
			Return(func(_ context.Context, _ *iface.QueryHarvestTrans) (*iface.ReplyHarvestTrans, error) {
				txs := make([][]byte, 0, quantity)
				for i := 0; i < quantity; i++ {
					txs = append(txs, []byte(fmt.Sprintf("REDACTED", i)))
				}

				allTxpoolTrans = append(allTxpoolTrans, txs...)

				invocations.Add(1)
				if invocations.Load() == invocationsToRevoke {
					revoke()
				}

				return &iface.ReplyHarvestTrans{Txs: txs}, nil
			})

		//
		m := NewApplicationTxpool(settings.StandardTxpoolSettings(), app)

		//
		//
		drain := [][]byte{}
		ch := m.TransferFlow(ctx)

		for txs := range ch {
			drain = append(drain, txs.ToSegmentOfOctets()...)
		}

		require.Subset(t, allTxpoolTrans, drain)
	})
}
