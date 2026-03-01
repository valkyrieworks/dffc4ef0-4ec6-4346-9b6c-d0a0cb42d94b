package statuschronize

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

var (
	_ p2p.Encapsulator = &SegmentSolicit{}
	_ p2p.Encapsulator = &SegmentReply{}
	_ p2p.Encapsulator = &ImagesSolicit{}
	_ p2p.Encapsulator = &ImagesReply{}
)

func (m *ImagesReply) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Artifact_Imagesreply{ImagesReply: m}
	return sm
}

func (m *ImagesSolicit) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Artifact_Imagessolicit{ImagesSolicit: m}
	return sm
}

func (m *SegmentReply) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Artifact_Fragmentreply{SegmentReply: m}
	return sm
}

func (m *SegmentSolicit) Enclose() proto.Message {
	sm := &Signal{}
	sm.Sum = &Artifact_Fragmentsolicit{SegmentSolicit: m}
	return sm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Artifact_Fragmentsolicit:
		return m.ObtainSegmentSolicit(), nil

	case *Artifact_Fragmentreply:
		return m.ObtainSegmentReply(), nil

	case *Artifact_Imagessolicit:
		return m.ObtainImagesSolicit(), nil

	case *Artifact_Imagesreply:
		return m.ObtainImagesReply(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
