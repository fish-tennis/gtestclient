package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/fish-tennis/gnet"
	"github.com/fish-tennis/gtestclient/logger"
	"github.com/fish-tennis/gtestclient/testclient"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	var (
		isDaemon = false
		serverAddr string
		mockClientAccountPrefix string
		mockClientNum int
		mockClientBeginId int
	)
	flag.StringVar(&serverAddr, "server", "127.0.0.1:10002", "testClient's ip:port")
	flag.IntVar(&mockClientNum, "num", 1, "num of mock client")
	flag.IntVar(&mockClientBeginId, "begin", 1, "begin id of mock client")
	flag.StringVar(&mockClientAccountPrefix, "prefix", "mock", "prefix of mock client's accountName")
	flag.BoolVar(&isDaemon, "d", false, "daemon mode")
	flag.Parse()
	logger.Info("server:%v num:%v prefix:%v beginId:%v isDaemon:%v", serverAddr, mockClientNum,
		mockClientAccountPrefix, mockClientBeginId, isDaemon)

	if isDaemon {
		daemon()
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	gnet.SetLogLevel(gnet.DebugLevel)
	rand.Seed(time.Now().UnixNano())

	testClient := new(testclient.TestClient)

	// context实现优雅的协程关闭通知
	ctx,cancel := context.WithCancel(context.Background())
	// 初始化
	if !testClient.Init(ctx, serverAddr, mockClientAccountPrefix, mockClientNum, mockClientBeginId) {
		panic("testClient init error")
	}
	// 运行
	testClient.Run(ctx)

	// 监听系统的kill信号
	signalKillNotify := make(chan os.Signal, 1)
	signal.Notify(signalKillNotify, os.Interrupt, os.Kill, syscall.SIGTERM)
	// 加一个控制台输入,以方便调试
	go func() {
		consoleReader := bufio.NewReader(os.Stdin)
		for {
			lineBytes, _, _ := consoleReader.ReadLine()
			line := strings.ToLower(string(lineBytes))
			logger.Info("line:%v", line)
			if line == "close" || line == "exit" {
				logger.Info("kill by console input")
				// 模拟一个kill信号,以方便退出流程
				signalKillNotify <- os.Kill
				return
			}
			// 机器人程序输入测试命令
			testClient.OnInputCmd(line)
		}
	}()
	// 阻塞等待系统关闭信号
	logger.Info("wait for kill signal")
	select {
	case <-signalKillNotify:
		logger.Info("signalKillNotify, cancel ctx")
		// 通知所有协程关闭,所有监听<-ctx.Done()的地方会收到通知
		cancel()
		break
	}
	// 清理
	testClient.Exit()
}

func daemon() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-d=true" {
			args[i] = "-d=false"
			break
		}
	}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Start()
	fmt.Println("[PID]", cmd.Process.Pid)
	os.Exit(0)
}
