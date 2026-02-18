package autofile

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cometos "github.com/valkyrieworks/utils/os"
)

func VerifySighup(t *testing.T) {
	origFolder, err := os.Getwd()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := os.Chdir(origFolder); err != nil {
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
	label := "REDACTED"
	af, err := AccessAutomaticEntry(label)
	require.NoError(t, err)
	require.True(t, filepath.IsAbs(af.Route))

	//
	_, err = af.Record([]byte("REDACTED"))
	require.NoError(t, err)
	_, err = af.Record([]byte("REDACTED"))
	require.NoError(t, err)

	//
	err = os.Rename(label, label+"REDACTED")
	require.NoError(t, err)

	//
	anotherFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(anotherFolder)
	err = os.Chdir(anotherFolder)
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
	err = af.End()
	require.NoError(t, err)

	//
	if content := cometos.ShouldReaderEntry(filepath.Join(dir, label+"REDACTED")); string(content) != "REDACTED" {
		t.Errorf("REDACTED", content)
	}
	if content := cometos.ShouldReaderEntry(filepath.Join(dir, label)); string(content) != "REDACTED" {
		t.Errorf("REDACTED", content)
	}

	//
	entries, err := os.ReadDir("REDACTED")
	require.NoError(t, err)
	assert.Empty(t, entries)
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

func VerifyAutomaticEntryVolume(t *testing.T) {
	//
	f, err := os.CreateTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	err = f.Close()
	require.NoError(t, err)

	//
	af, err := AccessAutomaticEntry(f.Name())
	require.NoError(t, err)

	//
	volume, err := af.Volume()
	require.Zero(t, volume)
	require.NoError(t, err)

	//
	data := []byte("REDACTED")
	_, err = af.Record(data)
	require.NoError(t, err)
	volume, err = af.Volume()
	require.EqualValues(t, len(data), volume)
	require.NoError(t, err)

	//
	err = af.End()
	require.NoError(t, err)
	err = os.Remove(f.Name())
	require.NoError(t, err)
	volume, err = af.Volume()
	require.EqualValues(t, 0, volume, "REDACTED")
	require.NoError(t, err)

	//
	_ = os.Remove(f.Name())
}
