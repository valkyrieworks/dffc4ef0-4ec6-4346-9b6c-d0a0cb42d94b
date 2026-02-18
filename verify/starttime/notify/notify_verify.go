package notifier_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/valkyrieworks/verify/starttime/shipment"
	"github.com/valkyrieworks/verify/starttime/notify"
	"github.com/valkyrieworks/kinds"
)

type emulateLedgerDepot struct {
	root   int64
	ledgers []*kinds.Ledger
}

func (m *emulateLedgerDepot) Level() int64 {
	return m.root + int64(len(m.ledgers))
}

func (m *emulateLedgerDepot) Root() int64 {
	return m.root
}

func (m *emulateLedgerDepot) ImportLedger(i int64) *kinds.Ledger {
	return m.ledgers[i-m.root]
}

func VerifyComposeNotify(t *testing.T) {
	t1 := time.Now()
	u := [16]byte(uuid.New())
	b1, err := shipment.NewOctets(&shipment.Shipment{
		Id:   u[:],
		Time: timestamppb.New(t1.Add(-10 * time.Second)),
		Volume: 1024,
	})
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	b2, err := shipment.NewOctets(&shipment.Shipment{
		Id:   u[:],
		Time: timestamppb.New(t1.Add(-4 * time.Second)),
		Volume: 1024,
	})
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	b3, err := shipment.NewOctets(&shipment.Shipment{
		Id:   u[:],
		Time: timestamppb.New(t1.Add(2 * time.Second)),
		Volume: 1024,
	})
	t2 := t1.Add(time.Second)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	s := &emulateLedgerDepot{
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
					Time: t1,
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
					Time: t2,
				},
				Data: kinds.Data{
					Txs: []kinds.Tx{},
				},
			},
		},
	}
	rs, err := notify.ComposeFromLedgerDepot(s)
	if err != nil {
		t.Fatalf("REDACTED", err)
	}
	if rs.FaultNumber() != 1 {
		t.Fatalf("REDACTED", 1, rs.FaultNumber())
	}
	rl := rs.Catalog()
	if len(rl) != 1 {
		t.Fatalf("REDACTED", 1, len(rl))
	}
	r := rl[0]
	if len(r.All) != 4 {
		t.Fatalf("REDACTED", 4, len(r.All))
	}
	if r.AdverseNumber != 2 {
		t.Fatalf("REDACTED", 2, r.AdverseNumber)
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
