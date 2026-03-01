package netp2p

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyFaults(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		internalFault := errors.New("REDACTED")
		fleetingFault := &FailureFleeting{Err: internalFault}

		for _, tt := range []struct {
			alias    string
			influx   any
			desireOKAY  bool
			desireFault error
		}{
			{
				alias:   "REDACTED",
				influx:  "REDACTED",
				desireOKAY: false,
			},
			{
				alias:   "REDACTED",
				influx:  errors.New("REDACTED"),
				desireOKAY: false,
			},
			{
				alias:    "REDACTED",
				influx:   fleetingFault,
				desireOKAY:  true,
				desireFault: internalFault,
			},
		} {
			t.Run(tt.alias, func(t *testing.T) {
				//
				got, ok := FleetingFailureOriginatingSome(tt.influx)

				//
				require.Equal(t, tt.desireOKAY, ok)
				if !tt.desireOKAY {
					assert.Nil(t, got)
					return
				}

				require.NotNil(t, got)
				assert.ErrorIs(t, got.Err, tt.desireFault)
			})
		}
	})
}
