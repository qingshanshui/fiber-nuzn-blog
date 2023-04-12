package web

import (
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"
)

type Home struct{}

func NewHomeService() *Home {
	return &Home{}
}

// Home 首页
func (t *Home) Home(currPage, pageSize int) fiber.Map {

	ma := models.NewArticle()
	// 首页文章列表
	ae := ma.GetArticleList(currPage, pageSize)
	// 热门推荐列表
	ht := ma.GetHotList()
	// 首页总条数
	tt := ma.GetArticleCount()

	return fiber.Map{
		"Ae": ae,
		"Ht": ht,
		"Tt": tt,
	}
}
