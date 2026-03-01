package agile_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	mocknode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier/simulate"
	dbs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot/db"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	successionUUID = verify.FallbackVerifySuccessionUUID
)

var (
	ctx      = context.Background()
	tokens     = producePrivateTokens(4)
	values     = tokens.TowardAssessors(20, 10)
	byteMoment, _ = time.Parse(time.RFC3339, "REDACTED")
	h1       = tokens.ProduceNotatedHeadline(successionUUID, 1, byteMoment, nil, values, values,
		digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))
	//
	items2 = values.DuplicateAdvanceNominatorUrgency(1)
	h2    = tokens.ProduceNotatedHeadlineFinalLedgerUUID(successionUUID, 2, byteMoment.Add(30*time.Minute), nil, items2, items2,
		digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens), kinds.LedgerUUID{Digest: h1.Digest()})
	//
	items3 = items2.DuplicateAdvanceNominatorUrgency(1)
	h3    = tokens.ProduceNotatedHeadlineFinalLedgerUUID(successionUUID, 3, byteMoment.Add(1*time.Hour), nil, items3, items3,
		digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens), kinds.LedgerUUID{Digest: h2.Digest()})
	relianceCycle  = 4 * time.Hour
	relianceChoices = agile.RelianceChoices{
		Cycle: 4 * time.Hour,
		Altitude: 1,
		Digest:   h1.Digest(),
	}
	itemAssign = map[int64]*kinds.AssessorAssign{
		1: values,
		2: items2,
		3: items3,
		4: values.DuplicateAdvanceNominatorUrgency(1),
	}
	headlineAssign = map[int64]*kinds.NotatedHeading{
		1: h1,
		//
		2: h2,
		//
		3: h3,
	}
	l1       = &kinds.AgileLedger{NotatedHeading: h1, AssessorAssign: values}
	l2       = &kinds.AgileLedger{NotatedHeading: h2, AssessorAssign: items2}
	completePeer = mocknode.New(
		successionUUID,
		headlineAssign,
		itemAssign,
	)
	lifelessPeer      = mocknode.FreshLifelessSimulate(successionUUID)
	ampleCompletePeer = mocknode.New(produceSimulatePeer(successionUUID, 10, 3, 0, byteMoment))
)

func VerifyCertifyRelianceChoices(t *testing.T) {
	verifyScenarios := []struct {
		err bool
		to  agile.RelianceChoices
	}{
		{
			false,
			relianceChoices,
		},
		{
			true,
			agile.RelianceChoices{
				Cycle: -1 * time.Hour,
				Altitude: 1,
				Digest:   h1.Digest(),
			},
		},
		{
			true,
			agile.RelianceChoices{
				Cycle: 1 * time.Hour,
				Altitude: 0,
				Digest:   h1.Digest(),
			},
		},
		{
			true,
			agile.RelianceChoices{
				Cycle: 1 * time.Hour,
				Altitude: 1,
				Digest:   []byte("REDACTED"),
			},
		},
	}

	for _, tc := range verifyScenarios {
		err := tc.to.CertifyFundamental()
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func VerifySimulate(t *testing.T) {
	l, _ := completePeer.AgileLedger(ctx, 3)
	assert.Equal(t, int64(3), l.Altitude)
}

func Verifycustomer_Orderedvalidation(t *testing.T) {
	freshTokens := producePrivateTokens(4)
	freshValues := freshTokens.TowardAssessors(10, 1)
	distinctValues, _ := kinds.ArbitraryAssessorAssign(10, 100)

	verifyScenarios := []struct {
		alias         string
		anotherHeadings map[int64]*kinds.NotatedHeading //
		values         map[int64]*kinds.AssessorAssign
		initializeFault      bool
		validateFault    bool
	}{
		{
			"REDACTED",
			headlineAssign,
			itemAssign,
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				//
				1: tokens.ProduceNotatedHeadline(successionUUID, 1, byteMoment.Add(1*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			},
			map[int64]*kinds.AssessorAssign{
				1: values,
			},
			true,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{},
			map[int64]*kinds.AssessorAssign{
				1: distinctValues,
			},
			true,
			true,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				1: h1,
			},
			map[int64]*kinds.AssessorAssign{
				1: distinctValues,
			},
			true,
			true,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				//
				1: h1,
				//
				2: tokens.ProduceNotatedHeadline(successionUUID, 2, byteMoment.Add(1*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(tokens)-1, len(tokens)),
				//
				3: tokens.ProduceNotatedHeadline(successionUUID, 3, byteMoment.Add(2*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
			},
			itemAssign,
			false,
			true,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				//
				1: h1,
				//
				2: tokens.ProduceNotatedHeadline(successionUUID, 2, byteMoment.Add(1*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
				//
				3: tokens.ProduceNotatedHeadline(successionUUID, 3, byteMoment.Add(2*time.Hour), nil, values, values,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), len(tokens)-1, len(tokens)),
			},
			itemAssign,
			false,
			true,
		},
		{
			"REDACTED",
			headlineAssign,
			map[int64]*kinds.AssessorAssign{
				1: values,
				2: values,
				3: freshValues,
			},
			false,
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.alias, func(t *testing.T) {
			c, err := agile.FreshCustomer(
				ctx,
				successionUUID,
				relianceChoices,
				mocknode.New(
					successionUUID,
					tc.anotherHeadings,
					tc.values,
				),
				[]supplier.Supplier{mocknode.New(
					successionUUID,
					tc.anotherHeadings,
					tc.values,
				)},
				dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
				agile.OrderedValidation(),
				agile.Tracer(log.VerifyingTracer()),
			)

			if tc.initializeFault {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 3, byteMoment.Add(3*time.Hour))
			if tc.validateFault {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Verifycustomer_Omittingvalidation(t *testing.T) {
	//
	freshTokens := producePrivateTokens(4)
	freshValues := freshTokens.TowardAssessors(10, 1)

	//
	passageTokens := tokens.Broaden(3)
	passageValues := passageTokens.TowardAssessors(10, 1)

	verifyScenarios := []struct {
		alias         string
		anotherHeadings map[int64]*kinds.NotatedHeading //
		values         map[int64]*kinds.AssessorAssign
		initializeFault      bool
		validateFault    bool
	}{
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				//
				1: h1,
				//
				3: h3,
			},
			itemAssign,
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				//
				1: h1,
				3: passageTokens.ProduceNotatedHeadline(successionUUID, 3, byteMoment.Add(2*time.Hour), nil, passageValues, passageValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(passageTokens)),
			},
			map[int64]*kinds.AssessorAssign{
				1: values,
				2: values,
				3: passageValues,
			},
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				//
				1: h1,
				//
				2: tokens.ProduceNotatedHeadline(successionUUID, 2, byteMoment.Add(1*time.Hour), nil, values, freshValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
				//
				3: freshTokens.ProduceNotatedHeadline(successionUUID, 3, byteMoment.Add(2*time.Hour), nil, freshValues, freshValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(freshTokens)),
			},
			map[int64]*kinds.AssessorAssign{
				1: values,
				2: values,
				3: freshValues,
			},
			false,
			false,
		},
		{
			"REDACTED",
			map[int64]*kinds.NotatedHeading{
				//
				1: h1,
				//
				2: tokens.ProduceNotatedHeadline(successionUUID, 2, byteMoment.Add(1*time.Hour), nil, values, freshValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, 0),
				//
				3: freshTokens.ProduceNotatedHeadline(successionUUID, 3, byteMoment.Add(2*time.Hour), nil, freshValues, freshValues,
					digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(freshTokens)),
			},
			map[int64]*kinds.AssessorAssign{
				1: values,
				2: values,
				3: freshValues,
			},
			false,
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.alias, func(t *testing.T) {
			c, err := agile.FreshCustomer(
				ctx,
				successionUUID,
				relianceChoices,
				mocknode.New(
					successionUUID,
					tc.anotherHeadings,
					tc.values,
				),
				[]supplier.Supplier{mocknode.New(
					successionUUID,
					tc.anotherHeadings,
					tc.values,
				)},
				dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
				agile.OmittingValidation(agile.FallbackRelianceStratum),
				agile.Tracer(log.VerifyingTracer()),
			)
			if tc.initializeFault {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 3, byteMoment.Add(3*time.Hour))
			if tc.validateFault {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

//
//
func VerifyCustomerAmplePartitionValidation(t *testing.T) {
	highlyAmpleCompletePeer := mocknode.New(produceSimulatePeer(successionUUID, 100, 3, 0, byteMoment))
	reliableAgileLedger, err := highlyAmpleCompletePeer.AgileLedger(ctx, 5)
	require.NoError(t, err)
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Cycle: 4 * time.Hour,
			Altitude: reliableAgileLedger.Altitude,
			Digest:   reliableAgileLedger.Digest(),
		},
		highlyAmpleCompletePeer,
		[]supplier.Supplier{highlyAmpleCompletePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.OmittingValidation(agile.FallbackRelianceStratum),
	)
	require.NoError(t, err)
	h, err := c.Revise(ctx, byteMoment.Add(100*time.Minute))
	assert.NoError(t, err)
	h2, err := highlyAmpleCompletePeer.AgileLedger(ctx, 100)
	require.NoError(t, err)
	assert.Equal(t, h, h2)
}

func VerifyCustomerPartitionAmongReliableHeadings(t *testing.T) {
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Cycle: 4 * time.Hour,
			Altitude: 1,
			Digest:   h1.Digest(),
		},
		completePeer,
		[]supplier.Supplier{completePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.OmittingValidation(agile.FallbackRelianceStratum),
	)
	require.NoError(t, err)

	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 3, byteMoment.Add(2*time.Hour))
	require.NoError(t, err)

	//
	_, err = c.ReliableAgileLedger(2)
	require.Error(t, err)

	//
	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 2, byteMoment.Add(1*time.Hour))
	assert.NoError(t, err)
}

func Verifycustomer_Sanitize(t *testing.T) {
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		completePeer,
		[]supplier.Supplier{completePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)
	_, err = c.ReliableAgileLedger(1)
	require.NoError(t, err)

	err = c.Sanitize()
	require.NoError(t, err)

	//
	l, err := c.ReliableAgileLedger(1)
	assert.Error(t, err)
	assert.Nil(t, l)
}

//
func VerifyCustomerRecoversReliableHeadlineSubsequentLaunch1(t *testing.T) {
	//
	{
		reliableDepot := dbs.New(dbm.FreshMemoryDatastore(), successionUUID)
		err := reliableDepot.PersistAgileLedger(l1)
		require.NoError(t, err)

		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			relianceChoices,
			completePeer,
			[]supplier.Supplier{completePeer},
			reliableDepot,
			agile.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		l, err := c.ReliableAgileLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), h1.Digest())
		assert.Equal(t, l.AssessorAssign.Digest(), h1.AssessorsDigest.Octets())
	}

	//
	{
		reliableDepot := dbs.New(dbm.FreshMemoryDatastore(), successionUUID)
		err := reliableDepot.PersistAgileLedger(l1)
		require.NoError(t, err)

		//
		heading1 := tokens.ProduceNotatedHeadline(successionUUID, 1, byteMoment.Add(1*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))

		leading := mocknode.New(
			successionUUID,
			map[int64]*kinds.NotatedHeading{
				//
				1: heading1,
			},
			itemAssign,
		)

		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			agile.RelianceChoices{
				Cycle: 4 * time.Hour,
				Altitude: 1,
				Digest:   heading1.Digest(),
			},
			leading,
			[]supplier.Supplier{leading},
			reliableDepot,
			agile.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		l, err := c.ReliableAgileLedger(1)
		assert.NoError(t, err)
		if assert.NotNil(t, l) {
			assert.Equal(t, l.Digest(), heading1.Digest())
			assert.NoError(t, l.CertifyFundamental(successionUUID))
		}
	}
}

//
func VerifyCustomerRecoversReliableHeadlineSubsequentLaunch2(t *testing.T) {
	//
	{
		reliableDepot := dbs.New(dbm.FreshMemoryDatastore(), successionUUID)
		err := reliableDepot.PersistAgileLedger(l1)
		require.NoError(t, err)

		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			agile.RelianceChoices{
				Cycle: 4 * time.Hour,
				Altitude: 2,
				Digest:   h2.Digest(),
			},
			completePeer,
			[]supplier.Supplier{completePeer},
			reliableDepot,
			agile.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ReliableAgileLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), h1.Digest())
		assert.NoError(t, l.CertifyFundamental(successionUUID))
	}

	//
	//
	{
		reliableDepot := dbs.New(dbm.FreshMemoryDatastore(), successionUUID)
		err := reliableDepot.PersistAgileLedger(l1)
		require.NoError(t, err)

		//
		varianceHeading1 := tokens.ProduceNotatedHeadline(successionUUID, 1, byteMoment.Add(1*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))

		varianceHeading2 := tokens.ProduceNotatedHeadline(successionUUID, 2, byteMoment.Add(2*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))

		leading := mocknode.New(
			successionUUID,
			map[int64]*kinds.NotatedHeading{
				1: varianceHeading1,
				2: varianceHeading2,
			},
			itemAssign,
		)

		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			agile.RelianceChoices{
				Cycle: 4 * time.Hour,
				Altitude: 2,
				Digest:   varianceHeading2.Digest(),
			},
			leading,
			[]supplier.Supplier{leading},
			reliableDepot,
			agile.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ReliableAgileLedger(1)
		assert.Error(t, err)
		assert.Nil(t, l)
	}
}

//
func VerifyCustomerRecoversReliableHeadlineSubsequentLaunch3(t *testing.T) {
	//
	{
		//
		reliableDepot := dbs.New(dbm.FreshMemoryDatastore(), successionUUID)
		err := reliableDepot.PersistAgileLedger(l1)
		require.NoError(t, err)

		err = reliableDepot.PersistAgileLedger(l2)
		require.NoError(t, err)

		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			relianceChoices,
			completePeer,
			[]supplier.Supplier{completePeer},
			reliableDepot,
			agile.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ReliableAgileLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), h1.Digest())
		assert.NoError(t, l.CertifyFundamental(successionUUID))

		//
		l, err = c.ReliableAgileLedger(2)
		assert.Error(t, err)
		assert.Nil(t, l)

		l, err = c.ReliableAgileLedger(3)
		assert.Error(t, err)
		assert.Nil(t, l)
	}

	//
	//
	{
		reliableDepot := dbs.New(dbm.FreshMemoryDatastore(), successionUUID)
		err := reliableDepot.PersistAgileLedger(l1)
		require.NoError(t, err)

		//
		heading1 := tokens.ProduceNotatedHeadline(successionUUID, 1, byteMoment.Add(1*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))

		heading2 := tokens.ProduceNotatedHeadline(successionUUID, 2, byteMoment.Add(2*time.Hour), nil, values, values,
			digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens))
		err = reliableDepot.PersistAgileLedger(&kinds.AgileLedger{
			NotatedHeading: heading2,
			AssessorAssign: values,
		})
		require.NoError(t, err)

		leading := mocknode.New(
			successionUUID,
			map[int64]*kinds.NotatedHeading{
				1: heading1,
			},
			itemAssign,
		)

		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			agile.RelianceChoices{
				Cycle: 4 * time.Hour,
				Altitude: 1,
				Digest:   heading1.Digest(),
			},
			leading,
			[]supplier.Supplier{leading},
			reliableDepot,
			agile.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		l, err := c.ReliableAgileLedger(1)
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, l.Digest(), heading1.Digest())
		assert.NoError(t, l.CertifyFundamental(successionUUID))

		//
		l, err = c.ReliableAgileLedger(2)
		assert.Error(t, err)
		assert.Nil(t, l)
	}
}

func Verifycustomer_Revise(t *testing.T) {
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		completePeer,
		[]supplier.Supplier{completePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)

	//
	l, err := c.Revise(ctx, byteMoment.Add(2*time.Hour))
	assert.NoError(t, err)
	if assert.NotNil(t, l) {
		assert.EqualValues(t, 3, l.Altitude)
		assert.NoError(t, l.CertifyFundamental(successionUUID))
	}
}

func Verifycustomer_Parallelism(t *testing.T) {
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		completePeer,
		[]supplier.Supplier{completePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)

	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 2, byteMoment.Add(2*time.Hour))
	require.NoError(t, err)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//
			//

			assert.Equal(t, successionUUID, c.SuccessionUUID())

			_, err := c.FinalReliableAltitude()
			assert.NoError(t, err)

			_, err = c.InitialReliableAltitude()
			assert.NoError(t, err)

			l, err := c.ReliableAgileLedger(1)
			assert.NoError(t, err)
			assert.NotNil(t, l)
		}()
	}

	wg.Wait()
}

func VerifyCustomerSwapsLeadingUsingAttestorConditionalLeadingEqualsInaccessible(t *testing.T) {
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		lifelessPeer,
		[]supplier.Supplier{completePeer, completePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumReissueEndeavors(1),
	)

	require.NoError(t, err)
	_, err = c.Revise(ctx, byteMoment.Add(2*time.Hour))
	require.NoError(t, err)

	assert.NotEqual(t, c.Leading(), lifelessPeer)
	assert.Equal(t, 2, len(c.Attestors()))
}

func Verifycustomer_Reversevalidation(t *testing.T) {
	{
		relianceHeadline, _ := ampleCompletePeer.AgileLedger(ctx, 6)
		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			agile.RelianceChoices{
				Cycle: 4 * time.Minute,
				Altitude: relianceHeadline.Altitude,
				Digest:   relianceHeadline.Digest(),
			},
			ampleCompletePeer,
			[]supplier.Supplier{ampleCompletePeer},
			dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
			agile.Tracer(log.VerifyingTracer()),
		)
		require.NoError(t, err)

		//
		h, err := c.ValidateAgileLedgerLocatedAltitude(ctx, 5, byteMoment.Add(6*time.Minute))
		require.NoError(t, err)
		if assert.NotNil(t, h) {
			assert.EqualValues(t, 5, h.Altitude)
		}

		//
		h, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 3, byteMoment.Add(8*time.Minute))
		assert.NoError(t, err)
		assert.NotNil(t, h)

		//
		h, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 5, byteMoment.Add(6*time.Minute))
		assert.NoError(t, err)
		assert.NotNil(t, h)

		//
		_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 9, byteMoment.Add(9*time.Minute))
		require.NoError(t, err)

		//
		_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 7, byteMoment.Add(9*time.Minute))
		assert.NoError(t, err)
		//
		_, err = c.ReliableAgileLedger(8)
		assert.Error(t, err)

		//
		//
		_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 8, byteMoment.Add(12*time.Minute))
		assert.Error(t, err)

	}
	{
		verifyScenarios := []struct {
			supplier supplier.Supplier
		}{
			{
				//
				mocknode.New(
					successionUUID,
					map[int64]*kinds.NotatedHeading{
						1: h1,
						2: tokens.ProduceNotatedHeadline(successionUUID, 1, byteMoment.Add(30*time.Minute), nil, values, values,
							digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
						3: h3,
					},
					itemAssign,
				),
			},
			{
				//
				mocknode.New(
					successionUUID,
					map[int64]*kinds.NotatedHeading{
						1: h1,
						2: tokens.ProduceNotatedHeadline(successionUUID, 2, byteMoment.Add(30*time.Minute), nil, values, values,
							digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(tokens)),
						3: h3,
					},
					itemAssign,
				),
			},
		}

		for idx, tc := range verifyScenarios {
			c, err := agile.FreshCustomer(
				ctx,
				successionUUID,
				agile.RelianceChoices{
					Cycle: 1 * time.Hour,
					Altitude: 3,
					Digest:   h3.Digest(),
				},
				tc.supplier,
				[]supplier.Supplier{tc.supplier},
				dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
				agile.Tracer(log.VerifyingTracer()),
			)
			require.NoError(t, err, idx)

			_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 2, byteMoment.Add(1*time.Hour).Add(1*time.Second))
			assert.Error(t, err, idx)
		}
	}
}

func Verifycustomer_Freshcustomerfromreliance(t *testing.T) {
	//
	db := dbs.New(dbm.FreshMemoryDatastore(), successionUUID)
	err := db.PersistAgileLedger(l1)
	require.NoError(t, err)

	c, err := agile.FreshCustomerOriginatingReliableDepot(
		successionUUID,
		relianceCycle,
		lifelessPeer,
		[]supplier.Supplier{lifelessPeer},
		db,
	)
	require.NoError(t, err)

	//
	//
	h, err := c.ReliableAgileLedger(1)
	assert.NoError(t, err)
	assert.EqualValues(t, l1.Altitude, h.Altitude)
}

func VerifyCustomerDeletesAttestorConditionalThatRelaysWeImpreciseHeadline(t *testing.T) {
	//
	flawedSource1 := mocknode.New(
		successionUUID,
		map[int64]*kinds.NotatedHeading{
			1: h1,
			2: tokens.ProduceNotatedHeadlineFinalLedgerUUID(successionUUID, 2, byteMoment.Add(30*time.Minute), nil, items2, items2,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"),
				len(tokens), len(tokens), kinds.LedgerUUID{Digest: h1.Digest()}),
		},
		map[int64]*kinds.AssessorAssign{
			1: values,
			2: items2,
		},
	)
	//
	flawedSource2 := mocknode.New(
		successionUUID,
		map[int64]*kinds.NotatedHeading{
			1: h1,
			2: h2,
		},
		map[int64]*kinds.AssessorAssign{
			1: values,
			2: items2,
		},
	)

	lb1, _ := flawedSource1.AgileLedger(ctx, 2)
	require.NotEqual(t, lb1.Digest(), l1.Digest())

	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		completePeer,
		[]supplier.Supplier{flawedSource1, flawedSource2},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumReissueEndeavors(1),
	)
	//
	require.NoError(t, err)
	assert.EqualValues(t, 2, len(c.Attestors()))

	//
	l, err := c.ValidateAgileLedgerLocatedAltitude(ctx, 2, byteMoment.Add(2*time.Hour))
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(c.Attestors()))
	//
	assert.EqualValues(t, 2, l.Altitude)

	//
	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 3, byteMoment.Add(2*time.Hour))
	if assert.Error(t, err) {
		assert.Equal(t, agile.FaultUnsuccessfulHeadlineIntersectAlluding, err)
	}
	//
	assert.EqualValues(t, 1, len(c.Attestors()))
}

func Verifycustomer_Relianceassessors(t *testing.T) {
	distinctValues, _ := kinds.ArbitraryAssessorAssign(10, 100)
	flawedItemAssignPeer := mocknode.New(
		successionUUID,
		map[int64]*kinds.NotatedHeading{
			1: h1,
			//
			//
			2: tokens.ProduceNotatedHeadlineFinalLedgerUUID(successionUUID, 2, byteMoment.Add(30*time.Minute), nil, values, values,
				digest("REDACTED"), digest("REDACTED"), digest("REDACTED"),
				0, len(tokens), kinds.LedgerUUID{Digest: h1.Digest()}),
			3: h3,
		},
		map[int64]*kinds.AssessorAssign{
			1: values,
			2: distinctValues,
			3: distinctValues,
		},
	)

	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		completePeer,
		[]supplier.Supplier{flawedItemAssignPeer, completePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)
	assert.Equal(t, 2, len(c.Attestors()))

	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 2, byteMoment.Add(2*time.Hour).Add(1*time.Second))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(c.Attestors()))
}

func VerifyCustomerShavesHeadingsAlsoAssessorGroupings(t *testing.T) {
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		completePeer,
		[]supplier.Supplier{completePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.ThinningExtent(1),
	)
	require.NoError(t, err)
	_, err = c.ReliableAgileLedger(1)
	require.NoError(t, err)

	h, err := c.Revise(ctx, byteMoment.Add(2*time.Hour))
	require.NoError(t, err)
	require.Equal(t, int64(3), h.Altitude)

	_, err = c.ReliableAgileLedger(1)
	assert.Error(t, err)
}

func VerifyCustomerAssureSoundHeadingsAlsoItemGroupings(t *testing.T) {
	blankItemAssign := &kinds.AssessorAssign{
		Assessors: nil,
		Nominator:   nil,
	}

	verifyScenarios := []struct {
		headings map[int64]*kinds.NotatedHeading
		values    map[int64]*kinds.AssessorAssign
		err     bool
	}{
		{
			headlineAssign,
			itemAssign,
			false,
		},
		{
			headlineAssign,
			map[int64]*kinds.AssessorAssign{
				1: values,
				2: values,
				3: nil,
			},
			true,
		},
		{
			map[int64]*kinds.NotatedHeading{
				1: h1,
				2: h2,
				3: nil,
			},
			itemAssign,
			true,
		},
		{
			headlineAssign,
			map[int64]*kinds.AssessorAssign{
				1: values,
				2: values,
				3: blankItemAssign,
			},
			true,
		},
	}

	for _, tc := range verifyScenarios {
		flawedPeer := mocknode.New(
			successionUUID,
			tc.headings,
			tc.values,
		)
		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			relianceChoices,
			flawedPeer,
			[]supplier.Supplier{flawedPeer, flawedPeer},
			dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
			agile.MaximumReissueEndeavors(1),
		)
		require.NoError(t, err)

		_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 3, byteMoment.Add(2*time.Hour))
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func VerifyCustomerOverseesScopes(t *testing.T) {
	p := mocknode.New(produceSimulatePeer(successionUUID, 100, 10, 1, byteMoment))
	produceLedger, err := p.AgileLedger(ctx, 1)
	require.NoError(t, err)

	//
	contextMomentOutput, abort := context.WithTimeout(ctx, 10*time.Millisecond)
	defer abort()
	_, err = agile.FreshCustomer(
		contextMomentOutput,
		successionUUID,
		agile.RelianceChoices{
			Cycle: 24 * time.Hour,
			Altitude: 1,
			Digest:   produceLedger.Digest(),
		},
		p,
		[]supplier.Supplier{p, p},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
	)
	require.Error(t, contextMomentOutput.Err())
	require.Error(t, err)
	require.True(t, errors.Is(err, context.DeadlineExceeded))

	//
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Cycle: 24 * time.Hour,
			Altitude: 1,
			Digest:   produceLedger.Digest(),
		},
		p,
		[]supplier.Supplier{p, p},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
	)
	require.NoError(t, err)

	//
	contextMomentOutputLedger, abort := context.WithTimeout(ctx, 10*time.Millisecond)
	defer abort()
	_, err = c.ValidateAgileLedgerLocatedAltitude(contextMomentOutputLedger, 100, byteMoment.Add(100*time.Minute))
	require.Error(t, contextMomentOutputLedger.Err())
	require.Error(t, err)
	require.True(t, errors.Is(err, context.DeadlineExceeded))

	//
	contextAbort, abort := context.WithCancel(ctx)
	defer abort()
	time.AfterFunc(10*time.Millisecond, abort)
	_, err = c.ValidateAgileLedgerLocatedAltitude(contextAbort, 100, byteMoment.Add(100*time.Minute))
	require.Error(t, contextAbort.Err())
	require.Error(t, err)
	require.True(t, errors.Is(err, context.Canceled))
}

//
//
func VerifyCustomerFaultsDistinctNominatorUrgencies(t *testing.T) {
	leading := mocknode.New(
		successionUUID,
		map[int64]*kinds.NotatedHeading{
			1: h1,
			2: h2,
		},
		map[int64]*kinds.AssessorAssign{
			1: values,
			2: items2,
		},
	)
	attestor := mocknode.New(
		successionUUID,
		map[int64]*kinds.NotatedHeading{
			1: h1,
			2: h2,
		},
		map[int64]*kinds.AssessorAssign{
			1: values,
			2: values,
		},
	)

	//
	//
	require.Equal(t, values.Digest(), items2.Digest())
	require.NotEqual(t, values.NominatorUrgencyDigest(), items2.NominatorUrgencyDigest())

	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		relianceChoices,
		completePeer,
		[]supplier.Supplier{leading, attestor},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumReissueEndeavors(1),
	)
	//
	require.NoError(t, err)
	assert.EqualValues(t, 2, len(c.Attestors()))

	//
	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 2, byteMoment.Add(2*time.Hour))
	require.Error(t, err)

	//
	assert.EqualValues(t, 2, len(c.Attestors()))
}
