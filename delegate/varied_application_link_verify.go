package delegate

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	abcistubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate/simulations"
)

func Verifyapplicationlinks_Initiate_Halt(t *testing.T) {
	exitChnl := make(<-chan struct{})

	customerOriginatorSimulate := &simulations.CustomerOriginator{}

	customerSimulate := &abcistubs.Customer{}
	customerSimulate.On("REDACTED", mock.Anything).Return().Times(4)
	customerSimulate.On("REDACTED").Return(nil).Times(4)
	customerSimulate.On("REDACTED").Return(nil).Times(4)
	customerSimulate.On("REDACTED").Return(exitChnl).Times(4)

	customerOriginatorSimulate.On("REDACTED").Return(customerSimulate, nil).Times(4)

	applicationLinks := FreshPlatformLinks(customerOriginatorSimulate, NooperationTelemetry())

	err := applicationLinks.Initiate()
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	err = applicationLinks.Halt()
	require.NoError(t, err)

	customerSimulate.AssertExpectations(t)
}

//
func Verifyapplicationlinks_Breakdown(t *testing.T) {
	ok := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for range c {
			close(ok)
		}
	}()

	exitChnl := make(chan struct{})
	var obtainExitChnl <-chan struct{} = exitChnl

	customerOriginatorSimulate := &simulations.CustomerOriginator{}

	customerSimulate := &abcistubs.Customer{}
	customerSimulate.On("REDACTED", mock.Anything).Return()
	customerSimulate.On("REDACTED").Return(nil)
	customerSimulate.On("REDACTED").Return(nil)

	customerSimulate.On("REDACTED").Return(obtainExitChnl)
	customerSimulate.On("REDACTED").Return(errors.New("REDACTED")).Once()

	customerOriginatorSimulate.On("REDACTED").Return(customerSimulate, nil)

	applicationLinks := FreshPlatformLinks(customerOriginatorSimulate, NooperationTelemetry())

	err := applicationLinks.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := applicationLinks.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	close(exitChnl)

	select {
	case <-ok:
		t.Log("REDACTED")
	case <-time.After(5 * time.Second):
		t.Fatal("REDACTED")
	}
}
