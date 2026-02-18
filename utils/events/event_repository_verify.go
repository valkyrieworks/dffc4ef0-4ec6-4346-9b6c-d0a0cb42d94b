package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Verifyeventshard_Purge(t *testing.T) {
	evsw := NewEventRouter()
	err := evsw.Begin()
	require.NoError(t, err)

	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED", func(data EventData) {
		//
		//
		require.FailNow(t, "REDACTED")
	})
	require.NoError(t, err)

	evc := NewEventRepository(evsw)
	evc.Purge()
	//
	evc.Purge()
	abort := true
	pass := false
	err = evsw.AppendObserverForEvent("REDACTED", "REDACTED", func(data EventData) {
		if abort {
			require.FailNow(t, "REDACTED")
		}
		pass = true
	})
	require.NoError(t, err)

	evc.TriggerEvent("REDACTED", struct{ int }{1})
	evc.TriggerEvent("REDACTED", struct{ int }{2})
	evc.TriggerEvent("REDACTED", struct{ int }{3})
	abort = false
	evc.Purge()
	assert.True(t, pass)
}
