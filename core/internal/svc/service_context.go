package svc

import (
	"github.com/g1ave/go-cloud-disk/core/internal/config"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/go-redis/redis/v9"
	"github.com/jinzhu/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     models.Init(c.Database.DSN),
		Redis:  models.InitRedis(c.Database.Redis),
	}
}
