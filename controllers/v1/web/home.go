package web

import (
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/service/web"
	"fiber-nuzn-blog/validator"
	web2 "fiber-nuzn-blog/validator/form/web"
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
	HomeRequestForm := web2.HomeRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &HomeRequestForm); err != nil {
		return err
	}
	// 分页调用
	t.PaginationInit(&HomeRequestForm.PaginationRequest)
	// 公共调用
	InitData := t.InitData()
	// 实际业务调用
	result := web.NewHomeService().Home(HomeRequestForm.CurrPage, HomeRequestForm.PageSize)
	return c.Render("web/index", fiber.Map{
		"HotArticleList": result["Ht"],             // 热门推荐
		"ArticleList":    result["Ae"],             // 文章列表
		"TotalCount":     result["Tt"],             // 首页总条数
		"CurrPage":       HomeRequestForm.CurrPage, // 当前页码
		"InitData":       InitData,
	}, "web/layout/index")
}
