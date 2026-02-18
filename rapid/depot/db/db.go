package db

import (
	"encoding/binary"
	"fmt"
	"regexp"
	"strconv"

	dbm "github.com/valkyrieworks/-db"
	cometfaults "github.com/valkyrieworks/kinds/faults"

	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/rapid/depot"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

var volumeKey = []byte("REDACTED")

type dbs struct {
	db     dbm.DB
	prefix string

	mtx  engineconnect.ReadwriteLock
	volume uint16
}

//
//
func New(db dbm.DB, prefix string) depot.Depot {
	volume := uint16(0)
	bz, err := db.Get(volumeKey)
	if err == nil && len(bz) > 0 {
		volume = unserializeVolume(bz)
	}

	return &dbs{db: db, prefix: prefix, volume: volume}
}

//
//
//
func (s *dbs) PersistRapidLedger(lb *kinds.RapidLedger) error {
	if lb.Level <= 0 {
		panic("REDACTED")
	}

	lbpb, err := lb.ToSchema()
	if err != nil {
		return cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
	}

	blockBz, err := lbpb.Serialize()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	b := s.db.NewGroup()
	defer b.End()
	if err = b.Set(s.blockKey(lb.Level), blockBz); err != nil {
		return err
	}
	if err = b.Set(volumeKey, serializeVolume(s.volume+1)); err != nil {
		return err
	}
	if err = b.RecordAlign(); err != nil {
		return err
	}
	s.volume++

	return nil
}

//
//
//
//
func (s *dbs) EraseRapidLedger(level int64) error {
	if level <= 0 {
		panic("REDACTED")
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	b := s.db.NewGroup()
	defer b.End()
	if err := b.Erase(s.blockKey(level)); err != nil {
		return err
	}
	if err := b.Set(volumeKey, serializeVolume(s.volume-1)); err != nil {
		return err
	}
	if err := b.RecordAlign(); err != nil {
		return err
	}
	s.volume--

	return nil
}

//
//
//
func (s *dbs) RapidLedger(level int64) (*kinds.RapidLedger, error) {
	if level <= 0 {
		panic("REDACTED")
	}

	bz, err := s.db.Get(s.blockKey(level))
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return nil, depot.ErrRapidLedgerNegateLocated
	}

	var lbpb engineproto.RapidLedger
	err = lbpb.Unserialize(bz)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	rapidLedger, err := kinds.RapidLedgerFromSchema(&lbpb)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return rapidLedger, err
}

//
//
//
func (s *dbs) FinalRapidLedgerLevel() (int64, error) {
	itr, err := s.db.InvertRepeater(
		s.blockKey(1),
		append(s.blockKey(1<<63-1), byte(0x00)),
	)
	if err != nil {
		panic(err)
	}
	defer itr.End()

	for itr.Sound() {
		key := itr.Key()
		_, level, ok := analyzeBlockKey(key)
		if ok {
			return level, nil
		}
		itr.Following()
	}

	return -1, itr.Fault()
}

//
//
//
func (s *dbs) InitialRapidLedgerLevel() (int64, error) {
	itr, err := s.db.Repeater(
		s.blockKey(1),
		append(s.blockKey(1<<63-1), byte(0x00)),
	)
	if err != nil {
		panic(err)
	}
	defer itr.End()

	for itr.Sound() {
		key := itr.Key()
		_, level, ok := analyzeBlockKey(key)
		if ok {
			return level, nil
		}
		itr.Following()
	}

	return -1, itr.Fault()
}

//
//
//
//
func (s *dbs) RapidLedgerPrior(level int64) (*kinds.RapidLedger, error) {
	if level <= 0 {
		panic("REDACTED")
	}

	itr, err := s.db.InvertRepeater(
		s.blockKey(1),
		s.blockKey(level),
	)
	if err != nil {
		panic(err)
	}
	defer itr.End()

	for itr.Sound() {
		key := itr.Key()
		_, presentLevel, ok := analyzeBlockKey(key)
		if ok {
			return s.RapidLedger(presentLevel)
		}
		itr.Following()
	}
	if err = itr.Fault(); err != nil {
		return nil, err
	}

	return nil, depot.ErrRapidLedgerNegateLocated
}

//
//
//
//
func (s *dbs) Trim(volume uint16) error {
	//
	s.mtx.RLock()
	sVolume := s.volume
	s.mtx.RUnlock()

	if sVolume <= volume { //
		return nil
	}
	countToTrim := sVolume - volume

	//
	itr, err := s.db.Repeater(
		s.blockKey(1),
		append(s.blockKey(1<<63-1), byte(0x00)),
	)
	if err != nil {
		return err
	}
	defer itr.End()

	b := s.db.NewGroup()
	defer b.End()

	trimmed := 0
	for itr.Sound() && countToTrim > 0 {
		key := itr.Key()
		_, level, ok := analyzeBlockKey(key)
		if ok {
			if err = b.Erase(s.blockKey(level)); err != nil {
				return err
			}
		}
		itr.Following()
		countToTrim--
		trimmed++
	}
	if err = itr.Fault(); err != nil {
		return err
	}

	err = b.RecordAlign()
	if err != nil {
		return err
	}

	//
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.volume -= uint16(trimmed)

	if writerErr := s.db.CollectionAlign(volumeKey, serializeVolume(s.volume)); writerErr != nil {
		return fmt.Errorf("REDACTED", writerErr)
	}

	return nil
}

//
//
//
func (s *dbs) Volume() uint16 {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.volume
}

func (s *dbs) blockKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", s.prefix, level))
}

var keyTemplate = regexp.MustCompile("REDACTED")

func analyzeKey(key []byte) (segment string, prefix string, level int64, ok bool) {
	subsection := keyTemplate.FindSubmatch(key)
	if subsection == nil {
		return "REDACTED", "REDACTED", 0, false
	}
	segment = string(subsection[1])
	prefix = string(subsection[2])
	level, err := strconv.ParseInt(string(subsection[3]), 10, 64)
	if err != nil {
		return "REDACTED", "REDACTED", 0, false
	}
	ok = true //
	return
}

func analyzeBlockKey(key []byte) (prefix string, level int64, ok bool) {
	var segment string
	segment, prefix, level, ok = analyzeKey(key)
	if segment != "REDACTED" {
		return "REDACTED", 0, false
	}
	return
}

func serializeVolume(volume uint16) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, volume)
	return bs
}

func unserializeVolume(bz []byte) uint16 {
	return binary.LittleEndian.Uint16(bz)
}
