// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: error_code_app_os.proto

package protoapi

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 错误码
type ErrorCode int32

const (
	// Success           （成功）
	ErrorCode_E_SUCCESS ErrorCode = 0
	// Unknown error     (未知错误)
	ErrorCode_E_UNKNOWN ErrorCode = 1
	// Server error      (服务器错误)
	ErrorCode_E_SERVER ErrorCode = 2
	// WebSocket error   (WebSocket 错误)
	ErrorCode_E_WEBSOCKET ErrorCode = 3
	// WebSocket write error    (WebSocket 写 错误)
	ErrorCode_E_WEBSOCKET_WRITE ErrorCode = 4
	// WebSocket read error     (WebSocket 读 错误)
	ErrorCode_E_WEBSOCKET_READ ErrorCode = 5
	// WebSocket timeout error  (WebSocket 超时 错误)
	ErrorCode_E_WEBSOCKET_TIMEOUT ErrorCode = 6
	// decode error      (解码错误)
	ErrorCode_E_DECODE ErrorCode = 7
	// encode error      (编码错误)
	ErrorCode_E_ENCODE ErrorCode = 8
	// AUTH error        (认证(鉴权)错误)
	ErrorCode_E_AUTH ErrorCode = 9
	// Not logged in     (未登录)
	ErrorCode_E_NOT_LOGGED_IN ErrorCode = 10
	// Operation failure (操作失败)
	ErrorCode_E_OPERATION_FAILURE ErrorCode = 11
	// Routing not exist (路由不存在)
	ErrorCode_E_ROUTING_NOT_EXIST ErrorCode = 12
	// DB error          (数据库错误)
	ErrorCode_E_DB ErrorCode = 13
	// Not support api   (不支持(无效的) API)
	ErrorCode_E_NOT_SUPPORT_API ErrorCode = 14
	// Invalid json      (无效的 json)
	ErrorCode_E_INVALID_JSON ErrorCode = 15
	// Encode json       (json 编码错误)
	ErrorCode_E_ENCODE_JSON ErrorCode = 16
	// Parameter empty   (参数为空)
	ErrorCode_E_PARAMETER_EMPTY ErrorCode = 17
	// Parameter error   (参数错误)
	ErrorCode_E_PARAMETER ErrorCode = 18
	// Resource error    (资源错误)
	ErrorCode_E_RESOURCE ErrorCode = 19
	// Resource repeat   (资源重复)
	ErrorCode_E_RESOURCE_REPEAT ErrorCode = 20
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:  "E_SUCCESS",
		1:  "E_UNKNOWN",
		2:  "E_SERVER",
		3:  "E_WEBSOCKET",
		4:  "E_WEBSOCKET_WRITE",
		5:  "E_WEBSOCKET_READ",
		6:  "E_WEBSOCKET_TIMEOUT",
		7:  "E_DECODE",
		8:  "E_ENCODE",
		9:  "E_AUTH",
		10: "E_NOT_LOGGED_IN",
		11: "E_OPERATION_FAILURE",
		12: "E_ROUTING_NOT_EXIST",
		13: "E_DB",
		14: "E_NOT_SUPPORT_API",
		15: "E_INVALID_JSON",
		16: "E_ENCODE_JSON",
		17: "E_PARAMETER_EMPTY",
		18: "E_PARAMETER",
		19: "E_RESOURCE",
		20: "E_RESOURCE_REPEAT",
	}
	ErrorCode_value = map[string]int32{
		"E_SUCCESS":           0,
		"E_UNKNOWN":           1,
		"E_SERVER":            2,
		"E_WEBSOCKET":         3,
		"E_WEBSOCKET_WRITE":   4,
		"E_WEBSOCKET_READ":    5,
		"E_WEBSOCKET_TIMEOUT": 6,
		"E_DECODE":            7,
		"E_ENCODE":            8,
		"E_AUTH":              9,
		"E_NOT_LOGGED_IN":     10,
		"E_OPERATION_FAILURE": 11,
		"E_ROUTING_NOT_EXIST": 12,
		"E_DB":                13,
		"E_NOT_SUPPORT_API":   14,
		"E_INVALID_JSON":      15,
		"E_ENCODE_JSON":       16,
		"E_PARAMETER_EMPTY":   17,
		"E_PARAMETER":         18,
		"E_RESOURCE":          19,
		"E_RESOURCE_REPEAT":   20,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_error_code_app_os_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_error_code_app_os_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_error_code_app_os_proto_rawDescGZIP(), []int{0}
}

// 客户端类型
type AppOS int32

const (
	AppOS_O_UNKNOWN AppOS = 0
	AppOS_O_IOS     AppOS = 1
	AppOS_O_ANDROID AppOS = 2
	AppOS_O_WEB     AppOS = 3
	AppOS_O_WINDOW  AppOS = 4
	AppOS_O_MAC     AppOS = 5
	AppOS_O_LINUX   AppOS = 6
)

// Enum value maps for AppOS.
var (
	AppOS_name = map[int32]string{
		0: "O_UNKNOWN",
		1: "O_IOS",
		2: "O_ANDROID",
		3: "O_WEB",
		4: "O_WINDOW",
		5: "O_MAC",
		6: "O_LINUX",
	}
	AppOS_value = map[string]int32{
		"O_UNKNOWN": 0,
		"O_IOS":     1,
		"O_ANDROID": 2,
		"O_WEB":     3,
		"O_WINDOW":  4,
		"O_MAC":     5,
		"O_LINUX":   6,
	}
)

func (x AppOS) Enum() *AppOS {
	p := new(AppOS)
	*p = x
	return p
}

func (x AppOS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppOS) Descriptor() protoreflect.EnumDescriptor {
	return file_error_code_app_os_proto_enumTypes[1].Descriptor()
}

func (AppOS) Type() protoreflect.EnumType {
	return &file_error_code_app_os_proto_enumTypes[1]
}

func (x AppOS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppOS.Descriptor instead.
func (AppOS) EnumDescriptor() ([]byte, []int) {
	return file_error_code_app_os_proto_rawDescGZIP(), []int{1}
}

var File_error_code_app_os_proto protoreflect.FileDescriptor

var file_error_code_app_os_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x70, 0x70,
	0x5f, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x61, 0x70, 0x69, 0x2a, 0x94, 0x03, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a,
	0x0b, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x12, 0x15,
	0x0a, 0x11, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x57, 0x52,
	0x49, 0x54, 0x45, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x4f,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x5f, 0x57, 0x45, 0x42, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x5f, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x08,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10,
	0x0a, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x5f,
	0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x5f, 0x44, 0x42, 0x10, 0x0d, 0x12, 0x15, 0x0a,
	0x11, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x41,
	0x50, 0x49, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x0f, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x5f, 0x45, 0x4e,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x10, 0x12, 0x15, 0x0a, 0x11, 0x45,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59,
	0x10, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45,
	0x52, 0x10, 0x12, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x10, 0x13, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x10, 0x14, 0x2a, 0x61, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x4f, 0x53, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x4f, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x4f, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x5f, 0x57, 0x49, 0x4e,
	0x44, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x5f, 0x4d, 0x41, 0x43, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x06, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x69, 0x78,
	0x68, 0x65, 0x2f, 0x6c, 0x61, 0x69, 0x78, 0x68, 0x65, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x6c, 0x61,
	0x69, 0x78, 0x68, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_code_app_os_proto_rawDescOnce sync.Once
	file_error_code_app_os_proto_rawDescData = file_error_code_app_os_proto_rawDesc
)

func file_error_code_app_os_proto_rawDescGZIP() []byte {
	file_error_code_app_os_proto_rawDescOnce.Do(func() {
		file_error_code_app_os_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_code_app_os_proto_rawDescData)
	})
	return file_error_code_app_os_proto_rawDescData
}

var file_error_code_app_os_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_error_code_app_os_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: protoapi.ErrorCode
	(AppOS)(0),     // 1: protoapi.AppOS
}
var file_error_code_app_os_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_code_app_os_proto_init() }
func file_error_code_app_os_proto_init() {
	if File_error_code_app_os_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_code_app_os_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_code_app_os_proto_goTypes,
		DependencyIndexes: file_error_code_app_os_proto_depIdxs,
		EnumInfos:         file_error_code_app_os_proto_enumTypes,
	}.Build()
	File_error_code_app_os_proto = out.File
	file_error_code_app_os_proto_rawDesc = nil
	file_error_code_app_os_proto_goTypes = nil
	file_error_code_app_os_proto_depIdxs = nil
}
