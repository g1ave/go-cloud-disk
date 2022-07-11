package handler

import (
	"net/http"

	"github.com/g1ave/go-cloud-disk/core/internal/logic"
	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileNameUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileNameUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileNameUpdateLogic(r.Context(), svcCtx)
		resp, err := l.FileNameUpdate(&req, r.Header.Get("userIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
