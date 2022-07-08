package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/models"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	user := new(models.UserBasic)
	db := l.svcCtx.DB
	result := db.Where("identity = ?", req.Identity).First(user)
	if err = result.Error; err != nil {
		return nil, err
	}
	resp = new(types.UserDetailResponse)
	resp.Name = user.Name
	resp.Email = user.Email
	return
}
