package delegate

import (
	"fmt"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/app"
)

//

//
type CustomerOriginator interface {
	//
	FreshIfaceCustomer() (abcicustomer.Customer, error)
}

//
//

type regionalCustomerOriginator struct {
	mtx *commitchronize.Exclusion
	app kinds.Platform
}

//
//
//
//
//
//
func FreshRegionalCustomerOriginator(app kinds.Platform) CustomerOriginator {
	return &regionalCustomerOriginator{
		mtx: new(commitchronize.Exclusion),
		app: app,
	}
}

func (l *regionalCustomerOriginator) FreshIfaceCustomer() (abcicustomer.Customer, error) {
	return abcicustomer.FreshRegionalCustomer(l.mtx, l.app), nil
}

//
//

type linkChronizeRegionalCustomerOriginator struct {
	app kinds.Platform
}

//
//
//
//
//
//
//
func FreshLinkChronizeRegionalCustomerOriginator(app kinds.Platform) CustomerOriginator {
	return &linkChronizeRegionalCustomerOriginator{
		app: app,
	}
}

func (c *linkChronizeRegionalCustomerOriginator) FreshIfaceCustomer() (abcicustomer.Customer, error) {
	//
	//
	return abcicustomer.FreshRegionalCustomer(nil, c.app), nil
}

//
//

type distantCustomerOriginator struct {
	location        string
	carrier   string
	shouldRelate bool
}

//
//
//
func FreshDistantCustomerOriginator(location, carrier string, shouldRelate bool) CustomerOriginator {
	return &distantCustomerOriginator{
		location:        location,
		carrier:   carrier,
		shouldRelate: shouldRelate,
	}
}

func (r *distantCustomerOriginator) FreshIfaceCustomer() (abcicustomer.Customer, error) {
	distantApplication, err := abcicustomer.FreshCustomer(r.location, r.carrier, r.shouldRelate)
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
func FallbackCustomerOriginator(location, carrier, datastorePath string) CustomerOriginator {
	switch location {
	case "REDACTED":
		return FreshRegionalCustomerOriginator(statedepot.FreshInsideRamPlatform())
	case "REDACTED":
		return FreshLinkChronizeRegionalCustomerOriginator(statedepot.FreshInsideRamPlatform())
	case "REDACTED":
		return FreshRegionalCustomerOriginator(statedepot.FreshEnduringPlatform(datastorePath))
	case "REDACTED":
		return FreshLinkChronizeRegionalCustomerOriginator(statedepot.FreshEnduringPlatform(datastorePath))
	case "REDACTED":
		app, err := e2e.FreshPlatform(e2e.FallbackSettings(datastorePath))
		if err != nil {
			panic(err)
		}
		return FreshRegionalCustomerOriginator(app)
	case "REDACTED":
		app, err := e2e.FreshPlatform(e2e.FallbackSettings(datastorePath))
		if err != nil {
			panic(err)
		}
		return FreshLinkChronizeRegionalCustomerOriginator(app)
	case "REDACTED":
		return FreshRegionalCustomerOriginator(kinds.FreshFoundationPlatform())
	default:
		shouldRelate := false //
		return FreshDistantCustomerOriginator(location, carrier, shouldRelate)
	}
}
