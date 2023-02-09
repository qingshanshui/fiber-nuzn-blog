package models

// NavBar 导航条表
type NavBar struct {
	BasesModel
	Uid   string `gorm:"not null;comment:文章uid，唯一id" json:"uid"` // uid
	Title string `gorm:"comment:导航栏名称" json:"title"`             // 名称
	Url   string `gorm:"comment:导航栏url" json:"url"`              // url
	Show  int    `gorm:"comment:显示隐藏，1隐藏 2显示" json:"show"`       // 显示隐藏，1隐藏 2显示
	Sort  int    `gorm:"comment:导航条排序" json:"sort"`              // 排序
}

// TableName 重命名表
func (u *NavBar) TableName() string {
	return "blog_navBar"
}

// NewNavBar new一个空的结构体
func NewNavBar() *NavBar {
	return &NavBar{}
}

// Create 创建
func (u *NavBar) Create() error {
	return u.DB().Create(u).Error
}

// Delete 删除
func (u *NavBar) Delete() error {
	return u.DB().Model(u).Delete(u).Error
}

// Update 修改
func (u *NavBar) Update(uid string) error {
	return u.DB().Model(u).Where("uid = ? ", uid).Updates(u).Error
}

// GetLinkByUid 通过uid查询 分类 详情
func (u *NavBar) GetLinkByUid(uid string) *NavBar {
	if err := u.DB().Where("uid = ?", uid).First(u).Error; err != nil {
		return nil
	}
	return u
}

// GetWebNavBarList 获取 web分类 列表
func (u *NavBar) GetWebNavBarList() []NavBar {
	var uArr []NavBar
	if err := u.DB().Where("`show` = ?", 2).Order("sort").Limit(100).Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}

// GetWebNavBarListAll 获取 admin分类 列表
func (u *NavBar) GetWebNavBarListAll() []NavBar {
	var uArr []NavBar
	if err := u.DB().Order("sort").Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}

// GetNavBarListCount 获取 分类列表 总数
func (u *NavBar) GetNavBarListCount() int64 {
	var count int64
	if err := u.DB().Model(u).Where("`show` = ?", 2).Count(&count).Error; err != nil {
		return 0
	}
	return count
}
