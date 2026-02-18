package pubsub_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"

	"github.com/valkyrieworks/utils/broadcast"
	"github.com/valkyrieworks/utils/broadcast/inquire"
)

func VerifyInstance(t *testing.T) {
	s := broadcast.NewHost()
	s.AssignTracer(log.VerifyingTracer())
	err := s.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := s.Halt(); err != nil {
			t.Error(err)
		}
	})

	ctx := context.Background()
	enrollment, err := s.Enrol(ctx, "REDACTED", inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)
	err = s.BroadcastWithEvents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", enrollment.Out())
}
