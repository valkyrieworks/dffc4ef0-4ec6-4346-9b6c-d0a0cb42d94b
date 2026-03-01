package kinds

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

func VerifyAgileLedgerCertifyFundamental(t *testing.T) {
	heading := createArbitraryHeading()
	endorse := arbitraryEndorse(time.Now())
	values, _ := ArbitraryAssessorAssign(5, 1)
	heading.Altitude = endorse.Altitude
	heading.FinalLedgerUUID = endorse.LedgerUUID
	heading.AssessorsDigest = values.Digest()
	heading.Edition.Ledger = edition.LedgerScheme
	items2, _ := ArbitraryAssessorAssign(3, 1)
	items3 := values.Duplicate()
	items3.Nominator = &Assessor{}
	endorse.LedgerUUID.Digest = heading.Digest()

	sh := &NotatedHeading{
		Heading: &heading,
		Endorse: endorse,
	}

	verifyScenarios := []struct {
		alias      string
		sh        *NotatedHeading
		values      *AssessorAssign
		anticipateFault bool
	}{
		{"REDACTED", sh, values, false},
		{"REDACTED", sh, items2, true},
		{"REDACTED", sh, items3, true},
		{"REDACTED", &NotatedHeading{Heading: &heading, Endorse: arbitraryEndorse(time.Now())}, values, true},
	}

	for _, tc := range verifyScenarios {
		agileLedger := AgileLedger{
			NotatedHeading: tc.sh,
			AssessorAssign: tc.values,
		}
		err := agileLedger.CertifyFundamental(heading.SuccessionUUID)
		if tc.anticipateFault {
			assert.Error(t, err, tc.alias)
		} else {
			assert.NoError(t, err, tc.alias)
		}
	}
}

func VerifyAgileLedgerSchemaformat(t *testing.T) {
	heading := createArbitraryHeading()
	endorse := arbitraryEndorse(time.Now())
	values, _ := ArbitraryAssessorAssign(5, 1)
	heading.Altitude = endorse.Altitude
	heading.FinalLedgerUUID = endorse.LedgerUUID
	heading.Edition.Ledger = edition.LedgerScheme
	heading.AssessorsDigest = values.Digest()
	items3 := values.Duplicate()
	items3.Nominator = &Assessor{}
	endorse.LedgerUUID.Digest = heading.Digest()

	sh := &NotatedHeading{
		Heading: &heading,
		Endorse: endorse,
	}

	verifyScenarios := []struct {
		alias       string
		sh         *NotatedHeading
		values       *AssessorAssign
		towardSchemaFault bool
		towardLedgerFault bool
	}{
		{"REDACTED", sh, values, false, false},
		{"REDACTED", &NotatedHeading{}, values, false, false},
		{"REDACTED", sh, &AssessorAssign{}, false, true},
		{"REDACTED", &NotatedHeading{}, &AssessorAssign{}, false, true},
	}

	for _, tc := range verifyScenarios {
		agileLedger := &AgileLedger{
			NotatedHeading: tc.sh,
			AssessorAssign: tc.values,
		}
		lbp, err := agileLedger.TowardSchema()
		if tc.towardSchemaFault {
			assert.Error(t, err, tc.alias)
		} else {
			assert.NoError(t, err, tc.alias)
		}

		lb, err := AgileLedgerOriginatingSchema(lbp)
		if tc.towardLedgerFault {
			assert.Error(t, err, tc.alias)
		} else {
			assert.NoError(t, err, tc.alias)
			assert.Equal(t, agileLedger, lb)
		}
	}
}

func VerifyNotatedHeadingCertifyFundamental(t *testing.T) {
	endorse := arbitraryEndorse(time.Now())
	successionUUID := "REDACTED"
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)
	h := Heading{
		Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: math.MaxInt64},
		SuccessionUUID:            successionUUID,
		Altitude:             endorse.Altitude,
		Moment:               timestamp,
		FinalLedgerUUID:        endorse.LedgerUUID,
		FinalEndorseDigest:     endorse.Digest(),
		DataDigest:           endorse.Digest(),
		AssessorsDigest:     endorse.Digest(),
		FollowingAssessorsDigest: endorse.Digest(),
		AgreementDigest:      endorse.Digest(),
		PlatformDigest:            endorse.Digest(),
		FinalOutcomesDigest:    endorse.Digest(),
		ProofDigest:       endorse.Digest(),
		NominatorLocation:    security.LocatorDigest([]byte("REDACTED")),
	}

	soundNotatedHeading := NotatedHeading{Heading: &h, Endorse: endorse}
	soundNotatedHeading.Endorse.LedgerUUID.Digest = soundNotatedHeading.Digest()
	unfitNotatedHeading := NotatedHeading{}

	verifyScenarios := []struct {
		verifyAlias  string
		shellHeading  *Heading
		shellEndorse  *Endorse
		anticipateFault bool
	}{
		{"REDACTED", soundNotatedHeading.Heading, soundNotatedHeading.Endorse, false},
		{"REDACTED", unfitNotatedHeading.Heading, soundNotatedHeading.Endorse, true},
		{"REDACTED", soundNotatedHeading.Heading, unfitNotatedHeading.Endorse, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			sh := NotatedHeading{
				Heading: tc.shellHeading,
				Endorse: tc.shellEndorse,
			}
			err := sh.CertifyFundamental(soundNotatedHeading.SuccessionUUID)
			assert.Equalf(
				t,
				tc.anticipateFault,
				err != nil,
				"REDACTED",
				err,
			)
		})
	}
}
