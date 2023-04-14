package models

import (
	"fiber-nuzn-blog/initalize"
	"gorm.io/gorm"
)

// Link 友链
type Link struct {
	*gorm.Model
	Uid      string `gorm:"not null;comment:友链uid，唯一id" json:"uid"` // uid
	Title    string `gorm:"comment:友链名称" json:"title"`              // 名称
	Url      string `gorm:"comment:友链url" json:"url"`               // url
	Describe string `gorm:"comment:友链描述" json:"describe"`           // 描述
	Show     int    `gorm:"comment:显示隐藏，1隐藏 2显示" json:"show"`       // 显示隐藏，1隐藏 2显示
	Sort     int    `gorm:"comment:导航条排序" json:"sort"`              // 排序
}

// TableName 重命名表
func (u *Link) TableName() string {
	return "blog_Link"
}

// NewLink new一个空的结构体
func NewLink() *Link {
	return &Link{}
}

// Create 创建
func (u *Link) Create() error {
	return initalize.DB.Create(u).Error
}

// Delete 删除
func (u *Link) Delete() error {
	return initalize.DB.Model(u).Delete(u).Error
}

// Update 修改
func (u *Link) Update(uid string) error {
	return initalize.DB.Model(u).Where("uid = ? ", uid).Updates(u).Error
}

// GetLinkByUid 通过uid查询 友链 详情
func (u *Link) GetLinkByUid(uid string) *Link {
	if err := initalize.DB.Where("uid = ?", uid).First(u).Error; err != nil {
		return nil
	}
	return u
}

// GetLinkList 获取 首页友链 列表
func (u *Link) GetLinkList() []Link {
	var uArr []Link
	if err := initalize.DB.Where("`show` = ?", 2).Order("sort desc").Limit(10).Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}

// GetAdminLinkList 获取 admin友链 列表
func (u *Link) GetAdminLinkList() []Link {
	var uArr []Link
	if err := initalize.DB.Order("sort desc").Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}
