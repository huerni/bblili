package handler

import (
	"net/http"

	"bblili/api/internal/logic"
	"bblili/api/internal/svc"
	"bblili/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetVideosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetVideosRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetVideosLogic(r.Context(), svcCtx)
		resp, err := l.GetVideos(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
