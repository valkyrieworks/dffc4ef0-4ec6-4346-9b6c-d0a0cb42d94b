package node

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

type specimenOutcome struct {
	Datum string `json:"datum"`
}

func VerifyMaximumUnlockLinkages(t *testing.T) {
	const max = 5 //

	//
	var unlock int32
	mux := http.NewServeMux()
	mux.HandleFunc("REDACTED", func(w http.ResponseWriter, r *http.Request) {
		if n := atomic.AddInt32(&unlock, 1); n > int32(max) {
			t.Errorf("REDACTED", n, max)
		}
		defer atomic.AddInt32(&unlock, -1)
		time.Sleep(10 * time.Millisecond)
		fmt.Fprint(w, "REDACTED")
	})
	settings := FallbackSettings()
	l, err := Overhear("REDACTED", max)
	require.NoError(t, err)
	defer l.Close()
	go Attend(l, mux, log.VerifyingTracer(), settings) //

	//
	endeavors := max * 2
	var wg sync.WaitGroup
	var unsuccessful int32
	for i := 0; i < endeavors; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := http.Client{Timeout: 3 * time.Second}
			r, err := c.Get("REDACTED" + l.Addr().String())
			if err != nil {
				atomic.AddInt32(&unsuccessful, 1)
				return
			}
			defer r.Body.Close()
		}()
	}
	wg.Wait()

	//
	//
	if int(unsuccessful) >= endeavors/2 {
		t.Errorf("REDACTED", unsuccessful, endeavors)
	}
}

func VerifyAttendTransportsec(t *testing.T) {
	ln, err := net.Listen("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer ln.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("REDACTED", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "REDACTED")
	})

	chnlFault := make(chan error, 1)
	go func() {
		//
		chnlFault <- AttendTransportsec(ln, mux, "REDACTED", "REDACTED", log.VerifyingTracer(), FallbackSettings())
	}()

	select {
	case err := <-chnlFault:
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

func VerifyRecordRemoteReplyHttpsvc(t *testing.T) {
	id := kinds.JsonifaceIntegerUUID(-1)

	//
	w := httptest.NewRecorder()
	err := RecordStorableRemoteReplyHttpsvc(w, kinds.FreshRemoteTriumphReply(id, &specimenOutcome{"REDACTED"}))
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
	err = RecordRemoteReplyHttpsvc(w,
		kinds.FreshRemoteTriumphReply(id, &specimenOutcome{"REDACTED"}),
		kinds.FreshRemoteTriumphReply(id, &specimenOutcome{"REDACTED"}))
	require.NoError(t, err)
	reply = w.Result()
	content, err = io.ReadAll(reply.Body)
	_ = reply.Body.Close()
	require.NoError(t, err)

	assert.Equal(t, 200, reply.StatusCode)
	assert.Equal(t, "REDACTED", reply.Header.Get("REDACTED"))
	assert.Equal(t, "REDACTED", string(content))
}

func VerifyRecordRemoteReplyHttpsvcFailure(t *testing.T) {
	w := httptest.NewRecorder()
	err := RecordRemoteReplyHttpsvcFailure(
		w,
		http.StatusInternalServerError,
		kinds.RemoteIntrinsicFailure(kinds.JsonifaceIntegerUUID(-1), errors.New("REDACTED")))
	require.NoError(t, err)
	reply := w.Result()
	content, err := io.ReadAll(reply.Body)
	_ = reply.Body.Close()
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, reply.StatusCode)
	assert.Equal(t, "REDACTED", reply.Header.Get("REDACTED"))
	assert.Equal(t, "REDACTED", string(content))
}
