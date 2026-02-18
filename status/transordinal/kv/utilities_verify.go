package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyIntegerInSection(t *testing.T) {
	assert.True(t, integerInSection(1, []int{1, 2, 3}))
	assert.False(t, integerInSection(4, []int{1, 2, 3}))
	assert.True(t, integerInSection(0, []int{0}))
	assert.False(t, integerInSection(0, []int{}))
}
