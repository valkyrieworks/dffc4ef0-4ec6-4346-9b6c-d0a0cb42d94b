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
	dataBYTE := HexOctets(bz)
	bz2, err := dataBYTE.Serialize()
	assert.Nil(t, err)
	assert.Equal(t, bz, bz2)

	var dataC2 HexOctets
	err = (&dataC2).Unserialize(bz)
	assert.Nil(t, err)
	assert.Equal(t, dataBYTE, dataC2)
}

//
func VerifyJSONSerialize(t *testing.T) {
	type VerifyRecord struct {
		B1 []byte
		B2 HexOctets
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
			ts := VerifyRecord{B1: tc.influx, B2: tc.influx}

			//
			jsonOctets, err := json.Marshal(ts)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, string(jsonOctets), tc.anticipated)

			//

			//
			ts2 := VerifyRecord{}
			err = json.Unmarshal(jsonOctets, &ts2)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, ts2.B1, tc.influx)
			assert.Equal(t, ts2.B2, HexOctets(tc.influx))
		})
	}
}

func Verifyhexoctets_String(t *testing.T) {
	hs := HexOctets([]byte("REDACTED"))
	if _, err := strconv.ParseInt(hs.String(), 16, 64); err != nil {
		t.Fatal(err)
	}
}
