package logic

import (
	"context"
	"github.com/spf13/cast"
	"main/app/service/user/rpc/crud/crud"

	"main/app/service/user/api/internal/svc"
	"main/app/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeNicknameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeNicknameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeNicknameLogic {
	return &ChangeNicknameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeNicknameLogic) ChangeNickname(req *types.ChangeNicknameReq) (resp *types.ChangeNicknameRes, err error) {
	userId := l.ctx.Value("user_id")
	res, _ := l.svcCtx.CrudRpcClient.ChangeNickName(l.ctx, &crud.ChangeNicknameReq{Id: cast.ToInt64(userId), Nickname: req.Nickname})

	return &types.ChangeNicknameRes{
		Code: int(res.Code),
		Msg:  res.Msg,
		Ok:   res.Ok,
	}, nil
}
