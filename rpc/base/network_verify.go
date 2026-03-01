package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

func VerifyInsecureCallOrigins(t *testing.T) {
	sw := p2p.CreateRouter(cfg.FallbackPeer2peerSettings(), 1,
		func(n int, sw *p2p.Router) *p2p.Router { return sw })
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	env := &Context{}
	env.Tracer = log.VerifyingTracer()
	env.Peer2peerNodes = sw

	verifyScenarios := []struct {
		origins []string
		equalsFault bool
	}{
		{[]string{}, true},
		{[]string{"REDACTED"}, false},
		{[]string{"REDACTED"}, true},
	}

	for _, tc := range verifyScenarios {
		res, err := env.InsecureCallOrigins(&remoteifacetypes.Env{}, tc.origins)
		if tc.equalsFault {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
		}
	}
}

func VerifyInsecureCallNodes(t *testing.T) {
	sw := p2p.CreateRouter(cfg.FallbackPeer2peerSettings(), 1,
		func(n int, sw *p2p.Router) *p2p.Router { return sw })
	sw.AssignLocationRegister(&p2p.LocationRegisterSimulate{
		Locations:        make(map[string]struct{}),
		MineLocations:     make(map[string]struct{}),
		SecludedLocations: make(map[string]struct{}),
	})
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	env := &Context{}
	env.Tracer = log.VerifyingTracer()
	env.Peer2peerNodes = sw

	verifyScenarios := []struct {
		nodes                               []string
		storage, absolute, secluded bool
		equalsFault                               bool
	}{
		{[]string{}, false, false, false, true},
		{[]string{"REDACTED"}, true, true, true, false},
		{[]string{"REDACTED"}, true, true, false, true},
	}

	for _, tc := range verifyScenarios {
		res, err := env.InsecureCallNodes(&remoteifacetypes.Env{}, tc.nodes, tc.storage, tc.absolute, tc.secluded)
		if tc.equalsFault {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
		}
	}
}
