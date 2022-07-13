package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {

	userFiles := make([]*types.UserFile, 0)
	size := req.Size
	if size == 0 {
		size = define.DefaultPageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	db := l.svcCtx.DB

	res := db.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity," +
			"user_repository.ext, + user_repository.name, repository_pool.path, repository_pool.size").
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at IS NULL").
		Limit(size).Offset(offset).Find(&userFiles)

	if err = res.Error; err != nil {
		return nil, err
	}
	resp = new(types.UserFileListResponse)
	resp.Files = userFiles
	resp.Count = len(userFiles)
	return
}
