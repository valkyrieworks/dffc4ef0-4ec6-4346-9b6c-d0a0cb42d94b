package primary

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/fetchmoment/content"
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
	frequency  uint64
	extent  uint64
}

func primary() {
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
	psb, err := content.MaximumUnfilledExtent()
	if err != nil {
		return err
	}
	if psb > cfg.Size {
		return fmt.Errorf("REDACTED")
	}
	return nil
}

func (f *CustomerBuilder) FreshCustomer(cfg loadtest.Config) (loadtest.Client, error) {
	return &TransferProducer{
		id:    f.ID,
		links: uint64(cfg.Connections),
		frequency:  uint64(cfg.Rate),
		extent:  uint64(cfg.Size),
	}, nil
}

func (c *TransferProducer) ComposeTransfer() ([]byte, error) {
	return content.FreshOctets(&content.Content{
		Linkages: c.links,
		Frequency:        c.frequency,
		Extent:        c.extent,
		Id:          c.id,
	})
}
