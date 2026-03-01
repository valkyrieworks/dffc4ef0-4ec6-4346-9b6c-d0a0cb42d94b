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

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	catalogutil "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var _ ordinalizer.LedgerOrdinalizer = (*PreventerOrdinalizer)(nil)

//
//
//
type PreventerOrdinalizer struct {
	depot dbm.DB

	//
	//
	incidentOrder int64
	log      log.Tracer
}

func New(depot dbm.DB) *PreventerOrdinalizer {
	return &PreventerOrdinalizer{
		depot: depot,
	}
}

func (idx *PreventerOrdinalizer) AssignTracer(l log.Tracer) {
	idx.log = l
}

//
//
func (idx *PreventerOrdinalizer) Has(altitude int64) (bool, error) {
	key, err := altitudeToken(altitude)
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
func (idx *PreventerOrdinalizer) Ordinal(bh kinds.IncidentDataFreshLedgerIncidents) error {
	cluster := idx.depot.FreshCluster()
	defer cluster.Shutdown()

	altitude := bh.Altitude

	//
	key, err := altitudeToken(altitude)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := cluster.Set(key, integer64towOctets(altitude)); err != nil {
		return err
	}

	//
	if err := idx.positionIncidents(cluster, bh.Incidents, altitude); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return cluster.PersistChronize()
}

//
//
//
//
//
func (idx *PreventerOrdinalizer) Lookup(ctx context.Context, q *inquire.Inquire) ([]int64, error) {
	outcomes := make([]int64, 0)
	select {
	case <-ctx.Done():
		return outcomes, nil

	default:
	}

	terms := q.Grammar()

	//
	omitPositions := make([]int, 0)

	var ok bool

	var altitudeDetails AltitudeDetails
	//
	//
	terms, altitudeDetails, ok = deconflictAltitude(terms)

	//
	//
	extents, scopePositions, altitudeScope := ordinalizer.ScanForeachExtentsUsingAltitude(terms)
	altitudeDetails.altitudeScope = altitudeScope

	//
	//
	//
	//
	//
	//
	if ok && altitudeDetails.solelyAltitudeEqual {
		ok, err := idx.Has(altitudeDetails.altitude)
		if err != nil {
			return nil, err
		}

		if ok {
			return []int64{altitudeDetails.altitude}, nil
		}

		return outcomes, nil
	}

	var elevationsStarted bool
	screenedElevations := make(map[string][]byte)
	if altitudeDetails.altitudeEqualOffset != -1 {
		omitPositions = append(omitPositions, altitudeDetails.altitudeEqualOffset)
	}

	if len(extents) > 0 {
		omitPositions = append(omitPositions, scopePositions...)

		for _, qr := range extents {
			//
			//
			//
			//
			if qr.Key == kinds.LedgerAltitudeToken && !altitudeDetails.solelyAltitudeScope {
				//
				//
				//
				//

				continue
			}
			heading, err := orderedcode.Append(nil, qr.Key)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", err)
			}

			if !elevationsStarted {
				screenedElevations, err = idx.alignScope(ctx, qr, heading, screenedElevations, true, altitudeDetails)
				if err != nil {
					return nil, err
				}

				elevationsStarted = true

				//
				//
				if len(screenedElevations) == 0 {
					break
				}
			} else {
				screenedElevations, err = idx.alignScope(ctx, qr, heading, screenedElevations, false, altitudeDetails)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	//
	for i, c := range terms {
		if integerInsideSection(i, omitPositions) {
			continue
		}

		initiateToken, err := orderedcode.Append(nil, c.Tag, c.Arg.Datum())
		if err != nil {
			return nil, err
		}

		if !elevationsStarted {
			screenedElevations, err = idx.align(ctx, c, initiateToken, screenedElevations, true, altitudeDetails)
			if err != nil {
				return nil, err
			}

			elevationsStarted = true

			//
			//
			if len(screenedElevations) == 0 {
				break
			}
		} else {
			screenedElevations, err = idx.align(ctx, c, initiateToken, screenedElevations, false, altitudeDetails)
			if err != nil {
				return nil, err
			}
		}
	}

	//
	outcomes = make([]int64, 0, len(screenedElevations))
	outcomeIndex := make(map[int64]struct{})

FOREACH_CYCLE:
	for _, hBz := range screenedElevations {
		h := integer64fromOctets(hBz)

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
			break FOREACH_CYCLE

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
func (idx *PreventerOrdinalizer) alignScope(
	ctx context.Context,
	qr ordinalizer.InquireScope,
	initiateToken []byte,
	screenedElevations map[string][]byte,
	initialExecute bool,
	altitudeDetails AltitudeDetails,
) (map[string][]byte, error) {
	//
	//
	if !initialExecute && len(screenedElevations) == 0 {
		return screenedElevations, nil
	}

	scratchElevations := make(map[string][]byte)

	it, err := dbm.TraverseHeading(idx.depot, initiateToken)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	defer it.Shutdown()

Cycle:
	for ; it.Sound(); it.Following() {
		var (
			incidentDatum string
			err        error
		)

		if qr.Key == kinds.LedgerAltitudeToken {
			incidentDatum, err = analyzeDatumOriginatingLeadingToken(it.Key())
		} else {
			incidentDatum, err = analyzeDatumOriginatingIncidentToken(it.Key())
		}

		if err != nil {
			continue
		}

		if _, ok := qr.SomeRestricted().(*big.Float); ok {
			v := new(big.Int)
			v, ok := v.SetString(incidentDatum, 10)
			var vF *big.Float
			if !ok {
				//
				//
				vF, _, err = big.ParseFloat(incidentDatum, 10, 125, big.ToNearestEven)
				if err != nil {
					continue Cycle
				}
			}

			if qr.Key != kinds.LedgerAltitudeToken {
				tokenAltitude, err := analyzeAltitudeOriginatingIncidentToken(it.Key())
				if err != nil {
					idx.log.Failure("REDACTED", err)
					continue Cycle
				}
				insideAltitude, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
				if err != nil {
					idx.log.Failure("REDACTED", err)
					continue Cycle
				}
				if !insideAltitude {
					continue Cycle
				}
			}

			var insideLimits bool
			var err error
			if !ok {
				insideLimits, err = catalogutil.InspectLimits(qr, vF)
			} else {
				insideLimits, err = catalogutil.InspectLimits(qr, v)
			}
			if err != nil {
				idx.log.Failure("REDACTED", err)
			} else if insideLimits {
				idx.assignScratchElevations(scratchElevations, it)
			}
		}

		select {
		case <-ctx.Done():
			break Cycle
		default:
		}
	}

	if err := it.Failure(); err != nil {
		return nil, err
	}

	if len(scratchElevations) == 0 || initialExecute {
		//
		//
		//
		//
		//
		//
		//
		return scratchElevations, nil
	}

	//
	//
FOREACH_CYCLE:
	for k, v := range screenedElevations {
		scratchAltitude := scratchElevations[k]

		//
		//
		if scratchAltitude == nil || !bytes.Equal(scratchAltitude, v) {
			delete(screenedElevations, k)

			select {
			case <-ctx.Done():
				break FOREACH_CYCLE
			default:
			}
		}
	}

	return screenedElevations, nil
}

func (idx *PreventerOrdinalizer) assignScratchElevations(scratchElevations map[string][]byte, it dbm.Traverser) {
	//
	//
	incidentOrder, _ := analyzeIncidentOrderOriginatingIncidentToken(it.Key())

	//
	datum := make([]byte, len(it.Datum()))
	copy(datum, it.Datum())

	scratchElevations[string(datum)+strconv.FormatInt(incidentOrder, 10)] = datum
}

//
//
//
//
//
//
func (idx *PreventerOrdinalizer) align(
	ctx context.Context,
	c grammar.Stipulation,
	initiateTokenByz []byte,
	screenedElevations map[string][]byte,
	initialExecute bool,
	altitudeDetails AltitudeDetails,
) (map[string][]byte, error) {
	//
	//
	if !initialExecute && len(screenedElevations) == 0 {
		return screenedElevations, nil
	}

	scratchElevations := make(map[string][]byte)

	switch c.Op {
	case grammar.TEq:
		it, err := dbm.TraverseHeading(idx.depot, initiateTokenByz)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		defer it.Shutdown()

		for ; it.Sound(); it.Following() {

			tokenAltitude, err := analyzeAltitudeOriginatingIncidentToken(it.Key())
			if err != nil {
				idx.log.Failure("REDACTED", err)
				continue
			}
			insideAltitude, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
			if err != nil {
				idx.log.Failure("REDACTED", err)
				continue
			}
			if !insideAltitude {
				continue
			}

			idx.assignScratchElevations(scratchElevations, it)

			if err := ctx.Err(); err != nil {
				break
			}
		}

		if err := it.Failure(); err != nil {
			return nil, err
		}

	case grammar.TYPPresent:
		heading, err := orderedcode.Append(nil, c.Tag)
		if err != nil {
			return nil, err
		}

		it, err := dbm.TraverseHeading(idx.depot, heading)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		defer it.Shutdown()

	CYCLE_PRESENT:
		for ; it.Sound(); it.Following() {

			tokenAltitude, err := analyzeAltitudeOriginatingIncidentToken(it.Key())
			if err != nil {
				idx.log.Failure("REDACTED", err)
				continue
			}
			insideAltitude, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
			if err != nil {
				idx.log.Failure("REDACTED", err)
				continue
			}
			if !insideAltitude {
				continue
			}

			idx.assignScratchElevations(scratchElevations, it)

			select {
			case <-ctx.Done():
				break CYCLE_PRESENT

			default:
			}
		}

		if err := it.Failure(); err != nil {
			return nil, err
		}

	case grammar.TYPIncludes:
		heading, err := orderedcode.Append(nil, c.Tag)
		if err != nil {
			return nil, err
		}

		it, err := dbm.TraverseHeading(idx.depot, heading)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		defer it.Shutdown()

	CYCLE_INCLUDES:
		for ; it.Sound(); it.Following() {
			incidentDatum, err := analyzeDatumOriginatingIncidentToken(it.Key())
			if err != nil {
				continue
			}

			if strings.Contains(incidentDatum, c.Arg.Datum()) {
				tokenAltitude, err := analyzeAltitudeOriginatingIncidentToken(it.Key())
				if err != nil {
					idx.log.Failure("REDACTED", err)
					continue
				}
				insideAltitude, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
				if err != nil {
					idx.log.Failure("REDACTED", err)
					continue
				}
				if !insideAltitude {
					continue
				}
				idx.assignScratchElevations(scratchElevations, it)
			}

			select {
			case <-ctx.Done():
				break CYCLE_INCLUDES

			default:
			}
		}
		if err := it.Failure(); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("REDACTED")
	}

	if len(scratchElevations) == 0 || initialExecute {
		//
		//
		//
		//
		//
		//
		//
		return scratchElevations, nil
	}

	//
	//
FOREACH_CYCLE:
	for k, v := range screenedElevations {
		scratchAltitude := scratchElevations[k]
		if scratchAltitude == nil || !bytes.Equal(scratchAltitude, v) {
			delete(screenedElevations, k)

			select {
			case <-ctx.Done():
				break FOREACH_CYCLE

			default:
			}
		}
	}

	return screenedElevations, nil
}

func (idx *PreventerOrdinalizer) positionIncidents(cluster dbm.Cluster, incidents []iface.Incident, altitude int64) error {
	altitudeByz := integer64towOctets(altitude)

	for _, incident := range incidents {
		idx.incidentOrder += 1
		//
		if len(incident.Kind) == 0 {
			continue
		}

		for _, property := range incident.Properties {
			if len(property.Key) == 0 {
				continue
			}

			//
			complexToken := incident.Kind + "REDACTED" + property.Key
			if complexToken == kinds.LedgerAltitudeToken {
				return fmt.Errorf("REDACTED", complexToken)
			}

			if property.ObtainOrdinal() {
				key, err := incidentToken(complexToken, property.Datum, altitude, idx.incidentOrder)
				if err != nil {
					return fmt.Errorf("REDACTED", err)
				}

				if err := cluster.Set(key, altitudeByz); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
