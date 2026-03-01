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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	"github.com/cosmos/gogoproto/proto"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	catalogutil "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	labelTokenDelimiter     = "REDACTED"
	labelTokenDelimiterCharacter = '/'
	incidentOrderDelimiter   = "REDACTED"
)

var _ transferordinal.TransferOrdinalizer = (*TransferOrdinal)(nil)

//
type TransferOrdinal struct {
	depot dbm.DB
	//
	incidentOrder int64

	log log.Tracer
}

//
func FreshTransferOrdinal(depot dbm.DB) *TransferOrdinal {
	return &TransferOrdinal{
		depot: depot,
	}
}

func (txi *TransferOrdinal) AssignTracer(l log.Tracer) {
	txi.log = l
}

//
//
func (txi *TransferOrdinal) Get(digest []byte) (*iface.TransferOutcome, error) {
	if len(digest) == 0 {
		return nil, transferordinal.FailureBlankDigest
	}

	crudeOctets, err := txi.depot.Get(digest)
	if err != nil {
		panic(err)
	}
	if crudeOctets == nil {
		return nil, nil
	}

	transferOutcome := new(iface.TransferOutcome)
	err = proto.Unmarshal(crudeOctets, transferOutcome)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return transferOutcome, nil
}

//
//
//
//
func (txi *TransferOrdinal) AppendCluster(b *transferordinal.Cluster) error {
	depotCluster := txi.depot.FreshCluster()
	defer depotCluster.Shutdown()

	for _, outcome := range b.Ops {
		digest := kinds.Tx(outcome.Tx).Digest()

		//
		err := txi.positionIncidents(outcome, digest, depotCluster)
		if err != nil {
			return err
		}

		//
		err = depotCluster.Set(tokenForeachAltitude(outcome), digest)
		if err != nil {
			return err
		}

		crudeOctets, err := proto.Marshal(outcome)
		if err != nil {
			return err
		}
		//
		err = depotCluster.Set(digest, crudeOctets)
		if err != nil {
			return err
		}
	}

	return depotCluster.PersistChronize()
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
func (txi *TransferOrdinal) Ordinal(outcome *iface.TransferOutcome) error {
	b := txi.depot.FreshCluster()
	defer b.Shutdown()

	digest := kinds.Tx(outcome.Tx).Digest()

	if !outcome.Outcome.EqualsOKAY() {
		agedOutcome, err := txi.Get(digest)
		if err != nil {
			return err
		}

		//
		//
		if agedOutcome != nil && agedOutcome.Outcome.Cipher == iface.CipherKindOKAY {
			return nil
		}
	}

	//
	err := txi.positionIncidents(outcome, digest, b)
	if err != nil {
		return err
	}

	//
	err = b.Set(tokenForeachAltitude(outcome), digest)
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

	return b.PersistChronize()
}

func (txi *TransferOrdinal) positionIncidents(outcome *iface.TransferOutcome, digest []byte, depot dbm.Cluster) error {
	for _, incident := range outcome.Outcome.Incidents {
		txi.incidentOrder += 1
		//
		if len(incident.Kind) == 0 {
			continue
		}

		for _, property := range incident.Properties {
			if len(property.Key) == 0 {
				continue
			}

			//
			complexLabel := fmt.Sprintf("REDACTED", incident.Kind, property.Key)
			//
			if complexLabel == kinds.TransferDigestToken || complexLabel == kinds.TransferAltitudeToken {
				return fmt.Errorf("REDACTED", complexLabel)
			}
			if property.ObtainOrdinal() {
				err := depot.Set(tokenForeachIncident(complexLabel, property.Datum, outcome, txi.incidentOrder), digest)
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
func (txi *TransferOrdinal) Lookup(ctx context.Context, q *inquire.Inquire) ([]*iface.TransferOutcome, error) {
	select {
	case <-ctx.Done():
		return make([]*iface.TransferOutcome, 0), nil

	default:
	}

	var digestsStarted bool
	screenedDigests := make(map[string][]byte)

	//
	terms := q.Grammar()

	//
	digest, ok, err := scanForeachDigest(terms)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	} else if ok {
		res, err := txi.Get(digest)
		switch {
		case err != nil:
			return []*iface.TransferOutcome{}, fmt.Errorf("REDACTED", err)
		case res == nil:
			return []*iface.TransferOutcome{}, nil
		default:
			return []*iface.TransferOutcome{res}, nil
		}
	}

	//
	omitCatalogs := make([]int, 0)
	var altitudeDetails AltitudeDetails

	//
	//
	terms, altitudeDetails = deconflictAltitude(terms)

	if !altitudeDetails.solelyAltitudeEqual {
		omitCatalogs = append(omitCatalogs, altitudeDetails.altitudeEqualOffset)
	}

	//
	//
	//
	extents, scopeCatalogs, altitudeScope := ordinalizer.ScanForeachExtentsUsingAltitude(terms)
	altitudeDetails.altitudeScope = altitudeScope
	if len(extents) > 0 {
		omitCatalogs = append(omitCatalogs, scopeCatalogs...)

		for _, qr := range extents {

			//
			//
			//
			//
			if qr.Key == kinds.TransferAltitudeToken && !altitudeDetails.solelyAltitudeScope {
				continue
			}
			if !digestsStarted {
				screenedDigests = txi.alignScope(ctx, qr, initiateToken(qr.Key), screenedDigests, true, altitudeDetails)
				digestsStarted = true

				//
				//
				if len(screenedDigests) == 0 {
					break
				}
			} else {
				screenedDigests = txi.alignScope(ctx, qr, initiateToken(qr.Key), screenedDigests, false, altitudeDetails)
			}
		}
	}

	//

	//
	for i, c := range terms {
		if integerInsideSection(i, omitCatalogs) {
			continue
		}

		if !digestsStarted {
			screenedDigests = txi.align(ctx, c, initiateTokenForeachStipulation(c, altitudeDetails.altitude), screenedDigests, true, altitudeDetails)
			digestsStarted = true

			//
			//
			if len(screenedDigests) == 0 {
				break
			}
		} else {
			screenedDigests = txi.align(ctx, c, initiateTokenForeachStipulation(c, altitudeDetails.altitude), screenedDigests, false, altitudeDetails)
		}
	}

	outcomes := make([]*iface.TransferOutcome, 0, len(screenedDigests))
	outcomeIndex := make(map[string]struct{})
OUTCOMES_CYCLE:
	for _, h := range screenedDigests {

		res, err := txi.Get(h)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", h, err)
		}
		digestText := string(h)
		if _, ok := outcomeIndex[digestText]; !ok {
			outcomeIndex[digestText] = struct{}{}
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

func scanForeachDigest(terms []grammar.Stipulation) (digest []byte, ok bool, err error) {
	for _, c := range terms {
		if c.Tag == kinds.TransferDigestToken {
			deserialized, err := hex.DecodeString(c.Arg.Datum())
			return deserialized, true, err
		}
	}
	return
}

func (*TransferOrdinal) collectionScratchDigests(scratchElevations map[string][]byte, key, datum []byte) {
	incidentOrder := deriveIncidentOrderOriginatingToken(key)

	//
	datumDuplicate := make([]byte, len(datum))
	copy(datumDuplicate, datum)

	scratchElevations[string(datumDuplicate)+incidentOrder] = datumDuplicate
}

//
//
//
//
//
func (txi *TransferOrdinal) align(
	ctx context.Context,
	c grammar.Stipulation,
	initiateTokenByz []byte,
	screenedDigests map[string][]byte,
	initialExecute bool,
	altitudeDetails AltitudeDetails,
) map[string][]byte {
	//
	//
	if !initialExecute && len(screenedDigests) == 0 {
		return screenedDigests
	}

	scratchDigests := make(map[string][]byte)

	switch c.Op {
	case grammar.TEq:
		it, err := dbm.TraverseHeading(txi.depot, initiateTokenByz)
		if err != nil {
			panic(err)
		}
		defer it.Shutdown()

	EQUAL_CYCLE:
		for ; it.Sound(); it.Following() {

			//
			//
			key := it.Key()
			tokenAltitude, err := deriveAltitudeOriginatingToken(key)
			if err != nil {
				txi.log.Failure("REDACTED", err)
				continue
			}
			insideLimits, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
			if err != nil {
				txi.log.Failure("REDACTED", err)
				continue
			}
			if !insideLimits {
				continue
			}
			txi.collectionScratchDigests(scratchDigests, key, it.Datum())
			//
			select {
			case <-ctx.Done():
				break EQUAL_CYCLE
			default:
			}
		}
		if err := it.Failure(); err != nil {
			panic(err)
		}

	case grammar.TYPPresent:
		//
		//
		it, err := dbm.TraverseHeading(txi.depot, initiateToken(c.Tag))
		if err != nil {
			panic(err)
		}
		defer it.Shutdown()

	PRESENT_CYCLE:
		for ; it.Sound(); it.Following() {
			key := it.Key()
			tokenAltitude, err := deriveAltitudeOriginatingToken(key)
			if err != nil {
				txi.log.Failure("REDACTED", err)
				continue
			}
			insideLimits, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
			if err != nil {
				txi.log.Failure("REDACTED", err)
				continue
			}
			if !insideLimits {
				continue
			}
			txi.collectionScratchDigests(scratchDigests, key, it.Datum())

			//
			select {
			case <-ctx.Done():
				break PRESENT_CYCLE
			default:
			}
		}
		if err := it.Failure(); err != nil {
			panic(err)
		}

	case grammar.TYPIncludes:
		//
		//
		//
		it, err := dbm.TraverseHeading(txi.depot, initiateToken(c.Tag))
		if err != nil {
			panic(err)
		}
		defer it.Shutdown()

	INCLUDES_CYCLE:
		for ; it.Sound(); it.Following() {
			if !equalsLabelToken(it.Key()) {
				continue
			}

			if strings.Contains(deriveDatumOriginatingToken(it.Key()), c.Arg.Datum()) {
				key := it.Key()
				tokenAltitude, err := deriveAltitudeOriginatingToken(key)
				if err != nil {
					txi.log.Failure("REDACTED", err)
					continue
				}
				insideLimits, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
				if err != nil {
					txi.log.Failure("REDACTED", err)
					continue
				}
				if !insideLimits {
					continue
				}
				txi.collectionScratchDigests(scratchDigests, key, it.Datum())
			}

			//
			select {
			case <-ctx.Done():
				break INCLUDES_CYCLE
			default:
			}
		}
		if err := it.Failure(); err != nil {
			panic(err)
		}
	default:
		panic("REDACTED")
	}

	if len(scratchDigests) == 0 || initialExecute {
		//
		//
		//
		//
		//
		//
		//
		return scratchDigests
	}

	//
	//
DISCARD_CYCLE:
	for k, v := range screenedDigests {
		scratchDigest := scratchDigests[k]
		if scratchDigest == nil || !bytes.Equal(scratchDigest, v) {
			delete(screenedDigests, k)

			//
			select {
			case <-ctx.Done():
				break DISCARD_CYCLE
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
func (txi *TransferOrdinal) alignScope(
	ctx context.Context,
	qr ordinalizer.InquireScope,
	initiateToken []byte,
	screenedDigests map[string][]byte,
	initialExecute bool,
	altitudeDetails AltitudeDetails,
) map[string][]byte {
	//
	//
	if !initialExecute && len(screenedDigests) == 0 {
		return screenedDigests
	}

	scratchDigests := make(map[string][]byte)

	it, err := dbm.TraverseHeading(txi.depot, initiateToken)
	if err != nil {
		panic(err)
	}
	defer it.Shutdown()
	ampleIntegerDatum := new(big.Int)

Cycle:
	for ; it.Sound(); it.Following() {
		//
		//
		//
		key := it.Key()
		if !equalsLabelToken(key) {
			continue
		}

		if _, ok := qr.SomeRestricted().(*big.Float); ok {
			datum := deriveDatumOriginatingToken(key)
			v, ok := ampleIntegerDatum.SetString(datum, 10)
			var vF *big.Float
			if !ok {
				vF, _, err = big.ParseFloat(datum, 10, 125, big.ToNearestEven)
				if err != nil {
					continue Cycle
				}

			}
			if qr.Key != kinds.TransferAltitudeToken {
				tokenAltitude, err := deriveAltitudeOriginatingToken(key)
				if err != nil {
					txi.log.Failure("REDACTED", err)
					continue
				}
				insideLimits, err := inspectAltitudeTerms(altitudeDetails, tokenAltitude)
				if err != nil {
					txi.log.Failure("REDACTED", err)
					continue
				}
				if !insideLimits {
					continue
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
				txi.log.Failure("REDACTED", err)
			} else if insideLimits {
				txi.collectionScratchDigests(scratchDigests, key, it.Datum())
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
	if err := it.Failure(); err != nil {
		panic(err)
	}

	if len(scratchDigests) == 0 || initialExecute {
		//
		//
		//
		//
		//
		//
		//
		return scratchDigests
	}

	//
	//
DISCARD_CYCLE:
	for k, v := range screenedDigests {
		scratchDigest := scratchDigests[k]
		if scratchDigest == nil || !bytes.Equal(scratchDigests[k], v) {
			delete(screenedDigests, k)

			//
			select {
			case <-ctx.Done():
				break DISCARD_CYCLE
			default:
			}
		}
	}

	return screenedDigests
}

//

func equalsLabelToken(key []byte) bool {
	//
	//
	//
	//
	countLabels := 0
	for i := 0; i < len(key); i++ {
		if key[i] == labelTokenDelimiterCharacter {
			countLabels++
			if countLabels >= 3 {
				return true
			}
		}
	}
	return false
}

func deriveAltitudeOriginatingToken(key []byte) (int64, error) {
	//
	//
	terminatePlace := bytes.LastIndexByte(key, labelTokenDelimiterCharacter)
	if terminatePlace == -1 {
		return 0, errors.New("REDACTED")
	}

	//
	initiatePlace := bytes.LastIndexByte(key[:terminatePlace-1], labelTokenDelimiterCharacter)
	if initiatePlace == -1 {
		return 0, errors.New("REDACTED")
	}

	//
	altitude, err := strconv.ParseInt(string(key[initiatePlace+1:terminatePlace]), 10, 64)
	if err != nil {
		return 0, err
	}
	return altitude, nil
}

func deriveDatumOriginatingToken(key []byte) string {
	//
	var catalogs []int
	for i, b := range key {
		if b == labelTokenDelimiterCharacter {
			catalogs = append(catalogs, i)
		}
	}

	//
	if len(catalogs) < 2 {
		return "REDACTED"
	}

	//
	datum := key[catalogs[0]+1 : catalogs[len(catalogs)-2]]

	//
	datum = bytes.TrimSpace(datum)

	//
	return string(datum)
}

func deriveIncidentOrderOriginatingToken(key []byte) string {
	fragments := strings.Split(string(key), labelTokenDelimiter)

	finalElement := fragments[len(fragments)-1]

	if strings.Contains(finalElement, incidentOrderDelimiter) {
		return strings.SplitN(finalElement, incidentOrderDelimiter, 2)[1]
	}
	return "REDACTED"
}

func tokenForeachIncident(key string, datum string, outcome *iface.TransferOutcome, incidentOrder int64) []byte {
	return []byte(fmt.Sprintf("REDACTED",
		key,
		datum,
		outcome.Altitude,
		outcome.Ordinal,
		incidentOrderDelimiter+strconv.FormatInt(incidentOrder, 10),
	))
}

func tokenForeachAltitude(outcome *iface.TransferOutcome) []byte {
	return []byte(fmt.Sprintf("REDACTED",
		kinds.TransferAltitudeToken,
		outcome.Altitude,
		outcome.Altitude,
		outcome.Ordinal,
		//
		//
		incidentOrderDelimiter+"REDACTED",
	))
}

func initiateTokenForeachStipulation(c grammar.Stipulation, altitude int64) []byte {
	if altitude > 0 {
		return initiateToken(c.Tag, c.Arg.Datum(), altitude)
	}
	return initiateToken(c.Tag, c.Arg.Datum())
}

func initiateToken(areas ...interface{}) []byte {
	var b bytes.Buffer
	for _, f := range areas {
		b.Write([]byte(fmt.Sprintf("REDACTED", f) + labelTokenDelimiter))
	}
	return b.Bytes()
}
