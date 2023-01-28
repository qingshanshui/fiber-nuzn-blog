package web

import (
	"beego_blog_mvc/models"
)

type MainController struct {
	BaseController
}

// Get 首页查询接口
func (c *MainController) Get() {
	c.InitData()

	pageSize, _ := c.GetInt("pageSize")     // 一页多少条
	pageNumber, _ := c.GetInt("pageNumber") //当前页

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

	c.Data["Article"] = article
	c.Data["TotalNumber"] = totalNumber
	c.Data["PageNumber"] = pageNumber
	c.Data["Acrid"] = Acrid
	c.Layout = "web/layout/index.html"
	c.TplName = "web/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "web/css.html"
	c.LayoutSections["LayoutJs"] = "web/js.html"
}
