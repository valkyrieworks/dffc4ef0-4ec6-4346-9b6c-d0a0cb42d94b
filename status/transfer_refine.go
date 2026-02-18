package status

import (
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/kinds"
)

//
//
func TransferPreInspect(status Status) txpool.PreInspectFunction {
	maximumOctets := status.AgreementOptions.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(kinds.MaximumLedgerVolumeOctets)
	}
	maximumDataOctets := kinds.MaximumDataOctetsNoProof(
		maximumOctets,
		status.Ratifiers.Volume(),
	)
	return txpool.PreInspectMaximumOctets(maximumDataOctets)
}

//
//
func TransferSubmitInspect(status Status) txpool.SubmitInspectFunction {
	return txpool.SubmitInspectMaximumFuel(status.AgreementOptions.Ledger.MaximumFuel)
}
