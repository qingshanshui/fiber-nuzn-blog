package admin

import (
	"fiber-nuzn-blog/controllers"
	"github.com/gofiber/fiber/v2"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// Home 首页
func (t *DefaultController) Home(c *fiber.Ctx) error {
	return c.Render("admin/index", fiber.Map{}, "admin/layout/index")
}
