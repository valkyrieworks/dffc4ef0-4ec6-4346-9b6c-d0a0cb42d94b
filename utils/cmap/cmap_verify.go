package cmap

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyRecurseKeysWithItems(t *testing.T) {
	cmap := NewCIndex()

	for i := 1; i <= 10; i++ {
		cmap.Set(fmt.Sprintf("REDACTED", i), fmt.Sprintf("REDACTED", i))
	}

	//
	assert.Equal(t, 10, cmap.Volume())
	assert.Equal(t, 10, len(cmap.Keys()))
	assert.Equal(t, 10, len(cmap.Items()))

	//
	for _, key := range cmap.Keys() {
		val := strings.ReplaceAll(key, "REDACTED", "REDACTED")
		assert.Equal(t, val, cmap.Get(key))
	}

	//
	keys := cmap.Keys()
	for i := 1; i <= 10; i++ {
		assert.Contains(t, keys, fmt.Sprintf("REDACTED", i), "REDACTED")
	}

	//
	cmap.Erase("REDACTED")

	assert.NotEqual(
		t,
		len(keys),
		len(cmap.Keys()),
		"REDACTED",
	)
}

func VerifyIncludes(t *testing.T) {
	cmap := NewCIndex()

	cmap.Set("REDACTED", "REDACTED")

	//
	assert.True(t, cmap.Has("REDACTED"))
	assert.Equal(t, "REDACTED", cmap.Get("REDACTED"))

	//
	assert.False(t, cmap.Has("REDACTED"))
	assert.Nil(t, cmap.Get("REDACTED"))
}

func CriterionCIndexHas(b *testing.B) {
	m := NewCIndex()
	for i := 0; i < 1000; i++ {
		m.Set(string(rune(i)), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Has(string(rune(i)))
	}
}
