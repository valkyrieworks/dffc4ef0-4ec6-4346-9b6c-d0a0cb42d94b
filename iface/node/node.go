/**
.

:
r
r
*/
package node

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

//
//
func FreshDaemon(schemaLocation, carrier string, app kinds.Platform) (facility.Facility, error) {
	var s facility.Facility
	var err error
	switch carrier {
	case "REDACTED":
		s = FreshPortDaemon(schemaLocation, app)
	case "REDACTED":
		s = FreshGRPSDaemon(schemaLocation, app)
	default:
		err = FaultUnfamiliarDaemonKind{DaemonKind: carrier}
	}
	return s, err
}
