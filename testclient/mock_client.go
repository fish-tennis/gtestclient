package testclient

import (
	"crypto/md5"
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

func (c *MockClient) getConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{
		SendPacketCacheCap: 256,
		SendBufferSize:     1024 * 10,
		RecvBufferSize:     1024 * 10,
		MaxPacketSize:      1024 * 10,
		RecvTimeout:        20,
		HeartBeatInterval:  5,
		WriteTimeout:       10,
		Codec:              _testClient.clientCodec,
		Handler:            _testClient.clientHandler,
		Path:               "/ws", // WebSocket才需要
		Scheme:             "ws",  // WebSocket才需要
	}
}

func (c *MockClient) connectServer(serverAddr string) bool {
	if !_testClient.useWebSocket {
		c.conn = GetNetMgr().NewConnector(_testClient.ctx, serverAddr, c.getConnectionConfig(),
			c.accountName)
	} else {
		c.conn = GetNetMgr().NewWsConnector(_testClient.ctx, serverAddr, c.getConnectionConfig(),
			c.accountName)
	}
	return c.conn != nil
}

// 账号协议不要发明文密码,而且要加一些混淆词,防止被"撞库"
func GetMd5Password(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password+"gserver")))
}

func (c *MockClient) start() {
	go func() {
		//defer func() {
		//	_testClient.removeMockClient(c.accountName)
		//	if err := recover(); err != nil {
		//		logger.Error(err.(error).Error())
		//	}
		//}()

		c.connectServer(_testClient.serverAddr)
		if c.conn == nil {
			_testClient.removeMockClient(c.accountName)
			return
		}
		c.Send(&pb.LoginReq{
			AccountName: c.accountName,
			Password:    GetMd5Password(c.accountName),
		})
		logger.Debug("LoginReq %v", c.accountName)
	}()
}

func (c *MockClient) Send(message proto.Message, opts ...SendOption) bool {
	clientCmd := GetCommandByProto(message)
	if clientCmd <= 0 {
		logger.Error("clientCmdNotFound messageName:%v", proto.MessageName(message))
		return false
	}
	if c.conn.Send(PacketCommand(clientCmd), message) {
		logger.Debug("Send %v messageName:%v", clientCmd, proto.MessageName(message))
		return true
	} else {
		logger.Error("SendError messageName:%v", proto.MessageName(message))
		return false
	}
}

func (c *MockClient) OnLoginRes(res *pb.LoginRes, errorCode int) {
	logger.Debug("onLoginRes:%v err:%v", res, errorCode)
	if errorCode == int(pb.ErrorCode_ErrorCode_NotReg) {
		// 自动注册账号
		// 这里是单纯的测试,账号和密码直接使用明文,实际项目需要做md5之类的处理
		c.Send(&pb.AccountReg{
			AccountName: c.accountName,
			Password:    c.accountName,
		})
	} else if errorCode == 0 {
		c.loginRes = res
		if !_testClient.useGate {
			// 客户端直连模式,需要断开和LoginServer的连接,并连接分配的GameServer
			c.conn.SetTag("")
			c.conn.Close()
			// 账号登录成功后,连接游戏服
			c.connectServer(res.GetGameServer().GetClientListenAddr())
			if c.conn == nil {
				logger.Error("%v connect game failed", c.accountName)
				_testClient.removeMockClient(c.accountName)
				return
			}
		}
		c.Send(&pb.PlayerEntryGameReq{
			AccountId:    c.loginRes.GetAccountId(),
			LoginSession: c.loginRes.GetLoginSession(),
			RegionId:     1,
		})
	} else {
		c.conn.Close()
	}
}

func (c *MockClient) OnAccountRes(res *pb.AccountRes, errorCode int) {
	logger.Debug("onAccountRes:%v err:%v", res, errorCode)
	if errorCode == 0 {
		c.Send(&pb.LoginReq{
			AccountName: c.accountName,
			Password:    GetMd5Password(c.accountName),
		})
	}
}

func (c *MockClient) OnPlayerEntryGameRes(res *pb.PlayerEntryGameRes, errorCode int) {
	logger.Debug("onPlayerEntryGameRes:%v", res)
	if errorCode == 0 {
		c.playerEntryGameRes = res
		// 玩家登录游戏服成功,模拟一个交互消息
		c.Send(&pb.GuildListReq{
			PageIndex: 0,
		})
		return
	}
	// 还没角色,则创建新角色
	if errorCode == int(pb.ErrorCode_ErrorCode_NoPlayer) {
		c.Send(&pb.CreatePlayerReq{
			AccountId:    c.loginRes.GetAccountId(),
			LoginSession: c.loginRes.GetLoginSession(),
			RegionId:     1,
			Name:         c.accountName,
			Gender:       1,
		})
		return
	}
	// 登录遇到问题,服务器提示客户端稍后重试
	if errorCode == int(pb.ErrorCode_ErrorCode_TryLater) {
		// 延迟重试
		time.AfterFunc(time.Second, func() {
			c.Send(&pb.PlayerEntryGameReq{
				AccountId:    c.loginRes.GetAccountId(),
				LoginSession: c.loginRes.GetLoginSession(),
				RegionId:     1,
			})
		})
	}
}

func (c *MockClient) OnCreatePlayerRes(res *pb.CreatePlayerRes, errorCode int) {
	logger.Debug("onCreatePlayerRes:%v", res)
	if errorCode == 0 {
		c.Send(&pb.PlayerEntryGameReq{
			AccountId:    c.loginRes.GetAccountId(),
			LoginSession: c.loginRes.GetLoginSession(),
			RegionId:     1,
		})
	}
}

func (c *MockClient) OnErrorRes(res *pb.ErrorRes) {
	logger.Debug("OnErrorRes cmd:%v id:%v str:%v", res.Command, res.ResultId, res.ResultStr)
}

func (c *MockClient) OnGateRouteClientPacketError(res *pb.GateRouteClientPacketError) {
	logger.Debug("OnGateRouteClientPacketError id:%v str:%v", res.ResultId, res.ResultStr)
	c.conn.Close()
}

func (c *MockClient) OnBaseInfoSync(res *pb.BaseInfoSync) {
	logger.Debug("OnBaseInfoSync:%v", res)
}

func (c *MockClient) OnMoneySync(res *pb.MoneySync) {
	logger.Debug("OnMoneySync:%v", res)
}

func (c *MockClient) OnBagsSync(res *pb.BagsSync) {
	logger.Debug("OnBagsSync:%v", res)
}

func (c *MockClient) OnElemContainerUpdate(res *pb.ElemContainerUpdate) {
	logger.Debug("OnElemContainerUpdate:%v", res)
}

func (c *MockClient) OnQuestSync(res *pb.QuestSync) {
	logger.Debug("OnQuestSync:%v", res)
}

func (c *MockClient) OnQuestUpdate(res *pb.QuestUpdate) {
	logger.Debug("OnQuestUpdate:%v", res)
}

func (c *MockClient) OnFinishQuestRes(res *pb.FinishQuestRes) {
	logger.Debug("OnFinishQuestRes:%v", res)
}

func (c *MockClient) OnGuildSync(res *pb.GuildSync) {
	logger.Debug("OnGuildSync:%v", res)
	if res.Data.GuildId > 0 {
		// 已有公会 获取公会数据
		c.Send(&pb.GuildDataViewReq{})
	}
}

func (c *MockClient) OnGuildCreateRes(res *pb.GuildCreateRes) {
	logger.Debug("OnGuildCreateRes:%v", res)
}

func (c *MockClient) OnGuildListRes(res *pb.GuildListRes) {
	logger.Debug("OnGuildListRes:%v", res)
	//if len(res.GuildInfos) > 0 {
	//	if c.playerEntryGameRes.GuildData.GuildId == 0 {
	//		// 申请加入公会
	//		c.conn.Send(PacketCommand(pb.CmdGuild_Cmd_GuildJoinReq), &pb.GuildJoinReq{
	//			Id: res.GuildInfos[0].Id,
	//		})
	//		logger.Debug("GuildJoinReq:%v", res.GuildInfos[0].Id)
	//	}
	//} else {
	//	// 没有公会 就创建一个
	//	c.conn.Send(PacketCommand(pb.CmdGuild_Cmd_GuildCreateReq), &pb.GuildCreateReq{
	//		Name: fmt.Sprintf("%v's guild", c.accountName),
	//		Intro: fmt.Sprintf("create by %v", c.accountName),
	//	})
	//}
}

func (c *MockClient) OnGuildJoinRes(res *pb.GuildJoinRes) {
	logger.Debug("OnGuildJoinRes:%v", res)
}

func (c *MockClient) OnGuildJoinReqTip(res *pb.GuildJoinReqTip) {
	logger.Debug("OnGuildJoinReqTip:%v", res)
	// 同意入会请求
	c.Send(&pb.GuildJoinAgreeReq{
		JoinPlayerId: res.PlayerId,
		IsAgree:      true,
	})
}

func (c *MockClient) OnGuildDataViewRes(res *pb.GuildDataViewRes) {
	logger.Debug("OnGuildDataViewRes:%v", res)
	for _, v := range res.GuildData.JoinRequests {
		// 同意入会请求
		c.Send(&pb.GuildJoinAgreeReq{
			JoinPlayerId: v.PlayerId,
			IsAgree:      true,
		})
	}
}

func (c *MockClient) OnGuildJoinAgreeRes(res *pb.GuildJoinAgreeRes) {
	logger.Debug("OnGuildJoinAgreeRes:%v", res)
}

func (c *MockClient) OnGuildJoinReqOpResult(res *pb.GuildJoinReqOpResult) {
	logger.Debug("OnGuildJoinReqOpResult:%v", res)
}

func (c *MockClient) OnActivitySync(res *pb.ActivitySync) {
	logger.Debug("OnActivitySync:%v", res)
}

func (c *MockClient) OnExchangeSync(res *pb.ExchangeSync) {
	logger.Debug("OnExchangeSync:%v", res)
}

func (c *MockClient) OnExchangeUpdate(res *pb.ExchangeUpdate) {
	logger.Debug("OnExchangeUpdate:%v", res)
}

func (c *MockClient) OnExchangeRes(res *pb.ExchangeRes) {
	logger.Debug("OnExchangeRes:%v", res)
}

// 测试命令
// 支持的命令,详见https://github.com/fish-tennis/gserver/gameserver/test_cmd.go
func (c *MockClient) OnInputCmd(cmd string) {
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
			c.Send(&pb.GuildCreateReq{
				Name:  fmt.Sprintf("%v's guild", c.accountName),
				Intro: fmt.Sprintf("create by %v", c.accountName),
			})
			return
		} else if cmdArgs[0] == "join" {
			if len(cmdArgs) != 2 {
				logger.Error("usage: guild join guildId")
				return
			}
			guildId, _ := strconv.ParseInt(cmdArgs[1], 10, 64)
			// 申请加入公会
			c.Send(&pb.GuildJoinReq{
				Id: guildId,
			})
			return
		}
	}
	// 发给服务器的测试命令
	c.Send(&pb.TestCmd{
		Cmd: cmd,
	})
}
