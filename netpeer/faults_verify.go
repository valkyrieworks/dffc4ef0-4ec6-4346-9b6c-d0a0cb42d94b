package netpeer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyFaults(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		deeperErr := errors.New("REDACTED")
		temporaryErr := &FaultTemporary{Err: deeperErr}

		for _, tt := range []struct {
			label    string
			influx   any
			desireOK  bool
			desireErr error
		}{
			{
				label:   "REDACTED",
				influx:  "REDACTED",
				desireOK: false,
			},
			{
				label:   "REDACTED",
				influx:  errors.New("REDACTED"),
				desireOK: false,
			},
			{
				label:    "REDACTED",
				influx:   temporaryErr,
				desireOK:  true,
				desireErr: deeperErr,
			},
		} {
			t.Run(tt.label, func(t *testing.T) {
				//
				got, ok := TemporaryFaultFromAny(tt.influx)

				//
				require.Equal(t, tt.desireOK, ok)
				if !tt.desireOK {
					assert.Nil(t, got)
					return
				}

				require.NotNil(t, got)
				assert.ErrorIs(t, got.Err, tt.desireErr)
			})
		}
	})
}
