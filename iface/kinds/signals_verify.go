package kinds

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"

	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

func VerifySerializeJSON(t *testing.T) {
	b, err := json.Marshal(&InvokeTransferOutcome{Code: 1})
	assert.NoError(t, err)
	//
	assert.True(t, strings.Contains(string(b), "REDACTED"))
	r1 := ReplyInspectTransfer{
		Code:      1,
		Data:      []byte("REDACTED"),
		FuelDesired: 43,
		Events: []Event{
			{
				Kind: "REDACTED",
				Properties: []EventProperty{
					{Key: "REDACTED", Item: "REDACTED"},
				},
			},
		},
	}
	b, err = json.Marshal(&r1)
	assert.Nil(t, err)

	var r2 ReplyInspectTransfer
	err = json.Unmarshal(b, &r2)
	assert.Nil(t, err)
	assert.Equal(t, r1, r2)
}

func VerifyRecordFetchSignalBasic(t *testing.T) {
	scenarios := []proto.Message{
		&QueryReverberate{
			Signal: "REDACTED",
		},
	}

	for _, c := range scenarios {
		buf := new(bytes.Buffer)
		err := RecordSignal(c, buf)
		assert.Nil(t, err)

		msg := new(QueryReverberate)
		err = FetchSignal(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}

func VerifyRecordFetchSignal(t *testing.T) {
	scenarios := []proto.Message{
		&engineproto.Heading{
			Level:  4,
			LedgerUID: "REDACTED",
		},
		//
	}

	for _, c := range scenarios {
		buf := new(bytes.Buffer)
		err := RecordSignal(c, buf)
		assert.Nil(t, err)

		msg := new(engineproto.Heading)
		err = FetchSignal(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}

func VerifyRecordFetchSignal2(t *testing.T) {
	expression := "REDACTED"
	scenarios := []proto.Message{
		&ReplyInspectTransfer{
			Data:      []byte(expression),
			Log:       expression,
			FuelDesired: 10,
			Events: []Event{
				{
					Kind: "REDACTED",
					Properties: []EventProperty{
						{Key: "REDACTED", Item: "REDACTED"},
					},
				},
			},
		},
		//
	}

	for _, c := range scenarios {
		buf := new(bytes.Buffer)
		err := RecordSignal(c, buf)
		assert.Nil(t, err)

		msg := new(ReplyInspectTransfer)
		err = FetchSignal(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}
