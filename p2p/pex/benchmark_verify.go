package pex

import (
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

func Assessmentaddrbook_digest(b *testing.B) {
	register := &locationRegister{
		mineLocations:          make(map[string]struct{}),
		secludedIDXDstore:        make(map[p2p.ID]struct{}),
		locationSearch:        make(map[p2p.ID]*recognizedLocator),
		flawedNodes:          make(map[p2p.ID]*recognizedLocator),
		recordRoute:          "REDACTED",
		reachabilityStringent: true,
	}
	register.initialize()
	msg := []byte("REDACTED")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = register.digest(msg)
	}
}
