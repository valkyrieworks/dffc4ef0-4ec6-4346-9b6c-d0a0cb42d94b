package chainchronize

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyHandlerAggregate(t *testing.T) {
	const awaitMoment = 5 * time.Second
	t.Run("REDACTED", func(t *testing.T) {
		//
		ts := freshAggregateStyleVerifyCollection(t, "REDACTED")

		//
		//
		var (
			supplier = freshHandler(t, ts.tracer, ts.producePaper, ts.privateItems, 4, usingCertainBallotMultiples())
			adherent = freshHandler(t, ts.tracer, ts.producePaper, ts.privateItems, 2, usingCertainBallotMultiples())
		)

		adherent.handler.aggregateStyleActivated = true
		adherent.handler.durationConditionRevise = aggregateStyleIntrinsicConditionRevise
		supplier.handler.durationConditionRevise = aggregateStyleIntrinsicConditionRevise

		ts.ledgerAbsorber.AssignUponAbsorb(func(ic agreement.AbsorbNominee) error {
			ts.tracer.Details("REDACTED", "REDACTED", ic.Altitude())
			return nil
		})

		//
		routers := p2p.CreateAssociatedRouters(ts.settings.P2P, 2, func(i int, s *p2p.Router) *p2p.Router {
			switch i {
			case 0:
				s.AppendHandler("REDACTED", supplier.handler)
			case 1:
				s.AppendHandler("REDACTED", adherent.handler)
				s.AppendHandler("REDACTED", ts.ledgerAbsorber)
			}

			return s
		}, p2p.Connect2routers)

		//
		adherentRouter := routers[1]

		//
		require.True(t, adherentRouter.EqualsActive())
		require.True(t, adherent.handler.EqualsActive())

		//
		//
		//
		//
		inspect := func() bool {
			return len(ts.ledgerAbsorber.Solicits()) == 1
		}

		//
		require.Eventually(t, inspect, awaitMoment, 100*time.Millisecond)

		//
		ledger := ts.ledgerAbsorber.Solicits()[0]
		require.Equal(t, int64(3), ledger.Altitude())

		//
		require.Equal(t, int64(4), adherent.handler.hub.Altitude())
	})

	t.Run("REDACTED", func(t *testing.T) {
		t.Skip()
		//
		ts := freshAggregateStyleVerifyCollection(t, "REDACTED")

		//
		//
		var (
			supplier = freshHandler(t, ts.tracer, ts.producePaper, ts.privateItems, 4, usingCertainBallotMultiples())
			adherent = freshHandler(t, ts.tracer, ts.producePaper, ts.privateItems, 2, usingCertainBallotMultiples())
		)

		adherent.handler.aggregateStyleActivated = true
		adherent.handler.durationConditionRevise = aggregateStyleIntrinsicConditionRevise
		supplier.handler.durationConditionRevise = aggregateStyleIntrinsicConditionRevise

		ts.ledgerAbsorber.AssignUponAbsorb(func(ic agreement.AbsorbNominee) error {
			ts.tracer.Details("REDACTED", "REDACTED", ic.Altitude())
			return agreement.FaultEarlierComprised
		})

		//
		routers := p2p.CreateAssociatedRouters(ts.settings.P2P, 2, func(i int, s *p2p.Router) *p2p.Router {
			switch i {
			case 0:
				s.AppendHandler("REDACTED", supplier.handler)
			case 1:
				s.AppendHandler("REDACTED", adherent.handler)
				s.AppendHandler("REDACTED", ts.ledgerAbsorber)
			}

			return s
		}, p2p.Connect2routers)

		//
		adherentRouter := routers[1]

		//
		require.True(t, adherentRouter.EqualsActive())
		require.True(t, adherent.handler.EqualsActive())

		//
		//
		//
		//
		inspect := func() bool {
			return len(ts.ledgerAbsorber.Solicits()) == 1
		}

		//
		require.Eventually(t, inspect, awaitMoment, 100*time.Millisecond)

		ledger := ts.ledgerAbsorber.Solicits()[0]
		require.Equal(t, int64(3), ledger.Altitude())

		//
		require.Equal(t, int64(4), adherent.handler.hub.Altitude())
	})
}

type aggregateStyleVerifyCollection struct {
	t             *testing.T
	ledgerAbsorber *ledgerAbsorberSimulate
	settings        *cfg.Settings
	producePaper        *kinds.OriginPaper
	privateItems      []kinds.PrivateAssessor
	tracer        log.Tracer
}

func freshAggregateStyleVerifyCollection(t *testing.T, alias string) *aggregateStyleVerifyCollection {
	tracer := log.VerifyingTracer()

	settings := verify.RestoreVerifyOrigin(alias)
	t.Cleanup(func() { _ = os.RemoveAll(settings.OriginPath) })

	producePaper, privateItems := arbitraryOriginPaper(1, false, 30)

	return &aggregateStyleVerifyCollection{
		t:             t,
		ledgerAbsorber: freshLedgerAbsorberSimulate(t),
		settings:        settings,
		producePaper:        producePaper,
		privateItems:      privateItems,
		tracer:        tracer,
	}
}

type ledgerAbsorberSimulate struct {
	*p2p.FoundationHandler

	t           *testing.T
	mu          chronize.Exclusion
	uponAbsorb    func(agreement.AbsorbNominee) error
	persistedInvocations []agreement.AbsorbNominee
}

var _ LedgerAbsorber = (*ledgerAbsorberSimulate)(nil)

func freshLedgerAbsorberSimulate(t *testing.T) *ledgerAbsorberSimulate {
	m := &ledgerAbsorberSimulate{
		t: t,
	}
	m.FoundationHandler = p2p.FreshFoundationHandler("REDACTED", m)
	return m
}

func (m *ledgerAbsorberSimulate) AbsorbAttestedLedger(ic agreement.AbsorbNominee) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.persistedInvocations = append(m.persistedInvocations, ic)

	if m.uponAbsorb == nil {
		return nil
	}

	return m.uponAbsorb(ic)
}

func (m *ledgerAbsorberSimulate) AssignUponAbsorb(uponAbsorb func(ic agreement.AbsorbNominee) error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.uponAbsorb = uponAbsorb
}

func (m *ledgerAbsorberSimulate) Solicits() []agreement.AbsorbNominee {
	m.mu.Lock()
	defer m.mu.Unlock()

	out := make([]agreement.AbsorbNominee, len(m.persistedInvocations))
	copy(out, m.persistedInvocations)
	return out
}
