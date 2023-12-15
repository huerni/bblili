package handler

import (
	"net/http"

	"bblili/api/internal/logic"
	"bblili/api/internal/svc"
	"bblili/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OperatorUserSanLianHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OperationSanLianRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOperatorUserSanLianLogic(r.Context(), svcCtx)
		resp, err := l.OperatorUserSanLian(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
