package models

import "github.com/jinzhu/gorm"

type UserRepository struct {
	gorm.Model
	Identity           string
	UserIdentity       string
	ParentId           int
	RepositoryIdentity string
	Ext                string
	Name               string
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
