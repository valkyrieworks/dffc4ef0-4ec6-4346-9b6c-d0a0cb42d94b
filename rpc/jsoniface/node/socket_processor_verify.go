package node

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

func VerifyWebterminalAdministratorProcessor(t *testing.T) {
	s := freshSocketDaemon()
	defer s.Close()

	//
	d := websocket.Dialer{}
	c, callAnswer, err := d.Dial("REDACTED"+s.Listener.Addr().String()+"REDACTED", nil)
	require.NoError(t, err)

	if got, desire := callAnswer.StatusCode, http.StatusSwitchingProtocols; got != desire {
		t.Errorf("REDACTED", got, desire)
	}

	//
	req, err := kinds.IndexTowardSolicit(
		kinds.JsonifaceTextUUID("REDACTED"),
		"REDACTED",
		map[string]any{"REDACTED": "REDACTED", "REDACTED": 10},
	)
	require.NoError(t, err)
	err = c.WriteJSON(req)
	require.NoError(t, err)

	var reply kinds.RemoteReply
	err = c.ReadJSON(&reply)
	require.NoError(t, err)
	require.Nil(t, reply.Failure)
	callAnswer.Body.Close()
}

func freshSocketDaemon() *httptest.Server {
	methodIndex := map[string]*RemoteMethod{
		"REDACTED": FreshSocketifaceMethod(func(ctx *kinds.Env, s string, i int) (string, error) { return "REDACTED", nil }, "REDACTED"),
	}
	wm := FreshWebterminalAdministrator(methodIndex)
	wm.AssignTracer(log.VerifyingTracer())

	mux := http.NewServeMux()
	mux.HandleFunc("REDACTED", wm.WebterminalProcessor)

	return httptest.NewServer(mux)
}
