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

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/intrinsic/verify"
	engineseed "github.com/valkyrieworks/utils/random"
	commitdepot "github.com/valkyrieworks/schema/consensuscore/depot"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
	"github.com/valkyrieworks/release"
)

//
//
type sanitizeFunction func()

//
//
func createVerifyExtensionEndorse(level int64, timestamp time.Time) *kinds.ExpandedEndorse {
	return createVerifyExtensionEndorseWithCountAutographs(level, timestamp, 1)
}

func createVerifyExtensionEndorseWithCountAutographs(level int64, timestamp time.Time, countAutographs int) *kinds.ExpandedEndorse {
	extensionEndorseAutographs := []kinds.ExpandedEndorseSignature{}
	for i := 0; i < countAutographs; i++ {
		extensionEndorseAutographs = append(extensionEndorseAutographs, kinds.ExpandedEndorseSignature{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
				RatifierLocation: engineseed.Octets(vault.LocationVolume),
				Timestamp:        timestamp,
				Autograph:        engineseed.Octets(64),
			},
			AdditionAutograph: []byte("REDACTED"),
		})
	}
	return &kinds.ExpandedEndorse{
		Level: level,
		LedgerUID: kinds.LedgerUID{
			Digest:          vault.CRandomOctets(32),
			SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: vault.CRandomOctets(32), Sum: 2},
		},
		ExpandedEndorsements: extensionEndorseAutographs,
	}
}

func createStatusAndLedgerDepot() (sm.Status, *LedgerDepot, sanitizeFunction) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	//
	//
	ledgerStore := dbm.NewMemoryStore()
	statusStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := statusDepot.ImportFromStoreOrOriginEntry(settings.OriginEntry())
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	return status, NewLedgerDepot(ledgerStore), func() { os.RemoveAll(settings.OriginFolder) }
}

func VerifyImportLedgerDepotStatus(t *testing.T) {
	type ledgerDepotVerify struct {
		verifyLabel string
		bss      *commitdepot.LedgerDepotStatus
		desire     commitdepot.LedgerDepotStatus
	}

	verifyScenarios := []ledgerDepotVerify{
		{
			"REDACTED", &commitdepot.LedgerDepotStatus{Root: 100, Level: 1000},
			commitdepot.LedgerDepotStatus{Root: 100, Level: 1000},
		},
		{"REDACTED", &commitdepot.LedgerDepotStatus{}, commitdepot.LedgerDepotStatus{}},
		{"REDACTED", &commitdepot.LedgerDepotStatus{Level: 1000}, commitdepot.LedgerDepotStatus{Root: 1, Level: 1000}},
	}

	for _, tc := range verifyScenarios {
		db := dbm.NewMemoryStore()
		group := db.NewGroup()
		PersistLedgerDepotStatusSegment(tc.bss, group)
		err := group.RecordAlign()
		require.NoError(t, err)
		retrieveBSJ := ImportLedgerDepotStatus(db)
		assert.Equal(t, tc.desire, retrieveBSJ, "REDACTED", tc.verifyLabel)
		err = group.End()
		require.NoError(t, err)
	}
}

func VerifyNewLedgerDepot(t *testing.T) {
	db := dbm.NewMemoryStore()
	bss := commitdepot.LedgerDepotStatus{Root: 100, Level: 10000}
	bz, _ := proto.Marshal(&bss)
	err := db.Set(ledgerDepotKey, bz)
	require.NoError(t, err)
	bs := NewLedgerDepot(db)
	require.Equal(t, int64(100), bs.Root(), "REDACTED")
	require.Equal(t, int64(10000), bs.Level(), "REDACTED")

	alarmCusers := []struct {
		data    []byte
		desireErr string
	}{
		{[]byte("REDACTED"), "REDACTED"},
		{[]byte("REDACTED"), "REDACTED"},
	}

	for i, tt := range alarmCusers {

		//
		_, _, alarmErr := doFn(func() (any, error) {
			err := db.Set(ledgerDepotKey, tt.data)
			require.NoError(t, err)
			_ = NewLedgerDepot(db)
			return nil, nil
		})
		require.NotNil(t, alarmErr, "REDACTED", i, tt.data)
		assert.Contains(t, fmt.Sprintf("REDACTED", alarmErr), tt.desireErr, "REDACTED", i, tt.data)
	}

	err = db.Set(ledgerDepotKey, []byte{})
	require.NoError(t, err)
	bs = NewLedgerDepot(db)
	assert.Equal(t, bs.Level(), int64(0), "REDACTED")
}

func newInRamLedgerDepot() (*LedgerDepot, dbm.DB) {
	db := dbm.NewMemoryStore()
	return NewLedgerDepot(db), db
}

//

func VerifyLedgerDepotPersistImportLedger(t *testing.T) {
	status, bs, sanitize := createStatusAndLedgerDepot()
	defer sanitize()
	require.Equal(t, bs.Root(), int64(0), "REDACTED")
	require.Equal(t, bs.Level(), int64(0), "REDACTED")

	//
	noLedgerLevels := []int64{0, -1, 100, 1000, 2}
	for i, level := range noLedgerLevels {
		if g := bs.ImportLedger(level); g != nil {
			t.Errorf("REDACTED", i, level)
		}
	}

	//
	txs := []kinds.Tx{make([]byte, kinds.LedgerSegmentVolumeOctets)} //
	ledger, err := status.CreateLedger(bs.Level()+1, txs, new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
	require.NoError(t, err)
	soundSectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	require.GreaterOrEqual(t, soundSectionCollection.Sum(), uint32(2))
	section2 := soundSectionCollection.FetchSegment(1)

	viewedEndorse := createVerifyExtensionEndorse(ledger.Level, engineclock.Now())
	bs.PersistLedgerWithExpandedEndorse(ledger, soundSectionCollection, viewedEndorse)
	require.EqualValues(t, 1, bs.Root(), "REDACTED")
	require.EqualValues(t, ledger.Level, bs.Level(), "REDACTED")

	partialSectionCollection := kinds.NewSegmentCollectionFromHeading(kinds.SegmentAssignHeading{Sum: 2})
	noncontiguousSectionCollection := kinds.NewSegmentCollectionFromHeading(kinds.SegmentAssignHeading{Sum: 0})
	_, err = noncontiguousSectionCollection.AppendSegment(section2)
	require.Error(t, err)

	header1 := kinds.Heading{
		Release:         cometrelease.Agreement{Ledger: release.LedgerProtocol},
		Level:          1,
		LedgerUID:         "REDACTED",
		Time:            engineclock.Now(),
		RecommenderLocation: engineseed.Octets(vault.LocationVolume),
	}

	//

	endorseAtH10 := createVerifyExtensionEndorse(10, engineclock.Now()).ToEndorse()
	records := []struct {
		ledger      *kinds.Ledger
		segments      *kinds.SegmentCollection
		viewedEndorse *kinds.ExpandedEndorse
		desireAlarm  string
		desireErr    bool

		taintLedgerInStore      bool
		taintEndorseInStore     bool
		taintViewedEndorseInStore bool
		purgeEndorseInStore       bool
		purgeViewedEndorseInStore   bool
	}{
		{
			ledger:      newLedger(header1, endorseAtH10),
			segments:      soundSectionCollection,
			viewedEndorse: viewedEndorse,
		},

		{
			ledger:     nil,
			desireAlarm: "REDACTED",
		},

		{
			ledger: newLedger( //
				kinds.Heading{
					Release:         cometrelease.Agreement{Ledger: release.LedgerProtocol},
					Level:          5,
					LedgerUID:         "REDACTED",
					Time:            engineclock.Now(),
					RecommenderLocation: engineseed.Octets(vault.LocationVolume),
				},
				createVerifyExtensionEndorse(5, engineclock.Now()).ToEndorse(),
			),
			segments:      soundSectionCollection,
			viewedEndorse: createVerifyExtensionEndorse(5, engineclock.Now()),
		},

		{
			ledger:      newLedger(header1, endorseAtH10),
			segments:      partialSectionCollection,
			desireAlarm:  "REDACTED", //
			viewedEndorse: createVerifyExtensionEndorse(10, engineclock.Now()),
		},

		{
			ledger:             newLedger(header1, endorseAtH10),
			segments:             soundSectionCollection,
			viewedEndorse:        viewedEndorse,
			taintEndorseInStore: true, //
			desireAlarm:         "REDACTED",
		},

		{
			ledger:            newLedger(header1, endorseAtH10),
			segments:            soundSectionCollection,
			viewedEndorse:       viewedEndorse,
			desireAlarm:        "REDACTED",
			taintLedgerInStore: true, //
		},

		{
			ledger:      newLedger(header1, endorseAtH10),
			segments:      soundSectionCollection,
			viewedEndorse: viewedEndorse,

			//
			purgeViewedEndorseInStore: true,
		},

		{
			ledger:      ledger,
			segments:      soundSectionCollection,
			viewedEndorse: viewedEndorse,

			taintViewedEndorseInStore: true,
			desireAlarm:             "REDACTED",
		},

		{
			ledger:      ledger,
			segments:      soundSectionCollection,
			viewedEndorse: viewedEndorse,

			//
			purgeEndorseInStore: true,
		},
	}

	type group struct {
		ledger  *kinds.Ledger
		endorse *kinds.Endorse
		meta   *kinds.LedgerMeta

		viewedEndorse *kinds.Endorse
	}

	for i, record := range records {

		bs, db := newInRamLedgerDepot()
		//
		res, err, alarmErr := doFn(func() (any, error) {
			bs.PersistLedgerWithExpandedEndorse(record.ledger, record.segments, record.viewedEndorse)
			if record.ledger == nil {
				return nil, nil
			}

			if record.taintLedgerInStore {
				err := db.Set(computeLedgerMetaKey(record.ledger.Level), []byte("REDACTED"))
				require.NoError(t, err)
			}
			byteLedger := bs.ImportLedger(record.ledger.Level)
			byteLedgerMeta := bs.ImportLedgerMeta(record.ledger.Level)

			if record.purgeViewedEndorseInStore {
				err := db.Erase(computeViewedEndorseKey(record.ledger.Level))
				require.NoError(t, err)
			}
			if record.taintViewedEndorseInStore {
				err := db.Set(computeViewedEndorseKey(record.ledger.Level), []byte("REDACTED"))
				require.NoError(t, err)
			}
			byteViewedEndorse := bs.ImportViewedEndorse(record.ledger.Level)

			endorseLevel := record.ledger.Level - 1
			if record.purgeEndorseInStore {
				err := db.Erase(computeLedgerEndorseKey(endorseLevel))
				require.NoError(t, err)
			}
			if record.taintEndorseInStore {
				err := db.Set(computeLedgerEndorseKey(endorseLevel), []byte("REDACTED"))
				require.NoError(t, err)
			}
			byteEndorse := bs.ImportLedgerEndorse(endorseLevel)
			return &group{
				ledger: byteLedger, viewedEndorse: byteViewedEndorse, endorse: byteEndorse,
				meta: byteLedgerMeta,
			}, nil
		})

		if subtractStr := record.desireAlarm; subtractStr != "REDACTED" {
			if alarmErr == nil {
				t.Errorf("REDACTED", i)
			} else if got := fmt.Sprintf("REDACTED", alarmErr); !strings.Contains(got, subtractStr) {
				t.Errorf("REDACTED", i, got, subtractStr)
			}
			continue
		}

		if record.desireErr {
			if err == nil {
				t.Errorf("REDACTED", i)
			}
			continue
		}

		assert.Nil(t, alarmErr, "REDACTED", i)
		assert.Nil(t, err, "REDACTED", i)
		qua, ok := res.(*group)
		if !ok || qua == nil {
			t.Errorf("REDACTED", i, res)
			continue
		}
		if record.purgeViewedEndorseInStore {
			assert.Nil(t, qua.viewedEndorse,
				"REDACTED")
		}
		if record.purgeEndorseInStore {
			assert.Nil(t, qua.endorse,
				"REDACTED")
		}
	}
}

//
//
//
func removePlugins(ec *kinds.ExpandedEndorse) bool {
	removed := false
	for idx := range ec.ExpandedEndorsements {
		if len(ec.ExpandedEndorsements[idx].Addition) > 0 || len(ec.ExpandedEndorsements[idx].AdditionAutograph) > 0 {
			removed = true
		}
		ec.ExpandedEndorsements[idx].Addition = nil
		ec.ExpandedEndorsements[idx].AdditionAutograph = nil
	}
	return removed
}

//
//
func VerifyPersistLedgerWithExpandedEndorseAlarmOnMissingAddition(t *testing.T) {
	for _, verifyInstance := range []struct {
		label           string
		distortEndorse func(*kinds.ExpandedEndorse)
		mustAlarm    bool
	}{
		{
			label:           "REDACTED",
			distortEndorse: func(_ *kinds.ExpandedEndorse) {},
			mustAlarm:    false,
		},
		{
			label: "REDACTED",
			distortEndorse: func(c *kinds.ExpandedEndorse) {
				removePlugins(c)
			},
			mustAlarm: true,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			status, bs, sanitize := createStatusAndLedgerDepot()
			defer sanitize()
			h := bs.Level() + 1
			ledger, err := status.CreateLedger(h, verify.CreateNTrans(h, 10), new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
			require.NoError(t, err)

			viewedEndorse := createVerifyExtensionEndorse(ledger.Level, engineclock.Now())
			ps, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
			require.NoError(t, err)
			verifyInstance.distortEndorse(viewedEndorse)
			if verifyInstance.mustAlarm {
				require.Panics(t, func() {
					bs.PersistLedgerWithExpandedEndorse(ledger, ps, viewedEndorse)
				})
			} else {
				bs.PersistLedgerWithExpandedEndorse(ledger, ps, viewedEndorse)
			}
		})
	}
}

//
//
//
func VerifyImportLedgerExpandedEndorse(t *testing.T) {
	for _, verifyInstance := range []struct {
		label         string
		persistExpanded bool
		anticipateOutcome bool
	}{
		{
			label:         "REDACTED",
			persistExpanded: false,
			anticipateOutcome: false,
		},
		{
			label:         "REDACTED",
			persistExpanded: true,
			anticipateOutcome: true,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			status, bs, sanitize := createStatusAndLedgerDepot()
			defer sanitize()
			h := bs.Level() + 1
			ledger, err := status.CreateLedger(h, verify.CreateNTrans(h, 10), new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
			require.NoError(t, err)
			viewedEndorse := createVerifyExtensionEndorse(ledger.Level, engineclock.Now())
			ps, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
			require.NoError(t, err)
			if verifyInstance.persistExpanded {
				bs.PersistLedgerWithExpandedEndorse(ledger, ps, viewedEndorse)
			} else {
				bs.PersistLedger(ledger, ps, viewedEndorse.ToEndorse())
			}
			res := bs.ImportLedgerExpandedEndorse(ledger.Level)
			if verifyInstance.anticipateOutcome {
				require.Equal(t, viewedEndorse, res)
			} else {
				require.Nil(t, res)
			}
		})
	}
}

func VerifyImportRootMeta(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	statusDepot := sm.NewDepot(dbm.NewMemoryStore(), sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := statusDepot.ImportFromStoreOrOriginEntry(settings.OriginEntry())
	require.NoError(t, err)
	bs := NewLedgerDepot(dbm.NewMemoryStore())

	for h := int64(1); h <= 10; h++ {
		ledger, err := status.CreateLedger(h, verify.CreateNTrans(h, 10), new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
		require.NoError(t, err)
		sectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
		require.NoError(t, err)
		viewedEndorse := createVerifyExtensionEndorse(h, engineclock.Now())
		bs.PersistLedgerWithExpandedEndorse(ledger, sectionCollection, viewedEndorse)
	}

	_, _, err = bs.TrimLedgers(4, status)
	require.NoError(t, err)

	rootLedger := bs.ImportRootMeta()
	assert.EqualValues(t, 4, rootLedger.Heading.Level)
	assert.EqualValues(t, 4, bs.Root())

	require.NoError(t, bs.RemoveNewestLedger())
	require.EqualValues(t, 9, bs.Level())
}

func VerifyImportLedgerSection(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")

	bs, db := newInRamLedgerDepot()
	const level, ordinal = 10, 1
	importSection := func() (any, error) {
		segment := bs.ImportLedgerSegment(level, ordinal)
		return segment, nil
	}

	status, err := sm.CreateOriginStatusFromEntry(settings.OriginEntry())
	require.NoError(t, err)

	//
	//
	res, _, alarmErr := doFn(importSection)
	require.Nil(t, alarmErr, "REDACTED")
	require.Nil(t, res, "REDACTED")

	//
	err = db.Set(computeLedgerSectionKey(level, ordinal), []byte("REDACTED"))
	require.NoError(t, err)
	res, _, alarmErr = doFn(importSection)
	require.NotNil(t, alarmErr, "REDACTED")
	require.Contains(t, alarmErr.Error(), "REDACTED")

	//
	ledger, err := status.CreateLedger(level, nil, new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
	require.NoError(t, err)
	sectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	section1 := sectionCollection.FetchSegment(0)

	pb1, err := section1.ToSchema()
	require.NoError(t, err)
	err = db.Set(computeLedgerSectionKey(level, ordinal), shouldSerialize(pb1))
	require.NoError(t, err)
	acquiredSection, _, alarmErr := doFn(importSection)
	require.Nil(t, alarmErr, "REDACTED")
	require.Nil(t, res, "REDACTED")

	//
	acquiredSectionJSON, err := json.Marshal(acquiredSection.(*kinds.Segment))
	require.NoError(t, err)
	section1json, err := json.Marshal(section1)
	require.NoError(t, err)
	require.JSONEq(t, string(acquiredSectionJSON), string(section1json),
		"REDACTED")
}

func VerifyTrimLedgers(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	statusDepot := sm.NewDepot(dbm.NewMemoryStore(), sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := statusDepot.ImportFromStoreOrOriginEntry(settings.OriginEntry())
	require.NoError(t, err)
	db := dbm.NewMemoryStore()
	bs := NewLedgerDepot(db)
	assert.EqualValues(t, 0, bs.Root())
	assert.EqualValues(t, 0, bs.Level())
	assert.EqualValues(t, 0, bs.Volume())

	//
	_, _, err = bs.TrimLedgers(1, status)
	require.Error(t, err)

	_, _, err = bs.TrimLedgers(0, status)
	require.Error(t, err)

	//
	for h := int64(1); h <= 1500; h++ {
		ledger, err := status.CreateLedger(h, verify.CreateNTrans(h, 10), new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
		require.NoError(t, err)
		sectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
		require.NoError(t, err)
		viewedEndorse := createVerifyExtensionEndorse(h, engineclock.Now())
		bs.PersistLedgerWithExpandedEndorse(ledger, sectionCollection, viewedEndorse)
	}

	assert.EqualValues(t, 1, bs.Root())
	assert.EqualValues(t, 1500, bs.Level())
	assert.EqualValues(t, 1500, bs.Volume())

	status.FinalLedgerTime = time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)
	status.FinalLedgerLevel = 1500

	status.AgreementOptions.Proof.MaximumDurationCountLedgers = 400
	status.AgreementOptions.Proof.MaximumDurationPeriod = 1 * time.Second

	//
	trimmed, proofPreserveLevel, err := bs.TrimLedgers(1200, status)
	require.NoError(t, err)
	assert.EqualValues(t, 1199, trimmed)
	assert.EqualValues(t, 1200, bs.Root())
	assert.EqualValues(t, 1500, bs.Level())
	assert.EqualValues(t, 301, bs.Volume())
	assert.EqualValues(t, 1100, proofPreserveLevel)

	require.NotNil(t, bs.ImportLedger(1200))
	require.Nil(t, bs.ImportLedger(1199))

	//
	//
	require.NotNil(t, bs.ImportLedgerMeta(1100))
	require.Nil(t, bs.ImportLedgerMeta(1099))
	require.NotNil(t, bs.ImportLedgerEndorse(1100))
	require.NotNil(t, bs.ImportLedgerExpandedEndorse(1100))
	//
	require.Nil(t, bs.ImportLedgerEndorse(1099))
	require.Nil(t, bs.ImportLedgerExpandedEndorse(1099))

	for i := int64(1); i < 1200; i++ {
		require.Nil(t, bs.ImportLedger(i))
	}
	for i := int64(1200); i <= 1500; i++ {
		require.NotNil(t, bs.ImportLedger(i))
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
	assert.EqualValues(t, 1300, bs.Root())

	//
	//
	require.NotNil(t, bs.ImportLedgerMeta(1100))
	require.Nil(t, bs.ImportLedgerMeta(1099))
	require.NotNil(t, bs.ImportLedgerEndorse(1100))
	require.NotNil(t, bs.ImportLedgerExpandedEndorse(1100))
	//
	require.Nil(t, bs.ImportLedgerEndorse(1099))
	require.Nil(t, bs.ImportLedgerExpandedEndorse(1099))

	//
	_, _, err = bs.TrimLedgers(1501, status)
	require.Error(t, err)

	//
	trimmed, _, err = bs.TrimLedgers(1500, status)
	require.NoError(t, err)
	assert.EqualValues(t, 200, trimmed)
	assert.Nil(t, bs.ImportLedger(1499))
	assert.NotNil(t, bs.ImportLedger(1500))
	assert.Nil(t, bs.ImportLedger(1501))
}

func VerifyImportLedgerMeta(t *testing.T) {
	bs, db := newInRamLedgerDepot()
	level := int64(10)
	importMeta := func() (any, error) {
		meta := bs.ImportLedgerMeta(level)
		return meta, nil
	}

	//
	//
	res, _, alarmErr := doFn(importMeta)
	require.Nil(t, alarmErr, "REDACTED")
	require.Nil(t, res, "REDACTED")

	//
	err := db.Set(computeLedgerMetaKey(level), []byte("REDACTED"))
	require.NoError(t, err)
	res, _, alarmErr = doFn(importMeta)
	require.NotNil(t, alarmErr, "REDACTED")
	require.Contains(t, alarmErr.Error(), "REDACTED")

	//
	meta := &kinds.LedgerMeta{Heading: kinds.Heading{
		Release: cometrelease.Agreement{
			Ledger: release.LedgerProtocol, App: 0,
		}, Level: 1, RecommenderLocation: engineseed.Octets(vault.LocationVolume),
	}}
	pbm := meta.ToSchema()
	err = db.Set(computeLedgerMetaKey(level), shouldSerialize(pbm))
	require.NoError(t, err)
	acquiredMeta, _, alarmErr := doFn(importMeta)
	require.Nil(t, alarmErr, "REDACTED")
	require.Nil(t, res, "REDACTED")
	pbdata := meta.ToSchema()
	if gdata, ok := acquiredMeta.(*kinds.LedgerMeta); ok {
		pbtainedMeta := gdata.ToSchema()
		require.Equal(t, shouldSerialize(pbdata), shouldSerialize(pbtainedMeta),
			"REDACTED")
	}
}

func VerifyImportLedgerMetaByDigest(t *testing.T) {
	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	statusDepot := sm.NewDepot(dbm.NewMemoryStore(), sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := statusDepot.ImportFromStoreOrOriginEntry(settings.OriginEntry())
	require.NoError(t, err)
	bs := NewLedgerDepot(dbm.NewMemoryStore())

	b1, err := status.CreateLedger(status.FinalLedgerLevel+1, verify.CreateNTrans(status.FinalLedgerLevel+1, 10), new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
	require.NoError(t, err)
	sectionCollection, err := b1.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	viewedEndorse := createVerifyExtensionEndorse(1, engineclock.Now())
	bs.PersistLedger(b1, sectionCollection, viewedEndorse.ToEndorse())

	rootLedger := bs.ImportLedgerMetaByDigest(b1.Digest())
	assert.EqualValues(t, b1.Level, rootLedger.Heading.Level)
	assert.EqualValues(t, b1.FinalLedgerUID, rootLedger.Heading.FinalLedgerUID)
	assert.EqualValues(t, b1.LedgerUID, rootLedger.Heading.LedgerUID)
}

func VerifyLedgerAcquireAtLevel(t *testing.T) {
	status, bs, sanitize := createStatusAndLedgerDepot()
	defer sanitize()
	require.Equal(t, bs.Level(), int64(0), "REDACTED")
	ledger, err := status.CreateLedger(bs.Level()+1, nil, new(kinds.Endorse), nil, status.Ratifiers.FetchRecommender().Location)
	require.NoError(t, err)

	sectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	viewedEndorse := createVerifyExtensionEndorse(ledger.Level, engineclock.Now())
	bs.PersistLedgerWithExpandedEndorse(ledger, sectionCollection, viewedEndorse)
	require.Equal(t, bs.Level(), ledger.Level, "REDACTED")

	ledgerAtLevel := bs.ImportLedger(bs.Level())
	b1, err := ledger.ToSchema()
	require.NoError(t, err)
	b2, err := ledgerAtLevel.ToSchema()
	require.NoError(t, err)
	bz1 := shouldSerialize(b1)
	bz2 := shouldSerialize(b2)
	require.Equal(t, bz1, bz2)
	require.Equal(t, ledger.Digest(), ledgerAtLevel.Digest(),
		"REDACTED")

	ledgerAtLevelAdd1 := bs.ImportLedger(bs.Level() + 1)
	require.Nil(t, ledgerAtLevelAdd1, "REDACTED")
	ledgerAtLevelAdd2 := bs.ImportLedger(bs.Level() + 2)
	require.Nil(t, ledgerAtLevelAdd2, "REDACTED")
}

func doFn(fn func() (any, error)) (res any, err error, alarmErr error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				alarmErr = e
			case string:
				alarmErr = fmt.Errorf("REDACTED", e)
			default:
				if st, ok := r.(fmt.Stringer); ok {
					alarmErr = fmt.Errorf("REDACTED", st)
				} else {
					alarmErr = fmt.Errorf("REDACTED", debug.Stack())
				}
			}
		}
	}()

	res, err = fn()
	return res, err, alarmErr
}

func newLedger(hdr kinds.Heading, finalEndorse *kinds.Endorse) *kinds.Ledger {
	return &kinds.Ledger{
		Heading:     hdr,
		FinalEndorse: finalEndorse,
	}
}
