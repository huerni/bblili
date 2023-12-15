package handler

import (
	"net/http"

	"bblili/api/internal/logic"
	"bblili/api/internal/svc"
	"bblili/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFileChunksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileChunksRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadFileChunksLogic(r.Context(), svcCtx)
		resp, err := l.UploadFileChunks(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
