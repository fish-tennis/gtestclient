package testclient

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log/slog"
	"os"
	"reflect"
)

var (
	// reflect.TypeOf(*pb.Xxx).Elem() -> packetCommand
	_messageTypeCmdMapping = make(map[reflect.Type]int32)
	_cmdMessageNameMapping = make(map[int32]string)
)

const (
	// 对应proto文件里的package导出名
	ProtoPackageName = "gserver"
)

func GetCommandByProto(protoMessage proto.Message) int32 {
	typ := reflect.TypeOf(protoMessage)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if cmd, ok := _messageTypeCmdMapping[typ]; ok {
		return cmd
	}
	slog.Warn("GetCommandByProtoErr", "messageName", proto.MessageName(protoMessage))
	return 0
}

func InitCommandMappingFromFile(file string) {
	mapping := loadCommandMapping(file)
	for messageName, messageId := range mapping {
		fullMessageName := GetFullMessageName(ProtoPackageName, messageName)
		messageType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(fullMessageName))
		if err != nil {
			slog.Warn("FindMessageByNameErr", "messageName", messageName, "id", messageId, "err", err)
			continue
		}
		if messageInfo, ok := messageType.(*protoimpl.MessageInfo); ok {
			typ := messageInfo.GoReflectType.Elem()
			_messageTypeCmdMapping[typ] = int32(messageId)
			_cmdMessageNameMapping[int32(messageId)] = messageName
			slog.Info("CommandMapping", "messageName", messageName, "id", messageId)
		}
	}
}

func loadCommandMapping(fileName string) map[string]int {
	mapping := make(map[string]int)
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		slog.Error("loadCommandMappingErr", "fileName", fileName, "err", err)
		return mapping
	}
	err = json.Unmarshal(fileData, &mapping)
	if err != nil {
		slog.Error("loadCommandMappingErr", "fileName", fileName, "err", err)
		return mapping
	}
	return mapping
}

func GetFullMessageName(packageName, messageName string) string {
	if ProtoPackageName == "" {
		return messageName
	}
	return ProtoPackageName + "." + messageName
}
