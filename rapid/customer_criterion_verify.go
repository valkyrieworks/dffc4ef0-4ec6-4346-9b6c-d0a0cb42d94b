package rapid_test

import (
	"context"
	"testing"
	"time"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/rapid"
	"github.com/valkyrieworks/rapid/source"
	mocknode "github.com/valkyrieworks/rapid/source/emulate"
	dbs "github.com/valkyrieworks/rapid/depot/db"
)

//
//
//
//
//
//
//
var (
	criterionCompleteMember = mocknode.New(generateEmulateMember(ledgerUID, 1000, 100, 1, byteTime))
	originLedger, _   = criterionCompleteMember.RapidLedger(context.Background(), 1)
)

func CriterionSeries(b *testing.B) {
	c, err := rapid.NewCustomer(
		context.Background(),
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 24 * time.Hour,
			Level: 1,
			Digest:   originLedger.Digest(),
		},
		criterionCompleteMember,
		[]source.Source{criterionCompleteMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
		rapid.OrderedValidation(),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.ValidateRapidLedgerAtLevel(context.Background(), 1000, byteTime.Add(1000*time.Minute))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func CriterionDivision(b *testing.B) {
	c, err := rapid.NewCustomer(
		context.Background(),
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 24 * time.Hour,
			Level: 1,
			Digest:   originLedger.Digest(),
		},
		criterionCompleteMember,
		[]source.Source{criterionCompleteMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.ValidateRapidLedgerAtLevel(context.Background(), 1000, byteTime.Add(1000*time.Minute))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func CriterionReverse(b *testing.B) {
	validatedLedger, _ := criterionCompleteMember.RapidLedger(context.Background(), 0)
	c, err := rapid.NewCustomer(
		context.Background(),
		ledgerUID,
		rapid.ValidateOptions{
			Duration: 24 * time.Hour,
			Level: validatedLedger.Level,
			Digest:   validatedLedger.Digest(),
		},
		criterionCompleteMember,
		[]source.Source{criterionCompleteMember},
		dbs.New(dbm.NewMemoryStore(), ledgerUID),
		rapid.Tracer(log.VerifyingTracer()),
	)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err = c.ValidateRapidLedgerAtLevel(context.Background(), 1, byteTime)
		if err != nil {
			b.Fatal(err)
		}
	}
}
