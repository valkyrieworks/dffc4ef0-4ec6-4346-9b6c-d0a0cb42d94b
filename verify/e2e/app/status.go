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
	statusRecordAlias     = "REDACTED"
	previousStatusRecordAlias = "REDACTED"
)

//
//
type marshaledStatus struct {
	Altitude uint64
	Items map[string]string
	Digest   []byte
}

//
type Status struct {
	sync.ReadwriteExclusion
	altitude uint64
	items map[string]string
	digest   []byte

	prevailingRecord string
	//
	precedingRecord    string
	endureDuration uint64
	primaryAltitude   uint64
}

//
func FreshStatus(dir string, endureDuration uint64) (*Status, error) {
	status := &Status{
		items:          make(map[string]string),
		prevailingRecord:     filepath.Join(dir, statusRecordAlias),
		precedingRecord:    filepath.Join(dir, previousStatusRecordAlias),
		endureDuration: endureDuration,
	}
	status.digest = digestElements(status.items, status.altitude)
	err := status.fetch()
	switch {
	case errors.Is(err, os.ErrNotExist):
	case err != nil:
		return nil, err
	}
	return status, nil
}

//
//
func (s *Status) fetch() error {
	bz, err := os.ReadFile(s.prevailingRecord)
	if err != nil {
		//
		if errors.Is(err, os.ErrNotExist) {
			bz, err = os.ReadFile(s.precedingRecord)
			if err != nil {
				return fmt.Errorf("REDACTED",
					s.precedingRecord, err)
			}
		} else {
			return fmt.Errorf("REDACTED", s.prevailingRecord, err)
		}
	}
	if err := json.Unmarshal(bz, s); err != nil {
		return fmt.Errorf("REDACTED", s.prevailingRecord, err)
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
	freshRecord := fmt.Sprintf("REDACTED", s.prevailingRecord)
	err = os.WriteFile(freshRecord, bz, 0o644) //
	if err != nil {
		return fmt.Errorf("REDACTED", s.prevailingRecord, err)
	}
	//
	if _, err := os.Stat(s.prevailingRecord); err == nil {
		if err := os.Rename(s.prevailingRecord, s.precedingRecord); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}
	//
	return os.Rename(freshRecord, s.prevailingRecord)
}

//
//
func (s *Status) ObtainDigest() []byte {
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
	altitude := s.altitude
	digest := make([]byte, len(s.digest))
	copy(digest, s.digest)
	return altitude, digest
}

//
//
func (s *Status) Disclose() ([]byte, uint64, []byte, error) {
	s.RLock()
	defer s.RUnlock()
	bz, err := json.Marshal(s.items)
	if err != nil {
		return nil, 0, nil, err
	}
	altitude := s.altitude
	statusDigest := digestElements(s.items, altitude)
	return bz, altitude, statusDigest, nil
}

//
//
func (s *Status) Ingest(altitude uint64, jsnOctets []byte) error {
	s.Lock()
	defer s.Unlock()
	items := map[string]string{}
	err := json.Unmarshal(jsnOctets, &items)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	s.altitude = altitude
	s.items = items
	s.digest = digestElements(items, altitude)
	return s.persist()
}

//
func (s *Status) Get(key string) string {
	s.RLock()
	defer s.RUnlock()
	return s.items[key]
}

//
func (s *Status) Set(key, datum string) {
	s.Lock()
	defer s.Unlock()
	if datum == "REDACTED" {
		delete(s.items, key)
	} else {
		s.items[key] = datum
	}
}

//
//
func (s *Status) Inquire(key string) (string, uint64) {
	s.RLock()
	defer s.RUnlock()
	altitude := s.altitude
	datum := s.items[key]
	return datum, altitude
}

//
func (s *Status) Culminate() []byte {
	s.Lock()
	defer s.Unlock()
	switch {
	case s.altitude > 0:
		s.altitude++
	case s.primaryAltitude > 0:
		s.altitude = s.primaryAltitude
	default:
		s.altitude = 1
	}
	s.digest = digestElements(s.items, s.altitude)
	return s.digest
}

//
func (s *Status) Endorse() (uint64, error) {
	s.Lock()
	defer s.Unlock()
	if s.endureDuration > 0 && s.altitude%s.endureDuration == 0 {
		err := s.persist()
		if err != nil {
			return 0, err
		}
	}
	return s.altitude, nil
}

func (s *Status) Revert() error {
	bz, err := os.ReadFile(s.precedingRecord)
	if err != nil {
		return fmt.Errorf("REDACTED", s.precedingRecord, err)
	}
	if err := json.Unmarshal(bz, s); err != nil {
		return fmt.Errorf("REDACTED", s.precedingRecord, err)
	}
	return nil
}

func (s *Status) DecodeJSN(b []byte) error {
	var ss marshaledStatus
	if err := json.Unmarshal(b, &ss); err != nil {
		return err
	}
	s.altitude = ss.Altitude
	s.items = ss.Items
	s.digest = ss.Digest
	return nil
}

func (s *Status) SerializeJSN() ([]byte, error) {
	ss := &marshaledStatus{
		Altitude: s.altitude,
		Items: s.items,
		Digest:   s.digest,
	}
	return json.Marshal(ss)
}

//
func digestElements(elements map[string]string, altitude uint64) []byte {
	tokens := make([]string, 0, len(elements))
	for key := range elements {
		tokens = append(tokens, key)
	}
	sort.Strings(tokens)

	digester := sha256.New()
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], altitude)
	_, _ = digester.Write(b[:])
	for _, key := range tokens {
		_, _ = digester.Write([]byte(key))
		_, _ = digester.Write([]byte{0})
		_, _ = digester.Write([]byte(elements[key]))
		_, _ = digester.Write([]byte{0})
	}
	return digester.Sum(nil)
}
