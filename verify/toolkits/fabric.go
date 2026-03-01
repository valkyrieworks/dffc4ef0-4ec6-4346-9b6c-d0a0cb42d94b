package toolkits

import (
	"net"
	"testing"
)

//
func ObtainReleaseChannels(t *testing.T, n int) []int {
	var (
		channels     = make([]int, 0, n)
		observers = make([]net.Listener, 0, n)
	)

	for i := 0; i < n; i++ {
		location, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
		if err != nil {
			t.Fatalf("REDACTED", err)
		}

		l, err := net.ListenTCP("REDACTED", location)
		if err != nil {
			t.Fatalf("REDACTED", err)
		}

		channel := l.Addr().(*net.TCPAddr).Port
		channels = append(channels, channel)

		//
		observers = append(observers, l)
	}

	//
	for _, l := range observers {
		if err := l.Close(); err != nil {
			t.Fatalf("REDACTED", err)
		}
	}

	return channels
}
