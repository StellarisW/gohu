package logic

import (
	"context"
	"main/app/common/log"
	"main/app/service/oauth/api/internal/svc"
	"main/app/service/oauth/api/internal/token"
	"main/app/service/oauth/api/internal/types"
	"main/app/utils/mapping"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenLogic) GetToken(req *types.GetTokenByAuthReq) (res *types.GetTokenByAuthRes, err error) {
	logger := log.GetSugaredLogger()

	tokenGranter := token.GetTokenGranter()
	accessToken, err := tokenGranter.Grant(l.ctx, token.GrantByAuth, req.Authorization)
	if err != nil {
		return &types.GetTokenByAuthRes{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Ok:   false,
		}, nil
	}

	res = &types.GetTokenByAuthRes{
		Code: http.StatusOK,
		Msg:  "get token successfully",
		Ok:   true,
		Data: types.GetTokenByAuthResData{AccessToken: &types.OAuth2Token{}},
	}
	err = mapping.Struct2Struct(accessToken, res.Data.AccessToken)
	if err != nil {
		logger.Errorf("mapping struct failed, err: %v", err)
		return &types.GetTokenByAuthRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
			Ok:   false,
		}, nil
	}
	return res, nil
}
