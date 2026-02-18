package native

import (
	"context"
	"fmt"
	"time"

	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	cmtinquire "github.com/valkyrieworks/utils/broadcast/inquire"
	nm "github.com/valkyrieworks/member"
	rpccustomer "github.com/valkyrieworks/rpc/customer"
	"github.com/valkyrieworks/rpc/core"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/kinds"
)

/**
c
.

:

d
r
y
.

e
.

.
T
u
l
.
*/
type Native struct {
	*kinds.EventBus
	Tracer log.Tracer
	ctx    *rpctypes.Context
	env    *core.Context
}

//
func New(member *nm.Member) *Native {
	env, err := member.SetupRPC()
	if err != nil {
		member.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	return &Native{
		EventBus: member.EventBus(),
		Tracer:   log.NewNoopTracer(),
		ctx:      &rpctypes.Context{},
		env:      env,
	}
}

var _ rpccustomer.Customer = (*Native)(nil)

//
func (c *Native) AssignTracer(l log.Tracer) {
	c.Tracer = l
}

func (c *Native) Status(context.Context) (*ctypes.OutcomeState, error) {
	return c.env.Status(c.ctx)
}

func (c *Native) IfaceDetails(context.Context) (*ctypes.OutcomeIfaceDetails, error) {
	return c.env.IfaceDetails(c.ctx)
}

func (c *Native) IfaceInquire(ctx context.Context, route string, data octets.HexOctets) (*ctypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireWithSettings(ctx, route, data, rpccustomer.StandardIfaceInquireSettings)
}

func (c *Native) IfaceInquireWithSettings(
	_ context.Context,
	route string,
	data octets.HexOctets,
	opts rpccustomer.IfaceInquireSettings,
) (*ctypes.OutcomeIfaceInquire, error) {
	return c.env.IfaceInquire(c.ctx, route, data, opts.Level, opts.Demonstrate)
}

func (c *Native) MulticastTransferEndorse(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransferEndorse, error) {
	return c.env.MulticastTransferEndorse(c.ctx, tx)
}

func (c *Native) MulticastTransferAsync(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferAsync(c.ctx, tx)
}

func (c *Native) MulticastTransferAlign(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferAlign(c.ctx, tx)
}

func (c *Native) UnattestedTrans(_ context.Context, ceiling *int) (*ctypes.OutcomeUnattestedTrans, error) {
	return c.env.UnattestedTrans(c.ctx, ceiling)
}

func (c *Native) CountUnattestedTrans(context.Context) (*ctypes.OutcomeUnattestedTrans, error) {
	return c.env.CountUnattestedTrans(c.ctx)
}

func (c *Native) InspectTransfer(_ context.Context, tx kinds.Tx) (*ctypes.OutcomeInspectTransfer, error) {
	return c.env.InspectTransfer(c.ctx, tx)
}

func (c *Native) NetDetails(context.Context) (*ctypes.OutcomeNetDetails, error) {
	return c.env.NetDetails(c.ctx)
}

func (c *Native) ExportAgreementStatus(context.Context) (*ctypes.OutcomeExportAgreementStatus, error) {
	return c.env.ExportAgreementStatus(c.ctx)
}

func (c *Native) AgreementStatus(context.Context) (*ctypes.OutcomeAgreementStatus, error) {
	return c.env.FetchAgreementStatus(c.ctx)
}

func (c *Native) AgreementOptions(_ context.Context, level *int64) (*ctypes.OutcomeAgreementOptions, error) {
	return c.env.AgreementOptions(c.ctx, level)
}

func (c *Native) Vitality(context.Context) (*ctypes.OutcomeVitality, error) {
	return c.env.Vitality(c.ctx)
}

func (c *Native) CallOrigins(_ context.Context, origins []string) (*ctypes.OutcomeCallOrigins, error) {
	return c.env.RiskyCallOrigins(c.ctx, origins)
}

func (c *Native) CallNodes(
	_ context.Context,
	nodes []string,
	durable,
	absolute,
	internal bool,
) (*ctypes.OutcomeCallNodes, error) {
	return c.env.RiskyCallNodes(c.ctx, nodes, durable, absolute, internal)
}

func (c *Native) LedgerchainDetails(_ context.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeLedgerchainDetails, error) {
	return c.env.LedgerchainDetails(c.ctx, minimumLevel, maximumLevel)
}

func (c *Native) Origin(context.Context) (*ctypes.OutcomeOrigin, error) {
	return c.env.Origin(c.ctx)
}

func (c *Native) OriginSegmented(_ context.Context, id uint) (*ctypes.OutcomeOriginSegment, error) {
	return c.env.OriginSegmented(c.ctx, id)
}

func (c *Native) Ledger(_ context.Context, level *int64) (*ctypes.OutcomeLedger, error) {
	return c.env.Ledger(c.ctx, level)
}

func (c *Native) LedgerByDigest(_ context.Context, digest []byte) (*ctypes.OutcomeLedger, error) {
	return c.env.LedgerByDigest(c.ctx, digest)
}

func (c *Native) LedgerOutcomes(_ context.Context, level *int64) (*ctypes.OutcomeLedgerOutcomes, error) {
	return c.env.LedgerOutcomes(c.ctx, level)
}

func (c *Native) Heading(_ context.Context, level *int64) (*ctypes.OutcomeHeading, error) {
	return c.env.Heading(c.ctx, level)
}

func (c *Native) HeadingByDigest(_ context.Context, digest octets.HexOctets) (*ctypes.OutcomeHeading, error) {
	return c.env.HeadingByDigest(c.ctx, digest)
}

func (c *Native) Endorse(_ context.Context, level *int64) (*ctypes.OutcomeEndorse, error) {
	return c.env.Endorse(c.ctx, level)
}

func (c *Native) Ratifiers(_ context.Context, level *int64, screen, eachScreen *int) (*ctypes.OutcomeRatifiers, error) {
	return c.env.Ratifiers(c.ctx, level, screen, eachScreen)
}

func (c *Native) Tx(_ context.Context, digest []byte, demonstrate bool) (*ctypes.OutcomeTransfer, error) {
	return c.env.Tx(c.ctx, digest, demonstrate)
}

func (c *Native) TransferScan(
	_ context.Context,
	inquire string,
	demonstrate bool,
	screen,
	eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeTransferScan, error) {
	return c.env.TransferScan(c.ctx, inquire, demonstrate, screen, eachScreen, arrangeBy)
}

func (c *Native) LedgerScan(
	_ context.Context,
	inquire string,
	screen, eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeLedgerScan, error) {
	return c.env.LedgerScan(c.ctx, inquire, screen, eachScreen, arrangeBy)
}

func (c *Native) MulticastProof(_ context.Context, ev kinds.Proof) (*ctypes.OutcomeMulticastProof, error) {
	return c.env.MulticastProof(c.ctx, ev)
}

func (c *Native) Enrol(
	ctx context.Context,
	enrollee,
	inquire string,
	outVolume ...int,
) (out <-chan ctypes.OutcomeEvent, err error) {
	q, err := cmtinquire.New(inquire)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	outCeiling := 1
	if len(outVolume) > 0 {
		outCeiling = outVolume[0]
	}

	var sub kinds.Enrollment
	if outCeiling > 0 {
		sub, err = c.EventBus.Enrol(ctx, enrollee, q, outCeiling)
	} else {
		sub, err = c.EnrolUnbuffered(ctx, enrollee, q)
	}
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	outchan := make(chan ctypes.OutcomeEvent, outCeiling)
	go c.eventsProcedure(sub, enrollee, q, outchan)

	return outchan, nil
}

func (c *Native) eventsProcedure(
	sub kinds.Enrollment,
	enrollee string,
	q cometbroadcast.Inquire,
	outchan chan<- ctypes.OutcomeEvent,
) {
	for {
		select {
		case msg := <-sub.Out():
			outcome := ctypes.OutcomeEvent{Inquire: q.String(), Data: msg.Data(), Events: msg.Events()}
			if cap(outchan) == 0 {
				outchan <- outcome
			} else {
				select {
				case outchan <- outcome:
				default:
					c.Tracer.Fault("REDACTED", "REDACTED", outcome, "REDACTED", outcome.Inquire)
				}
			}
		case <-sub.Revoked():
			if sub.Err() == cometbroadcast.ErrDeactivated {
				return
			}

			c.Tracer.Fault("REDACTED", "REDACTED", sub.Err(), "REDACTED", q.String())
			sub = c.reactivate(enrollee, q)
			if sub == nil { //
				return
			}
		case <-c.Exit():
			return
		}
	}
}

//
func (c *Native) reactivate(enrollee string, q cometbroadcast.Inquire) kinds.Enrollment {
	tries := 0
	for {
		if !c.IsActive() {
			return nil
		}

		sub, err := c.EventBus.Enrol(context.Background(), enrollee, q)
		if err == nil {
			return sub
		}

		tries++
		time.Sleep((10 << uint(tries)) * time.Millisecond) //
	}
}

func (c *Native) Deenroll(ctx context.Context, enrollee, inquire string) error {
	q, err := cmtinquire.New(inquire)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return c.EventBus.Deenroll(ctx, enrollee, q)
}

func (c *Native) DeenrollAll(ctx context.Context, enrollee string) error {
	return c.EventBus.DeenrollAll(ctx, enrollee)
}
