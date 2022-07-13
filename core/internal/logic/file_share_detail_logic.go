package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/jinzhu/gorm"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileShareDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileShareDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileShareDetailLogic {
	return &FileShareDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileShareDetailLogic) FileShareDetail(req *types.FileShareDetailsRequest) (resp *types.FileShareDetailsResponse, err error) {
	share := new(models.ShareBasic)
	db := l.svcCtx.DB
	res := db.First(share, "identity = ?", req.Identity).Update("click_num", gorm.Expr("click_num + ?", 1))
	if err = res.Error; err != nil {
		return nil, err
	}
	resp = new(types.FileShareDetailsResponse)
	res = db.Table("share_basic").
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Joins("LEFT join repository_pool on share_basic.repository_identity = repository_pool.identity").
		Joins("LEFT join user_repository on user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", req.Identity).Find(resp)
	err = res.Error
	return
}
