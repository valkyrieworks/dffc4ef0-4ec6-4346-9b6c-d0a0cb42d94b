package agreement

import (
	"testing"
	"time"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	cometconnect "github.com/valkyrieworks/schema/consensuscore/agreement"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//

//
//
func VerifyHandlerCorruptPreendorse(t *testing.T) {
	N := 4
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(true), newObjectDepot,
		func(c *cfg.Settings) {
			c.Agreement.DeadlineNominate = 3000 * time.Millisecond
			c.Agreement.DeadlinePreballot = 1000 * time.Millisecond
			c.Agreement.DeadlinePreendorse = 1000 * time.Millisecond
		})
	defer sanitize()

	for i := 0; i < N; i++ {
		timer := NewDeadlineTimer()
		timer.AssignTracer(css[i].Tracer)
		css[i].AssignDeadlineTimer(timer)

	}

	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, N)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	//
	faultyValueOrdinal := N - 1
	faultyValue := css[faultyValueOrdinal]
	faultyReader := handlers[faultyValueOrdinal]

	//
	//
	faultyValue.mtx.Lock()
	pv := faultyValue.privateRatifier
	faultyValue.executePreballot = func(int64, int32) {
		corruptExecutePreballotFunction(t, faultyValue, faultyReader.Router, pv)
	}
	faultyValue.mtx.Unlock()

	//
	//
	for i := 0; i < 10; i++ {
		deadlineWaitCluster(N, func(j int) {
			<-ledgersEnrollments[j].Out()
		})
	}
}

func corruptExecutePreballotFunction(t *testing.T, cs *Status, sw p2p.Router, pv kinds.PrivateRatifier) {
	//
	//
	//
	//
	go func() {
		cs.mtx.Lock()
		defer cs.mtx.Unlock()
		cs.privateRatifier = pv
		publicKey, err := cs.privateRatifier.FetchPublicKey()
		if err != nil {
			panic(err)
		}
		address := publicKey.Location()
		valueOrdinal, _ := cs.Ratifiers.FetchByLocation(address)

		//
		ledgerDigest := octets.HexOctets(engineseed.Octets(32))
		preendorse := &kinds.Ballot{
			RatifierLocation: address,
			RatifierOrdinal:   valueOrdinal,
			Level:           cs.Level,
			Cycle:            cs.Cycle,
			Timestamp:        cs.ballotTime(),
			Kind:             engineproto.PreendorseKind,
			LedgerUID: kinds.LedgerUID{
				Digest:          ledgerDigest,
				SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 1, Digest: engineseed.Octets(32)},
			},
		}
		p := preendorse.ToSchema()
		err = cs.privateRatifier.AttestBallot(cs.status.LedgerUID, p)
		if err != nil {
			t.Error(err)
		}
		preendorse.Autograph = p.Autograph
		preendorse.AdditionAutograph = p.AdditionAutograph
		cs.privateRatifier = nil //

		nodes := sw.Nodes().Clone()
		for _, node := range nodes {
			cs.Tracer.Details("REDACTED", "REDACTED", ledgerDigest, "REDACTED", node)
			node.Transmit(p2p.Packet{
				Signal:   &cometconnect.Ballot{Ballot: preendorse.ToSchema()},
				StreamUID: BallotStream,
			})
		}
	}()
}
