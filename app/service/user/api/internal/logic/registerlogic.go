package logic

import (
	"context"
	"main/app/service/user/api/internal/svc"
	"main/app/service/user/api/internal/types"
	"main/app/service/user/rpc/crud/crud"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	res, _ := l.svcCtx.CrudRpcClient.Register(l.ctx, &crud.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	return &types.RegisterRes{
		Code: res.Code,
		Msg:  res.Msg,
		Ok:   res.Ok,
	}, nil
}
