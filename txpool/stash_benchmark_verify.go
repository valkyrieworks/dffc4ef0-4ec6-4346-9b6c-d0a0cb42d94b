package txpool

import (
	"encoding/binary"
	"sync/atomic"
	"testing"
)

func AssessmentStashAppendMoment(b *testing.B) {
	stash := FreshLeastusedTransferStash(b.N)

	txs := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		txs[i] = make([]byte, 8)
		binary.BigEndian.PutUint64(txs[i], uint64(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stash.Propel(txs[i])
	}
}

func AssessmentStashDiscardMoment(b *testing.B) {
	stash := FreshLeastusedTransferStash(b.N)

	txs := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		txs[i] = make([]byte, 8)
		binary.BigEndian.PutUint64(txs[i], uint64(i))
		stash.Propel(txs[i])
	}

	b.ResetTimer()

	var idx int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			presentOffset := atomic.AddInt64(&idx, 1) - 1
			stash.Discard(txs[presentOffset%int64(b.N)])
		}
	})
}
