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

	abcinode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	ifacemimics "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	abcimaster "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
type sanitizeMethod func()

func freshTxpoolUsingApplicationSimulate(customer abcinode.Customer) (*CNCatalogTxpool, sanitizeMethod, error) {
	setting := verify.RestoreVerifyOrigin("REDACTED")

	mp, cu := freshTxpoolUsingApplicationAlsoSettingsSimulate(setting, customer)
	return mp, cu, nil
}

func freshTxpoolUsingApplicationAlsoSettingsSimulate(
	cfg *settings.Settings,
	customer abcinode.Customer,
) (*CNCatalogTxpool, sanitizeMethod) {
	applicationLinkMemory := customer
	applicationLinkMemory.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err := applicationLinkMemory.Initiate()
	if err != nil {
		panic(err)
	}

	mp := FreshCNCatalogTxpool(cfg.Txpool, applicationLinkMemory, 0)
	mp.AssignTracer(log.VerifyingTracer())

	return mp, func() { os.RemoveAll(cfg.OriginPath) }
}

func freshTxpoolUsingApplication(cc delegate.CustomerOriginator) (*CNCatalogTxpool, sanitizeMethod) {
	setting := verify.RestoreVerifyOrigin("REDACTED")

	mp, cu := freshTxpoolUsingApplicationAlsoSettings(cc, setting)
	return mp, cu
}

func freshTxpoolUsingApplicationAlsoSettings(cc delegate.CustomerOriginator, cfg *settings.Settings) (*CNCatalogTxpool, sanitizeMethod) {
	applicationLinkMemory, _ := cc.FreshIfaceCustomer()
	applicationLinkMemory.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err := applicationLinkMemory.Initiate()
	if err != nil {
		panic(err)
	}

	mp := FreshCNCatalogTxpool(cfg.Txpool, applicationLinkMemory, 0)
	mp.AssignTracer(log.VerifyingTracer())

	return mp, func() { os.RemoveAll(cfg.OriginPath) }
}

func assureNegativeTrigger(t *testing.T, ch <-chan struct{}, deadlineMSEC int) {
	clock := time.NewTimer(time.Duration(deadlineMSEC) * time.Millisecond)
	select {
	case <-ch:
		t.Fatal("REDACTED")
	case <-clock.C:
	}
}

func assureTrigger(t *testing.T, ch <-chan struct{}, deadlineMSEC int) {
	clock := time.NewTimer(time.Duration(deadlineMSEC) * time.Millisecond)
	select {
	case <-ch:
	case <-clock.C:
		t.Fatal("REDACTED")
	}
}

func invocationInspectTransfer(t *testing.T, mp Txpool, txs kinds.Txs, nodeUUID uint16) {
	transferDetails := TransferDetails{OriginatorUUID: nodeUUID}
	for i, tx := range txs {
		if err := mp.InspectTransfer(tx, nil, transferDetails); err != nil {
			//
			//
			//
			if EqualsPriorInspectFailure(err) {
				continue
			}
			t.Fatalf("REDACTED", err, i)
		}
	}
}

//
func FreshUnpredictableTrans(countTrans int, transferLength int) kinds.Txs {
	txs := make(kinds.Txs, countTrans)
	for i := 0; i < countTrans; i++ {
		transferOctets := statedepot.FreshUnpredictableTransfer(transferLength)
		txs[i] = transferOctets
	}
	return txs
}

//
//
func appendUnpredictableTrans(t *testing.T, mp Txpool, tally int, nodeUUID uint16) []kinds.Tx {
	t.Helper()
	txs := FreshUnpredictableTrans(tally, 20)
	invocationInspectTransfer(t, mp, txs, nodeUUID)
	return txs
}

func appendTrans(tb testing.TB, mp Txpool, initial, num int) []kinds.Tx {
	tb.Helper()
	txs := make([]kinds.Tx, 0, num)
	for i := initial; i < num; i++ {
		tx := statedepot.FreshTransferOriginatingUUID(i)
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(tb, err)
		txs = append(txs, tx)
	}
	return txs
}

func VerifyHarvestMaximumOctetsMaximumFuel(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	//
	appendUnpredictableTrans(t, mp, 1, UnfamiliarNodeUUID)
	tx0 := mp.TransLeading().Datum.(*txpoolTransfer)
	require.Equal(t, tx0.fuelDesired, int64(1), "REDACTED")
	//
	require.Equal(t, len(tx0.tx), 20, "REDACTED")
	mp.Purge()

	//
	//
	verifies := []struct {
		countTransTowardGenerate int
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
	for testcasePosition, tt := range verifies {
		appendUnpredictableTrans(t, mp, tt.countTransTowardGenerate, UnfamiliarNodeUUID)
		got := mp.HarvestMaximumOctetsMaximumFuel(tt.maximumOctets, tt.maximumFuel)
		assert.Equal(t, tt.anticipatedCountTrans, len(got), "REDACTED",
			len(got), tt.anticipatedCountTrans, testcasePosition)
		mp.Purge()
	}
}

func VerifyTxpoolCriteria(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()
	blankTransferList := []kinds.Tx{[]byte{}}

	nooperationPriorRefine := func(tx kinds.Tx) error { return nil }
	nooperationSubmitRefine := func(tx kinds.Tx, res *iface.ReplyInspectTransfer) error { return nil }

	//
	//
	verifies := []struct {
		countTransTowardGenerate int
		priorRefine      PriorInspectMethod
		submitRefine     RelayInspectMethod
		anticipatedCountTrans int
	}{
		{10, nooperationPriorRefine, nooperationSubmitRefine, 10},
		{10, PriorInspectMaximumOctets(10), nooperationSubmitRefine, 0},
		{10, PriorInspectMaximumOctets(22), nooperationSubmitRefine, 10},
		{10, nooperationPriorRefine, SubmitInspectMaximumFuel(-1), 10},
		{10, nooperationPriorRefine, SubmitInspectMaximumFuel(0), 0},
		{10, nooperationPriorRefine, SubmitInspectMaximumFuel(1), 10},
		{10, nooperationPriorRefine, SubmitInspectMaximumFuel(3000), 10},
		{10, PriorInspectMaximumOctets(10), SubmitInspectMaximumFuel(20), 0},
		{10, PriorInspectMaximumOctets(30), SubmitInspectMaximumFuel(20), 10},
		{10, PriorInspectMaximumOctets(22), SubmitInspectMaximumFuel(1), 10},
		{10, PriorInspectMaximumOctets(22), SubmitInspectMaximumFuel(0), 0},
	}
	for testcasePosition, tt := range verifies {
		err := mp.Revise(1, blankTransferList, ifaceReplies(len(blankTransferList), iface.CipherKindOKAY), tt.priorRefine, tt.submitRefine)
		require.NoError(t, err)
		appendUnpredictableTrans(t, mp, tt.countTransTowardGenerate, UnfamiliarNodeUUID)
		require.Equal(t, tt.anticipatedCountTrans, mp.Extent(), "REDACTED", testcasePosition)
		mp.Purge()
	}
}

func VerifyTxpoolRevise(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	//
	{
		tx1 := statedepot.FreshTransferOriginatingUUID(1)
		err := mp.Revise(1, []kinds.Tx{tx1}, ifaceReplies(1, iface.CipherKindOKAY), nil, nil)
		require.NoError(t, err)
		err = mp.InspectTransfer(tx1, nil, TransferDetails{})
		if assert.Error(t, err) {
			assert.Equal(t, FaultTransferInsideStash, err)
		}
	}

	//
	{
		tx2 := statedepot.FreshTransferOriginatingUUID(2)
		err := mp.InspectTransfer(tx2, nil, TransferDetails{})
		require.NoError(t, err)
		err = mp.Revise(1, []kinds.Tx{tx2}, ifaceReplies(1, iface.CipherKindOKAY), nil, nil)
		require.NoError(t, err)
		assert.Zero(t, mp.Extent())
	}

	//
	{
		tx3 := statedepot.FreshTransferOriginatingUUID(3)
		err := mp.InspectTransfer(tx3, nil, TransferDetails{})
		require.NoError(t, err)
		err = mp.Revise(1, []kinds.Tx{tx3}, ifaceReplies(1, 1), nil, nil)
		require.NoError(t, err)
		assert.Zero(t, mp.Extent())

		err = mp.InspectTransfer(tx3, nil, TransferDetails{})
		require.NoError(t, err)
	}
}

func VerifyTxpoolReviseExecutesNegationAlarmWheneverPlatformOmittedTransfer(t *testing.T) {
	var clbk abcinode.Clbk
	simulateCustomer := new(ifacemimics.Customer)
	simulateCustomer.On("REDACTED").Return(nil)
	simulateCustomer.On("REDACTED", mock.Anything)

	simulateCustomer.On("REDACTED").Return(nil).Times(4)
	simulateCustomer.On("REDACTED", mock.MatchedBy(func(cb abcinode.Clbk) bool { clbk = cb; return true }))
	simulateCustomer.On("REDACTED", mock.Anything).Return(nil)

	mp, sanitize, err := freshTxpoolUsingApplicationSimulate(simulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	txs := []kinds.Tx{[]byte{0x01}, []byte{0x02}, []byte{0x03}, []byte{0x04}}
	for _, tx := range txs {
		requestResult := freshRequestResult(tx, iface.CipherKindOKAY, iface.Inspecttranskind_Fresh)
		simulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestResult, nil)
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)

		//
		requestResult.ExecuteClbk()
	}
	require.Len(t, txs, mp.Extent())
	require.True(t, mp.reinspect.complete())

	//
	//
	err = mp.Revise(0, []kinds.Tx{txs[0]}, ifaceReplies(1, iface.CipherKindOKAY), nil, nil)
	require.Nil(t, err)

	//
	//
	//
	//
	//
	//
	reply := &iface.ReplyInspectTransfer{Cipher: iface.CipherKindOKAY}
	req := &iface.SolicitInspectTransfer{Tx: txs[1]}
	clbk(iface.TowardSolicitInspectTransfer(req), iface.TowardReplyInspectTransfer(reply))

	req = &iface.SolicitInspectTransfer{Tx: txs[3]}
	clbk(iface.TowardSolicitInspectTransfer(req), iface.TowardReplyInspectTransfer(reply))
	simulateCustomer.AssertExpectations(t)
}

func Verifypool_Retaininvalidtransinstash(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	wiring := settings.FallbackSettings()
	wiring.Txpool.RetainUnfitTransInsideStash = true
	mp, sanitize := freshTxpoolUsingApplicationAlsoSettings(cc, wiring)
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
		_, err = app.CulminateLedger(context.Background(), &iface.SolicitCulminateLedger{
			Txs: [][]byte{a, b},
		})
		require.NoError(t, err)
		err = mp.Revise(1, []kinds.Tx{a, b},
			[]*iface.InvokeTransferOutcome{{Cipher: iface.CipherKindOKAY}, {Cipher: 2}}, nil, nil)
		require.NoError(t, err)

		//
		err = mp.InspectTransfer(a, nil, TransferDetails{})
		if assert.Error(t, err) {
			assert.Equal(t, FaultTransferInsideStash, err)
		}

		//
		err = mp.InspectTransfer(b, nil, TransferDetails{})
		if assert.Error(t, err) {
			assert.Equal(t, FaultTransferInsideStash, err)
		}
	}

	//
	{
		a := make([]byte, 8)
		binary.BigEndian.PutUint64(a, 0)

		//
		mp.stash.Discard(a)

		err := mp.InspectTransfer(a, nil, TransferDetails{})
		require.NoError(t, err)
	}
}

func VerifyTransAccessible(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()
	mp.ActivateTransAccessible()

	deadlineMSEC := 500

	//
	assureNegativeTrigger(t, mp.TransAccessible(), deadlineMSEC)

	//
	txs := appendUnpredictableTrans(t, mp, 100, UnfamiliarNodeUUID)
	assureTrigger(t, mp.TransAccessible(), deadlineMSEC)
	assureNegativeTrigger(t, mp.TransAccessible(), deadlineMSEC)

	//
	//
	//
	ratifiedTrans, pendingTrans := txs[:50], txs[50:]
	if err := mp.Revise(1, ratifiedTrans, ifaceReplies(len(ratifiedTrans), iface.CipherKindOKAY), nil, nil); err != nil {
		t.Error(err)
	}
	assureTrigger(t, mp.TransAccessible(), deadlineMSEC)
	assureNegativeTrigger(t, mp.TransAccessible(), deadlineMSEC)

	//
	extraTrans := appendUnpredictableTrans(t, mp, 50, UnfamiliarNodeUUID)
	assureNegativeTrigger(t, mp.TransAccessible(), deadlineMSEC)

	//
	pendingTrans = append(pendingTrans, extraTrans...)
	ratifiedTrans = pendingTrans

	if err := mp.Revise(2, ratifiedTrans, ifaceReplies(len(ratifiedTrans), iface.CipherKindOKAY), nil, nil); err != nil {
		t.Error(err)
	}
	assureNegativeTrigger(t, mp.TransAccessible(), deadlineMSEC)

	//
	appendUnpredictableTrans(t, mp, 100, UnfamiliarNodeUUID)
	assureTrigger(t, mp.TransAccessible(), deadlineMSEC)
	assureNegativeTrigger(t, mp.TransAccessible(), deadlineMSEC)
}

func VerifySequentialHarvest(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)

	mp, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	applicationLinkConnection, _ := cc.FreshIfaceCustomer()
	applicationLinkConnection.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err := applicationLinkConnection.Initiate()
	require.Nil(t, err)

	stashIndex := make(map[string]struct{})
	dispatchTransScope := func(initiate, end int) {
		//
		for i := initiate; i < end; i++ {
			transferOctets := statedepot.FreshTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
			err := mp.InspectTransfer(transferOctets, nil, TransferDetails{})
			_, buffered := stashIndex[string(transferOctets)]
			if buffered {
				require.NotNil(t, err, "REDACTED")
			} else {
				require.Nil(t, err, "REDACTED")
			}
			stashIndex[string(transferOctets)] = struct{}{}

			//
			err = mp.InspectTransfer(transferOctets, nil, TransferDetails{})
			require.NotNil(t, err, "REDACTED")
		}
	}

	harvestInspect := func(exp int) {
		txs := mp.HarvestMaximumOctetsMaximumFuel(-1, -1)
		require.Equal(t, len(txs), exp, fmt.Sprintf("REDACTED", exp, len(txs)))
	}

	reviseScope := func(initiate, end int) {
		txs := make(kinds.Txs, end-initiate)
		for i := initiate; i < end; i++ {
			txs[i-initiate] = statedepot.FreshTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
		}
		if err := mp.Revise(0, txs, ifaceReplies(len(txs), iface.CipherKindOKAY), nil, nil); err != nil {
			t.Error(err)
		}
	}

	endorseScope := func(initiate, end int) {
		//
		txs := make([][]byte, end-initiate)
		for i := initiate; i < end; i++ {
			txs[i-initiate] = statedepot.FreshTransfer(fmt.Sprintf("REDACTED", i), "REDACTED")
		}

		res, err := applicationLinkConnection.CulminateLedger(context.Background(), &iface.SolicitCulminateLedger{Txs: txs})
		if err != nil {
			t.Errorf("REDACTED", err)
		}
		for _, transferOutcome := range res.TransferOutcomes {
			if transferOutcome.EqualsFault() {
				t.Errorf("REDACTED",
					transferOutcome.Cipher, transferOutcome.Data, transferOutcome.Log)
			}
		}
		if len(res.PlatformDigest) != 8 {
			t.Errorf("REDACTED", res.PlatformDigest)
		}

		_, err = applicationLinkConnection.Endorse(context.Background(), &iface.SolicitEndorse{})
		if err != nil {
			t.Errorf("REDACTED", err)
		}
	}

	//

	//
	dispatchTransScope(0, 100)

	//
	harvestInspect(100)

	//
	harvestInspect(100)

	//
	//
	dispatchTransScope(0, 1000)

	//
	harvestInspect(1000)

	//
	harvestInspect(1000)

	//
	endorseScope(0, 500)
	reviseScope(0, 500)

	//
	harvestInspect(500)

	//
	dispatchTransScope(900, 1100)

	//
	harvestInspect(600)
}

func Verifypool_Verifytransferlimits(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)

	txpooll, sanitize := freshTxpoolUsingApplication(cc)
	defer sanitize()

	maximumTransferExtent := txpooll.settings.MaximumTransferOctets

	verifyScenarios := []struct {
		len int
		err bool
	}{
		//
		0: {10, false},
		1: {1000, false},
		2: {1000000, false},

		//
		3: {maximumTransferExtent - 1, false},
		4: {maximumTransferExtent, false},
		5: {maximumTransferExtent + 1, true},
	}

	for i, verifyInstance := range verifyScenarios {
		scenarioText := fmt.Sprintf("REDACTED", i, verifyInstance.len)

		tx := commitrand.Octets(verifyInstance.len)

		err := txpooll.InspectTransfer(tx, nil, TransferDetails{})
		bv := gogotypes.BytesValue{Value: tx}
		bz, fault2 := bv.Marshal()
		require.NoError(t, fault2)
		require.Equal(t, len(bz), proto.Size(&bv), scenarioText)

		if !verifyInstance.err {
			require.NoError(t, err, scenarioText)
		} else {
			require.Equal(t, err, FaultTransferExcessivelyAmple{
				Max:    maximumTransferExtent,
				Existing: verifyInstance.len,
			}, scenarioText)
		}
	}
}

func VerifyTxpoolTransOctets(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)

	cfg := verify.RestoreVerifyOrigin("REDACTED")

	cfg.Txpool.MaximumTransOctets = 100
	mp, sanitize := freshTxpoolUsingApplicationAlsoSettings(cc, cfg)
	defer sanitize()

	//
	assert.EqualValues(t, 0, mp.ExtentOctets())

	//
	tx1 := statedepot.FreshUnpredictableTransfer(10)
	err := mp.InspectTransfer(tx1, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 10, mp.ExtentOctets())

	//
	err = mp.Revise(1, []kinds.Tx{tx1}, ifaceReplies(1, iface.CipherKindOKAY), nil, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 0, mp.ExtentOctets())

	//
	tx2 := statedepot.FreshUnpredictableTransfer(20)
	err = mp.InspectTransfer(tx2, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 20, mp.ExtentOctets())

	mp.Purge()
	assert.EqualValues(t, 0, mp.ExtentOctets())

	//
	tx3 := statedepot.FreshUnpredictableTransfer(100)
	err = mp.InspectTransfer(tx3, nil, TransferDetails{})
	require.NoError(t, err)

	tx4 := statedepot.FreshUnpredictableTransfer(10)
	err = mp.InspectTransfer(tx4, nil, TransferDetails{})
	if assert.Error(t, err) {
		assert.IsType(t, FaultTxpoolEqualsComplete{}, err)
	}

	//
	platform2 := statedepot.FreshInsideRamPlatform()
	cc = delegate.FreshRegionalCustomerOriginator(platform2)

	mp, sanitize = freshTxpoolUsingApplication(cc)
	defer sanitize()

	transferOctets := statedepot.FreshUnpredictableTransfer(10)

	err = mp.InspectTransfer(transferOctets, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 10, mp.ExtentOctets())

	applicationLinkConnection, _ := cc.FreshIfaceCustomer()
	applicationLinkConnection.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	err = applicationLinkConnection.Initiate()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := applicationLinkConnection.Halt(); err != nil {
			t.Error(err)
		}
	})

	res, err := applicationLinkConnection.CulminateLedger(context.Background(), &iface.SolicitCulminateLedger{Txs: [][]byte{transferOctets}})
	require.NoError(t, err)
	require.EqualValues(t, 0, res.TransferOutcomes[0].Cipher)
	require.NotEmpty(t, res.PlatformDigest)

	_, err = applicationLinkConnection.Endorse(context.Background(), &iface.SolicitEndorse{})
	require.NoError(t, err)

	//
	err = mp.Revise(1, []kinds.Tx{}, ifaceReplies(0, iface.CipherKindOKAY), nil, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 10, mp.ExtentOctets())

	//
	err = mp.InspectTransfer(tx1, nil, TransferDetails{})
	require.NoError(t, err)
	assert.EqualValues(t, 20, mp.ExtentOctets())
	assert.Error(t, mp.DiscardTransferViaToken(kinds.Tx([]byte{0x07}).Key()))
	assert.EqualValues(t, 20, mp.ExtentOctets())
	assert.NoError(t, mp.DiscardTransferViaToken(kinds.Tx(tx1).Key()))
	assert.EqualValues(t, 10, mp.ExtentOctets())
}

func VerifyTxpoolNegativeStashOverrun(t *testing.T) {
	mp, sanitize := freshTxpoolUsingAsyncronousLinkage(t)
	defer sanitize()

	//
	tx0 := statedepot.FreshTransferOriginatingUUID(0)
	err := mp.InspectTransfer(tx0, nil, TransferDetails{})
	require.NoError(t, err)
	err = mp.PurgeApplicationLink()
	require.NoError(t, err)

	//
	for i := 1; i <= mp.settings.StashExtent; i++ {
		err = mp.InspectTransfer(statedepot.FreshTransferOriginatingUUID(i), nil, TransferDetails{})
		require.NoError(t, err)
	}
	err = mp.PurgeApplicationLink()
	require.NoError(t, err)
	assert.False(t, mp.stash.Has(statedepot.FreshTransferOriginatingUUID(0)))

	//
	err = mp.InspectTransfer(tx0, nil, TransferDetails{})
	require.NoError(t, err)
	err = mp.PurgeApplicationLink()
	require.NoError(t, err)

	//
	detected := 0
	for e := mp.txs.Leading(); e != nil; e = e.Following() {
		if kinds.Tx.Key(e.Datum.(*txpoolTransfer).tx) == kinds.Tx.Key(tx0) {
			detected++
		}
	}
	assert.True(t, detected == 1)
}

//
//
//
//
func VerifyTxpoolDistantApplicationParallelism(t *testing.T) {
	mp, sanitize := freshTxpoolUsingAsyncronousLinkage(t)
	defer sanitize()

	//
	nthTrans := 10
	transferLength := 200
	txs := make([]kinds.Tx, nthTrans)
	for i := 0; i < nthTrans; i++ {
		txs[i] = statedepot.FreshUnpredictableTransfer(transferLength)
	}

	//
	n := mp.settings.Extent
	maximumNodes := 5
	for i := 0; i < n; i++ {
		nodeUUID := mrand.Intn(maximumNodes)
		transferCount := mrand.Intn(nthTrans)
		tx := txs[transferCount]

		//
		mp.InspectTransfer(tx, nil, TransferDetails{OriginatorUUID: uint16(nodeUUID)}) //
	}

	require.NoError(t, mp.PurgeApplicationLink())
}

func VerifyTxpoolParallelReviseAlsoAcceptInspectTransferReply(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)

	cfg := verify.RestoreVerifyOrigin("REDACTED")
	mp, sanitize := freshTxpoolUsingApplicationAlsoSettings(cc, cfg)
	defer sanitize()

	for h := 1; h <= 100; h++ {
		//
		//
		var wg sync.WaitGroup
		wg.Add(2)

		go func(h int) {
			defer wg.Done()

			conductRevise(t, mp, int64(h), []kinds.Tx{tx})
			require.Equal(t, int64(h), mp.altitude.Load(), "REDACTED")
		}(h)

		go func(h int) {
			defer wg.Done()

			tx := statedepot.FreshTransferOriginatingUUID(h)
			mp.resultClbkInitialMoment(tx, TransferDetails{}, iface.TowardReplyInspectTransfer(&iface.ReplyInspectTransfer{Cipher: iface.CipherKindOKAY}))
			require.Equal(t, h, mp.Extent(), "REDACTED")
		}(h)

		wg.Wait()
	}
}

func VerifyTxpoolAlertTransAccessible(t *testing.T) {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)

	cfg := verify.RestoreVerifyOrigin("REDACTED")
	mp, sanitize := freshTxpoolUsingApplicationAlsoSettings(cc, cfg)
	defer sanitize()

	mp.ActivateTransAccessible()
	assert.NotNil(t, mp.transAccessible)
	require.False(t, mp.alertedTransAccessible.Load())

	//
	tx := statedepot.FreshTransferOriginatingUUID(1)
	mp.resultClbkInitialMoment(tx, TransferDetails{}, iface.TowardReplyInspectTransfer(&iface.ReplyInspectTransfer{Cipher: iface.CipherKindOKAY}))
	require.Equal(t, 1, mp.Extent(), "REDACTED")
	require.True(t, mp.alertedTransAccessible.Load())
	require.Len(t, mp.TransAccessible(), 1)
	<-mp.TransAccessible()

	//
	mp.resultClbkInitialMoment(tx, TransferDetails{}, iface.TowardReplyInspectTransfer(&iface.ReplyInspectTransfer{Cipher: iface.CipherKindOKAY}))
	require.Equal(t, 1, mp.Extent())
	require.True(t, mp.alertedTransAccessible.Load())
	require.Empty(t, mp.TransAccessible())

	//
	err := mp.Revise(1, []kinds.Tx{tx}, ifaceReplies(1, iface.CipherKindOKAY), nil, nil)
	require.NoError(t, err)
	require.Zero(t, mp.Extent())
	require.False(t, mp.alertedTransAccessible.Load())
}

//
func VerifyTxpoolChronizeInspectTransferYieldFailure(t *testing.T) {
	simulateCustomer := new(ifacemimics.Customer)
	simulateCustomer.On("REDACTED").Return(nil)
	simulateCustomer.On("REDACTED", mock.Anything)
	simulateCustomer.On("REDACTED", mock.Anything)

	mp, sanitize, err := freshTxpoolUsingApplicationSimulate(simulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	tx := []byte{0x01}
	simulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(nil, errors.New("REDACTED")).Once()

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
func VerifyTxpoolChronizeReinspectTransferYieldFailure(t *testing.T) {
	simulateCustomer := new(ifacemimics.Customer)
	simulateCustomer.On("REDACTED").Return(nil)
	simulateCustomer.On("REDACTED", mock.Anything)
	simulateCustomer.On("REDACTED", mock.Anything)
	simulateCustomer.On("REDACTED").Return(nil)

	mp, sanitize, err := freshTxpoolUsingApplicationSimulate(simulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	txs := []kinds.Tx{[]byte{0x01}, []byte{0x02}}
	for _, tx := range txs {
		requestResult := freshRequestResult(tx, iface.CipherKindOKAY, iface.Inspecttranskind_Fresh)
		simulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestResult, nil).Once()
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)

		//
		requestResult.ExecuteClbk()
	}
	require.Len(t, txs, mp.Extent())

	//
	//
	requestOutcome0 := freshRequestResult(txs[0], iface.CipherKindOKAY, iface.Inspecttranskind_Reinspect)
	simulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestOutcome0, nil).Once()

	//
	simulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(nil, errors.New("REDACTED")).Once()

	//
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("REDACTED")
		}
	}()
	mp.reinspectTrans()
}

//
//
func VerifyTxpoolAsyncronousReinspectTransferYieldFailure(t *testing.T) {
	var clbk abcinode.Clbk
	simulateCustomer := new(ifacemimics.Customer)
	simulateCustomer.On("REDACTED").Return(nil)
	simulateCustomer.On("REDACTED", mock.Anything)
	simulateCustomer.On("REDACTED").Return(nil).Times(4)
	simulateCustomer.On("REDACTED", mock.MatchedBy(func(cb abcinode.Clbk) bool { clbk = cb; return true }))

	mp, sanitize, err := freshTxpoolUsingApplicationSimulate(simulateCustomer)
	require.NoError(t, err)
	defer sanitize()

	//
	txs := []kinds.Tx{[]byte{0x01}, []byte{0x02}, []byte{0x03}, []byte{0x04}}
	for _, tx := range txs {
		requestResult := freshRequestResult(tx, iface.CipherKindOKAY, iface.Inspecttranskind_Fresh)
		simulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(requestResult, nil).Once()
		err := mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)

		//
		requestResult.ExecuteClbk()
	}

	//
	require.Len(t, txs, mp.Extent())

	//
	require.True(t, mp.reinspect.complete())
	require.Nil(t, mp.reinspect.locator)
	require.Nil(t, mp.reinspect.end)
	require.False(t, mp.reinspect.equalsReexamining.Load())
	simulateCustomer.AssertExpectations(t)

	//
	simulateCustomer.On("REDACTED", mock.Anything, mock.Anything).Return(nil, nil).Times(4)

	//
	//
	//
	//
	simulateCustomer.On("REDACTED", mock.Anything).Run(func(_ mock.Arguments) {
		//
		requestOutcome1 := freshRequestResult(txs[0], iface.CipherKindOKAY, iface.Inspecttranskind_Reinspect)
		clbk(requestOutcome1.Solicit, requestOutcome1.Reply)
		//
		requestOutcome2 := freshRequestResult(txs[2], 1, iface.Inspecttranskind_Reinspect)
		clbk(requestOutcome2.Solicit, requestOutcome2.Reply)
	}).Return(nil)

	//
	mp.reinspectTrans()
	require.True(t, mp.reinspect.complete())
	require.False(t, mp.reinspect.equalsReexamining.Load())
	require.Nil(t, mp.reinspect.locator)
	require.NotNil(t, mp.reinspect.end)
	require.Equal(t, mp.reinspect.end, mp.txs.Rear())
	require.Equal(t, len(txs)-1, mp.Extent()) //
	require.Equal(t, int32(2), mp.reinspect.countAwaitingTrans.Load())

	simulateCustomer.AssertExpectations(t)
}

//
func VerifyTxpoolReinspectRivalry(t *testing.T) {
	mp, sanitize := freshTxpoolUsingAsyncronousLinkage(t)
	defer sanitize()

	//
	var err error
	txs := freshDistinctTrans(10)
	for _, tx := range txs {
		err = mp.InspectTransfer(tx, nil, TransferDetails{})
		require.NoError(t, err)
	}

	//
	conductRevise(t, mp, 1, txs[:1])

	//
	require.True(t, mp.reinspect.complete())
	require.Nil(t, mp.reinspect.locator)

	//
	//
	err = mp.InspectTransfer(txs[:1][0], nil, TransferDetails{})
	require.Equal(t, err, FaultTransferInsideStash)
	require.Zero(t, mp.reinspect.countAwaitingTrans.Load())
}

//
//
func VerifyTxpoolParallelInspectTransferAlsoRevise(t *testing.T) {
	mp, sanitize := freshTxpoolUsingAsyncronousLinkage(t)
	defer sanitize()

	maximumAltitude := 100
	var wg sync.WaitGroup
	wg.Add(1)

	//
	//
	go func() {
		defer wg.Done()

		time.Sleep(50 * time.Millisecond) //
		for h := 1; h <= maximumAltitude; h++ {
			if mp.Extent() == 0 {
				break
			}
			txs := mp.HarvestMaximumOctetsMaximumFuel(100, -1)
			conductRevise(t, mp, int64(h), txs)
		}
	}()

	//
	for h := 1; h <= maximumAltitude; h++ {
		err := mp.InspectTransfer(statedepot.FreshTransferOriginatingUUID(h), nil, TransferDetails{})
		require.NoError(t, err)
	}

	wg.Wait()

	//
	require.Zero(t, mp.Extent())
}

func freshTxpoolUsingAsyncronousLinkage(tb testing.TB) (*CNCatalogTxpool, sanitizeMethod) {
	tb.Helper()
	terminalRoute := fmt.Sprintf("REDACTED", commitrand.Str(6))
	app := statedepot.FreshInsideRamPlatform()
	_, node := freshDistantApplication(tb, terminalRoute, app)
	tb.Cleanup(func() {
		if err := node.Halt(); err != nil {
			tb.Error(err)
		}
	})
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	return freshTxpoolUsingApplicationAlsoSettings(delegate.FreshDistantCustomerOriginator(terminalRoute, "REDACTED", true), cfg)
}

//
func freshDistantApplication(tb testing.TB, location string, app iface.Platform) (abcinode.Customer, facility.Facility) {
	tb.Helper()
	customerOriginator, err := abcinode.FreshCustomer(location, "REDACTED", true)
	require.NoError(tb, err)

	//
	node := abcimaster.FreshPortDaemon(location, app)
	node.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	if err := node.Initiate(); err != nil {
		tb.Fatalf("REDACTED", err.Error())
	}

	return customerOriginator, node
}

func freshRequestResult(tx kinds.Tx, cipher uint32, solicitKind iface.InspectTransferKind) *abcinode.RequestResult { //
	requestResult := abcinode.FreshRequestResult(iface.TowardSolicitInspectTransfer(&iface.SolicitInspectTransfer{Tx: tx, Kind: solicitKind}))
	requestResult.Reply = iface.TowardReplyInspectTransfer(&iface.ReplyInspectTransfer{Cipher: cipher})
	return requestResult
}

func ifaceReplies(n int, cipher uint32) []*iface.InvokeTransferOutcome {
	replies := make([]*iface.InvokeTransferOutcome, 0, n)
	for i := 0; i < n; i++ {
		replies = append(replies, &iface.InvokeTransferOutcome{Cipher: cipher})
	}
	return replies
}

func conductRevise(tb testing.TB, mp Txpool, altitude int64, txs []kinds.Tx) {
	tb.Helper()
	mp.Secure()
	err := mp.PurgeApplicationLink()
	require.NoError(tb, err)
	err = mp.Revise(altitude, txs, ifaceReplies(len(txs), iface.CipherKindOKAY), nil, nil)
	require.NoError(tb, err)
	mp.Release()
}
