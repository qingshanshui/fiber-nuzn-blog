package NavbarController

import (
	"errors"
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/service/admin"
	"fiber-nuzn-blog/validator"
	admin2 "fiber-nuzn-blog/validator/form/admin"
	"github.com/gofiber/fiber/v2"
)

type NavbarController struct {
	controllers.Base
}

func NewNavbarController() *NavbarController {
	return &NavbarController{}
}

// Home 导航栏管理页面
func (t *NavbarController) Home(c *fiber.Ctx) error {
	// 实际业务调用
	result := admin.NewNavbarService().Home()
	// 渲染页面
	return c.Render("admin/navbar/index", fiber.Map{
		"navBarAll": result,
	}, "admin/layout/index")
}

// AddView 渲染添加分类
func (t *NavbarController) AddView(c *fiber.Ctx) error {
	return c.Render("admin/navbar/create", fiber.Map{}, "admin/layout/index")
}

// Add 添加分类的post
func (t *NavbarController) Add(c *fiber.Ctx) error {
	// 初始化参数结构体
	NavbarCreateResponseForm := admin2.NavbarRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &NavbarCreateResponseForm); err != nil {
		return err
	}
	// 实际业务调用
	err := admin.NewNavbarService().Add(NavbarCreateResponseForm)
	if err != nil {
		return c.JSON(t.Fail(errors.New("创建导航栏失败")))
	} else {
		return c.JSON(t.Ok(errors.New("创建导航栏成功").Error()))
	}
}

// EditView 渲染编辑分类
func (t *NavbarController) EditView(c *fiber.Ctx) error {
	// 接收参数
	id := c.FormValue("id")
	// 业务处理
	r := admin.NewNavbarService().EditView(id)
	return c.Render("admin/navbar/edit", fiber.Map{
		"Result": r,
	}, "admin/layout/index")
}

// Edit 编辑分类的post
func (t *NavbarController) Edit(c *fiber.Ctx) error {
	// 初始化参数结构体
	NavbarCreateResponseForm := admin2.NavbarEditRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &NavbarCreateResponseForm); err != nil {
		return err
	}
	// 实际业务调用
	err := admin.NewNavbarService().Edit(NavbarCreateResponseForm)
	if err != nil {
		return c.JSON(t.Fail(errors.New("编辑导航栏失败")))
	} else {
		return c.JSON(t.Ok(errors.New("编辑导航栏成功").Error()))
	}
}

// Del 删除分类
func (t *NavbarController) Del(c *fiber.Ctx) error {
	id := c.FormValue("id")
	// 实际业务调用
	err := admin.NewNavbarService().Del(id)
	if err != nil {
		return c.JSON(t.Fail(errors.New("删除导航栏失败")))
	} else {
		return c.JSON(t.Ok(errors.New("删除导航栏成功").Error()))
	}
}
