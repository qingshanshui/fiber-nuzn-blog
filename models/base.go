package models

import (
	"fiber-nuzn-blog/initalize"
	"gorm.io/gorm"
)

type BasesModel struct {
	gorm.Model
}

func (model *BasesModel) DB() *gorm.DB {
	return DB()
}

func DB() *gorm.DB {
	return initalize.DB
}
