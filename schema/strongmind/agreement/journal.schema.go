//
//

package agreement

import (
	fmt "fmt"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
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
type SignalDetails struct {
	Msg    Signal `protobuf:"octets,1,opt,name=msg,proto3" json:"msg"`
	NodeUUID string  `protobuf:"octets,2,opt,name=peer_id,json=peerId,proto3" json:"node_uuid,omitempty"`
}

func (m *SignalDetails) Restore()         { *m = SignalDetails{} }
func (m *SignalDetails) Text() string { return proto.CompactTextString(m) }
func (*SignalDetails) SchemaArtifact()    {}
func (*SignalDetails) Definition() ([]byte, []int) {
	return filedescriptor_ed0b60c2d348ab09, []int{0}
}
func (m *SignalDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *SignalDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Signalinfo.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignalDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Signalinfo.Merge(m, src)
}
func (m *SignalDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *SignalDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Signalinfo.DiscardUnknown(m)
}

var xxx_signaldetails_Signalinfo proto.InternalMessageInfo

func (m *SignalDetails) ObtainSignal() Signal {
	if m != nil {
		return m.Msg
	}
	return Signal{}
}

func (m *SignalDetails) ObtainNodeUUID() string {
	if m != nil {
		return m.NodeUUID
	}
	return "REDACTED"
}

//
type DeadlineDetails struct {
	Interval time.Duration `protobuf:"octets,1,opt,name=duration,proto3,stdduration" json:"interval"`
	Altitude   int64         `protobuf:"variableint,2,opt,name=height,proto3" json:"altitude,omitempty"`
	Iteration    int32         `protobuf:"variableint,3,opt,name=round,proto3" json:"iteration,omitempty"`
	Phase     uint32        `protobuf:"variableint,4,opt,name=step,proto3" json:"phase,omitempty"`
}

func (m *DeadlineDetails) Restore()         { *m = DeadlineDetails{} }
func (m *DeadlineDetails) Text() string { return proto.CompactTextString(m) }
func (*DeadlineDetails) SchemaArtifact()    {}
func (*DeadlineDetails) Definition() ([]byte, []int) {
	return filedescriptor_ed0b60c2d348ab09, []int{1}
}
func (m *DeadlineDetails) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *DeadlineDetails) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Alarminfo.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeadlineDetails) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Alarminfo.Merge(m, src)
}
func (m *DeadlineDetails) XXX_Extent() int {
	return m.Extent()
}
func (m *DeadlineDetails) XXX_Dropunfamiliar() {
	xxx_signaldetails_Alarminfo.DiscardUnknown(m)
}

var xxx_signaldetails_Alarminfo proto.InternalMessageInfo

func (m *DeadlineDetails) ObtainInterval() time.Duration {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *DeadlineDetails) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *DeadlineDetails) ObtainIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func (m *DeadlineDetails) ObtainPhase() uint32 {
	if m != nil {
		return m.Phase
	}
	return 0
}

//
//
type TerminateAltitude struct {
	Altitude int64 `protobuf:"variableint,1,opt,name=height,proto3" json:"altitude,omitempty"`
}

func (m *TerminateAltitude) Restore()         { *m = TerminateAltitude{} }
func (m *TerminateAltitude) Text() string { return proto.CompactTextString(m) }
func (*TerminateAltitude) SchemaArtifact()    {}
func (*TerminateAltitude) Definition() ([]byte, []int) {
	return filedescriptor_ed0b60c2d348ab09, []int{2}
}
func (m *TerminateAltitude) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *TerminateAltitude) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Finalheight.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TerminateAltitude) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Finalheight.Merge(m, src)
}
func (m *TerminateAltitude) XXX_Extent() int {
	return m.Extent()
}
func (m *TerminateAltitude) XXX_Dropunfamiliar() {
	xxx_signaldetails_Finalheight.DiscardUnknown(m)
}

var xxx_signaldetails_Finalheight proto.InternalMessageInfo

func (m *TerminateAltitude) ObtainAltitude() int64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

type JournalSignal struct {
	//
	//
	//
	//
	//
	Sum isjournalrecord_Total `protobuf_oneof:"sum"`
}

func (m *JournalSignal) Restore()         { *m = JournalSignal{} }
func (m *JournalSignal) Text() string { return proto.CompactTextString(m) }
func (*JournalSignal) SchemaArtifact()    {}
func (*JournalSignal) Definition() ([]byte, []int) {
	return filedescriptor_ed0b60c2d348ab09, []int{3}
}
func (m *JournalSignal) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *JournalSignal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Walrecord.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JournalSignal) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Walrecord.Merge(m, src)
}
func (m *JournalSignal) XXX_Extent() int {
	return m.Extent()
}
func (m *JournalSignal) XXX_Dropunfamiliar() {
	xxx_signaldetails_Walrecord.DiscardUnknown(m)
}

var xxx_signaldetails_Walrecord proto.InternalMessageInfo

type isjournalrecord_Total interface {
	isjournalrecord_Total()
	SerializeToward([]byte) (int, error)
	Extent() int
}

type Walrecord_Incidentiterationstate struct {
	IncidentDataIterationStatus *kinds.IncidentDataIterationStatus `protobuf:"octets,1,opt,name=event_data_round_state,json=eventDataRoundState,proto3,oneof" json:"incident_data_iteration_status,omitempty"`
}
type Walrecord_Signalinfo struct {
	SignalDetails *SignalDetails `protobuf:"octets,2,opt,name=msg_info,json=msgInfo,proto3,oneof" json:"signal_details,omitempty"`
}
type Walrecord_Alarminfo struct {
	DeadlineDetails *DeadlineDetails `protobuf:"octets,3,opt,name=timeout_info,json=timeoutInfo,proto3,oneof" json:"deadline_details,omitempty"`
}
type Walrecord_Finalheight struct {
	TerminateAltitude *TerminateAltitude `protobuf:"octets,4,opt,name=end_height,json=endHeight,proto3,oneof" json:"terminate_altitude,omitempty"`
}

func (*Walrecord_Incidentiterationstate) isjournalrecord_Total() {}
func (*Walrecord_Signalinfo) isjournalrecord_Total()             {}
func (*Walrecord_Alarminfo) isjournalrecord_Total()         {}
func (*Walrecord_Finalheight) isjournalrecord_Total()           {}

func (m *JournalSignal) ObtainTotal() isjournalrecord_Total {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *JournalSignal) ObtainIncidentDataIterationStatus() *kinds.IncidentDataIterationStatus {
	if x, ok := m.ObtainTotal().(*Walrecord_Incidentiterationstate); ok {
		return x.IncidentDataIterationStatus
	}
	return nil
}

func (m *JournalSignal) ObtainSignalDetails() *SignalDetails {
	if x, ok := m.ObtainTotal().(*Walrecord_Signalinfo); ok {
		return x.SignalDetails
	}
	return nil
}

func (m *JournalSignal) ObtainDeadlineDetails() *DeadlineDetails {
	if x, ok := m.ObtainTotal().(*Walrecord_Alarminfo); ok {
		return x.DeadlineDetails
	}
	return nil
}

func (m *JournalSignal) ObtainTerminateAltitude() *TerminateAltitude {
	if x, ok := m.ObtainTotal().(*Walrecord_Finalheight); ok {
		return x.TerminateAltitude
	}
	return nil
}

//
func (*JournalSignal) XXX_Oneofwrappers() []interface{} {
	return []interface{}{
		(*Walrecord_Incidentiterationstate)(nil),
		(*Walrecord_Signalinfo)(nil),
		(*Walrecord_Alarminfo)(nil),
		(*Walrecord_Finalheight)(nil),
	}
}

//
type ScheduledJournalSignal struct {
	Moment time.Time   `protobuf:"octets,1,opt,name=time,proto3,stdtime" json:"moment"`
	Msg  *JournalSignal `protobuf:"octets,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *ScheduledJournalSignal) Restore()         { *m = ScheduledJournalSignal{} }
func (m *ScheduledJournalSignal) Text() string { return proto.CompactTextString(m) }
func (*ScheduledJournalSignal) SchemaArtifact()    {}
func (*ScheduledJournalSignal) Definition() ([]byte, []int) {
	return filedescriptor_ed0b60c2d348ab09, []int{4}
}
func (m *ScheduledJournalSignal) XXX_Decode(b []byte) error {
	return m.Decode(b)
}
func (m *ScheduledJournalSignal) XXX_Serialize(b []byte, certain bool) ([]byte, error) {
	if certain {
		return xxx_signaldetails_Timedjournalrecord.Marshal(b, m, certain)
	} else {
		b = b[:cap(b)]
		n, err := m.SerializeTowardDimensionedReserve(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduledJournalSignal) XXX_Consolidate(src proto.Message) {
	xxx_signaldetails_Timedjournalrecord.Merge(m, src)
}
func (m *ScheduledJournalSignal) XXX_Extent() int {
	return m.Extent()
}
func (m *ScheduledJournalSignal) XXX_Dropunfamiliar() {
	xxx_signaldetails_Timedjournalrecord.DiscardUnknown(m)
}

var xxx_signaldetails_Timedjournalrecord proto.InternalMessageInfo

func (m *ScheduledJournalSignal) ObtainMoment() time.Time {
	if m != nil {
		return m.Moment
	}
	return time.Time{}
}

func (m *ScheduledJournalSignal) ObtainSignal() *JournalSignal {
	if m != nil {
		return m.Msg
	}
	return nil
}

func initialize() {
	proto.RegisterType((*SignalDetails)(nil), "REDACTED")
	proto.RegisterType((*DeadlineDetails)(nil), "REDACTED")
	proto.RegisterType((*TerminateAltitude)(nil), "REDACTED")
	proto.RegisterType((*JournalSignal)(nil), "REDACTED")
	proto.RegisterType((*ScheduledJournalSignal)(nil), "REDACTED")
}

func initialize() { proto.RegisterFile("REDACTED", filedescriptor_ed0b60c2d348ab09) }

var filedescriptor_ed0b60c2d348ab09 = []byte{
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

func (m *SignalDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *SignalDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *SignalDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if len(m.NodeUUID) > 0 {
		i -= len(m.NodeUUID)
		copy(deltaLocatedAN[i:], m.NodeUUID)
		i = encodeVariableintJournal(deltaLocatedAN, i, uint64(len(m.NodeUUID)))
		i--
		deltaLocatedAN[i] = 0x12
	}
	{
		extent, err := m.Msg.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
		if err != nil {
			return 0, err
		}
		i -= extent
		i = encodeVariableintJournal(deltaLocatedAN, i, uint64(extent))
	}
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func (m *DeadlineDetails) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *DeadlineDetails) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *DeadlineDetails) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Phase != 0 {
		i = encodeVariableintJournal(deltaLocatedAN, i, uint64(m.Phase))
		i--
		deltaLocatedAN[i] = 0x20
	}
	if m.Iteration != 0 {
		i = encodeVariableintJournal(deltaLocatedAN, i, uint64(m.Iteration))
		i--
		deltaLocatedAN[i] = 0x18
	}
	if m.Altitude != 0 {
		i = encodeVariableintJournal(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x10
	}
	n2, fault2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Interval, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Interval):])
	if fault2 != nil {
		return 0, fault2
	}
	i -= n2
	i = encodeVariableintJournal(deltaLocatedAN, i, uint64(n2))
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func (m *TerminateAltitude) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *TerminateAltitude) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *TerminateAltitude) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Altitude != 0 {
		i = encodeVariableintJournal(deltaLocatedAN, i, uint64(m.Altitude))
		i--
		deltaLocatedAN[i] = 0x8
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *JournalSignal) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *JournalSignal) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *JournalSignal) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			extent := m.Sum.Extent()
			i -= extent
			if _, err := m.Sum.SerializeToward(deltaLocatedAN[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(deltaLocatedAN) - i, nil
}

func (m *Walrecord_Incidentiterationstate) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Walrecord_Incidentiterationstate) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.IncidentDataIterationStatus != nil {
		{
			extent, err := m.IncidentDataIterationStatus.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintJournal(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0xa
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Walrecord_Signalinfo) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Walrecord_Signalinfo) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.SignalDetails != nil {
		{
			extent, err := m.SignalDetails.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintJournal(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Walrecord_Alarminfo) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Walrecord_Alarminfo) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.DeadlineDetails != nil {
		{
			extent, err := m.DeadlineDetails.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintJournal(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x1a
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *Walrecord_Finalheight) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *Walrecord_Finalheight) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	if m.TerminateAltitude != nil {
		{
			extent, err := m.TerminateAltitude.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintJournal(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x22
	}
	return len(deltaLocatedAN) - i, nil
}
func (m *ScheduledJournalSignal) Serialize() (deltaLocatedAN []byte, err error) {
	extent := m.Extent()
	deltaLocatedAN = make([]byte, extent)
	n, err := m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
	if err != nil {
		return nil, err
	}
	return deltaLocatedAN[:n], nil
}

func (m *ScheduledJournalSignal) SerializeToward(deltaLocatedAN []byte) (int, error) {
	extent := m.Extent()
	return m.SerializeTowardDimensionedReserve(deltaLocatedAN[:extent])
}

func (m *ScheduledJournalSignal) SerializeTowardDimensionedReserve(deltaLocatedAN []byte) (int, error) {
	i := len(deltaLocatedAN)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			extent, err := m.Msg.SerializeTowardDimensionedReserve(deltaLocatedAN[:i])
			if err != nil {
				return 0, err
			}
			i -= extent
			i = encodeVariableintJournal(deltaLocatedAN, i, uint64(extent))
		}
		i--
		deltaLocatedAN[i] = 0x12
	}
	n8, fault8 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Moment, deltaLocatedAN[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment):])
	if fault8 != nil {
		return 0, fault8
	}
	i -= n8
	i = encodeVariableintJournal(deltaLocatedAN, i, uint64(n8))
	i--
	deltaLocatedAN[i] = 0xa
	return len(deltaLocatedAN) - i, nil
}

func encodeVariableintJournal(deltaLocatedAN []byte, displacement int, v uint64) int {
	displacement -= sovJournal(v)
	foundation := displacement
	for v >= 1<<7 {
		deltaLocatedAN[displacement] = uint8(v&0x7f | 0x80)
		v >>= 7
		displacement++
	}
	deltaLocatedAN[displacement] = uint8(v)
	return foundation
}
func (m *SignalDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Msg.Extent()
	n += 1 + l + sovJournal(uint64(l))
	l = len(m.NodeUUID)
	if l > 0 {
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}

func (m *DeadlineDetails) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Interval)
	n += 1 + l + sovJournal(uint64(l))
	if m.Altitude != 0 {
		n += 1 + sovJournal(uint64(m.Altitude))
	}
	if m.Iteration != 0 {
		n += 1 + sovJournal(uint64(m.Iteration))
	}
	if m.Phase != 0 {
		n += 1 + sovJournal(uint64(m.Phase))
	}
	return n
}

func (m *TerminateAltitude) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovJournal(uint64(m.Altitude))
	}
	return n
}

func (m *JournalSignal) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Extent()
	}
	return n
}

func (m *Walrecord_Incidentiterationstate) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IncidentDataIterationStatus != nil {
		l = m.IncidentDataIterationStatus.Extent()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *Walrecord_Signalinfo) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignalDetails != nil {
		l = m.SignalDetails.Extent()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *Walrecord_Alarminfo) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeadlineDetails != nil {
		l = m.DeadlineDetails.Extent()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *Walrecord_Finalheight) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TerminateAltitude != nil {
		l = m.TerminateAltitude.Extent()
		n += 1 + l + sovJournal(uint64(l))
	}
	return n
}
func (m *ScheduledJournalSignal) Extent() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Moment)
	n += 1 + l + sovJournal(uint64(l))
	if m.Msg != nil {
		l = m.Msg.Extent()
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
func (m *SignalDetails) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunJournal
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Msg.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var textSize uint64
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				textSize |= uint64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			integerTextSize := int(textSize)
			if integerTextSize < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + integerTextSize
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeUUID = string(deltaLocatedAN[idxNdExc:submitOrdinal])
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitJournal(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeadlineDetails) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunJournal
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Interval, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 3:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Iteration = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Iteration |= int32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		case 4:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Phase = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Phase |= uint32(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitJournal(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TerminateAltitude) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunJournal
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 0 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			m.Altitude = 0
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				m.Altitude |= int64(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitJournal(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JournalSignal) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunJournal
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &kinds.IncidentDataIterationStatus{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Walrecord_Incidentiterationstate{v}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignalDetails{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Walrecord_Signalinfo{v}
			idxNdExc = submitOrdinal
		case 3:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeadlineDetails{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Walrecord_Alarminfo{v}
			idxNdExc = submitOrdinal
		case 4:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			v := &TerminateAltitude{}
			if err := v.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			m.Sum = &Walrecord_Finalheight{v}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitJournal(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduledJournalSignal) Decode(deltaLocatedAN []byte) error {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	for idxNdExc < l {
		priorOrdinal := idxNdExc
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return FaultIntegerOverrunJournal
			}
			if idxNdExc >= l {
				return io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= uint64(b&0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		attributeCount := int32(cable >> 3)
		cableKind := int(cable & 0x7)
		if cableKind == 4 {
			return fmt.Errorf("REDACTED")
		}
		if attributeCount <= 0 {
			return fmt.Errorf("REDACTED", attributeCount, cable)
		}
		switch attributeCount {
		case 1:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Moment, deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		case 2:
			if cableKind != 2 {
				return fmt.Errorf("REDACTED", cableKind)
			}
			var signallength int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				signallength |= int(b&0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if signallength < 0 {
				return FaultUnfitMagnitudeJournal
			}
			submitOrdinal := idxNdExc + signallength
			if submitOrdinal < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if submitOrdinal > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &JournalSignal{}
			}
			if err := m.Msg.Decode(deltaLocatedAN[idxNdExc:submitOrdinal]); err != nil {
				return err
			}
			idxNdExc = submitOrdinal
		default:
			idxNdExc = priorOrdinal
			omitted, err := omitJournal(deltaLocatedAN[idxNdExc:])
			if err != nil {
				return err
			}
			if (omitted < 0) || (idxNdExc+omitted) < 0 {
				return FaultUnfitMagnitudeJournal
			}
			if (idxNdExc + omitted) > l {
				return io.ErrUnexpectedEOF
			}
			idxNdExc += omitted
		}
	}

	if idxNdExc > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func omitJournal(deltaLocatedAN []byte) (n int, err error) {
	l := len(deltaLocatedAN)
	idxNdExc := 0
	intensity := 0
	for idxNdExc < l {
		var cable uint64
		for relocate := uint(0); ; relocate += 7 {
			if relocate >= 64 {
				return 0, FaultIntegerOverrunJournal
			}
			if idxNdExc >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := deltaLocatedAN[idxNdExc]
			idxNdExc++
			cable |= (uint64(b) & 0x7F) << relocate
			if b < 0x80 {
				break
			}
		}
		cableKind := int(cable & 0x7)
		switch cableKind {
		case 0:
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return 0, FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return 0, io.ErrUnexpectedEOF
				}
				idxNdExc++
				if deltaLocatedAN[idxNdExc-1] < 0x80 {
					break
				}
			}
		case 1:
			idxNdExc += 8
		case 2:
			var magnitude int
			for relocate := uint(0); ; relocate += 7 {
				if relocate >= 64 {
					return 0, FaultIntegerOverrunJournal
				}
				if idxNdExc >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := deltaLocatedAN[idxNdExc]
				idxNdExc++
				magnitude |= (int(b) & 0x7F) << relocate
				if b < 0x80 {
					break
				}
			}
			if magnitude < 0 {
				return 0, FaultUnfitMagnitudeJournal
			}
			idxNdExc += magnitude
		case 3:
			intensity++
		case 4:
			if intensity == 0 {
				return 0, FaultUnforeseenTerminateBelongingCollectionJournal
			}
			intensity--
		case 5:
			idxNdExc += 4
		default:
			return 0, fmt.Errorf("REDACTED", cableKind)
		}
		if idxNdExc < 0 {
			return 0, FaultUnfitMagnitudeJournal
		}
		if intensity == 0 {
			return idxNdExc, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	FaultUnfitMagnitudeJournal        = fmt.Errorf("REDACTED")
	FaultIntegerOverrunJournal          = fmt.Errorf("REDACTED")
	FaultUnforeseenTerminateBelongingCollectionJournal = fmt.Errorf("REDACTED")
)
