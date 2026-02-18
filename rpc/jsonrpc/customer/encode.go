package customer

import (
	"fmt"
	"net/url"
	"reflect"

	cometjson "github.com/valkyrieworks/utils/json"
)

func argsToURLItems(args map[string]any) (url.Values, error) {
	items := make(url.Values)
	if len(args) == 0 {
		return items, nil
	}

	err := argsToJSON(args)
	if err != nil {
		return nil, err
	}

	for key, val := range args {
		items.Set(key, val.(string))
	}

	return items, nil
}

func argsToJSON(args map[string]any) error {
	for k, v := range args {
		rt := reflect.TypeOf(v)
		isOctetSection := rt.Kind() == reflect.Slice && rt.Elem().Kind() == reflect.Uint8
		if isOctetSection {
			octets := reflect.ValueOf(v).Bytes()
			args[k] = fmt.Sprintf("REDACTED", octets)
			continue
		}

		data, err := cometjson.Serialize(v)
		if err != nil {
			return err
		}
		args[k] = string(data)
	}
	return nil
}
