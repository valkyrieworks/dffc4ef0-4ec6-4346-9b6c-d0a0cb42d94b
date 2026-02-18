package privatekey

import (
	"errors"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	cometnet "github.com/valkyrieworks/utils/net"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/kinds"
)

var (
	verifyDeadlineAllow = standardDeadlineAllowMoments * time.Second

	verifyDeadlineFetchRecord    = 100 * time.Millisecond
	verifyDeadlineFetchRecord2o3 = 60 * time.Millisecond //
)

type callerVerifyScenario struct {
	address   string
	caller SocketCaller
}

//
//
//
//
//
func VerifyNotaryDistantReprocessTCPSolely(t *testing.T) {
	var (
		endeavorChan = make(chan int)
		attempts   = 10
	)

	ln, err := net.Listen("REDACTED", "REDACTED")
	require.NoError(t, err)

	//
	go func(ln net.Listener, endeavorChan chan<- int) {
		tries := 0
		for {
			link, err := ln.Accept()
			require.NoError(t, err)

			err = link.Close()
			require.NoError(t, err)

			tries++

			if tries == attempts {
				endeavorChan <- tries
				break
			}
		}
	}(ln, endeavorChan)

	callerTerminus := NewNotaryCallerGateway(
		log.VerifyingTracer(),
		CallTCPFn(ln.Addr().String(), verifyDeadlineFetchRecord, ed25519.GeneratePrivateKey()),
	)
	NotaryCallerTerminusDeadlineFetchRecord(time.Millisecond)(callerTerminus)
	NotaryCallerTerminusLinkAttempts(attempts)(callerTerminus)

	ledgerUID := engineseed.Str(12)
	emulatePV := kinds.NewEmulatePV()
	notaryHost := NewNotaryHost(callerTerminus, ledgerUID, emulatePV)

	err = notaryHost.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := notaryHost.Halt(); err != nil {
			t.Error(err)
		}
	})

	select {
	case tries := <-endeavorChan:
		assert.Equal(t, attempts, tries)
	case <-time.After(1500 * time.Millisecond):
		t.Error("REDACTED")
	}
}

func VerifyReprocessLinkToDistantNotary(t *testing.T) {
	for _, tc := range fetchCallerVerifyScenarios(t) {
		var (
			tracer           = log.VerifyingTracer()
			ledgerUID          = engineseed.Str(12)
			emulatePV           = kinds.NewEmulatePV()
			terminusIsAccessChan = make(chan struct{})
			thisLinkDeadline  = verifyDeadlineFetchRecord
			observerTerminus = newNotaryObserverTerminus(tracer, tc.address, thisLinkDeadline)
		)

		callerTerminus := NewNotaryCallerGateway(
			tracer,
			tc.caller,
		)
		NotaryCallerTerminusDeadlineFetchRecord(verifyDeadlineFetchRecord)(callerTerminus)
		NotaryCallerTerminusLinkAttempts(10)(callerTerminus)

		notaryHost := NewNotaryHost(callerTerminus, ledgerUID, emulatePV)

		beginObserverTerminusAsync(t, observerTerminus, terminusIsAccessChan)
		t.Cleanup(func() {
			if err := observerTerminus.Halt(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, notaryHost.Begin())
		assert.True(t, notaryHost.IsActive())
		<-terminusIsAccessChan
		if err := notaryHost.Halt(); err != nil {
			t.Error(err)
		}

		callerTerminus2 := NewNotaryCallerGateway(
			tracer,
			tc.caller,
		)
		notaryServer2 := NewNotaryHost(callerTerminus2, ledgerUID, emulatePV)

		//
		require.NoError(t, notaryServer2.Begin())
		assert.True(t, notaryServer2.IsActive())
		t.Cleanup(func() {
			if err := notaryServer2.Halt(); err != nil {
				t.Error(err)
			}
		})

		//
		//
		//
		//
		//
		time.Sleep(verifyDeadlineFetchRecord * 2)
	}
}

func VerifyReplicatedObserveDecline(t *testing.T) {
	for _, tc := range fetchCallerVerifyScenarios(t) {
		var (
			tracer           = log.VerifyingTracer()
			ledgerUID          = engineseed.Str(12)
			emulatePV           = kinds.NewEmulatePV()
			terminusIsAccessChan = make(chan struct{})
			thisLinkDeadline  = verifyDeadlineFetchRecord
			observerTerminus = newNotaryObserverTerminus(tracer, tc.address, thisLinkDeadline)
		)
		observerTerminus.deadlineAllow = standardDeadlineAllowMoments / 2 * time.Second

		callerTerminus := NewNotaryCallerGateway(
			tracer,
			tc.caller,
		)
		NotaryCallerTerminusDeadlineFetchRecord(verifyDeadlineFetchRecord)(callerTerminus)
		NotaryCallerTerminusLinkAttempts(10)(callerTerminus)

		notaryHost := NewNotaryHost(callerTerminus, ledgerUID, emulatePV)

		beginObserverTerminusAsync(t, observerTerminus, terminusIsAccessChan)
		t.Cleanup(func() {
			if err := observerTerminus.Halt(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, notaryHost.Begin())
		assert.True(t, notaryHost.IsActive())

		<-terminusIsAccessChan
		if err := notaryHost.Halt(); err != nil {
			t.Error(err)
		}

		callerTerminus2 := NewNotaryCallerGateway(
			tracer,
			tc.caller,
		)
		notaryServer2 := NewNotaryHost(callerTerminus2, ledgerUID, emulatePV)

		//
		require.NoError(t, notaryServer2.Begin())
		assert.True(t, notaryServer2.IsActive())

		//
		for !observerTerminus.IsLinked() {
		}

		//
		time.Sleep(100 * time.Millisecond)
		observerTerminus.activateEstablish()
		time.Sleep(100 * time.Millisecond)
		observerTerminus.activateEstablish()
		time.Sleep(100 * time.Millisecond)
		observerTerminus.activateEstablish()

		//
		//
		//
		time.Sleep(3 * standardDeadlineAllowMoments * time.Second)
		t.Cleanup(func() {
			if err := notaryServer2.Halt(); err != nil {
				t.Error(err)
			}
		})

		//
		assert.LessOrEqual(t, observerTerminus.allowAbortNumber.Load(), uint32(1))

		//
		//
		//
		//
		//
		time.Sleep(verifyDeadlineFetchRecord * 2)
	}
}

func newNotaryObserverTerminus(tracer log.Tracer, address string, deadlineFetchRecord time.Duration) *NotaryObserverTerminus {
	schema, location := cometnet.ProtocolAndLocation(address)

	ln, err := net.Listen(schema, location)
	tracer.Details("REDACTED", "REDACTED", schema, "REDACTED", location)
	if err != nil {
		panic(err)
	}

	var observer net.Listener

	if schema == "REDACTED" {
		unixLn := NewUnixObserver(ln)
		UnixObserverDeadlineAllow(verifyDeadlineAllow)(unixLn)
		UnixObserverDeadlineFetchRecord(deadlineFetchRecord)(unixLn)
		observer = unixLn
	} else {
		tcpLn := NewTCPObserver(ln, ed25519.GeneratePrivateKey())
		TCPObserverDeadlineAllow(verifyDeadlineAllow)(tcpLn)
		TCPObserverDeadlineFetchRecord(deadlineFetchRecord)(tcpLn)
		observer = tcpLn
	}

	return NewNotaryObserverTerminus(
		tracer,
		observer,
		NotaryObserverTerminusDeadlineFetchRecord(verifyDeadlineFetchRecord),
	)
}

func beginObserverTerminusAsync(t *testing.T, sle *NotaryObserverTerminus, terminusIsAccessChan chan struct{}) {
	go func(sle *NotaryObserverTerminus) {
		require.NoError(t, sle.Begin())
		assert.True(t, sle.IsActive())
		close(terminusIsAccessChan)
	}(sle)
}

func fetchEmulateTermini(
	t *testing.T,
	address string,
	socketCaller SocketCaller,
) (*NotaryObserverTerminus, *NotaryCallerTerminus) {
	var (
		tracer           = log.VerifyingTracer()
		terminusIsAccessChan = make(chan struct{})

		callerTerminus = NewNotaryCallerGateway(
			tracer,
			socketCaller,
		)

		observerTerminus = newNotaryObserverTerminus(tracer, address, verifyDeadlineFetchRecord)
	)

	NotaryCallerTerminusDeadlineFetchRecord(verifyDeadlineFetchRecord)(callerTerminus)
	NotaryCallerTerminusLinkAttempts(1e6)(callerTerminus)

	beginObserverTerminusAsync(t, observerTerminus, terminusIsAccessChan)

	require.NoError(t, callerTerminus.Begin())
	assert.True(t, callerTerminus.IsActive())

	<-terminusIsAccessChan

	return observerTerminus, callerTerminus
}

func VerifyNotaryObserverTerminusDaemonCycle(t *testing.T) {
	observerTerminus := NewNotaryObserverTerminus(
		log.VerifyingTracer(),
		&verifyObserver{primaryErrors: 5},
	)

	require.NoError(t, observerTerminus.Begin())
	require.NoError(t, observerTerminus.WaitForLinkage(time.Second))
}

type verifyObserver struct {
	net.Observer
	primaryErrors int
}

func (l *verifyObserver) Allow() (net.Conn, error) {
	if l.primaryErrors > 0 {
		l.primaryErrors--

		return nil, errors.New("REDACTED")
	}

	return nil, nil //
}
