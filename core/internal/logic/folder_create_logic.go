package logic

import (
	"context"
	"errors"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FolderCreateLogic {
	return &FolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FolderCreateLogic) FolderCreate(req *types.FileFolderCreateRequest, userIdentity string) (resp *types.FileFolderCreateResponse, err error) {
	folder := new(models.UserRepository)
	db := l.svcCtx.DB
	res := db.Model(models.UserRepository{}).Where("parent_id = ? and name = ?", req.ParentId, req.Name).First(folder)
	if err = res.Error; err == nil {
		return nil, define.NameExistedErr
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	newFolder := &models.UserRepository{
		Identity:     uuid.NewV4().String(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	res = db.Create(newFolder)
	if err = res.Error; err != nil {
		return nil, err
	}
	return
}
