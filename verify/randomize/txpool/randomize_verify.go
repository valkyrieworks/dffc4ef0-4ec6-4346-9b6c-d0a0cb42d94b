package respond_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/randomize/txpool"
)

const verifydataScenariosPath = "REDACTED"

func VerifyTxpoolVerifydataScenarios(t *testing.T) {
	listings, err := os.ReadDir(verifydataScenariosPath)
	require.NoError(t, err)

	for _, e := range listings {
		record := e
		t.Run(record.Name(), func(t *testing.T) {
			defer func() {
				r := recover()
				require.Nilf(t, r, "REDACTED")
			}()
			f, err := os.Open(filepath.Join(verifydataScenariosPath, record.Name()))
			require.NoError(t, err)
			influx, err := io.ReadAll(f)
			require.NoError(t, err)
			txpooll.Randomize(influx)
		})
	}
}
