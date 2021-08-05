// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: metrics/types/types.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=Event,proto3" json:"Event,omitempty"`
}

func (x *StartEvent) Reset() {
	*x = StartEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_types_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartEvent) ProtoMessage() {}

func (x *StartEvent) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_types_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartEvent.ProtoReflect.Descriptor instead.
func (*StartEvent) Descriptor() ([]byte, []int) {
	return file_metrics_types_types_proto_rawDescGZIP(), []int{0}
}

func (x *StartEvent) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// Event is the basic type that is recorded by hotstuff.
// It contains the ID of the replica/client, the type (replica/client),
// the timestamp of the event, and the data.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Client    bool                   `protobuf:"varint,2,opt,name=Client,proto3" json:"Client,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_types_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_types_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_metrics_types_types_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Event) GetClient() bool {
	if x != nil {
		return x.Client
	}
	return false
}

func (x *Event) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type ThroughputMeasurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event    *Event               `protobuf:"bytes,1,opt,name=Event,proto3" json:"Event,omitempty"`
	Commits  uint64               `protobuf:"varint,2,opt,name=Commits,proto3" json:"Commits,omitempty"`
	Commands uint64               `protobuf:"varint,3,opt,name=Commands,proto3" json:"Commands,omitempty"`
	Duration *durationpb.Duration `protobuf:"bytes,4,opt,name=Duration,proto3" json:"Duration,omitempty"`
}

func (x *ThroughputMeasurement) Reset() {
	*x = ThroughputMeasurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_types_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThroughputMeasurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThroughputMeasurement) ProtoMessage() {}

func (x *ThroughputMeasurement) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_types_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThroughputMeasurement.ProtoReflect.Descriptor instead.
func (*ThroughputMeasurement) Descriptor() ([]byte, []int) {
	return file_metrics_types_types_proto_rawDescGZIP(), []int{2}
}

func (x *ThroughputMeasurement) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ThroughputMeasurement) GetCommits() uint64 {
	if x != nil {
		return x.Commits
	}
	return 0
}

func (x *ThroughputMeasurement) GetCommands() uint64 {
	if x != nil {
		return x.Commands
	}
	return 0
}

func (x *ThroughputMeasurement) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type LatencyMeasurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event    *Event  `protobuf:"bytes,1,opt,name=Event,proto3" json:"Event,omitempty"`
	Latency  float64 `protobuf:"fixed64,2,opt,name=Latency,proto3" json:"Latency,omitempty"`
	Variance float64 `protobuf:"fixed64,3,opt,name=Variance,proto3" json:"Variance,omitempty"`
}

func (x *LatencyMeasurement) Reset() {
	*x = LatencyMeasurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_types_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatencyMeasurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatencyMeasurement) ProtoMessage() {}

func (x *LatencyMeasurement) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_types_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatencyMeasurement.ProtoReflect.Descriptor instead.
func (*LatencyMeasurement) Descriptor() ([]byte, []int) {
	return file_metrics_types_types_proto_rawDescGZIP(), []int{3}
}

func (x *LatencyMeasurement) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *LatencyMeasurement) GetLatency() float64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *LatencyMeasurement) GetVariance() float64 {
	if x != nil {
		return x.Variance
	}
	return 0
}

var File_metrics_types_types_proto protoreflect.FileDescriptor

var file_metrics_types_types_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0xa8, 0x01, 0x0a, 0x15, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x70, 0x75, 0x74, 0x4d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x12, 0x4c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x62, 0x2f,
	0x68, 0x6f, 0x74, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metrics_types_types_proto_rawDescOnce sync.Once
	file_metrics_types_types_proto_rawDescData = file_metrics_types_types_proto_rawDesc
)

func file_metrics_types_types_proto_rawDescGZIP() []byte {
	file_metrics_types_types_proto_rawDescOnce.Do(func() {
		file_metrics_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_metrics_types_types_proto_rawDescData)
	})
	return file_metrics_types_types_proto_rawDescData
}

var file_metrics_types_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_metrics_types_types_proto_goTypes = []interface{}{
	(*StartEvent)(nil),            // 0: types.StartEvent
	(*Event)(nil),                 // 1: types.Event
	(*ThroughputMeasurement)(nil), // 2: types.ThroughputMeasurement
	(*LatencyMeasurement)(nil),    // 3: types.LatencyMeasurement
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 5: google.protobuf.Duration
}
var file_metrics_types_types_proto_depIdxs = []int32{
	1, // 0: types.StartEvent.Event:type_name -> types.Event
	4, // 1: types.Event.Timestamp:type_name -> google.protobuf.Timestamp
	1, // 2: types.ThroughputMeasurement.Event:type_name -> types.Event
	5, // 3: types.ThroughputMeasurement.Duration:type_name -> google.protobuf.Duration
	1, // 4: types.LatencyMeasurement.Event:type_name -> types.Event
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_metrics_types_types_proto_init() }
func file_metrics_types_types_proto_init() {
	if File_metrics_types_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metrics_types_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metrics_types_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metrics_types_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThroughputMeasurement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metrics_types_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatencyMeasurement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metrics_types_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metrics_types_types_proto_goTypes,
		DependencyIndexes: file_metrics_types_types_proto_depIdxs,
		MessageInfos:      file_metrics_types_types_proto_msgTypes,
	}.Build()
	File_metrics_types_types_proto = out.File
	file_metrics_types_types_proto_rawDesc = nil
	file_metrics_types_types_proto_goTypes = nil
	file_metrics_types_types_proto_depIdxs = nil
}
