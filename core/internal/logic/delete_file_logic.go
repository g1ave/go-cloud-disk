package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/models"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileLogic {
	return &DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req *types.DeleteFileRequest, userIdentity string) (resp *types.DeleteFileResponse, err error) {
	res := l.svcCtx.DB.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Delete(&models.UserRepository{})
	err = res.Error
	return
}
