package depot

import (
	"errors"
	"fmt"
	"strconv"

	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
	"github.com/cosmos/gogoproto/proto"
	lru "github.com/hashicorp/golang-lru/v2"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	commitstore "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/depot"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
const maximumLedgerFragmentsTowardCluster = 10

/**
.

:
k
t
s

s
g
)

.

s
.
*/
type LedgerDepot struct {
	db dbm.DB

	//
	//
	//
	//
	//
	//
	//
	//
	//
	mtx    commitchronize.ReadwriteExclusion
	foundation   int64
	altitude int64

	observedEndorseStash          *lru.Cache[int64, *kinds.Endorse]
	ledgerEndorseStash         *lru.Cache[int64, *kinds.Endorse]
	ledgerExpandedEndorseStash *lru.Cache[int64, *kinds.ExpandedEndorse]
}

//
//
func FreshLedgerDepot(db dbm.DB) *LedgerDepot {
	bs := FetchLedgerDepotStatus(db)
	byteDepot := &LedgerDepot{
		foundation:   bs.Foundation,
		altitude: bs.Altitude,
		db:     db,
	}
	byteDepot.appendStashes()
	return byteDepot
}

func (bs *LedgerDepot) appendStashes() {
	var err error
	//
	bs.ledgerEndorseStash, err = lru.New[int64, *kinds.Endorse](100)
	if err != nil {
		panic(err)
	}
	bs.ledgerExpandedEndorseStash, err = lru.New[int64, *kinds.ExpandedEndorse](100)
	if err != nil {
		panic(err)
	}
	bs.observedEndorseStash, err = lru.New[int64, *kinds.Endorse](100)
	if err != nil {
		panic(err)
	}
}

func (bs *LedgerDepot) EqualsBlank() bool {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	return bs.foundation == 0 && bs.altitude == 0
}

//
func (bs *LedgerDepot) Foundation() int64 {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	return bs.foundation
}

//
func (bs *LedgerDepot) Altitude() int64 {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	return bs.altitude
}

//
func (bs *LedgerDepot) Extent() int64 {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	if bs.altitude == 0 {
		return 0
	}
	return bs.altitude - bs.foundation + 1
}

//
func (bs *LedgerDepot) FetchFoundationSummary() *kinds.LedgerSummary {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	if bs.foundation == 0 {
		return nil
	}
	return bs.FetchLedgerSummary(bs.foundation)
}

//
//
func (bs *LedgerDepot) FetchLedger(altitude int64) *kinds.Ledger {
	ledgerSummary := bs.FetchLedgerSummary(altitude)
	if ledgerSummary == nil {
		return nil
	}

	pbb := new(commitchema.Ledger)
	buf := []byte{}
	for i := 0; i < int(ledgerSummary.LedgerUUID.FragmentAssignHeading.Sum); i++ {
		fragment := bs.FetchLedgerFragment(altitude, i)
		//
		//
		if fragment == nil {
			return nil
		}
		buf = append(buf, fragment.Octets...)
	}
	err := proto.Unmarshal(buf, pbb)
	if err != nil {
		//
		//
		panic(fmt.Sprintf("REDACTED", err))
	}

	ledger, err := kinds.LedgerOriginatingSchema(pbb)
	if err != nil {
		panic(strongminderrors.FaultSignalOriginatingSchema{SignalAlias: "REDACTED", Err: err})
	}

	return ledger
}

//
//
//
func (bs *LedgerDepot) FetchLedgerViaDigest(digest []byte) *kinds.Ledger {
	bz, err := bs.db.Get(reckonLedgerDigestToken(digest))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}

	s := string(bz)
	altitude, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", s, err))
	}
	return bs.FetchLedger(altitude)
}

//
//
//
func (bs *LedgerDepot) FetchLedgerFragment(altitude int64, ordinal int) *kinds.Fragment {
	protosection := new(commitchema.Fragment)

	bz, err := bs.db.Get(reckonLedgerFragmentToken(altitude, ordinal))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}

	err = proto.Unmarshal(bz, protosection)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	fragment, err := kinds.FragmentOriginatingSchema(protosection)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	return fragment
}

//
//
func (bs *LedgerDepot) FetchLedgerSummary(altitude int64) *kinds.LedgerSummary {
	protoassessment := new(commitchema.LedgerSummary)
	bz, err := bs.db.Get(reckonLedgerSummaryToken(altitude))
	if err != nil {
		panic(err)
	}

	if len(bz) == 0 {
		return nil
	}

	err = proto.Unmarshal(bz, protoassessment)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	ledgerSummary, err := kinds.LedgerSummaryOriginatingReliableSchema(protoassessment)
	if err != nil {
		panic(strongminderrors.FaultSignalOriginatingSchema{SignalAlias: "REDACTED", Err: err})
	}

	return ledgerSummary
}

//
//
func (bs *LedgerDepot) FetchLedgerSummaryViaDigest(digest []byte) *kinds.LedgerSummary {
	bz, err := bs.db.Get(reckonLedgerDigestToken(digest))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}

	s := string(bz)
	altitude, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", s, err))
	}
	return bs.FetchLedgerSummary(altitude)
}

//
//
//
//
func (bs *LedgerDepot) FetchLedgerEndorse(altitude int64) *kinds.Endorse {
	xchange, ok := bs.ledgerEndorseStash.Get(altitude)
	if ok {
		return xchange.Replicate()
	}
	pbc := new(commitchema.Endorse)
	bz, err := bs.db.Get(reckonLedgerEndorseToken(altitude))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}
	err = proto.Unmarshal(bz, pbc)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	endorse, err := kinds.EndorseOriginatingSchema(pbc)
	if err != nil {
		panic(strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err})
	}
	bs.ledgerEndorseStash.Add(altitude, endorse)
	return endorse.Replicate()
}

//
//
//
func (bs *LedgerDepot) FetchLedgerExpandedEndorse(altitude int64) *kinds.ExpandedEndorse {
	xchange, ok := bs.ledgerExpandedEndorseStash.Get(altitude)
	if ok {
		return xchange.Replicate()
	}
	protoencode := new(commitchema.ExpandedEndorse)
	bz, err := bs.db.Get(reckonAddnEndorseToken(altitude))
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	if len(bz) == 0 {
		return nil
	}
	err = proto.Unmarshal(bz, protoencode)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	addnEndorse, err := kinds.ExpandedEndorseOriginatingSchema(protoencode)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	bs.ledgerExpandedEndorseStash.Add(altitude, addnEndorse)
	return addnEndorse.Replicate()
}

//
//
//
func (bs *LedgerDepot) FetchObservedEndorse(altitude int64) *kinds.Endorse {
	xchange, ok := bs.observedEndorseStash.Get(altitude)
	if ok {
		return xchange.Replicate()
	}
	pbc := new(commitchema.Endorse)
	bz, err := bs.db.Get(reckonObservedEndorseToken(altitude))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}
	err = proto.Unmarshal(bz, pbc)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	endorse, err := kinds.EndorseOriginatingSchema(pbc)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	bs.observedEndorseStash.Add(altitude, endorse)
	return endorse.Replicate()
}

//
func (bs *LedgerDepot) TrimLedgers(altitude int64, status sm.Status) (uint64, int64, error) {
	if altitude <= 0 {
		return 0, -1, fmt.Errorf("REDACTED")
	}
	bs.mtx.RLock()
	if altitude > bs.altitude {
		bs.mtx.RUnlock()
		return 0, -1, fmt.Errorf("REDACTED", bs.altitude)
	}
	foundation := bs.foundation
	bs.mtx.RUnlock()
	if altitude < foundation {
		return 0, -1, fmt.Errorf("REDACTED",
			altitude, foundation)
	}

	trimmed := uint64(0)
	cluster := bs.db.FreshCluster()
	defer cluster.Shutdown()
	purge := func(cluster dbm.Cluster, foundation int64) error {
		//
		//
		bs.mtx.Lock()
		defer cluster.Shutdown()
		defer bs.mtx.Unlock()
		bs.foundation = foundation
		return bs.persistStatusAlsoPersistDatastore(cluster, "REDACTED")
	}

	proofMark := altitude
	for h := foundation; h < altitude; h++ {

		summary := bs.FetchLedgerSummary(h)
		if summary == nil { //
			continue
		}

		//
		//

		if proofMark == altitude && !proof.EqualsProofLapsed(status.FinalLedgerAltitude, status.FinalLedgerMoment, h, summary.Heading.Moment, status.AgreementSettings.Proof) {
			proofMark = h
		}

		//
		if h < proofMark {
			if err := cluster.Erase(reckonLedgerSummaryToken(h)); err != nil {
				return 0, -1, err
			}
		}
		if err := cluster.Erase(reckonLedgerDigestToken(summary.LedgerUUID.Digest)); err != nil {
			return 0, -1, err
		}
		//
		if h < proofMark {
			if err := cluster.Erase(reckonLedgerEndorseToken(h)); err != nil {
				return 0, -1, err
			}
		}
		if err := cluster.Erase(reckonObservedEndorseToken(h)); err != nil {
			return 0, -1, err
		}

		if h < proofMark {
			if err := cluster.Erase(reckonAddnEndorseToken(h)); err != nil {
				return 0, -1, err
			}
			bs.ledgerExpandedEndorseStash.Remove(h)
		}

		for p := 0; p < int(summary.LedgerUUID.FragmentAssignHeading.Sum); p++ {
			if err := cluster.Erase(reckonLedgerFragmentToken(h, p)); err != nil {
				return 0, -1, err
			}
		}
		trimmed++

		//
		if trimmed%1000 == 0 && trimmed > 0 {
			err := purge(cluster, h)
			if err != nil {
				return 0, -1, err
			}
			cluster = bs.db.FreshCluster()
			defer cluster.Shutdown()
		}
	}

	err := purge(cluster, altitude)
	if err != nil {
		return 0, -1, err
	}
	return trimmed, proofMark, nil
}

//
//
//
//
//
//
//
func (bs *LedgerDepot) PersistLedger(ledger *kinds.Ledger, ledgerFragments *kinds.FragmentAssign, observedEndorse *kinds.Endorse) {
	if ledger == nil {
		panic("REDACTED")
	}

	cluster := bs.db.FreshCluster()
	defer cluster.Shutdown()

	if err := bs.persistLedgerTowardCluster(ledger, ledgerFragments, observedEndorse, cluster); err != nil {
		panic(err)
	}

	bs.mtx.Lock()
	defer bs.mtx.Unlock()
	bs.altitude = ledger.Altitude
	if bs.foundation == 0 {
		bs.foundation = ledger.Altitude
	}

	//
	err := bs.persistStatusAlsoPersistDatastore(cluster, "REDACTED")
	if err != nil {
		panic(err)
	}
}

//
//
//
//
//
func (bs *LedgerDepot) PersistLedgerUsingExpandedEndorse(ledger *kinds.Ledger, ledgerFragments *kinds.FragmentAssign, observedExpandedEndorse *kinds.ExpandedEndorse) {
	if ledger == nil {
		panic("REDACTED")
	}
	if err := observedExpandedEndorse.AssureAdditions(true); err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	cluster := bs.db.FreshCluster()
	defer cluster.Shutdown()

	if err := bs.persistLedgerTowardCluster(ledger, ledgerFragments, observedExpandedEndorse.TowardEndorse(), cluster); err != nil {
		panic(err)
	}
	altitude := ledger.Altitude

	protoencode := observedExpandedEndorse.TowardSchema()
	addnEndorseOctets := shouldEncode(protoencode)
	if err := cluster.Set(reckonAddnEndorseToken(altitude), addnEndorseOctets); err != nil {
		panic(err)
	}

	bs.mtx.Lock()
	defer bs.mtx.Unlock()
	bs.altitude = altitude
	if bs.foundation == 0 {
		bs.foundation = altitude
	}

	//
	err := bs.persistStatusAlsoPersistDatastore(cluster, "REDACTED")
	if err != nil {
		panic(err)
	}
}

func (bs *LedgerDepot) persistLedgerTowardCluster(
	ledger *kinds.Ledger,
	ledgerFragments *kinds.FragmentAssign,
	observedEndorse *kinds.Endorse,
	cluster dbm.Cluster,
) error {
	if ledger == nil {
		panic("REDACTED")
	}

	altitude := ledger.Altitude
	digest := ledger.Digest()

	if g, w := altitude, bs.Altitude()+1; bs.Foundation() > 0 && g != w {
		return fmt.Errorf("REDACTED", w, g)
	}
	if !ledgerFragments.EqualsFinish() {
		return errors.New("REDACTED")
	}
	if altitude != observedEndorse.Altitude {
		return fmt.Errorf("REDACTED", altitude, observedEndorse.Altitude)
	}

	//
	//
	persistLedgerFragmentsTowardCluster := ledgerFragments.Tally() <= maximumLedgerFragmentsTowardCluster

	//
	//
	//
	//
	for i := 0; i < int(ledgerFragments.Sum()); i++ {
		fragment := ledgerFragments.ObtainFragment(i)
		bs.persistLedgerFragment(altitude, i, fragment, cluster, persistLedgerFragmentsTowardCluster)
	}

	//
	ledgerSummary := kinds.FreshLedgerSummary(ledger, ledgerFragments)
	pbm := ledgerSummary.TowardSchema()
	if pbm == nil {
		return errors.New("REDACTED")
	}
	summaryOctets := shouldEncode(pbm)
	if err := cluster.Set(reckonLedgerSummaryToken(altitude), summaryOctets); err != nil {
		return err
	}
	if err := cluster.Set(reckonLedgerDigestToken(digest), []byte(fmt.Sprintf("REDACTED", altitude))); err != nil {
		return err
	}

	//
	pbc := ledger.FinalEndorse.TowardSchema()
	ledgerEndorseOctets := shouldEncode(pbc)
	if err := cluster.Set(reckonLedgerEndorseToken(altitude-1), ledgerEndorseOctets); err != nil {
		return err
	}

	//
	//
	protoscope := observedEndorse.TowardSchema()
	observedEndorseOctets := shouldEncode(protoscope)
	if err := cluster.Set(reckonObservedEndorseToken(altitude), observedEndorseOctets); err != nil {
		return err
	}

	return nil
}

func (bs *LedgerDepot) persistLedgerFragment(altitude int64, ordinal int, fragment *kinds.Fragment, cluster dbm.Cluster, persistLedgerFragmentsTowardCluster bool) {
	pbp, err := fragment.TowardSchema()
	if err != nil {
		panic(strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err})
	}
	fragmentOctets := shouldEncode(pbp)
	if persistLedgerFragmentsTowardCluster {
		err = cluster.Set(reckonLedgerFragmentToken(altitude, ordinal), fragmentOctets)
	} else {
		err = bs.db.Set(reckonLedgerFragmentToken(altitude, ordinal), fragmentOctets)
	}
	if err != nil {
		panic(err)
	}
}

//
func (bs *LedgerDepot) persistStatusAlsoPersistDatastore(cluster dbm.Cluster, faultSignal string) error {
	bss := commitstore.LedgerDepotStatus{
		Foundation:   bs.foundation,
		Altitude: bs.altitude,
	}
	PersistLedgerDepotStatusCluster(&bss, cluster)

	err := cluster.PersistChronize()
	if err != nil {
		return fmt.Errorf("REDACTED",
			faultSignal, bs.foundation, bs.altitude, err)
	}
	return nil
}

//
func (bs *LedgerDepot) PersistObservedEndorse(altitude int64, observedEndorse *kinds.Endorse) error {
	pbc := observedEndorse.TowardSchema()
	observedEndorseOctets, err := proto.Marshal(pbc)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return bs.db.Set(reckonObservedEndorseToken(altitude), observedEndorseOctets)
}

func (bs *LedgerDepot) Shutdown() error {
	return bs.db.Shutdown()
}

//

func reckonLedgerSummaryToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

func reckonLedgerFragmentToken(altitude int64, fragmentPosition int) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude, fragmentPosition))
}

func reckonLedgerEndorseToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

func reckonObservedEndorseToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

func reckonAddnEndorseToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

func reckonLedgerDigestToken(digest []byte) []byte {
	return []byte(fmt.Sprintf("REDACTED", digest))
}

//

var ledgerDepotToken = []byte("REDACTED")

//
//
//
func PersistLedgerDepotStatus(bsj *commitstore.LedgerDepotStatus, db dbm.DB) {
	persistLedgerDepotStatusClusterIntrinsic(bsj, db, nil)
}

//
//
func PersistLedgerDepotStatusCluster(bsj *commitstore.LedgerDepotStatus, cluster dbm.Cluster) {
	persistLedgerDepotStatusClusterIntrinsic(bsj, nil, cluster)
}

func persistLedgerDepotStatusClusterIntrinsic(bsj *commitstore.LedgerDepotStatus, db dbm.DB, cluster dbm.Cluster) {
	octets, err := proto.Marshal(bsj)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if cluster != nil {
		err = cluster.Set(ledgerDepotToken, octets)
	} else {
		if db == nil {
			panic("REDACTED")
		}
		err = db.AssignChronize(ledgerDepotToken, octets)
	}
	if err != nil {
		panic(err)
	}
}

//
//
func FetchLedgerDepotStatus(db dbm.DB) commitstore.LedgerDepotStatus {
	octets, err := db.Get(ledgerDepotToken)
	if err != nil {
		panic(err)
	}

	if len(octets) == 0 {
		return commitstore.LedgerDepotStatus{
			Foundation:   0,
			Altitude: 0,
		}
	}

	var bsj commitstore.LedgerDepotStatus
	if err := proto.Unmarshal(octets, &bsj); err != nil {
		panic(fmt.Sprintf("REDACTED", octets))
	}

	//
	if bsj.Altitude > 0 && bsj.Foundation == 0 {
		bsj.Foundation = 1
	}
	return bsj
}

//
func shouldEncode(pb proto.Message) []byte {
	bz, err := proto.Marshal(pb)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	return bz
}

//

//
//
func (bs *LedgerDepot) EraseNewestLedger() error {
	bs.mtx.RLock()
	objectiveAltitude := bs.altitude
	bs.mtx.RUnlock()

	cluster := bs.db.FreshCluster()
	defer cluster.Shutdown()

	//
	//
	if summary := bs.FetchLedgerSummary(objectiveAltitude); summary != nil {
		if err := cluster.Erase(reckonLedgerDigestToken(summary.LedgerUUID.Digest)); err != nil {
			return err
		}
		for p := 0; p < int(summary.LedgerUUID.FragmentAssignHeading.Sum); p++ {
			if err := cluster.Erase(reckonLedgerFragmentToken(objectiveAltitude, p)); err != nil {
				return err
			}
		}
	}
	if err := cluster.Erase(reckonLedgerEndorseToken(objectiveAltitude)); err != nil {
		return err
	}
	if err := cluster.Erase(reckonObservedEndorseToken(objectiveAltitude)); err != nil {
		return err
	}
	//
	if err := cluster.Erase(reckonLedgerSummaryToken(objectiveAltitude)); err != nil {
		return err
	}

	bs.mtx.Lock()
	defer bs.mtx.Unlock()
	bs.altitude = objectiveAltitude - 1
	return bs.persistStatusAlsoPersistDatastore(cluster, "REDACTED")
}
