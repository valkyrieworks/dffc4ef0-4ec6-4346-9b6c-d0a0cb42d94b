package ledger

import (
	"errors"
	"fmt"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/status/ordinaler"
	ledgerordinalkv "github.com/valkyrieworks/status/ordinaler/ledger/kv"
	blockordinalvoid "github.com/valkyrieworks/status/ordinaler/ledger/void"
	"github.com/valkyrieworks/status/ordinaler/drain/psql"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/status/transordinal/kv"
	"github.com/valkyrieworks/status/transordinal/void"
)

//
//
func OrdinalerFromSettings(cfg *settings.Settings, storeSource settings.StoreSource, ledgerUID string) (
	transferOrdinal transordinal.TransOrdinaler, ledgerOrdinal ordinaler.LedgerOrdinaler, err error,
) {
	transferidx, ledgeridx, _, err := OrdinalerFromSettingsWithDeactivatedOrdinalers(cfg, storeSource, ledgerUID)
	return transferidx, ledgeridx, err
}

//
//
//
func OrdinalerFromSettingsWithDeactivatedOrdinalers(cfg *settings.Settings, storeSource settings.StoreSource, ledgerUID string) (
	transferOrdinal transordinal.TransOrdinaler, ledgerOrdinal ordinaler.LedgerOrdinaler, allOrdinalersDeactivated bool, err error,
) {
	switch cfg.TransOrdinal.Ordinaler {
	case "REDACTED":
		depot, err := storeSource(&settings.StoreContext{ID: "REDACTED", Settings: cfg})
		if err != nil {
			return nil, nil, false, err
		}

		return kv.NewTransOrdinal(depot), ledgerordinalkv.New(dbm.NewHeadingStore(depot, []byte("REDACTED"))), false, nil

	case "REDACTED":
		link := cfg.TransOrdinal.PsqlLink
		if link == "REDACTED" {
			return nil, nil, false, errors.New("REDACTED")
		}
		es, err := psql.NewEventDrain(cfg.TransOrdinal.PsqlLink, ledgerUID)
		if err != nil {
			return nil, nil, false, fmt.Errorf("REDACTED", err)
		}
		return es.TransOrdinaler(), es.LedgerOrdinaler(), false, nil

	default:
		return &void.TransOrdinal{}, &blockordinalvoid.ImpedimentOrdinaler{}, true, nil
	}
}
