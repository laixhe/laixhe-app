// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: message_send_request.proto

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

// 客户端发送消息-请求
type MessageSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalMessageId string          `protobuf:"bytes,1,opt,name=local_message_id,json=localMessageId,proto3" json:"local_message_id,omitempty"` // 客户端本地的消息ID(用于客户校验)(客户端生成)
	Content        *MessageContent `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                                       // 消息内容
}

func (x *MessageSendRequest) Reset() {
	*x = MessageSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_send_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendRequest) ProtoMessage() {}

func (x *MessageSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_send_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendRequest.ProtoReflect.Descriptor instead.
func (*MessageSendRequest) Descriptor() ([]byte, []int) {
	return file_message_send_request_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSendRequest) GetLocalMessageId() string {
	if x != nil {
		return x.LocalMessageId
	}
	return ""
}

func (x *MessageSendRequest) GetContent() *MessageContent {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_message_send_request_proto protoreflect.FileDescriptor

var file_message_send_request_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x1a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x12,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_send_request_proto_rawDescOnce sync.Once
	file_message_send_request_proto_rawDescData = file_message_send_request_proto_rawDesc
)

func file_message_send_request_proto_rawDescGZIP() []byte {
	file_message_send_request_proto_rawDescOnce.Do(func() {
		file_message_send_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_send_request_proto_rawDescData)
	})
	return file_message_send_request_proto_rawDescData
}

var file_message_send_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_send_request_proto_goTypes = []interface{}{
	(*MessageSendRequest)(nil), // 0: protoim.MessageSendRequest
	(*MessageContent)(nil),     // 1: protoim.MessageContent
}
var file_message_send_request_proto_depIdxs = []int32{
	1, // 0: protoim.MessageSendRequest.content:type_name -> protoim.MessageContent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_send_request_proto_init() }
func file_message_send_request_proto_init() {
	if File_message_send_request_proto != nil {
		return
	}
	file_message_content_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_send_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendRequest); i {
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
			RawDescriptor: file_message_send_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_send_request_proto_goTypes,
		DependencyIndexes: file_message_send_request_proto_depIdxs,
		MessageInfos:      file_message_send_request_proto_msgTypes,
	}.Build()
	File_message_send_request_proto = out.File
	file_message_send_request_proto_rawDesc = nil
	file_message_send_request_proto_goTypes = nil
	file_message_send_request_proto_depIdxs = nil
}