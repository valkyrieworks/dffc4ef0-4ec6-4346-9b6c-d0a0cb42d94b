package pex

import (
	"testing"

	"github.com/valkyrieworks/p2p"
)

func Benchmarkaddressbook_digest(b *testing.B) {
	registry := &addressRegistry{
		ourLocations:          make(map[string]struct{}),
		internalIDXDatastore:        make(map[p2p.ID]struct{}),
		addressSearch:        make(map[p2p.ID]*recognizedLocation),
		flawedNodes:          make(map[p2p.ID]*recognizedLocation),
		entryRoute:          "REDACTED",
		forwardingPrecise: true,
	}
	registry.init()
	msg := []byte("REDACTED")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = registry.digest(msg)
	}
}
