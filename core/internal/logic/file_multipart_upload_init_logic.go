package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/models"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMultipartUploadInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileMultipartUploadInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMultipartUploadInitLogic {
	return &FileMultipartUploadInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileMultipartUploadInitLogic) FileMultipartUploadInit(req *types.FileMultipartUploadInitRequest) (resp *types.FileMultipartUploadInitResponse, err error) {
	rp := new(models.RepositoryPool)
	db := l.svcCtx.DB
	res := db.First(rp, "hash = ?", req.Md5)
	resp = new(types.FileMultipartUploadInitResponse)
	if err = res.Error; err == nil {
		resp.Identity = rp.Identity
	} else {
		key, uploadId, err := l.CosMultipartUploadInit(req.Ext)
		if err != nil {
			return nil, err
		}
		resp.UploadId = uploadId
		resp.Key = key
	}
	return
}

func (l *FileMultipartUploadInitLogic) CosMultipartUploadInit(ext string) (string, string, error) {
	u, _ := url.Parse(define.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.svcCtx.Config.COS.SecretId,
			SecretKey: l.svcCtx.Config.COS.SecretKey,
		},
	})
	key := uuid.NewV4().String() + ext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	return key, v.UploadID, nil
}
