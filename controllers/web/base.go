package web

import (
	"beego_blog_mvc/models"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

// InitData 首页查询接口
func (c *BaseController) InitData() {
	// 导航栏
	mn := models.NewNavBar()
	results := mn.GetWebNavBarList()

	// 设置cookie
	cookie := c.Ctx.Input.Cookie("token")

	/*        站点统计      */
	// 文章数（article表的数量）
	ma := models.NewArticle()
	article := ma.GetArticleListCount()

	// 分类数（navBar表）
	navBar := mn.GetNavBarListCount()

	// 页面数（article表+navBar表）
	pages := article + navBar

	// 更新 （文章表最后一条内容）
	articleList := ma.GetArticleLast()

	// 获取友链（友情链接）
	ml := models.NewLink()
	linkAll := ml.GetLinkList()

	c.Data["Total"] = map[string]interface{}{
		"article":    article,
		"navBar":     navBar,
		"pages":      pages,
		"updateTime": articleList.UpdatedAt.Format("2006-01-02"),
	}
	c.Data["linkAll"] = linkAll
	c.Data["Sort"] = results
	c.Data["Cookie"] = cookie
}
