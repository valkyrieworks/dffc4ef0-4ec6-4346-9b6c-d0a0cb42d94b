package host

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

func VerifyWebchannelAdministratorManager(t *testing.T) {
	s := newWSHost()
	defer s.Close()

	//
	d := websocket.Dialer{}
	c, callReply, err := d.Dial("REDACTED"+s.Listener.Addr().String()+"REDACTED", nil)
	require.NoError(t, err)

	if got, desire := callReply.StatusCode, http.StatusSwitchingProtocols; got != desire {
		t.Errorf("REDACTED", got, desire)
	}

	//
	req, err := kinds.IndexToQuery(
		kinds.JsonrpcStringUID("REDACTED"),
		"REDACTED",
		map[string]any{"REDACTED": "REDACTED", "REDACTED": 10},
	)
	require.NoError(t, err)
	err = c.WriteJSON(req)
	require.NoError(t, err)

	var reply kinds.RPCAnswer
	err = c.ReadJSON(&reply)
	require.NoError(t, err)
	require.Nil(t, reply.Fault)
	callReply.Body.Close()
}

func newWSHost() *httptest.Server {
	functionIndex := map[string]*RPCFunction{
		"REDACTED": NewWsrpcFunction(func(ctx *kinds.Context, s string, i int) (string, error) { return "REDACTED", nil }, "REDACTED"),
	}
	wm := NewWebchannelOverseer(functionIndex)
	wm.AssignTracer(log.VerifyingTracer())

	mux := http.NewServeMux()
	mux.HandleFunc("REDACTED", wm.WebchannelManager)

	return httptest.NewServer(mux)
}
