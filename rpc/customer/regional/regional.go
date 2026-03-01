package regional

import (
	"context"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	tendermintinquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	nm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/peer"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
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
type Regional struct {
	*kinds.IncidentChannel
	Tracer log.Tracer
	ctx    *remoteifacetypes.Env
	env    *base.Context
}

//
func New(peer *nm.Peer) *Regional {
	env, err := peer.SetupRemote()
	if err != nil {
		peer.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	return &Regional{
		IncidentChannel: peer.IncidentChannel(),
		Tracer:   log.FreshNooperationTracer(),
		ctx:      &remoteifacetypes.Env{},
		env:      env,
	}
}

var _ customeriface.Customer = (*Regional)(nil)

//
func (c *Regional) AssignTracer(l log.Tracer) {
	c.Tracer = l
}

func (c *Regional) Condition(context.Context) (*ktypes.OutcomeCondition, error) {
	return c.env.Condition(c.ctx)
}

func (c *Regional) IfaceDetails(context.Context) (*ktypes.OutcomeIfaceDetails, error) {
	return c.env.IfaceDetails(c.ctx)
}

func (c *Regional) IfaceInquire(ctx context.Context, route string, data octets.HexadecimalOctets) (*ktypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireUsingChoices(ctx, route, data, customeriface.FallbackIfaceInquireChoices)
}

func (c *Regional) IfaceInquireUsingChoices(
	_ context.Context,
	route string,
	data octets.HexadecimalOctets,
	choices customeriface.IfaceInquireChoices,
) (*ktypes.OutcomeIfaceInquire, error) {
	return c.env.IfaceInquire(c.ctx, route, data, choices.Altitude, choices.Validate)
}

func (c *Regional) MulticastTransferEndorse(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransferEndorse, error) {
	return c.env.MulticastTransferEndorse(c.ctx, tx)
}

func (c *Regional) MulticastTransferAsyncronous(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferAsyncronous(c.ctx, tx)
}

func (c *Regional) MulticastTransferChronize(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.env.MulticastTransferChronize(c.ctx, tx)
}

func (c *Regional) PendingTrans(_ context.Context, threshold *int) (*ktypes.OutcomePendingTrans, error) {
	return c.env.PendingTrans(c.ctx, threshold)
}

func (c *Regional) CountPendingTrans(context.Context) (*ktypes.OutcomePendingTrans, error) {
	return c.env.CountPendingTrans(c.ctx)
}

func (c *Regional) InspectTransfer(_ context.Context, tx kinds.Tx) (*ktypes.OutcomeInspectTransfer, error) {
	return c.env.InspectTransfer(c.ctx, tx)
}

func (c *Regional) NetworkDetails(context.Context) (*ktypes.OutcomeNetworkDetails, error) {
	return c.env.NetworkDetails(c.ctx)
}

func (c *Regional) ExportAgreementStatus(context.Context) (*ktypes.OutcomeExportAgreementStatus, error) {
	return c.env.ExportAgreementStatus(c.ctx)
}

func (c *Regional) AgreementStatus(context.Context) (*ktypes.OutcomeAgreementStatus, error) {
	return c.env.ObtainAgreementStatus(c.ctx)
}

func (c *Regional) AgreementSettings(_ context.Context, altitude *int64) (*ktypes.OutcomeAgreementParameters, error) {
	return c.env.AgreementSettings(c.ctx, altitude)
}

func (c *Regional) Vitality(context.Context) (*ktypes.OutcomeVitality, error) {
	return c.env.Vitality(c.ctx)
}

func (c *Regional) CallOrigins(_ context.Context, origins []string) (*ktypes.OutcomeCallOrigins, error) {
	return c.env.InsecureCallOrigins(c.ctx, origins)
}

func (c *Regional) CallNodes(
	_ context.Context,
	nodes []string,
	enduring,
	absolute,
	secluded bool,
) (*ktypes.OutcomeCallNodes, error) {
	return c.env.InsecureCallNodes(c.ctx, nodes, enduring, absolute, secluded)
}

func (c *Regional) LedgerchainDetails(_ context.Context, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeLedgerchainDetails, error) {
	return c.env.LedgerchainDetails(c.ctx, minimumAltitude, maximumAltitude)
}

func (c *Regional) Inauguration(context.Context) (*ktypes.OutcomeInauguration, error) {
	return c.env.Inauguration(c.ctx)
}

func (c *Regional) InaugurationSegmented(_ context.Context, id uint) (*ktypes.OutcomeInaugurationSegment, error) {
	return c.env.InaugurationSegmented(c.ctx, id)
}

func (c *Regional) Ledger(_ context.Context, altitude *int64) (*ktypes.OutcomeLedger, error) {
	return c.env.Ledger(c.ctx, altitude)
}

func (c *Regional) LedgerViaDigest(_ context.Context, digest []byte) (*ktypes.OutcomeLedger, error) {
	return c.env.LedgerViaDigest(c.ctx, digest)
}

func (c *Regional) LedgerOutcomes(_ context.Context, altitude *int64) (*ktypes.OutcomeLedgerOutcomes, error) {
	return c.env.LedgerOutcomes(c.ctx, altitude)
}

func (c *Regional) Heading(_ context.Context, altitude *int64) (*ktypes.OutcomeHeadline, error) {
	return c.env.Heading(c.ctx, altitude)
}

func (c *Regional) HeadingViaDigest(_ context.Context, digest octets.HexadecimalOctets) (*ktypes.OutcomeHeadline, error) {
	return c.env.HeadingViaDigest(c.ctx, digest)
}

func (c *Regional) Endorse(_ context.Context, altitude *int64) (*ktypes.OutcomeEndorse, error) {
	return c.env.Endorse(c.ctx, altitude)
}

func (c *Regional) Assessors(_ context.Context, altitude *int64, screen, everyScreen *int) (*ktypes.OutcomeAssessors, error) {
	return c.env.Assessors(c.ctx, altitude, screen, everyScreen)
}

func (c *Regional) Tx(_ context.Context, digest []byte, ascertain bool) (*ktypes.OutcomeTransfer, error) {
	return c.env.Tx(c.ctx, digest, ascertain)
}

func (c *Regional) TransferLookup(
	_ context.Context,
	inquire string,
	ascertain bool,
	screen,
	everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeTransferLookup, error) {
	return c.env.TransferLookup(c.ctx, inquire, ascertain, screen, everyScreen, sequenceVia)
}

func (c *Regional) LedgerLookup(
	_ context.Context,
	inquire string,
	screen, everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeLedgerLookup, error) {
	return c.env.LedgerLookup(c.ctx, inquire, screen, everyScreen, sequenceVia)
}

func (c *Regional) MulticastProof(_ context.Context, ev kinds.Proof) (*ktypes.OutcomeMulticastProof, error) {
	return c.env.MulticastProof(c.ctx, ev)
}

func (c *Regional) Listen(
	ctx context.Context,
	listener,
	inquire string,
	outputVolume ...int,
) (out <-chan ktypes.OutcomeIncident, err error) {
	q, err := tendermintinquire.New(inquire)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	outputCeiling := 1
	if len(outputVolume) > 0 {
		outputCeiling = outputVolume[0]
	}

	var sub kinds.Listening
	if outputCeiling > 0 {
		sub, err = c.IncidentChannel.Listen(ctx, listener, q, outputCeiling)
	} else {
		sub, err = c.ListenUncached(ctx, listener, q)
	}
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	resultant := make(chan ktypes.OutcomeIncident, outputCeiling)
	go c.incidentsProcedure(sub, listener, q, resultant)

	return resultant, nil
}

func (c *Regional) incidentsProcedure(
	sub kinds.Listening,
	listener string,
	q tendermintpubsub.Inquire,
	resultant chan<- ktypes.OutcomeIncident,
) {
	for {
		select {
		case msg := <-sub.Out():
			outcome := ktypes.OutcomeIncident{Inquire: q.Text(), Data: msg.Data(), Incidents: msg.Incidents()}
			if cap(resultant) == 0 {
				resultant <- outcome
			} else {
				select {
				case resultant <- outcome:
				default:
					c.Tracer.Failure("REDACTED", "REDACTED", outcome, "REDACTED", outcome.Inquire)
				}
			}
		case <-sub.Aborted():
			if sub.Err() == tendermintpubsub.FaultUnlistened {
				return
			}

			c.Tracer.Failure("REDACTED", "REDACTED", sub.Err(), "REDACTED", q.Text())
			sub = c.relisten(listener, q)
			if sub == nil { //
				return
			}
		case <-c.Exit():
			return
		}
	}
}

//
func (c *Regional) relisten(listener string, q tendermintpubsub.Inquire) kinds.Listening {
	endeavors := 0
	for {
		if !c.EqualsActive() {
			return nil
		}

		sub, err := c.IncidentChannel.Listen(context.Background(), listener, q)
		if err == nil {
			return sub
		}

		endeavors++
		time.Sleep((10 << uint(endeavors)) * time.Millisecond) //
	}
}

func (c *Regional) Unlisten(ctx context.Context, listener, inquire string) error {
	q, err := tendermintinquire.New(inquire)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return c.IncidentChannel.Unlisten(ctx, listener, q)
}

func (c *Regional) UnlistenEvery(ctx context.Context, listener string) error {
	return c.IncidentChannel.UnlistenEvery(ctx, listener)
}
