package kinds

import (
	"context"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

const fallbackVolume = 0

type IncidentPipelineListener interface {
	Listen(ctx context.Context, listener string, inquire tendermintpubsub.Inquire, outputVolume ...int) (Listening, error)
	Unlisten(ctx context.Context, listener string, inquire tendermintpubsub.Inquire) error
	UnlistenEvery(ctx context.Context, listener string) error

	CountCustomers() int
	CountCustomerFeeds(customerUUID string) int
}

type Listening interface {
	Out() <-chan tendermintpubsub.Signal
	Aborted() <-chan struct{}
	Err() error
}

//
//
//
type IncidentChannel struct {
	facility.FoundationFacility
	broadcastlisten *tendermintpubsub.Daemon
}

//
func FreshIncidentPipeline() *IncidentChannel {
	return FreshIncidentPipelineUsingReserveVolume(fallbackVolume)
}

//
func FreshIncidentPipelineUsingReserveVolume(cap int) *IncidentChannel {
	//
	broadcastlisten := tendermintpubsub.FreshDaemon(tendermintpubsub.ReserveVolume(cap))
	b := &IncidentChannel{broadcastlisten: broadcastlisten}
	b.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", b)
	return b
}

func (b *IncidentChannel) AssignTracer(l log.Tracer) {
	b.FoundationFacility.AssignTracer(l)
	b.broadcastlisten.AssignTracer(l.Using("REDACTED", "REDACTED"))
}

func (b *IncidentChannel) UponInitiate() error {
	return b.broadcastlisten.Initiate()
}

func (b *IncidentChannel) UponHalt() {
	if err := b.broadcastlisten.Halt(); err != nil {
		b.broadcastlisten.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

func (b *IncidentChannel) CountCustomers() int {
	return b.broadcastlisten.CountCustomers()
}

func (b *IncidentChannel) CountCustomerFeeds(customerUUID string) int {
	return b.broadcastlisten.CountCustomerFeeds(customerUUID)
}

func (b *IncidentChannel) Listen(
	ctx context.Context,
	listener string,
	inquire tendermintpubsub.Inquire,
	outputVolume ...int,
) (Listening, error) {
	return b.broadcastlisten.Listen(ctx, listener, inquire, outputVolume...)
}

//
//
func (b *IncidentChannel) ListenUncached(
	ctx context.Context,
	listener string,
	inquire tendermintpubsub.Inquire,
) (Listening, error) {
	return b.broadcastlisten.ListenUncached(ctx, listener, inquire)
}

func (b *IncidentChannel) Unlisten(ctx context.Context, listener string, inquire tendermintpubsub.Inquire) error {
	return b.broadcastlisten.Unlisten(ctx, listener, inquire)
}

func (b *IncidentChannel) UnlistenEvery(ctx context.Context, listener string) error {
	return b.broadcastlisten.UnlistenEvery(ctx, listener)
}

func (b *IncidentChannel) Broadcast(incidentKind string, incidentData TEMPIncidentData) error {
	//
	ctx := context.Background()
	return b.broadcastlisten.BroadcastUsingIncidents(ctx, incidentData, map[string][]string{IncidentKindToken: {incidentKind}})
}

//
//
//
//
func (*IncidentChannel) certifyAlsoEncodeasstringIncidents(incidents []kinds.Incident) map[string][]string {
	outcome := make(map[string][]string)
	for _, incident := range incidents {
		if len(incident.Kind) == 0 {
			continue
		}
		heading := incident.Kind + "REDACTED"
		for _, property := range incident.Properties {
			if len(property.Key) == 0 {
				continue
			}

			complexLabel := heading + property.Key
			outcome[complexLabel] = append(outcome[complexLabel], property.Datum)
		}
	}

	return outcome
}

func (b *IncidentChannel) BroadcastIncidentFreshLedger(data IncidentDataFreshLedger) error {
	//
	ctx := context.Background()
	incidents := b.certifyAlsoEncodeasstringIncidents(data.OutcomeCulminateLedger.Incidents)

	//
	incidents[IncidentKindToken] = append(incidents[IncidentKindToken], IncidentFreshLedger)

	return b.broadcastlisten.BroadcastUsingIncidents(ctx, data, incidents)
}

func (b *IncidentChannel) BroadcastIncidentFreshLedgerIncidents(data IncidentDataFreshLedgerIncidents) error {
	//
	ctx := context.Background()

	incidents := b.certifyAlsoEncodeasstringIncidents(data.Incidents)

	//
	incidents[IncidentKindToken] = append(incidents[IncidentKindToken], IncidentFreshLedgerIncidents)

	return b.broadcastlisten.BroadcastUsingIncidents(ctx, data, incidents)
}

func (b *IncidentChannel) BroadcastIncidentFreshLedgerHeading(data IncidentDataFreshLedgerHeading) error {
	return b.Broadcast(IncidentFreshLedgerHeading, data)
}

func (b *IncidentChannel) BroadcastIncidentFreshProof(proof IncidentDataFreshProof) error {
	return b.Broadcast(IncidentFreshProof, proof)
}

func (b *IncidentChannel) BroadcastIncidentBallot(data IncidentDataBallot) error {
	return b.Broadcast(IncidentBallot, data)
}

func (b *IncidentChannel) BroadcastIncidentSoundLedger(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentSoundLedger, data)
}

//
//
//
func (b *IncidentChannel) BroadcastIncidentTransfer(data IncidentDataTransfer) error {
	//
	ctx := context.Background()

	incidents := b.certifyAlsoEncodeasstringIncidents(data.Outcome.Incidents)

	//
	incidents[IncidentKindToken] = append(incidents[IncidentKindToken], IncidentTransfer)
	incidents[TransferDigestToken] = append(incidents[TransferDigestToken], fmt.Sprintf("REDACTED", Tx(data.Tx).Digest()))
	incidents[TransferAltitudeToken] = append(incidents[TransferAltitudeToken], fmt.Sprintf("REDACTED", data.Altitude))

	return b.broadcastlisten.BroadcastUsingIncidents(ctx, data, incidents)
}

func (b *IncidentChannel) BroadcastIncidentFreshIterationPhase(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentFreshIterationPhase, data)
}

func (b *IncidentChannel) BroadcastIncidentDeadlineNominate(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentDeadlineNominate, data)
}

func (b *IncidentChannel) BroadcastIncidentDeadlinePause(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentDeadlinePause, data)
}

func (b *IncidentChannel) BroadcastIncidentFreshIteration(data IncidentDataFreshIteration) error {
	return b.Broadcast(IncidentFreshIteration, data)
}

func (b *IncidentChannel) BroadcastIncidentFinishNomination(data IncidentDataFinishNomination) error {
	return b.Broadcast(IncidentFinishedNomination, data)
}

func (b *IncidentChannel) BroadcastIncidentSpeck(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentSpeck, data)
}

func (b *IncidentChannel) BroadcastIncidentRelease(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentRelease, data)
}

func (b *IncidentChannel) BroadcastIncidentResecure(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentResecure, data)
}

func (b *IncidentChannel) BroadcastIncidentSecure(data IncidentDataIterationStatus) error {
	return b.Broadcast(IncidentSecure, data)
}

func (b *IncidentChannel) BroadcastIncidentAssessorAssignRevisions(data IncidentDataAssessorAssignRevisions) error {
	return b.Broadcast(IncidentAssessorAssignRevisions, data)
}

//
type NooperationIncidentPipeline struct{}

func (NooperationIncidentPipeline) Listen(
	context.Context,
	string,
	tendermintpubsub.Inquire,
	chan<- any,
) error {
	return nil
}

func (NooperationIncidentPipeline) Unlisten(context.Context, string, tendermintpubsub.Inquire) error {
	return nil
}

func (NooperationIncidentPipeline) UnlistenEvery(context.Context, string) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentFreshLedger(IncidentDataFreshLedger) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentFreshLedgerHeading(IncidentDataFreshLedgerHeading) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentFreshLedgerIncidents(IncidentDataFreshLedgerIncidents) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentFreshProof(IncidentDataFreshProof) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentBallot(IncidentDataBallot) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentTransfer(IncidentDataTransfer) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentFreshIterationPhase(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentDeadlineNominate(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentDeadlinePause(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentFreshIteration(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentFinishNomination(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentSpeck(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentRelease(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentResecure(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentSecure(IncidentDataIterationStatus) error {
	return nil
}

func (NooperationIncidentPipeline) BroadcastIncidentAssessorAssignRevisions(IncidentDataAssessorAssignRevisions) error {
	return nil
}
