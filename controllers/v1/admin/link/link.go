package LinkController

import (
	"errors"
	"fiber-nuzn-blog/controllers"
	serviceAdmin "fiber-nuzn-blog/service/admin"
	"fiber-nuzn-blog/validator"
	validatorForm "fiber-nuzn-blog/validator/form/admin"
	"github.com/gofiber/fiber/v2"
)

type LinkController struct {
	controllers.Base
}

func NewLinkController() *LinkController {
	return &LinkController{}
}

// Home 获取友链
func (t *LinkController) Home(c *fiber.Ctx) error {
	r := serviceAdmin.NewLinkService().Home()
	return c.Render("admin/link/index", fiber.Map{
		"linkAll": r,
	}, "admin/layout/index")
}

// AddView 渲染添加友链
func (t *LinkController) AddView(c *fiber.Ctx) error {
	return c.Render("admin/link/create", fiber.Map{}, "admin/layout/index")
}

// Add 添加友链的post
func (t *LinkController) Add(c *fiber.Ctx) error {
	// 初始化参数结构体
	LinkCreateRequestForm := validatorForm.LinkCreateRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &LinkCreateRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	err := serviceAdmin.NewLinkService().Add(LinkCreateRequestForm)
	if err != nil {
		return c.JSON(t.Fail(errors.New("创建友链失败")))
	} else {
		return c.JSON(t.Ok(errors.New("创建友链成功").Error()))
	}
}

// EditView 渲染 编辑友链
func (t *LinkController) EditView(c *fiber.Ctx) error {
	// 接收参数
	id := c.FormValue("id")
	// 业务处理
	r := serviceAdmin.NewLinkService().EditView(id)
	return c.Render("admin/link/edit", fiber.Map{
		"Result": r,
	}, "admin/layout/index")
}

// Edit 编辑友链的post
func (t *LinkController) Edit(c *fiber.Ctx) error {
	// 初始化参数结构体
	LinkEditRequestForm := validatorForm.LinkEditRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &LinkEditRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	err := serviceAdmin.NewLinkService().Edit(LinkEditRequestForm)
	if err != nil {
		return c.JSON(t.Fail(errors.New("编辑友链失败")))
	} else {
		return c.JSON(t.Ok(errors.New("编辑友链成功").Error()))
	}
}

// Del 删除友链
func (t *LinkController) Del(c *fiber.Ctx) error {
	// 接收参数
	id := c.FormValue("id")
	// 实际业务调用
	err := serviceAdmin.NewLinkService().Del(id)
	if err != nil {
		return c.JSON(t.Fail(errors.New("删除友链失败")))
	} else {
		return c.JSON(t.Ok(errors.New("删除友链成功").Error()))
	}
}
