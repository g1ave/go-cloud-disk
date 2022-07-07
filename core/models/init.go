package models

import (
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
