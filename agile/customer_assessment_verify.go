package agile_test

import (
	"context"
	"testing"
	"time"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	mocknode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier/simulate"
	dbs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot/db"
)

//
//
//
//
//
//
//
var (
	assessmentCompletePeer = mocknode.New(produceSimulatePeer(successionUUID, 1000, 100, 1, byteMoment))
	inaugurationLedger, _   = assessmentCompletePeer.AgileLedger(context.Background(), 1)
)

func AssessmentProgression(b *testing.B) {
	c, err := agile.FreshCustomer(
		context.Background(),
		successionUUID,
		agile.RelianceChoices{
			Cycle: 24 * time.Hour,
			Altitude: 1,
			Digest:   inaugurationLedger.Digest(),
		},
		assessmentCompletePeer,
		[]supplier.Supplier{assessmentCompletePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
		agile.OrderedValidation(),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.ValidateAgileLedgerLocatedAltitude(context.Background(), 1000, byteMoment.Add(1000*time.Minute))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func AssessmentPartition(b *testing.B) {
	c, err := agile.FreshCustomer(
		context.Background(),
		successionUUID,
		agile.RelianceChoices{
			Cycle: 24 * time.Hour,
			Altitude: 1,
			Digest:   inaugurationLedger.Digest(),
		},
		assessmentCompletePeer,
		[]supplier.Supplier{assessmentCompletePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.ValidateAgileLedgerLocatedAltitude(context.Background(), 1000, byteMoment.Add(1000*time.Minute))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func AssessmentReverse(b *testing.B) {
	reliableLedger, _ := assessmentCompletePeer.AgileLedger(context.Background(), 0)
	c, err := agile.FreshCustomer(
		context.Background(),
		successionUUID,
		agile.RelianceChoices{
			Cycle: 24 * time.Hour,
			Altitude: reliableLedger.Altitude,
			Digest:   reliableLedger.Digest(),
		},
		assessmentCompletePeer,
		[]supplier.Supplier{assessmentCompletePeer},
		dbs.New(dbm.FreshMemoryDatastore(), successionUUID),
		agile.Tracer(log.VerifyingTracer()),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.ValidateAgileLedgerLocatedAltitude(context.Background(), 1, byteMoment)
		if err != nil {
			b.Fatal(err)
		}
	}
}
