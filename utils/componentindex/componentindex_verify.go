package componentindex

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyTraverseTokensUsingItems(t *testing.T) {
	componentindex := FreshCNIndex()

	for i := 1; i <= 10; i++ {
		componentindex.Set(fmt.Sprintf("REDACTED", i), fmt.Sprintf("REDACTED", i))
	}

	//
	assert.Equal(t, 10, componentindex.Extent())
	assert.Equal(t, 10, len(componentindex.Tokens()))
	assert.Equal(t, 10, len(componentindex.Items()))

	//
	for _, key := range componentindex.Tokens() {
		val := strings.ReplaceAll(key, "REDACTED", "REDACTED")
		assert.Equal(t, val, componentindex.Get(key))
	}

	//
	tokens := componentindex.Tokens()
	for i := 1; i <= 10; i++ {
		assert.Contains(t, tokens, fmt.Sprintf("REDACTED", i), "REDACTED")
	}

	//
	componentindex.Erase("REDACTED")

	assert.NotEqual(
		t,
		len(tokens),
		len(componentindex.Tokens()),
		"REDACTED",
	)
}

func VerifyIncludes(t *testing.T) {
	componentindex := FreshCNIndex()

	componentindex.Set("REDACTED", "REDACTED")

	//
	assert.True(t, componentindex.Has("REDACTED"))
	assert.Equal(t, "REDACTED", componentindex.Get("REDACTED"))

	//
	assert.False(t, componentindex.Has("REDACTED"))
	assert.Nil(t, componentindex.Get("REDACTED"))
}

func AssessmentCNIndexOwns(b *testing.B) {
	m := FreshCNIndex()
	for i := 0; i < 1000; i++ {
		m.Set(string(rune(i)), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Has(string(rune(i)))
	}
}
