package kv

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func AssessmentTransferLookup(b *testing.B) {
	datastorePath, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		b.Errorf("REDACTED", err)
	}

	db, err := dbm.FreshProceedStratumDatastore("REDACTED", datastorePath)
	if err != nil {
		b.Errorf("REDACTED", err)
	}

	ordinalizer := FreshTransferOrdinal(db)

	for i := 0; i < 35000; i++ {
		incidents := []iface.Incident{
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{Key: "REDACTED", Datum: fmt.Sprintf("REDACTED", i%100), Ordinal: true},
					{Key: "REDACTED", Datum: "REDACTED", Ordinal: true},
				},
			},
		}

		transferByz := make([]byte, 8)
		if _, err := rand.Read(transferByz); err != nil {
			b.Errorf("REDACTED", err)
		}

		transferOutcome := &iface.TransferOutcome{
			Altitude: int64(i),
			Ordinal:  0,
			Tx:     kinds.Tx(string(transferByz)),
			Outcome: iface.InvokeTransferOutcome{
				Data:   []byte{0},
				Cipher:   iface.CipherKindOKAY,
				Log:    "REDACTED",
				Incidents: incidents,
			},
		}

		if err := ordinalizer.Ordinal(transferOutcome); err != nil {
			b.Errorf("REDACTED", err)
		}
	}

	transferInquire := inquire.ShouldAssemble("REDACTED")

	b.ResetTimer()

	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		if _, err := ordinalizer.Lookup(ctx, transferInquire); err != nil {
			b.Errorf("REDACTED", err)
		}
	}
}
