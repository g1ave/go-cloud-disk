package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/g1ave/go-cloud-disk/core/utils"
	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	name, password, email, code := req.Name, req.Password, req.Mail, req.Code
	db, rdb := l.svcCtx.DB, l.svcCtx.Redis
	var cnt int64
	result := db.Model(&models.UserBasic{}).Where("name = ?", name).Count(&cnt)
	if err = result.Error; err != nil {
		return nil, err
	}
	actualCode, err := rdb.Get(l.ctx, email).Result()
	if err == redis.Nil {
		return nil, define.MailMismatchErr
	} else if err != nil {
		return nil, err
	} else if actualCode != code {
		return nil, define.CodeMismatchErr
	}
	newUser := models.UserBasic{Name: name, Password: utils.Md5(password), Email: email, Identity: uuid.NewV4().String()}
	res := db.Create(&newUser)
	if err = res.Error; err != nil {
		return nil, err
	}
	return
}
