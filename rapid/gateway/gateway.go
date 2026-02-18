package gateway

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/valkyrieworks/utils/log"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	"github.com/valkyrieworks/rapid"
	lrpc "github.com/valkyrieworks/rapid/rpc"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	rpchost "github.com/valkyrieworks/rpc/jsonrpc/host"
)

//
type Gateway struct {
	Address     string //
	Settings   *rpchost.Settings
	Customer   *lrpc.Customer
	Tracer   log.Tracer
	Observer net.Listener
}

//
//
func NewGateway(
	rapidCustomer *rapid.Customer,
	acceptAddress, sourceAddress string,
	settings *rpchost.Settings,
	tracer log.Tracer,
	opts ...lrpc.Setting,
) (*Gateway, error) {
	rpcCustomer, err := rpchttp.NewWithDeadline(sourceAddress, "REDACTED", uint(settings.RecordDeadline.Seconds()))
	if err != nil {
		return nil, fmt.Errorf("REDACTED", sourceAddress, err)
	}

	return &Gateway{
		Address:   acceptAddress,
		Settings: settings,
		Customer: lrpc.NewCustomer(rpcCustomer, rapidCustomer, opts...),
		Tracer: tracer,
	}, nil
}

//
//
//
//
func (p *Gateway) AcceptAndHost() error {
	observer, mux, err := p.observe()
	if err != nil {
		return err
	}
	p.Observer = observer

	return rpchost.Attend(
		observer,
		mux,
		p.Tracer,
		p.Settings,
	)
}

//
//
//
func (p *Gateway) ObserveAndAttendTLS(tokenEntry, keyEntry string) error {
	observer, mux, err := p.observe()
	if err != nil {
		return err
	}
	p.Observer = observer

	return rpchost.AttendTLS(
		observer,
		mux,
		tokenEntry,
		keyEntry,
		p.Tracer,
		p.Settings,
	)
}

func (p *Gateway) observe() (net.Listener, *http.ServeMux, error) {
	mux := http.NewServeMux()

	//
	r := RPCPaths(p.Customer)
	rpchost.EnrollRPCRoutines(mux, r, p.Tracer)

	//
	wmTracer := p.Tracer.With("REDACTED", "REDACTED")
	wm := rpchost.NewWebchannelOverseer(r,
		rpchost.OnDetach(func(distantAddress string) {
			err := p.Customer.DeenrollAll(context.Background(), distantAddress)
			if err != nil && err != cometbroadcast.ErrEnrollmentNegateLocated {
				wmTracer.Fault("REDACTED", "REDACTED", distantAddress, "REDACTED", err)
			}
		}),
		rpchost.ScanCeiling(p.Settings.MaximumContentOctets),
	)
	wm.AssignTracer(wmTracer)
	mux.HandleFunc("REDACTED", wm.WebchannelManager)

	//
	if !p.Customer.IsActive() {
		if err := p.Customer.Begin(); err != nil {
			return nil, mux, fmt.Errorf("REDACTED", err)
		}
	}

	//
	observer, err := rpchost.Observe(p.Address, p.Settings.MaximumAccessLinks)
	if err != nil {
		return nil, mux, err
	}

	return observer, mux, nil
}
