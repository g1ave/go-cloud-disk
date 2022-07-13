package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/models"
	uuid "github.com/satori/go.uuid"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileShareCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileShareCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileShareCreateLogic {
	return &FileShareCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileShareCreateLogic) FileShareCreate(req *types.FileShareCreateRequest, userIdentity string) (resp *types.FileShareCreateResponse, err error) {
	db := l.svcCtx.DB
	repo := new(models.UserRepository)
	res := db.First(repo, "identity = ? AND user_identity = ?", req.UserRepoIdentity, userIdentity)
	if err = res.Error; err != nil {
		return nil, err
	}
	share := &models.ShareBasic{
		Identity:               uuid.NewV4().String(),
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepoIdentity,
		RepositoryIdentity:     repo.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	res = db.Create(share)
	if err = res.Error; err != nil {
		return nil, err
	}
	resp = new(types.FileShareCreateResponse)
	resp.Identity = share.Identity
	return
}
