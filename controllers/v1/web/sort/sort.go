package SortController

import (
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"
)

type SortController struct {
	controllers.Base
}

func NewSortController() *SortController {
	return &SortController{}
}

// Home 首页
func (t *SortController) Home(c *fiber.Ctx) error {
	InitData := t.InitData()
	// 当前分类id
	id := c.Params("id")
	// 一页多少条
	pageSize, _ := c.ParamsInt("pageSize")
	//当前页
	pageNumber, _ := c.ParamsInt("pageNumber")

	if pageSize == 0 {
		pageSize = 10
	}
	if pageNumber == 0 {
		pageNumber = 1
	}
	ma := models.NewArticle()
	// 总条数
	totalNumber := ma.GetSortToArticleListCount(id)
	// 文章列表
	article := ma.GetSortToArticleList(pageNumber, pageSize, id)

	return c.Render("web/sort/index", fiber.Map{
		"Article":     article,
		"TotalNumber": totalNumber,
		"PageNumber":  pageNumber,
		"article":     InitData["article"],
		"navBar":      InitData["navBar"],
		"pages":       InitData["pages"],
		"updateTime":  InitData["updateTime"],
		"linkAll":     InitData["linkAll"],
		"Sort":        InitData["Sort"],
	}, "web/layout/index")
}
