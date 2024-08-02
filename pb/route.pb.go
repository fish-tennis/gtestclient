// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: route.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 路由转发给玩家的消息
// server -> otherserver -> player
type RoutePlayerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error            string     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`                        // 错误码
	ToPlayerId       int64      `protobuf:"varint,2,opt,name=toPlayerId,proto3" json:"toPlayerId,omitempty"`             // 玩家id
	PacketCommand    int32      `protobuf:"varint,3,opt,name=packetCommand,proto3" json:"packetCommand,omitempty"`       // 消息号
	DirectSendClient bool       `protobuf:"varint,4,opt,name=directSendClient,proto3" json:"directSendClient,omitempty"` // 是否直接转发给客户端
	PendingMessageId int64      `protobuf:"varint,5,opt,name=pendingMessageId,proto3" json:"pendingMessageId,omitempty"` // 待处理消息id
	PacketData       *anypb.Any `protobuf:"bytes,6,opt,name=packetData,proto3" json:"packetData,omitempty"`              // 转发的消息
}

func (x *RoutePlayerMessage) Reset() {
	*x = RoutePlayerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutePlayerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutePlayerMessage) ProtoMessage() {}

func (x *RoutePlayerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutePlayerMessage.ProtoReflect.Descriptor instead.
func (*RoutePlayerMessage) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{0}
}

func (x *RoutePlayerMessage) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *RoutePlayerMessage) GetToPlayerId() int64 {
	if x != nil {
		return x.ToPlayerId
	}
	return 0
}

func (x *RoutePlayerMessage) GetPacketCommand() int32 {
	if x != nil {
		return x.PacketCommand
	}
	return 0
}

func (x *RoutePlayerMessage) GetDirectSendClient() bool {
	if x != nil {
		return x.DirectSendClient
	}
	return false
}

func (x *RoutePlayerMessage) GetPendingMessageId() int64 {
	if x != nil {
		return x.PendingMessageId
	}
	return 0
}

func (x *RoutePlayerMessage) GetPacketData() *anypb.Any {
	if x != nil {
		return x.PacketData
	}
	return nil
}

// 路由转发玩家的公会请求消息
// server -> otherserver -> guild
type GuildRoutePlayerMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromPlayerId   int64      `protobuf:"varint,1,opt,name=fromPlayerId,proto3" json:"fromPlayerId,omitempty"`    // 玩家id
	FromGuildId    int64      `protobuf:"varint,2,opt,name=fromGuildId,proto3" json:"fromGuildId,omitempty"`      // 玩家公会id
	FromServerId   int32      `protobuf:"varint,3,opt,name=fromServerId,proto3" json:"fromServerId,omitempty"`    // 玩家当前所在服务器id
	FromPlayerName string     `protobuf:"bytes,4,opt,name=fromPlayerName,proto3" json:"fromPlayerName,omitempty"` // 玩家名
	PacketCommand  int32      `protobuf:"varint,5,opt,name=packetCommand,proto3" json:"packetCommand,omitempty"`  // 消息号
	PacketData     *anypb.Any `protobuf:"bytes,6,opt,name=packetData,proto3" json:"packetData,omitempty"`         // 消息内容
}

func (x *GuildRoutePlayerMessageReq) Reset() {
	*x = GuildRoutePlayerMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuildRoutePlayerMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuildRoutePlayerMessageReq) ProtoMessage() {}

func (x *GuildRoutePlayerMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuildRoutePlayerMessageReq.ProtoReflect.Descriptor instead.
func (*GuildRoutePlayerMessageReq) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{1}
}

func (x *GuildRoutePlayerMessageReq) GetFromPlayerId() int64 {
	if x != nil {
		return x.FromPlayerId
	}
	return 0
}

func (x *GuildRoutePlayerMessageReq) GetFromGuildId() int64 {
	if x != nil {
		return x.FromGuildId
	}
	return 0
}

func (x *GuildRoutePlayerMessageReq) GetFromServerId() int32 {
	if x != nil {
		return x.FromServerId
	}
	return 0
}

func (x *GuildRoutePlayerMessageReq) GetFromPlayerName() string {
	if x != nil {
		return x.FromPlayerName
	}
	return ""
}

func (x *GuildRoutePlayerMessageReq) GetPacketCommand() int32 {
	if x != nil {
		return x.PacketCommand
	}
	return 0
}

func (x *GuildRoutePlayerMessageReq) GetPacketData() *anypb.Any {
	if x != nil {
		return x.PacketData
	}
	return nil
}

var File_route_proto protoreflect.FileDescriptor

var file_route_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x8a, 0x02, 0x0a, 0x1a, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x47, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d,
	0x47, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x66,
	0x72, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_route_proto_rawDescOnce sync.Once
	file_route_proto_rawDescData = file_route_proto_rawDesc
)

func file_route_proto_rawDescGZIP() []byte {
	file_route_proto_rawDescOnce.Do(func() {
		file_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_proto_rawDescData)
	})
	return file_route_proto_rawDescData
}

var file_route_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_route_proto_goTypes = []interface{}{
	(*RoutePlayerMessage)(nil),         // 0: gserver.RoutePlayerMessage
	(*GuildRoutePlayerMessageReq)(nil), // 1: gserver.GuildRoutePlayerMessageReq
	(*anypb.Any)(nil),                  // 2: google.protobuf.Any
}
var file_route_proto_depIdxs = []int32{
	2, // 0: gserver.RoutePlayerMessage.packetData:type_name -> google.protobuf.Any
	2, // 1: gserver.GuildRoutePlayerMessageReq.packetData:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_route_proto_init() }
func file_route_proto_init() {
	if File_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutePlayerMessage); i {
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
		file_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuildRoutePlayerMessageReq); i {
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
			RawDescriptor: file_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_route_proto_goTypes,
		DependencyIndexes: file_route_proto_depIdxs,
		MessageInfos:      file_route_proto_msgTypes,
	}.Build()
	File_route_proto = out.File
	file_route_proto_rawDesc = nil
	file_route_proto_goTypes = nil
	file_route_proto_depIdxs = nil
}
