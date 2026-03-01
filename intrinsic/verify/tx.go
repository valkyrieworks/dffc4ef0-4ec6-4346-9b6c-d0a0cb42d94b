package verify

import "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"

func CreateNTHTrans(altitude, n int64) kinds.Txs {
	txs := make([]kinds.Tx, n)
	for i := range txs {
		txs[i] = kinds.Tx([]byte{byte(altitude), byte(i / 256), byte(i % 256)})
	}
	return txs
}
