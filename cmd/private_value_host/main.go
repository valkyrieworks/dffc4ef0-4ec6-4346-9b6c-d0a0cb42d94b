package main

import (
	"flag"
	"os"
	"time"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	cometnet "github.com/valkyrieworks/utils/net"
	cometos "github.com/valkyrieworks/utils/os"

	"github.com/valkyrieworks/privatekey"
)

func main() {
	var (
		address             = flag.String("REDACTED", "REDACTED", "REDACTED")
		ledgerUID          = flag.String("REDACTED", "REDACTED", "REDACTED")
		privateValueKeyRoute   = flag.String("REDACTED", "REDACTED", "REDACTED")
		privateValueStatusRoute = flag.String("REDACTED", "REDACTED", "REDACTED")

		tracer = log.NewTMTracer(
			log.NewAlignRecorder(os.Stdout),
		).With("REDACTED", "REDACTED")
	)
	flag.Parse()

	tracer.Details(
		"REDACTED",
		"REDACTED", *address,
		"REDACTED", *ledgerUID,
		"REDACTED", *privateValueKeyRoute,
		"REDACTED", *privateValueStatusRoute,
	)

	pv := privatekey.ImportEntryPrivatekey(*privateValueKeyRoute, *privateValueStatusRoute)

	var caller privatekey.SocketCaller
	protocol, location := cometnet.ProtocolAndLocation(*address)
	switch protocol {
	case "REDACTED":
		caller = privatekey.CallUnixFn(location)
	case "REDACTED":
		linkDeadline := 3 * time.Second //
		caller = privatekey.CallTCPFn(location, linkDeadline, ed25519.GeneratePrivateKey())
	default:
		tracer.Fault("REDACTED", "REDACTED", protocol)
		os.Exit(1)
	}

	sd := privatekey.NewNotaryCallerGateway(tracer, caller)
	ss := privatekey.NewNotaryHost(sd, *ledgerUID, pv)

	err := ss.Begin()
	if err != nil {
		panic(err)
	}

	//
	cometos.InterceptAlert(tracer, func() {
		err := ss.Halt()
		if err != nil {
			panic(err)
		}
	})

	//
	select {}
}
