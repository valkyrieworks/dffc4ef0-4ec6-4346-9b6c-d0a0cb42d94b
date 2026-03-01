package facility

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type verifyFacility struct {
	FoundationFacility
}

func (verifyFacility) UponRestore() error {
	return nil
}

func VerifyFoundationFacilityPause(t *testing.T) {
	ts := &verifyFacility{}
	ts.FoundationFacility = *FreshFoundationFacility(nil, "REDACTED", ts)
	err := ts.Initiate()
	require.NoError(t, err)

	pauseConcluded := make(chan struct{})
	go func() {
		ts.Pause()
		pauseConcluded <- struct{}{}
	}()

	go ts.Halt() //

	select {
	case <-pauseConcluded:
		//
	case <-time.After(100 * time.Millisecond):
		t.Fatal("REDACTED")
	}
}

func VerifyFoundationFacilityRestore(t *testing.T) {
	ts := &verifyFacility{}
	ts.FoundationFacility = *FreshFoundationFacility(nil, "REDACTED", ts)
	err := ts.Initiate()
	require.NoError(t, err)

	err = ts.Restore()
	require.Error(t, err, "REDACTED")

	err = ts.Halt()
	require.NoError(t, err)

	err = ts.Restore()
	require.NoError(t, err)

	err = ts.Initiate()
	require.NoError(t, err)
}
