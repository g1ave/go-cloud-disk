package models

import "github.com/jinzhu/gorm"

type UserBasic struct {
	gorm.Model
	Identity string
	Name     string
	Password string
	Email    string
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
