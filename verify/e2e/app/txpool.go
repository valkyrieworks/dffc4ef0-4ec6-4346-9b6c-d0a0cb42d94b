package app

import (
	"bytes"
	"fmt"
	"slices"
	"sync"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
type ApplicationTxpool struct {
	txs    map[string]kinds.Tx
	tracer log.Tracer
	mu     sync.RWMutex
}

//
func FreshApplicationTxpool(tracer log.Tracer) *ApplicationTxpool {
	return &ApplicationTxpool{
		txs:    make(map[string]kinds.Tx),
		tracer: tracer,
	}
}

func (m *ApplicationTxpool) AppendTransfer(bz []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	tx := kinds.Tx(bz)
	transferDigest := fmt.Sprintf("REDACTED", tx.Digest())

	if _, ok := m.txs[transferDigest]; ok {
		m.tracer.Details("REDACTED", "REDACTED", transferDigest)
	} else {
		m.txs[transferDigest] = tx
		m.tracer.Details("REDACTED", "REDACTED", transferDigest)
	}
}

func (m *ApplicationTxpool) HarvestTrans(purge bool) kinds.Txs {
	m.mu.Lock()
	defer m.mu.Unlock()

	txs := make([]kinds.Tx, 0, len(m.txs))
	for _, tx := range m.txs {
		txs = append(txs, tx)
	}

	slices.SortFunc(txs, func(a, b kinds.Tx) int {
		return bytes.Compare(a, b)
	})

	if purge {
		m.txs = make(map[string]kinds.Tx)
	}

	return txs
}
