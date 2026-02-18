package gateway

import (
	"fmt"

	abciend "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/iface/kinds"
	engineconnect "github.com/valkyrieworks/utils/align"
	e2e "github.com/valkyrieworks/verify/e2e/app"
)

//

//
type CustomerOriginator interface {
	//
	NewIfaceCustomer() (abciend.Customer, error)
}

//
//

type nativeCustomerOriginator struct {
	mtx *engineconnect.Lock
	app kinds.Software
}

//
//
//
//
//
//
func NewNativeCustomerOriginator(app kinds.Software) CustomerOriginator {
	return &nativeCustomerOriginator{
		mtx: new(engineconnect.Lock),
		app: app,
	}
}

func (l *nativeCustomerOriginator) NewIfaceCustomer() (abciend.Customer, error) {
	return abciend.NewNativeCustomer(l.mtx, l.app), nil
}

//
//

type linkAlignNativeCustomerOriginator struct {
	app kinds.Software
}

//
//
//
//
//
//
//
func NewLinkAlignNativeCustomerOriginator(app kinds.Software) CustomerOriginator {
	return &linkAlignNativeCustomerOriginator{
		app: app,
	}
}

func (c *linkAlignNativeCustomerOriginator) NewIfaceCustomer() (abciend.Customer, error) {
	//
	//
	return abciend.NewNativeCustomer(nil, c.app), nil
}

//
//

type distantCustomerOriginator struct {
	address        string
	carrier   string
	shouldLink bool
}

//
//
//
func NewDistantCustomerOriginator(address, carrier string, shouldLink bool) CustomerOriginator {
	return &distantCustomerOriginator{
		address:        address,
		carrier:   carrier,
		shouldLink: shouldLink,
	}
}

func (r *distantCustomerOriginator) NewIfaceCustomer() (abciend.Customer, error) {
	distantApplication, err := abciend.NewCustomer(r.address, r.carrier, r.shouldLink)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return distantApplication, nil
}

//
//
//
//
//
//
//
//
//
func StandardCustomerOriginator(address, carrier, storeFolder string) CustomerOriginator {
	switch address {
	case "REDACTED":
		return NewNativeCustomerOriginator(objectdepot.NewInRamSoftware())
	case "REDACTED":
		return NewLinkAlignNativeCustomerOriginator(objectdepot.NewInRamSoftware())
	case "REDACTED":
		return NewNativeCustomerOriginator(objectdepot.NewDurableSoftware(storeFolder))
	case "REDACTED":
		return NewLinkAlignNativeCustomerOriginator(objectdepot.NewDurableSoftware(storeFolder))
	case "REDACTED":
		app, err := e2e.NewSoftware(e2e.StandardSettings(storeFolder))
		if err != nil {
			panic(err)
		}
		return NewNativeCustomerOriginator(app)
	case "REDACTED":
		app, err := e2e.NewSoftware(e2e.StandardSettings(storeFolder))
		if err != nil {
			panic(err)
		}
		return NewLinkAlignNativeCustomerOriginator(app)
	case "REDACTED":
		return NewNativeCustomerOriginator(kinds.NewRootSoftware())
	default:
		shouldLink := false //
		return NewDistantCustomerOriginator(address, carrier, shouldLink)
	}
}
