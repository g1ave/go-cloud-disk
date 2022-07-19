package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/g1ave/go-cloud-disk/core/internal/logic"
	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileMultipartUploadProcessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileMultipartUploadProcessRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		key := r.PostForm.Get("key")
		if key == "" {
			httpx.Error(w, errors.New("key is empty"))
			return
		}
		uploadId := r.PostForm.Get("upload_id")
		if uploadId == "" {
			httpx.Error(w, errors.New("upload_id is empty"))
			return
		}
		partNumber, err := strconv.Atoi(r.PostForm.Get("part_number"))
		if err != nil {
			httpx.Error(w, errors.New("part_number needs integer"))
			return
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewFileMultipartUploadProcessLogic(r.Context(), svcCtx)
		resp, err := l.FileMultipartUploadProcess(&req, key, uploadId, partNumber, file)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
