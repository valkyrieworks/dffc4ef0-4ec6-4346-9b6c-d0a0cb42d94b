package ordinalizer

import (
	"fmt"
	"math/big"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
)

//
//
func contrastDecimal(op1 *big.Float, op2 any) (int, bool, error) {
	switch actionItem := op2.(type) {
	case *big.Int:
		vF := new(big.Float)
		vF.SetInt(actionItem)
		cmp := op1.Cmp(vF)
		return cmp, false, nil

	case *big.Float:
		return op1.Cmp(actionItem), true, nil
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
	switch actionItem := op2.(type) {
	case *big.Int:
		return op1.Cmp(actionItem), false, nil
	case *big.Float:
		vF := new(big.Float)
		vF.SetInt(op1)
		return vF.Cmp(actionItem), true, nil
	default:
		return -1, false, fmt.Errorf("REDACTED", op2)
	}
}

func InspectLimits(extents ordinalizer.InquireScope, v any) (bool, error) {
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
	lesserRestricted := extents.LesserRestrictedDatum()
	higherRestricted := extents.HigherRestrictedDatum()

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
	switch verItem := v.(type) {
	case *big.Int:
		if lesserRestricted != nil {
			cmp, equalsDecimal, err := contrastInteger(verItem, lesserRestricted)
			if err != nil {
				return false, err
			}
			if cmp == -1 || (equalsDecimal && cmp == 0 && !extents.EncompassLesserRestricted) {
				return false, nil
			}
		}
		if higherRestricted != nil {
			cmp, equalsDecimal, err := contrastInteger(verItem, higherRestricted)
			if err != nil {
				return false, err
			}
			if cmp == 1 || (equalsDecimal && cmp == 0 && !extents.EncompassHigherRestricted) {
				return false, nil
			}
		}

	case *big.Float:
		if lesserRestricted != nil {
			cmp, equalsDecimal, err := contrastDecimal(verItem, lesserRestricted)
			if err != nil {
				return false, err
			}
			if cmp == -1 || (cmp == 0 && equalsDecimal && !extents.EncompassLesserRestricted) {
				return false, nil
			}
		}
		if higherRestricted != nil {
			cmp, equalsDecimal, err := contrastDecimal(verItem, higherRestricted)
			if err != nil {
				return false, err
			}
			if cmp == 1 || (cmp == 0 && equalsDecimal && !extents.EncompassHigherRestricted) {
				return false, nil
			}
		}

	default:
		return false, fmt.Errorf("REDACTED", v)
	}
	return true, nil
}
