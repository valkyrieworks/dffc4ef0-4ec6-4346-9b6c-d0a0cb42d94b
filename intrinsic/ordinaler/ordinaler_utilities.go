package ordinaler

import (
	"fmt"
	"math/big"

	"github.com/valkyrieworks/status/ordinaler"
)

//
//
func contrastFloat(op1 *big.Float, op2 any) (int, bool, error) {
	switch actValue := op2.(type) {
	case *big.Int:
		vF := new(big.Float)
		vF.SetInt(actValue)
		cmp := op1.Cmp(vF)
		return cmp, false, nil

	case *big.Float:
		return op1.Cmp(actValue), true, nil
	default:
		return -1, false, fmt.Errorf("REDACTED", op2)
	}
}

//
//
//
//
//
func contrastInteger(op1 *big.Int, op2 any) (int, bool, error) {
	switch actValue := op2.(type) {
	case *big.Int:
		return op1.Cmp(actValue), false, nil
	case *big.Float:
		vF := new(big.Float)
		vF.SetInt(op1)
		return vF.Cmp(actValue), true, nil
	default:
		return -1, false, fmt.Errorf("REDACTED", op2)
	}
}

func InspectLimits(spans ordinaler.InquireSpan, v any) (bool, error) {
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
	//
	lesserLimited := spans.LesserLimitedItem()
	upperLimited := spans.UpperLimitedItem()

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
	switch vValue := v.(type) {
	case *big.Int:
		if lesserLimited != nil {
			cmp, isFloat, err := contrastInteger(vValue, lesserLimited)
			if err != nil {
				return false, err
			}
			if cmp == -1 || (isFloat && cmp == 0 && !spans.EncompassLesserLimited) {
				return false, nil
			}
		}
		if upperLimited != nil {
			cmp, isFloat, err := contrastInteger(vValue, upperLimited)
			if err != nil {
				return false, err
			}
			if cmp == 1 || (isFloat && cmp == 0 && !spans.EncompassUpperLimited) {
				return false, nil
			}
		}

	case *big.Float:
		if lesserLimited != nil {
			cmp, isFloat, err := contrastFloat(vValue, lesserLimited)
			if err != nil {
				return false, err
			}
			if cmp == -1 || (cmp == 0 && isFloat && !spans.EncompassLesserLimited) {
				return false, nil
			}
		}
		if upperLimited != nil {
			cmp, isFloat, err := contrastFloat(vValue, upperLimited)
			if err != nil {
				return false, err
			}
			if cmp == 1 || (cmp == 0 && isFloat && !spans.EncompassUpperLimited) {
				return false, nil
			}
		}

	default:
		return false, fmt.Errorf("REDACTED", v)
	}
	return true, nil
}
