package host

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/utils/octets"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

func VerifyAnalyzeJSONIndex(t *testing.T) {
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
		"REDACTED":  &octets.HexOctets{},
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
		Item  any `json:"item"`
		Level any `json:"level"`
	}{
		Level: &tmp,
		Item:  &octets.HexOctets{},
	}
	err = json.Unmarshal(influx, &p3)
	if assert.Nil(t, err) {
		h, ok := p3.Level.(*int)
		if assert.True(t, ok, "REDACTED", p3.Level) {
			assert.Equal(t, 22, *h)
		}
		v, ok := p3.Item.(*octets.HexOctets)
		if assert.True(t, ok, "REDACTED", p3.Item) {
			assert.EqualValues(t, []byte{0x12, 0x34}, *v)
		}
	}

	//
	p4 := struct {
		Item  octets.HexOctets `json:"item"`
		Level int            `json:"level"`
	}{}
	err = json.Unmarshal(influx, &p4)
	if assert.Nil(t, err) {
		assert.EqualValues(t, 22, p4.Level)
		assert.EqualValues(t, []byte{0x12, 0x34}, p4.Item)
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

		var v octets.HexOctets
		err = json.Unmarshal(*p5["REDACTED"], &v)
		if assert.Nil(t, err) {
			assert.Equal(t, octets.HexOctets{0x12, 0x34}, v)
		}
	}
}

func VerifyAnalyzeJSONList(t *testing.T) {
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
	p2 := []any{&octets.HexOctets{}, &tmp}
	err = json.Unmarshal(influx, &p2)
	if assert.Nil(t, err) {
		v, ok := p2[0].(*octets.HexOctets)
		if assert.True(t, ok, "REDACTED", p2[0]) {
			assert.EqualValues(t, []byte{0x12, 0x34}, *v)
		}
		h, ok := p2[1].(*int)
		if assert.True(t, ok, "REDACTED", p2[1]) {
			assert.EqualValues(t, 22, *h)
		}
	}
}

func VerifyAnalyzeJsonrpc(t *testing.T) {
	demo := func(ctx *kinds.Context, level int, label string) {}
	invocation := NewRPCFunction(demo, "REDACTED")

	scenarios := []struct {
		raw    string
		level int64
		label   string
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
		values, err := jsonOptionsToArgs(invocation, data)
		if tc.abort {
			assert.NotNil(t, err, i)
		} else {
			assert.Nil(t, err, "REDACTED", i, err)
			if assert.Equal(t, 2, len(values), i) {
				assert.Equal(t, tc.level, values[0].Int(), i)
				assert.Equal(t, tc.label, values[1].String(), i)
			}
		}

	}
}

func VerifyAnalyzeURI(t *testing.T) {
	demo := func(ctx *kinds.Context, level int, label string) {}
	invocation := NewRPCFunction(demo, "REDACTED")

	scenarios := []struct {
		raw    []string
		level int64
		label   string
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
		values, err := httpOptionsToArgs(invocation, req)
		if tc.abort {
			assert.NotNil(t, err, i)
		} else {
			assert.Nil(t, err, "REDACTED", i, err)
			if assert.Equal(t, 2, len(values), i) {
				assert.Equal(t, tc.level, values[0].Int(), i)
				assert.Equal(t, tc.label, values[1].String(), i)
			}
		}

	}
}
