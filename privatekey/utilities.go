package privatekey

import (
	"errors"
	"fmt"
	"net"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	cometnet "github.com/valkyrieworks/utils/net"
)

//
//
//
func IsLinkDeadline(err error) bool {
	_, ok := errors.Unwrap(err).(deadlineFault)
	switch {
	case errors.As(err, &TerminusDeadlineFault{}):
		return true
	case ok:
		return true
	default:
		return false
	}
}

//
func NewNotaryObserver(acceptAddress string, tracer log.Tracer) (*NotaryObserverTerminus, error) {
	var observer net.Listener

	protocol, location := cometnet.ProtocolAndLocation(acceptAddress)
	ln, err := net.Listen(protocol, location)
	if err != nil {
		return nil, err
	}
	switch protocol {
	case "REDACTED":
		observer = NewUnixObserver(ln)
	case "REDACTED":
		//
		observer = NewTCPObserver(ln, ed25519.GeneratePrivateKey())
	default:
		return nil, fmt.Errorf(
			"REDACTED",
			protocol,
		)
	}

	pve := NewNotaryObserverTerminus(tracer.With("REDACTED", "REDACTED"), observer)

	return pve, nil
}

//
func FetchReleaseLocalhostAddressPort() string {
	port, err := cometnet.FetchReleasePort()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("REDACTED", port)
}
