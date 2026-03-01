package scratchfile

//

import (
	"bytes"
	"fmt"
	"os"
	testing "testing"

	"github.com/stretchr/testify/require"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

func VerifyPersistRecordIndivisible(t *testing.T) {
	var (
		data             = []byte(commitrand.Str(commitrand.Integern(2048)))
		old              = commitrand.Octets(commitrand.Integern(2048))
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

	if err = PersistRecordIndivisible(f.Name(), data, mode); err != nil {
		t.Fatal(err)
	}

	readerData, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, readerData) {
		t.Fatalf("REDACTED", data, readerData)
	}

	summary, err := os.Stat(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	if possess, desire := summary.Mode().Perm(), mode; possess != desire {
		t.Errorf("REDACTED", possess, desire)
	}
}

//
//
func VerifyPersistRecordIndivisibleReplicatedRecord(t *testing.T) {
	var (
		fallbackGerm    uint64 = 1
		verifyText            = "REDACTED"
		anticipatedText        = "REDACTED"

		recordTowardPersist = "REDACTED"
	)
	//
	indivisiblePersistRecordArbitrary = fallbackGerm
	initialRecordArbitrary := arbitraryPersistRecordEnding()
	indivisiblePersistRecordArbitrary = fallbackGerm
	filename := "REDACTED" + indivisiblePersistRecordHeading + initialRecordArbitrary
	f, err := os.OpenFile(filename, indivisiblePersistRecordMarker, 0o777)
	defer os.Remove(filename)
	//
	defer os.Remove(recordTowardPersist)

	require.NoError(t, err)
	_, err = f.WriteString(verifyText)
	require.NoError(t, err)
	err = PersistRecordIndivisible(recordTowardPersist, []byte(anticipatedText), 0o777)
	require.NoError(t, err)
	//
	initialIndivisibleRecordOctets, err := os.ReadFile(filename)
	require.NoError(t, err, "REDACTED")
	require.Equal(t, []byte(verifyText), initialIndivisibleRecordOctets, "REDACTED")
	//
	ensuingRecordOctets, err := os.ReadFile(recordTowardPersist)
	require.NoError(t, err, "REDACTED")
	require.Equal(t, []byte(anticipatedText), ensuingRecordOctets, "REDACTED")

	//
	//
	indivisiblePersistRecordArbitrary = fallbackGerm
	_ = arbitraryPersistRecordEnding()
	ordinalRecordArbitrary := arbitraryPersistRecordEnding()
	_, err = os.Stat("REDACTED" + indivisiblePersistRecordHeading + ordinalRecordArbitrary)
	require.True(t, os.IsNotExist(err), "REDACTED")
}

//
//
//
func VerifyPersistRecordIndivisibleMultipleReplicates(t *testing.T) {
	var (
		fallbackGerm    uint64 = 2
		verifyText            = "REDACTED"
		anticipatedText        = "REDACTED"

		recordTowardPersist = "REDACTED"
	)
	//
	indivisiblePersistRecordArbitrary = fallbackGerm
	for i := 0; i < indivisiblePersistRecordMaximumCountDisagreements+2; i++ {
		recordArbitrary := arbitraryPersistRecordEnding()
		filename := "REDACTED" + indivisiblePersistRecordHeading + recordArbitrary
		f, err := os.OpenFile(filename, indivisiblePersistRecordMarker, 0o777)
		require.Nil(t, err)
		_, err = fmt.Fprintf(f, verifyText, i)
		require.NoError(t, err)
		defer os.Remove(filename)
	}

	indivisiblePersistRecordArbitrary = fallbackGerm
	//
	defer os.Remove(recordTowardPersist)

	err := PersistRecordIndivisible(recordTowardPersist, []byte(anticipatedText), 0o777)
	require.NoError(t, err)
	//
	indivisiblePersistRecordArbitrary = fallbackGerm
	for i := 0; i < indivisiblePersistRecordMaximumCountDisagreements+2; i++ {
		recordArbitrary := arbitraryPersistRecordEnding()
		filename := "REDACTED" + indivisiblePersistRecordHeading + recordArbitrary
		initialIndivisibleRecordOctets, err := os.ReadFile(filename)
		require.Nil(t, err, "REDACTED")
		require.Equal(t, []byte(fmt.Sprintf(verifyText, i)), initialIndivisibleRecordOctets,
			"REDACTED", i)
	}

	//
	ensuingRecordOctets, err := os.ReadFile(recordTowardPersist)
	require.Nil(t, err, "REDACTED")
	require.Equal(t, []byte(anticipatedText), ensuingRecordOctets, "REDACTED")
}
