package emulate

/**
t
.

t
t
,
.

u
e
.
*/

import (
	"context"
	"reflect"

	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/rpc/customer"
	"github.com/valkyrieworks/rpc/core"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/kinds"
)

//
type Customer struct {
	customer.IfaceCustomer
	customer.AttestCustomer
	customer.LogbookCustomer
	customer.StateCustomer
	customer.EventsCustomer
	customer.ProofCustomer
	customer.TxpoolCustomer
	daemon.Daemon

	env *core.Context
}

func New() Customer {
	return Customer{
		env: &core.Context{},
	}
}

var _ customer.Customer = Customer{}

//
//
type Invoke struct {
	Label     string
	Args     any
	Reply any
	Fault    error
}

//
//
//
//
//
//
func (c Invoke) FetchReply(args any) (any, error) {
	//
	if c.Reply == nil {
		if c.Fault == nil {
			panic("REDACTED")
		}
		return nil, c.Fault
	}
	//
	if c.Fault == nil {
		return c.Reply, nil
	}
	//
	if reflect.DeepEqual(args, c.Args) {
		return c.Reply, nil
	}
	return nil, c.Fault
}

func (c Customer) Status(context.Context) (*ctypes.OutcomeState, error) {
	return c.env.Status(&rpctypes.Context{})
}

func (c Customer) IfaceDetails(context.Context) (*ctypes.OutcomeIfaceDetails, error) {
	return c.env.IfaceDetails(&rpctypes.Context{})
}

func (c Customer) IfaceInquire(ctx context.Context, route string, data octets.HexOctets) (*ctypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireWithSettings(ctx, route, data, customer.StandardIfaceInquireSettings)
}

func (c Customer) IfaceInquireWithSettings(
	_ context.Context,
	route string,
	data octets.HexOctets,
	opts customer.IfaceInquireSettings,
) (*ctypes.OutcomeIfaceInquire, error) {
	return c.env.IfaceInquire(&rpctypes.Context{}, route, data, opts.Level, opts.Demonstrate)
}

func (c Customer) MulticastTransferEndorse(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransferEndorse, error) {
	return c.env.MulticastTransferEndorse(&rpctypes.Context{}, tx)
}

func (c Customer) MulticastTransferAsync(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferAsync(&rpctypes.Context{}, tx)
}

func (c Customer) MulticastTransferAlign(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferAlign(&rpctypes.Context{}, tx)
}

func (c Customer) InspectTransfer(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeInspectTransfer, error) {
	return c.env.InspectTransfer(&rpctypes.Context{}, tx)
}

func (c Customer) NetDetails(_ context.Context) (*ctypes.OutcomeNetDetails, error) {
	return c.env.NetDetails(&rpctypes.Context{})
}

func (c Customer) AgreementStatus(_ context.Context) (*ctypes.OutcomeAgreementStatus, error) {
	return c.env.FetchAgreementStatus(&rpctypes.Context{})
}

func (c Customer) ExportAgreementStatus(_ context.Context) (*ctypes.OutcomeExportAgreementStatus, error) {
	return c.env.ExportAgreementStatus(&rpctypes.Context{})
}

func (c Customer) AgreementOptions(_ context.Context, level *int64) (*ctypes.OutcomeAgreementOptions, error) {
	return c.env.AgreementOptions(&rpctypes.Context{}, level)
}

func (c Customer) Vitality(_ context.Context) (*ctypes.OutcomeVitality, error) {
	return c.env.Vitality(&rpctypes.Context{})
}

func (c Customer) CallOrigins(_ context.Context, origins []string) (*ctypes.OutcomeCallOrigins, error) {
	return c.env.RiskyCallOrigins(&rpctypes.Context{}, origins)
}

func (c Customer) CallNodes(
	_ context.Context,
	nodes []string,
	durable,
	absolute,
	internal bool,
) (*ctypes.OutcomeCallNodes, error) {
	return c.env.RiskyCallNodes(&rpctypes.Context{}, nodes, durable, absolute, internal)
}

func (c Customer) LedgerchainDetails(_ context.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeLedgerchainDetails, error) {
	return c.env.LedgerchainDetails(&rpctypes.Context{}, minimumLevel, maximumLevel)
}

func (c Customer) Origin(context.Context) (*ctypes.OutcomeOrigin, error) {
	return c.env.Origin(&rpctypes.Context{})
}

func (c Customer) Ledger(_ context.Context, level *int64) (*ctypes.OutcomeLedger, error) {
	return c.env.Ledger(&rpctypes.Context{}, level)
}

func (c Customer) LedgerByDigest(_ context.Context, digest []byte) (*ctypes.OutcomeLedger, error) {
	return c.env.LedgerByDigest(&rpctypes.Context{}, digest)
}

func (c Customer) Endorse(_ context.Context, level *int64) (*ctypes.OutcomeEndorse, error) {
	return c.env.Endorse(&rpctypes.Context{}, level)
}

func (c Customer) Ratifiers(_ context.Context, level *int64, screen, eachScreen *int) (*ctypes.OutcomeRatifiers, error) {
	return c.env.Ratifiers(&rpctypes.Context{}, level, screen, eachScreen)
}

func (c Customer) MulticastProof(_ context.Context, ev kinds.Proof) (*ctypes.OutcomeMulticastProof, error) {
	return c.env.MulticastProof(&rpctypes.Context{}, ev)
}
