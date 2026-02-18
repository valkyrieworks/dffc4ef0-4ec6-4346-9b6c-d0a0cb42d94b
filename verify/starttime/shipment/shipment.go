package shipment

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"

	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	keyPrefix      = "REDACTED"
	maximumShipmentVolume = 4 * 1024 * 1024
)

//
//
//
func NewOctets(p *Shipment) ([]byte, error) {
	p.Stuffing = make([]byte, 1)
	if p.Time == nil {
		p.Time = timestamppb.Now()
	}
	us, err := ComputeUnfilledVolume(p)
	if err != nil {
		return nil, err
	}
	if p.Volume > maximumShipmentVolume {
		return nil, fmt.Errorf("REDACTED", p.Volume, maximumShipmentVolume)
	}
	pVolume := int(p.Volume) //
	if pVolume < us {
		return nil, fmt.Errorf("REDACTED", pVolume, us)
	}

	//
	p.Stuffing = make([]byte, (pVolume-us)/2)
	_, err = rand.Read(p.Stuffing)
	if err != nil {
		return nil, err
	}
	b, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	h := []byte(hex.EncodeToString(b))

	//
	//
	return append([]byte(keyPrefix), h...), nil
}

//
//
//
func FromOctets(b []byte) (*Shipment, error) {
	trH := bytes.TrimPrefix(b, []byte(keyPrefix))
	if bytes.Equal(b, trH) {
		return nil, fmt.Errorf("REDACTED", keyPrefix)
	}
	trB, err := hex.DecodeString(string(trH))
	if err != nil {
		return nil, err
	}

	p := &Shipment{}
	err = proto.Unmarshal(trB, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

//
//
func MaximumUnfilledVolume() (int, error) {
	p := &Shipment{
		Time:        timestamppb.Now(),
		Linkages: math.MaxUint64,
		Ratio:        math.MaxUint64,
		Volume:        math.MaxUint64,
		Stuffing:     make([]byte, 1),
	}
	return ComputeUnfilledVolume(p)
}

//
//
//
func ComputeUnfilledVolume(p *Shipment) (int, error) {
	if len(p.Stuffing) != 1 {
		return 0, fmt.Errorf("REDACTED", len(p.Stuffing))
	}
	b, err := proto.Marshal(p)
	if err != nil {
		return 0, err
	}
	h := []byte(hex.EncodeToString(b))
	return len(h) + len(keyPrefix), nil
}
