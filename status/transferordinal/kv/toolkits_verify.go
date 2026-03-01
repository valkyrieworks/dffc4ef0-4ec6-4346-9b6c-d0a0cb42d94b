package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyIntegerInsideSection(t *testing.T) {
	assert.True(t, integerInsideSection(1, []int{1, 2, 3}))
	assert.False(t, integerInsideSection(4, []int{1, 2, 3}))
	assert.True(t, integerInsideSection(0, []int{0}))
	assert.False(t, integerInsideSection(0, []int{}))
}
