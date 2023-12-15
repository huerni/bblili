package handler

import (
	"net/http"

	"bblili/api/internal/logic"
	"bblili/api/internal/svc"
	"bblili/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSuccessChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSuccessChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetSuccessChunkLogic(r.Context(), svcCtx)
		resp, err := l.GetSuccessChunk(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
