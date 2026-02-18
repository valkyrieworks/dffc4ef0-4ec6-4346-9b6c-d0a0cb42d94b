package depot

import (
	"errors"
	"fmt"
	"strconv"

	cometfaults "github.com/valkyrieworks/kinds/faults"
	"github.com/cosmos/gogoproto/proto"
	lru "github.com/hashicorp/golang-lru/v2"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/proof"
	engineconnect "github.com/valkyrieworks/utils/align"
	commitdepot "github.com/valkyrieworks/schema/consensuscore/depot"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
const maximumLedgerSectionsToGroup = 10

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
	mtx    engineconnect.ReadwriteLock
	root   int64
	level int64

	viewedEndorseRepository          *lru.Cache[int64, *kinds.Endorse]
	ledgerEndorseRepository         *lru.Cache[int64, *kinds.Endorse]
	ledgerExpandedEndorseRepository *lru.Cache[int64, *kinds.ExpandedEndorse]
}

//
//
func NewLedgerDepot(db dbm.DB) *LedgerDepot {
	bs := ImportLedgerDepotStatus(db)
	byteDepot := &LedgerDepot{
		root:   bs.Root,
		level: bs.Level,
		db:     db,
	}
	byteDepot.appendRepositories()
	return byteDepot
}

func (bs *LedgerDepot) appendRepositories() {
	var err error
	//
	bs.ledgerEndorseRepository, err = lru.New[int64, *kinds.Endorse](100)
	if err != nil {
		panic(err)
	}
	bs.ledgerExpandedEndorseRepository, err = lru.New[int64, *kinds.ExpandedEndorse](100)
	if err != nil {
		panic(err)
	}
	bs.viewedEndorseRepository, err = lru.New[int64, *kinds.Endorse](100)
	if err != nil {
		panic(err)
	}
}

func (bs *LedgerDepot) IsEmpty() bool {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	return bs.root == 0 && bs.level == 0
}

//
func (bs *LedgerDepot) Root() int64 {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	return bs.root
}

//
func (bs *LedgerDepot) Level() int64 {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	return bs.level
}

//
func (bs *LedgerDepot) Volume() int64 {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	if bs.level == 0 {
		return 0
	}
	return bs.level - bs.root + 1
}

//
func (bs *LedgerDepot) ImportRootMeta() *kinds.LedgerMeta {
	bs.mtx.RLock()
	defer bs.mtx.RUnlock()
	if bs.root == 0 {
		return nil
	}
	return bs.ImportLedgerMeta(bs.root)
}

//
//
func (bs *LedgerDepot) ImportLedger(level int64) *kinds.Ledger {
	ledgerMeta := bs.ImportLedgerMeta(level)
	if ledgerMeta == nil {
		return nil
	}

	pbb := new(engineproto.Ledger)
	buf := []byte{}
	for i := 0; i < int(ledgerMeta.LedgerUID.SegmentAssignHeading.Sum); i++ {
		segment := bs.ImportLedgerSegment(level, i)
		//
		//
		if segment == nil {
			return nil
		}
		buf = append(buf, segment.Octets...)
	}
	err := proto.Unmarshal(buf, pbb)
	if err != nil {
		//
		//
		panic(fmt.Sprintf("REDACTED", err))
	}

	ledger, err := kinds.LedgerFromSchema(pbb)
	if err != nil {
		panic(cometfaults.ErrMessageFromSchema{SignalLabel: "REDACTED", Err: err})
	}

	return ledger
}

//
//
//
func (bs *LedgerDepot) ImportLedgerByDigest(digest []byte) *kinds.Ledger {
	bz, err := bs.db.Get(computeLedgerDigestKey(digest))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}

	s := string(bz)
	level, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", s, err))
	}
	return bs.ImportLedger(level)
}

//
//
//
func (bs *LedgerDepot) ImportLedgerSegment(level int64, ordinal int) *kinds.Segment {
	pbsection := new(engineproto.Segment)

	bz, err := bs.db.Get(computeLedgerSectionKey(level, ordinal))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}

	err = proto.Unmarshal(bz, pbsection)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	segment, err := kinds.SegmentFromSchema(pbsection)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	return segment
}

//
//
func (bs *LedgerDepot) ImportLedgerMeta(level int64) *kinds.LedgerMeta {
	pblm := new(engineproto.LedgerMeta)
	bz, err := bs.db.Get(computeLedgerMetaKey(level))
	if err != nil {
		panic(err)
	}

	if len(bz) == 0 {
		return nil
	}

	err = proto.Unmarshal(bz, pblm)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	ledgerMeta, err := kinds.LedgerMetaFromValidatedSchema(pblm)
	if err != nil {
		panic(cometfaults.ErrMessageFromSchema{SignalLabel: "REDACTED", Err: err})
	}

	return ledgerMeta
}

//
//
func (bs *LedgerDepot) ImportLedgerMetaByDigest(digest []byte) *kinds.LedgerMeta {
	bz, err := bs.db.Get(computeLedgerDigestKey(digest))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil
	}

	s := string(bz)
	level, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", s, err))
	}
	return bs.ImportLedgerMeta(level)
}

//
//
//
//
func (bs *LedgerDepot) ImportLedgerEndorse(level int64) *kinds.Endorse {
	xfer, ok := bs.ledgerEndorseRepository.Get(level)
	if ok {
		return xfer.Replicate()
	}
	pbc := new(engineproto.Endorse)
	bz, err := bs.db.Get(computeLedgerEndorseKey(level))
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
	endorse, err := kinds.EndorseFromSchema(pbc)
	if err != nil {
		panic(cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err})
	}
	bs.ledgerEndorseRepository.Add(level, endorse)
	return endorse.Replicate()
}

//
//
//
func (bs *LedgerDepot) ImportLedgerExpandedEndorse(level int64) *kinds.ExpandedEndorse {
	xfer, ok := bs.ledgerExpandedEndorseRepository.Get(level)
	if ok {
		return xfer.Replicate()
	}
	pbac := new(engineproto.ExpandedEndorse)
	bz, err := bs.db.Get(computeExtensionEndorseKey(level))
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	if len(bz) == 0 {
		return nil
	}
	err = proto.Unmarshal(bz, pbac)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	extensionEndorse, err := kinds.ExpandedEndorseFromSchema(pbac)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	bs.ledgerExpandedEndorseRepository.Add(level, extensionEndorse)
	return extensionEndorse.Replicate()
}

//
//
//
func (bs *LedgerDepot) ImportViewedEndorse(level int64) *kinds.Endorse {
	xfer, ok := bs.viewedEndorseRepository.Get(level)
	if ok {
		return xfer.Replicate()
	}
	pbc := new(engineproto.Endorse)
	bz, err := bs.db.Get(computeViewedEndorseKey(level))
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

	endorse, err := kinds.EndorseFromSchema(pbc)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	bs.viewedEndorseRepository.Add(level, endorse)
	return endorse.Replicate()
}

//
func (bs *LedgerDepot) TrimLedgers(level int64, status sm.Status) (uint64, int64, error) {
	if level <= 0 {
		return 0, -1, fmt.Errorf("REDACTED")
	}
	bs.mtx.RLock()
	if level > bs.level {
		bs.mtx.RUnlock()
		return 0, -1, fmt.Errorf("REDACTED", bs.level)
	}
	root := bs.root
	bs.mtx.RUnlock()
	if level < root {
		return 0, -1, fmt.Errorf("REDACTED",
			level, root)
	}

	trimmed := uint64(0)
	group := bs.db.NewGroup()
	defer group.End()
	purge := func(group dbm.Group, root int64) error {
		//
		//
		bs.mtx.Lock()
		defer group.End()
		defer bs.mtx.Unlock()
		bs.root = root
		return bs.persistStatusAndRecordStore(group, "REDACTED")
	}

	proofSpot := level
	for h := root; h < level; h++ {

		meta := bs.ImportLedgerMeta(h)
		if meta == nil { //
			continue
		}

		//
		//

		if proofSpot == level && !proof.IsProofLapsed(status.FinalLedgerLevel, status.FinalLedgerTime, h, meta.Heading.Time, status.AgreementOptions.Proof) {
			proofSpot = h
		}

		//
		if h < proofSpot {
			if err := group.Erase(computeLedgerMetaKey(h)); err != nil {
				return 0, -1, err
			}
		}
		if err := group.Erase(computeLedgerDigestKey(meta.LedgerUID.Digest)); err != nil {
			return 0, -1, err
		}
		//
		if h < proofSpot {
			if err := group.Erase(computeLedgerEndorseKey(h)); err != nil {
				return 0, -1, err
			}
		}
		if err := group.Erase(computeViewedEndorseKey(h)); err != nil {
			return 0, -1, err
		}

		if h < proofSpot {
			if err := group.Erase(computeExtensionEndorseKey(h)); err != nil {
				return 0, -1, err
			}
			bs.ledgerExpandedEndorseRepository.Remove(h)
		}

		for p := 0; p < int(meta.LedgerUID.SegmentAssignHeading.Sum); p++ {
			if err := group.Erase(computeLedgerSectionKey(h, p)); err != nil {
				return 0, -1, err
			}
		}
		trimmed++

		//
		if trimmed%1000 == 0 && trimmed > 0 {
			err := purge(group, h)
			if err != nil {
				return 0, -1, err
			}
			group = bs.db.NewGroup()
			defer group.End()
		}
	}

	err := purge(group, level)
	if err != nil {
		return 0, -1, err
	}
	return trimmed, proofSpot, nil
}

//
//
//
//
//
//
//
func (bs *LedgerDepot) PersistLedger(ledger *kinds.Ledger, ledgerSegments *kinds.SegmentCollection, viewedEndorse *kinds.Endorse) {
	if ledger == nil {
		panic("REDACTED")
	}

	group := bs.db.NewGroup()
	defer group.End()

	if err := bs.persistLedgerToGroup(ledger, ledgerSegments, viewedEndorse, group); err != nil {
		panic(err)
	}

	bs.mtx.Lock()
	defer bs.mtx.Unlock()
	bs.level = ledger.Level
	if bs.root == 0 {
		bs.root = ledger.Level
	}

	//
	err := bs.persistStatusAndRecordStore(group, "REDACTED")
	if err != nil {
		panic(err)
	}
}

//
//
//
//
//
func (bs *LedgerDepot) PersistLedgerWithExpandedEndorse(ledger *kinds.Ledger, ledgerSegments *kinds.SegmentCollection, viewedExpandedEndorse *kinds.ExpandedEndorse) {
	if ledger == nil {
		panic("REDACTED")
	}
	if err := viewedExpandedEndorse.AssurePlugins(true); err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	group := bs.db.NewGroup()
	defer group.End()

	if err := bs.persistLedgerToGroup(ledger, ledgerSegments, viewedExpandedEndorse.ToEndorse(), group); err != nil {
		panic(err)
	}
	level := ledger.Level

	pbac := viewedExpandedEndorse.ToSchema()
	extensionEndorseOctets := shouldMarshal(pbac)
	if err := group.Set(computeExtensionEndorseKey(level), extensionEndorseOctets); err != nil {
		panic(err)
	}

	bs.mtx.Lock()
	defer bs.mtx.Unlock()
	bs.level = level
	if bs.root == 0 {
		bs.root = level
	}

	//
	err := bs.persistStatusAndRecordStore(group, "REDACTED")
	if err != nil {
		panic(err)
	}
}

func (bs *LedgerDepot) persistLedgerToGroup(
	ledger *kinds.Ledger,
	ledgerSegments *kinds.SegmentCollection,
	viewedEndorse *kinds.Endorse,
	group dbm.Group,
) error {
	if ledger == nil {
		panic("REDACTED")
	}

	level := ledger.Level
	digest := ledger.Digest()

	if g, w := level, bs.Level()+1; bs.Root() > 0 && g != w {
		return fmt.Errorf("REDACTED", w, g)
	}
	if !ledgerSegments.IsFinished() {
		return errors.New("REDACTED")
	}
	if level != viewedEndorse.Level {
		return fmt.Errorf("REDACTED", level, viewedEndorse.Level)
	}

	//
	//
	persistLedgerSectionsToGroup := ledgerSegments.Number() <= maximumLedgerSectionsToGroup

	//
	//
	//
	//
	for i := 0; i < int(ledgerSegments.Sum()); i++ {
		segment := ledgerSegments.FetchSegment(i)
		bs.persistLedgerSection(level, i, segment, group, persistLedgerSectionsToGroup)
	}

	//
	ledgerMeta := kinds.NewLedgerMeta(ledger, ledgerSegments)
	pbm := ledgerMeta.ToSchema()
	if pbm == nil {
		return errors.New("REDACTED")
	}
	metaOctets := shouldMarshal(pbm)
	if err := group.Set(computeLedgerMetaKey(level), metaOctets); err != nil {
		return err
	}
	if err := group.Set(computeLedgerDigestKey(digest), []byte(fmt.Sprintf("REDACTED", level))); err != nil {
		return err
	}

	//
	pbc := ledger.FinalEndorse.ToSchema()
	ledgerEndorseOctets := shouldMarshal(pbc)
	if err := group.Set(computeLedgerEndorseKey(level-1), ledgerEndorseOctets); err != nil {
		return err
	}

	//
	//
	pbcc := viewedEndorse.ToSchema()
	viewedEndorseOctets := shouldMarshal(pbcc)
	if err := group.Set(computeViewedEndorseKey(level), viewedEndorseOctets); err != nil {
		return err
	}

	return nil
}

func (bs *LedgerDepot) persistLedgerSection(level int64, ordinal int, segment *kinds.Segment, group dbm.Group, persistLedgerSectionsToGroup bool) {
	pbp, err := segment.ToSchema()
	if err != nil {
		panic(cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err})
	}
	sectionOctets := shouldMarshal(pbp)
	if persistLedgerSectionsToGroup {
		err = group.Set(computeLedgerSectionKey(level, ordinal), sectionOctets)
	} else {
		err = bs.db.Set(computeLedgerSectionKey(level, ordinal), sectionOctets)
	}
	if err != nil {
		panic(err)
	}
}

//
func (bs *LedgerDepot) persistStatusAndRecordStore(group dbm.Group, errMessage string) error {
	bss := commitdepot.LedgerDepotStatus{
		Root:   bs.root,
		Level: bs.level,
	}
	PersistLedgerDepotStatusGroup(&bss, group)

	err := group.RecordAlign()
	if err != nil {
		return fmt.Errorf("REDACTED",
			errMessage, bs.root, bs.level, err)
	}
	return nil
}

//
func (bs *LedgerDepot) PersistViewedEndorse(level int64, viewedEndorse *kinds.Endorse) error {
	pbc := viewedEndorse.ToSchema()
	viewedEndorseOctets, err := proto.Marshal(pbc)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return bs.db.Set(computeViewedEndorseKey(level), viewedEndorseOctets)
}

func (bs *LedgerDepot) End() error {
	return bs.db.End()
}

//

func computeLedgerMetaKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

func computeLedgerSectionKey(level int64, sectionOrdinal int) []byte {
	return []byte(fmt.Sprintf("REDACTED", level, sectionOrdinal))
}

func computeLedgerEndorseKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

func computeViewedEndorseKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

func computeExtensionEndorseKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

func computeLedgerDigestKey(digest []byte) []byte {
	return []byte(fmt.Sprintf("REDACTED", digest))
}

//

var ledgerDepotKey = []byte("REDACTED")

//
//
//
func PersistLedgerDepotStatus(bsj *commitdepot.LedgerDepotStatus, db dbm.DB) {
	persistLedgerDepotStatusGroupIntrinsic(bsj, db, nil)
}

//
//
func PersistLedgerDepotStatusGroup(bsj *commitdepot.LedgerDepotStatus, group dbm.Group) {
	persistLedgerDepotStatusGroupIntrinsic(bsj, nil, group)
}

func persistLedgerDepotStatusGroupIntrinsic(bsj *commitdepot.LedgerDepotStatus, db dbm.DB, group dbm.Group) {
	octets, err := proto.Marshal(bsj)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	if group != nil {
		err = group.Set(ledgerDepotKey, octets)
	} else {
		if db == nil {
			panic("REDACTED")
		}
		err = db.CollectionAlign(ledgerDepotKey, octets)
	}
	if err != nil {
		panic(err)
	}
}

//
//
func ImportLedgerDepotStatus(db dbm.DB) commitdepot.LedgerDepotStatus {
	octets, err := db.Get(ledgerDepotKey)
	if err != nil {
		panic(err)
	}

	if len(octets) == 0 {
		return commitdepot.LedgerDepotStatus{
			Root:   0,
			Level: 0,
		}
	}

	var bsj commitdepot.LedgerDepotStatus
	if err := proto.Unmarshal(octets, &bsj); err != nil {
		panic(fmt.Sprintf("REDACTED", octets))
	}

	//
	if bsj.Level > 0 && bsj.Root == 0 {
		bsj.Root = 1
	}
	return bsj
}

//
func shouldMarshal(pb proto.Message) []byte {
	bz, err := proto.Marshal(pb)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	return bz
}

//

//
//
func (bs *LedgerDepot) RemoveNewestLedger() error {
	bs.mtx.RLock()
	objectiveLevel := bs.level
	bs.mtx.RUnlock()

	group := bs.db.NewGroup()
	defer group.End()

	//
	//
	if meta := bs.ImportLedgerMeta(objectiveLevel); meta != nil {
		if err := group.Erase(computeLedgerDigestKey(meta.LedgerUID.Digest)); err != nil {
			return err
		}
		for p := 0; p < int(meta.LedgerUID.SegmentAssignHeading.Sum); p++ {
			if err := group.Erase(computeLedgerSectionKey(objectiveLevel, p)); err != nil {
				return err
			}
		}
	}
	if err := group.Erase(computeLedgerEndorseKey(objectiveLevel)); err != nil {
		return err
	}
	if err := group.Erase(computeViewedEndorseKey(objectiveLevel)); err != nil {
		return err
	}
	//
	if err := group.Erase(computeLedgerMetaKey(objectiveLevel)); err != nil {
		return err
	}

	bs.mtx.Lock()
	defer bs.mtx.Unlock()
	bs.level = objectiveLevel - 1
	return bs.persistStatusAndRecordStore(group, "REDACTED")
}
