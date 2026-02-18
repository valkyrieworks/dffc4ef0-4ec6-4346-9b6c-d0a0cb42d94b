package rapid_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/rapid"
	"github.com/valkyrieworks/rapid/source"
	mocknode "github.com/valkyrieworks/rapid/source/emulate"
	dbs "github.com/valkyrieworks/rapid/depot/db"
	"github.com/valkyrieworks/kinds"
)

func Verifylightnodeassaultevidence_Erratic(t *testing.T) {
	//
	var (
		newestLevel      = int64(10)
		valueVolume           = 5
		deviationLevel  = int64(6)
		leadingHeadings    = make(map[int64]*kinds.AttestedHeading, newestLevel)
		leadingRatifiers = make(map[int64]*kinds.RatifierAssign, newestLevel)
	)

	attestorHeadings, attestorRatifiers, seriesKeys := generateEmulateMemberWithKeys(ledgerUID, newestLevel, valueVolume, 2, byteTime)
	attestor := mocknode.New(ledgerUID, attestorHeadings, attestorRatifiers)
	falsifiedKeys := seriesKeys[deviationLevel-1].AlterKeys(3) //
	falsifiedValues := falsifiedKeys.ToRatifiers(2, 0)

	for level := int64(1); level <= newestLevel; level++ {
		if level < deviationLevel {
			leadingHeadings[level] = attestorHeadings[level]
			leadingRatifiers[level] = attestorRatifiers[level]
			continue
		}
		leadingHeadings[level] = falsifiedKeys.GenerateAttestedHeading(ledgerUID, level, byteTime.Add(time.Duration(level)*time.Minute),
			nil, falsifiedValues, falsifiedValues, digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(falsifiedKeys))
		leadingRatifiers[level] = falsifiedValues
	}
	leading := mocknode.New(ledgerUID, leadingHeadings, leadingRatifiers)

	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 4 * time.Hour,
			Level: 1,
			Digest:   leadingHeadings[1].Digest(),
		},
		leading,
		[]source.Source{attestor},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumReprocessTries(1),
	)
	require.NoError(t, err)

	//
	_, err = c.ValidateRapidLedgerAtLevel(ctx, 10, byteTime.Add(1*time.Hour))
	if assert.Error(t, err) {
		assert.Equal(t, rapid.ErrRapidCustomerAssault, err)
	}

	//
	evtVersusLeading := &kinds.RapidCustomerAssaultProof{
		//
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: leadingHeadings[10],
			RatifierAssign: leadingRatifiers[10],
		},
		SharedLevel: 4,
	}
	assert.True(t, attestor.HasProof(evtVersusLeading))

	evtVersusAttestor := &kinds.RapidCustomerAssaultProof{
		//
		//
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: attestorHeadings[7],
			RatifierAssign: attestorRatifiers[7],
		},
		SharedLevel: 4,
	}
	assert.True(t, leading.HasProof(evtVersusAttestor))
}

func Verifylightnodeassaultevidence_Ambiguity(t *testing.T) {
	validationSettings := map[string]rapid.Setting{
		"REDACTED": rapid.OrderedValidation(),
		"REDACTED":   rapid.OmittingValidation(rapid.StandardRelianceLayer),
	}

	for s, validationSetting := range validationSettings {
		t.Log("REDACTED", s)

		//
		var (
			newestLevel      = int64(10)
			valueVolume           = 5
			deviationLevel  = int64(6)
			leadingHeadings    = make(map[int64]*kinds.AttestedHeading, newestLevel)
			leadingRatifiers = make(map[int64]*kinds.RatifierAssign, newestLevel)
		)
		//
		attestorHeadings, attestorRatifiers, seriesKeys := generateEmulateMemberWithKeys(ledgerUID, newestLevel+2, valueVolume, 2, byteTime)
		attestor := mocknode.New(ledgerUID, attestorHeadings, attestorRatifiers)

		for level := int64(1); level <= newestLevel; level++ {
			if level < deviationLevel {
				leadingHeadings[level] = attestorHeadings[level]
				leadingRatifiers[level] = attestorRatifiers[level]
				continue
			}
			//
			//
			leadingHeadings[level] = seriesKeys[level].GenerateAttestedHeading(ledgerUID, level,
				byteTime.Add(time.Duration(level)*time.Minute), []kinds.Tx{[]byte("REDACTED")},
				attestorRatifiers[level], attestorRatifiers[level+1], digest("REDACTED"),
				digest("REDACTED"), digest("REDACTED"), 0, len(seriesKeys[level])-1)
			leadingRatifiers[level] = attestorRatifiers[level]
		}
		leading := mocknode.New(ledgerUID, leadingHeadings, leadingRatifiers)

		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			rapid.ValidateOptions{
				Duration: 4 * time.Hour,
				Level: 1,
				Digest:   leadingHeadings[1].Digest(),
			},
			leading,
			[]source.Source{attestor},
			dbs.New(dbm.NewMemoryStore(), ledgerUID),
			rapid.Tracer(log.VerifyingTracer()),
			rapid.MaximumReprocessTries(1),
			validationSetting,
		)
		require.NoError(t, err)

		//
		_, err = c.ValidateRapidLedgerAtLevel(ctx, 10, byteTime.Add(1*time.Hour))
		if assert.Error(t, err) {
			assert.Equal(t, rapid.ErrRapidCustomerAssault, err)
		}

		//
		//
		//
		evtVersusLeading := &kinds.RapidCustomerAssaultProof{
			ClashingLedger: &kinds.RapidLedger{
				AttestedHeading: leadingHeadings[deviationLevel],
				RatifierAssign: leadingRatifiers[deviationLevel],
			},
			SharedLevel: deviationLevel,
		}
		assert.True(t, attestor.HasProof(evtVersusLeading))

		evtVersusAttestor := &kinds.RapidCustomerAssaultProof{
			ClashingLedger: &kinds.RapidLedger{
				AttestedHeading: attestorHeadings[deviationLevel],
				RatifierAssign: attestorRatifiers[deviationLevel],
			},
			SharedLevel: deviationLevel,
		}
		assert.True(t, leading.HasProof(evtVersusAttestor))
	}
}

func Verifylightnodeassaultevidence_Advanceerratic(t *testing.T) {
	//
	//
	var (
		newestLevel      = int64(10)
		valueVolume           = 5
		falsifiedLevel      = int64(12)
		evidenceLevel       = int64(11)
		leadingHeadings    = make(map[int64]*kinds.AttestedHeading, falsifiedLevel)
		leadingRatifiers = make(map[int64]*kinds.RatifierAssign, falsifiedLevel)
	)

	attestorHeadings, attestorRatifiers, seriesKeys := generateEmulateMemberWithKeys(ledgerUID, newestLevel, valueVolume, 2, byteTime)

	//
	//
	for h := range attestorHeadings {
		leadingHeadings[h] = attestorHeadings[h]
		leadingRatifiers[h] = attestorRatifiers[h]
	}
	falsifiedKeys := seriesKeys[newestLevel].AlterKeys(3) //
	leadingRatifiers[falsifiedLevel] = falsifiedKeys.ToRatifiers(2, 0)
	leadingHeadings[falsifiedLevel] = falsifiedKeys.GenerateAttestedHeading(
		ledgerUID,
		falsifiedLevel,
		byteTime.Add(time.Duration(newestLevel+1)*time.Minute), //
		nil,
		leadingRatifiers[falsifiedLevel],
		leadingRatifiers[falsifiedLevel],
		digest("REDACTED"),
		digest("REDACTED"),
		digest("REDACTED"),
		0, len(falsifiedKeys),
	)

	attestor := mocknode.New(ledgerUID, attestorHeadings, attestorRatifiers)
	leading := mocknode.New(ledgerUID, leadingHeadings, leadingRatifiers)

	trailingAttestor := attestor.Clone(ledgerUID)

	//
	//
	accessory := leading

	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 4 * time.Hour,
			Level: 1,
			Digest:   leadingHeadings[1].Digest(),
		},
		leading,
		[]source.Source{attestor, accessory},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumTimerDeviation(1*time.Second),
		rapid.MaximumLedgerDelay(1*time.Second),
	)
	require.NoError(t, err)

	//
	//
	values := seriesKeys[newestLevel].ToRatifiers(2, 0)
	newBlock := &kinds.RapidLedger{
		AttestedHeading: seriesKeys[newestLevel].GenerateAttestedHeading(
			ledgerUID,
			evidenceLevel,
			byteTime.Add(time.Duration(evidenceLevel+1)*time.Minute), //
			nil,
			values,
			values,
			digest("REDACTED"),
			digest("REDACTED"),
			digest("REDACTED"),
			0, len(seriesKeys),
		),
		RatifierAssign: values,
	}
	go func() {
		time.Sleep(2 * time.Second)
		attestor.AppendRapidLedger(newBlock)
	}()

	//
	//
	_, err = c.Modify(ctx, byteTime.Add(time.Duration(falsifiedLevel)*time.Minute))
	if assert.Error(t, err) {
		assert.Equal(t, rapid.ErrRapidCustomerAssault, err)
	}

	//
	evtVersusLeading := &kinds.RapidCustomerAssaultProof{
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: leadingHeadings[falsifiedLevel],
			RatifierAssign: leadingRatifiers[falsifiedLevel],
		},
		SharedLevel: newestLevel,
	}
	assert.True(t, attestor.HasProof(evtVersusLeading))

	//
	//
	_, err = c.ValidateRapidLedgerAtLevel(ctx, falsifiedLevel, byteTime.Add(time.Duration(falsifiedLevel)*time.Minute))
	if assert.Error(t, err) {
		assert.Equal(t, rapid.ErrRapidCustomerAssault, err)
	}
	assert.True(t, attestor.HasProof(evtVersusLeading))

	//
	//
	c, err = rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 4 * time.Hour,
			Level: 1,
			Digest:   leadingHeadings[1].Digest(),
		},
		leading,
		[]source.Source{trailingAttestor, accessory},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumTimerDeviation(1*time.Second),
		rapid.MaximumLedgerDelay(1*time.Second),
	)
	require.NoError(t, err)

	_, err = c.Modify(ctx, byteTime.Add(time.Duration(falsifiedLevel)*time.Minute))
	assert.NoError(t, err)
}

//
//
//
func VerifyCustomerDeviantFootprints1(t *testing.T) {
	leading := mocknode.New(generateEmulateMember(ledgerUID, 10, 5, 2, byteTime))
	initialLedger, err := leading.RapidLedger(ctx, 1)
	require.NoError(t, err)
	attestor := mocknode.New(generateEmulateMember(ledgerUID, 10, 5, 2, byteTime))

	_, err = rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Level: 1,
			Digest:   initialLedger.Digest(),
			Duration: 4 * time.Hour,
		},
		leading,
		[]source.Source{attestor},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumReprocessTries(1),
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "REDACTED")
}

//
//
func VerifyCustomerDeviantFootprints2(t *testing.T) {
	leading := mocknode.New(generateEmulateMember(ledgerUID, 10, 5, 2, byteTime))
	initialLedger, err := leading.RapidLedger(ctx, 1)
	require.NoError(t, err)
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Level: 1,
			Digest:   initialLedger.Digest(),
			Duration: 4 * time.Hour,
		},
		leading,
		[]source.Source{inactiveMember, inactiveMember, leading},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumReprocessTries(1),
	)
	require.NoError(t, err)

	_, err = c.ValidateRapidLedgerAtLevel(ctx, 10, byteTime.Add(1*time.Hour))
	assert.NoError(t, err)
	assert.Equal(t, 3, len(c.Attestors()))
}

//
//
func VerifyCustomerDeviantFootprints3(t *testing.T) {
	_, leadingHeadings, leadingValues := generateEmulateMember(ledgerUID, 10, 5, 2, byteTime)
	leading := mocknode.New(ledgerUID, leadingHeadings, leadingValues)

	initialLedger, err := leading.RapidLedger(ctx, 1)
	require.NoError(t, err)

	_, emulateHeadings, emulateValues := generateEmulateMember(ledgerUID, 10, 5, 2, byteTime)
	emulateHeadings[1] = leadingHeadings[1]
	emulateValues[1] = leadingValues[1]
	attestor := mocknode.New(ledgerUID, emulateHeadings, emulateValues)

	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Level: 1,
			Digest:   initialLedger.Digest(),
			Duration: 4 * time.Hour,
		},
		leading,
		[]source.Source{attestor},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumReprocessTries(1),
	)
	require.NoError(t, err)

	_, err = c.ValidateRapidLedgerAtLevel(ctx, 10, byteTime.Add(1*time.Hour))
	assert.Error(t, err)
	assert.Equal(t, 1, len(c.Attestors()))
}

//
//
func VerifyCustomerDeviantFootprints4(t *testing.T) {
	_, leadingHeadings, leadingValues := generateEmulateMember(ledgerUID, 10, 5, 2, byteTime)
	leading := mocknode.New(ledgerUID, leadingHeadings, leadingValues)

	initialLedger, err := leading.RapidLedger(ctx, 1)
	require.NoError(t, err)

	_, emulateHeadings, emulateValues := generateEmulateMember(ledgerUID, 10, 5, 2, byteTime)
	attestor := leading.Clone(ledgerUID)
	attestor.AppendRapidLedger(&kinds.RapidLedger{
		AttestedHeading: emulateHeadings[10],
		RatifierAssign: emulateValues[10],
	})

	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Level: 1,
			Digest:   initialLedger.Digest(),
			Duration: 4 * time.Hour,
		},
		leading,
		[]source.Source{attestor},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)

	_, err = c.ValidateRapidLedgerAtLevel(ctx, 10, byteTime.Add(1*time.Hour))
	assert.Error(t, err)
	assert.Equal(t, 1, len(c.Attestors()))
}
