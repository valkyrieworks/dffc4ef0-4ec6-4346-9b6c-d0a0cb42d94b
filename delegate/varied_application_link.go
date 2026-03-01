package delegate

import (
	"fmt"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	endorsementlog "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

const (
	linkAgreement = "REDACTED"
	linkTxpool   = "REDACTED"
	linkInquire     = "REDACTED"
	linkImage  = "REDACTED"
)

//
//
type PlatformLinks interface {
	facility.Facility

	//
	Txpool() ApplicationLinkTxpool
	//
	Agreement() ApplicationLinkAgreement
	//
	Inquire() PlatformLinkInquire
	//
	Image() PlatformLinkImage
}

//
func FreshPlatformLinks(customerOriginator CustomerOriginator, telemetry *Telemetry) PlatformLinks {
	return FreshVariedApplicationLink(customerOriginator, telemetry)
}

//
//
//
//
//
type variedApplicationLink struct {
	facility.FoundationFacility

	telemetry       *Telemetry
	agreementLink ApplicationLinkAgreement
	txpoolLink   ApplicationLinkTxpool
	inquireLink     PlatformLinkInquire
	imageLink  PlatformLinkImage

	agreementLinkCustomer abcicustomer.Customer
	txpoolLinkCustomer   abcicustomer.Customer
	inquireLinkCustomer     abcicustomer.Customer
	imageLinkCustomer  abcicustomer.Customer

	customerOriginator CustomerOriginator
}

//
func FreshVariedApplicationLink(customerOriginator CustomerOriginator, telemetry *Telemetry) PlatformLinks {
	variedApplicationLink := &variedApplicationLink{
		telemetry:       telemetry,
		customerOriginator: customerOriginator,
	}
	variedApplicationLink.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", variedApplicationLink)
	return variedApplicationLink
}

func (app *variedApplicationLink) Txpool() ApplicationLinkTxpool {
	return app.txpoolLink
}

func (app *variedApplicationLink) Agreement() ApplicationLinkAgreement {
	return app.agreementLink
}

func (app *variedApplicationLink) Inquire() PlatformLinkInquire {
	return app.inquireLink
}

func (app *variedApplicationLink) Image() PlatformLinkImage {
	return app.imageLink
}

func (app *variedApplicationLink) UponInitiate() error {
	c, err := app.ifaceCustomerForeach(linkInquire)
	if err != nil {
		return err
	}
	app.inquireLinkCustomer = c
	app.inquireLink = FreshApplicationLinkInquire(c, app.telemetry)

	c, err = app.ifaceCustomerForeach(linkImage)
	if err != nil {
		app.haltEveryCustomers()
		return err
	}
	app.imageLinkCustomer = c
	app.imageLink = FreshApplicationLinkImage(c, app.telemetry)

	c, err = app.ifaceCustomerForeach(linkTxpool)
	if err != nil {
		app.haltEveryCustomers()
		return err
	}
	app.txpoolLinkCustomer = c
	app.txpoolLink = FreshApplicationLinkTxpool(c, app.telemetry)

	c, err = app.ifaceCustomerForeach(linkAgreement)
	if err != nil {
		app.haltEveryCustomers()
		return err
	}
	app.agreementLinkCustomer = c
	app.agreementLink = FreshApplicationLinkAgreement(c, app.telemetry)

	//
	go app.terminateTEMPUponCustomerFailure()

	return nil
}

func (app *variedApplicationLink) UponHalt() {
	app.haltEveryCustomers()
}

func (app *variedApplicationLink) terminateTEMPUponCustomerFailure() {
	terminateProc := func(link string, err error, tracer endorsementlog.Tracer) {
		tracer.Failure(
			fmt.Sprintf("REDACTED", link),
			"REDACTED", err)
		terminateFault := strongos.Terminate()
		if terminateFault != nil {
			tracer.Failure("REDACTED", "REDACTED", terminateFault)
		}
	}

	select {
	case <-app.agreementLinkCustomer.Exit():
		if err := app.agreementLinkCustomer.Failure(); err != nil {
			terminateProc(linkAgreement, err, app.Tracer)
		}
	case <-app.txpoolLinkCustomer.Exit():
		if err := app.txpoolLinkCustomer.Failure(); err != nil {
			terminateProc(linkTxpool, err, app.Tracer)
		}
	case <-app.inquireLinkCustomer.Exit():
		if err := app.inquireLinkCustomer.Failure(); err != nil {
			terminateProc(linkInquire, err, app.Tracer)
		}
	case <-app.imageLinkCustomer.Exit():
		if err := app.imageLinkCustomer.Failure(); err != nil {
			terminateProc(linkImage, err, app.Tracer)
		}
	}
}

func (app *variedApplicationLink) haltEveryCustomers() {
	if app.agreementLinkCustomer != nil {
		if err := app.agreementLinkCustomer.Halt(); err != nil {
			app.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	if app.txpoolLinkCustomer != nil {
		if err := app.txpoolLinkCustomer.Halt(); err != nil {
			app.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	if app.inquireLinkCustomer != nil {
		if err := app.inquireLinkCustomer.Halt(); err != nil {
			app.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	if app.imageLinkCustomer != nil {
		if err := app.imageLinkCustomer.Halt(); err != nil {
			app.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
}

func (app *variedApplicationLink) ifaceCustomerForeach(link string) (abcicustomer.Customer, error) {
	c, err := app.customerOriginator.FreshIfaceCustomer()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", link, err)
	}
	c.AssignTracer(app.Tracer.Using("REDACTED", "REDACTED", "REDACTED", link))
	if err := c.Initiate(); err != nil {
		return nil, fmt.Errorf("REDACTED", link, err)
	}
	return c, nil
}
