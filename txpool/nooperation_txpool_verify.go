package txpool

import (
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/stretchr/testify/assert"
)

var tx = kinds.Tx([]byte{0x01})

func Verifyunpooledtrans_Fundamental(t *testing.T) {
	mem := &NooperationTxpool{}

	assert.Equal(t, 0, mem.Extent())
	assert.Equal(t, int64(0), mem.ExtentOctets())

	err := mem.InspectTransfer(tx, nil, TransferDetails{})
	assert.Equal(t, faultNegationPermitted, err)

	err = mem.DiscardTransferViaToken(tx.Key())
	assert.Equal(t, faultNegationPermitted, err)

	txs := mem.HarvestMaximumOctetsMaximumFuel(0, 0)
	assert.Nil(t, txs)

	txs = mem.HarvestMaximumTrans(0)
	assert.Nil(t, txs)

	err = mem.PurgeApplicationLink()
	assert.NoError(t, err)

	err = mem.Revise(0, nil, nil, nil, nil)
	assert.NoError(t, err)

	transAccessible := mem.TransAccessible()
	assert.Nil(t, transAccessible)
}
