package web

import (
	"fiber-nuzn-blog/controllers"
	serviceWeb "fiber-nuzn-blog/service/web"
	"fiber-nuzn-blog/validator"
	validatorWeb "fiber-nuzn-blog/validator/form/web"
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
	// 初始化参数结构体
	HomeRequestForm := validatorWeb.HomeRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &HomeRequestForm); err != nil {
		return err
	}
	// 分页调用
	t.PaginationInit(&HomeRequestForm.PaginationRequest)
	// 公共调用
	InitData := t.InitData()
	// 实际业务调用
	result := serviceWeb.NewHomeService().Home(HomeRequestForm.CurrPage, HomeRequestForm.PageSize)
	result.CurrPage = HomeRequestForm.CurrPage
	result.InitData = InitData
	// 渲染页面
	return c.Render("web/index", result, "web/layout/index")
}
