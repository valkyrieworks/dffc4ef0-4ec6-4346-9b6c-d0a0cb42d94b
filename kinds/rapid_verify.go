package kinds

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/vault"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/release"
)

func VerifyRapidLedgerCertifySimple(t *testing.T) {
	heading := createRandomHeading()
	endorse := randomEndorse(time.Now())
	values, _ := RandomRatifierCollection(5, 1)
	heading.Level = endorse.Level
	heading.FinalLedgerUID = endorse.LedgerUID
	heading.RatifiersDigest = values.Digest()
	heading.Release.Ledger = release.LedgerProtocol
	values2, _ := RandomRatifierCollection(3, 1)
	nodes3 := values.Clone()
	nodes3.Recommender = &Ratifier{}
	endorse.LedgerUID.Digest = heading.Digest()

	sh := &AttestedHeading{
		Heading: &heading,
		Endorse: endorse,
	}

	verifyScenarios := []struct {
		label      string
		sh        *AttestedHeading
		values      *RatifierAssign
		anticipateErr bool
	}{
		{"REDACTED", sh, values, false},
		{"REDACTED", sh, values2, true},
		{"REDACTED", sh, nodes3, true},
		{"REDACTED", &AttestedHeading{Heading: &heading, Endorse: randomEndorse(time.Now())}, values, true},
	}

	for _, tc := range verifyScenarios {
		rapidLedger := RapidLedger{
			AttestedHeading: tc.sh,
			RatifierAssign: tc.values,
		}
		err := rapidLedger.CertifySimple(heading.LedgerUID)
		if tc.anticipateErr {
			assert.Error(t, err, tc.label)
		} else {
			assert.NoError(t, err, tc.label)
		}
	}
}

func VerifyRapidLedgerProtobuf(t *testing.T) {
	heading := createRandomHeading()
	endorse := randomEndorse(time.Now())
	values, _ := RandomRatifierCollection(5, 1)
	heading.Level = endorse.Level
	heading.FinalLedgerUID = endorse.LedgerUID
	heading.Release.Ledger = release.LedgerProtocol
	heading.RatifiersDigest = values.Digest()
	nodes3 := values.Clone()
	nodes3.Recommender = &Ratifier{}
	endorse.LedgerUID.Digest = heading.Digest()

	sh := &AttestedHeading{
		Heading: &heading,
		Endorse: endorse,
	}

	verifyScenarios := []struct {
		label       string
		sh         *AttestedHeading
		values       *RatifierAssign
		toSchemaErr bool
		toLedgerErr bool
	}{
		{"REDACTED", sh, values, false, false},
		{"REDACTED", &AttestedHeading{}, values, false, false},
		{"REDACTED", sh, &RatifierAssign{}, false, true},
		{"REDACTED", &AttestedHeading{}, &RatifierAssign{}, false, true},
	}

	for _, tc := range verifyScenarios {
		rapidLedger := &RapidLedger{
			AttestedHeading: tc.sh,
			RatifierAssign: tc.values,
		}
		lbp, err := rapidLedger.ToSchema()
		if tc.toSchemaErr {
			assert.Error(t, err, tc.label)
		} else {
			assert.NoError(t, err, tc.label)
		}

		lb, err := RapidLedgerFromSchema(lbp)
		if tc.toLedgerErr {
			assert.Error(t, err, tc.label)
		} else {
			assert.NoError(t, err, tc.label)
			assert.Equal(t, rapidLedger, lb)
		}
	}
}

func VerifyAttestedHeadingCertifySimple(t *testing.T) {
	endorse := randomEndorse(time.Now())
	ledgerUID := "REDACTED"
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)
	h := Heading{
		Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: math.MaxInt64},
		LedgerUID:            ledgerUID,
		Level:             endorse.Level,
		Time:               timestamp,
		FinalLedgerUID:        endorse.LedgerUID,
		FinalEndorseDigest:     endorse.Digest(),
		DataDigest:           endorse.Digest(),
		RatifiersDigest:     endorse.Digest(),
		FollowingRatifiersDigest: endorse.Digest(),
		AgreementDigest:      endorse.Digest(),
		ApplicationDigest:            endorse.Digest(),
		FinalOutcomesDigest:    endorse.Digest(),
		ProofDigest:       endorse.Digest(),
		RecommenderLocation:    vault.LocationDigest([]byte("REDACTED")),
	}

	soundAttestedHeading := AttestedHeading{Heading: &h, Endorse: endorse}
	soundAttestedHeading.Endorse.LedgerUID.Digest = soundAttestedHeading.Digest()
	corruptAttestedHeading := AttestedHeading{}

	verifyScenarios := []struct {
		verifyLabel  string
		shHeading  *Heading
		shEndorse  *Endorse
		anticipateErr bool
	}{
		{"REDACTED", soundAttestedHeading.Heading, soundAttestedHeading.Endorse, false},
		{"REDACTED", corruptAttestedHeading.Heading, soundAttestedHeading.Endorse, true},
		{"REDACTED", soundAttestedHeading.Heading, corruptAttestedHeading.Endorse, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			sh := AttestedHeading{
				Heading: tc.shHeading,
				Endorse: tc.shEndorse,
			}
			err := sh.CertifySimple(soundAttestedHeading.LedgerUID)
			assert.Equalf(
				t,
				tc.anticipateErr,
				err != nil,
				"REDACTED",
				err,
			)
		})
	}
}
