package strings

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func VerifyStringInSection(t *testing.T) {
	assert.True(t, StringInSection("REDACTED", []string{"REDACTED", "REDACTED", "REDACTED"}))
	assert.False(t, StringInSection("REDACTED", []string{"REDACTED", "REDACTED", "REDACTED"}))
	assert.True(t, StringInSection("REDACTED", []string{"REDACTED"}))
	assert.False(t, StringInSection("REDACTED", []string{}))
}

func VerifyIsAsciiContent(t *testing.T) {
	negateAsciiContent := []string{
		"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED",
	}
	for _, v := range negateAsciiContent {
		assert.False(t, IsAsciiContent(v), "REDACTED", v)
	}
	asciiContent := []string{
		"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED",
	}
	for _, v := range asciiContent {
		assert.True(t, IsAsciiContent(v), "REDACTED", v)
	}
}

func VerifyAsciiShave(t *testing.T) {
	assert.Equal(t, AsciiShave("REDACTED"), "REDACTED")
	assert.Equal(t, AsciiShave("REDACTED"), "REDACTED")
	assert.Equal(t, AsciiShave("REDACTED"), "REDACTED")
	assert.Equal(t, AsciiShave("REDACTED"), "REDACTED")
	assert.Panics(t, func() { AsciiShave("REDACTED") })
}

func VerifyStringSectionEquivalent(t *testing.T) {
	verifies := []struct {
		a    []string
		b    []string
		desire bool
	}{
		{[]string{"REDACTED", "REDACTED"}, []string{"REDACTED", "REDACTED"}, true},
		{[]string{"REDACTED"}, []string{"REDACTED"}, true},
		{[]string{"REDACTED"}, []string{"REDACTED"}, false},
		{[]string{"REDACTED", "REDACTED"}, []string{"REDACTED", "REDACTED"}, false},
		{[]string{"REDACTED"}, []string{"REDACTED", "REDACTED"}, false},
		{[]string{"REDACTED", "REDACTED"}, []string{"REDACTED"}, false},
	}
	for i, tt := range verifies {
		require.Equal(t, tt.desire, StringSectionEquivalent(tt.a, tt.b),
			"REDACTED", i)
	}
}
