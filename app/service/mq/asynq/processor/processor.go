package main

import (
	"context"
	apollo "main/app/common/config"
	"main/app/common/log"
	"main/app/service/mq/asynq/processor/internal/config"
	"main/app/service/mq/asynq/processor/internal/logic"
	"main/app/service/mq/asynq/processor/internal/svc"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

const mqName = "mq.asynq.processor"

var c config.Config

func main() {
	// 初始化日志管理器
	err := log.InitLogger()
	if err != nil {
		panic("initialize logger failed")
	}
	logger := log.GetSugaredLogger()

	logWriter, err := log.GetZapWriter()
	if err != nil {
		logger.Fatalf("get log writer failed")
	}
	logx.MustSetup(log.GetLogXConfig("mq-asynq-scheduler", "info"))
	logx.SetWriter(logWriter)

	// 初始化配置管理器
	configClient, err := apollo.NewConfigClient()
	if err != nil {
		logger.Fatalf("Initialize Apollo Client failed, err: %v", err)
	}

	// 初始化消息队列设置
	err = configClient.UnmarshalKey("mq.yaml", "asynq.processor", &c)
	if err != nil {
		logger.Fatalf("UnmarshalKey into service config failed, err: %v", err)
	}

	// 初始化log、trace
	err = c.SetUp()
	if err != nil {
		logger.Fatalf("initialize go-zero internal service failed, err: %v", err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	asynqProcessor := logic.NewProcessor(ctx, svcContext)
	mux := asynqProcessor.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logger.Fatalf("run asynq processor failed, err: %v", err)
		os.Exit(1)
	}
}
