package p2p

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Testnetaddress_String(t *testing.T) {
	tcpAddress, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
	require.Nil(t, err)

	netAddress := NewNetLocation("REDACTED", tcpAddress)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = netAddress.String()
		}()
	}

	wg.Wait()

	s := netAddress.String()
	require.Equal(t, "REDACTED", s)
}

func VerifyNewNetLocation(t *testing.T) {
	tcpAddress, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
	require.Nil(t, err)

	assert.Panics(t, func() {
		NewNetLocation("REDACTED", tcpAddress)
	})

	address := NewNetLocation("REDACTED", tcpAddress)
	assert.Equal(t, "REDACTED", address.String())

	assert.NotPanics(t, func() {
		NewNetLocation("REDACTED", &net.UDPAddr{IP: net.ParseIP("REDACTED"), Port: 8000})
	}, "REDACTED")
}

func VerifyNewNetLocationString(t *testing.T) {
	verifyScenarios := []struct {
		label     string
		address     string
		anticipated string
		accurate  bool
	}{
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},

		{
			"REDACTED",
			"REDACTED",
			"REDACTED",
			true,
		},
		{
			"REDACTED",
			"REDACTED",
			"REDACTED",
			true,
		},
		{
			"REDACTED",
			"REDACTED",
			"REDACTED",
			true,
		},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},

		//
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},

		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},

		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{
			"REDACTED",
			"REDACTED",
			"REDACTED",
			true,
		},

		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
		{"REDACTED", "REDACTED", "REDACTED", false},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			address, err := NewNetLocationString(tc.address)
			if tc.accurate {
				if assert.Nil(t, err, tc.address) {
					assert.Equal(t, tc.anticipated, address.String())
				}
			} else {
				assert.NotNil(t, err, tc.address)
			}
		})
	}
}

func VerifyNewNetLocationStrings(t *testing.T) {
	locations, faults := NewNetLocationStrings([]string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
	})
	assert.Len(t, faults, 1)
	assert.Equal(t, 2, len(locations))
}

func VerifyNewNetLocationIPPort(t *testing.T) {
	address := NewNetLocationIPPort(net.ParseIP("REDACTED"), 8080)
	assert.Equal(t, "REDACTED", address.String())
}

func VerifyNetLocationAttributes(t *testing.T) {
	//
	verifyScenarios := []struct {
		address     string
		sound    bool
		native    bool
		forwardable bool
	}{
		{"REDACTED", true, true, false},
		{"REDACTED", true, false, true},
	}

	for _, tc := range verifyScenarios {
		address, err := NewNetLocationString(tc.address)
		require.Nil(t, err)

		err = address.Sound()
		if tc.sound {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
		assert.Equal(t, tc.native, address.Native())
		assert.Equal(t, tc.forwardable, address.Forwardable())
	}
}

func VerifyNetLocationAccessibilityTo(t *testing.T) {
	//
	verifyScenarios := []struct {
		address         string
		another        string
		accessibility int
	}{
		{
			"REDACTED",
			"REDACTED",
			0,
		},
		{"REDACTED", "REDACTED", 1},
	}

	for _, tc := range verifyScenarios {
		address, err := NewNetLocationString(tc.address)
		require.Nil(t, err)

		another, err := NewNetLocationString(tc.another)
		require.Nil(t, err)

		assert.Equal(t, tc.accessibility, address.AccessibilityTo(another))
	}
}
