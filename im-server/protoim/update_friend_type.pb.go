// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: update_friend_type.proto

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

// 更新好友相关信息操作类型
type UpdateFriendType int32

const (
	UpdateFriendType_UF_ONLINE_FALSE UpdateFriendType = 0 // 下线
	UpdateFriendType_UF_ONLINE_TRUE  UpdateFriendType = 1 // 上线
	UpdateFriendType_UF_ADD          UpdateFriendType = 2 // 添加好友
	UpdateFriendType_UF_UPDATE       UpdateFriendType = 3 // 更新好友
	UpdateFriendType_UF_DELETE       UpdateFriendType = 4 // 删除好友
)

// Enum value maps for UpdateFriendType.
var (
	UpdateFriendType_name = map[int32]string{
		0: "UF_ONLINE_FALSE",
		1: "UF_ONLINE_TRUE",
		2: "UF_ADD",
		3: "UF_UPDATE",
		4: "UF_DELETE",
	}
	UpdateFriendType_value = map[string]int32{
		"UF_ONLINE_FALSE": 0,
		"UF_ONLINE_TRUE":  1,
		"UF_ADD":          2,
		"UF_UPDATE":       3,
		"UF_DELETE":       4,
	}
)

func (x UpdateFriendType) Enum() *UpdateFriendType {
	p := new(UpdateFriendType)
	*p = x
	return p
}

func (x UpdateFriendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateFriendType) Descriptor() protoreflect.EnumDescriptor {
	return file_update_friend_type_proto_enumTypes[0].Descriptor()
}

func (UpdateFriendType) Type() protoreflect.EnumType {
	return &file_update_friend_type_proto_enumTypes[0]
}

func (x UpdateFriendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateFriendType.Descriptor instead.
func (UpdateFriendType) EnumDescriptor() ([]byte, []int) {
	return file_update_friend_type_proto_rawDescGZIP(), []int{0}
}

var File_update_friend_type_proto protoreflect.FileDescriptor

var file_update_friend_type_proto_rawDesc = []byte{
	0x0a, 0x18, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x69, 0x6d, 0x2a, 0x65, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x46, 0x5f, 0x4f, 0x4e,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x46, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x46, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x55, 0x46, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x55,
	0x46, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_update_friend_type_proto_rawDescOnce sync.Once
	file_update_friend_type_proto_rawDescData = file_update_friend_type_proto_rawDesc
)

func file_update_friend_type_proto_rawDescGZIP() []byte {
	file_update_friend_type_proto_rawDescOnce.Do(func() {
		file_update_friend_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_update_friend_type_proto_rawDescData)
	})
	return file_update_friend_type_proto_rawDescData
}

var file_update_friend_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_update_friend_type_proto_goTypes = []interface{}{
	(UpdateFriendType)(0), // 0: protoim.UpdateFriendType
}
var file_update_friend_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_update_friend_type_proto_init() }
func file_update_friend_type_proto_init() {
	if File_update_friend_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_update_friend_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_update_friend_type_proto_goTypes,
		DependencyIndexes: file_update_friend_type_proto_depIdxs,
		EnumInfos:         file_update_friend_type_proto_enumTypes,
	}.Build()
	File_update_friend_type_proto = out.File
	file_update_friend_type_proto_rawDesc = nil
	file_update_friend_type_proto_goTypes = nil
	file_update_friend_type_proto_depIdxs = nil
}