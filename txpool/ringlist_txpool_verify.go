package txpool

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	mrand "math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	ifacecustomer "github.com/valkyrieworks/iface/customer"
	abcicliclients "github.com/valkyrieworks/iface/customer/simulations"
	"github.com/valkyrieworks/iface/instance/objectdepot"
	ifaceservice "github.com/valkyrieworks/iface/host"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/gateway"
	"github.com/valkyrieworks/kinds"
)

//
//
type sanitizeFunction func()

func newTxpoolWithApplicationEmulate(customer ifacecustomer.Customer) (*CCatalogTxpool, sanitizeFunction, error) {
	cfg := verify.RestoreVerifyOrigin("REDACTED")

	mp, cu := newTxpoolWithApplicationAndSettingsEmulate(cfg, customer)
	return mp, cu, nil
}

func newTxpoolWithApplicationAndSettingsEmulate(
	cfg *settings.Settings,
	customer ifacecustomer.Customer,
) (*CCatalogTxpool, sanitizeFunction) {
	applicationLinkMemory := customer
	applicationLinkMemory.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err := applicationLinkMemory.Begin()
	if err != nil {
		panic(err)
	}

	mp := NewCCatalogTxpool(cfg.Txpool, applicationLinkMemory, 0)
	mp.AssignTracer(log.VerifyingTracer())

	return mp, func() { os.RemoveAll(cfg.OriginFolder) }
}

func newTxpoolWithApplication(cc gateway.CustomerOriginator) (*CCatalogTxpool, sanitizeFunction) {
	cfg := verify.RestoreVerifyOrigin("REDACTED")

	mp, cu := newTxpoolWithApplicationAndSettings(cc, cfg)
	return mp, cu
}

func newTxpoolWithApplicationAndSettings(cc gateway.CustomerOriginator, cfg *settings.Settings) (*CCatalogTxpool, sanitizeFunction) {
	applicationLinkMemory, _ := cc.NewIfaceCustomer()
	applicationLinkMemory.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err := applicationLinkMemory.Begin()
	if err != nil {
		panic(err)
	}

	mp := NewCCatalogTxpool(cfg.Txpool, applicationLinkMemory, 0)
	mp.AssignTracer(log.VerifyingTracer())

	return mp, func() { os.RemoveAll(cfg.OriginFolder) }
}

func assureNoTrigger(t *testing.T, ch <-chan struct{}, deadlineMillis int) {
	clock := time.NewTimer(time.Duration(deadlineMillis) * time.Millisecond)
	select {
	case <-ch:
		t.Fatal("REDACTED")
	case <-clock.C:
	}
}

func assureTrigger(t *testing.T, ch <-chan struct{}, deadlineMillis int) {
	clock := time.NewTimer(time.Duration(deadlineMillis) * time.Millisecond)
	select {
	case <-ch:
	case <-clock.C:
		t.Fatal("REDACTED")
	}
}

func invokeInspectTransfer(t *testing.T, mp Txpool, txs kinds.Txs, nodeUID uint16) {
	transferDetails := TransferDetails{EmitterUID: nodeUID}
	for i, tx := range txs {
		if err := mp.InspectTransfer(tx, nil, transferDetails); err != nil {
			//
			//
			//
			if IsPreInspectFault(err) {
				continue
			}
			t.Fatalf("REDACTED", err, i)
		}
	}
}

//
func NewArbitraryTrans(countTrans int, transferSize int) kinds.Txs {
	txs := make(kinds.Txs, countTrans)
	for i := 0; i < countTrans; i++ {
		transferOctets := objectdepot.NewArbitraryTransfer(transferSize)
		txs[i] = transferOctets
	}
	return txs
}

//
//
func appendArbitraryTrans(t *testing.T, mp Txpool, tally int, nodeUID uint16) []kinds.Tx {
	t.Helper()
	txs := NewArbitraryTrans(tally, 20)
	invokeInspectTransfer(t, mp, txs, nodeUID)
	return txs
}

func appendTrans(tb testing.TB, mp Txpool, initial, num int) []kinds.Tx {
	tb.Helper()
	txs := make([]kinds.Tx, 0, num)
	for i := initial; i < num; i++ {
		tx := objectdepot.NewTransferFromUID(i)
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(tb, err)
		txs = append(txs, tx)
	}
	return txs
}

func VerifyHarvestMaximumOctetsMaximumFuel(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	//
	appendArbitraryTrans(t, mp, 1, UnclearNodeUID)
	tx0 := mp.TransHead().Item.(*txpoolTransfer)
	require.Equal(t, tx0.fuelDesired, int64(1), "REDACTED")
	//
	require.Equal(t, len(tx0.tx), 20, "REDACTED")
	mp.Purge()

	//
	//
	verifies := []struct {
		countTransToInstantiate int
		maximumOctets       int64
		maximumFuel         int64
		anticipatedCountTrans int
	}{
		{20, -1, -1, 20},
		{20, -1, 0, 0},
		{20, -1, 10, 10},
		{20, -1, 30, 20},
		{20, 0, -1, 0},
		{20, 0, 10, 0},
		{20, 10, 10, 0},
		{20, 24, 10, 1},
		{20, 240, 5, 5},
		{20, 240, -1, 10},
		{20, 240, 10, 10},
		{20, 240, 15, 10},
		{20, 20000, -1, 20},
		{20, 20000, 5, 5},
		{20, 20000, 30, 20},
	}
	for tcOrdinal, tt := range verifies {
		appendArbitraryTrans(t, mp, tt.countTransToInstantiate, UnclearNodeUID)
		got := mp.HarvestMaximumOctetsMaximumFuel(tt.maximumOctets, tt.maximumFuel)
		assert.Equal(t, tt.anticipatedCountTrans, len(got), "REDACTED",
			len(got), tt.anticipatedCountTrans, tcOrdinal)
		mp.Purge()
	}
}

func VerifyTxpoolScreens(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()
	emptyTransferArr := []kinds.Tx{[]byte{}}

	noopPreRefine := func(tx kinds.Tx) error { return nil }
	noopSubmitRefine := func(tx kinds.Tx, res *iface.ReplyInspectTransfer) error { return nil }

	//
	//
	verifies := []struct {
		countTransToInstantiate int
		preRefine      PreInspectFunction
		submitRefine     SubmitInspectFunction
		anticipatedCountTrans int
	}{
		{10, noopPreRefine, noopSubmitRefine, 10},
		{10, PreInspectMaximumOctets(10), noopSubmitRefine, 0},
		{10, PreInspectMaximumOctets(22), noopSubmitRefine, 10},
		{10, noopPreRefine, SubmitInspectMaximumFuel(-1), 10},
		{10, noopPreRefine, SubmitInspectMaximumFuel(0), 0},
		{10, noopPreRefine, SubmitInspectMaximumFuel(1), 10},
		{10, noopPreRefine, SubmitInspectMaximumFuel(3000), 10},
		{10, PreInspectMaximumOctets(10), SubmitInspectMaximumFuel(20), 0},
		{10, PreInspectMaximumOctets(30), SubmitInspectMaximumFuel(20), 10},
		{10, PreInspectMaximumOctets(22), SubmitInspectMaximumFuel(1), 10},
		{10, PreInspectMaximumOctets(22), SubmitInspectMaximumFuel(0), 0},
	}
	for tcOrdinal, tt := range verifies {
		err := mp.Modify(1, emptyTransferArr, ifaceReplies(len(emptyTransferArr), iface.CodeKindSuccess), tt.preRefine, tt.submitRefine)
		require.NoError(t, err)
		appendArbitraryTrans(t, mp, tt.countTransToInstantiate, UnclearNodeUID)
		require.Equal(t, tt.anticipatedCountTrans, mp.Volume(), "REDACTED", tcOrdinal)
		mp.Purge()
	}
}

func VerifyTxpoolModify(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	//
	{
		tx1 := objectdepot.NewTransferFromUID(1)
		err := mp.Modify(1, []kinds.Tx{tx1}, ifaceReplies(1, iface.CodeKindSuccess), nil, nil)
		require.NoError(t, err)
		err = mp.InspectTransfer(tx1, nil, TransferDetails{})
		if assert.Error(t, err) {
			assert.Equal(t, ErrTransferInRepository, err)
		}
	}

	//
	{
		tx2 := objectdepot.NewTransferFromUID(2)
		err := mp.InspectTransfer(tx2, nil, TransferDetails{})
		require.NoError(t, err)
		err = mp.Modify(1, []kinds.Tx{tx2}, ifaceReplies(1, iface.CodeKindSuccess), nil, nil)
		require.NoError(t, err)
		assert.Zero(t, mp.Volume())
	}

	//
	{
		tx3 := objectdepot.NewTransferFromUID(3)
		err := mp.InspectTransfer(tx3, nil, TransferDetails{})
		require.NoError(t, err)
		err = mp.Modify(1, []kinds.Tx{tx3}, ifaceReplies(1, 1), nil, nil)
		require.NoError(t, err)
		assert.Zero(t, mp.Volume())

		err = mp.InspectTransfer(tx3, nil, TransferDetails{})
		require.NoError(t, err)
	}
}

func VerifyTxpoolModifyDoesNegateAlarmWhenSoftwareSkippedTransfer(t *testing.T) {
	var callback ifacecustomer.Callback
	emulateCustomer := new(abcicliclients.Customer)
	emulateCustomer.On("REDACTED").Return(nil)
	emulateCustomer.On("REDACTED", mock.Anything)

	emulateCustomer.On("REDACTED").Return(nil).Times(4)
	emulateCustomer.On("REDACTED", mock.MatchedBy(func(cb ifacecustomer.Callback) bool { callback = cb; return true }))
	emulateCustomer.On("REDACTED", mock.Anything).Return(nil)

	mp, sanitize, err := newTxpoolWithApplicationEmulate(emulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	txs := []kinds.Tx{[]byte{0x01}, []byte{0x02}, []byte{0x03}, []byte{0x04}}
	for _, tx := range txs {
		requestOutput := newRequestOutput(tx, iface.CodeKindSuccess, iface.Transfercheckkind_New)
		emulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestOutput, nil)
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)

		//
		requestOutput.ExecuteCallback()
	}
	require.Len(t, txs, mp.Volume())
	require.True(t, mp.revalidate.done())

	//
	//
	err = mp.Modify(0, []kinds.Tx{txs[0]}, ifaceReplies(1, iface.CodeKindSuccess), nil, nil)
	require.Nil(t, err)

	//
	//
	//
	//
	//
	//
	reply := &iface.ReplyInspectTransfer{Code: iface.CodeKindSuccess}
	req := &iface.QueryInspectTransfer{Tx: txs[1]}
	callback(iface.ToQueryInspectTransfer(req), iface.ToReplyInspectTransfer(reply))

	req = &iface.QueryInspectTransfer{Tx: txs[3]}
	callback(iface.ToQueryInspectTransfer(req), iface.ToReplyInspectTransfer(reply))
	emulateCustomer.AssertExpectations(t)
}

func Verifyqueue_Keepinvalidtransincache(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	wsettings := settings.StandardSettings()
	wsettings.Txpool.RetainCorruptTransInRepository = true
	mp, sanitize := newTxpoolWithApplicationAndSettings(cc, wsettings)
	defer sanitize()

	//
	{
		a := make([]byte, 8)
		binary.BigEndian.PutUint64(a, 0)

		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, 1)

		err := mp.InspectTransfer(b, nil, TransferDetails{})
		require.NoError(t, err)

		//
		_, err = app.CompleteLedger(context.Background(), &iface.QueryCompleteLedger{
			Txs: [][]byte{a, b},
		})
		require.NoError(t, err)
		err = mp.Modify(1, []kinds.Tx{a, b},
			[]*iface.InvokeTransferOutcome{{Code: iface.CodeKindSuccess}, {Code: 2}}, nil, nil)
		require.NoError(t, err)

		//
		err = mp.InspectTransfer(a, nil, TransferDetails{})
		if assert.Error(t, err) {
			assert.Equal(t, ErrTransferInRepository, err)
		}

		//
		err = mp.InspectTransfer(b, nil, TransferDetails{})
		if assert.Error(t, err) {
			assert.Equal(t, ErrTransferInRepository, err)
		}
	}

	//
	{
		a := make([]byte, 8)
		binary.BigEndian.PutUint64(a, 0)

		//
		mp.repository.Delete(a)

		err := mp.InspectTransfer(a, nil, TransferDetails{})
		require.NoError(t, err)
	}
}

func VerifyTransAccessible(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()
	mp.ActivateTransAccessible()

	deadlineMillis := 500

	//
	assureNoTrigger(t, mp.TransAccessible(), deadlineMillis)

	//
	txs := appendArbitraryTrans(t, mp, 100, UnclearNodeUID)
	assureTrigger(t, mp.TransAccessible(), deadlineMillis)
	assureNoTrigger(t, mp.TransAccessible(), deadlineMillis)

	//
	//
	//
	confirmedTrans, outstandingTrans := txs[:50], txs[50:]
	if err := mp.Modify(1, confirmedTrans, ifaceReplies(len(confirmedTrans), iface.CodeKindSuccess), nil, nil); err != nil {
		t.Error(err)
	}
	assureTrigger(t, mp.TransAccessible(), deadlineMillis)
	assureNoTrigger(t, mp.TransAccessible(), deadlineMillis)

	//
	additionalTrans := appendArbitraryTrans(t, mp, 50, UnclearNodeUID)
	assureNoTrigger(t, mp.TransAccessible(), deadlineMillis)

	//
	outstandingTrans = append(outstandingTrans, additionalTrans...)
	confirmedTrans = outstandingTrans

	if err := mp.Modify(2, confirmedTrans, ifaceReplies(len(confirmedTrans), iface.CodeKindSuccess), nil, nil); err != nil {
		t.Error(err)
	}
	assureNoTrigger(t, mp.TransAccessible(), deadlineMillis)

	//
	appendArbitraryTrans(t, mp, 100, UnclearNodeUID)
	assureTrigger(t, mp.TransAccessible(), deadlineMillis)
	assureNoTrigger(t, mp.TransAccessible(), deadlineMillis)
}

func VerifySequentialHarvest(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)

	mp, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	applicationLinkConnect, _ := cc.NewIfaceCustomer()
	applicationLinkConnect.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err := applicationLinkConnect.Begin()
	require.Nil(t, err)

	repositoryIndex := make(map[string]struct{})
	transferTransScope := func(begin, end int) {
		//
		for i := begin; i < end; i++ {
			transferOctets := objectdepot.NewTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
			err := mp.InspectTransfer(transferOctets, nil, TransferDetails{})
			_, stored := repositoryIndex[string(transferOctets)]
			if stored {
				require.NotNil(t, err, "REDACTED")
			} else {
				require.Nil(t, err, "REDACTED")
			}
			repositoryIndex[string(transferOctets)] = struct{}{}

			//
			err = mp.InspectTransfer(transferOctets, nil, TransferDetails{})
			require.NotNil(t, err, "REDACTED")
		}
	}

	harvestInspect := func(exp int) {
		txs := mp.HarvestMaximumOctetsMaximumFuel(-1, -1)
		require.Equal(t, len(txs), exp, fmt.Sprintf("REDACTED", exp, len(txs)))
	}

	modifyScope := func(begin, end int) {
		txs := make(kinds.Txs, end-begin)
		for i := begin; i < end; i++ {
			txs[i-begin] = objectdepot.NewTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
		}
		if err := mp.Modify(0, txs, ifaceReplies(len(txs), iface.CodeKindSuccess), nil, nil); err != nil {
			t.Error(err)
		}
	}

	endorseScope := func(begin, end int) {
		//
		txs := make([][]byte, end-begin)
		for i := begin; i < end; i++ {
			txs[i-begin] = objectdepot.NewTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
		}

		res, err := applicationLinkConnect.CompleteLedger(context.Background(), &iface.QueryCompleteLedger{Txs: txs})
		if err != nil {
			t.Errorf("REDACTED", err)
		}
		for _, transOutcome := range res.TransOutcomes {
			if transOutcome.IsErr() {
				t.Errorf("REDACTED",
					transOutcome.Code, transOutcome.Data, transOutcome.Log)
			}
		}
		if len(res.ApplicationDigest) != 8 {
			t.Errorf("REDACTED", res.ApplicationDigest)
		}

		_, err = applicationLinkConnect.Endorse(context.Background(), &iface.QueryEndorse{})
		if err != nil {
			t.Errorf("REDACTED", err)
		}
	}

	//

	//
	transferTransScope(0, 100)

	//
	harvestInspect(100)

	//
	harvestInspect(100)

	//
	//
	transferTransScope(0, 1000)

	//
	harvestInspect(1000)

	//
	harvestInspect(1000)

	//
	endorseScope(0, 500)
	modifyScope(0, 500)

	//
	harvestInspect(500)

	//
	transferTransScope(900, 1100)

	//
	harvestInspect(600)
}

func Verifyqueue_Verifytxverifytxvolume(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)

	txpool, sanitize := newTxpoolWithApplication(cc)
	defer sanitize()

	maximumTransferVolume := txpool.settings.MaximumTransferOctets

	verifyScenarios := []struct {
		len int
		err bool
	}{
		//
		0: {10, false},
		1: {1000, false},
		2: {1000000, false},

		//
		3: {maximumTransferVolume - 1, false},
		4: {maximumTransferVolume, false},
		5: {maximumTransferVolume + 1, true},
	}

	for i, verifyInstance := range verifyScenarios {
		scenarioString := fmt.Sprintf("REDACTED", i, verifyInstance.len)

		tx := engineseed.Octets(verifyInstance.len)

		err := txpool.InspectTransfer(tx, nil, TransferDetails{})
		bv := gogotypes.BytesValue{Value: tx}
		bz, err2 := bv.Marshal()
		require.NoError(t, err2)
		require.Equal(t, len(bz), proto.Size(&bv), scenarioString)

		if !verifyInstance.err {
			require.NoError(t, err, scenarioString)
		} else {
			require.Equal(t, err, ErrTransferTooBulky{
				Max:    maximumTransferVolume,
				Factual: verifyInstance.len,
			}, scenarioString)
		}
	}
}

func VerifyTxpoolTransOctets(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)

	cfg := verify.RestoreVerifyOrigin("REDACTED")

	cfg.Txpool.MaximumTransOctets = 100
	mp, sanitize := newTxpoolWithApplicationAndSettings(cc, cfg)
	defer sanitize()

	//
	assert.EqualValues(t, 0, mp.VolumeOctets())

	//
	tx1 := objectdepot.NewArbitraryTransfer(10)
	err := mp.InspectTransfer(tx1, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 10, mp.VolumeOctets())

	//
	err = mp.Modify(1, []kinds.Tx{tx1}, ifaceReplies(1, iface.CodeKindSuccess), nil, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 0, mp.VolumeOctets())

	//
	tx2 := objectdepot.NewArbitraryTransfer(20)
	err = mp.InspectTransfer(tx2, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 20, mp.VolumeOctets())

	mp.Purge()
	assert.EqualValues(t, 0, mp.VolumeOctets())

	//
	tx3 := objectdepot.NewArbitraryTransfer(100)
	err = mp.InspectTransfer(tx3, nil, TransferDetails{})
	require.NoError(t, err)

	tx4 := objectdepot.NewArbitraryTransfer(10)
	err = mp.InspectTransfer(tx4, nil, TransferDetails{})
	if assert.Error(t, err) {
		assert.IsType(t, ErrTxpoolIsComplete{}, err)
	}

	//
	application2 := objectdepot.NewInRamSoftware()
	cc = gateway.NewNativeCustomerOriginator(application2)

	mp, sanitize = newTxpoolWithApplication(cc)
	defer sanitize()

	transferOctets := objectdepot.NewArbitraryTransfer(10)

	err = mp.InspectTransfer(transferOctets, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 10, mp.VolumeOctets())

	applicationLinkConnect, _ := cc.NewIfaceCustomer()
	applicationLinkConnect.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err = applicationLinkConnect.Begin()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := applicationLinkConnect.Halt(); err != nil {
			t.Error(err)
		}
	})

	res, err := applicationLinkConnect.CompleteLedger(context.Background(), &iface.QueryCompleteLedger{Txs: [][]byte{transferOctets}})
	require.NoError(t, err)
	require.EqualValues(t, 0, res.TransOutcomes[0].Code)
	require.NotEmpty(t, res.ApplicationDigest)

	_, err = applicationLinkConnect.Endorse(context.Background(), &iface.QueryEndorse{})
	require.NoError(t, err)

	//
	err = mp.Modify(1, []kinds.Tx{}, ifaceReplies(0, iface.CodeKindSuccess), nil, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 10, mp.VolumeOctets())

	//
	err = mp.InspectTransfer(tx1, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 20, mp.VolumeOctets())
	assert.Error(t, mp.DeleteTransferByKey(kinds.Tx([]byte{0x07}).Key()))
	assert.EqualValues(t, 20, mp.VolumeOctets())
	assert.NoError(t, mp.DeleteTransferByKey(kinds.Tx(tx1).Key()))
	assert.EqualValues(t, 10, mp.VolumeOctets())
}

func VerifyTxpoolNoRepositoryOverload(t *testing.T) {
	mp, sanitize := newTxpoolWithAsyncLinkage(t)
	defer sanitize()

	//
	tx0 := objectdepot.NewTransferFromUID(0)
	err := mp.InspectTransfer(tx0, nil, TransferDetails{})
	require.NoError(t, err)
	err = mp.PurgeApplicationLink()
	require.NoError(t, err)

	//
	for i := 1; i <= mp.settings.RepositoryVolume; i++ {
		err = mp.InspectTransfer(objectdepot.NewTransferFromUID(i), nil, TransferDetails{})
		require.NoError(t, err)
	}
	err = mp.PurgeApplicationLink()
	require.NoError(t, err)
	assert.False(t, mp.repository.Has(objectdepot.NewTransferFromUID(0)))

	//
	err = mp.InspectTransfer(tx0, nil, TransferDetails{})
	require.NoError(t, err)
	err = mp.PurgeApplicationLink()
	require.NoError(t, err)

	//
	located := 0
	for e := mp.txs.Head(); e != nil; e = e.Following() {
		if kinds.Tx.Key(e.Item.(*txpoolTransfer).tx) == kinds.Tx.Key(tx0) {
			located++
		}
	}
	assert.True(t, located == 1)
}

//
//
//
//
func VerifyTxpoolDistantApplicationParallelism(t *testing.T) {
	mp, sanitize := newTxpoolWithAsyncLinkage(t)
	defer sanitize()

	//
	nTrans := 10
	transferSize := 200
	txs := make([]kinds.Tx, nTrans)
	for i := 0; i < nTrans; i++ {
		txs[i] = objectdepot.NewArbitraryTransfer(transferSize)
	}

	//
	n := mp.settings.Volume
	maximumNodes := 5
	for i := 0; i < n; i++ {
		nodeUID := mrand.Intn(maximumNodes)
		transferCount := mrand.Intn(nTrans)
		tx := txs[transferCount]

		//
		mp.InspectTransfer(tx, nil, TransferDetails{EmitterUID: uint16(nodeUID)}) //
	}

	require.NoError(t, mp.PurgeApplicationLink())
}

func VerifyTxpoolParallelModifyAndAcceptInspectTransferReply(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)

	cfg := verify.RestoreVerifyOrigin("REDACTED")
	mp, sanitize := newTxpoolWithApplicationAndSettings(cc, cfg)
	defer sanitize()

	for h := 1; h <= 100; h++ {
		//
		//
		var wg sync.WaitGroup
		wg.Add(2)

		go func(h int) {
			defer wg.Done()

			doModify(t, mp, int64(h), []kinds.Tx{tx})
			require.Equal(t, int64(h), mp.level.Load(), "REDACTED")
		}(h)

		go func(h int) {
			defer wg.Done()

			tx := objectdepot.NewTransferFromUID(h)
			mp.outputCallbackfnInitialTime(tx, TransferDetails{}, iface.ToReplyInspectTransfer(&iface.ReplyInspectTransfer{Code: iface.CodeKindSuccess}))
			require.Equal(t, h, mp.Volume(), "REDACTED")
		}(h)

		wg.Wait()
	}
}

func VerifyTxpoolAlertTransAccessible(t *testing.T) {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)

	cfg := verify.RestoreVerifyOrigin("REDACTED")
	mp, sanitize := newTxpoolWithApplicationAndSettings(cc, cfg)
	defer sanitize()

	mp.ActivateTransAccessible()
	assert.NotNil(t, mp.transAccessible)
	require.False(t, mp.alertedTransAccessible.Load())

	//
	tx := objectdepot.NewTransferFromUID(1)
	mp.outputCallbackfnInitialTime(tx, TransferDetails{}, iface.ToReplyInspectTransfer(&iface.ReplyInspectTransfer{Code: iface.CodeKindSuccess}))
	require.Equal(t, 1, mp.Volume(), "REDACTED")
	require.True(t, mp.alertedTransAccessible.Load())
	require.Len(t, mp.TransAccessible(), 1)
	<-mp.TransAccessible()

	//
	mp.outputCallbackfnInitialTime(tx, TransferDetails{}, iface.ToReplyInspectTransfer(&iface.ReplyInspectTransfer{Code: iface.CodeKindSuccess}))
	require.Equal(t, 1, mp.Volume())
	require.True(t, mp.alertedTransAccessible.Load())
	require.Empty(t, mp.TransAccessible())

	//
	err := mp.Modify(1, []kinds.Tx{tx}, ifaceReplies(1, iface.CodeKindSuccess), nil, nil)
	require.NoError(t, err)
	require.Zero(t, mp.Volume())
	require.False(t, mp.alertedTransAccessible.Load())
}

//
func VerifyTxpoolAlignInspectTransferYieldFault(t *testing.T) {
	emulateCustomer := new(abcicliclients.Customer)
	emulateCustomer.On("REDACTED").Return(nil)
	emulateCustomer.On("REDACTED", mock.Anything)
	emulateCustomer.On("REDACTED", mock.Anything)

	mp, sanitize, err := newTxpoolWithApplicationEmulate(emulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	tx := []byte{0x01}
	emulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(nil, errors.New("REDACTED")).Once()

	//
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("REDACTED")
		}
	}()
	err = mp.InspectTransfer(tx, nil, TransferDetails{})
	require.NoError(t, err)
}

//
func VerifyTxpoolAlignRevalidateTransferYieldFault(t *testing.T) {
	emulateCustomer := new(abcicliclients.Customer)
	emulateCustomer.On("REDACTED").Return(nil)
	emulateCustomer.On("REDACTED", mock.Anything)
	emulateCustomer.On("REDACTED", mock.Anything)
	emulateCustomer.On("REDACTED").Return(nil)

	mp, sanitize, err := newTxpoolWithApplicationEmulate(emulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	txs := []kinds.Tx{[]byte{0x01}, []byte{0x02}}
	for _, tx := range txs {
		requestOutput := newRequestOutput(tx, iface.CodeKindSuccess, iface.Transfercheckkind_New)
		emulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestOutput, nil).Once()
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)

		//
		requestOutput.ExecuteCallback()
	}
	require.Len(t, txs, mp.Volume())

	//
	//
	requestOutput0 := newRequestOutput(txs[0], iface.CodeKindSuccess, iface.Transfercheckkind_Revalidate)
	emulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestOutput0, nil).Once()

	//
	emulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(nil, errors.New("REDACTED")).Once()

	//
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("REDACTED")
		}
	}()
	mp.revalidateTrans()
}

//
//
func VerifyTxpoolAsyncRevalidateTransferYieldFault(t *testing.T) {
	var callback ifacecustomer.Callback
	emulateCustomer := new(abcicliclients.Customer)
	emulateCustomer.On("REDACTED").Return(nil)
	emulateCustomer.On("REDACTED", mock.Anything)
	emulateCustomer.On("REDACTED").Return(nil).Times(4)
	emulateCustomer.On("REDACTED", mock.MatchedBy(func(cb ifacecustomer.Callback) bool { callback = cb; return true }))

	mp, sanitize, err := newTxpoolWithApplicationEmulate(emulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	txs := []kinds.Tx{[]byte{0x01}, []byte{0x02}, []byte{0x03}, []byte{0x04}}
	for _, tx := range txs {
		requestOutput := newRequestOutput(tx, iface.CodeKindSuccess, iface.Transfercheckkind_New)
		emulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestOutput, nil).Once()
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)

		//
		requestOutput.ExecuteCallback()
	}

	//
	require.Len(t, txs, mp.Volume())

	//
	require.True(t, mp.revalidate.done())
	require.Nil(t, mp.revalidate.pointer)
	require.Nil(t, mp.revalidate.end)
	require.False(t, mp.revalidate.isRevalidating.Load())
	emulateCustomer.AssertExpectations(t)

	//
	emulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(nil, nil).Times(4)

	//
	//
	//
	//
	emulateCustomer.On("REDACTED", mock.Anything).Run(func(_ mock.Arguments) {
		//
		requestOutput1 := newRequestOutput(txs[0], iface.CodeKindSuccess, iface.Transfercheckkind_Revalidate)
		callback(requestOutput1.Query, requestOutput1.Reply)
		//
		requestOutput2 := newRequestOutput(txs[2], 1, iface.Transfercheckkind_Revalidate)
		callback(requestOutput2.Query, requestOutput2.Reply)
	}).Return(nil)

	//
	mp.revalidateTrans()
	require.True(t, mp.revalidate.done())
	require.False(t, mp.revalidate.isRevalidating.Load())
	require.Nil(t, mp.revalidate.pointer)
	require.NotNil(t, mp.revalidate.end)
	require.Equal(t, mp.revalidate.end, mp.txs.Rear())
	require.Equal(t, len(txs)-1, mp.Volume()) //
	require.Equal(t, int32(2), mp.revalidate.countAwaitingTrans.Load())

	emulateCustomer.AssertExpectations(t)
}

//
func VerifyTxpoolRevalidateRivalry(t *testing.T) {
	mp, sanitize := newTxpoolWithAsyncLinkage(t)
	defer sanitize()

	//
	var err error
	txs := newDistinctTrans(10)
	for _, tx := range txs {
		err = mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)
	}

	//
	doModify(t, mp, 1, txs[:1])

	//
	require.True(t, mp.revalidate.done())
	require.Nil(t, mp.revalidate.pointer)

	//
	//
	err = mp.InspectTransfer(txs[:1][0], nil, TransferDetails{})
	require.Equal(t, err, ErrTransferInRepository)
	require.Zero(t, mp.revalidate.countAwaitingTrans.Load())
}

//
//
func VerifyTxpoolParallelInspectTransferAndModify(t *testing.T) {
	mp, sanitize := newTxpoolWithAsyncLinkage(t)
	defer sanitize()

	maximumLevel := 100
	var wg sync.WaitGroup
	wg.Add(1)

	//
	//
	go func() {
		defer wg.Done()

		time.Sleep(50 * time.Millisecond) //
		for h := 1; h <= maximumLevel; h++ {
			if mp.Volume() == 0 {
				break
			}
			txs := mp.HarvestMaximumOctetsMaximumFuel(100, -1)
			doModify(t, mp, int64(h), txs)
		}
	}()

	//
	for h := 1; h <= maximumLevel; h++ {
		err := mp.InspectTransfer(objectdepot.NewTransferFromUID(h), nil, TransferDetails{})
		require.NoError(t, err)
	}

	wg.Wait()

	//
	require.Zero(t, mp.Volume())
}

func newTxpoolWithAsyncLinkage(tb testing.TB) (*CCatalogTxpool, sanitizeFunction) {
	tb.Helper()
	socketRoute := fmt.Sprintf("REDACTED", engineseed.Str(6))
	app := objectdepot.NewInRamSoftware()
	_, host := newDistantApplication(tb, socketRoute, app)
	tb.Cleanup(func() {
		if err := host.Halt(); err != nil {
			tb.Error(err)
		}
	})
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	return newTxpoolWithApplicationAndSettings(gateway.NewDistantCustomerOriginator(socketRoute, "REDACTED", true), cfg)
}

//
func newDistantApplication(tb testing.TB, address string, app iface.Software) (ifacecustomer.Customer, daemon.Daemon) {
	tb.Helper()
	customerOriginator, err := ifacecustomer.NewCustomer(address, "REDACTED", true)
	require.NoError(tb, err)

	//
	host := ifaceservice.NewSocketHost(address, app)
	host.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	if err := host.Begin(); err != nil {
		tb.Fatalf("REDACTED", err.Error())
	}

	return customerOriginator, host
}

func newRequestOutput(tx kinds.Tx, code uint32, queryKind iface.InspectTransferKind) *ifacecustomer.RequestOutput { //
	requestOutput := ifacecustomer.NewRequestOutput(iface.ToQueryInspectTransfer(&iface.QueryInspectTransfer{Tx: tx, Kind: queryKind}))
	requestOutput.Reply = iface.ToReplyInspectTransfer(&iface.ReplyInspectTransfer{Code: code})
	return requestOutput
}

func ifaceReplies(n int, code uint32) []*iface.InvokeTransferOutcome {
	replies := make([]*iface.InvokeTransferOutcome, 0, n)
	for i := 0; i < n; i++ {
		replies = append(replies, &iface.InvokeTransferOutcome{Code: code})
	}
	return replies
}

func doModify(tb testing.TB, mp Txpool, level int64, txs []kinds.Tx) {
	tb.Helper()
	mp.Secure()
	err := mp.PurgeApplicationLink()
	require.NoError(tb, err)
	err = mp.Modify(level, txs, ifaceReplies(len(txs), iface.CodeKindSuccess), nil, nil)
	require.NoError(tb, err)
	mp.Release()
}
