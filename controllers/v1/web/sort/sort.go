package SortController

import (
	"fiber-nuzn-blog/controllers"
	serviceWeb "fiber-nuzn-blog/service/web"
	"fiber-nuzn-blog/validator"
	validatorWeb "fiber-nuzn-blog/validator/form/web"
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
	SortRequestForm := validatorWeb.SortRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &SortRequestForm); err != nil {
		return err
	}
	// 分页调用
	t.PaginationInit(&SortRequestForm.PaginationRequest)
	// 公共调用
	InitData := t.InitData()
	// 实际业务调用
	result := serviceWeb.NewSortService().Sort(SortRequestForm.CurrPage, SortRequestForm.PageSize, id)
	result.CurrPage = SortRequestForm.CurrPage
	result.InitData = InitData
	// 渲染页面
	return c.Render("web/sort/index", result, "web/layout/index")
}
