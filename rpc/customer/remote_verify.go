package cust_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	remotelocal "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/regional"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/customer"
	rpcoverify "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var ctx = context.Background()

func obtainHttpsvcCustomer() *rpchttpsvc.Httpsvc {
	remoteLocation := rpcoverify.FetchSettings().RPC.OverhearLocation
	c, err := rpchttpsvc.New(remoteLocation, "REDACTED")
	if err != nil {
		panic(err)
	}
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func obtainHttpsvcCustomerUsingDeadline(deadline uint) *rpchttpsvc.Httpsvc {
	remoteLocation := rpcoverify.FetchSettings().RPC.OverhearLocation
	c, err := rpchttpsvc.FreshUsingDeadline(remoteLocation, "REDACTED", deadline)
	if err != nil {
		panic(err)
	}
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func obtainRegionalCustomer() *remotelocal.Regional {
	return remotelocal.New(peer)
}

//
func ObtainCustomers() []customer.Customer {
	return []customer.Customer{
		obtainHttpsvcCustomer(),
		obtainRegionalCustomer(),
	}
}

func VerifyVoidBespokeHttpsvcCustomer(t *testing.T) {
	require.Panics(t, func() {
		_, _ = rpchttpsvc.FreshUsingCustomer("REDACTED", "REDACTED", nil)
	})
	require.Panics(t, func() {
		_, _ = customeriface.FreshUsingHttpsvcCustomer("REDACTED", nil)
	})
}

func VerifyBespokeHttpsvcCustomer(t *testing.T) {
	distant := rpcoverify.FetchSettings().RPC.OverhearLocation
	c, err := rpchttpsvc.FreshUsingCustomer(distant, "REDACTED", http.DefaultClient)
	require.Nil(t, err)
	condition, err := c.Condition(context.Background())
	require.NoError(t, err)
	require.NotNil(t, condition)
}

func VerifyCrossoriginActivated(t *testing.T) {
	source := rpcoverify.FetchSettings().RPC.CrossoriginPermittedSources[0]
	distant := strings.ReplaceAll(rpcoverify.FetchSettings().RPC.OverhearLocation, "REDACTED", "REDACTED")

	req, err := http.NewRequest("REDACTED", distant, nil)
	require.Nil(t, err, "REDACTED", err)
	req.Header.Set("REDACTED", source)
	c := &http.Client{}
	reply, err := c.Do(req)
	require.Nil(t, err, "REDACTED", err)
	defer reply.Body.Close()

	assert.Equal(t, reply.Header.Get("REDACTED"), source)
}

//
func VerifyCondition(t *testing.T) {
	for i, c := range ObtainCustomers() {
		pseudonym := rpcoverify.FetchSettings().Pseudonym
		condition, err := c.Condition(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.Equal(t, pseudonym, condition.PeerDetails.Pseudonym)
	}
}

//
func VerifyDetails(t *testing.T) {
	for i, c := range ObtainCustomers() {
		//
		//
		details, err := c.IfaceDetails(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		//
		//
		assert.True(t, strings.Contains(details.Reply.Data, "REDACTED"))
	}
}

func VerifyNetworkDetails(t *testing.T) {
	for i, c := range ObtainCustomers() {
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		fabricinfo, err := nc.NetworkDetails(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.True(t, fabricinfo.Observing)
		assert.Equal(t, 0, len(fabricinfo.Nodes))
	}
}

func VerifyExportAgreementStatus(t *testing.T) {
	for i, c := range ObtainCustomers() {
		//
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		consensus, err := nc.ExportAgreementStatus(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.NotEmpty(t, consensus.IterationStatus)
		assert.Empty(t, consensus.Nodes)
	}
}

func VerifyAgreementStatus(t *testing.T) {
	for i, c := range ObtainCustomers() {
		//
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		consensus, err := nc.AgreementStatus(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.NotEmpty(t, consensus.IterationStatus)
	}
}

func VerifyVitality(t *testing.T) {
	for i, c := range ObtainCustomers() {
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		_, err := nc.Vitality(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
	}
}

func VerifyInaugurationAlsoAssessors(t *testing.T) {
	for i, c := range ObtainCustomers() {

		//
		gen, err := c.Inauguration(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		//
		require.Equal(t, 1, len(gen.Inauguration.Assessors))
		genesisvalue := gen.Inauguration.Assessors[0]

		//
		h := int64(1)
		values, err := c.Assessors(context.Background(), &h, nil, nil)
		require.Nil(t, err, "REDACTED", i, err)
		require.Equal(t, 1, len(values.Assessors))
		require.Equal(t, 1, values.Tally)
		require.Equal(t, 1, values.Sum)
		val := values.Assessors[0]

		//
		assert.Equal(t, genesisvalue.Potency, val.BallotingPotency)
		assert.Equal(t, genesisvalue.PublicToken, val.PublicToken)
	}
}

func VerifyInaugurationSegmented(t *testing.T) {
	ctx := t.Context()

	for _, c := range ObtainCustomers() {
		initial, err := c.InaugurationSegmented(ctx, 0)
		require.NoError(t, err)

		deserialized := make([]string, 0, initial.SumSegments)
		for i := 0; i < initial.SumSegments; i++ {
			segment, err := c.InaugurationSegmented(ctx, uint(i))
			require.NoError(t, err)
			data, err := base64.StdEncoding.DecodeString(segment.Data)
			require.NoError(t, err)
			deserialized = append(deserialized, string(data))

		}
		doc := []byte(strings.Join(deserialized, "REDACTED"))

		var out kinds.OriginPaper
		require.NoError(t, strongmindjson.Decode(doc, &out),
			"REDACTED", initial, string(doc))
	}
}

func VerifyIfaceInquire(t *testing.T) {
	for i, c := range ObtainCustomers() {
		//
		k, v, tx := CreateTransferTokval()
		bresp, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.Nil(t, err, "REDACTED", i, err)
		appiface := bresp.Altitude + 1 //

		//
		err = customer.PauseForeachAltitude(c, appiface, nil)
		require.NoError(t, err)
		res, err := c.IfaceInquire(context.Background(), "REDACTED", k)
		queryresp := res.Reply
		if assert.Nil(t, err) && assert.True(t, queryresp.EqualsOKAY()) {
			assert.EqualValues(t, v, queryresp.Datum)
		}
	}
}

//
func VerifyApplicationInvocations(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)
	for i, c := range ObtainCustomers() {

		//
		s, err := c.Condition(context.Background())
		require.NoError(err)
		//
		sh := s.ChronizeDetails.NewestLedgerAltitude

		//
		h := sh + 20
		_, err = c.Ledger(context.Background(), &h)
		require.Error(err) //

		//
		k, v, tx := CreateTransferTokval()
		bresp, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.NoError(err)
		require.True(bresp.TransferOutcome.EqualsOKAY())
		txh := bresp.Altitude
		appiface := txh + 1 //

		//
		err = customer.PauseForeachAltitude(c, appiface, nil)
		require.NoError(err)

		_queryresp, err := c.IfaceInquireUsingChoices(context.Background(), "REDACTED", k, customer.IfaceInquireChoices{Validate: false})
		require.NoError(err)
		queryresp := _queryresp.Reply
		if assert.True(queryresp.EqualsOKAY()) {
			assert.Equal(k, queryresp.Key)
			assert.EqualValues(v, queryresp.Datum)
		}

		//
		ptx, err := c.Tx(context.Background(), bresp.Digest, true)
		require.NoError(err)
		assert.EqualValues(txh, ptx.Altitude)
		assert.EqualValues(tx, ptx.Tx)

		//
		ledger, err := c.Ledger(context.Background(), &appiface)
		require.NoError(err)
		platformDigest := ledger.Ledger.PlatformDigest
		assert.True(len(platformDigest) > 0)
		assert.EqualValues(appiface, ledger.Ledger.Altitude)

		ledgerViaDigest, err := c.LedgerViaDigest(context.Background(), ledger.LedgerUUID.Digest)
		require.NoError(err)
		require.Equal(ledger, ledgerViaDigest)

		//
		heading, err := c.Heading(context.Background(), &appiface)
		require.NoError(err)
		require.Equal(ledger.Ledger.Heading, *heading.Heading)

		headlineViaDigest, err := c.HeadingViaDigest(context.Background(), ledger.LedgerUUID.Digest)
		require.NoError(err)
		require.Equal(heading, headlineViaDigest)

		//
		ledgerOutcomes, err := c.LedgerOutcomes(context.Background(), &txh)
		require.Nil(err, "REDACTED", i, err)
		assert.Equal(txh, ledgerOutcomes.Altitude)
		if assert.Equal(1, len(ledgerOutcomes.TransOutcomes)) {
			//
			assert.EqualValues(0, ledgerOutcomes.TransOutcomes[0].Cipher)
		}

		//
		details, err := c.LedgerchainDetails(context.Background(), appiface, appiface)
		require.NoError(err)
		assert.True(details.FinalAltitude >= appiface)
		if assert.Equal(1, len(details.LedgerMetadata)) {
			finalSummary := details.LedgerMetadata[0]
			assert.EqualValues(appiface, finalSummary.Heading.Altitude)
			ledgerData := ledger.Ledger
			assert.Equal(ledgerData.PlatformDigest, finalSummary.Heading.PlatformDigest)
			assert.Equal(ledger.LedgerUUID, finalSummary.LedgerUUID)
		}

		//
		endorse, err := c.Endorse(context.Background(), &appiface)
		require.NoError(err)
		contractappDigest := endorse.PlatformDigest
		assert.Equal(platformDigest, contractappDigest)
		assert.NotNil(endorse.Endorse)

		//
		h = appiface - 1
		persist2, err := c.Endorse(context.Background(), &h)
		require.NoError(err)
		assert.Equal(ledger.Ledger.FinalEndorseDigest, persist2.Endorse.Digest())

		//
		_presence, err := c.IfaceInquireUsingChoices(context.Background(), "REDACTED", k, customer.IfaceInquireChoices{Validate: true})
		require.NoError(err)
		presence := _presence.Reply
		assert.True(presence.EqualsOKAY())

		//
	}
}

func VerifyMulticastTransferChronize(t *testing.T) {
	demand := require.New(t)

	//
	txpool := peer.Txpool()
	initializeTxpoolExtent := txpool.Extent()

	for i, c := range ObtainCustomers() {
		_, _, tx := CreateTransferTokval()
		bresp, err := c.MulticastTransferChronize(context.Background(), tx)
		require.Nil(err, "REDACTED", i, err)
		require.Equal(bresp.Cipher, iface.CipherKindOKAY) //

		require.Equal(initializeTxpoolExtent+1, txpool.Extent())

		txs := txpool.HarvestMaximumTrans(len(tx))
		require.EqualValues(tx, txs[0])
		txpool.Purge()
	}
}

func VerifyMulticastTransferEndorse(t *testing.T) {
	demand := require.New(t)

	txpool := peer.Txpool()
	for i, c := range ObtainCustomers() {
		_, _, tx := CreateTransferTokval()
		bresp, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.Nil(err, "REDACTED", i, err)
		require.True(bresp.InspectTransfer.EqualsOKAY())
		require.True(bresp.TransferOutcome.EqualsOKAY())

		require.Equal(0, txpool.Extent())
	}
}

func VerifyPendingTrans(t *testing.T) {
	_, _, tx := CreateTransferTokval()

	ch := make(chan *iface.ReplyInspectTransfer, 1)
	txpool := peer.Txpool()
	err := txpool.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) { ch <- reply }, txpooll.TransferDetails{})
	require.NoError(t, err)

	//
	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.Error("REDACTED")
	}

	for _, c := range ObtainCustomers() {
		mc := c.(customer.TxpoolCustomer)
		threshold := 1
		res, err := mc.PendingTrans(context.Background(), &threshold)
		require.NoError(t, err)

		assert.Equal(t, 1, res.Tally)
		assert.Equal(t, 1, res.Sum)
		assert.Equal(t, txpool.ExtentOctets(), res.SumOctets)
		assert.Exactly(t, kinds.Txs{tx}, kinds.Txs(res.Txs))
	}

	txpool.Purge()
}

func VerifyCountPendingTrans(t *testing.T) {
	_, _, tx := CreateTransferTokval()

	ch := make(chan *iface.ReplyInspectTransfer, 1)
	txpool := peer.Txpool()
	err := txpool.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) { ch <- reply }, txpooll.TransferDetails{})
	require.NoError(t, err)

	//
	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.Error("REDACTED")
	}

	txpoolExtent := txpool.Extent()
	for i, c := range ObtainCustomers() {
		mc, ok := c.(customer.TxpoolCustomer)
		require.True(t, ok, "REDACTED", i)
		res, err := mc.CountPendingTrans(context.Background())
		require.Nil(t, err, "REDACTED", i, err)

		assert.Equal(t, txpoolExtent, res.Tally)
		assert.Equal(t, txpoolExtent, res.Sum)
		assert.Equal(t, txpool.ExtentOctets(), res.SumOctets)
	}

	txpool.Purge()
}

func VerifyInspectTransfer(t *testing.T) {
	txpool := peer.Txpool()

	for _, c := range ObtainCustomers() {
		_, _, tx := CreateTransferTokval()

		res, err := c.InspectTransfer(context.Background(), tx)
		require.NoError(t, err)
		assert.Equal(t, iface.CipherKindOKAY, res.Cipher)

		assert.Equal(t, 0, txpool.Extent(), "REDACTED")
	}
}

func VerifyTransfer(t *testing.T) {
	//
	c := obtainHttpsvcCustomer()
	_, _, tx := CreateTransferTokval()
	bresp, err := c.MulticastTransferEndorse(context.Background(), tx)
	require.Nil(t, err, "REDACTED", err)

	transferAltitude := bresp.Altitude
	transferDigest := bresp.Digest

	alternateTransferDigest := kinds.Tx("REDACTED").Digest()

	scenarios := []struct {
		sound bool
		ascertain bool
		digest  []byte
	}{
		//
		{true, false, transferDigest},
		{true, true, transferDigest},
		{false, false, alternateTransferDigest},
		{false, true, alternateTransferDigest},
		{false, false, nil},
		{false, true, nil},
	}

	for i, c := range ObtainCustomers() {
		for j, tc := range scenarios {
			t.Logf("REDACTED", i, j)

			//
			//
			ptx, err := c.Tx(context.Background(), tc.digest, tc.ascertain)

			if !tc.sound {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err, "REDACTED", err)
				assert.EqualValues(t, transferAltitude, ptx.Altitude)
				assert.EqualValues(t, tx, ptx.Tx)
				assert.Zero(t, ptx.Ordinal)
				assert.True(t, ptx.TransferOutcome.EqualsOKAY())
				assert.EqualValues(t, transferDigest, ptx.Digest)

				//
				attestation := ptx.Attestation
				if tc.ascertain && assert.EqualValues(t, tx, attestation.Data) {
					assert.NoError(t, attestation.Attestation.Validate(attestation.OriginDigest, transferDigest))
				}
			}
		}
	}
}

func VerifyTransferLookupUsingDeadline(t *testing.T) {
	//
	deadlineCustomer := obtainHttpsvcCustomerUsingDeadline(10)

	_, _, tx := CreateTransferTokval()
	_, err := deadlineCustomer.MulticastTransferEndorse(context.Background(), tx)
	require.NoError(t, err)

	//
	outcome, err := deadlineCustomer.TransferLookup(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
	require.Nil(t, err)
	require.Greater(t, len(outcome.Txs), 0, "REDACTED")
}

//
//
func VerifyLedgerLookup(t *testing.T) {
	c := obtainHttpsvcCustomer()

	//
	for i := 0; i < 10; i++ {
		_, _, tx := CreateTransferTokval()

		_, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.NoError(t, err)
	}
	require.NoError(t, customer.PauseForeachAltitude(c, 5, nil))
	outcome, err := c.LedgerLookup(context.Background(), "REDACTED", nil, nil, "REDACTED")
	require.NoError(t, err)
	ledgerNumber := len(outcome.Ledgers)
	//
	//
	//
	//

	//
	require.Equal(t, ledgerNumber, 0)
}

func VerifyTransferLookup(t *testing.T) {
	c := obtainHttpsvcCustomer()

	//
	for i := 0; i < 10; i++ {
		_, _, tx := CreateTransferTokval()
		_, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.NoError(t, err)
	}

	//
	//
	outcome, err := c.TransferLookup(context.Background(), "REDACTED", true, nil, nil, "REDACTED")
	require.NoError(t, err)
	transferNumber := len(outcome.Txs)

	//
	locate := outcome.Txs[len(outcome.Txs)-1]
	alternateTransferDigest := kinds.Tx("REDACTED").Digest()

	for _, c := range ObtainCustomers() {

		//
		outcome, err := c.TransferLookup(context.Background(), fmt.Sprintf("REDACTED", locate.Digest), true, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 1)
		require.Equal(t, locate.Digest, outcome.Txs[0].Digest)

		ptx := outcome.Txs[0]
		assert.EqualValues(t, locate.Altitude, ptx.Altitude)
		assert.EqualValues(t, locate.Tx, ptx.Tx)
		assert.Zero(t, ptx.Ordinal)
		assert.True(t, ptx.TransferOutcome.EqualsOKAY())
		assert.EqualValues(t, locate.Digest, ptx.Digest)

		//
		if assert.EqualValues(t, locate.Tx, ptx.Attestation.Data) {
			assert.NoError(t, ptx.Attestation.Attestation.Validate(ptx.Attestation.OriginDigest, locate.Digest))
		}

		//
		outcome, err = c.TransferLookup(context.Background(), fmt.Sprintf("REDACTED", locate.Altitude), true, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 1)

		//
		outcome, err = c.TransferLookup(context.Background(), fmt.Sprintf("REDACTED", alternateTransferDigest), false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 0)

		//
		outcome, err = c.TransferLookup(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Greater(t, len(outcome.Txs), 0, "REDACTED")

		//
		outcome, err = c.TransferLookup(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Greater(t, len(outcome.Txs), 0, "REDACTED")

		//
		outcome, err = c.TransferLookup(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Equal(t, len(outcome.Txs), 0, "REDACTED")

		//
		outcome, err = c.TransferLookup(context.Background(),
			"REDACTED", true, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Greater(t, len(outcome.Txs), 0, "REDACTED")

		//
		everyScreen := 1
		outcome, err = c.TransferLookup(context.Background(), "REDACTED", true, nil, &everyScreen, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 0)

		//
		outcome, err = c.TransferLookup(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		for k := 0; k < len(outcome.Txs)-1; k++ {
			require.LessOrEqual(t, outcome.Txs[k].Altitude, outcome.Txs[k+1].Altitude)
			require.LessOrEqual(t, outcome.Txs[k].Ordinal, outcome.Txs[k+1].Ordinal)
		}

		outcome, err = c.TransferLookup(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		for k := 0; k < len(outcome.Txs)-1; k++ {
			require.GreaterOrEqual(t, outcome.Txs[k].Altitude, outcome.Txs[k+1].Altitude)
			require.GreaterOrEqual(t, outcome.Txs[k].Ordinal, outcome.Txs[k+1].Ordinal)
		}
		//
		everyScreen = 3
		var (
			observed      = map[int64]bool{}
			maximumAltitude int64
			displays     = int(math.Ceil(float64(transferNumber) / float64(everyScreen)))
		)

		sumTransfer := 0
		for screen := 1; screen <= displays; screen++ {

			outcome, err := c.TransferLookup(context.Background(), "REDACTED", true, &screen, &everyScreen, "REDACTED")
			require.NoError(t, err)
			if screen < displays {
				require.Len(t, outcome.Txs, everyScreen)
			} else {
				require.LessOrEqual(t, len(outcome.Txs), everyScreen)
			}
			sumTransfer += len(outcome.Txs)
			for _, tx := range outcome.Txs {
				require.False(t, observed[tx.Altitude],
					"REDACTED", tx.Altitude, screen)
				require.Greater(t, tx.Altitude, maximumAltitude,
					"REDACTED", tx.Altitude, maximumAltitude, screen)
				observed[tx.Altitude] = true
				maximumAltitude = tx.Altitude
			}
		}
		require.Equal(t, transferNumber, sumTransfer)
		require.Len(t, observed, transferNumber)
	}
}

func VerifyGroupedJsonifaceInvocations(t *testing.T) {
	c := obtainHttpsvcCustomer()
	verifyGroupedJsonifaceInvocations(t, c)
}

func verifyGroupedJsonifaceInvocations(t *testing.T, c *rpchttpsvc.Httpsvc) {
	k1, v1, tx1 := CreateTransferTokval()
	k2, v2, tx2 := CreateTransferTokval()

	cluster := c.FreshCluster()
	r1, err := cluster.MulticastTransferEndorse(context.Background(), tx1)
	require.NoError(t, err)
	r2, err := cluster.MulticastTransferEndorse(context.Background(), tx2)
	require.NoError(t, err)
	require.Equal(t, 2, cluster.Tally())
	byteresults, err := cluster.Transmit(ctx)
	require.NoError(t, err)
	require.Len(t, byteresults, 2)
	require.Equal(t, 0, cluster.Tally())

	byteresult1, ok := byteresults[0].(*ktypes.OutcomeMulticastTransferEndorse)
	require.True(t, ok)
	require.Equal(t, *byteresult1, *r1)
	byteresult2, ok := byteresults[1].(*ktypes.OutcomeMulticastTransferEndorse)
	require.True(t, ok)
	require.Equal(t, *byteresult2, *r2)
	appiface := strongarithmetic.MaximumInt64n(byteresult1.Altitude, byteresult2.Altitude) + 1

	err = customer.PauseForeachAltitude(c, appiface, nil)
	require.NoError(t, err)

	q1, err := cluster.IfaceInquire(context.Background(), "REDACTED", k1)
	require.NoError(t, err)
	q2, err := cluster.IfaceInquire(context.Background(), "REDACTED", k2)
	require.NoError(t, err)
	require.Equal(t, 2, cluster.Tally())
	queryresults, err := cluster.Transmit(ctx)
	require.NoError(t, err)
	require.Len(t, queryresults, 2)
	require.Equal(t, 0, cluster.Tally())

	queryresult1, ok := queryresults[0].(*ktypes.OutcomeIfaceInquire)
	require.True(t, ok)
	require.Equal(t, *queryresult1, *q1)
	queryresult2, ok := queryresults[1].(*ktypes.OutcomeIfaceInquire)
	require.True(t, ok)
	require.Equal(t, *queryresult2, *q2)

	require.Equal(t, queryresult1.Reply.Key, k1)
	require.Equal(t, queryresult2.Reply.Key, k2)
	require.Equal(t, queryresult1.Reply.Datum, v1)
	require.Equal(t, queryresult2.Reply.Datum, v2)
}

func VerifyGroupedJsonifaceInvocationsRevocation(t *testing.T) {
	c := obtainHttpsvcCustomer()
	_, _, tx1 := CreateTransferTokval()
	_, _, tx2 := CreateTransferTokval()

	cluster := c.FreshCluster()
	_, err := cluster.MulticastTransferEndorse(context.Background(), tx1)
	require.NoError(t, err)
	_, err = cluster.MulticastTransferEndorse(context.Background(), tx2)
	require.NoError(t, err)
	//
	require.Equal(t, 2, cluster.Tally())
	//
	require.Equal(t, 2, cluster.Flush())
	//
	require.Equal(t, 0, cluster.Tally())
}

func VerifyRelayingBlankSolicitCluster(t *testing.T) {
	c := obtainHttpsvcCustomer()
	cluster := c.FreshCluster()
	_, err := cluster.Transmit(ctx)
	require.Error(t, err, "REDACTED")
}

func VerifyPurgingBlankSolicitCluster(t *testing.T) {
	c := obtainHttpsvcCustomer()
	cluster := c.FreshCluster()
	require.Zero(t, cluster.Flush(), "REDACTED")
}

func VerifyParallelJsonifaceGrouping(t *testing.T) {
	var wg sync.WaitGroup
	c := obtainHttpsvcCustomer()
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			verifyGroupedJsonifaceInvocations(t, c)
		}()
	}
	wg.Wait()
}
