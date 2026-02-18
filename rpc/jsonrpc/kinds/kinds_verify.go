package kinds

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpecimenOutcome struct {
	Item string
}

type replyVerify struct {
	id       jsonrpcuid
	anticipated string
}

var replyVerifies = []replyVerify{
	{JsonrpcStringUID("REDACTED"), "REDACTED"},
	{JsonrpcStringUID("REDACTED"), "REDACTED"},
	{JsonrpcStringUID("REDACTED"), "REDACTED"},
	{JsonrpcStringUID("REDACTED"), "REDACTED"},
	{JsonrpcIntegerUID(-1), "REDACTED"},
	{JsonrpcIntegerUID(0), "REDACTED"},
	{JsonrpcIntegerUID(1), "REDACTED"},
	{JsonrpcIntegerUID(100), "REDACTED"},
}

func VerifyReplies(t *testing.T) {
	affirm := assert.New(t)
	for _, tt := range replyVerifies {
		jsonuid := tt.id
		a := NewRPCSuccessReply(jsonuid, &SpecimenOutcome{"REDACTED"})
		b, _ := json.Marshal(a)
		s := fmt.Sprintf("REDACTED", tt.anticipated)
		assert.Equal(s, string(b))

		d := RPCAnalyzeFault(errors.New("REDACTED"))
		e, _ := json.Marshal(d)
		f := "REDACTED"
		assert.Equal(f, string(e))

		g := RPCProcedureNegateLocatedFault(jsonuid)
		h, _ := json.Marshal(g)
		i := fmt.Sprintf("REDACTED", tt.anticipated)
		assert.Equal(string(h), i)
	}
}

func VerifyDecodeReplies(t *testing.T) {
	affirm := assert.New(t)
	for _, tt := range replyVerifies {
		reply := &RPCAnswer{}
		err := json.Unmarshal(
			[]byte(fmt.Sprintf("REDACTED", tt.anticipated)),
			reply,
		)
		assert.Nil(err)
		a := NewRPCSuccessReply(tt.id, &SpecimenOutcome{"REDACTED"})
		assert.Equal(*reply, a)
	}
	reply := &RPCAnswer{}
	err := json.Unmarshal([]byte("REDACTED"), reply)
	assert.NotNil(err)
}

func VerifyRPCFault(t *testing.T) {
	assert.Equal(t, "REDACTED",
		fmt.Sprintf("REDACTED", &RPCFault{
			Code:    12,
			Signal: "REDACTED",
			Data:    "REDACTED",
		}))

	assert.Equal(t, "REDACTED",
		fmt.Sprintf("REDACTED", &RPCFault{
			Code:    12,
			Signal: "REDACTED",
		}))
}
