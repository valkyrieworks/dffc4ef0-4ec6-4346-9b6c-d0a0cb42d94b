package octets

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
func VerifySerialize(t *testing.T) {
	bz := []byte("REDACTED")
	dataBYTE := HexadecimalOctets(bz)
	bz2, err := dataBYTE.Serialize()
	assert.Nil(t, err)
	assert.Equal(t, bz, bz2)

	var dataByte2 HexadecimalOctets
	err = (&dataByte2).Decode(bz)
	assert.Nil(t, err)
	assert.Equal(t, dataBYTE, dataByte2)
}

//
func VerifyJSNSerialize(t *testing.T) {
	type VerifyLayout struct {
		B1 []byte
		B2 HexadecimalOctets
	}

	scenarios := []struct {
		influx    []byte
		anticipated string
	}{
		{[]byte("REDACTED"), "REDACTED"},
		{[]byte("REDACTED"), "REDACTED"},
		{[]byte("REDACTED"), "REDACTED"},
	}

	for i, tc := range scenarios {

		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			ts := VerifyLayout{B1: tc.influx, B2: tc.influx}

			//
			jsnOctets, err := json.Marshal(ts)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, string(jsnOctets), tc.anticipated)

			//

			//
			ts2 := VerifyLayout{}
			err = json.Unmarshal(jsnOctets, &ts2)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, ts2.B1, tc.influx)
			assert.Equal(t, ts2.B2, HexadecimalOctets(tc.influx))
		})
	}
}

func Testhexoctets_Text(t *testing.T) {
	hs := HexadecimalOctets([]byte("REDACTED"))
	if _, err := strconv.ParseInt(hs.Text(), 16, 64); err != nil {
		t.Fatal(err)
	}
}
