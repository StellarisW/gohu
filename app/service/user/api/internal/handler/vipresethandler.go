package handler

import (
	"main/app/common/log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"main/app/service/user/api/internal/logic"
	"main/app/service/user/api/internal/svc"
	"main/app/service/user/api/internal/types"
)

func VipResetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VipResetReq
		logger := log.GetSugaredLogger()

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			logger.Errorf("Parse http args failed, \nrequest: \n%v, \nerr: \n%v", r, err)
			return
		}
		logger.Debugf("recv args: %v", req)

		l := logic.NewVipResetLogic(r.Context(), svcCtx)

		res, err := l.VipReset(&req)
		if err != nil {
			logger.Errorf("Process logic failed, err: %v", err)
		}

		logger.Info("response: %v", res)
		httpx.WriteJson(w, int(res.Code), res)
	}
}
