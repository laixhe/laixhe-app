// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: app_os.proto

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

type AppOs int32

const (
	AppOs_OsUnknown AppOs = 0 // 未知
	AppOs_Web       AppOs = 1 // web
	AppOs_Ios       AppOs = 2 // ios
	AppOs_Android   AppOs = 3 // android
	AppOs_Pc        AppOs = 4 // pc
)

// Enum value maps for AppOs.
var (
	AppOs_name = map[int32]string{
		0: "OsUnknown",
		1: "Web",
		2: "Ios",
		3: "Android",
		4: "Pc",
	}
	AppOs_value = map[string]int32{
		"OsUnknown": 0,
		"Web":       1,
		"Ios":       2,
		"Android":   3,
		"Pc":        4,
	}
)

func (x AppOs) Enum() *AppOs {
	p := new(AppOs)
	*p = x
	return p
}

func (x AppOs) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppOs) Descriptor() protoreflect.EnumDescriptor {
	return file_app_os_proto_enumTypes[0].Descriptor()
}

func (AppOs) Type() protoreflect.EnumType {
	return &file_app_os_proto_enumTypes[0]
}

func (x AppOs) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppOs.Descriptor instead.
func (AppOs) EnumDescriptor() ([]byte, []int) {
	return file_app_os_proto_rawDescGZIP(), []int{0}
}

var File_app_os_proto protoreflect.FileDescriptor

var file_app_os_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x70, 0x70, 0x5f, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6d, 0x2a, 0x3d, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x4f, 0x73,
	0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x73, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x57, 0x65, 0x62, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x6f, 0x73, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x10, 0x03, 0x12, 0x06,
	0x0a, 0x02, 0x50, 0x63, 0x10, 0x04, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_os_proto_rawDescOnce sync.Once
	file_app_os_proto_rawDescData = file_app_os_proto_rawDesc
)

func file_app_os_proto_rawDescGZIP() []byte {
	file_app_os_proto_rawDescOnce.Do(func() {
		file_app_os_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_os_proto_rawDescData)
	})
	return file_app_os_proto_rawDescData
}

var file_app_os_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_os_proto_goTypes = []interface{}{
	(AppOs)(0), // 0: protoim.AppOs
}
var file_app_os_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_os_proto_init() }
func file_app_os_proto_init() {
	if File_app_os_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_os_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_os_proto_goTypes,
		DependencyIndexes: file_app_os_proto_depIdxs,
		EnumInfos:         file_app_os_proto_enumTypes,
	}.Build()
	File_app_os_proto = out.File
	file_app_os_proto_rawDesc = nil
	file_app_os_proto_goTypes = nil
	file_app_os_proto_depIdxs = nil
}
