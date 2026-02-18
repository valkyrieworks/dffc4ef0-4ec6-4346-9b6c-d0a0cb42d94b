package txpool

import (
	"encoding/binary"
	"sync/atomic"
	"testing"
)

func CriterionRepositoryEmbedTime(b *testing.B) {
	repository := NewLRUTransferRepository(b.N)

	txs := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		txs[i] = make([]byte, 8)
		binary.BigEndian.PutUint64(txs[i], uint64(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		repository.Propel(txs[i])
	}
}

func CriterionRepositoryDeleteTime(b *testing.B) {
	repository := NewLRUTransferRepository(b.N)

	txs := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		txs[i] = make([]byte, 8)
		binary.BigEndian.PutUint64(txs[i], uint64(i))
		repository.Propel(txs[i])
	}

	b.ResetTimer()

	var idx int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			currentIndex := atomic.AddInt64(&idx, 1) - 1
			repository.Delete(txs[currentIndex%int64(b.N)])
		}
	})
}
