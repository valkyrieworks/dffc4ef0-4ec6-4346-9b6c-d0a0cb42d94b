package kinds

import (
	"bytes"
	"encoding/json"

	"github.com/cosmos/gogoproto/jsonpb"
)

const (
	CodeKindSuccess uint32 = 0

	//
	//
	CodeKindReprocess uint32 = 32_000
)

//
func (r ReplyInspectTransfer) IsOK() bool {
	return r.Code == CodeKindSuccess
}

//
func (r ReplyInspectTransfer) IsErr() bool {
	return r.Code != CodeKindSuccess
}

//
func (r InvokeTransferOutcome) IsOK() bool {
	return r.Code == CodeKindSuccess
}

//
func (r InvokeTransferOutcome) IsErr() bool {
	return r.Code != CodeKindSuccess
}

//
func (r ReplyInquire) IsOK() bool {
	return r.Code == CodeKindSuccess
}

//
func (r ReplyInquire) IsErr() bool {
	return r.Code != CodeKindSuccess
}

//
func (r ReplyHandleNomination) IsApproved() bool {
	return r.Status == Responseprocessnomination_ALLOW
}

//
func (r ReplyHandleNomination) IsStateUnclear() bool {
	return r.Status == Responseprocessnomination_UNCLEAR
}

func (r ReplyValidateBallotAddition) IsApproved() bool {
	return r.Status == Responseverifyballotextension_ALLOW
}

//
func (r ReplyValidateBallotAddition) IsStateUnclear() bool {
	return r.Status == Responseverifyballotextension_UNCLEAR
}

//
//

var (
	jsonfmtFormatter = jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
	}
	jsonfmtParser = jsonpb.Unmarshaler{}
)

func (r *ReplyInspectTransfer) SerializeJSON() ([]byte, error) {
	s, err := jsonfmtFormatter.MarshalToString(r)
	return []byte(s), err
}

func (r *ReplyInspectTransfer) UnserializeJSON(b []byte) error {
	scanner := bytes.NewBuffer(b)
	return jsonfmtParser.Unmarshal(scanner, r)
}

func (r *InvokeTransferOutcome) SerializeJSON() ([]byte, error) {
	s, err := jsonfmtFormatter.MarshalToString(r)
	return []byte(s), err
}

func (r *InvokeTransferOutcome) UnserializeJSON(b []byte) error {
	scanner := bytes.NewBuffer(b)
	return jsonfmtParser.Unmarshal(scanner, r)
}

func (r *ReplyInquire) SerializeJSON() ([]byte, error) {
	s, err := jsonfmtFormatter.MarshalToString(r)
	return []byte(s), err
}

func (r *ReplyInquire) UnserializeJSON(b []byte) error {
	scanner := bytes.NewBuffer(b)
	return jsonfmtParser.Unmarshal(scanner, r)
}

func (r *ReplyEndorse) SerializeJSON() ([]byte, error) {
	s, err := jsonfmtFormatter.MarshalToString(r)
	return []byte(s), err
}

func (r *ReplyEndorse) UnserializeJSON(b []byte) error {
	scanner := bytes.NewBuffer(b)
	return jsonfmtParser.Unmarshal(scanner, r)
}

func (r *EventProperty) SerializeJSON() ([]byte, error) {
	s, err := jsonfmtFormatter.MarshalToString(r)
	return []byte(s), err
}

func (r *EventProperty) UnserializeJSON(b []byte) error {
	scanner := bytes.NewBuffer(b)
	return jsonfmtParser.Unmarshal(scanner, r)
}

//
//

//
//
type jsonEpochTripper interface {
	json.Marshaler
	json.Unmarshaler
}

var (
	_ jsonEpochTripper = (*ReplyEndorse)(nil)
	_ jsonEpochTripper = (*ReplyInquire)(nil)
	_ jsonEpochTripper = (*InvokeTransferOutcome)(nil)
	_ jsonEpochTripper = (*ReplyInspectTransfer)(nil)
)

var _ jsonEpochTripper = (*EventProperty)(nil)

//
//
func CertainInvokeTransferOutcome(reply *InvokeTransferOutcome) *InvokeTransferOutcome {
	return &InvokeTransferOutcome{
		Code:      reply.Code,
		Data:      reply.Data,
		FuelDesired: reply.FuelDesired,
		FuelApplied:   reply.FuelApplied,
	}
}

//
//
//
//
func SerializeTransferOutcomes(r []*InvokeTransferOutcome) ([][]byte, error) {
	s := make([][]byte, len(r))
	for i, e := range r {
		d := CertainInvokeTransferOutcome(e)
		b, err := d.Serialize()
		if err != nil {
			return nil, err
		}
		s[i] = b
	}
	return s, nil
}

//
//
