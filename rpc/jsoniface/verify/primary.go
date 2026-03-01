package primary

import (
	"fmt"
	"net/http"
	"os"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	rpchandler "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

var paths = map[string]*rpchandler.RemoteMethod{
	"REDACTED": rpchandler.FreshRemoteMethod(GreetingGlobe, "REDACTED"),
}

func GreetingGlobe(_ *remoteifacetypes.Env, alias string, num int) (Outcome, error) {
	return Outcome{fmt.Sprintf("REDACTED", alias, num)}, nil
}

type Outcome struct {
	Outcome string
}

func primary() {
	var (
		mux    = http.NewServeMux()
		tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))
	)

	//
	strongos.EnsnareGesture(tracer, func() {})

	rpchandler.EnrollRemoteRoutines(mux, paths, tracer)
	settings := rpchandler.FallbackSettings()
	observer, err := rpchandler.Overhear("REDACTED", settings.MaximumInitiateLinks)
	if err != nil {
		strongos.Quit(err.Error())
	}

	if err = rpchandler.Attend(observer, mux, tracer, settings); err != nil {
		strongos.Quit(err.Error())
	}
}
