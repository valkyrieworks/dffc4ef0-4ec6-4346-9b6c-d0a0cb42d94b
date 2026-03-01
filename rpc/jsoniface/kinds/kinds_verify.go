package kinds

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpecimenOutcome struct {
	Datum string
}

type replyVerify struct {
	id       jsonrpcuuid
	anticipated string
}

var replyVerifies = []replyVerify{
	{JsonifaceTextUUID("REDACTED"), "REDACTED"},
	{JsonifaceTextUUID("REDACTED"), "REDACTED"},
	{JsonifaceTextUUID("REDACTED"), "REDACTED"},
	{JsonifaceTextUUID("REDACTED"), "REDACTED"},
	{JsonifaceIntegerUUID(-1), "REDACTED"},
	{JsonifaceIntegerUUID(0), "REDACTED"},
	{JsonifaceIntegerUUID(1), "REDACTED"},
	{JsonifaceIntegerUUID(100), "REDACTED"},
}

func VerifyReplies(t *testing.T) {
	affirm := assert.New(t)
	for _, tt := range replyVerifies {
		jsonuuid := tt.id
		a := FreshRemoteTriumphReply(jsonuuid, &SpecimenOutcome{"REDACTED"})
		b, _ := json.Marshal(a)
		s := fmt.Sprintf("REDACTED", tt.anticipated)
		assert.Equal(s, string(b))

		d := RemoteAnalyzeFailure(errors.New("REDACTED"))
		e, _ := json.Marshal(d)
		f := "REDACTED"
		assert.Equal(f, string(e))

		g := RemoteProcedureNegationDetectedFailure(jsonuuid)
		h, _ := json.Marshal(g)
		i := fmt.Sprintf("REDACTED", tt.anticipated)
		assert.Equal(string(h), i)
	}
}

func VerifyDeformatReplies(t *testing.T) {
	affirm := assert.New(t)
	for _, tt := range replyVerifies {
		reply := &RemoteReply{}
		err := json.Unmarshal(
			[]byte(fmt.Sprintf("REDACTED", tt.anticipated)),
			reply,
		)
		assert.Nil(err)
		a := FreshRemoteTriumphReply(tt.id, &SpecimenOutcome{"REDACTED"})
		assert.Equal(*reply, a)
	}
	reply := &RemoteReply{}
	err := json.Unmarshal([]byte("REDACTED"), reply)
	assert.NotNil(err)
}

func VerifyRemoteFailure(t *testing.T) {
	assert.Equal(t, "REDACTED",
		fmt.Sprintf("REDACTED", &RemoteFailure{
			Cipher:    12,
			Signal: "REDACTED",
			Data:    "REDACTED",
		}))

	assert.Equal(t, "REDACTED",
		fmt.Sprintf("REDACTED", &RemoteFailure{
			Cipher:    12,
			Signal: "REDACTED",
		}))
}
