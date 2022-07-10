package svc

import (
	"github.com/g1ave/go-cloud-disk/core/internal/config"
	"github.com/g1ave/go-cloud-disk/core/internal/middleware"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/go-redis/redis/v9"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     models.Init(c.Database.DSN),
		Redis:  models.InitRedis(c.Database.Redis),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
