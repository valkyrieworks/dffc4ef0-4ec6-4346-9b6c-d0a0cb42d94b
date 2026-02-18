package txpool

import (
	"sync/atomic"
	"testing"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/gateway"
	"github.com/stretchr/testify/require"
)

func CriterionHarvest(b *testing.B) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	mp.settings.Volume = 100_000_000 //
	appendTrans(b, mp, 0, 10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mp.HarvestMaximumOctetsMaximumFuel(100_000_000, -1)
	}
}

func CriterionInspectTransfer(b *testing.B) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	mp.settings.Volume = 100_000_000
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tx := objectdepot.NewTransferFromUID(i)
		b.StartTimer()

		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(b, err, i)
	}
}

func CriterionConcurrentInspectTransfer(b *testing.B) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	mp.settings.Volume = 100_000_000
	var transcount uint64
	following := func() uint64 {
		return atomic.AddUint64(&transcount, 1)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tx := objectdepot.NewTransferFromUID(int(following()))
			err := mp.InspectTransfer(tx, nil, TransferDetails{})
			require.NoError(b, err, tx)
		}
	})
}

func CriterionInspectReplicatedTransfer(b *testing.B) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	mp.settings.Volume = 2

	tx := objectdepot.NewTransferFromUID(1)
	if err := mp.InspectTransfer(tx, nil, TransferDetails{}); err != nil {
		b.Fatal(err)
	}
	err := mp.PurgeApplicationLink()
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.ErrorAs(b, err, &ErrTransferInRepository, "REDACTED")
	}
}

func CriterionModify(b *testing.B) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	countTrans := 1000
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		txs := appendTrans(b, mp, i*countTrans, countTrans)
		require.Equal(b, len(txs), mp.Volume(), len(txs))
		b.StartTimer()

		doModify(b, mp, int64(i), txs)
		require.Zero(b, mp.Volume())
	}
}

func CriterionModifyAndRevalidate(b *testing.B) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	countTrans := 1000
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		mp.Purge()
		txs := appendTrans(b, mp, i*countTrans, countTrans)
		require.Equal(b, len(txs), mp.Volume(), len(txs))
		b.StartTimer()

		//
		doModify(b, mp, int64(i), txs[:countTrans/2])
	}
}

func CriterionModifyExternalCustomer(b *testing.B) {
	mp, sanitize := newTxpoolWithAsyncLinkage(b)
	defer sanitize()

	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		tx := objectdepot.NewTransferFromUID(i)
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(b, err)
		err = mp.PurgeApplicationLink()
		require.NoError(b, err)
		require.Equal(b, 1, mp.Volume())
		b.StartTimer()

		txs := mp.HarvestMaximumTrans(mp.Volume())
		doModify(b, mp, int64(i), txs)
	}
}
