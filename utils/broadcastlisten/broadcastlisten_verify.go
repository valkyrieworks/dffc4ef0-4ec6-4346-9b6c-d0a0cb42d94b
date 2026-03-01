package broadlisten_test

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
)

const (
	customerUUID = "REDACTED"
)

func VerifyListen(t *testing.T) {
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
	listening, err := s.Listen(ctx, customerUUID, inquire.All)
	require.NoError(t, err)

	assert.Equal(t, 1, s.CountCustomers())
	assert.Equal(t, 1, s.CountCustomerFeeds(customerUUID))

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening.Out())

	broadcasted := make(chan struct{})
	go func() {
		defer close(broadcasted)

		err := s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)

		err = s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)

		err = s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)
	}()

	select {
	case <-broadcasted:
		affirmAccept(t, "REDACTED", listening.Out())
		affirmAborted(t, listening, broadcastlisten.FaultOutputBelongingVolume)
	case <-time.After(3 * time.Second):
		t.Fatal("REDACTED")
	}
}

func VerifyListenUsingVolume(t *testing.T) {
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
	assert.Panics(t, func() {
		_, err = s.Listen(ctx, customerUUID, inquire.All, -1)
		require.NoError(t, err)
	})
	assert.Panics(t, func() {
		_, err = s.Listen(ctx, customerUUID, inquire.All, 0)
		require.NoError(t, err)
	})
	listening, err := s.Listen(ctx, customerUUID, inquire.All, 1)
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening.Out())
}

func VerifyListenUncached(t *testing.T) {
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
	listening, err := s.ListenUncached(ctx, customerUUID, inquire.All)
	require.NoError(t, err)

	broadcasted := make(chan struct{})
	go func() {
		defer close(broadcasted)

		err := s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)

		err = s.Broadcast(ctx, "REDACTED")
		assert.NoError(t, err)
	}()

	select {
	case <-broadcasted:
		t.Fatal("REDACTED")
	case <-time.After(3 * time.Second):
		affirmAccept(t, "REDACTED", listening.Out())
		affirmAccept(t, "REDACTED", listening.Out())
	}
}

func VerifyGradualCustomerEqualsDiscardedUsingFaultOutputBelongingVolume(t *testing.T) {
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
	listening, err := s.Listen(ctx, customerUUID, inquire.All)
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)

	affirmAborted(t, listening, broadcastlisten.FaultOutputBelongingVolume)
}

func VerifyDistinctCustomers(t *testing.T) {
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
	listening1, err := s.Listen(ctx, "REDACTED", inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)
	err = s.BroadcastUsingIncidents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening1.Out())

	listening2, err := s.Listen(
		ctx,
		"REDACTED",
		inquire.ShouldAssemble("REDACTED"),
	)
	require.NoError(t, err)
	err = s.BroadcastUsingIncidents(
		ctx,
		"REDACTED",
		map[string][]string{"REDACTED": {"REDACTED"}, "REDACTED": {"REDACTED"}},
	)
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening1.Out())
	affirmAccept(t, "REDACTED", listening2.Out())

	listening3, err := s.Listen(
		ctx,
		"REDACTED",
		inquire.ShouldAssemble("REDACTED"),
	)
	require.NoError(t, err)
	err = s.BroadcastUsingIncidents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	assert.Zero(t, len(listening3.Out()))
}

func VerifyListenReplicatedTokens(t *testing.T) {
	ctx := context.Background()
	s := broadcastlisten.FreshDaemon()
	s.AssignTracer(log.VerifyingTracer())
	require.NoError(t, s.Initiate())
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
		sub, err := s.Listen(ctx, fmt.Sprintf("REDACTED", i), inquire.ShouldAssemble(tc.inquire))
		require.NoError(t, err)

		err = s.BroadcastUsingIncidents(
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

func VerifyCustomerListensBis(t *testing.T) {
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
	q := inquire.ShouldAssemble("REDACTED")

	listening1, err := s.Listen(ctx, customerUUID, q)
	require.NoError(t, err)
	err = s.BroadcastUsingIncidents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening1.Out())

	listening2, err := s.Listen(ctx, customerUUID, q)
	require.Error(t, err)
	require.Nil(t, listening2)

	err = s.BroadcastUsingIncidents(ctx, "REDACTED", map[string][]string{"REDACTED": {"REDACTED"}})
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening1.Out())
}

func VerifyUnlisten(t *testing.T) {
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
	listening, err := s.Listen(ctx, customerUUID, inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)
	err = s.Unlisten(ctx, customerUUID, inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	assert.Zero(t, len(listening.Out()), "REDACTED")

	affirmAborted(t, listening, broadcastlisten.FaultUnlistened)
}

func VerifyCustomerUnlistensBis(t *testing.T) {
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
	_, err = s.Listen(ctx, customerUUID, inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)
	err = s.Unlisten(ctx, customerUUID, inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)

	err = s.Unlisten(ctx, customerUUID, inquire.ShouldAssemble("REDACTED"))
	assert.Equal(t, broadcastlisten.FaultListeningNegationDetected, err)
	err = s.UnlistenEvery(ctx, customerUUID)
	assert.Equal(t, broadcastlisten.FaultListeningNegationDetected, err)
}

func VerifyRelisten(t *testing.T) {
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
	_, err = s.Listen(ctx, customerUUID, inquire.All)
	require.NoError(t, err)
	err = s.Unlisten(ctx, customerUUID, inquire.All)
	require.NoError(t, err)
	listening, err := s.Listen(ctx, customerUUID, inquire.All)
	require.NoError(t, err)

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	affirmAccept(t, "REDACTED", listening.Out())
}

func VerifyUnlistenEvery(t *testing.T) {
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
	listening1, err := s.Listen(ctx, customerUUID, inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)
	listening2, err := s.Listen(ctx, customerUUID, inquire.ShouldAssemble("REDACTED"))
	require.NoError(t, err)

	err = s.UnlistenEvery(ctx, customerUUID)
	require.NoError(t, err)

	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	assert.Zero(t, len(listening1.Out()), "REDACTED")
	assert.Zero(t, len(listening2.Out()), "REDACTED")

	affirmAborted(t, listening1, broadcastlisten.FaultUnlistened)
	affirmAborted(t, listening2, broadcastlisten.FaultUnlistened)
}

func VerifyReserveVolume(t *testing.T) {
	s := broadcastlisten.FreshDaemon(broadcastlisten.ReserveVolume(2))
	s.AssignTracer(log.VerifyingTracer())

	assert.Equal(t, 2, s.ReserveVolume())

	ctx := context.Background()
	err := s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)
	err = s.Broadcast(ctx, "REDACTED")
	require.NoError(t, err)

	ctx, abort := context.WithTimeout(ctx, 10*time.Millisecond)
	defer abort()
	err = s.Broadcast(ctx, "REDACTED")
	if assert.Error(t, err) {
		assert.Equal(t, context.DeadlineExceeded, err)
	}
}

func Verify10nodes(b *testing.B)   { assessmentNTHCustomers(10, b) }
func Verify100nodes(b *testing.B)  { assessmentNTHCustomers(100, b) }
func Verify1000nodes(b *testing.B) { assessmentNTHCustomers(1000, b) }

func Verify10nodesSingleInquire(b *testing.B)   { assessmentNTHCustomersSingleInquire(10, b) }
func Verify100nodesSingleInquire(b *testing.B)  { assessmentNTHCustomersSingleInquire(100, b) }
func Verify1000nodesSingleInquire(b *testing.B) { assessmentNTHCustomersSingleInquire(1000, b) }

func assessmentNTHCustomers(n int, b *testing.B) {
	s := broadcastlisten.FreshDaemon()
	err := s.Initiate()
	require.NoError(b, err)

	b.Cleanup(func() {
		if err := s.Halt(); err != nil {
			b.Error(err)
		}
	})

	ctx := context.Background()
	for i := 0; i < n; i++ {
		listening, err := s.Listen(
			ctx,
			customerUUID,
			inquire.ShouldAssemble(fmt.Sprintf("REDACTED", i)),
		)
		if err != nil {
			b.Fatal(err)
		}
		go func() {
			for {
				select {
				case <-listening.Out():
					continue
				case <-listening.Aborted():
					return
				}
			}
		}()
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = s.BroadcastUsingIncidents(
			ctx,
			"REDACTED",
			map[string][]string{"REDACTED": {"REDACTED"}, "REDACTED": {string(rune(i))}},
		)
		require.NoError(b, err)
	}
}

func assessmentNTHCustomersSingleInquire(n int, b *testing.B) {
	s := broadcastlisten.FreshDaemon()
	err := s.Initiate()
	require.NoError(b, err)
	b.Cleanup(func() {
		if err := s.Halt(); err != nil {
			b.Error(err)
		}
	})

	ctx := context.Background()
	q := inquire.ShouldAssemble("REDACTED")
	for i := 0; i < n; i++ {
		listening, err := s.Listen(ctx, customerUUID, q)
		if err != nil {
			b.Fatal(err)
		}
		go func() {
			for {
				select {
				case <-listening.Out():
					continue
				case <-listening.Aborted():
					return
				}
			}
		}()
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = s.BroadcastUsingIncidents(ctx, "REDACTED", map[string][]string{
			"REDACTED":   {"REDACTED"},
			"REDACTED": {"REDACTED"},
		})
		require.NoError(b, err)
	}
}

//

func affirmAccept(t *testing.T, anticipated any, ch <-chan broadcastlisten.Signal, signalAlsoArguments ...any) {
	select {
	case existing := <-ch:
		assert.Equal(t, anticipated, existing.Data(), signalAlsoArguments...)
	case <-time.After(1 * time.Second):
		t.Errorf("REDACTED", anticipated)
		debug.PrintStack()
	}
}

func affirmAborted(t *testing.T, listening *broadcastlisten.Listening, err error) {
	_, ok := <-listening.Aborted()
	assert.False(t, ok)
	assert.Equal(t, err, listening.Err())
}
