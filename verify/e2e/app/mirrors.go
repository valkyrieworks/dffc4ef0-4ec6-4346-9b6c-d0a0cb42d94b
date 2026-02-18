package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sync"

	iface "github.com/valkyrieworks/iface/kinds"
)

const (
	mirrorSegmentVolume = 1e6

	//
	maximumMirrorNumber = 10
)

//
//
//
type MirrorDepot struct {
	sync.ReadwriteLock
	dir      string
	metainfo []iface.Mirror
}

//
func NewMirrorDepot(dir string) (*MirrorDepot, error) {
	depot := &MirrorDepot{dir: dir}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	if err := depot.importMetainfo(); err != nil {
		return nil, err
	}
	return depot, nil
}

//
//
func (s *MirrorDepot) importMetainfo() error {
	entry := filepath.Join(s.dir, "REDACTED")
	metainfo := []iface.Mirror{}

	bz, err := os.ReadFile(entry)
	switch {
	case errors.Is(err, os.ErrNotExist):
	case err != nil:
		return fmt.Errorf("REDACTED", entry, err)
	}
	if len(bz) != 0 {
		err = json.Unmarshal(bz, &metainfo)
		if err != nil {
			return fmt.Errorf("REDACTED", entry, err)
		}
	}
	s.metainfo = metainfo
	return nil
}

//
//
func (s *MirrorDepot) persistMetainfo() error {
	bz, err := json.Marshal(s.metainfo)
	if err != nil {
		return err
	}

	//
	newEntry := filepath.Join(s.dir, "REDACTED")
	entry := filepath.Join(s.dir, "REDACTED")
	err = os.WriteFile(newEntry, bz, 0o644) //
	if err != nil {
		return err
	}
	return os.Rename(newEntry, entry)
}

//
func (s *MirrorDepot) Instantiate(status *Status) (iface.Mirror, error) {
	s.Lock()
	defer s.Unlock()
	bz, level, statusDigest, err := status.Publish()
	if err != nil {
		return iface.Mirror{}, err
	}
	mirror := iface.Mirror{
		Level: level,
		Layout: 1,
		Digest:   statusDigest,
		Segments: octetSegments(bz),
	}
	err = os.WriteFile(filepath.Join(s.dir, fmt.Sprintf("REDACTED", level)), bz, 0o644) //
	if err != nil {
		return iface.Mirror{}, err
	}
	s.metainfo = append(s.metainfo, mirror)
	err = s.persistMetainfo()
	if err != nil {
		return iface.Mirror{}, err
	}
	return mirror, nil
}

//
func (s *MirrorDepot) Trim(n int) error {
	s.Lock()
	defer s.Unlock()
	//
	//
	i := 0
	for ; i < len(s.metainfo)-n; i++ {
		h := s.metainfo[i].Level
		if err := os.Remove(filepath.Join(s.dir, fmt.Sprintf("REDACTED", h))); err != nil {
			return err
		}
	}

	//
	trimmed := make([]iface.Mirror, len(s.metainfo[i:]))
	copy(trimmed, s.metainfo[i:])
	s.metainfo = trimmed
	return nil
}

//
func (s *MirrorDepot) Catalog() ([]*iface.Mirror, error) {
	s.RLock()
	defer s.RUnlock()
	mirrors := make([]*iface.Mirror, len(s.metainfo))
	for idx := range s.metainfo {
		mirrors[idx] = &s.metainfo[idx]
	}
	return mirrors, nil
}

//
func (s *MirrorDepot) ImportSegment(level uint64, layout uint32, segment uint32) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()
	for _, mirror := range s.metainfo {
		if mirror.Level == level && mirror.Layout == layout {
			bz, err := os.ReadFile(filepath.Join(s.dir, fmt.Sprintf("REDACTED", level)))
			if err != nil {
				return nil, err
			}
			return octetSegment(bz, segment), nil
		}
	}
	return nil, nil
}

//
func octetSegment(bz []byte, ordinal uint32) []byte {
	begin := int(ordinal * mirrorSegmentVolume)
	end := int((ordinal + 1) * mirrorSegmentVolume)
	switch {
	case begin >= len(bz):
		return nil
	case end >= len(bz):
		return bz[begin:]
	default:
		return bz[begin:end]
	}
}

//
func octetSegments(bz []byte) uint32 {
	return uint32(math.Ceil(float64(len(bz)) / mirrorSegmentVolume))
}
