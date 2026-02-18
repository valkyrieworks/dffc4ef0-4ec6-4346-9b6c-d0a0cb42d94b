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

	"github.com/valkyrieworks/utils/log"
	engineconnect "github.com/valkyrieworks/utils/align"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

var wsInvocationDeadline = 5 * time.Second

type mineManager struct {
	endLinkAfterRead bool
	mtx                engineconnect.ReadwriteLock
}

var converter = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *mineManager) AttendHTTP(w http.ResponseWriter, r *http.Request) {
	link, err := converter.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer link.Close()
	for {
		signalKind, in, err := link.ReadMessage()
		if err != nil {
			return
		}

		var req kinds.RPCQuery
		err = json.Unmarshal(in, &req)
		if err != nil {
			panic(err)
		}

		h.mtx.RLock()
		if h.endLinkAfterRead {
			if err := link.Close(); err != nil {
				panic(err)
			}
		}
		h.mtx.RUnlock()

		res := json.RawMessage("REDACTED")
		emptyReplyOctets, _ := json.Marshal(kinds.RPCAnswer{Outcome: res, ID: req.ID})
		if err := link.WriteMessage(signalKind, emptyReplyOctets); err != nil {
			return
		}
	}
}

func VerifyWSCustomerReestablishesAfterReadBreakdown(t *testing.T) {
	var wg sync.WaitGroup

	//
	h := &mineManager{}
	s := httptest.NewServer(h)
	defer s.Close()

	c := beginCustomer(t, "REDACTED"+s.Listener.Addr().String())
	defer c.Halt() //

	wg.Add(1)
	go invocationGroupDoneOnOutcome(t, c, &wg)

	h.mtx.Lock()
	h.endLinkAfterRead = true
	h.mtx.Unlock()

	//
	invocation(t, "REDACTED", c)

	//
	time.Sleep(10 * time.Millisecond)
	h.mtx.Lock()
	h.endLinkAfterRead = false
	h.mtx.Unlock()

	//
	invocation(t, "REDACTED", c)

	wg.Wait()
}

func VerifyWSCustomerReestablishesAfterRecordBreakdown(t *testing.T) {
	var wg sync.WaitGroup

	//
	h := &mineManager{}
	s := httptest.NewServer(h)

	c := beginCustomer(t, "REDACTED"+s.Listener.Addr().String())
	defer c.Halt() //

	wg.Add(2)
	go invocationGroupDoneOnOutcome(t, c, &wg)

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

func VerifyWSCustomerReestablishBreakdown(t *testing.T) {
	//
	h := &mineManager{}
	s := httptest.NewServer(h)

	c := beginCustomer(t, "REDACTED"+s.Listener.Addr().String())
	defer c.Halt() //

	go func() {
		for {
			select {
			case <-c.RepliesChan:
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
	ctx, revoke := context.WithTimeout(context.Background(), wsInvocationDeadline)
	defer revoke()
	if err := c.Invoke(ctx, "REDACTED", make(map[string]any)); err != nil {
		t.Error(err)
	}

	//
	time.Sleep(10 * time.Millisecond)

	done := make(chan struct{})
	go func() {
		//
		invocation(t, "REDACTED", c)
		close(done)
	}()

	//
	select {
	case <-done:
		t.Fatal("REDACTED")
	case <-time.After(5 * time.Second):
		t.Log("REDACTED")
	}
}

func VerifyNegateHaltingOnHalt(t *testing.T) {
	deadline := 2 * time.Second
	s := httptest.NewServer(&mineManager{})
	c := beginCustomer(t, "REDACTED"+s.Listener.Addr().String())
	c.Invoke(context.Background(), "REDACTED", make(map[string]any)) //
	//
	time.Sleep(time.Second)
	passChan := make(chan struct{})
	go func() {
		//
		//
		err := c.Halt()
		require.NoError(t, err)
		passChan <- struct{}{}
	}()
	select {
	case <-passChan:
		//
	case <-time.After(deadline):
		t.Fatalf("REDACTED",
			deadline.Seconds())
	}
}

func beginCustomer(t *testing.T, address string) *WSCustomer {
	c, err := NewWS(address, "REDACTED")
	require.Nil(t, err)
	err = c.Begin()
	require.Nil(t, err)
	c.AssignTracer(log.VerifyingTracer())
	return c
}

func invocation(t *testing.T, procedure string, c *WSCustomer) {
	err := c.Invoke(context.Background(), procedure, make(map[string]any))
	require.NoError(t, err)
}

func invocationGroupDoneOnOutcome(t *testing.T, c *WSCustomer, wg *sync.WaitGroup) {
	for {
		select {
		case reply := <-c.RepliesChan:
			if reply.Fault != nil {
				t.Errorf("REDACTED", reply.Fault)
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
