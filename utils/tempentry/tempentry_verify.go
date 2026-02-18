package tempentry

//

import (
	"bytes"
	"fmt"
	"os"
	testing "testing"

	"github.com/stretchr/testify/require"

	engineseed "github.com/valkyrieworks/utils/random"
)

func VerifyRecordEntryAtomic(t *testing.T) {
	var (
		data             = []byte(engineseed.Str(engineseed.Intn(2048)))
		old              = engineseed.Octets(engineseed.Intn(2048))
		mode os.FileMode = 0o600
	)

	f, err := os.CreateTemp("REDACTED", "REDACTED")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	if err = os.WriteFile(f.Name(), old, 0o600); err != nil {
		t.Fatal(err)
	}

	if err = RecordEntryAtomic(f.Name(), data, mode); err != nil {
		t.Fatal(err)
	}

	readerData, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, readerData) {
		t.Fatalf("REDACTED", data, readerData)
	}

	status, err := os.Stat(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	if possess, desire := status.Mode().Perm(), mode; possess != desire {
		t.Errorf("REDACTED", possess, desire)
	}
}

//
//
func VerifyRecordEntryAtomicReplicatedEntry(t *testing.T) {
	var (
		standardOrigin    uint64 = 1
		verifyString            = "REDACTED"
		anticipatedString        = "REDACTED"

		entryToRecord = "REDACTED"
	)
	//
	atomicRecordEntryRandom = standardOrigin
	initialEntryRandom := randomRecordEntryPostfix()
	atomicRecordEntryRandom = standardOrigin
	filename := "REDACTED" + atomicRecordEntryPrefix + initialEntryRandom
	f, err := os.OpenFile(filename, atomicRecordEntryMark, 0o777)
	defer os.Remove(filename)
	//
	defer os.Remove(entryToRecord)

	require.NoError(t, err)
	_, err = f.WriteString(verifyString)
	require.NoError(t, err)
	err = RecordEntryAtomic(entryToRecord, []byte(anticipatedString), 0o777)
	require.NoError(t, err)
	//
	initialAtomicEntryOctets, err := os.ReadFile(filename)
	require.NoError(t, err, "REDACTED")
	require.Equal(t, []byte(verifyString), initialAtomicEntryOctets, "REDACTED")
	//
	ensuingEntryOctets, err := os.ReadFile(entryToRecord)
	require.NoError(t, err, "REDACTED")
	require.Equal(t, []byte(anticipatedString), ensuingEntryOctets, "REDACTED")

	//
	//
	atomicRecordEntryRandom = standardOrigin
	_ = randomRecordEntryPostfix()
	momentEntryRandom := randomRecordEntryPostfix()
	_, err = os.Stat("REDACTED" + atomicRecordEntryPrefix + momentEntryRandom)
	require.True(t, os.IsNotExist(err), "REDACTED")
}

//
//
//
func VerifyRecordEntryAtomicNumerousReplicates(t *testing.T) {
	var (
		standardOrigin    uint64 = 2
		verifyString            = "REDACTED"
		anticipatedString        = "REDACTED"

		entryToRecord = "REDACTED"
	)
	//
	atomicRecordEntryRandom = standardOrigin
	for i := 0; i < atomicRecordEntryMaximumCountClashes+2; i++ {
		entryRandom := randomRecordEntryPostfix()
		filename := "REDACTED" + atomicRecordEntryPrefix + entryRandom
		f, err := os.OpenFile(filename, atomicRecordEntryMark, 0o777)
		require.Nil(t, err)
		_, err = fmt.Fprintf(f, verifyString, i)
		require.NoError(t, err)
		defer os.Remove(filename)
	}

	atomicRecordEntryRandom = standardOrigin
	//
	defer os.Remove(entryToRecord)

	err := RecordEntryAtomic(entryToRecord, []byte(anticipatedString), 0o777)
	require.NoError(t, err)
	//
	atomicRecordEntryRandom = standardOrigin
	for i := 0; i < atomicRecordEntryMaximumCountClashes+2; i++ {
		entryRandom := randomRecordEntryPostfix()
		filename := "REDACTED" + atomicRecordEntryPrefix + entryRandom
		initialAtomicEntryOctets, err := os.ReadFile(filename)
		require.Nil(t, err, "REDACTED")
		require.Equal(t, []byte(fmt.Sprintf(verifyString, i)), initialAtomicEntryOctets,
			"REDACTED", i)
	}

	//
	ensuingEntryOctets, err := os.ReadFile(entryToRecord)
	require.Nil(t, err, "REDACTED")
	require.Equal(t, []byte(anticipatedString), ensuingEntryOctets, "REDACTED")
}
