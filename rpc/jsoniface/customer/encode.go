package customer

import (
	"fmt"
	"net/url"
	"reflect"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

func argumentsTowardWebrouteItems(arguments map[string]any) (url.Values, error) {
	items := make(url.Values)
	if len(arguments) == 0 {
		return items, nil
	}

	err := argumentsTowardJSN(arguments)
	if err != nil {
		return nil, err
	}

	for key, val := range arguments {
		items.Set(key, val.(string))
	}

	return items, nil
}

func argumentsTowardJSN(arguments map[string]any) error {
	for k, v := range arguments {
		rt := reflect.TypeOf(v)
		equalsOctetSection := rt.Kind() == reflect.Slice && rt.Elem().Kind() == reflect.Uint8
		if equalsOctetSection {
			octets := reflect.ValueOf(v).Bytes()
			arguments[k] = fmt.Sprintf("REDACTED", octets)
			continue
		}

		data, err := strongmindjson.Serialize(v)
		if err != nil {
			return err
		}
		arguments[k] = string(data)
	}
	return nil
}
