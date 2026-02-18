package record_test

import (
	"encoding/hex"
	"math"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/blockreplication"
	bcschema "github.com/valkyrieworks/schema/consensuscore/blockreplication"
	"github.com/valkyrieworks/kinds"
)

func TestBcBlockRequestMessageValidateBasic(t *testing.T) {
	testCases := []struct {
		testName      string
		requestHeight int64
		expectErr     bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			request := bcschema.BlockRequest{Height: tc.requestHeight}
			assert.Equal(t, tc.expectErr, blockreplication.ValidateMsg(&request) != nil, "REDACTED")
		})
	}
}

func TestBcNoBlockResponseMessageValidateBasic(t *testing.T) {
	testCases := []struct {
		testName          string
		nonResponseHeight int64
		expectErr         bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			nonResponse := bcschema.NoBlockResponse{Height: tc.nonResponseHeight}
			assert.Equal(t, tc.expectErr, blockreplication.ValidateMsg(&nonResponse) != nil, "REDACTED")
		})
	}
}

func TestBcStatusRequestMessageValidateBasic(t *testing.T) {
	request := bcschema.StatusRequest{}
	assert.NoError(t, blockreplication.ValidateMsg(&request))
}

func TestBcStatusResponseMessageValidateBasic(t *testing.T) {
	testCases := []struct {
		testName       string
		responseHeight int64
		expectErr      bool
	}{
		{"REDACTED", 0, false},
		{"REDACTED", 1, false},
		{"REDACTED", -1, true},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			response := bcschema.StatusResponse{Height: tc.responseHeight}
			assert.Equal(t, tc.expectErr, blockreplication.ValidateMsg(&response) != nil, "REDACTED")
		})
	}
}

//
func TestBlocksyncMessageVectors(t *testing.T) {
	block := kinds.MakeBlock(int64(3), []kinds.Tx{kinds.Tx("REDACTED")}, nil, nil)
	block.Version.Block = 11 //

	bpb, err := block.ToProto()
	require.NoError(t, err)

	testCases := []struct {
		testName string
		bmsg     proto.Message
		expBytes string
	}{
		{"REDACTED", &bcschema.Message{Sum: &bcschema.Message_BlockRequest{
			BlockRequest: &bcschema.BlockRequest{Height: 1},
		}}, "REDACTED"},
		{
			"REDACTED", &bcschema.Message{Sum: &bcschema.Message_BlockRequest{
				BlockRequest: &bcschema.BlockRequest{Height: math.MaxInt64},
			}},
			"REDACTED",
		},
		{"REDACTED", &bcschema.Message{Sum: &bcschema.Message_BlockResponse{
			BlockResponse: &bcschema.BlockResponse{Block: bpb},
		}}, "REDACTED"},
		{"REDACTED", &bcschema.Message{Sum: &bcschema.Message_NoBlockResponse{
			NoBlockResponse: &bcschema.NoBlockResponse{Height: 1},
		}}, "REDACTED"},
		{
			"REDACTED", &bcschema.Message{Sum: &bcschema.Message_NoBlockResponse{
				NoBlockResponse: &bcschema.NoBlockResponse{Height: math.MaxInt64},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &bcschema.Message{Sum: &bcschema.Message_StatusRequest{
				StatusRequest: &bcschema.StatusRequest{},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &bcschema.Message{Sum: &bcschema.Message_StatusResponse{
				StatusResponse: &bcschema.StatusResponse{Height: 1, Base: 2},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &bcschema.Message{Sum: &bcschema.Message_StatusResponse{
				StatusResponse: &bcschema.StatusResponse{Height: math.MaxInt64, Base: math.MaxInt64},
			}},
			"REDACTED",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.testName, func(t *testing.T) {
			bz, _ := proto.Marshal(tc.bmsg)

			require.Equal(t, tc.expBytes, hex.EncodeToString(bz))
		})
	}
}
