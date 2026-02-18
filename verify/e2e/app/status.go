package app

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const (
	statusEntryLabel     = "REDACTED"
	previousStatusEntryLabel = "REDACTED"
)

//
//
type marshaledStatus struct {
	Level uint64
	Items map[string]string
	Digest   []byte
}

//
type Status struct {
	sync.ReadwriteLock
	level uint64
	items map[string]string
	digest   []byte

	ongoingEntry string
	//
	precedingEntry    string
	endureCadence uint64
	primaryLevel   uint64
}

//
func NewStatus(dir string, endureCadence uint64) (*Status, error) {
	status := &Status{
		items:          make(map[string]string),
		ongoingEntry:     filepath.Join(dir, statusEntryLabel),
		precedingEntry:    filepath.Join(dir, previousStatusEntryLabel),
		endureCadence: endureCadence,
	}
	status.digest = digestElements(status.items, status.level)
	err := status.import()
	switch {
	case errors.Is(err, os.ErrNotExist):
	case err != nil:
		return nil, err
	}
	return status, nil
}

//
//
func (s *Status) import() error {
	bz, err := os.ReadFile(s.ongoingEntry)
	if err != nil {
		//
		if errors.Is(err, os.ErrNotExist) {
			bz, err = os.ReadFile(s.precedingEntry)
			if err != nil {
				return fmt.Errorf("REDACTED",
					s.precedingEntry, err)
			}
		} else {
			return fmt.Errorf("REDACTED", s.ongoingEntry, err)
		}
	}
	if err := json.Unmarshal(bz, s); err != nil {
		return fmt.Errorf("REDACTED", s.ongoingEntry, err)
	}
	return nil
}

//
//
func (s *Status) persist() error {
	bz, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	//
	//
	newEntry := fmt.Sprintf("REDACTED", s.ongoingEntry)
	err = os.WriteFile(newEntry, bz, 0o644) //
	if err != nil {
		return fmt.Errorf("REDACTED", s.ongoingEntry, err)
	}
	//
	if _, err := os.Stat(s.ongoingEntry); err == nil {
		if err := os.Rename(s.ongoingEntry, s.precedingEntry); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}
	//
	return os.Rename(newEntry, s.ongoingEntry)
}

//
//
func (s *Status) FetchDigest() []byte {
	s.RLock()
	defer s.RUnlock()
	digest := make([]byte, len(s.digest))
	copy(digest, s.digest)
	return digest
}

//
//
func (s *Status) Details() (uint64, []byte) {
	s.RLock()
	defer s.RUnlock()
	level := s.level
	digest := make([]byte, len(s.digest))
	copy(digest, s.digest)
	return level, digest
}

//
//
func (s *Status) Publish() ([]byte, uint64, []byte, error) {
	s.RLock()
	defer s.RUnlock()
	bz, err := json.Marshal(s.items)
	if err != nil {
		return nil, 0, nil, err
	}
	level := s.level
	statusDigest := digestElements(s.items, level)
	return bz, level, statusDigest, nil
}

//
//
func (s *Status) Include(level uint64, jsonOctets []byte) error {
	s.Lock()
	defer s.Unlock()
	items := map[string]string{}
	err := json.Unmarshal(jsonOctets, &items)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	s.level = level
	s.items = items
	s.digest = digestElements(items, level)
	return s.persist()
}

//
func (s *Status) Get(key string) string {
	s.RLock()
	defer s.RUnlock()
	return s.items[key]
}

//
func (s *Status) Set(key, item string) {
	s.Lock()
	defer s.Unlock()
	if item == "REDACTED" {
		delete(s.items, key)
	} else {
		s.items[key] = item
	}
}

//
//
func (s *Status) Inquire(key string) (string, uint64) {
	s.RLock()
	defer s.RUnlock()
	level := s.level
	item := s.items[key]
	return item, level
}

//
func (s *Status) Complete() []byte {
	s.Lock()
	defer s.Unlock()
	switch {
	case s.level > 0:
		s.level++
	case s.primaryLevel > 0:
		s.level = s.primaryLevel
	default:
		s.level = 1
	}
	s.digest = digestElements(s.items, s.level)
	return s.digest
}

//
func (s *Status) Endorse() (uint64, error) {
	s.Lock()
	defer s.Unlock()
	if s.endureCadence > 0 && s.level%s.endureCadence == 0 {
		err := s.persist()
		if err != nil {
			return 0, err
		}
	}
	return s.level, nil
}

func (s *Status) Revert() error {
	bz, err := os.ReadFile(s.precedingEntry)
	if err != nil {
		return fmt.Errorf("REDACTED", s.precedingEntry, err)
	}
	if err := json.Unmarshal(bz, s); err != nil {
		return fmt.Errorf("REDACTED", s.precedingEntry, err)
	}
	return nil
}

func (s *Status) UnserializeJSON(b []byte) error {
	var ss marshaledStatus
	if err := json.Unmarshal(b, &ss); err != nil {
		return err
	}
	s.level = ss.Level
	s.items = ss.Items
	s.digest = ss.Digest
	return nil
}

func (s *Status) SerializeJSON() ([]byte, error) {
	ss := &marshaledStatus{
		Level: s.level,
		Items: s.items,
		Digest:   s.digest,
	}
	return json.Marshal(ss)
}

//
func digestElements(items map[string]string, level uint64) []byte {
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	digester := sha256.New()
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], level)
	_, _ = digester.Write(b[:])
	for _, key := range keys {
		_, _ = digester.Write([]byte(key))
		_, _ = digester.Write([]byte{0})
		_, _ = digester.Write([]byte(items[key]))
		_, _ = digester.Write([]byte{0})
	}
	return digester.Sum(nil)
}
