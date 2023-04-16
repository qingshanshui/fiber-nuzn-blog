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
	result.InitData = InitData
	// 渲染页面
	return c.Render("web/category/index", result, "web/layout/index")
}
