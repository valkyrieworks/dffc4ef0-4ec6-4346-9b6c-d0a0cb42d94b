package kinds

import (
	//
	//

	"crypto/rand"
	"encoding/hex"
	"math"
	"os"
	"reflect"
	"testing"
	"time"

	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

func VerifyPrimary(m *testing.M) {
	cipher := m.Run()
	os.Exit(cipher)
}

func VerifyLedgerAppendProof(t *testing.T) {
	txs := []Tx{Tx("REDACTED"), Tx("REDACTED")}
	finalUUID := createLedgerUUIDUnpredictable()
	h := int64(3)

	ballotAssign, _, values := arbitraryBallotAssign(h-1, 1, commitchema.PreendorseKind, 10, 1, false)
	addnEndorse, err := CreateAddnEndorse(finalUUID, h-1, 1, ballotAssign, values, time.Now(), false)
	require.NoError(t, err)

	ev, err := FreshSimulateReplicatedBallotProofUsingAssessor(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	occurenceCatalog := []Proof{ev}

	ledger := CreateLedger(h, txs, addnEndorse.TowardEndorse(), occurenceCatalog)
	require.NotNil(t, ledger)
	require.Equal(t, 1, len(ledger.Proof.Proof))
	require.NotNil(t, ledger.ProofDigest)
}

func VerifyLedgerCertifyFundamental(t *testing.T) {
	require.Error(t, (*Ledger)(nil).CertifyFundamental())

	txs := []Tx{Tx("REDACTED"), Tx("REDACTED")}
	finalUUID := createLedgerUUIDUnpredictable()
	h := int64(3)

	ballotAssign, itemAssign, values := arbitraryBallotAssign(h-1, 1, commitchema.PreendorseKind, 10, 1, false)
	addnEndorse, err := CreateAddnEndorse(finalUUID, h-1, 1, ballotAssign, values, time.Now(), false)
	require.NoError(t, err)
	endorse := addnEndorse.TowardEndorse()

	ev, err := FreshSimulateReplicatedBallotProofUsingAssessor(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	occurenceCatalog := []Proof{ev}

	verifyScenarios := []struct {
		verifyAlias      string
		distortLedger func(*Ledger)
		expirationFault        bool
	}{
		{"REDACTED", func(blk *Ledger) {}, false},
		{"REDACTED", func(blk *Ledger) { blk.NominatorLocation = itemAssign.ObtainNominator().Location }, false},
		{"REDACTED", func(blk *Ledger) { blk.Altitude = -1 }, true},
		{"REDACTED", func(blk *Ledger) {
			blk.FinalEndorse.Notations = endorse.Notations[:endorse.Extent()/2]
			blk.FinalEndorse.digest = nil //
		}, true},
		{"REDACTED", func(blk *Ledger) { blk.FinalEndorseDigest = []byte("REDACTED") }, true},
		{"REDACTED", func(blk *Ledger) {
			blk.Txs[0] = Tx("REDACTED")
			blk.digest = nil //
		}, true},
		{"REDACTED", func(blk *Ledger) {
			blk.DataDigest = commitrand.Octets(len(blk.DataDigest))
		}, true},
		{"REDACTED", func(blk *Ledger) {
			blk.ProofDigest = []byte("REDACTED")
		}, true},
		{"REDACTED", func(blk *Ledger) {
			blk.Edition.Ledger = 1
		}, true},
	}
	for i, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			ledger := CreateLedger(h, txs, endorse, occurenceCatalog)
			ledger.NominatorLocation = itemAssign.ObtainNominator().Location
			tc.distortLedger(ledger)
			err = ledger.CertifyFundamental()
			assert.Equal(t, tc.expirationFault, err != nil, "REDACTED", i, err)
		})
	}
}

func VerifyLedgerDigest(t *testing.T) {
	assert.Nil(t, (*Ledger)(nil).Digest())
	assert.Nil(t, CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil).Digest())
}

func VerifyLedgerCreateFragmentAssign(t *testing.T) {
	bps, err := (*Ledger)(nil).CreateFragmentAssign(2)
	assert.Error(t, err)
	assert.Nil(t, bps)

	fragmentAssign, err := CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil).CreateFragmentAssign(1024)
	require.NoError(t, err)
	assert.NotNil(t, fragmentAssign)
	assert.EqualValues(t, 1, fragmentAssign.Sum())
}

func VerifyLedgerCreateFragmentAssignUsingProof(t *testing.T) {
	bps, err := (*Ledger)(nil).CreateFragmentAssign(2)
	assert.Error(t, err)
	assert.Nil(t, bps)

	finalUUID := createLedgerUUIDUnpredictable()
	h := int64(3)

	ballotAssign, _, values := arbitraryBallotAssign(h-1, 1, commitchema.PreendorseKind, 10, 1, false)
	addnEndorse, err := CreateAddnEndorse(finalUUID, h-1, 1, ballotAssign, values, time.Now(), false)
	require.NoError(t, err)

	ev, err := FreshSimulateReplicatedBallotProofUsingAssessor(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	occurenceCatalog := []Proof{ev}

	fragmentAssign, err := CreateLedger(h, []Tx{Tx("REDACTED")}, addnEndorse.TowardEndorse(), occurenceCatalog).CreateFragmentAssign(512)
	require.NoError(t, err)

	assert.NotNil(t, fragmentAssign)
	assert.EqualValues(t, 4, fragmentAssign.Sum())
}

func VerifyLedgerDigestsToward(t *testing.T) {
	assert.False(t, (*Ledger)(nil).DigestsToward(nil))

	finalUUID := createLedgerUUIDUnpredictable()
	h := int64(3)
	ballotAssign, itemAssign, values := arbitraryBallotAssign(h-1, 1, commitchema.PreendorseKind, 10, 1, false)
	addnEndorse, err := CreateAddnEndorse(finalUUID, h-1, 1, ballotAssign, values, time.Now(), false)
	require.NoError(t, err)

	ev, err := FreshSimulateReplicatedBallotProofUsingAssessor(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	occurenceCatalog := []Proof{ev}

	ledger := CreateLedger(h, []Tx{Tx("REDACTED")}, addnEndorse.TowardEndorse(), occurenceCatalog)
	ledger.AssessorsDigest = itemAssign.Digest()
	assert.False(t, ledger.DigestsToward([]byte{}))
	assert.False(t, ledger.DigestsToward([]byte("REDACTED")))
	assert.True(t, ledger.DigestsToward(ledger.Digest()))
}

func VerifyLedgerExtent(t *testing.T) {
	extent := CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil).Extent()
	if extent <= 0 {
		t.Fatal("REDACTED")
	}
}

func VerifyLedgerText(t *testing.T) {
	assert.Equal(t, "REDACTED", (*Ledger)(nil).Text())
	assert.Equal(t, "REDACTED", (*Ledger)(nil).TextFormatted("REDACTED"))
	assert.Equal(t, "REDACTED", (*Ledger)(nil).TextBrief())

	ledger := CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil)
	assert.NotEqual(t, "REDACTED", ledger.Text())
	assert.NotEqual(t, "REDACTED", ledger.TextFormatted("REDACTED"))
	assert.NotEqual(t, "REDACTED", ledger.TextBrief())
}

func createLedgerUUIDUnpredictable() LedgerUUID {
	var (
		ledgerDigest   = make([]byte, tenderminthash.Extent)
		fragmentAssignDigest = make([]byte, tenderminthash.Extent)
	)
	rand.Read(ledgerDigest)   //
	rand.Read(fragmentAssignDigest) //
	return LedgerUUID{ledgerDigest, FragmentAssignHeading{123, fragmentAssignDigest}}
}

func createLedgerUUID(digest []byte, fragmentAssignExtent uint32, fragmentAssignDigest []byte) LedgerUUID {
	var (
		h   = make([]byte, tenderminthash.Extent)
		psH = make([]byte, tenderminthash.Extent)
	)
	copy(h, digest)
	copy(psH, fragmentAssignDigest)
	return LedgerUUID{
		Digest: h,
		FragmentAssignHeading: FragmentAssignHeading{
			Sum: fragmentAssignExtent,
			Digest:  psH,
		},
	}
}

var voidOctets []byte

//
var blankOctets = []byte{
	0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8,
	0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b,
	0x78, 0x52, 0xb8, 0x55,
}

func VerifyVoidHeadingDigestNotCollapse(t *testing.T) {
	assert.Equal(t, voidOctets, []byte((*Heading)(nil).Digest()))
	assert.Equal(t, voidOctets, []byte((new(Heading)).Digest()))
}

func VerifyVoidDataDigestNotCollapse(t *testing.T) {
	assert.Equal(t, blankOctets, []byte((*Data)(nil).Digest()))
	assert.Equal(t, blankOctets, []byte(new(Data).Digest()))
}

func VerifyEndorse(t *testing.T) {
	finalUUID := createLedgerUUIDUnpredictable()
	h := int64(3)
	ballotAssign, _, values := arbitraryBallotAssign(h-1, 1, commitchema.PreendorseKind, 10, 1, true)
	addnEndorse, err := CreateAddnEndorse(finalUUID, h-1, 1, ballotAssign, values, time.Now(), true)
	require.NoError(t, err)

	assert.Equal(t, h-1, addnEndorse.Altitude)
	assert.EqualValues(t, 1, addnEndorse.Iteration)
	assert.Equal(t, commitchema.PreendorseKind, commitchema.AttestedSignalKind(addnEndorse.Kind()))
	if addnEndorse.Extent() <= 0 {
		t.Fatalf("REDACTED", addnEndorse, addnEndorse.Extent())
	}

	require.NotNil(t, addnEndorse.DigitSeries())
	assert.Equal(t, digits.FreshDigitCollection(10).Extent(), addnEndorse.DigitSeries().Extent())

	assert.Equal(t, ballotAssign.ObtainViaOrdinal(0), addnEndorse.ObtainViaOrdinal(0))
	assert.True(t, addnEndorse.EqualsEndorse())
}

func VerifyEndorseCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		verifyAlias       string
		distortEndorse func(*Endorse)
		anticipateFault      bool
	}{
		{"REDACTED", func(com *Endorse) {}, false},
		{"REDACTED", func(com *Endorse) { com.Notations[0].Notation = []byte{0} }, false},
		{"REDACTED", func(com *Endorse) { com.Altitude = int64(-100) }, true},
		{"REDACTED", func(com *Endorse) { com.Iteration = -100 }, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			com := arbitraryEndorse(time.Now())
			tc.distortEndorse(com)
			assert.Equal(t, tc.anticipateFault, com.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func VerifyMaximumEndorseOctets(t *testing.T) {
	//
	//
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)

	cs := EndorseSignature{
		LedgerUUIDMarker:      LedgerUUIDMarkerVoid,
		AssessorLocation: security.LocatorDigest([]byte("REDACTED")),
		Timestamp:        timestamp,
		Notation:        security.CHARArbitraryOctets(MaximumSigningExtent),
	}

	bufferSignature := cs.TowardSchema()
	//
	assert.EqualValues(t, MaximumEndorseSignatureOctets, bufferSignature.Extent())

	//
	endorse := &Endorse{
		Altitude: math.MaxInt64,
		Iteration:  math.MaxInt32,
		LedgerUUID: LedgerUUID{
			Digest: tenderminthash.Sum([]byte("REDACTED")),
			FragmentAssignHeading: FragmentAssignHeading{
				Sum: math.MaxInt32,
				Digest:  tenderminthash.Sum([]byte("REDACTED")),
			},
		},
		Notations: []EndorseSignature{cs},
	}

	pb := endorse.TowardSchema()

	assert.EqualValues(t, MaximumEndorseOctets(1), int64(pb.Extent()))

	//
	for i := 1; i < MaximumBallotsTally; i++ {
		endorse.Notations = append(endorse.Notations, cs)
	}

	pb = endorse.TowardSchema()

	assert.EqualValues(t, MaximumEndorseOctets(MaximumBallotsTally), int64(pb.Extent()))
}

func VerifyHeadingDigest(t *testing.T) {
	verifyScenarios := []struct {
		description       string
		heading     *Heading
		anticipateDigest octets.HexadecimalOctets
	}{
		{"REDACTED", &Heading{
			Edition:            strongmindedition.Agreement{Ledger: 1, App: 2},
			SuccessionUUID:            "REDACTED",
			Altitude:             3,
			Moment:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			FinalLedgerUUID:        createLedgerUUID(make([]byte, tenderminthash.Extent), 6, make([]byte, tenderminthash.Extent)),
			FinalEndorseDigest:     tenderminthash.Sum([]byte("REDACTED")),
			DataDigest:           tenderminthash.Sum([]byte("REDACTED")),
			AssessorsDigest:     tenderminthash.Sum([]byte("REDACTED")),
			FollowingAssessorsDigest: tenderminthash.Sum([]byte("REDACTED")),
			AgreementDigest:      tenderminthash.Sum([]byte("REDACTED")),
			PlatformDigest:            tenderminthash.Sum([]byte("REDACTED")),
			FinalOutcomesDigest:    tenderminthash.Sum([]byte("REDACTED")),
			ProofDigest:       tenderminthash.Sum([]byte("REDACTED")),
			NominatorLocation:    security.LocatorDigest([]byte("REDACTED")),
		}, hexadecimalOctetsOriginatingText("REDACTED")},
		{"REDACTED", nil, nil},
		{"REDACTED", &Heading{
			Edition:            strongmindedition.Agreement{Ledger: 1, App: 2},
			SuccessionUUID:            "REDACTED",
			Altitude:             3,
			Moment:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			FinalLedgerUUID:        createLedgerUUID(make([]byte, tenderminthash.Extent), 6, make([]byte, tenderminthash.Extent)),
			FinalEndorseDigest:     tenderminthash.Sum([]byte("REDACTED")),
			DataDigest:           tenderminthash.Sum([]byte("REDACTED")),
			AssessorsDigest:     nil,
			FollowingAssessorsDigest: tenderminthash.Sum([]byte("REDACTED")),
			AgreementDigest:      tenderminthash.Sum([]byte("REDACTED")),
			PlatformDigest:            tenderminthash.Sum([]byte("REDACTED")),
			FinalOutcomesDigest:    tenderminthash.Sum([]byte("REDACTED")),
			ProofDigest:       tenderminthash.Sum([]byte("REDACTED")),
			NominatorLocation:    security.LocatorDigest([]byte("REDACTED")),
		}, nil},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.anticipateDigest, tc.heading.Digest())

			//
			//
			if tc.heading != nil && tc.anticipateDigest != nil {
				octetSegments := [][]byte{}

				s := reflect.ValueOf(*tc.heading)
				for i := 0; i < s.NumField(); i++ {
					f := s.Field(i)

					assert.False(t, f.IsZero(), "REDACTED",
						s.Type().Field(i).Name)

					switch f := f.Interface().(type) {
					case int64, octets.HexadecimalOctets, string:
						octetSegments = append(octetSegments, codecSerialize(f))
					case time.Time:
						bz, err := gogotypes.StdTimeMarshal(f)
						require.NoError(t, err)
						octetSegments = append(octetSegments, bz)
					case strongmindedition.Agreement:
						bz, err := f.Serialize()
						require.NoError(t, err)
						octetSegments = append(octetSegments, bz)
					case LedgerUUID:
						bufferbi := f.TowardSchema()
						bz, err := bufferbi.Serialize()
						require.NoError(t, err)
						octetSegments = append(octetSegments, bz)
					default:
						t.Errorf("REDACTED", f)
					}
				}
				assert.Equal(t,
					octets.HexadecimalOctets(hashmap.DigestOriginatingOctetSegments(octetSegments)), tc.heading.Digest())
			}
		})
	}
}

func VerifyMaximumHeadingOctets(t *testing.T) {
	//
	//
	//
	//
	maximumSuccessionUUID := "REDACTED"
	for i := 0; i < MaximumSuccessionUUIDSize; i++ {
		maximumSuccessionUUID += "REDACTED"
	}

	//
	//
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)

	h := Heading{
		Edition:            strongmindedition.Agreement{Ledger: math.MaxInt64, App: math.MaxInt64},
		SuccessionUUID:            maximumSuccessionUUID,
		Altitude:             math.MaxInt64,
		Moment:               timestamp,
		FinalLedgerUUID:        createLedgerUUID(make([]byte, tenderminthash.Extent), math.MaxInt32, make([]byte, tenderminthash.Extent)),
		FinalEndorseDigest:     tenderminthash.Sum([]byte("REDACTED")),
		DataDigest:           tenderminthash.Sum([]byte("REDACTED")),
		AssessorsDigest:     tenderminthash.Sum([]byte("REDACTED")),
		FollowingAssessorsDigest: tenderminthash.Sum([]byte("REDACTED")),
		AgreementDigest:      tenderminthash.Sum([]byte("REDACTED")),
		PlatformDigest:            tenderminthash.Sum([]byte("REDACTED")),
		FinalOutcomesDigest:    tenderminthash.Sum([]byte("REDACTED")),
		ProofDigest:       tenderminthash.Sum([]byte("REDACTED")),
		NominatorLocation:    security.LocatorDigest([]byte("REDACTED")),
	}

	bz, err := h.TowardSchema().Serialize()
	require.NoError(t, err)

	assert.EqualValues(t, MaximumHeadingOctets, int64(len(bz)))
}

func arbitraryEndorse(now time.Time) *Endorse {
	finalUUID := createLedgerUUIDUnpredictable()
	h := int64(3)
	ballotAssign, _, values := arbitraryBallotAssign(h-1, 1, commitchema.PreendorseKind, 10, 1, false)
	addnEndorse, err := CreateAddnEndorse(finalUUID, h-1, 1, ballotAssign, values, now, false)
	if err != nil {
		panic(err)
	}
	return addnEndorse.TowardEndorse()
}

func hexadecimalOctetsOriginatingText(s string) octets.HexadecimalOctets {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return octets.HexadecimalOctets(b)
}

func VerifyLedgerMaximumDataOctets(t *testing.T) {
	verifyScenarios := []struct {
		maximumOctets      int64
		valuesTally     int
		proofOctets int64
		alarms        bool
		outcome        int64
	}{
		0: {-10, 1, 0, true, 0},
		1: {10, 1, 0, true, 0},
		2: {841, 1, 0, true, 0},
		3: {875, 1, 0, false, 0},
		4: {876, 1, 0, false, 1},
		5: {1019, 2, 0, false, 0},
	}

	for i, tc := range verifyScenarios {
		if tc.alarms {
			assert.Panics(t, func() {
				MaximumDataOctets(tc.maximumOctets, tc.proofOctets, tc.valuesTally)
			}, "REDACTED", i)
		} else {
			assert.Equal(t,
				tc.outcome,
				MaximumDataOctets(tc.maximumOctets, tc.proofOctets, tc.valuesTally),
				"REDACTED", i)
		}
	}
}

func VerifyLedgerMaximumDataOctetsNegativeProof(t *testing.T) {
	verifyScenarios := []struct {
		maximumOctets  int64
		valuesTally int
		alarms    bool
		outcome    int64
	}{
		0: {-10, 1, true, 0},
		1: {10, 1, true, 0},
		2: {841, 1, true, 0},
		3: {875, 1, false, 0},
		4: {876, 1, false, 1},
	}

	for i, tc := range verifyScenarios {
		if tc.alarms {
			assert.Panics(t, func() {
				MaximumDataOctetsNegativeProof(tc.maximumOctets, tc.valuesTally)
			}, "REDACTED", i)
		} else {
			assert.Equal(t,
				tc.outcome,
				MaximumDataOctetsNegativeProof(tc.maximumOctets, tc.valuesTally),
				"REDACTED", i)
		}
	}
}

//
//
//
//
func VerifyBallotAssignTowardExpandedEndorse(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias             string
		encompassAddition bool
	}{
		{
			alias:             "REDACTED",
			encompassAddition: false,
		},
		{
			alias:             "REDACTED",
			encompassAddition: true,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			ledgerUUID := createLedgerUUIDUnpredictable()

			itemAssign, values := ArbitraryAssessorAssign(10, 1)
			var ballotAssign *BallotAssign
			if verifyInstance.encompassAddition {
				ballotAssign = FreshExpandedBallotAssign("REDACTED", 3, 1, commitchema.PreendorseKind, itemAssign)
			} else {
				ballotAssign = FreshBallotAssign("REDACTED", 3, 1, commitchema.PreendorseKind, itemAssign)
			}
			for i := 0; i < len(values); i++ {
				publicToken, err := values[i].ObtainPublicToken()
				require.NoError(t, err)
				ballot := &Ballot{
					AssessorLocation: publicToken.Location(),
					AssessorOrdinal:   int32(i),
					Altitude:           3,
					Iteration:            1,
					Kind:             commitchema.PreendorseKind,
					LedgerUUID:          ledgerUUID,
					Timestamp:        time.Now(),
				}
				v := ballot.TowardSchema()
				err = values[i].AttestBallot(ballotAssign.SuccessionUUID(), v)
				require.NoError(t, err)
				ballot.Notation = v.Notation
				if verifyInstance.encompassAddition {
					ballot.AdditionNotation = v.AdditionNotation
				}
				appended, err := ballotAssign.AppendBallot(ballot)
				require.NoError(t, err)
				require.True(t, appended)
			}
			var verAltitude int64
			if verifyInstance.encompassAddition {
				verAltitude = 1
			}
			ec := ballotAssign.CreateExpandedEndorse(IfaceParameters{BallotAdditionsActivateAltitude: verAltitude})

			for i := int32(0); int(i) < len(values); i++ {
				ballot1 := ballotAssign.ObtainViaOrdinal(i)
				ballot2 := ec.ObtainExpandedBallot(i)

				ballot1bz, err := ballot1.TowardSchema().Serialize()
				require.NoError(t, err)
				ballot2bz, err := ballot2.TowardSchema().Serialize()
				require.NoError(t, err)
				assert.Equal(t, ballot1bz, ballot2bz)
			}
		})
	}
}

//
//
//
func towardBallotAssign(ec *ExpandedEndorse, successionUUID string, values *AssessorAssign) *BallotAssign {
	ballotAssign := FreshBallotAssign(successionUUID, ec.Altitude, ec.Iteration, commitchema.PreendorseKind, values)
	ec.appendSignaturesTowardBallotAssign(ballotAssign)
	return ballotAssign
}

//
//
//
//
func VerifyExpandedEndorseTowardBallotAssign(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias             string
		encompassAddition bool
	}{
		{
			alias:             "REDACTED",
			encompassAddition: false,
		},
		{
			alias:             "REDACTED",
			encompassAddition: true,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			finalUUID := createLedgerUUIDUnpredictable()
			h := int64(3)

			ballotAssign, itemAssign, values := arbitraryBallotAssign(h-1, 1, commitchema.PreendorseKind, 10, 1, true)
			addnEndorse, err := CreateAddnEndorse(finalUUID, h-1, 1, ballotAssign, values, time.Now(), true)
			assert.NoError(t, err)

			if !verifyInstance.encompassAddition {
				for i := 0; i < len(values); i++ {
					v := ballotAssign.ObtainViaOrdinal(int32(i))
					v.Addition = nil
					v.AdditionNotation = nil
					addnEndorse.ExpandedNotations[i].Addition = nil
					addnEndorse.ExpandedNotations[i].AdditionNotation = nil
				}
			}

			successionUUID := ballotAssign.SuccessionUUID()
			var ballotGroup2 *BallotAssign
			if verifyInstance.encompassAddition {
				ballotGroup2 = addnEndorse.TowardExpandedBallotAssign(successionUUID, itemAssign)
			} else {
				ballotGroup2 = towardBallotAssign(addnEndorse, successionUUID, itemAssign)
			}

			for i := int32(0); int(i) < len(values); i++ {
				ballot1 := ballotAssign.ObtainViaOrdinal(i)
				ballot2 := ballotGroup2.ObtainViaOrdinal(i)
				ballot3 := addnEndorse.ObtainExpandedBallot(i)

				ballot1bz, err := ballot1.TowardSchema().Serialize()
				require.NoError(t, err)
				ballot2bz, err := ballot2.TowardSchema().Serialize()
				require.NoError(t, err)
				ballot3bz, err := ballot3.TowardSchema().Serialize()
				require.NoError(t, err)
				assert.Equal(t, ballot1bz, ballot2bz)
				assert.Equal(t, ballot1bz, ballot3bz)
			}
		})
	}
}

func VerifyEndorseTowardBallotAssignUsingBallotsForeachVoidLedger(t *testing.T) {
	ledgerUUID := createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED"))

	const (
		altitude = int64(3)
		iteration  = 0
	)

	type endorseBallotVerify struct {
		ledgerIDXDstore      []LedgerUUID
		countBallots      []int //
		countAssessors int
		sound         bool
	}

	verifyScenarios := []endorseBallotVerify{
		{[]LedgerUUID{ledgerUUID, {}}, []int{67, 33}, 100, true},
	}

	for _, tc := range verifyScenarios {
		ballotAssign, itemAssign, values := arbitraryBallotAssign(altitude-1, iteration, commitchema.PreendorseKind, tc.countAssessors, 1, false)

		vi := int32(0)
		for n := range tc.ledgerIDXDstore {
			for i := 0; i < tc.countBallots[n]; i++ {
				publicToken, err := values[vi].ObtainPublicToken()
				require.NoError(t, err)
				ballot := &Ballot{
					AssessorLocation: publicToken.Location(),
					AssessorOrdinal:   vi,
					Altitude:           altitude - 1,
					Iteration:            iteration,
					Kind:             commitchema.PreendorseKind,
					LedgerUUID:          tc.ledgerIDXDstore[n],
					Timestamp:        committime.Now(),
				}

				appended, err := attestAppendBallot(values[vi], ballot, ballotAssign)
				assert.NoError(t, err)
				assert.True(t, appended)

				vi++
			}
		}

		verAltitudeArgument := IfaceParameters{BallotAdditionsActivateAltitude: 0}
		if tc.sound {
			addnEndorse := ballotAssign.CreateExpandedEndorse(verAltitudeArgument) //
			assert.NotNil(t, addnEndorse)
			err := itemAssign.ValidateEndorse(ballotAssign.SuccessionUUID(), ledgerUUID, altitude-1, addnEndorse.TowardEndorse())
			assert.Nil(t, err)
		} else {
			assert.Panics(t, func() { ballotAssign.CreateExpandedEndorse(verAltitudeArgument) })
		}
	}
}

func VerifyLedgerUUIDCertifyFundamental(t *testing.T) {
	soundLedgerUUID := LedgerUUID{
		Digest: octets.HexadecimalOctets{},
		FragmentAssignHeading: FragmentAssignHeading{
			Sum: 1,
			Digest:  octets.HexadecimalOctets{},
		},
	}

	unfitLedgerUUID := LedgerUUID{
		Digest: []byte{0},
		FragmentAssignHeading: FragmentAssignHeading{
			Sum: 1,
			Digest:  []byte{0},
		},
	}

	verifyScenarios := []struct {
		verifyAlias             string
		ledgerUUIDDigest          octets.HexadecimalOctets
		ledgerUUIDFragmentAssignHeading FragmentAssignHeading
		anticipateFault            bool
	}{
		{"REDACTED", soundLedgerUUID.Digest, soundLedgerUUID.FragmentAssignHeading, false},
		{"REDACTED", unfitLedgerUUID.Digest, soundLedgerUUID.FragmentAssignHeading, true},
		{"REDACTED", soundLedgerUUID.Digest, unfitLedgerUUID.FragmentAssignHeading, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			ledgerUUID := LedgerUUID{
				Digest:          tc.ledgerUUIDDigest,
				FragmentAssignHeading: tc.ledgerUUIDFragmentAssignHeading,
			}
			assert.Equal(t, tc.anticipateFault, ledgerUUID.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func VerifyLedgerSchemaBuffer(t *testing.T) {
	h := commitrand.Int63n()
	c1 := arbitraryEndorse(time.Now())
	b1 := CreateLedger(h, []Tx{Tx([]byte{1})}, &Endorse{Notations: []EndorseSignature{}}, []Proof{})
	b1.NominatorLocation = commitrand.Octets(security.LocatorExtent)

	b2 := CreateLedger(h, []Tx{Tx([]byte{1})}, c1, []Proof{})
	b2.NominatorLocation = commitrand.Octets(security.LocatorExtent)
	proofMoment := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	evi, err := FreshSimulateReplicatedBallotProof(h, proofMoment, "REDACTED")
	require.NoError(t, err)
	b2.Proof = ProofData{Proof: ProofCatalog{evi}}
	b2.ProofDigest = b2.Proof.Digest()

	b3 := CreateLedger(h, []Tx{}, c1, []Proof{})
	b3.NominatorLocation = commitrand.Octets(security.LocatorExtent)
	verifyScenarios := []struct {
		msg      string
		b1       *Ledger
		expirationPhrase  bool
		expirationPhase2 bool
	}{
		{"REDACTED", nil, false, false},
		{"REDACTED", b1, true, true},
		{"REDACTED", b2, true, true},
		{"REDACTED", b3, true, true},
	}
	for _, tc := range verifyScenarios {
		pb, err := tc.b1.TowardSchema()
		if tc.expirationPhrase {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		ledger, err := LedgerOriginatingSchema(pb)
		if tc.expirationPhase2 {
			require.NoError(t, err, tc.msg)
			require.EqualValues(t, tc.b1.Heading, ledger.Heading, tc.msg)
			require.EqualValues(t, tc.b1.Data, ledger.Data, tc.msg)
			require.EqualValues(t, tc.b1.Proof.Proof, ledger.Proof.Proof, tc.msg)
			require.EqualValues(t, *tc.b1.FinalEndorse, *ledger.FinalEndorse, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func VerifyDataSchemaBuffer(t *testing.T) {
	data := &Data{Txs: Txs{Tx([]byte{1}), Tx([]byte{2}), Tx([]byte{3})}}
	datum2 := &Data{Txs: Txs{}}
	verifyScenarios := []struct {
		msg     string
		item1   *Data
		expirationPhrase bool
	}{
		{"REDACTED", data, true},
		{"REDACTED", datum2, true},
	}
	for _, tc := range verifyScenarios {
		schemaData := tc.item1.TowardSchema()
		d, err := DataOriginatingSchema(&schemaData)
		if tc.expirationPhrase {
			require.NoError(t, err, tc.msg)
			require.EqualValues(t, tc.item1, &d, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

//
func VerifyProofDataSchemaBuffer(t *testing.T) {
	const successionUUID = "REDACTED"
	ev, err := FreshSimulateReplicatedBallotProof(math.MaxInt64, time.Now(), successionUUID)
	require.NoError(t, err)
	data := &ProofData{Proof: ProofCatalog{ev}}
	_ = data.OctetExtent()
	verifyScenarios := []struct {
		msg      string
		item1    *ProofData
		expirationPhase1 bool
		expirationPhase2 bool
	}{
		{"REDACTED", data, true, true},
		{"REDACTED", &ProofData{Proof: ProofCatalog{}}, true, true},
		{"REDACTED", nil, false, false},
	}

	for _, tc := range verifyScenarios {
		schemaData, err := tc.item1.TowardSchema()
		if tc.expirationPhase1 {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		evidenceDelta := new(ProofData)
		err = evidenceDelta.OriginatingSchema(schemaData)
		if tc.expirationPhase2 {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.item1, evidenceDelta, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func createArbitraryHeading() Heading {
	successionUUID := "REDACTED"
	t := time.Now()
	altitude := commitrand.Int63n()
	arbitraryOctets := commitrand.Octets(tenderminthash.Extent)
	arbitraryLocator := commitrand.Octets(security.LocatorExtent)
	h := Heading{
		Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1},
		SuccessionUUID:            successionUUID,
		Altitude:             altitude,
		Moment:               t,
		FinalLedgerUUID:        LedgerUUID{},
		FinalEndorseDigest:     arbitraryOctets,
		DataDigest:           arbitraryOctets,
		AssessorsDigest:     arbitraryOctets,
		FollowingAssessorsDigest: arbitraryOctets,
		AgreementDigest:      arbitraryOctets,
		PlatformDigest:            arbitraryOctets,

		FinalOutcomesDigest: arbitraryOctets,

		ProofDigest:    arbitraryOctets,
		NominatorLocation: arbitraryLocator,
	}

	return h
}

func VerifyHeadingSchema(t *testing.T) {
	h1 := createArbitraryHeading()
	tc := []struct {
		msg     string
		h1      *Heading
		expirationPhrase bool
	}{
		{"REDACTED", &h1, true},
		{"REDACTED", &Heading{}, false},
	}

	for _, tt := range tc {
		t.Run(tt.msg, func(t *testing.T) {
			pb := tt.h1.TowardSchema()
			h, err := HeadingOriginatingSchema(pb)
			if tt.expirationPhrase {
				require.NoError(t, err, tt.msg)
				require.Equal(t, tt.h1, &h, tt.msg)
			} else {
				require.Error(t, err, tt.msg)
			}
		})
	}
}

func VerifyLedgerUUIDSchemaBuffer(t *testing.T) {
	ledgerUUID := createLedgerUUID([]byte("REDACTED"), 2, []byte("REDACTED"))
	verifyScenarios := []struct {
		msg     string
		offer1    *LedgerUUID
		expirationPhrase bool
	}{
		{"REDACTED", &ledgerUUID, true},
		{"REDACTED", &LedgerUUID{}, true},
		{"REDACTED", nil, false},
	}
	for _, tc := range verifyScenarios {
		schemaLedgerUUID := tc.offer1.TowardSchema()

		bi, err := LedgerUUIDOriginatingSchema(&schemaLedgerUUID)
		if tc.expirationPhrase {
			require.NoError(t, err)
			require.Equal(t, tc.offer1, bi, tc.msg)
		} else {
			require.NotEqual(t, tc.offer1, bi, tc.msg)
		}
	}
}

func VerifyNotatedHeadingSchemaBuffer(t *testing.T) {
	endorse := arbitraryEndorse(time.Now())
	h := createArbitraryHeading()

	sh := NotatedHeading{Heading: &h, Endorse: endorse}

	verifyScenarios := []struct {
		msg     string
		sh1     *NotatedHeading
		expirationPhrase bool
	}{
		{"REDACTED", &NotatedHeading{}, true},
		{"REDACTED", &sh, true},
		{"REDACTED", nil, false},
	}
	for _, tc := range verifyScenarios {
		schemaNotatedHeading := tc.sh1.TowardSchema()

		sh, err := NotatedHeadingOriginatingSchema(schemaNotatedHeading)

		if tc.expirationPhrase {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.sh1, sh, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func VerifyLedgerUUIDMatches(t *testing.T) {
	var (
		ledgerUUID          = createLedgerUUID([]byte("REDACTED"), 2, []byte("REDACTED"))
		ledgerUUIDReplicated = createLedgerUUID([]byte("REDACTED"), 2, []byte("REDACTED"))
		ledgerUUIDDistinct = createLedgerUUID([]byte("REDACTED"), 2, []byte("REDACTED"))
		ledgerUUIDBlank     = LedgerUUID{}
	)

	assert.True(t, ledgerUUID.Matches(ledgerUUIDReplicated))
	assert.False(t, ledgerUUID.Matches(ledgerUUIDDistinct))
	assert.False(t, ledgerUUID.Matches(ledgerUUIDBlank))
	assert.True(t, ledgerUUIDBlank.Matches(ledgerUUIDBlank)) //
	assert.False(t, ledgerUUIDBlank.Matches(ledgerUUIDDistinct))
}
