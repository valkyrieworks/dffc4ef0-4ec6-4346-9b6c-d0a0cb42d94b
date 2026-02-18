package integration_t_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/kinds"
)

//
//
func Verifyvalidator_Collections(t *testing.T) {
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		if member.Style == e2e.StyleOrigin {
			return
		}

		customer, err := member.Customer()
		require.NoError(t, err)
		state, err := customer.Status(ctx)
		require.NoError(t, err)

		initial := state.AlignDetails.OldestLedgerLevel
		final := state.AlignDetails.NewestLedgerLevel

		//
		if member.PreserveLedgers > 0 {
			initial++
		}

		valueSequence := newRatifierSequence(*member.Verifychain)
		valueSequence.Augment(initial - member.Verifychain.PrimaryLevel)

		for h := initial; h <= final; h++ {
			ratifiers := []*kinds.Ratifier{}
			eachScreen := 100
			for screen := 1; ; screen++ {
				reply, err := customer.Ratifiers(ctx, &(h), &(screen), &eachScreen)
				require.NoError(t, err)
				ratifiers = append(ratifiers, reply.Ratifiers...)
				if len(ratifiers) == reply.Sum {
					break
				}
			}
			require.Equal(t, valueSequence.Set.Ratifiers, ratifiers,
				"REDACTED", h)
			valueSequence.Augment(1)
		}
	})
}

//
//
func Verifyvalidator_Nominate(t *testing.T) {
	ledgers := acquireLedgerSeries(t)
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		if member.Style != e2e.StyleRatifier {
			return
		}
		location := member.PrivatekeyKey.PublicKey().Location()
		valueSequence := newRatifierSequence(*member.Verifychain)

		anticipateNumber := 0
		nominateNumber := 0
		for _, ledger := range ledgers {
			if bytes.Equal(valueSequence.Set.Recommender.Location, location) {
				anticipateNumber++
				if bytes.Equal(ledger.RecommenderLocation, location) {
					nominateNumber++
				}
			}
			valueSequence.Augment(1)
		}

		require.False(t, nominateNumber == 0 && anticipateNumber > 0,
			"REDACTED", anticipateNumber)
		if anticipateNumber > 5 {
			require.GreaterOrEqual(t, nominateNumber, 3, "REDACTED")
		}
	})
}

//
//
func Verifyvalidator_Attest(t *testing.T) {
	ledgers := acquireLedgerSeries(t)
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		if member.Style != e2e.StyleRatifier {
			return
		}
		location := member.PrivatekeyKey.PublicKey().Location()
		valueSequence := newRatifierSequence(*member.Verifychain)

		anticipateNumber := 0
		attestNumber := 0
		for _, ledger := range ledgers[1:] { //
			attested := false
			for _, sig := range ledger.FinalEndorse.Endorsements {
				if bytes.Equal(sig.RatifierLocation, location) {
					attested = true
					break
				}
			}
			if valueSequence.Set.HasLocation(location) {
				anticipateNumber++
				if attested {
					attestNumber++
				}
			} else {
				require.False(t, attested, "REDACTED", ledger.FinalEndorse.Level)
			}
			valueSequence.Augment(1)
		}

		require.False(t, attestNumber == 0 && anticipateNumber > 0,
			"REDACTED", anticipateNumber)
		if anticipateNumber > 7 {
			require.GreaterOrEqual(t, attestNumber, 3, "REDACTED", anticipateNumber)
		}
	})
}

//
//
type ratifierSequence struct {
	Set     *kinds.RatifierAssign
	level  int64
	refreshes map[int64]map[*e2e.Member]int64
}

func newRatifierSequence(verifychain e2e.Verifychain) *ratifierSequence {
	valueIndex := verifychain.Ratifiers                  //
	if v, ok := verifychain.RatifierRefreshes[0]; ok { //
		valueIndex = v
	}
	return &ratifierSequence{
		level:  verifychain.PrimaryLevel,
		Set:     kinds.NewRatifierCollection(createValues(valueIndex)),
		refreshes: verifychain.RatifierRefreshes,
	}
}

func (s *ratifierSequence) Augment(levels int64) {
	for i := int64(0); i < levels; i++ {
		s.level++
		if s.level > 2 {
			//
			//
			if modify, ok := s.refreshes[s.level-2]; ok {
				if err := s.Set.ModifyWithAlterCollection(createValues(modify)); err != nil {
					panic(err)
				}
			}
		}
		s.Set.AugmentRecommenderUrgency(1)
	}
}

func createValues(valueIndex map[*e2e.Member]int64) []*kinds.Ratifier {
	values := make([]*kinds.Ratifier, 0, len(valueIndex))
	for member, energy := range valueIndex {
		values = append(values, kinds.NewRatifier(member.PrivatekeyKey.PublicKey(), energy))
	}
	return values
}
