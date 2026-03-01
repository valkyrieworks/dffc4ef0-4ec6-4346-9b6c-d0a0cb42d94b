package agile_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	maximumTimerDeviation = 10 * time.Second
)

func VerifyValidateContiguousHeadings(t *testing.T) {
	const (
		successionUUID    = "REDACTED"
		finalAltitude = 1
		followingAltitude = 2
	)

	var (
		tokens = producePrivateTokens(4)
		//
		values     = tokens.TowardAssessors(20, 10)
		byteMoment, _ = time.Parse(time.RFC3339, "REDACTED")
		heading   = tokens.ProduceNotatedHeadline(successionUUID, finalAltitude, byteMoment, nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))
	)

	verifyScenarios := []struct {
		freshHeadline      *kinds.NotatedHeading
		freshValues        *kinds.AssessorAssign
		relyingCycle time.Duration
		now            time.Time
		expirationFault         error
		expirationFaultString     string
	}{
		//
		0: {
			heading,
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		1: {
			tokens.ProduceNotatedHeadline("REDACTED", followingAltitude, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		2: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(-1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		3: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(3*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		4: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude,
				byteMoment.Add(2*time.Hour).Add(maximumTimerDeviation).Add(-1*time.Millisecond), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		5: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		6: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 1, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		7: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(tokens)-1, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			agile.FaultUnfitHeadline{Rationale: kinds.FaultNegationAmpleBallotingPotencyNotated{Got: 50, Required: 93}},
			"REDACTED",
		},
		//
		8: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(1*time.Hour), nil, tokens.TowardAssessors(10, 1), values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			tokens.TowardAssessors(10, 1),
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		9: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			tokens.TowardAssessors(10, 1),
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		10: {
			tokens.ProduceNotatedHeadline(successionUUID, followingAltitude, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			tokens.TowardAssessors(10, 1),
			1 * time.Hour,
			byteMoment.Add(1 * time.Hour),
			nil,
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {

		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			err := agile.ValidateContiguous(heading, tc.freshHeadline, tc.freshValues, tc.relyingCycle, tc.now, maximumTimerDeviation)
			switch {
			case tc.expirationFault != nil && assert.Error(t, err):
				assert.Equal(t, tc.expirationFault, err)
			case tc.expirationFaultString != "REDACTED":
				assert.Contains(t, err.Error(), tc.expirationFaultString)
			default:
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyValidateUnContiguousHeadings(t *testing.T) {
	const (
		successionUUID    = "REDACTED"
		finalAltitude = 1
	)

	var (
		tokens = producePrivateTokens(4)
		//
		values     = tokens.TowardAssessors(20, 10)
		byteMoment, _ = time.Parse(time.RFC3339, "REDACTED")
		heading   = tokens.ProduceNotatedHeadline(successionUUID, finalAltitude, byteMoment, nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))

		//
		coupleTrinity     = tokens[1:]
		coupleTrinityValues = coupleTrinity.TowardAssessors(30, 10)

		//
		singleTertiary     = tokens[len(tokens)-1:]
		singleTertiaryValues = singleTertiary.TowardAssessors(50, 10)

		//
		inferiorOverSingleTertiary     = tokens[0:1]
		inferiorOverSingleTertiaryValues = inferiorOverSingleTertiary.TowardAssessors(20, 10)
	)

	verifyScenarios := []struct {
		freshHeadline      *kinds.NotatedHeading
		freshValues        *kinds.AssessorAssign
		relyingCycle time.Duration
		now            time.Time
		expirationFault         error
		expirationFaultString     string
	}{
		//
		0: {
			tokens.ProduceNotatedHeadline(successionUUID, 3, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		1: {
			tokens.ProduceNotatedHeadline(successionUUID, 4, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 1, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		2: {
			tokens.ProduceNotatedHeadline(successionUUID, 5, byteMoment.Add(1*time.Hour), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(tokens)-1, len(tokens)),
			values,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			agile.FaultUnfitHeadline{kinds.FaultNegationAmpleBallotingPotencyNotated{Got: 50, Required: 93}},
			"REDACTED",
		},
		//
		3: {
			coupleTrinity.ProduceNotatedHeadline(successionUUID, 5, byteMoment.Add(1*time.Hour), nil, coupleTrinityValues, coupleTrinityValues,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(coupleTrinity)),
			coupleTrinityValues,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		4: {
			singleTertiary.ProduceNotatedHeadline(successionUUID, 5, byteMoment.Add(1*time.Hour), nil, singleTertiaryValues, singleTertiaryValues,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(singleTertiary)),
			singleTertiaryValues,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			nil,
			"REDACTED",
		},
		//
		5: {
			inferiorOverSingleTertiary.ProduceNotatedHeadline(successionUUID, 5, byteMoment.Add(1*time.Hour), nil, inferiorOverSingleTertiaryValues, inferiorOverSingleTertiaryValues,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(inferiorOverSingleTertiary)),
			inferiorOverSingleTertiaryValues,
			3 * time.Hour,
			byteMoment.Add(2 * time.Hour),
			agile.FaultFreshItemAssignCannotExistReliable{kinds.FaultNegationAmpleBallotingPotencyNotated{Got: 20, Required: 46}},
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {

		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			err := agile.ValidateUnContiguous(heading, values, tc.freshHeadline, tc.freshValues, tc.relyingCycle,
				tc.now, maximumTimerDeviation,
				agile.FallbackRelianceStratum)

			switch {
			case tc.expirationFault != nil && assert.Error(t, err):
				assert.Equal(t, tc.expirationFault, err)
			case tc.expirationFaultString != "REDACTED":
				assert.Contains(t, err.Error(), tc.expirationFaultString)
			default:
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyValidateYieldsFailureConditionalRelianceStratumEqualsUnfit(t *testing.T) {
	const (
		successionUUID    = "REDACTED"
		finalAltitude = 1
	)

	var (
		tokens = producePrivateTokens(4)
		//
		values     = tokens.TowardAssessors(20, 10)
		byteMoment, _ = time.Parse(time.RFC3339, "REDACTED")
		heading   = tokens.ProduceNotatedHeadline(successionUUID, finalAltitude, byteMoment, nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))
	)

	err := agile.Validate(heading, values, heading, values, 2*time.Hour, time.Now(), maximumTimerDeviation,
		strongarithmetic.Portion{Dividend: 2, Divisor: 1})
	assert.Error(t, err)
}

func VerifyCertifyRelianceStratum(t *testing.T) {
	verifyScenarios := []struct {
		lvl   strongarithmetic.Portion
		sound bool
	}{
		//
		0: {strongarithmetic.Portion{Dividend: 1, Divisor: 1}, true},
		1: {strongarithmetic.Portion{Dividend: 1, Divisor: 3}, true},
		2: {strongarithmetic.Portion{Dividend: 2, Divisor: 3}, true},
		3: {strongarithmetic.Portion{Dividend: 3, Divisor: 3}, true},
		4: {strongarithmetic.Portion{Dividend: 4, Divisor: 5}, true},

		//
		5: {strongarithmetic.Portion{Dividend: 6, Divisor: 5}, false},
		6: {strongarithmetic.Portion{Dividend: 0, Divisor: 1}, false},
		7: {strongarithmetic.Portion{Dividend: 0, Divisor: 0}, false},
		8: {strongarithmetic.Portion{Dividend: 1, Divisor: 0}, false},
	}

	for _, tc := range verifyScenarios {
		err := agile.CertifyRelianceStratum(tc.lvl)
		if !tc.sound {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
