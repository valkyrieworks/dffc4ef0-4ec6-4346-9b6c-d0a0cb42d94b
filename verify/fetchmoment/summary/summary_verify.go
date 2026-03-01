package sumry_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/fetchmoment/content"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/fetchmoment/summary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

type simulateLedgerDepot struct {
	foundation   int64
	ledgers []*kinds.Ledger
}

func (m *simulateLedgerDepot) Altitude() int64 {
	return m.foundation + int64(len(m.ledgers))
}

func (m *simulateLedgerDepot) Foundation() int64 {
	return m.foundation
}

func (m *simulateLedgerDepot) FetchLedger(i int64) *kinds.Ledger {
	return m.ledgers[i-m.foundation]
}

func VerifyComposeSummary(t *testing.T) {
	t1 := time.Now()
	u := [16]byte(uuid.New())
	b1, err := content.FreshOctets(&content.Content{
		Id:   u[:],
		Moment: timestamppb.New(t1.Add(-10 * time.Second)),
		Extent: 1024,
	})
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	b2, err := content.FreshOctets(&content.Content{
		Id:   u[:],
		Moment: timestamppb.New(t1.Add(-4 * time.Second)),
		Extent: 1024,
	})
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	b3, err := content.FreshOctets(&content.Content{
		Id:   u[:],
		Moment: timestamppb.New(t1.Add(2 * time.Second)),
		Extent: 1024,
	})
	t2 := t1.Add(time.Second)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	s := &simulateLedgerDepot{
		ledgers: []*kinds.Ledger{
			{
				Data: kinds.Data{
					Txs: []kinds.Tx{b1, b2},
				},
			},
			{
				//
				//
				Heading: kinds.Heading{
					Moment: t1,
				},
				Data: kinds.Data{
					Txs: []kinds.Tx{[]byte("REDACTED")},
				},
			},
			{
				Data: kinds.Data{
					Txs: []kinds.Tx{b3, b3},
				},
			},
			{
				Heading: kinds.Heading{
					Moment: t2,
				},
				Data: kinds.Data{
					Txs: []kinds.Tx{},
				},
			},
		},
	}
	rs, err := summary.ComposeOriginatingLedgerDepot(s)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if rs.FailureTally() != 1 {
		t.Fatalf("REDACTED", 1, rs.FailureTally())
	}
	rl := rs.Catalog()
	if len(rl) != 1 {
		t.Fatalf("REDACTED", 1, len(rl))
	}
	r := rl[0]
	if len(r.All) != 4 {
		t.Fatalf("REDACTED", 4, len(r.All))
	}
	if r.AdverseTally != 2 {
		t.Fatalf("REDACTED", 2, r.AdverseTally)
	}
	if r.Avg != 3*time.Second {
		t.Fatalf("REDACTED", 3*time.Second, r.Avg)
	}
	if r.Min != -time.Second {
		t.Fatalf("REDACTED", time.Second, r.Min)
	}
	if r.Max != 10*time.Second {
		t.Fatalf("REDACTED", 10*time.Second, r.Max)
	}
	//
	//
	anticipatedStandardDevelop := 5228129047 * time.Nanosecond
	if r.StandardDevelop != anticipatedStandardDevelop {
		t.Fatalf("REDACTED", anticipatedStandardDevelop, r.StandardDevelop)
	}
}
