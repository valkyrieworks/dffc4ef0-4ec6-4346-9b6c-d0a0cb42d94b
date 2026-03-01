package simulate

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type Customer struct {
	customer.IfaceCustomer
	customer.AttestCustomer
	customer.ChronicleCustomer
	customer.ConditionCustomer
	customer.IncidentsCustomer
	customer.ProofCustomer
	customer.TxpoolCustomer
	facility.Facility

	env *base.Context
}

func New() Customer {
	return Customer{
		env: &base.Context{},
	}
}

var _ customer.Customer = Customer{}

//
//
type Invocation struct {
	Alias     string
	Arguments     any
	Reply any
	Failure    error
}

//
//
//
//
//
//
func (c Invocation) ObtainReply(arguments any) (any, error) {
	//
	if c.Reply == nil {
		if c.Failure == nil {
			panic("REDACTED")
		}
		return nil, c.Failure
	}
	//
	if c.Failure == nil {
		return c.Reply, nil
	}
	//
	if reflect.DeepEqual(arguments, c.Arguments) {
		return c.Reply, nil
	}
	return nil, c.Failure
}

func (c Customer) Condition(context.Context) (*ktypes.OutcomeCondition, error) {
	return c.env.Condition(&remoteifacetypes.Env{})
}

func (c Customer) IfaceDetails(context.Context) (*ktypes.OutcomeIfaceDetails, error) {
	return c.env.IfaceDetails(&remoteifacetypes.Env{})
}

func (c Customer) IfaceInquire(ctx context.Context, route string, data octets.HexadecimalOctets) (*ktypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireUsingChoices(ctx, route, data, customer.FallbackIfaceInquireChoices)
}

func (c Customer) IfaceInquireUsingChoices(
	_ context.Context,
	route string,
	data octets.HexadecimalOctets,
	choices customer.IfaceInquireChoices,
) (*ktypes.OutcomeIfaceInquire, error) {
	return c.env.IfaceInquire(&remoteifacetypes.Env{}, route, data, choices.Altitude, choices.Validate)
}

func (c Customer) MulticastTransferEndorse(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransferEndorse, error) {
	return c.env.MulticastTransferEndorse(&remoteifacetypes.Env{}, tx)
}

func (c Customer) MulticastTransferAsyncronous(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferAsyncronous(&remoteifacetypes.Env{}, tx)
}

func (c Customer) MulticastTransferChronize(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferChronize(&remoteifacetypes.Env{}, tx)
}

func (c Customer) InspectTransfer(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeInspectTransfer, error) {
	return c.env.InspectTransfer(&remoteifacetypes.Env{}, tx)
}

func (c Customer) NetworkDetails(_ context.Context) (*ktypes.OutcomeNetworkDetails, error) {
	return c.env.NetworkDetails(&remoteifacetypes.Env{})
}

func (c Customer) AgreementStatus(_ context.Context) (*ktypes.OutcomeAgreementStatus, error) {
	return c.env.ObtainAgreementStatus(&remoteifacetypes.Env{})
}

func (c Customer) ExportAgreementStatus(_ context.Context) (*ktypes.OutcomeExportAgreementStatus, error) {
	return c.env.ExportAgreementStatus(&remoteifacetypes.Env{})
}

func (c Customer) AgreementSettings(_ context.Context, altitude *int64) (*ktypes.OutcomeAgreementParameters, error) {
	return c.env.AgreementSettings(&remoteifacetypes.Env{}, altitude)
}

func (c Customer) Vitality(_ context.Context) (*ktypes.OutcomeVitality, error) {
	return c.env.Vitality(&remoteifacetypes.Env{})
}

func (c Customer) CallOrigins(_ context.Context, origins []string) (*ktypes.OutcomeCallOrigins, error) {
	return c.env.InsecureCallOrigins(&remoteifacetypes.Env{}, origins)
}

func (c Customer) CallNodes(
	_ context.Context,
	nodes []string,
	enduring,
	absolute,
	secluded bool,
) (*ktypes.OutcomeCallNodes, error) {
	return c.env.InsecureCallNodes(&remoteifacetypes.Env{}, nodes, enduring, absolute, secluded)
}

func (c Customer) LedgerchainDetails(_ context.Context, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeLedgerchainDetails, error) {
	return c.env.LedgerchainDetails(&remoteifacetypes.Env{}, minimumAltitude, maximumAltitude)
}

func (c Customer) Inauguration(context.Context) (*ktypes.OutcomeInauguration, error) {
	return c.env.Inauguration(&remoteifacetypes.Env{})
}

func (c Customer) Ledger(_ context.Context, altitude *int64) (*ktypes.OutcomeLedger, error) {
	return c.env.Ledger(&remoteifacetypes.Env{}, altitude)
}

func (c Customer) LedgerViaDigest(_ context.Context, digest []byte) (*ktypes.OutcomeLedger, error) {
	return c.env.LedgerViaDigest(&remoteifacetypes.Env{}, digest)
}

func (c Customer) Endorse(_ context.Context, altitude *int64) (*ktypes.OutcomeEndorse, error) {
	return c.env.Endorse(&remoteifacetypes.Env{}, altitude)
}

func (c Customer) Assessors(_ context.Context, altitude *int64, screen, everyScreen *int) (*ktypes.OutcomeAssessors, error) {
	return c.env.Assessors(&remoteifacetypes.Env{}, altitude, screen, everyScreen)
}

func (c Customer) MulticastProof(_ context.Context, ev kinds.Proof) (*ktypes.OutcomeMulticastProof, error) {
	return c.env.MulticastProof(&remoteifacetypes.Env{}, ev)
}
