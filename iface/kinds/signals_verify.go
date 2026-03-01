package kinds

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

func VerifySerializeJSN(t *testing.T) {
	b, err := json.Marshal(&InvokeTransferOutcome{Cipher: 1})
	assert.NoError(t, err)
	//
	assert.True(t, strings.Contains(string(b), "REDACTED"))
	r1 := ReplyInspectTransfer{
		Cipher:      1,
		Data:      []byte("REDACTED"),
		FuelDesired: 43,
		Incidents: []Incident{
			{
				Kind: "REDACTED",
				Properties: []IncidentProperty{
					{Key: "REDACTED", Datum: "REDACTED"},
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

func VerifyPersistFetchArtifactPlain(t *testing.T) {
	scenarios := []proto.Message{
		&SolicitReverberate{
			Signal: "REDACTED",
		},
	}

	for _, c := range scenarios {
		buf := new(bytes.Buffer)
		err := PersistArtifact(c, buf)
		assert.Nil(t, err)

		msg := new(SolicitReverberate)
		err = FetchArtifact(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}

func VerifyPersistFetchArtifact(t *testing.T) {
	scenarios := []proto.Message{
		&commitchema.Heading{
			Altitude:  4,
			SuccessionUUID: "REDACTED",
		},
		//
	}

	for _, c := range scenarios {
		buf := new(bytes.Buffer)
		err := PersistArtifact(c, buf)
		assert.Nil(t, err)

		msg := new(commitchema.Heading)
		err = FetchArtifact(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}

func VerifyPersistFetchSignal2(t *testing.T) {
	expression := "REDACTED"
	scenarios := []proto.Message{
		&ReplyInspectTransfer{
			Data:      []byte(expression),
			Log:       expression,
			FuelDesired: 10,
			Incidents: []Incident{
				{
					Kind: "REDACTED",
					Properties: []IncidentProperty{
						{Key: "REDACTED", Datum: "REDACTED"},
					},
				},
			},
		},
		//
	}

	for _, c := range scenarios {
		buf := new(bytes.Buffer)
		err := PersistArtifact(c, buf)
		assert.Nil(t, err)

		msg := new(ReplyInspectTransfer)
		err = FetchArtifact(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}
