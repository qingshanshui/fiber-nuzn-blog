package web

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
)

// HomeRequest 主页入口参数
type HomeRequest struct {
	form.PaginationRequest
}

// HomeResponse 主页入口参数
type HomeResponse struct {
	form.PaginationResponse
	HotArticleList []models.Article // 热门推荐
	ArticleList    []models.Article // 文章列表
	InitData       form.InitData    // 初始化参数
}
