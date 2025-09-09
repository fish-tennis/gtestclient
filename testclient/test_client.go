package testclient

import (
	"context"
	"fmt"
	. "github.com/fish-tennis/gnet"
	"github.com/fish-tennis/gtestclient/logger"
	"sync"
)

var (
	_testClient *TestClient
)

// 客户端测试
// 可以管理多个模拟的客户端
type TestClient struct {
	ctx context.Context
	wg  sync.WaitGroup

	clientCodec   Codec
	clientHandler *MockClientHandler

	useGate                 bool // 是否使用网关模式
	useWebSocket            bool // 是否使用WebSocket
	serverAddr              string
	mockClientAccountPrefix string
	mockClientNum           int
	mockClientBeginId       int

	mockClients      map[string]*MockClient
	mockClientsMutex sync.RWMutex
}

func (t *TestClient) Init(ctx context.Context, useGate bool, useWebsocket bool, serverAddr string,
	mockClientAccountPrefix string, mockClientNum int, mockClientBeginId int) bool {
	_testClient = t
	t.ctx = ctx
	t.useGate = useGate
	t.useWebSocket = useWebsocket
	t.serverAddr = serverAddr
	t.mockClientAccountPrefix = mockClientAccountPrefix
	t.mockClientNum = mockClientNum
	t.mockClientBeginId = mockClientBeginId
	t.mockClients = make(map[string]*MockClient)

	InitCommandMappingFromFile("cfgdata/message_command_mapping.json")
	if !t.useWebSocket {
		t.clientCodec = NewProtoCodec(nil)
	} else {
		t.clientCodec = NewSimpleProtoCodec()
	}
	t.clientHandler = NewMockClientHandler(t.clientCodec)
	t.clientHandler.autoRegister()
	t.clientHandler.SetOnDisconnectedFunc(func(connection Connection) {
		accountName := connection.GetTag().(string)
		mockClient := _testClient.getMockClientByAccountName(accountName)
		if mockClient == nil {
			return
		}
		_testClient.removeMockClient(accountName)
		logger.Debug("client disconnect %v", accountName)
	})

	for i := 0; i < t.mockClientNum; i++ {
		accountName := fmt.Sprintf("%v%v", t.mockClientAccountPrefix, t.mockClientBeginId+i)
		client := newMockClient(accountName)
		t.addMockClient(client)
		client.start()
	}
	return true
}

func (t *TestClient) Run(ctx context.Context) {
}

func (t *TestClient) OnUpdate(ctx context.Context, updateCount int64) {
}

func (t *TestClient) Exit() {
	t.mockClientsMutex.Lock()
	defer t.mockClientsMutex.Unlock()
	for _, mockClient := range t.mockClients {
		if mockClient.conn != nil {
			mockClient.conn.Close()
		}
	}
}

func (t *TestClient) getMockClientByAccountName(accountName string) *MockClient {
	t.mockClientsMutex.RLock()
	defer t.mockClientsMutex.RUnlock()
	return t.mockClients[accountName]
}

func (t *TestClient) addMockClient(client *MockClient) {
	t.mockClientsMutex.Lock()
	defer t.mockClientsMutex.Unlock()
	t.mockClients[client.accountName] = client
}

func (t *TestClient) removeMockClient(accountName string) {
	t.mockClientsMutex.Lock()
	defer t.mockClientsMutex.Unlock()
	delete(t.mockClients, accountName)
}

func (t *TestClient) OnInputCmd(cmd string) {
	t.mockClientsMutex.RLock()
	defer t.mockClientsMutex.RUnlock()
	for _, mockClient := range t.mockClients {
		mockClient.OnInputCmd(cmd)
	}
}
