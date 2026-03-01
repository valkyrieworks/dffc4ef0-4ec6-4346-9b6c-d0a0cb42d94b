package kinds

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	tendermintinquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
)

func VerifyIncidentPipelineBroadcastIncidentTransfer(t *testing.T) {
	incidentPipeline := FreshIncidentPipeline()
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	tx := Tx("REDACTED")
	outcome := iface.InvokeTransferOutcome{
		Data: []byte("REDACTED"),
		Incidents: []iface.Incident{
			{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED"}}},
		},
	}

	//
	inquire := fmt.Sprintf("REDACTED", tx.Digest())
	transUnder, err := incidentPipeline.Listen(context.Background(), "REDACTED", tendermintinquire.ShouldAssemble(inquire))
	require.NoError(t, err)

	complete := make(chan struct{})
	go func() {
		msg := <-transUnder.Out()
		edt := msg.Data().(IncidentDataTransfer)
		assert.Equal(t, int64(1), edt.Altitude)
		assert.Equal(t, uint32(0), edt.Ordinal)
		assert.EqualValues(t, tx, edt.Tx)
		assert.Equal(t, outcome, edt.Outcome)
		close(complete)
	}()

	err = incidentPipeline.BroadcastIncidentTransfer(IncidentDataTransfer{iface.TransferOutcome{
		Altitude: 1,
		Ordinal:  0,
		Tx:     tx,
		Outcome: outcome,
	}})
	assert.NoError(t, err)

	select {
	case <-complete:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyIncidentPipelineBroadcastIncidentFreshLedger(t *testing.T) {
	incidentPipeline := FreshIncidentPipeline()
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	ledger := CreateLedger(0, []Tx{}, nil, []Proof{})
	outcomeCulminateLedger := iface.ReplyCulminateLedger{
		Incidents: []iface.Incident{
			{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED"}}},
		},
	}

	//
	inquire := "REDACTED"
	ledgersUnder, err := incidentPipeline.Listen(context.Background(), "REDACTED", tendermintinquire.ShouldAssemble(inquire))
	require.NoError(t, err)

	complete := make(chan struct{})
	go func() {
		msg := <-ledgersUnder.Out()
		edt := msg.Data().(IncidentDataFreshLedger)
		assert.Equal(t, ledger, edt.Ledger)
		assert.Equal(t, outcomeCulminateLedger, edt.OutcomeCulminateLedger)
		close(complete)
	}()

	var ps *FragmentAssign
	ps, err = ledger.CreateFragmentAssign(LedgerFragmentExtentOctets)
	require.NoError(t, err)

	err = incidentPipeline.BroadcastIncidentFreshLedger(IncidentDataFreshLedger{
		Ledger: ledger,
		LedgerUUID: LedgerUUID{
			Digest:          ledger.Digest(),
			FragmentAssignHeading: ps.Heading(),
		},
		OutcomeCulminateLedger: outcomeCulminateLedger,
	})
	assert.NoError(t, err)

	select {
	case <-complete:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyIncidentPipelineBroadcastIncidentTransferReplicatedTokens(t *testing.T) {
	incidentPipeline := FreshIncidentPipeline()
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	tx := Tx("REDACTED")
	outcome := iface.InvokeTransferOutcome{
		Data: []byte("REDACTED"),
		Incidents: []iface.Incident{
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{Key: "REDACTED", Datum: "REDACTED"},
					{Key: "REDACTED", Datum: "REDACTED"},
					{Key: "REDACTED", Datum: "REDACTED"},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{Key: "REDACTED", Datum: "REDACTED"},
					{Key: "REDACTED", Datum: "REDACTED"},
					{Key: "REDACTED", Datum: "REDACTED"},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{Key: "REDACTED", Datum: "REDACTED"},
					{Key: "REDACTED", Datum: "REDACTED"},
					{Key: "REDACTED", Datum: "REDACTED"},
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
		sub, err := incidentPipeline.Listen(context.Background(), fmt.Sprintf("REDACTED", i), tendermintinquire.ShouldAssemble(tc.inquire))
		require.NoError(t, err)

		complete := make(chan struct{})

		go func() {
			select {
			case msg := <-sub.Out():
				data := msg.Data().(IncidentDataTransfer)
				assert.Equal(t, int64(1), data.Altitude)
				assert.Equal(t, uint32(0), data.Ordinal)
				assert.EqualValues(t, tx, data.Tx)
				assert.Equal(t, outcome, data.Outcome)
				close(complete)
			case <-time.After(1 * time.Second):
				return
			}
		}()

		err = incidentPipeline.BroadcastIncidentTransfer(IncidentDataTransfer{iface.TransferOutcome{
			Altitude: 1,
			Ordinal:  0,
			Tx:     tx,
			Outcome: outcome,
		}})
		assert.NoError(t, err)

		select {
		case <-complete:
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

func VerifyIncidentPipelineBroadcastIncidentFreshLedgerHeading(t *testing.T) {
	incidentPipeline := FreshIncidentPipeline()
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	ledger := CreateLedger(0, []Tx{}, nil, []Proof{})
	//
	inquire := "REDACTED"
	headingsUnder, err := incidentPipeline.Listen(context.Background(), "REDACTED", tendermintinquire.ShouldAssemble(inquire))
	require.NoError(t, err)

	complete := make(chan struct{})
	go func() {
		msg := <-headingsUnder.Out()
		edt := msg.Data().(IncidentDataFreshLedgerHeading)
		assert.Equal(t, ledger.Heading, edt.Heading)
		close(complete)
	}()

	err = incidentPipeline.BroadcastIncidentFreshLedgerHeading(IncidentDataFreshLedgerHeading{
		Heading: ledger.Heading,
	})
	assert.NoError(t, err)

	select {
	case <-complete:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyIncidentPipelineBroadcastIncidentFreshLedgerIncidents(t *testing.T) {
	incidentPipeline := FreshIncidentPipeline()
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	inquire := "REDACTED"
	headingsUnder, err := incidentPipeline.Listen(context.Background(), "REDACTED", tendermintinquire.ShouldAssemble(inquire))
	require.NoError(t, err)

	complete := make(chan struct{})
	go func() {
		msg := <-headingsUnder.Out()
		edt := msg.Data().(IncidentDataFreshLedgerIncidents)
		assert.Equal(t, int64(1), edt.Altitude)
		close(complete)
	}()

	err = incidentPipeline.BroadcastIncidentFreshLedgerIncidents(IncidentDataFreshLedgerIncidents{
		Altitude: 1,
		Incidents: []iface.Incident{{
			Kind: "REDACTED",
			Properties: []iface.IncidentProperty{{
				Key:   "REDACTED",
				Datum: "REDACTED",
			}},
		}},
	})
	assert.NoError(t, err)

	select {
	case <-complete:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyIncidentPipelineBroadcastIncidentFreshProof(t *testing.T) {
	incidentPipeline := FreshIncidentPipeline()
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	ev, err := FreshSimulateReplicatedBallotProof(1, time.Now(), "REDACTED")
	require.NoError(t, err)

	inquire := "REDACTED"
	occurenceUnder, err := incidentPipeline.Listen(context.Background(), "REDACTED", tendermintinquire.ShouldAssemble(inquire))
	require.NoError(t, err)

	complete := make(chan struct{})
	go func() {
		msg := <-occurenceUnder.Out()
		edt := msg.Data().(IncidentDataFreshProof)
		assert.Equal(t, ev, edt.Proof)
		assert.Equal(t, int64(4), edt.Altitude)
		close(complete)
	}()

	err = incidentPipeline.BroadcastIncidentFreshProof(IncidentDataFreshProof{
		Proof: ev,
		Altitude:   4,
	})
	assert.NoError(t, err)

	select {
	case <-complete:
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyIncidentPipelineBroadcast(t *testing.T) {
	incidentPipeline := FreshIncidentPipeline()
	err := incidentPipeline.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			t.Error(err)
		}
	})

	const countIncidentsAnticipated = 15

	sub, err := incidentPipeline.Listen(context.Background(), "REDACTED", tendermintinquire.All, countIncidentsAnticipated)
	require.NoError(t, err)

	complete := make(chan struct{})
	go func() {
		countIncidents := 0
		for range sub.Out() {
			countIncidents++
			if countIncidents >= countIncidentsAnticipated {
				close(complete)
				return
			}
		}
	}()

	err = incidentPipeline.Broadcast(IncidentFreshLedgerHeading, IncidentDataFreshLedgerHeading{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentFreshLedger(IncidentDataFreshLedger{Ledger: &Ledger{Heading: Heading{Altitude: 1}}})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentFreshLedgerHeading(IncidentDataFreshLedgerHeading{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentFreshLedgerIncidents(IncidentDataFreshLedgerIncidents{Altitude: 1})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentBallot(IncidentDataBallot{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentFreshIterationPhase(IncidentDataIterationStatus{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentDeadlineNominate(IncidentDataIterationStatus{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentDeadlinePause(IncidentDataIterationStatus{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentFreshIteration(IncidentDataFreshIteration{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentFinishNomination(IncidentDataFinishNomination{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentSpeck(IncidentDataIterationStatus{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentRelease(IncidentDataIterationStatus{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentResecure(IncidentDataIterationStatus{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentSecure(IncidentDataIterationStatus{})
	require.NoError(t, err)
	err = incidentPipeline.BroadcastIncidentAssessorAssignRevisions(IncidentDataAssessorAssignRevisions{})
	require.NoError(t, err)

	select {
	case <-complete:
	case <-time.After(1 * time.Second):
		t.Fatalf("REDACTED", countIncidentsAnticipated)
	}
}

func AssessmentIncidentPipeline(b *testing.B) {
	measurements := []struct {
		alias        string
		countCustomers  int
		arbitraryInquiries bool
		arbitraryIncidents  bool
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
		b.Run(bm.alias, func(b *testing.B) {
			assessmentIncidentPipeline(bm.countCustomers, bm.arbitraryInquiries, bm.arbitraryIncidents, b)
		})
	}
}

func assessmentIncidentPipeline(countCustomers int, arbitraryInquiries bool, arbitraryIncidents bool, b *testing.B) {
	//
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	incidentPipeline := FreshIncidentPipelineUsingReserveVolume(0) //
	err := incidentPipeline.Initiate()
	if err != nil {
		b.Error(err)
	}
	b.Cleanup(func() {
		if err := incidentPipeline.Halt(); err != nil {
			b.Error(err)
		}
	})

	ctx := context.Background()
	q := IncidentInquireFreshLedger

	for i := 0; i < countCustomers; i++ {
		if arbitraryInquiries {
			q = arbitraryInquire(rnd)
		}
		sub, err := incidentPipeline.Listen(ctx, fmt.Sprintf("REDACTED", i), q)
		if err != nil {
			b.Fatal(err)
		}
		go func() {
			for {
				select {
				case <-sub.Out():
				case <-sub.Aborted():
					return
				}
			}
		}()
	}

	incidentKind := IncidentFreshLedger

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if arbitraryIncidents {
			incidentKind = arbitraryIncident(rnd)
		}

		err := incidentPipeline.Broadcast(incidentKind, IncidentDataText("REDACTED"))
		if err != nil {
			b.Error(err)
		}
	}
}

var incidents = []string{
	IncidentFreshLedger,
	IncidentFreshLedgerHeading,
	IncidentFreshLedgerIncidents,
	IncidentFreshIteration,
	IncidentFreshIterationPhase,
	IncidentDeadlineNominate,
	IncidentFinishedNomination,
	IncidentSpeck,
	IncidentRelease,
	IncidentSecure,
	IncidentResecure,
	IncidentDeadlinePause,
	IncidentBallot,
}

func arbitraryIncident(r *rand.Rand) string {
	return incidents[r.Intn(len(incidents))]
}

var inquiries = []tendermintpubsub.Inquire{
	IncidentInquireFreshLedger,
	IncidentInquireFreshLedgerHeading,
	IncidentInquireFreshLedgerIncidents,
	IncidentInquireFreshIteration,
	IncidentInquireFreshIterationPhase,
	IncidentInquireDeadlineNominate,
	IncidentInquireFinishNomination,
	IncidentInquireSpeck,
	IncidentInquireRelease,
	IncidentInquireSecure,
	IncidentInquireResecure,
	IncidentInquireDeadlinePause,
	IncidentInquireBallot,
}

func arbitraryInquire(r *rand.Rand) tendermintpubsub.Inquire {
	return inquiries[r.Intn(len(inquiries))]
}
