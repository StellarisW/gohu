package svc

import (
	"github.com/go-redis/redis/v8"
	apollo "main/app/common/config"
	"main/app/common/log"
	"main/app/service/comment/dao/query"
	"main/app/service/comment/rpc/info/internal/config"
)

type ServiceContext struct {
	Config config.Config

	CommentModel *query.Query
	Rdb          *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger := log.GetSugaredLogger()

	db, err := apollo.GetMysqlDB("question.yaml")
	if err != nil {
		logger.Fatalf("initialize mysql failed, err: %v", err)
	}

	rdb, err := apollo.GetRedisClient("question.yaml")
	if err != nil {
		logger.Fatalf("initialize redis failed, err: %v", err)
	}

	return &ServiceContext{
		Config: c,

		CommentModel: query.Use(db),
		Rdb:          rdb,
	}
}
