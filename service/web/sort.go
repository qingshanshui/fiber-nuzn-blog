package web

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
	"fiber-nuzn-blog/validator/form/web"
)

type Sort struct{}

func NewSortService() *Sort {
	return &Sort{}
}

// Sort 分类
func (t *Sort) Sort(currPage, pageSize int, id string) web.SortResponse {

	ma := models.NewArticle()
	// 首页文章列表
	ae := ma.GetSortToArticleList(currPage, pageSize, id)
	// 首页总条数
	tt := ma.GetSortToArticleListCount(id)

	return web.SortResponse{
		PaginationResponse: form.PaginationResponse{
			TotalCount: tt,
		},
		ArticleList: ae,
	}
}
