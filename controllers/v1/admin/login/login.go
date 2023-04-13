package LoginController

import (
	"errors"
	"fiber-nuzn-blog/controllers"
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
	// 接收参数
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println(username, password, "===========")

	// 业务处理
	if username == "" || password == "" {
		return c.JSON(t.Fail(errors.New("用户名密码不能为空")))
	}
	mu := models.NewUser()
	if err := mu.Login(username, password); err != nil {
		return c.JSON(t.Fail(err))
	}
	if mu.Uid != "" {
		token, _ := utils.CreateToken(mu.Uid, viper.GetString("Jwt.Secret"))
		return c.JSON(map[string]interface{}{
			"code":  true,
			"token": token,
			"msg":   "登录成功",
		})
	} else {
		return c.JSON(t.Fail(errors.New("用户名密码错误")))
	}
}

// Logout 退出
func (t *LoginController) Logout(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.Redirect("/", 301)
}
