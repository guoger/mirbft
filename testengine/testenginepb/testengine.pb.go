// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testengine.proto

package testenginepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mirbftpb "github.com/IBM/mirbft/mirbftpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogEntryPreamble_Type int32

const (
	LogEntryPreamble_UNKNOWN  LogEntryPreamble_Type = 0
	LogEntryPreamble_SCENARIO LogEntryPreamble_Type = 1
	LogEntryPreamble_EVENT    LogEntryPreamble_Type = 2
)

var LogEntryPreamble_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCENARIO",
	2: "EVENT",
}
var LogEntryPreamble_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"SCENARIO": 1,
	"EVENT":    2,
}

func (x LogEntryPreamble_Type) String() string {
	return proto.EnumName(LogEntryPreamble_Type_name, int32(x))
}
func (LogEntryPreamble_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{6, 0}
}

type ScenarioConfig struct {
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	InitialNetworkConfig *mirbftpb.NetworkConfig `protobuf:"bytes,3,opt,name=initial_network_config,json=initialNetworkConfig,proto3" json:"initial_network_config,omitempty"`
	NodeConfigs          []*NodeConfig           `protobuf:"bytes,4,rep,name=node_configs,json=nodeConfigs,proto3" json:"node_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ScenarioConfig) Reset()         { *m = ScenarioConfig{} }
func (m *ScenarioConfig) String() string { return proto.CompactTextString(m) }
func (*ScenarioConfig) ProtoMessage()    {}
func (*ScenarioConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{0}
}
func (m *ScenarioConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScenarioConfig.Unmarshal(m, b)
}
func (m *ScenarioConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScenarioConfig.Marshal(b, m, deterministic)
}
func (dst *ScenarioConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScenarioConfig.Merge(dst, src)
}
func (m *ScenarioConfig) XXX_Size() int {
	return xxx_messageInfo_ScenarioConfig.Size(m)
}
func (m *ScenarioConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ScenarioConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ScenarioConfig proto.InternalMessageInfo

func (m *ScenarioConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScenarioConfig) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ScenarioConfig) GetInitialNetworkConfig() *mirbftpb.NetworkConfig {
	if m != nil {
		return m.InitialNetworkConfig
	}
	return nil
}

func (m *ScenarioConfig) GetNodeConfigs() []*NodeConfig {
	if m != nil {
		return m.NodeConfigs
	}
	return nil
}

type NodeConfig struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HeartbeatTicks       int32    `protobuf:"varint,2,opt,name=heartbeat_ticks,json=heartbeatTicks,proto3" json:"heartbeat_ticks,omitempty"`
	SuspectTicks         int32    `protobuf:"varint,3,opt,name=suspect_ticks,json=suspectTicks,proto3" json:"suspect_ticks,omitempty"`
	NewEpochTimeoutTicks int32    `protobuf:"varint,4,opt,name=new_epoch_timeout_ticks,json=newEpochTimeoutTicks,proto3" json:"new_epoch_timeout_ticks,omitempty"`
	TickInterval         int32    `protobuf:"varint,5,opt,name=tick_interval,json=tickInterval,proto3" json:"tick_interval,omitempty"`
	LinkLatency          int32    `protobuf:"varint,6,opt,name=link_latency,json=linkLatency,proto3" json:"link_latency,omitempty"`
	ReadyLatency         int32    `protobuf:"varint,7,opt,name=ready_latency,json=readyLatency,proto3" json:"ready_latency,omitempty"`
	ProcessLatency       int32    `protobuf:"varint,8,opt,name=process_latency,json=processLatency,proto3" json:"process_latency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{1}
}
func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (dst *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(dst, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NodeConfig) GetHeartbeatTicks() int32 {
	if m != nil {
		return m.HeartbeatTicks
	}
	return 0
}

func (m *NodeConfig) GetSuspectTicks() int32 {
	if m != nil {
		return m.SuspectTicks
	}
	return 0
}

func (m *NodeConfig) GetNewEpochTimeoutTicks() int32 {
	if m != nil {
		return m.NewEpochTimeoutTicks
	}
	return 0
}

func (m *NodeConfig) GetTickInterval() int32 {
	if m != nil {
		return m.TickInterval
	}
	return 0
}

func (m *NodeConfig) GetLinkLatency() int32 {
	if m != nil {
		return m.LinkLatency
	}
	return 0
}

func (m *NodeConfig) GetReadyLatency() int32 {
	if m != nil {
		return m.ReadyLatency
	}
	return 0
}

func (m *NodeConfig) GetProcessLatency() int32 {
	if m != nil {
		return m.ProcessLatency
	}
	return 0
}

type Request struct {
	ClientId             []byte   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ReqNo                uint64   `protobuf:"varint,2,opt,name=req_no,json=reqNo,proto3" json:"req_no,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Digest               []byte   `protobuf:"bytes,5,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *Request) GetReqNo() uint64 {
	if m != nil {
		return m.ReqNo
	}
	return 0
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Request) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Request) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Batch struct {
	Source               uint64              `protobuf:"varint,1,opt,name=source,proto3" json:"source,omitempty"`
	Epoch                uint64              `protobuf:"varint,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	SeqNo                uint64              `protobuf:"varint,3,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	Digest               []byte              `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	Requests             []*mirbftpb.Request `protobuf:"bytes,5,rep,name=requests,proto3" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{3}
}
func (m *Batch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Batch.Unmarshal(m, b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
}
func (dst *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(dst, src)
}
func (m *Batch) XXX_Size() int {
	return xxx_messageInfo_Batch.Size(m)
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetSource() uint64 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *Batch) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Batch) GetSeqNo() uint64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

func (m *Batch) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Batch) GetRequests() []*mirbftpb.Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

type Checkpoint struct {
	SeqNo                uint64   `protobuf:"varint,1,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{4}
}
func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (dst *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(dst, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetSeqNo() uint64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

func (m *Checkpoint) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Event struct {
	Target uint64 `protobuf:"varint,1,opt,name=target,proto3" json:"target,omitempty"`
	Time   uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*Event_Apply_
	//	*Event_Receive_
	//	*Event_Process_
	//	*Event_Propose_
	//	*Event_Tick_
	Type                 isEvent_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{5}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTarget() uint64 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *Event) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_Apply_ struct {
	Apply *Event_Apply `protobuf:"bytes,3,opt,name=apply,proto3,oneof"`
}

type Event_Receive_ struct {
	Receive *Event_Receive `protobuf:"bytes,4,opt,name=receive,proto3,oneof"`
}

type Event_Process_ struct {
	Process *Event_Process `protobuf:"bytes,5,opt,name=process,proto3,oneof"`
}

type Event_Propose_ struct {
	Propose *Event_Propose `protobuf:"bytes,6,opt,name=propose,proto3,oneof"`
}

type Event_Tick_ struct {
	Tick *Event_Tick `protobuf:"bytes,7,opt,name=tick,proto3,oneof"`
}

func (*Event_Apply_) isEvent_Type() {}

func (*Event_Receive_) isEvent_Type() {}

func (*Event_Process_) isEvent_Type() {}

func (*Event_Propose_) isEvent_Type() {}

func (*Event_Tick_) isEvent_Type() {}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Event) GetApply() *Event_Apply {
	if x, ok := m.GetType().(*Event_Apply_); ok {
		return x.Apply
	}
	return nil
}

func (m *Event) GetReceive() *Event_Receive {
	if x, ok := m.GetType().(*Event_Receive_); ok {
		return x.Receive
	}
	return nil
}

func (m *Event) GetProcess() *Event_Process {
	if x, ok := m.GetType().(*Event_Process_); ok {
		return x.Process
	}
	return nil
}

func (m *Event) GetPropose() *Event_Propose {
	if x, ok := m.GetType().(*Event_Propose_); ok {
		return x.Propose
	}
	return nil
}

func (m *Event) GetTick() *Event_Tick {
	if x, ok := m.GetType().(*Event_Tick_); ok {
		return x.Tick
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Apply_)(nil),
		(*Event_Receive_)(nil),
		(*Event_Process_)(nil),
		(*Event_Propose_)(nil),
		(*Event_Tick_)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// type
	switch x := m.Type.(type) {
	case *Event_Apply_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Apply); err != nil {
			return err
		}
	case *Event_Receive_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Receive); err != nil {
			return err
		}
	case *Event_Process_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Process); err != nil {
			return err
		}
	case *Event_Propose_:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Propose); err != nil {
			return err
		}
	case *Event_Tick_:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tick); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Type has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 3: // type.apply
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Apply)
		err := b.DecodeMessage(msg)
		m.Type = &Event_Apply_{msg}
		return true, err
	case 4: // type.receive
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Receive)
		err := b.DecodeMessage(msg)
		m.Type = &Event_Receive_{msg}
		return true, err
	case 5: // type.process
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Process)
		err := b.DecodeMessage(msg)
		m.Type = &Event_Process_{msg}
		return true, err
	case 6: // type.propose
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Propose)
		err := b.DecodeMessage(msg)
		m.Type = &Event_Propose_{msg}
		return true, err
	case 7: // type.tick
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Tick)
		err := b.DecodeMessage(msg)
		m.Type = &Event_Tick_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// type
	switch x := m.Type.(type) {
	case *Event_Apply_:
		s := proto.Size(x.Apply)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Receive_:
		s := proto.Size(x.Receive)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Process_:
		s := proto.Size(x.Process)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Propose_:
		s := proto.Size(x.Propose)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Tick_:
		s := proto.Size(x.Tick)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Event_Apply struct {
	Preprocessed         []*Request    `protobuf:"bytes,1,rep,name=preprocessed,proto3" json:"preprocessed,omitempty"`
	Processed            []*Batch      `protobuf:"bytes,2,rep,name=processed,proto3" json:"processed,omitempty"`
	Checkpoints          []*Checkpoint `protobuf:"bytes,3,rep,name=checkpoints,proto3" json:"checkpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Event_Apply) Reset()         { *m = Event_Apply{} }
func (m *Event_Apply) String() string { return proto.CompactTextString(m) }
func (*Event_Apply) ProtoMessage()    {}
func (*Event_Apply) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{5, 0}
}
func (m *Event_Apply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Apply.Unmarshal(m, b)
}
func (m *Event_Apply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Apply.Marshal(b, m, deterministic)
}
func (dst *Event_Apply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Apply.Merge(dst, src)
}
func (m *Event_Apply) XXX_Size() int {
	return xxx_messageInfo_Event_Apply.Size(m)
}
func (m *Event_Apply) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Apply.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Apply proto.InternalMessageInfo

func (m *Event_Apply) GetPreprocessed() []*Request {
	if m != nil {
		return m.Preprocessed
	}
	return nil
}

func (m *Event_Apply) GetProcessed() []*Batch {
	if m != nil {
		return m.Processed
	}
	return nil
}

func (m *Event_Apply) GetCheckpoints() []*Checkpoint {
	if m != nil {
		return m.Checkpoints
	}
	return nil
}

type Event_Process struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_Process) Reset()         { *m = Event_Process{} }
func (m *Event_Process) String() string { return proto.CompactTextString(m) }
func (*Event_Process) ProtoMessage()    {}
func (*Event_Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{5, 1}
}
func (m *Event_Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Process.Unmarshal(m, b)
}
func (m *Event_Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Process.Marshal(b, m, deterministic)
}
func (dst *Event_Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Process.Merge(dst, src)
}
func (m *Event_Process) XXX_Size() int {
	return xxx_messageInfo_Event_Process.Size(m)
}
func (m *Event_Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Process proto.InternalMessageInfo

type Event_Propose struct {
	RequestData          *mirbftpb.RequestData `protobuf:"bytes,1,opt,name=request_data,json=requestData,proto3" json:"request_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Event_Propose) Reset()         { *m = Event_Propose{} }
func (m *Event_Propose) String() string { return proto.CompactTextString(m) }
func (*Event_Propose) ProtoMessage()    {}
func (*Event_Propose) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{5, 2}
}
func (m *Event_Propose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Propose.Unmarshal(m, b)
}
func (m *Event_Propose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Propose.Marshal(b, m, deterministic)
}
func (dst *Event_Propose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Propose.Merge(dst, src)
}
func (m *Event_Propose) XXX_Size() int {
	return xxx_messageInfo_Event_Propose.Size(m)
}
func (m *Event_Propose) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Propose.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Propose proto.InternalMessageInfo

func (m *Event_Propose) GetRequestData() *mirbftpb.RequestData {
	if m != nil {
		return m.RequestData
	}
	return nil
}

type Event_Receive struct {
	Source               uint64        `protobuf:"varint,1,opt,name=source,proto3" json:"source,omitempty"`
	Msg                  *mirbftpb.Msg `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Event_Receive) Reset()         { *m = Event_Receive{} }
func (m *Event_Receive) String() string { return proto.CompactTextString(m) }
func (*Event_Receive) ProtoMessage()    {}
func (*Event_Receive) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{5, 3}
}
func (m *Event_Receive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Receive.Unmarshal(m, b)
}
func (m *Event_Receive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Receive.Marshal(b, m, deterministic)
}
func (dst *Event_Receive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Receive.Merge(dst, src)
}
func (m *Event_Receive) XXX_Size() int {
	return xxx_messageInfo_Event_Receive.Size(m)
}
func (m *Event_Receive) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Receive.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Receive proto.InternalMessageInfo

func (m *Event_Receive) GetSource() uint64 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *Event_Receive) GetMsg() *mirbftpb.Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Event_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_Tick) Reset()         { *m = Event_Tick{} }
func (m *Event_Tick) String() string { return proto.CompactTextString(m) }
func (*Event_Tick) ProtoMessage()    {}
func (*Event_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{5, 4}
}
func (m *Event_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Tick.Unmarshal(m, b)
}
func (m *Event_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Tick.Marshal(b, m, deterministic)
}
func (dst *Event_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Tick.Merge(dst, src)
}
func (m *Event_Tick) XXX_Size() int {
	return xxx_messageInfo_Event_Tick.Size(m)
}
func (m *Event_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Tick proto.InternalMessageInfo

type LogEntryPreamble struct {
	Type                 LogEntryPreamble_Type `protobuf:"varint,1,opt,name=type,proto3,enum=testenginepb.LogEntryPreamble_Type" json:"type,omitempty"`
	Length               int32                 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LogEntryPreamble) Reset()         { *m = LogEntryPreamble{} }
func (m *LogEntryPreamble) String() string { return proto.CompactTextString(m) }
func (*LogEntryPreamble) ProtoMessage()    {}
func (*LogEntryPreamble) Descriptor() ([]byte, []int) {
	return fileDescriptor_testengine_56032773bb6ee9b8, []int{6}
}
func (m *LogEntryPreamble) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntryPreamble.Unmarshal(m, b)
}
func (m *LogEntryPreamble) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntryPreamble.Marshal(b, m, deterministic)
}
func (dst *LogEntryPreamble) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntryPreamble.Merge(dst, src)
}
func (m *LogEntryPreamble) XXX_Size() int {
	return xxx_messageInfo_LogEntryPreamble.Size(m)
}
func (m *LogEntryPreamble) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntryPreamble.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntryPreamble proto.InternalMessageInfo

func (m *LogEntryPreamble) GetType() LogEntryPreamble_Type {
	if m != nil {
		return m.Type
	}
	return LogEntryPreamble_UNKNOWN
}

func (m *LogEntryPreamble) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func init() {
	proto.RegisterType((*ScenarioConfig)(nil), "testenginepb.ScenarioConfig")
	proto.RegisterType((*NodeConfig)(nil), "testenginepb.NodeConfig")
	proto.RegisterType((*Request)(nil), "testenginepb.Request")
	proto.RegisterType((*Batch)(nil), "testenginepb.Batch")
	proto.RegisterType((*Checkpoint)(nil), "testenginepb.Checkpoint")
	proto.RegisterType((*Event)(nil), "testenginepb.Event")
	proto.RegisterType((*Event_Apply)(nil), "testenginepb.Event.Apply")
	proto.RegisterType((*Event_Process)(nil), "testenginepb.Event.Process")
	proto.RegisterType((*Event_Propose)(nil), "testenginepb.Event.Propose")
	proto.RegisterType((*Event_Receive)(nil), "testenginepb.Event.Receive")
	proto.RegisterType((*Event_Tick)(nil), "testenginepb.Event.Tick")
	proto.RegisterType((*LogEntryPreamble)(nil), "testenginepb.LogEntryPreamble")
	proto.RegisterEnum("testenginepb.LogEntryPreamble_Type", LogEntryPreamble_Type_name, LogEntryPreamble_Type_value)
}

func init() { proto.RegisterFile("testengine.proto", fileDescriptor_testengine_56032773bb6ee9b8) }

var fileDescriptor_testengine_56032773bb6ee9b8 = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xc1, 0x72, 0x1b, 0x45,
	0x13, 0xf6, 0x4a, 0x5a, 0xc9, 0xea, 0xdd, 0x38, 0xfa, 0xe7, 0xb7, 0x13, 0xa1, 0x50, 0x85, 0x51,
	0x0e, 0xf1, 0x01, 0x44, 0x45, 0x14, 0x05, 0x81, 0x53, 0x6c, 0x54, 0x65, 0x17, 0x89, 0x92, 0x9a,
	0x18, 0x38, 0x6e, 0xad, 0x56, 0x1d, 0x69, 0x4a, 0xab, 0xd9, 0xf5, 0xcc, 0xc8, 0x2e, 0x3d, 0x01,
	0x67, 0x8a, 0x13, 0x8f, 0xc1, 0x89, 0xf7, 0xe0, 0x89, 0xa8, 0xe9, 0x19, 0x69, 0x25, 0x12, 0xdf,
	0xa6, 0xbf, 0xfe, 0xba, 0xa7, 0xbb, 0xe7, 0xeb, 0x5d, 0xe8, 0x18, 0xd4, 0x06, 0xe5, 0x4c, 0x48,
	0x1c, 0x94, 0xaa, 0x30, 0x05, 0x8b, 0x2b, 0xa4, 0x9c, 0xf4, 0x4e, 0x96, 0x42, 0x4d, 0xde, 0x9b,
	0x72, 0xf2, 0x95, 0x3b, 0x38, 0x52, 0xff, 0x9f, 0x00, 0x8e, 0xde, 0x65, 0x28, 0x53, 0x25, 0x8a,
	0x8b, 0x42, 0xbe, 0x17, 0x33, 0xc6, 0xa0, 0x21, 0xd3, 0x25, 0x76, 0x83, 0xd3, 0xe0, 0xac, 0xcd,
	0xe9, 0xcc, 0x4e, 0x21, 0x9a, 0xa2, 0xce, 0x94, 0x28, 0x8d, 0x28, 0x64, 0xb7, 0x46, 0xae, 0x5d,
	0x88, 0xbd, 0x86, 0x47, 0x42, 0x0a, 0x23, 0xd2, 0x3c, 0x91, 0x68, 0xee, 0x0a, 0xb5, 0x48, 0x32,
	0xca, 0xd7, 0xad, 0x9f, 0x06, 0x67, 0xd1, 0xf0, 0xf1, 0x60, 0x53, 0xc0, 0x60, 0xec, 0xfc, 0xee,
	0x3a, 0x7e, 0xec, 0xc3, 0xf6, 0x50, 0xf6, 0x03, 0xc4, 0xb2, 0x98, 0xa2, 0xcf, 0xa1, 0xbb, 0x8d,
	0xd3, 0xfa, 0x59, 0x34, 0xec, 0x0e, 0x76, 0x7b, 0x1a, 0x8c, 0x8b, 0x29, 0xfa, 0x2c, 0x91, 0xdc,
	0x9e, 0x75, 0xff, 0xef, 0x1a, 0x40, 0xe5, 0x63, 0x47, 0x50, 0x13, 0x53, 0x6a, 0xa7, 0xc1, 0x6b,
	0x62, 0xca, 0x9e, 0xc1, 0xc3, 0x39, 0xa6, 0xca, 0x4c, 0x30, 0x35, 0x89, 0x11, 0xd9, 0x42, 0x53,
	0x43, 0x21, 0x3f, 0xda, 0xc2, 0xd7, 0x16, 0x65, 0x4f, 0xe1, 0x81, 0x5e, 0xe9, 0x12, 0xb3, 0x0d,
	0xad, 0x4e, 0xb4, 0xd8, 0x83, 0x8e, 0xf4, 0x0d, 0x3c, 0x96, 0x78, 0x97, 0x60, 0x59, 0x64, 0xf3,
	0xc4, 0x88, 0x25, 0x16, 0xab, 0x0d, 0xbd, 0x41, 0xf4, 0x63, 0x89, 0x77, 0x23, 0xeb, 0xbd, 0x76,
	0xce, 0x6d, 0x6e, 0x4b, 0x4a, 0x84, 0x34, 0xa8, 0x6e, 0xd3, 0xbc, 0x1b, 0xba, 0xdc, 0x16, 0xbc,
	0xf2, 0x18, 0xfb, 0x1c, 0xe2, 0x5c, 0xc8, 0x45, 0x92, 0xa7, 0x06, 0x65, 0xb6, 0xee, 0x36, 0x89,
	0x13, 0x59, 0xec, 0x95, 0x83, 0x6c, 0x1e, 0x85, 0xe9, 0x74, 0xbd, 0xe5, 0xb4, 0x5c, 0x1e, 0x02,
	0x37, 0xa4, 0x67, 0xf0, 0xb0, 0x54, 0x45, 0x86, 0x5a, 0x6f, 0x69, 0x87, 0xae, 0x63, 0x0f, 0x7b,
	0x62, 0xff, 0xb7, 0x00, 0x5a, 0x1c, 0x6f, 0x56, 0xa8, 0x0d, 0x7b, 0x02, 0xed, 0x2c, 0x17, 0x28,
	0x4d, 0xe2, 0xa7, 0x17, 0xf3, 0x43, 0x07, 0x5c, 0x4d, 0xd9, 0x09, 0x34, 0x15, 0xde, 0x24, 0xb2,
	0xa0, 0xd1, 0x35, 0x78, 0xa8, 0xf0, 0x66, 0x5c, 0x58, 0xed, 0x4c, 0x53, 0x93, 0xd2, 0xa0, 0x62,
	0x4e, 0x67, 0xf6, 0x29, 0xb4, 0xb5, 0x98, 0xc9, 0xd4, 0xac, 0x14, 0xd2, 0x48, 0x62, 0x5e, 0x01,
	0xec, 0x11, 0x34, 0xa7, 0x62, 0x86, 0xda, 0xd0, 0x00, 0x62, 0xee, 0xad, 0xfe, 0x1f, 0x01, 0x84,
	0xe7, 0xa9, 0xc9, 0xe6, 0x96, 0xa1, 0x8b, 0x95, 0xca, 0xd0, 0x3f, 0xa1, 0xb7, 0xd8, 0x31, 0x84,
	0x34, 0xf4, 0x4d, 0x05, 0x64, 0xd8, 0xc2, 0xb4, 0x2b, 0xac, 0xee, 0x60, 0x4d, 0x85, 0x55, 0xd7,
	0x34, 0x76, 0xaf, 0x61, 0x5f, 0xc2, 0xa1, 0x72, 0xfd, 0xea, 0x6e, 0x48, 0x1a, 0xfb, 0x5f, 0x25,
	0x54, 0x3f, 0x09, 0xbe, 0xa5, 0xf4, 0x5f, 0x00, 0x5c, 0xcc, 0x31, 0x5b, 0x94, 0x85, 0x90, 0x66,
	0xe7, 0xae, 0x60, 0xf7, 0xae, 0x63, 0x08, 0x6f, 0xd3, 0x7c, 0x85, 0x54, 0x58, 0xcc, 0x9d, 0xd1,
	0xff, 0x33, 0x84, 0x70, 0x74, 0x8b, 0xd2, 0xd8, 0x5a, 0x4c, 0xaa, 0x66, 0x68, 0x36, 0x0d, 0x39,
	0xcb, 0x0e, 0xcf, 0xea, 0xc7, 0xf7, 0x43, 0x67, 0xf6, 0x1c, 0xc2, 0xb4, 0x2c, 0xf3, 0xb5, 0xdf,
	0xa2, 0x4f, 0xf6, 0x17, 0x80, 0xf2, 0x0d, 0x5e, 0x5a, 0xc2, 0xe5, 0x01, 0x77, 0x4c, 0xf6, 0x2d,
	0xb4, 0x14, 0x66, 0x28, 0x6e, 0xdd, 0xb4, 0xa3, 0xe1, 0x93, 0x8f, 0x05, 0x71, 0x47, 0xb9, 0x3c,
	0xe0, 0x1b, 0xb6, 0x0d, 0xf4, 0x72, 0xa0, 0xb7, 0xb8, 0x27, 0xf0, 0xad, 0xa3, 0xd8, 0x40, 0xcf,
	0xf6, 0x81, 0x65, 0xa1, 0x91, 0x14, 0x7a, 0x7f, 0xa0, 0xa5, 0xf8, 0x40, 0x7b, 0x64, 0x03, 0xdb,
	0x71, 0xb6, 0x20, 0xcd, 0x7e, 0xb0, 0xdd, 0x2e, 0xca, 0x6e, 0xcb, 0xe5, 0x01, 0x27, 0x5e, 0xef,
	0xaf, 0x00, 0x42, 0xea, 0x96, 0xbd, 0x80, 0xb8, 0x54, 0xe8, 0x0b, 0x40, 0xab, 0x4f, 0xfb, 0x76,
	0x27, 0xfb, 0x19, 0x36, 0xef, 0xb7, 0x47, 0x65, 0xcf, 0xa1, 0x5d, 0xc5, 0xd5, 0x28, 0xee, 0xff,
	0xfb, 0x71, 0xa4, 0x3b, 0x5e, 0xb1, 0xd8, 0xf7, 0x10, 0x65, 0xdb, 0x67, 0xb7, 0x9f, 0x81, 0x8f,
	0x7c, 0x8c, 0x2a, 0x5d, 0xf0, 0x5d, 0x72, 0xaf, 0x0d, 0x2d, 0x3f, 0xb2, 0xde, 0x05, 0x1d, 0xa9,
	0xf3, 0xef, 0x20, 0xf6, 0xa2, 0x4a, 0x68, 0x61, 0x02, 0x9a, 0xc0, 0xc9, 0x07, 0xda, 0xfb, 0x31,
	0x35, 0x29, 0x8f, 0x54, 0x65, 0xf4, 0xce, 0xed, 0x86, 0xba, 0x07, 0xbb, 0x6f, 0x33, 0x3e, 0x83,
	0xfa, 0x52, 0xcf, 0x48, 0x47, 0xd1, 0xf0, 0x41, 0x95, 0xf3, 0xb5, 0x9e, 0x71, 0xeb, 0xe9, 0x35,
	0xa1, 0x61, 0xe7, 0x7a, 0xde, 0x84, 0x86, 0x59, 0x97, 0xd8, 0xff, 0x3d, 0x80, 0xce, 0xab, 0x62,
	0x36, 0x92, 0x46, 0xad, 0xdf, 0x2a, 0x4c, 0x97, 0x93, 0xdc, 0xca, 0x81, 0x9c, 0x94, 0xfb, 0x68,
	0xf8, 0x74, 0xbf, 0xdb, 0xff, 0xb2, 0x07, 0xd7, 0xeb, 0x12, 0x39, 0x05, 0xd8, 0xb2, 0x72, 0x94,
	0x33, 0x33, 0xf7, 0x9f, 0x55, 0x6f, 0xf5, 0xbf, 0x80, 0x86, 0x65, 0xb1, 0x08, 0x5a, 0x3f, 0x8f,
	0x7f, 0x1a, 0xbf, 0xf9, 0x75, 0xdc, 0x39, 0x60, 0x31, 0x1c, 0xbe, 0xbb, 0x18, 0x8d, 0x5f, 0xf2,
	0xab, 0x37, 0x9d, 0x80, 0xb5, 0x21, 0x1c, 0xfd, 0x32, 0x1a, 0x5f, 0x77, 0x6a, 0x93, 0x26, 0xfd,
	0xa0, 0xbe, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x63, 0xfe, 0xb9, 0xd9, 0x06, 0x00, 0x00,
}
