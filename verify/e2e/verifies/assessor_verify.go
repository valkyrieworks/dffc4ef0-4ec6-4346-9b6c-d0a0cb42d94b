package end2end_typ_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
func Testverifier_Groupings(t *testing.T) {
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		if peer.Style == e2e.StyleGerm {
			return
		}

		customer, err := peer.Customer()
		require.NoError(t, err)
		condition, err := customer.Condition(ctx)
		require.NoError(t, err)

		initial := condition.ChronizeDetails.InitialLedgerAltitude
		final := condition.ChronizeDetails.NewestLedgerAltitude

		//
		if peer.PreserveLedgers > 0 {
			initial++
		}

		itemTimeline := freshAssessorTimeline(*peer.Simnet)
		itemTimeline.Advance(initial - peer.Simnet.PrimaryAltitude)

		for h := initial; h <= final; h++ {
			assessors := []*kinds.Assessor{}
			everyScreen := 100
			for screen := 1; ; screen++ {
				reply, err := customer.Assessors(ctx, &(h), &(screen), &everyScreen)
				require.NoError(t, err)
				assessors = append(assessors, reply.Assessors...)
				if len(assessors) == reply.Sum {
					break
				}
			}
			require.Equal(t, itemTimeline.Set.Assessors, assessors,
				"REDACTED", h)
			itemTimeline.Advance(1)
		}
	})
}

//
//
func Testverifier_Nominate(t *testing.T) {
	ledgers := acquireLedgerSuccession(t)
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		if peer.Style != e2e.StyleAssessor {
			return
		}
		location := peer.PrivatevalueToken.PublicToken().Location()
		itemTimeline := freshAssessorTimeline(*peer.Simnet)

		anticipateTally := 0
		nominateTally := 0
		for _, ledger := range ledgers {
			if bytes.Equal(itemTimeline.Set.Nominator.Location, location) {
				anticipateTally++
				if bytes.Equal(ledger.NominatorLocation, location) {
					nominateTally++
				}
			}
			itemTimeline.Advance(1)
		}

		require.False(t, nominateTally == 0 && anticipateTally > 0,
			"REDACTED", anticipateTally)
		if anticipateTally > 5 {
			require.GreaterOrEqual(t, nominateTally, 3, "REDACTED")
		}
	})
}

//
//
func Testverifier_Attest(t *testing.T) {
	ledgers := acquireLedgerSuccession(t)
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		if peer.Style != e2e.StyleAssessor {
			return
		}
		location := peer.PrivatevalueToken.PublicToken().Location()
		itemTimeline := freshAssessorTimeline(*peer.Simnet)

		anticipateTally := 0
		attestTally := 0
		for _, ledger := range ledgers[1:] { //
			notated := false
			for _, sig := range ledger.FinalEndorse.Notations {
				if bytes.Equal(sig.AssessorLocation, location) {
					notated = true
					break
				}
			}
			if itemTimeline.Set.OwnsLocation(location) {
				anticipateTally++
				if notated {
					attestTally++
				}
			} else {
				require.False(t, notated, "REDACTED", ledger.FinalEndorse.Altitude)
			}
			itemTimeline.Advance(1)
		}

		require.False(t, attestTally == 0 && anticipateTally > 0,
			"REDACTED", anticipateTally)
		if anticipateTally > 7 {
			require.GreaterOrEqual(t, attestTally, 3, "REDACTED", anticipateTally)
		}
	})
}

//
//
type assessorTimeline struct {
	Set     *kinds.AssessorAssign
	altitude  int64
	revisions map[int64]map[*e2e.Peer]int64
}

func freshAssessorTimeline(simnet e2e.Simnet) *assessorTimeline {
	itemIndex := simnet.Assessors                  //
	if v, ok := simnet.AssessorRevisions[0]; ok { //
		itemIndex = v
	}
	return &assessorTimeline{
		altitude:  simnet.PrimaryAltitude,
		Set:     kinds.FreshAssessorAssign(createValues(itemIndex)),
		revisions: simnet.AssessorRevisions,
	}
}

func (s *assessorTimeline) Advance(elevations int64) {
	for i := int64(0); i < elevations; i++ {
		s.altitude++
		if s.altitude > 2 {
			//
			//
			if revise, ok := s.revisions[s.altitude-2]; ok {
				if err := s.Set.ReviseUsingModifyAssign(createValues(revise)); err != nil {
					panic(err)
				}
			}
		}
		s.Set.AdvanceNominatorUrgency(1)
	}
}

func createValues(itemIndex map[*e2e.Peer]int64) []*kinds.Assessor {
	values := make([]*kinds.Assessor, 0, len(itemIndex))
	for peer, potency := range itemIndex {
		values = append(values, kinds.FreshAssessor(peer.PrivatevalueToken.PublicToken(), potency))
	}
	return values
}
