//

package verifies

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	rpchandler "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

func RandomizeRemotecalljsnremoteDaemon(f *testing.F) {
	type arguments struct {
		S string `json:"s"`
		I int    `json:"i"`
	}
	remoteMethodIndex := map[string]*rpchandler.RemoteMethod{
		"REDACTED": rpchandler.FreshRemoteMethod(func(ctx *remoteifacetypes.Env, arguments *arguments, choices ...rpchandler.Selection) (string, error) {
			return "REDACTED", nil
		}, "REDACTED"),
	}

	mux := http.NewServeMux()
	rpchandler.EnrollRemoteRoutines(mux, remoteMethodIndex, log.FreshNooperationTracer())
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) == 0 {
			return
		}

		req, err := http.NewRequest("REDACTED", "REDACTED", bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		res := rec.Result()
		chunk, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
		if len(chunk) == 0 {
			return
		}

		if emissionJSNEqualsSection(chunk) {
			var obtain []remoteifacetypes.RemoteReply
			if err := json.Unmarshal(chunk, &obtain); err != nil {
				panic(err)
			}
			return
		}
		var obtain remoteifacetypes.RemoteReply
		if err := json.Unmarshal(chunk, &obtain); err != nil {
			panic(err)
		}
	})
}

func emissionJSNEqualsSection(influx []byte) bool {
	var section []json.RawMessage
	return json.Unmarshal(influx, &section) == nil
}
