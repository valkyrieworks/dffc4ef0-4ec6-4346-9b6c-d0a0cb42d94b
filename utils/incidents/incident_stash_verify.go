package incidents

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Verifyincidentstash_Purge(t *testing.T) {
	incidentctl := FreshIncidentRouter()
	err := incidentctl.Initiate()
	require.NoError(t, err)

	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED", func(data IncidentData) {
		//
		//
		require.FailNow(t, "REDACTED")
	})
	require.NoError(t, err)

	evc := FreshIncidentStash(incidentctl)
	evc.Purge()
	//
	evc.Purge()
	abort := true
	phrase := false
	err = incidentctl.AppendObserverForeachIncident("REDACTED", "REDACTED", func(data IncidentData) {
		if abort {
			require.FailNow(t, "REDACTED")
		}
		phrase = true
	})
	require.NoError(t, err)

	evc.TriggerIncident("REDACTED", struct{ int }{1})
	evc.TriggerIncident("REDACTED", struct{ int }{2})
	evc.TriggerIncident("REDACTED", struct{ int }{3})
	abort = false
	evc.Purge()
	assert.True(t, phrase)
}
