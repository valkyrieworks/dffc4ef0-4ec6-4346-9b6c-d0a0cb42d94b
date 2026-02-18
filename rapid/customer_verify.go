package rapid_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/rapid"
	"github.com/valkyrieworks/rapid/source"
	mocknode "github.com/valkyrieworks/rapid/source/emulate"
	dbs "github.com/valkyrieworks/rapid/depot/db"
	"github.com/valkyrieworks/kinds"
)

const (
	ledgerUID = verify.StandardVerifyLedgerUID
)

var (
	ctx      = context.Background()
	keys     = generatePrivateKeys(4)
	values     = keys.ToRatifiers(20, 10)
	byteTime, _ = time.Parse(time.RFC3339, "REDACTED")
	h1       = keys.GenerateAttestedHeading(ledgerUID, 1, byteTime, nil, values, values,
		digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))
	//
	values2 = values.CloneAugmentRecommenderUrgency(1)
	h2    = keys.GenerateAttestedHeadingFinalLedgerUID(ledgerUID, 2, byteTime.Add(30*time.Minute), nil, values2, values2,
		digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys), kinds.LedgerUID{Digest: h1.Digest()})
	//
	nodes3 = values2.CloneAugmentRecommenderUrgency(1)
	h3    = keys.GenerateAttestedHeadingFinalLedgerUID(ledgerUID, 3, byteTime.Add(1*time.Hour), nil, nodes3, nodes3,
		digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys), kinds.LedgerUID{Digest: h2.Digest()})
	relianceDuration  = 4 * time.Hour
	validateOptions = rapid.ValidateOptions{
		Duration: 4 * time.Hour,
		Level: 1,
		Digest:   h1.Digest(),
	}
	valueCollection = map[int64]*kinds.RatifierAssign{
		1: values,
		2: values2,
		3: nodes3,
		4: values.CloneAugmentRecommenderUrgency(1),
	}
	headingCollection = map[int64]*kinds.AttestedHeading{
		1: h1,
		//
		2: h2,
		//
		3: h3,
	}
	l1       = &kinds.RapidLedger{AttestedHeading: h1, RatifierAssign: values}
	l2       = &kinds.RapidLedger{AttestedHeading: h2, RatifierAssign: values2}
	completeMember = mocknode.New(
		ledgerUID,
		headingCollection,
		valueCollection,
	)
	inactiveMember      = mocknode.NewInactiveEmulate(ledgerUID)
	bulkyCompleteMember = mocknode.New(generateEmulateMember(ledgerUID, 10, 3, 0, byteTime))
)

func VerifyCertifyRelianceSettings(t *testing.T) {
	verifyScenarios := []struct {
		err bool
		to  rapid.ValidateOptions
	}{
		{
			false,
			validateOptions,
		},
		{
			true,
			rapid.ValidateOptions{
				Duration: -1 * time.Hour,
				Level: 1,
				Digest:   h1.Digest(),
			},
		},
		{
			true,
			rapid.ValidateOptions{
				Duration: 1 * time.Hour,
				Level: 0,
				Digest:   h1.Digest(),
			},
		},
		{
			true,
			rapid.ValidateOptions{
				Duration: 1 * time.Hour,
				Level: 1,
				Digest:   []byte("REDACTED"),
			},
		},
	}

	for _, tc := range verifyScenarios {
		err := tc.to.CertifySimple()
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func VerifyEmulate(t *testing.T) {
	l, _ := completeMember.RapidLedger(ctx, 3)
	assert.Equal(t, int64(3), l.Level)
}

func Verifyclient_Orderedvalidation(t *testing.T) {
	newKeys := generatePrivateKeys(4)
	newValues := newKeys.ToRatifiers(10, 1)
	distinctValues, _ := kinds.RandomRatifierCollection(10, 100)

	verifyScenarios := []struct {
		label         string
		anotherHeadings map[int64]*kinds.AttestedHeading //
		values         map[int64]*kinds.RatifierAssign
		initErr      bool
		validateErr    bool
	}{
		{
			"REDACTED",
			headingCollection,
			valueCollection,
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				//
				1: keys.GenerateAttestedHeading(ledgerUID, 1, byteTime.Add(1*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			},
			map[int64]*kinds.RatifierAssign{
				1: values,
			},
			true,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{},
			map[int64]*kinds.RatifierAssign{
				1: distinctValues,
			},
			true,
			true,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				1: h1,
			},
			map[int64]*kinds.RatifierAssign{
				1: distinctValues,
			},
			true,
			true,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				//
				1: h1,
				//
				2: keys.GenerateAttestedHeading(ledgerUID, 2, byteTime.Add(1*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(keys)-1, len(keys)),
				//
				3: keys.GenerateAttestedHeading(ledgerUID, 3, byteTime.Add(2*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
			},
			valueCollection,
			false,
			true,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				//
				1: h1,
				//
				2: keys.GenerateAttestedHeading(ledgerUID, 2, byteTime.Add(1*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
				//
				3: keys.GenerateAttestedHeading(ledgerUID, 3, byteTime.Add(2*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(keys)-1, len(keys)),
			},
			valueCollection,
			false,
			true,
		},
		{
			"REDACTED",
			headingCollection,
			map[int64]*kinds.RatifierAssign{
				1: values,
				2: values,
				3: newValues,
			},
			false,
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.label, func(t *testing.T) {
			c, err := rapid.NewCustomer(
				ctx,
				ledgerUID,
				validateOptions,
				mocknode.New(
					ledgerUID,
					tc.anotherHeadings,
					tc.values,
				),
				[]source.Source{mocknode.New(
					ledgerUID,
					tc.anotherHeadings,
					tc.values,
				)},
				dbs.New(dbm.NewMemoryStore(), ledgerUID),
				rapid.OrderedValidation(),
				rapid.Tracer(log.VerifyingTracer()),
			)

			if tc.initErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			_, err = c.ValidateRapidLedgerAtLevel(ctx, 3, byteTime.Add(3*time.Hour))
			if tc.validateErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Verifyclient_Omittingvalidation(t *testing.T) {
	//
	newKeys := generatePrivateKeys(4)
	newValues := newKeys.ToRatifiers(10, 1)

	//
	passageKeys := keys.Expand(3)
	passageValues := passageKeys.ToRatifiers(10, 1)

	verifyScenarios := []struct {
		label         string
		anotherHeadings map[int64]*kinds.AttestedHeading //
		values         map[int64]*kinds.RatifierAssign
		initErr      bool
		validateErr    bool
	}{
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				//
				1: h1,
				//
				3: h3,
			},
			valueCollection,
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				//
				1: h1,
				3: passageKeys.GenerateAttestedHeading(ledgerUID, 3, byteTime.Add(2*time.Hour), nil, passageValues, passageValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(passageKeys)),
			},
			map[int64]*kinds.RatifierAssign{
				1: values,
				2: values,
				3: passageValues,
			},
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				//
				1: h1,
				//
				2: keys.GenerateAttestedHeading(ledgerUID, 2, byteTime.Add(1*time.Hour), nil, values, newValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
				//
				3: newKeys.GenerateAttestedHeading(ledgerUID, 3, byteTime.Add(2*time.Hour), nil, newValues, newValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(newKeys)),
			},
			map[int64]*kinds.RatifierAssign{
				1: values,
				2: values,
				3: newValues,
			},
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.AttestedHeading{
				//
				1: h1,
				//
				2: keys.GenerateAttestedHeading(ledgerUID, 2, byteTime.Add(1*time.Hour), nil, values, newValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, 0),
				//
				3: newKeys.GenerateAttestedHeading(ledgerUID, 3, byteTime.Add(2*time.Hour), nil, newValues, newValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(newKeys)),
			},
			map[int64]*kinds.RatifierAssign{
				1: values,
				2: values,
				3: newValues,
			},
			false,
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.label, func(t *testing.T) {
			c, err := rapid.NewCustomer(
				ctx,
				ledgerUID,
				validateOptions,
				mocknode.New(
					ledgerUID,
					tc.anotherHeadings,
					tc.values,
				),
				[]source.Source{mocknode.New(
					ledgerUID,
					tc.anotherHeadings,
					tc.values,
				)},
				dbs.New(dbm.NewMemoryStore(), ledgerUID),
				rapid.OmittingValidation(rapid.StandardRelianceLayer),
				rapid.Tracer(log.VerifyingTracer()),
			)
			if tc.initErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			_, err = c.ValidateRapidLedgerAtLevel(ctx, 3, byteTime.Add(3*time.Hour))
			if tc.validateErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

//
//
func VerifyCustomerBulkyDivisionValidation(t *testing.T) {
	highlyBulkyCompleteMember := mocknode.New(generateEmulateMember(ledgerUID, 100, 3, 0, byteTime))
	validatedRapidLedger, err := highlyBulkyCompleteMember.RapidLedger(ctx, 5)
	require.NoError(t, err)
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 4 * time.Hour,
			Level: validatedRapidLedger.Level,
			Digest:   validatedRapidLedger.Digest(),
		},
		highlyBulkyCompleteMember,
		[]source.Source{highlyBulkyCompleteMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.OmittingValidation(rapid.StandardRelianceLayer),
	)
	require.NoError(t, err)
	h, err := c.Modify(ctx, byteTime.Add(100*time.Minute))
	assert.NoError(t, err)
	h2, err := highlyBulkyCompleteMember.RapidLedger(ctx, 100)
	require.NoError(t, err)
	assert.Equal(t, h, h2)
}

func VerifyCustomerDivisionAmongValidatedHeadings(t *testing.T) {
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 4 * time.Hour,
			Level: 1,
			Digest:   h1.Digest(),
		},
		completeMember,
		[]source.Source{completeMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.OmittingValidation(rapid.StandardRelianceLayer),
	)
	require.NoError(t, err)

	_, err = c.ValidateRapidLedgerAtLevel(ctx, 3, byteTime.Add(2*time.Hour))
	require.NoError(t, err)

	//
	_, err = c.ValidatedRapidLedger(2)
	require.Error(t, err)

	//
	_, err = c.ValidateRapidLedgerAtLevel(ctx, 2, byteTime.Add(1*time.Hour))
	assert.NoError(t, err)
}

func Verifyclient_Sanitize(t *testing.T) {
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		completeMember,
		[]source.Source{completeMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)
	_, err = c.ValidatedRapidLedger(1)
	require.NoError(t, err)

	err = c.Sanitize()
	require.NoError(t, err)

	//
	l, err := c.ValidatedRapidLedger(1)
	assert.Error(t, err)
	assert.Nil(t, l)
}

//
func VerifyCustomerRecoversValidatedHeadingAfterActivation1(t *testing.T) {
	//
	{
		validatedDepot := dbs.New(dbm.NewMemoryStore(), ledgerUID)
		err := validatedDepot.PersistRapidLedger(l1)
		require.NoError(t, err)

		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			validateOptions,
			completeMember,
			[]source.Source{completeMember},
			validatedDepot,
			rapid.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		l, err := c.ValidatedRapidLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), h1.Digest())
		assert.Equal(t, l.RatifierAssign.Digest(), h1.RatifiersDigest.Octets())
	}

	//
	{
		validatedDepot := dbs.New(dbm.NewMemoryStore(), ledgerUID)
		err := validatedDepot.PersistRapidLedger(l1)
		require.NoError(t, err)

		//
		header1 := keys.GenerateAttestedHeading(ledgerUID, 1, byteTime.Add(1*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))

		leading := mocknode.New(
			ledgerUID,
			map[int64]*kinds.AttestedHeading{
				//
				1: header1,
			},
			valueCollection,
		)

		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			rapid.ValidateOptions{
				Duration: 4 * time.Hour,
				Level: 1,
				Digest:   header1.Digest(),
			},
			leading,
			[]source.Source{leading},
			validatedDepot,
			rapid.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		l, err := c.ValidatedRapidLedger(1)
		assert.NoError(t, err)
		if assert.NotNil(t, l) {
			assert.Equal(t, l.Digest(), header1.Digest())
			assert.NoError(t, l.CertifySimple(ledgerUID))
		}
	}
}

//
func VerifyCustomerRecoversValidatedHeadingAfterActivation2(t *testing.T) {
	//
	{
		validatedDepot := dbs.New(dbm.NewMemoryStore(), ledgerUID)
		err := validatedDepot.PersistRapidLedger(l1)
		require.NoError(t, err)

		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			rapid.ValidateOptions{
				Duration: 4 * time.Hour,
				Level: 2,
				Digest:   h2.Digest(),
			},
			completeMember,
			[]source.Source{completeMember},
			validatedDepot,
			rapid.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ValidatedRapidLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), h1.Digest())
		assert.NoError(t, l.CertifySimple(ledgerUID))
	}

	//
	//
	{
		validatedDepot := dbs.New(dbm.NewMemoryStore(), ledgerUID)
		err := validatedDepot.PersistRapidLedger(l1)
		require.NoError(t, err)

		//
		varyHeader1 := keys.GenerateAttestedHeading(ledgerUID, 1, byteTime.Add(1*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))

		varyHeader2 := keys.GenerateAttestedHeading(ledgerUID, 2, byteTime.Add(2*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))

		leading := mocknode.New(
			ledgerUID,
			map[int64]*kinds.AttestedHeading{
				1: varyHeader1,
				2: varyHeader2,
			},
			valueCollection,
		)

		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			rapid.ValidateOptions{
				Duration: 4 * time.Hour,
				Level: 2,
				Digest:   varyHeader2.Digest(),
			},
			leading,
			[]source.Source{leading},
			validatedDepot,
			rapid.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ValidatedRapidLedger(1)
		assert.Error(t, err)
		assert.Nil(t, l)
	}
}

//
func VerifyCustomerRecoversValidatedHeadingAfterActivation3(t *testing.T) {
	//
	{
		//
		validatedDepot := dbs.New(dbm.NewMemoryStore(), ledgerUID)
		err := validatedDepot.PersistRapidLedger(l1)
		require.NoError(t, err)

		err = validatedDepot.PersistRapidLedger(l2)
		require.NoError(t, err)

		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			validateOptions,
			completeMember,
			[]source.Source{completeMember},
			validatedDepot,
			rapid.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ValidatedRapidLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), h1.Digest())
		assert.NoError(t, l.CertifySimple(ledgerUID))

		//
		l, err = c.ValidatedRapidLedger(2)
		assert.Error(t, err)
		assert.Nil(t, l)

		l, err = c.ValidatedRapidLedger(3)
		assert.Error(t, err)
		assert.Nil(t, l)
	}

	//
	//
	{
		validatedDepot := dbs.New(dbm.NewMemoryStore(), ledgerUID)
		err := validatedDepot.PersistRapidLedger(l1)
		require.NoError(t, err)

		//
		header1 := keys.GenerateAttestedHeading(ledgerUID, 1, byteTime.Add(1*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))

		header2 := keys.GenerateAttestedHeading(ledgerUID, 2, byteTime.Add(2*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys))
		err = validatedDepot.PersistRapidLedger(&kinds.RapidLedger{
			AttestedHeading: header2,
			RatifierAssign: values,
		})
		require.NoError(t, err)

		leading := mocknode.New(
			ledgerUID,
			map[int64]*kinds.AttestedHeading{
				1: header1,
			},
			valueCollection,
		)

		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			rapid.ValidateOptions{
				Duration: 4 * time.Hour,
				Level: 1,
				Digest:   header1.Digest(),
			},
			leading,
			[]source.Source{leading},
			validatedDepot,
			rapid.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ValidatedRapidLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), header1.Digest())
		assert.NoError(t, l.CertifySimple(ledgerUID))

		//
		l, err = c.ValidatedRapidLedger(2)
		assert.Error(t, err)
		assert.Nil(t, l)
	}
}

func Verifyclient_Modify(t *testing.T) {
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		completeMember,
		[]source.Source{completeMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)

	//
	l, err := c.Modify(ctx, byteTime.Add(2*time.Hour))
	assert.NoError(t, err)
	if assert.NotNil(t, l) {
		assert.EqualValues(t, 3, l.Level)
		assert.NoError(t, l.CertifySimple(ledgerUID))
	}
}

func Verifyclient_Parallelism(t *testing.T) {
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		completeMember,
		[]source.Source{completeMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)

	_, err = c.ValidateRapidLedgerAtLevel(ctx, 2, byteTime.Add(2*time.Hour))
	require.NoError(t, err)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//
			//

			assert.Equal(t, ledgerUID, c.LedgerUID())

			_, err := c.FinalValidatedLevel()
			assert.NoError(t, err)

			_, err = c.InitialValidatedLevel()
			assert.NoError(t, err)

			l, err := c.ValidatedRapidLedger(1)
			assert.NoError(t, err)
			assert.NotNil(t, l)
		}()
	}

	wg.Wait()
}

func VerifyCustomerSubstitutesLeadingWithAttestorIfLeadingIsInaccessible(t *testing.T) {
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		inactiveMember,
		[]source.Source{completeMember, completeMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumReprocessTries(1),
	)

	require.NoError(t, err)
	_, err = c.Modify(ctx, byteTime.Add(2*time.Hour))
	require.NoError(t, err)

	assert.NotEqual(t, c.Leading(), inactiveMember)
	assert.Equal(t, 2, len(c.Attestors()))
}

func Verifyclient_Reversevalidation(t *testing.T) {
	{
		relianceHeading, _ := bulkyCompleteMember.RapidLedger(ctx, 6)
		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			rapid.ValidateOptions{
				Duration: 4 * time.Minute,
				Level: relianceHeading.Level,
				Digest:   relianceHeading.Digest(),
			},
			bulkyCompleteMember,
			[]source.Source{bulkyCompleteMember},
			dbs.New(dbm.NewMemoryStore(), ledgerUID),
			rapid.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		h, err := c.ValidateRapidLedgerAtLevel(ctx, 5, byteTime.Add(6*time.Minute))
		require.NoError(t, err)
		if assert.NotNil(t, h) {
			assert.EqualValues(t, 5, h.Level)
		}

		//
		h, err = c.ValidateRapidLedgerAtLevel(ctx, 3, byteTime.Add(8*time.Minute))
		assert.NoError(t, err)
		assert.NotNil(t, h)

		//
		h, err = c.ValidateRapidLedgerAtLevel(ctx, 5, byteTime.Add(6*time.Minute))
		assert.NoError(t, err)
		assert.NotNil(t, h)

		//
		_, err = c.ValidateRapidLedgerAtLevel(ctx, 9, byteTime.Add(9*time.Minute))
		require.NoError(t, err)

		//
		_, err = c.ValidateRapidLedgerAtLevel(ctx, 7, byteTime.Add(9*time.Minute))
		assert.NoError(t, err)
		//
		_, err = c.ValidatedRapidLedger(8)
		assert.Error(t, err)

		//
		//
		_, err = c.ValidateRapidLedgerAtLevel(ctx, 8, byteTime.Add(12*time.Minute))
		assert.Error(t, err)

	}
	{
		verifyScenarios := []struct {
			source source.Source
		}{
			{
				//
				mocknode.New(
					ledgerUID,
					map[int64]*kinds.AttestedHeading{
						1: h1,
						2: keys.GenerateAttestedHeading(ledgerUID, 1, byteTime.Add(30*time.Minute), nil, values, values,
							digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
						3: h3,
					},
					valueCollection,
				),
			},
			{
				//
				mocknode.New(
					ledgerUID,
					map[int64]*kinds.AttestedHeading{
						1: h1,
						2: keys.GenerateAttestedHeading(ledgerUID, 2, byteTime.Add(30*time.Minute), nil, values, values,
							digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(keys)),
						3: h3,
					},
					valueCollection,
				),
			},
		}

		for idx, tc := range verifyScenarios {
			c, err := rapid.NewCustomer(
				ctx,
				ledgerUID,
				rapid.ValidateOptions{
					Duration: 1 * time.Hour,
					Level: 3,
					Digest:   h3.Digest(),
				},
				tc.source,
				[]source.Source{tc.source},
				dbs.New(dbm.NewMemoryStore(), ledgerUID),
				rapid.Tracer(log.VerifyingTracer()),
			)
			require.NoError(t, err, idx)

			_, err = c.ValidateRapidLedgerAtLevel(ctx, 2, byteTime.Add(1*time.Hour).Add(1*time.Second))
			assert.Error(t, err, idx)
		}
	}
}

func Verifyclient_Newclientfromvalidatedrepository(t *testing.T) {
	//
	db := dbs.New(dbm.NewMemoryStore(), ledgerUID)
	err := db.PersistRapidLedger(l1)
	require.NoError(t, err)

	c, err := rapid.NewCustomerFromValidatedDepot(
		ledgerUID,
		relianceDuration,
		inactiveMember,
		[]source.Source{inactiveMember},
		db,
	)
	require.NoError(t, err)

	//
	//
	h, err := c.ValidatedRapidLedger(1)
	assert.NoError(t, err)
	assert.EqualValues(t, l1.Level, h.Level)
}

func VerifyCustomerDeletesAttestorIfItDispatchesWeInvalidHeading(t *testing.T) {
	//
	flawedSource1 := mocknode.New(
		ledgerUID,
		map[int64]*kinds.AttestedHeading{
			1: h1,
			2: keys.GenerateAttestedHeadingFinalLedgerUID(ledgerUID, 2, byteTime.Add(30*time.Minute), nil, values2, values2,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"),
				len(keys), len(keys), kinds.LedgerUID{Digest: h1.Digest()}),
		},
		map[int64]*kinds.RatifierAssign{
			1: values,
			2: values2,
		},
	)
	//
	flawedSource2 := mocknode.New(
		ledgerUID,
		map[int64]*kinds.AttestedHeading{
			1: h1,
			2: h2,
		},
		map[int64]*kinds.RatifierAssign{
			1: values,
			2: values2,
		},
	)

	lb1, _ := flawedSource1.RapidLedger(ctx, 2)
	require.NotEqual(t, lb1.Digest(), l1.Digest())

	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		completeMember,
		[]source.Source{flawedSource1, flawedSource2},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumReprocessTries(1),
	)
	//
	require.NoError(t, err)
	assert.EqualValues(t, 2, len(c.Attestors()))

	//
	l, err := c.ValidateRapidLedgerAtLevel(ctx, 2, byteTime.Add(2*time.Hour))
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(c.Attestors()))
	//
	assert.EqualValues(t, 2, l.Level)

	//
	_, err = c.ValidateRapidLedgerAtLevel(ctx, 3, byteTime.Add(2*time.Hour))
	if assert.Error(t, err) {
		assert.Equal(t, rapid.ErrErroredHeadingIntersectPointing, err)
	}
	//
	assert.EqualValues(t, 1, len(c.Attestors()))
}

func Verifyclient_Validatedratifiersgroup(t *testing.T) {
	distinctValues, _ := kinds.RandomRatifierCollection(10, 100)
	flawedValueCollectionMember := mocknode.New(
		ledgerUID,
		map[int64]*kinds.AttestedHeading{
			1: h1,
			//
			//
			2: keys.GenerateAttestedHeadingFinalLedgerUID(ledgerUID, 2, byteTime.Add(30*time.Minute), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"),
				0, len(keys), kinds.LedgerUID{Digest: h1.Digest()}),
			3: h3,
		},
		map[int64]*kinds.RatifierAssign{
			1: values,
			2: distinctValues,
			3: distinctValues,
		},
	)

	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		completeMember,
		[]source.Source{flawedValueCollectionMember, completeMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)
	assert.Equal(t, 2, len(c.Attestors()))

	_, err = c.ValidateRapidLedgerAtLevel(ctx, 2, byteTime.Add(2*time.Hour).Add(1*time.Second))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(c.Attestors()))
}

func VerifyCustomerTrimsHeadingsAndRatifierCollections(t *testing.T) {
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		completeMember,
		[]source.Source{completeMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.TrimmingVolume(1),
	)
	require.NoError(t, err)
	_, err = c.ValidatedRapidLedger(1)
	require.NoError(t, err)

	h, err := c.Modify(ctx, byteTime.Add(2*time.Hour))
	require.NoError(t, err)
	require.Equal(t, int64(3), h.Level)

	_, err = c.ValidatedRapidLedger(1)
	assert.Error(t, err)
}

func VerifyCustomerAssureSoundHeadingsAndValueCollections(t *testing.T) {
	emptyValueCollection := &kinds.RatifierAssign{
		Ratifiers: nil,
		Recommender:   nil,
	}

	verifyScenarios := []struct {
		headings map[int64]*kinds.AttestedHeading
		values    map[int64]*kinds.RatifierAssign
		err     bool
	}{
		{
			headingCollection,
			valueCollection,
			false,
		},
		{
			headingCollection,
			map[int64]*kinds.RatifierAssign{
				1: values,
				2: values,
				3: nil,
			},
			true,
		},
		{
			map[int64]*kinds.AttestedHeading{
				1: h1,
				2: h2,
				3: nil,
			},
			valueCollection,
			true,
		},
		{
			headingCollection,
			map[int64]*kinds.RatifierAssign{
				1: values,
				2: values,
				3: emptyValueCollection,
			},
			true,
		},
	}

	for _, tc := range verifyScenarios {
		flawedMember := mocknode.New(
			ledgerUID,
			tc.headings,
			tc.values,
		)
		c, err := rapid.NewCustomer(
			ctx,
			ledgerUID,
			validateOptions,
			flawedMember,
			[]source.Source{flawedMember, flawedMember},
			dbs.New(dbm.NewMemoryStore(), ledgerUID),
			rapid.MaximumReprocessTries(1),
		)
		require.NoError(t, err)

		_, err = c.ValidateRapidLedgerAtLevel(ctx, 3, byteTime.Add(2*time.Hour))
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func VerifyCustomerManagersScopes(t *testing.T) {
	p := mocknode.New(generateEmulateMember(ledgerUID, 100, 10, 1, byteTime))
	generateLedger, err := p.RapidLedger(ctx, 1)
	require.NoError(t, err)

	//
	ctxTimeOut, revoke := context.WithTimeout(ctx, 10*time.Millisecond)
	defer revoke()
	_, err = rapid.NewCustomer(
		ctxTimeOut,
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 24 * time.Hour,
			Level: 1,
			Digest:   generateLedger.Digest(),
		},
		p,
		[]source.Source{p, p},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
	)
	require.Error(t, ctxTimeOut.Err())
	require.Error(t, err)
	require.True(t, errors.Is(err, context.DeadlineExceeded))

	//
	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 24 * time.Hour,
			Level: 1,
			Digest:   generateLedger.Digest(),
		},
		p,
		[]source.Source{p, p},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
	)
	require.NoError(t, err)

	//
	ctxTimeOutLedger, revoke := context.WithTimeout(ctx, 10*time.Millisecond)
	defer revoke()
	_, err = c.ValidateRapidLedgerAtLevel(ctxTimeOutLedger, 100, byteTime.Add(100*time.Minute))
	require.Error(t, ctxTimeOutLedger.Err())
	require.Error(t, err)
	require.True(t, errors.Is(err, context.DeadlineExceeded))

	//
	ctxRevoke, revoke := context.WithCancel(ctx)
	defer revoke()
	time.AfterFunc(10*time.Millisecond, revoke)
	_, err = c.ValidateRapidLedgerAtLevel(ctxRevoke, 100, byteTime.Add(100*time.Minute))
	require.Error(t, ctxRevoke.Err())
	require.Error(t, err)
	require.True(t, errors.Is(err, context.Canceled))
}

//
//
func VerifyCustomerFaultsDistinctRecommenderUrgencies(t *testing.T) {
	leading := mocknode.New(
		ledgerUID,
		map[int64]*kinds.AttestedHeading{
			1: h1,
			2: h2,
		},
		map[int64]*kinds.RatifierAssign{
			1: values,
			2: values2,
		},
	)
	attestor := mocknode.New(
		ledgerUID,
		map[int64]*kinds.AttestedHeading{
			1: h1,
			2: h2,
		},
		map[int64]*kinds.RatifierAssign{
			1: values,
			2: values,
		},
	)

	//
	//
	require.Equal(t, values.Digest(), values2.Digest())
	require.NotEqual(t, values.RecommenderUrgencyDigest(), values2.RecommenderUrgencyDigest())

	c, err := rapid.NewCustomer(
		ctx,
		ledgerUID,
		validateOptions,
		completeMember,
		[]source.Source{leading, attestor},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.MaximumReprocessTries(1),
	)
	//
	require.NoError(t, err)
	assert.EqualValues(t, 2, len(c.Attestors()))

	//
	_, err = c.ValidateRapidLedgerAtLevel(ctx, 2, byteTime.Add(2*time.Hour))
	require.Error(t, err)

	//
	assert.EqualValues(t, 2, len(c.Attestors()))
}
