package texts

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func VerifyTextInsideSection(t *testing.T) {
	assert.True(t, TextInsideSection("REDACTED", []string{"REDACTED", "REDACTED", "REDACTED"}))
	assert.False(t, TextInsideSection("REDACTED", []string{"REDACTED", "REDACTED", "REDACTED"}))
	assert.True(t, TextInsideSection("REDACTED", []string{"REDACTED"}))
	assert.False(t, TextInsideSection("REDACTED", []string{}))
}

func VerifyEqualsCODETxt(t *testing.T) {
	negationCODETxt := []string{
		"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED",
	}
	for _, v := range negationCODETxt {
		assert.False(t, EqualsCODETxt(v), "REDACTED", v)
	}
	codeTxt := []string{
		"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED",
	}
	for _, v := range codeTxt {
		assert.True(t, EqualsCODETxt(v), "REDACTED", v)
	}
}

func VerifyCODEShave(t *testing.T) {
	assert.Equal(t, CODEShave("REDACTED"), "REDACTED")
	assert.Equal(t, CODEShave("REDACTED"), "REDACTED")
	assert.Equal(t, CODEShave("REDACTED"), "REDACTED")
	assert.Equal(t, CODEShave("REDACTED"), "REDACTED")
	assert.Panics(t, func() { CODEShave("REDACTED") })
}

func VerifyTextSectionEquivalent(t *testing.T) {
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
		require.Equal(t, tt.desire, TextSectionEquivalent(tt.a, tt.b),
			"REDACTED", i)
	}
}
