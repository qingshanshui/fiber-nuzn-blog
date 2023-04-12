package models

// Article 文章表
type Article struct {
	BasesModel
	Uid         string `gorm:"not null;comment:文章uid，唯一id" json:"uid"`            // id
	Sort        int    `gorm:"comment:文章排序" json:"sort"`                          // 排序
	Title       string `gorm:"not null;comment:文章标题" json:"title"`                // 标题
	Content     string `gorm:"comment:内容;type:longtext" json:"content"`           // 内容
	ContentMk   string `gorm:"comment:mk内容;type:longtext" json:"content_mk"`      // mk内容
	ContentHtml string `gorm:"comment:html内容;type:longtext" json:"content_html"`  // html内容
	Pic         string `gorm:"comment:文章主图" json:"pic"`                           // 主图
	Show        int    `gorm:"default:2;comment:文章是否显示，1为否，2为是，默认为2" json:"show"` // 是否显示，1为否，2为是
	NavBarId    string `gorm:"comment:分类id" json:"nav_bar_id"`                    // 导航栏uid
	UserId      string `gorm:"comment:用户id" json:"user_id"`                       // 用户uid
	UserName    string `gorm:"comment:用户昵称" json:"user_name"`                     // 用户昵称
	Hot         int64  `gorm:"default:0;comment:文章点击数" json:"hot"`                // 文章点击数
}

// TableName 自定义表名
func (u *Article) TableName() string {
	return "blog_article"
}

// NewArticle new一个空的结构体
func NewArticle() *Article {
	return &Article{}
}

// Create 创建
func (u *Article) Create() error {
	return u.DB().Create(u).Error
}

// Delete 删除
func (u *Article) Delete() error {
	return u.DB().Model(u).Delete(u).Error
}

// Update 修改
func (u *Article) Update(uid string) error {
	return u.DB().Model(u).Where("uid = ? ", uid).Updates(u).Error
}

// AddHot 文章点击数增加
func (u *Article) AddHot() {
	u.DB().Model(u).Update("hot", u.Hot+1)
}

// GetArticleByUid 通过uid查询文章详情
func (u *Article) GetArticleByUid(uid string) []Article {
	var a []Article
	if err := u.DB().Where("uid = ?", uid).Find(&a).Error; err != nil {
		return nil
	}
	return a
}

// GetArticleListAll 获取 全部文章 列表
func (u *Article) GetArticleListAll(pageNumber, pageSize int) []Article {
	var uArr []Article
	if err := u.DB().Offset((pageNumber - 1) * pageSize).Limit(pageSize).Order("created_at desc").Limit(10).Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}

// GetArticleList 获取 首页文章 列表
func (u *Article) GetArticleList(pageNumber, pageSize int) []Article {
	var uArr []Article
	if err := u.DB().Where("`show` = ?", 2).Offset((pageNumber - 1) * pageSize).Limit(pageSize).Order("created_at desc").Limit(10).Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}

// GetSortToArticleList 获取 分类页文章 列表
func (u *Article) GetSortToArticleList(pageNumber, pageSize int, id string) []Article {
	var uArr []Article
	if err := u.DB().Where("nav_bar_id = ? and `show` = ?", id, 2).Offset((pageNumber - 1) * pageSize).Limit(pageSize).Order("created_at desc").Limit(10).Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}

// GetHotList 获取 热门推荐 文章列表
func (u *Article) GetHotList() []Article {
	var uArr []Article
	if err := u.DB().Order("RAND()").Limit(5).Find(&uArr).Error; err != nil {
		return nil
	}
	return uArr
}

// GetArticleCount 获取 文章列表 总条数
func (u *Article) GetArticleCount() int64 {
	var count int64
	if err := u.DB().Model(u).Where("`show` = ?", 2).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// GetSortToArticleListCount 获取 分类文章列表 总数
func (u *Article) GetSortToArticleListCount(id string) int64 {
	var count int64
	if err := u.DB().Model(u).Where("nav_bar_id = ? and `show` = ?", id, 2).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// GetArticleLast 获取最后一篇文章
func (u *Article) GetArticleLast() *Article {
	if err := u.DB().Last(u).Error; err != nil {
		return nil
	}
	return u
}
