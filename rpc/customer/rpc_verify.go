package agent_test

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

	iface "github.com/valkyrieworks/iface/kinds"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	cometmath "github.com/valkyrieworks/utils/math"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/rpc/customer"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	rpcnative "github.com/valkyrieworks/rpc/customer/native"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpccustomer "github.com/valkyrieworks/rpc/jsonrpc/customer"
	rpctest "github.com/valkyrieworks/rpc/verify"
	"github.com/valkyrieworks/kinds"
)

var ctx = context.Background()

func fetchHTTPCustomer() *rpchttp.HTTP {
	rpcAddress := rpctest.FetchSettings().RPC.AcceptLocation
	c, err := rpchttp.New(rpcAddress, "REDACTED")
	if err != nil {
		panic(err)
	}
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func fetchHTTPCustomerWithDeadline(deadline uint) *rpchttp.HTTP {
	rpcAddress := rpctest.FetchSettings().RPC.AcceptLocation
	c, err := rpchttp.NewWithDeadline(rpcAddress, "REDACTED", deadline)
	if err != nil {
		panic(err)
	}
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func fetchNativeCustomer() *rpcnative.Native {
	return rpcnative.New(member)
}

//
func FetchAgents() []customer.Customer {
	return []customer.Customer{
		fetchHTTPCustomer(),
		fetchNativeCustomer(),
	}
}

func VerifyNullBespokeHTTPCustomer(t *testing.T) {
	require.Panics(t, func() {
		_, _ = rpchttp.NewWithCustomer("REDACTED", "REDACTED", nil)
	})
	require.Panics(t, func() {
		_, _ = rpccustomer.NewWithHTTPCustomer("REDACTED", nil)
	})
}

func VerifyBespokeHTTPCustomer(t *testing.T) {
	external := rpctest.FetchSettings().RPC.AcceptLocation
	c, err := rpchttp.NewWithCustomer(external, "REDACTED", http.DefaultClient)
	require.Nil(t, err)
	state, err := c.Status(context.Background())
	require.NoError(t, err)
	require.NotNil(t, state)
}

func VerifyCorsActivated(t *testing.T) {
	source := rpctest.FetchSettings().RPC.CORSPermittedSources[0]
	external := strings.ReplaceAll(rpctest.FetchSettings().RPC.AcceptLocation, "REDACTED", "REDACTED")

	req, err := http.NewRequest("REDACTED", external, nil)
	require.Nil(t, err, "REDACTED", err)
	req.Header.Set("REDACTED", source)
	c := &http.Client{}
	reply, err := c.Do(req)
	require.Nil(t, err, "REDACTED", err)
	defer reply.Body.Close()

	assert.Equal(t, reply.Header.Get("REDACTED"), source)
}

//
func VerifyState(t *testing.T) {
	for i, c := range FetchAgents() {
		moniker := rpctest.FetchSettings().Moniker
		state, err := c.Status(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.Equal(t, moniker, state.MemberDetails.Moniker)
	}
}

//
func VerifyDetails(t *testing.T) {
	for i, c := range FetchAgents() {
		//
		//
		details, err := c.IfaceDetails(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		//
		//
		assert.True(t, strings.Contains(details.Reply.Data, "REDACTED"))
	}
}

func VerifyNetDetails(t *testing.T) {
	for i, c := range FetchAgents() {
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		netdata, err := nc.NetDetails(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.True(t, netdata.Observing)
		assert.Equal(t, 0, len(netdata.Nodes))
	}
}

func VerifyExportAgreementStatus(t *testing.T) {
	for i, c := range FetchAgents() {
		//
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		const, err := nc.ExportAgreementStatus(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.NotEmpty(t, const.EpochStatus)
		assert.Empty(t, const.Nodes)
	}
}

func VerifyAgreementStatus(t *testing.T) {
	for i, c := range FetchAgents() {
		//
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		const, err := nc.AgreementStatus(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		assert.NotEmpty(t, const.EpochStatus)
	}
}

func VerifyVitality(t *testing.T) {
	for i, c := range FetchAgents() {
		nc, ok := c.(customer.FabricCustomer)
		require.True(t, ok, "REDACTED", i)
		_, err := nc.Vitality(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
	}
}

func VerifyOriginAndRatifiers(t *testing.T) {
	for i, c := range FetchAgents() {

		//
		gen, err := c.Origin(context.Background())
		require.Nil(t, err, "REDACTED", i, err)
		//
		require.Equal(t, 1, len(gen.Origin.Ratifiers))
		gvalue := gen.Origin.Ratifiers[0]

		//
		h := int64(1)
		values, err := c.Ratifiers(context.Background(), &h, nil, nil)
		require.Nil(t, err, "REDACTED", i, err)
		require.Equal(t, 1, len(values.Ratifiers))
		require.Equal(t, 1, values.Number)
		require.Equal(t, 1, values.Sum)
		val := values.Ratifiers[0]

		//
		assert.Equal(t, gvalue.Energy, val.PollingEnergy)
		assert.Equal(t, gvalue.PublicKey, val.PublicKey)
	}
}

func VerifyOriginSegmented(t *testing.T) {
	ctx := t.Context()

	for _, c := range FetchAgents() {
		initial, err := c.OriginSegmented(ctx, 0)
		require.NoError(t, err)

		parsed := make([]string, 0, initial.SumSegments)
		for i := 0; i < initial.SumSegments; i++ {
			segment, err := c.OriginSegmented(ctx, uint(i))
			require.NoError(t, err)
			data, err := base64.StdEncoding.DecodeString(segment.Data)
			require.NoError(t, err)
			parsed = append(parsed, string(data))

		}
		doc := []byte(strings.Join(parsed, "REDACTED"))

		var out kinds.OriginPaper
		require.NoError(t, cometjson.Unserialize(doc, &out),
			"REDACTED", initial, string(doc))
	}
}

func VerifyIfaceInquire(t *testing.T) {
	for i, c := range FetchAgents() {
		//
		k, v, tx := CreateTransferObject()
		bout, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.Nil(t, err, "REDACTED", i, err)
		appheading := bout.Level + 1 //

		//
		err = customer.WaitForLevel(c, appheading, nil)
		require.NoError(t, err)
		res, err := c.IfaceInquire(context.Background(), "REDACTED", k)
		inquiryout := res.Reply
		if assert.Nil(t, err) && assert.True(t, inquiryout.IsOK()) {
			assert.EqualValues(t, v, inquiryout.Item)
		}
	}
}

//
func VerifyApplicationInvocations(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)
	for i, c := range FetchAgents() {

		//
		s, err := c.Status(context.Background())
		require.NoError(err)
		//
		sh := s.AlignDetails.NewestLedgerLevel

		//
		h := sh + 20
		_, err = c.Ledger(context.Background(), &h)
		require.Error(err) //

		//
		k, v, tx := CreateTransferObject()
		bout, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.NoError(err)
		require.True(bout.TransOutcome.IsOK())
		txh := bout.Level
		appheading := txh + 1 //

		//
		err = customer.WaitForLevel(c, appheading, nil)
		require.NoError(err)

		_inquiryout, err := c.IfaceInquireWithSettings(context.Background(), "REDACTED", k, customer.IfaceInquireSettings{Demonstrate: false})
		require.NoError(err)
		inquiryout := _inquiryout.Reply
		if assert.True(inquiryout.IsOK()) {
			assert.Equal(k, inquiryout.Key)
			assert.EqualValues(v, inquiryout.Item)
		}

		//
		ptx, err := c.Tx(context.Background(), bout.Digest, true)
		require.NoError(err)
		assert.EqualValues(txh, ptx.Level)
		assert.EqualValues(tx, ptx.Tx)

		//
		ledger, err := c.Ledger(context.Background(), &appheading)
		require.NoError(err)
		applicationDigest := ledger.Ledger.ApplicationDigest
		assert.True(len(applicationDigest) > 0)
		assert.EqualValues(appheading, ledger.Ledger.Level)

		ledgerByDigest, err := c.LedgerByDigest(context.Background(), ledger.LedgerUID.Digest)
		require.NoError(err)
		require.Equal(ledger, ledgerByDigest)

		//
		heading, err := c.Heading(context.Background(), &appheading)
		require.NoError(err)
		require.Equal(ledger.Ledger.Heading, *heading.Heading)

		headingByDigest, err := c.HeadingByDigest(context.Background(), ledger.LedgerUID.Digest)
		require.NoError(err)
		require.Equal(heading, headingByDigest)

		//
		ledgerOutcomes, err := c.LedgerOutcomes(context.Background(), &txh)
		require.Nil(err, "REDACTED", i, err)
		assert.Equal(txh, ledgerOutcomes.Level)
		if assert.Equal(1, len(ledgerOutcomes.TransOutcomes)) {
			//
			assert.EqualValues(0, ledgerOutcomes.TransOutcomes[0].Code)
		}

		//
		details, err := c.LedgerchainDetails(context.Background(), appheading, appheading)
		require.NoError(err)
		assert.True(details.FinalLevel >= appheading)
		if assert.Equal(1, len(details.LedgerMetadata)) {
			finalMeta := details.LedgerMetadata[0]
			assert.EqualValues(appheading, finalMeta.Heading.Level)
			ledgerData := ledger.Ledger
			assert.Equal(ledgerData.ApplicationDigest, finalMeta.Heading.ApplicationDigest)
			assert.Equal(ledger.LedgerUID, finalMeta.LedgerUID)
		}

		//
		endorse, err := c.Endorse(context.Background(), &appheading)
		require.NoError(err)
		capplicationDigest := endorse.ApplicationDigest
		assert.Equal(applicationDigest, capplicationDigest)
		assert.NotNil(endorse.Endorse)

		//
		h = appheading - 1
		confirm2, err := c.Endorse(context.Background(), &h)
		require.NoError(err)
		assert.Equal(ledger.Ledger.FinalEndorseDigest, confirm2.Endorse.Digest())

		//
		_pout, err := c.IfaceInquireWithSettings(context.Background(), "REDACTED", k, customer.IfaceInquireSettings{Demonstrate: true})
		require.NoError(err)
		pout := _pout.Reply
		assert.True(pout.IsOK())

		//
	}
}

func VerifyMulticastTransferAlign(t *testing.T) {
	demand := require.New(t)

	//
	txpool := member.Txpool()
	initTxpoolVolume := txpool.Volume()

	for i, c := range FetchAgents() {
		_, _, tx := CreateTransferObject()
		bout, err := c.MulticastTransferAlign(context.Background(), tx)
		require.Nil(err, "REDACTED", i, err)
		require.Equal(bout.Code, iface.CodeKindSuccess) //

		require.Equal(initTxpoolVolume+1, txpool.Volume())

		txs := txpool.HarvestMaximumTrans(len(tx))
		require.EqualValues(tx, txs[0])
		txpool.Purge()
	}
}

func VerifyMulticastTransferEndorse(t *testing.T) {
	demand := require.New(t)

	txpool := member.Txpool()
	for i, c := range FetchAgents() {
		_, _, tx := CreateTransferObject()
		bout, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.Nil(err, "REDACTED", i, err)
		require.True(bout.InspectTransfer.IsOK())
		require.True(bout.TransOutcome.IsOK())

		require.Equal(0, txpool.Volume())
	}
}

func VerifyUnattestedTrans(t *testing.T) {
	_, _, tx := CreateTransferObject()

	ch := make(chan *iface.ReplyInspectTransfer, 1)
	txpool := member.Txpool()
	err := txpool.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) { ch <- reply }, txpool.TransferDetails{})
	require.NoError(t, err)

	//
	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.Error("REDACTED")
	}

	for _, c := range FetchAgents() {
		mc := c.(customer.TxpoolCustomer)
		ceiling := 1
		res, err := mc.UnattestedTrans(context.Background(), &ceiling)
		require.NoError(t, err)

		assert.Equal(t, 1, res.Number)
		assert.Equal(t, 1, res.Sum)
		assert.Equal(t, txpool.VolumeOctets(), res.SumOctets)
		assert.Exactly(t, kinds.Txs{tx}, kinds.Txs(res.Txs))
	}

	txpool.Purge()
}

func VerifyCountUnattestedTrans(t *testing.T) {
	_, _, tx := CreateTransferObject()

	ch := make(chan *iface.ReplyInspectTransfer, 1)
	txpool := member.Txpool()
	err := txpool.InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) { ch <- reply }, txpool.TransferDetails{})
	require.NoError(t, err)

	//
	select {
	case <-ch:
	case <-time.After(5 * time.Second):
		t.Error("REDACTED")
	}

	txpoolVolume := txpool.Volume()
	for i, c := range FetchAgents() {
		mc, ok := c.(customer.TxpoolCustomer)
		require.True(t, ok, "REDACTED", i)
		res, err := mc.CountUnattestedTrans(context.Background())
		require.Nil(t, err, "REDACTED", i, err)

		assert.Equal(t, txpoolVolume, res.Number)
		assert.Equal(t, txpoolVolume, res.Sum)
		assert.Equal(t, txpool.VolumeOctets(), res.SumOctets)
	}

	txpool.Purge()
}

func VerifyInspectTransfer(t *testing.T) {
	txpool := member.Txpool()

	for _, c := range FetchAgents() {
		_, _, tx := CreateTransferObject()

		res, err := c.InspectTransfer(context.Background(), tx)
		require.NoError(t, err)
		assert.Equal(t, iface.CodeKindSuccess, res.Code)

		assert.Equal(t, 0, txpool.Volume(), "REDACTED")
	}
}

func VerifyTransfer(t *testing.T) {
	//
	c := fetchHTTPCustomer()
	_, _, tx := CreateTransferObject()
	bout, err := c.MulticastTransferEndorse(context.Background(), tx)
	require.Nil(t, err, "REDACTED", err)

	transferLevel := bout.Level
	transferDigest := bout.Digest

	otherTransferDigest := kinds.Tx("REDACTED").Digest()

	scenarios := []struct {
		sound bool
		demonstrate bool
		digest  []byte
	}{
		//
		{true, false, transferDigest},
		{true, true, transferDigest},
		{false, false, otherTransferDigest},
		{false, true, otherTransferDigest},
		{false, false, nil},
		{false, true, nil},
	}

	for i, c := range FetchAgents() {
		for j, tc := range scenarios {
			t.Logf("REDACTED", i, j)

			//
			//
			ptx, err := c.Tx(context.Background(), tc.digest, tc.demonstrate)

			if !tc.sound {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err, "REDACTED", err)
				assert.EqualValues(t, transferLevel, ptx.Level)
				assert.EqualValues(t, tx, ptx.Tx)
				assert.Zero(t, ptx.Ordinal)
				assert.True(t, ptx.TransOutcome.IsOK())
				assert.EqualValues(t, transferDigest, ptx.Digest)

				//
				evidence := ptx.Attestation
				if tc.demonstrate && assert.EqualValues(t, tx, evidence.Data) {
					assert.NoError(t, evidence.Attestation.Validate(evidence.OriginDigest, transferDigest))
				}
			}
		}
	}
}

func VerifyTransferScanWithDeadline(t *testing.T) {
	//
	deadlineCustomer := fetchHTTPCustomerWithDeadline(10)

	_, _, tx := CreateTransferObject()
	_, err := deadlineCustomer.MulticastTransferEndorse(context.Background(), tx)
	require.NoError(t, err)

	//
	outcome, err := deadlineCustomer.TransferScan(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
	require.Nil(t, err)
	require.Greater(t, len(outcome.Txs), 0, "REDACTED")
}

//
//
func VerifyLedgerScan(t *testing.T) {
	c := fetchHTTPCustomer()

	//
	for i := 0; i < 10; i++ {
		_, _, tx := CreateTransferObject()

		_, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.NoError(t, err)
	}
	require.NoError(t, customer.WaitForLevel(c, 5, nil))
	outcome, err := c.LedgerScan(context.Background(), "REDACTED", nil, nil, "REDACTED")
	require.NoError(t, err)
	ledgerNumber := len(outcome.Ledgers)
	//
	//
	//
	//

	//
	require.Equal(t, ledgerNumber, 0)
}

func VerifyTransferScan(t *testing.T) {
	c := fetchHTTPCustomer()

	//
	for i := 0; i < 10; i++ {
		_, _, tx := CreateTransferObject()
		_, err := c.MulticastTransferEndorse(context.Background(), tx)
		require.NoError(t, err)
	}

	//
	//
	outcome, err := c.TransferScan(context.Background(), "REDACTED", true, nil, nil, "REDACTED")
	require.NoError(t, err)
	transferNumber := len(outcome.Txs)

	//
	locate := outcome.Txs[len(outcome.Txs)-1]
	otherTransferDigest := kinds.Tx("REDACTED").Digest()

	for _, c := range FetchAgents() {

		//
		outcome, err := c.TransferScan(context.Background(), fmt.Sprintf("REDACTED", locate.Digest), true, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 1)
		require.Equal(t, locate.Digest, outcome.Txs[0].Digest)

		ptx := outcome.Txs[0]
		assert.EqualValues(t, locate.Level, ptx.Level)
		assert.EqualValues(t, locate.Tx, ptx.Tx)
		assert.Zero(t, ptx.Ordinal)
		assert.True(t, ptx.TransOutcome.IsOK())
		assert.EqualValues(t, locate.Digest, ptx.Digest)

		//
		if assert.EqualValues(t, locate.Tx, ptx.Attestation.Data) {
			assert.NoError(t, ptx.Attestation.Attestation.Validate(ptx.Attestation.OriginDigest, locate.Digest))
		}

		//
		outcome, err = c.TransferScan(context.Background(), fmt.Sprintf("REDACTED", locate.Level), true, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 1)

		//
		outcome, err = c.TransferScan(context.Background(), fmt.Sprintf("REDACTED", otherTransferDigest), false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 0)

		//
		outcome, err = c.TransferScan(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Greater(t, len(outcome.Txs), 0, "REDACTED")

		//
		outcome, err = c.TransferScan(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Greater(t, len(outcome.Txs), 0, "REDACTED")

		//
		outcome, err = c.TransferScan(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Equal(t, len(outcome.Txs), 0, "REDACTED")

		//
		outcome, err = c.TransferScan(context.Background(),
			"REDACTED", true, nil, nil, "REDACTED")
		require.Nil(t, err)
		require.Greater(t, len(outcome.Txs), 0, "REDACTED")

		//
		eachScreen := 1
		outcome, err = c.TransferScan(context.Background(), "REDACTED", true, nil, &eachScreen, "REDACTED")
		require.Nil(t, err)
		require.Len(t, outcome.Txs, 0)

		//
		outcome, err = c.TransferScan(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		for k := 0; k < len(outcome.Txs)-1; k++ {
			require.LessOrEqual(t, outcome.Txs[k].Level, outcome.Txs[k+1].Level)
			require.LessOrEqual(t, outcome.Txs[k].Ordinal, outcome.Txs[k+1].Ordinal)
		}

		outcome, err = c.TransferScan(context.Background(), "REDACTED", false, nil, nil, "REDACTED")
		require.Nil(t, err)
		for k := 0; k < len(outcome.Txs)-1; k++ {
			require.GreaterOrEqual(t, outcome.Txs[k].Level, outcome.Txs[k+1].Level)
			require.GreaterOrEqual(t, outcome.Txs[k].Ordinal, outcome.Txs[k+1].Ordinal)
		}
		//
		eachScreen = 3
		var (
			viewed      = map[int64]bool{}
			maximumLevel int64
			sections     = int(math.Ceil(float64(transferNumber) / float64(eachScreen)))
		)

		sumTransfer := 0
		for screen := 1; screen <= sections; screen++ {

			outcome, err := c.TransferScan(context.Background(), "REDACTED", true, &screen, &eachScreen, "REDACTED")
			require.NoError(t, err)
			if screen < sections {
				require.Len(t, outcome.Txs, eachScreen)
			} else {
				require.LessOrEqual(t, len(outcome.Txs), eachScreen)
			}
			sumTransfer += len(outcome.Txs)
			for _, tx := range outcome.Txs {
				require.False(t, viewed[tx.Level],
					"REDACTED", tx.Level, screen)
				require.Greater(t, tx.Level, maximumLevel,
					"REDACTED", tx.Level, maximumLevel, screen)
				viewed[tx.Level] = true
				maximumLevel = tx.Level
			}
		}
		require.Equal(t, transferNumber, sumTransfer)
		require.Len(t, viewed, transferNumber)
	}
}

func VerifySegmentedJsonrpcInvocations(t *testing.T) {
	c := fetchHTTPCustomer()
	verifySegmentedJsonrpcInvocations(t, c)
}

func verifySegmentedJsonrpcInvocations(t *testing.T, c *rpchttp.HTTP) {
	k1, v1, tx1 := CreateTransferObject()
	k2, v2, tx2 := CreateTransferObject()

	group := c.NewGroup()
	r1, err := group.MulticastTransferEndorse(context.Background(), tx1)
	require.NoError(t, err)
	r2, err := group.MulticastTransferEndorse(context.Background(), tx2)
	require.NoError(t, err)
	require.Equal(t, 2, group.Number())
	boutcomes, err := group.Transmit(ctx)
	require.NoError(t, err)
	require.Len(t, boutcomes, 2)
	require.Equal(t, 0, group.Number())

	boutcome1, ok := boutcomes[0].(*ctypes.OutcomeMulticastTransferEndorse)
	require.True(t, ok)
	require.Equal(t, *boutcome1, *r1)
	boutcome2, ok := boutcomes[1].(*ctypes.OutcomeMulticastTransferEndorse)
	require.True(t, ok)
	require.Equal(t, *boutcome2, *r2)
	appheading := cometmath.MaximumInt64(boutcome1.Level, boutcome2.Level) + 1

	err = customer.WaitForLevel(c, appheading, nil)
	require.NoError(t, err)

	q1, err := group.IfaceInquire(context.Background(), "REDACTED", k1)
	require.NoError(t, err)
	q2, err := group.IfaceInquire(context.Background(), "REDACTED", k2)
	require.NoError(t, err)
	require.Equal(t, 2, group.Number())
	inquiryoutcomes, err := group.Transmit(ctx)
	require.NoError(t, err)
	require.Len(t, inquiryoutcomes, 2)
	require.Equal(t, 0, group.Number())

	inquiryoutcome1, ok := inquiryoutcomes[0].(*ctypes.OutcomeIfaceInquire)
	require.True(t, ok)
	require.Equal(t, *inquiryoutcome1, *q1)
	inquiryoutcome2, ok := inquiryoutcomes[1].(*ctypes.OutcomeIfaceInquire)
	require.True(t, ok)
	require.Equal(t, *inquiryoutcome2, *q2)

	require.Equal(t, inquiryoutcome1.Reply.Key, k1)
	require.Equal(t, inquiryoutcome2.Reply.Key, k2)
	require.Equal(t, inquiryoutcome1.Reply.Item, v1)
	require.Equal(t, inquiryoutcome2.Reply.Item, v2)
}

func VerifySegmentedJsonrpcInvocationsAbort(t *testing.T) {
	c := fetchHTTPCustomer()
	_, _, tx1 := CreateTransferObject()
	_, _, tx2 := CreateTransferObject()

	group := c.NewGroup()
	_, err := group.MulticastTransferEndorse(context.Background(), tx1)
	require.NoError(t, err)
	_, err = group.MulticastTransferEndorse(context.Background(), tx2)
	require.NoError(t, err)
	//
	require.Equal(t, 2, group.Number())
	//
	require.Equal(t, 2, group.Flush())
	//
	require.Equal(t, 0, group.Number())
}

func VerifyDispatchingEmptyQuerySegment(t *testing.T) {
	c := fetchHTTPCustomer()
	group := c.NewGroup()
	_, err := group.Transmit(ctx)
	require.Error(t, err, "REDACTED")
}

func VerifyPurgingEmptyQuerySegment(t *testing.T) {
	c := fetchHTTPCustomer()
	group := c.NewGroup()
	require.Zero(t, group.Flush(), "REDACTED")
}

func VerifyParallelJsonrpcSegmenting(t *testing.T) {
	var wg sync.WaitGroup
	c := fetchHTTPCustomer()
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			verifySegmentedJsonrpcInvocations(t, c)
		}()
	}
	wg.Wait()
}
