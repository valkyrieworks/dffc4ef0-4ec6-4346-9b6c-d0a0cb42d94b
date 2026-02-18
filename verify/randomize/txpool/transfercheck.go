package handler

import (
	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/settings"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/gateway"
)

var txpool txpool.Txpool

func init() {
	app := objectdepot.NewInRamSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	applicationLinkMemory, _ := cc.NewIfaceCustomer()
	err := applicationLinkMemory.Begin()
	if err != nil {
		panic(err)
	}

	cfg := settings.StandardTxpoolSettings()
	cfg.Multicast = false
	txpool = txpool.NewCCatalogTxpool(cfg, applicationLinkMemory, 0)
}

func Randomize(data []byte) int {
	err := txpool.InspectTransfer(data, nil, txpool.TransferDetails{})
	if err != nil {
		return 0
	}

	return 1
}
