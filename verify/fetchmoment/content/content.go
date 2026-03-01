package content

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
	tokenHeading      = "REDACTED"
	maximumContentExtent = 4 * 1024 * 1024
)

//
//
//
func FreshOctets(p *Content) ([]byte, error) {
	p.Filling = make([]byte, 1)
	if p.Moment == nil {
		p.Moment = timestamppb.Now()
	}
	us, err := ComputeUnfilledExtent(p)
	if err != nil {
		return nil, err
	}
	if p.Extent > maximumContentExtent {
		return nil, fmt.Errorf("REDACTED", p.Extent, maximumContentExtent)
	}
	paramExtent := int(p.Extent) //
	if paramExtent < us {
		return nil, fmt.Errorf("REDACTED", paramExtent, us)
	}

	//
	p.Filling = make([]byte, (paramExtent-us)/2)
	_, err = rand.Read(p.Filling)
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
	return append([]byte(tokenHeading), h...), nil
}

//
//
//
func OriginatingOctets(b []byte) (*Content, error) {
	trH := bytes.TrimPrefix(b, []byte(tokenHeading))
	if bytes.Equal(b, trH) {
		return nil, fmt.Errorf("REDACTED", tokenHeading)
	}
	trB, err := hex.DecodeString(string(trH))
	if err != nil {
		return nil, err
	}

	p := &Content{}
	err = proto.Unmarshal(trB, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

//
//
func MaximumUnfilledExtent() (int, error) {
	p := &Content{
		Moment:        timestamppb.Now(),
		Linkages: math.MaxUint64,
		Frequency:        math.MaxUint64,
		Extent:        math.MaxUint64,
		Filling:     make([]byte, 1),
	}
	return ComputeUnfilledExtent(p)
}

//
//
//
func ComputeUnfilledExtent(p *Content) (int, error) {
	if len(p.Filling) != 1 {
		return 0, fmt.Errorf("REDACTED", len(p.Filling))
	}
	b, err := proto.Marshal(p)
	if err != nil {
		return 0, err
	}
	h := []byte(hex.EncodeToString(b))
	return len(h) + len(tokenHeading), nil
}
