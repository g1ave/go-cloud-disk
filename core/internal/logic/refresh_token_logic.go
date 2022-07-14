package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/utils"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenRequest, authorization string) (resp *types.RefreshTokenResponse, err error) {
	uc, err := utils.ParseToken(authorization)
	if err != nil {
		return nil, err
	}
	// generate new token
	token, err := utils.GenerateNewToke(uc.Id, uc.Name, uc.Identity, define.TokenExpiredTime)
	if err != nil {
		return nil, err
	}
	refreshToken, err := utils.GenerateNewToke(uc.Id, uc.Name, uc.Identity, define.RefreshTokenExpiredTime)
	if err != nil {
		return nil, err
	}
	resp = new(types.RefreshTokenResponse)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
