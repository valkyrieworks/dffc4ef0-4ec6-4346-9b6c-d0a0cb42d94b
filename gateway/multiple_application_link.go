package gateway

import (
	"fmt"

	abciend "github.com/valkyrieworks/iface/customer"
	cmttrace "github.com/valkyrieworks/utils/log"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/utils/daemon"
)

const (
	linkAgreement = "REDACTED"
	linkTxpool   = "REDACTED"
	linkInquire     = "REDACTED"
	linkMirror  = "REDACTED"
)

//
//
type ApplicationLinks interface {
	daemon.Daemon

	//
	Txpool() ApplicationLinkTxpool
	//
	Agreement() ApplicationLinkAgreement
	//
	Inquire() ApplicationLinkInquire
	//
	Mirror() ApplicationLinkMirror
}

//
func NewApplicationLinks(customerOriginator CustomerOriginator, stats *Stats) ApplicationLinks {
	return NewMultipleApplicationLink(customerOriginator, stats)
}

//
//
//
//
//
type multipleApplicationLink struct {
	daemon.RootDaemon

	stats       *Stats
	agreementLink ApplicationLinkAgreement
	txpoolLink   ApplicationLinkTxpool
	inquireLink     ApplicationLinkInquire
	mirrorLink  ApplicationLinkMirror

	agreementLinkCustomer abciend.Customer
	txpoolLinkCustomer   abciend.Customer
	inquireLinkCustomer     abciend.Customer
	mirrorLinkCustomer  abciend.Customer

	customerOriginator CustomerOriginator
}

//
func NewMultipleApplicationLink(customerOriginator CustomerOriginator, stats *Stats) ApplicationLinks {
	multipleApplicationLink := &multipleApplicationLink{
		stats:       stats,
		customerOriginator: customerOriginator,
	}
	multipleApplicationLink.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", multipleApplicationLink)
	return multipleApplicationLink
}

func (app *multipleApplicationLink) Txpool() ApplicationLinkTxpool {
	return app.txpoolLink
}

func (app *multipleApplicationLink) Agreement() ApplicationLinkAgreement {
	return app.agreementLink
}

func (app *multipleApplicationLink) Inquire() ApplicationLinkInquire {
	return app.inquireLink
}

func (app *multipleApplicationLink) Mirror() ApplicationLinkMirror {
	return app.mirrorLink
}

func (app *multipleApplicationLink) OnBegin() error {
	c, err := app.ifaceCustomerFor(linkInquire)
	if err != nil {
		return err
	}
	app.inquireLinkCustomer = c
	app.inquireLink = NewApplicationLinkInquire(c, app.stats)

	c, err = app.ifaceCustomerFor(linkMirror)
	if err != nil {
		app.haltAllAgents()
		return err
	}
	app.mirrorLinkCustomer = c
	app.mirrorLink = NewApplicationLinkMirror(c, app.stats)

	c, err = app.ifaceCustomerFor(linkTxpool)
	if err != nil {
		app.haltAllAgents()
		return err
	}
	app.txpoolLinkCustomer = c
	app.txpoolLink = NewApplicationLinkTxpool(c, app.stats)

	c, err = app.ifaceCustomerFor(linkAgreement)
	if err != nil {
		app.haltAllAgents()
		return err
	}
	app.agreementLinkCustomer = c
	app.agreementLink = NewApplicationLinkAgreement(c, app.stats)

	//
	go app.abortTMOnCustomerFault()

	return nil
}

func (app *multipleApplicationLink) OnHalt() {
	app.haltAllAgents()
}

func (app *multipleApplicationLink) abortTMOnCustomerFault() {
	abortFn := func(link string, err error, tracer cmttrace.Tracer) {
		tracer.Fault(
			fmt.Sprintf("REDACTED", link),
			"REDACTED", err)
		abortErr := cometos.Abort()
		if abortErr != nil {
			tracer.Fault("REDACTED", "REDACTED", abortErr)
		}
	}

	select {
	case <-app.agreementLinkCustomer.Exit():
		if err := app.agreementLinkCustomer.Fault(); err != nil {
			abortFn(linkAgreement, err, app.Tracer)
		}
	case <-app.txpoolLinkCustomer.Exit():
		if err := app.txpoolLinkCustomer.Fault(); err != nil {
			abortFn(linkTxpool, err, app.Tracer)
		}
	case <-app.inquireLinkCustomer.Exit():
		if err := app.inquireLinkCustomer.Fault(); err != nil {
			abortFn(linkInquire, err, app.Tracer)
		}
	case <-app.mirrorLinkCustomer.Exit():
		if err := app.mirrorLinkCustomer.Fault(); err != nil {
			abortFn(linkMirror, err, app.Tracer)
		}
	}
}

func (app *multipleApplicationLink) haltAllAgents() {
	if app.agreementLinkCustomer != nil {
		if err := app.agreementLinkCustomer.Halt(); err != nil {
			app.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	if app.txpoolLinkCustomer != nil {
		if err := app.txpoolLinkCustomer.Halt(); err != nil {
			app.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	if app.inquireLinkCustomer != nil {
		if err := app.inquireLinkCustomer.Halt(); err != nil {
			app.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	if app.mirrorLinkCustomer != nil {
		if err := app.mirrorLinkCustomer.Halt(); err != nil {
			app.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
}

func (app *multipleApplicationLink) ifaceCustomerFor(link string) (abciend.Customer, error) {
	c, err := app.customerOriginator.NewIfaceCustomer()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", link, err)
	}
	c.AssignTracer(app.Tracer.With("REDACTED", "REDACTED", "REDACTED", link))
	if err := c.Begin(); err != nil {
		return nil, fmt.Errorf("REDACTED", link, err)
	}
	return c, nil
}
