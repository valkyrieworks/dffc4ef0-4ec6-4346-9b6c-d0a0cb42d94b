package txpool

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/stretchr/testify/require"
)

func VerifyStashDiscard(t *testing.T) {
	stash := FreshLeastusedTransferStash(100)
	countTrans := 10

	txs := make([][]byte, countTrans)
	for i := 0; i < countTrans; i++ {
		//
		transferOctets := make([]byte, 32)
		_, err := rand.Read(transferOctets)
		require.NoError(t, err)

		txs[i] = transferOctets
		stash.Propel(transferOctets)

		//
		require.Equal(t, i+1, len(stash.stashIndex))
		require.Equal(t, i+1, stash.catalog.Len())
	}

	for i := 0; i < countTrans; i++ {
		stash.Discard(txs[i])
		//
		require.Equal(t, countTrans-(i+1), len(stash.stashIndex))
		require.Equal(t, countTrans-(i+1), stash.catalog.Len())
	}
}

func VerifyStashSubsequentRevise(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	//
	//
	//
	verifies := []struct {
		countTransTowardGenerate int
		reviseIndexes  []int
		againAppendIndexes   []int
		transInsideStash     []int
	}{
		{1, []int{}, []int{1}, []int{1, 0}},    //
		{2, []int{1}, []int{}, []int{1, 0}},    //
		{2, []int{2}, []int{}, []int{2, 1, 0}}, //
		{2, []int{1}, []int{1}, []int{1, 0}},   //
	}
	for testcasePosition, tc := range verifies {
		for i := 0; i < tc.countTransTowardGenerate; i++ {
			tx := statedepot.FreshTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
			err := mp.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				require.False(t, reply.EqualsFault())
			}, TransferDetails{})
			require.NoError(t, err)
		}

		reviseTrans := []kinds.Tx{}
		for _, v := range tc.reviseIndexes {
			tx := statedepot.FreshTransfer(fmt.Sprintf("REDACTED", v), "REDACTED")
			reviseTrans = append(reviseTrans, tx)
		}
		err := mp.Revise(int64(testcasePosition), reviseTrans, ifaceReplies(len(reviseTrans), iface.CipherKindOKAY), nil, nil)
		require.NoError(t, err)

		for _, v := range tc.againAppendIndexes {
			tx := statedepot.FreshTransfer(fmt.Sprintf("REDACTED", v), "REDACTED")
			_ = mp.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				require.False(t, reply.EqualsFault())
			}, TransferDetails{})
		}

		stash := mp.stash.(*LeastusedTransferStash)
		peer := stash.ObtainCatalog().Front()
		tally := 0
		for peer != nil {
			require.NotEqual(t, len(tc.transInsideStash), tally,
				"REDACTED", testcasePosition)

			peerItem := peer.Value.(kinds.TransferToken)
			expirationTransfer := statedepot.FreshTransfer(fmt.Sprintf("REDACTED", tc.transInsideStash[len(tc.transInsideStash)-tally-1]), "REDACTED")
			anticipatedByz := sha256.Sum256(expirationTransfer)
			//
			//
			//
			//
			//
			//
			//

			require.EqualValues(t, anticipatedByz, peerItem, "REDACTED", tally, testcasePosition)
			tally++
			peer = peer.Next()
		}
		require.Equal(t, len(tc.transInsideStash), tally,
			"REDACTED", testcasePosition)
		mp.Purge()
	}
}
