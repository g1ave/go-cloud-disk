package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/models"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFileLogic {
	return &MoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFileLogic) MoveFile(req *types.MoveFileRequest, userIdentity string) (resp *types.MoveFileResponse, err error) {
	parent := new(models.UserRepository)
	db := l.svcCtx.DB
	res := db.First(parent, "identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity)
	if err = res.Error; err != nil {
		return nil, err
	}
	repo := new(models.UserRepository)
	res = db.First(repo, "identity = ?", req.Identity).Update("parent_id", parent.ID)
	err = res.Error
	return
}
