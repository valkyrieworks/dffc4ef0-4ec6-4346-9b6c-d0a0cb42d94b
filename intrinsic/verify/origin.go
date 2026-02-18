package verify

import (
	"time"

	"github.com/valkyrieworks/kinds"
)

func OriginPaper(
	moment time.Time,
	ratifiers []*kinds.Ratifier,
	agreementOptions *kinds.AgreementOptions,
	ledgerUID string,
) *kinds.OriginPaper {
	originRatifiers := make([]kinds.OriginRatifier, len(ratifiers))

	for i := range ratifiers {
		originRatifiers[i] = kinds.OriginRatifier{
			Energy:  ratifiers[i].PollingEnergy,
			PublicKey: ratifiers[i].PublicKey,
		}
	}

	return &kinds.OriginPaper{
		OriginMoment:     moment,
		PrimaryLevel:   1,
		LedgerUID:         ledgerUID,
		Ratifiers:      originRatifiers,
		AgreementOptions: agreementOptions,
	}
}
