package models

import (
	"beego_blog_mvc/global"
	"gorm.io/gorm"
)

type BasesModel struct {
	gorm.Model
}

func (model *BasesModel) DB() *gorm.DB {
	return DB()
}

func DB() *gorm.DB {
	return global.DB
}
