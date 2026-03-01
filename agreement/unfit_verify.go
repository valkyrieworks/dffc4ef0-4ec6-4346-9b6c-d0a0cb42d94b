package agreement

import (
	"testing"
	"time"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	strongmindcons "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/agreement"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//

//
//
func VerifyHandlerUnfitPreendorse(t *testing.T) {
	N := 4
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(true), freshTokvalDepot,
		func(c *cfg.Settings) {
			c.Agreement.DeadlineNominate = 3000 * time.Millisecond
			c.Agreement.DeadlinePreballot = 1000 * time.Millisecond
			c.Agreement.DeadlinePreendorse = 1000 * time.Millisecond
		})
	defer sanitize()

	for i := 0; i < N; i++ {
		metronome := FreshDeadlineMetronome()
		metronome.AssignTracer(css[i].Tracer)
		css[i].AssignDeadlineMetronome(metronome)

	}

	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, N)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	//
	byzantineItemOffset := N - 1
	byzantineItem := css[byzantineItemOffset]
	byzantineReader := engines[byzantineItemOffset]

	//
	//
	byzantineItem.mtx.Lock()
	pv := byzantineItem.privateAssessor
	byzantineItem.performPreballot = func(int64, int32) {
		unfitPerformPreballotMethod(t, byzantineItem, byzantineReader.Router, pv)
	}
	byzantineItem.mtx.Unlock()

	//
	//
	for i := 0; i < 10; i++ {
		deadlinePauseCluster(N, func(j int) {
			<-ledgersSubscriptions[j].Out()
		})
	}
}

func unfitPerformPreballotMethod(t *testing.T, cs *Status, sw p2p.Router, pv kinds.PrivateAssessor) {
	//
	//
	//
	//
	go func() {
		cs.mtx.Lock()
		defer cs.mtx.Unlock()
		cs.privateAssessor = pv
		publicToken, err := cs.privateAssessor.ObtainPublicToken()
		if err != nil {
			panic(err)
		}
		location := publicToken.Location()
		itemOrdinal, _ := cs.Assessors.ObtainViaLocation(location)

		//
		ledgerDigest := octets.HexadecimalOctets(commitrand.Octets(32))
		preendorse := &kinds.Ballot{
			AssessorLocation: location,
			AssessorOrdinal:   itemOrdinal,
			Altitude:           cs.Altitude,
			Iteration:            cs.Iteration,
			Timestamp:        cs.ballotMoment(),
			Kind:             commitchema.PreendorseKind,
			LedgerUUID: kinds.LedgerUUID{
				Digest:          ledgerDigest,
				FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 1, Digest: commitrand.Octets(32)},
			},
		}
		p := preendorse.TowardSchema()
		err = cs.privateAssessor.AttestBallot(cs.status.SuccessionUUID, p)
		if err != nil {
			t.Error(err)
		}
		preendorse.Notation = p.Notation
		preendorse.AdditionNotation = p.AdditionNotation
		cs.privateAssessor = nil //

		nodes := sw.Nodes().Duplicate()
		for _, node := range nodes {
			cs.Tracer.Details("REDACTED", "REDACTED", ledgerDigest, "REDACTED", node)
			node.Transmit(p2p.Wrapper{
				Signal:   &strongmindcons.Ballot{Ballot: preendorse.TowardSchema()},
				ConduitUUID: BallotConduit,
			})
		}
	}()
}
