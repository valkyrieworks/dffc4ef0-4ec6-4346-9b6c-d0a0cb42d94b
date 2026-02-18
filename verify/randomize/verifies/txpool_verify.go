//

package verifies

import (
	"testing"

	ifacecustomer "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/settings"
	engineconnect "github.com/valkyrieworks/utils/align"
	txpool "github.com/valkyrieworks/txpool"
)

func RandomizeTxpool(f *testing.F) {
	app := objectdepot.NewInRamSoftware()
	mtx := new(engineconnect.Lock)
	link := ifacecustomer.NewNativeCustomer(mtx, app)
	err := link.Begin()
	if err != nil {
		panic(err)
	}

	cfg := settings.StandardTxpoolSettings()
	cfg.Multicast = false

	mp := txpool.NewCCatalogTxpool(cfg, link, 0)

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.InspectTransfer(data, nil, txpool.TransferDetails{})
	})
}
