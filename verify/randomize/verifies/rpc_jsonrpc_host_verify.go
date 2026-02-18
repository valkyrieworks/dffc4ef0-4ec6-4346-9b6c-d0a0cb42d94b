//

package verifies

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/valkyrieworks/utils/log"
	rpchost "github.com/valkyrieworks/rpc/jsonrpc/host"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

func RandomizeRpcjsonrpcHost(f *testing.F) {
	type args struct {
		S string `json:"s"`
		I int    `json:"i"`
	}
	rpcFunctionIndex := map[string]*rpchost.RPCFunction{
		"REDACTED": rpchost.NewRPCFunction(func(ctx *rpctypes.Context, args *args, options ...rpchost.Setting) (string, error) {
			return "REDACTED", nil
		}, "REDACTED"),
	}

	mux := http.NewServeMux()
	rpchost.EnrollRPCRoutines(mux, rpcFunctionIndex, log.NewNoopTracer())
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
		binary, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
		if len(binary) == 0 {
			return
		}

		if resultJSONIsSection(binary) {
			var receive []rpctypes.RPCAnswer
			if err := json.Unmarshal(binary, &receive); err != nil {
				panic(err)
			}
			return
		}
		var receive rpctypes.RPCAnswer
		if err := json.Unmarshal(binary, &receive); err != nil {
			panic(err)
		}
	})
}

func resultJSONIsSection(influx []byte) bool {
	var section []json.RawMessage
	return json.Unmarshal(influx, &section) == nil
}
