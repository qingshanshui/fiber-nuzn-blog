package web

import (
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"
)

type Sort struct{}

func NewSortService() *Sort {
	return &Sort{}
}

// Sort 分类
func (t *Sort) Sort(currPage, pageSize int, id string) fiber.Map {

	ma := models.NewArticle()
	// 首页文章列表
	ae := ma.GetSortToArticleList(currPage, pageSize, id)
	// 首页总条数
	tt := ma.GetSortToArticleListCount(id)

	return fiber.Map{
		"Ae": ae,
		"Tt": tt,
	}
}
