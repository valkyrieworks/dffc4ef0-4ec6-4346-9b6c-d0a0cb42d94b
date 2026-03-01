package statedepot

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	abcimaster "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

const (
	verifyToken   = "REDACTED"
	verifyDatum = "REDACTED"
)

func VerifyTokvalDepotTokval(t *testing.T) {
	ctx := t.Context()

	statedepot := FreshInsideRamPlatform()
	tx := []byte(verifyToken + "REDACTED" + verifyDatum)
	verifyTokvalDepot(ctx, t, statedepot, tx, verifyToken, verifyDatum)
	tx = []byte(verifyToken + "REDACTED" + verifyDatum)
	verifyTokvalDepot(ctx, t, statedepot, tx, verifyToken, verifyDatum)
}

func verifyTokvalDepot(ctx context.Context, t *testing.T, app kinds.Platform, tx []byte, key, datum string) {
	inspectTransferAnswer, err := app.InspectTransfer(ctx, &kinds.SolicitInspectTransfer{Tx: tx})
	require.NoError(t, err)
	require.Equal(t, uint32(0), inspectTransferAnswer.Cipher)

	ppAnswer, err := app.ArrangeNomination(ctx, &kinds.SolicitArrangeNomination{Txs: [][]byte{tx}})
	require.NoError(t, err)
	require.Len(t, ppAnswer.Txs, 1)
	req := &kinds.SolicitCulminateLedger{Altitude: 1, Txs: ppAnswer.Txs}
	ar, err := app.CulminateLedger(ctx, req)
	require.NoError(t, err)
	require.Equal(t, 1, len(ar.TransferOutcomes))
	require.False(t, ar.TransferOutcomes[0].EqualsFault())
	//
	_, err = app.Endorse(ctx, &kinds.SolicitEndorse{})
	require.NoError(t, err)

	details, err := app.Details(ctx, &kinds.SolicitDetails{})
	require.NoError(t, err)
	require.NotZero(t, details.FinalLedgerAltitude)

	//
	outcomeInquire, err := app.Inquire(ctx, &kinds.SolicitInquire{
		Route: "REDACTED",
		Data: []byte(key),
	})
	require.NoError(t, err)
	require.Equal(t, CipherKindOKAY, outcomeInquire.Cipher)
	require.Equal(t, key, string(outcomeInquire.Key))
	require.Equal(t, datum, string(outcomeInquire.Datum))
	require.EqualValues(t, details.FinalLedgerAltitude, outcomeInquire.Altitude)

	//
	outcomeInquire, err = app.Inquire(ctx, &kinds.SolicitInquire{
		Route:  "REDACTED",
		Data:  []byte(key),
		Validate: true,
	})
	require.NoError(t, err)
	require.EqualValues(t, CipherKindOKAY, outcomeInquire.Cipher)
	require.Equal(t, key, string(outcomeInquire.Key))
	require.Equal(t, datum, string(outcomeInquire.Datum))
	require.EqualValues(t, details.FinalLedgerAltitude, outcomeInquire.Altitude)
}

func VerifyEnduringTokvalDepotBlankTransfer(t *testing.T) {
	ctx := t.Context()

	statedepot := FreshEnduringPlatform(t.TempDir())
	tx := []byte("REDACTED")
	requestInspect := kinds.SolicitInspectTransfer{Tx: tx}
	resultInspect, err := statedepot.InspectTransfer(ctx, &requestInspect)
	require.NoError(t, err)
	require.Equal(t, resultInspect.Cipher, CipherKindUnfitTransferLayout)

	txs := make([][]byte, 0, 4)
	txs = append(txs, []byte("REDACTED"), []byte("REDACTED"), []byte("REDACTED"), []byte("REDACTED"))
	requestArrange := kinds.SolicitArrangeNomination{Txs: txs, MaximumTransferOctets: 10 * 1024}
	resultArrange, err := statedepot.ArrangeNomination(ctx, &requestArrange)
	require.NoError(t, err)
	require.Equal(t, len(requestArrange.Txs)-1, len(resultArrange.Txs), "REDACTED")
}

func VerifyEnduringTokvalDepotTokval(t *testing.T) {
	ctx := t.Context()

	statedepot := FreshEnduringPlatform(t.TempDir())
	key := verifyToken
	datum := verifyDatum
	verifyTokvalDepot(ctx, t, statedepot, FreshTransfer(key, datum), key, datum)
}

func VerifyEnduringTokvalDepotDetails(t *testing.T) {
	ctx := t.Context()

	statedepot := FreshEnduringPlatform(t.TempDir())
	require.NoError(t, InitializeTokvalDepot(ctx, statedepot))
	altitude := int64(0)

	resultDetails, err := statedepot.Details(ctx, &kinds.SolicitDetails{})
	require.NoError(t, err)
	if resultDetails.FinalLedgerAltitude != altitude {
		t.Fatalf("REDACTED", altitude, resultDetails.FinalLedgerAltitude)
	}

	//
	altitude = int64(1)
	digest := []byte("REDACTED")
	if _, err := statedepot.CulminateLedger(ctx, &kinds.SolicitCulminateLedger{Digest: digest, Altitude: altitude}); err != nil {
		t.Fatal(err)
	}

	_, err = statedepot.Endorse(ctx, &kinds.SolicitEndorse{})
	require.NoError(t, err)

	resultDetails, err = statedepot.Details(ctx, &kinds.SolicitDetails{})
	require.NoError(t, err)
	require.Equal(t, altitude, resultDetails.FinalLedgerAltitude)
}

//
func VerifyItemRevisions(t *testing.T) {
	ctx := t.Context()

	statedepot := FreshInsideRamPlatform()

	//
	sum := 10
	nthInitialize := 5
	values := ArbitraryValues(sum)
	//
	_, err := statedepot.InitializeSuccession(ctx, &kinds.SolicitInitializeSuccession{
		Assessors: values[:nthInitialize],
	})
	require.NoError(t, err)

	items1, items2 := values[:nthInitialize], statedepot.obtainAssessors()
	valuesEquivalent(t, items1, items2)

	var v1, v2, v3 kinds.AssessorRevise

	//
	v1, v2 = values[nthInitialize], values[nthInitialize+1]
	variance := []kinds.AssessorRevise{v1, v2}
	tx1 := CreateItemAssignModifyTransfer(v1.PublicToken, v1.Potency)
	tx2 := CreateItemAssignModifyTransfer(v2.PublicToken, v2.Potency)

	createExecuteLedger(ctx, t, statedepot, 1, variance, tx1, tx2)

	items1, items2 = values[:nthInitialize+2], statedepot.obtainAssessors()
	valuesEquivalent(t, items1, items2)

	//
	v1, v2, v3 = values[nthInitialize-2], values[nthInitialize-1], values[nthInitialize]
	v1.Potency = 0
	v2.Potency = 0
	v3.Potency = 0
	variance = []kinds.AssessorRevise{v1, v2, v3}
	tx1 = CreateItemAssignModifyTransfer(v1.PublicToken, v1.Potency)
	tx2 = CreateItemAssignModifyTransfer(v2.PublicToken, v2.Potency)
	tx3 := CreateItemAssignModifyTransfer(v3.PublicToken, v3.Potency)

	createExecuteLedger(ctx, t, statedepot, 2, variance, tx1, tx2, tx3)

	items1 = append(values[:nthInitialize-2], values[nthInitialize+1]) //
	items2 = statedepot.obtainAssessors()
	valuesEquivalent(t, items1, items2)

	//
	v1 = values[0]
	if v1.Potency == 5 {
		v1.Potency = 6
	} else {
		v1.Potency = 5
	}
	variance = []kinds.AssessorRevise{v1}
	tx1 = CreateItemAssignModifyTransfer(v1.PublicToken, v1.Potency)

	createExecuteLedger(ctx, t, statedepot, 3, variance, tx1)

	items1 = append([]kinds.AssessorRevise{v1}, items1[1:]...)
	items2 = statedepot.obtainAssessors()
	valuesEquivalent(t, items1, items2)
}

func VerifyInspectTransfer(t *testing.T) {
	ctx := t.Context()
	statedepot := FreshInsideRamPlatform()

	val := ArbitraryItem()

	verifyScenarios := []struct {
		expirationCipher uint32
		tx      []byte
	}{
		{CipherKindOKAY, FreshTransfer("REDACTED", "REDACTED")},
		{CipherKindUnfitTransferLayout, []byte("REDACTED")},
		{CipherKindOKAY, []byte("REDACTED")},
		{CipherKindUnfitTransferLayout, []byte("REDACTED")},
		{CipherKindUnfitTransferLayout, []byte("REDACTED")},
		{CipherKindOKAY, []byte("REDACTED")},
		{CipherKindUnfitTransferLayout, []byte("REDACTED")},
		{CipherKindUnfitTransferLayout, []byte("REDACTED")},
		{CipherKindOKAY, CreateItemAssignModifyTransfer(val.PublicToken, 10)},
	}

	for idx, tc := range verifyScenarios {
		reply, err := statedepot.InspectTransfer(ctx, &kinds.SolicitInspectTransfer{Tx: tc.tx})
		require.NoError(t, err, idx)
		fmt.Println(string(tc.tx))
		require.Equal(t, tc.expirationCipher, reply.Cipher, idx)
	}
}

func VerifyCustomerDaemon(t *testing.T) {
	ctx := t.Context()
	//
	statedepot := FreshInsideRamPlatform()
	customer, _, err := createCustomerDaemon(t, statedepot, "REDACTED", "REDACTED")
	require.NoError(t, err)
	executeCustomerVerifies(ctx, t, customer)

	//
	statedepot = FreshInsideRamPlatform()
	gnode, _, err := createCustomerDaemon(t, statedepot, t.TempDir(), "REDACTED")
	require.NoError(t, err)
	executeCustomerVerifies(ctx, t, gnode)
}

func createExecuteLedger(
	ctx context.Context,
	t *testing.T,
	statedepot kinds.Platform,
	altitudeInteger int,
	variance []kinds.AssessorRevise,
	txs ...[]byte,
) {
	//
	altitude := int64(altitudeInteger)
	digest := []byte("REDACTED")
	resultCulminateLedger, err := statedepot.CulminateLedger(ctx, &kinds.SolicitCulminateLedger{
		Digest:   digest,
		Altitude: altitude,
		Txs:    txs,
	})
	require.NoError(t, err)

	_, err = statedepot.Endorse(ctx, &kinds.SolicitEndorse{})
	require.NoError(t, err)

	valuesEquivalent(t, variance, resultCulminateLedger.AssessorRevisions)
}

//
func valuesEquivalent(t *testing.T, items1, items2 []kinds.AssessorRevise) {
	t.Helper()
	if len(items1) != len(items2) {
		t.Fatalf("REDACTED", len(items2), len(items1))
	}
	sort.Sort(kinds.AssessorRevisions(items1))
	sort.Sort(kinds.AssessorRevisions(items2))
	for i, v1 := range items1 {
		v2 := items2[i]
		if !v1.PublicToken.Equivalent(v2.PublicToken) ||
			v1.Potency != v2.Potency {
			t.Fatalf("REDACTED", i, v2.PublicToken, v2.Potency, v1.PublicToken, v1.Potency)
		}
	}
}

func createCustomerDaemon(t *testing.T, app kinds.Platform, alias, carrier string) (abcicustomer.Customer, facility.Facility, error) {
	//
	location := fmt.Sprintf("REDACTED", alias)
	tracer := log.VerifyingTracer()

	node, err := abcimaster.FreshDaemon(location, carrier, app)
	require.NoError(t, err)
	node.AssignTracer(tracer.Using("REDACTED", "REDACTED"))
	if err := node.Initiate(); err != nil {
		return nil, nil, err
	}

	t.Cleanup(func() {
		if err := node.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	customer, err := abcicustomer.FreshCustomer(location, carrier, false)
	require.NoError(t, err)
	customer.AssignTracer(tracer.Using("REDACTED", "REDACTED"))
	if err := customer.Initiate(); err != nil {
		return nil, nil, err
	}

	t.Cleanup(func() {
		if err := customer.Halt(); err != nil {
			t.Error(err)
		}
	})

	return customer, node, nil
}

func executeCustomerVerifies(ctx context.Context, t *testing.T, customer abcicustomer.Customer) {
	//
	tx := []byte(verifyToken + "REDACTED" + verifyDatum)
	verifyTokvalDepot(ctx, t, customer, tx, verifyToken, verifyDatum)
	tx = []byte(verifyToken + "REDACTED" + verifyDatum)
	verifyTokvalDepot(ctx, t, customer, tx, verifyToken, verifyDatum)
}

func VerifyTransferComposition(t *testing.T) {
	require.Len(t, FreshUnpredictableTransfer(20), 20)
	require.Len(t, FreshUnpredictableTrans(10), 10)
}
