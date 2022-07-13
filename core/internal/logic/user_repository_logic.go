package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/models"
	uuid "github.com/satori/go.uuid"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositoryLogic {
	return &UserRepositoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositoryLogic) UserRepository(req *types.UserRepoSaveRequeset, userIdentity string) (resp *types.UserRepoSaveResponse, err error) {
	up := &models.UserRepository{
		Identity:           uuid.NewV4().String(),
		UserIdentity:       userIdentity,
		ParentId:           int(req.ParentId),
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	res := l.svcCtx.DB.Create(up)
	if err = res.Error; err != nil {
		return nil, err
	}
	return
}
