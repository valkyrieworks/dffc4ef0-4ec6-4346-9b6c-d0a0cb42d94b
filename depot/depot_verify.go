package depot

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitstore "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/depot"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//
//
type sanitizeMethod func()

//
//
func createVerifyAddnEndorse(altitude int64, timestamp time.Time) *kinds.ExpandedEndorse {
	return createVerifyAddnEndorseUsingCountSignatures(altitude, timestamp, 1)
}

func createVerifyAddnEndorseUsingCountSignatures(altitude int64, timestamp time.Time, countSignatures int) *kinds.ExpandedEndorse {
	addnEndorseSignatures := []kinds.ExpandedEndorseSignature{}
	for i := 0; i < countSignatures; i++ {
		addnEndorseSignatures = append(addnEndorseSignatures, kinds.ExpandedEndorseSignature{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
				AssessorLocation: commitrand.Octets(security.LocatorExtent),
				Timestamp:        timestamp,
				Notation:        commitrand.Octets(64),
			},
			AdditionNotation: []byte("REDACTED"),
		})
	}
	return &kinds.ExpandedEndorse{
		Altitude: altitude,
		LedgerUUID: kinds.LedgerUUID{
			Digest:          security.CHARArbitraryOctets(32),
			FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: security.CHARArbitraryOctets(32), Sum: 2},
		},
		ExpandedNotations: addnEndorseSignatures,
	}
}

func createStatusAlsoLedgerDepot() (sm.Status, *LedgerDepot, sanitizeMethod) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	//
	//
	ledgerDatastore := dbm.FreshMemoryDatastore()
	statusDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status, err := statusDepot.FetchOriginatingDatastoreEitherInaugurationRecord(settings.InaugurationRecord())
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	return status, FreshLedgerDepot(ledgerDatastore), func() { os.RemoveAll(settings.OriginPath) }
}

func VerifyFetchLedgerDepotStatus(t *testing.T) {
	type ledgerDepotVerify struct {
		verifyAlias string
		bss      *commitstore.LedgerDepotStatus
		desire     commitstore.LedgerDepotStatus
	}

	verifyScenarios := []ledgerDepotVerify{
		{
			"REDACTED", &commitstore.LedgerDepotStatus{Foundation: 100, Altitude: 1000},
			commitstore.LedgerDepotStatus{Foundation: 100, Altitude: 1000},
		},
		{"REDACTED", &commitstore.LedgerDepotStatus{}, commitstore.LedgerDepotStatus{}},
		{"REDACTED", &commitstore.LedgerDepotStatus{Altitude: 1000}, commitstore.LedgerDepotStatus{Foundation: 1, Altitude: 1000}},
	}

	for _, tc := range verifyScenarios {
		db := dbm.FreshMemoryDatastore()
		cluster := db.FreshCluster()
		PersistLedgerDepotStatusCluster(tc.bss, cluster)
		err := cluster.PersistChronize()
		require.NoError(t, err)
		retrieveBytesjsn := FetchLedgerDepotStatus(db)
		assert.Equal(t, tc.desire, retrieveBytesjsn, "REDACTED", tc.verifyAlias)
		err = cluster.Shutdown()
		require.NoError(t, err)
	}
}

func VerifyFreshLedgerDepot(t *testing.T) {
	db := dbm.FreshMemoryDatastore()
	bss := commitstore.LedgerDepotStatus{Foundation: 100, Altitude: 10000}
	bz, _ := proto.Marshal(&bss)
	err := db.Set(ledgerDepotToken, bz)
	require.NoError(t, err)
	bs := FreshLedgerDepot(db)
	require.Equal(t, int64(100), bs.Foundation(), "REDACTED")
	require.Equal(t, int64(10000), bs.Altitude(), "REDACTED")

	alarmCausees := []struct {
		data    []byte
		desireFault string
	}{
		{[]byte("REDACTED"), "REDACTED"},
		{[]byte("REDACTED"), "REDACTED"},
	}

	for i, tt := range alarmCausees {

		//
		_, _, alarmFault := conductProc(func() (any, error) {
			err := db.Set(ledgerDepotToken, tt.data)
			require.NoError(t, err)
			_ = FreshLedgerDepot(db)
			return nil, nil
		})
		require.NotNil(t, alarmFault, "REDACTED", i, tt.data)
		assert.Contains(t, fmt.Sprintf("REDACTED", alarmFault), tt.desireFault, "REDACTED", i, tt.data)
	}

	err = db.Set(ledgerDepotToken, []byte{})
	require.NoError(t, err)
	bs = FreshLedgerDepot(db)
	assert.Equal(t, bs.Altitude(), int64(0), "REDACTED")
}

func freshInsideRamLedgerDepot() (*LedgerDepot, dbm.DB) {
	db := dbm.FreshMemoryDatastore()
	return FreshLedgerDepot(db), db
}

//

func VerifyLedgerDepotPersistFetchLedger(t *testing.T) {
	status, bs, sanitize := createStatusAlsoLedgerDepot()
	defer sanitize()
	require.Equal(t, bs.Foundation(), int64(0), "REDACTED")
	require.Equal(t, bs.Altitude(), int64(0), "REDACTED")

	//
	negativeLedgerElevations := []int64{0, -1, 100, 1000, 2}
	for i, altitude := range negativeLedgerElevations {
		if g := bs.FetchLedger(altitude); g != nil {
			t.Errorf("REDACTED", i, altitude)
		}
	}

	//
	txs := []kinds.Tx{make([]byte, kinds.LedgerFragmentExtentOctets)} //
	ledger, err := status.CreateLedger(bs.Altitude()+1, txs, new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
	require.NoError(t, err)
	soundFragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	require.GreaterOrEqual(t, soundFragmentAssign.Sum(), uint32(2))
	section2 := soundFragmentAssign.ObtainFragment(1)

	observedEndorse := createVerifyAddnEndorse(ledger.Altitude, committime.Now())
	bs.PersistLedgerUsingExpandedEndorse(ledger, soundFragmentAssign, observedEndorse)
	require.EqualValues(t, 1, bs.Foundation(), "REDACTED")
	require.EqualValues(t, ledger.Altitude, bs.Altitude(), "REDACTED")

	partialFragmentAssign := kinds.FreshFragmentAssignOriginatingHeading(kinds.FragmentAssignHeading{Sum: 2})
	noncontiguousFragmentAssign := kinds.FreshFragmentAssignOriginatingHeading(kinds.FragmentAssignHeading{Sum: 0})
	_, err = noncontiguousFragmentAssign.AppendFragment(section2)
	require.Error(t, err)

	heading1 := kinds.Heading{
		Edition:         strongmindedition.Agreement{Ledger: edition.LedgerScheme},
		Altitude:          1,
		SuccessionUUID:         "REDACTED",
		Moment:            committime.Now(),
		NominatorLocation: commitrand.Octets(security.LocatorExtent),
	}

	//

	endorseLocatedLevel10 := createVerifyAddnEndorse(10, committime.Now()).TowardEndorse()
	groups := []struct {
		ledger      *kinds.Ledger
		fragments      *kinds.FragmentAssign
		observedEndorse *kinds.ExpandedEndorse
		desireAlarm  string
		desireFault    bool

		invalidLedgerInsideDatastore      bool
		invalidEndorseInsideDatastore     bool
		invalidObservedEndorseInsideDatastore bool
		removeEndorseInsideDatastore       bool
		removeObservedEndorseInsideDatastore   bool
	}{
		{
			ledger:      freshLedger(heading1, endorseLocatedLevel10),
			fragments:      soundFragmentAssign,
			observedEndorse: observedEndorse,
		},

		{
			ledger:     nil,
			desireAlarm: "REDACTED",
		},

		{
			ledger: freshLedger( //
				kinds.Heading{
					Edition:         strongmindedition.Agreement{Ledger: edition.LedgerScheme},
					Altitude:          5,
					SuccessionUUID:         "REDACTED",
					Moment:            committime.Now(),
					NominatorLocation: commitrand.Octets(security.LocatorExtent),
				},
				createVerifyAddnEndorse(5, committime.Now()).TowardEndorse(),
			),
			fragments:      soundFragmentAssign,
			observedEndorse: createVerifyAddnEndorse(5, committime.Now()),
		},

		{
			ledger:      freshLedger(heading1, endorseLocatedLevel10),
			fragments:      partialFragmentAssign,
			desireAlarm:  "REDACTED", //
			observedEndorse: createVerifyAddnEndorse(10, committime.Now()),
		},

		{
			ledger:             freshLedger(heading1, endorseLocatedLevel10),
			fragments:             soundFragmentAssign,
			observedEndorse:        observedEndorse,
			invalidEndorseInsideDatastore: true, //
			desireAlarm:         "REDACTED",
		},

		{
			ledger:            freshLedger(heading1, endorseLocatedLevel10),
			fragments:            soundFragmentAssign,
			observedEndorse:       observedEndorse,
			desireAlarm:        "REDACTED",
			invalidLedgerInsideDatastore: true, //
		},

		{
			ledger:      freshLedger(heading1, endorseLocatedLevel10),
			fragments:      soundFragmentAssign,
			observedEndorse: observedEndorse,

			//
			removeObservedEndorseInsideDatastore: true,
		},

		{
			ledger:      ledger,
			fragments:      soundFragmentAssign,
			observedEndorse: observedEndorse,

			invalidObservedEndorseInsideDatastore: true,
			desireAlarm:             "REDACTED",
		},

		{
			ledger:      ledger,
			fragments:      soundFragmentAssign,
			observedEndorse: observedEndorse,

			//
			removeEndorseInsideDatastore: true,
		},
	}

	type quartet struct {
		ledger  *kinds.Ledger
		endorse *kinds.Endorse
		summary   *kinds.LedgerSummary

		observedEndorse *kinds.Endorse
	}

	for i, group := range groups {

		bs, db := freshInsideRamLedgerDepot()
		//
		res, err, alarmFault := conductProc(func() (any, error) {
			bs.PersistLedgerUsingExpandedEndorse(group.ledger, group.fragments, group.observedEndorse)
			if group.ledger == nil {
				return nil, nil
			}

			if group.invalidLedgerInsideDatastore {
				err := db.Set(reckonLedgerSummaryToken(group.ledger.Altitude), []byte("REDACTED"))
				require.NoError(t, err)
			}
			byteLedger := bs.FetchLedger(group.ledger.Altitude)
			byteLedgerSummary := bs.FetchLedgerSummary(group.ledger.Altitude)

			if group.removeObservedEndorseInsideDatastore {
				err := db.Erase(reckonObservedEndorseToken(group.ledger.Altitude))
				require.NoError(t, err)
			}
			if group.invalidObservedEndorseInsideDatastore {
				err := db.Set(reckonObservedEndorseToken(group.ledger.Altitude), []byte("REDACTED"))
				require.NoError(t, err)
			}
			byteObservedEndorse := bs.FetchObservedEndorse(group.ledger.Altitude)

			endorseAltitude := group.ledger.Altitude - 1
			if group.removeEndorseInsideDatastore {
				err := db.Erase(reckonLedgerEndorseToken(endorseAltitude))
				require.NoError(t, err)
			}
			if group.invalidEndorseInsideDatastore {
				err := db.Set(reckonLedgerEndorseToken(endorseAltitude), []byte("REDACTED"))
				require.NoError(t, err)
			}
			byteEndorse := bs.FetchLedgerEndorse(endorseAltitude)
			return &quartet{
				ledger: byteLedger, observedEndorse: byteObservedEndorse, endorse: byteEndorse,
				summary: byteLedgerSummary,
			}, nil
		})

		if underTxt := group.desireAlarm; underTxt != "REDACTED" {
			if alarmFault == nil {
				t.Errorf("REDACTED", i)
			} else if got := fmt.Sprintf("REDACTED", alarmFault); !strings.Contains(got, underTxt) {
				t.Errorf("REDACTED", i, got, underTxt)
			}
			continue
		}

		if group.desireFault {
			if err == nil {
				t.Errorf("REDACTED", i)
			}
			continue
		}

		assert.Nil(t, alarmFault, "REDACTED", i)
		assert.Nil(t, err, "REDACTED", i)
		qua, ok := res.(*quartet)
		if !ok || qua == nil {
			t.Errorf("REDACTED", i, res)
			continue
		}
		if group.removeObservedEndorseInsideDatastore {
			assert.Nil(t, qua.observedEndorse,
				"REDACTED")
		}
		if group.removeEndorseInsideDatastore {
			assert.Nil(t, qua.endorse,
				"REDACTED")
		}
	}
}

//
//
//
func extractAdditions(ec *kinds.ExpandedEndorse) bool {
	extracted := false
	for idx := range ec.ExpandedNotations {
		if len(ec.ExpandedNotations[idx].Addition) > 0 || len(ec.ExpandedNotations[idx].AdditionNotation) > 0 {
			extracted = true
		}
		ec.ExpandedNotations[idx].Addition = nil
		ec.ExpandedNotations[idx].AdditionNotation = nil
	}
	return extracted
}

//
//
func VerifyPersistLedgerUsingExpandedEndorseAlarmUponMissingAddition(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias           string
		distortEndorse func(*kinds.ExpandedEndorse)
		mustAlarm    bool
	}{
		{
			alias:           "REDACTED",
			distortEndorse: func(_ *kinds.ExpandedEndorse) {},
			mustAlarm:    false,
		},
		{
			alias: "REDACTED",
			distortEndorse: func(c *kinds.ExpandedEndorse) {
				extractAdditions(c)
			},
			mustAlarm: true,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			status, bs, sanitize := createStatusAlsoLedgerDepot()
			defer sanitize()
			h := bs.Altitude() + 1
			ledger, err := status.CreateLedger(h, verify.CreateNTHTrans(h, 10), new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
			require.NoError(t, err)

			observedEndorse := createVerifyAddnEndorse(ledger.Altitude, committime.Now())
			ps, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
			require.NoError(t, err)
			verifyInstance.distortEndorse(observedEndorse)
			if verifyInstance.mustAlarm {
				require.Panics(t, func() {
					bs.PersistLedgerUsingExpandedEndorse(ledger, ps, observedEndorse)
				})
			} else {
				bs.PersistLedgerUsingExpandedEndorse(ledger, ps, observedEndorse)
			}
		})
	}
}

//
//
//
func VerifyFetchLedgerExpandedEndorse(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias         string
		persistExpanded bool
		anticipateOutcome bool
	}{
		{
			alias:         "REDACTED",
			persistExpanded: false,
			anticipateOutcome: false,
		},
		{
			alias:         "REDACTED",
			persistExpanded: true,
			anticipateOutcome: true,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			status, bs, sanitize := createStatusAlsoLedgerDepot()
			defer sanitize()
			h := bs.Altitude() + 1
			ledger, err := status.CreateLedger(h, verify.CreateNTHTrans(h, 10), new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
			require.NoError(t, err)
			observedEndorse := createVerifyAddnEndorse(ledger.Altitude, committime.Now())
			ps, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
			require.NoError(t, err)
			if verifyInstance.persistExpanded {
				bs.PersistLedgerUsingExpandedEndorse(ledger, ps, observedEndorse)
			} else {
				bs.PersistLedger(ledger, ps, observedEndorse.TowardEndorse())
			}
			res := bs.FetchLedgerExpandedEndorse(ledger.Altitude)
			if verifyInstance.anticipateOutcome {
				require.Equal(t, observedEndorse, res)
			} else {
				require.Nil(t, res)
			}
		})
	}
}

func VerifyFetchFoundationSummary(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	statusDepot := sm.FreshDepot(dbm.FreshMemoryDatastore(), sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status, err := statusDepot.FetchOriginatingDatastoreEitherInaugurationRecord(settings.InaugurationRecord())
	require.NoError(t, err)
	bs := FreshLedgerDepot(dbm.FreshMemoryDatastore())

	for h := int64(1); h <= 10; h++ {
		ledger, err := status.CreateLedger(h, verify.CreateNTHTrans(h, 10), new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
		require.NoError(t, err)
		fragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
		require.NoError(t, err)
		observedEndorse := createVerifyAddnEndorse(h, committime.Now())
		bs.PersistLedgerUsingExpandedEndorse(ledger, fragmentAssign, observedEndorse)
	}

	_, _, err = bs.TrimLedgers(4, status)
	require.NoError(t, err)

	foundationLedger := bs.FetchFoundationSummary()
	assert.EqualValues(t, 4, foundationLedger.Heading.Altitude)
	assert.EqualValues(t, 4, bs.Foundation())

	require.NoError(t, bs.EraseNewestLedger())
	require.EqualValues(t, 9, bs.Altitude())
}

func VerifyFetchLedgerFragment(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")

	bs, db := freshInsideRamLedgerDepot()
	const altitude, ordinal = 10, 1
	fetchFragment := func() (any, error) {
		fragment := bs.FetchLedgerFragment(altitude, ordinal)
		return fragment, nil
	}

	status, err := sm.CreateInaugurationStatusOriginatingRecord(settings.InaugurationRecord())
	require.NoError(t, err)

	//
	//
	res, _, alarmFault := conductProc(fetchFragment)
	require.Nil(t, alarmFault, "REDACTED")
	require.Nil(t, res, "REDACTED")

	//
	err = db.Set(reckonLedgerFragmentToken(altitude, ordinal), []byte("REDACTED"))
	require.NoError(t, err)
	res, _, alarmFault = conductProc(fetchFragment)
	require.NotNil(t, alarmFault, "REDACTED")
	require.Contains(t, alarmFault.Error(), "REDACTED")

	//
	ledger, err := status.CreateLedger(altitude, nil, new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
	require.NoError(t, err)
	fragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	section1 := fragmentAssign.ObtainFragment(0)

	pb1, err := section1.TowardSchema()
	require.NoError(t, err)
	err = db.Set(reckonLedgerFragmentToken(altitude, ordinal), shouldEncode(pb1))
	require.NoError(t, err)
	attainedFragment, _, alarmFault := conductProc(fetchFragment)
	require.Nil(t, alarmFault, "REDACTED")
	require.Nil(t, res, "REDACTED")

	//
	attainedFragmentJSN, err := json.Marshal(attainedFragment.(*kinds.Fragment))
	require.NoError(t, err)
	section1jsn, err := json.Marshal(section1)
	require.NoError(t, err)
	require.JSONEq(t, string(attainedFragmentJSN), string(section1jsn),
		"REDACTED")
}

func VerifyTrimLedgers(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	statusDepot := sm.FreshDepot(dbm.FreshMemoryDatastore(), sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status, err := statusDepot.FetchOriginatingDatastoreEitherInaugurationRecord(settings.InaugurationRecord())
	require.NoError(t, err)
	db := dbm.FreshMemoryDatastore()
	bs := FreshLedgerDepot(db)
	assert.EqualValues(t, 0, bs.Foundation())
	assert.EqualValues(t, 0, bs.Altitude())
	assert.EqualValues(t, 0, bs.Extent())

	//
	_, _, err = bs.TrimLedgers(1, status)
	require.Error(t, err)

	_, _, err = bs.TrimLedgers(0, status)
	require.Error(t, err)

	//
	for h := int64(1); h <= 1500; h++ {
		ledger, err := status.CreateLedger(h, verify.CreateNTHTrans(h, 10), new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
		require.NoError(t, err)
		fragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
		require.NoError(t, err)
		observedEndorse := createVerifyAddnEndorse(h, committime.Now())
		bs.PersistLedgerUsingExpandedEndorse(ledger, fragmentAssign, observedEndorse)
	}

	assert.EqualValues(t, 1, bs.Foundation())
	assert.EqualValues(t, 1500, bs.Altitude())
	assert.EqualValues(t, 1500, bs.Extent())

	status.FinalLedgerMoment = time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)
	status.FinalLedgerAltitude = 1500

	status.AgreementSettings.Proof.MaximumLifespanCountLedgers = 400
	status.AgreementSettings.Proof.MaximumLifespanInterval = 1 * time.Second

	//
	trimmed, proofPreserveAltitude, err := bs.TrimLedgers(1200, status)
	require.NoError(t, err)
	assert.EqualValues(t, 1199, trimmed)
	assert.EqualValues(t, 1200, bs.Foundation())
	assert.EqualValues(t, 1500, bs.Altitude())
	assert.EqualValues(t, 301, bs.Extent())
	assert.EqualValues(t, 1100, proofPreserveAltitude)

	require.NotNil(t, bs.FetchLedger(1200))
	require.Nil(t, bs.FetchLedger(1199))

	//
	//
	require.NotNil(t, bs.FetchLedgerSummary(1100))
	require.Nil(t, bs.FetchLedgerSummary(1099))
	require.NotNil(t, bs.FetchLedgerEndorse(1100))
	require.NotNil(t, bs.FetchLedgerExpandedEndorse(1100))
	//
	require.Nil(t, bs.FetchLedgerEndorse(1099))
	require.Nil(t, bs.FetchLedgerExpandedEndorse(1099))

	for i := int64(1); i < 1200; i++ {
		require.Nil(t, bs.FetchLedger(i))
	}
	for i := int64(1200); i <= 1500; i++ {
		require.NotNil(t, bs.FetchLedger(i))
	}

	//
	_, _, err = bs.TrimLedgers(1199, status)
	require.Error(t, err)

	//
	trimmed, _, err = bs.TrimLedgers(1200, status)
	require.NoError(t, err)
	assert.EqualValues(t, 0, trimmed)

	//
	trimmed, _, err = bs.TrimLedgers(1300, status)
	require.NoError(t, err)
	assert.EqualValues(t, 100, trimmed)
	assert.EqualValues(t, 1300, bs.Foundation())

	//
	//
	require.NotNil(t, bs.FetchLedgerSummary(1100))
	require.Nil(t, bs.FetchLedgerSummary(1099))
	require.NotNil(t, bs.FetchLedgerEndorse(1100))
	require.NotNil(t, bs.FetchLedgerExpandedEndorse(1100))
	//
	require.Nil(t, bs.FetchLedgerEndorse(1099))
	require.Nil(t, bs.FetchLedgerExpandedEndorse(1099))

	//
	_, _, err = bs.TrimLedgers(1501, status)
	require.Error(t, err)

	//
	trimmed, _, err = bs.TrimLedgers(1500, status)
	require.NoError(t, err)
	assert.EqualValues(t, 200, trimmed)
	assert.Nil(t, bs.FetchLedger(1499))
	assert.NotNil(t, bs.FetchLedger(1500))
	assert.Nil(t, bs.FetchLedger(1501))
}

func VerifyFetchLedgerSummary(t *testing.T) {
	bs, db := freshInsideRamLedgerDepot()
	altitude := int64(10)
	fetchSummary := func() (any, error) {
		summary := bs.FetchLedgerSummary(altitude)
		return summary, nil
	}

	//
	//
	res, _, alarmFault := conductProc(fetchSummary)
	require.Nil(t, alarmFault, "REDACTED")
	require.Nil(t, res, "REDACTED")

	//
	err := db.Set(reckonLedgerSummaryToken(altitude), []byte("REDACTED"))
	require.NoError(t, err)
	res, _, alarmFault = conductProc(fetchSummary)
	require.NotNil(t, alarmFault, "REDACTED")
	require.Contains(t, alarmFault.Error(), "REDACTED")

	//
	summary := &kinds.LedgerSummary{Heading: kinds.Heading{
		Edition: strongmindedition.Agreement{
			Ledger: edition.LedgerScheme, App: 0,
		}, Altitude: 1, NominatorLocation: commitrand.Octets(security.LocatorExtent),
	}}
	pbm := summary.TowardSchema()
	err = db.Set(reckonLedgerSummaryToken(altitude), shouldEncode(pbm))
	require.NoError(t, err)
	attainedSummary, _, alarmFault := conductProc(fetchSummary)
	require.Nil(t, alarmFault, "REDACTED")
	require.Nil(t, res, "REDACTED")
	protosummary := summary.TowardSchema()
	if genesisummary, ok := attainedSummary.(*kinds.LedgerSummary); ok {
		protoattainedSummary := genesisummary.TowardSchema()
		require.Equal(t, shouldEncode(protosummary), shouldEncode(protoattainedSummary),
			"REDACTED")
	}
}

func VerifyFetchLedgerSummaryViaDigest(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	statusDepot := sm.FreshDepot(dbm.FreshMemoryDatastore(), sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status, err := statusDepot.FetchOriginatingDatastoreEitherInaugurationRecord(settings.InaugurationRecord())
	require.NoError(t, err)
	bs := FreshLedgerDepot(dbm.FreshMemoryDatastore())

	b1, err := status.CreateLedger(status.FinalLedgerAltitude+1, verify.CreateNTHTrans(status.FinalLedgerAltitude+1, 10), new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
	require.NoError(t, err)
	fragmentAssign, err := b1.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	observedEndorse := createVerifyAddnEndorse(1, committime.Now())
	bs.PersistLedger(b1, fragmentAssign, observedEndorse.TowardEndorse())

	foundationLedger := bs.FetchLedgerSummaryViaDigest(b1.Digest())
	assert.EqualValues(t, b1.Altitude, foundationLedger.Heading.Altitude)
	assert.EqualValues(t, b1.FinalLedgerUUID, foundationLedger.Heading.FinalLedgerUUID)
	assert.EqualValues(t, b1.SuccessionUUID, foundationLedger.Heading.SuccessionUUID)
}

func VerifyLedgerAcquireLocatedAltitude(t *testing.T) {
	status, bs, sanitize := createStatusAlsoLedgerDepot()
	defer sanitize()
	require.Equal(t, bs.Altitude(), int64(0), "REDACTED")
	ledger, err := status.CreateLedger(bs.Altitude()+1, nil, new(kinds.Endorse), nil, status.Assessors.ObtainNominator().Location)
	require.NoError(t, err)

	fragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	observedEndorse := createVerifyAddnEndorse(ledger.Altitude, committime.Now())
	bs.PersistLedgerUsingExpandedEndorse(ledger, fragmentAssign, observedEndorse)
	require.Equal(t, bs.Altitude(), ledger.Altitude, "REDACTED")

	ledgerLocatedAltitude := bs.FetchLedger(bs.Altitude())
	b1, err := ledger.TowardSchema()
	require.NoError(t, err)
	b2, err := ledgerLocatedAltitude.TowardSchema()
	require.NoError(t, err)
	bz1 := shouldEncode(b1)
	bz2 := shouldEncode(b2)
	require.Equal(t, bz1, bz2)
	require.Equal(t, ledger.Digest(), ledgerLocatedAltitude.Digest(),
		"REDACTED")

	ledgerLocatedAltitudeAdd1 := bs.FetchLedger(bs.Altitude() + 1)
	require.Nil(t, ledgerLocatedAltitudeAdd1, "REDACTED")
	ledgerLocatedAltitudeAdd2 := bs.FetchLedger(bs.Altitude() + 2)
	require.Nil(t, ledgerLocatedAltitudeAdd2, "REDACTED")
}

func conductProc(fn func() (any, error)) (res any, err error, alarmFault error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				alarmFault = e
			case string:
				alarmFault = fmt.Errorf("REDACTED", e)
			default:
				if st, ok := r.(fmt.Stringer); ok {
					alarmFault = fmt.Errorf("REDACTED", st)
				} else {
					alarmFault = fmt.Errorf("REDACTED", debug.Stack())
				}
			}
		}
	}()

	res, err = fn()
	return res, err, alarmFault
}

func freshLedger(hdr kinds.Heading, finalEndorse *kinds.Endorse) *kinds.Ledger {
	return &kinds.Ledger{
		Heading:     hdr,
		FinalEndorse: finalEndorse,
	}
}
