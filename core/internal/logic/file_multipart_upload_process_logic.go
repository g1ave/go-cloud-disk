package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMultipartUploadProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileMultipartUploadProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMultipartUploadProcessLogic {
	return &FileMultipartUploadProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileMultipartUploadProcessLogic) FileMultipartUploadProcess(req *types.FileMultipartUploadProcessRequest, key, uploadId string, partNumber int, file io.Reader) (resp *types.FileMultipartUploadProcessResponse, err error) {
	u, _ := url.Parse(define.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.svcCtx.Config.COS.SecretId,
			SecretKey: l.svcCtx.Config.COS.SecretKey,
		},
	})
	res, err := client.Object.UploadPart(
		context.Background(), key, uploadId, 1, file, nil,
	)
	if err != nil {
		return nil, err
	}
	partETag := res.Header.Get("ETag")
	resp.ETag = partETag
	return
}
