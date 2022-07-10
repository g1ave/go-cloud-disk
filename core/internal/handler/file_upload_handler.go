package handler

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"path"

	"github.com/g1ave/go-cloud-disk/core/internal/logic"
	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
		}

		// determine if the file existed
		buf := make([]byte, fileHeader.Size)
		_, err = file.Read(buf)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(buf))
		rp := new(models.RepositoryPool)
		res := svcCtx.DB.First(&rp, "hash = ?", hash)

		if err = res.Error; err == nil {
			httpx.OkJson(w, &types.FileUploadReply{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			httpx.Error(w, err)
			return
		}

		// upload file to Cos
		u, _ := url.Parse(define.BucketURL)
		b := &cos.BaseURL{BucketURL: u}
		client := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  svcCtx.Config.COS.SecretId,
				SecretKey: svcCtx.Config.COS.SecretKey,
			},
		})
		fileExt := path.Ext(fileHeader.Filename)
		key := uuid.NewV4().String() + fileExt
		file, _, _ = r.FormFile("file")
		_, err = client.Object.Put(
			context.Background(), key, file, nil,
		)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		req.Size = fileHeader.Size
		req.Ext = fileExt
		req.Hash = hash
		req.Path = define.BucketURL + "/" + key
		req.Name = fileHeader.Filename

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
