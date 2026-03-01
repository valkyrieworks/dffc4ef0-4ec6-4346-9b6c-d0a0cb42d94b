package handler

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
)

var txpool txpooll.Txpool

func initialize() {
	app := statedepot.FreshInsideRamPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	applicationLinkMemory, _ := cc.FreshIfaceCustomer()
	err := applicationLinkMemory.Initiate()
	if err != nil {
		panic(err)
	}

	cfg := settings.FallbackTxpoolSettings()
	cfg.Multicast = false
	txpool = txpooll.FreshCNCatalogTxpool(cfg, applicationLinkMemory, 0)
}

func Randomize(data []byte) int {
	err := txpool.InspectTransfer(data, nil, txpooll.TransferDetails{})
	if err != nil {
		return 0
	}

	return 1
}
