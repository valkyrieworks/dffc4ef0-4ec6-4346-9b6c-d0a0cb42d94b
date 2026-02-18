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

func VerifyArgumentToJSON(t *testing.T) {
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
		args := map[string]any{"REDACTED": tc.influx}
		err := argsToJSON(args)
		require.Nil(err, "REDACTED", i, err)
		require.Equal(1, len(args), "REDACTED", i)
		data, ok := args["REDACTED"].(string)
		require.True(ok, "REDACTED", i, args["REDACTED"])
		assert.Equal(tc.anticipated, data, "REDACTED", i)
	}
}
