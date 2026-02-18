package verify

import "github.com/valkyrieworks/kinds"

func CreateNTrans(level, n int64) kinds.Txs {
	txs := make([]kinds.Tx, n)
	for i := range txs {
		txs[i] = kinds.Tx([]byte{byte(level), byte(i / 256), byte(i % 256)})
	}
	return txs
}
