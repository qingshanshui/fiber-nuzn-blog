package LoginController

import (
	"fiber-nuzn-blog/controllers"
	serviceAdmin "fiber-nuzn-blog/service/admin"
	"fiber-nuzn-blog/validator"
	validatorForm "fiber-nuzn-blog/validator/form/admin"
	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	controllers.Base
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

// Get 登录页面
func (t *LoginController) Get(c *fiber.Ctx) error {
	return c.Render("admin/login/index", nil)
}

// Post 登录
func (t *LoginController) Post(c *fiber.Ctx) error {
	// 初始化参数结构体
	LoginRequestForm := validatorForm.LoginRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &LoginRequestForm); err != nil {
		return c.JSON(t.Fail(err))
	}
	// 实际业务调用
	token, err := serviceAdmin.NewLoginService().Login(LoginRequestForm)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(token))
}

// Logout 退出
func (t *LoginController) Logout(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.Redirect("/", 301)
}
