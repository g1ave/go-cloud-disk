package logic

import (
	"context"
	"errors"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/jinzhu/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileNameUpdateLogic {
	return &FileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileNameUpdateLogic) FileNameUpdate(req *types.FileNameUpdateRequest, userIdentity string) (resp *types.FileNameUpdateResponse, err error) {
	// determine if the new name existed
	db := l.svcCtx.DB
	ur := new(models.UserRepository)
	res := db.Where("name = ? and parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).First(ur)
	if err = res.Error; err == nil {
		return nil, define.NameExistedErr
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	// update filename
	res = db.Model(models.UserRepository{}).
		Where("repository_identity = ? AND user_identity = ?", req.Identity, userIdentity).Update("name", req.Name)
	if err = res.Error; err != nil {
		return nil, err
	}
	return
}
