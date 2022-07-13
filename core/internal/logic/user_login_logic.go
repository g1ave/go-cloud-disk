package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/g1ave/go-cloud-disk/core/utils"
	"log"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user := &models.UserBasic{}
	db := l.svcCtx.DB
	result := db.Where("name = ? and password = ?", req.Name, utils.Md5(req.Password)).First(user)
	if err2 := result.Error; err2 != nil {
		log.Println(err2)
		return nil, err2
	}
	token, err := utils.GenerateNewToke(user.ID, user.Name, user.Identity, 3000)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token
	return
}
