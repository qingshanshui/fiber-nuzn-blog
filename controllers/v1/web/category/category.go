package CategoryController

import (
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	controllers.Base
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// Home 首页
func (t *CategoryController) Home(c *fiber.Ctx) error {

	InitData := t.InitData()
	uid := c.Params("id")
	fmt.Println(uid, "------")
	ma := models.NewArticle()
	// 通过uid 获取详情
	Category := ma.GetArticleByUid(uid)
	if Category == nil {
		return c.Render("web/category/404", fiber.Map{}, "web/layout/home")
	}
	ma.AddHot()
	fmt.Println(InitData["navBar"], "============")
	return c.Render("web/category/index", fiber.Map{
		"Article":     Category,
		"ArticleTime": Category.UpdatedAt.Format("2006-01-02"),
		"article":     InitData["article"],
		"navBar":      InitData["navBar"],
		"pages":       InitData["pages"],
		"updateTime":  InitData["updateTime"],
		"linkAll":     InitData["linkAll"],
		"Sort":        InitData["Sort"],
	}, "web/layout/index")
}
