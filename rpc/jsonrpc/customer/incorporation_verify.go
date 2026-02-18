//

//
//

package customer

import (
	"bytes"
	"errors"
	"net"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/log"
)

func TestWSClientReconnectWithJitter(t *testing.T) {
	n := 8
	maxReconnectAttempts := 3
	//
	maxSleepTime := time.Second * time.Duration(((1<<uint(maxReconnectAttempts))-1)+maxReconnectAttempts)

	errNotConnected := errors.New("REDACTED")
	clientMap := make(map[int]*WSClient)
	buf := new(bytes.Buffer)
	logger := log.NewTMLogger(buf)
	for i := 0; i < n; i++ {
		c, err := NewWS("REDACTED", "REDACTED")
		require.Nil(t, err)
		c.Dialer = func(string, string) (net.Conn, error) {
			return nil, errNotConnected
		}
		c.SetLogger(logger)
		c.maxReconnectAttempts = maxReconnectAttempts
		//
		//
		//
		clientMap[i] = c
		//
		go c.reconnect()
	}

	stopCount := 0
	time.Sleep(maxSleepTime)
	for key, c := range clientMap {
		if !c.IsActive() {
			delete(clientMap, key)
			stopCount++
		}
	}
	require.Equal(t, stopCount, n, "REDACTED")

	//
	backoffDurRegexp := regexp.MustCompile("REDACTED")
	matches := backoffDurRegexp.FindAll(buf.Bytes(), -1)
	seenMap := make(map[string]int)
	for i, match := range matches {
		if origIndex, seen := seenMap[string(match)]; seen {
			t.Errorf("REDACTED", i, match, origIndex)
		} else {
			seenMap[string(match)] = i
		}
	}
}
