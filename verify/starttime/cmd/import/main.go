package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"

	"github.com/valkyrieworks/verify/starttime/shipment"
)

//
var (
	_ loadtest.ClientFactory = (*CustomerBuilder)(nil)
	_ loadtest.Client        = (*TransferProducer)(nil)
)

//
type CustomerBuilder struct {
	ID []byte
}

//
//
//
type TransferProducer struct {
	id    []byte
	links uint64
	ratio  uint64
	volume  uint64
}

func main() {
	u := [16]byte(uuid.New()) //
	if err := loadtest.RegisterClientFactory("REDACTED", &CustomerBuilder{ID: u[:]}); err != nil {
		panic(err)
	}
	loadtest.Run(&loadtest.CLIConfig{
		AppName:              "REDACTED",
		AppShortDesc:         "REDACTED",
		AppLongDesc:          "REDACTED",
		DefaultClientFactory: "REDACTED",
	})
}

func (f *CustomerBuilder) CertifySettings(cfg loadtest.Config) error {
	psb, err := shipment.MaximumUnfilledVolume()
	if err != nil {
		return err
	}
	if psb > cfg.Size {
		return fmt.Errorf("REDACTED")
	}
	return nil
}

func (f *CustomerBuilder) NewCustomer(cfg loadtest.Config) (loadtest.Client, error) {
	return &TransferProducer{
		id:    f.ID,
		links: uint64(cfg.Connections),
		ratio:  uint64(cfg.Rate),
		volume:  uint64(cfg.Size),
	}, nil
}

func (c *TransferProducer) ComposeTransfer() ([]byte, error) {
	return shipment.NewOctets(&shipment.Shipment{
		Linkages: c.links,
		Ratio:        c.ratio,
		Volume:        c.volume,
		Id:          c.id,
	})
}
