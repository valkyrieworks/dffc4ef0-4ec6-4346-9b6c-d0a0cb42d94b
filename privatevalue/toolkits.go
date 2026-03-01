package privatevalue

import (
	"errors"
	"fmt"
	"net"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
)

//
//
//
func EqualsLinkDeadline(err error) bool {
	_, ok := errors.Unwrap(err).(deadlineFailure)
	switch {
	case errors.As(err, &GatewayDeadlineFailure{}):
		return true
	case ok:
		return true
	default:
		return false
	}
}

//
func FreshEndorserObserver(overhearLocation string, tracer log.Tracer) (*EndorserObserverGateway, error) {
	var observer net.Listener

	scheme, location := strongmindnet.SchemeAlsoLocation(overhearLocation)
	ln, err := net.Listen(scheme, location)
	if err != nil {
		return nil, err
	}
	switch scheme {
	case "REDACTED":
		observer = FreshPosixObserver(ln)
	case "REDACTED":
		//
		observer = FreshTcpsocketObserver(ln, edwards25519.ProducePrivateToken())
	default:
		return nil, fmt.Errorf(
			"REDACTED",
			scheme,
		)
	}

	pve := FreshEndorserObserverGateway(tracer.Using("REDACTED", "REDACTED"), observer)

	return pve, nil
}

//
func ObtainReleaseLocalmachineLocationChannel() string {
	channel, err := strongmindnet.ObtainLiberateChannel()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("REDACTED", channel)
}
