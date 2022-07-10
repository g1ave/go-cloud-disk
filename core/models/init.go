package models

import (
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func Init(dataSource string) *gorm.DB {
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		log.Println("Gorm DB Connect Error", err)
		return nil
	}
	db.AutoMigrate(&UserBasic{}, &RepositoryPool{}, &UserRepository{}, &ShareBasic{})
	return db
}

func InitRedis(dataSource string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     dataSource,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
