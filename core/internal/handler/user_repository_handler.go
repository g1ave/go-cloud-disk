package handler

import (
	"net/http"

	"github.com/g1ave/go-cloud-disk/core/internal/logic"
	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRepositoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRepoSaveRequeset
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRepositoryLogic(r.Context(), svcCtx)
		resp, err := l.UserRepository(&req, r.Header.Get("userIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
