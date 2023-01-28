package CategoryController

import (
	"beego_blog_mvc/controllers/web"
	"beego_blog_mvc/models"
)

type CategoryController struct {
	web.BaseController
}

// Category 文章详情
func (c *CategoryController) Category() {
	c.InitData()
	uid := c.Ctx.Input.Param(":id")
	ma := models.NewArticle()
	// 通过uid 获取详情
	Category := ma.GetArticleByUid(uid)
	if Category == nil {
		c.Layout = "web/layout/index.html"
		c.TplName = "web/category/404.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["LayoutCss"] = "web/category/css.html"
		c.LayoutSections["LayoutJs"] = "web/category/js.html"
		return
	}
	ma.AddHot()
	c.Data["Article"] = Category
	c.Data["ArticleTime"] = Category.UpdatedAt.Format("2006-01-02")
	c.Layout = "web/layout/index.html"
	c.TplName = "web/category/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "web/category/css.html"
	c.LayoutSections["LayoutJs"] = "web/category/js.html"
}
