package models

import "github.com/jinzhu/gorm"

type RepositoryPool struct {
	gorm.Model
	Identity string
	Hash     string
	Name     string
	Ext      string
	Size     int64
	Path     string
}
