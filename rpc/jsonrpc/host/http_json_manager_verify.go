package host

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

func verifyMultiplexer() *http.ServeMux {
	functionIndex := map[string]*RPCFunction{
		"REDACTED":     NewRPCFunction(func(ctx *kinds.Context, s string, i int) (string, error) { return "REDACTED", nil }, "REDACTED"),
		"REDACTED": NewRPCFunction(func(ctx *kinds.Context, h int) (string, error) { return "REDACTED", nil }, "REDACTED", Storable("REDACTED")),
	}
	mux := http.NewServeMux()
	buf := new(bytes.Buffer)
	tracer := log.NewTMTracer(buf)
	EnrollRPCRoutines(mux, functionIndex, tracer)

	return mux
}

func stateOK(code int) bool { return code >= 200 && code <= 299 }

//
//
//
func VerifyRPCOptions(t *testing.T) {
	mux := verifyMultiplexer()
	verifies := []struct {
		shipment    string
		desireErr    string
		anticipatedUID any
	}{
		//
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},
		//
		{"REDACTED", "REDACTED", nil},
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},

		//
		//

		//
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonrpcStringUID("REDACTED")},
	}

	for i, tt := range verifies {
		req, _ := http.NewRequest("REDACTED", "REDACTED", strings.NewReader(tt.shipment))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		//
		assert.NotZero(t, res.StatusCode, "REDACTED", i)
		binary, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("REDACTED", i, err)
			continue
		}

		receive := new(kinds.RPCAnswer)
		assert.Nil(t, json.Unmarshal(binary, receive), "REDACTED", i, binary)
		assert.NotEqual(t, receive, new(kinds.RPCAnswer), "REDACTED", i)
		assert.Equal(t, tt.anticipatedUID, receive.ID, "REDACTED", i)
		if tt.desireErr == "REDACTED" {
			assert.Nil(t, receive.Fault, "REDACTED", i)
		} else {
			assert.True(t, receive.Fault.Code < 0, "REDACTED", i)
			//
			assert.Contains(t, receive.Fault.Signal+receive.Fault.Data, tt.desireErr, "REDACTED", i)
		}
	}
}

func VerifyJsonrpcuid(t *testing.T) {
	mux := verifyMultiplexer()
	verifies := []struct {
		shipment    string
		desireErr    bool
		anticipatedUID any
	}{
		//
		{"REDACTED", false, kinds.JsonrpcStringUID("REDACTED")},
		{"REDACTED", false, kinds.JsonrpcStringUID("REDACTED")},
		{"REDACTED", false, kinds.JsonrpcIntegerUID(0)},
		{"REDACTED", false, kinds.JsonrpcIntegerUID(1)},
		{"REDACTED", false, kinds.JsonrpcIntegerUID(1)},
		{"REDACTED", false, kinds.JsonrpcIntegerUID(-1)},

		//
		{"REDACTED", true, nil},
		{"REDACTED", true, nil},
	}

	for i, tt := range verifies {
		req, _ := http.NewRequest("REDACTED", "REDACTED", strings.NewReader(tt.shipment))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		res := rec.Result()
		//
		assert.NotZero(t, res.StatusCode, "REDACTED", i)
		binary, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("REDACTED", i, err)
			continue
		}
		res.Body.Close()

		receive := new(kinds.RPCAnswer)
		err = json.Unmarshal(binary, receive)
		assert.Nil(t, err, "REDACTED", i, binary)
		if !tt.desireErr {
			assert.NotEqual(t, receive, new(kinds.RPCAnswer), "REDACTED", i)
			assert.Equal(t, tt.anticipatedUID, receive.ID, "REDACTED", i)
			assert.Nil(t, receive.Fault, "REDACTED", i)
		} else {
			assert.True(t, receive.Fault.Code < 0, "REDACTED", i)
		}
	}
}

func VerifyRPCBulletin(t *testing.T) {
	mux := verifyMultiplexer()
	content := strings.NewReader("REDACTED")
	req, _ := http.NewRequest("REDACTED", "REDACTED", content)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res := rec.Result()

	//
	require.True(t, stateOK(res.StatusCode), "REDACTED")
	binary, err := io.ReadAll(res.Body)
	res.Body.Close()
	require.Nil(t, err, "REDACTED")
	require.Equal(t, len(binary), 0, "REDACTED")
}

func VerifyRPCBulletinInGroup(t *testing.T) {
	mux := verifyMultiplexer()
	verifies := []struct {
		shipment     string
		anticipateNumber int
	}{
		{
			`REDACTED[
REDACTED,
REDACTED}
REDACTED`,
			1,
		},
		{
			`REDACTED[
REDACTED,
REDACTED,
REDACTED,
REDACTED}
REDACTED`,
			2,
		},
	}
	for i, tt := range verifies {
		req, _ := http.NewRequest("REDACTED", "REDACTED", strings.NewReader(tt.shipment))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		res := rec.Result()
		//
		assert.True(t, stateOK(res.StatusCode), "REDACTED", i)
		binary, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("REDACTED", i, err)
			continue
		}
		res.Body.Close()

		var replies []kinds.RPCAnswer
		//
		err = json.Unmarshal(binary, &replies)
		if err != nil {
			//
			if tt.anticipateNumber > 1 {
				t.Errorf("REDACTED", i, binary)
				continue
			}
			//
			var reply kinds.RPCAnswer
			err = json.Unmarshal(binary, &reply)
			if err != nil {
				t.Errorf("REDACTED", i, binary)
				continue
			}
			//
			replies = []kinds.RPCAnswer{reply}
		}
		if tt.anticipateNumber != len(replies) {
			t.Errorf("REDACTED", i, tt.anticipateNumber, len(replies), binary)
			continue
		}
		for _, reply := range replies {
			assert.NotEqual(t, reply, new(kinds.RPCAnswer), "REDACTED", i)
		}
	}
}

func VerifyUnclearRPCRoute(t *testing.T) {
	mux := verifyMultiplexer()
	req, _ := http.NewRequest("REDACTED", "REDACTED", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res := rec.Result()

	//
	require.Equal(t, http.StatusNotFound, res.StatusCode, "REDACTED")
	res.Body.Close()
}

func VerifyRPCReplyRepository(t *testing.T) {
	mux := verifyMultiplexer()
	content := strings.NewReader("REDACTED")
	req, _ := http.NewRequest("REDACTED", "REDACTED", content)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res := rec.Result()

	//
	require.True(t, stateOK(res.StatusCode), "REDACTED")
	require.Equal(t, "REDACTED", res.Header.Get("REDACTED"))

	_, err := io.ReadAll(res.Body)
	res.Body.Close()
	require.Nil(t, err, "REDACTED")

	//
	content = strings.NewReader("REDACTED")
	req, _ = http.NewRequest("REDACTED", "REDACTED", content)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res = rec.Result()

	//
	require.True(t, stateOK(res.StatusCode), "REDACTED")
	require.Equal(t, "REDACTED", res.Header.Get("REDACTED"))

	_, err = io.ReadAll(res.Body)

	res.Body.Close()
	require.Nil(t, err, "REDACTED")

	//
	content = strings.NewReader("REDACTED")
	req, _ = http.NewRequest("REDACTED", "REDACTED", content)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res = rec.Result()

	//
	require.True(t, stateOK(res.StatusCode), "REDACTED")
	require.Equal(t, "REDACTED", res.Header.Get("REDACTED"))

	_, err = io.ReadAll(res.Body)

	res.Body.Close()
	require.Nil(t, err, "REDACTED")
}
