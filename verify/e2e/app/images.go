package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sync"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

const (
	imageSegmentExtent = 1e6

	//
	maximumImageTally = 10
)

//
//
//
type ImageDepot struct {
	sync.ReadwriteExclusion
	dir      string
	attributes []iface.Image
}

//
func FreshImageDepot(dir string) (*ImageDepot, error) {
	depot := &ImageDepot{dir: dir}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	if err := depot.fetchAttributes(); err != nil {
		return nil, err
	}
	return depot, nil
}

//
//
func (s *ImageDepot) fetchAttributes() error {
	record := filepath.Join(s.dir, "REDACTED")
	attributes := []iface.Image{}

	bz, err := os.ReadFile(record)
	switch {
	case errors.Is(err, os.ErrNotExist):
	case err != nil:
		return fmt.Errorf("REDACTED", record, err)
	}
	if len(bz) != 0 {
		err = json.Unmarshal(bz, &attributes)
		if err != nil {
			return fmt.Errorf("REDACTED", record, err)
		}
	}
	s.attributes = attributes
	return nil
}

//
//
func (s *ImageDepot) persistAttributes() error {
	bz, err := json.Marshal(s.attributes)
	if err != nil {
		return err
	}

	//
	freshRecord := filepath.Join(s.dir, "REDACTED")
	record := filepath.Join(s.dir, "REDACTED")
	err = os.WriteFile(freshRecord, bz, 0o644) //
	if err != nil {
		return err
	}
	return os.Rename(freshRecord, record)
}

//
func (s *ImageDepot) Generate(status *Status) (iface.Image, error) {
	s.Lock()
	defer s.Unlock()
	bz, altitude, statusDigest, err := status.Disclose()
	if err != nil {
		return iface.Image{}, err
	}
	image := iface.Image{
		Altitude: altitude,
		Layout: 1,
		Digest:   statusDigest,
		Segments: octetSegments(bz),
	}
	err = os.WriteFile(filepath.Join(s.dir, fmt.Sprintf("REDACTED", altitude)), bz, 0o644) //
	if err != nil {
		return iface.Image{}, err
	}
	s.attributes = append(s.attributes, image)
	err = s.persistAttributes()
	if err != nil {
		return iface.Image{}, err
	}
	return image, nil
}

//
func (s *ImageDepot) Trim(n int) error {
	s.Lock()
	defer s.Unlock()
	//
	//
	i := 0
	for ; i < len(s.attributes)-n; i++ {
		h := s.attributes[i].Altitude
		if err := os.Remove(filepath.Join(s.dir, fmt.Sprintf("REDACTED", h))); err != nil {
			return err
		}
	}

	//
	trimmed := make([]iface.Image, len(s.attributes[i:]))
	copy(trimmed, s.attributes[i:])
	s.attributes = trimmed
	return nil
}

//
func (s *ImageDepot) Catalog() ([]*iface.Image, error) {
	s.RLock()
	defer s.RUnlock()
	images := make([]*iface.Image, len(s.attributes))
	for idx := range s.attributes {
		images[idx] = &s.attributes[idx]
	}
	return images, nil
}

//
func (s *ImageDepot) FetchSegment(altitude uint64, layout uint32, segment uint32) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()
	for _, image := range s.attributes {
		if image.Altitude == altitude && image.Layout == layout {
			bz, err := os.ReadFile(filepath.Join(s.dir, fmt.Sprintf("REDACTED", altitude)))
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
	initiate := int(ordinal * imageSegmentExtent)
	end := int((ordinal + 1) * imageSegmentExtent)
	switch {
	case initiate >= len(bz):
		return nil
	case end >= len(bz):
		return bz[initiate:]
	default:
		return bz[initiate:end]
	}
}

//
func octetSegments(bz []byte) uint32 {
	return uint32(math.Ceil(float64(len(bz)) / imageSegmentExtent))
}
