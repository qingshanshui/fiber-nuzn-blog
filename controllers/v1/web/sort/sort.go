package SortController

import (
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/service/web"
	"fiber-nuzn-blog/validator"
	web2 "fiber-nuzn-blog/validator/form/web"
	"github.com/gofiber/fiber/v2"
)

type SortController struct {
	controllers.Base
}

func NewSortController() *SortController {
	return &SortController{}
}

// Sort 分类页
func (t *SortController) Sort(c *fiber.Ctx) error {
	// 当前分类id
	id := c.Params("id")
	// 初始化参数结构体
	SortRequestForm := web2.SortRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &SortRequestForm); err != nil {
		return err
	}
	// 分页调用
	t.PaginationInit(&SortRequestForm.PaginationRequest)
	// 公共调用
	InitData := t.InitData()
	// 实际业务调用
	result := web.NewSortService().Sort(SortRequestForm.CurrPage, SortRequestForm.PageSize, id)

	return c.Render("web/sort/index", fiber.Map{
		"ArticleList": result["Ae"],             // 文章列表
		"TotalCount":  result["Tt"],             // 首页总条数
		"CurrPage":    SortRequestForm.CurrPage, // 当前页码
		"InitData":    InitData,
	}, "web/layout/index")
}
