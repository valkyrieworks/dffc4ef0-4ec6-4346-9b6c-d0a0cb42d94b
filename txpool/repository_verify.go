package txpool

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/gateway"
	"github.com/valkyrieworks/kinds"
	"github.com/stretchr/testify/require"
)

func VerifyRepositoryDelete(t *testing.T) {
	repository := NewLRUTransferRepository(100)
	countTrans := 10

	txs := make([][]byte, countTrans)
	for i := 0; i < countTrans; i++ {
		//
		transferOctets := make([]byte, 32)
		_, err := rand.Read(transferOctets)
		require.NoError(t, err)

		txs[i] = transferOctets
		repository.Propel(transferOctets)

		//
		require.Equal(t, i+1, len(repository.repositoryIndex))
		require.Equal(t, i+1, repository.catalog.Len())
	}

	for i := 0; i < countTrans; i++ {
		repository.Delete(txs[i])
		//
		require.Equal(t, countTrans-(i+1), len(repository.repositoryIndex))
		require.Equal(t, countTrans-(i+1), repository.catalog.Len())
	}
}

func VerifyRepositoryAfterModify(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	//
	//
	//
	verifies := []struct {
		countTransToInstantiate int
		modifyOrdinals  []int
		reAppendOrdinals   []int
		transInRepository     []int
	}{
		{1, []int{}, []int{1}, []int{1, 0}},    //
		{2, []int{1}, []int{}, []int{1, 0}},    //
		{2, []int{2}, []int{}, []int{2, 1, 0}}, //
		{2, []int{1}, []int{1}, []int{1, 0}},   //
	}
	for tcOrdinal, tc := range verifies {
		for i := 0; i < tc.countTransToInstantiate; i++ {
			tx := objectdepot.NewTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
			err := mp.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				require.False(t, reply.IsErr())
			}, TransferDetails{})
			require.NoError(t, err)
		}

		modifyTrans := []kinds.Tx{}
		for _, v := range tc.modifyOrdinals {
			tx := objectdepot.NewTransfer(fmt.Sprintf("REDACTED", v), "REDACTED")
			modifyTrans = append(modifyTrans, tx)
		}
		err := mp.Modify(int64(tcOrdinal), modifyTrans, ifaceReplies(len(modifyTrans), iface.CodeKindSuccess), nil, nil)
		require.NoError(t, err)

		for _, v := range tc.reAppendOrdinals {
			tx := objectdepot.NewTransfer(fmt.Sprintf("REDACTED", v), "REDACTED")
			_ = mp.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				require.False(t, reply.IsErr())
			}, TransferDetails{})
		}

		repository := mp.repository.(*LRUTransferRepository)
		member := repository.FetchCatalog().Front()
		tally := 0
		for member != nil {
			require.NotEqual(t, len(tc.transInRepository), tally,
				"REDACTED", tcOrdinal)

			memberValue := member.Value.(kinds.TransferKey)
			expirationTransfer := objectdepot.NewTransfer(fmt.Sprintf("REDACTED", tc.transInRepository[len(tc.transInRepository)-tally-1]), "REDACTED")
			anticipatedBz := sha256.Sum256(expirationTransfer)
			//
			//
			//
			//
			//
			//
			//

			require.EqualValues(t, anticipatedBz, memberValue, "REDACTED", tally, tcOrdinal)
			tally++
			member = member.Next()
		}
		require.Equal(t, len(tc.transInRepository), tally,
			"REDACTED", tcOrdinal)
		mp.Purge()
	}
}
