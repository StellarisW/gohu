package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	apollo "main/app/common/config"
	"main/app/common/log"
	"main/app/common/middleware"
	"main/app/service/question/api/internal/config"
	"main/app/service/question/rpc/crud/crud"
	"main/app/service/question/rpc/info/info"
)

type ServiceContext struct {
	Config config.Config

	AuthMiddleware rest.Middleware

	CrudRpcClient crud.Crud
	InfoRpcClient info.Info
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger := log.GetSugaredLogger()

	rdb, err := apollo.GetRedisClient("user.yaml")
	if err != nil {
		logger.Fatalf("initialize redis failed, err: %v", err)
	}

	cookieConfig, err := apollo.NewCookieConfig()
	if err != nil {
		logger.Fatalf("get cookieConfig failed, err: %v", err)
	}

	domain, err := apollo.GetDomain()
	if err != nil {
		logger.Fatalf("get domain failed, err: %v", err)
	}

	return &ServiceContext{
		Config: c,

		AuthMiddleware: middleware.NewAuthMiddleware(domain, cookieConfig, rdb).Handle,

		CrudRpcClient: crud.NewCrud(zrpc.MustNewClient(c.CrudRpcClientConf)),
		InfoRpcClient: info.NewInfo(zrpc.MustNewClient(c.InfoRpcClientConf)),
	}
}
