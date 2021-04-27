// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: ssl_gc_ci.proto

package ci

import (
	api "github.com/RoboCup-SSL/ssl-game-controller/internal/app/api"
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	tracker "github.com/RoboCup-SSL/ssl-game-controller/internal/app/tracker"
	vision "github.com/RoboCup-SSL/ssl-game-controller/internal/app/vision"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The input format to the GC
type CiInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// New unix timestamp in [ns] for the GC
	Timestamp *int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// New tracker packet with ball and robot data
	TrackerPacket *tracker.TrackerWrapperPacket `protobuf:"bytes,2,opt,name=tracker_packet,json=trackerPacket" json:"tracker_packet,omitempty"`
	// (UI) API input
	ApiInputs []*api.Input `protobuf:"bytes,3,rep,name=api_inputs,json=apiInputs" json:"api_inputs,omitempty"`
	// Update geometry
	Geometry *vision.SSL_GeometryData `protobuf:"bytes,4,opt,name=geometry" json:"geometry,omitempty"`
}

func (x *CiInput) Reset() {
	*x = CiInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_ci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CiInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CiInput) ProtoMessage() {}

func (x *CiInput) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_ci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CiInput.ProtoReflect.Descriptor instead.
func (*CiInput) Descriptor() ([]byte, []int) {
	return file_ssl_gc_ci_proto_rawDescGZIP(), []int{0}
}

func (x *CiInput) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *CiInput) GetTrackerPacket() *tracker.TrackerWrapperPacket {
	if x != nil {
		return x.TrackerPacket
	}
	return nil
}

func (x *CiInput) GetApiInputs() []*api.Input {
	if x != nil {
		return x.ApiInputs
	}
	return nil
}

func (x *CiInput) GetGeometry() *vision.SSL_GeometryData {
	if x != nil {
		return x.Geometry
	}
	return nil
}

// The output format of the GC response
type CiOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latest referee message
	RefereeMsg *state.Referee `protobuf:"bytes,1,opt,name=referee_msg,json=refereeMsg" json:"referee_msg,omitempty"`
}

func (x *CiOutput) Reset() {
	*x = CiOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_ci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CiOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CiOutput) ProtoMessage() {}

func (x *CiOutput) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_ci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CiOutput.ProtoReflect.Descriptor instead.
func (*CiOutput) Descriptor() ([]byte, []int) {
	return file_ssl_gc_ci_proto_rawDescGZIP(), []int{1}
}

func (x *CiOutput) GetRefereeMsg() *state.Referee {
	if x != nil {
		return x.RefereeMsg
	}
	return nil
}

var File_ssl_gc_ci_proto protoreflect.FileDescriptor

var file_ssl_gc_ci_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb,
	0x01, 0x0a, 0x07, 0x43, 0x69, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x09, 0x61, 0x70, 0x69, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x2d, 0x0a,
	0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x53, 0x53, 0x4c, 0x5f, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0x35, 0x0a, 0x08,
	0x43, 0x69, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x65,
	0x4d, 0x73, 0x67, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73,
	0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63,
	0x69,
}

var (
	file_ssl_gc_ci_proto_rawDescOnce sync.Once
	file_ssl_gc_ci_proto_rawDescData = file_ssl_gc_ci_proto_rawDesc
)

func file_ssl_gc_ci_proto_rawDescGZIP() []byte {
	file_ssl_gc_ci_proto_rawDescOnce.Do(func() {
		file_ssl_gc_ci_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_ci_proto_rawDescData)
	})
	return file_ssl_gc_ci_proto_rawDescData
}

var file_ssl_gc_ci_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ssl_gc_ci_proto_goTypes = []interface{}{
	(*CiInput)(nil),                      // 0: CiInput
	(*CiOutput)(nil),                     // 1: CiOutput
	(*tracker.TrackerWrapperPacket)(nil), // 2: TrackerWrapperPacket
	(*api.Input)(nil),                    // 3: Input
	(*vision.SSL_GeometryData)(nil),      // 4: SSL_GeometryData
	(*state.Referee)(nil),                // 5: Referee
}
var file_ssl_gc_ci_proto_depIdxs = []int32{
	2, // 0: CiInput.tracker_packet:type_name -> TrackerWrapperPacket
	3, // 1: CiInput.api_inputs:type_name -> Input
	4, // 2: CiInput.geometry:type_name -> SSL_GeometryData
	5, // 3: CiOutput.referee_msg:type_name -> Referee
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ssl_gc_ci_proto_init() }
func file_ssl_gc_ci_proto_init() {
	if File_ssl_gc_ci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_ci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CiInput); i {
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
		file_ssl_gc_ci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CiOutput); i {
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
			RawDescriptor: file_ssl_gc_ci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_ci_proto_goTypes,
		DependencyIndexes: file_ssl_gc_ci_proto_depIdxs,
		MessageInfos:      file_ssl_gc_ci_proto_msgTypes,
	}.Build()
	File_ssl_gc_ci_proto = out.File
	file_ssl_gc_ci_proto_rawDesc = nil
	file_ssl_gc_ci_proto_goTypes = nil
	file_ssl_gc_ci_proto_depIdxs = nil
}
