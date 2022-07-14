package logic

import (
	"context"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/models"
	uuid "github.com/satori/go.uuid"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileShareSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileShareSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileShareSaveLogic {
	return &FileShareSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileShareSaveLogic) FileShareSave(req *types.FileShareSaveRequest, userIdentity string) (resp *types.FileShareSaveResponse, err error) {
	repo := new(models.RepositoryPool)
	db := l.svcCtx.DB

	// determine if the repository exists
	res := db.First(repo, "identity = ?", req.RepoIdentity)
	if err = res.Error; err != nil {
		return nil, err
	}

	// determine if the folder user want to save exists
	res = db.First(new(models.UserRepository), "id = ? AND user_identity = ?", req.ParentId, userIdentity)
	if err = res.Error; err != nil {
		return nil, define.FolderNotExistsErr
	}

	// determine if the name user want to save exists
	res = db.First(&models.UserRepository{}, "user_identity = ? AND parent_id = ? AND name = ?", userIdentity, req.ParentId, req.Name)
	if err = res.Error; err == nil {
		return nil, define.NameExistedErr
	}
	// save repository
	userRepo := &models.UserRepository{
		Identity:           uuid.NewV4().String(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: repo.Identity,
		Ext:                repo.Ext,
		Name:               req.Name,
	}
	res = db.Create(userRepo)
	if err = res.Error; err != nil {
		return nil, err
	}
	resp = new(types.FileShareSaveResponse)
	resp.Identity = userRepo.Identity
	return
}
