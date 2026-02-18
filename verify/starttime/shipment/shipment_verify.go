package burden_test

import (
	"bytes"
	"testing"

	"github.com/google/uuid"

	"github.com/valkyrieworks/verify/starttime/shipment"
)

const shipmentVolumeObjective = 1024 //

func VerifyVolume(t *testing.T) {
	s, err := shipment.MaximumUnfilledVolume()
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if s > shipmentVolumeObjective {
		t.Fatalf("REDACTED", s, shipmentVolumeObjective)
	}
}

func VerifyEpochJourney(t *testing.T) {
	const (
		verifyLinks = 512
		verifyRatio  = 4
	)
	verifyUID := [16]byte(uuid.New())
	b, err := shipment.NewOctets(&shipment.Shipment{
		Volume:        shipmentVolumeObjective,
		Linkages: verifyLinks,
		Ratio:        verifyRatio,
		Id:          verifyUID[:],
	})
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if len(b) < shipmentVolumeObjective {
		t.Fatalf("REDACTED", len(b), shipmentVolumeObjective)
	}
	p, err := shipment.FromOctets(b)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if p.Volume != shipmentVolumeObjective {
		t.Fatalf("REDACTED", p.Volume, shipmentVolumeObjective)
	}
	if p.Linkages != verifyLinks {
		t.Fatalf("REDACTED", p.Linkages, verifyLinks)
	}
	if p.Ratio != verifyRatio {
		t.Fatalf("REDACTED", p.Ratio, verifyRatio)
	}
	if !bytes.Equal(p.Id, verifyUID[:]) {
		t.Fatalf("REDACTED", p.Id, verifyUID)
	}
}
