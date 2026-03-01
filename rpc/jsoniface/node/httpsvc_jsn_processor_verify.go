package node

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

func verifyMultiplexer() *http.ServeMux {
	methodIndex := map[string]*RemoteMethod{
		"REDACTED":     FreshRemoteMethod(func(ctx *kinds.Env, s string, i int) (string, error) { return "REDACTED", nil }, "REDACTED"),
		"REDACTED": FreshRemoteMethod(func(ctx *kinds.Env, h int) (string, error) { return "REDACTED", nil }, "REDACTED", Storable("REDACTED")),
	}
	mux := http.NewServeMux()
	buf := new(bytes.Buffer)
	tracer := log.FreshTEMPTracer(buf)
	EnrollRemoteRoutines(mux, methodIndex, tracer)

	return mux
}

func conditionOKAY(cipher int) bool { return cipher >= 200 && cipher <= 299 }

//
//
//
func VerifyRemoteParameters(t *testing.T) {
	mux := verifyMultiplexer()
	verifies := []struct {
		content    string
		desireFault    string
		anticipatedUUID any
	}{
		//
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},
		//
		{"REDACTED", "REDACTED", nil},
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},

		//
		//

		//
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},
		{"REDACTED", "REDACTED", kinds.JsonifaceTextUUID("REDACTED")},
	}

	for i, tt := range verifies {
		req, _ := http.NewRequest("REDACTED", "REDACTED", strings.NewReader(tt.content))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		//
		assert.NotZero(t, res.StatusCode, "REDACTED", i)
		chunk, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("REDACTED", i, err)
			continue
		}

		obtain := new(kinds.RemoteReply)
		assert.Nil(t, json.Unmarshal(chunk, obtain), "REDACTED", i, chunk)
		assert.NotEqual(t, obtain, new(kinds.RemoteReply), "REDACTED", i)
		assert.Equal(t, tt.anticipatedUUID, obtain.ID, "REDACTED", i)
		if tt.desireFault == "REDACTED" {
			assert.Nil(t, obtain.Failure, "REDACTED", i)
		} else {
			assert.True(t, obtain.Failure.Cipher < 0, "REDACTED", i)
			//
			assert.Contains(t, obtain.Failure.Signal+obtain.Failure.Data, tt.desireFault, "REDACTED", i)
		}
	}
}

func VerifyJsonrpcuuid(t *testing.T) {
	mux := verifyMultiplexer()
	verifies := []struct {
		content    string
		desireFault    bool
		anticipatedUUID any
	}{
		//
		{"REDACTED", false, kinds.JsonifaceTextUUID("REDACTED")},
		{"REDACTED", false, kinds.JsonifaceTextUUID("REDACTED")},
		{"REDACTED", false, kinds.JsonifaceIntegerUUID(0)},
		{"REDACTED", false, kinds.JsonifaceIntegerUUID(1)},
		{"REDACTED", false, kinds.JsonifaceIntegerUUID(1)},
		{"REDACTED", false, kinds.JsonifaceIntegerUUID(-1)},

		//
		{"REDACTED", true, nil},
		{"REDACTED", true, nil},
	}

	for i, tt := range verifies {
		req, _ := http.NewRequest("REDACTED", "REDACTED", strings.NewReader(tt.content))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		res := rec.Result()
		//
		assert.NotZero(t, res.StatusCode, "REDACTED", i)
		chunk, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("REDACTED", i, err)
			continue
		}
		res.Body.Close()

		obtain := new(kinds.RemoteReply)
		err = json.Unmarshal(chunk, obtain)
		assert.Nil(t, err, "REDACTED", i, chunk)
		if !tt.desireFault {
			assert.NotEqual(t, obtain, new(kinds.RemoteReply), "REDACTED", i)
			assert.Equal(t, tt.anticipatedUUID, obtain.ID, "REDACTED", i)
			assert.Nil(t, obtain.Failure, "REDACTED", i)
		} else {
			assert.True(t, obtain.Failure.Cipher < 0, "REDACTED", i)
		}
	}
}

func VerifyRemoteBulletin(t *testing.T) {
	mux := verifyMultiplexer()
	content := strings.NewReader("REDACTED")
	req, _ := http.NewRequest("REDACTED", "REDACTED", content)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res := rec.Result()

	//
	require.True(t, conditionOKAY(res.StatusCode), "REDACTED")
	chunk, err := io.ReadAll(res.Body)
	res.Body.Close()
	require.Nil(t, err, "REDACTED")
	require.Equal(t, len(chunk), 0, "REDACTED")
}

func VerifyRemoteBulletinInsideCluster(t *testing.T) {
	mux := verifyMultiplexer()
	verifies := []struct {
		content     string
		anticipateTally int
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
		req, _ := http.NewRequest("REDACTED", "REDACTED", strings.NewReader(tt.content))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		res := rec.Result()
		//
		assert.True(t, conditionOKAY(res.StatusCode), "REDACTED", i)
		chunk, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("REDACTED", i, err)
			continue
		}
		res.Body.Close()

		var replies []kinds.RemoteReply
		//
		err = json.Unmarshal(chunk, &replies)
		if err != nil {
			//
			if tt.anticipateTally > 1 {
				t.Errorf("REDACTED", i, chunk)
				continue
			}
			//
			var reply kinds.RemoteReply
			err = json.Unmarshal(chunk, &reply)
			if err != nil {
				t.Errorf("REDACTED", i, chunk)
				continue
			}
			//
			replies = []kinds.RemoteReply{reply}
		}
		if tt.anticipateTally != len(replies) {
			t.Errorf("REDACTED", i, tt.anticipateTally, len(replies), chunk)
			continue
		}
		for _, reply := range replies {
			assert.NotEqual(t, reply, new(kinds.RemoteReply), "REDACTED", i)
		}
	}
}

func VerifyUnfamiliarRemoteRoute(t *testing.T) {
	mux := verifyMultiplexer()
	req, _ := http.NewRequest("REDACTED", "REDACTED", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res := rec.Result()

	//
	require.Equal(t, http.StatusNotFound, res.StatusCode, "REDACTED")
	res.Body.Close()
}

func VerifyRemoteReplyStash(t *testing.T) {
	mux := verifyMultiplexer()
	content := strings.NewReader("REDACTED")
	req, _ := http.NewRequest("REDACTED", "REDACTED", content)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	res := rec.Result()

	//
	require.True(t, conditionOKAY(res.StatusCode), "REDACTED")
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
	require.True(t, conditionOKAY(res.StatusCode), "REDACTED")
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
	require.True(t, conditionOKAY(res.StatusCode), "REDACTED")
	require.Equal(t, "REDACTED", res.Header.Get("REDACTED"))

	_, err = io.ReadAll(res.Body)

	res.Body.Close()
	require.Nil(t, err, "REDACTED")
}
