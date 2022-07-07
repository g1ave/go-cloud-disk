package models

import "github.com/jinzhu/gorm"

type UserRepository struct {
	gorm.Model
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
}
