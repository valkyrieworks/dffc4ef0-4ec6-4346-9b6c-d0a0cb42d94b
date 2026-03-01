package db

import (
	"encoding/binary"
	"fmt"
	"regexp"
	"strconv"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var extentToken = []byte("REDACTED")

type dbs struct {
	db     dbm.DB
	heading string

	mtx  commitchronize.ReadwriteExclusion
	extent uint16
}

//
//
func New(db dbm.DB, heading string) depot.Depot {
	extent := uint16(0)
	bz, err := db.Get(extentToken)
	if err == nil && len(bz) > 0 {
		extent = decodeExtent(bz)
	}

	return &dbs{db: db, heading: heading, extent: extent}
}

//
//
//
func (s *dbs) PersistAgileLedger(lb *kinds.AgileLedger) error {
	if lb.Altitude <= 0 {
		panic("REDACTED")
	}

	ldgrpb, err := lb.TowardSchema()
	if err != nil {
		return strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
	}

	ldgrByz, err := ldgrpb.Serialize()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	b := s.db.FreshCluster()
	defer b.Shutdown()
	if err = b.Set(s.ldgrToken(lb.Altitude), ldgrByz); err != nil {
		return err
	}
	if err = b.Set(extentToken, serializeExtent(s.extent+1)); err != nil {
		return err
	}
	if err = b.PersistChronize(); err != nil {
		return err
	}
	s.extent++

	return nil
}

//
//
//
//
func (s *dbs) EraseAgileLedger(altitude int64) error {
	if altitude <= 0 {
		panic("REDACTED")
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	b := s.db.FreshCluster()
	defer b.Shutdown()
	if err := b.Erase(s.ldgrToken(altitude)); err != nil {
		return err
	}
	if err := b.Set(extentToken, serializeExtent(s.extent-1)); err != nil {
		return err
	}
	if err := b.PersistChronize(); err != nil {
		return err
	}
	s.extent--

	return nil
}

//
//
//
func (s *dbs) AgileLedger(altitude int64) (*kinds.AgileLedger, error) {
	if altitude <= 0 {
		panic("REDACTED")
	}

	bz, err := s.db.Get(s.ldgrToken(altitude))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil, depot.FaultAgileLedgerNegationDetected
	}

	var ldgrpb commitchema.AgileLedger
	err = ldgrpb.Decode(bz)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	agileLedger, err := kinds.AgileLedgerOriginatingSchema(&ldgrpb)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return agileLedger, err
}

//
//
//
func (s *dbs) FinalAgileLedgerAltitude() (int64, error) {
	itr, err := s.db.InvertTraverser(
		s.ldgrToken(1),
		append(s.ldgrToken(1<<63-1), byte(0x00)),
	)
	if err != nil {
		panic(err)
	}
	defer itr.Shutdown()

	for itr.Sound() {
		key := itr.Key()
		_, altitude, ok := analyzeLdgrToken(key)
		if ok {
			return altitude, nil
		}
		itr.Following()
	}

	return -1, itr.Failure()
}

//
//
//
func (s *dbs) InitialAgileLedgerAltitude() (int64, error) {
	itr, err := s.db.Traverser(
		s.ldgrToken(1),
		append(s.ldgrToken(1<<63-1), byte(0x00)),
	)
	if err != nil {
		panic(err)
	}
	defer itr.Shutdown()

	for itr.Sound() {
		key := itr.Key()
		_, altitude, ok := analyzeLdgrToken(key)
		if ok {
			return altitude, nil
		}
		itr.Following()
	}

	return -1, itr.Failure()
}

//
//
//
//
func (s *dbs) AgileLedgerPrior(altitude int64) (*kinds.AgileLedger, error) {
	if altitude <= 0 {
		panic("REDACTED")
	}

	itr, err := s.db.InvertTraverser(
		s.ldgrToken(1),
		s.ldgrToken(altitude),
	)
	if err != nil {
		panic(err)
	}
	defer itr.Shutdown()

	for itr.Sound() {
		key := itr.Key()
		_, extantAltitude, ok := analyzeLdgrToken(key)
		if ok {
			return s.AgileLedger(extantAltitude)
		}
		itr.Following()
	}
	if err = itr.Failure(); err != nil {
		return nil, err
	}

	return nil, depot.FaultAgileLedgerNegationDetected
}

//
//
//
//
func (s *dbs) Trim(extent uint16) error {
	//
	s.mtx.RLock()
	strExtent := s.extent
	s.mtx.RUnlock()

	if strExtent <= extent { //
		return nil
	}
	countTowardTrim := strExtent - extent

	//
	itr, err := s.db.Traverser(
		s.ldgrToken(1),
		append(s.ldgrToken(1<<63-1), byte(0x00)),
	)
	if err != nil {
		return err
	}
	defer itr.Shutdown()

	b := s.db.FreshCluster()
	defer b.Shutdown()

	trimmed := 0
	for itr.Sound() && countTowardTrim > 0 {
		key := itr.Key()
		_, altitude, ok := analyzeLdgrToken(key)
		if ok {
			if err = b.Erase(s.ldgrToken(altitude)); err != nil {
				return err
			}
		}
		itr.Following()
		countTowardTrim--
		trimmed++
	}
	if err = itr.Failure(); err != nil {
		return err
	}

	err = b.PersistChronize()
	if err != nil {
		return err
	}

	//
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.extent -= uint16(trimmed)

	if wrFault := s.db.AssignChronize(extentToken, serializeExtent(s.extent)); wrFault != nil {
		return fmt.Errorf("REDACTED", wrFault)
	}

	return nil
}

//
//
//
func (s *dbs) Extent() uint16 {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.extent
}

func (s *dbs) ldgrToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", s.heading, altitude))
}

var tokenTemplate = regexp.MustCompile("REDACTED")

func analyzeToken(key []byte) (fragment string, heading string, altitude int64, ok bool) {
	undermatch := tokenTemplate.FindSubmatch(key)
	if undermatch == nil {
		return "REDACTED", "REDACTED", 0, false
	}
	fragment = string(undermatch[1])
	heading = string(undermatch[2])
	altitude, err := strconv.ParseInt(string(undermatch[3]), 10, 64)
	if err != nil {
		return "REDACTED", "REDACTED", 0, false
	}
	ok = true //
	return
}

func analyzeLdgrToken(key []byte) (heading string, altitude int64, ok bool) {
	var fragment string
	fragment, heading, altitude, ok = analyzeToken(key)
	if fragment != "REDACTED" {
		return "REDACTED", 0, false
	}
	return
}

func serializeExtent(extent uint16) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, extent)
	return bs
}

func decodeExtent(bz []byte) uint16 {
	return binary.LittleEndian.Uint16(bz)
}
