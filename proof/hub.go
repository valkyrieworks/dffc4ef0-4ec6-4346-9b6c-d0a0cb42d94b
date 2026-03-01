package proof

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"

	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	linkedlist "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/linkedlist"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	foundationTokenRatified = byte(0x00)
	foundationTokenAwaiting   = byte(0x01)
)

//
type Hub struct {
	tracer log.Tracer

	proofDepot dbm.DB
	proofCatalog  *linkedlist.CNCatalog //
	proofExtent  uint32       //

	//
	statusDatastore sm.Depot
	//
	ledgerDepot LedgerDepot

	mtx sync.Mutex
	//
	status sm.Status
	//
	//
	//
	agreementReserve []replicatedBallotAssign

	thinningAltitude int64
	thinningMoment   time.Time
}

//
//
func FreshHub(proofDatastore dbm.DB, statusDatastore sm.Depot, ledgerDepot LedgerDepot) (*Hub, error) {
	status, err := statusDatastore.Fetch()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	hub := &Hub{
		statusDatastore:         statusDatastore,
		ledgerDepot:      ledgerDepot,
		status:           status,
		tracer:          log.FreshNooperationTracer(),
		proofDepot:   proofDatastore,
		proofCatalog:    linkedlist.New(),
		agreementReserve: make([]replicatedBallotAssign, 0),
	}

	//
	//
	hub.thinningAltitude, hub.thinningMoment = hub.discardLapsedAwaitingProof()
	occurenceCatalog, _, err := hub.catalogProof(foundationTokenAwaiting, -1)
	if err != nil {
		return nil, err
	}
	atomic.StoreUint32(&hub.proofExtent, uint32(len(occurenceCatalog)))
	for _, ev := range occurenceCatalog {
		hub.proofCatalog.PropelRear(ev)
	}

	return hub, nil
}

//
func (incidentpool *Hub) AwaitingProof(maximumOctets int64) ([]kinds.Proof, int64) {
	if incidentpool.Extent() == 0 {
		return []kinds.Proof{}, 0
	}
	proof, extent, err := incidentpool.catalogProof(foundationTokenAwaiting, maximumOctets)
	if err != nil {
		incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
	}
	return proof, extent
}

//
//
//
//
//
//
//
func (incidentpool *Hub) Revise(status sm.Status, ev kinds.ProofCatalog) {
	//
	if status.FinalLedgerAltitude <= incidentpool.status.FinalLedgerAltitude {
		panic(fmt.Sprintf(
			"REDACTED",
			status.FinalLedgerAltitude,
			incidentpool.status.FinalLedgerAltitude,
		))
	}
	incidentpool.tracer.Diagnose("REDACTED", "REDACTED", status.FinalLedgerAltitude,
		"REDACTED", status.FinalLedgerMoment)

	//
	//
	incidentpool.handleAgreementReserve(status)
	//
	incidentpool.reviseStatus(status)

	//
	incidentpool.labelProofLikeRatified(ev)

	//
	if incidentpool.Extent() > 0 && status.FinalLedgerAltitude > incidentpool.thinningAltitude &&
		status.FinalLedgerMoment.After(incidentpool.thinningMoment) {
		incidentpool.thinningAltitude, incidentpool.thinningMoment = incidentpool.discardLapsedAwaitingProof()
	}
}

//
func (incidentpool *Hub) AppendProof(ev kinds.Proof) error {
	incidentpool.tracer.Details("REDACTED", "REDACTED", ev)

	//
	if incidentpool.equalsAwaiting(ev) {
		incidentpool.tracer.Details("REDACTED", "REDACTED", ev)
		return nil
	}

	//
	if incidentpool.equalsRatified(ev) {
		//
		//
		incidentpool.tracer.Details("REDACTED", "REDACTED", ev)
		return nil
	}

	//
	err := incidentpool.validate(ev)
	if err != nil {
		return kinds.FreshFaultUnfitProof(ev, err)
	}

	//
	if err := incidentpool.appendAwaitingProof(ev); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	incidentpool.proofCatalog.PropelRear(ev)

	incidentpool.tracer.Details("REDACTED", "REDACTED", ev)

	return nil
}

//
//
//
//
//
//
//
//
//
func (incidentpool *Hub) DiscloseDiscordantBallots(ballotAN, ballotBYTE *kinds.Ballot) {
	incidentpool.mtx.Lock()
	defer incidentpool.mtx.Unlock()
	incidentpool.agreementReserve = append(incidentpool.agreementReserve, replicatedBallotAssign{
		BallotAN: ballotAN,
		BallotBYTE: ballotBYTE,
	})
}

//
//
//
//
func (incidentpool *Hub) InspectProof(occurenceCatalog kinds.ProofCatalog) error {
	digests := make([][]byte, len(occurenceCatalog))
	for idx, ev := range occurenceCatalog {

		_, equalsAgileOccurence := ev.(*kinds.AgileCustomerOnslaughtProof)

		//
		//
		if equalsAgileOccurence || !incidentpool.equalsAwaiting(ev) {
			//
			if incidentpool.equalsRatified(ev) {
				return &kinds.FaultUnfitProof{Proof: ev, Rationale: errors.New("REDACTED")}
			}

			err := incidentpool.validate(ev)
			if err != nil {
				return err
			}

			if err := incidentpool.appendAwaitingProof(ev); err != nil {
				//
				//
				incidentpool.tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", ev)
			}

			incidentpool.tracer.Details("REDACTED", "REDACTED", ev)
		}

		//
		digests[idx] = ev.Digest()
		for i := idx - 1; i >= 0; i-- {
			if bytes.Equal(digests[i], digests[idx]) {
				return &kinds.FaultUnfitProof{Proof: ev, Rationale: errors.New("REDACTED")}
			}
		}
	}

	return nil
}

//
func (incidentpool *Hub) ProofLeading() *linkedlist.CNComponent {
	return incidentpool.proofCatalog.Leading()
}

//
func (incidentpool *Hub) ProofPauseChnl() <-chan struct{} {
	return incidentpool.proofCatalog.PauseChnl()
}

//
func (incidentpool *Hub) AssignTracer(l log.Tracer) {
	incidentpool.tracer = l
}

//
func (incidentpool *Hub) Extent() uint32 {
	return atomic.LoadUint32(&incidentpool.proofExtent)
}

//
func (incidentpool *Hub) Status() sm.Status {
	incidentpool.mtx.Lock()
	defer incidentpool.mtx.Unlock()
	return incidentpool.status
}

func (incidentpool *Hub) Shutdown() error {
	return incidentpool.proofDepot.Shutdown()
}

//
//
func (incidentpool *Hub) equalsLapsed(altitude int64, moment time.Time) bool {
	var (
		parameters       = incidentpool.Status().AgreementSettings.Proof
		lifespanInterval  = incidentpool.Status().FinalLedgerMoment.Sub(moment)
		lifespanCountLedgers = incidentpool.Status().FinalLedgerAltitude - altitude
	)
	return lifespanCountLedgers > parameters.MaximumLifespanCountLedgers &&
		lifespanInterval > parameters.MaximumLifespanInterval
}

//
func (incidentpool *Hub) equalsRatified(proof kinds.Proof) bool {
	key := tokenRatified(proof)
	ok, err := incidentpool.proofDepot.Has(key)
	if err != nil {
		incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
	}
	return ok
}

//
func (incidentpool *Hub) equalsAwaiting(proof kinds.Proof) bool {
	key := tokenAwaiting(proof)
	ok, err := incidentpool.proofDepot.Has(key)
	if err != nil {
		incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
	}
	return ok
}

func (incidentpool *Hub) appendAwaitingProof(ev kinds.Proof) error {
	proofschema, err := kinds.ProofTowardSchema(ev)
	if err != nil {
		return strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
	}

	occurenceOctets, err := proofschema.Serialize()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	key := tokenAwaiting(ev)

	err = incidentpool.proofDepot.Set(key, occurenceOctets)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	atomic.AddUint32(&incidentpool.proofExtent, 1)
	return nil
}

func (incidentpool *Hub) discardAwaitingProof(proof kinds.Proof) {
	key := tokenAwaiting(proof)
	if err := incidentpool.proofDepot.Erase(key); err != nil {
		incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
	} else {
		atomic.AddUint32(&incidentpool.proofExtent, ^uint32(0))
		incidentpool.tracer.Diagnose("REDACTED", "REDACTED", proof)
	}
}

//
//
func (incidentpool *Hub) labelProofLikeRatified(proof kinds.ProofCatalog) {
	ledgerProofIndex := make(map[string]struct{}, len(proof))
	for _, ev := range proof {
		if incidentpool.equalsAwaiting(ev) {
			incidentpool.discardAwaitingProof(ev)
			ledgerProofIndex[occurenceIndexToken(ev)] = struct{}{}
		}

		//
		//
		key := tokenRatified(ev)

		h := gogotypes.Int64Value{Value: ev.Altitude()}
		occurenceOctets, err := proto.Marshal(&h)
		if err != nil {
			incidentpool.tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", key)
			continue
		}

		if err := incidentpool.proofDepot.Set(key, occurenceOctets); err != nil {
			incidentpool.tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", key)
		}
	}

	//
	if len(ledgerProofIndex) != 0 {
		incidentpool.discardProofOriginatingCatalog(ledgerProofIndex)
	}
}

//
//
func (incidentpool *Hub) catalogProof(headingToken byte, maximumOctets int64) ([]kinds.Proof, int64, error) {
	var (
		occurenceExtent    int64
		sumExtent int64
		proof  []kinds.Proof
		occurenceCatalog    commitchema.ProofCatalog //
	)

	count, err := dbm.TraverseHeading(incidentpool.proofDepot, []byte{headingToken})
	if err != nil {
		return nil, sumExtent, fmt.Errorf("REDACTED", err)
	}
	defer count.Shutdown()
	for ; count.Sound(); count.Following() {
		var proofschema commitchema.Proof
		err := proofschema.Decode(count.Datum())
		if err != nil {
			return proof, sumExtent, err
		}
		occurenceCatalog.Proof = append(occurenceCatalog.Proof, proofschema)
		occurenceExtent = int64(occurenceCatalog.Extent())
		if maximumOctets != -1 && occurenceExtent > maximumOctets {
			if err := count.Failure(); err != nil {
				return proof, sumExtent, err
			}
			return proof, sumExtent, nil
		}

		ev, err := kinds.ProofOriginatingSchema(&proofschema)
		if err != nil {
			return nil, sumExtent, err
		}

		sumExtent = occurenceExtent
		proof = append(proof, ev)
	}

	if err := count.Failure(); err != nil {
		return proof, sumExtent, err
	}
	return proof, sumExtent, nil
}

func (incidentpool *Hub) discardLapsedAwaitingProof() (int64, time.Time) {
	count, err := dbm.TraverseHeading(incidentpool.proofDepot, []byte{foundationTokenAwaiting})
	if err != nil {
		incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
		return incidentpool.Status().FinalLedgerAltitude, incidentpool.Status().FinalLedgerMoment
	}
	defer count.Shutdown()
	ledgerProofIndex := make(map[string]struct{})
	for ; count.Sound(); count.Following() {
		ev, err := octetsTowardOccurence(count.Datum())
		if err != nil {
			incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
			continue
		}
		if !incidentpool.equalsLapsed(ev.Altitude(), ev.Moment()) {
			if len(ledgerProofIndex) != 0 {
				incidentpool.discardProofOriginatingCatalog(ledgerProofIndex)
			}

			//
			return ev.Altitude() + incidentpool.Status().AgreementSettings.Proof.MaximumLifespanCountLedgers + 1,
				ev.Moment().Add(incidentpool.Status().AgreementSettings.Proof.MaximumLifespanInterval).Add(time.Second)
		}
		incidentpool.discardAwaitingProof(ev)
		ledgerProofIndex[occurenceIndexToken(ev)] = struct{}{}
	}
	//
	if len(ledgerProofIndex) != 0 {
		incidentpool.discardProofOriginatingCatalog(ledgerProofIndex)
	}
	return incidentpool.Status().FinalLedgerAltitude, incidentpool.Status().FinalLedgerMoment
}

func (incidentpool *Hub) discardProofOriginatingCatalog(
	ledgerProofIndex map[string]struct{},
) {
	for e := incidentpool.proofCatalog.Leading(); e != nil; e = e.Following() {
		//
		ev := e.Datum.(kinds.Proof)
		if _, ok := ledgerProofIndex[occurenceIndexToken(ev)]; ok {
			incidentpool.proofCatalog.Discard(e)
			e.UncouplePrevious()
		}
	}
}

func (incidentpool *Hub) reviseStatus(status sm.Status) {
	incidentpool.mtx.Lock()
	defer incidentpool.mtx.Unlock()
	incidentpool.status = status
}

//
//
//
//
func (incidentpool *Hub) handleAgreementReserve(status sm.Status) {
	incidentpool.mtx.Lock()
	defer incidentpool.mtx.Unlock()
	for _, ballotAssign := range incidentpool.agreementReserve {

		//
		//
		var (
			dve *kinds.ReplicatedBallotProof
			err error
		)
		switch {
		case ballotAssign.BallotAN.Altitude == status.FinalLedgerAltitude:
			dve, err = kinds.FreshReplicatedBallotProof(
				ballotAssign.BallotAN,
				ballotAssign.BallotBYTE,
				status.FinalLedgerMoment,
				status.FinalAssessors,
			)

		case ballotAssign.BallotAN.Altitude < status.FinalLedgerAltitude:
			var itemAssign *kinds.AssessorAssign
			itemAssign, err = incidentpool.statusDatastore.FetchAssessors(ballotAssign.BallotAN.Altitude)
			if err != nil {
				incidentpool.tracer.Failure("REDACTED", "REDACTED",
					ballotAssign.BallotAN.Altitude, "REDACTED", err,
				)
				continue
			}
			ledgerSummary := incidentpool.ledgerDepot.FetchLedgerSummary(ballotAssign.BallotAN.Altitude)
			if ledgerSummary == nil {
				incidentpool.tracer.Failure("REDACTED", "REDACTED", ballotAssign.BallotAN.Altitude)
				continue
			}
			dve, err = kinds.FreshReplicatedBallotProof(
				ballotAssign.BallotAN,
				ballotAssign.BallotBYTE,
				ledgerSummary.Heading.Moment,
				itemAssign,
			)

		default:
			//
			//
			//
			incidentpool.tracer.Failure("REDACTED",
				"REDACTED", ballotAssign.BallotAN.Altitude,
				"REDACTED", status.FinalLedgerAltitude)
			continue
		}
		if err != nil {
			incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
			continue
		}

		//
		if incidentpool.equalsAwaiting(dve) {
			incidentpool.tracer.Details("REDACTED", "REDACTED", dve)
			continue
		}

		//
		if incidentpool.equalsRatified(dve) {
			incidentpool.tracer.Details("REDACTED", "REDACTED", dve)
			continue
		}

		if err := incidentpool.appendAwaitingProof(dve); err != nil {
			incidentpool.tracer.Failure("REDACTED", "REDACTED", err)
			continue
		}

		incidentpool.proofCatalog.PropelRear(dve)

		incidentpool.tracer.Details("REDACTED", "REDACTED", dve)
	}
	//
	incidentpool.agreementReserve = make([]replicatedBallotAssign, 0)
}

type replicatedBallotAssign struct {
	BallotAN *kinds.Ballot
	BallotBYTE *kinds.Ballot
}

func octetsTowardOccurence(occurenceOctets []byte) (kinds.Proof, error) {
	var proofschema commitchema.Proof
	err := proofschema.Decode(occurenceOctets)
	if err != nil {
		return &kinds.ReplicatedBallotProof{}, err
	}

	return kinds.ProofOriginatingSchema(&proofschema)
}

func occurenceIndexToken(ev kinds.Proof) string {
	return string(ev.Digest())
}

//
func bE(h int64) string {
	return fmt.Sprintf("REDACTED", h)
}

func tokenRatified(proof kinds.Proof) []byte {
	return append([]byte{foundationTokenRatified}, tokenEnding(proof)...)
}

func tokenAwaiting(proof kinds.Proof) []byte {
	return append([]byte{foundationTokenAwaiting}, tokenEnding(proof)...)
}

func tokenEnding(proof kinds.Proof) []byte {
	return []byte(fmt.Sprintf("REDACTED", bE(proof.Altitude()), proof.Digest()))
}
