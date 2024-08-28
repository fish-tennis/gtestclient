package testclient

import (
	"fmt"
	. "github.com/fish-tennis/gnet"
	"github.com/fish-tennis/gtestclient/logger"
	"github.com/fish-tennis/gtestclient/pb"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
	"time"
)

// 模拟客户端
type MockClient struct {
	accountName string
	conn        Connection

	loginRes           *pb.LoginRes
	playerEntryGameRes *pb.PlayerEntryGameRes
}

func newMockClient(accountName string) *MockClient {
	return &MockClient{
		accountName: accountName,
	}
}

func (this *MockClient) getConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{
		SendPacketCacheCap: 16,
		SendBufferSize:     1024 * 10,
		RecvBufferSize:     1024 * 10,
		MaxPacketSize:      1024 * 10,
		RecvTimeout:        0,
		HeartBeatInterval:  5,
		WriteTimeout:       0,
		Codec:              _testClient.clientCodec,
		Handler:            _testClient.clientHandler,
		Path:               "/ws", // WebSocket才需要
		Scheme:             "ws",  // WebSocket才需要
	}
}

func (this *MockClient) connectServer(serverAddr string) bool {
	if !_testClient.useWebSocket {
		this.conn = GetNetMgr().NewConnector(_testClient.ctx, serverAddr, this.getConnectionConfig(),
			this.accountName)
	} else {
		this.conn = GetNetMgr().NewWsConnector(_testClient.ctx, serverAddr, this.getConnectionConfig(),
			this.accountName)
	}
	return this.conn != nil
}

func (this *MockClient) start() {
	go func() {
		//defer func() {
		//	_testClient.removeMockClient(this.accountName)
		//	if err := recover(); err != nil {
		//		logger.Error(err.(error).Error())
		//	}
		//}()

		this.connectServer(_testClient.serverAddr)
		if this.conn == nil {
			_testClient.removeMockClient(this.accountName)
			return
		}
		this.Send(&pb.LoginReq{
			AccountName: this.accountName,
			Password:    this.accountName,
		})
		logger.Debug("LoginReq %v", this.accountName)
	}()
}

func (this *MockClient) Send(message proto.Message, opts ...SendOption) bool {
	clientCmd := GetClientCommandByProto(message)
	if clientCmd <= 0 {
		logger.Error("clientCmdNotFound messageName:%v", proto.MessageName(message))
		return false
	}
	if this.conn.Send(PacketCommand(clientCmd), message) {
		logger.Debug("Send %v messageName:%v", clientCmd, proto.MessageName(message))
		return true
	} else {
		logger.Error("SendError messageName:%v", proto.MessageName(message))
		return false
	}
}

func (this *MockClient) OnLoginRes(res *pb.LoginRes) {
	logger.Debug("onLoginRes:%v", res)
	if res.Error == "NotReg" {
		// 自动注册账号
		// 这里是单纯的测试,账号和密码直接使用明文,实际项目需要做md5之类的处理
		this.Send(&pb.AccountReg{
			AccountName: this.accountName,
			Password:    this.accountName,
		})
	} else if res.Error == "" {
		this.loginRes = res
		if !_testClient.useGate {
			// 客户端直连模式,需要断开和LoginServer的连接,并连接分配的GameServer
			this.conn.SetTag("")
			this.conn.Close()
			// 账号登录成功后,连接游戏服
			this.connectServer(res.GetGameServer().GetClientListenAddr())
			if this.conn == nil {
				logger.Error("%v connect game failed", this.accountName)
				_testClient.removeMockClient(this.accountName)
				return
			}
		}
		this.Send(&pb.PlayerEntryGameReq{
			AccountId:    this.loginRes.GetAccountId(),
			LoginSession: this.loginRes.GetLoginSession(),
			RegionId:     1,
		})
	} else {
		this.conn.Close()
	}
}

func (this *MockClient) OnAccountRes(res *pb.AccountRes) {
	logger.Debug("onAccountRes:%v", res)
	if res.Error == "" {
		this.Send(&pb.LoginReq{
			AccountName: this.accountName,
			Password:    this.accountName,
		})
	}
}

func (this *MockClient) OnPlayerEntryGameRes(res *pb.PlayerEntryGameRes) {
	logger.Debug("onPlayerEntryGameRes:%v", res)
	if res.Error == "" {
		this.playerEntryGameRes = res
		// 玩家登录游戏服成功,模拟一个交互消息
		this.Send(&pb.CoinReq{
			AddCoin: 1,
		})
		this.Send(&pb.GuildListReq{
			PageIndex: 0,
		})
		if res.GuildData.GuildId > 0 {
			// 已有公会 获取公会数据
			this.Send(&pb.GuildDataViewReq{})
		}
		return
	}
	// 还没角色,则创建新角色
	if res.Error == "NoPlayer" {
		this.Send(&pb.CreatePlayerReq{
			AccountId:    this.loginRes.GetAccountId(),
			LoginSession: this.loginRes.GetLoginSession(),
			RegionId:     1,
			Name:         this.accountName,
			Gender:       1,
		})
		return
	}
	// 登录遇到问题,服务器提示客户端稍后重试
	if res.Error == "TryLater" {
		// 延迟重试
		time.AfterFunc(time.Second, func() {
			this.Send(&pb.PlayerEntryGameReq{
				AccountId:    this.loginRes.GetAccountId(),
				LoginSession: this.loginRes.GetLoginSession(),
				RegionId:     1,
			})
		})
	}
}

func (this *MockClient) OnCreatePlayerRes(res *pb.CreatePlayerRes) {
	logger.Debug("onCreatePlayerRes:%v", res)
	if res.Error == "" {
		this.Send(&pb.PlayerEntryGameReq{
			AccountId:    this.loginRes.GetAccountId(),
			LoginSession: this.loginRes.GetLoginSession(),
			RegionId:     1,
		})
	}
}

func (this *MockClient) OnErrorRes(res *pb.ErrorRes) {
	logger.Debug("OnErrorRes cmd:%v id:%v str:%v", res.Command, res.ResultId, res.ResultStr)
}

func (this *MockClient) OnCoinRes(res *pb.CoinRes) {
	logger.Debug("OnCoinRes:%v", res)
}

func (this *MockClient) OnFinishQuestRes(res *pb.FinishQuestRes) {
	logger.Debug("OnFinishQuestRes:%v", res)
}

func (this *MockClient) OnGuildCreateRes(res *pb.GuildCreateRes) {
	logger.Debug("OnGuildCreateRes:%v", res)
}

func (this *MockClient) OnGuildListRes(res *pb.GuildListRes) {
	logger.Debug("OnGuildListRes:%v", res)
	//if len(res.GuildInfos) > 0 {
	//	if this.playerEntryGameRes.GuildData.GuildId == 0 {
	//		// 申请加入公会
	//		this.conn.Send(PacketCommand(pb.CmdGuild_Cmd_GuildJoinReq), &pb.GuildJoinReq{
	//			Id: res.GuildInfos[0].Id,
	//		})
	//		logger.Debug("GuildJoinReq:%v", res.GuildInfos[0].Id)
	//	}
	//} else {
	//	// 没有公会 就创建一个
	//	this.conn.Send(PacketCommand(pb.CmdGuild_Cmd_GuildCreateReq), &pb.GuildCreateReq{
	//		Name: fmt.Sprintf("%v's guild", this.accountName),
	//		Intro: fmt.Sprintf("create by %v", this.accountName),
	//	})
	//}
}

func (this *MockClient) OnGuildJoinRes(res *pb.GuildJoinRes) {
	logger.Debug("OnGuildJoinRes:%v", res)
}

func (this *MockClient) OnGuildJoinReqTip(res *pb.GuildJoinReqTip) {
	logger.Debug("OnGuildJoinReqTip:%v", res)
	// 同意入会请求
	this.Send(&pb.GuildJoinAgreeReq{
		JoinPlayerId: res.PlayerId,
		IsAgree:      true,
	})
}

func (this *MockClient) OnGuildDataViewRes(res *pb.GuildDataViewRes) {
	logger.Debug("OnGuildDataViewRes:%v", res)
	for _, v := range res.GuildData.JoinRequests {
		// 同意入会请求
		this.Send(&pb.GuildJoinAgreeReq{
			JoinPlayerId: v.PlayerId,
			IsAgree:      true,
		})
	}
}

func (this *MockClient) OnGuildJoinAgreeRes(res *pb.GuildJoinAgreeRes) {
	logger.Debug("OnGuildJoinAgreeRes:%v", res)
}

func (this *MockClient) OnGuildJoinReqOpResult(res *pb.GuildJoinReqOpResult) {
	logger.Debug("OnGuildJoinReqOpResult:%v", res)
}

// 测试命令
// 支持的命令,详见https://github.com/fish-tennis/gserver/gameserver/test_cmd.go
func (this *MockClient) OnInputCmd(cmd string) {
	// 本地测试命令
	cmdStrs := strings.Split(cmd, " ")
	if len(cmdStrs) == 0 {
		return
	}
	cmdKey := strings.ToLower(cmdStrs[0])
	cmdArgs := cmdStrs[1:]
	if cmdKey == "guild" && len(cmdArgs) >= 1 {
		if cmdArgs[0] == "create" {
			// 创建公会
			this.Send(&pb.GuildCreateReq{
				Name:  fmt.Sprintf("%v's guild", this.accountName),
				Intro: fmt.Sprintf("create by %v", this.accountName),
			})
			return
		} else if cmdArgs[0] == "join" {
			if len(cmdArgs) != 2 {
				logger.Error("usage: guild join guildId")
				return
			}
			guildId, _ := strconv.ParseInt(cmdArgs[1], 10, 64)
			// 申请加入公会
			this.Send(&pb.GuildJoinReq{
				Id: guildId,
			})
			return
		}
	}
	// 发给服务器的测试命令
	this.Send(&pb.TestCmd{
		Cmd: cmd,
	})
}
