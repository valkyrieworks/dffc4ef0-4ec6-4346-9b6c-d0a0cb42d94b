package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Tx []byte

type Foo struct {
	Bar int
	Baz string
}

func VerifyArgumentTowardJSN(t *testing.T) {
	affirm := assert.New(t)
	demand := require.New(t)

	scenarios := []struct {
		influx    any
		anticipated string
	}{
		{[]byte("REDACTED"), "REDACTED"},
		{Tx("REDACTED"), "REDACTED"},
		{Foo{7, "REDACTED"}, "REDACTED"},
	}

	for i, tc := range scenarios {
		arguments := map[string]any{"REDACTED": tc.influx}
		err := argumentsTowardJSN(arguments)
		require.Nil(err, "REDACTED", i, err)
		require.Equal(1, len(arguments), "REDACTED", i)
		data, ok := arguments["REDACTED"].(string)
		require.True(ok, "REDACTED", i, arguments["REDACTED"])
		assert.Equal(tc.anticipated, data, "REDACTED", i)
	}
}
