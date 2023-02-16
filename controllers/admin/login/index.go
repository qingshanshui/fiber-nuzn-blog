package LoginController

import (
	"beego_blog_mvc/controllers"
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	controllers.Nuzn
}

// Get 登录页面
func (c *LoginController) Get() {
	c.TplName = "admin/login/index.html"
}

// Post 登录
func (c *LoginController) Post() {
	// 接收参数
	username := c.GetString("username")
	password := c.GetString("password")

	// 业务处理
	if username == "" || password == "" {
		c.Fail("用户名密码不能为空")
		return
	}
	mu := models.NewUser()
	if err := mu.Login(username, password); err != nil {
		c.Fail(err)
		return
	}
	if mu.Uid != "" {
		Secret, _ := beego.AppConfig.String("Secret")
		token, _ := utils.CreateToken(*mu, Secret)
		c.Data["json"] = map[string]interface{}{
			"code":  true,
			"token": token,
			"msg":   "登录成功",
		}
		c.ServeJSON()
	} else {
		c.Fail("用户名密码错误")
		return
	}
}

func (c *LoginController) Logout() {
	c.Ctx.SetCookie("token", "", -1)
	c.Redirect("/", 301)
}
