package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMultipartUploadCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileMultipartUploadCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMultipartUploadCompleteLogic {
	return &FileMultipartUploadCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileMultipartUploadCompleteLogic) FileMultipartUploadComplete(req *types.FileMultipartUploadCompleteRequest) (resp *types.FileMultipartUploadCompleteResponse, err error) {
	u, _ := url.Parse(define.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.svcCtx.Config.COS.SecretId,
			SecretKey: l.svcCtx.Config.COS.SecretKey,
		},
	})
	opt := &cos.CompleteMultipartUploadOptions{}
	for _, part := range req.FileParts {
		opt.Parts = append(opt.Parts, cos.Object{
			ETag:       part.ETag,
			PartNumber: part.PartNumber,
		})
	}
	_, _, err = client.Object.CompleteMultipartUpload(
		context.Background(), req.Key, req.UploadId, opt,
	)
	if err != nil {
		return nil, err
	}
	return
}
