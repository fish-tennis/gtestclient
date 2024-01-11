# gtestclient
test client for [gserver](https://github.com/fish-tennis/gserver)

测试客户端,可以用来测试逻辑或者模拟机器人来进行压力测试

## 命令行参数
-gate: 是否使用网关模式

-server: 服务器地址(gate为true时,填写网关服务器地址,gate为false时,填写登录服务器地址)

-num: 客户端数量

-begin: 客户端起始id

-prefix: 客户端账号名前缀

-d: 是否以守护进程模式运行

## pb
从github.com/fish-tennis/gserver/pb拷贝

## 测试命令
加经验值 addexp 数值

加物品 itemCfgId 物品数量

完成任务 finishquest 任务配置id

模拟一个战斗事件 fight 是否是Pvp(true|false) 是否获胜(true|false)

创建一个公会 guild create

申请加入公会 guild join 公会id
