// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: error_code.proto

package pb

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

type ErrorCode int32

const (
	ErrorCode_ErrorCode_OK            ErrorCode = 0
	ErrorCode_ErrorCode_NotReg        ErrorCode = 11 // 未注册
	ErrorCode_ErrorCode_PasswordError ErrorCode = 12 // 密码错误
	ErrorCode_ErrorCode_DbErr         ErrorCode = 13 // 数据库错误
	ErrorCode_ErrorCode_NameDuplicate ErrorCode = 14 // 重名
	ErrorCode_ErrorCode_HasLogin      ErrorCode = 15
	ErrorCode_ErrorCode_SessionError  ErrorCode = 16
	ErrorCode_ErrorCode_NoPlayer      ErrorCode = 17
	ErrorCode_ErrorCode_TryLater      ErrorCode = 18
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:  "ErrorCode_OK",
		11: "ErrorCode_NotReg",
		12: "ErrorCode_PasswordError",
		13: "ErrorCode_DbErr",
		14: "ErrorCode_NameDuplicate",
		15: "ErrorCode_HasLogin",
		16: "ErrorCode_SessionError",
		17: "ErrorCode_NoPlayer",
		18: "ErrorCode_TryLater",
	}
	ErrorCode_value = map[string]int32{
		"ErrorCode_OK":            0,
		"ErrorCode_NotReg":        11,
		"ErrorCode_PasswordError": 12,
		"ErrorCode_DbErr":         13,
		"ErrorCode_NameDuplicate": 14,
		"ErrorCode_HasLogin":      15,
		"ErrorCode_SessionError":  16,
		"ErrorCode_NoPlayer":      17,
		"ErrorCode_TryLater":      18,
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
	return file_error_code_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_error_code_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_error_code_proto_rawDescGZIP(), []int{0}
}

var File_error_code_proto protoreflect.FileDescriptor

var file_error_code_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2a, 0xe6, 0x01, 0x0a, 0x09,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x67, 0x10,
	0x0b, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0c, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x44, 0x62, 0x45, 0x72,
	0x72, 0x10, 0x0d, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x10, 0x0e,
	0x12, 0x16, 0x0a, 0x12, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x48, 0x61,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x0f, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x5f, 0x4e, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x11, 0x12, 0x16, 0x0a, 0x12,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x54, 0x72, 0x79, 0x4c, 0x61, 0x74,
	0x65, 0x72, 0x10, 0x12, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_code_proto_rawDescOnce sync.Once
	file_error_code_proto_rawDescData = file_error_code_proto_rawDesc
)

func file_error_code_proto_rawDescGZIP() []byte {
	file_error_code_proto_rawDescOnce.Do(func() {
		file_error_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_code_proto_rawDescData)
	})
	return file_error_code_proto_rawDescData
}

var file_error_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_code_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: gserver.ErrorCode
}
var file_error_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_code_proto_init() }
func file_error_code_proto_init() {
	if File_error_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_code_proto_goTypes,
		DependencyIndexes: file_error_code_proto_depIdxs,
		EnumInfos:         file_error_code_proto_enumTypes,
	}.Build()
	File_error_code_proto = out.File
	file_error_code_proto_rawDesc = nil
	file_error_code_proto_goTypes = nil
	file_error_code_proto_depIdxs = nil
}