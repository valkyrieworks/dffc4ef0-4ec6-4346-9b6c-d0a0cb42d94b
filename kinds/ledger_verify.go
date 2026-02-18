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

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/utils/bits"
	"github.com/valkyrieworks/utils/octets"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	engineclock "github.com/valkyrieworks/kinds/moment"
	"github.com/valkyrieworks/release"
)

func VerifyMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func VerifyLedgerAppendProof(t *testing.T) {
	txs := []Tx{Tx("REDACTED"), Tx("REDACTED")}
	finalUID := createLedgerUIDArbitrary()
	h := int64(3)

	ballotCollection, _, values := randomBallotCollection(h-1, 1, engineproto.PreendorseKind, 10, 1, false)
	extensionEndorse, err := CreateExtensionEndorse(finalUID, h-1, 1, ballotCollection, values, time.Now(), false)
	require.NoError(t, err)

	ev, err := NewEmulateReplicatedBallotProofWithRatifier(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	evtCatalog := []Proof{ev}

	ledger := CreateLedger(h, txs, extensionEndorse.ToEndorse(), evtCatalog)
	require.NotNil(t, ledger)
	require.Equal(t, 1, len(ledger.Proof.Proof))
	require.NotNil(t, ledger.ProofDigest)
}

func VerifyLedgerCertifySimple(t *testing.T) {
	require.Error(t, (*Ledger)(nil).CertifySimple())

	txs := []Tx{Tx("REDACTED"), Tx("REDACTED")}
	finalUID := createLedgerUIDArbitrary()
	h := int64(3)

	ballotCollection, valueCollection, values := randomBallotCollection(h-1, 1, engineproto.PreendorseKind, 10, 1, false)
	extensionEndorse, err := CreateExtensionEndorse(finalUID, h-1, 1, ballotCollection, values, time.Now(), false)
	require.NoError(t, err)
	endorse := extensionEndorse.ToEndorse()

	ev, err := NewEmulateReplicatedBallotProofWithRatifier(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	evtCatalog := []Proof{ev}

	verifyScenarios := []struct {
		verifyLabel      string
		distortLedger func(*Ledger)
		expirationErr        bool
	}{
		{"REDACTED", func(blk *Ledger) {}, false},
		{"REDACTED", func(blk *Ledger) { blk.RecommenderLocation = valueCollection.FetchRecommender().Location }, false},
		{"REDACTED", func(blk *Ledger) { blk.Level = -1 }, true},
		{"REDACTED", func(blk *Ledger) {
			blk.FinalEndorse.Endorsements = endorse.Endorsements[:endorse.Volume()/2]
			blk.FinalEndorse.digest = nil //
		}, true},
		{"REDACTED", func(blk *Ledger) { blk.FinalEndorseDigest = []byte("REDACTED") }, true},
		{"REDACTED", func(blk *Ledger) {
			blk.Txs[0] = Tx("REDACTED")
			blk.digest = nil //
		}, true},
		{"REDACTED", func(blk *Ledger) {
			blk.DataDigest = engineseed.Octets(len(blk.DataDigest))
		}, true},
		{"REDACTED", func(blk *Ledger) {
			blk.ProofDigest = []byte("REDACTED")
		}, true},
		{"REDACTED", func(blk *Ledger) {
			blk.Release.Ledger = 1
		}, true},
	}
	for i, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			ledger := CreateLedger(h, txs, endorse, evtCatalog)
			ledger.RecommenderLocation = valueCollection.FetchRecommender().Location
			tc.distortLedger(ledger)
			err = ledger.CertifySimple()
			assert.Equal(t, tc.expirationErr, err != nil, "REDACTED", i, err)
		})
	}
}

func VerifyLedgerDigest(t *testing.T) {
	assert.Nil(t, (*Ledger)(nil).Digest())
	assert.Nil(t, CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil).Digest())
}

func VerifyLedgerCreateSectionCollection(t *testing.T) {
	bps, err := (*Ledger)(nil).CreateSegmentAssign(2)
	assert.Error(t, err)
	assert.Nil(t, bps)

	sectionCollection, err := CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil).CreateSegmentAssign(1024)
	require.NoError(t, err)
	assert.NotNil(t, sectionCollection)
	assert.EqualValues(t, 1, sectionCollection.Sum())
}

func VerifyLedgerCreateSectionCollectionWithProof(t *testing.T) {
	bps, err := (*Ledger)(nil).CreateSegmentAssign(2)
	assert.Error(t, err)
	assert.Nil(t, bps)

	finalUID := createLedgerUIDArbitrary()
	h := int64(3)

	ballotCollection, _, values := randomBallotCollection(h-1, 1, engineproto.PreendorseKind, 10, 1, false)
	extensionEndorse, err := CreateExtensionEndorse(finalUID, h-1, 1, ballotCollection, values, time.Now(), false)
	require.NoError(t, err)

	ev, err := NewEmulateReplicatedBallotProofWithRatifier(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	evtCatalog := []Proof{ev}

	sectionCollection, err := CreateLedger(h, []Tx{Tx("REDACTED")}, extensionEndorse.ToEndorse(), evtCatalog).CreateSegmentAssign(512)
	require.NoError(t, err)

	assert.NotNil(t, sectionCollection)
	assert.EqualValues(t, 4, sectionCollection.Sum())
}

func VerifyLedgerDigestsTo(t *testing.T) {
	assert.False(t, (*Ledger)(nil).DigestsTo(nil))

	finalUID := createLedgerUIDArbitrary()
	h := int64(3)
	ballotCollection, valueCollection, values := randomBallotCollection(h-1, 1, engineproto.PreendorseKind, 10, 1, false)
	extensionEndorse, err := CreateExtensionEndorse(finalUID, h-1, 1, ballotCollection, values, time.Now(), false)
	require.NoError(t, err)

	ev, err := NewEmulateReplicatedBallotProofWithRatifier(h, time.Now(), values[0], "REDACTED")
	require.NoError(t, err)
	evtCatalog := []Proof{ev}

	ledger := CreateLedger(h, []Tx{Tx("REDACTED")}, extensionEndorse.ToEndorse(), evtCatalog)
	ledger.RatifiersDigest = valueCollection.Digest()
	assert.False(t, ledger.DigestsTo([]byte{}))
	assert.False(t, ledger.DigestsTo([]byte("REDACTED")))
	assert.True(t, ledger.DigestsTo(ledger.Digest()))
}

func VerifyLedgerVolume(t *testing.T) {
	volume := CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil).Volume()
	if volume <= 0 {
		t.Fatal("REDACTED")
	}
}

func VerifyLedgerString(t *testing.T) {
	assert.Equal(t, "REDACTED", (*Ledger)(nil).String())
	assert.Equal(t, "REDACTED", (*Ledger)(nil).StringIndented("REDACTED"))
	assert.Equal(t, "REDACTED", (*Ledger)(nil).StringBrief())

	ledger := CreateLedger(int64(3), []Tx{Tx("REDACTED")}, nil, nil)
	assert.NotEqual(t, "REDACTED", ledger.String())
	assert.NotEqual(t, "REDACTED", ledger.StringIndented("REDACTED"))
	assert.NotEqual(t, "REDACTED", ledger.StringBrief())
}

func createLedgerUIDArbitrary() LedgerUID {
	var (
		ledgerDigest   = make([]byte, comethash.Volume)
		sectionCollectionDigest = make([]byte, comethash.Volume)
	)
	rand.Read(ledgerDigest)   //
	rand.Read(sectionCollectionDigest) //
	return LedgerUID{ledgerDigest, SegmentAssignHeading{123, sectionCollectionDigest}}
}

func createLedgerUID(digest []byte, sectionCollectionVolume uint32, sectionCollectionDigest []byte) LedgerUID {
	var (
		h   = make([]byte, comethash.Volume)
		psH = make([]byte, comethash.Volume)
	)
	copy(h, digest)
	copy(psH, sectionCollectionDigest)
	return LedgerUID{
		Digest: h,
		SegmentAssignHeading: SegmentAssignHeading{
			Sum: sectionCollectionVolume,
			Digest:  psH,
		},
	}
}

var nullOctets []byte

//
var emptyOctets = []byte{
	0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8,
	0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b,
	0x78, 0x52, 0xb8, 0x55,
}

func VerifyNullHeadingDigestDoesntCollapse(t *testing.T) {
	assert.Equal(t, nullOctets, []byte((*Heading)(nil).Digest()))
	assert.Equal(t, nullOctets, []byte((new(Heading)).Digest()))
}

func VerifyNullDataDigestDoesntCollapse(t *testing.T) {
	assert.Equal(t, emptyOctets, []byte((*Data)(nil).Digest()))
	assert.Equal(t, emptyOctets, []byte(new(Data).Digest()))
}

func VerifyEndorse(t *testing.T) {
	finalUID := createLedgerUIDArbitrary()
	h := int64(3)
	ballotCollection, _, values := randomBallotCollection(h-1, 1, engineproto.PreendorseKind, 10, 1, true)
	extensionEndorse, err := CreateExtensionEndorse(finalUID, h-1, 1, ballotCollection, values, time.Now(), true)
	require.NoError(t, err)

	assert.Equal(t, h-1, extensionEndorse.Level)
	assert.EqualValues(t, 1, extensionEndorse.Cycle)
	assert.Equal(t, engineproto.PreendorseKind, engineproto.AttestedMessageKind(extensionEndorse.Kind()))
	if extensionEndorse.Volume() <= 0 {
		t.Fatalf("REDACTED", extensionEndorse, extensionEndorse.Volume())
	}

	require.NotNil(t, extensionEndorse.BitList())
	assert.Equal(t, bits.NewBitList(10).Volume(), extensionEndorse.BitList().Volume())

	assert.Equal(t, ballotCollection.FetchByOrdinal(0), extensionEndorse.FetchByOrdinal(0))
	assert.True(t, extensionEndorse.IsEndorse())
}

func VerifyEndorseCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		verifyLabel       string
		distortEndorse func(*Endorse)
		anticipateErr      bool
	}{
		{"REDACTED", func(com *Endorse) {}, false},
		{"REDACTED", func(com *Endorse) { com.Endorsements[0].Autograph = []byte{0} }, false},
		{"REDACTED", func(com *Endorse) { com.Level = int64(-100) }, true},
		{"REDACTED", func(com *Endorse) { com.Cycle = -100 }, true},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			com := randomEndorse(time.Now())
			tc.distortEndorse(com)
			assert.Equal(t, tc.anticipateErr, com.CertifySimple() != nil, "REDACTED")
		})
	}
}

func VerifyMaximumEndorseOctets(t *testing.T) {
	//
	//
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)

	cs := EndorseSignature{
		LedgerUIDMark:      LedgerUIDMarkNull,
		RatifierLocation: vault.LocationDigest([]byte("REDACTED")),
		Timestamp:        timestamp,
		Autograph:        vault.CRandomOctets(MaximumAutographVolume),
	}

	pbSignature := cs.ToSchema()
	//
	assert.EqualValues(t, MaximumEndorseSignatureOctets, pbSignature.Volume())

	//
	endorse := &Endorse{
		Level: math.MaxInt64,
		Cycle:  math.MaxInt32,
		LedgerUID: LedgerUID{
			Digest: comethash.Sum([]byte("REDACTED")),
			SegmentAssignHeading: SegmentAssignHeading{
				Sum: math.MaxInt32,
				Digest:  comethash.Sum([]byte("REDACTED")),
			},
		},
		Endorsements: []EndorseSignature{cs},
	}

	pb := endorse.ToSchema()

	assert.EqualValues(t, MaximumEndorseOctets(1), int64(pb.Volume()))

	//
	for i := 1; i < MaximumBallotsTally; i++ {
		endorse.Endorsements = append(endorse.Endorsements, cs)
	}

	pb = endorse.ToSchema()

	assert.EqualValues(t, MaximumEndorseOctets(MaximumBallotsTally), int64(pb.Volume()))
}

func VerifyHeadingDigest(t *testing.T) {
	verifyScenarios := []struct {
		note       string
		heading     *Heading
		anticipateDigest octets.HexOctets
	}{
		{"REDACTED", &Heading{
			Release:            cometrelease.Agreement{Ledger: 1, App: 2},
			LedgerUID:            "REDACTED",
			Level:             3,
			Time:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			FinalLedgerUID:        createLedgerUID(make([]byte, comethash.Volume), 6, make([]byte, comethash.Volume)),
			FinalEndorseDigest:     comethash.Sum([]byte("REDACTED")),
			DataDigest:           comethash.Sum([]byte("REDACTED")),
			RatifiersDigest:     comethash.Sum([]byte("REDACTED")),
			FollowingRatifiersDigest: comethash.Sum([]byte("REDACTED")),
			AgreementDigest:      comethash.Sum([]byte("REDACTED")),
			ApplicationDigest:            comethash.Sum([]byte("REDACTED")),
			FinalOutcomesDigest:    comethash.Sum([]byte("REDACTED")),
			ProofDigest:       comethash.Sum([]byte("REDACTED")),
			RecommenderLocation:    vault.LocationDigest([]byte("REDACTED")),
		}, hexOctetsFromString("REDACTED")},
		{"REDACTED", nil, nil},
		{"REDACTED", &Heading{
			Release:            cometrelease.Agreement{Ledger: 1, App: 2},
			LedgerUID:            "REDACTED",
			Level:             3,
			Time:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			FinalLedgerUID:        createLedgerUID(make([]byte, comethash.Volume), 6, make([]byte, comethash.Volume)),
			FinalEndorseDigest:     comethash.Sum([]byte("REDACTED")),
			DataDigest:           comethash.Sum([]byte("REDACTED")),
			RatifiersDigest:     nil,
			FollowingRatifiersDigest: comethash.Sum([]byte("REDACTED")),
			AgreementDigest:      comethash.Sum([]byte("REDACTED")),
			ApplicationDigest:            comethash.Sum([]byte("REDACTED")),
			FinalOutcomesDigest:    comethash.Sum([]byte("REDACTED")),
			ProofDigest:       comethash.Sum([]byte("REDACTED")),
			RecommenderLocation:    vault.LocationDigest([]byte("REDACTED")),
		}, nil},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.note, func(t *testing.T) {
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
					case int64, octets.HexOctets, string:
						octetSegments = append(octetSegments, cdcEncode(f))
					case time.Time:
						bz, err := gogotypes.StdTimeMarshal(f)
						require.NoError(t, err)
						octetSegments = append(octetSegments, bz)
					case cometrelease.Agreement:
						bz, err := f.Serialize()
						require.NoError(t, err)
						octetSegments = append(octetSegments, bz)
					case LedgerUID:
						pbbi := f.ToSchema()
						bz, err := pbbi.Serialize()
						require.NoError(t, err)
						octetSegments = append(octetSegments, bz)
					default:
						t.Errorf("REDACTED", f)
					}
				}
				assert.Equal(t,
					octets.HexOctets(merkle.DigestFromOctetSegments(octetSegments)), tc.heading.Digest())
			}
		})
	}
}

func VerifyMaximumHeadingOctets(t *testing.T) {
	//
	//
	//
	//
	maximumSeriesUID := "REDACTED"
	for i := 0; i < MaximumSeriesUIDSize; i++ {
		maximumSeriesUID += "REDACTED"
	}

	//
	//
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)

	h := Heading{
		Release:            cometrelease.Agreement{Ledger: math.MaxInt64, App: math.MaxInt64},
		LedgerUID:            maximumSeriesUID,
		Level:             math.MaxInt64,
		Time:               timestamp,
		FinalLedgerUID:        createLedgerUID(make([]byte, comethash.Volume), math.MaxInt32, make([]byte, comethash.Volume)),
		FinalEndorseDigest:     comethash.Sum([]byte("REDACTED")),
		DataDigest:           comethash.Sum([]byte("REDACTED")),
		RatifiersDigest:     comethash.Sum([]byte("REDACTED")),
		FollowingRatifiersDigest: comethash.Sum([]byte("REDACTED")),
		AgreementDigest:      comethash.Sum([]byte("REDACTED")),
		ApplicationDigest:            comethash.Sum([]byte("REDACTED")),
		FinalOutcomesDigest:    comethash.Sum([]byte("REDACTED")),
		ProofDigest:       comethash.Sum([]byte("REDACTED")),
		RecommenderLocation:    vault.LocationDigest([]byte("REDACTED")),
	}

	bz, err := h.ToSchema().Serialize()
	require.NoError(t, err)

	assert.EqualValues(t, MaximumHeadingOctets, int64(len(bz)))
}

func randomEndorse(now time.Time) *Endorse {
	finalUID := createLedgerUIDArbitrary()
	h := int64(3)
	ballotCollection, _, values := randomBallotCollection(h-1, 1, engineproto.PreendorseKind, 10, 1, false)
	extensionEndorse, err := CreateExtensionEndorse(finalUID, h-1, 1, ballotCollection, values, now, false)
	if err != nil {
		panic(err)
	}
	return extensionEndorse.ToEndorse()
}

func hexOctetsFromString(s string) octets.HexOctets {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return octets.HexOctets(b)
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

func VerifyLedgerMaximumDataOctetsNoProof(t *testing.T) {
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
				MaximumDataOctetsNoProof(tc.maximumOctets, tc.valuesTally)
			}, "REDACTED", i)
		} else {
			assert.Equal(t,
				tc.outcome,
				MaximumDataOctetsNoProof(tc.maximumOctets, tc.valuesTally),
				"REDACTED", i)
		}
	}
}

//
//
//
//
func VerifyBallotCollectionToExpandedEndorse(t *testing.T) {
	for _, verifyInstance := range []struct {
		label             string
		encompassAddition bool
	}{
		{
			label:             "REDACTED",
			encompassAddition: false,
		},
		{
			label:             "REDACTED",
			encompassAddition: true,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			ledgerUID := createLedgerUIDArbitrary()

			valueCollection, values := RandomRatifierCollection(10, 1)
			var ballotCollection *BallotCollection
			if verifyInstance.encompassAddition {
				ballotCollection = NewExpandedBallotCollection("REDACTED", 3, 1, engineproto.PreendorseKind, valueCollection)
			} else {
				ballotCollection = NewBallotCollection("REDACTED", 3, 1, engineproto.PreendorseKind, valueCollection)
			}
			for i := 0; i < len(values); i++ {
				publicKey, err := values[i].FetchPublicKey()
				require.NoError(t, err)
				ballot := &Ballot{
					RatifierLocation: publicKey.Location(),
					RatifierOrdinal:   int32(i),
					Level:           3,
					Cycle:            1,
					Kind:             engineproto.PreendorseKind,
					LedgerUID:          ledgerUID,
					Timestamp:        time.Now(),
				}
				v := ballot.ToSchema()
				err = values[i].AttestBallot(ballotCollection.LedgerUID(), v)
				require.NoError(t, err)
				ballot.Autograph = v.Autograph
				if verifyInstance.encompassAddition {
					ballot.AdditionAutograph = v.AdditionAutograph
				}
				appended, err := ballotCollection.AppendBallot(ballot)
				require.NoError(t, err)
				require.True(t, appended)
			}
			var veLevel int64
			if verifyInstance.encompassAddition {
				veLevel = 1
			}
			ec := ballotCollection.CreateExpandedEndorse(IfaceOptions{BallotPluginsActivateLevel: veLevel})

			for i := int32(0); int(i) < len(values); i++ {
				vote1 := ballotCollection.FetchByOrdinal(i)
				ballot2 := ec.FetchExpandedBallot(i)

				vote1bz, err := vote1.ToSchema().Serialize()
				require.NoError(t, err)
				vote2bz, err := ballot2.ToSchema().Serialize()
				require.NoError(t, err)
				assert.Equal(t, vote1bz, vote2bz)
			}
		})
	}
}

//
//
//
func toBallotCollection(ec *ExpandedEndorse, ledgerUID string, values *RatifierAssign) *BallotCollection {
	ballotCollection := NewBallotCollection(ledgerUID, ec.Level, ec.Cycle, engineproto.PreendorseKind, values)
	ec.appendAutographsToBallotCollection(ballotCollection)
	return ballotCollection
}

//
//
//
//
func VerifyExpandedEndorseToBallotCollection(t *testing.T) {
	for _, verifyInstance := range []struct {
		label             string
		encompassAddition bool
	}{
		{
			label:             "REDACTED",
			encompassAddition: false,
		},
		{
			label:             "REDACTED",
			encompassAddition: true,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			finalUID := createLedgerUIDArbitrary()
			h := int64(3)

			ballotCollection, valueCollection, values := randomBallotCollection(h-1, 1, engineproto.PreendorseKind, 10, 1, true)
			extensionEndorse, err := CreateExtensionEndorse(finalUID, h-1, 1, ballotCollection, values, time.Now(), true)
			assert.NoError(t, err)

			if !verifyInstance.encompassAddition {
				for i := 0; i < len(values); i++ {
					v := ballotCollection.FetchByOrdinal(int32(i))
					v.Addition = nil
					v.AdditionAutograph = nil
					extensionEndorse.ExpandedEndorsements[i].Addition = nil
					extensionEndorse.ExpandedEndorsements[i].AdditionAutograph = nil
				}
			}

			ledgerUID := ballotCollection.LedgerUID()
			var ballotSet2 *BallotCollection
			if verifyInstance.encompassAddition {
				ballotSet2 = extensionEndorse.ToExpandedBallotCollection(ledgerUID, valueCollection)
			} else {
				ballotSet2 = toBallotCollection(extensionEndorse, ledgerUID, valueCollection)
			}

			for i := int32(0); int(i) < len(values); i++ {
				vote1 := ballotCollection.FetchByOrdinal(i)
				ballot2 := ballotSet2.FetchByOrdinal(i)
				vote3 := extensionEndorse.FetchExpandedBallot(i)

				vote1bz, err := vote1.ToSchema().Serialize()
				require.NoError(t, err)
				vote2bz, err := ballot2.ToSchema().Serialize()
				require.NoError(t, err)
				vote3bz, err := vote3.ToSchema().Serialize()
				require.NoError(t, err)
				assert.Equal(t, vote1bz, vote2bz)
				assert.Equal(t, vote1bz, vote3bz)
			}
		})
	}
}

func VerifyEndorseToBallotCollectionWithBallotsForNullLedger(t *testing.T) {
	ledgerUID := createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED"))

	const (
		level = int64(3)
		epoch  = 0
	)

	type endorseBallotVerify struct {
		ledgerIDXDatastore      []LedgerUID
		countBallots      []int //
		countRatifiers int
		sound         bool
	}

	verifyScenarios := []endorseBallotVerify{
		{[]LedgerUID{ledgerUID, {}}, []int{67, 33}, 100, true},
	}

	for _, tc := range verifyScenarios {
		ballotCollection, valueCollection, values := randomBallotCollection(level-1, epoch, engineproto.PreendorseKind, tc.countRatifiers, 1, false)

		vi := int32(0)
		for n := range tc.ledgerIDXDatastore {
			for i := 0; i < tc.countBallots[n]; i++ {
				publicKey, err := values[vi].FetchPublicKey()
				require.NoError(t, err)
				ballot := &Ballot{
					RatifierLocation: publicKey.Location(),
					RatifierOrdinal:   vi,
					Level:           level - 1,
					Cycle:            epoch,
					Kind:             engineproto.PreendorseKind,
					LedgerUID:          tc.ledgerIDXDatastore[n],
					Timestamp:        engineclock.Now(),
				}

				appended, err := attestAppendBallot(values[vi], ballot, ballotCollection)
				assert.NoError(t, err)
				assert.True(t, appended)

				vi++
			}
		}

		veLevelArgument := IfaceOptions{BallotPluginsActivateLevel: 0}
		if tc.sound {
			extensionEndorse := ballotCollection.CreateExpandedEndorse(veLevelArgument) //
			assert.NotNil(t, extensionEndorse)
			err := valueCollection.ValidateEndorse(ballotCollection.LedgerUID(), ledgerUID, level-1, extensionEndorse.ToEndorse())
			assert.Nil(t, err)
		} else {
			assert.Panics(t, func() { ballotCollection.CreateExpandedEndorse(veLevelArgument) })
		}
	}
}

func VerifyLedgerUIDCertifySimple(t *testing.T) {
	soundLedgerUID := LedgerUID{
		Digest: octets.HexOctets{},
		SegmentAssignHeading: SegmentAssignHeading{
			Sum: 1,
			Digest:  octets.HexOctets{},
		},
	}

	corruptLedgerUID := LedgerUID{
		Digest: []byte{0},
		SegmentAssignHeading: SegmentAssignHeading{
			Sum: 1,
			Digest:  []byte{0},
		},
	}

	verifyScenarios := []struct {
		verifyLabel             string
		ledgerUIDDigest          octets.HexOctets
		ledgerUIDSectionCollectionHeading SegmentAssignHeading
		anticipateErr            bool
	}{
		{"REDACTED", soundLedgerUID.Digest, soundLedgerUID.SegmentAssignHeading, false},
		{"REDACTED", corruptLedgerUID.Digest, soundLedgerUID.SegmentAssignHeading, true},
		{"REDACTED", soundLedgerUID.Digest, corruptLedgerUID.SegmentAssignHeading, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			ledgerUID := LedgerUID{
				Digest:          tc.ledgerUIDDigest,
				SegmentAssignHeading: tc.ledgerUIDSectionCollectionHeading,
			}
			assert.Equal(t, tc.anticipateErr, ledgerUID.CertifySimple() != nil, "REDACTED")
		})
	}
}

func VerifyLedgerSchemaBuffer(t *testing.T) {
	h := engineseed.Int63()
	c1 := randomEndorse(time.Now())
	b1 := CreateLedger(h, []Tx{Tx([]byte{1})}, &Endorse{Endorsements: []EndorseSignature{}}, []Proof{})
	b1.RecommenderLocation = engineseed.Octets(vault.LocationVolume)

	b2 := CreateLedger(h, []Tx{Tx([]byte{1})}, c1, []Proof{})
	b2.RecommenderLocation = engineseed.Octets(vault.LocationVolume)
	proofTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	evi, err := NewEmulateReplicatedBallotProof(h, proofTime, "REDACTED")
	require.NoError(t, err)
	b2.Proof = ProofData{Proof: ProofCatalog{evi}}
	b2.ProofDigest = b2.Proof.Digest()

	b3 := CreateLedger(h, []Tx{}, c1, []Proof{})
	b3.RecommenderLocation = engineseed.Octets(vault.LocationVolume)
	verifyScenarios := []struct {
		msg      string
		b1       *Ledger
		expirationPass  bool
		expirationPass2 bool
	}{
		{"REDACTED", nil, false, false},
		{"REDACTED", b1, true, true},
		{"REDACTED", b2, true, true},
		{"REDACTED", b3, true, true},
	}
	for _, tc := range verifyScenarios {
		pb, err := tc.b1.ToSchema()
		if tc.expirationPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		ledger, err := LedgerFromSchema(pb)
		if tc.expirationPass2 {
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
	data2 := &Data{Txs: Txs{}}
	verifyScenarios := []struct {
		msg     string
		data1   *Data
		expirationPass bool
	}{
		{"REDACTED", data, true},
		{"REDACTED", data2, true},
	}
	for _, tc := range verifyScenarios {
		schemaData := tc.data1.ToSchema()
		d, err := DataFromSchema(&schemaData)
		if tc.expirationPass {
			require.NoError(t, err, tc.msg)
			require.EqualValues(t, tc.data1, &d, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

//
func VerifyProofDataSchemaBuffer(t *testing.T) {
	const ledgerUID = "REDACTED"
	ev, err := NewEmulateReplicatedBallotProof(math.MaxInt64, time.Now(), ledgerUID)
	require.NoError(t, err)
	data := &ProofData{Proof: ProofCatalog{ev}}
	_ = data.OctetVolume()
	verifyScenarios := []struct {
		msg      string
		data1    *ProofData
		expirationPass1 bool
		expirationPass2 bool
	}{
		{"REDACTED", data, true, true},
		{"REDACTED", &ProofData{Proof: ProofCatalog{}}, true, true},
		{"REDACTED", nil, false, false},
	}

	for _, tc := range verifyScenarios {
		schemaData, err := tc.data1.ToSchema()
		if tc.expirationPass1 {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		eviD := new(ProofData)
		err = eviD.FromSchema(schemaData)
		if tc.expirationPass2 {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.data1, eviD, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func createRandomHeading() Heading {
	ledgerUID := "REDACTED"
	t := time.Now()
	level := engineseed.Int63()
	randomOctets := engineseed.Octets(comethash.Volume)
	randomLocation := engineseed.Octets(vault.LocationVolume)
	h := Heading{
		Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1},
		LedgerUID:            ledgerUID,
		Level:             level,
		Time:               t,
		FinalLedgerUID:        LedgerUID{},
		FinalEndorseDigest:     randomOctets,
		DataDigest:           randomOctets,
		RatifiersDigest:     randomOctets,
		FollowingRatifiersDigest: randomOctets,
		AgreementDigest:      randomOctets,
		ApplicationDigest:            randomOctets,

		FinalOutcomesDigest: randomOctets,

		ProofDigest:    randomOctets,
		RecommenderLocation: randomLocation,
	}

	return h
}

func VerifyHeadingSchema(t *testing.T) {
	h1 := createRandomHeading()
	tc := []struct {
		msg     string
		h1      *Heading
		expirationPass bool
	}{
		{"REDACTED", &h1, true},
		{"REDACTED", &Heading{}, false},
	}

	for _, tt := range tc {
		t.Run(tt.msg, func(t *testing.T) {
			pb := tt.h1.ToSchema()
			h, err := HeadingFromSchema(pb)
			if tt.expirationPass {
				require.NoError(t, err, tt.msg)
				require.Equal(t, tt.h1, &h, tt.msg)
			} else {
				require.Error(t, err, tt.msg)
			}
		})
	}
}

func VerifyLedgerUIDSchemaBuffer(t *testing.T) {
	ledgerUID := createLedgerUID([]byte("REDACTED"), 2, []byte("REDACTED"))
	verifyScenarios := []struct {
		msg     string
		bid1    *LedgerUID
		expirationPass bool
	}{
		{"REDACTED", &ledgerUID, true},
		{"REDACTED", &LedgerUID{}, true},
		{"REDACTED", nil, false},
	}
	for _, tc := range verifyScenarios {
		schemaLedgerUID := tc.bid1.ToSchema()

		bi, err := LedgerUIDFromSchema(&schemaLedgerUID)
		if tc.expirationPass {
			require.NoError(t, err)
			require.Equal(t, tc.bid1, bi, tc.msg)
		} else {
			require.NotEqual(t, tc.bid1, bi, tc.msg)
		}
	}
}

func VerifyAttestedHeadingSchemaBuffer(t *testing.T) {
	endorse := randomEndorse(time.Now())
	h := createRandomHeading()

	sh := AttestedHeading{Heading: &h, Endorse: endorse}

	verifyScenarios := []struct {
		msg     string
		sh1     *AttestedHeading
		expirationPass bool
	}{
		{"REDACTED", &AttestedHeading{}, true},
		{"REDACTED", &sh, true},
		{"REDACTED", nil, false},
	}
	for _, tc := range verifyScenarios {
		schemaAttestedHeading := tc.sh1.ToSchema()

		sh, err := AttestedHeadingFromSchema(schemaAttestedHeading)

		if tc.expirationPass {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.sh1, sh, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func VerifyLedgerUIDMatches(t *testing.T) {
	var (
		ledgerUID          = createLedgerUID([]byte("REDACTED"), 2, []byte("REDACTED"))
		ledgerUIDReplicated = createLedgerUID([]byte("REDACTED"), 2, []byte("REDACTED"))
		ledgerUIDDistinct = createLedgerUID([]byte("REDACTED"), 2, []byte("REDACTED"))
		ledgerUIDEmpty     = LedgerUID{}
	)

	assert.True(t, ledgerUID.Matches(ledgerUIDReplicated))
	assert.False(t, ledgerUID.Matches(ledgerUIDDistinct))
	assert.False(t, ledgerUID.Matches(ledgerUIDEmpty))
	assert.True(t, ledgerUIDEmpty.Matches(ledgerUIDEmpty)) //
	assert.False(t, ledgerUIDEmpty.Matches(ledgerUIDDistinct))
}
