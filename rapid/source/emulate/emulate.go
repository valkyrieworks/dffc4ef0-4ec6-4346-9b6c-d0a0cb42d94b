package emulate

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/valkyrieworks/rapid/source"
	"github.com/valkyrieworks/kinds"
)

type Emulate struct {
	ledgerUID string

	mtx              sync.Mutex
	headings          map[int64]*kinds.AttestedHeading
	values             map[int64]*kinds.RatifierAssign
	proofToNotify map[string]kinds.Proof //
	newestLevel     int64
}

var _ source.Source = (*Emulate)(nil)

//
//
func New(ledgerUID string, headings map[int64]*kinds.AttestedHeading, values map[int64]*kinds.RatifierAssign) *Emulate {
	level := int64(0)
	for h := range headings {
		if h > level {
			level = h
		}
	}
	return &Emulate{
		ledgerUID:          ledgerUID,
		headings:          headings,
		values:             values,
		proofToNotify: make(map[string]kinds.Proof),
		newestLevel:     level,
	}
}

//
func (p *Emulate) LedgerUID() string {
	return p.ledgerUID
}

func (p *Emulate) String() string {
	var headings strings.Builder
	for _, h := range p.headings {
		fmt.Fprintf(&headings, "REDACTED", h.Level, h.Digest())
	}

	var values strings.Builder
	for _, v := range p.values {
		fmt.Fprintf(&values, "REDACTED", v.Digest())
	}

	return fmt.Sprintf("REDACTED", headings.String(), values.String())
}

func (p *Emulate) RapidLedger(ctx context.Context, level int64) (*kinds.RapidLedger, error) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	//
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(10 * time.Millisecond):
	}

	var lb *kinds.RapidLedger

	if level > p.newestLevel {
		return nil, source.ErrLevelTooElevated
	}

	if level == 0 && len(p.headings) > 0 {
		level = p.newestLevel
	}

	if _, ok := p.headings[level]; ok {
		sh := p.headings[level]
		values := p.values[level]
		lb = &kinds.RapidLedger{
			AttestedHeading: sh,
			RatifierAssign: values,
		}
	}
	if lb == nil {
		return nil, source.ErrRapidLedgerNegateLocated
	}
	if lb.AttestedHeading == nil || lb.RatifierAssign == nil {
		return nil, source.ErrFlawedRapidLedger{Cause: errors.New("REDACTED")}
	}
	if err := lb.CertifySimple(lb.LedgerUID); err != nil {
		return nil, source.ErrFlawedRapidLedger{Cause: err}
	}
	return lb, nil
}

func (p *Emulate) NotifyProof(_ context.Context, ev kinds.Proof) error {
	p.proofToNotify[string(ev.Digest())] = ev
	return nil
}

func (p *Emulate) HasProof(ev kinds.Proof) bool {
	_, ok := p.proofToNotify[string(ev.Digest())]
	return ok
}

func (p *Emulate) AppendRapidLedger(lb *kinds.RapidLedger) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if err := lb.CertifySimple(lb.LedgerUID); err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	p.headings[lb.Level] = lb.AttestedHeading
	p.values[lb.Level] = lb.RatifierAssign
	if lb.Level > p.newestLevel {
		p.newestLevel = lb.Level
	}
}

func (p *Emulate) Clone(id string) *Emulate {
	return New(id, p.headings, p.values)
}
