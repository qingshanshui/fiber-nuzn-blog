package admin

import (
	"beego_blog_mvc/controllers"
)

type HomeController struct {
	controllers.Nuzn
}

// Get 管理首页
func (c *HomeController) Get() {
	c.Layout = "admin/layout/index.html"
	c.TplName = "admin/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutCss"] = "admin/css.html"
	c.LayoutSections["LayoutJs"] = "admin/js.html"
}
