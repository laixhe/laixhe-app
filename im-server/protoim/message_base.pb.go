// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: message_base.proto

package protoim

import (
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

// 基本信息结构
type MessageBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  CMD    `protobuf:"varint,1,opt,name=cmd,proto3,enum=protoim.CMD" json:"cmd,omitempty"` // 指令值
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                 // 原pd数据
}

func (x *MessageBase) Reset() {
	*x = MessageBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBase) ProtoMessage() {}

func (x *MessageBase) ProtoReflect() protoreflect.Message {
	mi := &file_message_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBase.ProtoReflect.Descriptor instead.
func (*MessageBase) Descriptor() ([]byte, []int) {
	return file_message_base_proto_rawDescGZIP(), []int{0}
}

func (x *MessageBase) GetCmd() CMD {
	if x != nil {
		return x.Cmd
	}
	return CMD_UNKNOWN
}

func (x *MessageBase) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_message_base_proto protoreflect.FileDescriptor

var file_message_base_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x1a, 0x09, 0x63,
	0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x2e, 0x43,
	0x4d, 0x44, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_message_base_proto_rawDescOnce sync.Once
	file_message_base_proto_rawDescData = file_message_base_proto_rawDesc
)

func file_message_base_proto_rawDescGZIP() []byte {
	file_message_base_proto_rawDescOnce.Do(func() {
		file_message_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_base_proto_rawDescData)
	})
	return file_message_base_proto_rawDescData
}

var file_message_base_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_base_proto_goTypes = []interface{}{
	(*MessageBase)(nil), // 0: protoim.MessageBase
	(CMD)(0),            // 1: protoim.CMD
}
var file_message_base_proto_depIdxs = []int32{
	1, // 0: protoim.MessageBase.cmd:type_name -> protoim.CMD
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_base_proto_init() }
func file_message_base_proto_init() {
	if File_message_base_proto != nil {
		return
	}
	file_cmd_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageBase); i {
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
			RawDescriptor: file_message_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_base_proto_goTypes,
		DependencyIndexes: file_message_base_proto_depIdxs,
		MessageInfos:      file_message_base_proto_msgTypes,
	}.Build()
	File_message_base_proto = out.File
	file_message_base_proto_rawDesc = nil
	file_message_base_proto_goTypes = nil
	file_message_base_proto_depIdxs = nil
}
