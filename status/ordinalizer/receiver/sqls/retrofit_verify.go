package sqls

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
)

var (
	_ ordinalizer.LedgerOrdinalizer = RetrofitLedgerOrdinalizer{}
	_ transferordinal.TransferOrdinalizer    = RetrofitTransferOrdinalizer{}
)
