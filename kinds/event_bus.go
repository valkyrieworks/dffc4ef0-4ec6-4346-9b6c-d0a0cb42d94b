package kinds

import (
	"context"
	"fmt"

	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/log"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	"github.com/valkyrieworks/utils/daemon"
)

const standardVolume = 0

type EventBusEnrollee interface {
	Enrol(ctx context.Context, enrollee string, inquire cometbroadcast.Inquire, outVolume ...int) (Enrollment, error)
	Deenroll(ctx context.Context, enrollee string, inquire cometbroadcast.Inquire) error
	DeenrollAll(ctx context.Context, enrollee string) error

	CountAgents() int
	CountCustomerRegistrations(customerUID string) int
}

type Enrollment interface {
	Out() <-chan cometbroadcast.Signal
	Revoked() <-chan struct{}
	Err() error
}

//
//
//
type EventBus struct {
	daemon.RootDaemon
	broadcast *cometbroadcast.Host
}

//
func NewEventBus() *EventBus {
	return NewEventBusWithBufferVolume(standardVolume)
}

//
func NewEventBusWithBufferVolume(cap int) *EventBus {
	//
	broadcast := cometbroadcast.NewHost(cometbroadcast.BufferVolume(cap))
	b := &EventBus{broadcast: broadcast}
	b.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", b)
	return b
}

func (b *EventBus) AssignTracer(l log.Tracer) {
	b.RootDaemon.AssignTracer(l)
	b.broadcast.AssignTracer(l.With("REDACTED", "REDACTED"))
}

func (b *EventBus) OnBegin() error {
	return b.broadcast.Begin()
}

func (b *EventBus) OnHalt() {
	if err := b.broadcast.Halt(); err != nil {
		b.broadcast.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

func (b *EventBus) CountAgents() int {
	return b.broadcast.CountAgents()
}

func (b *EventBus) CountCustomerRegistrations(customerUID string) int {
	return b.broadcast.CountCustomerRegistrations(customerUID)
}

func (b *EventBus) Enrol(
	ctx context.Context,
	enrollee string,
	inquire cometbroadcast.Inquire,
	outVolume ...int,
) (Enrollment, error) {
	return b.broadcast.Enrol(ctx, enrollee, inquire, outVolume...)
}

//
//
func (b *EventBus) EnrolUnbuffered(
	ctx context.Context,
	enrollee string,
	inquire cometbroadcast.Inquire,
) (Enrollment, error) {
	return b.broadcast.EnrolUnbuffered(ctx, enrollee, inquire)
}

func (b *EventBus) Deenroll(ctx context.Context, enrollee string, inquire cometbroadcast.Inquire) error {
	return b.broadcast.Deenroll(ctx, enrollee, inquire)
}

func (b *EventBus) DeenrollAll(ctx context.Context, enrollee string) error {
	return b.broadcast.DeenrollAll(ctx, enrollee)
}

func (b *EventBus) Broadcast(eventKind string, eventData TMEventData) error {
	//
	ctx := context.Background()
	return b.broadcast.BroadcastWithEvents(ctx, eventData, map[string][]string{EventKindKey: {eventKind}})
}

//
//
//
//
func (*EventBus) certifyAndEncodejsonEvents(events []kinds.Event) map[string][]string {
	outcome := make(map[string][]string)
	for _, event := range events {
		if len(event.Kind) == 0 {
			continue
		}
		prefix := event.Kind + "REDACTED"
		for _, property := range event.Properties {
			if len(property.Key) == 0 {
				continue
			}

			compoundLabel := prefix + property.Key
			outcome[compoundLabel] = append(outcome[compoundLabel], property.Item)
		}
	}

	return outcome
}

func (b *EventBus) BroadcastEventNewLedger(data EventDataNewLedger) error {
	//
	ctx := context.Background()
	events := b.certifyAndEncodejsonEvents(data.OutcomeCompleteLedger.Events)

	//
	events[EventKindKey] = append(events[EventKindKey], EventNewLedger)

	return b.broadcast.BroadcastWithEvents(ctx, data, events)
}

func (b *EventBus) BroadcastEventNewLedgerEvents(data EventDataNewLedgerEvents) error {
	//
	ctx := context.Background()

	events := b.certifyAndEncodejsonEvents(data.Events)

	//
	events[EventKindKey] = append(events[EventKindKey], EventNewLedgerEvents)

	return b.broadcast.BroadcastWithEvents(ctx, data, events)
}

func (b *EventBus) BroadcastEventNewLedgerHeading(data EventDataNewLedgerHeading) error {
	return b.Broadcast(EventNewLedgerHeading, data)
}

func (b *EventBus) BroadcastEventNewProof(proof EventDataNewProof) error {
	return b.Broadcast(EventNewProof, proof)
}

func (b *EventBus) BroadcastEventBallot(data EventDataBallot) error {
	return b.Broadcast(EventBallot, data)
}

func (b *EventBus) BroadcastEventSoundLedger(data EventDataDurationStatus) error {
	return b.Broadcast(EventSoundLedger, data)
}

//
//
//
func (b *EventBus) BroadcastEventTransfer(data EventDataTransfer) error {
	//
	ctx := context.Background()

	events := b.certifyAndEncodejsonEvents(data.Outcome.Events)

	//
	events[EventKindKey] = append(events[EventKindKey], EventTransfer)
	events[TransferDigestKey] = append(events[TransferDigestKey], fmt.Sprintf("REDACTED", Tx(data.Tx).Digest()))
	events[TransferLevelKey] = append(events[TransferLevelKey], fmt.Sprintf("REDACTED", data.Level))

	return b.broadcast.BroadcastWithEvents(ctx, data, events)
}

func (b *EventBus) BroadcastEventNewEpochPhase(data EventDataDurationStatus) error {
	return b.Broadcast(EventNewDurationPhase, data)
}

func (b *EventBus) BroadcastEventDeadlineNominate(data EventDataDurationStatus) error {
	return b.Broadcast(EventDeadlineNominate, data)
}

func (b *EventBus) BroadcastEventDeadlineWait(data EventDataDurationStatus) error {
	return b.Broadcast(EventDeadlineWait, data)
}

func (b *EventBus) BroadcastEventNewEpoch(data EventDataNewEpoch) error {
	return b.Broadcast(EventNewEpoch, data)
}

func (b *EventBus) BroadcastEventFinishedNomination(data EventDataFinishedNomination) error {
	return b.Broadcast(EventFinishedNomination, data)
}

func (b *EventBus) BroadcastEventPolka(data EventDataDurationStatus) error {
	return b.Broadcast(EventPolka, data)
}

func (b *EventBus) BroadcastEventRelease(data EventDataDurationStatus) error {
	return b.Broadcast(EventRelease, data)
}

func (b *EventBus) BroadcastEventResecure(data EventDataDurationStatus) error {
	return b.Broadcast(EventResecure, data)
}

func (b *EventBus) BroadcastEventSecure(data EventDataDurationStatus) error {
	return b.Broadcast(EventSecure, data)
}

func (b *EventBus) BroadcastEventRatifierCollectionRefreshes(data EventDataRatifierCollectionRefreshes) error {
	return b.Broadcast(EventRatifierCollectionRefreshes, data)
}

//
type NoopEventBus struct{}

func (NoopEventBus) Enrol(
	context.Context,
	string,
	cometbroadcast.Inquire,
	chan<- any,
) error {
	return nil
}

func (NoopEventBus) Deenroll(context.Context, string, cometbroadcast.Inquire) error {
	return nil
}

func (NoopEventBus) DeenrollAll(context.Context, string) error {
	return nil
}

func (NoopEventBus) BroadcastEventNewLedger(EventDataNewLedger) error {
	return nil
}

func (NoopEventBus) BroadcastEventNewLedgerHeading(EventDataNewLedgerHeading) error {
	return nil
}

func (NoopEventBus) BroadcastEventNewLedgerEvents(EventDataNewLedgerEvents) error {
	return nil
}

func (NoopEventBus) BroadcastEventNewProof(EventDataNewProof) error {
	return nil
}

func (NoopEventBus) BroadcastEventBallot(EventDataBallot) error {
	return nil
}

func (NoopEventBus) BroadcastEventTransfer(EventDataTransfer) error {
	return nil
}

func (NoopEventBus) BroadcastEventNewEpochPhase(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventDeadlineNominate(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventDeadlineWait(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventNewEpoch(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventFinishedNomination(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventPolka(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventRelease(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventResecure(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventSecure(EventDataDurationStatus) error {
	return nil
}

func (NoopEventBus) BroadcastEventRatifierCollectionRefreshes(EventDataRatifierCollectionRefreshes) error {
	return nil
}
