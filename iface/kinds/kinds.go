package kinds

import (
	"bytes"
	"encoding/json"

	"github.com/cosmos/gogoproto/jsonpb"
)

const (
	CipherKindOKAY uint32 = 0

	//
	//
	CipherKindReissue uint32 = 32_000
)

//
func (r ReplyInspectTransfer) EqualsOKAY() bool {
	return r.Cipher == CipherKindOKAY
}

//
func (r ReplyInspectTransfer) EqualsFault() bool {
	return r.Cipher != CipherKindOKAY
}

//
func (r InvokeTransferOutcome) EqualsOKAY() bool {
	return r.Cipher == CipherKindOKAY
}

//
func (r InvokeTransferOutcome) EqualsFault() bool {
	return r.Cipher != CipherKindOKAY
}

//
func (r ReplyInquire) EqualsOKAY() bool {
	return r.Cipher == CipherKindOKAY
}

//
func (r ReplyInquire) EqualsFault() bool {
	return r.Cipher != CipherKindOKAY
}

//
func (r ReplyHandleNomination) EqualsApproved() bool {
	return r.Condition == Responseexecuteitem_EMBRACE
}

//
func (r ReplyHandleNomination) EqualsConditionUnfamiliar() bool {
	return r.Condition == Responseexecuteitem_UNFAMILIAR
}

func (r ReplyValidateBallotAddition) EqualsApproved() bool {
	return r.Condition == Responsecertifyballotaddition_EMBRACE
}

//
func (r ReplyValidateBallotAddition) EqualsConditionUnfamiliar() bool {
	return r.Condition == Responsecertifyballotaddition_UNFAMILIAR
}

//
//

var (
	jsonpbSerializer = jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
	}
	jsonpbDeserializer = jsonpb.Unmarshaler{}
)

func (r *ReplyInspectTransfer) SerializeJSN() ([]byte, error) {
	s, err := jsonpbSerializer.MarshalToString(r)
	return []byte(s), err
}

func (r *ReplyInspectTransfer) DecodeJSN(b []byte) error {
	fetcher := bytes.NewBuffer(b)
	return jsonpbDeserializer.Unmarshal(fetcher, r)
}

func (r *InvokeTransferOutcome) SerializeJSN() ([]byte, error) {
	s, err := jsonpbSerializer.MarshalToString(r)
	return []byte(s), err
}

func (r *InvokeTransferOutcome) DecodeJSN(b []byte) error {
	fetcher := bytes.NewBuffer(b)
	return jsonpbDeserializer.Unmarshal(fetcher, r)
}

func (r *ReplyInquire) SerializeJSN() ([]byte, error) {
	s, err := jsonpbSerializer.MarshalToString(r)
	return []byte(s), err
}

func (r *ReplyInquire) DecodeJSN(b []byte) error {
	fetcher := bytes.NewBuffer(b)
	return jsonpbDeserializer.Unmarshal(fetcher, r)
}

func (r *ReplyEndorse) SerializeJSN() ([]byte, error) {
	s, err := jsonpbSerializer.MarshalToString(r)
	return []byte(s), err
}

func (r *ReplyEndorse) DecodeJSN(b []byte) error {
	fetcher := bytes.NewBuffer(b)
	return jsonpbDeserializer.Unmarshal(fetcher, r)
}

func (r *IncidentProperty) SerializeJSN() ([]byte, error) {
	s, err := jsonpbSerializer.MarshalToString(r)
	return []byte(s), err
}

func (r *IncidentProperty) DecodeJSN(b []byte) error {
	fetcher := bytes.NewBuffer(b)
	return jsonpbDeserializer.Unmarshal(fetcher, r)
}

//
//

//
//
type jsnIterationMessenger interface {
	json.Marshaler
	json.Unmarshaler
}

var (
	_ jsnIterationMessenger = (*ReplyEndorse)(nil)
	_ jsnIterationMessenger = (*ReplyInquire)(nil)
	_ jsnIterationMessenger = (*InvokeTransferOutcome)(nil)
	_ jsnIterationMessenger = (*ReplyInspectTransfer)(nil)
)

var _ jsnIterationMessenger = (*IncidentProperty)(nil)

//
//
func CertainInvokeTransferOutcome(reply *InvokeTransferOutcome) *InvokeTransferOutcome {
	return &InvokeTransferOutcome{
		Cipher:      reply.Cipher,
		Data:      reply.Data,
		FuelDesired: reply.FuelDesired,
		FuelUtilized:   reply.FuelUtilized,
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
