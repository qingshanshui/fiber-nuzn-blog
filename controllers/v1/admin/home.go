package admin

import (
	"fiber-nuzn-blog/controllers"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	controllers.Base
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

// Home 首页
func (t *HomeController) Home(c *fiber.Ctx) error {
	return c.Render("admin/index", nil, "admin/layout/index")
}
