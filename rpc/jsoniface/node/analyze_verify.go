package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

func VerifyAnalyzeJSNIndex(t *testing.T) {
	influx := []byte("REDACTED")

	//
	var p1 map[string]any
	err := json.Unmarshal(influx, &p1)
	if assert.Nil(t, err) {
		h, ok := p1["REDACTED"].(float64)
		if assert.True(t, ok, "REDACTED", p1["REDACTED"]) {
			assert.EqualValues(t, 22, h)
		}
		v, ok := p1["REDACTED"].(string)
		if assert.True(t, ok, "REDACTED", p1["REDACTED"]) {
			assert.EqualValues(t, "REDACTED", v)
		}
	}

	//
	tmp := 0
	p2 := map[string]any{
		"REDACTED":  &octets.HexadecimalOctets{},
		"REDACTED": &tmp,
	}
	err = json.Unmarshal(influx, &p2)
	if assert.Nil(t, err) {
		h, ok := p2["REDACTED"].(float64)
		if assert.True(t, ok, "REDACTED", p2["REDACTED"]) {
			assert.EqualValues(t, 22, h)
		}
		v, ok := p2["REDACTED"].(string)
		if assert.True(t, ok, "REDACTED", p2["REDACTED"]) {
			assert.EqualValues(t, "REDACTED", v)
		}
	}

	//
	//
	tmp = 0
	p3 := struct {
		Datum  any `json:"datum"`
		Altitude any `json:"altitude"`
	}{
		Altitude: &tmp,
		Datum:  &octets.HexadecimalOctets{},
	}
	err = json.Unmarshal(influx, &p3)
	if assert.Nil(t, err) {
		h, ok := p3.Altitude.(*int)
		if assert.True(t, ok, "REDACTED", p3.Altitude) {
			assert.Equal(t, 22, *h)
		}
		v, ok := p3.Datum.(*octets.HexadecimalOctets)
		if assert.True(t, ok, "REDACTED", p3.Datum) {
			assert.EqualValues(t, []byte{0x12, 0x34}, *v)
		}
	}

	//
	p4 := struct {
		Datum  octets.HexadecimalOctets `json:"datum"`
		Altitude int            `json:"altitude"`
	}{}
	err = json.Unmarshal(influx, &p4)
	if assert.Nil(t, err) {
		assert.EqualValues(t, 22, p4.Altitude)
		assert.EqualValues(t, []byte{0x12, 0x34}, p4.Datum)
	}

	//
	//
	var p5 map[string]*json.RawMessage
	err = json.Unmarshal(influx, &p5)
	if assert.Nil(t, err) {
		var h int
		err = json.Unmarshal(*p5["REDACTED"], &h)
		if assert.Nil(t, err) {
			assert.Equal(t, 22, h)
		}

		var v octets.HexadecimalOctets
		err = json.Unmarshal(*p5["REDACTED"], &v)
		if assert.Nil(t, err) {
			assert.Equal(t, octets.HexadecimalOctets{0x12, 0x34}, v)
		}
	}
}

func VerifyAnalyzeJSNSeries(t *testing.T) {
	influx := []byte("REDACTED")

	//
	var p1 []any
	err := json.Unmarshal(influx, &p1)
	if assert.Nil(t, err) {
		v, ok := p1[0].(string)
		if assert.True(t, ok, "REDACTED", p1[0]) {
			assert.EqualValues(t, "REDACTED", v)
		}
		h, ok := p1[1].(float64)
		if assert.True(t, ok, "REDACTED", p1[1]) {
			assert.EqualValues(t, 22, h)
		}
	}

	//
	tmp := 0
	p2 := []any{&octets.HexadecimalOctets{}, &tmp}
	err = json.Unmarshal(influx, &p2)
	if assert.Nil(t, err) {
		v, ok := p2[0].(*octets.HexadecimalOctets)
		if assert.True(t, ok, "REDACTED", p2[0]) {
			assert.EqualValues(t, []byte{0x12, 0x34}, *v)
		}
		h, ok := p2[1].(*int)
		if assert.True(t, ok, "REDACTED", p2[1]) {
			assert.EqualValues(t, 22, *h)
		}
	}
}

func VerifyAnalyzeJsoniface(t *testing.T) {
	prototype := func(ctx *kinds.Env, altitude int, alias string) {}
	invocation := FreshRemoteMethod(prototype, "REDACTED")

	scenarios := []struct {
		raw    string
		altitude int64
		alias   string
		abort   bool
	}{
		//
		{"REDACTED", 7, "REDACTED", false},
		{"REDACTED", 22, "REDACTED", false},
		//
		{"REDACTED", 0, "REDACTED", false},
		//
		{"REDACTED", 0, "REDACTED", true},
		{"REDACTED", 0, "REDACTED", true},
		{"REDACTED", 0, "REDACTED", true},
	}
	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		data := []byte(tc.raw)
		values, err := jsnParametersTowardArguments(invocation, data)
		if tc.abort {
			assert.NotNil(t, err, i)
		} else {
			assert.Nil(t, err, "REDACTED", i, err)
			if assert.Equal(t, 2, len(values), i) {
				assert.Equal(t, tc.altitude, values[0].Int(), i)
				assert.Equal(t, tc.alias, values[1].String(), i)
			}
		}

	}
}

func VerifyAnalyzeURL(t *testing.T) {
	prototype := func(ctx *kinds.Env, altitude int, alias string) {}
	invocation := FreshRemoteMethod(prototype, "REDACTED")

	scenarios := []struct {
		raw    []string
		altitude int64
		alias   string
		abort   bool
	}{
		//
		{[]string{"REDACTED", "REDACTED"}, 7, "REDACTED", false},
		{[]string{"REDACTED", "REDACTED"}, 22, "REDACTED", false},
		{[]string{"REDACTED", "REDACTED"}, -10, "REDACTED", false},
		//
		{[]string{"REDACTED", "REDACTED"}, 7, "REDACTED", false},
		{[]string{"REDACTED", "REDACTED"}, -10, "REDACTED", false},
		//
		{[]string{"REDACTED", "REDACTED"}, -10, "REDACTED", true},
	}
	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		//
		url := fmt.Sprintf(
			"REDACTED",
			tc.raw[0], tc.raw[1])
		req, err := http.NewRequest("REDACTED", url, nil)
		assert.NoError(t, err)
		values, err := httpsvcParametersTowardArguments(invocation, req)
		if tc.abort {
			assert.NotNil(t, err, i)
		} else {
			assert.Nil(t, err, "REDACTED", i, err)
			if assert.Equal(t, 2, len(values), i) {
				assert.Equal(t, tc.altitude, values[0].Int(), i)
				assert.Equal(t, tc.alias, values[1].String(), i)
			}
		}

	}
}
