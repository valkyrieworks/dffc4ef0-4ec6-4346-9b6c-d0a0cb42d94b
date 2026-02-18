package kv

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/google/orderedcode"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	ordinalutil "github.com/valkyrieworks/intrinsic/ordinaler"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/utils/broadcast/inquire/grammar"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/kinds"
)

var _ ordinaler.LedgerOrdinaler = (*ImpedimentOrdinaler)(nil)

//
//
//
type ImpedimentOrdinaler struct {
	depot dbm.DB

	//
	//
	eventSeq int64
	log      log.Tracer
}

func New(depot dbm.DB) *ImpedimentOrdinaler {
	return &ImpedimentOrdinaler{
		depot: depot,
	}
}

func (idx *ImpedimentOrdinaler) AssignTracer(l log.Tracer) {
	idx.log = l
}

//
//
func (idx *ImpedimentOrdinaler) Has(level int64) (bool, error) {
	key, err := levelKey(level)
	if err != nil {
		return false, fmt.Errorf("REDACTED", err)
	}

	return idx.depot.Has(key)
}

//
//
//
//
//
func (idx *ImpedimentOrdinaler) Ordinal(bh kinds.EventDataNewLedgerEvents) error {
	group := idx.depot.NewGroup()
	defer group.End()

	level := bh.Level

	//
	key, err := levelKey(level)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := group.Set(key, int64toOctets(level)); err != nil {
		return err
	}

	//
	if err := idx.ordinalEvents(group, bh.Events, level); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return group.RecordAlign()
}

//
//
//
//
//
func (idx *ImpedimentOrdinaler) Scan(ctx context.Context, q *inquire.Inquire) ([]int64, error) {
	outcomes := make([]int64, 0)
	select {
	case <-ctx.Done():
		return outcomes, nil

	default:
	}

	states := q.Grammar()

	//
	omitListings := make([]int, 0)

	var ok bool

	var levelDetails LevelDetails
	//
	//
	states, levelDetails, ok = deduplicateLevel(states)

	//
	//
	spans, scopeListings, levelScope := ordinaler.PeekForSpansWithLevel(states)
	levelDetails.levelScope = levelScope

	//
	//
	//
	//
	//
	//
	if ok && levelDetails.solelyLevelEqual {
		ok, err := idx.Has(levelDetails.level)
		if err != nil {
			return nil, err
		}

		if ok {
			return []int64{levelDetails.level}, nil
		}

		return outcomes, nil
	}

	var levelsSetup bool
	screenedLevels := make(map[string][]byte)
	if levelDetails.levelEqualIndex != -1 {
		omitListings = append(omitListings, levelDetails.levelEqualIndex)
	}

	if len(spans) > 0 {
		omitListings = append(omitListings, scopeListings...)

		for _, qr := range spans {
			//
			//
			//
			//
			if qr.Key == kinds.LedgerLevelKey && !levelDetails.solelyLevelScope {
				//
				//
				//
				//

				continue
			}
			prefix, err := orderedcode.Append(nil, qr.Key)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", err)
			}

			if !levelsSetup {
				screenedLevels, err = idx.alignScope(ctx, qr, prefix, screenedLevels, true, levelDetails)
				if err != nil {
					return nil, err
				}

				levelsSetup = true

				//
				//
				if len(screenedLevels) == 0 {
					break
				}
			} else {
				screenedLevels, err = idx.alignScope(ctx, qr, prefix, screenedLevels, false, levelDetails)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	//
	for i, c := range states {
		if integerInSection(i, omitListings) {
			continue
		}

		beginKey, err := orderedcode.Append(nil, c.Tag, c.Arg.Item())
		if err != nil {
			return nil, err
		}

		if !levelsSetup {
			screenedLevels, err = idx.align(ctx, c, beginKey, screenedLevels, true, levelDetails)
			if err != nil {
				return nil, err
			}

			levelsSetup = true

			//
			//
			if len(screenedLevels) == 0 {
				break
			}
		} else {
			screenedLevels, err = idx.align(ctx, c, beginKey, screenedLevels, false, levelDetails)
			if err != nil {
				return nil, err
			}
		}
	}

	//
	outcomes = make([]int64, 0, len(screenedLevels))
	outcomeIndex := make(map[int64]struct{})

FOR_CYCLE:
	for _, hBz := range screenedLevels {
		h := int64fromOctets(hBz)

		ok, err := idx.Has(h)
		if err != nil {
			return nil, err
		}
		if ok {
			if _, ok := outcomeIndex[h]; !ok {
				outcomeIndex[h] = struct{}{}
				outcomes = append(outcomes, h)
			}
		}

		select {
		case <-ctx.Done():
			break FOR_CYCLE

		default:
		}
	}

	sort.Slice(outcomes, func(i, j int) bool { return outcomes[i] < outcomes[j] })

	return outcomes, nil
}

//
//
//
//
//
//
func (idx *ImpedimentOrdinaler) alignScope(
	ctx context.Context,
	qr ordinaler.InquireSpan,
	beginKey []byte,
	screenedLevels map[string][]byte,
	initialRun bool,
	levelDetails LevelDetails,
) (map[string][]byte, error) {
	//
	//
	if !initialRun && len(screenedLevels) == 0 {
		return screenedLevels, nil
	}

	tempLevels := make(map[string][]byte)

	it, err := dbm.RecursePrefix(idx.depot, beginKey)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	defer it.End()

Cycle:
	for ; it.Sound(); it.Following() {
		var (
			eventItem string
			err        error
		)

		if qr.Key == kinds.LedgerLevelKey {
			eventItem, err = analyzeItemFromLeadingKey(it.Key())
		} else {
			eventItem, err = analyzeItemFromEventKey(it.Key())
		}

		if err != nil {
			continue
		}

		if _, ok := qr.AnyLimited().(*big.Float); ok {
			v := new(big.Int)
			v, ok := v.SetString(eventItem, 10)
			var vF *big.Float
			if !ok {
				//
				//
				vF, _, err = big.ParseFloat(eventItem, 10, 125, big.ToNearestEven)
				if err != nil {
					continue Cycle
				}
			}

			if qr.Key != kinds.LedgerLevelKey {
				keyLevel, err := analyzeLevelFromEventKey(it.Key())
				if err != nil {
					idx.log.Fault("REDACTED", err)
					continue Cycle
				}
				insideLevel, err := inspectLevelStates(levelDetails, keyLevel)
				if err != nil {
					idx.log.Fault("REDACTED", err)
					continue Cycle
				}
				if !insideLevel {
					continue Cycle
				}
			}

			var insideLimits bool
			var err error
			if !ok {
				insideLimits, err = ordinalutil.InspectLimits(qr, vF)
			} else {
				insideLimits, err = ordinalutil.InspectLimits(qr, v)
			}
			if err != nil {
				idx.log.Fault("REDACTED", err)
			} else if insideLimits {
				idx.collectionTempLevels(tempLevels, it)
			}
		}

		select {
		case <-ctx.Done():
			break Cycle
		default:
		}
	}

	if err := it.Fault(); err != nil {
		return nil, err
	}

	if len(tempLevels) == 0 || initialRun {
		//
		//
		//
		//
		//
		//
		//
		return tempLevels, nil
	}

	//
	//
FOR_CYCLE:
	for k, v := range screenedLevels {
		tempLevel := tempLevels[k]

		//
		//
		if tempLevel == nil || !bytes.Equal(tempLevel, v) {
			delete(screenedLevels, k)

			select {
			case <-ctx.Done():
				break FOR_CYCLE
			default:
			}
		}
	}

	return screenedLevels, nil
}

func (idx *ImpedimentOrdinaler) collectionTempLevels(tempLevels map[string][]byte, it dbm.Repeater) {
	//
	//
	eventSeq, _ := analyzeEventSeqFromEventKey(it.Key())

	//
	item := make([]byte, len(it.Item()))
	copy(item, it.Item())

	tempLevels[string(item)+strconv.FormatInt(eventSeq, 10)] = item
}

//
//
//
//
//
//
func (idx *ImpedimentOrdinaler) align(
	ctx context.Context,
	c grammar.State,
	beginKeyBz []byte,
	screenedLevels map[string][]byte,
	initialRun bool,
	levelDetails LevelDetails,
) (map[string][]byte, error) {
	//
	//
	if !initialRun && len(screenedLevels) == 0 {
		return screenedLevels, nil
	}

	tempLevels := make(map[string][]byte)

	switch c.Op {
	case grammar.TEq:
		it, err := dbm.RecursePrefix(idx.depot, beginKeyBz)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		defer it.End()

		for ; it.Sound(); it.Following() {

			keyLevel, err := analyzeLevelFromEventKey(it.Key())
			if err != nil {
				idx.log.Fault("REDACTED", err)
				continue
			}
			insideLevel, err := inspectLevelStates(levelDetails, keyLevel)
			if err != nil {
				idx.log.Fault("REDACTED", err)
				continue
			}
			if !insideLevel {
				continue
			}

			idx.collectionTempLevels(tempLevels, it)

			if err := ctx.Err(); err != nil {
				break
			}
		}

		if err := it.Fault(); err != nil {
			return nil, err
		}

	case grammar.TPresent:
		prefix, err := orderedcode.Append(nil, c.Tag)
		if err != nil {
			return nil, err
		}

		it, err := dbm.RecursePrefix(idx.depot, prefix)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		defer it.End()

	CYCLE_PRESENT:
		for ; it.Sound(); it.Following() {

			keyLevel, err := analyzeLevelFromEventKey(it.Key())
			if err != nil {
				idx.log.Fault("REDACTED", err)
				continue
			}
			insideLevel, err := inspectLevelStates(levelDetails, keyLevel)
			if err != nil {
				idx.log.Fault("REDACTED", err)
				continue
			}
			if !insideLevel {
				continue
			}

			idx.collectionTempLevels(tempLevels, it)

			select {
			case <-ctx.Done():
				break CYCLE_PRESENT

			default:
			}
		}

		if err := it.Fault(); err != nil {
			return nil, err
		}

	case grammar.TIncludes:
		prefix, err := orderedcode.Append(nil, c.Tag)
		if err != nil {
			return nil, err
		}

		it, err := dbm.RecursePrefix(idx.depot, prefix)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		defer it.End()

	CYCLE_INCLUDES:
		for ; it.Sound(); it.Following() {
			eventItem, err := analyzeItemFromEventKey(it.Key())
			if err != nil {
				continue
			}

			if strings.Contains(eventItem, c.Arg.Item()) {
				keyLevel, err := analyzeLevelFromEventKey(it.Key())
				if err != nil {
					idx.log.Fault("REDACTED", err)
					continue
				}
				insideLevel, err := inspectLevelStates(levelDetails, keyLevel)
				if err != nil {
					idx.log.Fault("REDACTED", err)
					continue
				}
				if !insideLevel {
					continue
				}
				idx.collectionTempLevels(tempLevels, it)
			}

			select {
			case <-ctx.Done():
				break CYCLE_INCLUDES

			default:
			}
		}
		if err := it.Fault(); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("REDACTED")
	}

	if len(tempLevels) == 0 || initialRun {
		//
		//
		//
		//
		//
		//
		//
		return tempLevels, nil
	}

	//
	//
FOR_CYCLE:
	for k, v := range screenedLevels {
		tempLevel := tempLevels[k]
		if tempLevel == nil || !bytes.Equal(tempLevel, v) {
			delete(screenedLevels, k)

			select {
			case <-ctx.Done():
				break FOR_CYCLE

			default:
			}
		}
	}

	return screenedLevels, nil
}

func (idx *ImpedimentOrdinaler) ordinalEvents(group dbm.Group, events []iface.Event, level int64) error {
	levelBz := int64toOctets(level)

	for _, event := range events {
		idx.eventSeq += 1
		//
		if len(event.Kind) == 0 {
			continue
		}

		for _, property := range event.Properties {
			if len(property.Key) == 0 {
				continue
			}

			//
			compoundKey := event.Kind + "REDACTED" + property.Key
			if compoundKey == kinds.LedgerLevelKey {
				return fmt.Errorf("REDACTED", compoundKey)
			}

			if property.FetchOrdinal() {
				key, err := eventKey(compoundKey, property.Item, level, idx.eventSeq)
				if err != nil {
					return fmt.Errorf("REDACTED", err)
				}

				if err := group.Set(key, levelBz); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
