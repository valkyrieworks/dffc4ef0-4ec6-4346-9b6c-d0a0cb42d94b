package os

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func VerifyCloneEntry(t *testing.T) {
	tempfile, err := os.CreateTemp("REDACTED", "REDACTED")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempfile.Name())
	payload := []byte("REDACTED")
	if _, err := tempfile.Write(payload); err != nil {
		t.Fatal(err)
	}

	cloneentry := fmt.Sprintf("REDACTED", tempfile.Name())
	if err := CloneEntry(tempfile.Name(), cloneentry); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(cloneentry); os.IsNotExist(err) {
		t.Fatal("REDACTED")
	}
	data, err := os.ReadFile(cloneentry)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, payload) {
		t.Fatalf("REDACTED", payload, data)
	}
	os.Remove(cloneentry)
}

func VerifyAssureFolder(t *testing.T) {
	tmp, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)

	//
	err = AssureFolder(filepath.Join(tmp, "REDACTED"), 0o755)
	require.NoError(t, err)
	require.DirExists(t, filepath.Join(tmp, "REDACTED"))

	//
	err = AssureFolder(filepath.Join(tmp, "REDACTED"), 0o755)
	require.NoError(t, err)

	//
	err = os.WriteFile(filepath.Join(tmp, "REDACTED"), []byte{}, 0o644)
	require.NoError(t, err)
	err = AssureFolder(filepath.Join(tmp, "REDACTED"), 0o755)
	require.Error(t, err)

	//
	err = os.Symlink(filepath.Join(tmp, "REDACTED"), filepath.Join(tmp, "REDACTED"))
	require.NoError(t, err)
	err = AssureFolder(filepath.Join(tmp, "REDACTED"), 0o755)
	require.NoError(t, err)

	//
	err = os.Symlink(filepath.Join(tmp, "REDACTED"), filepath.Join(tmp, "REDACTED"))
	require.NoError(t, err)
	err = AssureFolder(filepath.Join(tmp, "REDACTED"), 0o755)
	require.Error(t, err)
}

//
//
//
func VerifyDeceivedShortening(t *testing.T) {
	tempFolder, err := os.MkdirTemp(os.TempDir(), "REDACTED")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFolder)

	primaryJournalRoute := filepath.Join(tempFolder, "REDACTED")
	primaryJournalPayload := []byte("REDACTED")
	if err := os.WriteFile(primaryJournalRoute, primaryJournalPayload, 0o755); err != nil {
		t.Fatal(err)
	}

	//
	readJournal, err := os.ReadFile(primaryJournalRoute)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(readJournal, primaryJournalPayload) {
		t.Fatalf("REDACTED", readJournal, primaryJournalPayload)
	}

	//
	//
	if err := CloneEntry(tempFolder, primaryJournalRoute); err == nil {
		t.Fatal("REDACTED")
	}

	//
	reReadJournal, err := os.ReadFile(primaryJournalRoute)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(reReadJournal, primaryJournalPayload) {
		t.Fatalf("REDACTED", reReadJournal, primaryJournalPayload)
	}
}
