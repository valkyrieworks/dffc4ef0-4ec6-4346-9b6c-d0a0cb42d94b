package os

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func VerifyDuplicateRecord(t *testing.T) {
	tempfile, err := os.CreateTemp("REDACTED", "REDACTED")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempfile.Name())
	subject := []byte("REDACTED")
	if _, err := tempfile.Write(subject); err != nil {
		t.Fatal(err)
	}

	duplicaterecord := fmt.Sprintf("REDACTED", tempfile.Name())
	if err := DuplicateRecord(tempfile.Name(), duplicaterecord); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(duplicaterecord); os.IsNotExist(err) {
		t.Fatal("REDACTED")
	}
	data, err := os.ReadFile(duplicaterecord)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, subject) {
		t.Fatalf("REDACTED", subject, data)
	}
	os.Remove(duplicaterecord)
}

func VerifyAssurePath(t *testing.T) {
	tmp, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)

	//
	err = AssurePath(filepath.Join(tmp, "REDACTED"), 0o755)
	require.NoError(t, err)
	require.DirExists(t, filepath.Join(tmp, "REDACTED"))

	//
	err = AssurePath(filepath.Join(tmp, "REDACTED"), 0o755)
	require.NoError(t, err)

	//
	err = os.WriteFile(filepath.Join(tmp, "REDACTED"), []byte{}, 0o644)
	require.NoError(t, err)
	err = AssurePath(filepath.Join(tmp, "REDACTED"), 0o755)
	require.Error(t, err)

	//
	err = os.Symlink(filepath.Join(tmp, "REDACTED"), filepath.Join(tmp, "REDACTED"))
	require.NoError(t, err)
	err = AssurePath(filepath.Join(tmp, "REDACTED"), 0o755)
	require.NoError(t, err)

	//
	err = os.Symlink(filepath.Join(tmp, "REDACTED"), filepath.Join(tmp, "REDACTED"))
	require.NoError(t, err)
	err = AssurePath(filepath.Join(tmp, "REDACTED"), 0o755)
	require.Error(t, err)
}

//
//
//
func VerifyDeceivedShortening(t *testing.T) {
	scratchPath, err := os.MkdirTemp(os.TempDir(), "REDACTED")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(scratchPath)

	authenticJournalRoute := filepath.Join(scratchPath, "REDACTED")
	authenticJournalSubject := []byte("REDACTED")
	if err := os.WriteFile(authenticJournalRoute, authenticJournalSubject, 0o755); err != nil {
		t.Fatal(err)
	}

	//
	fetchJournal, err := os.ReadFile(authenticJournalRoute)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(fetchJournal, authenticJournalSubject) {
		t.Fatalf("REDACTED", fetchJournal, authenticJournalSubject)
	}

	//
	//
	if err := DuplicateRecord(scratchPath, authenticJournalRoute); err == nil {
		t.Fatal("REDACTED")
	}

	//
	againFetchJournal, err := os.ReadFile(authenticJournalRoute)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(againFetchJournal, authenticJournalSubject) {
		t.Fatalf("REDACTED", againFetchJournal, authenticJournalSubject)
	}
}
