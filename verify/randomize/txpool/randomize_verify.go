package respond_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	txpool "github.com/valkyrieworks/verify/randomize/txpool"
)

const mockdataScenariosFolder = "REDACTED"

func VerifyTxpoolMockdataScenarios(t *testing.T) {
	records, err := os.ReadDir(mockdataScenariosFolder)
	require.NoError(t, err)

	for _, e := range records {
		entry := e
		t.Run(entry.Name(), func(t *testing.T) {
			defer func() {
				r := recover()
				require.Nilf(t, r, "REDACTED")
			}()
			f, err := os.Open(filepath.Join(mockdataScenariosFolder, entry.Name()))
			require.NoError(t, err)
			influx, err := io.ReadAll(f)
			require.NoError(t, err)
			txpool.Randomize(influx)
		})
	}
}
