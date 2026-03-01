package txpool

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	ifacemimic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer/simulations"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func VerifyApplicationTxpool(t *testing.T) {
	tx := func(v string) kinds.Tx { return kinds.Tx(v) }

	t.Run("REDACTED", func(t *testing.T) {
		//
		appended := atomic.Uint64{}

		//
		app := ifacemimic.FreshCustomer(t)
		app.
			On("REDACTED", mock.Anything, mock.Anything).
			Return(func(_ context.Context, req *iface.SolicitAppendTransfer) (*iface.ReplyAppendTransfer, error) {
				if string(req.Tx) == "REDACTED" {
					t.Logf("REDACTED")
					return &iface.ReplyAppendTransfer{Cipher: iface.CipherKindReissue}, nil
				}

				appended.Add(1)
				return &iface.ReplyAppendTransfer{Cipher: iface.CipherKindOKAY}, nil
			})

		//
		m := FreshApplicationTxpool(settings.FallbackTxpoolSettings(), app)

		//
		txs := []kinds.Tx{tx("REDACTED"), tx("REDACTED"), tx("REDACTED"), tx("REDACTED")}

		//
		faultone := m.AppendTransfer(txs[0])
		fault2 := m.AppendTransfer(txs[1])
		fault3 := m.AppendTransfer(txs[0]) //
		fault4 := m.AppendTransfer(txs[2]) //
		fault5 := m.AppendTransfer(txs[3]) //

		//
		require.NoError(t, faultone)
		require.NoError(t, fault2)

		require.ErrorIs(t, fault3, FaultObservedTransfer)
		require.ErrorIs(t, fault4, FaultBlankTransfer)

		require.ErrorContains(t, fault5, "REDACTED")
		require.False(t, m.observed.Has(txs[3]), "REDACTED")

		require.Equal(t, uint64(2), appended.Load())

		t.Run("REDACTED", func(t *testing.T) {
			for _, tt := range []struct {
				alias        string
				tx          kinds.Tx
				faultIncludes string
				negativeReact  bool
				affirm      func(t *testing.T, res *iface.ReplyInspectTransfer)
			}{
				{
					alias:        "REDACTED",
					tx:          tx("REDACTED"),
					faultIncludes: "REDACTED",
				},
				{
					alias: "REDACTED",
					tx:   tx("REDACTED"),
					affirm: func(t *testing.T, res *iface.ReplyInspectTransfer) {
						require.Equal(t, iface.CipherKindReissue, res.Cipher)
					},
				},
				{
					alias: "REDACTED",
					tx:   tx("REDACTED"),
					affirm: func(t *testing.T, res *iface.ReplyInspectTransfer) {
						require.Equal(t, iface.CipherKindOKAY, res.Cipher)
					},
				},
				{
					alias: "REDACTED",
					tx:   tx("REDACTED"),
				},
			} {
				t.Run(tt.alias, func(t *testing.T) {
					//
					var (
						outcome       = atomic.Pointer[iface.ReplyInspectTransfer]{}
						clbk     = func(res *iface.ReplyInspectTransfer) { outcome.Store(res) }
						assureOutcome = func() bool { return outcome.Load() != nil }
					)

					if tt.negativeReact {
						clbk = nil
					}

					//
					err := m.InspectTransfer(tt.tx, clbk, TransferDetails{})

					//
					if tt.faultIncludes != "REDACTED" {
						require.ErrorContains(t, err, tt.faultIncludes)
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
		const invocationsTowardAbort = 4

		//
		ctx, abort := context.WithCancel(context.Background())
		invocations := atomic.Uint64{}

		//
		everyTxpoolTrans := [][]byte{}

		app := ifacemimic.FreshCustomer(t)
		app.
			On("REDACTED", mock.Anything, mock.Anything).
			Return(func(_ context.Context, _ *iface.SolicitHarvestTrans) (*iface.ReplyHarvestTrans, error) {
				txs := make([][]byte, 0, quantity)
				for i := 0; i < quantity; i++ {
					txs = append(txs, []byte(fmt.Sprintf("REDACTED", i)))
				}

				everyTxpoolTrans = append(everyTxpoolTrans, txs...)

				invocations.Add(1)
				if invocations.Load() == invocationsTowardAbort {
					abort()
				}

				return &iface.ReplyHarvestTrans{Txs: txs}, nil
			})

		//
		m := FreshApplicationTxpool(settings.FallbackTxpoolSettings(), app)

		//
		//
		receiver := [][]byte{}
		ch := m.TransferInflux(ctx)

		for txs := range ch {
			receiver = append(receiver, txs.TowardSegmentBelongingOctets()...)
		}

		require.Subset(t, everyTxpoolTrans, receiver)
	})
}
