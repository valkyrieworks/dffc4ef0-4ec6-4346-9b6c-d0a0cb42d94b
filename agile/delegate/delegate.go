package delegate

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	airpc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/rpc"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	rpchandler "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
)

//
type Delegate struct {
	Location     string //
	Settings   *rpchandler.Settings
	Customer   *airpc.Customer
	Tracer   log.Tracer
	Observer net.Listener
}

//
//
func FreshDelegate(
	agileCustomer *agile.Customer,
	overhearLocation, supplierLocation string,
	settings *rpchandler.Settings,
	tracer log.Tracer,
	choices ...airpc.Selection,
) (*Delegate, error) {
	ifaceCustomer, err := rpchttpsvc.FreshUsingDeadline(supplierLocation, "REDACTED", uint(settings.PersistDeadline.Seconds()))
	if err != nil {
		return nil, fmt.Errorf("REDACTED", supplierLocation, err)
	}

	return &Delegate{
		Location:   overhearLocation,
		Settings: settings,
		Customer: airpc.FreshCustomer(ifaceCustomer, agileCustomer, choices...),
		Tracer: tracer,
	}, nil
}

//
//
//
//
func (p *Delegate) OverhearAlsoAttend() error {
	observer, mux, err := p.overhear()
	if err != nil {
		return err
	}
	p.Observer = observer

	return rpchandler.Attend(
		observer,
		mux,
		p.Tracer,
		p.Settings,
	)
}

//
//
//
func (p *Delegate) OverhearAlsoAttendTransportsec(licenseRecord, tokenRecord string) error {
	observer, mux, err := p.overhear()
	if err != nil {
		return err
	}
	p.Observer = observer

	return rpchandler.AttendTransportsec(
		observer,
		mux,
		licenseRecord,
		tokenRecord,
		p.Tracer,
		p.Settings,
	)
}

func (p *Delegate) overhear() (net.Listener, *http.ServeMux, error) {
	mux := http.NewServeMux()

	//
	r := RemotePaths(p.Customer)
	rpchandler.EnrollRemoteRoutines(mux, r, p.Tracer)

	//
	watermarkTracer := p.Tracer.Using("REDACTED", "REDACTED")
	wm := rpchandler.FreshWebterminalAdministrator(r,
		rpchandler.UponDetach(func(distantLocation string) {
			err := p.Customer.UnlistenEvery(context.Background(), distantLocation)
			if err != nil && err != tendermintpubsub.FaultListeningNegationDetected {
				watermarkTracer.Failure("REDACTED", "REDACTED", distantLocation, "REDACTED", err)
			}
		}),
		rpchandler.RetrieveThreshold(p.Settings.MaximumContentOctets),
	)
	wm.AssignTracer(watermarkTracer)
	mux.HandleFunc("REDACTED", wm.WebterminalProcessor)

	//
	if !p.Customer.EqualsActive() {
		if err := p.Customer.Initiate(); err != nil {
			return nil, mux, fmt.Errorf("REDACTED", err)
		}
	}

	//
	observer, err := rpchandler.Overhear(p.Location, p.Settings.MaximumInitiateLinks)
	if err != nil {
		return nil, mux, err
	}

	return observer, mux, nil
}
