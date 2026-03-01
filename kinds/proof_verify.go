package kinds

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

var fallbackBallotMoment = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

func VerifyProofCatalog(t *testing.T) {
	ev := unpredictableReplicatedBallotProof(t)
	evl := ProofCatalog([]Proof{ev})

	assert.NotNil(t, evl.Digest())
	assert.True(t, evl.Has(ev))
	assert.False(t, evl.Has(&ReplicatedBallotProof{}))
}

func unpredictableReplicatedBallotProof(t *testing.T) *ReplicatedBallotProof {
	val := FreshSimulatePRV()
	ledgerUUID := createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	ledgerUuid2 := createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	const successionUUID = "REDACTED"
	return &ReplicatedBallotProof{
		BallotAN:            CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUUID, fallbackBallotMoment),
		BallotBYTE:            CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUuid2, fallbackBallotMoment.Add(1*time.Minute)),
		SumBallotingPotency: 30,
		AssessorPotency:   10,
		Timestamp:        fallbackBallotMoment,
	}
}

func VerifyReplicatedBallotProof(t *testing.T) {
	const altitude = int64(13)
	ev, err := FreshSimulateReplicatedBallotProof(altitude, time.Now(), "REDACTED")
	require.NoError(t, err)
	assert.Equal(t, ev.Digest(), tenderminthash.Sum(ev.Octets()))
	assert.NotNil(t, ev.Text())
	assert.Equal(t, ev.Altitude(), altitude)
}

func VerifyReplicatedBallotProofCertification(t *testing.T) {
	val := FreshSimulatePRV()
	ledgerUUID := createLedgerUUID(tenderminthash.Sum([]byte("REDACTED")), math.MaxInt32, tenderminthash.Sum([]byte("REDACTED")))
	ledgerUuid2 := createLedgerUUID(tenderminthash.Sum([]byte("REDACTED")), math.MaxInt32, tenderminthash.Sum([]byte("REDACTED")))
	const successionUUID = "REDACTED"

	verifyScenarios := []struct {
		verifyAlias         string
		distortProof func(*ReplicatedBallotProof)
		anticipateFault        bool
	}{
		{"REDACTED", func(ev *ReplicatedBallotProof) {}, false},
		{"REDACTED", func(ev *ReplicatedBallotProof) { ev.BallotAN = nil }, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) { ev.BallotBYTE = nil }, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) {
			ev.BallotAN = nil
			ev.BallotBYTE = nil
		}, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) {
			ev.BallotAN = CreateBallotNegativeFailure(t, val, successionUUID, math.MaxInt32, math.MaxInt64, math.MaxInt32, 0, ledgerUuid2, fallbackBallotMoment)
		}, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) {
			exchange := ev.BallotAN.Duplicate()
			ev.BallotAN = ev.BallotBYTE.Duplicate()
			ev.BallotBYTE = exchange
		}, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			ballot1 := CreateBallotNegativeFailure(t, val, successionUUID, math.MaxInt32, math.MaxInt64, math.MaxInt32, 0x02, ledgerUUID, fallbackBallotMoment)
			ballot2 := CreateBallotNegativeFailure(t, val, successionUUID, math.MaxInt32, math.MaxInt64, math.MaxInt32, 0x02, ledgerUuid2, fallbackBallotMoment)
			itemAssign := FreshAssessorAssign([]*Assessor{val.DeriveWithinAssessor(10)})
			ev, err := FreshReplicatedBallotProof(ballot1, ballot2, fallbackBallotMoment, itemAssign)
			require.NoError(t, err)
			tc.distortProof(ev)
			assert.Equal(t, tc.anticipateFault, ev.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func VerifyAgileCustomerOnslaughtProofFundamental(t *testing.T) {
	altitude := int64(5)
	sharedAltitude := altitude - 1
	nthAssessors := 10
	ballotAssign, itemAssign, privateItems := arbitraryBallotAssign(altitude, 1, commitchema.PreendorseKind, nthAssessors, 1, false)
	heading := createHeadingUnpredictable()
	heading.Altitude = altitude
	ledgerUUID := createLedgerUUID(tenderminthash.Sum([]byte("REDACTED")), math.MaxInt32, tenderminthash.Sum([]byte("REDACTED")))
	addnEndorse, err := CreateAddnEndorse(ledgerUUID, altitude, 1, ballotAssign, privateItems, fallbackBallotMoment, false)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	agilecustomerattackevidence := &AgileCustomerOnslaughtProof{
		DiscordantLedger: &AgileLedger{
			NotatedHeading: &NotatedHeading{
				Heading: heading,
				Endorse: endorse,
			},
			AssessorAssign: itemAssign,
		},
		SharedAltitude:        sharedAltitude,
		SumBallotingPotency:    itemAssign.SumBallotingPotency(),
		Timestamp:           heading.Moment,
		TreacherousAssessors: itemAssign.Assessors[:nthAssessors/2],
	}
	assert.NotNil(t, agilecustomerattackevidence.Text())
	assert.NotNil(t, agilecustomerattackevidence.Digest())
	assert.Equal(t, agilecustomerattackevidence.Altitude(), sharedAltitude) //
	assert.NotNil(t, agilecustomerattackevidence.Octets())

	//
	verifyScenarios := []struct {
		verifyAlias         string
		distortProof func(*AgileCustomerOnslaughtProof)
	}{
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) { ev.DiscordantLedger.Heading = createHeadingUnpredictable() }},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) {
			ev.SharedAltitude = altitude + 1
		}},
	}

	for _, tc := range verifyScenarios {
		agilecustomerattackevidence := &AgileCustomerOnslaughtProof{
			DiscordantLedger: &AgileLedger{
				NotatedHeading: &NotatedHeading{
					Heading: heading,
					Endorse: endorse,
				},
				AssessorAssign: itemAssign,
			},
			SharedAltitude:        sharedAltitude,
			SumBallotingPotency:    itemAssign.SumBallotingPotency(),
			Timestamp:           heading.Moment,
			TreacherousAssessors: itemAssign.Assessors[:nthAssessors/2],
		}
		digest := agilecustomerattackevidence.Digest()
		tc.distortProof(agilecustomerattackevidence)
		assert.NotEqual(t, digest, agilecustomerattackevidence.Digest(), tc.verifyAlias)
	}
}

func VerifyAgileCustomerOnslaughtProofCertification(t *testing.T) {
	altitude := int64(5)
	sharedAltitude := altitude - 1
	nthAssessors := 10
	ballotAssign, itemAssign, privateItems := arbitraryBallotAssign(altitude, 1, commitchema.PreendorseKind, nthAssessors, 1, false)
	heading := createHeadingUnpredictable()
	heading.Altitude = altitude
	heading.AssessorsDigest = itemAssign.Digest()
	ledgerUUID := createLedgerUUID(heading.Digest(), math.MaxInt32, tenderminthash.Sum([]byte("REDACTED")))
	addnEndorse, err := CreateAddnEndorse(ledgerUUID, altitude, 1, ballotAssign, privateItems, time.Now(), false)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	agilecustomerattackevidence := &AgileCustomerOnslaughtProof{
		DiscordantLedger: &AgileLedger{
			NotatedHeading: &NotatedHeading{
				Heading: heading,
				Endorse: endorse,
			},
			AssessorAssign: itemAssign,
		},
		SharedAltitude:        sharedAltitude,
		SumBallotingPotency:    itemAssign.SumBallotingPotency(),
		Timestamp:           heading.Moment,
		TreacherousAssessors: itemAssign.Assessors[:nthAssessors/2],
	}
	assert.NoError(t, agilecustomerattackevidence.CertifyFundamental())

	verifyScenarios := []struct {
		verifyAlias         string
		distortProof func(*AgileCustomerOnslaughtProof)
		anticipateFault        bool
	}{
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) {}, false},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) { ev.SharedAltitude = -10 }, true},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) {
			ev.SharedAltitude = altitude + 1
		}, true},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) {
			ev.SharedAltitude = altitude
		}, false},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) { ev.DiscordantLedger.Heading = nil }, true},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) { ev.DiscordantLedger = nil }, true},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) {
			ev.DiscordantLedger.AssessorAssign = &AssessorAssign{}
		}, true},
		{"REDACTED", func(ev *AgileCustomerOnslaughtProof) {
			ev.SumBallotingPotency = -1
		}, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			agilecustomerattackevidence := &AgileCustomerOnslaughtProof{
				DiscordantLedger: &AgileLedger{
					NotatedHeading: &NotatedHeading{
						Heading: heading,
						Endorse: endorse,
					},
					AssessorAssign: itemAssign,
				},
				SharedAltitude:        sharedAltitude,
				SumBallotingPotency:    itemAssign.SumBallotingPotency(),
				Timestamp:           heading.Moment,
				TreacherousAssessors: itemAssign.Assessors[:nthAssessors/2],
			}
			tc.distortProof(agilecustomerattackevidence)
			if tc.anticipateFault {
				assert.Error(t, agilecustomerattackevidence.CertifyFundamental(), tc.verifyAlias)
			} else {
				assert.NoError(t, agilecustomerattackevidence.CertifyFundamental(), tc.verifyAlias)
			}
		})
	}
}

func VerifySimulateProofCertifyFundamental(t *testing.T) {
	validProof, err := FreshSimulateReplicatedBallotProof(int64(1), time.Now(), "REDACTED")
	require.NoError(t, err)
	assert.Nil(t, validProof.CertifyFundamental())
}

func createHeadingUnpredictable() *Heading {
	return &Heading{
		Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1},
		SuccessionUUID:            commitrand.Str(12),
		Altitude:             int64(commitrand.Uint16()) + 1,
		Moment:               time.Now(),
		FinalLedgerUUID:        createLedgerUUIDUnpredictable(),
		FinalEndorseDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
		DataDigest:           security.CHARArbitraryOctets(tenderminthash.Extent),
		AssessorsDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
		FollowingAssessorsDigest: security.CHARArbitraryOctets(tenderminthash.Extent),
		AgreementDigest:      security.CHARArbitraryOctets(tenderminthash.Extent),
		PlatformDigest:            security.CHARArbitraryOctets(tenderminthash.Extent),
		FinalOutcomesDigest:    security.CHARArbitraryOctets(tenderminthash.Extent),
		ProofDigest:       security.CHARArbitraryOctets(tenderminthash.Extent),
		NominatorLocation:    security.CHARArbitraryOctets(security.LocatorExtent),
	}
}

func VerifyProofSchema(t *testing.T) {
	//
	val := FreshSimulatePRV()
	ledgerUUID := createLedgerUUID(tenderminthash.Sum([]byte("REDACTED")), math.MaxInt32, tenderminthash.Sum([]byte("REDACTED")))
	ledgerUuid2 := createLedgerUUID(tenderminthash.Sum([]byte("REDACTED")), math.MaxInt32, tenderminthash.Sum([]byte("REDACTED")))
	const successionUUID = "REDACTED"
	v := CreateBallotNegativeFailure(t, val, successionUUID, math.MaxInt32, math.MaxInt64, 1, 0x01, ledgerUUID, fallbackBallotMoment)
	v2 := CreateBallotNegativeFailure(t, val, successionUUID, math.MaxInt32, math.MaxInt64, 2, 0x01, ledgerUuid2, fallbackBallotMoment)

	//
	const altitude int64 = 37

	var (
		heading1 = createHeadingUnpredictable()
		heading2 = createHeadingUnpredictable()
	)

	heading1.Altitude = altitude
	heading1.FinalLedgerUUID = ledgerUUID
	heading1.SuccessionUUID = successionUUID

	heading2.Altitude = altitude
	heading2.FinalLedgerUUID = ledgerUUID
	heading2.SuccessionUUID = successionUUID

	verifies := []struct {
		verifyAlias     string
		proof     Proof
		towardSchemaFault   bool
		originatingSchemaFault bool
	}{
		{"REDACTED", nil, true, true},
		{"REDACTED", &ReplicatedBallotProof{}, false, true},
		{"REDACTED", &ReplicatedBallotProof{BallotAN: v, BallotBYTE: nil}, false, true},
		{"REDACTED", &ReplicatedBallotProof{BallotAN: nil, BallotBYTE: v}, false, true},
		{"REDACTED", &ReplicatedBallotProof{BallotAN: v2, BallotBYTE: v}, false, false},
	}
	for _, tt := range verifies {
		t.Run(tt.verifyAlias, func(t *testing.T) {
			pb, err := ProofTowardSchema(tt.proof)
			if tt.towardSchemaFault {
				assert.Error(t, err, tt.verifyAlias)
				return
			}
			assert.NoError(t, err, tt.verifyAlias)

			evi, err := ProofOriginatingSchema(pb)
			if tt.originatingSchemaFault {
				assert.Error(t, err, tt.verifyAlias)
				return
			}
			require.Equal(t, tt.proof, evi, tt.verifyAlias)
		})
	}
}
