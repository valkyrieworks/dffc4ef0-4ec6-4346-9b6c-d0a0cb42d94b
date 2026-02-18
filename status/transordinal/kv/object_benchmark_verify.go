package kv

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/kinds"
)

func CriterionTransferScan(b *testing.B) {
	storeFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		b.Errorf("REDACTED", err)
	}

	db, err := dbm.NewGoLayerStore("REDACTED", storeFolder)
	if err != nil {
		b.Errorf("REDACTED", err)
	}

	ordinaler := NewTransOrdinal(db)

	for i := 0; i < 35000; i++ {
		events := []iface.Event{
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{Key: "REDACTED", Item: fmt.Sprintf("REDACTED", i%100), Ordinal: true},
					{Key: "REDACTED", Item: "REDACTED", Ordinal: true},
				},
			},
		}

		transferBz := make([]byte, 8)
		if _, err := rand.Read(transferBz); err != nil {
			b.Errorf("REDACTED", err)
		}

		transOutcome := &iface.TransOutcome{
			Level: int64(i),
			Ordinal:  0,
			Tx:     kinds.Tx(string(transferBz)),
			Outcome: iface.InvokeTransferOutcome{
				Data:   []byte{0},
				Code:   iface.CodeKindSuccess,
				Log:    "REDACTED",
				Events: events,
			},
		}

		if err := ordinaler.Ordinal(transOutcome); err != nil {
			b.Errorf("REDACTED", err)
		}
	}

	transferInquire := inquire.ShouldBuild("REDACTED")

	b.ResetTimer()

	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		if _, err := ordinaler.Scan(ctx, transferInquire); err != nil {
			b.Errorf("REDACTED", err)
		}
	}
}
