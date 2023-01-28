package SortController

import (
	"beego_blog_mvc/controllers/web"
	"beego_blog_mvc/models"
)

type SortController struct {
	web.BaseController
}

// Get 分类列表
func (c *SortController) Get() {
	c.InitData()
	// 当前分类id
	id := c.Ctx.Input.Param(":id")
	// 一页多少条
	pageSize, _ := c.GetInt("pageSize")
	//当前页
	pageNumber, _ := c.GetInt("pageNumber")

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
	c.Data["Article"] = article
	c.Data["TotalNumber"] = totalNumber
	c.Data["PageNumber"] = pageNumber
	c.Layout = "web/layout/index.html"
	c.TplName = "web/sort/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "web/sort/css.html"
	c.LayoutSections["LayoutJs"] = "web/sort/js.html"
}
