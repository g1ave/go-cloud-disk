package models

import "github.com/jinzhu/gorm"

type ShareBasic struct {
	gorm.Model
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string
	RepositoryIdentity     string
	ExpiredTime            int
	ClickNum               int
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
