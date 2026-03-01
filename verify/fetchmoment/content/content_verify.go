package payld_test

import (
	"bytes"
	"testing"

	"github.com/google/uuid"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/fetchmoment/content"
)

const contentExtentObjective = 1024 //

func VerifyExtent(t *testing.T) {
	s, err := content.MaximumUnfilledExtent()
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if s > contentExtentObjective {
		t.Fatalf("REDACTED", s, contentExtentObjective)
	}
}

func VerifyIterationJourney(t *testing.T) {
	const (
		verifyLinks = 512
		verifyFrequency  = 4
	)
	verifyUUID := [16]byte(uuid.New())
	b, err := content.FreshOctets(&content.Content{
		Extent:        contentExtentObjective,
		Linkages: verifyLinks,
		Frequency:        verifyFrequency,
		Id:          verifyUUID[:],
	})
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if len(b) < contentExtentObjective {
		t.Fatalf("REDACTED", len(b), contentExtentObjective)
	}
	p, err := content.OriginatingOctets(b)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if p.Extent != contentExtentObjective {
		t.Fatalf("REDACTED", p.Extent, contentExtentObjective)
	}
	if p.Linkages != verifyLinks {
		t.Fatalf("REDACTED", p.Linkages, verifyLinks)
	}
	if p.Frequency != verifyFrequency {
		t.Fatalf("REDACTED", p.Frequency, verifyFrequency)
	}
	if !bytes.Equal(p.Id, verifyUUID[:]) {
		t.Fatalf("REDACTED", p.Id, verifyUUID)
	}
}
