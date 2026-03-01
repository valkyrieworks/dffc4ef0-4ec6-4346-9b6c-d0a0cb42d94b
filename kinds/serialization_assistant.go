package kinds

import (
	gogotypes "github.com/cosmos/gogoproto/types"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
)

//
//
func codecSerialize(record any) []byte {
	if record != nil && !equalsClassedVoid(record) && !equalsBlank(record) {
		switch record := record.(type) {
		case string:
			i := gogotypes.StringValue{
				Value: record,
			}
			bz, err := i.Marshal()
			if err != nil {
				return nil
			}
			return bz
		case int64:
			i := gogotypes.Int64Value{
				Value: record,
			}
			bz, err := i.Marshal()
			if err != nil {
				return nil
			}
			return bz
		case octets.HexadecimalOctets:
			i := gogotypes.BytesValue{
				Value: record,
			}
			bz, err := i.Marshal()
			if err != nil {
				return nil
			}
			return bz
		default:
			return nil
		}
	}

	return nil
}
