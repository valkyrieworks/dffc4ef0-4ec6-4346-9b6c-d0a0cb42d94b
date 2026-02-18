package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/valkyrieworks/utils/log"
	cometos "github.com/valkyrieworks/utils/os"
	rpchost "github.com/valkyrieworks/rpc/jsonrpc/host"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

var paths = map[string]*rpchost.RPCFunction{
	"REDACTED": rpchost.NewRPCFunction(GreetingDomain, "REDACTED"),
}

func GreetingDomain(_ *rpctypes.Context, label string, num int) (Outcome, error) {
	return Outcome{fmt.Sprintf("REDACTED", label, num)}, nil
}

type Outcome struct {
	Outcome string
}

func main() {
	var (
		mux    = http.NewServeMux()
		tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))
	)

	//
	cometos.InterceptAlert(tracer, func() {})

	rpchost.EnrollRPCRoutines(mux, paths, tracer)
	settings := rpchost.StandardSettings()
	observer, err := rpchost.Observe("REDACTED", settings.MaximumAccessLinks)
	if err != nil {
		cometos.Quit(err.Error())
	}

	if err = rpchost.Attend(observer, mux, tracer, settings); err != nil {
		cometos.Quit(err.Error())
	}
}
