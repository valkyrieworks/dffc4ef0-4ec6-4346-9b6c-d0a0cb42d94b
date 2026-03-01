package autosave

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
)

func VerifySignalhup(t *testing.T) {
	sourcePath, err := os.Getwd()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := os.Chdir(sourcePath); err != nil {
			t.Error(err)
		}
	})

	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})
	err = os.Chdir(dir)
	require.NoError(t, err)

	//
	alias := "REDACTED"
	af, err := UnlockAutomaticRecord(alias)
	require.NoError(t, err)
	require.True(t, filepath.IsAbs(af.Route))

	//
	_, err = af.Record([]byte("REDACTED"))
	require.NoError(t, err)
	_, err = af.Record([]byte("REDACTED"))
	require.NoError(t, err)

	//
	err = os.Rename(alias, alias+"REDACTED")
	require.NoError(t, err)

	//
	anotherPath, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(anotherPath)
	err = os.Chdir(anotherPath)
	require.NoError(t, err)

	//
	err = syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	require.NoError(t, err)

	//
	time.Sleep(time.Millisecond * 10)

	//
	_, err = af.Record([]byte("REDACTED"))
	require.NoError(t, err)
	_, err = af.Record([]byte("REDACTED"))
	require.NoError(t, err)
	err = af.Shutdown()
	require.NoError(t, err)

	//
	if content := strongos.ShouldFetchRecord(filepath.Join(dir, alias+"REDACTED")); string(content) != "REDACTED" {
		t.Errorf("REDACTED", content)
	}
	if content := strongos.ShouldFetchRecord(filepath.Join(dir, alias)); string(content) != "REDACTED" {
		t.Errorf("REDACTED", content)
	}

	//
	records, err := os.ReadDir("REDACTED")
	require.NoError(t, err)
	assert.Empty(t, records)
}

//
//
//
//
//
//
//
//

//
//
//
//
//
//
//

//
//
//
//
//
//
//
//
//

func VerifyAutomaticRecordExtent(t *testing.T) {
	//
	f, err := os.CreateTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	err = f.Close()
	require.NoError(t, err)

	//
	af, err := UnlockAutomaticRecord(f.Name())
	require.NoError(t, err)

	//
	extent, err := af.Extent()
	require.Zero(t, extent)
	require.NoError(t, err)

	//
	data := []byte("REDACTED")
	_, err = af.Record(data)
	require.NoError(t, err)
	extent, err = af.Extent()
	require.EqualValues(t, len(data), extent)
	require.NoError(t, err)

	//
	err = af.Shutdown()
	require.NoError(t, err)
	err = os.Remove(f.Name())
	require.NoError(t, err)
	extent, err = af.Extent()
	require.EqualValues(t, 0, extent, "REDACTED")
	require.NoError(t, err)

	//
	_ = os.Remove(f.Name())
}
