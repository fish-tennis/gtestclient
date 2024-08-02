// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: cmd_client.proto

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

type CmdClient int32

const (
	CmdClient_CmdClient_None CmdClient = 0 // 解决"The first enum value must be zero in proto3."的报错
	// Login
	CmdClient_Cmd_LoginReq           CmdClient = 1001
	CmdClient_Cmd_LoginRes           CmdClient = 1002
	CmdClient_Cmd_AccountReg         CmdClient = 1003
	CmdClient_Cmd_AccountRes         CmdClient = 1004
	CmdClient_Cmd_PlayerEntryGameReq CmdClient = 1005
	CmdClient_Cmd_PlayerEntryGameRes CmdClient = 1006
	CmdClient_Cmd_CreatePlayerReq    CmdClient = 1007
	CmdClient_Cmd_CreatePlayerRes    CmdClient = 1008
	// Money
	CmdClient_Cmd_CoinReq CmdClient = 1101 // 请求加coin的测试消息
	CmdClient_Cmd_CoinRes CmdClient = 1102
	// Quest
	CmdClient_Cmd_FinishQuestReq CmdClient = 1201 // 完成任务
	CmdClient_Cmd_FinishQuestRes CmdClient = 1202 // 完成任务的返回结果
	// Guild
	CmdClient_Cmd_GuildListReq         CmdClient = 2001 // 公会列表查看
	CmdClient_Cmd_GuildListRes         CmdClient = 2002
	CmdClient_Cmd_GuildCreateReq       CmdClient = 2003 // 创建公会
	CmdClient_Cmd_GuildCreateRes       CmdClient = 2004
	CmdClient_Cmd_GuildJoinReq         CmdClient = 2005 // 请求加入某个公会
	CmdClient_Cmd_GuildJoinRes         CmdClient = 2006
	CmdClient_Cmd_GuildJoinAgreeReq    CmdClient = 2007 // 公会管理员处理入会请求
	CmdClient_Cmd_GuildJoinAgreeRes    CmdClient = 2008
	CmdClient_Cmd_GuildDataViewReq     CmdClient = 2009 // 查看本公会的数据
	CmdClient_Cmd_GuildDataViewRes     CmdClient = 2010
	CmdClient_Cmd_GuildJoinReqTip      CmdClient = 2012 // 公会入户申请的提示
	CmdClient_Cmd_GuildJoinReqOpResult CmdClient = 2013 // 自己的入会申请的操作结果
)

// Enum value maps for CmdClient.
var (
	CmdClient_name = map[int32]string{
		0:    "CmdClient_None",
		1001: "Cmd_LoginReq",
		1002: "Cmd_LoginRes",
		1003: "Cmd_AccountReg",
		1004: "Cmd_AccountRes",
		1005: "Cmd_PlayerEntryGameReq",
		1006: "Cmd_PlayerEntryGameRes",
		1007: "Cmd_CreatePlayerReq",
		1008: "Cmd_CreatePlayerRes",
		1101: "Cmd_CoinReq",
		1102: "Cmd_CoinRes",
		1201: "Cmd_FinishQuestReq",
		1202: "Cmd_FinishQuestRes",
		2001: "Cmd_GuildListReq",
		2002: "Cmd_GuildListRes",
		2003: "Cmd_GuildCreateReq",
		2004: "Cmd_GuildCreateRes",
		2005: "Cmd_GuildJoinReq",
		2006: "Cmd_GuildJoinRes",
		2007: "Cmd_GuildJoinAgreeReq",
		2008: "Cmd_GuildJoinAgreeRes",
		2009: "Cmd_GuildDataViewReq",
		2010: "Cmd_GuildDataViewRes",
		2012: "Cmd_GuildJoinReqTip",
		2013: "Cmd_GuildJoinReqOpResult",
	}
	CmdClient_value = map[string]int32{
		"CmdClient_None":           0,
		"Cmd_LoginReq":             1001,
		"Cmd_LoginRes":             1002,
		"Cmd_AccountReg":           1003,
		"Cmd_AccountRes":           1004,
		"Cmd_PlayerEntryGameReq":   1005,
		"Cmd_PlayerEntryGameRes":   1006,
		"Cmd_CreatePlayerReq":      1007,
		"Cmd_CreatePlayerRes":      1008,
		"Cmd_CoinReq":              1101,
		"Cmd_CoinRes":              1102,
		"Cmd_FinishQuestReq":       1201,
		"Cmd_FinishQuestRes":       1202,
		"Cmd_GuildListReq":         2001,
		"Cmd_GuildListRes":         2002,
		"Cmd_GuildCreateReq":       2003,
		"Cmd_GuildCreateRes":       2004,
		"Cmd_GuildJoinReq":         2005,
		"Cmd_GuildJoinRes":         2006,
		"Cmd_GuildJoinAgreeReq":    2007,
		"Cmd_GuildJoinAgreeRes":    2008,
		"Cmd_GuildDataViewReq":     2009,
		"Cmd_GuildDataViewRes":     2010,
		"Cmd_GuildJoinReqTip":      2012,
		"Cmd_GuildJoinReqOpResult": 2013,
	}
)

func (x CmdClient) Enum() *CmdClient {
	p := new(CmdClient)
	*p = x
	return p
}

func (x CmdClient) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdClient) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_client_proto_enumTypes[0].Descriptor()
}

func (CmdClient) Type() protoreflect.EnumType {
	return &file_cmd_client_proto_enumTypes[0]
}

func (x CmdClient) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdClient.Descriptor instead.
func (CmdClient) EnumDescriptor() ([]byte, []int) {
	return file_cmd_client_proto_rawDescGZIP(), []int{0}
}

var File_cmd_client_proto protoreflect.FileDescriptor

var file_cmd_client_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6d, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2a, 0xe8, 0x04, 0x0a, 0x09,
	0x43, 0x6d, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6d, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0c, 0x43, 0x6d, 0x64, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x10, 0xe9, 0x07,
	0x12, 0x11, 0x0a, 0x0c, 0x43, 0x6d, 0x64, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x10, 0xea, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x6d, 0x64, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x67, 0x10, 0xeb, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x6d, 0x64, 0x5f,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x10, 0xec, 0x07, 0x12, 0x1b, 0x0a,
	0x16, 0x43, 0x6d, 0x64, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x10, 0xed, 0x07, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x6d,
	0x64, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x10, 0xee, 0x07, 0x12, 0x18, 0x0a, 0x13, 0x43, 0x6d, 0x64, 0x5f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x10, 0xef,
	0x07, 0x12, 0x18, 0x0a, 0x13, 0x43, 0x6d, 0x64, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x10, 0xf0, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x43,
	0x6d, 0x64, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x10, 0xcd, 0x08, 0x12, 0x10, 0x0a,
	0x0b, 0x43, 0x6d, 0x64, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x10, 0xce, 0x08, 0x12,
	0x17, 0x0a, 0x12, 0x43, 0x6d, 0x64, 0x5f, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x10, 0xb1, 0x09, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x6d, 0x64, 0x5f,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x10, 0xb2,
	0x09, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x6d, 0x64, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x10, 0xd1, 0x0f, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x6d, 0x64, 0x5f,
	0x47, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x10, 0xd2, 0x0f, 0x12,
	0x17, 0x0a, 0x12, 0x43, 0x6d, 0x64, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x10, 0xd3, 0x0f, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x6d, 0x64, 0x5f,
	0x47, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x10, 0xd4,
	0x0f, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x6d, 0x64, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x10, 0xd5, 0x0f, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x6d, 0x64, 0x5f,
	0x47, 0x75, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x10, 0xd6, 0x0f, 0x12,
	0x1a, 0x0a, 0x15, 0x43, 0x6d, 0x64, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x69, 0x6e,
	0x41, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x10, 0xd7, 0x0f, 0x12, 0x1a, 0x0a, 0x15, 0x43,
	0x6d, 0x64, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x10, 0xd8, 0x0f, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x6d, 0x64, 0x5f, 0x47,
	0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x10,
	0xd9, 0x0f, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x6d, 0x64, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x10, 0xda, 0x0f, 0x12, 0x18, 0x0a,
	0x13, 0x43, 0x6d, 0x64, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x54, 0x69, 0x70, 0x10, 0xdc, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x43, 0x6d, 0x64, 0x5f, 0x47,
	0x75, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x4f, 0x70, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x10, 0xdd, 0x0f, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_client_proto_rawDescOnce sync.Once
	file_cmd_client_proto_rawDescData = file_cmd_client_proto_rawDesc
)

func file_cmd_client_proto_rawDescGZIP() []byte {
	file_cmd_client_proto_rawDescOnce.Do(func() {
		file_cmd_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_client_proto_rawDescData)
	})
	return file_cmd_client_proto_rawDescData
}

var file_cmd_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_client_proto_goTypes = []interface{}{
	(CmdClient)(0), // 0: gserver.CmdClient
}
var file_cmd_client_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_client_proto_init() }
func file_cmd_client_proto_init() {
	if File_cmd_client_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_client_proto_goTypes,
		DependencyIndexes: file_cmd_client_proto_depIdxs,
		EnumInfos:         file_cmd_client_proto_enumTypes,
	}.Build()
	File_cmd_client_proto = out.File
	file_cmd_client_proto_rawDesc = nil
	file_cmd_client_proto_goTypes = nil
	file_cmd_client_proto_depIdxs = nil
}
