package kinds

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/iface/kinds"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	cmtinquire "github.com/valkyrieworks/utils/broadcast/inquire"
)

func VerifyEventBusBroadcastEventTransfer(t *testing.T) {
	eventBus := NewEventBus()
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	tx := Tx("REDACTED")
	outcome := iface.InvokeTransferOutcome{
		Data: []byte("REDACTED"),
		Events: []iface.Event{
			{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED"}}},
		},
	}

	//
	inquire := fmt.Sprintf("REDACTED", tx.Digest())
	transSubtract, err := eventBus.Enrol(context.Background(), "REDACTED", cmtinquire.ShouldBuild(inquire))
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		msg := <-transSubtract.Out()
		edt := msg.Data().(EventDataTransfer)
		assert.Equal(t, int64(1), edt.Level)
		assert.Equal(t, uint32(0), edt.Ordinal)
		assert.EqualValues(t, tx, edt.Tx)
		assert.Equal(t, outcome, edt.Outcome)
		close(done)
	}()

	err = eventBus.BroadcastEventTransfer(EventDataTransfer{iface.TransOutcome{
		Level: 1,
		Ordinal:  0,
		Tx:     tx,
		Outcome: outcome,
	}})
	assert.NoError(t, err)

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyEventBusBroadcastEventNewLedger(t *testing.T) {
	eventBus := NewEventBus()
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	ledger := CreateLedger(0, []Tx{}, nil, []Proof{})
	outcomeCompleteLedger := iface.ReplyCompleteLedger{
		Events: []iface.Event{
			{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED"}}},
		},
	}

	//
	inquire := "REDACTED"
	ledgersSubtract, err := eventBus.Enrol(context.Background(), "REDACTED", cmtinquire.ShouldBuild(inquire))
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		msg := <-ledgersSubtract.Out()
		edt := msg.Data().(EventDataNewLedger)
		assert.Equal(t, ledger, edt.Ledger)
		assert.Equal(t, outcomeCompleteLedger, edt.OutcomeCompleteLedger)
		close(done)
	}()

	var ps *SegmentCollection
	ps, err = ledger.CreateSegmentAssign(LedgerSegmentVolumeOctets)
	require.NoError(t, err)

	err = eventBus.BroadcastEventNewLedger(EventDataNewLedger{
		Ledger: ledger,
		LedgerUID: LedgerUID{
			Digest:          ledger.Digest(),
			SegmentAssignHeading: ps.Heading(),
		},
		OutcomeCompleteLedger: outcomeCompleteLedger,
	})
	assert.NoError(t, err)

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyEventBusBroadcastEventTransferReplicatedKeys(t *testing.T) {
	eventBus := NewEventBus()
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	tx := Tx("REDACTED")
	outcome := iface.InvokeTransferOutcome{
		Data: []byte("REDACTED"),
		Events: []iface.Event{
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{Key: "REDACTED", Item: "REDACTED"},
					{Key: "REDACTED", Item: "REDACTED"},
					{Key: "REDACTED", Item: "REDACTED"},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{Key: "REDACTED", Item: "REDACTED"},
					{Key: "REDACTED", Item: "REDACTED"},
					{Key: "REDACTED", Item: "REDACTED"},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{Key: "REDACTED", Item: "REDACTED"},
					{Key: "REDACTED", Item: "REDACTED"},
					{Key: "REDACTED", Item: "REDACTED"},
				},
			},
		},
	}

	verifyScenarios := []struct {
		inquire         string
		anticipateOutcomes bool
	}{
		{
			"REDACTED",
			false,
		},
		{
			"REDACTED",
			true,
		},
		{
			"REDACTED",
			true,
		},
		{
			"REDACTED",
			true,
		},
		{
			"REDACTED",
			false,
		},
	}

	for i, tc := range verifyScenarios {
		sub, err := eventBus.Enrol(context.Background(), fmt.Sprintf("REDACTED", i), cmtinquire.ShouldBuild(tc.inquire))
		require.NoError(t, err)

		done := make(chan struct{})

		go func() {
			select {
			case msg := <-sub.Out():
				data := msg.Data().(EventDataTransfer)
				assert.Equal(t, int64(1), data.Level)
				assert.Equal(t, uint32(0), data.Ordinal)
				assert.EqualValues(t, tx, data.Tx)
				assert.Equal(t, outcome, data.Outcome)
				close(done)
			case <-time.After(1 * time.Second):
				return
			}
		}()

		err = eventBus.BroadcastEventTransfer(EventDataTransfer{iface.TransOutcome{
			Level: 1,
			Ordinal:  0,
			Tx:     tx,
			Outcome: outcome,
		}})
		assert.NoError(t, err)

		select {
		case <-done:
			if !tc.anticipateOutcomes {
				require.Fail(t, "REDACTED")
			}
		case <-time.After(1 * time.Second):
			if tc.anticipateOutcomes {
				require.Fail(t, "REDACTED")
			}
		}
	}
}

func VerifyEventBusBroadcastEventNewLedgerHeading(t *testing.T) {
	eventBus := NewEventBus()
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	ledger := CreateLedger(0, []Tx{}, nil, []Proof{})
	//
	inquire := "REDACTED"
	headingsSubtract, err := eventBus.Enrol(context.Background(), "REDACTED", cmtinquire.ShouldBuild(inquire))
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		msg := <-headingsSubtract.Out()
		edt := msg.Data().(EventDataNewLedgerHeading)
		assert.Equal(t, ledger.Heading, edt.Heading)
		close(done)
	}()

	err = eventBus.BroadcastEventNewLedgerHeading(EventDataNewLedgerHeading{
		Heading: ledger.Heading,
	})
	assert.NoError(t, err)

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyEventBusBroadcastEventNewLedgerEvents(t *testing.T) {
	eventBus := NewEventBus()
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	inquire := "REDACTED"
	headingsSubtract, err := eventBus.Enrol(context.Background(), "REDACTED", cmtinquire.ShouldBuild(inquire))
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		msg := <-headingsSubtract.Out()
		edt := msg.Data().(EventDataNewLedgerEvents)
		assert.Equal(t, int64(1), edt.Level)
		close(done)
	}()

	err = eventBus.BroadcastEventNewLedgerEvents(EventDataNewLedgerEvents{
		Level: 1,
		Events: []iface.Event{{
			Kind: "REDACTED",
			Properties: []iface.EventProperty{{
				Key:   "REDACTED",
				Item: "REDACTED",
			}},
		}},
	})
	assert.NoError(t, err)

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyEventBusBroadcastEventNewProof(t *testing.T) {
	eventBus := NewEventBus()
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	ev, err := NewEmulateReplicatedBallotProof(1, time.Now(), "REDACTED")
	require.NoError(t, err)

	inquire := "REDACTED"
	evtSubtract, err := eventBus.Enrol(context.Background(), "REDACTED", cmtinquire.ShouldBuild(inquire))
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		msg := <-evtSubtract.Out()
		edt := msg.Data().(EventDataNewProof)
		assert.Equal(t, ev, edt.Proof)
		assert.Equal(t, int64(4), edt.Level)
		close(done)
	}()

	err = eventBus.BroadcastEventNewProof(EventDataNewProof{
		Proof: ev,
		Level:   4,
	})
	assert.NoError(t, err)

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyEventBusBroadcast(t *testing.T) {
	eventBus := NewEventBus()
	err := eventBus.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			t.Error(err)
		}
	})

	const countEventsAnticipated = 15

	sub, err := eventBus.Enrol(context.Background(), "REDACTED", cmtinquire.All, countEventsAnticipated)
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		countEvents := 0
		for range sub.Out() {
			countEvents++
			if countEvents >= countEventsAnticipated {
				close(done)
				return
			}
		}
	}()

	err = eventBus.Broadcast(EventNewLedgerHeading, EventDataNewLedgerHeading{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventNewLedger(EventDataNewLedger{Ledger: &Ledger{Heading: Heading{Level: 1}}})
	require.NoError(t, err)
	err = eventBus.BroadcastEventNewLedgerHeading(EventDataNewLedgerHeading{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventNewLedgerEvents(EventDataNewLedgerEvents{Level: 1})
	require.NoError(t, err)
	err = eventBus.BroadcastEventBallot(EventDataBallot{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventNewEpochPhase(EventDataDurationStatus{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventDeadlineNominate(EventDataDurationStatus{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventDeadlineWait(EventDataDurationStatus{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventNewEpoch(EventDataNewEpoch{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventFinishedNomination(EventDataFinishedNomination{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventPolka(EventDataDurationStatus{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventRelease(EventDataDurationStatus{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventResecure(EventDataDurationStatus{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventSecure(EventDataDurationStatus{})
	require.NoError(t, err)
	err = eventBus.BroadcastEventRatifierCollectionRefreshes(EventDataRatifierCollectionRefreshes{})
	require.NoError(t, err)

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatalf("REDACTED", countEventsAnticipated)
	}
}

func CriterionEventBus(b *testing.B) {
	measurements := []struct {
		label        string
		countAgents  int
		randomInquiries bool
		randomEvents  bool
	}{
		{"REDACTED", 10, false, false},
		{"REDACTED", 100, false, false},
		{"REDACTED", 1000, false, false},

		{"REDACTED", 10, true, false},
		{"REDACTED", 100, true, false},
		{"REDACTED", 1000, true, false},

		{"REDACTED", 10, true, true},
		{"REDACTED", 100, true, true},
		{"REDACTED", 1000, true, true},

		{"REDACTED", 10, false, true},
		{"REDACTED", 100, false, true},
		{"REDACTED", 1000, false, true},
	}

	for _, bm := range measurements {
		b.Run(bm.label, func(b *testing.B) {
			criterionEventBus(bm.countAgents, bm.randomInquiries, bm.randomEvents, b)
		})
	}
}

func criterionEventBus(countAgents int, randomInquiries bool, randomEvents bool, b *testing.B) {
	//
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	eventBus := NewEventBusWithBufferVolume(0) //
	err := eventBus.Begin()
	if err != nil {
		b.Error(err)
	}
	b.Cleanup(func() {
		if err := eventBus.Halt(); err != nil {
			b.Error(err)
		}
	})

	ctx := context.Background()
	q := EventInquireNewLedger

	for i := 0; i < countAgents; i++ {
		if randomInquiries {
			q = randomInquire(rnd)
		}
		sub, err := eventBus.Enrol(ctx, fmt.Sprintf("REDACTED", i), q)
		if err != nil {
			b.Fatal(err)
		}
		go func() {
			for {
				select {
				case <-sub.Out():
				case <-sub.Revoked():
					return
				}
			}
		}()
	}

	eventKind := EventNewLedger

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if randomEvents {
			eventKind = randomEvent(rnd)
		}

		err := eventBus.Broadcast(eventKind, EventDataString("REDACTED"))
		if err != nil {
			b.Error(err)
		}
	}
}

var events = []string{
	EventNewLedger,
	EventNewLedgerHeading,
	EventNewLedgerEvents,
	EventNewEpoch,
	EventNewDurationPhase,
	EventDeadlineNominate,
	EventFinishedNomination,
	EventPolka,
	EventRelease,
	EventSecure,
	EventResecure,
	EventDeadlineWait,
	EventBallot,
}

func randomEvent(r *rand.Rand) string {
	return events[r.Intn(len(events))]
}

var inquiries = []cometbroadcast.Inquire{
	EventInquireNewLedger,
	EventInquireNewLedgerHeading,
	EventInquireNewLedgerEvents,
	EventInquireNewEpoch,
	EventInquireNewEpochPhase,
	EventInquireDeadlineNominate,
	EventInquireFinishedNomination,
	EventInquirePolka,
	EventInquireRelease,
	EventInquireSecure,
	EventInquireResecure,
	EventInquireDeadlineWait,
	EventInquireBallot,
}

func randomInquire(r *rand.Rand) cometbroadcast.Inquire {
	return inquiries[r.Intn(len(inquiries))]
}
