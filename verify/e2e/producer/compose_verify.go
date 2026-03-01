//

package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
)

//
func TestGenerator(t *testing.T) {
	cfg := &generateConfig{
		randSource: rand.New(rand.NewSource(randomSeed)),
	}
	manifests, err := Generate(cfg)
	require.NoError(t, err)

	for idx, m := range manifests {
		t.Run(fmt.Sprintf("REDACTED", idx), func(t *testing.T) {
			infra, err := e2e.NewDockerInfrastructureData(m)
			require.NoError(t, err)
			_, err = e2e.NewTestnetFromManifest(m, filepath.Join(t.TempDir(), fmt.Sprintf("REDACTED", idx)), infra)
			require.NoError(t, err)
		})
	}
}
