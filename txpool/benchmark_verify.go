package txpool

import (
	"sync/atomic"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	"github.com/stretchr/testify/require"
)

func AssessmentHarvest(b *testing.B) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	mp.settings.Extent = 100_000_000 //
	appendTrans(b, mp, 0, 10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mp.HarvestMaximumOctetsMaximumFuel(100_000_000, -1)
	}
}

func AssessmentInspectTransfer(b *testing.B) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	mp.settings.Extent = 100_000_000
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tx := statedepot.FreshTransferOriginatingUUID(i)
		b.StartTimer()

		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(b, err, i)
	}
}

func AssessmentConcurrentInspectTransfer(b *testing.B) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	mp.settings.Extent = 100_000_000
	var transcount uint64
	following := func() uint64 {
		return atomic.AddUint64(&transcount, 1)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tx := statedepot.FreshTransferOriginatingUUID(int(following()))
			err := mp.InspectTransfer(tx, nil, TransferDetails{})
			require.NoError(b, err, tx)
		}
	})
}

func AssessmentInspectReplicatedTransfer(b *testing.B) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	mp.settings.Extent = 2

	tx := statedepot.FreshTransferOriginatingUUID(1)
	if err := mp.InspectTransfer(tx, nil, TransferDetails{}); err != nil {
		b.Fatal(err)
	}
	err := mp.PurgeApplicationLink()
	require.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.ErrorAs(b, err, &FaultTransferInsideStash, "REDACTED")
	}
}

func AssessmentRevise(b *testing.B) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	countTrans := 1000
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		txs := appendTrans(b, mp, i*countTrans, countTrans)
		require.Equal(b, len(txs), mp.Extent(), len(txs))
		b.StartTimer()

		conductRevise(b, mp, int64(i), txs)
		require.Zero(b, mp.Extent())
	}
}

func AssessmentReviseAlsoReinspect(b *testing.B) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	countTrans := 1000
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		mp.Purge()
		txs := appendTrans(b, mp, i*countTrans, countTrans)
		require.Equal(b, len(txs), mp.Extent(), len(txs))
		b.StartTimer()

		//
		conductRevise(b, mp, int64(i), txs[:countTrans/2])
	}
}

func AssessmentReviseDistantCustomer(b *testing.B) {
	mp, sanitize := freshTxpoolUsingAsyncronousLinkage(b)
	defer sanitize()

	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		tx := statedepot.FreshTransferOriginatingUUID(i)
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(b, err)
		err = mp.PurgeApplicationLink()
		require.NoError(b, err)
		require.Equal(b, 1, mp.Extent())
		b.StartTimer()

		txs := mp.HarvestMaximumTrans(mp.Extent())
		conductRevise(b, mp, int64(i), txs)
	}
}
