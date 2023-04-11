package controllers

import (
	"fiber-nuzn-blog/models"
	"github.com/gofiber/fiber/v2"
)

// 返回的 code 码值
const (
	ERROR   = 0
	SUCCESS = 1000
)

// Base 公共方法
type Base struct{}

// Ok 成功
func (r *Base) Ok(data interface{}) fiber.Map {
	return fiber.Map{
		"code": SUCCESS,
		"data": data,
		"msg":  "操作成功",
	}
}

// Fail 失败
func (r *Base) Fail(err error, code ...int) fiber.Map {
	return fiber.Map{
		"code": If(code),
		"data": err.Error(),
		"msg":  "操作失败",
	}
}

func If(code []int) int {
	if code != nil {
		return code[0]
	}
	return ERROR
}

// InitData 首页查询接口
func (r *Base) InitData() map[string]interface{} {
	// 导航栏
	mn := models.NewNavBar()
	results := mn.GetWebNavBarList()

	// 设置cookie
	//cookie := c.Ctx.Input.Cookie("token")

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
	return map[string]interface{}{
		"article":    article,
		"navBar":     navBar,
		"pages":      pages,
		"updateTime": articleList.UpdatedAt.Format("2006-01-02"),
		"linkAll":    linkAll,
		"Sort":       results,
	}
	//c.Data["Cookie"] = cookie
}
