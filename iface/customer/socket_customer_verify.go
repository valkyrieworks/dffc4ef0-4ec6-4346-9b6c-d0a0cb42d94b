package abcic_test

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

	abciend "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/host"
	"github.com/valkyrieworks/iface/kinds"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/utils/daemon"
)

func VerifyInvocations(t *testing.T) {
	ctx := t.Context()
	app := kinds.RootSoftware{}

	_, c := configureCustomerHost(t, app)

	reply := make(chan error, 1)
	go func() {
		res, err := c.Replicate(ctx, "REDACTED")
		require.NoError(t, err)
		require.NotNil(t, res)
		reply <- c.Fault()
	}()

	select {
	case <-time.After(time.Second):
		require.Fail(t, "REDACTED")
	case err, ok := <-reply:
		require.True(t, ok, "REDACTED")
		assert.NoError(t, err, "REDACTED")
	}
}

func VerifyPendingAsyncInvocations(t *testing.T) {
	app := gradualApplication{}

	s, c := configureCustomerHost(t, app)

	reply := make(chan error, 1)
	go func() {
		//
		requestresponse, err := c.InspectTransferAsync(context.Background(), &kinds.QueryInspectTransfer{})
		require.NoError(t, err)
		//
		//
		time.Sleep(50 * time.Millisecond)
		//
		err = s.Halt()
		require.NoError(t, err)

		//
		requestresponse.Wait()
		reply <- c.Fault()
	}()

	select {
	case <-time.After(time.Second):
		require.Fail(t, "REDACTED")
	case err, ok := <-reply:
		require.True(t, ok, "REDACTED")
		assert.Error(t, err, "REDACTED")
	}
}

func VerifyBatch(t *testing.T) {
	const countTrans = 700000
	//
	socketEntry := fmt.Sprintf("REDACTED", rand.Int31n(1<<30))
	defer os.Remove(socketEntry)
	socket := fmt.Sprintf("REDACTED", socketEntry)
	app := kinds.NewRootSoftware()
	//
	host := host.NewSocketHost(socket, app)
	t.Cleanup(func() {
		if err := host.Halt(); err != nil {
			t.Log(err)
		}
	})
	err := host.Begin()
	require.NoError(t, err)

	//
	customer := abciend.NewSocketCustomer(socket, false)

	t.Cleanup(func() {
		if err := customer.Halt(); err != nil {
			t.Log(err)
		}
	})

	err = customer.Begin()
	require.NoError(t, err)

	//
	rfb := &kinds.QueryCompleteLedger{Txs: make([][]byte, countTrans)}
	for tally := 0; tally < countTrans; tally++ {
		rfb.Txs[tally] = []byte("REDACTED")
	}
	//
	res, err := customer.CompleteLedger(context.Background(), rfb)
	require.NoError(t, err)
	require.Equal(t, countTrans, len(res.TransOutcomes), "REDACTED")
	for _, tx := range res.TransOutcomes {
		require.Equal(t, uint32(0), tx.Code, "REDACTED")
	}

	//
	err = customer.Purge(context.Background())
	require.NoError(t, err)
}

func configureCustomerHost(t *testing.T, app kinds.Software) (
	daemon.Daemon, abciend.Customer,
) {
	t.Helper()

	//
	port := 20000 + engineseed.Int32()%10000
	address := fmt.Sprintf("REDACTED", port)

	s := host.NewSocketHost(address, app)
	err := s.Begin()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := s.Halt(); err != nil {
			t.Log(err)
		}
	})

	c := abciend.NewSocketCustomer(address, true)
	err = c.Begin()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := c.Halt(); err != nil {
			t.Log(err)
		}
	})

	return s, c
}

type gradualApplication struct {
	kinds.RootSoftware
}

func (gradualApplication) InspectTransfer(context.Context, *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	time.Sleep(time.Second)
	return &kinds.ReplyInspectTransfer{}, nil
}

//
//
//
//
func VerifyCallbackExecutedWhenCollectionTardy(t *testing.T) {
	ctx := t.Context()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	app := obstructedIfaceSoftware{
		wg: wg,
	}
	_, c := configureCustomerHost(t, app)
	requestOutput, err := c.InspectTransferAsync(ctx, &kinds.QueryInspectTransfer{})
	require.NoError(t, err)

	done := make(chan struct{})
	cb := func(_ *kinds.Reply) {
		close(done)
	}
	requestOutput.CollectionCallback(cb)
	app.wg.Done()
	<-done

	var invoked bool
	cb = func(_ *kinds.Reply) {
		invoked = true
	}
	requestOutput.CollectionCallback(cb)
	require.True(t, invoked)
}

type obstructedIfaceSoftware struct {
	wg *sync.WaitGroup
	kinds.RootSoftware
}

func (b obstructedIfaceSoftware) InspectTransferAsync(ctx context.Context, r *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	b.wg.Wait()
	return b.InspectTransfer(ctx, r)
}

//
//
func VerifyCallbackExecutedWhenCollectionPremature(t *testing.T) {
	ctx := t.Context()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	app := obstructedIfaceSoftware{
		wg: wg,
	}
	_, c := configureCustomerHost(t, app)
	requestOutput, err := c.InspectTransferAsync(ctx, &kinds.QueryInspectTransfer{})
	require.NoError(t, err)

	done := make(chan struct{})
	cb := func(_ *kinds.Reply) {
		close(done)
	}
	requestOutput.CollectionCallback(cb)
	app.wg.Done()

	invoked := func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	}
	require.Eventually(t, invoked, time.Second, time.Millisecond*25)
}
