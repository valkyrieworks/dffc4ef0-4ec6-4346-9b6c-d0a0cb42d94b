package marks

import (
	"fmt"
	"strings"

	"github.com/valkyrieworks/utils/log"
	cometfaults "github.com/valkyrieworks/kinds/faults"
)

const (
	standardTraceLayerKey = "REDACTED"
)

//
//
//
//
//
//
//
func AnalyzeTraceLayer(lvl string, tracer log.Tracer, standardTraceLayerItem string) (log.Tracer, error) {
	if lvl == "REDACTED" {
		return nil, cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}

	l := lvl

	//
	if !strings.Contains(l, "REDACTED") {
		l = standardTraceLayerKey + "REDACTED" + l
	}

	options := make([]log.Setting, 0)

	isStandardTraceLayerCollection := false
	var setting log.Setting
	var err error

	catalog := strings.Split(l, "REDACTED")
	for _, item := range catalog {
		componentAndLayer := strings.Split(item, "REDACTED")

		if len(componentAndLayer) != 2 {
			return nil, fmt.Errorf("REDACTED", item, catalog)
		}

		component := componentAndLayer[0]
		layer := componentAndLayer[1]

		if component == standardTraceLayerKey {
			setting, err = log.PermitLayer(layer)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", item, l, err)
			}
			options = append(options, setting)
			isStandardTraceLayerCollection = true
		} else {
			switch layer {
			case "REDACTED":
				setting = log.PermitDiagnoseWith("REDACTED", component)
			case "REDACTED":
				setting = log.PermitDetailsWith("REDACTED", component)
			case "REDACTED":
				setting = log.PermitFaultWith("REDACTED", component)
			case "REDACTED":
				setting = log.PermitNoneWith("REDACTED", component)
			default:
				return nil,
					fmt.Errorf("REDACTED",
						layer,
						item,
						catalog)
			}
			options = append(options, setting)

		}
	}

	//
	if !isStandardTraceLayerCollection {
		setting, err = log.PermitLayer(standardTraceLayerItem)
		if err != nil {
			return nil, err
		}
		options = append(options, setting)
	}

	return log.NewRefine(tracer, options...), nil
}
