package CategoryController

import (
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/service/web"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	controllers.Base
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// Category 详情页
func (t *CategoryController) Category(c *fiber.Ctx) error {
	uid := c.Params("id")
	// 公共调用
	InitData := t.InitData()
	// 实际业务调用
	result := web.NewCategoryService().Category(uid)

	if result == nil {
		return c.Render("web/category/404", fiber.Map{}, "web/layout/index")
	}

	return c.Render("web/category/index", fiber.Map{
		"ArticleCategory":     result["Cy"],
		"ArticleCategoryTime": result["Cm"],
		"InitData":            InitData,
	}, "web/layout/index")
}
