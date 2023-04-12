package web

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
)

// SortRequest 分类页入口参数
type SortRequest struct {
	form.PaginationRequest
}

// SortResponse 分类页出口参数
type SortResponse struct {
	form.PaginationResponse
	ArticleList []models.Article // 文章列表
	InitData    form.InitData    // 初始化参数
}
