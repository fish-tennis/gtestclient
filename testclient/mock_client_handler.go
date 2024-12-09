package testclient

import (
	"fmt"
	. "github.com/fish-tennis/gnet"
	"github.com/fish-tennis/gtestclient/logger"
	"github.com/fish-tennis/gtestclient/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"reflect"
	"strings"
	"time"
)

type MockClientHandler struct {
	DefaultConnectionHandler
	methods map[PacketCommand]reflect.Method
}

func NewMockClientHandler(protoCodec Codec) *MockClientHandler {
	handler := &MockClientHandler{
		DefaultConnectionHandler: *NewDefaultConnectionHandler(protoCodec),
		methods:                  make(map[PacketCommand]reflect.Method),
	}
	handler.RegisterHeartBeat(func() Packet {
		return NewProtoPacket(PacketCommand(pb.CmdClient_Cmd_HeartBeatReq), &pb.HeartBeatReq{
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
		})
	})
	handler.Register(PacketCommand(pb.CmdClient_Cmd_HeartBeatRes), func(connection Connection, packet Packet) {
	}, new(pb.HeartBeatRes))
	handler.SetUnRegisterHandler(func(connection Connection, packet Packet) {
		logger.Debug("un register %v", string(packet.Message().ProtoReflect().Descriptor().Name()))
	})
	return handler
}

func (h *MockClientHandler) OnRecvPacket(connection Connection, packet Packet) {
	if connection.GetTag() != nil {
		accountName := connection.GetTag().(string)
		mockClient := _testClient.getMockClientByAccountName(accountName)
		if mockClient == nil {
			return
		}
		if protoPacket, ok := packet.(*ProtoPacket); ok {
			handlerMethod, ok2 := h.methods[protoPacket.Command()]
			if ok2 {
				if handlerMethod.Type.NumIn() == 2 {
					// func (h *MockClient) OnMessageName(res *pb.MessageName)
					handlerMethod.Func.Call([]reflect.Value{
						reflect.ValueOf(mockClient),
						reflect.ValueOf(protoPacket.Message()),
					})
				} else {
					// func (h *MockClient) OnMessageName(res *pb.MessageName, errorCode int)
					handlerMethod.Func.Call([]reflect.Value{
						reflect.ValueOf(mockClient),
						reflect.ValueOf(protoPacket.Message()),
						reflect.ValueOf(int(protoPacket.ErrorCode())),
					})
				}
				return
			}
		}
		h.DefaultConnectionHandler.OnRecvPacket(connection, packet)
	}
}

// 通过反射自动注册消息回调
func (h *MockClientHandler) autoRegister() {
	typ := reflect.TypeOf(&MockClient{})
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		// 支持2种格式的回调
		// func (h *MockClient) OnMessageName(res *pb.MessageName)
		// func (h *MockClient) OnMessageName(res *pb.MessageName, errorCode int)
		if method.Type.NumIn() != 2 && method.Type.NumIn() != 3 {
			continue
		}
		if !strings.HasPrefix(method.Name, "On") {
			continue
		}
		protoArg := method.Type.In(1)
		if !strings.HasPrefix(protoArg.String(), "*pb.") {
			continue
		}
		if method.Type.NumIn() == 3 {
			// func (h *MockClient) OnMessageName(res *pb.MessageName, errorCode int)
			errorCodeArg := method.Type.In(2)
			if errorCodeArg.Kind() != reflect.Int {
				continue
			}
		}
		// 消息名,如: LoginRes
		messageName := protoArg.String()[strings.LastIndex(protoArg.String(), ".")+1:]
		// 函数名必须是onLoginRes
		if method.Name != fmt.Sprintf("On%v", messageName) {
			logger.Debug("methodName not match:%v", method.Name)
			continue
		}
		// Cmd_LoginRes
		enumValueName := fmt.Sprintf("Cmd_%v", messageName)
		var messageId int32
		protoregistry.GlobalTypes.RangeEnums(func(enumType protoreflect.EnumType) bool {
			// gserver.Login.CmdLogin
			enumValueDescriptor := enumType.Descriptor().Values().ByName(protoreflect.Name(enumValueName))
			if enumValueDescriptor != nil {
				messageId = int32(enumValueDescriptor.Number())
				return false
			}
			return true
		})
		if messageId == 0 {
			continue
		}
		cmd := PacketCommand(messageId)
		// 注册消息回调到组件上
		h.methods[cmd] = method
		// 注册消息的构造函数
		h.DefaultConnectionHandler.Register(cmd, nil, reflect.New(protoArg.Elem()).Interface().(proto.Message))
		logger.Debug("register %v %v hasErrorCode:%v", messageId, method.Name, method.Type.NumIn() == 3)
	}
}
