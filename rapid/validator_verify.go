package rapid_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cometmath "github.com/valkyrieworks/utils/math"
	"github.com/valkyrieworks/rapid"
	"github.com/valkyrieworks/kinds"
)

const (
	maximumTimerDeviation = 10 * time.Second
)

func VerifyValidateNeighboringHeadings(t *testing.T) {
	const (
		ledgerUID    = "REDACTED"
		finalLevel = 1
		followingLevel = 2
	)

	var (
		keys = generatePrivateKeys(4)
		//
		values     = keys.ToRatifiers(20, 10)
		byteTime, _ = time.Parse(time.RFC3339, "REDACTED")
		heading   = keys.GenerateAttestedHeading(ledgerUID, finalLevel, byteTime, nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))
	)

	verifyScenarios := []struct {
		newHeading      *kinds.AttestedHeading
		newValues        *kinds.RatifierAssign
		validatingDuration time.Duration
		now            time.Time
		expirationErr         error
		expirationErrContent     string
	}{
		//
		0: {
			heading,
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		1: {
			keys.GenerateAttestedHeading("REDACTED", followingLevel, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		2: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(-1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		3: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(3*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		4: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel,
				byteTime.Add(2*time.Hour).Add(maximumTimerDeviation).Add(-1*time.Millisecond), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		5: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		6: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 1, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		7: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(keys)-1, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			rapid.ErrCorruptHeading{Cause: kinds.ErrNoSufficientPollingEnergyAttested{Got: 50, Required: 93}},
			"REDACTED",
		},
		//
		8: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(1*time.Hour), nil, keys.ToRatifiers(10, 1), values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			keys.ToRatifiers(10, 1),
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		9: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			keys.ToRatifiers(10, 1),
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		10: {
			keys.GenerateAttestedHeading(ledgerUID, followingLevel, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			keys.ToRatifiers(10, 1),
			1 * time.Hour,
			byteTime.Add(1 * time.Hour),
			nil,
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {

		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			err := rapid.ValidateNeighboring(heading, tc.newHeading, tc.newValues, tc.validatingDuration, tc.now, maximumTimerDeviation)
			switch {
			case tc.expirationErr != nil && assert.Error(t, err):
				assert.Equal(t, tc.expirationErr, err)
			case tc.expirationErrContent != "REDACTED":
				assert.Contains(t, err.Error(), tc.expirationErrContent)
			default:
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyValidateNotNeighboringHeadings(t *testing.T) {
	const (
		ledgerUID    = "REDACTED"
		finalLevel = 1
	)

	var (
		keys = generatePrivateKeys(4)
		//
		values     = keys.ToRatifiers(20, 10)
		byteTime, _ = time.Parse(time.RFC3339, "REDACTED")
		heading   = keys.GenerateAttestedHeading(ledgerUID, finalLevel, byteTime, nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))

		//
		dualThirds     = keys[1:]
		dualThirdsValues = dualThirds.ToRatifiers(30, 10)

		//
		oneTertiary     = keys[len(keys)-1:]
		oneTertiaryValues = oneTertiary.ToRatifiers(50, 10)

		//
		lowerThanOneTertiary     = keys[0:1]
		lowerThanOneTertiaryValues = lowerThanOneTertiary.ToRatifiers(20, 10)
	)

	verifyScenarios := []struct {
		newHeading      *kinds.AttestedHeading
		newValues        *kinds.RatifierAssign
		validatingDuration time.Duration
		now            time.Time
		expirationErr         error
		expirationErrContent     string
	}{
		//
		0: {
			keys.GenerateAttestedHeading(ledgerUID, 3, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		1: {
			keys.GenerateAttestedHeading(ledgerUID, 4, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 1, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		2: {
			keys.GenerateAttestedHeading(ledgerUID, 5, byteTime.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(keys)-1, len(keys)),
			values,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			rapid.ErrCorruptHeading{kinds.ErrNoSufficientPollingEnergyAttested{Got: 50, Required: 93}},
			"REDACTED",
		},
		//
		3: {
			dualThirds.GenerateAttestedHeading(ledgerUID, 5, byteTime.Add(1*time.Hour), nil, dualThirdsValues, dualThirdsValues,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(dualThirds)),
			dualThirdsValues,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		4: {
			oneTertiary.GenerateAttestedHeading(ledgerUID, 5, byteTime.Add(1*time.Hour), nil, oneTertiaryValues, oneTertiaryValues,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(oneTertiary)),
			oneTertiaryValues,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		5: {
			lowerThanOneTertiary.GenerateAttestedHeading(ledgerUID, 5, byteTime.Add(1*time.Hour), nil, lowerThanOneTertiaryValues, lowerThanOneTertiaryValues,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(lowerThanOneTertiary)),
			lowerThanOneTertiaryValues,
			3 * time.Hour,
			byteTime.Add(2 * time.Hour),
			rapid.ErrNewValueCollectionCannotBeValidated{kinds.ErrNoSufficientPollingEnergyAttested{Got: 20, Required: 46}},
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {

		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			err := rapid.ValidateNotNeighboring(heading, values, tc.newHeading, tc.newValues, tc.validatingDuration,
				tc.now, maximumTimerDeviation,
				rapid.StandardRelianceLayer)

			switch {
			case tc.expirationErr != nil && assert.Error(t, err):
				assert.Equal(t, tc.expirationErr, err)
			case tc.expirationErrContent != "REDACTED":
				assert.Contains(t, err.Error(), tc.expirationErrContent)
			default:
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyValidateYieldsFaultIfRelianceLayerIsCorrupt(t *testing.T) {
	const (
		ledgerUID    = "REDACTED"
		finalLevel = 1
	)

	var (
		keys = generatePrivateKeys(4)
		//
		values     = keys.ToRatifiers(20, 10)
		byteTime, _ = time.Parse(time.RFC3339, "REDACTED")
		heading   = keys.GenerateAttestedHeading(ledgerUID, finalLevel, byteTime, nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))
	)

	err := rapid.Validate(heading, values, heading, values, 2*time.Hour, time.Now(), maximumTimerDeviation,
		cometmath.Portion{Dividend: 2, Divisor: 1})
	assert.Error(t, err)
}

func VerifyCertifyRelianceLayer(t *testing.T) {
	verifyScenarios := []struct {
		lvl   cometmath.Portion
		sound bool
	}{
		//
		0: {cometmath.Portion{Dividend: 1, Divisor: 1}, true},
		1: {cometmath.Portion{Dividend: 1, Divisor: 3}, true},
		2: {cometmath.Portion{Dividend: 2, Divisor: 3}, true},
		3: {cometmath.Portion{Dividend: 3, Divisor: 3}, true},
		4: {cometmath.Portion{Dividend: 4, Divisor: 5}, true},

		//
		5: {cometmath.Portion{Dividend: 6, Divisor: 5}, false},
		6: {cometmath.Portion{Dividend: 0, Divisor: 1}, false},
		7: {cometmath.Portion{Dividend: 0, Divisor: 0}, false},
		8: {cometmath.Portion{Dividend: 1, Divisor: 0}, false},
	}

	for _, tc := range verifyScenarios {
		err := rapid.CertifyRelianceLayer(tc.lvl)
		if !tc.sound {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
