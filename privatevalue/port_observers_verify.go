package privatevalue

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
)

//
//

func freshPrivateToken() edwards25519.PrivateToken {
	return edwards25519.ProducePrivateToken()
}

//
//

type observerVerifyInstance struct {
	definition string //
	observer    net.Listener
	caller      PortCaller
}

//
//
func verifyPosixLocation() (string, error) {
	f, err := os.CreateTemp("REDACTED", "REDACTED")
	if err != nil {
		return "REDACTED", err
	}
	location := f.Name()
	f.Close()
	os.Remove(location)
	return location, nil
}

func tcpsocketObserverVerifyInstance(t *testing.T, deadlineEmbrace, deadlineRetrievePersist time.Duration) observerVerifyInstance {
	ln, err := net.Listen("REDACTED", "REDACTED")
	if err != nil {
		t.Fatal(err)
	}

	tcpsocketLink := FreshTcpsocketObserver(ln, freshPrivateToken())
	TcpsocketObserverDeadlineEmbrace(deadlineEmbrace)(tcpsocketLink)
	TcpsocketObserverDeadlineFetchPersist(deadlineRetrievePersist)(tcpsocketLink)
	return observerVerifyInstance{
		definition: "REDACTED",
		observer:    tcpsocketLink,
		caller:      CallStreamProc(ln.Addr().String(), verifyDeadlineRetrievePersist, freshPrivateToken()),
	}
}

func posixObserverVerifyInstance(t *testing.T, deadlineEmbrace, deadlineRetrievePersist time.Duration) observerVerifyInstance {
	location, err := verifyPosixLocation()
	if err != nil {
		t.Fatal(err)
	}
	ln, err := net.Listen("REDACTED", location)
	if err != nil {
		t.Fatal(err)
	}

	posixLink := FreshPosixObserver(ln)
	PosixObserverDeadlineEmbrace(deadlineEmbrace)(posixLink)
	PosixObserverDeadlineFetchPersist(deadlineRetrievePersist)(posixLink)
	return observerVerifyInstance{
		definition: "REDACTED",
		observer:    posixLink,
		caller:      CallPosixProc(location),
	}
}

func observerVerifyScenarios(t *testing.T, deadlineEmbrace, deadlineRetrievePersist time.Duration) []observerVerifyInstance {
	return []observerVerifyInstance{
		tcpsocketObserverVerifyInstance(t, deadlineEmbrace, deadlineRetrievePersist),
		posixObserverVerifyInstance(t, deadlineEmbrace, deadlineRetrievePersist),
	}
}

func VerifyObserverDeadlineEmbrace(t *testing.T) {
	for _, tc := range observerVerifyScenarios(t, time.Millisecond, time.Second) {
		_, err := tc.observer.Accept()
		actionFault, ok := err.(*net.OpError)
		if !ok {
			t.Fatalf("REDACTED", tc.definition, err)
		}

		if possess, desire := actionFault.Op, "REDACTED"; possess != desire {
			t.Errorf("REDACTED", tc.definition, possess, desire)
		}
	}
}

func VerifyObserverDeadlineFetchPersist(t *testing.T) {
	const (
		//
		deadlineEmbrace = time.Second
		//
		//
		//
		deadlineRetrievePersist = 10 * time.Millisecond
	)

	for _, tc := range observerVerifyScenarios(t, deadlineEmbrace, deadlineRetrievePersist) {
		go func(caller PortCaller) {
			link, err := caller()
			if err != nil {
				panic(err)
			}
			//
			time.Sleep(2 * deadlineRetrievePersist)
			link.Close()
		}(tc.caller)

		c, err := tc.observer.Accept()
		if err != nil {
			t.Fatal(err)
		}

		//
		msg := make([]byte, 200)
		_, err = c.Read(msg)
		actionFault, ok := err.(*net.OpError)
		if !ok {
			t.Fatalf("REDACTED", tc.definition, err)
		}

		if possess, desire := actionFault.Op, "REDACTED"; possess != desire {
			t.Errorf("REDACTED", tc.definition, possess, desire)
		}

		if !actionFault.Timeout() {
			t.Errorf("REDACTED", tc.definition, actionFault)
		}
	}
}
