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

func (this *TestClient) Init(ctx context.Context, useGate bool, useWebsocket bool, serverAddr string,
	mockClientAccountPrefix string, mockClientNum int, mockClientBeginId int) bool {
	_testClient = this
	this.ctx = ctx
	this.useGate = useGate
	this.useWebSocket = useWebsocket
	this.serverAddr = serverAddr
	this.mockClientAccountPrefix = mockClientAccountPrefix
	this.mockClientNum = mockClientNum
	this.mockClientBeginId = mockClientBeginId
	this.mockClients = make(map[string]*MockClient)

	if !this.useWebSocket {
		this.clientCodec = NewProtoCodec(nil)
	} else {
		this.clientCodec = NewSimpleProtoCodec()
	}
	this.clientHandler = NewMockClientHandler(this.clientCodec)
	this.clientHandler.autoRegister()
	this.clientHandler.SetOnDisconnectedFunc(func(connection Connection) {
		accountName := connection.GetTag().(string)
		mockClient := _testClient.getMockClientByAccountName(accountName)
		if mockClient == nil {
			return
		}
		_testClient.removeMockClient(accountName)
		logger.Debug("client disconnect %v", accountName)
	})

	for i := 0; i < this.mockClientNum; i++ {
		accountName := fmt.Sprintf("%v%v", this.mockClientAccountPrefix, this.mockClientBeginId+i)
		client := newMockClient(accountName)
		this.addMockClient(client)
		client.start()
	}
	return true
}

func (this *TestClient) Run(ctx context.Context) {
}

func (this *TestClient) OnUpdate(ctx context.Context, updateCount int64) {
}

func (this *TestClient) Exit() {
	this.mockClientsMutex.Lock()
	defer this.mockClientsMutex.Unlock()
	for _, mockClient := range this.mockClients {
		if mockClient.conn != nil {
			mockClient.conn.Close()
		}
	}
}

func (this *TestClient) getMockClientByAccountName(accountName string) *MockClient {
	this.mockClientsMutex.RLock()
	defer this.mockClientsMutex.RUnlock()
	return this.mockClients[accountName]
}

func (this *TestClient) addMockClient(client *MockClient) {
	this.mockClientsMutex.Lock()
	defer this.mockClientsMutex.Unlock()
	this.mockClients[client.accountName] = client
}

func (this *TestClient) removeMockClient(accountName string) {
	this.mockClientsMutex.Lock()
	defer this.mockClientsMutex.Unlock()
	delete(this.mockClients, accountName)
}

func (this *TestClient) OnInputCmd(cmd string) {
	this.mockClientsMutex.RLock()
	defer this.mockClientsMutex.RUnlock()
	for _, mockClient := range this.mockClients {
		mockClient.OnInputCmd(cmd)
	}
}
