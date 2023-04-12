package web

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
)

// CategoryRequest 详情页入口参数
type CategoryRequest struct{}

// CategoryResponse 详情页出口参数
type CategoryResponse struct {
	ArticleCategory     models.Article // 文章详情
	ArticleCategoryTime string         // 文章时间
	InitData            form.InitData  // 初始化参数
}
