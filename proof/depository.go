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

	cometfaults "github.com/valkyrieworks/kinds/faults"

	dbm "github.com/valkyrieworks/-db"

	ringlist "github.com/valkyrieworks/utils/ringlist"
	"github.com/valkyrieworks/utils/log"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

const (
	rootKeyConfirmed = byte(0x00)
	rootKeyAwaiting   = byte(0x01)
)

//
type Depository struct {
	tracer log.Tracer

	proofDepot dbm.DB
	proofCatalog  *ringlist.CCatalog //
	proofVolume  uint32       //

	//
	statusStore sm.Depot
	//
	ledgerDepot LedgerDepot

	mtx sync.Mutex
	//
	status sm.Status
	//
	//
	//
	agreementBuffer []replicatedBallotCollection

	trimmingLevel int64
	trimmingTime   time.Time
}

//
//
func NewDepository(proofStore dbm.DB, statusStore sm.Depot, ledgerDepot LedgerDepot) (*Depository, error) {
	status, err := statusStore.Import()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	depository := &Depository{
		statusStore:         statusStore,
		ledgerDepot:      ledgerDepot,
		status:           status,
		tracer:          log.NewNoopTracer(),
		proofDepot:   proofStore,
		proofCatalog:    ringlist.New(),
		agreementBuffer: make([]replicatedBallotCollection, 0),
	}

	//
	//
	depository.trimmingLevel, depository.trimmingTime = depository.deleteLapsedAwaitingProof()
	evtCatalog, _, err := depository.catalogProof(rootKeyAwaiting, -1)
	if err != nil {
		return nil, err
	}
	atomic.StoreUint32(&depository.proofVolume, uint32(len(evtCatalog)))
	for _, ev := range evtCatalog {
		depository.proofCatalog.PropelRear(ev)
	}

	return depository, nil
}

//
func (eventpool *Depository) AwaitingProof(maximumOctets int64) ([]kinds.Proof, int64) {
	if eventpool.Volume() == 0 {
		return []kinds.Proof{}, 0
	}
	proof, volume, err := eventpool.catalogProof(rootKeyAwaiting, maximumOctets)
	if err != nil {
		eventpool.tracer.Fault("REDACTED", "REDACTED", err)
	}
	return proof, volume
}

//
//
//
//
//
//
//
func (eventpool *Depository) Modify(status sm.Status, ev kinds.ProofCatalog) {
	//
	if status.FinalLedgerLevel <= eventpool.status.FinalLedgerLevel {
		panic(fmt.Sprintf(
			"REDACTED",
			status.FinalLedgerLevel,
			eventpool.status.FinalLedgerLevel,
		))
	}
	eventpool.tracer.Diagnose("REDACTED", "REDACTED", status.FinalLedgerLevel,
		"REDACTED", status.FinalLedgerTime)

	//
	//
	eventpool.handleAgreementBuffer(status)
	//
	eventpool.modifyStatus(status)

	//
	eventpool.stampProofAsConfirmed(ev)

	//
	if eventpool.Volume() > 0 && status.FinalLedgerLevel > eventpool.trimmingLevel &&
		status.FinalLedgerTime.After(eventpool.trimmingTime) {
		eventpool.trimmingLevel, eventpool.trimmingTime = eventpool.deleteLapsedAwaitingProof()
	}
}

//
func (eventpool *Depository) AppendProof(ev kinds.Proof) error {
	eventpool.tracer.Details("REDACTED", "REDACTED", ev)

	//
	if eventpool.isAwaiting(ev) {
		eventpool.tracer.Details("REDACTED", "REDACTED", ev)
		return nil
	}

	//
	if eventpool.isConfirmed(ev) {
		//
		//
		eventpool.tracer.Details("REDACTED", "REDACTED", ev)
		return nil
	}

	//
	err := eventpool.validate(ev)
	if err != nil {
		return kinds.NewErrCorruptProof(ev, err)
	}

	//
	if err := eventpool.appendAwaitingProof(ev); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	eventpool.proofCatalog.PropelRear(ev)

	eventpool.tracer.Details("REDACTED", "REDACTED", ev)

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
func (eventpool *Depository) NotifyClashingBallots(ballotA, ballotBYTE *kinds.Ballot) {
	eventpool.mtx.Lock()
	defer eventpool.mtx.Unlock()
	eventpool.agreementBuffer = append(eventpool.agreementBuffer, replicatedBallotCollection{
		BallotA: ballotA,
		BallotBYTE: ballotBYTE,
	})
}

//
//
//
//
func (eventpool *Depository) InspectProof(evtCatalog kinds.ProofCatalog) error {
	digests := make([][]byte, len(evtCatalog))
	for idx, ev := range evtCatalog {

		_, isRapidEvt := ev.(*kinds.RapidCustomerAssaultProof)

		//
		//
		if isRapidEvt || !eventpool.isAwaiting(ev) {
			//
			if eventpool.isConfirmed(ev) {
				return &kinds.ErrCorruptProof{Proof: ev, Cause: errors.New("REDACTED")}
			}

			err := eventpool.validate(ev)
			if err != nil {
				return err
			}

			if err := eventpool.appendAwaitingProof(ev); err != nil {
				//
				//
				eventpool.tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", ev)
			}

			eventpool.tracer.Details("REDACTED", "REDACTED", ev)
		}

		//
		digests[idx] = ev.Digest()
		for i := idx - 1; i >= 0; i-- {
			if bytes.Equal(digests[i], digests[idx]) {
				return &kinds.ErrCorruptProof{Proof: ev, Cause: errors.New("REDACTED")}
			}
		}
	}

	return nil
}

//
func (eventpool *Depository) ProofHead() *ringlist.CComponent {
	return eventpool.proofCatalog.Head()
}

//
func (eventpool *Depository) ProofWaitChan() <-chan struct{} {
	return eventpool.proofCatalog.WaitChan()
}

//
func (eventpool *Depository) AssignTracer(l log.Tracer) {
	eventpool.tracer = l
}

//
func (eventpool *Depository) Volume() uint32 {
	return atomic.LoadUint32(&eventpool.proofVolume)
}

//
func (eventpool *Depository) Status() sm.Status {
	eventpool.mtx.Lock()
	defer eventpool.mtx.Unlock()
	return eventpool.status
}

func (eventpool *Depository) End() error {
	return eventpool.proofDepot.End()
}

//
//
func (eventpool *Depository) isLapsed(level int64, moment time.Time) bool {
	var (
		options       = eventpool.Status().AgreementOptions.Proof
		eraPeriod  = eventpool.Status().FinalLedgerTime.Sub(moment)
		eraCountLedgers = eventpool.Status().FinalLedgerLevel - level
	)
	return eraCountLedgers > options.MaximumDurationCountLedgers &&
		eraPeriod > options.MaximumDurationPeriod
}

//
func (eventpool *Depository) isConfirmed(proof kinds.Proof) bool {
	key := keyConfirmed(proof)
	ok, err := eventpool.proofDepot.Has(key)
	if err != nil {
		eventpool.tracer.Fault("REDACTED", "REDACTED", err)
	}
	return ok
}

//
func (eventpool *Depository) isAwaiting(proof kinds.Proof) bool {
	key := keyAwaiting(proof)
	ok, err := eventpool.proofDepot.Has(key)
	if err != nil {
		eventpool.tracer.Fault("REDACTED", "REDACTED", err)
	}
	return ok
}

func (eventpool *Depository) appendAwaitingProof(ev kinds.Proof) error {
	evschema, err := kinds.ProofToSchema(ev)
	if err != nil {
		return cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
	}

	evtOctets, err := evschema.Serialize()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	key := keyAwaiting(ev)

	err = eventpool.proofDepot.Set(key, evtOctets)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	atomic.AddUint32(&eventpool.proofVolume, 1)
	return nil
}

func (eventpool *Depository) deleteAwaitingProof(proof kinds.Proof) {
	key := keyAwaiting(proof)
	if err := eventpool.proofDepot.Erase(key); err != nil {
		eventpool.tracer.Fault("REDACTED", "REDACTED", err)
	} else {
		atomic.AddUint32(&eventpool.proofVolume, ^uint32(0))
		eventpool.tracer.Diagnose("REDACTED", "REDACTED", proof)
	}
}

//
//
func (eventpool *Depository) stampProofAsConfirmed(proof kinds.ProofCatalog) {
	ledgerProofIndex := make(map[string]struct{}, len(proof))
	for _, ev := range proof {
		if eventpool.isAwaiting(ev) {
			eventpool.deleteAwaitingProof(ev)
			ledgerProofIndex[evtIndexKey(ev)] = struct{}{}
		}

		//
		//
		key := keyConfirmed(ev)

		h := gogotypes.Int64Value{Value: ev.Level()}
		evtOctets, err := proto.Marshal(&h)
		if err != nil {
			eventpool.tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", key)
			continue
		}

		if err := eventpool.proofDepot.Set(key, evtOctets); err != nil {
			eventpool.tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", key)
		}
	}

	//
	if len(ledgerProofIndex) != 0 {
		eventpool.deleteProofFromCatalog(ledgerProofIndex)
	}
}

//
//
func (eventpool *Depository) catalogProof(prefixKey byte, maximumOctets int64) ([]kinds.Proof, int64, error) {
	var (
		evtVolume    int64
		sumVolume int64
		proof  []kinds.Proof
		evtCatalog    engineproto.ProofCatalog //
	)

	iterator, err := dbm.RecursePrefix(eventpool.proofDepot, []byte{prefixKey})
	if err != nil {
		return nil, sumVolume, fmt.Errorf("REDACTED", err)
	}
	defer iterator.End()
	for ; iterator.Sound(); iterator.Following() {
		var evschema engineproto.Proof
		err := evschema.Unserialize(iterator.Item())
		if err != nil {
			return proof, sumVolume, err
		}
		evtCatalog.Proof = append(evtCatalog.Proof, evschema)
		evtVolume = int64(evtCatalog.Volume())
		if maximumOctets != -1 && evtVolume > maximumOctets {
			if err := iterator.Fault(); err != nil {
				return proof, sumVolume, err
			}
			return proof, sumVolume, nil
		}

		ev, err := kinds.ProofFromSchema(&evschema)
		if err != nil {
			return nil, sumVolume, err
		}

		sumVolume = evtVolume
		proof = append(proof, ev)
	}

	if err := iterator.Fault(); err != nil {
		return proof, sumVolume, err
	}
	return proof, sumVolume, nil
}

func (eventpool *Depository) deleteLapsedAwaitingProof() (int64, time.Time) {
	iterator, err := dbm.RecursePrefix(eventpool.proofDepot, []byte{rootKeyAwaiting})
	if err != nil {
		eventpool.tracer.Fault("REDACTED", "REDACTED", err)
		return eventpool.Status().FinalLedgerLevel, eventpool.Status().FinalLedgerTime
	}
	defer iterator.End()
	ledgerProofIndex := make(map[string]struct{})
	for ; iterator.Sound(); iterator.Following() {
		ev, err := octetsToEvt(iterator.Item())
		if err != nil {
			eventpool.tracer.Fault("REDACTED", "REDACTED", err)
			continue
		}
		if !eventpool.isLapsed(ev.Level(), ev.Time()) {
			if len(ledgerProofIndex) != 0 {
				eventpool.deleteProofFromCatalog(ledgerProofIndex)
			}

			//
			return ev.Level() + eventpool.Status().AgreementOptions.Proof.MaximumDurationCountLedgers + 1,
				ev.Time().Add(eventpool.Status().AgreementOptions.Proof.MaximumDurationPeriod).Add(time.Second)
		}
		eventpool.deleteAwaitingProof(ev)
		ledgerProofIndex[evtIndexKey(ev)] = struct{}{}
	}
	//
	if len(ledgerProofIndex) != 0 {
		eventpool.deleteProofFromCatalog(ledgerProofIndex)
	}
	return eventpool.Status().FinalLedgerLevel, eventpool.Status().FinalLedgerTime
}

func (eventpool *Depository) deleteProofFromCatalog(
	ledgerProofIndex map[string]struct{},
) {
	for e := eventpool.proofCatalog.Head(); e != nil; e = e.Following() {
		//
		ev := e.Item.(kinds.Proof)
		if _, ok := ledgerProofIndex[evtIndexKey(ev)]; ok {
			eventpool.proofCatalog.Delete(e)
			e.UnplugPrevious()
		}
	}
}

func (eventpool *Depository) modifyStatus(status sm.Status) {
	eventpool.mtx.Lock()
	defer eventpool.mtx.Unlock()
	eventpool.status = status
}

//
//
//
//
func (eventpool *Depository) handleAgreementBuffer(status sm.Status) {
	eventpool.mtx.Lock()
	defer eventpool.mtx.Unlock()
	for _, ballotCollection := range eventpool.agreementBuffer {

		//
		//
		var (
			dve *kinds.ReplicatedBallotProof
			err error
		)
		switch {
		case ballotCollection.BallotA.Level == status.FinalLedgerLevel:
			dve, err = kinds.NewReplicatedBallotProof(
				ballotCollection.BallotA,
				ballotCollection.BallotBYTE,
				status.FinalLedgerTime,
				status.FinalRatifiers,
			)

		case ballotCollection.BallotA.Level < status.FinalLedgerLevel:
			var valueCollection *kinds.RatifierAssign
			valueCollection, err = eventpool.statusStore.ImportRatifiers(ballotCollection.BallotA.Level)
			if err != nil {
				eventpool.tracer.Fault("REDACTED", "REDACTED",
					ballotCollection.BallotA.Level, "REDACTED", err,
				)
				continue
			}
			ledgerMeta := eventpool.ledgerDepot.ImportLedgerMeta(ballotCollection.BallotA.Level)
			if ledgerMeta == nil {
				eventpool.tracer.Fault("REDACTED", "REDACTED", ballotCollection.BallotA.Level)
				continue
			}
			dve, err = kinds.NewReplicatedBallotProof(
				ballotCollection.BallotA,
				ballotCollection.BallotBYTE,
				ledgerMeta.Heading.Time,
				valueCollection,
			)

		default:
			//
			//
			//
			eventpool.tracer.Fault("REDACTED",
				"REDACTED", ballotCollection.BallotA.Level,
				"REDACTED", status.FinalLedgerLevel)
			continue
		}
		if err != nil {
			eventpool.tracer.Fault("REDACTED", "REDACTED", err)
			continue
		}

		//
		if eventpool.isAwaiting(dve) {
			eventpool.tracer.Details("REDACTED", "REDACTED", dve)
			continue
		}

		//
		if eventpool.isConfirmed(dve) {
			eventpool.tracer.Details("REDACTED", "REDACTED", dve)
			continue
		}

		if err := eventpool.appendAwaitingProof(dve); err != nil {
			eventpool.tracer.Fault("REDACTED", "REDACTED", err)
			continue
		}

		eventpool.proofCatalog.PropelRear(dve)

		eventpool.tracer.Details("REDACTED", "REDACTED", dve)
	}
	//
	eventpool.agreementBuffer = make([]replicatedBallotCollection, 0)
}

type replicatedBallotCollection struct {
	BallotA *kinds.Ballot
	BallotBYTE *kinds.Ballot
}

func octetsToEvt(evtOctets []byte) (kinds.Proof, error) {
	var evschema engineproto.Proof
	err := evschema.Unserialize(evtOctets)
	if err != nil {
		return &kinds.ReplicatedBallotProof{}, err
	}

	return kinds.ProofFromSchema(&evschema)
}

func evtIndexKey(ev kinds.Proof) string {
	return string(ev.Digest())
}

//
func bE(h int64) string {
	return fmt.Sprintf("REDACTED", h)
}

func keyConfirmed(proof kinds.Proof) []byte {
	return append([]byte{rootKeyConfirmed}, keyExtension(proof)...)
}

func keyAwaiting(proof kinds.Proof) []byte {
	return append([]byte{rootKeyAwaiting}, keyExtension(proof)...)
}

func keyExtension(proof kinds.Proof) []byte {
	return []byte(fmt.Sprintf("REDACTED", bE(proof.Level()), proof.Digest()))
}
