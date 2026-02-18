package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyPagingScreen(t *testing.T) {
	scenarios := []struct {
		sumNumber int
		eachScreen    int
		screen       int
		newScreen    int
		expirationErr     bool
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
		p, err := certifyScreen(&c.screen, c.eachScreen, c.sumNumber)
		if c.expirationErr {
			assert.Error(t, err)
			continue
		}

		assert.Equal(t, c.newScreen, p, fmt.Sprintf("REDACTED", c))
	}

	//
	p, err := certifyScreen(nil, 1, 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, p)
	}
}

func VerifyPagingEachScreen(t *testing.T) {
	scenarios := []struct {
		sumNumber int
		eachScreen    int
		newEachScreen int
	}{
		{5, 0, standardEachScreen},
		{5, 1, 1},
		{5, 2, 2},
		{5, standardEachScreen, standardEachScreen},
		{5, maximumEachScreen - 1, maximumEachScreen - 1},
		{5, maximumEachScreen, maximumEachScreen},
		{5, maximumEachScreen + 1, maximumEachScreen},
	}
	env := &Context{}
	for _, c := range scenarios {
		p := env.certifyEachScreen(&c.eachScreen)
		assert.Equal(t, c.newEachScreen, p, fmt.Sprintf("REDACTED", c))
	}

	//
	p := env.certifyEachScreen(nil)
	assert.Equal(t, standardEachScreen, p)
}
