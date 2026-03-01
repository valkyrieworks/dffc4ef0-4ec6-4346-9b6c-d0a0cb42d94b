package agile_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	mocknode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier/simulate"
	dbs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot/db"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func Verifylightnodeonslaughtattestation_Insane(t *testing.T) {
	//
	var (
		newestAltitude      = int64(10)
		itemExtent           = 5
		deviationAltitude  = int64(6)
		leadingHeadings    = make(map[int64]*kinds.NotatedHeading, newestAltitude)
		leadingAssessors = make(map[int64]*kinds.AssessorAssign, newestAltitude)
	)

	attestorHeadings, attestorAssessors, successionTokens := produceSimulatePeerUsingTokens(successionUUID, newestAltitude, itemExtent, 2, byteMoment)
	attestor := mocknode.New(successionUUID, attestorHeadings, attestorAssessors)
	fabricatedTokens := successionTokens[deviationAltitude-1].AlterationTokens(3) //
	fabricatedValues := fabricatedTokens.TowardAssessors(2, 0)

	for altitude := int64(1); altitude <= newestAltitude; altitude++ {
		if altitude < deviationAltitude {
			leadingHeadings[altitude] = attestorHeadings[altitude]
			leadingAssessors[altitude] = attestorAssessors[altitude]
			continue
		}
		leadingHeadings[altitude] = fabricatedTokens.ProduceNotatedHeadline(successionUUID, altitude, byteMoment.Add(time.Duration(altitude)*time.Minute),
			nil, fabricatedValues, fabricatedValues, digest("REDACTED"), digest("REDACTED"), digest("REDACTED"), 0, len(fabricatedTokens))
		leadingAssessors[altitude] = fabricatedValues
	}
	leading := mocknode.New(successionUUID, leadingHeadings, leadingAssessors)

	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Cycle: 4 * time.Hour,
			Altitude: 1,
			Digest:   leadingHeadings[1].Digest(),
		},
		leading,
		[]supplier.Supplier{attestor},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumReissueEndeavors(1),
	)
	require.NoError(t, err)

	//
	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 10, byteMoment.Add(1*time.Hour))
	if assert.Error(t, err) {
		assert.Equal(t, agile.FaultAgileCustomerOnslaught, err)
	}

	//
	occurenceVersusLeading := &kinds.AgileCustomerOnslaughtProof{
		//
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: leadingHeadings[10],
			AssessorAssign: leadingAssessors[10],
		},
		SharedAltitude: 4,
	}
	assert.True(t, attestor.OwnsProof(occurenceVersusLeading))

	occurenceVersusAttestor := &kinds.AgileCustomerOnslaughtProof{
		//
		//
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: attestorHeadings[7],
			AssessorAssign: attestorAssessors[7],
		},
		SharedAltitude: 4,
	}
	assert.True(t, leading.OwnsProof(occurenceVersusAttestor))
}

func Verifylightnodeonslaughtattestation_Ambiguity(t *testing.T) {
	validationChoices := map[string]agile.Selection{
		"REDACTED": agile.OrderedValidation(),
		"REDACTED":   agile.OmittingValidation(agile.FallbackRelianceStratum),
	}

	for s, validationSelection := range validationChoices {
		t.Log("REDACTED", s)

		//
		var (
			newestAltitude      = int64(10)
			itemExtent           = 5
			deviationAltitude  = int64(6)
			leadingHeadings    = make(map[int64]*kinds.NotatedHeading, newestAltitude)
			leadingAssessors = make(map[int64]*kinds.AssessorAssign, newestAltitude)
		)
		//
		attestorHeadings, attestorAssessors, successionTokens := produceSimulatePeerUsingTokens(successionUUID, newestAltitude+2, itemExtent, 2, byteMoment)
		attestor := mocknode.New(successionUUID, attestorHeadings, attestorAssessors)

		for altitude := int64(1); altitude <= newestAltitude; altitude++ {
			if altitude < deviationAltitude {
				leadingHeadings[altitude] = attestorHeadings[altitude]
				leadingAssessors[altitude] = attestorAssessors[altitude]
				continue
			}
			//
			//
			leadingHeadings[altitude] = successionTokens[altitude].ProduceNotatedHeadline(successionUUID, altitude,
				byteMoment.Add(time.Duration(altitude)*time.Minute), []kinds.Tx{[]byte("REDACTED")},
				attestorAssessors[altitude], attestorAssessors[altitude+1], digest("REDACTED"),
				digest("REDACTED"), digest("REDACTED"), 0, len(successionTokens[altitude])-1)
			leadingAssessors[altitude] = attestorAssessors[altitude]
		}
		leading := mocknode.New(successionUUID, leadingHeadings, leadingAssessors)

		c, err := agile.FreshCustomer(
			ctx,
			successionUUID,
			agile.RelianceChoices{
				Cycle: 4 * time.Hour,
				Altitude: 1,
				Digest:   leadingHeadings[1].Digest(),
			},
			leading,
			[]supplier.Supplier{attestor},
			dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
			agile.Tracer(log.VerifyingTracer()),
			agile.MaximumReissueEndeavors(1),
			validationSelection,
		)
		require.NoError(t, err)

		//
		_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 10, byteMoment.Add(1*time.Hour))
		if assert.Error(t, err) {
			assert.Equal(t, agile.FaultAgileCustomerOnslaught, err)
		}

		//
		//
		//
		occurenceVersusLeading := &kinds.AgileCustomerOnslaughtProof{
			DiscordantLedger: &kinds.AgileLedger{
				NotatedHeading: leadingHeadings[deviationAltitude],
				AssessorAssign: leadingAssessors[deviationAltitude],
			},
			SharedAltitude: deviationAltitude,
		}
		assert.True(t, attestor.OwnsProof(occurenceVersusLeading))

		occurenceVersusAttestor := &kinds.AgileCustomerOnslaughtProof{
			DiscordantLedger: &kinds.AgileLedger{
				NotatedHeading: attestorHeadings[deviationAltitude],
				AssessorAssign: attestorAssessors[deviationAltitude],
			},
			SharedAltitude: deviationAltitude,
		}
		assert.True(t, leading.OwnsProof(occurenceVersusAttestor))
	}
}

func Verifylightnodeonslaughtattestation_Advanceinsane(t *testing.T) {
	//
	//
	var (
		newestAltitude      = int64(10)
		itemExtent           = 5
		fabricatedAltitude      = int64(12)
		attestationAltitude       = int64(11)
		leadingHeadings    = make(map[int64]*kinds.NotatedHeading, fabricatedAltitude)
		leadingAssessors = make(map[int64]*kinds.AssessorAssign, fabricatedAltitude)
	)

	attestorHeadings, attestorAssessors, successionTokens := produceSimulatePeerUsingTokens(successionUUID, newestAltitude, itemExtent, 2, byteMoment)

	//
	//
	for h := range attestorHeadings {
		leadingHeadings[h] = attestorHeadings[h]
		leadingAssessors[h] = attestorAssessors[h]
	}
	fabricatedTokens := successionTokens[newestAltitude].AlterationTokens(3) //
	leadingAssessors[fabricatedAltitude] = fabricatedTokens.TowardAssessors(2, 0)
	leadingHeadings[fabricatedAltitude] = fabricatedTokens.ProduceNotatedHeadline(
		successionUUID,
		fabricatedAltitude,
		byteMoment.Add(time.Duration(newestAltitude+1)*time.Minute), //
		nil,
		leadingAssessors[fabricatedAltitude],
		leadingAssessors[fabricatedAltitude],
		digest("REDACTED"),
		digest("REDACTED"),
		digest("REDACTED"),
		0, len(fabricatedTokens),
	)

	attestor := mocknode.New(successionUUID, attestorHeadings, attestorAssessors)
	leading := mocknode.New(successionUUID, leadingHeadings, leadingAssessors)

	delayingAttestor := attestor.Duplicate(successionUUID)

	//
	//
	accessory := leading

	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Cycle: 4 * time.Hour,
			Altitude: 1,
			Digest:   leadingHeadings[1].Digest(),
		},
		leading,
		[]supplier.Supplier{attestor, accessory},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumTimerDeviation(1*time.Second),
		agile.MaximumLedgerDelay(1*time.Second),
	)
	require.NoError(t, err)

	//
	//
	values := successionTokens[newestAltitude].TowardAssessors(2, 0)
	freshLdgr := &kinds.AgileLedger{
		NotatedHeading: successionTokens[newestAltitude].ProduceNotatedHeadline(
			successionUUID,
			attestationAltitude,
			byteMoment.Add(time.Duration(attestationAltitude+1)*time.Minute), //
			nil,
			values,
			values,
			digest("REDACTED"),
			digest("REDACTED"),
			digest("REDACTED"),
			0, len(successionTokens),
		),
		AssessorAssign: values,
	}
	go func() {
		time.Sleep(2 * time.Second)
		attestor.AppendAgileLedger(freshLdgr)
	}()

	//
	//
	_, err = c.Revise(ctx, byteMoment.Add(time.Duration(fabricatedAltitude)*time.Minute))
	if assert.Error(t, err) {
		assert.Equal(t, agile.FaultAgileCustomerOnslaught, err)
	}

	//
	occurenceVersusLeading := &kinds.AgileCustomerOnslaughtProof{
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: leadingHeadings[fabricatedAltitude],
			AssessorAssign: leadingAssessors[fabricatedAltitude],
		},
		SharedAltitude: newestAltitude,
	}
	assert.True(t, attestor.OwnsProof(occurenceVersusLeading))

	//
	//
	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, fabricatedAltitude, byteMoment.Add(time.Duration(fabricatedAltitude)*time.Minute))
	if assert.Error(t, err) {
		assert.Equal(t, agile.FaultAgileCustomerOnslaught, err)
	}
	assert.True(t, attestor.OwnsProof(occurenceVersusLeading))

	//
	//
	c, err = agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Cycle: 4 * time.Hour,
			Altitude: 1,
			Digest:   leadingHeadings[1].Digest(),
		},
		leading,
		[]supplier.Supplier{delayingAttestor, accessory},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumTimerDeviation(1*time.Second),
		agile.MaximumLedgerDelay(1*time.Second),
	)
	require.NoError(t, err)

	_, err = c.Revise(ctx, byteMoment.Add(time.Duration(fabricatedAltitude)*time.Minute))
	assert.NoError(t, err)
}

//
//
//
func VerifyCustomerDeviatingLogging1(t *testing.T) {
	leading := mocknode.New(produceSimulatePeer(successionUUID, 10, 5, 2, byteMoment))
	initialLedger, err := leading.AgileLedger(ctx, 1)
	require.NoError(t, err)
	attestor := mocknode.New(produceSimulatePeer(successionUUID, 10, 5, 2, byteMoment))

	_, err = agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Altitude: 1,
			Digest:   initialLedger.Digest(),
			Cycle: 4 * time.Hour,
		},
		leading,
		[]supplier.Supplier{attestor},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumReissueEndeavors(1),
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "REDACTED")
}

//
//
func VerifyCustomerDeviatingLogging2(t *testing.T) {
	leading := mocknode.New(produceSimulatePeer(successionUUID, 10, 5, 2, byteMoment))
	initialLedger, err := leading.AgileLedger(ctx, 1)
	require.NoError(t, err)
	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Altitude: 1,
			Digest:   initialLedger.Digest(),
			Cycle: 4 * time.Hour,
		},
		leading,
		[]supplier.Supplier{inactivePeer, inactivePeer, leading},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumReissueEndeavors(1),
	)
	require.NoError(t, err)

	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 10, byteMoment.Add(1*time.Hour))
	assert.NoError(t, err)
	assert.Equal(t, 3, len(c.Attestors()))
}

//
//
func VerifyCustomerDeviatingLogging3(t *testing.T) {
	_, leadingHeadings, leadingValues := produceSimulatePeer(successionUUID, 10, 5, 2, byteMoment)
	leading := mocknode.New(successionUUID, leadingHeadings, leadingValues)

	initialLedger, err := leading.AgileLedger(ctx, 1)
	require.NoError(t, err)

	_, simulateHeadings, simulateValues := produceSimulatePeer(successionUUID, 10, 5, 2, byteMoment)
	simulateHeadings[1] = leadingHeadings[1]
	simulateValues[1] = leadingValues[1]
	attestor := mocknode.New(successionUUID, simulateHeadings, simulateValues)

	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Altitude: 1,
			Digest:   initialLedger.Digest(),
			Cycle: 4 * time.Hour,
		},
		leading,
		[]supplier.Supplier{attestor},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.MaximumReissueEndeavors(1),
	)
	require.NoError(t, err)

	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 10, byteMoment.Add(1*time.Hour))
	assert.Error(t, err)
	assert.Equal(t, 1, len(c.Attestors()))
}

//
//
func VerifyCustomerDeviatingLogging4(t *testing.T) {
	_, leadingHeadings, leadingValues := produceSimulatePeer(successionUUID, 10, 5, 2, byteMoment)
	leading := mocknode.New(successionUUID, leadingHeadings, leadingValues)

	initialLedger, err := leading.AgileLedger(ctx, 1)
	require.NoError(t, err)

	_, simulateHeadings, simulateValues := produceSimulatePeer(successionUUID, 10, 5, 2, byteMoment)
	attestor := leading.Duplicate(successionUUID)
	attestor.AppendAgileLedger(&kinds.AgileLedger{
		NotatedHeading: simulateHeadings[10],
		AssessorAssign: simulateValues[10],
	})

	c, err := agile.FreshCustomer(
		ctx,
		successionUUID,
		agile.RelianceChoices{
			Altitude: 1,
			Digest:   initialLedger.Digest(),
			Cycle: 4 * time.Hour,
		},
		leading,
		[]supplier.Supplier{attestor},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
	)
	require.NoError(t, err)

	_, err = c.ValidateAgileLedgerLocatedAltitude(ctx, 10, byteMoment.Add(1*time.Hour))
	assert.Error(t, err)
	assert.Equal(t, 1, len(c.Attestors()))
}
