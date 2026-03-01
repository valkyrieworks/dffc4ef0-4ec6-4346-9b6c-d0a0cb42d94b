package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyPagingScreen(t *testing.T) {
	scenarios := []struct {
		sumTally int
		everyScreen    int
		screen       int
		freshScreen    int
		expirationFault     bool
	}{
		{0, 10, 1, 1, false},

		{0, 10, 0, 1, false},
		{0, 10, 1, 1, false},
		{0, 10, 2, 0, true},

		{5, 10, -1, 0, true},
		{5, 10, 0, 1, false},
		{5, 10, 1, 1, false},
		{5, 10, 2, 0, true},
		{5, 10, 2, 0, true},

		{5, 5, 1, 1, false},
		{5, 5, 2, 0, true},
		{5, 5, 3, 0, true},

		{5, 3, 2, 2, false},
		{5, 3, 3, 0, true},

		{5, 2, 2, 2, false},
		{5, 2, 3, 3, false},
		{5, 2, 4, 0, true},
	}

	for _, c := range scenarios {
		p, err := certifyScreen(&c.screen, c.everyScreen, c.sumTally)
		if c.expirationFault {
			assert.Error(t, err)
			continue
		}

		assert.Equal(t, c.freshScreen, p, fmt.Sprintf("REDACTED", c))
	}

	//
	p, err := certifyScreen(nil, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, p)
	}
}

func VerifyPagingEveryScreen(t *testing.T) {
	scenarios := []struct {
		sumTally int
		everyScreen    int
		freshEveryScreen int
	}{
		{5, 0, fallbackEveryScreen},
		{5, 1, 1},
		{5, 2, 2},
		{5, fallbackEveryScreen, fallbackEveryScreen},
		{5, maximumEveryScreen - 1, maximumEveryScreen - 1},
		{5, maximumEveryScreen, maximumEveryScreen},
		{5, maximumEveryScreen + 1, maximumEveryScreen},
	}
	env := &Context{}
	for _, c := range scenarios {
		p := env.certifyEveryScreen(&c.everyScreen)
		assert.Equal(t, c.freshEveryScreen, p, fmt.Sprintf("REDACTED", c))
	}

	//
	p := env.certifyEveryScreen(nil)
	assert.Equal(t, fallbackEveryScreen, p)
}
