package delegate

import (
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//
//
//
var SolicitDetails = &iface.SolicitDetails{
	Edition:      edition.TEMPBaseSemaphoreEdtn,
	LedgerEdition: edition.LedgerScheme,
	Peer2peerEdition:   edition.Peer2peerScheme,
	IfaceEdition:  edition.IfaceEdition,
}
