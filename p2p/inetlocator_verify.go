package p2p

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Verifyfabricaddress_Text(t *testing.T) {
	tcpsocketLocation, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
	require.Nil(t, err)

	networkLocation := FreshNetworkLocator("REDACTED", tcpsocketLocation)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = networkLocation.Text()
		}()
	}

	wg.Wait()

	s := networkLocation.Text()
	require.Equal(t, "REDACTED", s)
}

func VerifyFreshNetworkLocator(t *testing.T) {
	tcpsocketLocation, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
	require.Nil(t, err)

	assert.Panics(t, func() {
		FreshNetworkLocator("REDACTED", tcpsocketLocation)
	})

	location := FreshNetworkLocator("REDACTED", tcpsocketLocation)
	assert.Equal(t, "REDACTED", location.Text())

	assert.NotPanics(t, func() {
		FreshNetworkLocator("REDACTED", &net.UDPAddr{IP: net.ParseIP("REDACTED"), Port: 8000})
	}, "REDACTED")
}

func VerifyFreshNetworkLocatorText(t *testing.T) {
	verifyScenarios := []struct {
		alias     string
		location     string
		anticipated string
		precise  bool
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
		t.Run(tc.alias, func(t *testing.T) {
			location, err := FreshNetworkLocatorText(tc.location)
			if tc.precise {
				if assert.Nil(t, err, tc.location) {
					assert.Equal(t, tc.anticipated, location.Text())
				}
			} else {
				assert.NotNil(t, err, tc.location)
			}
		})
	}
}

func VerifyFreshNetworkLocatorTexts(t *testing.T) {
	locations, errors := FreshNetworkLocatorTexts([]string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
	})
	assert.Len(t, errors, 1)
	assert.Equal(t, 2, len(locations))
}

func VerifyFreshNetworkLocatorINETChannel(t *testing.T) {
	location := FreshNetworkLocatorINETChannel(net.ParseIP("REDACTED"), 8080)
	assert.Equal(t, "REDACTED", location.Text())
}

func VerifyNetworkLocatorAttributes(t *testing.T) {
	//
	verifyScenarios := []struct {
		location     string
		sound    bool
		regional    bool
		directable bool
	}{
		{"REDACTED", true, true, false},
		{"REDACTED", true, false, true},
	}

	for _, tc := range verifyScenarios {
		location, err := FreshNetworkLocatorText(tc.location)
		require.Nil(t, err)

		err = location.Sound()
		if tc.sound {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
		assert.Equal(t, tc.regional, location.Regional())
		assert.Equal(t, tc.directable, location.Directable())
	}
}

func VerifyNetworkLocatorAccessibilityToward(t *testing.T) {
	//
	verifyScenarios := []struct {
		location         string
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
		location, err := FreshNetworkLocatorText(tc.location)
		require.Nil(t, err)

		another, err := FreshNetworkLocatorText(tc.another)
		require.Nil(t, err)

		assert.Equal(t, tc.accessibility, location.AccessibilityToward(another))
	}
}
