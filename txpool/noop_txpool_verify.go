package txpool

import (
	"testing"

	"github.com/valkyrieworks/kinds"
	"github.com/stretchr/testify/assert"
)

var tx = kinds.Tx([]byte{0x01})

func Verifynomempool_Simple(t *testing.T) {
	mem := &NoopTxpool{}

	assert.Equal(t, 0, mem.Volume())
	assert.Equal(t, int64(0), mem.VolumeOctets())

	err := mem.InspectTransfer(tx, nil, TransferDetails{})
	assert.Equal(t, errNegatePermitted, err)

	err = mem.DeleteTransferByKey(tx.Key())
	assert.Equal(t, errNegatePermitted, err)

	txs := mem.HarvestMaximumOctetsMaximumFuel(0, 0)
	assert.Nil(t, txs)

	txs = mem.HarvestMaximumTrans(0)
	assert.Nil(t, txs)

	err = mem.PurgeApplicationLink()
	assert.NoError(t, err)

	err = mem.Modify(0, nil, nil, nil, nil)
	assert.NoError(t, err)

	transAccessible := mem.TransAccessible()
	assert.Nil(t, transAccessible)
}
