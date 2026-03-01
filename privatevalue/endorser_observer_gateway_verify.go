package privatevalue

import (
	"errors"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var (
	verifyDeadlineEmbrace = fallbackDeadlineEmbraceMoments * time.Second

	verifyDeadlineRetrievePersist    = 100 * time.Millisecond
	verifyDeadlineRetrievePersist2o3 = 60 * time.Millisecond //
)

type callerVerifyScenario struct {
	location   string
	caller PortCaller
}

//
//
//
//
//
func VerifyEndorserDistantReissueTcpsocketSolely(t *testing.T) {
	var (
		effortChnl = make(chan int)
		attempts   = 10
	)

	ln, err := net.Listen("REDACTED", "REDACTED")
	require.NoError(t, err)

	//
	go func(ln net.Listener, effortChnl chan<- int) {
		endeavors := 0
		for {
			link, err := ln.Accept()
			require.NoError(t, err)

			err = link.Close()
			require.NoError(t, err)

			endeavors++

			if endeavors == attempts {
				effortChnl <- endeavors
				break
			}
		}
	}(ln, effortChnl)

	callerGateway := FreshEndorserCallerGateway(
		log.VerifyingTracer(),
		CallStreamProc(ln.Addr().String(), verifyDeadlineRetrievePersist, edwards25519.ProducePrivateToken()),
	)
	EndorserCallerGatewayDeadlineRetrievePersist(time.Millisecond)(callerGateway)
	EndorserCallerGatewayLinkAttempts(attempts)(callerGateway)

	successionUUID := commitrand.Str(12)
	simulatePRV := kinds.FreshSimulatePRV()
	endorserDaemon := FreshEndorserDaemon(callerGateway, successionUUID, simulatePRV)

	err = endorserDaemon.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := endorserDaemon.Halt(); err != nil {
			t.Error(err)
		}
	})

	select {
	case endeavors := <-effortChnl:
		assert.Equal(t, attempts, endeavors)
	case <-time.After(1500 * time.Millisecond):
		t.Error("REDACTED")
	}
}

func VerifyReissueLinkTowardDistantEndorser(t *testing.T) {
	for _, tc := range obtainCallerVerifyScenarios(t) {
		var (
			tracer           = log.VerifyingTracer()
			successionUUID          = commitrand.Str(12)
			simulatePRV           = kinds.FreshSimulatePRV()
			gatewayEqualsUnlockChnl = make(chan struct{})
			thatLinkDeadline  = verifyDeadlineRetrievePersist
			observerGateway = freshEndorserObserverGateway(tracer, tc.location, thatLinkDeadline)
		)

		callerGateway := FreshEndorserCallerGateway(
			tracer,
			tc.caller,
		)
		EndorserCallerGatewayDeadlineRetrievePersist(verifyDeadlineRetrievePersist)(callerGateway)
		EndorserCallerGatewayLinkAttempts(10)(callerGateway)

		endorserDaemon := FreshEndorserDaemon(callerGateway, successionUUID, simulatePRV)

		initiateObserverGatewayAsyncronous(t, observerGateway, gatewayEqualsUnlockChnl)
		t.Cleanup(func() {
			if err := observerGateway.Halt(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, endorserDaemon.Initiate())
		assert.True(t, endorserDaemon.EqualsActive())
		<-gatewayEqualsUnlockChnl
		if err := endorserDaemon.Halt(); err != nil {
			t.Error(err)
		}

		callerInterface2 := FreshEndorserCallerGateway(
			tracer,
			tc.caller,
		)
		endorserFacility2 := FreshEndorserDaemon(callerInterface2, successionUUID, simulatePRV)

		//
		require.NoError(t, endorserFacility2.Initiate())
		assert.True(t, endorserFacility2.EqualsActive())
		t.Cleanup(func() {
			if err := endorserFacility2.Halt(); err != nil {
				t.Error(err)
			}
		})

		//
		//
		//
		//
		//
		time.Sleep(verifyDeadlineRetrievePersist * 2)
	}
}

func VerifyReplicatedOverhearDecline(t *testing.T) {
	for _, tc := range obtainCallerVerifyScenarios(t) {
		var (
			tracer           = log.VerifyingTracer()
			successionUUID          = commitrand.Str(12)
			simulatePRV           = kinds.FreshSimulatePRV()
			gatewayEqualsUnlockChnl = make(chan struct{})
			thatLinkDeadline  = verifyDeadlineRetrievePersist
			observerGateway = freshEndorserObserverGateway(tracer, tc.location, thatLinkDeadline)
		)
		observerGateway.deadlineEmbrace = fallbackDeadlineEmbraceMoments / 2 * time.Second

		callerGateway := FreshEndorserCallerGateway(
			tracer,
			tc.caller,
		)
		EndorserCallerGatewayDeadlineRetrievePersist(verifyDeadlineRetrievePersist)(callerGateway)
		EndorserCallerGatewayLinkAttempts(10)(callerGateway)

		endorserDaemon := FreshEndorserDaemon(callerGateway, successionUUID, simulatePRV)

		initiateObserverGatewayAsyncronous(t, observerGateway, gatewayEqualsUnlockChnl)
		t.Cleanup(func() {
			if err := observerGateway.Halt(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, endorserDaemon.Initiate())
		assert.True(t, endorserDaemon.EqualsActive())

		<-gatewayEqualsUnlockChnl
		if err := endorserDaemon.Halt(); err != nil {
			t.Error(err)
		}

		callerInterface2 := FreshEndorserCallerGateway(
			tracer,
			tc.caller,
		)
		endorserFacility2 := FreshEndorserDaemon(callerInterface2, successionUUID, simulatePRV)

		//
		require.NoError(t, endorserFacility2.Initiate())
		assert.True(t, endorserFacility2.EqualsActive())

		//
		for !observerGateway.EqualsAssociated() {
		}

		//
		time.Sleep(100 * time.Millisecond)
		observerGateway.activateRelate()
		time.Sleep(100 * time.Millisecond)
		observerGateway.activateRelate()
		time.Sleep(100 * time.Millisecond)
		observerGateway.activateRelate()

		//
		//
		//
		time.Sleep(3 * fallbackDeadlineEmbraceMoments * time.Second)
		t.Cleanup(func() {
			if err := endorserFacility2.Halt(); err != nil {
				t.Error(err)
			}
		})

		//
		assert.LessOrEqual(t, observerGateway.embraceMishapTally.Load(), uint32(1))

		//
		//
		//
		//
		//
		time.Sleep(verifyDeadlineRetrievePersist * 2)
	}
}

func freshEndorserObserverGateway(tracer log.Tracer, location string, deadlineRetrievePersist time.Duration) *EndorserObserverGateway {
	schema, location := strongmindnet.SchemeAlsoLocation(location)

	ln, err := net.Listen(schema, location)
	tracer.Details("REDACTED", "REDACTED", schema, "REDACTED", location)
	if err != nil {
		panic(err)
	}

	var observer net.Listener

	if schema == "REDACTED" {
		posixLink := FreshPosixObserver(ln)
		PosixObserverDeadlineEmbrace(verifyDeadlineEmbrace)(posixLink)
		PosixObserverDeadlineFetchPersist(deadlineRetrievePersist)(posixLink)
		observer = posixLink
	} else {
		tcpsocketLink := FreshTcpsocketObserver(ln, edwards25519.ProducePrivateToken())
		TcpsocketObserverDeadlineEmbrace(verifyDeadlineEmbrace)(tcpsocketLink)
		TcpsocketObserverDeadlineFetchPersist(deadlineRetrievePersist)(tcpsocketLink)
		observer = tcpsocketLink
	}

	return FreshEndorserObserverGateway(
		tracer,
		observer,
		EndorserObserverGatewayDeadlineRetrievePersist(verifyDeadlineRetrievePersist),
	)
}

func initiateObserverGatewayAsyncronous(t *testing.T, sle *EndorserObserverGateway, gatewayEqualsUnlockChnl chan struct{}) {
	go func(sle *EndorserObserverGateway) {
		require.NoError(t, sle.Initiate())
		assert.True(t, sle.EqualsActive())
		close(gatewayEqualsUnlockChnl)
	}(sle)
}

func obtainSimulateTerminals(
	t *testing.T,
	location string,
	portCaller PortCaller,
) (*EndorserObserverGateway, *EndorserCallerGateway) {
	var (
		tracer           = log.VerifyingTracer()
		gatewayEqualsUnlockChnl = make(chan struct{})

		callerGateway = FreshEndorserCallerGateway(
			tracer,
			portCaller,
		)

		observerGateway = freshEndorserObserverGateway(tracer, location, verifyDeadlineRetrievePersist)
	)

	EndorserCallerGatewayDeadlineRetrievePersist(verifyDeadlineRetrievePersist)(callerGateway)
	EndorserCallerGatewayLinkAttempts(1e6)(callerGateway)

	initiateObserverGatewayAsyncronous(t, observerGateway, gatewayEqualsUnlockChnl)

	require.NoError(t, callerGateway.Initiate())
	assert.True(t, callerGateway.EqualsActive())

	<-gatewayEqualsUnlockChnl

	return observerGateway, callerGateway
}

func VerifyEndorserObserverGatewayFacilityCycle(t *testing.T) {
	observerGateway := FreshEndorserObserverGateway(
		log.VerifyingTracer(),
		&verifyObserver{primaryErrors: 5},
	)

	require.NoError(t, observerGateway.Initiate())
	require.NoError(t, observerGateway.PauseForeachLinkage(time.Second))
}

type verifyObserver struct {
	net.Observer
	primaryErrors int
}

func (l *verifyObserver) Embrace() (net.Conn, error) {
	if l.primaryErrors > 0 {
		l.primaryErrors--

		return nil, errors.New("REDACTED")
	}

	return nil, nil //
}
