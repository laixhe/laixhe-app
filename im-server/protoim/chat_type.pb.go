// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: chat_type.proto

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

type ChatType int32

const (
	ChatType_USE_CHAT   ChatType = 0 // 私聊
	ChatType_GROUP_CHAT ChatType = 1 // 群聊
)

// Enum value maps for ChatType.
var (
	ChatType_name = map[int32]string{
		0: "USE_CHAT",
		1: "GROUP_CHAT",
	}
	ChatType_value = map[string]int32{
		"USE_CHAT":   0,
		"GROUP_CHAT": 1,
	}
)

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}

func (x ChatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatType) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_type_proto_enumTypes[0].Descriptor()
}

func (ChatType) Type() protoreflect.EnumType {
	return &file_chat_type_proto_enumTypes[0]
}

func (x ChatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatType.Descriptor instead.
func (ChatType) EnumDescriptor() ([]byte, []int) {
	return file_chat_type_proto_rawDescGZIP(), []int{0}
}

var File_chat_type_proto protoreflect.FileDescriptor

var file_chat_type_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x2a, 0x28, 0x0a, 0x08, 0x43, 0x68,
	0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x5f, 0x43, 0x48,
	0x41, 0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x48,
	0x41, 0x54, 0x10, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_type_proto_rawDescOnce sync.Once
	file_chat_type_proto_rawDescData = file_chat_type_proto_rawDesc
)

func file_chat_type_proto_rawDescGZIP() []byte {
	file_chat_type_proto_rawDescOnce.Do(func() {
		file_chat_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_type_proto_rawDescData)
	})
	return file_chat_type_proto_rawDescData
}

var file_chat_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chat_type_proto_goTypes = []interface{}{
	(ChatType)(0), // 0: protoim.ChatType
}
var file_chat_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_type_proto_init() }
func file_chat_type_proto_init() {
	if File_chat_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_type_proto_goTypes,
		DependencyIndexes: file_chat_type_proto_depIdxs,
		EnumInfos:         file_chat_type_proto_enumTypes,
	}.Build()
	File_chat_type_proto = out.File
	file_chat_type_proto_rawDesc = nil
	file_chat_type_proto_goTypes = nil
	file_chat_type_proto_depIdxs = nil
}
