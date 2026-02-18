package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

func VerifyRiskyCallOrigins(t *testing.T) {
	sw := p2p.CreateRouter(cfg.StandardP2PSettings(), 1,
		func(n int, sw *p2p.Router) *p2p.Router { return sw })
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	env := &Context{}
	env.Tracer = log.VerifyingTracer()
	env.P2PNodes = sw

	verifyScenarios := []struct {
		origins []string
		isErr bool
	}{
		{[]string{}, true},
		{[]string{"REDACTED"}, false},
		{[]string{"REDACTED"}, true},
	}

	for _, tc := range verifyScenarios {
		res, err := env.RiskyCallOrigins(&rpctypes.Context{}, tc.origins)
		if tc.isErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
		}
	}
}

func VerifyRiskyCallNodes(t *testing.T) {
	sw := p2p.CreateRouter(cfg.StandardP2PSettings(), 1,
		func(n int, sw *p2p.Router) *p2p.Router { return sw })
	sw.CollectionAddressRegistry(&p2p.AddressRegistryEmulate{
		Locations:        make(map[string]struct{}),
		OurLocations:     make(map[string]struct{}),
		InternalLocations: make(map[string]struct{}),
	})
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	env := &Context{}
	env.Tracer = log.VerifyingTracer()
	env.P2PNodes = sw

	verifyScenarios := []struct {
		nodes                               []string
		durability, absolute, internal bool
		isErr                               bool
	}{
		{[]string{}, false, false, false, true},
		{[]string{"REDACTED"}, true, true, true, false},
		{[]string{"REDACTED"}, true, true, false, true},
	}

	for _, tc := range verifyScenarios {
		res, err := env.RiskyCallNodes(&rpctypes.Context{}, tc.nodes, tc.durability, tc.absolute, tc.internal)
		if tc.isErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
		}
	}
}
