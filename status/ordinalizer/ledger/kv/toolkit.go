package kv

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"

	"github.com/google/orderedcode"

	catalogutil "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

type AltitudeDetails struct {
	altitudeScope     ordinalizer.InquireScope
	altitude          int64
	altitudeEqualOffset     int
	solelyAltitudeScope bool
	solelyAltitudeEqual    bool
}

func integerInsideSection(a int, catalog []int) bool {
	for _, b := range catalog {
		if b == a {
			return true
		}
	}

	return false
}

func integer64fromOctets(bz []byte) int64 {
	v, _ := binary.Varint(bz)
	return v
}

func integer64towOctets(i int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, i)
	return buf[:n]
}

func altitudeToken(altitude int64) ([]byte, error) {
	return orderedcode.Append(
		nil,
		kinds.LedgerAltitudeToken,
		altitude,
	)
}

func incidentToken(complexToken, incidentDatum string, altitude int64, incidentOrder int64) ([]byte, error) {
	return orderedcode.Append(
		nil,
		complexToken,
		incidentDatum,
		altitude,
		incidentOrder,
	)
}

func analyzeDatumOriginatingLeadingToken(key []byte) (string, error) {
	var (
		complexToken string
		altitude       int64
	)

	pending, err := orderedcode.Parse(string(key), &complexToken, &altitude)
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", err)
	}

	if len(pending) != 0 {
		return "REDACTED", fmt.Errorf("REDACTED", pending)
	}

	return strconv.FormatInt(altitude, 10), nil
}

func analyzeDatumOriginatingIncidentToken(key []byte) (string, error) {
	var (
		complexToken, incidentDatum string
		altitude                   int64
	)

	_, err := orderedcode.Parse(string(key), &complexToken, &incidentDatum, &altitude)
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", err)
	}

	return incidentDatum, nil
}

func analyzeAltitudeOriginatingIncidentToken(key []byte) (int64, error) {
	var (
		complexToken, incidentDatum string
		altitude                   int64
	)

	_, err := orderedcode.Parse(string(key), &complexToken, &incidentDatum, &altitude)
	if err != nil {
		return -1, fmt.Errorf("REDACTED", err)
	}

	return altitude, nil
}

func analyzeIncidentOrderOriginatingIncidentToken(key []byte) (int64, error) {
	var (
		complexToken, incidentDatum string
		altitude                   int64
		incidentOrder                 int64
	)

	pending, err := orderedcode.Parse(string(key), &complexToken, &incidentDatum, &altitude)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}

	//
	//
	//
	//
	//
	//

	if len(pending) == 0 { //
		return 0, fmt.Errorf("REDACTED")
	}
	var typ string
	pending2, err := orderedcode.Parse(pending, &typ) //
	if err != nil {                                       //
		pending, fault2 := orderedcode.Parse(string(key), &complexToken, &incidentDatum, &altitude, &incidentOrder)
		if fault2 != nil || len(pending) != 0 { //
			return 0, fmt.Errorf("REDACTED", err, fault2)
		}
	} else if len(pending2) != 0 { //
		pending, fault2 := orderedcode.Parse(pending2, &incidentOrder) //
		//
		//
		if fault2 != nil || len(pending) != 0 { //
			return 0, fmt.Errorf("REDACTED", fault2)
		}
	}
	return incidentOrder, nil
}

//
//
//
//
func deconflictAltitude(terms []grammar.Stipulation) (deconflictTerms []grammar.Stipulation, altitudeDetails AltitudeDetails, detected bool) {
	altitudeDetails.altitudeEqualOffset = -1
	altitudeScopePresent := false
	var altitudeStipulation []grammar.Stipulation
	altitudeDetails.solelyAltitudeEqual = true
	altitudeDetails.solelyAltitudeScope = true
	for _, c := range terms {
		if c.Tag == kinds.LedgerAltitudeToken {
			if c.Op == grammar.TEq {
				if detected || altitudeScopePresent {
					continue
				}
				hashDecimal := c.Arg.Numeral()
				if hashDecimal != nil {
					h, _ := hashDecimal.Int64()
					altitudeDetails.altitude = h
					altitudeStipulation = append(altitudeStipulation, c)
					detected = true
				}
			} else {
				altitudeDetails.solelyAltitudeEqual = false
				altitudeScopePresent = true
				deconflictTerms = append(deconflictTerms, c)
			}
		} else {
			altitudeDetails.solelyAltitudeScope = false
			altitudeDetails.solelyAltitudeEqual = false
			deconflictTerms = append(deconflictTerms, c)
		}
	}
	if !altitudeScopePresent && len(altitudeStipulation) != 0 {
		altitudeDetails.altitudeEqualOffset = len(deconflictTerms)
		altitudeDetails.solelyAltitudeScope = false
		deconflictTerms = append(deconflictTerms, altitudeStipulation...)
	} else {
		//
		//
		altitudeDetails.altitudeEqualOffset = -1
		altitudeDetails.altitude = 0
		altitudeDetails.solelyAltitudeEqual = false
		detected = false
	}
	return deconflictTerms, altitudeDetails, detected
}

func inspectAltitudeTerms(altitudeDetails AltitudeDetails, tokenAltitude int64) (bool, error) {
	if altitudeDetails.altitudeScope.Key != "REDACTED" {
		insideLimits, err := catalogutil.InspectLimits(altitudeDetails.altitudeScope, big.NewInt(tokenAltitude))
		if err != nil || !insideLimits {
			return false, err
		}
	} else if altitudeDetails.altitude != 0 && tokenAltitude != altitudeDetails.altitude {
		return false, nil
	}

	return true, nil
}
