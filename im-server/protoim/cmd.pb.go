// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: cmd.proto

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

// 指令类型(路由)
type CMD int32

const (
	CMD_C_UNKNOWN                    CMD = 0    // 未知
	CMD_C_ERROR                      CMD = 1    // 错误
	CMD_C_PING                       CMD = 2    // 心跳-ping
	CMD_C_PONG                       CMD = 3    // 心跳-pong
	CMD_C_USER_LOGIN_REQUEST         CMD = 1000 // 用户登录-请求
	CMD_C_USER_LOGIN_RESPONSE        CMD = 1001 // 用户登录-响应
	CMD_C_GET_USER_REQUEST           CMD = 1002 // 获取用户信息-请求
	CMD_C_GET_USER_RESPONSE          CMD = 1003 // 获取用户信息-响应
	CMD_C_GET_FRIENDS_REQUEST        CMD = 1004 // 获取好友列表-请求
	CMD_C_GET_FRIENDS_RESPONSE       CMD = 1005 // 获取好友列表-响应
	CMD_C_UPDATE_FRIENDS             CMD = 1006 // 更新好友相关信息
	CMD_C_GET_CONVERSATIONS_REQUEST  CMD = 1007 // 消息会话列表-请求
	CMD_C_GET_CONVERSATIONS_RESPONSE CMD = 1008 // 消息会话列表-响应
	CMD_C_MESSAGE_REQUEST            CMD = 1009 // 消息相关(客户端发送-服务端接收)-请求
	CMD_C_MESSAGE_RESPONSE           CMD = 1010 // 消息相关(服务端发送-客户端接收)-响应
	CMD_C_MESSAGE_FEEDBACK           CMD = 1011 // 消息相关-反馈(客户端发送给服务端接收后响应接收该消息，则反)
)

// Enum value maps for CMD.
var (
	CMD_name = map[int32]string{
		0:    "C_UNKNOWN",
		1:    "C_ERROR",
		2:    "C_PING",
		3:    "C_PONG",
		1000: "C_USER_LOGIN_REQUEST",
		1001: "C_USER_LOGIN_RESPONSE",
		1002: "C_GET_USER_REQUEST",
		1003: "C_GET_USER_RESPONSE",
		1004: "C_GET_FRIENDS_REQUEST",
		1005: "C_GET_FRIENDS_RESPONSE",
		1006: "C_UPDATE_FRIENDS",
		1007: "C_GET_CONVERSATIONS_REQUEST",
		1008: "C_GET_CONVERSATIONS_RESPONSE",
		1009: "C_MESSAGE_REQUEST",
		1010: "C_MESSAGE_RESPONSE",
		1011: "C_MESSAGE_FEEDBACK",
	}
	CMD_value = map[string]int32{
		"C_UNKNOWN":                    0,
		"C_ERROR":                      1,
		"C_PING":                       2,
		"C_PONG":                       3,
		"C_USER_LOGIN_REQUEST":         1000,
		"C_USER_LOGIN_RESPONSE":        1001,
		"C_GET_USER_REQUEST":           1002,
		"C_GET_USER_RESPONSE":          1003,
		"C_GET_FRIENDS_REQUEST":        1004,
		"C_GET_FRIENDS_RESPONSE":       1005,
		"C_UPDATE_FRIENDS":             1006,
		"C_GET_CONVERSATIONS_REQUEST":  1007,
		"C_GET_CONVERSATIONS_RESPONSE": 1008,
		"C_MESSAGE_REQUEST":            1009,
		"C_MESSAGE_RESPONSE":           1010,
		"C_MESSAGE_FEEDBACK":           1011,
	}
)

func (x CMD) Enum() *CMD {
	p := new(CMD)
	*p = x
	return p
}

func (x CMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMD) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[0].Descriptor()
}

func (CMD) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[0]
}

func (x CMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMD.Descriptor instead.
func (CMD) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

var File_cmd_proto protoreflect.FileDescriptor

var file_cmd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x69, 0x6d, 0x2a, 0x82, 0x03, 0x0a, 0x03, 0x43, 0x4d, 0x44, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x50, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x19, 0x0a, 0x14, 0x43, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xe8, 0x07, 0x12, 0x1a, 0x0a, 0x15, 0x43,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x10, 0xe9, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x5f, 0x47, 0x45, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xea, 0x07,
	0x12, 0x18, 0x0a, 0x13, 0x43, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xeb, 0x07, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x5f,
	0x47, 0x45, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x10, 0xec, 0x07, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x5f, 0x47, 0x45, 0x54, 0x5f,
	0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x10, 0xed, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x10, 0xee, 0x07, 0x12, 0x20, 0x0a, 0x1b, 0x43, 0x5f,
	0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xef, 0x07, 0x12, 0x21, 0x0a, 0x1c,
	0x43, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xf0, 0x07, 0x12,
	0x16, 0x0a, 0x11, 0x43, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0xf1, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x5f, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xf2, 0x07,
	0x12, 0x17, 0x0a, 0x12, 0x43, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x45,
	0x45, 0x44, 0x42, 0x41, 0x43, 0x4b, 0x10, 0xf3, 0x07, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_proto_rawDescOnce sync.Once
	file_cmd_proto_rawDescData = file_cmd_proto_rawDesc
)

func file_cmd_proto_rawDescGZIP() []byte {
	file_cmd_proto_rawDescOnce.Do(func() {
		file_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_proto_rawDescData)
	})
	return file_cmd_proto_rawDescData
}

var file_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_proto_goTypes = []interface{}{
	(CMD)(0), // 0: protoim.CMD
}
var file_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_proto_init() }
func file_cmd_proto_init() {
	if File_cmd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_proto_goTypes,
		DependencyIndexes: file_cmd_proto_depIdxs,
		EnumInfos:         file_cmd_proto_enumTypes,
	}.Build()
	File_cmd_proto = out.File
	file_cmd_proto_rawDesc = nil
	file_cmd_proto_goTypes = nil
	file_cmd_proto_depIdxs = nil
}
