//
//

package agreement

import (
	fmt "fmt"
	kinds "github.com/valkyrieworks/schema/consensuscore/kinds"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

//
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

//
//
//
//
const _ = proto.GoGoProtoPackageIsVersion3 //

//
type MessageDetails struct {
	Msg    Signal `protobuf:"octets,1,opt,name=msg,proto3" json:"msg"`
	NodeUID string  `protobuf:"octets,2,opt,name=peer_id,json=peerId,proto3" json:"node_uid,omitempty"`
}

func (m *MessageDetails) Restore()         { *m = MessageDetails{} }
func (m *MessageDetails) String() string { return proto.CompactTextString(m) }
func (*MessageDetails) SchemaSignal()    {}
func (*MessageDetails) Definition() ([]byte, []int) {
	return filedefinition_ed0b60c2d348ab09, []int{0}
}
func (m *MessageDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *MessageDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Signaldetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Signaldetails.Merge(m, src)
}
func (m *MessageDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *MessageDetails) XXX_Omitunclear() {
	xxx_messagedata_Signaldetails.DiscardUnknown(m)
}

var xxx_messagedata_Signaldetails proto.InternalMessageInfo

func (m *MessageDetails) FetchMessage() Signal {
	if m != nil {
		return m.Msg
	}
	return Signal{}
}

func (m *MessageDetails) FetchNodeUID() string {
	if m != nil {
		return m.NodeUID
	}
	return "REDACTED"
}

//
type DeadlineDetails struct {
	Period time.Duration `protobuf:"octets,1,opt,name=duration,proto3,stdduration" json:"period"`
	Level   int64         `protobuf:"variableint,2,opt,name=height,proto3" json:"level,omitempty"`
	Cycle    int32         `protobuf:"variableint,3,opt,name=round,proto3" json:"epoch,omitempty"`
	Phase     uint32        `protobuf:"variableint,4,opt,name=step,proto3" json:"phase,omitempty"`
}

func (m *DeadlineDetails) Restore()         { *m = DeadlineDetails{} }
func (m *DeadlineDetails) String() string { return proto.CompactTextString(m) }
func (*DeadlineDetails) SchemaSignal()    {}
func (*DeadlineDetails) Definition() ([]byte, []int) {
	return filedefinition_ed0b60c2d348ab09, []int{1}
}
func (m *DeadlineDetails) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *DeadlineDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Deadlinedetails.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeadlineDetails) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Deadlinedetails.Merge(m, src)
}
func (m *DeadlineDetails) XXX_Volume() int {
	return m.Volume()
}
func (m *DeadlineDetails) XXX_Omitunclear() {
	xxx_messagedata_Deadlinedetails.DiscardUnknown(m)
}

var xxx_messagedata_Deadlinedetails proto.InternalMessageInfo

func (m *DeadlineDetails) FetchPeriod() time.Duration {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *DeadlineDetails) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *DeadlineDetails) FetchDuration() int32 {
	if m != nil {
		return m.Cycle
	}
	return 0
}

func (m *DeadlineDetails) FetchPhase() uint32 {
	if m != nil {
		return m.Phase
	}
	return 0
}

//
//
type TerminateLevel struct {
	Level int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"level,omitempty"`
}

func (m *TerminateLevel) Restore()         { *m = TerminateLevel{} }
func (m *TerminateLevel) String() string { return proto.CompactTextString(m) }
func (*TerminateLevel) SchemaSignal()    {}
func (*TerminateLevel) Definition() ([]byte, []int) {
	return filedefinition_ed0b60c2d348ab09, []int{2}
}
func (m *TerminateLevel) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *TerminateLevel) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Finallayer.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TerminateLevel) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Finallayer.Merge(m, src)
}
func (m *TerminateLevel) XXX_Volume() int {
	return m.Volume()
}
func (m *TerminateLevel) XXX_Omitunclear() {
	xxx_messagedata_Finallayer.DiscardUnknown(m)
}

var xxx_messagedata_Finallayer proto.InternalMessageInfo

func (m *TerminateLevel) FetchLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

type JournalSignal struct {
	//
	//
	//
	//
	//
	Sum isjournalsignal_Total `protobuf_oneof:"sum"`
}

func (m *JournalSignal) Restore()         { *m = JournalSignal{} }
func (m *JournalSignal) String() string { return proto.CompactTextString(m) }
func (*JournalSignal) SchemaSignal()    {}
func (*JournalSignal) Definition() ([]byte, []int) {
	return filedefinition_ed0b60c2d348ab09, []int{3}
}
func (m *JournalSignal) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *JournalSignal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Journalsignal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JournalSignal) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Journalsignal.Merge(m, src)
}
func (m *JournalSignal) XXX_Volume() int {
	return m.Volume()
}
func (m *JournalSignal) XXX_Omitunclear() {
	xxx_messagedata_Journalsignal.DiscardUnknown(m)
}

var xxx_messagedata_Journalsignal proto.InternalMessageInfo

type isjournalsignal_Total interface {
	isjournalsignal_Total()
	SerializeTo([]byte) (int, error)
	Volume() int
}

type Journalsignal_Signaldatadurationstate struct {
	EventDataDurationStatus *kinds.EventDataDurationStatus `protobuf:"octets,1,opt,name=event_data_round_state,json=eventDataRoundState,proto3,oneof" json:"event_data_epoch_status,omitempty"`
}
type Journalsignal_Signaldetails struct {
	MessageDetails *MessageDetails `protobuf:"octets,2,opt,name=msg_info,json=msgInfo,proto3,oneof" json:"message_details,omitempty"`
}
type Journalsignal_Deadlinedetails struct {
	DeadlineDetails *DeadlineDetails `protobuf:"octets,3,opt,name=timeout_info,json=timeoutInfo,proto3,oneof" json:"deadline_details,omitempty"`
}
type Journalsignal_Finallayer struct {
	TerminateLevel *TerminateLevel `protobuf:"octets,4,opt,name=end_height,json=endHeight,proto3,oneof" json:"terminate_level,omitempty"`
}

func (*Journalsignal_Signaldatadurationstate) isjournalsignal_Total() {}
func (*Journalsignal_Signaldetails) isjournalsignal_Total()             {}
func (*Journalsignal_Deadlinedetails) isjournalsignal_Total()         {}
func (*Journalsignal_Finallayer) isjournalsignal_Total()           {}

func (m *JournalSignal) FetchTotal() isjournalsignal_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *JournalSignal) FetchEventDataEpochStatus() *kinds.EventDataDurationStatus {
	if x, ok := m.FetchTotal().(*Journalsignal_Signaldatadurationstate); ok {
		return x.EventDataDurationStatus
	}
	return nil
}

func (m *JournalSignal) FetchMessageDetails() *MessageDetails {
	if x, ok := m.FetchTotal().(*Journalsignal_Signaldetails); ok {
		return x.MessageDetails
	}
	return nil
}

func (m *JournalSignal) FetchDeadlineDetails() *DeadlineDetails {
	if x, ok := m.FetchTotal().(*Journalsignal_Deadlinedetails); ok {
		return x.DeadlineDetails
	}
	return nil
}

func (m *JournalSignal) FetchTerminateLevel() *TerminateLevel {
	if x, ok := m.FetchTotal().(*Journalsignal_Finallayer); ok {
		return x.TerminateLevel
	}
	return nil
}

//
func (*JournalSignal) XXX_Variantcontainers() []interface{} {
	return []interface{}{
		(*Journalsignal_Signaldatadurationstate)(nil),
		(*Journalsignal_Signaldetails)(nil),
		(*Journalsignal_Deadlinedetails)(nil),
		(*Journalsignal_Finallayer)(nil),
	}
}

//
type ScheduledJournalSignal struct {
	Time time.Time   `protobuf:"octets,1,opt,name=time,proto3,stdtime" json:"moment"`
	Msg  *JournalSignal `protobuf:"octets,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *ScheduledJournalSignal) Restore()         { *m = ScheduledJournalSignal{} }
func (m *ScheduledJournalSignal) String() string { return proto.CompactTextString(m) }
func (*ScheduledJournalSignal) SchemaSignal()    {}
func (*ScheduledJournalSignal) Definition() ([]byte, []int) {
	return filedefinition_ed0b60c2d348ab09, []int{4}
}
func (m *ScheduledJournalSignal) XXX_Unserialize(b []byte) error {
	return m.Unserialize(b)
}
func (m *ScheduledJournalSignal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_messagedata_Timedjournalsignal.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeToDimensionedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduledJournalSignal) XXX_Coalesce(src proto.Message) {
	xxx_messagedata_Timedjournalsignal.Merge(m, src)
}
func (m *ScheduledJournalSignal) XXX_Volume() int {
	return m.Volume()
}
func (m *ScheduledJournalSignal) XXX_Omitunclear() {
	xxx_messagedata_Timedjournalsignal.DiscardUnknown(m)
}

var xxx_messagedata_Timedjournalsignal proto.InternalMessageInfo

func (m *ScheduledJournalSignal) FetchTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *ScheduledJournalSignal) FetchMessage() *JournalSignal {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageDetails)(nil), "REDACTED")
	proto.RegisterType((*DeadlineDetails)(nil), "REDACTED")
	proto.RegisterType((*TerminateLevel)(nil), "REDACTED")
	proto.RegisterType((*JournalSignal)(nil), "REDACTED")
	proto.RegisterType((*ScheduledJournalSignal)(nil), "REDACTED")
}

func init() { proto.RegisterFile("REDACTED", filedefinition_ed0b60c2d348ab09) }

var filedefinition_ed0b60c2d348ab09 = []byte{
	//
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0xce, 0x6c, 0xff, 0x9f, 0xfe, 0x7e, 0x08, 0xb1, 0x2c, 0xb5, 0xb0, 0x69, 0xec, 0x22, 0xf4,
	0x2a, 0x81, 0x15, 0x51, 0xbc, 0x51, 0x4b, 0x57, 0x5a, 0x70, 0x41, 0xc7, 0x05, 0x41, 0x84, 0x90,
	0x36, 0xa7, 0x69, 0x60, 0x33, 0x53, 0x32, 0x13, 0xc5, 0x2b, 0x5f, 0xa1, 0x97, 0xbe, 0x89, 0xaf,
	0xb0, 0x97, 0x7b, 0xe9, 0xd5, 0x2a, 0xed, 0x8b, 0x48, 0x66, 0xd2, 0x36, 0xb8, 0xf1, 0x6e, 0xce,
	0x9c, 0xef, 0x9c, 0xef, 0x9c, 0xef, 0x9b, 0x01, 0x4b, 0x22, 0x0b, 0x30, 0x89, 0x23, 0x26, 0xdd,
	0x39, 0x67, 0x02, 0x99, 0x48, 0x85, 0xfb, 0xc5, 0xbf, 0x72, 0x56, 0x09, 0x97, 0xdc, 0xec, 0x1c,
	0xf2, 0xce, 0x3e, 0xdf, 0xeb, 0x84, 0x3c, 0xe4, 0x0a, 0xe0, 0x66, 0x27, 0x8d, 0xed, 0xd9, 0xa5,
	0xbd, 0xe4, 0xd7, 0x15, 0x8a, 0x1c, 0x71, 0x52, 0x40, 0xa8, 0x7b, 0x17, 0x3f, 0x23, 0x93, 0xbb,
	0xb4, 0x15, 0x72, 0x1e, 0x5e, 0xa1, 0xab, 0xa2, 0x59, 0xba, 0x70, 0x83, 0x34, 0xf1, 0x65, 0xc4,
	0x59, 0x9e, 0xef, 0xff, 0x9d, 0x97, 0x51, 0x8c, 0x42, 0xfa, 0xf1, 0x4a, 0x03, 0x06, 0x08, 0x8d,
	0x0b, 0x11, 0x4e, 0xd9, 0x82, 0x9b, 0x4f, 0xa0, 0x12, 0x8b, 0xb0, 0x4b, 0x6c, 0x32, 0x6c, 0x9f,
	0x9d, 0x38, 0x65, 0x6b, 0x38, 0x17, 0x28, 0x84, 0x1f, 0xe2, 0xa8, 0x7a, 0x7d, 0xdb, 0x37, 0x68,
	0x86, 0x37, 0x4f, 0xa1, 0xb1, 0x42, 0x4c, 0xbc, 0x28, 0xe8, 0x1e, 0xd9, 0x64, 0xd8, 0x1a, 0xc1,
	0xe6, 0xb6, 0x5f, 0x7f, 0x8b, 0x98, 0x4c, 0xc7, 0xb4, 0x9e, 0xa5, 0xa6, 0xc1, 0x60, 0x4d, 0xa0,
	0x7d, 0x19, 0xc5, 0xc8, 0x53, 0xa9, 0xb8, 0x5e, 0x40, 0x73, 0x37, 0x69, 0x4e, 0xf8, 0xc0, 0xd1,
	0xa3, 0x3a, 0xbb, 0x51, 0x9d, 0x71, 0x0e, 0x18, 0x35, 0x33, 0xb2, 0xef, 0xbf, 0xfa, 0x84, 0xee,
	0x8b, 0xcc, 0x63, 0xa8, 0x2f, 0x31, 0x0a, 0x97, 0x52, 0x91, 0x56, 0x68, 0x1e, 0x99, 0x1d, 0xa8,
	0x25, 0x3c, 0x65, 0x41, 0xb7, 0x62, 0x93, 0x61, 0x8d, 0xea, 0xc0, 0x34, 0xa1, 0x2a, 0x24, 0xae,
	0xba, 0x55, 0x9b, 0x0c, 0xff, 0xa7, 0xea, 0x3c, 0x38, 0x85, 0xd6, 0x39, 0x0b, 0x26, 0xba, 0xec,
	0xd0, 0x8e, 0x14, 0xdb, 0x0d, 0x7e, 0x1c, 0x01, 0x7c, 0x78, 0xf5, 0x26, 0x5f, 0xdb, 0xfc, 0x04,
	0xc7, 0x4a, 0x7e, 0x2f, 0xf0, 0xa5, 0xef, 0xa9, 0xde, 0x9e, 0x90, 0xbe, 0xc4, 0x7c, 0x89, 0x47,
	0x45, 0xd5, 0xb4, 0x8d, 0xe7, 0x19, 0x7e, 0xec, 0x4b, 0x9f, 0x66, 0xe8, 0xf7, 0x19, 0x78, 0x62,
	0xd0, 0xfb, 0x78, 0xf7, 0xda, 0x7c, 0x0e, 0xcd, 0x58, 0x84, 0x5e, 0xc4, 0x16, 0x5c, 0x6d, 0xf5,
	0x6f, 0x17, 0xb4, 0x63, 0x13, 0x83, 0x36, 0xe2, 0xdc, 0xbc, 0xd7, 0xf0, 0x9f, 0xd4, 0xfa, 0xea,
	0xfa, 0x8a, 0xaa, 0x7f, 0x58, 0x5e, 0x5f, 0x70, 0x62, 0x62, 0xd0, 0xb6, 0x2c, 0x18, 0xf3, 0x12,
	0x00, 0x59, 0xe0, 0xe5, 0x62, 0x54, 0x55, 0x97, 0x7e, 0x79, 0x97, 0xbd, 0x7a, 0x13, 0x83, 0xb6,
	0x70, 0x17, 0x8c, 0x6a, 0x50, 0x11, 0x69, 0x3c, 0xf8, 0x06, 0xf7, 0x32, 0x9a, 0xa0, 0xa0, 0xde,
	0x33, 0xa8, 0x66, 0x54, 0xb9, 0x56, 0xbd, 0x3b, 0x86, 0x5f, 0xee, 0xde, 0xa6, 0x76, 0x7c, 0x9d,
	0x39, 0xae, 0x2a, 0xcc, 0x33, 0xfd, 0x34, 0xb5, 0x28, 0x76, 0xf9, 0x38, 0x07, 0x22, 0xf5, 0x2e,
	0x47, 0xef, 0xae, 0x37, 0x16, 0xb9, 0xd9, 0x58, 0xe4, 0xf7, 0xc6, 0x22, 0xeb, 0xad, 0x65, 0xdc,
	0x6c, 0x2d, 0xe3, 0xe7, 0xd6, 0x32, 0x3e, 0x3e, 0x0d, 0x23, 0xb9, 0x4c, 0x67, 0xce, 0x9c, 0xc7,
	0xee, 0x9c, 0xc7, 0x28, 0x67, 0x0b, 0x79, 0x38, 0xe8, 0x4f, 0x5a, 0xf6, 0x31, 0x67, 0x75, 0x95,
	0x7b, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x81, 0x69, 0x90, 0x03, 0x04, 0x00, 0x00,
}

func (m *MessageDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *MessageDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeUID) > 0 {
		i -= len(m.NodeUID)
		copy(dAtA[i:], m.NodeUID)
		i = encodeVariableintJournal(dAtA, i, uint64(len(m.NodeUID)))
		i--
		dAtA[i] = 0x12
	}
	{
		volume, err := m.Msg.SerializeToDimensionedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= volume
		i = encodeVariableintJournal(dAtA, i, uint64(volume))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DeadlineDetails) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeadlineDetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *DeadlineDetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Phase != 0 {
		i = encodeVariableintJournal(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x20
	}
	if m.Cycle != 0 {
		i = encodeVariableintJournal(dAtA, i, uint64(m.Cycle))
		i--
		dAtA[i] = 0x18
	}
	if m.Level != 0 {
		i = encodeVariableintJournal(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Period, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Period):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVariableintJournal(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TerminateLevel) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TerminateLevel) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *TerminateLevel) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Level != 0 {
		i = encodeVariableintJournal(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *JournalSignal) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JournalSignal) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *JournalSignal) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			volume := m.Sum.Volume()
			i -= volume
			if _, err := m.Sum.SerializeTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Journalsignal_Signaldatadurationstate) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Journalsignal_Signaldatadurationstate) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EventDataDurationStatus != nil {
		{
			volume, err := m.EventDataDurationStatus.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintJournal(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Journalsignal_Signaldetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Journalsignal_Signaldetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MessageDetails != nil {
		{
			volume, err := m.MessageDetails.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintJournal(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Journalsignal_Deadlinedetails) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Journalsignal_Deadlinedetails) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeadlineDetails != nil {
		{
			volume, err := m.DeadlineDetails.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintJournal(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Journalsignal_Finallayer) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *Journalsignal_Finallayer) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TerminateLevel != nil {
		{
			volume, err := m.TerminateLevel.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintJournal(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *ScheduledJournalSignal) Serialize() (dAtA []byte, err error) {
	volume := m.Volume()
	dAtA = make([]byte, volume)
	n, err := m.SerializeToDimensionedBuffer(dAtA[:volume])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledJournalSignal) SerializeTo(dAtA []byte) (int, error) {
	volume := m.Volume()
	return m.SerializeToDimensionedBuffer(dAtA[:volume])
}

func (m *ScheduledJournalSignal) SerializeToDimensionedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			volume, err := m.Msg.SerializeToDimensionedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= volume
			i = encodeVariableintJournal(dAtA, i, uint64(volume))
		}
		i--
		dAtA[i] = 0x12
	}
	n8, err8 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if err8 != nil {
		return 0, err8
	}
	i -= n8
	i = encodeVariableintJournal(dAtA, i, uint64(n8))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVariableintJournal(dAtA []byte, displacement int, v uint64) int {
	displacement -= sovJournal(v)
	root := displacement
	for v >= 1<<7 {
		dAtA[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	dAtA[displacement] = uint8(v)
	return root
}
func (m *MessageDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Msg.Volume()
	n += 1 + l + sovJournal(uint64(l))
	l = len(m.NodeUID)
	if l > 0 {
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}

func (m *DeadlineDetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Period)
	n += 1 + l + sovJournal(uint64(l))
	if m.Level != 0 {
		n += 1 + sovJournal(uint64(m.Level))
	}
	if m.Cycle != 0 {
		n += 1 + sovJournal(uint64(m.Cycle))
	}
	if m.Phase != 0 {
		n += 1 + sovJournal(uint64(m.Phase))
	}
	return n
}

func (m *TerminateLevel) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Level != 0 {
		n += 1 + sovJournal(uint64(m.Level))
	}
	return n
}

func (m *JournalSignal) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Volume()
	}
	return n
}

func (m *Journalsignal_Signaldatadurationstate) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventDataDurationStatus != nil {
		l = m.EventDataDurationStatus.Volume()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *Journalsignal_Signaldetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MessageDetails != nil {
		l = m.MessageDetails.Volume()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *Journalsignal_Deadlinedetails) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeadlineDetails != nil {
		l = m.DeadlineDetails.Volume()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *Journalsignal_Finallayer) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TerminateLevel != nil {
		l = m.TerminateLevel.Volume()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *ScheduledJournalSignal) Volume() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovJournal(uint64(l))
	if m.Msg != nil {
		l = m.Msg.Volume()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}

func sovJournal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJournal(x uint64) (n int) {
	return sovJournal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MessageDetails) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadJournal
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Msg.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var stringSize uint64
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				stringSize |= uint64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			integerStringSize := int(stringSize)
			if integerStringSize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + integerStringSize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeUID = string(dAtA[idxNdEx:submitOrdinal])
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitJournal(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentJournal
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeadlineDetails) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadJournal
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Period, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Cycle = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Cycle |= int32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Phase = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Phase |= uint32(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitJournal(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentJournal
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TerminateLevel) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadJournal
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Level = 0
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				m.Level |= int64(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdEx = preOrdinal
			skippy, err := omitJournal(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentJournal
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JournalSignal) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadJournal
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &kinds.EventDataDurationStatus{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Journalsignal_Signaldatadurationstate{v}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &MessageDetails{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Journalsignal_Signaldetails{v}
			idxNdEx = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeadlineDetails{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Journalsignal_Deadlinedetails{v}
			idxNdEx = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &TerminateLevel{}
			if err := v.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Journalsignal_Finallayer{v}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitJournal(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentJournal
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduledJournalSignal) Unserialize(dAtA []byte) error {
	l := len(dAtA)
	idxNdEx := 0
	for idxNdEx < l {
		preOrdinal := idxNdEx
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return ErrIntegerOverloadJournal
			}
			if idxNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= uint64(b&0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		fieldCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if fieldCount <= 0 {
			return fmt.Errorf("REDACTED", fieldCount, cable)
		}
		switch fieldCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var messagesize int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				messagesize |= int(b&0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if messagesize < 0 {
				return ErrCorruptExtentJournal
			}
			submitOrdinal := idxNdEx + messagesize
			if submitOrdinal < 0 {
				return ErrCorruptExtentJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &JournalSignal{}
			}
			if err := m.Msg.Unserialize(dAtA[idxNdEx:submitOrdinal]); err != nil {
				return err
			}
			idxNdEx = submitOrdinal
		default:
			idxNdEx = preOrdinal
			skippy, err := omitJournal(dAtA[idxNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (idxNdEx+skippy) < 0 {
				return ErrCorruptExtentJournal
			}
			if (idxNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdEx += skippy
		}
	}

	if idxNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func omitJournal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	idxNdEx := 0
	intensity := 0
	for idxNdEx < l {
		var cable uint64
		for displace := uint(0); ; displace += 7 {
			if displace >= 64 {
				return 0, ErrIntegerOverloadJournal
			}
			if idxNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[idxNdEx]
			idxNdEx++
			cable |= (uint64(b) & 0x7F) << displace
			if b < 0x80 {
				break
			}
		}
		cableKind := int(cable & 0x7)
		switch cableKind {
		case 0:
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return 0, ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				idxNdEx++
				if dAtA[idxNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			idxNdEx += 8
		case 2:
			var extent int
			for displace := uint(0); ; displace += 7 {
				if displace >= 64 {
					return 0, ErrIntegerOverloadJournal
				}
				if idxNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[idxNdEx]
				idxNdEx++
				extent |= (int(b) & 0x7F) << displace
				if b < 0x80 {
					break
				}
			}
			if extent < 0 {
				return 0, ErrCorruptExtentJournal
			}
			idxNdEx += extent
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, ErrUnforeseenTerminateOfClusterJournal
			}
			intensity--
		case 5:
			idxNdEx += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdEx < 0 {
			return 0, ErrCorruptExtentJournal
		}
		if intensity == 0 {
			return idxNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrCorruptExtentJournal        = fmt.Errorf("REDACTED")
	ErrIntegerOverloadJournal          = fmt.Errorf("REDACTED")
	ErrUnforeseenTerminateOfClusterJournal = fmt.Errorf("REDACTED")
)
