package utilities

import (
	"net"
	"testing"
)

//
func FetchReleasePorts(t *testing.T, n int) []int {
	var (
		ports     = make([]int, 0, n)
		observers = make([]net.Listener, 0, n)
	)

	for i := 0; i < n; i++ {
		address, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
		if err != nil {
			t.Fatalf("REDACTED", err)
		}

		l, err := net.ListenTCP("REDACTED", address)
		if err != nil {
			t.Fatalf("REDACTED", err)
		}

		port := l.Addr().(*net.TCPAddr).Port
		ports = append(ports, port)

		//
		observers = append(observers, l)
	}

	//
	for _, l := range observers {
		if err := l.Close(); err != nil {
			t.Fatalf("REDACTED", err)
		}
	}

	return ports
}
