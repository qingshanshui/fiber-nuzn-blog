package form

import (
	"fiber-nuzn-blog/models"
)

// PaginationRequest 分页入参
type PaginationRequest struct {
	PageSize int `form:"pageSize" json:"pageSize"` //  一页/条数
	CurrPage int `form:"currPage" json:"currPage"` //  当前页码
}

// PaginationResponse 分页出参
type PaginationResponse struct {
	TotalCount int64 `form:"totalCount" json:"totalCount"` //  总条数
	CurrPage   int   `form:"currPage" json:"currPage"`     //  当前页码
}

// InitData 初始参数
type InitData struct {
	ArticleTotalCount int64           // 文章总条数
	NavBarCount       int64           // 分类总条数
	PagesCount        int64           // 总页面数量
	UpdateTime        string          //最后文章更新时间
	LinkAllList       []models.Link   // 友情链接列表
	WebNavBarList     []models.NavBar // 导航栏列表
}
