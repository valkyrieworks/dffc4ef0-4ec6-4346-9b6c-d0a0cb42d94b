package privatekey

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/valkyrieworks/vault/ed25519"
)

//
//

func newPrivateKey() ed25519.PrivateKey {
	return ed25519.GeneratePrivateKey()
}

//
//

type observerVerifyScenario struct {
	summary string //
	observer    net.Listener
	caller      SocketCaller
}

//
//
func verifyUnixAddress() (string, error) {
	f, err := os.CreateTemp("REDACTED", "REDACTED")
	if err != nil {
		return "REDACTED", err
	}
	address := f.Name()
	f.Close()
	os.Remove(address)
	return address, nil
}

func tcpObserverVerifyScenario(t *testing.T, deadlineAllow, deadlineScanRecord time.Duration) observerVerifyScenario {
	ln, err := net.Listen("REDACTED", "REDACTED")
	if err != nil {
		t.Fatal(err)
	}

	tcpLn := NewTCPObserver(ln, newPrivateKey())
	TCPObserverDeadlineAllow(deadlineAllow)(tcpLn)
	TCPObserverDeadlineScanRecord(deadlineScanRecord)(tcpLn)
	return observerVerifyScenario{
		summary: "REDACTED",
		observer:    tcpLn,
		caller:      CallTCPFn(ln.Addr().String(), verifyDeadlineScanRecord, newPrivateKey()),
	}
}

func unixObserverVerifyScenario(t *testing.T, deadlineAllow, deadlineScanRecord time.Duration) observerVerifyScenario {
	address, err := verifyUnixAddress()
	if err != nil {
		t.Fatal(err)
	}
	ln, err := net.Listen("REDACTED", address)
	if err != nil {
		t.Fatal(err)
	}

	unixLn := NewUnixObserver(ln)
	UnixObserverDeadlineAllow(deadlineAllow)(unixLn)
	UnixObserverDeadlineScanRecord(deadlineScanRecord)(unixLn)
	return observerVerifyScenario{
		summary: "REDACTED",
		observer:    unixLn,
		caller:      CallUnixFn(address),
	}
}

func observerVerifyScenarios(t *testing.T, deadlineAllow, deadlineScanRecord time.Duration) []observerVerifyScenario {
	return []observerVerifyScenario{
		tcpObserverVerifyScenario(t, deadlineAllow, deadlineScanRecord),
		unixObserverVerifyScenario(t, deadlineAllow, deadlineScanRecord),
	}
}

func VerifyObserverDeadlineAllow(t *testing.T) {
	for _, tc := range observerVerifyScenarios(t, time.Millisecond, time.Second) {
		_, err := tc.observer.Accept()
		actErr, ok := err.(*net.OpError)
		if !ok {
			t.Fatalf("REDACTED", tc.summary, err)
		}

		if possess, desire := actErr.Op, "REDACTED"; possess != desire {
			t.Errorf("REDACTED", tc.summary, possess, desire)
		}
	}
}

func VerifyObserverDeadlineScanRecord(t *testing.T) {
	const (
		//
		deadlineAllow = time.Second
		//
		//
		//
		deadlineScanRecord = 10 * time.Millisecond
	)

	for _, tc := range observerVerifyScenarios(t, deadlineAllow, deadlineScanRecord) {
		go func(caller SocketCaller) {
			link, err := caller()
			if err != nil {
				panic(err)
			}
			//
			time.Sleep(2 * deadlineScanRecord)
			link.Close()
		}(tc.caller)

		c, err := tc.observer.Accept()
		if err != nil {
			t.Fatal(err)
		}

		//
		msg := make([]byte, 200)
		_, err = c.Read(msg)
		actErr, ok := err.(*net.OpError)
		if !ok {
			t.Fatalf("REDACTED", tc.summary, err)
		}

		if possess, desire := actErr.Op, "REDACTED"; possess != desire {
			t.Errorf("REDACTED", tc.summary, possess, desire)
		}

		if !actErr.Timeout() {
			t.Errorf("REDACTED", tc.summary, actErr)
		}
	}
}
