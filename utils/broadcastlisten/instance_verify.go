package broadlisten_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
)

func VerifyInstance(t *testing.T) {
	s := broadcastlisten.FreshDaemon()
	s.AssignTracer(log.VerifyingTracer())
	err := s.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := s.Halt(); err != nil {
			t.Error(err)
		}
	})

	ctx := context.Background()
	listening, err := s.Listen(ctx, "REDACTED", inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)
	err = s.BroadcastUsingIncidents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening.Out())
}
