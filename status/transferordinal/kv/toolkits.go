package kv

import (
	"fmt"
	"math/big"

	catalogutil "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/ordinalizer"
	cmtgrammar "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/google/orderedcode"
)

type AltitudeDetails struct {
	altitudeScope     ordinalizer.InquireScope
	altitude          int64
	altitudeEqualOffset     int
	solelyAltitudeScope bool
	solelyAltitudeEqual    bool
}

//
func integerInsideSection(a int, catalog []int) bool {
	for _, b := range catalog {
		if b == a {
			return true
		}
	}
	return false
}

func AnalyzeIncidentOrderOriginatingIncidentToken(key []byte) (int64, error) {
	var (
		complexToken, typ, incidentDatum string
		altitude                        int64
		incidentOrder                      int64
	)

	lingering, err := orderedcode.Parse(string(key), &complexToken, &incidentDatum, &altitude, &typ, &incidentOrder)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}

	if len(lingering) != 0 {
		return 0, fmt.Errorf("REDACTED", lingering)
	}

	return incidentOrder, nil
}

func deconflictAltitude(terms []cmtgrammar.Stipulation) (deconflictTerms []cmtgrammar.Stipulation, altitudeDetails AltitudeDetails) {
	altitudeDetails.altitudeEqualOffset = -1
	altitudeScopePresent := false
	detected := false
	var altitudeStipulation []cmtgrammar.Stipulation
	altitudeDetails.solelyAltitudeEqual = true
	altitudeDetails.solelyAltitudeScope = true
	for _, c := range terms {
		if c.Tag == kinds.TransferAltitudeToken {
			if c.Op == cmtgrammar.TEq {
				if altitudeScopePresent || detected {
					continue
				}
				hashDecimal := c.Arg.Numeral()
				if hashDecimal != nil {
					h, _ := hashDecimal.Int64()
					altitudeDetails.altitude = h
					detected = true
					altitudeStipulation = append(altitudeStipulation, c)
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
	}
	return deconflictTerms, altitudeDetails
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
