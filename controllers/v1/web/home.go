package web

import (
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	controllers.Base
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

// Home 首页
func (t *HomeController) Home(c *fiber.Ctx) error {
	InitData := t.InitData()
	pageSize, _ := c.ParamsInt("pageSize")     // 一页多少条
	pageNumber, _ := c.ParamsInt("pageNumber") //当前页

	if pageSize == 0 {
		pageSize = 10
	}
	if pageNumber == 0 {
		pageNumber = 1
	}
	ma := models.NewArticle()
	// 总条数
	totalNumber := ma.GetArticleListCount()
	// 首页列表
	article := ma.GetArticleList(pageNumber, pageSize)
	// 热门推荐
	Acrid := ma.GetAcridList()
	return c.Render("web/index", fiber.Map{
		"Acrid":       Acrid,
		"PageNumber":  pageNumber,
		"TotalNumber": totalNumber,
		"Article":     article,
		"article":     InitData["article"],
		"navBar":      InitData["navBar"],
		"pages":       InitData["pages"],
		"updateTime":  InitData["updateTime"],
		"linkAll":     InitData["linkAll"],
		"Sort":        InitData["Sort"],
	}, "web/layout/index")
}
