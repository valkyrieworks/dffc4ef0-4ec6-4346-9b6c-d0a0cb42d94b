package host

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

type specimenOutcome struct {
	Item string `json:"item"`
}

func VerifyMaximumAccessLinkages(t *testing.T) {
	const max = 5 //

	//
	var access int32
	mux := http.NewServeMux()
	mux.HandleFunc("REDACTED", func(w http.ResponseWriter, r *http.Request) {
		if n := atomic.AddInt32(&access, 1); n > int32(max) {
			t.Errorf("REDACTED", n, max)
		}
		defer atomic.AddInt32(&access, -1)
		time.Sleep(10 * time.Millisecond)
		fmt.Fprint(w, "REDACTED")
	})
	settings := StandardSettings()
	l, err := Observe("REDACTED", max)
	require.NoError(t, err)
	defer l.Close()
	go Attend(l, mux, log.VerifyingTracer(), settings) //

	//
	tries := max * 2
	var wg sync.WaitGroup
	var errored int32
	for i := 0; i < tries; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := http.Client{Timeout: 3 * time.Second}
			r, err := c.Get("REDACTED" + l.Addr().String())
			if err != nil {
				atomic.AddInt32(&errored, 1)
				return
			}
			defer r.Body.Close()
		}()
	}
	wg.Wait()

	//
	//
	if int(errored) >= tries/2 {
		t.Errorf("REDACTED", errored, tries)
	}
}

func VerifyAttendTLS(t *testing.T) {
	ln, err := net.Listen("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer ln.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("REDACTED", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "REDACTED")
	})

	chanErr := make(chan error, 1)
	go func() {
		//
		chanErr <- AttendTLS(ln, mux, "REDACTED", "REDACTED", log.VerifyingTracer(), StandardSettings())
	}()

	select {
	case err := <-chanErr:
		require.NoError(t, err)
	case <-time.After(100 * time.Millisecond):
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}
	res, err := c.Get("REDACTED" + ln.Addr().String())
	require.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	content, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	assert.Equal(t, []byte("REDACTED"), content)
}

func VerifyRecordRPCReplyHTTP(t *testing.T) {
	id := kinds.JsonrpcIntegerUID(-1)

	//
	w := httptest.NewRecorder()
	err := RecordStorableRPCReplyHTTP(w, kinds.NewRPCSuccessReply(id, &specimenOutcome{"REDACTED"}))
	require.NoError(t, err)
	reply := w.Result()
	content, err := io.ReadAll(reply.Body)
	_ = reply.Body.Close()
	require.NoError(t, err)
	assert.Equal(t, 200, reply.StatusCode)
	assert.Equal(t, "REDACTED", reply.Header.Get("REDACTED"))
	assert.Equal(t, "REDACTED", reply.Header.Get("REDACTED"))
	assert.Equal(t, "REDACTED", string(content))

	//
	w = httptest.NewRecorder()
	err = RecordRPCReplyHTTP(w,
		kinds.NewRPCSuccessReply(id, &specimenOutcome{"REDACTED"}),
		kinds.NewRPCSuccessReply(id, &specimenOutcome{"REDACTED"}))
	require.NoError(t, err)
	reply = w.Result()
	content, err = io.ReadAll(reply.Body)
	_ = reply.Body.Close()
	require.NoError(t, err)

	assert.Equal(t, 200, reply.StatusCode)
	assert.Equal(t, "REDACTED", reply.Header.Get("REDACTED"))
	assert.Equal(t, "REDACTED", string(content))
}

func VerifyRecordRPCReplyHTTPFault(t *testing.T) {
	w := httptest.NewRecorder()
	err := RecordRPCReplyHTTPFault(
		w,
		http.StatusInternalServerError,
		kinds.RPCIntrinsicFault(kinds.JsonrpcIntegerUID(-1), errors.New("REDACTED")))
	require.NoError(t, err)
	reply := w.Result()
	content, err := io.ReadAll(reply.Body)
	_ = reply.Body.Close()
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, reply.StatusCode)
	assert.Equal(t, "REDACTED", reply.Header.Get("REDACTED"))
	assert.Equal(t, "REDACTED", string(content))
}
