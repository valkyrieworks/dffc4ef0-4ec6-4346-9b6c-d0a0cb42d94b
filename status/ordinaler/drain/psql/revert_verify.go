package psql

import (
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/status/transordinal"
)

var (
	_ ordinaler.LedgerOrdinaler = RevertLedgerOrdinaler{}
	_ transordinal.TransOrdinaler    = RevertTransferOrdinaler{}
)
