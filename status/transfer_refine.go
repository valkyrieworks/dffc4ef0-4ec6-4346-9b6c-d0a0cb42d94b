package status

import (
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
func TransferPriorInspect(status Status) txpooll.PriorInspectMethod {
	maximumOctets := status.AgreementSettings.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(kinds.MaximumLedgerExtentOctets)
	}
	maximumDataOctets := kinds.MaximumDataOctetsNegativeProof(
		maximumOctets,
		status.Assessors.Extent(),
	)
	return txpooll.PriorInspectMaximumOctets(maximumDataOctets)
}

//
//
func TransferRelayInspect(status Status) txpooll.RelayInspectMethod {
	return txpooll.SubmitInspectMaximumFuel(status.AgreementSettings.Ledger.MaximumFuel)
}
