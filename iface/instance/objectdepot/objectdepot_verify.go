package objectdepot

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"

	abciend "github.com/valkyrieworks/iface/customer"
	ifaceservice "github.com/valkyrieworks/iface/host"
	"github.com/valkyrieworks/iface/kinds"
)

const (
	verifyKey   = "REDACTED"
	verifyItem = "REDACTED"
)

func VerifyObjectDepotObject(t *testing.T) {
	ctx := t.Context()

	objectdepot := NewInRamSoftware()
	tx := []byte(verifyKey + "REDACTED" + verifyItem)
	verifyObjectDepot(ctx, t, objectdepot, tx, verifyKey, verifyItem)
	tx = []byte(verifyKey + "REDACTED" + verifyItem)
	verifyObjectDepot(ctx, t, objectdepot, tx, verifyKey, verifyItem)
}

func verifyObjectDepot(ctx context.Context, t *testing.T, app kinds.Software, tx []byte, key, item string) {
	inspectTransferReply, err := app.InspectTransfer(ctx, &kinds.QueryInspectTransfer{Tx: tx})
	require.NoError(t, err)
	require.Equal(t, uint32(0), inspectTransferReply.Code)

	ppReply, err := app.ArrangeNomination(ctx, &kinds.QueryArrangeNomination{Txs: [][]byte{tx}})
	require.NoError(t, err)
	require.Len(t, ppReply.Txs, 1)
	req := &kinds.QueryCompleteLedger{Level: 1, Txs: ppReply.Txs}
	ar, err := app.CompleteLedger(ctx, req)
	require.NoError(t, err)
	require.Equal(t, 1, len(ar.TransOutcomes))
	require.False(t, ar.TransOutcomes[0].IsErr())
	//
	_, err = app.Endorse(ctx, &kinds.QueryEndorse{})
	require.NoError(t, err)

	details, err := app.Details(ctx, &kinds.QueryDetails{})
	require.NoError(t, err)
	require.NotZero(t, details.FinalLedgerLevel)

	//
	outcomeInquire, err := app.Inquire(ctx, &kinds.QueryInquire{
		Route: "REDACTED",
		Data: []byte(key),
	})
	require.NoError(t, err)
	require.Equal(t, CodeKindSuccess, outcomeInquire.Code)
	require.Equal(t, key, string(outcomeInquire.Key))
	require.Equal(t, item, string(outcomeInquire.Item))
	require.EqualValues(t, details.FinalLedgerLevel, outcomeInquire.Level)

	//
	outcomeInquire, err = app.Inquire(ctx, &kinds.QueryInquire{
		Route:  "REDACTED",
		Data:  []byte(key),
		Demonstrate: true,
	})
	require.NoError(t, err)
	require.EqualValues(t, CodeKindSuccess, outcomeInquire.Code)
	require.Equal(t, key, string(outcomeInquire.Key))
	require.Equal(t, item, string(outcomeInquire.Item))
	require.EqualValues(t, details.FinalLedgerLevel, outcomeInquire.Level)
}

func VerifyDurableObjectDepotEmptyTransfer(t *testing.T) {
	ctx := t.Context()

	objectdepot := NewDurableSoftware(t.TempDir())
	tx := []byte("REDACTED")
	requestInspect := kinds.QueryInspectTransfer{Tx: tx}
	outputInspect, err := objectdepot.InspectTransfer(ctx, &requestInspect)
	require.NoError(t, err)
	require.Equal(t, outputInspect.Code, CodeKindCorruptTransferLayout)

	txs := make([][]byte, 0, 4)
	txs = append(txs, []byte("REDACTED"), []byte("REDACTED"), []byte("REDACTED"), []byte("REDACTED"))
	requestArrange := kinds.QueryArrangeNomination{Txs: txs, MaximumTransferOctets: 10 * 1024}
	outputArrange, err := objectdepot.ArrangeNomination(ctx, &requestArrange)
	require.NoError(t, err)
	require.Equal(t, len(requestArrange.Txs)-1, len(outputArrange.Txs), "REDACTED")
}

func VerifyDurableObjectDepotObject(t *testing.T) {
	ctx := t.Context()

	objectdepot := NewDurableSoftware(t.TempDir())
	key := verifyKey
	item := verifyItem
	verifyObjectDepot(ctx, t, objectdepot, NewTransfer(key, item), key, item)
}

func VerifyDurableObjectDepotDetails(t *testing.T) {
	ctx := t.Context()

	objectdepot := NewDurableSoftware(t.TempDir())
	require.NoError(t, InitObjectDepot(ctx, objectdepot))
	level := int64(0)

	outputDetails, err := objectdepot.Details(ctx, &kinds.QueryDetails{})
	require.NoError(t, err)
	if outputDetails.FinalLedgerLevel != level {
		t.Fatalf("REDACTED", level, outputDetails.FinalLedgerLevel)
	}

	//
	level = int64(1)
	digest := []byte("REDACTED")
	if _, err := objectdepot.CompleteLedger(ctx, &kinds.QueryCompleteLedger{Digest: digest, Level: level}); err != nil {
		t.Fatal(err)
	}

	_, err = objectdepot.Endorse(ctx, &kinds.QueryEndorse{})
	require.NoError(t, err)

	outputDetails, err = objectdepot.Details(ctx, &kinds.QueryDetails{})
	require.NoError(t, err)
	require.Equal(t, level, outputDetails.FinalLedgerLevel)
}

//
func VerifyValueRefreshes(t *testing.T) {
	ctx := t.Context()

	objectdepot := NewInRamSoftware()

	//
	sum := 10
	nInit := 5
	values := RandomValues(sum)
	//
	_, err := objectdepot.InitSeries(ctx, &kinds.QueryInitSeries{
		Ratifiers: values[:nInit],
	})
	require.NoError(t, err)

	values1, values2 := values[:nInit], objectdepot.fetchRatifiers()
	valuesEquivalent(t, values1, values2)

	var v1, v2, v3 kinds.RatifierModify

	//
	v1, v2 = values[nInit], values[nInit+1]
	vary := []kinds.RatifierModify{v1, v2}
	tx1 := CreateValueCollectionAlterTransfer(v1.PublicKey, v1.Energy)
	tx2 := CreateValueCollectionAlterTransfer(v2.PublicKey, v2.Energy)

	createExecuteLedger(ctx, t, objectdepot, 1, vary, tx1, tx2)

	values1, values2 = values[:nInit+2], objectdepot.fetchRatifiers()
	valuesEquivalent(t, values1, values2)

	//
	v1, v2, v3 = values[nInit-2], values[nInit-1], values[nInit]
	v1.Energy = 0
	v2.Energy = 0
	v3.Energy = 0
	vary = []kinds.RatifierModify{v1, v2, v3}
	tx1 = CreateValueCollectionAlterTransfer(v1.PublicKey, v1.Energy)
	tx2 = CreateValueCollectionAlterTransfer(v2.PublicKey, v2.Energy)
	tx3 := CreateValueCollectionAlterTransfer(v3.PublicKey, v3.Energy)

	createExecuteLedger(ctx, t, objectdepot, 2, vary, tx1, tx2, tx3)

	values1 = append(values[:nInit-2], values[nInit+1]) //
	values2 = objectdepot.fetchRatifiers()
	valuesEquivalent(t, values1, values2)

	//
	v1 = values[0]
	if v1.Energy == 5 {
		v1.Energy = 6
	} else {
		v1.Energy = 5
	}
	vary = []kinds.RatifierModify{v1}
	tx1 = CreateValueCollectionAlterTransfer(v1.PublicKey, v1.Energy)

	createExecuteLedger(ctx, t, objectdepot, 3, vary, tx1)

	values1 = append([]kinds.RatifierModify{v1}, values1[1:]...)
	values2 = objectdepot.fetchRatifiers()
	valuesEquivalent(t, values1, values2)
}

func VerifyInspectTransfer(t *testing.T) {
	ctx := t.Context()
	objectdepot := NewInRamSoftware()

	val := RandomValue()

	verifyScenarios := []struct {
		expirationCode uint32
		tx      []byte
	}{
		{CodeKindSuccess, NewTransfer("REDACTED", "REDACTED")},
		{CodeKindCorruptTransferLayout, []byte("REDACTED")},
		{CodeKindSuccess, []byte("REDACTED")},
		{CodeKindCorruptTransferLayout, []byte("REDACTED")},
		{CodeKindCorruptTransferLayout, []byte("REDACTED")},
		{CodeKindSuccess, []byte("REDACTED")},
		{CodeKindCorruptTransferLayout, []byte("REDACTED")},
		{CodeKindCorruptTransferLayout, []byte("REDACTED")},
		{CodeKindSuccess, CreateValueCollectionAlterTransfer(val.PublicKey, 10)},
	}

	for idx, tc := range verifyScenarios {
		reply, err := objectdepot.InspectTransfer(ctx, &kinds.QueryInspectTransfer{Tx: tc.tx})
		require.NoError(t, err, idx)
		fmt.Println(string(tc.tx))
		require.Equal(t, tc.expirationCode, reply.Code, idx)
	}
}

func VerifyCustomerHost(t *testing.T) {
	ctx := t.Context()
	//
	objectdepot := NewInRamSoftware()
	customer, _, err := createCustomerHost(t, objectdepot, "REDACTED", "REDACTED")
	require.NoError(t, err)
	runCustomerVerifies(ctx, t, customer)

	//
	objectdepot = NewInRamSoftware()
	gcustomer, _, err := createCustomerHost(t, objectdepot, t.TempDir(), "REDACTED")
	require.NoError(t, err)
	runCustomerVerifies(ctx, t, gcustomer)
}

func createExecuteLedger(
	ctx context.Context,
	t *testing.T,
	objectdepot kinds.Software,
	levelInteger int,
	vary []kinds.RatifierModify,
	txs ...[]byte,
) {
	//
	level := int64(levelInteger)
	digest := []byte("REDACTED")
	outputCompleteLedger, err := objectdepot.CompleteLedger(ctx, &kinds.QueryCompleteLedger{
		Digest:   digest,
		Level: level,
		Txs:    txs,
	})
	require.NoError(t, err)

	_, err = objectdepot.Endorse(ctx, &kinds.QueryEndorse{})
	require.NoError(t, err)

	valuesEquivalent(t, vary, outputCompleteLedger.RatifierRefreshes)
}

//
func valuesEquivalent(t *testing.T, values1, values2 []kinds.RatifierModify) {
	t.Helper()
	if len(values1) != len(values2) {
		t.Fatalf("REDACTED", len(values2), len(values1))
	}
	sort.Sort(kinds.RatifierRefreshes(values1))
	sort.Sort(kinds.RatifierRefreshes(values2))
	for i, v1 := range values1 {
		v2 := values2[i]
		if !v1.PublicKey.Equivalent(v2.PublicKey) ||
			v1.Energy != v2.Energy {
			t.Fatalf("REDACTED", i, v2.PublicKey, v2.Energy, v1.PublicKey, v1.Energy)
		}
	}
}

func createCustomerHost(t *testing.T, app kinds.Software, label, carrier string) (abciend.Customer, daemon.Daemon, error) {
	//
	address := fmt.Sprintf("REDACTED", label)
	tracer := log.VerifyingTracer()

	host, err := ifaceservice.NewHost(address, carrier, app)
	require.NoError(t, err)
	host.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := host.Begin(); err != nil {
		return nil, nil, err
	}

	t.Cleanup(func() {
		if err := host.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	customer, err := abciend.NewCustomer(address, carrier, false)
	require.NoError(t, err)
	customer.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := customer.Begin(); err != nil {
		return nil, nil, err
	}

	t.Cleanup(func() {
		if err := customer.Halt(); err != nil {
			t.Error(err)
		}
	})

	return customer, host, nil
}

func runCustomerVerifies(ctx context.Context, t *testing.T, customer abciend.Customer) {
	//
	tx := []byte(verifyKey + "REDACTED" + verifyItem)
	verifyObjectDepot(ctx, t, customer, tx, verifyKey, verifyItem)
	tx = []byte(verifyKey + "REDACTED" + verifyItem)
	verifyObjectDepot(ctx, t, customer, tx, verifyKey, verifyItem)
}

func VerifyTransferGenesis(t *testing.T) {
	require.Len(t, NewArbitraryTransfer(20), 20)
	require.Len(t, NewArbitraryTrans(10), 10)
}
