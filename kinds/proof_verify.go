package kinds

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/release"
)

var standardBallotTime = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

func VerifyProofCatalog(t *testing.T) {
	ev := arbitraryReplicatedBallotProof(t)
	evl := ProofCatalog([]Proof{ev})

	assert.NotNil(t, evl.Digest())
	assert.True(t, evl.Has(ev))
	assert.False(t, evl.Has(&ReplicatedBallotProof{}))
}

func arbitraryReplicatedBallotProof(t *testing.T) *ReplicatedBallotProof {
	val := NewEmulatePV()
	ledgerUID := createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	ledgerUidtwo := createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	const ledgerUID = "REDACTED"
	return &ReplicatedBallotProof{
		BallotA:            CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerUID, standardBallotTime),
		BallotBYTE:            CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerUidtwo, standardBallotTime.Add(1*time.Minute)),
		SumPollingEnergy: 30,
		RatifierEnergy:   10,
		Timestamp:        standardBallotTime,
	}
}

func VerifyReplicatedBallotProof(t *testing.T) {
	const level = int64(13)
	ev, err := NewEmulateReplicatedBallotProof(level, time.Now(), "REDACTED")
	require.NoError(t, err)
	assert.Equal(t, ev.Digest(), comethash.Sum(ev.Octets()))
	assert.NotNil(t, ev.String())
	assert.Equal(t, ev.Level(), level)
}

func VerifyReplicatedBallotProofVerification(t *testing.T) {
	val := NewEmulatePV()
	ledgerUID := createLedgerUID(comethash.Sum([]byte("REDACTED")), math.MaxInt32, comethash.Sum([]byte("REDACTED")))
	ledgerUidtwo := createLedgerUID(comethash.Sum([]byte("REDACTED")), math.MaxInt32, comethash.Sum([]byte("REDACTED")))
	const ledgerUID = "REDACTED"

	verifyScenarios := []struct {
		verifyLabel         string
		distortProof func(*ReplicatedBallotProof)
		anticipateErr        bool
	}{
		{"REDACTED", func(ev *ReplicatedBallotProof) {}, false},
		{"REDACTED", func(ev *ReplicatedBallotProof) { ev.BallotA = nil }, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) { ev.BallotBYTE = nil }, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) {
			ev.BallotA = nil
			ev.BallotBYTE = nil
		}, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) {
			ev.BallotA = CreateBallotNoFault(t, val, ledgerUID, math.MaxInt32, math.MaxInt64, math.MaxInt32, 0, ledgerUidtwo, standardBallotTime)
		}, true},
		{"REDACTED", func(ev *ReplicatedBallotProof) {
			exchange := ev.BallotA.Clone()
			ev.BallotA = ev.BallotBYTE.Clone()
			ev.BallotBYTE = exchange
		}, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			vote1 := CreateBallotNoFault(t, val, ledgerUID, math.MaxInt32, math.MaxInt64, math.MaxInt32, 0x02, ledgerUID, standardBallotTime)
			ballot2 := CreateBallotNoFault(t, val, ledgerUID, math.MaxInt32, math.MaxInt64, math.MaxInt32, 0x02, ledgerUidtwo, standardBallotTime)
			valueCollection := NewRatifierCollection([]*Ratifier{val.RetrieveTowardRatifier(10)})
			ev, err := NewReplicatedBallotProof(vote1, ballot2, standardBallotTime, valueCollection)
			require.NoError(t, err)
			tc.distortProof(ev)
			assert.Equal(t, tc.anticipateErr, ev.CertifySimple() != nil, "REDACTED")
		})
	}
}

func VerifyRapidCustomerAssaultProofSimple(t *testing.T) {
	level := int64(5)
	sharedLevel := level - 1
	nRatifiers := 10
	ballotCollection, valueCollection, privateValues := randomBallotCollection(level, 1, engineproto.PreendorseKind, nRatifiers, 1, false)
	heading := createHeadingArbitrary()
	heading.Level = level
	ledgerUID := createLedgerUID(comethash.Sum([]byte("REDACTED")), math.MaxInt32, comethash.Sum([]byte("REDACTED")))
	extensionEndorse, err := CreateExtensionEndorse(ledgerUID, level, 1, ballotCollection, privateValues, standardBallotTime, false)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	rce := &RapidCustomerAssaultProof{
		ClashingLedger: &RapidLedger{
			AttestedHeading: &AttestedHeading{
				Heading: heading,
				Endorse: endorse,
			},
			RatifierAssign: valueCollection,
		},
		SharedLevel:        sharedLevel,
		SumPollingEnergy:    valueCollection.SumPollingEnergy(),
		Timestamp:           heading.Time,
		FaultyRatifiers: valueCollection.Ratifiers[:nRatifiers/2],
	}
	assert.NotNil(t, rce.String())
	assert.NotNil(t, rce.Digest())
	assert.Equal(t, rce.Level(), sharedLevel) //
	assert.NotNil(t, rce.Octets())

	//
	verifyScenarios := []struct {
		verifyLabel         string
		distortProof func(*RapidCustomerAssaultProof)
	}{
		{"REDACTED", func(ev *RapidCustomerAssaultProof) { ev.ClashingLedger.Heading = createHeadingArbitrary() }},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) {
			ev.SharedLevel = level + 1
		}},
	}

	for _, tc := range verifyScenarios {
		rce := &RapidCustomerAssaultProof{
			ClashingLedger: &RapidLedger{
				AttestedHeading: &AttestedHeading{
					Heading: heading,
					Endorse: endorse,
				},
				RatifierAssign: valueCollection,
			},
			SharedLevel:        sharedLevel,
			SumPollingEnergy:    valueCollection.SumPollingEnergy(),
			Timestamp:           heading.Time,
			FaultyRatifiers: valueCollection.Ratifiers[:nRatifiers/2],
		}
		digest := rce.Digest()
		tc.distortProof(rce)
		assert.NotEqual(t, digest, rce.Digest(), tc.verifyLabel)
	}
}

func VerifyRapidCustomerAssaultProofVerification(t *testing.T) {
	level := int64(5)
	sharedLevel := level - 1
	nRatifiers := 10
	ballotCollection, valueCollection, privateValues := randomBallotCollection(level, 1, engineproto.PreendorseKind, nRatifiers, 1, false)
	heading := createHeadingArbitrary()
	heading.Level = level
	heading.RatifiersDigest = valueCollection.Digest()
	ledgerUID := createLedgerUID(heading.Digest(), math.MaxInt32, comethash.Sum([]byte("REDACTED")))
	extensionEndorse, err := CreateExtensionEndorse(ledgerUID, level, 1, ballotCollection, privateValues, time.Now(), false)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	rce := &RapidCustomerAssaultProof{
		ClashingLedger: &RapidLedger{
			AttestedHeading: &AttestedHeading{
				Heading: heading,
				Endorse: endorse,
			},
			RatifierAssign: valueCollection,
		},
		SharedLevel:        sharedLevel,
		SumPollingEnergy:    valueCollection.SumPollingEnergy(),
		Timestamp:           heading.Time,
		FaultyRatifiers: valueCollection.Ratifiers[:nRatifiers/2],
	}
	assert.NoError(t, rce.CertifySimple())

	verifyScenarios := []struct {
		verifyLabel         string
		distortProof func(*RapidCustomerAssaultProof)
		anticipateErr        bool
	}{
		{"REDACTED", func(ev *RapidCustomerAssaultProof) {}, false},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) { ev.SharedLevel = -10 }, true},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) {
			ev.SharedLevel = level + 1
		}, true},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) {
			ev.SharedLevel = level
		}, false},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) { ev.ClashingLedger.Heading = nil }, true},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) { ev.ClashingLedger = nil }, true},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) {
			ev.ClashingLedger.RatifierAssign = &RatifierAssign{}
		}, true},
		{"REDACTED", func(ev *RapidCustomerAssaultProof) {
			ev.SumPollingEnergy = -1
		}, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			rce := &RapidCustomerAssaultProof{
				ClashingLedger: &RapidLedger{
					AttestedHeading: &AttestedHeading{
						Heading: heading,
						Endorse: endorse,
					},
					RatifierAssign: valueCollection,
				},
				SharedLevel:        sharedLevel,
				SumPollingEnergy:    valueCollection.SumPollingEnergy(),
				Timestamp:           heading.Time,
				FaultyRatifiers: valueCollection.Ratifiers[:nRatifiers/2],
			}
			tc.distortProof(rce)
			if tc.anticipateErr {
				assert.Error(t, rce.CertifySimple(), tc.verifyLabel)
			} else {
				assert.NoError(t, rce.CertifySimple(), tc.verifyLabel)
			}
		})
	}
}

func VerifyEmulateProofCertifySimple(t *testing.T) {
	validProof, err := NewEmulateReplicatedBallotProof(int64(1), time.Now(), "REDACTED")
	require.NoError(t, err)
	assert.Nil(t, validProof.CertifySimple())
}

func createHeadingArbitrary() *Heading {
	return &Heading{
		Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1},
		LedgerUID:            engineseed.Str(12),
		Level:             int64(engineseed.Uint16()) + 1,
		Time:               time.Now(),
		FinalLedgerUID:        createLedgerUIDArbitrary(),
		FinalEndorseDigest:     vault.CRandomOctets(comethash.Volume),
		DataDigest:           vault.CRandomOctets(comethash.Volume),
		RatifiersDigest:     vault.CRandomOctets(comethash.Volume),
		FollowingRatifiersDigest: vault.CRandomOctets(comethash.Volume),
		AgreementDigest:      vault.CRandomOctets(comethash.Volume),
		ApplicationDigest:            vault.CRandomOctets(comethash.Volume),
		FinalOutcomesDigest:    vault.CRandomOctets(comethash.Volume),
		ProofDigest:       vault.CRandomOctets(comethash.Volume),
		RecommenderLocation:    vault.CRandomOctets(vault.LocationVolume),
	}
}

func VerifyProofSchema(t *testing.T) {
	//
	val := NewEmulatePV()
	ledgerUID := createLedgerUID(comethash.Sum([]byte("REDACTED")), math.MaxInt32, comethash.Sum([]byte("REDACTED")))
	ledgerUidtwo := createLedgerUID(comethash.Sum([]byte("REDACTED")), math.MaxInt32, comethash.Sum([]byte("REDACTED")))
	const ledgerUID = "REDACTED"
	v := CreateBallotNoFault(t, val, ledgerUID, math.MaxInt32, math.MaxInt64, 1, 0x01, ledgerUID, standardBallotTime)
	v2 := CreateBallotNoFault(t, val, ledgerUID, math.MaxInt32, math.MaxInt64, 2, 0x01, ledgerUidtwo, standardBallotTime)

	//
	const level int64 = 37

	var (
		header1 = createHeadingArbitrary()
		header2 = createHeadingArbitrary()
	)

	header1.Level = level
	header1.FinalLedgerUID = ledgerUID
	header1.LedgerUID = ledgerUID

	header2.Level = level
	header2.FinalLedgerUID = ledgerUID
	header2.LedgerUID = ledgerUID

	verifies := []struct {
		verifyLabel     string
		proof     Proof
		toSchemaErr   bool
		fromSchemaErr bool
	}{
		{"REDACTED", nil, true, true},
		{"REDACTED", &ReplicatedBallotProof{}, false, true},
		{"REDACTED", &ReplicatedBallotProof{BallotA: v, BallotBYTE: nil}, false, true},
		{"REDACTED", &ReplicatedBallotProof{BallotA: nil, BallotBYTE: v}, false, true},
		{"REDACTED", &ReplicatedBallotProof{BallotA: v2, BallotBYTE: v}, false, false},
	}
	for _, tt := range verifies {
		t.Run(tt.verifyLabel, func(t *testing.T) {
			pb, err := ProofToSchema(tt.proof)
			if tt.toSchemaErr {
				assert.Error(t, err, tt.verifyLabel)
				return
			}
			assert.NoError(t, err, tt.verifyLabel)

			evi, err := ProofFromSchema(pb)
			if tt.fromSchemaErr {
				assert.Error(t, err, tt.verifyLabel)
				return
			}
			require.Equal(t, tt.proof, evi, tt.verifyLabel)
		})
	}
}
