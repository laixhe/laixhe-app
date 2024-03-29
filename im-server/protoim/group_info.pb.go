// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: group_info.proto

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

// 群信息
type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId   string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GroupName string `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Avatar    string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_group_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupInfo.ProtoReflect.Descriptor instead.
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return file_group_info_proto_rawDescGZIP(), []int{0}
}

func (x *GroupInfo) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupInfo) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_group_info_proto protoreflect.FileDescriptor

var file_group_info_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x22, 0x5d, 0x0a, 0x09, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_info_proto_rawDescOnce sync.Once
	file_group_info_proto_rawDescData = file_group_info_proto_rawDesc
)

func file_group_info_proto_rawDescGZIP() []byte {
	file_group_info_proto_rawDescOnce.Do(func() {
		file_group_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_info_proto_rawDescData)
	})
	return file_group_info_proto_rawDescData
}

var file_group_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_group_info_proto_goTypes = []interface{}{
	(*GroupInfo)(nil), // 0: protoim.GroupInfo
}
var file_group_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_group_info_proto_init() }
func file_group_info_proto_init() {
	if File_group_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupInfo); i {
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
			RawDescriptor: file_group_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_group_info_proto_goTypes,
		DependencyIndexes: file_group_info_proto_depIdxs,
		MessageInfos:      file_group_info_proto_msgTypes,
	}.Build()
	File_group_info_proto = out.File
	file_group_info_proto_rawDesc = nil
	file_group_info_proto_goTypes = nil
	file_group_info_proto_depIdxs = nil
}
