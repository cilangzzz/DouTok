package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/comment/handler"
	"github.com/TremblingV5/DouTok/applications/comment/rpc"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/kitex_gen/comment/commentservice"
	"github.com/TremblingV5/DouTok/pkg/DouTokContext"
	"github.com/TremblingV5/DouTok/pkg/DouTokLogger"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/services"
	"go.uber.org/zap"
)

type Config struct {
	configStruct.BaseConfig `envPrefix:"DOUTOK_COMMENT_"`
	Jwt                     configStruct.Jwt    `envPrefix:"DOUTOK_COMMENT_"`
	Logger                  configStruct.Logger `envPrefix:"DOUTOK_COMMENT_"`
}

var (
	logger *zap.Logger
	config = &Config{}
)

func init() {

	ctx := context.Background()
	_, err := configurator.Load(config, "DOUTOK_COMMENT", "comment")
	logger = DouTokLogger.InitLogger(config.Logger)
	DouTokContext.DefaultLogger = logger
	DouTokContext.AddLoggerToContext(ctx, logger)
	if err != nil {
		logger.Fatal("could not load env variables", zap.Error(err), zap.Any("config", config))
	}
}

func main() {
	options, shutdown := services.InitRPCServerArgs(constants.COMMENT_SERVER_NAME, config.BaseConfig)
	defer shutdown()

	svr := commentservice.NewServer(
		handler.New(rpc.New(services.InitRPCClientArgs(constants.COMMENT_SERVER_NAME, config.Etcd))),
		options...,
	)

	if err := svr.Run(); err != nil {
		logger.Fatal("run server err", zap.Error(err))
	}
}
