//

package verifies

import (
	"testing"

	abcinode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	txpool "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
)

func RandomizeTxpool(f *testing.F) {
	app := statedepot.FreshInsideRamPlatform()
	mtx := new(commitchronize.Exclusion)
	link := abcinode.FreshRegionalCustomer(mtx, app)
	err := link.Initiate()
	if err != nil {
		panic(err)
	}

	cfg := settings.FallbackTxpoolSettings()
	cfg.Multicast = false

	mp := txpool.FreshCNCatalogTxpool(cfg, link, 0)

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.InspectTransfer(data, nil, txpool.TransferDetails{})
	})
}
