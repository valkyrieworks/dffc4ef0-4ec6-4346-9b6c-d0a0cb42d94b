package ledger

import (
	"errors"
	"fmt"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	ledgeridxkv "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger/kv"
	ledgerpositionnull "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger/nothing"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/receiver/sqls"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/kv"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/nothing"
)

//
//
func OrdinalizerOriginatingSettings(cfg *settings.Settings, datastoreSupplier settings.DatastoreSupplier, successionUUID string) (
	transferOffset transferordinal.TransferOrdinalizer, ledgerOffset ordinalizer.LedgerOrdinalizer, err error,
) {
	transoffset, ldgoffset, _, err := OrdinalizerOriginatingSettingsUsingDeactivatedOrdinalizers(cfg, datastoreSupplier, successionUUID)
	return transoffset, ldgoffset, err
}

//
//
//
func OrdinalizerOriginatingSettingsUsingDeactivatedOrdinalizers(cfg *settings.Settings, datastoreSupplier settings.DatastoreSupplier, successionUUID string) (
	transferOffset transferordinal.TransferOrdinalizer, ledgerOffset ordinalizer.LedgerOrdinalizer, everyOrdinalizersDeactivated bool, err error,
) {
	switch cfg.TransferOrdinal.Ordinalizer {
	case "REDACTED":
		depot, err := datastoreSupplier(&settings.DatastoreScope{ID: "REDACTED", Settings: cfg})
		if err != nil {
			return nil, nil, false, err
		}

		return kv.FreshTransferOrdinal(depot), ledgeridxkv.New(dbm.FreshHeadingDatastore(depot, []byte("REDACTED"))), false, nil

	case "REDACTED":
		link := cfg.TransferOrdinal.SqlsLink
		if link == "REDACTED" {
			return nil, nil, false, errors.New("REDACTED")
		}
		es, err := sqls.FreshIncidentReceiver(cfg.TransferOrdinal.SqlsLink, successionUUID)
		if err != nil {
			return nil, nil, false, fmt.Errorf("REDACTED", err)
		}
		return es.TransferOrdinalizer(), es.LedgerOrdinalizer(), false, nil

	default:
		return &nothing.TransferOrdinal{}, &ledgerpositionnull.PreventerOrdinalizer{}, true, nil
	}
}
