package gateway

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	abciemulators "github.com/valkyrieworks/iface/customer/simulations"
	"github.com/valkyrieworks/gateway/simulations"
)

func Testapplicationlinks_Begin_Halt(t *testing.T) {
	exitChan := make(<-chan struct{})

	customerOriginatorEmulate := &simulations.CustomerOriginator{}

	customerEmulate := &abciemulators.Customer{}
	customerEmulate.On("REDACTED", mock.Anything).Return().Times(4)
	customerEmulate.On("REDACTED").Return(nil).Times(4)
	customerEmulate.On("REDACTED").Return(nil).Times(4)
	customerEmulate.On("REDACTED").Return(exitChan).Times(4)

	customerOriginatorEmulate.On("REDACTED").Return(customerEmulate, nil).Times(4)

	applicationLinks := NewApplicationLinks(customerOriginatorEmulate, NoopStats())

	err := applicationLinks.Begin()
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	err = applicationLinks.Halt()
	require.NoError(t, err)

	customerEmulate.AssertExpectations(t)
}

//
func Testapplicationlinks_Breakdown(t *testing.T) {
	ok := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for range c {
			close(ok)
		}
	}()

	exitChan := make(chan struct{})
	var receiveExitChan <-chan struct{} = exitChan

	customerOriginatorEmulate := &simulations.CustomerOriginator{}

	customerEmulate := &abciemulators.Customer{}
	customerEmulate.On("REDACTED", mock.Anything).Return()
	customerEmulate.On("REDACTED").Return(nil)
	customerEmulate.On("REDACTED").Return(nil)

	customerEmulate.On("REDACTED").Return(receiveExitChan)
	customerEmulate.On("REDACTED").Return(errors.New("REDACTED")).Once()

	customerOriginatorEmulate.On("REDACTED").Return(customerEmulate, nil)

	applicationLinks := NewApplicationLinks(customerOriginatorEmulate, NoopStats())

	err := applicationLinks.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := applicationLinks.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	close(exitChan)

	select {
	case <-ok:
		t.Log("REDACTED")
	case <-time.After(5 * time.Second):
		t.Fatal("REDACTED")
	}
}
