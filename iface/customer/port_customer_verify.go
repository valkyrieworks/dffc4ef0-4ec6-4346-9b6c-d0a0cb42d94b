package ifacec_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

func VerifyInvocations(t *testing.T) {
	ctx := t.Context()
	app := kinds.FoundationPlatform{}

	_, c := configureCustomerDaemon(t, app)

	reply := make(chan error, 1)
	go func() {
		res, err := c.Reverberate(ctx, "REDACTED")
		require.NoError(t, err)
		require.NotNil(t, res)
		reply <- c.Failure()
	}()

	select {
	case <-time.After(time.Second):
		require.Fail(t, "REDACTED")
	case err, ok := <-reply:
		require.True(t, ok, "REDACTED")
		assert.NoError(t, err, "REDACTED")
	}
}

func VerifyLingeringAsyncronousInvocations(t *testing.T) {
	app := gradualApplication{}

	s, c := configureCustomerDaemon(t, app)

	reply := make(chan error, 1)
	go func() {
		//
		requestresponse, err := c.InspectTransferAsyncronous(context.Background(), &kinds.SolicitInspectTransfer{})
		require.NoError(t, err)
		//
		//
		time.Sleep(50 * time.Millisecond)
		//
		err = s.Halt()
		require.NoError(t, err)

		//
		requestresponse.Wait()
		reply <- c.Failure()
	}()

	select {
	case <-time.After(time.Second):
		require.Fail(t, "REDACTED")
	case err, ok := <-reply:
		require.True(t, ok, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}
}

func VerifyLump(t *testing.T) {
	const countTrans = 700000
	//
	portRecord := fmt.Sprintf("REDACTED", rand.Int31n(1<<30))
	defer os.Remove(portRecord)
	port := fmt.Sprintf("REDACTED", portRecord)
	app := kinds.FreshFoundationPlatform()
	//
	node := node.FreshPortDaemon(port, app)
	t.Cleanup(func() {
		if err := node.Halt(); err != nil {
			t.Log(err)
		}
	})
	err := node.Initiate()
	require.NoError(t, err)

	//
	customer := abcicustomer.FreshPortCustomer(port, false)

	t.Cleanup(func() {
		if err := customer.Halt(); err != nil {
			t.Log(err)
		}
	})

	err = customer.Initiate()
	require.NoError(t, err)

	//
	rfb := &kinds.SolicitCulminateLedger{Txs: make([][]byte, countTrans)}
	for tally := 0; tally < countTrans; tally++ {
		rfb.Txs[tally] = []byte("REDACTED")
	}
	//
	res, err := customer.CulminateLedger(context.Background(), rfb)
	require.NoError(t, err)
	require.Equal(t, countTrans, len(res.TransferOutcomes), "REDACTED")
	for _, tx := range res.TransferOutcomes {
		require.Equal(t, uint32(0), tx.Cipher, "REDACTED")
	}

	//
	err = customer.Purge(context.Background())
	require.NoError(t, err)
}

func configureCustomerDaemon(t *testing.T, app kinds.Platform) (
	facility.Facility, abcicustomer.Customer,
) {
	t.Helper()

	//
	channel := 20000 + commitrand.Integer32()%10000
	location := fmt.Sprintf("REDACTED", channel)

	s := node.FreshPortDaemon(location, app)
	err := s.Initiate()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := s.Halt(); err != nil {
			t.Log(err)
		}
	})

	c := abcicustomer.FreshPortCustomer(location, true)
	err = c.Initiate()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := c.Halt(); err != nil {
			t.Log(err)
		}
	})

	return s, c
}

type gradualApplication struct {
	kinds.FoundationPlatform
}

func (gradualApplication) InspectTransfer(context.Context, *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	time.Sleep(time.Second)
	return &kinds.ReplyInspectTransfer{}, nil
}

//
//
//
//
func VerifyClbkExecutedWheneverAssignTardy(t *testing.T) {
	ctx := t.Context()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	app := obstructedIfacePlatform{
		wg: wg,
	}
	_, c := configureCustomerDaemon(t, app)
	requestResult, err := c.InspectTransferAsyncronous(ctx, &kinds.SolicitInspectTransfer{})
	require.NoError(t, err)

	complete := make(chan struct{})
	cb := func(_ *kinds.Reply) {
		close(complete)
	}
	requestResult.AssignClbk(cb)
	app.wg.Done()
	<-complete

	var invoked bool
	cb = func(_ *kinds.Reply) {
		invoked = true
	}
	requestResult.AssignClbk(cb)
	require.True(t, invoked)
}

type obstructedIfacePlatform struct {
	wg *sync.WaitGroup
	kinds.FoundationPlatform
}

func (b obstructedIfacePlatform) InspectTransferAsyncronous(ctx context.Context, r *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	b.wg.Wait()
	return b.InspectTransfer(ctx, r)
}

//
//
func VerifyClbkExecutedWheneverAssignPremature(t *testing.T) {
	ctx := t.Context()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	app := obstructedIfacePlatform{
		wg: wg,
	}
	_, c := configureCustomerDaemon(t, app)
	requestResult, err := c.InspectTransferAsyncronous(ctx, &kinds.SolicitInspectTransfer{})
	require.NoError(t, err)

	complete := make(chan struct{})
	cb := func(_ *kinds.Reply) {
		close(complete)
	}
	requestResult.AssignClbk(cb)
	app.wg.Done()

	invoked := func() bool {
		select {
		case <-complete:
			return true
		default:
			return false
		}
	}
	require.Eventually(t, invoked, time.Second, time.Millisecond*25)
}
