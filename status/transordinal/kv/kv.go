package kv

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/valkyrieworks/utils/log"

	"github.com/cosmos/gogoproto/proto"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	ordinalutil "github.com/valkyrieworks/intrinsic/ordinaler"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/utils/broadcast/inquire/grammar"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/kinds"
)

const (
	markerKeyDelimiter     = "REDACTED"
	markerKeyDelimiterRune = '/'
	eventSeqDelimiter   = "REDACTED"
)

var _ transordinal.TransOrdinaler = (*TransOrdinal)(nil)

//
type TransOrdinal struct {
	depot dbm.DB
	//
	eventSeq int64

	log log.Tracer
}

//
func NewTransOrdinal(depot dbm.DB) *TransOrdinal {
	return &TransOrdinal{
		depot: depot,
	}
}

func (txi *TransOrdinal) AssignTracer(l log.Tracer) {
	txi.log = l
}

//
//
func (txi *TransOrdinal) Get(digest []byte) (*iface.TransOutcome, error) {
	if len(digest) == 0 {
		return nil, transordinal.FaultEmptyDigest
	}

	crudeOctets, err := txi.depot.Get(digest)
	if err != nil {
		panic(err)
	}
	if crudeOctets == nil {
		return nil, nil
	}

	transOutcome := new(iface.TransOutcome)
	err = proto.Unmarshal(crudeOctets, transOutcome)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return transOutcome, nil
}

//
//
//
//
func (txi *TransOrdinal) AppendGroup(b *transordinal.Group) error {
	depotGroup := txi.depot.NewGroup()
	defer depotGroup.End()

	for _, outcome := range b.Ops {
		digest := kinds.Tx(outcome.Tx).Digest()

		//
		err := txi.ordinalEvents(outcome, digest, depotGroup)
		if err != nil {
			return err
		}

		//
		err = depotGroup.Set(keyForLevel(outcome), digest)
		if err != nil {
			return err
		}

		crudeOctets, err := proto.Marshal(outcome)
		if err != nil {
			return err
		}
		//
		err = depotGroup.Set(digest, crudeOctets)
		if err != nil {
			return err
		}
	}

	return depotGroup.RecordAlign()
}

//
//
//
//
//
//
//
//
//
func (txi *TransOrdinal) Ordinal(outcome *iface.TransOutcome) error {
	b := txi.depot.NewGroup()
	defer b.End()

	digest := kinds.Tx(outcome.Tx).Digest()

	if !outcome.Outcome.IsOK() {
		agedOutcome, err := txi.Get(digest)
		if err != nil {
			return err
		}

		//
		//
		if agedOutcome != nil && agedOutcome.Outcome.Code == iface.CodeKindSuccess {
			return nil
		}
	}

	//
	err := txi.ordinalEvents(outcome, digest, b)
	if err != nil {
		return err
	}

	//
	err = b.Set(keyForLevel(outcome), digest)
	if err != nil {
		return err
	}

	crudeOctets, err := proto.Marshal(outcome)
	if err != nil {
		return err
	}
	//
	err = b.Set(digest, crudeOctets)
	if err != nil {
		return err
	}

	return b.RecordAlign()
}

func (txi *TransOrdinal) ordinalEvents(outcome *iface.TransOutcome, digest []byte, depot dbm.Group) error {
	for _, event := range outcome.Outcome.Events {
		txi.eventSeq += 1
		//
		if len(event.Kind) == 0 {
			continue
		}

		for _, property := range event.Properties {
			if len(property.Key) == 0 {
				continue
			}

			//
			compoundLabel := fmt.Sprintf("REDACTED", event.Kind, property.Key)
			//
			if compoundLabel == kinds.TransferDigestKey || compoundLabel == kinds.TransferLevelKey {
				return fmt.Errorf("REDACTED", compoundLabel)
			}
			if property.FetchOrdinal() {
				err := depot.Set(keyForEvent(compoundLabel, property.Item, outcome, txi.eventSeq), digest)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

//
//
//
//
//
//
//
//
//
//
//
func (txi *TransOrdinal) Scan(ctx context.Context, q *inquire.Inquire) ([]*iface.TransOutcome, error) {
	select {
	case <-ctx.Done():
		return make([]*iface.TransOutcome, 0), nil

	default:
	}

	var digestsSetup bool
	screenedDigests := make(map[string][]byte)

	//
	states := q.Grammar()

	//
	digest, ok, err := scanForDigest(states)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	} else if ok {
		res, err := txi.Get(digest)
		switch {
		case err != nil:
			return []*iface.TransOutcome{}, fmt.Errorf("REDACTED", err)
		case res == nil:
			return []*iface.TransOutcome{}, nil
		default:
			return []*iface.TransOutcome{res}, nil
		}
	}

	//
	omitListings := make([]int, 0)
	var levelDetails LevelDetails

	//
	//
	states, levelDetails = deduplicateLevel(states)

	if !levelDetails.solelyLevelEqual {
		omitListings = append(omitListings, levelDetails.levelEqualOrdinal)
	}

	//
	//
	//
	spans, spanListings, levelSpan := ordinaler.ScanForSpansWithLevel(states)
	levelDetails.levelSpan = levelSpan
	if len(spans) > 0 {
		omitListings = append(omitListings, spanListings...)

		for _, qr := range spans {

			//
			//
			//
			//
			if qr.Key == kinds.TransferLevelKey && !levelDetails.solelyLevelSpan {
				continue
			}
			if !digestsSetup {
				screenedDigests = txi.alignSpan(ctx, qr, beginKey(qr.Key), screenedDigests, true, levelDetails)
				digestsSetup = true

				//
				//
				if len(screenedDigests) == 0 {
					break
				}
			} else {
				screenedDigests = txi.alignSpan(ctx, qr, beginKey(qr.Key), screenedDigests, false, levelDetails)
			}
		}
	}

	//

	//
	for i, c := range states {
		if integerInSection(i, omitListings) {
			continue
		}

		if !digestsSetup {
			screenedDigests = txi.align(ctx, c, beginKeyForStatus(c, levelDetails.level), screenedDigests, true, levelDetails)
			digestsSetup = true

			//
			//
			if len(screenedDigests) == 0 {
				break
			}
		} else {
			screenedDigests = txi.align(ctx, c, beginKeyForStatus(c, levelDetails.level), screenedDigests, false, levelDetails)
		}
	}

	outcomes := make([]*iface.TransOutcome, 0, len(screenedDigests))
	outcomeIndex := make(map[string]struct{})
OUTCOMES_CYCLE:
	for _, h := range screenedDigests {

		res, err := txi.Get(h)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", h, err)
		}
		digestString := string(h)
		if _, ok := outcomeIndex[digestString]; !ok {
			outcomeIndex[digestString] = struct{}{}
			outcomes = append(outcomes, res)
		}
		//
		select {
		case <-ctx.Done():
			break OUTCOMES_CYCLE
		default:
		}
	}

	return outcomes, nil
}

func scanForDigest(states []grammar.Status) (digest []byte, ok bool, err error) {
	for _, c := range states {
		if c.Tag == kinds.TransferDigestKey {
			parsed, err := hex.DecodeString(c.Arg.Item())
			return parsed, true, err
		}
	}
	return
}

func (*TransOrdinal) collectionTempDigests(tempLevels map[string][]byte, key, item []byte) {
	eventSeq := retrieveEventSeqFromKey(key)

	//
	itemClone := make([]byte, len(item))
	copy(itemClone, item)

	tempLevels[string(itemClone)+eventSeq] = itemClone
}

//
//
//
//
//
func (txi *TransOrdinal) align(
	ctx context.Context,
	c grammar.Status,
	beginKeyBz []byte,
	screenedDigests map[string][]byte,
	initialRun bool,
	levelDetails LevelDetails,
) map[string][]byte {
	//
	//
	if !initialRun && len(screenedDigests) == 0 {
		return screenedDigests
	}

	tempDigests := make(map[string][]byte)

	switch c.Op {
	case grammar.TEq:
		it, err := dbm.RecursePrefix(txi.depot, beginKeyBz)
		if err != nil {
			panic(err)
		}
		defer it.End()

	EQUAL_CYCLE:
		for ; it.Sound(); it.Following() {

			//
			//
			key := it.Key()
			keyLevel, err := retrieveLevelFromKey(key)
			if err != nil {
				txi.log.Fault("REDACTED", err)
				continue
			}
			insideLimits, err := inspectLevelStates(levelDetails, keyLevel)
			if err != nil {
				txi.log.Fault("REDACTED", err)
				continue
			}
			if !insideLimits {
				continue
			}
			txi.collectionTempDigests(tempDigests, key, it.Item())
			//
			select {
			case <-ctx.Done():
				break EQUAL_CYCLE
			default:
			}
		}
		if err := it.Fault(); err != nil {
			panic(err)
		}

	case grammar.TPresent:
		//
		//
		it, err := dbm.RecursePrefix(txi.depot, beginKey(c.Tag))
		if err != nil {
			panic(err)
		}
		defer it.End()

	PRESENT_CYCLE:
		for ; it.Sound(); it.Following() {
			key := it.Key()
			keyLevel, err := retrieveLevelFromKey(key)
			if err != nil {
				txi.log.Fault("REDACTED", err)
				continue
			}
			insideLimits, err := inspectLevelStates(levelDetails, keyLevel)
			if err != nil {
				txi.log.Fault("REDACTED", err)
				continue
			}
			if !insideLimits {
				continue
			}
			txi.collectionTempDigests(tempDigests, key, it.Item())

			//
			select {
			case <-ctx.Done():
				break PRESENT_CYCLE
			default:
			}
		}
		if err := it.Fault(); err != nil {
			panic(err)
		}

	case grammar.TIncludes:
		//
		//
		//
		it, err := dbm.RecursePrefix(txi.depot, beginKey(c.Tag))
		if err != nil {
			panic(err)
		}
		defer it.End()

	INCLUDES_CYCLE:
		for ; it.Sound(); it.Following() {
			if !isMarkerKey(it.Key()) {
				continue
			}

			if strings.Contains(retrieveItemFromKey(it.Key()), c.Arg.Item()) {
				key := it.Key()
				keyLevel, err := retrieveLevelFromKey(key)
				if err != nil {
					txi.log.Fault("REDACTED", err)
					continue
				}
				insideLimits, err := inspectLevelStates(levelDetails, keyLevel)
				if err != nil {
					txi.log.Fault("REDACTED", err)
					continue
				}
				if !insideLimits {
					continue
				}
				txi.collectionTempDigests(tempDigests, key, it.Item())
			}

			//
			select {
			case <-ctx.Done():
				break INCLUDES_CYCLE
			default:
			}
		}
		if err := it.Fault(); err != nil {
			panic(err)
		}
	default:
		panic("REDACTED")
	}

	if len(tempDigests) == 0 || initialRun {
		//
		//
		//
		//
		//
		//
		//
		return tempDigests
	}

	//
	//
DELETE_CYCLE:
	for k, v := range screenedDigests {
		tempDigest := tempDigests[k]
		if tempDigest == nil || !bytes.Equal(tempDigest, v) {
			delete(screenedDigests, k)

			//
			select {
			case <-ctx.Done():
				break DELETE_CYCLE
			default:
			}
		}
	}

	return screenedDigests
}

//
//
//
//
//
func (txi *TransOrdinal) alignSpan(
	ctx context.Context,
	qr ordinaler.InquireSpan,
	beginKey []byte,
	screenedDigests map[string][]byte,
	initialRun bool,
	levelDetails LevelDetails,
) map[string][]byte {
	//
	//
	if !initialRun && len(screenedDigests) == 0 {
		return screenedDigests
	}

	tempDigests := make(map[string][]byte)

	it, err := dbm.RecursePrefix(txi.depot, beginKey)
	if err != nil {
		panic(err)
	}
	defer it.End()
	largeIntegerItem := new(big.Int)

Cycle:
	for ; it.Sound(); it.Following() {
		//
		//
		//
		key := it.Key()
		if !isMarkerKey(key) {
			continue
		}

		if _, ok := qr.AnyLimited().(*big.Float); ok {
			item := retrieveItemFromKey(key)
			v, ok := largeIntegerItem.SetString(item, 10)
			var vF *big.Float
			if !ok {
				vF, _, err = big.ParseFloat(item, 10, 125, big.ToNearestEven)
				if err != nil {
					continue Cycle
				}

			}
			if qr.Key != kinds.TransferLevelKey {
				keyLevel, err := retrieveLevelFromKey(key)
				if err != nil {
					txi.log.Fault("REDACTED", err)
					continue
				}
				insideLimits, err := inspectLevelStates(levelDetails, keyLevel)
				if err != nil {
					txi.log.Fault("REDACTED", err)
					continue
				}
				if !insideLimits {
					continue
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
				txi.log.Fault("REDACTED", err)
			} else if insideLimits {
				txi.collectionTempDigests(tempDigests, key, it.Item())
			}

			//
			//
			//
			//
			//
			//
		}

		//
		select {
		case <-ctx.Done():
			break Cycle
		default:
		}
	}
	if err := it.Fault(); err != nil {
		panic(err)
	}

	if len(tempDigests) == 0 || initialRun {
		//
		//
		//
		//
		//
		//
		//
		return tempDigests
	}

	//
	//
DELETE_CYCLE:
	for k, v := range screenedDigests {
		tempDigest := tempDigests[k]
		if tempDigest == nil || !bytes.Equal(tempDigests[k], v) {
			delete(screenedDigests, k)

			//
			select {
			case <-ctx.Done():
				break DELETE_CYCLE
			default:
			}
		}
	}

	return screenedDigests
}

//

func isMarkerKey(key []byte) bool {
	//
	//
	//
	//
	countMarkers := 0
	for i := 0; i < len(key); i++ {
		if key[i] == markerKeyDelimiterRune {
			countMarkers++
			if countMarkers >= 3 {
				return true
			}
		}
	}
	return false
}

func retrieveLevelFromKey(key []byte) (int64, error) {
	//
	//
	terminatePlace := bytes.LastIndexByte(key, markerKeyDelimiterRune)
	if terminatePlace == -1 {
		return 0, errors.New("REDACTED")
	}

	//
	beginPlace := bytes.LastIndexByte(key[:terminatePlace-1], markerKeyDelimiterRune)
	if beginPlace == -1 {
		return 0, errors.New("REDACTED")
	}

	//
	level, err := strconv.ParseInt(string(key[beginPlace+1:terminatePlace]), 10, 64)
	if err != nil {
		return 0, err
	}
	return level, nil
}

func retrieveItemFromKey(key []byte) string {
	//
	var ordinals []int
	for i, b := range key {
		if b == markerKeyDelimiterRune {
			ordinals = append(ordinals, i)
		}
	}

	//
	if len(ordinals) < 2 {
		return "REDACTED"
	}

	//
	item := key[ordinals[0]+1 : ordinals[len(ordinals)-2]]

	//
	item = bytes.TrimSpace(item)

	//
	return string(item)
}

func retrieveEventSeqFromKey(key []byte) string {
	segments := strings.Split(string(key), markerKeyDelimiter)

	finalElement := segments[len(segments)-1]

	if strings.Contains(finalElement, eventSeqDelimiter) {
		return strings.SplitN(finalElement, eventSeqDelimiter, 2)[1]
	}
	return "REDACTED"
}

func keyForEvent(key string, item string, outcome *iface.TransOutcome, eventSeq int64) []byte {
	return []byte(fmt.Sprintf("REDACTED",
		key,
		item,
		outcome.Level,
		outcome.Ordinal,
		eventSeqDelimiter+strconv.FormatInt(eventSeq, 10),
	))
}

func keyForLevel(outcome *iface.TransOutcome) []byte {
	return []byte(fmt.Sprintf("REDACTED",
		kinds.TransferLevelKey,
		outcome.Level,
		outcome.Level,
		outcome.Ordinal,
		//
		//
		eventSeqDelimiter+"REDACTED",
	))
}

func beginKeyForStatus(c grammar.Status, level int64) []byte {
	if level > 0 {
		return beginKey(c.Tag, c.Arg.Item(), level)
	}
	return beginKey(c.Tag, c.Arg.Item())
}

func beginKey(attributes ...interface{}) []byte {
	var b bytes.Buffer
	for _, f := range attributes {
		b.Write([]byte(fmt.Sprintf("REDACTED", f) + markerKeyDelimiter))
	}
	return b.Bytes()
}
