package simulate

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

type Simulate struct {
	successionUUID string

	mtx              sync.Mutex
	headings          map[int64]*kinds.NotatedHeading
	values             map[int64]*kinds.AssessorAssign
	proofTowardNotify map[string]kinds.Proof //
	newestAltitude     int64
}

var _ supplier.Supplier = (*Simulate)(nil)

//
//
func New(successionUUID string, headings map[int64]*kinds.NotatedHeading, values map[int64]*kinds.AssessorAssign) *Simulate {
	altitude := int64(0)
	for h := range headings {
		if h > altitude {
			altitude = h
		}
	}
	return &Simulate{
		successionUUID:          successionUUID,
		headings:          headings,
		values:             values,
		proofTowardNotify: make(map[string]kinds.Proof),
		newestAltitude:     altitude,
	}
}

//
func (p *Simulate) SuccessionUUID() string {
	return p.successionUUID
}

func (p *Simulate) Text() string {
	var headings strings.Builder
	for _, h := range p.headings {
		fmt.Fprintf(&headings, "REDACTED", h.Altitude, h.Digest())
	}

	var values strings.Builder
	for _, v := range p.values {
		fmt.Fprintf(&values, "REDACTED", v.Digest())
	}

	return fmt.Sprintf("REDACTED", headings.String(), values.String())
}

func (p *Simulate) AgileLedger(ctx context.Context, altitude int64) (*kinds.AgileLedger, error) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	//
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(10 * time.Millisecond):
	}

	var lb *kinds.AgileLedger

	if altitude > p.newestAltitude {
		return nil, supplier.FaultAltitudeExcessivelyTall
	}

	if altitude == 0 && len(p.headings) > 0 {
		altitude = p.newestAltitude
	}

	if _, ok := p.headings[altitude]; ok {
		sh := p.headings[altitude]
		values := p.values[altitude]
		lb = &kinds.AgileLedger{
			NotatedHeading: sh,
			AssessorAssign: values,
		}
	}
	if lb == nil {
		return nil, supplier.FaultAgileLedgerNegationDetected
	}
	if lb.NotatedHeading == nil || lb.AssessorAssign == nil {
		return nil, supplier.FaultFlawedAgileLedger{Rationale: errors.New("REDACTED")}
	}
	if err := lb.CertifyFundamental(lb.SuccessionUUID); err != nil {
		return nil, supplier.FaultFlawedAgileLedger{Rationale: err}
	}
	return lb, nil
}

func (p *Simulate) NotifyProof(_ context.Context, ev kinds.Proof) error {
	p.proofTowardNotify[string(ev.Digest())] = ev
	return nil
}

func (p *Simulate) OwnsProof(ev kinds.Proof) bool {
	_, ok := p.proofTowardNotify[string(ev.Digest())]
	return ok
}

func (p *Simulate) AppendAgileLedger(lb *kinds.AgileLedger) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if err := lb.CertifyFundamental(lb.SuccessionUUID); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	p.headings[lb.Altitude] = lb.NotatedHeading
	p.values[lb.Altitude] = lb.AssessorAssign
	if lb.Altitude > p.newestAltitude {
		p.newestAltitude = lb.Altitude
	}
}

func (p *Simulate) Duplicate(id string) *Simulate {
	return New(id, p.headings, p.values)
}
