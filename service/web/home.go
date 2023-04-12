package web

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
	"fiber-nuzn-blog/validator/form/web"
)

type Home struct{}

func NewHomeService() *Home {
	return &Home{}
}

// Home 首页
func (t *Home) Home(currPage, pageSize int) web.HomeResponse {

	ma := models.NewArticle()
	// 首页文章列表
	ae := ma.GetArticleList(currPage, pageSize)
	// 热门推荐列表
	ht := ma.GetHotList()
	// 首页总条数
	tt := ma.GetArticleCount()

	return web.HomeResponse{
		PaginationResponse: form.PaginationResponse{
			TotalCount: tt,
		},
		HotArticleList: ht,
		ArticleList:    ae,
	}
}
