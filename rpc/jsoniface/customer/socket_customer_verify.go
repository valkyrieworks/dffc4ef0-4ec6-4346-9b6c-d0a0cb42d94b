package customer

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

var socketInvocationDeadline = 5 * time.Second

type mineProcessor struct {
	shutdownLinkSubsequentFetch bool
	mtx                commitchronize.ReadwriteExclusion
}

var enhancer = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *mineProcessor) AttendHttpsvc(w http.ResponseWriter, r *http.Request) {
	link, err := enhancer.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer link.Close()
	for {
		signalKind, in, err := link.ReadMessage()
		if err != nil {
			return
		}

		var req kinds.RemoteSolicit
		err = json.Unmarshal(in, &req)
		if err != nil {
			panic(err)
		}

		h.mtx.RLock()
		if h.shutdownLinkSubsequentFetch {
			if err := link.Close(); err != nil {
				panic(err)
			}
		}
		h.mtx.RUnlock()

		res := json.RawMessage("REDACTED")
		blankAnswerOctets, _ := json.Marshal(kinds.RemoteReply{Outcome: res, ID: req.ID})
		if err := link.WriteMessage(signalKind, blankAnswerOctets); err != nil {
			return
		}
	}
}

func VerifySocketCustomerReestablishesSubsequentFetchBreakdown(t *testing.T) {
	var wg sync.WaitGroup

	//
	h := &mineProcessor{}
	s := httptest.NewServer(h)
	defer s.Close()

	c := initiateCustomer(t, "REDACTED"+s.Listener.Addr().String())
	defer c.Halt() //

	wg.Add(1)
	go invocationGroupCompleteUponOutcome(t, c, &wg)

	h.mtx.Lock()
	h.shutdownLinkSubsequentFetch = true
	h.mtx.Unlock()

	//
	invocation(t, "REDACTED", c)

	//
	time.Sleep(10 * time.Millisecond)
	h.mtx.Lock()
	h.shutdownLinkSubsequentFetch = false
	h.mtx.Unlock()

	//
	invocation(t, "REDACTED", c)

	wg.Wait()
}

func VerifySocketCustomerReestablishesSubsequentRecordBreakdown(t *testing.T) {
	var wg sync.WaitGroup

	//
	h := &mineProcessor{}
	s := httptest.NewServer(h)

	c := initiateCustomer(t, "REDACTED"+s.Listener.Addr().String())
	defer c.Halt() //

	wg.Add(2)
	go invocationGroupCompleteUponOutcome(t, c, &wg)

	//
	if err := c.link.Close(); err != nil {
		t.Error(err)
	}

	//
	invocation(t, "REDACTED", c)

	//
	time.Sleep(10 * time.Millisecond)

	//
	invocation(t, "REDACTED", c)

	wg.Wait()
}

func VerifySocketCustomerReestablishBreakdown(t *testing.T) {
	//
	h := &mineProcessor{}
	s := httptest.NewServer(h)

	c := initiateCustomer(t, "REDACTED"+s.Listener.Addr().String())
	defer c.Halt() //

	go func() {
		for {
			select {
			case <-c.RepliesChnl:
			case <-c.Exit():
				return
			}
		}
	}()

	//
	if err := c.link.Close(); err != nil {
		t.Error(err)
	}
	s.Close()

	//
	//
	ctx, abort := context.WithTimeout(context.Background(), socketInvocationDeadline)
	defer abort()
	if err := c.Invocation(ctx, "REDACTED", make(map[string]any)); err != nil {
		t.Error(err)
	}

	//
	time.Sleep(10 * time.Millisecond)

	complete := make(chan struct{})
	go func() {
		//
		invocation(t, "REDACTED", c)
		close(complete)
	}()

	//
	select {
	case <-complete:
		t.Fatal("REDACTED")
	case <-time.After(5 * time.Second):
		t.Log("REDACTED")
	}
}

func VerifyNegationHaltingUponHalt(t *testing.T) {
	deadline := 2 * time.Second
	s := httptest.NewServer(&mineProcessor{})
	c := initiateCustomer(t, "REDACTED"+s.Listener.Addr().String())
	c.Invocation(context.Background(), "REDACTED", make(map[string]any)) //
	//
	time.Sleep(time.Second)
	phraseChnl := make(chan struct{})
	go func() {
		//
		//
		err := c.Halt()
		require.NoError(t, err)
		phraseChnl <- struct{}{}
	}()
	select {
	case <-phraseChnl:
		//
	case <-time.After(deadline):
		t.Fatalf("REDACTED",
			deadline.Seconds())
	}
}

func initiateCustomer(t *testing.T, location string) *SocketCustomer {
	c, err := FreshSocket(location, "REDACTED")
	require.Nil(t, err)
	err = c.Initiate()
	require.Nil(t, err)
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func invocation(t *testing.T, procedure string, c *SocketCustomer) {
	err := c.Invocation(context.Background(), procedure, make(map[string]any))
	require.NoError(t, err)
}

func invocationGroupCompleteUponOutcome(t *testing.T, c *SocketCustomer, wg *sync.WaitGroup) {
	for {
		select {
		case reply := <-c.RepliesChnl:
			if reply.Failure != nil {
				t.Errorf("REDACTED", reply.Failure)
				return
			}
			if reply.Outcome != nil {
				wg.Done()
			}
		case <-c.Exit():
			return
		}
	}
}
