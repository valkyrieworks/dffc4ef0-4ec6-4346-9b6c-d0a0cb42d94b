package pubsub_test

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"

	"github.com/valkyrieworks/utils/broadcast"
	"github.com/valkyrieworks/utils/broadcast/inquire"
)

const (
	customerUID = "REDACTED"
)

func VerifyEnrol(t *testing.T) {
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
	enrollment, err := s.Enrol(ctx, customerUID, inquire.All)
	require.NoError(t, err)

	assert.Equal(t, 1, s.CountAgents())
	assert.Equal(t, 1, s.CountCustomerRegistrations(customerUID))

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", enrollment.Out())

	issued := make(chan struct{})
	go func() {
		defer close(issued)

		err := s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)

		err = s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)

		err = s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)
	}()

	select {
	case <-issued:
		affirmAccept(t, "REDACTED", enrollment.Out())
		affirmAborted(t, enrollment, broadcast.ErrOutOfAbility)
	case <-time.After(3 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyEnrolWithAbility(t *testing.T) {
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
	assert.Panics(t, func() {
		_, err = s.Enrol(ctx, customerUID, inquire.All, -1)
		require.NoError(t, err)
	})
	assert.Panics(t, func() {
		_, err = s.Enrol(ctx, customerUID, inquire.All, 0)
		require.NoError(t, err)
	})
	enrollment, err := s.Enrol(ctx, customerUID, inquire.All, 1)
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", enrollment.Out())
}

func VerifyEnrolUnbuffered(t *testing.T) {
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
	enrollment, err := s.EnrolUnbuffered(ctx, customerUID, inquire.All)
	require.NoError(t, err)

	issued := make(chan struct{})
	go func() {
		defer close(issued)

		err := s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)

		err = s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)
	}()

	select {
	case <-issued:
		t.Fatal("REDACTED")
	case <-time.After(3 * time.Second):
		affirmAccept(t, "REDACTED", enrollment.Out())
		affirmAccept(t, "REDACTED", enrollment.Out())
	}
}

func VerifyGradualCustomerIsDeletedWithErrOutOfAbility(t *testing.T) {
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
	enrollment, err := s.Enrol(ctx, customerUID, inquire.All)
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)

	affirmAborted(t, enrollment, broadcast.ErrOutOfAbility)
}

func VerifyDistinctAgents(t *testing.T) {
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
	feed1, err := s.Enrol(ctx, "REDACTED", inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)
	err = s.BroadcastWithEvents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", feed1.Out())

	feed2, err := s.Enrol(
		ctx,
		"REDACTED",
		inquire.ShouldBuild("REDACTED"),
	)
	require.NoError(t, err)
	err = s.BroadcastWithEvents(
		ctx,
		"REDACTED",
		map[string][]string{"REDACTED": {"REDACTED"}, "REDACTED": {"REDACTED"}},
	)
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", feed1.Out())
	affirmAccept(t, "REDACTED", feed2.Out())

	feed3, err := s.Enrol(
		ctx,
		"REDACTED",
		inquire.ShouldBuild("REDACTED"),
	)
	require.NoError(t, err)
	err = s.BroadcastWithEvents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	assert.Zero(t, len(feed3.Out()))
}

func VerifyEnrolReplicatedKeys(t *testing.T) {
	ctx := context.Background()
	s := broadcast.NewHost()
	s.AssignTracer(log.VerifyingTracer())
	require.NoError(t, s.Begin())
	t.Cleanup(func() {
		if err := s.Halt(); err != nil {
			t.Error(err)
		}
	})

	verifyScenarios := []struct {
		inquire    string
		anticipated any
	}{
		{
			"REDACTED",
			"REDACTED",
		},
		{
			"REDACTED",
			"REDACTED",
		},
		{
			"REDACTED",
			"REDACTED",
		},
		{
			"REDACTED",
			nil,
		},
	}

	for i, tc := range verifyScenarios {
		sub, err := s.Enrol(ctx, fmt.Sprintf("REDACTED", i), inquire.ShouldBuild(tc.inquire))
		require.NoError(t, err)

		err = s.BroadcastWithEvents(
			ctx,
			"REDACTED",
			map[string][]string{
				"REDACTED":  {"REDACTED", "REDACTED", "REDACTED"},
				"REDACTED": {"REDACTED", "REDACTED", "REDACTED"},
			},
		)
		require.NoError(t, err)

		if tc.anticipated != nil {
			affirmAccept(t, tc.anticipated, sub.Out())
		} else {
			require.Zero(t, len(sub.Out()))
		}
	}
}

func VerifyCustomerRegistersTwice(t *testing.T) {
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
	q := inquire.ShouldBuild("REDACTED")

	feed1, err := s.Enrol(ctx, customerUID, q)
	require.NoError(t, err)
	err = s.BroadcastWithEvents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", feed1.Out())

	feed2, err := s.Enrol(ctx, customerUID, q)
	require.Error(t, err)
	require.Nil(t, feed2)

	err = s.BroadcastWithEvents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", feed1.Out())
}

func VerifyDeenroll(t *testing.T) {
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
	enrollment, err := s.Enrol(ctx, customerUID, inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)
	err = s.Deenroll(ctx, customerUID, inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	assert.Zero(t, len(enrollment.Out()), "REDACTED")

	affirmAborted(t, enrollment, broadcast.ErrDeactivated)
}

func VerifyCustomerUnregistersTwice(t *testing.T) {
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
	_, err = s.Enrol(ctx, customerUID, inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)
	err = s.Deenroll(ctx, customerUID, inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)

	err = s.Deenroll(ctx, customerUID, inquire.ShouldBuild("REDACTED"))
	assert.Equal(t, broadcast.ErrEnrollmentNegateLocated, err)
	err = s.DeenrollAll(ctx, customerUID)
	assert.Equal(t, broadcast.ErrEnrollmentNegateLocated, err)
}

func VerifyReactivate(t *testing.T) {
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
	_, err = s.Enrol(ctx, customerUID, inquire.All)
	require.NoError(t, err)
	err = s.Deenroll(ctx, customerUID, inquire.All)
	require.NoError(t, err)
	enrollment, err := s.Enrol(ctx, customerUID, inquire.All)
	require.NoError(t, err)

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", enrollment.Out())
}

func VerifyDeenrollAll(t *testing.T) {
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
	feed1, err := s.Enrol(ctx, customerUID, inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)
	feed2, err := s.Enrol(ctx, customerUID, inquire.ShouldBuild("REDACTED"))
	require.NoError(t, err)

	err = s.DeenrollAll(ctx, customerUID)
	require.NoError(t, err)

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	assert.Zero(t, len(feed1.Out()), "REDACTED")
	assert.Zero(t, len(feed2.Out()), "REDACTED")

	affirmAborted(t, feed1, broadcast.ErrDeactivated)
	affirmAborted(t, feed2, broadcast.ErrDeactivated)
}

func VerifyFrameAbility(t *testing.T) {
	s := broadcast.NewHost(broadcast.BufferVolume(2))
	s.AssignTracer(log.VerifyingTracer())

	assert.Equal(t, 2, s.BufferVolume())

	ctx := context.Background()
	err := s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)

	ctx, revoke := context.WithTimeout(ctx, 10*time.Millisecond)
	defer revoke()
	err = s.Broadcast(ctx, "REDACTED")
	if assert.Error(t, err) {
		assert.Equal(t, context.DeadlineExceeded, err)
	}
}

func Benchmark10customers(b *testing.B)   { criterionNAgents(10, b) }
func Benchmark100customers(b *testing.B)  { criterionNAgents(100, b) }
func Benchmark1000customers(b *testing.B) { criterionNAgents(1000, b) }

func Benchmark10customersOneInquire(b *testing.B)   { criterionNAgentsOneInquire(10, b) }
func Benchmark100customersOneInquire(b *testing.B)  { criterionNAgentsOneInquire(100, b) }
func Benchmark1000customersOneInquire(b *testing.B) { criterionNAgentsOneInquire(1000, b) }

func criterionNAgents(n int, b *testing.B) {
	s := broadcast.NewHost()
	err := s.Begin()
	require.NoError(b, err)

	b.Cleanup(func() {
		if err := s.Halt(); err != nil {
			b.Error(err)
		}
	})

	ctx := context.Background()
	for i := 0; i < n; i++ {
		enrollment, err := s.Enrol(
			ctx,
			customerUID,
			inquire.ShouldBuild(fmt.Sprintf("REDACTED", i)),
		)
		if err != nil {
			b.Fatal(err)
		}
		go func() {
			for {
				select {
				case <-enrollment.Out():
					continue
				case <-enrollment.Revoked():
					return
				}
			}
		}()
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = s.BroadcastWithEvents(
			ctx,
			"REDACTED",
			map[string][]string{"REDACTED": {"REDACTED"}, "REDACTED": {string(rune(i))}},
		)
		require.NoError(b, err)
	}
}

func criterionNAgentsOneInquire(n int, b *testing.B) {
	s := broadcast.NewHost()
	err := s.Begin()
	require.NoError(b, err)
	b.Cleanup(func() {
		if err := s.Halt(); err != nil {
			b.Error(err)
		}
	})

	ctx := context.Background()
	q := inquire.ShouldBuild("REDACTED")
	for i := 0; i < n; i++ {
		enrollment, err := s.Enrol(ctx, customerUID, q)
		if err != nil {
			b.Fatal(err)
		}
		go func() {
			for {
				select {
				case <-enrollment.Out():
					continue
				case <-enrollment.Revoked():
					return
				}
			}
		}()
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = s.BroadcastWithEvents(ctx, "REDACTED", map[string][]string{
			"REDACTED":   {"REDACTED"},
			"REDACTED": {"REDACTED"},
		})
		require.NoError(b, err)
	}
}

//

func affirmAccept(t *testing.T, anticipated any, ch <-chan broadcast.Signal, messageAndArgs ...any) {
	select {
	case factual := <-ch:
		assert.Equal(t, anticipated, factual.Data(), messageAndArgs...)
	case <-time.After(1 * time.Second):
		t.Errorf("REDACTED", anticipated)
		debug.PrintStack()
	}
}

func affirmAborted(t *testing.T, enrollment *broadcast.Enrollment, err error) {
	_, ok := <-enrollment.Revoked()
	assert.False(t, ok)
	assert.Equal(t, err, enrollment.Err())
}
