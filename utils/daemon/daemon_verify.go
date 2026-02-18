package daemon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type verifyDaemon struct {
	RootDaemon
}

func (verifyDaemon) OnRestore() error {
	return nil
}

func VerifyRootDaemonWait(t *testing.T) {
	ts := &verifyDaemon{}
	ts.RootDaemon = *NewRootDaemon(nil, "REDACTED", ts)
	err := ts.Begin()
	require.NoError(t, err)

	waitCompleted := make(chan struct{})
	go func() {
		ts.Wait()
		waitCompleted <- struct{}{}
	}()

	go ts.Halt() //

	select {
	case <-waitCompleted:
		//
	case <-time.After(100 * time.Millisecond):
		t.Fatal("REDACTED")
	}
}

func VerifyRootDaemonRestore(t *testing.T) {
	ts := &verifyDaemon{}
	ts.RootDaemon = *NewRootDaemon(nil, "REDACTED", ts)
	err := ts.Begin()
	require.NoError(t, err)

	err = ts.Restore()
	require.Error(t, err, "REDACTED")

	err = ts.Halt()
	require.NoError(t, err)

	err = ts.Restore()
	require.NoError(t, err)

	err = ts.Begin()
	require.NoError(t, err)
}
