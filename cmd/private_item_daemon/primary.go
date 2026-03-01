package primary

import (
	"flag"
	"os"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
)

func primary() {
	var (
		location             = flag.String("REDACTED", "REDACTED", "REDACTED")
		successionUUID          = flag.String("REDACTED", "REDACTED", "REDACTED")
		privateItemTokenRoute   = flag.String("REDACTED", "REDACTED", "REDACTED")
		privateItemStatusRoute = flag.String("REDACTED", "REDACTED", "REDACTED")

		tracer = log.FreshTEMPTracer(
			log.FreshChronizePersistor(os.Stdout),
		).Using("REDACTED", "REDACTED")
	)
	flag.Parse()

	tracer.Details(
		"REDACTED",
		"REDACTED", *location,
		"REDACTED", *successionUUID,
		"REDACTED", *privateItemTokenRoute,
		"REDACTED", *privateItemStatusRoute,
	)

	pv := privatevalue.FetchRecordPRV(*privateItemTokenRoute, *privateItemStatusRoute)

	var caller privatevalue.PortCaller
	scheme, location := strongmindnet.SchemeAlsoLocation(*location)
	switch scheme {
	case "REDACTED":
		caller = privatevalue.CallPosixProc(location)
	case "REDACTED":
		linkDeadline := 3 * time.Second //
		caller = privatevalue.CallStreamProc(location, linkDeadline, edwards25519.ProducePrivateToken())
	default:
		tracer.Failure("REDACTED", "REDACTED", scheme)
		os.Exit(1)
	}

	sd := privatevalue.FreshEndorserCallerGateway(tracer, caller)
	ss := privatevalue.FreshEndorserDaemon(sd, *successionUUID, pv)

	err := ss.Initiate()
	if err != nil {
		panic(err)
	}

	//
	strongos.EnsnareGesture(tracer, func() {
		err := ss.Halt()
		if err != nil {
			panic(err)
		}
	})

	//
	select {}
}
