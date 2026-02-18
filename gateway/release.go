package gateway

import (
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/release"
)

//
//
//
var QueryDetails = &iface.QueryDetails{
	Release:      release.TMCoreSemaphoreRev,
	LedgerRelease: release.LedgerProtocol,
	P2PRelease:   release.P2PProtocol,
	IfaceRelease:  release.IfaceRelease,
}
