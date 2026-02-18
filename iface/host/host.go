/**
.

:
r
r
*/
package host

import (
	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/daemon"
)

//
//
func NewHost(schemaAddress, carrier string, app kinds.Software) (daemon.Daemon, error) {
	var s daemon.Daemon
	var err error
	switch carrier {
	case "REDACTED":
		s = NewSocketHost(schemaAddress, app)
	case "REDACTED":
		s = NewGRPCHost(schemaAddress, app)
	default:
		err = ErrUnclearHostKind{HostKind: carrier}
	}
	return s, err
}
