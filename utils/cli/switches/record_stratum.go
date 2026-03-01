package switches

import (
	"fmt"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
)

const (
	fallbackReportStratumToken = "REDACTED"
)

//
//
//
//
//
//
//
func AnalyzeRecordStratum(lvl string, tracer log.Tracer, fallbackReportStratumDatum string) (log.Tracer, error) {
	if lvl == "REDACTED" {
		return nil, strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}

	l := lvl

	//
	if !strings.Contains(l, "REDACTED") {
		l = fallbackReportStratumToken + "REDACTED" + l
	}

	choices := make([]log.Selection, 0)

	equalsFallbackReportStratumAssign := false
	var selection log.Selection
	var err error

	catalog := strings.Split(l, "REDACTED")
	for _, record := range catalog {
		componentAlsoStratum := strings.Split(record, "REDACTED")

		if len(componentAlsoStratum) != 2 {
			return nil, fmt.Errorf("REDACTED", record, catalog)
		}

		component := componentAlsoStratum[0]
		stratum := componentAlsoStratum[1]

		if component == fallbackReportStratumToken {
			selection, err = log.PermitStratum(stratum)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", record, l, err)
			}
			choices = append(choices, selection)
			equalsFallbackReportStratumAssign = true
		} else {
			switch stratum {
			case "REDACTED":
				selection = log.PermitDiagnoseUsing("REDACTED", component)
			case "REDACTED":
				selection = log.PermitDetailsUsing("REDACTED", component)
			case "REDACTED":
				selection = log.PermitFailureUsing("REDACTED", component)
			case "REDACTED":
				selection = log.PermitNilUsing("REDACTED", component)
			default:
				return nil,
					fmt.Errorf("REDACTED",
						stratum,
						record,
						catalog)
			}
			choices = append(choices, selection)

		}
	}

	//
	if !equalsFallbackReportStratumAssign {
		selection, err = log.PermitStratum(fallbackReportStratumDatum)
		if err != nil {
			return nil, err
		}
		choices = append(choices, selection)
	}

	return log.FreshRefine(tracer, choices...), nil
}
