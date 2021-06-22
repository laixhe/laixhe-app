// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: get_friends_request.proto

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

// 获取好友列表-请求
type GetFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFriendsRequest) Reset() {
	*x = GetFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_friends_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendsRequest) ProtoMessage() {}

func (x *GetFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_friends_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendsRequest.ProtoReflect.Descriptor instead.
func (*GetFriendsRequest) Descriptor() ([]byte, []int) {
	return file_get_friends_request_proto_rawDescGZIP(), []int{0}
}

var File_get_friends_request_proto protoreflect.FileDescriptor

var file_get_friends_request_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x69, 0x6d, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_friends_request_proto_rawDescOnce sync.Once
	file_get_friends_request_proto_rawDescData = file_get_friends_request_proto_rawDesc
)

func file_get_friends_request_proto_rawDescGZIP() []byte {
	file_get_friends_request_proto_rawDescOnce.Do(func() {
		file_get_friends_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_friends_request_proto_rawDescData)
	})
	return file_get_friends_request_proto_rawDescData
}

var file_get_friends_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_get_friends_request_proto_goTypes = []interface{}{
	(*GetFriendsRequest)(nil), // 0: protoim.GetFriendsRequest
}
var file_get_friends_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_get_friends_request_proto_init() }
func file_get_friends_request_proto_init() {
	if File_get_friends_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_get_friends_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendsRequest); i {
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
			RawDescriptor: file_get_friends_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_friends_request_proto_goTypes,
		DependencyIndexes: file_get_friends_request_proto_depIdxs,
		MessageInfos:      file_get_friends_request_proto_msgTypes,
	}.Build()
	File_get_friends_request_proto = out.File
	file_get_friends_request_proto_rawDesc = nil
	file_get_friends_request_proto_goTypes = nil
	file_get_friends_request_proto_depIdxs = nil
}